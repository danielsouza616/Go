package main
import (
    "fmt"
    "os"
//    "strconv"
)

func main(){
    if len(os.Args) <= 1 {
         fmt.Println("Sem argumentos recebidos!")
         os.Exit(1)
    }
fmt.Println("argumentos recebidos!")
fmt.Println(os.Args[0])
fmt.Println(os.Args[1])
}
