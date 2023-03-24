package commonUtils

import (
	"fmt"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/bigInt"
	"testing"
)

func TestCommon(t *testing.T) {

	fmt.Println(CastToString(bigInt.MustFromString("1234567890123456789012345678901234567890")))
	fmt.Println(CastToString(2.4, 3))
	fmt.Println(CastToString(float32(2.2), 2))
	fmt.Println(CastToString(bigInt.MustFromString("1234567890123456789012345678901234567890")))
	fmt.Println(CastToString(int8(2)))

	s := bigInt.MustFromString("3098832384303375").Add(bigInt.MustFromString("3098832384303375").Int, bigInt.MustFromString("1087").Int)
	fmt.Println(s)
	fmt.Println(s.Sub(s, bigInt.MustFromString("106041580380816075").Int))
}
