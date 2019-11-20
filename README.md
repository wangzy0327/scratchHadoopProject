# scratchHadoopProject
scratchHadoop

- 博客: [基于scratch构建轻量快速镜像](https://blog.csdn.net/qq_32062657/article/details/102950926)

```
与emptyImageProject项目类似，结合hadoop-cluster-docker共同使用
```

- hadoop-cluster-docker项目: [基于docker容器组构建hadoop集群](https://github.com/wangzy0327/hadoop-cluster-docker)

### 常用命令
- 静态编译go语言
```
编译为amd64指令集架构,项目名为scratchHadoopProject(在GOPATH路径src目录下),编译后的二进制文件为scratchHadoop
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o scratchHadoop -ldflags '-s' scratchHadoopProject/

编译为mips64le指令集架构(龙芯),项目名为scratchHadoopProject(在GOPATH路径src目录下),编译后的二进制文件为scratchHadoop0
CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -a -o scratchHadoop0 -ldflags '-s' scratchHadoopProject/

编译为arm64指令集架构,项目名为scratchHadoopProject(在GOPATH路径src目录下),编译后的二进制文件为scratchHadoop1
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -o scratchHadoop1 -ldflags '-s' scratchHadoopProject/

```

- 查看go语言支持的平台
```
go tool dist list

```

- 基于scratch构建Dockerfile运行
```
docker run -v /home/wzy/hadoop-cluster-docker:/home/wzy/hadoop-cluster-docker scratchhadoop /scratchHadoop --help

Usage of /scratchHadoop:
  -input string
    	hadoop input (default "input/input1")
  -output string
    	hadoop output (default "output/output1")

```
