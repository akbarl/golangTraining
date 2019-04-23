package models

import (
	"entities"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductModel struct {
	Db         *mgo.Database
	Collection string
}

func (productModel ProductModel) FindAll() (products []entities.Product, err error) {
	err = productModel.Db.C(productModel.Collection).Find(bson.M{}).All(&products)
	return products, err
}

func (productModel ProductModel) Search1(min float64) (products []entities.Product, err error) {
	err = productModel.Db.C(productModel.Collection).Find(bson.M{
		"price": bson.M{"$gte": min},
	}).All(&products)
	return products, err
}

func (productModel ProductModel) Search2(min, max float64) (products []entities.Product, err error) {
	err = productModel.Db.C(productModel.Collection).Find(bson.M{
		"$and": []bson.M{
			bson.M{"price": bson.M{"$gte": min}},
			bson.M{"price": bson.M{"$lte": max}},
		},
	}).All(&products)
	return products, err
}

//startWith
func (productModel ProductModel) Like1(keyword string) (products []entities.Product, err error) {
	err = productModel.Db.C(productModel.Collection).Find(bson.M{
		"name": bson.RegEx{Pattern: "^" + keyword, Options: "i"},
	}).All(&products)
	return products, err
}

//endWith
func (productModel ProductModel) Like2(keyword string) (products []entities.Product, err error) {
	err = productModel.Db.C(productModel.Collection).Find(bson.M{
		"name": bson.RegEx{Pattern: keyword + "$", Options: "i"},
	}).All(&products)
	return products, err
}

//sortASC
func (productModel ProductModel) SortASC() (products []entities.Product, err error) {
	err = productModel.Db.C(productModel.Collection).Find(bson.M{}).Sort("price").All(&products)
	return products, err
}

//sortDESC
func (productModel ProductModel) SortDESC() (products []entities.Product, err error) {
	err = productModel.Db.C(productModel.Collection).Find(bson.M{}).Sort("-price").All(&products)
	return products, err
}

//limit
func (productModel ProductModel) Limit(n int) (products []entities.Product, err error) {
	err = productModel.Db.C(productModel.Collection).Find(bson.M{}).Limit(n).All(&products)
	return products, err
}

//findById
func (productModel ProductModel) FindById(id string) (product entities.Product, err error) {
	err = productModel.Db.C(productModel.Collection).FindId(bson.ObjectIdHex(id)).One(&product)
	return product, err
}

func (productModel ProductModel) Create(product *entities.Product) error {
	err := productModel.Db.C(productModel.Collection).Insert(&product)
	return err
}

func (productModel ProductModel) Update(product *entities.Product) error {
	err := productModel.Db.C(productModel.Collection).UpdateId(product.Id, &product)
	return err
}

func (productModel ProductModel) Delete(id string) error {
	err := productModel.Db.C(productModel.Collection).RemoveId(bson.ObjectIdHex(id))
	return err
}
