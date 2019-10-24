package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"syscall"
	"time"
)

func logger(sigLog <-chan procInfo, sigExec <-chan procInfo) {
	for {
		select {
		case info := <-sigLog:
			fileName := targetProg + time.Now().String()
			ioutil.WriteFile(coverageDir+fileName, info.args, 0644)
			fmt.Printf("\nJACKPOT!!\n% 0x\n", info.args)
			log.Printf("Found at: %s\n", time.Now())
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		case info := <-sigExec:
			updateCov(info)
		}
	}
}

func updateCov(info procInfo) {
	fileName := targetProg + "." + strconv.Itoa(info.pid) + ".sancov"
	covData, err := ioutil.ReadFile(coverageDir + "tmp/" + fileName)
	checkErr(err)

	for i := 8; i < len(covData); i += 8 {
		if !covAddrs[string(covData[i:i+8])] {
			covAddrs[string(covData[i:i+8])] = true
			copy(baseTest, info.args)
			fmt.Printf("\n\nNEW COVADDR:% X MUT: % X\n", covData[i:i+8], info.args)
		}
	}

	os.Remove(coverageDir + "tmp/" + fileName)

}
