package servicefabric

// ApplicationItemsPage encapsulates the paged response
// model for Applications in the Service Fabric API
type ApplicationItemsPage struct {
	ContinuationToken *string           `json:"ContinuationToken"`
	Items             []ApplicationItem `json:"Items"`
}

// ApplicationItem encapsulates the embedded model for
// ApplicationItems within the ApplicationItemsPage model
type ApplicationItem struct {
	HealthState string `json:"HealthState"`
	ID          string `json:"Id"`
	Name        string `json:"Name"`
	Parameters  []*struct {
		Key   string `json:"Key"`
		Value string `json:"Value"`
	} `json:"Parameters"`
	Status      string `json:"Status"`
	TypeName    string `json:"TypeName"`
	TypeVersion string `json:"TypeVersion"`
}

// ServiceItemsPage encapsulates the paged response
// model for Services in the Service Fabric API
type ServiceItemsPage struct {
	ContinuationToken *string       `json:"ContinuationToken"`
	Items             []ServiceItem `json:"Items"`
}

// ServiceItem encapsulates the embedded model for
// ServiceItems within the ServiceItemsPaage model
type ServiceItem struct {
	HasPersistedState bool   `json:"HasPersistedState"`
	HealthState       string `json:"HealthState"`
	ID                string `json:"Id"`
	IsServiceGroup    bool   `json:"IsServiceGroup"`
	ManifestVersion   string `json:"ManifestVersion"`
	Name              string `json:"Name"`
	ServiceKind       string `json:"ServiceKind"`
	ServiceStatus     string `json:"ServiceStatus"`
	TypeName          string `json:"TypeName"`
}

// ServiceItemExtended provides a flattened view
// of the service with details of the application
// it belongs too and the replicas/partitions
type ServiceItemExtended struct {
	ServiceItem
	HasHTTPEndpoint bool
	IsHealthy       bool
	Application     ApplicationItem
	Partitions      []PartitionItemExtended
}

// PartitionItemsPage encapsulates the paged response
// model for ParititonItems in the Service Fabric API
type PartitionItemsPage struct {
	ContinuationToken *string         `json:"ContinuationToken"`
	Items             []PartitionItem `json:"Items"`
}

// PartitionItem encapsulates the service information
// returned for each PartitionItem under the service
type PartitionItem struct {
	CurrentConfigurationEpoch struct {
		ConfigurationVersion string `json:"ConfigurationVersion"`
		DataLossVersion      string `json:"DataLossVersion"`
	} `json:"CurrentConfigurationEpoch"`
	HealthState          string `json:"HealthState"`
	MinReplicaSetSize    int64  `json:"MinReplicaSetSize"`
	PartitionInformation struct {
		HighKey              string `json:"HighKey"`
		ID                   string `json:"Id"`
		LowKey               string `json:"LowKey"`
		ServicePartitionKind string `json:"ServicePartitionKind"`
	} `json:"PartitionInformation"`
	PartitionStatus      string `json:"PartitionStatus"`
	ServiceKind          string `json:"ServiceKind"`
	TargetReplicaSetSize int64  `json:"TargetReplicaSetSize"`
}

// PartitionItemExtended provides a flattened view
// of a services partitions
type PartitionItemExtended struct {
	PartitionItem
	HasReplicas  bool
	Replicas     []ReplicaItem
	HasInstances bool
	Instances    []InstanceItem
}

// ReplicaItemBase shared data used
// in both replicas and instances
type ReplicaItemBase struct {
	Address                      string `json:"Address"`
	HealthState                  string `json:"HealthState"`
	LastInBuildDurationInSeconds string `json:"LastInBuildDurationInSeconds"`
	NodeName                     string `json:"NodeName"`
	ReplicaRole                  string `json:"ReplicaRole"`
	ReplicaStatus                string `json:"ReplicaStatus"`
	ServiceKind                  string `json:"ServiceKind"`
}

// ReplicaItemsPage encapsulates the response
// model for Replicas in the Service Fabric API
type ReplicaItemsPage struct {
	ContinuationToken *string       `json:"ContinuationToken"`
	Items             []ReplicaItem `json:"Items"`
}

// ReplicaItem holds replica specific data
type ReplicaItem struct {
	*ReplicaItemBase
	ID string `json:"ReplicaId"`
}

// ReplicaInstance interface provides a unified interface
// over replicas and instances
type ReplicaInstance interface {
	GetReplicaData() (string, *ReplicaItemBase)
}

// GetReplicaData returns replica data
func (m *ReplicaItem) GetReplicaData() (string, *ReplicaItemBase) {
	return m.ID, m.ReplicaItemBase
}

// InstanceItemsPage encapsulates the response
// model for Instances in the Service Fabric API
type InstanceItemsPage struct {
	ContinuationToken *string        `json:"ContinuationToken"`
	Items             []InstanceItem `json:"Items"`
}

// InstanceItem hold instance specific data
type InstanceItem struct {
	*ReplicaItemBase
	ID string `json:"InstanceId"`
}

// GetReplicaData returns replica data from an instance
func (m *InstanceItem) GetReplicaData() (string, *ReplicaItemBase) {
	return m.ID, m.ReplicaItemBase
}

// ServiceType encapsulates the response model for
// Service types in the Service Fabric API
type ServiceType struct {
	ServiceTypeDescription struct {
		IsStateful               bool           `json:"IsStateful"`
		ServiceTypeName          string         `json:"ServiceTypeName"`
		PlacementConstraints     string         `json:"PlacementConstraints"`
		HasPersistedState        bool           `json:"HasPersistedState"`
		Kind                     string         `json:"Kind"`
		Extensions               []KeyValuePair `json:"Extensions"`
		LoadMetrics              []interface{}  `json:"LoadMetrics"`
		ServicePlacementPolicies []interface{}  `json:"ServicePlacementPolicies"`
	} `json:"ServiceTypeDescription"`
	ServiceManifestVersion string `json:"ServiceManifestVersion"`
	ServiceManifestName    string `json:"ServiceManifestName"`
	IsServiceGroup         bool   `json:"IsServiceGroup"`
}

// ServiceExtensionData encapsulates the extension model
// for any extensions defined in a Service's manifest
type ServiceExtensionData map[string]map[string]string

// KeyValuePair represents a key value pair structure
type KeyValuePair struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}
