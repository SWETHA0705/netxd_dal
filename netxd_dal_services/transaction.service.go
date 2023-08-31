package netxddalservices

import (
	"context"
	"fmt"

	netxddalinterfaces "github.com/SWETHA0705/netxd_dal/netxd_dal_interfaces"
	netxddalmodels "github.com/SWETHA0705/netxd_dal/netxd_dal_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Transaction struct{
	TransactionCollection *mongo.Collection
	CustomerCollection *mongo.Collection
	ctx context.Context
}

func InitialiseTransactionService(transactioncollection *mongo.Collection, customercollection *mongo.Collection, ctx context.Context)(netxddalinterfaces.Itransaction){
	return &Transaction{transactioncollection,customercollection,ctx}
}

func (s* Transaction)Transaction(FromId int64,ToId int64, amount int64 )(string,error){
	
	fmt.Println(FromId)
	_, err2 := s.CustomerCollection.UpdateOne(context.Background(), bson.M{"CustomerId": ToId}, bson.M{"$inc": bson.M{"balance": amount}})
	_, err2 = s.CustomerCollection.UpdateOne(context.Background(), bson.M{"CustomerId": FromId}, bson.M{"$inc": bson.M{"balance": -amount}})
	if err2!= nil{
		return "",err2
	}
	res, _ := s.TransactionCollection.InsertOne(context.Background(), &netxddalmodels.Transaction{
		ID:          primitive.NewObjectID(),
		FromAccount: FromId,
		ToAccount:   ToId,
		Amount:      amount,
	})
	var newUser *netxddalmodels.Transaction
		query := bson.M{"_id": res.InsertedID}

		err3 := s.TransactionCollection.FindOne(s.ctx, query).Decode(&newUser)
		if err3 != nil {
			return "", err3
		}
		return "success", nil
	}
	

	
