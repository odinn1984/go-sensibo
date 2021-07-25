package models

type ACState struct {
	ID      string
	Status  string
	ACState struct {
		Timestamp SensiboTime
		ACStateData
	}
	ChangedProperties []string
	Reason            string
	FailureReason     string
}
