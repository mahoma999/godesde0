package files
import(
  "github.com/mahoma999/godesde0/ejercicios"
  "os"
  "bufio"
  "io/ioutil"
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
func LeerArchivo(){
  archivo,err:=ioutil.ReadFile("./files/txt/tabla.txt")
  if err!=nil{
    fmt.Println("error al leer el archivo"+err.Error())
  }
  fmt.Println(string(archivo))
  
}
func LeerArchivo2(){
  archivo,err:=os.Open("./files/txt/tabla.txt")
  if err!=nil{
    fmt.Println("error al leer el archivo"+err.Error())
  }
  scanner:=bufio.NewScanner(archivo)
  for scanner.Scan(){
    linea:=scanner.Text()
    fmt.Println("> ",linea)
  }
  archivo.Close()
}
