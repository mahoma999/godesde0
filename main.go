package main
import(
  "fmt"
  //"runtime"
  //"github.com/mahoma999/godesde0/variables"
  "github.com/mahoma999/godesde0/ejercicios"
  
)
func main(){
  /*estado,texto:=variables.ConviertoaText(9)
  fmt.Println(estado)
  fmt.Println(texto)
  if os:=runtime.GOOS;os=="linux" || os=="OS X."{
    fmt.Println("no es windows,es",os)

  }else{
    fmt.Println("Es windpws")
  }
  */
  numero,texto:=ejercicios.ConvaNum("ffff")
  fmt.Printf("El numero es %d %s",numero,texto)

}
