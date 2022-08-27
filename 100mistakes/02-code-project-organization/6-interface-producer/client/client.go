package client

import "github.com/phial3/100mistakes/02-code-project-organization/6-interface-producer/store"

type customersGetter interface {
	GetAllCustomers() ([]store.Customer, error)
}
