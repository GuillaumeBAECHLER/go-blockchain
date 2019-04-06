package blockchain

import (
	"strings"
	"testing"
	"time"
)

func TestBlock(t *testing.T) {

	var genesisBlock Block
	genesisBlock = Genesis()
	data := "Hello world !"
	newBlock := MineBlock(genesisBlock, data)

	cases := []struct {
		label, in, want string
	}{
		{"data", newBlock.Data, data},
		{"lastBlockHash", newBlock.LastBlockHash, genesisBlock.Hash},
		{"hash", newBlock.Hash[0:newBlock.Difficulty], strings.Repeat("0", newBlock.Difficulty)},
	}
	for _, c := range cases {
		if c.in != c.want {
			t.Errorf("%q = %q, expected %q", c.label, c.in, c.want)
		}
	}

	if newDifficulty := AdjustDifficulty(newBlock, time.Now().Add(time.Hour*1)); newDifficulty != (newBlock.Difficulty - 1) {
		t.Errorf("Difficulty should decrease if time is too long. Have %q, expected %q", newDifficulty, (newBlock.Difficulty - 1))
	}

	if newDifficulty = AdjustDifficulty(newBlock, time.Now().Add(time.Second*1)); newDifficulty != (newBlock.Difficulty + 1) {
		t.Errorf("Difficulty should increase if time is too short. Have %q, expected %q", newDifficulty, (newBlock.Difficulty + 1))
	}
}
