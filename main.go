package main

import "fmt"

//henter ekstern modul rsc.io/quote
import "rsc.io/quote"

func  GetHello()string {
	return quote.Hello()

}
func  GetGlass()string {
	return quote.Glass()

}
func GetGo()string {
	return quote.Go()

}
func  GetOpt()string {
	return quote.Opt()

}
func main(){
	fmt.Println(GetHello())
	fmt.Println(GetGlass())
	fmt.Println(GetGo())
	fmt.Println(GetOpt())
}
