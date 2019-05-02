package main

import (

	mysql "lib/mysql/db"
)

func main() {

	mongo.Get() //calling method of package "lib/mongodb/db"
	mysql.Get() //calling method of package "lib/mysql/db"
}
