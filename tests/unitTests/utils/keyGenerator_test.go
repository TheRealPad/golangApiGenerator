package utils

import (
	"httpServer/src/utils"
	"testing"
)

func TestKeyGenerator_1(t *testing.T) {
	size := 10
	key, err := utils.GenerateKey(size)
	if err != nil {
		t.Errorf("Error after generating key")
	}
	if len(key) != size {
		t.Errorf("Bad key generation, wrong size - ")
	}
}

func TestKeyGenerator_2(t *testing.T) {
	size := 100
	key, err := utils.GenerateKey(size)
	if err != nil {
		t.Errorf("Error after generating key")
	}
	if len(key) != size {
		t.Errorf("Bad key generation, wrong size - ")
	}
}

func TestKeyGenerator_3(t *testing.T) {
	size := 42
	key, err := utils.GenerateKey(size)
	if err != nil {
		t.Errorf("Error after generating key")
	}
	if len(key) != size {
		t.Errorf("Bad key generation, wrong size - ")
	}
}
