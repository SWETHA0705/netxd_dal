package netxddalinterfaces

type Itransaction interface{
	Transaction(int64,int64,int64)(string,error)
}