package lib

import "github.com/fatih/color"

var (
	//Version of ydict
	Version = "0.1"
	logo    = `
██╗   ██╗██████╗ ██╗ ██████╗████████╗
╚██╗ ██╔╝██╔══██╗██║██╔════╝╚══██╔══╝
 ╚████╔╝ ██║  ██║██║██║        ██║   
  ╚██╔╝  ██║  ██║██║██║        ██║   
   ██║   ██████╔╝██║╚██████╗   ██║   
   ╚═╝   ╚═════╝ ╚═╝ ╚═════╝   ╚═╝   

YDict V%s
https://github.com/TimothyYe/ydict

`
)

func DisplayLogo() {
	color.Cyan(logo, Version)
}
