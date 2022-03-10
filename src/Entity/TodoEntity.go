package Entity

type TaskCore struct {
	Id          string `bson:"id" json:"id"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Done        bool   `bson:"done" json:"done"`
}

type TodoStructure struct {
	TaskCore `bson:"inline"`
	SubTasks []TaskCore `bson:"subTasks" json:"subTasks"`
}
