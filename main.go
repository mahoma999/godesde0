package main
import(
  "fmt"
  //"runtime"
  //"github.com/mahoma999/godesde0/variables"
  //"github.com/mahoma999/godesde0/ejercicios"
  //"github.com/mahoma999/godesde0/teclado"
  //"github.com/mahoma999/godesde0/files"
  //"github.com/mahoma999/godesde0/funciones"
  //"github.com/mahoma999/godesde0/arreglos_slice"
  //"github.com/mahoma999/godesde0/mapas"
  //"github.com/mahoma999/godesde0/users"
  "github.com/mahoma999/godesde0/modelos"
  e "github.com/mahoma999/godesde0/ejercicio_interface"
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
  numero,texto:=ejercicios.ConvaNum("ffff")
  fmt.Printf("El numero es %d %s",numero,texto)
  */
  //teclado.IngresoNumero()
 
  //fmt.Println( ejercicios.Multiplicar())
  fmt.Println("interfaces")
  //files.GrabarTabla()
  //files.SumaTabla()
  //files.LeerArchivo()
  //funciones.Calculos()
  //funciones.LlamarTabla()
  //arreglosslice.MuestrArreglo()
  //arreglosslice.MuestroSlice()
  //arreglosslice.Capacidad()
  //mapas.MostrarMapas()
  //users.AltaUsuario()
  var oscar* modelos.Hombre
  oscar =new(modelos.Hombre)
  e.HumanoRespirando(oscar)
  var veronica* modelos.Mujer
  veronica= new(modelos.Mujer)
  e.HumanoRespirando(veronica)

}
