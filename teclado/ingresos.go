package teclado
import(
  "fmt"
  "os"
  "bufio"
  "strconv"
)
var numero1 int
var numero2 int
var leyendo string
var err error
func IngresoNumero(){
  scanner:=bufio.NewScanner(os.Stdin)
  fmt.Println("ingreso numero1:")
  if scanner.Scan(){
    numero1,err= strconv.Atoi(scanner.Text())
    if err!=nil{
      panic("El dato es incorrecto"+err.Error())
    }
  }
  fmt.Println("ingreso numero2:")
  if scanner.Scan(){
    numero2,err= strconv.Atoi(scanner.Text())
    if err!=nil{
      panic("El dato es incorrecto"+err.Error())
    }
  }
  fmt.Println("ingreso leyemda:")
  if scanner.Scan(){
    leyendo=scanner.Text()
  }

  fmt.Println(leyendo,numero1*numero2)
}
