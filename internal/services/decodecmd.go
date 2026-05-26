package services

import "github.com/johngithiyon/Deskassist/internal/services/commands"


func Decodecmd(cmd string,arg string) {
   
	    if cmd == "open" {
		   
			 commands.Open(arg)
			   
		} else if cmd == "google" {
             
			 commands.Searchcmd(arg)
		 	
		} else if cmd == "openbrowser" {
			 
			commands.Openbrowser(arg)
		}
        
	     
}