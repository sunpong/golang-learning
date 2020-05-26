package main

import "fmt"

// https://itnext.io/how-it-works-kubectl-exec-e31325daa910
// https://erkanerol.github.io/post/how-kubectl-exec-works/
// 分析 kubectl exec 的工作原理

// 学习实践 kubectl 的 plugin https://github.com/kubernetes/sample-cli-plugin
// https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/sample-cli-plugin

// 编辑可执行文件，编译时 以 kubeclt-xxx开始，然后把二进制拷贝到PATH下，使用kubectl xxx 调用
// https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/

func main() {

	fmt.Println("hello kubectl plugin")
	
}

