package application

type Client struct {
	ID       string
	Email    string
	CPF      string
	Document []byte
	Address  Address
}
