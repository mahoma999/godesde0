package ejercicios
import(
  "fmt"
  "os"
  "bufio"
  "strconv"
)
func Multiplicar()string{
  scanner:=bufio.NewScanner(os.Stdin)
  var numero int
  var err error
  var text string
  for{
    fmt.Println("Da un numero")
    if scanner.Scan(){
      numero,err=strconv.Atoi(scanner.Text())
      if(err!=nil){
        continue;
      }else{
        break
      }
    }
  }
  for i:=1;i<=10;i++{
    text+=fmt.Sprintf("%d X %d = %d\n",i,numero,i*numero)
  }
    return text
}
