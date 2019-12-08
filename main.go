package main

import (
	"fmt"
	"shop/cmd"
)
11111111
func main(){

	err := cmd.ShopStart()
	if err != nil{
		fmt.Println("start error: ",err)
	}
}
