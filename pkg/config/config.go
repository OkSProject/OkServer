package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type SiteInfo struct {
	Domain string `yaml:"domain"`
	Root   string `yaml:"root"`
}

type SiteList struct {
	Sites []SiteInfo `yaml:"sites"`
}

func GetLogDir() string {
	log_dir := os.Getenv("OKSERVER_LOG_DIR")
	if log_dir == "" {
		log_dir = "./log/"
	}
	return log_dir
}

func GetSitesFile() string {
	sites_file := "./pkg/sites/sites.yaml"
	if sites_file == "" {
		sites_file = "./pkg/sites/sites.yaml"
	}
	return sites_file
}

func GetRootDir() string {
	// Read sites.yaml
	sites_file := GetSitesFile()
	site_yaml, err := os.ReadFile(sites_file)
	if err != nil {
		log.Fatalf("Error opening sites.yaml: ", err)
	}

	// Parse sites.yaml into the SiteInfo struct.
	var site_list SiteList
	err = yaml.Unmarshal(site_yaml, &site_list)
	if err != nil {
		log.Fatalf("Error parsing sites.yaml: %v", err)
	}

	/*
		Makes first entry the default, and determines how many entries are in sites.yaml
		using len().
	*/
	if len(site_list.Sites) > 0 {
		return site_list.Sites[0].Root
	}

	// Return main environmental root if no other entries are detected.
	root_dir := os.Getenv("OKSERVER_ROOT_DIR")
	return root_dir
}

func GetListenPort() string {
	listen_port := os.Getenv("OKSERVER_PORT")
	if listen_port == "" {
		listen_port = "8080"
	}
	return listen_port
}
