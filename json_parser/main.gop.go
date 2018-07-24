package json_parser

import (
"fmt"
"regexp"
"strings"
//"encoding/json"
"os/exec"
"os"
)

func main() {
	logRex := regexp.MustCompile(`(?P<datetime>.*)\sIP\s(?P<serverIp>.*)\.(?P<serverPort>\d{2,5})\s\>\s(?P<clientIp>.*)\.(?P<clientPort>\d{2,5})\:.*length\s(?P<length>\d*)`)

	match := logRex.FindStringSubmatch(os.Args[1])//"2018-07-24 18:46:03.467264 IP 172.17.0.2.8000 > 172.17.0.1.36866: Flags [P.], seq 2726929042:2726931422, ack 3983119563, win 235, options [nop,nop,TS val 2247612 ecr 2247578], length 2380")

	result := make(map[string]string)
	for i, name := range logRex.SubexpNames() {
		fmt.Println(name, match)
		if i != 0 && name != "" {
			result[name] = match[i]

			if name == "ServerPort" {
				nlp := exec.Command("netstat", "-nlp | grep "+match[i]+" | awk '{ print $7 }'")  //"11/./sc_serv"
				nlpout, _ := nlp.Output()

				pid := strings.Split(string(nlpout), "/")[0]
				pwdx := exec.Command("pwdx", pid) //strings.Split(pwdx[pid], "/")
				pwdxout, _ := pwdx.Output()
				pidpath := strings.Split(string(pwdxout), "/")

				fmt.Println(pidpath)
				contractAndStreamId := strings.Split(pidpath[len(pidpath)-1], "_")
				result["contractCode"] = contractAndStreamId[0]

				fmt.Println(result)
				//var streamId int
				//if streamId = 1; 1 < len(contractAndStreamId) {
				//	streamId = contractAndStreamId[1]
				//}

				//result["streamId"] = streamId
			}
		}
	}

	//js, _ := json.Marshal(result)

	//fmt.Println(","+string(js))
	//fmt.Println(result)
}