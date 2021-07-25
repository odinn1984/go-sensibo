package models

type DeviceTimer struct {
	ID                     string
	IsEnabled              bool
	ACState                ACStateData
	CausedBy               CausedBy
	CreateTime             string
	CreateTimeSecondsAgo   int
	LastScheduledInstances []struct {
		Type                 string
		TargetTime           string
		TargetTimeSecondsAgo int
		Status               string
		ScheduleId           string
		LastExecutions       []string
	}
	TargetTime               string
	TargetTimeSecondsFromNow int
}
