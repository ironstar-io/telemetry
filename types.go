package telemetry

import "time"

// User carries global settings and usage information for a single installation of Tokaido
type User struct {
	// Timestamp is the submission date and time in UTC
	Timestamp time.Time `firestore:"timestamp,serverTimestamp" json:"timestamp"`
	// TelemetryID is the users telemetry UUID
	TelemetryID string `firestore:"telemetry_id" json:"telemetry_id"`
}

// UserCheckin is a subcollection of a user that records data with each user 'up'
type UserCheckin struct {
	// Timestamp is the submission date and time in UTC
	Timestamp time.Time `firestore:"timestamp" json:"timestamp"`
	// OperatingSystemType ...
	OperatingSystemType string `firestore:"operating_system_type" json:"operating_system_type"`
	// OperatingSystemVersion ...
	OperatingSystemVersion string `firestore:"operating_system_version" json:"operating_system_version"`
	// Xdebug is set to true if the user uses xdebug in any of their projects
	Xdebug bool `firestore:"xdebug" json:"xdebug"`
	// ProjectCount is the number of Tokaido projects configured on the users' system
	ProjectCount int `firestore:"project_count" json:"project_count"`
	// TokaidoVersion ...
	TokaidoVersion string `firestore:"tokaido_version" json:"tokaido_version"`
	// SyncStrategy being either 'docker', 'unison', or 'fusion'
	SyncStrategy string `firestore:"sync_strategy" json:"sync_strategy"`
	// Country being the user's detected country such as "USA", or "Australia"
	Country string `firestore:"country" json:"country"`
}

// DrupalProject is a collection of DrupalProjectCheckins
type DrupalProject struct {
	// Timestamp is the submission date and time in UTC
	Timestamp time.Time `firestore:"timestamp,serverTimestamp" json:"timestamp"`
	// ProjectID is the users telemetry UUID
	ProjectID string `firestore:"project_id" json:"project_id"`
}

// DrupalProjectCheckin carries project-level configuration entries submitted each time `up` is run on a project
type DrupalProjectCheckin struct {
	// Timestamp is the submission date and time in UTC
	Timestamp time.Time `firestore:"timestamp" json:"timestamp"`
	// TelemetryID is the UUID of the Tokaido installation sending this update
	TelemetryID string `firestore:"telemetry_id" json:"telemetry_id"`
	// PhpVersion ...
	PhpVersion string `firestore:"php_version" json:"php_version"`
	// Mailhog is true if mailhog is enabled
	Mailhog bool `firestore:"mailhog" json:"mailhog"`
	// Adminer is true if adminer is enabled
	Adminer bool `firestore:"adminer" json:"adminer"`
	// Solr is true if solr is enabled
	Solr bool `firestore:"solr" json:"solr"`
	// Redis is true if redis is enabled
	Redis bool `firestore:"redis" json:"redis"`
	// Memcache is true if memcache is enabled
	Memcache bool `firestore:"memcache" json:"memcache"`
	// DatabaseEngine in use, such as 'mysql' or 'mariadb'
	DatabaseEngine string `firestore:"database_engine" json:"database_engine"`
	// DrupalVersion is the detected Drupal Major version, such as 8 or 9.
	DrupalVersion string `firestore:"drupal_version" json:"drupal_version"`
	// Stability identifies the 'edge', 'stable', or 'experimental' release set
	Stability string `firestore:"stability" json:"stability"`
	// PHPMemory is the PHP memory limit
	PHPMemory string `firestore:"php_memorylimit" json:"php_memorylimit"`
	// Duration is the number of seconds it took to complete the startup
	Duration int `firestore:"duration" json:"startup_seconds"`
}

// CommandCheckin carries information about which commands were run on a system
type CommandCheckin struct {
	// Timestamp is the submission date and time in UTC
	Timestamp time.Time `firestore:"timestamp,serverTimestamp" json:"timestamp"`
	// ProjectID is the project's UUID
	ProjectID string `firestore:"project_id" json:"project_id"`
	// TelemetryID is the UUID of the Tokaido installation sending this document
	TelemetryID string `firestore:"telemetry_id" json:"telemetry_id"`
	// Command is the name of the command that was run
	Command string `firestore:"command" json:"command"`
}

// SurveyResponse provides feedback using `tok survey`
type SurveyResponse struct {
	// Timestamp is the submission date and time in UTC
	Timestamp time.Time `firestore:"timestamp,serverTimestamp" json:"timestamp"`
	// TelemetryID is the UUID of the Tokaido installation sending this document
	TelemetryID string `firestore:"telemetry_id" json:"telemetry_id"`
	// TokaidoVersion ...
	TokaidoVersion string `firestore:"tokaido_version" json:"tokaido_version"`
	// Satisfaction on a scale of 1-5, 5 being most satisfied
	Satisfaction string `firestore:"satisfaction" json:"satisfaction"`
	// NextFeature is the name of planned features the user most wants to see
	NextFeature string `firestore:"nextfeature" json:"nextfeature"`
	// Message is a freeform feedback message from the user
	Message string `firestore:"message" json:"message"`
}
