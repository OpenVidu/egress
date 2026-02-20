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
	// Enables system-wide CPU monitoring instead of tracking only egress subprocess
	// CPU usage. When enabled, the monitor uses the actual idle CPU of the host to
	// determine if new egress requests can be accepted. This is useful when other
	// CPU-intensive processes run on the same host alongside egress.
	UseGlobalCpuMonitoring bool `yaml:"use_global_cpu_monitoring,omitempty"`
	// Minimum available disk space in MB required to accept new egress requests.
	// Default: 512 MB. Set to a negative value (e.g., -1) to disable disk space checking.
	MinDiskSpaceMB float64 `yaml:"min_disk_space_mb,omitempty"`
}

// END OPENVIDU BLOCK
