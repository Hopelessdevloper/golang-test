package main
import ("fmt")

type rating float64
type Game struct{
    name string     
}

type gamelist make(map[Game]rating)

//func addGame(addlist gamelist)
func main(){
   CoD := Game {
       name: "Call Of Duty",
   }
   
   activision := gamelist
    
   activision.gamelist[CoD.Game] = 5.5
 //  activision.gamelist["CoD:MW"] = 4.5
   
   fmt.Println(activision.gamelist)
}