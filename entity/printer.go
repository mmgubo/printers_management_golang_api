package entity

type Printer struct {
	ID    int64  `json:"id"`
	Name string `json:"name"`
	Ip  string `json:"ip"`
	Status  string `json:"status"`
}
