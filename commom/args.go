package commom

type Req struct {
	Name  string
	Psw   int
	Token string
}

type AuthReq struct {
	Name string
	Psw  int
}

type OperationReq struct {
	Req
	Amount float32
}

type GetBalanceRes struct {
	Balance float32
}
