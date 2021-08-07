package util

import "os"

func GetConfigFile(mode string) string {
	if mode == "prod" {
		if FileExist("conf/app-prod.conf") {
			return "conf/app-prod.conf"
		} else {
			return "conf/app.conf"
		}
	} else if mode == "dev" {
		if FileExist("conf/app-dev.conf") {
			return "conf/app-dev.conf"
		} else {
			return "conf/app.conf"
		}
	} else {
		return "conf/app.conf"
	}
}


func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}