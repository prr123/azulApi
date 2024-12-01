//

package main

import (
	"fmt"
	"log"
	"os"

	apiLib "goDemo/api/apiLib"
	cliutil "github.com/prr123/utility/utilLib"
)


func main() {

    numarg := len(os.Args)
    flags:=[]string{"dbg", "cmd", "app", "val"}

    useStr := " /cmd={list,get,add,upd,rm} /app=appname </val={yamlfile}> [/dbg]"
    helpStr := "api vault program"

    if numarg > len(flags) +1 {
        fmt.Println("too many arguments in cl!")
        fmt.Println("usage: %s %s\n", os.Args[0], useStr)
        os.Exit(-1)
    }

    if numarg == 1 || (numarg > 1 && os.Args[1] == "help") {
        fmt.Printf("help: %s\n", helpStr)
        fmt.Printf("usage is: %s %s\n", os.Args[0], useStr)
        os.Exit(1)
    }

    flagMap, err := cliutil.ParseFlags(os.Args, flags)
    if err != nil {log.Fatalf("util.ParseFlags: %v\n", err)}

    dbg:= false
    _, ok := flagMap["dbg"]
    if ok {dbg = true}


    cmdStr := ""
    cmdval, ok := flagMap["cmd"]
    if !ok {
        log.Fatalf("error -- no command flag provided!\n")
    } else {
        if cmdval.(string) == "none" {log.Fatalf("error -- no command provided!\n")}
        cmdStr = cmdval.(string)
    }

    appStr := ""
   	appval, ok := flagMap["app"]
    if !ok {
        log.Fatalf("error -- no app flag provided!\n")
    } else {
        if appval.(string) == "none" {log.Fatalf("error -- no app name provided!\n")}
        appStr = appval.(string)
    }

    valStr := "list"
    val, vok := flagMap["val"]
    if vok {
        if val.(string) == "none" {log.Fatalf("error -- no yaml file provided!\n")}
        valStr = val.(string)
    }

    valFilnam := "yaml/" + valStr + ".yaml"

	if ok:= apiLib.VerifyCmd(cmdStr); !ok {log.Fatalf("error -- invalid command: %s\n", cmdStr)}
	if ok:= apiLib.VerifyApp(appStr); !ok {log.Fatalf("error -- invalid app: %s\n", appStr)}

    if dbg {
        fmt.Printf("cmd: %s\n", cmdStr)
        fmt.Printf("app: %s\n", appStr)
		fmt.Printf("val: %s\n", valFilnam)
//			fmt.Printf("no valFilnam\n")
    }

	applist, err := apiLib.GetList(valFilnam)
	if err != nil {log.Fatalf("error -- GetList: %v\n", err)}

	if dbg {apiLib.PrintList(applist)}

	tokVal, res := apiLib.FindToken("nchsbox", applist)

	fmt.Printf("token: %s, res %t\n", tokVal, res)

	fmt.Println("*** api success ***")
}

