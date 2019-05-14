package entities

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type Mobile struct {
	Id       bson.ObjectId `bson:"_id" json:"id"`
	Name     string        `bson:"name" json:"name"`
	Price    float64       `bson:"price" json:"price"`
	Quantity int           `bson:"quantity" json:"quantity"`
	Status   bool          `bson:"status" json:"status"`
}

func (mobile Mobile) ToString() string {
	return fmt.Sprintf("id: %s\nname: %s\nprice: %0.1f\nquantity: %d\nstatus: %t",
		mobile.Id.Hex(), mobile.Name, mobile.Price, mobile.Quantity, mobile.Status)
}
