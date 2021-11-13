package sub

import "fmt"

var Subhoge1 int = 1000
var subhoge2 int = 2000

func Add(atai1 int, atai2 int) int {
	return atai1 + atai2 + A
}

// strings.Join関数のコードを抜粋
func Join(a []string, sep string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}
	n := len(sep) * (len(a) - 1)
	fmt.Printf("(1) len(sep):%v len(a):%v n:%v\n", len(sep), len(a), n)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
		fmt.Printf("(2) %v \n", a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	fmt.Printf("(3.0) %v [%v]\n", bp, string(b))
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
		fmt.Printf("(3) %v \n", bp)
	}
	return string(b)
}
