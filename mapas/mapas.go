package mapas
import "fmt"
func MostrarMapas(){
  paises:=make(map[string]string)
  //fmt.Println(paises)
  paises["mexico"]="df"
  paises["argentina"]="buenos aires"
  fmt.Println(paises)
  fmt.Println(paises["mexico"])
  campeonato:=map[string]int{
    "barcelona":5,
    "chivas":9,
    "mexico":9,
  }
  fmt.Println(campeonato)
  for equipo,puntaje:= range campeonato{
    fmt.Printf("el equipo es %s con su puntaje %d\n",equipo,puntaje)
  }
  delete(campeonato,"chivas")
  fmt.Println(campeonato)

  puntaje,existe:=campeonato["munich"]
  fmt.Printf("El puntaje %d ,equpi existe %t\n",puntaje,existe)

  puntaj,exist:=campeonato["mexico"]
  fmt.Printf("El puntaje %d ,equpi existe %t\n",puntaj,exist)


}
