**工具功能**
```
  主要针对prometheus 云原生项目模块 Alertmanager 模块发送feishu/lark 告警webhook 消息功能
```

**项目初始化**
```go基础环境
go version 1.23.0
```

```方式1
go mod  init op-prome-alert #项目初始化
go mod tidy  #安装项目依赖
```
```方式2
go mod  init op-prome-alert #项目初始化
go get -u github.com/gin-gonic/gin #逐个安装依赖
go get -u github.com/go-playground/validator/v10 #逐个安装依赖
go get -u github.com/sirupsen/logrus #逐个安装依赖
```

**客户端测试**

```发送 单个告警场景
curl --location --request POST 'http://127.0.0.1:8080/report_alert?lark_webhook="your webhook 注意不支持带认证" \
--header 'Content-Type: application/json' \
--data-raw '{
           "status": "firing",
           "alerts": [
             {
               "status": "firing",
               "labels": {
                 "alertname": "KubeProxyDown",
                 "severity": "critical",
                 "prometheus": "monitoring/kube-prometheus-stack-prometheus"
               },
               "annotations": {
                 "description": "KubeProxy has disappeared from Prometheus target discovery.",
                 "runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/kubernetes/kubeproxydown",
                 "summary": "Target disappeared from Prometheus target discovery."
               },
               "startsAt": "2024-12-01T03:09:00Z",
               "endsAt": "0001-01-01T00:00:00Z",
               "generatorURL": "http://kube-prometheus-stack-prometheus.monitoring:9090/graph?g0.expr=absent%28up%7Bjob%3D%22kube-proxy%22%7D+%3D%3D+1%29&g0.tab=1"
             }
           ],
           "commonLabels": {
             "alertname": "KubeProxyDown",
             "severity": "critical",
             "prometheus": "monitoring/kube-prometheus-stack-prometheus"
           },
           "commonAnnotations": {
             "description": "KubeProxy has disappeared from Prometheus target discovery.",
             "runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/kubernetes/kubeproxydown",
             "summary": "Target disappeared from Prometheus target discovery."
           },
           "externalURL": "http://kube-prometheus-stack-alertmanager.monitoring:9093",
           "groupKey": "{}:{alertname=\"KubeProxyDown\"}",
           "groupLabels": {
             "alertname": "KubeProxyDown"
           },
           "receiver": "default-receiver",
           "status": "firing",
           "truncatedAlerts": 0,
           "version": "4"
         }'\''
}'
```
```发送2条告警场景
curl --location --request POST 'http://127.0.0.1:8080/report_alert?lark_webhook='your webhook 机器人注意不支持携带认证' \
--header 'Content-Type: application/json' \
--data-raw '{
           "status": "firing",
           "alerts": [
             {
               "status": "firing",
               "labels": {
                 "alertname": "KubeProxyDown",
                 "severity": "critical",
                 "prometheus": "monitoring/kube-prometheus-stack-prometheus"
               },
               "annotations": {
                 "description": "KubeProxy has disappeared from Prometheus target discovery.",
                 "runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/kubernetes/kubeproxydown",
                 "summary": "Target disappeared from Prometheus target discovery."
               },
               "startsAt": "2024-12-01T03:09:00Z",
               "endsAt": "0001-01-01T00:00:00Z",
               "generatorURL": "http://kube-prometheus-stack-prometheus.monitoring:9090/graph?g0.expr=absent%28up%7Bjob%3D%22kube-proxy%22%7D+%3D%3D+1%29&g0.tab=1"
             },
             {
               "status": "firing",
               "labels": {
                 "alertname": "NodeDown",
                 "severity": "critical",
                 "prometheus": "monitoring/kube-prometheus-stack-prometheus"
               },
               "annotations": {
                 "description": "Node has gone down and is unreachable.",
                 "runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/kubernetes/nodedown",
                 "summary": "Target node unreachable."
               },
               "startsAt": "2024-12-01T03:15:00Z",
               "endsAt": "0001-01-01T00:00:00Z",
               "generatorURL": "http://kube-prometheus-stack-prometheus.monitoring:9090/graph?g0.expr=absent%28up%7Bjob%3D%22node%22%7D+%3D%3D+1%29&g0.tab=1"
             }
           ],
           "commonLabels": {
             "alertname": "KubeProxyDown",
             "severity": "critical",
             "prometheus": "monitoring/kube-prometheus-stack-prometheus"
           },
           "commonAnnotations": {
             "description": "KubeProxy has disappeared from Prometheus target discovery.",
             "runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/kubernetes/kubeproxydown",
             "summary": "Target disappeared from Prometheus target discovery."
           },
           "externalURL": "http://kube-prometheus-stack-alertmanager.monitoring:9093",
           "groupKey": "{}:{alertname=\"KubeProxyDown\"}",
           "groupLabels": {
             "alertname": "KubeProxyDown"
           },
           "receiver": "default-receiver",
           "status": "firing",
           "truncatedAlerts": 0,
           "version": "4"
         }

}'
```
```发送单条告警恢复数据

curl --location --request POST 'http://127.0.0.1:8080/report_alert?lark_webhook=your webhook 机器人注意不支持携带认证' \
--header 'Content-Type: application/json' \
--data-raw '{
           "status": "resolved",
           "alerts": [
             {
               "status": "resolved",
               "labels": {
                 "alertname": "KubeProxyDown",
                 "severity": "critical",
                 "prometheus": "monitoring/kube-prometheus-stack-prometheus"
               },
               "annotations": {
                 "description": "KubeProxy has disappeared from Prometheus target discovery.",
                 "runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/kubernetes/kubeproxydown",
                 "summary": "Target disappeared from Prometheus target discovery."
               },
               "startsAt": "2024-12-01T03:09:00Z",
               "endsAt": "2024-12-01T04:09:00Z",
               "generatorURL": "http://kube-prometheus-stack-prometheus.monitoring:9090/graph?g0.expr=absent%28up%7Bjob%3D%22kube-proxy%22%7D+%3D%3D+1%29&g0.tab=1"
             }
           ],
           "commonLabels": {
             "alertname": "KubeProxyDown",
             "severity": "critical",
             "prometheus": "monitoring/kube-prometheus-stack-prometheus"
           },
           "commonAnnotations": {
             "description": "KubeProxy has disappeared from Prometheus target discovery.",
             "runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/kubernetes/kubeproxydown",
             "summary": "Target disappeared from Prometheus target discovery."
           },
           "externalURL": "http://kube-prometheus-stack-alertmanager.monitoring:9093",
           "groupKey": "{}:{alertname=\"KubeProxyDown\"}",
           "groupLabels": {
             "alertname": "KubeProxyDown"
           },
           "receiver": "default-receiver",
           "status": "resolved",
           "truncatedAlerts": 0,
           "version": "4"
         }
}'
```

```发送2条告警恢复数据
curl --location --request POST 'http://127.0.0.1:8080/report_alert?lark_webhook=your webhook 机器人注意不支持携带认证' \
--header 'Content-Type: application/json' \
--data-raw '{
           "status": "resolved",
           "alerts": [
             {
               "status": "resolved",
               "labels": {
                 "alertname": "KubeProxyDown",
                 "severity": "critical",
                 "prometheus": "monitoring/kube-prometheus-stack-prometheus"
               },
               "annotations": {
                 "description": "KubeProxy has disappeared from Prometheus target discovery.",
                 "runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/kubernetes/kubeproxydown",
                 "summary": "Target disappeared from Prometheus target discovery."
               },
               "startsAt": "2024-12-01T03:09:00Z",
               "endsAt": "2024-12-01T04:09:00Z",
               "generatorURL": "http://kube-prometheus-stack-prometheus.monitoring:9090/graph?g0.expr=absent%28up%7Bjob%3D%22kube-proxy%22%7D+%3D%3D+1%29&g0.tab=1"
             },
             {
               "status": "resolved",
               "labels": {
                 "alertname": "NodeDown",
                 "severity": "critical",
                 "prometheus": "monitoring/kube-prometheus-stack-prometheus"
               },
               "annotations": {
                 "description": "Node has gone down and is unreachable.",
                 "runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/kubernetes/nodedown",
                 "summary": "Target node unreachable."
               },
               "startsAt": "2024-12-01T03:15:00Z",
               "endsAt": "2024-12-01T04:15:00Z",
               "generatorURL": "http://kube-prometheus-stack-prometheus.monitoring:9090/graph?g0.expr=absent%28up%7Bjob%3D%22node%22%7D+%3D%3D+1%29&g0.tab=1"
             }
           ],
           "commonLabels": {
             "alertname": "KubeProxyDown",
             "severity": "critical",
             "prometheus": "monitoring/kube-prometheus-stack-prometheus"
           },
           "commonAnnotations": {
             "description": "KubeProxy has disappeared from Prometheus target discovery.",
             "runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/kubernetes/kubeproxydown",
             "summary": "Target disappeared from Prometheus target discovery."
           },
           "externalURL": "http://kube-prometheus-stack-alertmanager.monitoring:9093",
           "groupKey": "{}:{alertname=\"KubeProxyDown\"}",
           "groupLabels": {
             "alertname": "KubeProxyDown"
           },
           "receiver": "default-receiver",
           "status": "resolved",
           "truncatedAlerts": 0,
           "version": "4"
         }
}'
```

**告警显示**

![带告警模版实例](https://github.com/rocklinux9527/go-ops-tools/blob/master/op-prome-lark/assets/multiple-alarm-firing.png)

![带告警模版实例](https://github.com/rocklinux9527/go-ops-tools/blob/master/op-prome-lark/assets/multiple-alarm-resolved.png)

![带告警模版实例](https://github.com/rocklinux9527/go-ops-tools/blob/master/op-prome-lark/assets/single-alarm-firing.png)

![带告警模版实例](https://github.com/rocklinux9527/go-ops-tools/blob/master/op-prome-lark/assets/single-alarm-resolved.png)



