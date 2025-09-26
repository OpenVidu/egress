// BEGIN OPENVIDU BLOCK
package openviduconfig

type OpenViduConfig struct {
	// controls how StartEgressAffinity is computed by the egress server
	// supported values:
	//  - "binpack": default behavior (0 if no capacity, 0.5 if idle, 1 if busy)
	//  - "cpuload": prefer nodes with more available CPU (normalized 0..1)
	AllocationStrategy string `yaml:"allocation_strategy,omitempty"`
}

// END OPENVIDU BLOCK
