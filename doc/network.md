# Kubernetes Service 随笔

## 环境信息
version: v1.18.2  
cni: flannel  
proxyMode: iptables

## iptables
从 **iptables** 说起

iptables: is a user-space utility program that allows a system administrator to configure the IP packet filter rules of the Linux kernel firewall, implemented as different Netfilter modules

[Iptables](https://en.wikipedia.org/wiki/Iptables)  
[Netfilter](https://en.wikipedia.org/wiki/Netfilter)

### iptables 四表五链的工作流程
```
tables: raw, mangle, filter, nat
chains: PREROUTING, FORWARD, INPUT, OUTPUT, POSTROUTING
```

![tables](./pictures/tablesflow.png)

**PREROUTING**
1. 数据包到达网络设备
2. 进入 iptables 的 raw 表，此时报文还未进入内核，属于原始报文
3. 进入 iptables 的 mangle 表，在 mangle 表里一般用来对报文做MARK
4. 进入 iptables 的 nat 表，一般用来做 dnat 功能（修改报文的目的ip）

**路由** 






**FORWARD**


**POSTROUTING**




## pod 的网络通信




## kubernetes service 实现原理分析
