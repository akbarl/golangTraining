package demostruct

import (
	"entities"
	"fmt"
	"sort"
	"strings"
)

func Demo1() {
	var product entities.Product
	product.Id = "p01"
	product.Name = "name 1"
	product.Price = 4.5
	product.Quantity = 2
	product.Status = true
	fmt.Println("Product 1 info")
	fmt.Println("id: ", product.Id)
	fmt.Println("name: ", product.Name)
	fmt.Println("price: ", product.Price)
	fmt.Println("Quatity: ", product.Quantity)
	fmt.Println("Totol: ", product.Price*float64(product.Quantity))
}

func Demo2() {
	product := entities.Product{
		Id:       "p01",
		Name:     "name 1",
		Price:    4.5,
		Quantity: 2,
		Status:   true,
	}

	fmt.Println("Product 1 info")
	fmt.Println("id: ", product.Id)
	fmt.Println("name: ", product.Name)
	fmt.Println("price: ", product.Price)
	fmt.Println("Quatity: ", product.Quantity)
	fmt.Println("Totol: ", product.Price*float64(product.Quantity))

}

func Display(product entities.Product) {
	fmt.Println("id: ", product.Id)
	fmt.Println("name: ", product.Name)
	fmt.Println("price: ", product.Price)
	fmt.Println("Quatity: ", product.Quantity)
	fmt.Println("Totol: ", product.Price*float64(product.Quantity))
	fmt.Println("=======================")
}

func DisplayList(products []entities.Product) {
	for _, product := range products {
		Display(product)
	}
}
func Demo3() {
	product := entities.Product{
		Id:       "p01",
		Name:     "Name 1",
		Price:    2,
		Quantity: 8,
		Status:   true,
	}

	p := &product
	fmt.Println("id: ", p.Id)
	fmt.Println("name: ", (*p).Name)
	p.Quantity = 20
	Display(product)

	changeValue2(p)
	Display(product)
}

func changeValue2(p *entities.Product) {
	p.Price = 6
}

func Demo4() {
	products := []entities.Product{}
	products = append(products, entities.Product{
		Id:       "p01",
		Name:     "name 1",
		Price:    4.5,
		Quantity: 2,
		Status:   true,
	})

	products = append(products, entities.Product{
		Id:       "p02",
		Name:     "name 2",
		Price:    2,
		Quantity: 2,
		Status:   true,
	})

	products = append(products, entities.Product{
		Id:       "p03",
		Name:     "name 3",
		Price:    3,
		Quantity: 2,
		Status:   true,
	})

	fmt.Println("Product List 1")
	for i := 0; i < len(products); i++ {
		Display(products[i])
		fmt.Println("----------------------------")
	}

	fmt.Println("Product list 2")
	for _, product := range products {
		Display(product)
		fmt.Println("----------------------------")
	}

	fmt.Println("Total price of all products: ", SumPrice(products))
	maxProduct, minProduct := FindMaxAndMin(products)
	fmt.Println("Max product: ")
	Display(maxProduct)
	fmt.Println("Min product: ")
	Display(minProduct)

	foundProduct := findWithKeyWord(products, "1")
	fmt.Println("Found product: ", foundProduct)

	fmt.Println("Sorted product by price")
	DisplayList(sortProductByPrice(products))
}

func displayProducts(products []entities.Product) {
	for _, product := range products {
		Display(product)
		fmt.Println("----------------------------")
	}
}

/*
Su dung ham thuc hien cac yeu cau sau:
1. Tinh tong tien tat ca san pham. Tra ve tong tien
2. Tim thong tin san pham co gia tri max, min. Tra ve thong tin 2 san pham do
3. Liet ke danh sach cac san pham co ten chua 1 keyword. Tra ve danh sach cac
san pham do.
4. su dung package sort sap xep cac san pham theo gia giam dan
5. Dem co bao nhieu san pham co gia nam trong khoang tu min den max.
Tra ve ket qua
*/

func SumPrice(products []entities.Product) (sumPrice float64) {
	for _, product := range products {
		sumPrice += product.Price * float64(product.Quantity)
	}
	return sumPrice
}

func FindMaxAndMin(products []entities.Product) (maxProduct entities.Product, minProduct entities.Product) {
	maxProductId := 0
	minProductId := 0
	for index, product := range products {
		if products[maxProductId].Price < product.Price {
			maxProductId = index
		}
		if products[minProductId].Price > product.Price {
			minProductId = index
		}
	}
	return products[maxProductId], products[minProductId]
}

func findWithKeyWord(products []entities.Product, keyword string) (result []entities.Product) {
	for _, product := range products {
		if strings.Contains(product.Name, keyword) {
			result = append(result, product)
		}
	}
	return result
}

func sortProductByPrice(products []entities.Product) []entities.Product {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price > products[j].Price
	})
	return products
}

func Demo5() {
	product := entities.Product{
		Id:       "p01",
		Name:     "name 1",
		Price:    4.5,
		Quantity: 2,
		Status:   true,
		Category: entities.Category{
			Id:   "c1",
			Name: "Category 1",
		},
		Comments: []entities.Comment{
			entities.Comment{
				Id:      "1",
				Title:   "Title 1",
				Content: "Content 1",
			},
			entities.Comment{
				Id:      "2",
				Title:   "Title 2",
				Content: "Content 2",
			},
			entities.Comment{
				Id:      "3",
				Title:   "Title 3",
				Content: "Content 3",
			},
			entities.Comment{
				Id:      "4",
				Title:   "Title 4",
				Content: "Content 4",
			},
		},
	}

	fmt.Println(product)
}
