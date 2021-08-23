package builder

type Demand interface {
	ListDemands() string
}

type demand struct {}

func NewDemand() Demand {
	return &demand{}
}

func (d *demand) ListDemands() string {
	return "SELECT * FROM demands;"
}