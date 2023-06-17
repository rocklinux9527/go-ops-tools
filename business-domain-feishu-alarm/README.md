**1.部署**
```
bash build.sh   #编译二进制可执行文件 
``` 
**2.使用**
```
cd ./cmd  && chmod +x domain-alarm-linux-64  && ./domain-alarm-linux-64  #注意配置文件config.yaml 必须在同一目录中

```
**3.飞书**
```
  主要使用飞书机器人 webhook方式secret方式进行鉴权验证(需要配置正确的地址和secret)
```
**4.消息模版**
```
URL域名监控
检查周期: 10min
检测状态: 503
检查时间: 2023-06-17 18:56:41
URL域名: https://www.51-devops.com/503
域名用途: 业务域名-503
```
