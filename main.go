package main

import (
	"go-train/database"
	"go-train/middleware"
	"go-train/routes"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 初始化數據庫
	database.ConnectDB()

	// 初始化 Gin
	r := gin.Default()
	r.Use(middleware.Logger())
	routes.SetupRoutes(r)

	// 啟動伺服器
	r.Run(":8080")
}

//import "fmt"
//
//func main() {
//	fmt.Println("Hello, World!")
//	var a int32 = 10
//	fmt.Println("Value of a:", a)
//	fmt.Println("capacity of a:", cap([]int{1, 2, 3}))
//	var b int64 = 20
//	fmt.Println("Value of b:", b)
//	var c float64 = 30.5
//	fmt.Println("Value of c:", c)
//	var d string = "Hello, Go!"
//	fmt.Println("Value of d:", d)
//	var e bool = true
//	fmt.Println("Value of e:", e)
//	var f []int = []int{1, 2, 3, 4, 5}
//	fmt.Println("Value of f:", f)
//	var g map[string]int = map[string]int{"one": 1, "two": 2}
//	fmt.Println("Value of g:", g)
//	var h struct {
//		Name string
//		Age  int
//	}
//	fmt.Println("Value of h:", h)
//
//	var myMap map[string]uint16 = make(map[string]uint16) // Initialize a map with string keys and uint16 values
//	myMap["One"] = 1
//	myMap["Two"] = 2
//	// fmt.Println("Value of myMap:", myMap)
//
//	for key, value := range myMap {
//		fmt.Printf("Key: %s, Value: %d\n", key, value)
//	}
//
//}
