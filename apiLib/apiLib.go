package apiLib

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
    appList := []string{"namcheap", "cloudflare", "nchsb", "*", "all"}
    for i:=0; i< len(appList); i++ {
        if appStr == appList[i] {
            return true
        }
    }
    return false
}


