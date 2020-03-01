package main

import "flag"

func main() {
	hosts := flag.String("h", "./hosts.txt", "host or ip list file path.")
	cmds := flag.String("c", "./cmds.txt", "cmd list you wanna execute on hosts.")
	scp := flag.String("s", "./goods/*.*", "all files in folder `./goods` will be transfer to hosts.")
	//jsons := flag.String("j", "./allinone.json", "Json file for fit all your demand.")
	flag.Parse()
	select {}
}
