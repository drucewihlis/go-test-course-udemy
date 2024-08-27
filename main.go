// go run .
package main

import "fmt"

func main() {
    n := 3

    _, msg := isPrime(n)
    fmt.Println(msg)
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
