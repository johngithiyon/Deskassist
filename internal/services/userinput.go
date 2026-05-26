package services

import (
	"fmt"
	"log"

	err "github.com/johngithiyon/Deskassist/internal/errors"
)


func GetInput() (string,string,error) {

	  var command string
	  var argument string  
	
	  log.Println("Enter What I Have To Do !")
	  log.Println("Enter like [command] [argument]")

	  lines,inputerr := fmt.Scanf("%s %s",&command,&argument)

	  if lines != 2 {
	  
		  log.Println(err.ErrInputlines)
		  return "","",err.ErrInputlines
			
		}

		if inputerr != nil {

			 log.Println(inputerr)
			 return "","",inputerr
		}

	  return command,argument,nil 
}