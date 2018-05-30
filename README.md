# 说明
框架代码来自于https://github.com/anishathalye/6.824-golabs-2017-experimental
/src/labrpc是用go channel仿真的不可靠网络上的rpc
/src/raft/test_test.go 是测试程序
leader选举在/src/raft/raft.go中实现

# 环境配置和运行
## 安装go
https://www.linode.com/docs/development/go/install-go-on-ubuntu/

##将leaderElection目录添加到GOPATH
export GOPATH=$PWD/leaderElection/:$GOPATH

##切换目录到源码下，执行测试程序
cd leaderElection/src/raft
###运行测试用例
go test
###运行测试1000次
for i in $(seq 1 1000); do go test; done
###开启日志
日志默认关闭，util.go中const Debug = 1开启日志

#可视化部分
可视化部分是python程序visual-le.py
##安装wxpython 
###Debian
apt-get install python-wxgtk2.8
###MAC OS
brew install wxpython
##运行
python  visual-le.py
