package main


import (
	log "github.com/sirupsen/logrus"
	"fmt"
	"github.com/coolboy43/mydocker/container"
	"os/exec"
)

func commitContainer(containerName, imageName string){
	mntURL := fmt.Sprintf(container.MntUrl, containerName)
	mntURL += "/"

	imageTar := container.RootUrl + "/" + imageName + ".tar"

	if _, err := exec.Command("tar", "-czf", imageTar, "-C", mntURL, ".").CombinedOutput(); err != nil {
		log.Errorf("Tar folder %s error %v", mntURL, err)
	}
}
