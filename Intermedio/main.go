package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int//Explicito
	x = 8
	y := 7//Implicito
	fmt.Println(x)
	fmt.Println(y)
	myValue, err := strconv.ParseInt("NaN",0,64)
	if err != nil{//Catch de errores explicitos
		fmt.Printf("%v\n",err)
	}else{
		fmt.Println(myValue)
	}
	m := make(map[string]int)//Map equivalente a los diccionarios
	m["Alan"]=10
	fmt.Println(m["Alan"])
	s := []int{1,2,3}//Slice equivalente a las listas
	s = append(s, 51)//agregar un valor
	for index,value := range(s){
		fmt.Printf("%d:%d\n",index,value)
	}
	//c := make(chan int)//Se crea un canal
	//go doSomething(c)//Se crea una go routine
	//<-c//Se comunica el canal

	g:= 25
	fmt.Println(g)
	h:=&g//referencia
	fmt.Println(h)//Imprime la direccion de referencia
	fmt.Println(*h)//Imprime el valor al que apunta  

}

func doSomething(c chan int){
	time.Sleep(3*time.Second)
	fmt.Println("Done")
	c<-1//Devuelve un int al canal

}