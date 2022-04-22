package main

import "os"

func main() {
	os.ReadFile()
	file, _ := os.Open()
	file.ReadAt()
	fs, _ := file.Stat()

}
