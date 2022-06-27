package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {

	command()
	lookpath()
	stdout()
	start()
	output()
}

func command() {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}

func lookpath() {
	path, err := exec.LookPath("sqlite3")
	if err != nil {
		log.Fatal("installing fortune is in your future")
	}
	fmt.Printf("fortune is available at %s\n", path)
}

func stdout() {
	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var person struct {
		Name string
		Age  int
	}
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
}

func start() {
	cmd := exec.Command("sleep", "5")
	err := cmd.Start() //并不会等待该命令完成即返回
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	//Wait会阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的。
	//Wait方法会在命令返回后释放相关的资源。
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func output() {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}
