package database

type DatabaseConfig struct {
	Type         string `xml:"type"`
	HostName     string `xml:"host"`
	HostPort     int    `xml:"port"`
	UserName     string `xml:"user"`
	Password     string `xml:"password"`
	DataBaseName string `xml:"databaseName"`
}

type IDataBase interface {
	Insert(object interface{}) (interface{}, error)
	Update(object interface{}) error
	Delete(Object interface{}) error
	Find(object interface{}) (interface{}, error)
}

type DataBase struct {
	IDataBase
}

func (this *DataBase) Insert(object interface{}) (interface{}, error) {
	return nil, nil
}

func (this *DataBase) Update(object interface{}) error {
	return nil
}

func (this *DataBase) Delete(object interface{}) error {
	return nil
}

func (this *DataBase) Find(object interface{}) (interface{}, error)  {
	return nil, nil
}
