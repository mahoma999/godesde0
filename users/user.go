package users
import(
  "fmt"
  "time"
  "github.com/mahoma999/godesde0/modelos"
)
func AltaUsuario(){
  u:=new(modelos.User)
  u.AddUser(9,"oscar",time.Now(),true)
  fmt.Println(u)

}
