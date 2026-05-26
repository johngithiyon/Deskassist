package appstartup


import (
	"github.com/johngithiyon/Deskassist/internal/services"
)



func Appstartup() {
	 
	for(true) {

		cmd,arg,inputerr  := services.GetInput()
	
		
		if inputerr != nil {
			return 
		}
	
		services.Decodecmd(cmd,arg)
	
	}
}