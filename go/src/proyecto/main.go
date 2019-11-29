package main

import (
	"fmt"
	//"proyecto/otra"
	//"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	//_"github.com/go-sql-drive/mysql" 
	"errors"
)

func main(){

	/*
	fmt.Println("Hello world")
	fmt.Println()

	c := otra.SumarValores(4,5)
	fmt.Println(c)
*/
	autoX := NewAuto("Ford","Nuevo", "a3")
	autoZ := NewAuto("Chevrolet","Viejo", "b6")
	autoY := NewAuto("lalala","Masomeno", "b4")
	agenciaX := NewAgencia()
	agenciaX.AgregarAuto(autoX)
	agenciaX.AgregarAuto(autoZ)
	agenciaX.AgregarAuto(autoY)
	
	/*autoA := agenciaX.DevolverAuto(1)
	
	fmt.Println(autoA.GetModelo())
	agenciaX.EliminarAuto(1)
	autoB := agenciaX.DevolverAuto(1)
	autoH := agenciaX.DevolverAutoByPatente("b4")
	fmt.Println(autoH.GetModelo())
	
var b := 5
	var a := 10

	var c *int  // el tipo de c es un puntero a entero
	c= &b
	fmt.Println(&b)   // lugar de memoria de b
	fmt.Println(c)   //el valor de c es el lugar de memoria de b
	fmt.Println(&c)  // lugar de memoria de c
// para imprimir dos cosas, la "," es como el "+"
*/
/*
	myString := "hola como andas"
	fmt.Println(myString)
	var a []byte 
	a, _ = json.Marshal(myString)
	fmt.Println(a)


	var otroString string
	_=json.Unmarshal(a, &otroString)
	fmt.Println(otroString)
	autoL := NewAuto("Ford","Nuevo", "a3")

	autoAMarshalled, _ := json.Marshal(autoL)
	fmt.Println("AUTO" , autoAMarshalled)

*/

router := gin.Default()
router.GET("/cars/:patente", agenciaX.GetAutoByPatente)

if err := router.Run(); err !=nil{
	fmt.Println(err)
}


}
/*
func 	AddCar(ctx *gin.Context) {
    body, err := ioutil.ReadAll(ctx.Request.Body)
    if err != nil {
	  fmt.Println("error")
	  return
    }

    var car Car
    if err = json.Unmarshal(body, &car); err != nil {
		fmt.Println("error")
		return
    }

    a.carDatabase.AddCar(car)

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "car added",
    })
}
*/




func (a *Agencia) GetAutoByPatente(ctx *gin.Context){
	patente := ctx.Param("patente")
	
	 auto, err:= a.DevolverAutoByPatente(patente)
	
	 if err !=nil{
	ctx.JSON(http.StatusNotFound, gin.H{
		"error": err.Error(),
	})
	
	return
	
	}
		ctx.JSON(http.StatusOK, gin.H{
	
			"Auto" : auto,
	
	
		})
		
	}








type Agencia struct{
	 autos [] Auto
}


type Auto struct{
	marca string
	modelo string
	patente string
}

func (a *Auto) GetModelo() string{
	return a.modelo
}

func (a *Auto) GetPatente() string{
	return a.patente
}


func (a *Auto) SetModelo(modelo string) {
	a.modelo = modelo
}

func (a *Agencia) setModeloAuto(auto Auto, modelo string){
	for i:= 0 ; i<len(a.autos); i++{
		if a.autos[i]  == auto {
			auto.SetModelo(modelo)
		}
	}
}



func (a *Agencia) DevolverAutoByPatente (patente string) (*Auto, error){
	for _, auto:= range a.autos{
		if auto.patente == patente{
			return &auto, nil
		}
	}
	return nil, errors.New("auto no encontrado")
}	


func (a *Agencia) AgregarAuto( auto Auto){
a.autos = append(a.autos, auto)
}

func NewAuto(marca string, modelo string, patente string) Auto{
	return Auto{
		marca: marca,
		modelo: modelo,
		patente: patente,
	}
}
/*
func (a *Agencia) EliminarAuto(pos int){

	if pos ==0 {
		if len(a.autos)>1{
		a.autos = a.autos[1:len(a.autos)]
	}else{
		var autos [] Auto
		a.autos = autos
	}
	}
	else if pos == len(a.autos)-1{
		a.autos = a.autos[0:len(a.autos)-1]
	}
	else{
	aux := a.autos[0:pos]
	aux2 := a.autos[pos+1:len(a.autos)]
	
	for i :=0; i < len(aux2); i++{
		aux = append(aux, aux2[i])
	}
	a.autos = aux
}
}*/


func (a *Agencia) DevolverAuto (pos int) Auto{
	return a.autos[pos]
}

func NewAgencia() Agencia{
	var autos [] Auto
	return Agencia	{
		autos: autos,
	}
}


