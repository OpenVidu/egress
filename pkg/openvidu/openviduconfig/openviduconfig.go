// BEGIN OPENVIDU BLOCK
package openviduconfig

type OpenViduConfig struct {
	// Controls how StartEgressAffinity is computed by the egress server.
	// Supported values:
	//  - "binpack": original behavior (0 if no capacity, 0.5 if idle, 1 if busy)
	//  - "cpuload": prefer nodes with more available CPU (normalized 0..1)
	AllocationStrategy string `yaml:"allocation_strategy,omitempty"`
	// Disables the parent process ability to kill egress subprocesses.
	// By default if a high CPU usage is detected for a sustained period of time
	// the most CPU consuming egress will be killed.
	DisableCpuOverloadKiller bool `yaml:"disable_cpu_overload_killer,omitempty"`
}

// END OPENVIDU BLOCK
