package interfaces

type Adm interface {
	InitSession() (string, error)
	Auth(name string, psw int) error
	CreateAccount(name string, psw int, token string) error
	DeleteAccount(name string, psw int, token string) error
	GetBalance(name string, psw int, token string) (float32, error)
	Deposit(name string, psw int, amount float32, token string) error
	Withdraw(name string, psw int, amount float32, token string) error
}
