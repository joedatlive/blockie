package main

import "time"
import "fmt"
import "crypto/sha256"
import "encoding/hex"
import "net/http"
import "google.golang.org/appengine"
import "encoding/json"

type Block struct {
	Hash string
	Index int
	Timestamp time.Time
	PreviousHash []byte
	Data string
}

//http handler below
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Best Blockie Ever!") // Response to request
  }

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
	
	// pass in this block's data to hasher and return hash
	blockhash := hasher(index, previousHash, ts, data)

	block := Block{
		Hash: blockhash,
		Index:index,
		Timestamp: ts,
		PreviousHash: previousHash,
		Data: data,
	}

	b,err := json.Marshal(block)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	//starts web server 
	http.HandleFunc("/blockweb", indexHandler) // set endpoint handler
		appengine.Main() // starts the server to receive requests
	
}
