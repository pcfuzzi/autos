package main

import "fmt"

func main() {
    s := "Hallo Teilnehmer."
    var vowels, consonants int
    for _, c := range s {
        switch c {
        case 'a', 'e', 'i', 'o', 'u':
            vowels++
        case ' ', '.':
            // Do nothing.
        default:
            consonants++
        }
    }
    fmt.Printf("Es sind Vokale: %d\n", vowels)
    fmt.Printf("Es sind Konsonaten: %d\n", consonants)
}
