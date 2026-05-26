package appstartup

import (
	"log"

	"github.com/johngithiyon/Deskassist/internal/services"
)



func Appstartup() {

	log.Println("To Get Help Use: deskassist --help")
	 
	for(true) {

		cmd,arg,inputerr  := services.GetInput()	
		
		if inputerr != nil {
			return 
		}
	
		services.Decodecmd(cmd,arg)
	
	}
}