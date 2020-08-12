package main

import (
	"fmt"
    "math/rand"
    "time"
)

func random(min int, max int) int {
    return rand.Intn(max-min) + min
}

func Get_ItemAI(item string) string{
    var AI_Answer string 
	rand.Seed(time.Now().UnixNano())
	randomNum := random(1, 4)
	if  randomNum==1{ 
		AI_Answer="Rock"  
		
	
	}
	if randomNum==2	{
		AI_Answer="Paper" 
		
	}
	if randomNum==3	{
		AI_Answer="Scissors" 
	 
	}
	fmt.Print("Your item = ",item)
	fmt.Print(" VS ")
	fmt.Print("Item of AI = ",AI_Answer)
	fmt.Print("\n")
	if item==AI_Answer{
	fmt.Println("X| DRAW! |X")
	}
	if item=="Rock"&&AI_Answer=="Paper"{
		fmt.Println("YOU LOSE!!!!!!")
	}
	if item=="Paper"&&AI_Answer=="Rock"{
		fmt.Println("YOU WIN!!!!!!")
	}
	if item=="Scissors"&&AI_Answer=="Rock"{
		fmt.Println("YOU LOSE!!!!!!")
	}
	if item=="Rock"&&AI_Answer=="Scissors"{
		fmt.Println("YOU WIN!!!!!!")
	}
	if item=="Scissors"&&AI_Answer=="Paper"{
		fmt.Println("YOU WIN!!!!!!")
	}
	if item=="Paper"&&AI_Answer=="Scissors"{
		fmt.Println("YOU LOSE!!!!!!")
	}
return "0"
}

func main() {
	fmt.Println("Input your item:")
	var Human_item string
	fmt.Scanf("%s",&Human_item)
	Get_ItemAI(Human_item)
}
