package commands

import (
	"log"
	"os/exec"
)

func Openfile(dirname string) {

         _,filerr := exec.Command("xdg-open",dirname).Output()

		 if filerr != nil {
			   
			   log.Println(filerr)
			   return
		 }

	
}