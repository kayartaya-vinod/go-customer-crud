package main

import (
	"fmt"
	"strconv"

	"github.com/kayartaya-vinod/go-customer-crud/model"
)

var custMgr model.CustomerManager

func main() {
	custMgr = model.CustomerManager{}
	custMgr.LoadFromFile()

	for {
		fmt.Println("Customer dashboard")
		fmt.Println("==================")
		fmt.Println("1. Add new customer")
		fmt.Println("2. Search by id")
		fmt.Println("3. Display all customers")
		fmt.Println("4. Search customers by city")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		var choice string
		fmt.Scanln(&choice)

		ans, _ := strconv.Atoi(choice)
		if ans == 5 {
			break
		}

		switch ans {
		case 1:
			addNewCustomer()
		case 2:
			searchCustomerById()
		case 3:
			displayAllCustomers()
		case 4:
			searchCustomerByCity()
		}
	}
}

func addNewCustomer() {
	var customer model.Customer
	fmt.Print("Enter customer details")
	fmt.Println("------------------------------------------")
	fmt.Print("Name         : ")
	fmt.Scanln(&customer.Name)
	fmt.Print("Email address: ")
	fmt.Scanln(&customer.Email)
	fmt.Print("Phone number : ")
	fmt.Scanln(&customer.Phone)
	fmt.Print("City         : ")
	fmt.Scanln(&customer.City)

	customer = custMgr.AddCustomer(customer)
	fmt.Println("Customer details added with id", customer.Id)

}
func searchCustomerById() {
	var input string
	fmt.Print("Enter id of the customer to search: ")
	fmt.Scanln(&input)
	id, _ := strconv.Atoi(input)
	customer, found := custMgr.GetCustomer(id)
	if found {
		fmt.Println("Search result: ")
		fmt.Println("------------------------------------------")
		fmt.Printf("ID            : %v\n", customer.Id)
		fmt.Printf("Name          : %v\n", customer.Name)
		fmt.Printf("Email         : %v\n", customer.Email)
		fmt.Printf("Phone         : %v\n", customer.Phone)
		fmt.Printf("City          : %v\n", customer.City)
		fmt.Println("------------------------------------------")
	} else {
		fmt.Printf("No customer data found for id %v\n", id)
	}
}
func displayAllCustomers() {
	customers := custMgr.GetAll()
	displayAsList(customers)
}

func searchCustomerByCity() {
	var city string
	fmt.Print("Enter city to search customers: ")
	fmt.Scanln(&city)
	customers := custMgr.GetByCity(city)
	if len(customers) == 0 {
		fmt.Println("No customers found in city:", city)
	} else {
		displayAsList(customers)
	}
}

func displayAsList(customers []model.Customer) {
	fmt.Printf("%5s %-20s %-30s %-20s %-20s\n", "ID", "Name", "Email", "Phone", "City")
	fmt.Println("-----------------------------------------------------------------------------------------")
	for _, c := range customers {
		fmt.Printf("%5d %-20s %-30s %-20s %-20s\n", c.Id, c.Name, c.Email, c.Phone, c.City)
	}
	fmt.Println("-----------------------------------------------------------------------------------------")

}
