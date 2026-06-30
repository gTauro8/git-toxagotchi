package domain

import "time"

type EventType string

const (
	EventCommit          EventType = "commit"
	EventPush            EventType = "push"
	EventMerge           EventType = "merge"
	EventTestPassed      EventType = "test_passed"
	EventTestFailed      EventType = "test_failed"
	EventIdle            EventType = "idle"
	EventDependencyAdded EventType = "dependency_added"
	EventSecretDetected  EventType = "secret_detected"
	EventForcePush       EventType = "force_push"
)

type Event struct {
	Type      EventType
	Timestamp time.Time
	Metadata  map[string]interface{}
}

type CommitAnalysis struct {
	MessageQuality  int
	DiffSize        int
	FilesChanged    int
	TodosAdded      int
	SecretsDetected bool
	DepsAdded       int
	Message         string
}
