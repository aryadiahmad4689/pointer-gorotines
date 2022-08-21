package main_test

import (
	"sync"
	"testing"
)

type User struct {
	S *sync.WaitGroup
	m sync.RWMutex
}

type Nilai struct {
	Nilai int
}

func (u *User) ChangeNilai1(chanel chan *Nilai) {
	chanel <- &Nilai{Nilai: 100}
}

func (u *User) ChangeNilai2(chanel chan *Nilai) {
	chanel <- &Nilai{Nilai: 200}
}

func (u *User) ChangeNilai3(chanel chan *Nilai) {
	chanel <- &Nilai{Nilai: 300}
}

func BenchmarkChange(t *testing.B) {
	// var wg = &sync.WaitGroup{}
	var chanel1 = make(chan *Nilai, 1)
	var u = &User{}
	go u.ChangeNilai1(chanel1)
	<-chanel1
	go u.ChangeNilai2(chanel1)
	<-chanel1
	go u.ChangeNilai3(chanel1)
	<-chanel1
	defer close(chanel1)
}
