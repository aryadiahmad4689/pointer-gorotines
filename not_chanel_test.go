package main_test

import (
	"testing"
)

func (u *User) ChangeNilai11(n *Nilai) {

}

func (u *User) ChangeNilai22(n *Nilai) {
	n.Nilai += 100
}

func (u *User) ChangeNilai33(n *Nilai) {
	n.Nilai += 100
}

func BenchmarkChange1(t *testing.B) {
	// var wg = &sync.WaitGroup{}
	var u = &User{}
	nilai := &Nilai{Nilai: 100}
	u.ChangeNilai11(nilai)
	u.ChangeNilai22(nilai)
	u.ChangeNilai33(nilai)
	// fmt.Println("ini noroutine", nilai)

}
