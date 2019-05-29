/*
	
	ALUMNO: Ernesto Z. Chura Flores
        ALUMNO: Cristian Omar Tacora Claros
	CURSO: Paralelismo y Sistemas Distribuidos
	OJO: La finalizacion de entrada de datos es con '0' (cero)
	
*/
package main

import(

	"fmt"
	"sync"
)

var temp sync.WaitGroup

func mostrar_array(ar[] int){

	for i := 0; i < len(ar); i++{

		fmt.Printf("%v\t",ar[i])
	}
	fmt.Println("\n---------------------------")
}

func Burbuja(ListaDesordenada []int){

 	var auxiliar int
	for i := 0; i < len(ListaDesordenada); i++ {
	  	for j := 0; j < len(ListaDesordenada); j++ {
	   		if ListaDesordenada[i] > ListaDesordenada[j] {
	    		auxiliar = ListaDesordenada[i]
	    		ListaDesordenada[i] = ListaDesordenada[j]
	    		ListaDesordenada[j] = auxiliar
	   		}
	  	}
	}
	mostrar_array(ListaDesordenada)
	temp.Done()
}

func rutinas(arreglo []int) {
	tam := len(arreglo)/4
	go Burbuja(arreglo[:tam])
	go Burbuja(arreglo[tam:tam*2])
	go Burbuja(arreglo[tam*2:tam*3])
	go Burbuja(arreglo[tam*3:])
	Burbuja(arreglo)

}

func main(){

	var numero int
	i := 0
	
	arreglo:= [50] int{}	
	// --------------- INSERTAMOS NUMEROS ---------------
	for{
		fmt.Scanf("%d\n",&numero)
		if(numero == 0){
			break
		}else{
			arreglo[i] = numero
			i++
		}
	}
	//----------------- MOSTRAR ------------------
	fmt.Printf("\n----------------- MATRIZ DESORDENADA -----------------\n\n")
	for j := 0; j < i; j++{

		fmt.Printf("%v\t",arreglo[j])
	}
	fmt.Printf("\n")
	temp.Add(5)
	slice := arreglo[:i]
	fmt.Printf("\n----------------- ORDENANDO MATRIZ -----------------\n\n")
	rutinas(slice)
	temp.Wait()
}
