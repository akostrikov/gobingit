package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	content := "what is up, doc?"
	var store bytes.Buffer

	store.WriteString("blob ")
	store.WriteString(strconv.Itoa(len(content)))
	store.WriteString(string('\u0000'))
	store.WriteString(content)
	fmt.Printf("Store string is: %s\n", store.String())

	h := sha1.New()
	io.WriteString(h, store.String())
	shaValue := h.Sum(nil)
	fmt.Printf("Hex print: %x\n", h.Sum(nil))

	src := shaValue
	hexValue := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(hexValue, src)

	fmt.Printf("Hex encode: %s\n", hexValue)

	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(store.Bytes())
	w.Close()
	fmt.Println(b.String())

	path := ".git/objects/" + string(hexValue[0:2]) + "/" + string(hexValue[2:len(hexValue)])
	fmt.Println(path)
	dir, _ := filepath.Split(path)
	err := os.MkdirAll(dir, os.ModePerm) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(path, b.Bytes(), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	//git write-tree
	//c82b154409d60ff285aacd55ff340c0fbb0901d2
	//git cat-file -p c82b154409d60ff285aacd55ff340c0fbb0901d2
	//100644 blob 4de9dff31397e3881f074ea8ba5b62fa6d23fbd8	.gitignore

	//echo 'first commit' | git commit-tree c82b154409d60ff285aacd55ff340c0fbb0901d2
	//519b6efacb56a32930a7b22dafe903aea8a76114
	//git cat-file -p 519b6efacb56a32930a7b22dafe903aea8a76114
	//tree c82b154409d60ff285aacd55ff340c0fbb0901d2
	//author Alexandr Kostrikov <alexandr.kostrikov@gmail.com> 1493995910 +0300
	//committer Alexandr Kostrikov <alexandr.kostrikov@gmail.com> 1493995910 +0300
	//
	//first commit

	//git log
	//fatal: your current branch 'master' does not have any commits yet
	//echo "519b6efacb56a32930a7b22dafe903aea8a76114" > .git/refs/heads/master
	//git log
	//commit 519b6efacb56a32930a7b22dafe903aea8a76114
	//Author: Alexandr Kostrikov <alexandr.kostrikov@gmail.com>
	//Date:   Fri May 5 17:51:50 2017 +0300
	//
	//first commit
}
