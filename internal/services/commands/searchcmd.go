package commands

import (
	"log"
	"os/exec"
)


func Searchcmd(arg string ) {
 
	     searchout,searcherr := exec.Command("xdg-open","https://google.com/search?q="+arg).Output()

         if searcherr != nil {
			log.Println(searcherr)
			return 
		 }

		 log.Println(string(searchout))
	     
}