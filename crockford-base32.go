package main

import (
	"bufio"
	"encoding/base32"
	"flag"
	"fmt"
	"io"
	"os"
)

const crockfordAlphabet = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

func main() {
    decodeFlag := flag.Bool("d", false, "Decode input")
    lineLengthFlag := flag.Int("l", 64, "Line length (0 for no line breaks)")
    flag.Parse()

    if *decodeFlag {
        decode(os.Stdin, os.Stdout)
    } else {
        encode(os.Stdin, os.Stdout, *lineLengthFlag)
    }
}

func encode(input io.Reader, output io.Writer, lineLength int) {
    data, err := io.ReadAll(input)
    if err != nil {
        fmt.Println("Error reading input:", err)
        os.Exit(1)
    }
    encoded := base32.NewEncoding(crockfordAlphabet).WithPadding(base32.NoPadding).EncodeToString(data)
    if lineLength > 0 {
        for i := 0; i < len(encoded); i += lineLength {
            j := i + lineLength
            if j > len(encoded) {
                j = len(encoded)
            }
            fmt.Fprintln(output, encoded[i:j])
        }
    } else {
        fmt.Fprintln(output, encoded)
    }
}

func decode(input io.Reader, output io.Writer) {
    encoded := ""
    scanner := bufio.NewScanner(input)
    for scanner.Scan() {
        encoded += scanner.Text()
    }
    decoded, err := base32.NewEncoding(crockfordAlphabet).WithPadding(base32.NoPadding).DecodeString(encoded)
    if err != nil {
        fmt.Println("Error decoding input:", err)
        os.Exit(1)
    }
    fmt.Fprintln(output, string(decoded))
    if scanner.Err() != nil {
        fmt.Println("Error reading input:", scanner.Err())
        os.Exit(1)
    }
}
