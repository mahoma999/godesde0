package funciones
import "fmt"
func Calculos(){
  var numero1,numero2 int
  numero2=9
  numero1=9
  sumna:=func(numero1 int,numero2 int)int{
    return numero2+numero1
  }
  fmt.Printf("la suma es :%d",sumna(numero2,numero1))
}
