package metrics

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	. "github.com/mlabouardy/mon-put-instance-data/services"
	"github.com/shirou/gopsutil/mem"
)

// Swap metric entity
type Swap struct{}

// Collect Swap usage
func (d Swap) Collect(instanceID string, c CloudWatchService) {
	swapMetrics, err := mem.SwapMemory()
	if err != nil {
		log.Fatal(err)
	}

	dimensionKey := "InstanceId"
	dimensions := []cloudwatch.Dimension{
		cloudwatch.Dimension{
			Name:  &dimensionKey,
			Value: &instanceID,
		},
	}

	swapUtilizationData := constructMetricDatum("SwapUtilization", swapMetrics.UsedPercent, cloudwatch.StandardUnitPercent, dimensions)
	c.Publish(swapUtilizationData, "CustomMetrics")

	swapUsedData := constructMetricDatum("SwapUsed", float64(swapMetrics.Used), cloudwatch.StandardUnitBytes, dimensions)
	c.Publish(swapUsedData, "CustomMetrics")

	swapFreeData := constructMetricDatum("SwapFree", float64(swapMetrics.Free), cloudwatch.StandardUnitBytes, dimensions)
	c.Publish(swapFreeData, "CustomMetrics")

	log.Printf("Swap - Utilization:%v%% Used:%v Free:%v\n", swapMetrics.UsedPercent, swapMetrics.Used, swapMetrics.Free)
}
