package netxddalinterfaces

import netxddal "github.com/SWETHA0705/netxd_dal/netxd_dal_models"


type ICustomer interface{
	CreateCustomer(customer *netxddal.Customer)(*netxddal.Customer,error)
}