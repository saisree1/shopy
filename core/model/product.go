package model

type Product struct {
	Id       string
	Name     string
	Price    float64
	Category string
	Quantity int
}

func getDefaultProducts() map[string]Product {
	return map[string]Product{
		"p1": {
			Id:       "p1",
			Name:     "Reynolds Pen",
			Price:    10.00,
			Category: "Regualr",
			Quantity: 100,
		},
		"p2": {
			Id:       "p2",
			Name:     "FastTrack Watch",
			Price:    1999.0,
			Category: "Premium",
			Quantity: 20,
		},
		"p3": {
			Id:       "p3",
			Name:     "Paragon Shoes",
			Price:    500.00,
			Category: "Budget",
			Quantity: 100,
		},
		"p4": {
			Id:       "p4",
			Name:     "Fossil Watch",
			Price:    19990.0,
			Category: "Premium",
			Quantity: 25,
		},
		"p5": {
			Id:       "p5",
			Name:     "Lenovo Laptop",
			Price:    60000.00,
			Category: "Budget",
			Quantity: 100,
		},
		"p6": {
			Id:       "p6",
			Name:     "Air Jordans",
			Price:    10999.99,
			Category: "Premium",
			Quantity: 15,
		},
		"p7": {
			Id:       "p7",
			Name:     "Reynolds Pen",
			Price:    10.00,
			Category: "Regualr",
			Quantity: 100,
		},
		"p8": {
			Id:       "p8",
			Name:     "Titan Watch",
			Price:    999.0,
			Category: "Budget",
			Quantity: 200,
		},
		"p9": {
			Id:       "p9",
			Name:     "Fountain Pen",
			Price:    450.00,
			Category: "Premium",
			Quantity: 10,
		},
		"p10": {
			Id:       "p10",
			Name:     "Bata Shoes",
			Price:    1099.0,
			Category: "Regular",
			Quantity: 300,
		},
		"p11": {
			Id:       "p11",
			Name:     "Track Pant",
			Price:    150.00,
			Category: "Budget",
			Quantity: 175,
		},
	}
}
