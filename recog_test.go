package imgo

import (
	"testing"
)

func TestCosineSimilarity(t *testing.T) {
	cos, err := CosineSimilarity("testpic/aa.png", "testpic/cc.png")

	if err != nil {
		t.Error(err)
	}

	t.Log(cos)
	if cos > 0.5 {
		t.Fail()
	}
}
