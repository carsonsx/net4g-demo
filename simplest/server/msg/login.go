package msg

type UserLogin struct {
	Username string
	Password string
}

type UserLoginReply struct {
	Code int
	Msg string
}

type UserOnline struct {
	UseId    int
}

type UserOffline struct {
	UseId    int
}
