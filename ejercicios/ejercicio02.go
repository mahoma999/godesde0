package ejercicios
import(
  "fmt"
  "os"
  "bufio"
  "strconv"
)
func Multiplicar(){
  scanner:=bufio.NewScanner(os.Stdin)
  var numero int
  var err error
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
    fmt.Printf("%d X %d = %d\n",i,numero,i*numero)
  }

}
