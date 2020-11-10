package tailer

import (
	"testing"
	"os"
	"fmt"
)

const testfilename = "/tmp/followtest"
const tstring ="Test string\n"

func TestTailer1(t *testing.T) {
	tfile,_:=os.Create(testfilename)
	defer func() {
		tfile.Close()
		os.Remove(testfilename)
	}()
	fl:=NewFollower(testfilename)
	fmt.Fprintf(tfile,tstring)
	readstring:=fl.Tail()
	if readstring!=tstring {
		t.Errorf("Read back failed: read %q want %q",readstring,tstring)
	}
	ch:=make(chan int)
	go func() {
		ch<-0
		readstring=fl.Tail()
		if readstring!=tstring {
			t.Errorf("Read back failed: read %q want %q",readstring,tstring)
		}
	}()
	<-ch
	fmt.Fprintf(tfile,tstring)
}

func TestTailer2(t *testing.T) {
	fl:=NewFollower(testfilename)
	tfile,_:=os.Create(testfilename)
	defer func() {
		tfile.Close()
		os.Remove(testfilename)
	}()
	fmt.Fprintf(tfile,tstring )
	readstring:=fl.Tail()
	if readstring!=tstring {
		t.Errorf("Read back failed: read %q want %q",readstring,tstring)
	}
}

func TestTailer3(t *testing.T) {
	fl:=NewFollower(testfilename)
	tfile,_:=os.Create(testfilename)
	defer func() {
		tfile.Close()
		os.Remove(testfilename)
	}()
	fmt.Fprintf(tfile,tstring )
	readstring:=fl.Tail()
	if readstring!=tstring {
		t.Errorf("Read back failed: read %q want %q",readstring,tstring)
	}
	tfile.Close()
	tfile,_=os.Create(testfilename)
	ch:=make(chan int)
	go func() {
		ch<-0
		readstring=fl.Tail()
		if readstring!=tstring {
			t.Errorf("Read back failed: read %q want %q",readstring,tstring)
		}
	}()
	<-ch
	fmt.Fprintf(tfile,tstring)
}
