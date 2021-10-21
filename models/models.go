package models

type Item struct {
	Id          int64  `json:"id"`
	BaNo        string `json:"ba_no"`
	Sqn         string `json:"sqn"`
	VehicleType string `json:"vt"`
	Type        string `json:"type"`
	Remarks     string `json:"remarks"`
}

type Vehicle struct {
	Item
	Efc              int64  `json:"efc"`
	EngineHours      int64  `json:"eh"`
	Kilometers       int64  `json:"km"`
	SeriesInspection string `json:"si"`
	Tag              string `json:"tag"`
	TM1              string `json:"tm_1"`
	TM2              string `json:"tm_2"`
	CMSIn            string `json:"cms_in"`
	CMSOut           string `json:"cms_out"`
	WorkshopIn       string `json:"workshop_in"`
	WorkshopOut      string `json:"workshop_out"`
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
	ControlNumber     string `json:"cn"`
	DemandNumber      string `json:"dn"`
	Depot             string `json:"depot"`
	Status            string `json:"status"`
	EquipmentDemanded string `json:"ed"`
}

type Ids struct {
	Ids []int64 `json:"ids"`
}

type User struct {
	Username       string `json:"username"`
	Password       string `json:"password,omitempty"`
	Authentication string `json:"authentication,omitempty"`
}
