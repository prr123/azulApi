package apiLib

import (
	"fmt"
	"os"

	db "github.com/prr123/pogLib/pogLib"
	"github.com/goccy/go-yaml"

)
type ApiLib struct {
	pogdb *db.PogDB
	Dbg bool
}

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


func InitApi(dbdir string)(api *ApiLib, err error) {

	var apiObj ApiLib
	apiObj.pogdb, err = db.InitPogDb(dbdir)
	if err != nil {return nil, fmt.Errorf("InitApi-InitPodDb: %v\n", err)}
	return &apiObj, nil
}

func (api *ApiLib) ProcCmd(cmdStr, appNam, valStr string) (error) {

//  list := ul.list
    switch cmdStr {
    // list users
    case "list":
        if appNam == "*" || appNam == "all" {
            appList, err := api.ListAllApps()
			if err != nil {return fmt.Errorf("ListAllApps: %v\n",err)}
            for i, appNam := range appList  {
                fmt.Printf("--%d: %s\n", i, appNam)
            }
            return nil
        }
        ok, err:= api.ListApp(appNam)
        if err != nil {return fmt.Errorf("ListApp: %v", err)}
        if ok {
            fmt.Printf("dbg -- api: %s found!\n", appNam)
        } else {
            fmt.Printf("dbg -- api: %s not found!\n", appNam)
        }
        return nil

    case "get":
        fmt.Printf("dbg -- Cmd: get; App: %s\n", appNam)
        token, err := api.GetToken(appNam)
        if err != nil {return fmt.Errorf("GetToken: %v", err)}
        fmt.Printf("dbg -- App: %s Token: %s\n", appNam, token)
        return nil

    case "add":
        fmt.Printf("dbg -- Cmd: add; App: %s\n", appNam)
/*
		fmt.Print("App>")
		var appNamNew string
		fmt.Scanln(&appNamNew)
		fmt.Print("Token>")
		var token string
		fmt.Scanln(&token)
*/
        err := api.AddApp(appNam, valStr)
        if err != nil {return fmt.Errorf("AddApp: %v", err)}
        return nil

    case "rm":
        fmt.Printf("dbg -- Cmd: rm; App: %s\n", appNam)
        err := api.RmApp(appNam)
        if err != nil {return fmt.Errorf("RmApp: %v", err)}
        return nil

    case "upd":
        fmt.Printf("dbg -- Cmd: upd; App: %s\n", appNam)
        err := api.UpdApp(appNam, valStr)
        if err != nil {return fmt.Errorf("UpApp: %v", err)}
        return nil

    default:
        return fmt.Errorf("unknown command: %s\n", cmdStr)

    }

    return nil
}

func (api *ApiLib) ListApp(appNam string) (bool, error) {
	db := api.pogdb
	ok, err := db.HasKey(appNam)
	if err != nil {return false, fmt.Errorf("DbHas: %v\n", err)}
    return ok, nil
}

func (api *ApiLib) ListAllApps() (appList []string, err error) {

	db := api.pogdb
	appnum, err := db.DbCount()
	if err != nil {return appList, fmt.Errorf("DbCount: %v\n", err)}

    appList = make([]string, appnum)

	count:=0
    for i:=0; i<appnum; i++  {
		app,_, end, err := db.NextItem()
		if err != nil {return appList, fmt.Errorf("NextItem: %v\n", err)}
		if end {break}
        appList[i] = string(app)
		count++
    }
    return appList[:count], nil
}

func (api *ApiLib) GetToken(appNam string) (string, error){

	db := api.pogdb
    token, err := db.Read(appNam)
    if err != nil {return "", fmt.Errorf("GetToken: %v", err)}

    return string(token), nil
}

func (api *ApiLib) UpdApp(appNam, valStr string) (error){

	db := api.pogdb
	err := db.Upd(appNam, []byte(valStr))
    if err != nil {return fmt.Errorf("dbUpd: %v", err)}

    return nil
}

func (api *ApiLib) AddApp(appNam, token string) (error){

	db := api.pogdb
	err := db.Add(appNam, []byte(token))
    if err != nil {return fmt.Errorf("dbAdd: %v", err)}
    return nil
}

func (api *ApiLib) RmApp(appNam string) (error){

	db := api.pogdb
	err := db.Del(appNam)
    if err != nil {return fmt.Errorf("dbDel: %v", err)}
    return nil
}

func (api *ApiLib) DbClose() (error){
	db := api.pogdb
	err := db.Close()
    if err != nil {return fmt.Errorf("dbClose: %v", err)}
    return nil
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
