package database

type DatabaseConfig struct {
	Type         string `xml:"type"`
	HostName     string `xml:"host"`
	HostPort     int    `xml:"port"`
	UserName     string `xml:"user"`
	Password     string `xml:"password"`
	DataBaseName string `xml:"databaseName"`
}