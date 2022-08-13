package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "Sushil", City: "Kathmandu", ZipCode: "111111", DateOfBirth: "2000-01-01", Status: "1"},
		{Id: "2", Name: "Puja", City: "Kathmandu", ZipCode: "00000", DateOfBirth: "2001-01-01", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}
