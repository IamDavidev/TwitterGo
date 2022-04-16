package main

import (
	"fmt"
	"os"
)

func main() {
	var envMongo = os.Getenv("MONGO_DB_URI")
	fmt.Println("Go Twitter Clone ", envMongo)

}
