package users

type Book struct {
	Title string
	Cost  float64
}

type password struct {
	PublicKey  int
	PrivateKey string
}

type User struct {
	Name    string
	Age     int
	Address string
	Book
	password
}
