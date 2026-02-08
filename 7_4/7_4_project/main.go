type Warehouse struct {
	Name     string
	ID       string
	District District
	Capacity Capacity
	Stock    Stock
	Director Director
}

type District struct {
	ID            string
	Name          string
	Warehouses    []Warehouse
	TotalArea     float64
	TotalCapacity int
}

type Capacity struct {
	Area            float64
	StorageBinCount int
	EmptyStorageBin int
}

type Stock struct {
	WeeklyDemand int
	CurrentStock int
}

type Director struct {
	ID         string
	Name       string
	StaffCount int
	Salary     float64
}

