package main

import (
	"github.com/johngithiyon/Deskassist/internal/services"
)

func main() {

	for(true) {

	cmd,arg,inputerr  := services.GetInput()

	
	if inputerr != nil {
		return 
	}

    services.Decodecmd(cmd,arg)

}

}