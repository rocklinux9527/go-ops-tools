# ops-tools
***1.功能描述***
```
1.发送POST 访问Api接口
2.记录请求结果日志.
3.命令行输入定义参数-参数访问访问Api接口
```
***2.访问数据结构***
```
2.1.访问gitlab 接口 数据结构 (request)-App接口数据结构类似略...

{
 "LinkData":[],
 "ModelIdentify":"gitlab",
 "Data": {
 "GitlabIdentify":"1235",
 "GitlabAddress":"https://github.com/rocklinux9527/op-kube-manage-api.git",
 "GitlabName": "op-kube-manage-api",
 "GitlabDescribe": "测试应用"
 }
}

2.2.访问gitlab 接口 数据结构 (response)
{
"code": 1
"data": ""
"msg": "success"
}
```

***3.运行程序命令***
```
  chmod +x ./cmd/first-apply-order-linux
 ./cmd/first-apply-order-linux -appid 1345 -languageCode 5 -deployK8s "" -deployType "2"  -gitAddress "https://github.com/rocklinux9527/op-kube-manage-api.git"  -appName op-cicd-test -gray 0  -describe "测试应用"
```




