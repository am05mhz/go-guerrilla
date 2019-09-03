package guerrilla

import "time"

var (
	// Version ...
	Version string
	// Commit ...
	Commit string
	// BuildTime ...
	BuildTime string

	// StartTime ...
	StartTime time.Time
	// ConfigLoadTime ...
	ConfigLoadTime time.Time
)

func init() {
	// If version, commit, or build time are not set, make that clear.
	const unknown = "unknown"
	if Version == "" {
		Version = unknown
	}
	if Commit == "" {
		Commit = unknown
	}
	if BuildTime == "" {
		BuildTime = unknown
	}

	StartTime = time.Now()
}
