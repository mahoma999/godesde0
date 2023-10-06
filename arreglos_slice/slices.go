package arreglosslice
import "fmt"
var tablaS[] int=[]int{2,5,4}
var arreglo[10] int=[10]int{1,2,3,4,5,6,7,8,9,10}
func MuestroSlice(){
  fmt.Println(tablaS)
  numero:=arreglo[3:]
  numero1:=arreglo[:5]
  numero2:=arreglo[6:7]
  fmt.Println(numero)
  fmt.Println(numero1)
  fmt.Println(numero2)

}
func Capacidad(){
  elementos:=make([]int,5,20)
  fmt.Printf("tamaño %d capacidad %d",len(elementos),cap(elementos))

  numeros:=make([]int,0,0)
  for i:=0;i<100;i++{
    numeros = append(numeros,i)
  }
  fmt.Printf("\ntamaño %d capacidad %d",len(numeros),cap(numeros))
}
