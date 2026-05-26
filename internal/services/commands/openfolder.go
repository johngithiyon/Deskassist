package commands

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/johngithiyon/Deskassist/internal/errors"
)


func Openfolder(arg string ) {
	 
      cmd :=  exec.Command("find","/", "-type","d", "-name", arg)
	 
	  output,runerr := cmd.Output()

	  if runerr != nil {
		   log.Println(runerr)
	  }

	  log.Println(string(output))

	  var folderpath string 
	  
	  log.Println("Enter folder path")
	  
	  lines,scanerr := fmt.Scanf("%s",&folderpath)

	  if lines != 1 {
		    log.Println(errors.ErrInputlines)
			return 
	  }

	  if scanerr != nil {
		     log.Println(scanerr)
			 return 
	  }

	 _,openerr :=  exec.Command("xdg-open",folderpath).Output()


	 if openerr != nil {
		  
		    log.Println(openerr)
			return 
	 }


}