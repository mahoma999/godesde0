package ejercicio_interface
import (
  "fmt"
  "github.com/mahoma999/godesde0/interfaces"
)
func HumanoRespirando(humano interfaces.Humano){
  humano.Respirar()
  fmt.Printf("soy un %s estoy respirando\n",humano.Sexo())

}
