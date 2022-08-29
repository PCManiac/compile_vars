package compile_vars

import (
	_ "github.com/PCManiac/logrus_init"
	log "github.com/sirupsen/logrus"
)

var (
	version    string
	build_time string
)

func init() {
	if version == "" {
		version = "not set"
	}
	if build_time == "" {
		build_time = "not set"
	}

	log.WithFields(log.Fields{
		"version":    version,
		"build_time": build_time,
	}).Info("Initializing.")
}

func GetVersion() string {
	return version
}

func GetBuildTime() string {
	return build_time
}
