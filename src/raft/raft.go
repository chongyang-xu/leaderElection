package raft

//
// this is an outline of the API that raft must expose to
// the service (or tester). see comments below for
// each of these functions for more details.
//
// rf = Make(...)
//   create a new Raft server.
// rf.Start(command interface{}) (index, term, isleader)
//   start agreement on a new log entry
// rf.GetState() (term, isLeader)
//   ask a Raft for its current term, and whether it thinks it is leader
// ApplyMsg
//   each time a new entry is committed to the log, each Raft peer
//   should send an ApplyMsg to the service (or tester)
//   in the same server.
//

import "sync"
import "labrpc"
import "time"
import "math/rand"

// import "bytes"
// import "encoding/gob"



//
// as each Raft peer becomes aware that successive log entries are
// committed, the peer should send an ApplyMsg to the service (or
// tester) on the same server, via the applyCh passed to Make().
//
type ApplyMsg struct {
	Index       int
	Command     interface{}
	UseSnapshot bool   // ignore for lab2; only used in lab3
	Snapshot    []byte // ignore for lab2; only used in lab3
}

//
// A Go object implementing a single Raft peer.
//
type Raft struct {
	mu        sync.Mutex          // Lock to protect shared access to this peer's state
	peers     []*labrpc.ClientEnd // RPC end points of all peers
	persister *Persister          // Object to hold this peer's persisted state
	me        int                 // this peer's index into peers[]

	// Your data here (2A, 2B, 2C).
	// Look at the paper's Figure 2 for a description of what
	// state a Raft server must maintain.
	/**/
	currentTerm int
	votedFor int
	log	[]LogEntry

	commitIndex int
	lastApplied int

	nextIndex []int
	matchIndex []int

	role int
	electTimer *time.Timer
	discoverLeader  chan int
	newTerm    chan int
	winElect    chan int
	/**/
}
/**/
const(
	follower = 0
	candicate = 1
	leader = 2
)

type LogEntry struct{
	Term int
	Command interface{}
}
/**/

// return currentTerm and whether this server
// believes it is the leader.
func (rf *Raft) GetState() (int, bool) {

	//var term int
	//var isleader bool
	
	// Your code here (2A).
	/**/
	rf.mu.Lock()
	defer rf.mu.Unlock()
	return rf.currentTerm, rf.role==leader
	/**/
	//return term, isleader
}

//
// save Raft's persistent state to stable storage,
// where it can later be retrieved after a crash and restart.
// see paper's Figure 2 for a description of what should be persistent.
//
func (rf *Raft) persist() {
	// Your code here (2C).
	// Example:
	// w := new(bytes.Buffer)
	// e := gob.NewEncoder(w)
	// e.Encode(rf.xxx)
	// e.Encode(rf.yyy)
	// data := w.Bytes()
	// rf.persister.SaveRaftState(data)
}

//
// restore previously persisted state.
//
func (rf *Raft) readPersist(data []byte) {
	// Your code here (2C).
	// Example:
	// r := bytes.NewBuffer(data)
	// d := gob.NewDecoder(r)
	// d.Decode(&rf.xxx)
	// d.Decode(&rf.yyy)
	if data == nil || len(data) < 1 { // bootstrap without any state?
		return
	}
}




//
// example RequestVote RPC arguments structure.
// field names must start with capital letters!
//
type RequestVoteArgs struct {
	// Your data here (2A, 2B).
	/**/
	Term int
	CandicateId int
	LastLogIndex int
	LastLogTerm int
	/**/
}

//
// example RequestVote RPC reply structure.
// field names must start with capital letters!
//
type RequestVoteReply struct {
	// Your data here (2A).
	/**/
	Term int
	VoteGranted bool
	/**/
}

//
// example RequestVote RPC handler.
//
func (rf *Raft) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {
	// Your code here (2A, 2B).
	/**/
	//DPrintf("Invoke RequestVote\n")
	curTerm := 0
	voted := -1
	//var log []LogEntry

	rf.withLock(func(){
		curTerm  = rf.currentTerm
		voted = rf.votedFor
		//log = rf.log
	})

	switch{
	case args.Term < curTerm :
		reply.Term = curTerm
		reply.VoteGranted = false
		return
	case args.Term > curTerm://执行到这儿，发送channl之前宕机怎么办
		rf.withLock(func(){
			rf.currentTerm = args.Term
			rf.role = follower
			rf.votedFor = -1
			curTerm = args.Term
			voted = -1
		})
		rf.newTerm <- 1
		fallthrough
	case args.Term == curTerm:
		cond1 := voted == -1 || voted == args.CandicateId
		//cond2 := args.LastLogTerm > rf.log[len(rf.log)-1].Term || ( args.LastLogTerm == rf.log[len(rf.log)-1].Term && args.LastLogIndex >= len(rf.log) )
		cond2 := true
		if cond1 && cond2 {
			reply.Term = curTerm
			reply.VoteGranted = true
			rf.withLock(func(){rf.votedFor = args.CandicateId})
			rf.withLock( func(){ rf.resetOrNewElectTimer(randomTimeout())} )
			return
		}else{
			reply.Term = curTerm
			reply.VoteGranted = false
			return
		}
	}
	/**/
}

//
// example code to send a RequestVote RPC to a server.
// server is the index of the target server in rf.peers[].
// expects RPC arguments in args.
// fills in *reply with RPC reply, so caller should
// pass &reply.
// the types of the args and reply passed to Call() must be
// the same as the types of the arguments declared in the
// handler function (including whether they are pointers).
//
// The labrpc package simulates a lossy network, in which servers
// may be unreachable, and in which requests and replies may be lost.
// Call() sends a request and waits for a reply. If a reply arrives
// within a timeout interval, Call() returns true; otherwise
// Call() returns false. Thus Call() may not return for a while.
// A false return can be caused by a dead server, a live server that
// can't be reached, a lost request, or a lost reply.
//
// Call() is guaranteed to return (perhaps after a delay) *except* if the
// handler function on the server side does not return.  Thus there
// is no need to implement your own timeouts around Call().
//
// look at the comments in ../labrpc/labrpc.go for more details.
//
// if you're having trouble getting RPC to work, check that you've
// capitalized all field names in structs passed over RPC, and
// that the caller passes the address of the reply struct with &, not
// the struct itself.
//
func (rf *Raft) sendRequestVote(server int, args *RequestVoteArgs, reply *RequestVoteReply) bool {
	ok := rf.peers[server].Call("Raft.RequestVote", args, reply)
	return ok
}


/**/
type AppendEntriesArgs struct{
	Term int
	LeaderId int
	PrevLogIndex int
	PreLogTerm int
	Entries []LogEntry
	LeaderCommit int
}
type AppendEntriesReply struct{
	Term int
	Success bool
}

func (rf *Raft) AppendEntries(args* AppendEntriesArgs, reply* AppendEntriesReply){
	//DPrintf("Invoke AppendEntries\n")
	curTerm := 0
	ifcan:=false
	rf.withLock(func(){
		curTerm = rf.currentTerm
		ifcan = (rf.role==candicate)
	})
	switch{
	case args.Term < curTerm:
		reply.Term = curTerm
		reply.Success=false
		return
	case args.Term > curTerm:
		rf.withLock( func(){ rf.resetOrNewElectTimer(randomTimeout())} )
		rf.withLock(func(){
			rf.currentTerm = args.Term
			rf.role = follower
			//rf.votedFor = args.CandicateId
			reply.Term = args.Term
			reply.Success = true
		})
		rf.newTerm <- 1
		return
	case args.Term == curTerm:
		rf.withLock( func(){ rf.resetOrNewElectTimer(randomTimeout())} )
		reply.Term = curTerm
		reply.Success = true
		if ifcan {rf.discoverLeader <- 1}
		return
	}
}
func (rf *Raft) sendAppendEntries(server int, args* AppendEntriesArgs, reply* AppendEntriesReply) bool{
	ok := rf.peers[server].Call("Raft.AppendEntries", args, reply)
	return ok
}
/**/



//
// the service using Raft (e.g. a k/v server) wants to start
// agreement on the next command to be appended to Raft's log. if this
// server isn't the leader, returns false. otherwise start the
// agreement and return immediately. there is no guarantee that this
// command will ever be committed to the Raft log, since the leader
// may fail or lose an election.
//
// the first return value is the index that the command will appear at
// if it's ever committed. the second return value is the current
// term. the third return value is true if this server believes it is
// the leader.
//
func (rf *Raft) Start(command interface{}) (int, int, bool) {
	index := -1
	term := -1
	isLeader := true

	// Your code here (2B).


	return index, term, isLeader
}

//
// the tester calls Kill() when a Raft instance won't
// be needed again. you are not required to do anything
// in Kill(), but it might be convenient to (for example)
// turn off debug output from this instance.
//
func (rf *Raft) Kill() {
	// Your code here, if desired.
}

//
// the service or tester wants to create a Raft server. the ports
// of all the Raft servers (including this one) are in peers[]. this
// server's port is peers[me]. all the servers' peers[] arrays
// have the same order. persister is a place for this server to
// save its persistent state, and also initially holds the most
// recent saved state, if any. applyCh is a channel on which the
// tester or service expects Raft to send ApplyMsg messages.
// Make() must return quickly, so it should start goroutines
// for any long-running work.
//
func Make(peers []*labrpc.ClientEnd, me int,
	persister *Persister, applyCh chan ApplyMsg) *Raft {
	rf := &Raft{}
	rf.peers = peers
	rf.persister = persister
	rf.me = me

	// Your initialization code here (2A, 2B, 2C).
	/**/
	rf.role = follower
	rf.currentTerm = 0
	rf.votedFor = -1
	rf.log = make([]LogEntry, 100)
	
	rf.commitIndex = 0
	rf.lastApplied = 0
	//reintialzied after election
	rf.nextIndex = make([]int,len(peers))
	rf.matchIndex =make([]int, len(peers))

	/**/
	// initialize from state persisted before a crash
	rf.readPersist(persister.ReadRaftState())
	
	rf.electTimer = time.NewTimer(randomTimeout())
	rf.discoverLeader = make(chan int)
	rf.newTerm = make(chan int)
	rf.winElect = make(chan int)

	//state machine
	go func(){
	for{
	DPrintf("%d into for{}\n", rf.me)
	start:

		/*
			check commit
		*/
		role := 0
		rf.withLock(func(){role=rf.role})
		switch role{
			//channel 用作转换条件合不合适，状态转移是瞬时的，channel有思索风险
			case follower:
				DPrintf("peer %d, ->follower\n", rf.me)
				for{
				select{
					case <-rf.newTerm:
						DPrintf("peer %d, follower: newTerm\n", rf.me)
						goto start
					case <-rf.electTimer.C:
						DPrintf("peer %d, follower: elec timeout\n", rf.me)
						rf.withLock(func(){rf.role=candicate})
						goto start
				}//end select
				}//end for
			case candicate:
				DPrintf("peer %d, ->candicate\n",rf.me)
				rf.withLock(func(){rf.currentTerm++})
				rf.withLock(func(){rf.votedFor = rf.me})
				rf.withLock(func(){rf.resetOrNewElectTimer(randomTimeout())} )
				//send requestVote
				votes := 1
				//var mux sync.Mutex
				for i := range rf.peers{
					if i == rf.me {
						continue
					}
					go func(server int){
						curTerm := 0
						cand := 0
						rf.withLock(func(){
							curTerm = rf.currentTerm
							cand	= rf.me
						})
						args := &RequestVoteArgs{Term:curTerm, CandicateId:cand}
						reply := RequestVoteReply{}
						for ok:=false; !ok;{
							//发送期间，rf的term变了怎么办：
							//ans:这种情况说明term落后了，仍用最初的term发出即可，请求会被拒绝
							//DPrintf("peer %d, sendRequestVote\n", rf.me)
							ok = rf.sendRequestVote(server, args, &reply)
							//DPrintf("peer %d, sendRequestVote return %v\n", rf.me, ok)
							//RPC 的行为？阻塞？重传？
							//if !ok {
							//	time.Sleep(fixedTimeout())
							//}
						}
						//再读边界情况 student tutorial
						//等到rpc响应的时候，role已经变了，怎么办？
						//ans:目前来看，没有影响，重新选举时候是新的vote变量副本，旧副本上的vote修改不影响	
						/*curTerm := 0
						curTerm = 0
						equal := false
						rf.withLock(func(){
							curTerm = rf.currentTerm
							equal = reply.Term== rf.currentTerm
						})
*/
						rf.withLock(func(){ curTerm = rf.currentTerm })
						//收到请求时，term可能已经变了，不是当前term的投票没有意义
						switch{
						case reply.Term < curTerm:
							//do what?
							return
						case reply.Term > curTerm:
							rf.withLock(func(){
								rf.currentTerm = reply.Term
								rf.votedFor = -1
								rf.role = follower
							})
							rf.newTerm <- 1
							return
						//是否换成选举成功channel，channel会不会清不干净？
						case reply.Term==curTerm && reply.VoteGranted:
							win := false
							rf.withLock(func(){
								votes++
								win = (votes > int(len(rf.peers)/2))
							})
							if win {
								rf.winElect <- 1
							}
						}
					}(i)
				}
				//transition
				for{
				select{
					case <-rf.newTerm:
						DPrintf("peer %d, candicate: newTerm\n", rf.me)
						goto start
					case <-rf.winElect:
						DPrintf("peer %d, candicate: voted sucess\n", rf.me)
						rf.withLock(func(){rf.role=leader})
						goto start
					case <-rf.discoverLeader:
						DPrintf("peer %d, candicate: disc leader\n", rf.me)
						rf.withLock(func(){rf.role=follower})
						goto start
					case <-rf.electTimer.C:
						DPrintf("peer %d, candicate: elect timeout\n", rf.me)
						goto start
				}//end select
				}//end for
			case leader:
				DPrintf("peer %d, ->leader\n",rf.me)
				/*
					initial empty heartbeat?
				*/
				for i := range rf.peers{
					if i==rf.me {
						continue
					}
					go func(server int){
						curTerm :=0
						leaderId :=0
						rf.withLock(func(){
							curTerm = rf.currentTerm
							leaderId = rf.me
						})
						args := &AppendEntriesArgs{Term:curTerm, LeaderId:leaderId}
						reply := AppendEntriesReply{}
						ok := rf.sendAppendEntries(server, args, &reply)
						if ok {
							curTerm := 0
							rf.withLock(func(){curTerm = rf.currentTerm})
							if reply.Term > curTerm{
								rf.withLock(func(){
									rf.currentTerm = reply.Term
									rf.role = follower
									rf.votedFor = -1
								})
								rf.newTerm <- 1
								//timeout是否合适？
								rf.withLock( func(){ rf.resetOrNewElectTimer(randomTimeout())} )
								return
							}
						}
					}(i)
				}

				for{
				select{
					//case <- command:
					case <-rf.newTerm:
						DPrintf("peer %d, leader: new Term\n", rf.me)
						goto start
					default:
						/*
						for{
							switch{
								case cond1
								case cond2
								default:
									headtbeat
							}
						}
						*/
						//heartbeat
						for i := range rf.peers{
							if i==rf.me {
								continue
							}
							go func(server int){
								curTerm :=0
								leaderId :=0
								rf.withLock(func(){
									curTerm = rf.currentTerm
									leaderId = rf.me
								})
								args := &AppendEntriesArgs{Term:curTerm, LeaderId:leaderId}
								reply := AppendEntriesReply{}
								ok := rf.sendAppendEntries(server, args, &reply)
								if ok {
									curTerm := 0
									rf.withLock(func(){curTerm = rf.currentTerm})
									if reply.Term > curTerm{
										rf.withLock(func(){
											rf.currentTerm = reply.Term
											rf.role = follower
											rf.votedFor = -1
										})
										rf.newTerm <- 1
										//在这里重置timer是否合适？
										rf.withLock( func(){ rf.resetOrNewElectTimer(randomTimeout())} )
										return
									}
								}
							}(i)
						}
						//不加sleep会发送过多消息
						time.Sleep(fixedTimeout())
				}//end select
				}//end for
		}
	}
	}()

	return rf
}


//
//timeout settings
//
//RPC timeout是多少
const(
	election_timeout = 300
	time_interval = 300
	fixed_timeout = 20
)

func randomTimeout() time.Duration {
	return time.Duration(election_timeout+rand.Intn(time_interval)) * time.Millisecond 
}

func fixedTimeout() time.Duration {
	return time.Duration(fixed_timeout) * time.Millisecond 
}

//
//helpful func
//
/*
func (rf *Raft) safeSetRole(role int){
	rf.mu.Lock()
	defer rf.mu.Unlock()
	rf.role = role
}

func (rf *Raft) safeGetRole() int{
	rf.mu.Lock()
	defer rf.mu.Unlock()
	return rf.role
}
*/
func (rf *Raft) resetOrNewElectTimer(d time.Duration) {
	//may have bug here, re-look timer api carefully
	//DPrintf("peer %d, resetTimer %v\n", rf.me, d)
	rf.electTimer.Stop()
	rf.electTimer.Reset(d)
}
/*
func (rf *Raft) safeIncTerm(){
	rf.mu.Lock()
	defer rf.mu.Unlock()
	rf.currentTerm++
}

func (rf *Raft) safeUpdateTerm(newTerm int){
	rf.mu.Lock()
	defer rf.mu.Unlock()
	rf.currentTerm = newTerm
}

func (rf* Raft) safeVotedFor(peer int){
	rf.mu.Lock()
	defer rf.mu.Unlock()
	rf.votedFor = peer 
}
*/
func (rf* Raft) withLock(f func()){
	rf.mu.Lock()
	defer rf.mu.Unlock()
	f()
}
