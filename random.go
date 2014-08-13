package main

import "fmt"
import "flag"
import "math/rand"

func main() {
  var opt_no_alpha bool
  var opt_no_digit bool
  var opt_no_unsco bool
  var opt_length   int
  var opt_num      int

  flag.BoolVar(&opt_no_alpha, "A", false, "Do not use any alpabets")
  flag.BoolVar(&opt_no_digit, "D", false, "Do not use any digits")
  flag.BoolVar(&opt_no_unsco, "U", false, "Do not use any digits")
  flag.IntVar(&opt_length, "l", 12, "Length of generated strings")
  flag.IntVar(&opt_num   , "n", 5,  "Number of generated strings")
  flag.Parse()

  m := make(map[byte]bool)
  for i := 0; i < flag.NArg(); i++ {
    str := flag.Arg(i)
    for j := 0; j < len(str); j++ {
      m[str[j]] = true
    }
  }

  if !opt_no_alpha {
    for c := 'a';  c <= 'z';  c++ {
      m[byte(c)] = true
    }
    for c := 'A';  c <= 'Z';  c++ {
      m[byte(c)] = true
    }
  }

  if !opt_no_digit {
    for c := '0';  c <= '9';  c++ {
      m[byte(c)] = true
    }
  }

  if !opt_no_unsco {
    m[byte('_')] = true
  }

  chars := make([]byte, 0, len(m))
  for c := range m {
    chars = append(chars, c)
  }

  for i := 0; i < opt_num;  i++ {
    for j := 0; j < opt_length; j++ {
      fmt.Printf("%c", chars[rand.Intn(len(chars))])
    }
    fmt.Println();
  }
}
