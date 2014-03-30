package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    if len(os.Args) == 2 {
        PrintFileToStdout(os.Args[1])
        os.Exit(0)
    } else if len(os.Args) == 1 {
        scanner := bufio.NewScanner(os.Stdin)

        for scanner.Scan() {
            fmt.Println(scanner.Text())
        }
        if err := scanner.Err(); err != nil {
            fmt.Fprintln(os.Stderr, "error:", err)
            os.Exit(1)
        }
    } else if len(os.Args) > 2 {
       for i := 1; i < len(os.Args); i++ {
           PrintFileToStdout(os.Args[i])
       }
    }
}

func PrintFileToStdout(fileName string) {
        file, err := os.Open(fileName)
        if err != nil {
            os.Exit(1)
        }
        defer file.Close()

        stat, err := file.Stat()
        if err != nil {
            return
        }

        bs := make([]byte, stat.Size())
        _, err = file.Read(bs)
        if err != nil {
            return
        }

        str := string(bs)
        fmt.Print(str)
}
