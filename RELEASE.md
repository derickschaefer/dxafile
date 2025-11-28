# dxafile v2.0 - Usability Release: Friendly Column Names

## What's New

Version 2.0 introduces **human-readable column names** in CSV output, making data analysis much easier without needing a reference guide.

---

## Changes Overview

### Before (v1.x):
```csv
id1,id2,id3,date,value_0,value_1,value_2,mass_0_total,mass_0_left,...
```

### After (v2.0):
```csv
Last_Name,First_Name,Patient_ID,Measure_Date,Head_BMD,Arms_BMD,Legs_BMD,Arms_Bone_Mass_Total,Arms_Bone_Mass_Left,...
```

---

## Detailed Changes by File Type

### 1. Total Body Format

**Old format:**
- `id1, id2, id3, date, value_0, value_1, value_2, ...`

**New format:**
- `Last_Name, First_Name, Patient_ID, Measure_Date, Head_BMD, Arms_BMD, Legs_BMD, ...`

**Column name examples:**
- `Head_BMD`, `Arms_BMD`, `Legs_BMD` (Bone Mineral Density)
- `Head_BMC`, `Arms_BMC`, `Legs_BMC` (Bone Mineral Content)
- `Head_Area`, `Arms_Area`, `Legs_Area` (Scan areas)
- `Head_T_Score`, `Arms_T_Score` (T-scores for osteoporosis assessment)
- `Head_Z_Score`, `Arms_Z_Score` (Age-matched Z-scores)
- `Head_Average_Height`, `Arms_Average_Height` (Bone thickness)
- `Head_Average_Width`, `Arms_Average_Width` (Scan width)

**Total columns:** 119 named columns covering all body regions and measurement types

---

### 2. Body Composition Format

**Old format:**
- `id1, id2, id3, date, mass_0_total, mass_0_left, mass_0_right, mass_0_delta, ...`

**New format:**
- `Last_Name, First_Name, Patient_ID, Measure_Date, Arms_Bone_Mass_Total, Arms_Bone_Mass_Left, Arms_Bone_Mass_Right, Arms_Bone_Mass_Delta, ...`

**Mass measurement categories (each with Total/Left/Right/Delta):**
- **Bone Mass:** `Arms_Bone_Mass_*`, `Legs_Bone_Mass_*`, `Trunk_Bone_Mass_*`, `Android_Bone_Mass_*`, `Gynoid_Bone_Mass_*`, `Total_Bone_Mass_*`, `TBLH_Bone_Mass_*`
- **Fat Mass:** `Arms_Fat_Mass_*`, `Legs_Fat_Mass_*`, `Trunk_Fat_Mass_*`, `Android_Fat_Mass_*`, `Gynoid_Fat_Mass_*`, `Total_Fat_Mass_*`, `TBLH_Fat_Mass_*`
- **Lean Mass:** `Arms_Lean_Mass_*`, `Legs_Lean_Mass_*`, `Trunk_Lean_Mass_*`, etc.
- **Tissue Mass:** `Arms_Tissue_Mass_*`, `Legs_Tissue_Mass_*`, `Trunk_Tissue_Mass_*`, etc.
- **Fat-Free Mass:** `Arms_Fat_Free_Mass_*`, `Legs_Fat_Free_Mass_*`, `Trunk_Fat_Free_Mass_*`, etc.
- **Total Mass:** `Arms_Total_Mass_*`, `Legs_Total_Mass_*`, `Trunk_Total_Mass_*`, etc.

**Percentage measurement categories (each with Total/Left/Right/Delta):**
- **Region %Fat:** `Arms_Region_Percent_Fat_*`, `Legs_Region_Percent_Fat_*`, `Trunk_Region_Percent_Fat_*`, etc.
- **Tissue %Fat:** `Arms_Tissue_Percent_Fat_*`, `Legs_Tissue_Percent_Fat_*`, `Trunk_Tissue_Percent_Fat_*`, etc.

**Pattern example:**
```
Arms_Fat_Mass_Total    # Total fat mass in both arms
Arms_Fat_Mass_Left     # Fat mass in left arm
Arms_Fat_Mass_Right    # Fat mass in right arm
Arms_Fat_Mass_Delta    # Difference (Left - Right) indicating asymmetry
```

---

### 3. Core Scan Format

**Old format:**
- `id1, id2, id3, date, vat_mass_lbs, vat_volume_in3`

**New format:**
- `Last_Name, First_Name, Patient_ID, Measure_Date, VAT_Mass_lbs, VAT_Volume_in3`

**Changes:**
- Base columns now have friendly names
- VAT columns capitalized for consistency

---

## Naming Conventions

### Standard Patterns:

1. **Underscores replace spaces:** `Arms Fat Mass` → `Arms_Fat_Mass`
2. **Special characters replaced:** `Region %Fat` → `Region_Percent_Fat`
3. **Standard abbreviations kept:** `BMD`, `BMC`, `VAT`, `TBLH`
4. **Left/Right/Total/Delta suffixes:** Consistently applied across all measurements
5. **Capitalization:** Each word capitalized for readability

### Common Abbreviations:

| Abbreviation | Meaning |
|--------------|---------|
| BMD | Bone Mineral Density |
| BMC | Bone Mineral Content |
| VAT | Visceral Adipose Tissue |
| TBLH | Total Body Less Head |
| T_Score | Comparison to young adult mean |
| Z_Score | Comparison to age-matched mean |

### Anatomical Regions:

| Region | Description |
|--------|-------------|
| Arms | Both arms combined |
| Legs | Both legs combined |
| Trunk | Torso/chest/abdomen |
| Android | Upper abdomen (apple-shaped fat) |
| Gynoid | Hips and thighs (pear-shaped fat) |
| Total | Entire body |
| TBLH | Total body excluding head |

---

## Benefits

### 1. **Self-Documenting**
No need to consult a reference guide to understand what each column contains.

### 2. **Excel/Google Sheets Friendly**
```
=AVERAGE(Arms_Fat_Mass_Total:Legs_Fat_Mass_Total)
```
Much clearer than:
```
=AVERAGE(mass_7_total:mass_8_total)
```

### 3. **Python/Pandas Friendly**
```python
df = pd.read_csv('bodycomp.csv')

# Clear and readable
arms_fat = df['Arms_Fat_Mass_Total']
asymmetry = df['Arms_Fat_Mass_Delta']

# Easy filtering
high_fat = df[df['Total_Region_Percent_Fat_Total'] > 30]
```

### 4. **R Friendly**
```r
df <- read.csv('bodycomp.csv')

# Clear column references
arms_fat <- df$Arms_Fat_Mass_Total
total_bmd <- df$Total_BMD

# Easy to plot
plot(df$Measure_Date, df$Total_Region_Percent_Fat_Total)
```

### 5. **SQL Import Friendly**
Column names are valid SQL identifiers without needing quotes or escaping.

---

## Migration Guide

### If you have existing scripts using v1.x format:

#### Python/Pandas:
```python
# Old v1.x code:
arms_fat = df['mass_7_total']

# New v2.0 code:
arms_fat = df['Arms_Fat_Mass_Total']

# Or use column position (still works):
arms_fat = df.iloc[:, 11]  # 12th column (after id1,id2,id3,date + 7 mass groups)
```

#### Excel Formulas:
```
Old: =AVERAGE(E2:E100)  # had to remember E is value_0
New: =AVERAGE(Head_BMD)  # self-documenting
```

#### R:
```r
# Old v1.x code:
arms_fat <- df$mass_7_total

# New v2.0 code:
arms_fat <- df$Arms_Fat_Mass_Total
```

---

## Backward Compatibility

### JSON Output
**No changes** - JSON output still uses the original field names from the data structures:
- `"id1"`, `"id2"`, `"id3"`, `"date"`
- `"mass"`, `"percent"`
- `"values"`

JSON structure remains stable for existing integrations.

### CSV Output
**Breaking change** - Column names have changed. 

If you need the old format, you can:
1. Keep using v1.x release
2. Use column positions instead of names
3. Create a mapping file to translate old→new names

---

## Examples

### Example 1: Body Composition CSV Header
```csv
Last_Name,First_Name,Patient_ID,Measure_Date,
Arms_Bone_Mass_Total,Arms_Bone_Mass_Left,Arms_Bone_Mass_Right,Arms_Bone_Mass_Delta,
Legs_Bone_Mass_Total,Legs_Bone_Mass_Left,Legs_Bone_Mass_Right,Legs_Bone_Mass_Delta,
Trunk_Bone_Mass_Total,Trunk_Bone_Mass_Left,Trunk_Bone_Mass_Right,Trunk_Bone_Mass_Delta,
Android_Bone_Mass_Total,Android_Bone_Mass_Left,Android_Bone_Mass_Right,Android_Bone_Mass_Delta,
...
```

### Example 2: Total Body CSV Header
```csv
Last_Name,First_Name,Patient_ID,Measure_Date,
Head_BMD,Arms_BMD,Legs_BMD,Trunk_BMD,Ribs_BMD,Pelvis_BMD,Spine_BMD,
Arm_Left_BMD,Leg_Left_BMD,Arm_Right_BMD,Leg_Right_BMD,Total_BMD,TBLH_BMD,
...
```

### Example 3: Core Scan CSV Header
```csv
Last_Name,First_Name,Patient_ID,Measure_Date,VAT_Mass_lbs,VAT_Volume_in3
```

---

## Real-World Usage

### Analyzing Body Composition in Python:
```python
import pandas as pd
import matplotlib.pyplot as plt

# Load data with friendly column names
df = pd.read_csv('bodycomp.csv')

# Calculate total body fat percentage over time
plt.plot(pd.to_datetime(df['Measure_Date']), 
         df['Total_Region_Percent_Fat_Total'])
plt.xlabel('Date')
plt.ylabel('Body Fat %')
plt.title('Body Fat Percentage Trend')
plt.show()

# Identify asymmetries
asymmetry_threshold = 0.5  # lbs
asymmetric_arms = df[abs(df['Arms_Fat_Mass_Delta']) > asymmetry_threshold]
print(f"Found {len(asymmetric_arms)} scans with arm asymmetry > {asymmetry_threshold} lbs")
```

### Analyzing BMD in R:
```r
library(ggplot2)

df <- read.csv('totalbody.csv')

# Plot BMD trends
ggplot(df, aes(x=as.Date(Measure_Date, "%m/%d/%Y"), y=Total_BMD)) +
  geom_line() +
  geom_point() +
  labs(title="Bone Mineral Density Over Time",
       x="Date", y="Total BMD (g/cm²)")

# Check for osteoporosis risk
at_risk <- df[df$Total_T_Score < -2.5, ]
print(paste("Number of scans indicating osteoporosis:", nrow(at_risk)))
```

### Excel PivotTable:
1. Import CSV into Excel
2. Insert → PivotTable
3. Drag `Patient_ID` to Rows
4. Drag `Total_Region_Percent_Fat_Total` to Values (Average)
5. Drag `Measure_Date` to Columns
6. Instant analysis of body fat trends by patient!

---

## Technical Details

### Implementation:
- **Body Composition:** 42 measurement categories × 4 columns each = 168 potential columns
- **Total Body:** 119 measurement columns (7 categories × 17 body regions)
- **Core Scan:** 2 measurement columns

### Code changes:
- Updated `writeCSVBodyComp()` in `output.go`
- Updated `writeCSVTotalBody()` in `output.go`
- Updated `writeCSVCoreScan()` in `output.go`
- Added `sanitizeColumnName()` helper function (for future use)

---

## Testing Recommendations

After upgrading to v2.0:

1. **Convert a test file:**
   ```bash
   dxafile test_scan.txt -f csv -o test_output.csv
   ```

2. **Verify column names:**
   ```bash
   head -1 test_output.csv
   ```

3. **Compare record counts:**
   ```bash
   # Should have same number of data rows as v1.x
   wc -l test_output.csv
   ```

4. **Update analysis scripts** to use new column names

5. **Test your analysis pipeline** end-to-end

---

## Upgrade Path

### Immediate Upgrade (Recommended):
```bash
# Back up existing v1.x binary
mv dxafile dxafile_v1_backup

# Install v2.0
# ... download/build new binary ...

# Test with sample file
dxafile sample.txt -f csv -o test.csv
head -1 test.csv  # verify friendly names
```

### Gradual Migration:
```bash
# Keep both versions
mv dxafile dxafile_v2
# Keep dxafile_v1_backup for old scripts

# Use v2 for new analysis
dxafile_v2 newscan.txt -f csv

# Use v1 for legacy scripts (until updated)
dxafile_v1_backup oldscan.txt -f csv
```

---

## Questions & Support

**Q: Can I get the old column names back?**  
A: Not directly, but you can use v1.x release for old format, or use column positions instead of names.

**Q: Does this affect JSON output?**  
A: No, JSON output is unchanged.

**Q: What about the --dry-run output?**  
A: No changes - dry-run still shows file type and record count.

**Q: Are there performance impacts?**  
A: No performance difference - only column naming changed.

**Q: Can I customize column names?**  
A: Not currently, but you can request this feature on GitHub.

---

## Release Notes

**Version:** 2.0.0  
**Release Date:** November 2025  
**Type:** Major release (breaking changes in CSV format)

**Breaking Changes:**
- CSV column names changed from generic (value_N, mass_N) to descriptive names
- Base columns renamed: id1→Last_Name, id2→First_Name, id3→Patient_ID, date→Measure_Date

**Non-Breaking:**
- JSON output unchanged
- CLI interface unchanged
- File detection unchanged
- --dry-run output unchanged

**Additions:**
- 119 friendly column names for Total Body format
- 168+ friendly column names for Body Composition format
- Improved column names for Core Scan format

---

**Upgrade today for better data analysis workflows!**
