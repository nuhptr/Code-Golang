package database

var connect string

func init() {
	connect = "MongoDB"
}

func GetDatabase() string {
	return connect
}