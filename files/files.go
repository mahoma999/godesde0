package files
import(
  "github.com/mahoma999/godesde0/ejercicios"
  "os"
  //"bufio"
  //"ioutil"
  "fmt"
)
var filename string="./files/txt/tabla.txt"
func GrabarTabla(){
  var texto string=ejercicios.Multiplicar()
  archivo,err:=os.Create(filename)
  if err!=nil{
    fmt.Println("error al crear el archico",err.Error())
    return
  }
  fmt.Fprintln(archivo,texto)
  archivo.Close()
}

func SumaTabla(){
  var texto string=ejercicios.Multiplicar()
  if Append(filename,texto) == false{
    fmt.Println("error l concatenar el archvio")
  }

}
func Append(file string,texto string)bool{
  arch,err:=os.OpenFile(filename,os.O_WRONLY|os.O_APPEND,0644)

  if err!=nil{
    fmt.Println("hubo un error",err.Error())
    return false
  }
  _,err=arch.WriteString(texto)
  if err!=nil{
    fmt.Println("hubo un error write",err.Error())
    return false
  }
  arch.Close()
  return true
}
