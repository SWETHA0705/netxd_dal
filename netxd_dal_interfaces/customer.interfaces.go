package netxddalinterfaces

import netxddal "netxd_project/dal/netxd_dal/netxd_dal_models"


type ICustomer interface{
	CreateCustomer(customer *netxddal.Customer)(*netxddal.Customer,error)
}