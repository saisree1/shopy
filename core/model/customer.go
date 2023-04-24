package model

type Customer struct {
	Id           string
	Name         string
	MobileNumber string
}

func getDefaultCustomers() map[string]Customer {
	return map[string]Customer{
		"c1": {
			Id:           "c1",
			Name:         "Raju",
			MobileNumber: "9898989898",
		},
		"c2": {
			Id:           "c2",
			Name:         "Hari",
			MobileNumber: "9898989899",
		},
	}
}
