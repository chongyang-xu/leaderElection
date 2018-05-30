package raft

//
// Raft tests.
//
// we will use the original test_test.go to test your code for grading.
// so, while you can modify this code to help you debug, please
// test with the original before submitting.
//

import "testing"
import "fmt"
import "time"

// The tester generously allows solutions to complete elections in one second
// (much more than the paper's range of timeouts).
const RaftElectionTimeout = 1000 * time.Millisecond

func TestReElection(t *testing.T) {
	servers := 3
	cfg := make_config(t, servers, false)
	defer cfg.cleanup()

	fmt.Printf("Test (2A): election after network failure ...\n")
	DPrintf("[TESTCASE][Initial Election]\n")
	leader1 := cfg.checkOneLeader()

	// if the leader disconnects, a new one should be elected.
	cfg.disconnect(leader1)
	DPrintf("[TESTCASE][Disconnect %d]\n", leader1)
	cfg.checkOneLeader()

	// if the old leader rejoins, that shouldn't
	// disturb the old leader.
	cfg.connect(leader1)
	DPrintf("[TESTCASE][Connect %d]\n", leader1)
	leader2 := cfg.checkOneLeader()

	// if there's no quorum, no leader should
	// be elected.
	cfg.disconnect(leader2)
	cfg.disconnect((leader2 + 1) % servers)
	DPrintf("[TESTCASE][Disconnect %d,%d]\n", leader2, (leader2+1)%servers)
	time.Sleep(2 * RaftElectionTimeout)
	cfg.checkNoLeader()

	// if a quorum arises, it should elect a leader.
	cfg.connect((leader2 + 1) % servers)
	DPrintf("[TESTCASE][Connect %d]\n", (leader2+1) % servers)
	cfg.checkOneLeader()

	// re-join of last node shouldn't prevent leader from existing.
	cfg.connect(leader2)
	DPrintf("[TESTCASE][Connect %d]\n", leader2)
	cfg.checkOneLeader()

	fmt.Printf("  ... Passed\n")
}

