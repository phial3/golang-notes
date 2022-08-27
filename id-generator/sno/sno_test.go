package sno_test

import (
	"fmt"
	"testing"
	"time"
)

import (
	"github.com/muyo/sno"
)

func TestGenerateId(t *testing.T) {
	snoId := sno.New(0)
	snoId1 := sno.New(10)
	snoId2 := sno.New(100)
	snoId3 := sno.New(255)
	fmt.Printf("snoId : %v\n", snoId)
	fmt.Printf("snoId1 : %v\n", snoId1)
	fmt.Printf("snoId2 : %v\n", snoId2)
	fmt.Printf("snoId3 : %v\n", snoId3)
}

func TestGenerateIdWithTime(t *testing.T) {
	snoId := sno.NewWithTime(0, time.Now())
	fmt.Printf("snoId : %v\n", snoId)
}
