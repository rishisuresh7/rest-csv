package models

type Item struct {
	Id      string `json:"id"`
	BaNo    string `json:"baNo"`
	CDR     string `json:"cdr"`
	Driver  string `json:"driver"`
	Oper    string `json:"oper"`
	Tm_1    string `json:"tm_1"`
	Tm_2    string `json:"tm_2"`
	Demand  string `json:"demand"`
	Fault   string `json:"fault"`
	Remarks string `json:"remarks"`
}

type Ids struct {
	Ids []int64 `json:"ids"`
}

type User struct {
	Username       string `json:"username"`
	Password       string `json:"password,omitempty"`
	Authentication string `json:"authentication,omitempty"`
}
