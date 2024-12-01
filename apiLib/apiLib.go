package apiLib

import (
	"fmt"
	"os"
	"github.com/goccy/go-yaml"
)


func VerifyCmd(cmdStr string) (bool) {
    cmdList := []string{"list", "get", "add", "upd", "rm"}
    for i:=0; i< len(cmdList); i++ {
        if cmdStr == cmdList[i] {
            return true
        }
    }
    return false
}

func VerifyApp(appStr string) (bool) {

    appList := []string{"namcheap", "cloudflare", "nchsbox", "*", "all"}
    for i:=0; i< len(appList); i++ {
        if appStr == appList[i] {
            return true
        }
    }
    return false
}

type prov struct {
	Nam string
	Val provVal
}

type provVal struct {
	Token string
}

type provMap map[string]string

func GetList(yamlFil string) (list map[string]provMap, err error) {

//	var ProvList []prov

	list = make(map[string]provMap)

	ldat, err := os.ReadFile(yamlFil)
	if err != nil {return list, fmt.Errorf("Read List: %v", err)}

//	fmt.Printf("dbg -- %s\n", ldat)

	err = yaml.Unmarshal(ldat,&list)
	if err != nil {return list, fmt.Errorf("UnMarshal: %v", err)}

	return list, nil
}

func FindToken(app string, list map[string]provMap) (string, bool){
	appVal, ok:= list[app]
	if !ok {return "", false}
	
	token, _ := appVal["token"]

	return token, true
}

func PrintList(list map[string]provMap) {

	fmt.Printf("*** api list: %d providers ***\n", len(list))
	for nam, val := range list {
		fmt.Printf("provider: %s\n", nam)
		for key, kval := range val {
			fmt.Printf(" %s : %s\n", key, kval)
		}
	}
	fmt.Printf("*** end api list providers ***\n")
}
