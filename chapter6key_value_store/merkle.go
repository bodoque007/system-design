package main

import (
	"crypto/sha1"
)

type KV struct {
	Key string
	Value string
}


type Bucket struct {
	Entries []KV
}

type MerkleNode struct {
	Hash []byte
	Left *MerkleNode
	Right *MerkleNode
}


func hashBucket(bucket *Bucket) []byte {
	h := sha1.New()
	for _, kv := range bucket.Entries {
		h.Write([]byte(kv.Key + kv.Value))
	}
	return h.Sum(nil)
}

func buildMerkleTreeFromLeaves(leaves []*MerkleNode) *MerkleNode {
	if len(leaves) == 1 {
		return leaves[0];
	}
	var parents []*MerkleNode
	for i := 0; i < len(leaves); i += 2 {
		left := leaves[i]
		var right *MerkleNode
		if i+1 < len(leaves) {
			right = leaves[i+1]
		}
		h := sha1.New()
		h.Write(left.Hash)
		if right != nil {
			h.Write(right.Hash)
		}
		parents = append(parents, &MerkleNode{
			Hash: h.Sum(nil),
			Left: left, 
			Right: right,
		})
	}
	return buildMerkleTreeFromLeaves(parents)
}

func main() {
	buckets := []Bucket{
		{Entries: []KV{{"k1", "v1"}, {"k2", "v2"}}},
		{Entries: []KV{{"k3", "v3"}, {"k4", "v4"}}},
	}
	leaves := []*MerkleNode{}
	for _, bucket := range buckets {
		leaves = append(leaves, &MerkleNode{
			Hash: hashBucket(&bucket),
			Left: nil,
			Right: nil,
		})
	}
	merkleTree := buildMerkleTreeFromLeaves(leaves)

}