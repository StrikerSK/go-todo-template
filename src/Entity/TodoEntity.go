package Entity

type TaskCore struct {
	Id          string `bson:"id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Done        bool   `bson:"done"`
}

type SubTaskStruct struct {
	Id          string `bson:"id" json:"id"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Done        bool   `bson:"done" json:"done"`
}

type TodoStructure struct {
	Id          string          `bson:"id" json:"id"`
	Name        string          `bson:"name" json:"name"`
	Description string          `bson:"description" json:"description"`
	SubTasks    []SubTaskStruct `bson:"subTasks" json:"subTasks"`
	Done        bool            `bson:"done" json:"done"`
}
