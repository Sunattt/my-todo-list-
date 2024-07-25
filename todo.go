package todo_list

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UserLists struct {
	Id     int
	UserId int
	ListId int
}

type TodoItems struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItems struct {
	Id     int
	ListId int
	ItemId int
}
