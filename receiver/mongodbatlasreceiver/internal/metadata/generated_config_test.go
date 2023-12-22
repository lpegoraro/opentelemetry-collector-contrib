// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestMetricsBuilderConfig(t *testing.T) {
	tests := []struct {
		name string
		want MetricsBuilderConfig
	}{
		{
			name: "default",
			want: DefaultMetricsBuilderConfig(),
		},
		{
			name: "all_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					MongodbatlasDbCounts:                                  MetricConfig{Enabled: true},
					MongodbatlasDbSize:                                    MetricConfig{Enabled: true},
					MongodbatlasDiskPartitionIopsAverage:                  MetricConfig{Enabled: true},
					MongodbatlasDiskPartitionIopsMax:                      MetricConfig{Enabled: true},
					MongodbatlasDiskPartitionLatencyAverage:               MetricConfig{Enabled: true},
					MongodbatlasDiskPartitionLatencyMax:                   MetricConfig{Enabled: true},
					MongodbatlasDiskPartitionSpaceAverage:                 MetricConfig{Enabled: true},
					MongodbatlasDiskPartitionSpaceMax:                     MetricConfig{Enabled: true},
					MongodbatlasDiskPartitionUsageAverage:                 MetricConfig{Enabled: true},
					MongodbatlasDiskPartitionUsageMax:                     MetricConfig{Enabled: true},
					MongodbatlasDiskPartitionUtilizationAverage:           MetricConfig{Enabled: true},
					MongodbatlasDiskPartitionUtilizationMax:               MetricConfig{Enabled: true},
					MongodbatlasProcessAsserts:                            MetricConfig{Enabled: true},
					MongodbatlasProcessBackgroundFlush:                    MetricConfig{Enabled: true},
					MongodbatlasProcessCacheIo:                            MetricConfig{Enabled: true},
					MongodbatlasProcessCacheSize:                          MetricConfig{Enabled: true},
					MongodbatlasProcessConnections:                        MetricConfig{Enabled: true},
					MongodbatlasProcessCPUChildrenNormalizedUsageAverage:  MetricConfig{Enabled: true},
					MongodbatlasProcessCPUChildrenNormalizedUsageMax:      MetricConfig{Enabled: true},
					MongodbatlasProcessCPUChildrenUsageAverage:            MetricConfig{Enabled: true},
					MongodbatlasProcessCPUChildrenUsageMax:                MetricConfig{Enabled: true},
					MongodbatlasProcessCPUNormalizedUsageAverage:          MetricConfig{Enabled: true},
					MongodbatlasProcessCPUNormalizedUsageMax:              MetricConfig{Enabled: true},
					MongodbatlasProcessCPUUsageAverage:                    MetricConfig{Enabled: true},
					MongodbatlasProcessCPUUsageMax:                        MetricConfig{Enabled: true},
					MongodbatlasProcessCursors:                            MetricConfig{Enabled: true},
					MongodbatlasProcessDbDocumentRate:                     MetricConfig{Enabled: true},
					MongodbatlasProcessDbOperationsRate:                   MetricConfig{Enabled: true},
					MongodbatlasProcessDbOperationsTime:                   MetricConfig{Enabled: true},
					MongodbatlasProcessDbQueryExecutorScanned:             MetricConfig{Enabled: true},
					MongodbatlasProcessDbQueryTargetingScannedPerReturned: MetricConfig{Enabled: true},
					MongodbatlasProcessDbStorage:                          MetricConfig{Enabled: true},
					MongodbatlasProcessGlobalLock:                         MetricConfig{Enabled: true},
					MongodbatlasProcessIndexBtreeMissRatio:                MetricConfig{Enabled: true},
					MongodbatlasProcessIndexCounters:                      MetricConfig{Enabled: true},
					MongodbatlasProcessJournalingCommits:                  MetricConfig{Enabled: true},
					MongodbatlasProcessJournalingDataFiles:                MetricConfig{Enabled: true},
					MongodbatlasProcessJournalingWritten:                  MetricConfig{Enabled: true},
					MongodbatlasProcessMemoryUsage:                        MetricConfig{Enabled: true},
					MongodbatlasProcessNetworkIo:                          MetricConfig{Enabled: true},
					MongodbatlasProcessNetworkRequests:                    MetricConfig{Enabled: true},
					MongodbatlasProcessOplogRate:                          MetricConfig{Enabled: true},
					MongodbatlasProcessOplogTime:                          MetricConfig{Enabled: true},
					MongodbatlasProcessPageFaults:                         MetricConfig{Enabled: true},
					MongodbatlasProcessRestarts:                           MetricConfig{Enabled: true},
					MongodbatlasProcessTickets:                            MetricConfig{Enabled: true},
					MongodbatlasSystemCPUNormalizedUsageAverage:           MetricConfig{Enabled: true},
					MongodbatlasSystemCPUNormalizedUsageMax:               MetricConfig{Enabled: true},
					MongodbatlasSystemCPUUsageAverage:                     MetricConfig{Enabled: true},
					MongodbatlasSystemCPUUsageMax:                         MetricConfig{Enabled: true},
					MongodbatlasSystemFtsCPUNormalizedUsage:               MetricConfig{Enabled: true},
					MongodbatlasSystemFtsCPUUsage:                         MetricConfig{Enabled: true},
					MongodbatlasSystemFtsDiskUsed:                         MetricConfig{Enabled: true},
					MongodbatlasSystemFtsMemoryUsage:                      MetricConfig{Enabled: true},
					MongodbatlasSystemMemoryUsageAverage:                  MetricConfig{Enabled: true},
					MongodbatlasSystemMemoryUsageMax:                      MetricConfig{Enabled: true},
					MongodbatlasSystemNetworkIoAverage:                    MetricConfig{Enabled: true},
					MongodbatlasSystemNetworkIoMax:                        MetricConfig{Enabled: true},
					MongodbatlasSystemPagingIoAverage:                     MetricConfig{Enabled: true},
					MongodbatlasSystemPagingIoMax:                         MetricConfig{Enabled: true},
					MongodbatlasSystemPagingUsageAverage:                  MetricConfig{Enabled: true},
					MongodbatlasSystemPagingUsageMax:                      MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					MongodbAtlasClusterName:     ResourceAttributeConfig{Enabled: true},
					MongodbAtlasDbName:          ResourceAttributeConfig{Enabled: true},
					MongodbAtlasDiskPartition:   ResourceAttributeConfig{Enabled: true},
					MongodbAtlasHostName:        ResourceAttributeConfig{Enabled: true},
					MongodbAtlasOrgName:         ResourceAttributeConfig{Enabled: true},
					MongodbAtlasProcessID:       ResourceAttributeConfig{Enabled: true},
					MongodbAtlasProcessPort:     ResourceAttributeConfig{Enabled: true},
					MongodbAtlasProcessTypeName: ResourceAttributeConfig{Enabled: true},
					MongodbAtlasProjectID:       ResourceAttributeConfig{Enabled: true},
					MongodbAtlasProjectName:     ResourceAttributeConfig{Enabled: true},
					MongodbAtlasProviderName:    ResourceAttributeConfig{Enabled: true},
					MongodbAtlasRegionName:      ResourceAttributeConfig{Enabled: true},
					MongodbAtlasUserAlias:       ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					MongodbatlasDbCounts:                                  MetricConfig{Enabled: false},
					MongodbatlasDbSize:                                    MetricConfig{Enabled: false},
					MongodbatlasDiskPartitionIopsAverage:                  MetricConfig{Enabled: false},
					MongodbatlasDiskPartitionIopsMax:                      MetricConfig{Enabled: false},
					MongodbatlasDiskPartitionLatencyAverage:               MetricConfig{Enabled: false},
					MongodbatlasDiskPartitionLatencyMax:                   MetricConfig{Enabled: false},
					MongodbatlasDiskPartitionSpaceAverage:                 MetricConfig{Enabled: false},
					MongodbatlasDiskPartitionSpaceMax:                     MetricConfig{Enabled: false},
					MongodbatlasDiskPartitionUsageAverage:                 MetricConfig{Enabled: false},
					MongodbatlasDiskPartitionUsageMax:                     MetricConfig{Enabled: false},
					MongodbatlasDiskPartitionUtilizationAverage:           MetricConfig{Enabled: false},
					MongodbatlasDiskPartitionUtilizationMax:               MetricConfig{Enabled: false},
					MongodbatlasProcessAsserts:                            MetricConfig{Enabled: false},
					MongodbatlasProcessBackgroundFlush:                    MetricConfig{Enabled: false},
					MongodbatlasProcessCacheIo:                            MetricConfig{Enabled: false},
					MongodbatlasProcessCacheSize:                          MetricConfig{Enabled: false},
					MongodbatlasProcessConnections:                        MetricConfig{Enabled: false},
					MongodbatlasProcessCPUChildrenNormalizedUsageAverage:  MetricConfig{Enabled: false},
					MongodbatlasProcessCPUChildrenNormalizedUsageMax:      MetricConfig{Enabled: false},
					MongodbatlasProcessCPUChildrenUsageAverage:            MetricConfig{Enabled: false},
					MongodbatlasProcessCPUChildrenUsageMax:                MetricConfig{Enabled: false},
					MongodbatlasProcessCPUNormalizedUsageAverage:          MetricConfig{Enabled: false},
					MongodbatlasProcessCPUNormalizedUsageMax:              MetricConfig{Enabled: false},
					MongodbatlasProcessCPUUsageAverage:                    MetricConfig{Enabled: false},
					MongodbatlasProcessCPUUsageMax:                        MetricConfig{Enabled: false},
					MongodbatlasProcessCursors:                            MetricConfig{Enabled: false},
					MongodbatlasProcessDbDocumentRate:                     MetricConfig{Enabled: false},
					MongodbatlasProcessDbOperationsRate:                   MetricConfig{Enabled: false},
					MongodbatlasProcessDbOperationsTime:                   MetricConfig{Enabled: false},
					MongodbatlasProcessDbQueryExecutorScanned:             MetricConfig{Enabled: false},
					MongodbatlasProcessDbQueryTargetingScannedPerReturned: MetricConfig{Enabled: false},
					MongodbatlasProcessDbStorage:                          MetricConfig{Enabled: false},
					MongodbatlasProcessGlobalLock:                         MetricConfig{Enabled: false},
					MongodbatlasProcessIndexBtreeMissRatio:                MetricConfig{Enabled: false},
					MongodbatlasProcessIndexCounters:                      MetricConfig{Enabled: false},
					MongodbatlasProcessJournalingCommits:                  MetricConfig{Enabled: false},
					MongodbatlasProcessJournalingDataFiles:                MetricConfig{Enabled: false},
					MongodbatlasProcessJournalingWritten:                  MetricConfig{Enabled: false},
					MongodbatlasProcessMemoryUsage:                        MetricConfig{Enabled: false},
					MongodbatlasProcessNetworkIo:                          MetricConfig{Enabled: false},
					MongodbatlasProcessNetworkRequests:                    MetricConfig{Enabled: false},
					MongodbatlasProcessOplogRate:                          MetricConfig{Enabled: false},
					MongodbatlasProcessOplogTime:                          MetricConfig{Enabled: false},
					MongodbatlasProcessPageFaults:                         MetricConfig{Enabled: false},
					MongodbatlasProcessRestarts:                           MetricConfig{Enabled: false},
					MongodbatlasProcessTickets:                            MetricConfig{Enabled: false},
					MongodbatlasSystemCPUNormalizedUsageAverage:           MetricConfig{Enabled: false},
					MongodbatlasSystemCPUNormalizedUsageMax:               MetricConfig{Enabled: false},
					MongodbatlasSystemCPUUsageAverage:                     MetricConfig{Enabled: false},
					MongodbatlasSystemCPUUsageMax:                         MetricConfig{Enabled: false},
					MongodbatlasSystemFtsCPUNormalizedUsage:               MetricConfig{Enabled: false},
					MongodbatlasSystemFtsCPUUsage:                         MetricConfig{Enabled: false},
					MongodbatlasSystemFtsDiskUsed:                         MetricConfig{Enabled: false},
					MongodbatlasSystemFtsMemoryUsage:                      MetricConfig{Enabled: false},
					MongodbatlasSystemMemoryUsageAverage:                  MetricConfig{Enabled: false},
					MongodbatlasSystemMemoryUsageMax:                      MetricConfig{Enabled: false},
					MongodbatlasSystemNetworkIoAverage:                    MetricConfig{Enabled: false},
					MongodbatlasSystemNetworkIoMax:                        MetricConfig{Enabled: false},
					MongodbatlasSystemPagingIoAverage:                     MetricConfig{Enabled: false},
					MongodbatlasSystemPagingIoMax:                         MetricConfig{Enabled: false},
					MongodbatlasSystemPagingUsageAverage:                  MetricConfig{Enabled: false},
					MongodbatlasSystemPagingUsageMax:                      MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					MongodbAtlasClusterName:     ResourceAttributeConfig{Enabled: false},
					MongodbAtlasDbName:          ResourceAttributeConfig{Enabled: false},
					MongodbAtlasDiskPartition:   ResourceAttributeConfig{Enabled: false},
					MongodbAtlasHostName:        ResourceAttributeConfig{Enabled: false},
					MongodbAtlasOrgName:         ResourceAttributeConfig{Enabled: false},
					MongodbAtlasProcessID:       ResourceAttributeConfig{Enabled: false},
					MongodbAtlasProcessPort:     ResourceAttributeConfig{Enabled: false},
					MongodbAtlasProcessTypeName: ResourceAttributeConfig{Enabled: false},
					MongodbAtlasProjectID:       ResourceAttributeConfig{Enabled: false},
					MongodbAtlasProjectName:     ResourceAttributeConfig{Enabled: false},
					MongodbAtlasProviderName:    ResourceAttributeConfig{Enabled: false},
					MongodbAtlasRegionName:      ResourceAttributeConfig{Enabled: false},
					MongodbAtlasUserAlias:       ResourceAttributeConfig{Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadMetricsBuilderConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(MetricConfig{}, ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadMetricsBuilderConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}

func TestResourceAttributesConfig(t *testing.T) {
	tests := []struct {
		name string
		want ResourceAttributesConfig
	}{
		{
			name: "default",
			want: DefaultResourceAttributesConfig(),
		},
		{
			name: "all_set",
			want: ResourceAttributesConfig{
				MongodbAtlasClusterName:     ResourceAttributeConfig{Enabled: true},
				MongodbAtlasDbName:          ResourceAttributeConfig{Enabled: true},
				MongodbAtlasDiskPartition:   ResourceAttributeConfig{Enabled: true},
				MongodbAtlasHostName:        ResourceAttributeConfig{Enabled: true},
				MongodbAtlasOrgName:         ResourceAttributeConfig{Enabled: true},
				MongodbAtlasProcessID:       ResourceAttributeConfig{Enabled: true},
				MongodbAtlasProcessPort:     ResourceAttributeConfig{Enabled: true},
				MongodbAtlasProcessTypeName: ResourceAttributeConfig{Enabled: true},
				MongodbAtlasProjectID:       ResourceAttributeConfig{Enabled: true},
				MongodbAtlasProjectName:     ResourceAttributeConfig{Enabled: true},
				MongodbAtlasProviderName:    ResourceAttributeConfig{Enabled: true},
				MongodbAtlasRegionName:      ResourceAttributeConfig{Enabled: true},
				MongodbAtlasUserAlias:       ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				MongodbAtlasClusterName:     ResourceAttributeConfig{Enabled: false},
				MongodbAtlasDbName:          ResourceAttributeConfig{Enabled: false},
				MongodbAtlasDiskPartition:   ResourceAttributeConfig{Enabled: false},
				MongodbAtlasHostName:        ResourceAttributeConfig{Enabled: false},
				MongodbAtlasOrgName:         ResourceAttributeConfig{Enabled: false},
				MongodbAtlasProcessID:       ResourceAttributeConfig{Enabled: false},
				MongodbAtlasProcessPort:     ResourceAttributeConfig{Enabled: false},
				MongodbAtlasProcessTypeName: ResourceAttributeConfig{Enabled: false},
				MongodbAtlasProjectID:       ResourceAttributeConfig{Enabled: false},
				MongodbAtlasProjectName:     ResourceAttributeConfig{Enabled: false},
				MongodbAtlasProviderName:    ResourceAttributeConfig{Enabled: false},
				MongodbAtlasRegionName:      ResourceAttributeConfig{Enabled: false},
				MongodbAtlasUserAlias:       ResourceAttributeConfig{Enabled: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadResourceAttributesConfig(t *testing.T, name string) ResourceAttributesConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	sub, err = sub.Sub("resource_attributes")
	require.NoError(t, err)
	cfg := DefaultResourceAttributesConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
