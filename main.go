package main

import (
	"scratchHadoopProject/http"
        "scratchHadoopProject/polling"
        "fmt"
	"flag"
        "time"
)


var input string
var output string

func Init(){
	flag.StringVar(&input,"input","input/input1","hadoop input")
	flag.StringVar(&output,"output","output/output1","hadoop output")
}

func main(){
        Init()
        flag.Parse()
	timeUnixNano := time.Now().UnixNano()
	timeUnixMirco := float64(timeUnixNano)/1000000000
	//fmt.Printf("纳秒时间为：%d\n",timeUnixNano)
	fmt.Printf("容器启动后时间戳为：%.3f\n",timeUnixMirco)
	var host string = "http://10.18.127.4:8081/hadoop"
	uuidStr := http.HttpPost(host,input,output)
	//var output = "/mnt/xfs/pipeline_server/output/workflowArr.log"
	//fmt.Printf("发送请求后时间戳(调整后)为:%d\n",time.Now().UTC().Unix()-ZONE_DIFF)
	fmt.Printf("发送请求后时间戳(调整后)为:%.3f\n",float64(time.Now().UnixNano())/1000000000)
	//fmt.Println(time.Now().UTC().Format(polling.TIME_LAYOUT))
	//fmt.Println(time.Now().UTC().Unix())
        fmt.Println("uuid str : "+uuidStr)
	polling.Track(uuidStr, output)
}



