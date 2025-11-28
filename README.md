# dxafile - Updated Usage Examples

## New Features Added

### 1. Comprehensive Help System
The `--help` flag now provides thorough documentation for new users.

### 2. Dry-Run Mode
The `--dry-run` flag analyzes files without conversion, showing:
- File type detected
- Number of records
- What the output filename would be

---

## Help System Examples

### Example 1: Show Full Help
```bash
dxafile --help
# or
dxafile -h
# or just
dxafile
```

**Output:**
```
dxafile - DEXA Scanner File Converter

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

For more information, visit: https://github.com/derickschaefer/dxafile
```

---

## Dry-Run Mode Examples

### Example 2: Analyze Body Composition File
```bash
dxafile patient_bodycomp.txt --dry-run
```

**Output:**
```
File Analysis:
  Input File:   patient_bodycomp.txt
  Format Type:  Body Composition (Fat Mass/Percentage)
  Record Count: 15
  Output Would: patient_bodycomp.txt.json
```

### Example 3: Analyze with CSV Format
```bash
dxafile totalbody_scan.txt -f csv --dry-run
```

**Output:**
```
File Analysis:
  Input File:   totalbody_scan.txt
  Format Type:  Total Body (BMD Measurements)
  Record Count: 8
  Output Would: totalbody_scan.txt.csv
```

### Example 4: Analyze with Custom Output Path
```bash
dxafile vat_data.txt -f csv -o results/vat_analysis.csv -d
```

**Output:**
```
File Analysis:
  Input File:   vat_data.txt
  Format Type:  Core Scan (VAT Measurements)
  Record Count: 23
  Output Would: results/vat_analysis.csv
```

---

## Short Flag Examples

### Example 5: Using Short Flags
All flags now have short versions for faster typing:

```bash
# Long form
dxafile scan.txt --format=csv --output=result.csv

# Short form (equivalent)
dxafile scan.txt -f csv -o result.csv

# Dry-run short form
dxafile scan.txt -d

# Help short form
dxafile -h
```

---

## Enhanced Success Messages

### Example 6: Conversion with Record Count
```bash
dxafile patient_001.txt -f json
```

**Output:**
```
Successfully converted 12 records
Output file: /absolute/path/to/patient_001.txt.json
```

### Example 7: CSV Conversion
```bash
dxafile bodycomp.txt -f csv -o analysis/data.csv
```

**Output:**
```
Successfully converted 45 records
Output file: /absolute/path/to/analysis/data.csv
```

---

## Workflow Integration Examples

### Example 8: Pre-Check Before Batch Processing
```bash
# First, analyze all files to see what you're working with
for file in scans/*.txt; do
    echo "=== $file ==="
    dxafile "$file" --dry-run
    echo ""
done
```

**Output:**
```
=== scans/patient_001.txt ===
File Analysis:
  Input File:   scans/patient_001.txt
  Format Type:  Body Composition (Fat Mass/Percentage)
  Record Count: 12
  Output Would: scans/patient_001.txt.json

=== scans/patient_002.txt ===
File Analysis:
  Input File:   scans/patient_002.txt
  Format Type:  Total Body (BMD Measurements)
  Record Count: 8
  Output Would: scans/patient_002.txt.json

...
```

### Example 9: Conditional Processing Based on Record Count
```bash
#!/bin/bash
# Only convert files with more than 5 records

for file in *.txt; do
    # Get the dry-run output
    output=$(dxafile "$file" -d)
    
    # Extract record count (this is a simplified example)
    count=$(echo "$output" | grep "Record Count:" | awk '{print $3}')
    
    if [ "$count" -gt 5 ]; then
        echo "Converting $file ($count records)..."
        dxafile "$file" -f csv
    else
        echo "Skipping $file (only $count records)"
    fi
done
```

### Example 10: Audit Before Processing
```bash
# Create audit report of all files before conversion
echo "DEXA File Audit Report" > audit.txt
echo "Generated: $(date)" >> audit.txt
echo "==================" >> audit.txt
echo "" >> audit.txt

for file in data/*.txt; do
    echo "File: $file" >> audit.txt
    dxafile "$file" -d | grep -E "(Format Type|Record Count)" >> audit.txt
    echo "" >> audit.txt
done

echo "Audit complete. Review audit.txt before proceeding with conversions."
```

---

## Error Handling Examples

### Example 11: No Input File
```bash
dxafile
```

**Output:**
```
dxafile - DEXA Scanner File Converter

DESCRIPTION:
    Converts DEXA (Dual-Energy X-ray Absorptiometry) scanner export files from 
    UTF-16 LE BOM format into standard JSON or CSV formats.
    
[... full help displayed ...]
```

### Example 12: Wrong Number of Files
```bash
dxafile file1.txt file2.txt
```

**Output:**
```
Error: Expected exactly one input file
Usage: dxafile <file> [options]
Run 'dxafile --help' for more information
```

### Example 13: Invalid Format
```bash
dxafile scan.txt -f xml
```

**Output:**
```
Error: Invalid format 'xml'. Use 'json' or 'csv'
```

### Example 14: File Not Found
```bash
dxafile nonexistent.txt --dry-run
```

**Output:**
```
Error opening input file: open nonexistent.txt: no such file or directory
```

---

## Advanced Use Cases

### Example 15: Dry-Run with Piping
```bash
# Extract just the record counts from multiple files
for file in *.txt; do
    dxafile "$file" -d | grep "Record Count"
done | sort -t: -k2 -n
```

### Example 16: Generate Summary Report
```bash
#!/bin/bash
# Generate summary of all DEXA files in directory

echo "File Name,Type,Records" > summary.csv

for file in *.txt; do
    info=$(dxafile "$file" -d 2>/dev/null)
    
    if [ $? -eq 0 ]; then
        type=$(echo "$info" | grep "Format Type:" | cut -d: -f2 | xargs)
        count=$(echo "$info" | grep "Record Count:" | awk '{print $3}')
        echo "\"$file\",\"$type\",$count" >> summary.csv
    fi
done

echo "Summary report generated: summary.csv"
```

### Example 17: Quality Control - Find Small Files
```bash
# Find files with suspiciously few records (possible errors)
echo "Files with fewer than 5 records:"
for file in data/*.txt; do
    count=$(dxafile "$file" -d 2>/dev/null | grep "Record Count:" | awk '{print $3}')
    if [ ! -z "$count" ] && [ "$count" -lt 5 ]; then
        echo "  $file: $count records"
    fi
done
```

---

## Key Improvements Summary

1. **Comprehensive Help** (`-h` or `--help`)
   - Detailed description of tool purpose
   - All options explained
   - Multiple usage examples
   - File format detection rules
   - Input/output format specifications

2. **Dry-Run Mode** (`-d` or `--dry-run`)
   - Validates file can be parsed
   - Shows detected format type
   - Displays record count
   - Shows intended output path
   - No files written to disk

3. **Short Flags**
   - `-f` for `--format`
   - `-o` for `--output`
   - `-d` for `--dry-run`
   - `-h` for `--help`

4. **Better Error Messages**
   - Clear error descriptions
   - Helpful usage hints
   - Suggestions to use `--help`

5. **Enhanced Success Output**
   - Shows record count after conversion
   - Displays full output path
   - More informative than before

These improvements make the tool much more user-friendly for new users while maintaining full functionality for experienced users.
