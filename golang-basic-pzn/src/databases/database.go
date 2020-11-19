package databases

var connect string

func init() {
	connect = "MongoDB"
}

func GetDatabase() string {
	return connect
}