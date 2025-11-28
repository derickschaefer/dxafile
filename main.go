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
    var dryRun bool
    var help bool

    pflag.StringVarP(&format, "format", "f", "json", "Output format: json or csv")
    pflag.StringVarP(&output, "output", "o", "", "Output file path (default: <input>.<format>)")
    pflag.BoolVarP(&dryRun, "dry-run", "d", false, "Show file type and record count without converting")
    pflag.BoolVarP(&help, "help", "h", false, "Show detailed help information")
    pflag.Parse()

    // Show help if requested or no arguments provided
    if help || pflag.NArg() == 0 {
        showHelp()
        os.Exit(0)
    }

    // Validate argument count
    if pflag.NArg() != 1 {
        fmt.Println("Error: Expected exactly one input file")
        fmt.Println("Usage: dxafile <file> [options]")
        fmt.Println("Run 'dxafile --help' for more information")
        os.Exit(1)
    }

    inputFile := pflag.Arg(0)

    // Validate format
    switch format {
    case "json", "csv":
        // Valid format
    default:
        fmt.Printf("Error: Invalid format '%s'. Use 'json' or 'csv'\n", format)
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

    // Get record count
    recordCount := getRecordCount(records)

    // If dry-run, show info and exit
    if dryRun {
        fmt.Println("File Analysis:")
        fmt.Printf("  Input File:   %s\n", inputFile)
        fmt.Printf("  Format Type:  %s\n", formatTypeName(dxaType))
        fmt.Printf("  Record Count: %d\n", recordCount)
        fmt.Printf("  Output Would: %s\n", output)
        os.Exit(0)
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
    fmt.Printf("Successfully converted %d records\n", recordCount)
    fmt.Printf("Output file: %s\n", absOut)
}

// showHelp displays comprehensive usage information
func showHelp() {
    fmt.Println(`dxafile - DEXA Scanner File Converter

DESCRIPTION:
    Converts DEXA (Dual-Energy X-ray Absorptiometry) scanner export files from 
    UTF-16 LE BOM format into standard JSON or CSV formats.
    
    Automatically detects and handles three DEXA format types:
      • Body Composition - Fat mass/percentage by body region
      • Total Body       - Bone mineral density (BMD) measurements  
      • Core Scan        - Visceral adipose tissue (VAT) measurements

USAGE:
    dxafile <input_file> [options]

OPTIONS:
    -f, --format <type>     Output format: json or csv (default: json)
    -o, --output <path>     Output file path (default: <input>.<format>)
    -d, --dry-run           Analyze file without converting (shows type & count)
    -h, --help              Show this help message

EXAMPLES:
    # Convert to JSON (auto-named output)
    dxafile scan_data.txt

    # Convert to CSV with custom output path
    dxafile scan_data.txt --format=csv --output=results/data.csv

    # Short flags work too
    dxafile scan_data.txt -f csv -o output.csv

    # Analyze file without converting
    dxafile scan_data.txt --dry-run

    # Batch convert all .txt files to CSV
    for file in *.txt; do dxafile "$file" -f csv; done

FILE FORMAT DETECTION:
    The tool automatically detects the DEXA format by examining the header:
      • Body Composition: Header contains "arms fat mass"
      • Total Body:       Header contains "head bmd"
      • Core Scan:        Header contains "vat mass"

INPUT FORMAT:
    • Encoding: UTF-16 Little Endian with BOM
    • Structure: Tab-delimited text
    • Common ID fields: ID1, ID2, ID3, Date
    • Data fields: Vary by DEXA format type

OUTPUT FORMATS:
    JSON: Pretty-printed with 2-space indentation
    CSV:  Headers included, format-specific column layout

NOTES:
    • Input files are not modified (read-only)
    • Output files are overwritten if they already exist
    • Empty lines in input are automatically skipped
    • Malformed lines generate descriptive error messages

For more information, visit: https://github.com/derickschaefer/dxafile`)
}

// formatTypeName returns a human-readable name for the DXA type
func formatTypeName(t DXAType) string {
    switch t {
    case DXATypeBodyComp:
        return "Body Composition (Fat Mass/Percentage)"
    case DXATypeTotalBody:
        return "Total Body (BMD Measurements)"
    case DXATypeCoreScan:
        return "Core Scan (VAT Measurements)"
    default:
        return "Unknown"
    }
}

// getRecordCount returns the number of records in the parsed data
func getRecordCount(records interface{}) int {
    switch r := records.(type) {
    case []BodyFatRecord:
        return len(r)
    case []TotalBodyRecord:
        return len(r)
    case []CoreScanRecord:
        return len(r)
    default:
        return 0
    }
}
