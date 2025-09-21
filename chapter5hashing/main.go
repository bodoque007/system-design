package main

import (
	"fmt"
	"hash/fnv"
)

type BST struct {
	hash uint32
	server string
	left, right *BST
}

// Never used, just for the BST to be complete lol
func find(bst *BST, hash uint32) *BST {
	if bst == nil {
		return nil
	}
	if hash == bst.hash {
		return bst
	} else if bst.hash < hash {
		return find(bst.right, hash)
	} else {
		return find(bst.left, hash)
	}
}

func insert(bst *BST, server string) *BST {
	hash := hashKey(server)
	if bst == nil {
		return &BST{
			hash:   hash,
			server: server,
			left:   nil,
			right:  nil,
		}
	}
	
	if hash == bst.hash {
		// Kind of extreme edge case, in case we do find the hash, we just update the node with the new server.
		bst.server = server
		return bst
	} else if hash < bst.hash {
		bst.left = insert(bst.left, server)
	} else {
		bst.right = insert(bst.right, server)
	}
	
	return bst
}

// Find the server with hash greater than key closest to the key
func findServer(bst *BST, key string) *BST {
	h := hashKey(key)
	var candidate *BST 
	cur := bst
	for cur != nil {
		if cur.hash >= h {
			// current node is a candidate, but there could be another server with a key smaller than current yet still bigger than input key.
			candidate = cur
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	if candidate == nil {
		// If no server has a greater key, wrap around and return the smallest one (leftmost node in the BST)
		candidate = findLeftmost(bst)
	}
	return candidate
}


func findLeftmost(bst *BST) *BST {
	if bst == nil {
		return nil 
	}
	if bst.left == nil {
		return bst 
	}
	return findLeftmost(bst.left)
}

func delete(bst *BST, server string) *BST {
    if bst == nil {
        return nil
    }
    
    hash := hashKey(server)
    
    if hash < bst.hash {
        bst.left = delete(bst.left, server)
    } else if hash > bst.hash {
        bst.right = delete(bst.right, server)
    } else {
        // bst.hash == hash, found node
        if bst.left == nil {
            return bst.right
        } else if bst.right == nil {
            return bst.left
        } else {
            // two children: replace with inorder successor (could find the rightmost of the left child subtree too, both are correct and mantain invariant)
            successor := findLeftmost(bst.right)
            bst.hash = successor.hash
            bst.server = successor.server
            bst.right = delete(bst.right, successor.server)
        }
    }
    
    return bst
}

func hashKey(key string) uint32 {
    h := fnv.New32a()
    h.Write([]byte(key))
    return h.Sum32()
}

func main() {
	var root *BST
	servers := []string{"Server1", "Server2", "Server3", "Server4"}

	for _, s := range servers {
		root = insert(root, s)
	}

	keys := []string{"key1", "key2", "key3"}

	for _, k := range keys {
		server := findServer(root, k)
		fmt.Printf("%s maps to server %s\n", k, server.server)
	}

}