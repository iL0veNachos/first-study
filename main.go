package main

import (
	"fmt"
	"study/feature"
	"study/feature2"
	simpleconnection "study/feature_postgres/simple_connection"
)

func main() {

	fmt.Println("Hello from main")
	feature.Feature()
	feature2.Feature2()
	fmt.Println("okey")
	simpleconnection.CheckConnect()
}
