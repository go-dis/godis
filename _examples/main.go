package main

import ("os"
        "github.com/go-dis/godis")


func main() {
	client := godis.New()
	client.Login("UR_TOKEN")
}