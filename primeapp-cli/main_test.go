// go test .
// go test -v -run Test_isPrime
// go test -v -run Test_alpha
// go test -cover .
// go test -coverprofile=coverage.out .
// go tool cover -html=coverage.out  // should open browser, does not work for me
// go tool cover -html=coverage.out -o coverage.html // generates html

// go test -v -coverprofile=coverage.out . && go tool cover -html=coverage.out -o coverage.html

package main

import (
    "bufio"
    "os"
    "testing"
    "io"
    "strings"
    "bytes"
)

func Test_alpha_isPrime(t *testing.T) { // aplha is arranging 2 tests into a test suite
    primeTests := []struct {
        name string
        testNum int
        expected bool
        msg string
    }{
        {"prime", 7, true, "7 is a prime!"},
        {"not prime", 8, false, "8 is not prime, cause it is divisible by 2!"},
        {"zero", 0, false, "0 is not prime, by definition!"},
        {"one", 1, false, "1 is not prime, by definition!"},
        {"negative", -1, false, "negs are not prime by definition!"},
    }

    for _, e := range primeTests {
        result, msg := isPrime(e.testNum)
        if e.expected && !result {
            t.Errorf("%s: expected true but got false", e.name)
        }

        if !e.expected && result {
            t.Errorf("%s: expected false but got true", e.name)
        }

        if e.msg != msg {
            t.Errorf("%s: expected `%s` but got `%s`", e.name, e.msg, msg)
        }
    }
}

func Test_alpha_prompt(t *testing.T) {
    // save a copy of os.Stdout
    oldOut := os.Stdout

    // create a read and write pipe
    r, w, _ := os.Pipe()

    // set os.Stdout to our write pipe
    os.Stdout = w
    prompt()

    // close our writer
    _ = w.Close()

    // reset os.Stdout to what it was before
    os.Stdout = oldOut

    // read the output of our prompt() func from our read pipe
    out, _ := io.ReadAll(r)

    // perform our test
    if string(out) != "-> " {
        t.Errorf("incorrect prompt: expected `-> ` but got %s", string(out))
    }
}

func Test_intro(t *testing.T) {
    // save a copy of os.Stdout
    oldOut := os.Stdout

    // create a read and write pipe
    r, w, _ := os.Pipe()

    // set os.Stdout to our write pipe
    os.Stdout = w
    intro()

    // close our writer
    _ = w.Close()

    // reset os.Stdout to what it was before
    os.Stdout = oldOut

    // read the output of our prompt() func from our read pipe
    out, _ := io.ReadAll(r)

    // perform our test
    if !strings.Contains(string(out), "Enter a whole number") {
        t.Errorf("intro text is not correct;got %s", string(out))
    }
}

func Test_checkNumbers(t *testing.T) {
    tests := []struct {
        name string
        input string
        expected string
    } {
        {name: "empty", input: "", expected: "Please enter a whole number!"},
        {name: "zero", input: "0", expected: "0 is not prime, by definition!"},
        {name: "one", input: "1", expected: "1 is not prime, by definition!"},
        {name: "two", input: "2", expected: "2 is a prime!"},
        {name: "negative", input: "-1", expected: "negs are not prime by definition!"},
        {name: "types", input: "three", expected: "Please enter a whole number!"},
        {name: "decimal", input: "1.1", expected: "Please enter a whole number!"},
        {name: "quit", input: "q", expected: ""},
        {name: "Q", input: "Q", expected: ""},
        {name: "slavic", input: "ёпта", expected: "Please enter a whole number!"},
    }

    for _, e := range tests {
        input := strings.NewReader(e.input)
        reader := bufio.NewScanner(input)
        res, _ := checkNumbers(reader)

        if !strings.EqualFold(res, e.expected) {
            t.Errorf("%s: expected %s, but got %s", e.name, e.expected, res)
        }
    }
}

func Test_readUserInput(t *testing.T) {
       // to test this func, we need a channel, and an instance of an io.NewReader
       doneChan := make(chan bool)

       // create a ref to a bytes.Buffer
       var stdin bytes.Buffer

       stdin.Write([]byte("1\nq\n"))

       go readUserInput(&stdin, doneChan)
       <-doneChan
       close(doneChan)
}
