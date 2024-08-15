package chat

type Chat struct {
	Id int32
	Name string
}

type ChatInfo struct {
	Name string
	UserIds []int32
}