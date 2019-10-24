package main

import (
	"fmt"
	"os"
	"os/exec"
)

func worker(sigLog chan<- procInfo, sigExec chan<- procInfo, sigWork <-chan []byte) {

	var pInfo procInfo
	count := 0
	for testCase := range sigWork {
		count++
		pInfo.args = make([]byte, len(testCase))
		copy(pInfo.args, testCase)

		cmd := exec.Command("./"+targetProg, string(testCase))
		cmd.Env = append(os.Environ(), "ASAN_OPTIONS=coverage=1:coverage_dir="+coverageDir+"tmp")
		err := cmd.Run()
		pInfo.pid = cmd.ProcessState.Pid()
		if err != nil {
			sigLog <- pInfo
			fmt.Println(err)
		} else {
			sigExec <- pInfo
		}
		fmt.Printf("\r#%d || MUT:% 0#x", count, testCase[0:5])
	}
}
