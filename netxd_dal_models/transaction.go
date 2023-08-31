package netxddalmodels

import "go.mongodb.org/mongo-driver/bson/primitive"



type Transaction struct{
	ID primitive.ObjectID`json:"_id" bson:"_id"`
	FromAccount int64`json:"from_acc" bson:"from_acc"`
	ToAccount int64 `json:"to_acc" bson:"to_acc"`
	Amount int64 `json:"amount" bson:"amount"`
	//Balance int64 `json:"balance" bson:"balance"`
}

