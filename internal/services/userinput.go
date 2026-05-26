package services

import (
	"bufio"
	"log"
	"os"
	"strings"
	//err "github.com/johngithiyon/Deskassist/internal/errors"
)


func GetInput() (string,string,error) {

	  log.Println("Enter What I Have To Do !")
	  log.Println("Enter like [command] [argument]")

	 input := bufio.NewReader(os.Stdin)

	 inputstring,readerr := input.ReadString('\n')

	 userdata := strings.SplitN(inputstring," ",2)

	  if readerr != nil {
		  log.Println(readerr)
	  }

	  cmd := strings.TrimSpace(userdata[0])
	  arg := strings.TrimSpace(userdata[1])

	  return cmd,arg,nil 
}