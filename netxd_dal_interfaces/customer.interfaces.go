package netxddalinterfaces

import netxddalmodels "github.com/SWETHA0705/netxd_dal/netxd_dal_models"



type ICustomer interface{
	CreateCustomer(customer *netxddalmodels.Customer)(*netxddalmodels.Customer,error)
}