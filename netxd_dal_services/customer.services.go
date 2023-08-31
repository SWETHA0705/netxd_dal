package netxddalservices

import (
	"context"
	netxddalinterfaces " github.com/SWETHA0705/netxd_dal/netxd_dal_interfaces"
	netxddal "github.com/SWETHA0705/netxd_dal/netxd_dal_models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//netxddal "dal/netxd_dal/netxd_dal_models"
)

type CustomerService struct{
	CustomerCollection *mongo.Collection
	ctx context.Context
}

func InitialiseCustomerService(collection *mongo.Collection,ctx context.Context)netxddalinterfaces.ICustomer{
	return &CustomerService{ collection,ctx}
}

func(c * CustomerService) CreateCustomer(user *netxddal.Customer)(*netxddal.Customer,error){
	
	user.CreatedAt = time.Now().Format(time.RFC3339)
    user.UpdatedAt = user.CreatedAt
    user.IsActive = "true"
	res, err := c.CustomerCollection.InsertOne(c.ctx, user)
    if err != nil {
        return nil,err
    }
    
	var newUser *netxddal.Customer
	query := bson.M{"_id": res.InsertedID}

	err = c.CustomerCollection.FindOne(c.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

