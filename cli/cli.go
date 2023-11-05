package cli

import (
	"fmt"
	"github.com/GabiHert/t2-fppd/agencyclient"
	"github.com/GabiHert/t2-fppd/machineclient"
	"github.com/manifoldco/promptui"
	"strconv"
)

type Cli struct {
	agency      *agencyclient.Agency
	cashMachine *machineclient.CashMachine
	isAgency    bool
	stop        chan bool
}

func (c *Cli) Run() {
	selectPrompt := promptui.Select{
		Label: "Select your operator",
		Items: []string{"cash-machine", "agency"},
	}

	_, result, _ := selectPrompt.Run()
	c.isAgency = result == "agency"

	var actions []string
	if c.isAgency {
		actions = []string{"Auth", "Withdraw", "Deposit",
			"Create account", "Delete account", "Get balance", "Cancel", "Abort"}
	} else {
		actions = []string{"Auth", "Withdraw", "Deposit",
			"Get balance", "Cancel", "Abort"}
	}

	selectPrompt = promptui.Select{Label: "Select your action",
		Items: actions}

	_, result, _ = selectPrompt.Run()

	switch result {
	case "Auth":
		c.auth()
	case "Withdraw":
		c.withdraw()
	case "Deposit":
		c.deposit()
	case "Create account":
		c.createAccount()
	case "Delete account":
		c.deleteAccount()
	case "Get balance":
		c.getBalance()
	case "Cancel":
		return
	case "Abort":
		c.stop <- true
	}

}

func getNameAndPassword() (name string, password int) {
	prompt := promptui.Prompt{
		Label: "Name",
	}

	name, _ = prompt.Run()

	prompt = promptui.Prompt{Label: "Password"}

	pswS, _ := prompt.Run()

	password, _ = strconv.Atoi(pswS)
	return
}

func (c *Cli) auth() {
	name, psw := getNameAndPassword()
	if c.isAgency {
		err := c.agency.Auth(name, psw)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		err := c.cashMachine.Auth(name, psw)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	fmt.Println("Auth successful")
}
func (c *Cli) withdraw() {
	name, psw := getNameAndPassword()
	prompt := promptui.Prompt{Label: "Amount"}

	amountS, _ := prompt.Run()

	amount, _ := strconv.ParseFloat(amountS, 32)
	if c.isAgency {
		err := c.agency.Withdraw(name, psw, float32(amount))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		err := c.cashMachine.Withdraw(name, psw, float32(amount))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	fmt.Println("Withdraw successful")

}
func (c *Cli) deposit() {
	name, psw := getNameAndPassword()
	prompt := promptui.Prompt{Label: "Amount"}

	amountS, _ := prompt.Run()

	amount, _ := strconv.ParseFloat(amountS, 32)
	if c.isAgency {
		err := c.agency.Deposit(name, psw, float32(amount))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		err := c.cashMachine.Deposit(name, psw, float32(amount))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	fmt.Println("Deposit successful")
}
func (c *Cli) createAccount() {
	name, psw := getNameAndPassword()
	err := c.agency.CreateAccount(name, psw)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Account created successfully")

}
func (c *Cli) deleteAccount() {
	name, psw := getNameAndPassword()
	err := c.agency.DeleteAccount(name, psw)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Account deleted successfully")

}

func (c *Cli) getBalance() {
	name, psw := getNameAndPassword()
	var (
		balance float32
		err     error
	)
	if c.isAgency {
		balance, err = c.agency.GetBalance(name, psw)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		balance, err = c.cashMachine.GetBalance(name, psw)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	fmt.Printf("Balance: %2.f\n", balance)
}

func NewCli(
	agency *agencyclient.Agency,
	machine *machineclient.CashMachine,
	stop chan bool) *Cli {
	return &Cli{
		stop:        stop,
		agency:      agency,
		cashMachine: machine,
	}
}
