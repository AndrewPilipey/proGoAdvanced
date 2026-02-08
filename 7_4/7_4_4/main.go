type Table struct {
    Number     int
    Guests     []string
    IsReserved bool
    BillAmount float64
}

type Waiter struct {
    Name   string
    Tables []Table
}


table := Table{1, []{"Гость"}, true, 31.50}
waiter := Waiter{"Анна", []Table{table}}