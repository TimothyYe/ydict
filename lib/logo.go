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

func DisplayUsage() {
	logo = ""
	color.Cyan(logo, Version)
	color.Cyan("Usage:")
	color.Cyan("ydict <word(s) to query>        Query the word(s)")
	color.Cyan("ydict -v <word(s) to query>     Query with speech")
	color.Cyan("ydict -m <word(s) to query>     Query with more example sentences")
	color.Cyan("ydict -q <word(s) to query>     Query with quiet mode, don't show spinner")
	color.Cyan("ydict -c <word(s) to query>     Query with local cache")
	color.Cyan("ydict -clear                    Clear local cache")
	color.Cyan("ydict -h                        For help")
}
