#!/bin/bash

FILE="$1"

if [[ -z "$FILE" ]]; then
    echo "Usage: $0 <filename>"
    exit 1
fi

echo "============================================================"
echo " FILE INFO"
echo "============================================================"
file -i "$FILE"
echo

echo "============================================================"
echo " FIRST 5 LINES (literal with invisible characters)"
echo "============================================================"
sed -n '1,5p' "$FILE" | cat -A
echo

echo "============================================================"
echo " FIRST 5 LINES (tabs shown as [TAB])"
echo "============================================================"
sed -n '1,5p' "$FILE" | sed -e 's/\t/[TAB]/g'
echo

echo "============================================================"
echo " TOKEN COUNT PER EACH OF FIRST 5 LINES"
echo "============================================================"
for i in {1..5}; do
    line=$(sed -n "${i}p" "$FILE")
    count=$(echo "$line" | wc -w)
    echo "Line $i: $count tokens"
done
echo

echo "============================================================"
echo " FIRST DATA ROW SPLIT BY TABS (showing actual fields)"
echo "============================================================"
sed -n '2p' "$FILE" | tr '\t' '\n' | cat -n
echo

echo "============================================================"
echo " HEX DUMP (first 200 bytes of file)"
echo "============================================================"
xxd -g 1 -c 40 -n 200 "$FILE"
echo

echo "============================================================"
echo " HEX OF FIRST NON-HEADER DATA ROW"
echo "============================================================"
# detect second line (first row after header)
DATA_LINE=$(sed -n '2p' "$FILE")
echo "$DATA_LINE" | xxd -g 1 -c 40
echo

echo "============================================================"
echo " DONE "
echo "============================================================"
