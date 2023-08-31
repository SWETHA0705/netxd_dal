package netxddalmodels


type Customer struct{
	CustomerId int `json:"CustomerId" bson:"CustomerId"`
	FirstName string `json:"FirstName" bson:"FirstName"`
	LastName string `json:"LastName" bson:"LastName"`
	BankId  int `json:"BankId" bson:"BankId"`
	Balance int `json:"Balance" bson:"Balance"`
	CreatedAt string `json:"CreatedAt" bson:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt bson:"UpdatedAt"`
	IsActive string `json:"IsActive" bson:"IsActive"`

}