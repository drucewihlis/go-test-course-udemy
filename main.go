// go run .
package main

import (
    "bufio"
    "fmt"
    "os"
    "io"
    "strconv"
    "strings"
)

func main() {
    // print a welcome
    intro()

    // create a channel to indicate when the user wants to quit
    doneChan := make(chan bool)

    // start a goroutine to read user input and run program
    go readUserInput(os.Stdin, doneChan)

    // block until doneChan gets a value
    <-doneChan

    // close the channel
    close(doneChan)

    // say goodbye
    fmt.Println("Goodbye.")
}

func readUserInput(in io.Reader, doneChan chan bool) {
    scanner := bufio.NewScanner(in)

    for {
        res, done := checkNumbers(scanner)

        if done {
            doneChan <- true
            return
        }

        fmt.Println(res)
        prompt()
    }
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
    // read user input
    scanner.Scan()

    // check to see if user wants to quit
    if strings.EqualFold(scanner.Text(), "q") {
        return "", true
    }

    // try to convert what user typed into int
    numToCheck, err := strconv.Atoi(scanner.Text())
    if err != nil {
        return "Please enter a whole number!", false
    }

    _, msg := isPrime(numToCheck)

    return msg, false
}

func intro() {
    fmt.Println("Is it Prime?")
    fmt.Println("------------")
    fmt.Println("Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit")
    prompt()
}

func prompt() {
    fmt.Print("-> ")
}

func isPrime(n int) (bool, string) {
    // 0 and 1 are not prime by def
    if n == 0 || n == 1 {
        return false, fmt.Sprintf("%d is not prime, by definition!", n)
    }

    // negs are not prime
    if n < 0 {
        return false, "negs are not prime by definition!"
    }

    // use modulus repeatedly to see if we have prime
    for i := 2 ; i <= n/2; i++ {
        if n%i == 0 {
            // not prime
            return false, fmt.Sprintf("%d is not prime, cause it is divisible by %d!", n, i)
        }
    }

    return true, fmt.Sprintf("%d is a prime!", n)
}
