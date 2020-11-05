package db

var Ins DataBase

func NewDataBase() DataBase {
	d := new(DataBaseImpl)
	return d
}

func GetDataBase() DataBase {
	return Ins
}

func InitDB() {
	Ins = NewDataBase()
	Ins.Open()
}
