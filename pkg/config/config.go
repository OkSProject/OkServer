package config

import "os"

func GetLogDir() string {
	d := os.Getenv("OKSERVER_LOG_DIR")
	if d == "" {
		d = "./log/"
	}
	return d
}

func GetAssetDir() string {
	d := os.Getenv("OKSERVER_ASSET_DIR")
	if d == "" {
		d = "./www/html/"
	}
	return d
}

func GetListenPort() string {
	p := os.Getenv("OKSERVER_PORT")
	if p == "" {
		p = "8080"
	}
	return p
}
