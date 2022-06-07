package model

import (
	"encoding/json"
	"os"
)

type Customer struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	City  string `json:"city"`
}

// not required at present
type CustomerOps interface {
	AddCustomer(Customer) Customer
	GetCustomer(int) (Customer, bool)
	UpdateCustomer(Customer) Customer
	DeleteCustomer(int)
	GetAll() []Customer
	GetByCity(string) []Customer
}

// implements methods in CustomerOps
type CustomerManager struct {
	customers []Customer
}

func generateId(customers []Customer) (newId int) {
	if len(customers) == 0 {
		newId = 1
		return
	}
	for _, c := range customers {
		if c.Id > newId {
			newId = c.Id
		}
	}
	newId++
	return
}

func (this *CustomerManager) AddCustomer(customer Customer) Customer {
	customer.Id = generateId(this.customers)
	this.customers = append(this.customers, customer)
	this.SaveToFile() // same will be called in update and delete operations
	return customer
}

func (this *CustomerManager) GetCustomer(id int) (customer Customer, found bool) {
	for _, c := range this.customers {
		if id == c.Id {
			customer = c
			found = true
			return
		}
	}
	found = false
	return
}

func (this *CustomerManager) GetAll() []Customer {
	return this.customers
}

func (this *CustomerManager) GetByCity(city string) (customers []Customer) {
	for _, c := range this.customers {
		if city == c.City {
			customers = append(customers, c)
		}
	}
	return
}

func (this *CustomerManager) SaveToFile() {
	customersJson, _ := json.Marshal(this.customers)
	os.WriteFile("customers.json", customersJson, 0644)
}

func (this *CustomerManager) LoadFromFile() {
	bytes, _ := os.ReadFile("customers.json")
	json.Unmarshal(bytes, &this.customers)
}
