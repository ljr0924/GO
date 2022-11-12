package server

type Msg struct {
	User *User
	Msg  string
}

func NewMsg(user *User, msg string) *Msg {
	return &Msg{
		User: user,
		Msg:  msg,
	}
}
