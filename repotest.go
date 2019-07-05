package main

import (
	"fmt"
	"io"
	//"strings"
	"time"
	"bytes"
	"io/ioutil"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/filemode"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/format/index"
)


func main(){


	
	repo, err := git.PlainOpen(".")
    if err != nil {
        panic(err)
    }

	w, err := repo.Worktree()
	if err != nil {
		panic(err)
	}
	
	idx, err := repo.Storer.Index()
	if err != nil {
		panic(err)
	}

	var  fileName string = "repotest.go"
	var  fileType uint32 = 0100664
	
// git hash-object ...
	//reader := strings.NewReader("last test of encapsulation??")
	//byteArr, err := ioutil.ReadFile("testCad.dwg")
	//reader := bytes.NewReader(byteArr)
	// django 에서 go 서버로 오는건 byte 배열 ...
	
	byteArr, err := ioutil.ReadFile(fileName)
	reader := bytes.NewReader(byteArr)
	
	obj :=  repo.Storer.NewEncodedObject()
	obj.SetType(plumbing.BlobObject)
	obj.SetSize(int64(reader.Len()))

	writer, err := obj.Writer()
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(writer, reader); err != nil {
		panic(err)
	}

	
	h, err :=  repo.Storer.SetEncodedObject(obj)
	if err != nil {
		panic(err)
	}
//end

// 현재 index 에 있는 걸 github 창에서 띄워준다고 봐도 무방함 ...
//idx.Remove("testCad.dwg")
//idx.Remove("testCad.dxf")
// git update-index 파일명을 인자로 받아야함..

e, err := idx.Entry(fileName)

if err == index.ErrEntryNotFound {
	e = idx.Add(fileName)	
}




e.Hash = h
e.ModifiedAt = time.Now()
//e.Mode, err = filemode.NewFromOSFileMode(info.Mode())
//클라이언트에서 파일모드 보내주는 걸로 ....
e.Mode = filemode.FileMode(fileType)

if e.Mode == 100644 {
	e.Size = uint32(int64(reader.Len()))
}




repo.Storer.SetIndex(idx)
//end




	commit, err := w.Commit("this is a last file test", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "hosuk",
			Email: "kirklayer@gmail.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		panic(err)
	}

	commitObj, err := repo.CommitObject(commit)
	if err != nil {
		panic(err)
	}
	fmt.Println(commitObj)
	
	fmt.Printf("hello \n")
	
}








