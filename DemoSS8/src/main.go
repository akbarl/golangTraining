package main

import (
	"config"
	"entities"
	"fmt"
	"models"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	Demo10()
}

func Demo1() {
	db, err := config.GetConnect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db:         db,
			Collection: "product",
		}

		products, err2 := productModel.FindAll()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println("**********************")
				fmt.Println(product.ToString())
			}
		}
	}
}

func Demo2() {
	db, err := config.GetConnect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db:         db,
			Collection: "product",
		}

		products, err2 := productModel.Search1(5)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println("**********************")
				fmt.Println(product.ToString())
			}
		}
	}
}

func Demo3() {
	db, err := config.GetConnect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db:         db,
			Collection: "product",
		}

		products, err2 := productModel.Search2(0, 5)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println("**********************")
				fmt.Println(product.ToString())
			}
		}
	}
}

func Demo4() {
	db, err := config.GetConnect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db:         db,
			Collection: "product",
		}

		products, err2 := productModel.Like1("lap")
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println("**********************")
				fmt.Println(product.ToString())
			}
		}
	}
}

func Demo5() {
	db, err := config.GetConnect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db:         db,
			Collection: "product",
		}

		products, err2 := productModel.Like2("2")
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println("**********************")
				fmt.Println(product.ToString())
			}
		}
	}
}

func Demo6() {
	db, err := config.GetConnect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db:         db,
			Collection: "product",
		}

		products, err2 := productModel.SortDESC()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println("**********************")
				fmt.Println(product.ToString())
			}
		}
	}
}

func Demo7() {
	db, err := config.GetConnect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db:         db,
			Collection: "product",
		}

		product, err2 := productModel.FindById("5cbebcf78f9b4012422d32ad")
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println("**********************")
			fmt.Println(product.ToString())
		}
	}
}

func Demo8() {
	product := entities.Product{
		Id:       bson.NewObjectId(),
		Name:     "TV300",
		Price:    555,
		Quantity: 33,
		Status:   true,
	}

	db, err := config.GetConnect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db:         db,
			Collection: "product",
		}
		err2 := productModel.Create(&product)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println(product.ToString())
		}
	}
}

func Demo9() {
	db, err := config.GetConnect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db:         db,
			Collection: "product",
		}
		err2 := productModel.Delete("5cbebcf78f9b4012422d32ad")
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println("Deleted")
		}
	}
}

func Demo10() {
	db, err := config.GetConnect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db:         db,
			Collection: "product",
		}
		product, err2 := productModel.FindById("5cbebcf78f9b4012422d32ac")
		if err2 != nil {
			fmt.Println(err2)
		} else {
			product.Name = "edited"

			err3 := productModel.Update(&product)
			if err3 != nil {
				fmt.Println(err3)
			} else {
				fmt.Println("Updated")
			}
		}
	}
}
