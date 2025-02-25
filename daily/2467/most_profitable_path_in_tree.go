package _467

import "math"

/*
	There is an undirected tree with n nodes labeled from 0 to n - 1, rooted at node 0. You are given a 2D integer array edges of length n - 1 where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.

	At every node i, there is a gate. You are also given an array of even integers amount, where amount[i] represents:

	the price needed to open the gate at node i, if amount[i] is negative, or,
	the cash reward obtained on opening the gate at node i, otherwise.
	The game goes on as follows:

	Initially, Alice is at node 0 and Bob is at node bob.
	At every second, Alice and Bob each move to an adjacent node. Alice moves towards some leaf node, while Bob moves towards node 0.
	For every node along their path, Alice and Bob either spend money to open the gate at that node, or accept the reward. Note that:
	If the gate is already open, no price will be required, nor will there be any cash reward.
	If Alice and Bob reach the node simultaneously, they share the price/reward for opening the gate there. In other words, if the price to open the gate is c, then both Alice and Bob pay c / 2 each. Similarly, if the reward at the gate is c, both of them receive c / 2 each.
	If Alice reaches a leaf node, she stops moving. Similarly, if Bob reaches node 0, he stops moving. Note that these events are independent of each other.
	Return the maximum net income Alice can have if she travels towards the optimal leaf node.

	Example 1

	Input: edges = [[0,1],[1,2],[1,3],[3,4]], bob = 3, amount = [-2,4,2,-4,6]
	Output: 6
	Explanation:
	The above diagram represents the given tree. The game goes as follows:
	- Alice is initially on node 0, Bob on node 3. They open the gates of their respective nodes.
	  Alice's net income is now -2.
	- Both Alice and Bob move to node 1.
	  Since they reach here simultaneously, they open the gate together and share the reward.
	  Alice's net income becomes -2 + (4 / 2) = 0.
	- Alice moves on to node 3. Since Bob already opened its gate, Alice's income remains unchanged.
	  Bob moves on to node 0, and stops moving.
	- Alice moves on to node 4 and opens the gate there. Her net income becomes 0 + 6 = 6.
	Now, neither Alice nor Bob can make any further moves, and the game ends.
	It is not possible for Alice to get a higher net income.

	URL: https://leetcode.com/problems/most-profitable-path-in-a-tree/description/?envType=daily-question&envId=2025-02-24

*/

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(amount)

	// Create adjacency list representation of the tree
	tree := make([][]int, n)
	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		tree[edge[1]] = append(tree[edge[1]], edge[0])
	}

	// Initialize distanceFromBob array with a large value
	distanceFromBob := make([]int, n)
	for i := range distanceFromBob {
		distanceFromBob[i] = n
	}

	return findPaths(0, -1, 0, bob, amount, tree, distanceFromBob)
}

// findPaths performs DFS to calculate the maximum profit
func findPaths(sourceNode, parentNode, time, bob int, amount []int, tree [][]int, distanceFromBob []int) int {
	maxIncome := 0
	maxChild := math.MinInt32

	// If Bob is at the current node, set distance to 0, otherwise set to max
	if sourceNode == bob {
		distanceFromBob[sourceNode] = 0
	} else {
		distanceFromBob[sourceNode] = len(amount)
	}

	// Traverse adjacent nodes
	for _, adjacentNode := range tree[sourceNode] {
		if adjacentNode != parentNode {
			maxChild = max(maxChild, findPaths(adjacentNode, sourceNode, time+1, bob, amount, tree, distanceFromBob))
			distanceFromBob[sourceNode] = min(distanceFromBob[sourceNode], distanceFromBob[adjacentNode]+1)
		}
	}

	// Alice reaches the node first
	if distanceFromBob[sourceNode] > time {
		maxIncome += amount[sourceNode]
	} else if distanceFromBob[sourceNode] == time { // Alice and Bob reach simultaneously
		maxIncome += amount[sourceNode] / 2
	}

	// If leaf node, return the max income
	if maxChild == math.MinInt32 {
		return maxIncome
	} else {
		return maxIncome + maxChild
	}
}
