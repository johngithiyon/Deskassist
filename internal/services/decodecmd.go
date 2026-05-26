package services

import "github.com/johngithiyon/Deskassist/internal/services/commands"


func Decodecmd(cmd string,arg string) {
   
	    if cmd == "openfolder" {
		   
			 commands.Openfolder(arg)
			   
		} else if cmd == "google" {
             
			 commands.Searchcmd(arg)
		 	
		} else if cmd == "openbrowser" {
			 
			commands.Openbrowser(arg)

		} else if cmd == "openfile" {
			  
			 commands.Openfile(arg)
		} else if cmd=="deskassist" && arg == "--help"{
			
			 commands.Helpcmd()		  
		}
        
	     
}