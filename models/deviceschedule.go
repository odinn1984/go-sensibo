package models

type DeviceSchedule struct {
	ID        string
	IsEnabled bool
	ACState   struct {
		ACStateData
		Extra struct {
			Scheduler struct {
				ClimateReact         bool `json:"climate_react"`
				Motion               string
				On                   bool
				ClimateReactSettings SmartMode `json:"climate_react_settings"`
				PureBoost            string    `json:"pure_boost"`
			}
		}
		HorizontalSwing string
		Light           string
	}
	CausedBy               CausedBy
	CreateTime             string
	CreateTimeSecondsAgo   int
	RecurringDays          []string
	TargetTimeLocal        string
	TimeZone               string
	PodUID                 string
	NextTime               string
	NextTimeSecondsFromNow int
}
