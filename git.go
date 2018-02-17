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

	var tree_content bytes.Buffer
	tree_content.WriteString("100644 .gitignore")
	tree_content.WriteString(string('\u0000'))
	tree_content.Write(shaValue)
	
	var tree_store bytes.Buffer
	tree_store.WriteString("tree ")
	tree_store.WriteString(strconv.Itoa(tree_content.Len()))
	tree_store.WriteString(string('\u0000'))
	tree_store.Write(tree_content.Bytes())

	h2 := sha1.New()
	io.WriteString(h2, tree_store.String())
	shaValue2 := h2.Sum(nil)

	src2 := shaValue2
	hexValue2 := make([]byte, hex.EncodedLen(len(src2)))
	hex.Encode(hexValue2, src2)

	tree_path := ".git/objects/" + string(hexValue2[0:2]) + "/" + string(hexValue2[2:len(hexValue)])

	var b2 bytes.Buffer
	w2 := zlib.NewWriter(&b2)
	w2.Write(tree_store.Bytes())
	w2.Close()
	fmt.Println(b2.String())

	fmt.Println(tree_path)
	dir2, _ := filepath.Split(tree_path)
	err2 := os.MkdirAll(dir2, os.ModePerm) // For read access.
	if err2 != nil {log.Fatal(err2)}

	err = ioutil.WriteFile(tree_path, b2.Bytes(), os.ModePerm)
	if err != nil {log.Fatal(err)}

    commit_multiline := "tree 6398107bf9b91a96e55e90959994958705325d06\n"+
	"author Alexandr Kostrikov <alexandr.kostrikov@gmail.com> 1494000784 +0300\n" +
	"committer Alexandr Kostrikov <alexandr.kostrikov@gmail.com> 1494000784 +0300\n" +
	"\n"  +"commit message\n"
	var commit_store bytes.Buffer
	commit_store.WriteString("commit ")
	commit_store.WriteString(strconv.Itoa(len(commit_multiline)))
	commit_store.WriteString(string('\u0000'))
	commit_store.WriteString(commit_multiline)

	h3 := sha1.New()
	io.WriteString(h3, commit_store.String())
	shaValue3 := h3.Sum(nil)

	src3 := shaValue3
	hexValue3 := make([]byte, hex.EncodedLen(len(src3)))
	hex.Encode(hexValue3, src3)

	commit_path := ".git/objects/" + string(hexValue3[0:2]) + "/" + string(hexValue3[2:len(hexValue3)])

	fmt.Println(commit_path)

	var b3 bytes.Buffer
	w3 := zlib.NewWriter(&b3)
	w3.Write(commit_store.Bytes())
	w3.Close()
	fmt.Println(b3.String())

	dir3, _ := filepath.Split(commit_path)
	err3 := os.MkdirAll(dir3, os.ModePerm)
	if err3 != nil {log.Fatal(err3)}

	err = ioutil.WriteFile(commit_path, b3.Bytes(), os.ModePerm)
	if err != nil {log.Fatal(err)}

	err = ioutil.WriteFile(".git/refs/heads/master", hexValue3, os.ModePerm)
	if err != nil {log.Fatal(err)}
	//sha1 = Digest::SHA1.hexdigest(tree_store)

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
