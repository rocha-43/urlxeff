package main

import (
	"flag"
	"fmt"

	"urlxeff/mdirectory"
	"urlxeff/msub"
	"urlxeff/sdirectory"
	"urlxeff/ssub"
)

func main() {
	var target = flag.String("t", "", "[!] Supply a target")
	var targetWordlist = flag.String("tw", "", "[!] Supply a target list (.txt)")
	var wordlist = flag.String("wl", "", "[!] Supply a wordlist.")
	var mode = flag.String("m", "default", "")
	var outFileName = flag.String("n", "", "[!] Need a name for the output file.")

	flag.Parse()

	help := `
 _    _ _____  _         _____ _____            ______ _______ ______ _____  
| |  | |  __ \| |       / ____|  __ \     /\   |  ____|__   __|  ____|  __ \ 
| |  | | |__) | |      | |    | |__) |   /  \  | |__     | |  | |__  | |__) |
| |  | |  _  /| |      | |    |  _  /   / /\ \ |  __|    | |  |  __| |  _  / 
| |__| | | \ \| |____  | |____| | \ \  / ____ \| |       | |  | |____| | \ \ 
 \____/|_|  \_\______|  \_____|_|  \_\/_/    \_\_|       |_|  |______|_|  \_\

USAGE:

-t target
-tw target wordlist (.txt)
-wl wordlist (.txt)
-m mode (single-directory | multiple-directory | single-sub | multiple-sub)
-n output file name (just the name. Will be a .txt)
	`

	switch *mode {
	case "single-directory":
		sdirectory.Sdirectory(*target, *wordlist, *outFileName)
	case "multiple-directory":
		mdirectory.Mdirectory(*targetWordlist, *wordlist, *outFileName)
	case "single-sub":
		ssub.Ssubdomain(*target, *wordlist, *outFileName)
	case "multiple-sub":
		msub.Msub(*targetWordlist, *wordlist, *outFileName)
	default:
		fmt.Println(help)
	}
}
