package xz

import (
	"crypto/sha256"
	"fmt"
	"os"
	"testing"
)

func TestRawReader(T *testing.T){
	f,err := os.Open("testdata/compressed.bin")
	if err != nil{
		T.Fatalf("Failed To Open file due to : %v\n",err)
		os.Exit(-1)
	}
	var opts [5]byte
	f.Read(opts[:])
	dec,err := NewReaderRaw(f,opts)
	if err!=nil{
		fmt.Printf("Errors Decompressing: %v\n",err)
		os.Exit(-1)
	}
	total:=0
	var allContents []byte
	for ;;{
		var data [2048]byte
		n, err := dec.Read(data[:])
		if err!=nil || n==0 {
			fmt.Println("Read All");
			break;
		}
		allContents=append(allContents,data[:]...)
		total+=n
	}
	expected:="58728f2bae65d1c03bebf4f4fd4c7baae662c347b63d35114cd639281757e432"
	computed:=fmt.Sprintf("%x",sha256.Sum256(allContents))
	if expected!=computed{
		T.Fatalf("Decompression Of Raw Stream Failed. Total size:%d. %s != %s",total,expected,computed)
	}
}
