package goConfig

import (
	"fmt"
	"testing"
)

type testSub struct {
	S1 int        `cfg:"field S1" cfgDefault:"1"`
	S2 int        `cfg:"field S2"`
	S3 string     `cfg:"field S3"`
	S4 testSubSub `cfg:"Sub Sub"`
}
type testSubSub struct {
	SS1 int    `cfg:"field SS1" cfgDefault:"2"`
	SS2 int    `cfg:"field SS2"`
	SS3 string `cfg:"field SS3"`
}
type testAux struct {
	A int     `cfg:"field a"`
	B string  `cfg:"field b"`
	S testSub `cfg:"Sub"`
}

func TestParseTags(t *testing.T) {
	s := &testAux{A: 1, S: testSub{S1: 1, S2: 2, S3: "test"}}
	err := parseTags(s)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", s)

	s1 := "test"
	err = parseTags(s1)
	if err == nil {
		t.Fatal("Error expected")
	}

	err = parseTags(&s1)
	if err == nil {
		t.Fatal("Error expected")
	}
}

/*
{
  "domain": "www.example.com",
  "mongodb": {
    "host": "localhost",
    "port": 27017
  }
*/

type mongoDB struct {
	Host string
	port int
}

type configTest struct {
	Domain  string
	MongoDB mongoDB
}

func TestLoad(t *testing.T) {
	config := configTest{}
	err := Load(&config)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("\n\nTestLoad: %#v\n\n", config)

}
