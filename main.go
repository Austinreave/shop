package main

import (
	"fmt"
	"shop/cmd"
)
func main(){

	err := cmd.ShopStart()
	if err != nil{
		fmt.Println("start error: ",err)
	}
}
