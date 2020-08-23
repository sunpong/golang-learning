# Golang learning

## Kubernetes
1. 写一个控制器，当一个应用启动的时候，保证其他的node上也有其镜像
2. Prometheus 部署 [kube-prometheus](https://github.com/coreos/kube-prometheus)
3. kubectl plugin + https://github.com/ahmetb/kubectx
1. 容器内部访问 k8s api https://www.jianshu.com/p/b1a723033a3c
2.  k8s 首先看 event 和 ns
3.  自定义crd开发
4.  k8s ingress 跨域 https://blog.csdn.net/u012375924/article/details/94360425
5.  K8s pod 出外网
6.  Docker 集成 ovs [ovs](https://docs.openvswitch.org/en/latest/intro/install/general/#obtaining-open-vswitch-sources)

## Golang
1. log 的使用 — logrus
2. 并发 goroutine —— ok
3. slice 排序 —— ok
4. channel 实现并发和锁，（排他锁）—— half ok
5. go-restful
6. cobra
