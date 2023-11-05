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

type Error struct {
	Message string
}

func (e Error) Error() string {
	return e.Message
}

type Res struct {
	Err *Error
}

type GetBalanceRes struct {
	Res
	Balance float32
}
