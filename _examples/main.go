package main

import ("os"
        "github.com/go-dis/godis")


func main() {
	client := godis.New()
	client.Login("OTMxMTc2MjA4Nzc4MDE0NzQw.YeAnbA.fDM3RWwsPc7IFquodoYkhUpt66c")
}