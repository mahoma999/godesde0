package variables
import(
"fmt"
"time"
"strconv"
) 
var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time
func RestoVariables(){
  Nombre="pedro"
  Estado=true
  Sueldo=12.99
  Fecha= time.Now()
  fmt.Println(Nombre)
  fmt.Println(Estado)
  fmt.Println(Sueldo)
  fmt.Println(Fecha)

}
func ConviertoaText(numero int)(bool,string){
  var texto string
  texto=strconv.Itoa( numero )
  return true,texto

}
