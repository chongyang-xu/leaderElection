# 说明
框架代码来自于https://github.com/anishathalye/6.824-golabs-2017-experimental

/src/labrpc是用go channel 仿真的不可靠网络上的rpc

/src/raft/test_test.go 是提供的测试程序



实现leader选举需要填充一个文件

/src/raft/raft.go

# 运行
export GOPATH=$PWD/leaderElection/:$GOPATH

cd leaderElection/src/raft

go test -run 2A
