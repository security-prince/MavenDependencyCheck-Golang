package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var result_dir = "depcheckreport"
var project_dir = "projects"
var dependency_check = "dependency-check"

// "main() Downloads Dependency-check and creates directories"

func main() {

	if _, err := os.Stat("result_dir"); os.IsNotExist(err) {
		fmt.Println("The result directory not present so creating directory")
		os.Mkdir(result_dir, 0777)

	}
	if _, err := os.Stat("project_dir"); os.IsNotExist(err) {
		fmt.Println("The projects directory not present so creating directory")
		os.Mkdir(project_dir, 0777)
	}
	if _, err := os.Stat("dependency_check_launcher"); os.IsNotExist(err) {
		os.Mkdir(dependency_check, 0777)
	}
	os.Chdir("dependency_check_launcher")
	depcheck_download := exec.Command("wget", "https://dl.bintray.com/jeremy-long/owasp/dependency-check-5.3.0-release.zip")
	log.Printf("Downloading dependency check...")
	depcheck_download.Run()
	log.Println("Dependency Check Downloaded")
	unzip_launcher := exec.Command("unzip", "dependency-check-5.3.0-release.zip")
	log.Printf("Extracting Dependency Check...")
	unzip_launcher.Run()

}
