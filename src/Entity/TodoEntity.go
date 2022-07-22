package Entity

type TaskCore struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Done        bool   `bson:"done" json:"done"`
}

type TodoEntity struct {
	Id       string `bson:"id" json:"id"`
	TaskCore `bson:"inline"`
	SubTasks []TaskCore `bson:"subTasks" json:"subTasks"`
}
