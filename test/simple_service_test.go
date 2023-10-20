package test

import (
	"fmt"
	"memetpaten/go-rest/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService := simple.InitializeService()
	fmt.Println(simpleService.SimpleRepository)
}
