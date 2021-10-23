package models

type Item struct {
	Id          int64  `json:"id"`
	Sqn         string `json:"squadron"`
	VehicleType string `json:"vehicleType"`
	BaNo        string `json:"baNumber"`
	Type        string `json:"type"`
	Remarks     string `json:"remarks"`
}

type Vehicle struct {
	Item
	Kilometers       int64  `json:"kilometers"`
	EngineHours      int64  `json:"engineHours"`
	Efc              int64  `json:"efc"`
	TM1              string `json:"tm1"`
	TM2              string `json:"tm2"`
	CMSIn            string `json:"cmsIn"`
	CMSOut           string `json:"cmsOut"`
	WorkshopIn       string `json:"workshopIn"`
	WorkshopOut      string `json:"workshopOut"`
	MR1              string `json:"mr1"`
	MR2              string `json:"mr2"`
	FDFiring         string  `json:"fdFiring"`
	SeriesInspection string `json:"seriesInspection"`
	Trg              string `json:"trg"`
}

type Alert struct {
	Id         int64  `json:"id"`
	Name       string `json:"alertName"`
	BaNo       string `json:"ba_number"`
	AlertField string `json:"fieldName"`
	LastValue  string `json:"lastValue"`
	NextValue  string `json:"nextValue"`
	Notify     bool   `json:"-"`
	Remarks    string `json:"remarks"`
}

type Notification struct {
	AlertId        int64  `json:"alertId"`
	VehicleId      int64  `json:"vehicleId"`
	AlertName      string `json:"alertName"`
	BaNo           string `json:"baNumber"`
	VehicleType    string `json:"vehicleType"`
	AlertField     string `json:"fieldName"`
	LastValue      string `json:"lastValue"`
	NextValue      string `json:"nextValue"`
	VehicleRemarks string `json:"vehicleRemarks"`
	AlertRemarks   string `json:"alertRemarks"`
}

type Demand struct {
	Item
	ControlNumber     string `json:"controlNumber"`
	ControlDate       string `json:"controlDate"`
	DemandNumber      string `json:"demandNumber"`
	DemandDate        string `json:"demandDate"`
	Depot             string `json:"depot"`
	Status            string `json:"status"`
	EquipmentDemanded string `json:"equipmentDemanded"`
}

type Ids struct {
	Ids []int64 `json:"ids"`
}

type User struct {
	Username       string `json:"username"`
	Password       string `json:"password,omitempty"`
	Authentication string `json:"authentication,omitempty"`
}
