package httpmodels

import (
	"github.com/FlipsideCrypto/databricks-sdk-golang/azure/clusters/models"
)

type ListNodeTypesRespItem struct {
	NodeTypeID     string                        `json:"node_type_id,omitempty" url:"node_type_id,omitempty"`
	MemoryMb       int32                         `json:"memory_mb,omitempty" url:"memory_mb,omitempty"`
	NumCores       float32                       `json:"num_cores,omitempty" url:"num_cores,omitempty"`
	Description    string                        `json:"description,omitempty" url:"description,omitempty"`
	InstanceTypeID string                        `json:"instance_type_id,omitempty" url:"instance_type_id,omitempty"`
	IsDeprecated   bool                          `json:"is_deprecated,omitempty" url:"is_deprecated,omitempty"`
	NodeInfo       *models.ClusterCloudProviderNodeInfo `json:"node_info,omitempty" url:"node_info,omitempty"`
}
