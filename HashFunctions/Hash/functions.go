package hash
import (
	"crypto/sha256"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func ExampleNew(Message string) {
	h := sha256.New()
	h.Write([]byte(Message))
	fmt.Printf("%x", h.Sum(nil))
	
}

func SHA1() {
	data := []byte("This page intentionally left blank.")
	fmt.Printf("% x", sha1.Sum(data))
	// Output: af 06 49 23 bb f2 30 15 96 aa c4 c2 73 ba 32 17 8e bc 4a 96
}


func ExampleNew_file() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))
}