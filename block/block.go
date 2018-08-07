package main

import "time"
import "fmt"
import "crypto/sha256"
import "encoding/hex"
import "net/http"
import "google.golang.org/appengine"
//import "google.golang.org/appengine/datastore"
//import "log"
import "encoding/json"

type Block struct {
	Hash string
	Index int
	Timestamp time.Time
	PreviousHash []byte
	Data string
}

type Data struct {
	NodeName string
	Owner string
	Account string
	Package string
	Zone string
}

//http handler below
func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			//we will enumerate blocks here
			fmt.Fprintln(w,"This will return a list of blocks.  Someday.")
		case "POST":
			//we will create a block upon POST
			ts := timeStamp()
			
			//these are temporary data
			previousHash := []byte("ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad")
			index := 1
			
			//get the data from the for POST
			//Call to ParseForm makes form fields available.
			err := r.ParseForm()
			if err != nil {
				// Handle error here via logging and then return
				return            
			}

			n := r.Form.Get("nodeName")
			o := r.Form.Get("owner")
			a := r.Form.Get("account")
			p := r.Form.Get("package")
			z := r.Form.Get("zone")

			// Put URL values into a data struct
			data := Data{
				NodeName: n,
				Owner: o,
				Account: a,
				Package: p,
				Zone: z,
			} 
			
			//convert to json
			d,err := json.Marshal(data)
			if err != nil {
				fmt.Println(err)
				return
			}
			dstr :=string(d)

			// pass in this block's data to hasher and return hash
			blockhash := hasher(index, previousHash, ts, dstr)

			//create a block with the data and needed meta-data	
			block := Block{
				Hash: blockhash,
				Index:index,
				Timestamp: ts,
				PreviousHash: previousHash,
				Data: dstr,
			}

			//put the block in json
			b,err := json.Marshal(block)
			if err != nil {
				fmt.Println(err)
				return
			}
			//lets try to connect to Google Cloud Storage Bucket
		/*	ctx := appengine.NewContext(r)
			key := datastore.NewIncompleteKey(ctx, "Block", nil)

			if _, err := datastore.Put(ctx, key, &b); err != nil {
				log.Fatalf("datastore.Put: %v", err)
				return
		}
		*/

			// Send the block info back to the web client for the user
			// we will convert this to read from the data store...
			fmt.Fprintln(w, string(b))
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}	
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
	

	//starts web server 
	http.HandleFunc("/block", indexHandler) // set endpoint handler
		appengine.Main() // starts the server to receive requests
	
}
