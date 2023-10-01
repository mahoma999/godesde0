package ejercicios
import "strconv"
func ConvaNum(cadena string)(int,string){
  entero,error:=strconv.Atoi(cadena)
  if error!=nil{
    return 0,"hubo un error"+error.Error()
  }
  if entero > 100{
    return entero,"es mayor"
  }else{
    return entero,"es menor que cien"
  }
}
