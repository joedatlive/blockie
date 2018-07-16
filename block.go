// a fuction which takes input and creates a block

package main

import "time"
import "fmt"
import "crypto/SHA256"
import "encoding/hex"
import "google.golang.org/appengine" // Required external App Enginer lib

func timeStamp() time.Time {
	ts := time.Now().UTC()
	return time.Time(ts)
}

func hasher(index int, previousHash []byte, timeStamp time.Time, data string) string {
	
	//hash this block's data
	h := sha256.New()
	h.Write([]byte(data))
	//convert to string
	sha256_hash := hex.EncodeToString(h.Sum(nil))
	//return string of hash in hex of this block's data
	return string(sha256_hash)
}



func main() {
	//set variables
	ts := timeStamp()
	previousHash := []byte("ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad")
	data := "abc"
	index := 1
	
	// pass in this block's data to hasher and retrun hash
	block := hasher(index, previousHash, ts, data)
	
	//print block data
	fmt.Println("Block: ", block)
	fmt.Println("Block data")
	fmt.Println("Index: ", index)
	fmt.Println("Previous hash: ", previousHash)
	fmt.Println("Timestamp: ", ts.Format(time.RFC3339))
	fmt.Println("Data: ", data)
	
}

//func Block (string h) int index, time.Time ts, string data, string previousHash