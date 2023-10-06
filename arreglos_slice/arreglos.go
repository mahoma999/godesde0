package arreglosslice
import "fmt"
var tabla [10]int = [10]int{1,3}
func MuestrArreglo(){
  tabla[6]=6
  tabla[9]=9
  fmt.Println(tabla[6])
  fmt.Println(tabla[9])
  fmt.Println(tabla)
  for i:=0;i<len(tabla);i++{
    fmt.Println(tabla[i])
  }
}
