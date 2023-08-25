***1.工具部署***
```
  1.1.go 初始化
      go mod init k8s-format-yaml
  
  1.2.下载依赖
      go mod tidy
      
  1.3.环境编译
      bash -x build.sh 
```
***2.工具运行***
```
   2.1 cmd/k8s-format-yaml-linux -sourcePath source.yaml -targetPath test-format.yaml
```
