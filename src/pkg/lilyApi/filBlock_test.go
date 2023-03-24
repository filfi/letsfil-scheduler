package lilyApi

import "testing"

func TestGetFilBlockReward(t *testing.T) {
	// t.Log(GetFilBlockReward("f01000", 100))
	minerId := "t01130"
	height := 169336
	ret, err := GetFilBlockReward(minerId, uint64(height))
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)

}
