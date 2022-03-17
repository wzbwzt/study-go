//go:build ingore
// +build ingore

package main

func aaa() (done func(), err error) {
	return func() {
		print("aaa: done")
	}, nil
}

func bbb() (done func(), _ error) {
	done, err := aaa()
	return func() {
		print("bbb: surprise!")
		done() //导致递归执行，变成了一个递归函数
	}, err
}

func main() {
	done, _ := bbb()
	done()
}
