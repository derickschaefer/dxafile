package main

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"

    "github.com/spf13/pflag"
)

func main() {
    var format string
    var output string

    pflag.StringVar(&format, "format", "json", "Output format: json or csv")
    pflag.StringVar(&output, "output", "", "Output file (optional). If omitted, auto-named <input>.<format>")
    pflag.Parse()

    if pflag.NArg() != 1 {
        fmt.Println("Usage: kfile <file> --format=<json|csv> [--output=<file>]")
        os.Exit(1)
    }

    inputFile := pflag.Arg(0)

    // Validate format
    switch format {
    case "json", "csv":
    default:
        fmt.Println("Invalid --format value. Use json or csv.")
        os.Exit(1)
    }

    // Auto-name output if not provided
    if output == "" {
        ext := "." + format
        output = inputFile + ext
    }

    // Open input file
    in, err := os.Open(inputFile)
    if err != nil {
        fmt.Println("Error opening input file:", err)
        os.Exit(1)
    }
    defer in.Close()

    // Parse the file
    dxaType, records, err := ParseFile(in)
    if err != nil {
        fmt.Println("Error parsing file:", err)
        os.Exit(1)
    }

    // Create output file
    out, err := os.Create(output)
    if err != nil {
        fmt.Println("Error creating output file:", err)
        os.Exit(1)
    }
    defer out.Close()

    buf := bufio.NewWriter(out)

    // Write output depending on format
    switch format {
    case "json":
        err = OutputJSON(buf, dxaType, records)
    case "csv":
        err = OutputCSV(buf, dxaType, records)
    }

    if err != nil {
        fmt.Println("Error writing output:", err)
        os.Exit(1)
    }

    buf.Flush()

    absOut, _ := filepath.Abs(output)
    fmt.Println("Wrote:", absOut)
}
