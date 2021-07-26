// Copyright 2021 To Levan Giguashvili. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package models

// Device holds extended information about the device.
type Device struct {
	IsGeofenceOnEnterEnabledForThisUser             bool
	IsClimateReactGeofenceOnEnterEnabledForThisUser bool
	IsMotionGeofenceOnEnterEnabled                  bool
	IsOwner                                         bool
	ID                                              string
	QRID                                            string
	TemperatureUnit                                 string
	Room                                            Room
	ACState                                         struct {
		Timestamp SensiboTime
		ACStateData
	}
	Location struct {
		ID                    string
		Name                  string
		LatLon                []float64
		Address               []string
		Country               string
		CreateTime            SensiboTime
		UpdateTime            SensiboTime
		GeofenceTriggerRadius float64
		Subscription          string
		Occupancy             string
	}
	ConnectionStatus struct {
		IsAlive  bool
		lastSeen SensiboTime
	}
	FirmwareVersion                      string
	FirmwareType                         string
	ProductModel                         string
	ConfigGroup                          string
	CurrentlyAvailableFirmwareVersion    string
	CleanFiltersNotificationEnabled      bool
	ShouldShowFilterCleaningNotification bool
	IsGeofenceOnExitEnabled              bool
	IsClimateReactGeofenceOnExitEnabled  bool
	IsMotionGeofenceOnExitEnabled        bool
	SensorsCalibration                   struct {
		Temperature float64
		Humidity    float64
	}
	MotionSensors   []string
	Tags            []string
	Timer           DeviceTimer
	Schedules       []DeviceSchedule
	MotionConfig    string
	FiltersCleaning struct {
		ACOnSecondsSinceLastFiltersClean float64
		FiltersCleanSecondsThreshold     float64
		LastFiltersCleanTime             string
		ShouldCleanFilters               bool
	}
	RoomIsOccupied         string
	MainMeasurementsSensor string
	PureBoostConfig        string
	WarrantyEligible       string
	Features               []string
	RunningHealthcheck     string
	HomekitSupported       bool
	RemoteCapabilities     struct {
		Modes map[string]Mode
	}
	Remote struct {
		Toggle bool
		Window bool
	}
	RemoteFlavor       string
	RemoteAlternatives []string
	SmartMode          ClimateReact
	Measurements       struct {
		Temperature    float64
		Humidity       float64
		Time           SensiboTime
		Rssi           float64
		BatteryVoltage string
		Piezo          []string
		Pm25           float64
		Tvoc           float64
		Co2            float64
	}
	AccessPofloat64 struct {
		SSID     string
		Password string
	}
	MacAddress string
}

// Room information.
type Room struct {
	UID  string
	Name string
	Icon string
}

// SensiboTime holds general time information data structure that is re-usable.
type SensiboTime struct {
	Time       string
	SecondsAgo float64
}

// Mode of the AC unit.
type Mode struct {
	Temperatures    map[string]Temperature
	FanLevels       []string
	Swing           []string
	HorizontalSwing []string
	Light           []string
}

// Temperature histogram.
type Temperature struct {
	IsNative bool
	Values   []float64
}

// ClimateReactState holds extended AC data for climate react.
type ClimateReactState struct {
	ACStateData
	HorizontalSwing string
	Light           string
}

// ClimateReact holds climate react data.
type ClimateReact struct {
	Enabled                  bool
	Type                     string
	DeviceUID                string
	LowTemperatureThreshold  float64
	HighTemperatureThreshold float64
	LowTemperatureState      ClimateReactState
	HighTemperatureState     ClimateReactState
	LowTemperatureWebhook    string
	HighTemperatureWebhook   string
}
