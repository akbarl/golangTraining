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
		Category: entities.Category{
			Id:   "c1",
			Name: "category 1",
		},
		Comments: []entities.Comment{
			entities.Comment{Id: "cm1", Content: "content 1"},
			entities.Comment{Id: "cm2", Content: "content 2"},
			entities.Comment{Id: "cm3", Content: "content 3"},
			entities.Comment{Id: "cm4", Content: "content 4"},
		},
	}

	result, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(result))
	}
}

func Demo2() {
	s := `{"id":"p01","name":"name 1","price":4.5,"category":{"id":"c1","name":"category 1","content":""},"comments":[{"Id":"cm1","Content":"content 1"},{"Id":"cm2","Content":"content 2"},{"Id":"cm3","Content":"content 3"},{"Id":"cm4","Content":"content 4"}]}`
	var product entities.Product
	json.Unmarshal([]byte(s), &product)
	fmt.Println("Id: ", product.Id)
	fmt.Println("Name: ", product.Name)
	fmt.Println("Price: ", product.Price)
	for _, comment := range product.Comments {
		fmt.Println()
		fmt.Printf("comment: %s", comment.Content)
		fmt.Println()
	}
}
