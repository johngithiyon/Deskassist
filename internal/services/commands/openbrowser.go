package commands

import (
	"log"
	"os/exec"
)

func Openbrowser(arg string) {

      
	     if arg == "chrome" {
	 
	     _,browerr := exec.Command("google-chrome").Output()

		 if browerr != nil {
			 log.Println(browerr)
			 return 
		 }  

       } else if arg == "firefox" {
		_,browerr := exec.Command(arg).Output()

		if browerr != nil {
			log.Println(browerr)
			return 
		}  
		       
	   } else if arg == "edge" {
		  
		_,browerr := exec.Command("microsoft-edge").Output()

		if browerr != nil {
			log.Println(browerr)
			return 
		}  
		     
	   } else {
		   log.Println("Not found in the system")
		   return 
	   }



}