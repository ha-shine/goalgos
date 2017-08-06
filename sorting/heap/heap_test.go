package heap

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ha-shine/goalgos/util"
)

func TestBenchmark(t *testing.T) {
	items := []string{}
	consume := func(s string) {
		s = strings.TrimSpace(s)
		if s != "" {
			items = append(items, strings.TrimSpace(s))
		}
	}
	util.ReadFile("../../tale.txt", consume)
	Sort(items)
	fmt.Println(items[:10])
}
