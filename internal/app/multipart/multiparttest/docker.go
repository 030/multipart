package multiparttest

import (
	"strconv"
	"sync"

	"github.com/030/mij"
	log "github.com/sirupsen/logrus"
)

func Setup(containers []mij.DockerImage) error {
	log.SetLevel(log.DebugLevel)
	var wg sync.WaitGroup
	for _, container := range containers {
		wg.Add(1)
		go func(m mij.DockerImage) {
			defer wg.Done()
			if err := m.Run(); err != nil {
				panic(err)
			}
		}(container)
	}
	wg.Wait()
	return nil
}

func Shutdown(containers []mij.DockerImage) error {
	var wg sync.WaitGroup
	for _, container := range containers {
		wg.Add(1)
		go func(m mij.DockerImage) {
			defer wg.Done()
			if err := m.Stop(); err != nil {
				panic(err)
			}
		}(container)
	}
	wg.Wait()
	return nil
}

func Image(port int) mij.DockerImage {
	return mij.DockerImage{
		Name:                     "sonatype/nexus3",
		PortExternal:             port,
		PortInternal:             8081,
		Version:                  "3.16.0",
		ContainerName:            "nexus" + strconv.Itoa(port),
		LogFile:                  "/nexus-data/log/nexus.log",
		LogFileStringHealthCheck: "Started Sonatype Nexus OSS",
	}
}
