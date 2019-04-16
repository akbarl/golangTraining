package marshal

import (
	"encoding/json"
	"entities"
	"fmt"
)

func Demo1() {
	product := entities.Product{
		Id:    "p01",
		Name:  "name 1",
		Price: 4.5,
	}

	result, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(result))
	}
}
