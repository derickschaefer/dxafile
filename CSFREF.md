# DEXA File Converter - CSV Output Reference Guide

## Overview
This document explains the CSV output structure for each of the three DEXA scanner file types. The converter automatically detects the file type and generates appropriate CSV columns.

---

## Common Fields (All File Types)

All three file types share these first four columns:

| Column | Description |
|--------|-------------|
| `id1` | Last Name or Primary Patient Identifier |
| `id2` | First Name or Secondary Identifier |
| `id3` | Patient ID or Tertiary Identifier |
| `date` | Scan/Measurement Date |

---

## 1. TOTAL BODY FORMAT

**Detected by:** Header contains "head bmd"  
**File Type:** Total Body (BMD Measurements)  
**Columns:** id1, id2, id3, date, value_0, value_1, value_2, ... value_N

### Column Index Reference

The numeric indices (value_0, value_1, etc.) correspond to these measurements in order:

#### BMD (Bone Mineral Density) - values_0 through value_16
| Index | Measurement | Description |
|-------|-------------|-------------|
| value_0 | Head BMD | Bone mineral density of head |
| value_1 | Arms BMD | Bone mineral density of both arms |
| value_2 | Legs BMD | Bone mineral density of both legs |
| value_3 | Trunk BMD | Bone mineral density of trunk |
| value_4 | Ribs BMD | Bone mineral density of ribs |
| value_5 | Pelvis BMD | Bone mineral density of pelvis |
| value_6 | Spine BMD | Bone mineral density of spine |
| value_7 | Arm Left BMD | Bone mineral density of left arm |
| value_8 | Leg Left BMD | Bone mineral density of left leg |
| value_9 | Arm Right BMD | Bone mineral density of right arm |
| value_10 | Leg Right BMD | Bone mineral density of right leg |
| value_11 | Total BMD | Total body bone mineral density |
| value_12 | TBLH BMD | Total body less head bone mineral density |
| value_13 | Trunk Left BMD | Bone mineral density of left trunk |
| value_14 | Total Left BMD | Total left side bone mineral density |
| value_15 | Trunk Right BMD | Bone mineral density of right trunk |
| value_16 | Total Right BMD | Total right side bone mineral density |

#### BMC (Bone Mineral Content) - values_17 through value_33
| Index | Measurement | Description |
|-------|-------------|-------------|
| value_17 | Head BMC | Bone mineral content of head (grams) |
| value_18 | Arms BMC | Bone mineral content of both arms |
| value_19 | Legs BMC | Bone mineral content of both legs |
| value_20 | Trunk BMC | Bone mineral content of trunk |
| value_21 | Ribs BMC | Bone mineral content of ribs |
| value_22 | Pelvis BMC | Bone mineral content of pelvis |
| value_23 | Spine BMC | Bone mineral content of spine |
| value_24 | Arm Left BMC | Bone mineral content of left arm |
| value_25 | Leg Left BMC | Bone mineral content of left leg |
| value_26 | Arm Right BMC | Bone mineral content of right arm |
| value_27 | Leg Right BMC | Bone mineral content of right leg |
| value_28 | Total BMC | Total body bone mineral content |
| value_29 | TBLH BMC | Total body less head bone mineral content |
| value_30 | Trunk Left BMC | Bone mineral content of left trunk |
| value_31 | Total Left BMC | Total left side bone mineral content |
| value_32 | Trunk Right BMC | Bone mineral content of right trunk |
| value_33 | Total Right BMC | Total right side bone mineral content |

#### Area Measurements - values_34 through value_50
| Index | Measurement | Description |
|-------|-------------|-------------|
| value_34 | Head Area | Scan area of head (cm²) |
| value_35 | Arms Area | Scan area of both arms |
| value_36 | Legs Area | Scan area of both legs |
| value_37 | Trunk Area | Scan area of trunk |
| value_38 | Ribs Area | Scan area of ribs |
| value_39 | Pelvis Area | Scan area of pelvis |
| value_40 | Spine Area | Scan area of spine |
| value_41 | Arm Left Area | Scan area of left arm |
| value_42 | Leg Left Area | Scan area of left leg |
| value_43 | Arm Right Area | Scan area of right arm |
| value_44 | Leg Right Area | Scan area of right leg |
| value_45 | Total Area | Total body scan area |
| value_46 | TBLH Area | Total body less head scan area |
| value_47 | Trunk Left Area | Scan area of left trunk |
| value_48 | Total Left Area | Total left side scan area |
| value_49 | Trunk Right Area | Scan area of right trunk |
| value_50 | Total Right Area | Total right side scan area |

#### T-Scores - values_51 through value_67
| Index | Measurement | Description |
|-------|-------------|-------------|
| value_51 | Head T-Score | Standard deviations from young adult mean (head) |
| value_52 | Arms T-Score | Standard deviations from young adult mean (arms) |
| value_53 | Legs T-Score | Standard deviations from young adult mean (legs) |
| value_54 | Trunk T-Score | Standard deviations from young adult mean (trunk) |
| value_55 | Ribs T-Score | Standard deviations from young adult mean (ribs) |
| value_56 | Pelvis T-Score | Standard deviations from young adult mean (pelvis) |
| value_57 | Spine T-Score | Standard deviations from young adult mean (spine) |
| value_58 | Arm Left T-Score | Standard deviations from young adult mean (left arm) |
| value_59 | Leg Left T-Score | Standard deviations from young adult mean (left leg) |
| value_60 | Arm Right T-Score | Standard deviations from young adult mean (right arm) |
| value_61 | Leg Right T-Score | Standard deviations from young adult mean (right leg) |
| value_62 | Total T-Score | Standard deviations from young adult mean (total body) |
| value_63 | TBLH T-Score | Standard deviations from young adult mean (TBLH) |
| value_64 | Trunk Left T-Score | Standard deviations from young adult mean (left trunk) |
| value_65 | Total Left T-Score | Standard deviations from young adult mean (total left) |
| value_66 | Trunk Right T-Score | Standard deviations from young adult mean (right trunk) |
| value_67 | Total Right T-Score | Standard deviations from young adult mean (total right) |

#### Z-Scores - values_68 through value_84
| Index | Measurement | Description |
|-------|-------------|-------------|
| value_68 | Head Z-Score | Standard deviations from age-matched mean (head) |
| value_69 | Arms Z-Score | Standard deviations from age-matched mean (arms) |
| value_70 | Legs Z-Score | Standard deviations from age-matched mean (legs) |
| value_71 | Trunk Z-Score | Standard deviations from age-matched mean (trunk) |
| value_72 | Ribs Z-Score | Standard deviations from age-matched mean (ribs) |
| value_73 | Pelvis Z-Score | Standard deviations from age-matched mean (pelvis) |
| value_74 | Spine Z-Score | Standard deviations from age-matched mean (spine) |
| value_75 | Arm Left Z-Score | Standard deviations from age-matched mean (left arm) |
| value_76 | Leg Left Z-Score | Standard deviations from age-matched mean (left leg) |
| value_77 | Arm Right Z-Score | Standard deviations from age-matched mean (right arm) |
| value_78 | Leg Right Z-Score | Standard deviations from age-matched mean (right leg) |
| value_79 | Total Z-Score | Standard deviations from age-matched mean (total body) |
| value_80 | TBLH Z-Score | Standard deviations from age-matched mean (TBLH) |
| value_81 | Trunk Left Z-Score | Standard deviations from age-matched mean (left trunk) |
| value_82 | Total Left Z-Score | Standard deviations from age-matched mean (total left) |
| value_83 | Trunk Right Z-Score | Standard deviations from age-matched mean (right trunk) |
| value_84 | Total Right Z-Score | Standard deviations from age-matched mean (total right) |

#### Average Height - values_85 through value_101
| Index | Measurement | Description |
|-------|-------------|-------------|
| value_85 | Head Average Height | Average thickness of bone in head region |
| value_86 | Arms Average Height | Average thickness of bone in arms |
| value_87 | Legs Average Height | Average thickness of bone in legs |
| value_88 | Trunk Average Height | Average thickness of bone in trunk |
| value_89 | Ribs Average Height | Average thickness of bone in ribs |
| value_90 | Pelvis Average Height | Average thickness of bone in pelvis |
| value_91 | Spine Average Height | Average thickness of bone in spine |
| value_92 | Arm Left Average Height | Average thickness of bone in left arm |
| value_93 | Leg Left Average Height | Average thickness of bone in left leg |
| value_94 | Arm Right Average Height | Average thickness of bone in right arm |
| value_95 | Leg Right Average Height | Average thickness of bone in right leg |
| value_96 | Total Average Height | Average thickness of bone in total body |
| value_97 | TBLH Average Height | Average thickness of bone in TBLH |
| value_98 | Trunk Left Average Height | Average thickness of bone in left trunk |
| value_99 | Total Left Average Height | Average thickness of bone in total left |
| value_100 | Trunk Right Average Height | Average thickness of bone in right trunk |
| value_101 | Total Right Average Height | Average thickness of bone in total right |

#### Average Width - values_102 through value_118
| Index | Measurement | Description |
|-------|-------------|-------------|
| value_102 | Head Average Width | Average width of scanned region (head) |
| value_103 | Arms Average Width | Average width of scanned region (arms) |
| value_104 | Legs Average Width | Average width of scanned region (legs) |
| value_105 | Trunk Average Width | Average width of scanned region (trunk) |
| value_106 | Ribs Average Width | Average width of scanned region (ribs) |
| value_107 | Pelvis Average Width | Average width of scanned region (pelvis) |
| value_108 | Spine Average Width | Average width of scanned region (spine) |
| value_109 | Arm Left Average Width | Average width of scanned region (left arm) |
| value_110 | Leg Left Average Width | Average width of scanned region (left leg) |
| value_111 | Arm Right Average Width | Average width of scanned region (right arm) |
| value_112 | Leg Right Average Width | Average width of scanned region (right leg) |
| value_113 | Total Average Width | Average width of scanned region (total body) |
| value_114 | TBLH Average Width | Average width of scanned region (TBLH) |
| value_115 | Trunk Left Average Width | Average width of scanned region (left trunk) |
| value_116 | Total Left Average Width | Average width of scanned region (total left) |
| value_117 | Trunk Right Average Width | Average width of scanned region (right trunk) |
| value_118 | Total Right Average Width | Average width of scanned region (total right) |

**Note:** Some indices may be empty/blank in the output if not measured in a particular scan.

---

## 2. CORE SCAN FORMAT

**Detected by:** Header contains "vat mass"  
**File Type:** Core Scan (VAT Measurements)  
**Columns:** id1, id2, id3, date, vat_mass_lbs, vat_volume_in3

### Column Reference

| Column | Description | Units |
|--------|-------------|-------|
| `vat_mass_lbs` | Visceral Adipose Tissue Mass | Pounds (lbs) |
| `vat_volume_in3` | Visceral Adipose Tissue Volume | Cubic inches (in³) |

**What is VAT?**  
Visceral Adipose Tissue (VAT) is the abdominal fat stored around internal organs. It's a key health indicator associated with metabolic syndrome, cardiovascular disease risk, and type 2 diabetes. Unlike subcutaneous fat (under the skin), VAT is metabolically active and poses greater health risks.

---

## 3. BODY COMPOSITION FORMAT

**Detected by:** Header contains "arms fat mass"  
**File Type:** Body Composition (Fat Mass/Percentage)  
**Structure:** Each measurement type has 4 values: Total, Left, Right, Delta

### Column Structure Pattern

Body Composition CSV uses a repeating pattern of 4 columns per measurement:

```
<measurement_type>_<index>_total   - Combined left + right
<measurement_type>_<index>_left    - Left side measurement
<measurement_type>_<index>_right   - Right side measurement
<measurement_type>_<index>_delta   - Difference (left - right)
```

### Mass Measurements (mass_0 through mass_N)

Each "mass" group represents a body region measurement in pounds (lbs):

#### Bone Mass (mass_0 through mass_4)
| Index | Region | Description |
|-------|--------|-------------|
| mass_0 | Arms Bone Mass | Bone mass in both arms |
| mass_1 | Legs Bone Mass | Bone mass in both legs |
| mass_2 | Trunk Bone Mass | Bone mass in trunk/torso |
| mass_3 | Android Bone Mass | Bone mass in android region (upper abdomen) |
| mass_4 | Gynoid Bone Mass | Bone mass in gynoid region (hips/thighs) |
| mass_5 | Total Bone Mass | Total body bone mass |
| mass_6 | TBLH Bone Mass | Total body less head bone mass |

#### Fat Mass (mass_7 through mass_13)
| Index | Region | Description |
|-------|--------|-------------|
| mass_7 | Arms Fat Mass | Fat mass in both arms |
| mass_8 | Legs Fat Mass | Fat mass in both legs |
| mass_9 | Trunk Fat Mass | Fat mass in trunk/torso |
| mass_10 | Android Fat Mass | Fat mass in android region (upper abdomen) |
| mass_11 | Gynoid Fat Mass | Fat mass in gynoid region (hips/thighs) |
| mass_12 | Total Fat Mass | Total body fat mass |
| mass_13 | TBLH Fat Mass | Total body less head fat mass |

#### Lean Mass (mass_14 through mass_20)
| Index | Region | Description |
|-------|--------|-------------|
| mass_14 | Arms Lean Mass | Lean (muscle) mass in both arms |
| mass_15 | Legs Lean Mass | Lean (muscle) mass in both legs |
| mass_16 | Trunk Lean Mass | Lean (muscle) mass in trunk/torso |
| mass_17 | Android Lean Mass | Lean mass in android region |
| mass_18 | Gynoid Lean Mass | Lean mass in gynoid region |
| mass_19 | Total Lean Mass | Total body lean mass |
| mass_20 | TBLH Lean Mass | Total body less head lean mass |

#### Tissue Mass (mass_21 through mass_27)
| Index | Region | Description |
|-------|--------|-------------|
| mass_21 | Arms Tissue Mass | Total soft tissue mass in both arms |
| mass_22 | Legs Tissue Mass | Total soft tissue mass in both legs |
| mass_23 | Trunk Tissue Mass | Total soft tissue mass in trunk/torso |
| mass_24 | Android Tissue Mass | Total soft tissue mass in android region |
| mass_25 | Gynoid Tissue Mass | Total soft tissue mass in gynoid region |
| mass_26 | Total Tissue Mass | Total body soft tissue mass |
| mass_27 | TBLH Tissue Mass | Total body less head soft tissue mass |

#### Fat-Free Mass (mass_28 through mass_34)
| Index | Region | Description |
|-------|--------|-------------|
| mass_28 | Arms Fat-Free Mass | Fat-free mass in both arms (lean + bone) |
| mass_29 | Legs Fat-Free Mass | Fat-free mass in both legs |
| mass_30 | Trunk Fat-Free Mass | Fat-free mass in trunk/torso |
| mass_31 | Android Fat-Free Mass | Fat-free mass in android region |
| mass_32 | Gynoid Fat-Free Mass | Fat-free mass in gynoid region |
| mass_33 | Total Fat-Free Mass | Total body fat-free mass |
| mass_34 | TBLH Fat-Free Mass | Total body less head fat-free mass |

#### Total Mass (mass_35 through mass_41)
| Index | Region | Description |
|-------|--------|-------------|
| mass_35 | Arms Total Mass | Total mass in both arms (all components) |
| mass_36 | Legs Total Mass | Total mass in both legs |
| mass_37 | Trunk Total Mass | Total mass in trunk/torso |
| mass_38 | Android Total Mass | Total mass in android region |
| mass_39 | Gynoid Total Mass | Total mass in gynoid region |
| mass_40 | Total Total Mass | Total body mass |
| mass_41 | TBLH Total Mass | Total body less head mass |

### Percentage Measurements (pct_0 through pct_N)

Each "pct" group represents body fat percentage for the same regions:

#### Region %Fat (pct_0 through pct_6)
| Index | Region | Description |
|-------|--------|-------------|
| pct_0 | Arms Region %Fat | Body fat percentage in arms |
| pct_1 | Legs Region %Fat | Body fat percentage in legs |
| pct_2 | Trunk Region %Fat | Body fat percentage in trunk |
| pct_3 | Android Region %Fat | Body fat percentage in android region |
| pct_4 | Gynoid Region %Fat | Body fat percentage in gynoid region |
| pct_5 | Total Region %Fat | Total body fat percentage |
| pct_6 | TBLH Region %Fat | Total body less head fat percentage |

#### Tissue %Fat (pct_7 through pct_13)
| Index | Region | Description |
|-------|--------|-------------|
| pct_7 | Arms Tissue %Fat | Fat percentage of soft tissue in arms |
| pct_8 | Legs Tissue %Fat | Fat percentage of soft tissue in legs |
| pct_9 | Trunk Tissue %Fat | Fat percentage of soft tissue in trunk |
| pct_10 | Android Tissue %Fat | Fat percentage of soft tissue in android region |
| pct_11 | Gynoid Tissue %Fat | Fat percentage of soft tissue in gynoid region |
| pct_12 | Total Tissue %Fat | Fat percentage of total body soft tissue |
| pct_13 | TBLH Tissue %Fat | Fat percentage of TBLH soft tissue |

#### Additional Hydration Metrics (may appear at end)
Some scans may include:
- TBW (Total Body Water)
- ICW (Intracellular Water)
- ECW (Extracellular Water)
- TBW Device (Device used for measurement)

### Understanding the Delta Values

The "delta" column shows asymmetry between left and right sides:
- **Positive delta:** Left side has more than right
- **Negative delta:** Right side has more than left
- **Zero delta:** Perfectly symmetrical

Significant asymmetry (large delta values) may indicate:
- Muscle imbalance
- Dominant side development (common in athletes)
- Injury recovery differences
- Potential medical conditions requiring further evaluation

---

## Key Terminology

### Anatomical Regions

- **Android:** Upper abdominal region (typically above the pelvis, below the ribcage)
- **Gynoid:** Hip and thigh region (typically includes hips, buttocks, upper thighs)
- **TBLH:** Total Body Less Head (useful for tracking body composition changes)
- **Trunk:** Torso region including chest, abdomen, and back

### Measurement Types

- **BMD (Bone Mineral Density):** Amount of mineral matter per square centimeter of bone (g/cm²)
- **BMC (Bone Mineral Content):** Total amount of mineral in bone (grams)
- **Lean Mass:** Muscle, organs, and other non-fat, non-bone tissue
- **Fat-Free Mass:** Lean mass + bone mass
- **Tissue Mass:** All soft tissue (fat + lean, excludes bone)
- **T-Score:** Compares BMD to healthy 30-year-old adult average
- **Z-Score:** Compares BMD to age-matched average

### Clinical Significance

**T-Score Interpretation (for osteoporosis assessment):**
- Above -1.0: Normal bone density
- -1.0 to -2.5: Low bone mass (osteopenia)
- Below -2.5: Osteoporosis

**Body Fat Percentage Guidelines (general):**
- Essential fat: 2-5% (male), 10-13% (female)
- Athletes: 6-13% (male), 14-20% (female)
- Fitness: 14-17% (male), 21-24% (female)
- Average: 18-24% (male), 25-31% (female)
- Obese: 25%+ (male), 32%+ (female)

**VAT (Visceral Fat) Guidelines:**
- Low Risk: < 100 cm² (roughly < 0.7 lbs)
- Increased Risk: 100-160 cm² (roughly 0.7-1.1 lbs)
- High Risk: > 160 cm² (roughly > 1.1 lbs)

---

## Units Summary

| Measurement Type | Units |
|------------------|-------|
| Mass (all types) | Pounds (lbs) |
| BMD | g/cm² (grams per square centimeter) |
| BMC | Grams (g) |
| Area | Square centimeters (cm²) |
| VAT Volume | Cubic inches (in³) |
| Percentages | % (0-100 scale) |
| T-Scores & Z-Scores | Standard deviations (can be negative) |

---

## CSV Import Tips

### Excel/Google Sheets
1. Import as CSV with comma delimiter
2. Enable "Treat consecutive delimiters as one" if needed
3. Set numeric columns to number format (not text)
4. For percentages: multiply by 1 (already in % form)

### Python (pandas)
```python
import pandas as pd

# Body Composition
df = pd.read_csv('bodycomp.csv')

# Access specific measurement
arms_fat_total = df['mass_7_total']
arms_fat_left = df['mass_7_left']
arms_fat_right = df['mass_7_right']
asymmetry = df['mass_7_delta']

# Total Body
df = pd.read_csv('totalbody.csv')
total_bmd = df['value_11']  # Total BMD

# Core Scan
df = pd.read_csv('corescan.csv')
vat_mass = df['vat_mass_lbs']
```

### R
```r
# Body Composition
df <- read.csv('bodycomp.csv')

# Access by column name
arms_fat <- df$mass_7_total

# Total Body - access by position
total_bmd <- df$value_11

# Core Scan
vat <- read.csv('corescan.csv')
vat_mass <- vat$vat_mass_lbs
```

---

## Troubleshooting

**Problem:** Too many columns, Excel truncates  
**Solution:** Use Import Data wizard or split into multiple sheets

**Problem:** Scientific notation in large CSVs  
**Solution:** Format cells as Number with 2 decimal places

**Problem:** Empty columns in Body Composition  
**Solution:** Normal - some records have fewer measurement blocks than others

**Problem:** Can't match value_N to measurement  
**Solution:** Use this reference guide or check original header order

---

## Additional Resources

For questions about clinical interpretation of DEXA results, consult:
- Radiologist or physician who ordered the scan
- Endocrinologist (for bone density concerns)
- Sports medicine specialist (for body composition)
- Registered dietitian (for body composition goals)

For technical questions about the converter:
- GitHub: https://github.com/derickschaefer/dxafile
- Report issues or suggest improvements via GitHub Issues

---

**Document Version:** 1.0  
**Last Updated:** November 2025  
**Converter Version:** dxafile v1.0
