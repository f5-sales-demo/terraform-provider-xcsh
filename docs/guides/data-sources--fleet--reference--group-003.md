---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-3332203333233132-1030130000012312-3123212312110121-2322211011010333-1211031101203022-0020230332221321-2010331033230222-0033232303032220"></a>

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.security_style` property

Type: `"string"`. Computed.

Security Style. Security style for new volumes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2323110203001112-2012323102131010-0301302021133110-0112111313120201-2233223101232301-3321203000113031-0101112233231302-1033023330300301"></a>

<a id="canonical-3123122003212321-0330210112000133-2130333002330330-1213202012212031-3030132221233112-3202020022033333-0313112102302121-1213310130331223"></a>

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_dir` property

Type: `"bool"`. Computed.

Access to Snapshot Directory. Access to the .snapshot directory.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1303122323302033-3131333030232321-2313002231333131-3311131001220301-1210110231203220-1301012230123221-2231000202122122-0303032002230331"></a>

<a id="canonical-1301322202210022-2020113123320022-2311301202233022-1331011121110311-3121131200320323-0121130323033132-2023231122010130-0320032010000002"></a>

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Snapshot policy to use.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1312131131123032-2020120012332023-3102333320300113-0320221030020130-2012111112110103-2020313121010311-1031310022310111-0302223221110023"></a>

<a id="canonical-0100220221013003-2303203320011333-0000113331121302-1020320022313231-0100102112213313-3123033000303101-3221021030202020-0320300223302111"></a>

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_reserve` property

Type: `"string"`. Computed.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Additional upstream details:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3131013312012312-1221200322320120-2103233022302222-0330201200313313-0322313102201012-3331202100112033-1202232021301002-1010221033233211"></a>

<a id="canonical-3030211030210130-2212330032232012-2220211331023321-0220221223111301-0002112013023123-0211320131313102-0333231222003031-3132011310312121"></a>

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.space_reserve` property

Type: `"string"`. Computed.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-2222231031120220-3323023000012211-2310110332110121-3110130203033022-3113023330310310-0230121012100103-1223312211323302-3131220302202021"></a>

<a id="canonical-0003021301333111-1110031013013322-2101212300310122-3311130012123021-1221211020230200-0233000132311002-3010310102112030-3300201011320112"></a>

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.split_on_clone` property

Type: `"bool"`. Computed.

Split a clone from its parent upon creation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2003323003031210-3331110123011303-0333223000000320-3123133313300202-3102212001210132-2330102220330023-0210032300113001-1011213222202033"></a>

<a id="canonical-3213223321011332-3230320002313323-0301022302000331-0313202012113122-0033122320300313-2120312111222201-1320132213133030-0133032202102213"></a>

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.tiering_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Tiering policy to use. "none" is default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2331313110233112-1333312133031210-2332203033231031-2120132011321132-0130231222231221-2010210021031212-1333212213203222-3210222231310203"></a>

<a id="canonical-0121332311201113-3022021223311201-2100202132102303-1232222222033133-3333000131203130-3312300323031010-2310120233012013-1103102203203002"></a>

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.unix_permissions` property

Type: `"number"`. Computed.

Unix permission mode for new volumes. All allowed 777.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3113121113310310-2213300013020110-3121130113032000-0201301203122212-1202103110001302-2031103010102301-0200203322033132-3030200300000132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-002.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--fleet--reference--group-002.md#canonical-2110120020013122-0302220310310103-1003013222131120-1301111323321323-2112112320103320-0131001012111321-0012200230303231-1111121213230133)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](data-sources--fleet--reference--group-002.md#canonical-1211033321203322-2132121023112301-1231232232311122-2110303320220002-0122310010221301-3323203132223102-1121213222001002-1103112012220111)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos

<a id="canonical-0332332220231030-0323322002200213-1231110312111111-0121233021032223-1110300032110103-3303230002322032-0330230022121001-1012200231021333"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1031002301120210-1211113332013001-0221212030021230-1213323311322101-0103111130010301-1132312332210131-1123103220103002-1121031232101023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-002.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults

<a id="canonical-3120220122313122-3321001102020332-2132201023202013-1001003321001101-2333303320313313-3102211332313313-3031321010311233-2122233011223210"></a>

Type: `"single"`. Computed.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

<a id="canonical-1122230031200031-2210230232011312-3223111200312003-3013300201030011-3200321133012323-2212001213111101-1101313031222222-2223113121102100"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults`

<a id="canonical-2223031010002302-3101233322331303-3301000120232131-2110231230033001-1132130011332202-0320001313233201-2110300210201133-3200122100132003"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.adaptive_qos_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3110212201321300-0333001203312101-2013211131100310-0311233112002320-2012011200131223-1123001112100212-1003201112320003-1033123020110010"></a>

<a id="canonical-1101101312120103-2103123113103233-3213010110030021-3123000200331110-1201202222332333-3132113001132223-2331111323332013-1130102311231003"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.encryption` property

Type: `"bool"`. Computed.

Enable Encryption. Enable NetApp volume encryption.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3000221123203222-3300232012131211-0233002213021312-2333111003222213-2230223123232122-1202100011113032-0012301301021111-1120131100021221"></a>

<a id="canonical-3032100331101020-2231101020000010-1211223200303031-3110300133311113-0123202110313210-3121030233231013-1020111020001312-1303020002220313"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.export_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Export policy to use.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_qos](data-sources--fleet--reference--group-003.md#canonical-2201230221322113-3100212201331313-2122002222033101-0232333201101302-2223232110333230-1323231320121001-0133012121022011-3011221133303100): complete subsection reference.

<a id="canonical-2213300111320222-2003130111233203-0200002221303123-2101201300102232-2022311333321233-2101331132101320-3330022310311023-0233122032330303"></a>

<a id="canonical-3322312022201122-1021333120122202-3203323121011011-0303031312330211-3232003330300123-0233301202101131-0113222111312210-3220012010130001"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.qos_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1332011221032211-3112212301230301-0032113121023101-3102202022302112-2313110212013233-0122213221033310-2222220231013232-1101123332013023"></a>

<a id="canonical-2230311310222303-1122300133023111-0330031102203000-3331302210332222-1313212230123031-1013000311030310-2233130211112102-1131100303103120"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.security_style` property

Type: `"string"`. Computed.

Security Style. Security style for new volumes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1333313030112300-0123000330132300-3313022130311233-2111333223003130-1120033330021323-2301212323320322-2202303321301112-3201213003200033"></a>

<a id="canonical-0303313110103001-1123022302203020-3030020313200213-1113132111302202-0020200122121101-2222331323020131-3322321223321131-3202200202232000"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_dir` property

Type: `"bool"`. Computed.

Access to Snapshot Directory. Access to the .snapshot directory.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2012221032212021-3111203233312123-3230033223333031-2001212222230220-1131233321223300-1120012202320210-0000001023221022-1233313020110303"></a>

<a id="canonical-0012112220213220-3311033212020310-2202230201202110-0220202131201010-0110223130313303-2132133101200302-0101122201111122-3210033312122011"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Snapshot policy to use.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3210113120113130-1302003203220102-1301322313023200-1203210302320230-1000331222221113-0202222231021122-0121110203013100-2030230202000222"></a>

<a id="canonical-0201300212301203-0220013223332222-3002022221223102-0113331121023323-2331022230021130-0002221323332031-1110302111110120-1123031202012001"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_reserve` property

Type: `"string"`. Computed.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Additional upstream details:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0030331122233001-3012301030303032-0202131032313222-3313322111323201-1031132322131310-3012132302323103-2021231030232021-2113102323323322"></a>

<a id="canonical-2121212101203222-2230232301011223-2011222111222100-3102011003321233-1223330131323300-1032212301101032-3233010130301012-1211230331122011"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.space_reserve` property

Type: `"string"`. Computed.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-1322331331203332-3023110133211300-3131212012101011-2223322332121320-3112030321111222-3321300001011101-1102130011203213-3031120013201031"></a>

<a id="canonical-3113232001202021-2021301103121322-3222032302202101-2021131210321320-2021121120231110-1121323203113010-0210113030130221-0010023210012002"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.split_on_clone` property

Type: `"bool"`. Computed.

Split a clone from its parent upon creation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0323011110310122-1331301002221123-3020113030121221-0003330131212322-0003103102133031-0333020212203002-2333201120311323-3220230211113002"></a>

<a id="canonical-2212231231023010-2333013203201320-3021213120222331-0311113120312331-3312121000201123-3023223222200310-1221031213122012-2002223331100100"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.tiering_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Tiering policy to use. "none" is default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2203210122321211-2332220031010113-3222322310010033-1021322210033003-0133121332112012-3030111213321120-3212103032012001-3312231312101323"></a>

<a id="canonical-2333303331330200-3111033120013112-3111332212202101-0311032113021310-0323011322213022-0120200112320233-2121322300302033-1313033203330030"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.unix_permissions` property

Type: `"number"`. Computed.

Unix permission mode for new volumes. All allowed 777.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2201230221322113-3100212201331313-2122002222033101-0232333201101302-2223232110333230-1323231320121001-0133012121022011-3011221133303100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-002.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-1031002301120210-1211113332013001-0221212030021230-1213323311322101-0103111130010301-1132312332210131-1123103220103002-1121031232101023)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos

<a id="canonical-2021001022300031-3202100022231310-0300011112223000-3013303023101202-3120133011203220-1232312001320012-1101010021010320-1223000203133202"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san

<a id="canonical-0312112002310313-0011230320033000-0221223220022300-0123133033010010-0331311100212213-2001331113201230-3132333113002002-0311003130300301"></a>

Type: `"single"`. Computed.

Configuration of storage backend for NetApp ONTAP SAN.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-chap_choice": "[\"no_chap\",\"use_chap\"]",
  "x-ves-oneof-field-data_lif": "[\"data_lif_dns_name\",\"data_lif_ip\"]",
  "x-ves-oneof-field-management_lif": "[\"management_lif_dns_name\",\"management_lif_ip\"]"
}
```

<a id="canonical-3102222112122100-3331311013202113-1330213022312020-3221232033232021-2021123132301110-2302233211102030-3110122002101030-3111201230210122"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san`

<a id="canonical-1103102323312322-1132320211201101-0100201201230033-3003012001303301-1121023102010221-2013113032121030-2122102102120222-1311111311322002"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_certificate` property

Type: `"string"`. Computed.

Please Enter base64-encoded value of client certificate. Used for certificate-based auth.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [client_private_key](data-sources--fleet--reference--group-003.md#canonical-1302022310332322-2210300011312230-0302201131212100-2011210313122032-0300202010231330-3313111122220121-0333031121210010-3322300012301302): complete subsection reference.

<a id="canonical-1223011231032031-3320201301010320-1212323222010312-3100102111303203-0111130310121120-1213013121213201-3132310111210203-2022213301201310"></a>

<a id="canonical-0133303300013300-2231332001331100-2210130212121233-2312130202030030-2312121211112210-2300323031301213-3030032302301032-2031331103211213"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_dns_name` property

Type: `"string"`. Computed.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1213302220302210-2213013233012032-0033220221130121-0330121333331331-0302201300211201-1112011013213202-3030311030323131-0112212103200302"></a>

<a id="canonical-0113332323033231-3033331112001321-1101213201321312-2010021300211110-2300230102331112-2033021121232033-2023110300230230-2013313120022231"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_ip` property

Type: `"string"`. Computed.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-2010323130231333-1023302320210111-2123330310131313-3333011123033023-1133133031201000-2321203301301021-2103022231110210-1321030331211003"></a>

<a id="canonical-1322201132321320-1322300322210313-3132011000101230-3201321231012102-0011133233331130-0013201210102112-2033230311111121-2331202201022303"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.igroup_name` property

Type: `"string"`. Computed.

Name of the igroup for SAN volumes to use.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0222003121002303-1132113230233322-2101231313301301-2013213121330333-0231003300112203-0302210120022030-3020103211321322-1110032332121003"></a>

<a id="canonical-2111000311131323-2121030132120300-1202320121122231-3100301223111002-1201033013332013-1203332031331103-0111231210001333-0223320323220131"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.labels` property

Type: `["map", "string"]`. Computed.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-3222131122000132-2121230300303330-1122203111112130-2331232021333331-3122100301113323-0000302131130221-2222002222211300-1022303212302231"></a>

<a id="canonical-0222313231021313-2332212310031131-1113002201300200-3122020120001213-2333030231220012-3313101312333121-1112011020130122-3000012202013202"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_aggregate_usage` property

Type: `"number"`. Computed.

Fail provisioning if usage is above this percentage. Not enforced by default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-0133123200301012-0120230111120202-0030200213020320-2132121010212232-1230210211113322-0311332323031013-1313320102131101-0312111031102122"></a>

<a id="canonical-1021131013003123-2202221321211310-0013203233312313-1220111201313333-1302313220111301-1210103021011103-2001103031133112-1200331322013013"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_volume_size` property

Type: `"number"`. Computed.

Fail provisioning if requested volume size in GBi is above this value. Not enforced by default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3310233010320233-2100333322313023-3220103023333232-2112103123313330-1023022301221002-2233022002021122-3200101010322033-3020201001323031"></a>

<a id="canonical-3233121211033022-2110210222011312-1120212231030120-0132301103102310-1132112222223131-1211013001110000-1023030011110113-0332003310113302"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_dns_name` property

Type: `"string"`. Computed.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1000220103221121-0321233031310222-3211232303013303-1020032031132101-1011330302010231-3032103023031321-3202301033110210-0103301113310032"></a>

<a id="canonical-2300220021133031-3133322301133130-1022120312300020-3103310002321031-0031233333220330-3220123023222132-0220211222210221-0333003001230132"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_ip` property

Type: `"string"`. Computed.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

- [no_chap](data-sources--fleet--reference--group-003.md#canonical-3313201320201133-2313002020200110-1323011320323131-2311012102232011-1201313313233301-0323203012311221-2032102101000000-1203131310313310): complete subsection reference.

- [password](data-sources--fleet--reference--group-003.md#canonical-1120102000102001-2032321311131031-1110201233301132-2222110002303111-1003233110303022-0232203211001321-3022023132001121-0023311012331213): complete subsection reference.

<a id="canonical-2020022010103223-2021130102210223-3133323031212321-3013332032332021-0013230203210011-1010302103032132-2220313301332002-3000002101331313"></a>

<a id="canonical-3011301103222203-0102022103102332-2203330203233131-1301121303311111-0203221111231231-0323031320221132-1001101110230103-1001233001223323"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.region` property

Type: `"string"`. Computed.

Backend Region. Virtual Pool Region.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [storage](data-sources--fleet--reference--group-003.md#canonical-0203301222213122-0311200133322320-2330230003102220-0010021312213121-1003010303003013-1100202112013010-1112331202022330-1012101223313000): complete subsection reference.

<a id="canonical-0211322203302201-2300121302203201-3133221110121030-0033122031303030-2031031031223231-1002210021211121-3032200221120123-3001310123301101"></a>

<a id="canonical-1102300020231203-0302212121331220-0121303113001201-0222111002221322-3122313201222301-2202331203123113-0331020222331013-1302002032311203"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_driver_name` property

Type: `"string"`. Computed.

\[Enum: ontap-san|ontap-san-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-san\`, \`ontap-san-economy\`, \`ontap-nas-flexgroup\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ontap-san",
    "ontap-san-economy",
    "ontap-nas-flexgroup"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-san\\\",\\\"ontap-san-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-san\\\",\\\"ontap-san-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  }
}
```

<a id="canonical-0322122121233033-3001031213030200-3002030113102013-3110230203230110-0231000221210030-1321222321220313-3201003112211322-1221133301322123"></a>

<a id="canonical-1002213230121131-2031120210121312-1010212333111333-1020320003233110-0033332203221303-3222321101333123-1002132121122210-3001103022230130"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_prefix` property

Type: `"string"`. Computed.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 80,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 80,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "80",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "80",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0133200121312230-3113011211332221-1102303012011010-2233221332110000-2313021102100330-2213210012012211-2313021120322331-2111132110313330"></a>

<a id="canonical-2223022032001132-0313310210111323-0023122212010223-2013102330131302-1313131313331322-0001312202322331-3101322310123112-1032321100122320"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.svm` property

Type: `"string"`. Computed.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3312110300012211-1112031120233122-1103331132210021-1133011111131232-1212110011221330-3311233220113312-0121232102130320-3200131222013010"></a>

<a id="canonical-3320111133012103-1120131210330222-3131303001123321-1323132132232001-3220311203120010-1310212103100100-0001120321200021-3331312020111023"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.trusted_ca_certificate` property

Type: `"string"`. Computed.

Please Enter base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [use_chap](data-sources--fleet--reference--group-003.md#canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223): complete subsection reference.

<a id="canonical-2033003311022231-2100213321002033-3013221033210111-3310100003012033-2001012002010201-1010200313133320-0112303303223033-1302012333123013"></a>

<a id="canonical-2032112031213301-3130211303303201-0132221022323131-1233133202000121-2310133112233012-0233020011321300-1113331213311331-2300131331203222"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.username` property

Type: `"string"`. Computed.

Username. Username to connect to the cluster/SVM.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [volume_defaults](data-sources--fleet--reference--group-003.md#canonical-0312231301000023-1223223123130323-1300021202220113-1200230003223230-0121123321012302-1111201303220322-3001223112113203-2120323230132130): complete subsection reference.

<a id="canonical-1302022310332322-2210300011312230-0302201131212100-2011210313122032-0300202010231330-3313111122220121-0333031121210010-3322300012301302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key

<a id="canonical-3031003302120210-1112103033332012-2322330201031200-2232132110032223-0033200211101111-3100223210012331-0233223020222332-1103232300111231"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-2003331212010330-3332003032232030-3002202002123321-3200300302111102-3303302231202000-1230322322201300-3111302020123000-0112311103032231"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key`

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-3222030100002311-0011211111301020-1200320221123312-1020200020121232-0322000012302113-1203030000200313-1021010202032220-2222030133030010): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-3221203331201322-2001320000331111-2011101102210002-1121330001202201-3002200013300231-3111332102131332-0022013300100031-2012302021332301): complete subsection reference.

<a id="canonical-3222030100002311-0011211111301020-1200320221123312-1020200020121232-0322000012302113-1203030000200313-1021010202032220-2222030133030010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--fleet--reference--group-003.md#canonical-1302022310332322-2210300011312230-0302201131212100-2011210313122032-0300202010231330-3313111122220121-0333031121210010-3322300012301302)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info

<a id="canonical-3200233233132022-2111132032132010-2303203113011311-3101221233200013-2100121322133213-0210030311010220-2002222103030121-1113323122100031"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3201210220322122-2301333022112202-3210122122213303-1003031320220232-2200221330011123-2123102110033221-1221311321122223-1000322200221323"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info`

<a id="canonical-1110023111312112-1231021123032003-1110211030300021-1230232130122322-1331132003000320-2210131221201221-3012133233002220-3203203121010223"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0003321100231310-2010123033102232-3132333113201000-1312223331300120-1033311313322221-1303120021223003-1300102132220210-3101001201310121"></a>

<a id="canonical-3032122001203032-1113203302301100-2033333213100311-0300001022001030-3333131101023103-1010312022021330-0013313113322312-1113021023011310"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1201000000021200-0331323302302310-1003021011230300-2110012333013200-2300111332112113-1132032111001232-0232231221321130-0123333033300031"></a>

<a id="canonical-0223220100230001-3001211312030220-2121022110322311-3002013123221122-2202230030023031-0221032133033203-0333131210121122-0020122010322310"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3221203331201322-2001320000331111-2011101102210002-1121330001202201-3002200013300231-3111332102131332-0022013300100031-2012302021332301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--fleet--reference--group-003.md#canonical-1302022310332322-2210300011312230-0302201131212100-2011210313122032-0300202010231330-3313111122220121-0333031121210010-3322300012301302)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info

<a id="canonical-1323002212231221-3031200330111102-1201032232101100-0132302300021133-3201311113123001-2331222030313221-0201031310312331-1131230222322201"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1200013201012031-2332323010221023-2031122121102010-3203210010031331-2120011130310202-2102033120213321-0302020321122011-2300111022232310"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info`

<a id="canonical-2103313223212303-1112220332121103-3003133033030230-3030302132003320-3121100112312031-1033230013310333-0123301002022112-3133132211303100"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0232202332113011-1113232303133010-1201101302202131-3113031120131213-3133313220323232-1320133020022231-2201231310322021-1121021222230021"></a>

<a id="canonical-3101322133020212-3100200103300120-1112012031310202-0332322333030010-0233200211321333-0120320120023122-2303212312320002-1021313020201312"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3313201320201133-2313002020200110-1323011320323131-2311012102232011-1201313313233301-0323203012311221-2032102101000000-1203131310313310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap

<a id="canonical-0310031112220200-0123213310312233-3212003112311332-2030013013221131-1302002123030320-3021012122132210-3301133303230322-1013013303310122"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120102000102001-2032321311131031-1110201233301132-2222110002303111-1003233110303022-0232203211001321-3022023132001121-0023311012331213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password

<a id="canonical-0310232031000322-2020333220332233-3102000221213313-3330031001223123-0013203031031021-1002202203222203-0221000322201112-2223022211013200"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-1333001330023233-3322230003120030-3300002032133020-2102312111102122-1010132220111232-3221301303032013-1033013113230133-0003223021100123"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password`

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-3133112300100001-2200211002101323-2112032312301132-1032203130010222-3100032002101311-0212023330122312-3012221231200212-0103213130302323): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-3102200112033112-2132102311221003-0031231023020300-2333323202030130-0103122031133233-2223001132020233-0121102213220120-3000003302100031): complete subsection reference.

<a id="canonical-3133112300100001-2200211002101323-2112032312301132-1032203130010222-3100032002101311-0212023330122312-3012221231200212-0103213130302323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--fleet--reference--group-003.md#canonical-1120102000102001-2032321311131031-1110201233301132-2222110002303111-1003233110303022-0232203211001321-3022023132001121-0023311012331213)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info

<a id="canonical-1011032010132111-1210033133110121-1211320232203110-3101201121202321-0102231300121111-0120310322130210-2210000002030013-1023222120222301"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2023111003310133-1313013220222203-0201233232222222-2321213221323012-3312033013110300-0020013231020131-2103330022110203-1131233333133332"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info`

<a id="canonical-3021103221322131-1213221210233200-0030213231113000-2311013000332121-2131311002003011-2103211010220220-2222311310110100-1111311302301102"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0222222312313033-3133322100001012-1231132131002202-1013021310210311-3021112311102221-2033333010311201-1120013323231001-2213312332021123"></a>

<a id="canonical-1221200301120023-1113310023231013-2131111212100020-0103012232320312-0230031031203203-1201213231201103-3011303113112021-1312111301331031"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1311121320101323-2023130201222003-1231113013033131-2200120011331231-2331331110122302-0100022032123312-0020003332313002-0200213003031333"></a>

<a id="canonical-0301320300030020-1130310302223322-3321013210223103-0013112113131121-1012013131030012-1032012011333311-0232301100233103-3202021223321323"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3102200112033112-2132102311221003-0031231023020300-2333323202030130-0103122031133233-2223001132020233-0121102213220120-3000003302100031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--fleet--reference--group-003.md#canonical-1120102000102001-2032321311131031-1110201233301132-2222110002303111-1003233110303022-0232203211001321-3022023132001121-0023311012331213)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info

<a id="canonical-3231010021101212-2330203013213232-0333103330222012-1312023210230211-2221020033302012-0121121131222112-1013322011201003-2301211312132103"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3233301323121313-3120022312030212-3213231203312232-0111211110013022-0303213221232313-3023202213330002-2311021010111102-0022323131212313"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info`

<a id="canonical-3302311001112122-0023202123033212-3231102203230031-1232032102132122-3210030220020013-0323103131303101-1131320333133010-1231121320313321"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1002323110330003-1303310332322020-1312023031110032-1020230110223320-0212312120311031-2200132223213012-2120220033102100-0023101031101302"></a>

<a id="canonical-2112102110332132-0031123011130100-1333013031101120-3333302311331130-1133033031011011-0023032230332100-1102111013111033-0100321102330133"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0203301222213122-0311200133322320-2330230003102220-0010021312213121-1003010303003013-1100202112013010-1112331202022330-1012101223313000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage

<a id="canonical-0310312313110012-0202213323202030-3012112021002203-2301311202031310-2103221132110001-0023130211002111-0210010331321201-1200210032113110"></a>

Type: `"list"`. Computed.

List of Virtual Storage Pool definitions which are referred back by Storage Class label match
selection.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1103311021021320-0113033131313300-1230120323220111-3120301303103302-3323323302011202-3310320221311230-1330023013331313-2232121320111333"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage`

<a id="canonical-0132220222322321-2032001131231003-1212230231132231-3302231220001233-3313323222000001-1211131031132002-2303313112002232-2123023323200031"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.labels` property

Type: `["map", "string"]`. Computed.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match
selection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [volume_defaults](data-sources--fleet--reference--group-003.md#canonical-0100022222011211-1020313330202211-0130320331202011-1032331232231310-1221120203231123-2000021330110310-2302032001301201-2122232132023331): complete subsection reference.

<a id="canonical-3010122312301113-3020310312300202-3033213301101001-0200202230333111-1303333323302230-1332000233331021-3222213231220222-1300022312133010"></a>

<a id="canonical-0002012330133012-2301311100001211-3211222131301022-0030231311030112-2111022231012123-3122200210330102-1203100021032022-0031121220000203"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.zone` property

Type: `"string"`. Computed.

Virtual Pool Zone. Virtual Storage Pool zone definition.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0100022222011211-1020313330202211-0130320331202011-1032331232231310-1221120203231123-2000021330110310-2302032001301201-2122232132023331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--fleet--reference--group-003.md#canonical-0203301222213122-0311200133322320-2330230003102220-0010021312213121-1003010303003013-1100202112013010-1112331202022330-1012101223313000)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults

<a id="canonical-2203103131332010-2222303201333330-3323333211012333-1320201310300112-3120133211123031-2122033311121220-2001122113131130-3311303210030310"></a>

Type: `"single"`. Computed.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

<a id="canonical-0322010122123212-1022312300320000-0330132233233231-1130131210323313-0032322023132330-1323223111232213-3320111312233003-3123220121032201"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults`

<a id="canonical-2013320102233201-3300111130201133-0011223320001200-0010210032310022-2321000111002213-2232022010133321-0312331301210202-1132023010030333"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.adaptive_qos_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3330000302201033-1120323012301320-3302033320011131-3220113032320022-2310123013203023-2102010002303032-3113031231122213-2032122022110203"></a>

<a id="canonical-1320210102303213-1323210112313021-0111123133321222-2030233001302003-2232233223032122-1200323322121330-0212121000012332-2222211222333123"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.encryption` property

Type: `"bool"`. Computed.

Enable Encryption. Enable NetApp volume encryption.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2101020230122302-3200133102212201-2011220101220020-2300330233212323-2203132311010020-0203002210202303-2002222113310032-1330123110313022"></a>

<a id="canonical-0301200213013223-3112211102101130-2201101000300021-2300220331010212-2032102112013231-3022322012202310-0311121111230303-0321013331200032"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.export_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Export policy to use.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_qos](data-sources--fleet--reference--group-003.md#canonical-2033121012023220-2332020112110101-0321031031111313-3200123113133023-2100333003111122-2133231032311332-3013313230331031-2230330130112112): complete subsection reference.

<a id="canonical-3010230100021132-0200112020110321-0002330300103020-2313223300322312-2011320003031203-0332012320210031-3032222303311300-3301112211003130"></a>

<a id="canonical-3130212201033033-2013101032111211-2010133322111011-2010203133133001-0102130331221220-1213110021033223-1232221121110002-0103322102023103"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.qos_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0221000231020020-1332131130032312-2330311131020220-1231102320021312-1112310200203312-0320200232131301-0321033201202230-3221201021220012"></a>

<a id="canonical-0311310132131101-2210123122101102-3223133131130200-3012102331100320-1023031030220233-2032022211131033-0220122323030232-2100132221031113"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.security_style` property

Type: `"string"`. Computed.

Security Style. Security style for new volumes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1123232011332130-0011201032221000-0211223223223103-1120202213011013-0331203301010112-2223011021023321-2213023232202223-1330211323001300"></a>

<a id="canonical-1031021120201322-1202310303132201-3000010100031302-2121320011320003-2003100300211111-3230000212313201-2201020320220211-3002312101033233"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_dir` property

Type: `"bool"`. Computed.

Access to Snapshot Directory. Access to the .snapshot directory.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1102211102123332-1112230032312300-1313233212211231-1212133131200020-0333113120301323-1122122222132031-2210020331323023-0011312223223201"></a>

<a id="canonical-1010302110211122-0102032112201011-0212211321331320-1200222011121130-1012312122020200-2012230022133320-3211121100021231-3312220031102031"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Snapshot policy to use.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2002300030332321-3033310313110003-2002101330113031-2311001310032112-2003003201021131-2100203203012100-0303210220011123-2210211003101101"></a>

<a id="canonical-3221123211012120-1230320223033120-1011210312331211-2130103011212202-0203033121023001-0200113311022332-0300013001100202-3120000320101110"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_reserve` property

Type: `"string"`. Computed.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Additional upstream details:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3002023203302231-2221001032023310-2211222202231023-3020032100330233-3030003201330023-3101130302023023-2202210130311310-3030000332031022"></a>

<a id="canonical-2301303101301130-0301112133132031-0322323103211023-3230021333000011-3003011232210120-2133200200022123-3313323310212230-1330320001010010"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.space_reserve` property

Type: `"string"`. Computed.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-2320313331131302-1112030310210210-0230101233021223-2211101021030133-3101033331211320-2011033131232112-3302221120330201-3001031212120113"></a>

<a id="canonical-1131012132021033-1322200221212033-3213222232122123-2103321012201313-2312222033012332-0223110002223132-1323103332103123-0013031022122323"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.split_on_clone` property

Type: `"bool"`. Computed.

Split a clone from its parent upon creation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1030303301321303-0302111101300323-1021221100110002-2030111031101201-0322331311323320-2031313011331000-3302203021211212-2301132032231002"></a>

<a id="canonical-0120103221121031-1112100001130301-2220121212312103-1002010022121133-1210003132110232-2023323002322322-2130303222033302-2033000322010232"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.tiering_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Tiering policy to use. "none" is default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2012300231311000-3021200322013233-2313312033023211-2130003133132132-2012112312011210-0202102231223003-1123213130300332-1202111130113001"></a>

<a id="canonical-2122221010230111-0001022230330210-3100200221102323-0102302210311112-1123103300021210-1311211320131332-1111222333032023-0033201010010100"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.unix_permissions` property

Type: `"number"`. Computed.

Unix permission mode for new volumes. All allowed 777.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2033121012023220-2332020112110101-0321031031111313-3200123113133023-2100333003111122-2133231032311332-3013313230331031-2230330130112112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--fleet--reference--group-003.md#canonical-0203301222213122-0311200133322320-2330230003102220-0010021312213121-1003010303003013-1100202112013010-1112331202022330-1012101223313000)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-0100022222011211-1020313330202211-0130320331202011-1032331232231310-1221120203231123-2000021330110310-2302032001301201-2122232132023331)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos

<a id="canonical-2132313213220021-3232233203203002-1111201021233232-3200200321100203-3333021312131130-3132023011210130-0201023132223020-0213202033021210"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap

<a id="canonical-3022032202232031-2330032132010002-2220122221011220-2323310200322320-2323222110012203-3302020322231021-3023011223311111-2203231010130110"></a>

Type: `"single"`. Computed.

Device NetApp Backend ONTAP SAN CHAP configuration OPTIONS for enabled CHAP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2320330312330121-3022322211030221-3101031020221202-3120000120211003-0302303222023233-3222201231131122-1323002202322332-2232112312202331"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap`

- [chap_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-2222302212123213-1032300331301122-3210031123122321-2010312311333111-1323331323202131-0103320000013002-2310212230020120-1210311133011322): complete subsection reference.

- [chap_target_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-3310332213203121-2032233103313131-0111301002333130-3323212301221301-3320230030331300-2033210302113211-1123113110010333-1310330231220030): complete subsection reference.

<a id="canonical-2131322311223003-1221000333013200-3013031110012322-3213230321322110-3231130002311102-0000331300131010-3331222203103330-1100210003121301"></a>

<a id="canonical-0313003001312201-0132113320302333-2333003311221000-0312321310132333-2111031030230210-0100101131131211-1033031222023000-3132233230032110"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_username` property

Type: `"string"`. Computed.

Target username. Required if useCHAP=true.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3133233032230213-1222110300230203-0123211033313023-1120202133101201-3021103202131313-0313201022132200-2001122200012300-0311003232302201"></a>

<a id="canonical-2022323222201131-3011200220212201-1010212300202313-3310213111222021-2211320221001330-0302220202200203-0000312001201010-2331231330322320"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_username` property

Type: `"string"`. Computed.

Inbound username. Required if useCHAP=true.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2222302212123213-1032300331301122-3210031123122321-2010312311333111-1323331323202131-0103320000013002-2310212230020120-1210311133011322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret

<a id="canonical-1121122311111333-2010100230213332-2302011132320211-3220012023210112-1010310010132010-0131110031200101-1020110123220331-2021130301303103"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-0013013321032021-0131010233001101-2230000013321012-3323103023330233-1001301002311131-0111203111223331-1331230031120120-3112031122323021"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret`

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-2020211030320032-0322010313303130-0100222030313123-2032033001123011-1213103330003331-1320230100311212-0131130132220300-1021102203013301): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-0112301002112002-2212002033302312-3101021132222000-0102320320013121-2303232230000130-3001211103013021-2032311020113303-0002202113300332): complete subsection reference.

<a id="canonical-2020211030320032-0322010313303130-0100222030313123-2032033001123011-1213103330003331-1320230100311212-0131130132220300-1021102203013301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-2222302212123213-1032300331301122-3210031123122321-2010312311333111-1323331323202131-0103320000013002-2310212230020120-1210311133011322)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info

<a id="canonical-1210312020222122-0311102012211233-0021003031210230-0131220221032203-1113101200321321-2112312121030212-0112300123322021-0130212223322123"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2203210301203310-2301201213000230-2232111121100001-0130000100102201-0313111011001001-2103030321213112-1023332201320333-3211001302122011"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info`

<a id="canonical-1120220122032302-3032231022031201-2013133023021022-0020021013132013-2012013320112300-3310232103032333-2321010330013200-2131123232332222"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1301222302320302-1033103201301010-0013032000032113-2333132021002022-2110311221301021-2123130322322300-3333132132331232-2010121002220231"></a>

<a id="canonical-2330213200212001-0232312033232323-2113211010112220-3111232231312101-2002230001001130-3031333323102211-3111310202011323-2201102010213023"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3133211210222110-1303113032323320-3021121023221203-2212300223123323-0323311023313213-1220312210223000-0320210331013030-2232200130220230"></a>

<a id="canonical-1021302301203220-0011320010120121-3323112121121212-1032200203021131-2033222111130102-3233102123033112-0022010131321211-1301201121312023"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0112301002112002-2212002033302312-3101021132222000-0102320320013121-2303232230000130-3001211103013021-2032311020113303-0002202113300332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-2222302212123213-1032300331301122-3210031123122321-2010312311333111-1323331323202131-0103320000013002-2310212230020120-1210311133011322)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info

<a id="canonical-2223020212022331-3131033330011233-3232030101122132-1333030132122021-0031133120030033-2012201301110021-0012122110021011-3003030022302031"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3103030303223011-2231202303033121-1200322230221120-3102232012130133-3323221301112333-2330101303030103-0033022000003233-0021223002303111"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info`

<a id="canonical-3113003011211123-0330012032002031-3113301030003211-3230130333202022-0030123030102020-2123031311212233-3112300212311101-0131001120133333"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0221013221131232-1313102331122231-2322101201212133-0320331010130130-3002330132033011-3130022311212231-1200112123123321-0310212123102121"></a>

<a id="canonical-2022221033320203-0220111122203020-3113121000321301-3020012333031220-0021330302013311-2223301101100020-3102311110021221-2120223120131012"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3310332213203121-2032233103313131-0111301002333130-3323212301221301-3320230030331300-2033210302113211-1123113110010333-1310330231220030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret

<a id="canonical-2111032202300332-0202103323100031-0223020330201333-1011031112200332-0333131202301203-0133301002033001-3102120001132110-3122312230312033"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-3322010223012100-1023301102303001-1300032202020102-0232012111233131-0232001010121223-0102210023322313-3321332013202121-3220321233030020"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret`

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-3322321230331020-0100100110023131-0111023101130232-3213312223203302-2201122113320010-2230002030222100-2312221020031100-1011330023231132): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-1310022120000032-1223310030013111-2232132100112231-2213123112022221-2012011031030031-3221030232210233-1030312121231330-0210300320102301): complete subsection reference.

<a id="canonical-3322321230331020-0100100110023131-0111023101130232-3213312223203302-2201122113320010-2230002030222100-2312221020031100-1011330023231132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-3310332213203121-2032233103313131-0111301002333130-3323212301221301-3320230030331300-2033210302113211-1123113110010333-1310330231220030)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info

<a id="canonical-2311100313330320-2301332312000222-0132012020012220-3312112032210300-0311211333202033-1123022020032013-0321320230230133-0133210020313031"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1020302200331130-1210021023100303-3200133131302330-3003211320231121-2222321131101133-3010033033120113-3013121021223303-0121232003333101"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info`

<a id="canonical-2211320130111310-1113210033212101-3030102032300111-0131130203300231-3111210302033221-0211213112213233-0123001102223320-3113203303300310"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1012211210321310-1321131033011022-0300322113231233-3210331033023011-2001221103330130-1213022100232211-1020321310331111-3032231012313210"></a>

<a id="canonical-3012130330023122-2331031020030003-3030223202021102-1023011022122001-1301220213332230-1302032131202331-3012100222231200-0133313022022213"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3310131330110201-1330211300211302-0012022130321133-2011213013112032-2003333211310323-1122330203113230-3222200301023030-1013022113011332"></a>

<a id="canonical-1321310232232311-1030230012323121-1330311323133103-3311021122210030-2122311001333203-1311013320320121-3330233102010003-3020200131300013"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1310022120000032-1223310030013111-2232132100112231-2213123112022221-2012011031030031-3221030232210233-1030312121231330-0210300320102301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-3310332213203121-2032233103313131-0111301002333130-3323212301221301-3320230030331300-2033210302113211-1123113110010333-1310330231220030)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info

<a id="canonical-1011000022331322-2202333231231210-0330313312000132-3313012122010202-3233212323000010-3223133011120330-1322222332322100-1331230010220312"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1321021030230002-2132200122333030-0211112101223203-3333100322021213-0102121310012210-2021313203203033-3111322212320000-0100120112011022"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info`

<a id="canonical-3130210001220331-0233213130111131-0012120132323313-0311211332030320-1100220033012120-1020303113232212-3311233132001000-2320202111032102"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2013110222210210-1213300101120031-2233101311031203-0110310031232100-1311030201012233-0123330130010330-3311003321113101-3233110103031131"></a>

<a id="canonical-0322011120023133-3100221003003222-0103231230223301-1011303003300230-3303101201320123-0132233112020332-2202203100201203-2001200003111130"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0312231301000023-1223223123130323-1300021202220113-1200230003223230-0121123321012302-1111201303220322-3001223112113203-2120323230132130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults

<a id="canonical-1332120030303100-1011030102300031-1320033110110031-2221101323321313-1001122120023202-1122303033122221-2011320210112110-3301302212233320"></a>

Type: `"single"`. Computed.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

<a id="canonical-0023121211121031-1021021203300320-0203313231213020-0111002132220330-2213132100332323-0311030033003001-1300103111212303-2111210113103100"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults`

<a id="canonical-0223312230221031-2223213122223030-3201123231233102-1211211210013000-0000101231233033-1122112233103021-0211132103033023-3030330123123131"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.adaptive_qos_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3011310020100202-3001033230102321-3010133231100130-2120311321122211-2001003321333132-0030130133301023-2023310113301230-1123311130320231"></a>

<a id="canonical-0320320311113030-3021103010222001-2233013311203320-3303332122102021-2131001332110330-0201012023311311-0022100321121113-3222321113332132"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.encryption` property

Type: `"bool"`. Computed.

Enable Encryption. Enable NetApp volume encryption.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1331313020000122-2332221302010221-0120000130322233-2112010230022122-0222233113312112-3331123001232100-1033222001000120-0131002032012223"></a>

<a id="canonical-0033300200011232-0130301033331022-1000012112110113-3003033120103102-2133301023330333-0322102223000121-1100332101303120-3021333202133123"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.export_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Export policy to use.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_qos](data-sources--fleet--reference--group-003.md#canonical-0112010323101023-2331002330103032-1223131001001030-0221301023032130-0013032111330233-1231122211211221-0032203001033102-2302001301303210): complete subsection reference.

<a id="canonical-1001113023012001-3131000323101332-2012021220333012-0003122123313331-0303213333202120-2031203313001111-1320221032022110-0310231121103212"></a>

<a id="canonical-3132332302331232-2102303002101123-1211000021322200-3032212012021310-0113021110301113-2002222102232122-0332021322131130-3302213031223011"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.qos_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2323313100102031-3223022212130221-3323123312200312-0100010032031300-1021202031130002-2033120300302300-1100133201212101-3011101000332110"></a>

<a id="canonical-2221130300310021-2200323032112210-1023313230233303-2311322113122031-3012210020323233-2222233302113013-2111003223021021-3031301322231020"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.security_style` property

Type: `"string"`. Computed.

Security Style. Security style for new volumes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0303201022200003-3011211131102101-1332033220300030-3331223031121001-2211311003130023-3332233122023222-3030201033101011-2320130020313210"></a>

<a id="canonical-2001011132111210-1013200110220102-3120312331233131-3022132023111033-2320010321331211-2010002101120221-0131203332330203-3332121132212001"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_dir` property

Type: `"bool"`. Computed.

Access to Snapshot Directory. Access to the .snapshot directory.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2330002311031200-1323302032003100-3221003230330001-3231333232100011-0213330222321311-3030110000231210-0223011003330213-3321020023303103"></a>

<a id="canonical-0310212201233111-1100113120311310-2320000122333000-2010313102222222-2022203233220110-1023023302322323-1101321313032200-0311221301223023"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Snapshot policy to use.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0010203023233122-0111302032300213-1300023003311020-0013222202311101-2312322231222022-1203011221233230-3120032300110000-1101230000031130"></a>

<a id="canonical-2302212123331202-3302312012301013-2100330110030223-1312010031010122-3210011210110233-1220121111211023-3010311223222133-2312210220333201"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_reserve` property

Type: `"string"`. Computed.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Additional upstream details:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3032231002021200-2033111103132301-2331300103000022-3202013221031012-3300211032033002-2303300120022223-3131133023310111-0001221122132232"></a>

<a id="canonical-2013321311223203-2123302212120200-3131113010133220-1132202023300033-1331223323333223-3201002132111033-0322120230000210-1302123232321203"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.space_reserve` property

Type: `"string"`. Computed.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-0132231321211230-3002222301200020-1221220123322223-0113001131301220-2113312122022311-2213200320001023-0230013133022022-3211131033002002"></a>

<a id="canonical-1112233133010120-0110200012300311-2131020130003102-3232022112012131-3212211233023013-0210231001023311-3102212203132100-0102300312212111"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.split_on_clone` property

Type: `"bool"`. Computed.

Split a clone from its parent upon creation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2002101112132233-2132010020221320-1032230122321301-1231312010330200-1333220131012313-1310020323212332-3001333210011011-3123031102202001"></a>

<a id="canonical-0232233221303000-1110030113221310-3223101012031120-3301001332333332-1233102212312010-0110233000301123-1132001030333030-0333321333321322"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.tiering_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Tiering policy to use. "none" is default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3333032202313302-3232311230323102-0130002213122033-1002321012023232-3233102321222120-1230131003203310-0220111330012033-3033031132022130"></a>

<a id="canonical-2203223320331323-1212113010123321-2023200321323001-0022310010100032-3012123223322033-1323121031030302-0021110213132211-3322211102023323"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.unix_permissions` property

Type: `"number"`. Computed.

Unix permission mode for new volumes. All allowed 777.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0112010323101023-2331002330103032-1223131001001030-0221301023032130-0013032111330233-1231122211211221-0032203001033102-2302001301303210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-0312231301000023-1223223123130323-1300021202220113-1200230003223230-0121123321012302-1111201303220322-3001223112113203-2120323230132130)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos

<a id="canonical-1103111130302023-0013203012111220-0020223120232110-0230113320122320-0310232321323221-3302022113031001-1123232211313310-3113322022133202"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- storage_device_list.storage_devices.pure_service_orchestrator

<a id="canonical-3103200021113021-2113212020022021-3033203132002300-2030331010033303-1110331103303232-2221010220120101-2001201013001121-1002013130103233"></a>

Type: `"single"`. Computed.

Device configuration for Pure Storage Service Orchestrator.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0012322311113230-1323012232331202-1210232231323222-1013201033223223-0132010202201222-3322313332310320-2210312132102123-3111102001302222"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator`

- [arrays](data-sources--fleet--reference--group-003.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223): complete subsection reference.

<a id="canonical-2101031202010310-1031221221321122-0112033111020322-1313301032211103-1233131223333332-0202020103320112-0202332120012200-2220202203231332"></a>

<a id="canonical-3223201301232000-3013300023110300-0311213211321330-2322101121120303-2213022011323022-0130211322030312-2310030213203100-3112130301131222"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.cluster_id` property

Type: `"string"`. Computed.

ClusterID is added as a prefix for all volumes created by this PSO installation. ClusterID is also
used to identify the volumes used by the datastore, pso-db. ClusterID MUST BE UNIQUE for multiple
K8s clusters running on top of the same storage arrays. Characters allowed: alphanumeric and
underscores.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 22,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 22,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z0-9_]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "22",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9_]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "22",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9_]*$"
  }
}
```

<a id="canonical-2100311003310111-1001222033122302-2211102021322333-1223100012322100-0233230012221110-3002332033331023-3321111002210030-3330021013101222"></a>

<a id="canonical-0130230010312302-3000332230132312-2121120132023210-0321020210310221-0013331320301302-1230013100121321-0201233311320202-3132320032122122"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.enable_storage_topology` property

Type: `"bool"`. Computed.

This option is to enable/disable the csi topology feature for pso-csi.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1323013031203322-3200122311011320-1310311133223331-2232312032311312-1122300022321220-3123113032020031-2122031111110210-1303021330101313"></a>

<a id="canonical-0003131202030103-0302021232133332-0031330231203333-0021001112212132-3310202201210313-0303210312131322-2011123201321220-2313111202201300"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.enable_strict_topology` property

Type: `"bool"`. Computed.

This option is to enable/disable the strict csi topology feature for pso-csi.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-003.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays

<a id="canonical-1220133332301021-2220302313013310-2232231112111331-3331213003133231-0002230133132323-0301013231023201-3323320200020303-0002121312333132"></a>

Type: `"single"`. Computed.

Arrays Configuration. Device configuration for PSO Arrays.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2130232020111113-2031301113101320-1300110202202303-3021111312131022-1010303012133011-0110211110101030-3102201133113122-2302312130122133"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays`

- [flash_array](data-sources--fleet--reference--group-003.md#canonical-3100333230102320-2201210300121102-1332331302002131-1233110231203031-3013323101032112-2023313222111010-1131233313003323-0221211230200313): complete subsection reference.

- [flash_blade](data-sources--fleet--reference--group-003.md#canonical-2032231100232113-1121332300210032-2300122032301201-1022223222230332-0320022001320313-0322210021221233-1130002200302111-0110333133311303): complete subsection reference.

<a id="canonical-3100333230102320-2201210300121102-1332331302002131-1233110231203031-3013323101032112-2023313222111010-1131233313003323-0221211230200313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-003.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-003.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array

<a id="canonical-0302201022230330-0012011321221330-0120310021103012-0300303211221032-0002221201013200-0302113310212312-0101323330202020-2223032211113330"></a>

Type: `"single"`. Computed.

Specify what storage flash arrays should be managed the plugin.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0320302230001202-3322223233321100-1030330010311222-0220313201230322-3030121112132301-3211230001101211-0313322330230122-1020030000202223"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array`

<a id="canonical-2131223011212022-0110323200102202-3102103122033213-1220231130221110-0000320233110000-0232013123213010-0031221003021212-0011211101013011"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_opt` property

Type: `"string"`. Computed.

Block volume default mkfs OPTIONS. Not recommended to change!

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1130112130001332-3011220300332020-3031033330012131-3111221300123102-0111120122221201-3010230320223201-3310302032111100-3223203232330311"></a>

<a id="canonical-3101300303222121-0321203322002300-3311112012232211-3023023200202203-0213211201210023-1023032221303333-2113302120332033-3023223113003113"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_type` property

Type: `"string"`. Computed.

\[Enum: xfs|ext4\] Block volume default filesystem type. Not recommended to change!. Possible values
are \`xfs\`, \`ext4\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "xfs",
    "ext4"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"xfs\\\",\\\"ext4\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"xfs\\\",\\\"ext4\\\"]"
  }
}
```

<a id="canonical-2111332201122310-1111021121110001-1020230131023332-0111303001130013-1321332210011010-3011021231333313-2330002203310222-2233032111212102"></a>

<a id="canonical-2000000131100020-3210201312320232-3230022212211110-1120123221030130-2001201111223232-3320202021002110-1001110232313331-3133200123010120"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_mount_opts` property

Type: `["list", "string"]`. Computed.

Block volume default filesystem mount OPTIONS. Not recommended to change!

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1312021032030033-0330322302300111-1231130221323323-1203030311310311-1310320211220302-0201200231101221-3301133210231120-3012120331002102"></a>

<a id="canonical-3321320113020113-3101200010333112-0220303200313013-3122133133101310-2102030123322223-2131312102323311-0002032201221321-0110112323311010"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.disable_preempt_attachments` property

Type: `"bool"`. Computed.

Disable Preempt Attachments. Enable/Disable attachment preemption!

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [flash_arrays](data-sources--fleet--reference--group-003.md#canonical-3302231123231130-1200133220000322-1120123000331211-2120021120300333-0312210101021321-3311211303232112-3012230333112203-3302201320332102): complete subsection reference.

<a id="canonical-3112101121232013-1311320131013102-1213323211210312-3003330003321133-1023101001021133-1032312121322322-2220212123121302-2303030132133312"></a>

<a id="canonical-2033223211101023-1132200120022110-2201012322133332-2130132232130331-3333221102302333-1311231123313003-3233112112231010-0121102110301112"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.iscsi_login_timeout` property

Type: `"number"`. Computed.

ISCSI login timeout in seconds. Not recommended to change!

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-0302122123033100-0300322222122110-0131213232020222-0100030303311300-2100323000202203-0301101023301203-3112320301113213-1000301301033213"></a>

<a id="canonical-1002000233231232-2111231111020223-1022120333201221-3020112003331031-1002021300011222-0023031003010023-1030120212011113-2130023003010221"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.san_type` property

Type: `"string"`. Computed.

\[Enum: ISCSI|FC\] Block volume access protocol, either ISCSI or FC. Possible values are \`ISCSI\`,
\`FC\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ISCSI",
    "FC"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ISCSI\\\",\\\"FC\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ISCSI\\\",\\\"FC\\\"]"
  }
}
```

<a id="canonical-3302231123231130-1200133220000322-1120123000331211-2120021120300333-0312210101021321-3311211303232112-3012230333112203-3302201320332102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-003.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-003.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--reference--group-003.md#canonical-3100333230102320-2201210300121102-1332331302002131-1233110231203031-3013323101032112-2023313222111010-1131233313003323-0221211230200313)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays

<a id="canonical-3103113323013021-2013123110001301-0032021233130100-0021032201110132-2311200320133212-3200330121131233-0010210220220301-1222312111012121"></a>

Type: `"list"`. Computed.

For FlashArrays you must set the 'mgmt\_endpoint' and 'api\_token'.

Additional upstream details:

For FlashArrays you must set the "mgmt\_endpoint" and "api\_token"

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0103120103112332-3110333122331132-2223310021122202-0001111212302010-0020123323103000-2113020133303323-2133301002001233-1321213001101020"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays`

- [api_token](data-sources--fleet--reference--group-003.md#canonical-3321023201313323-3303221121321200-2200021032301011-3013011202132111-3220321023021033-2022001332013210-2231012012223321-3222320330103321): complete subsection reference.

<a id="canonical-1112222310111031-2023132301131211-2021102110323103-2132021132303230-0011321132112203-0332222000221312-0313301331033311-3031333302322022"></a>

<a id="canonical-0333123121311212-1021331101110000-2312232221321010-1021023233121322-0331302002213320-1331310010011011-2311300021303123-3133132310321223"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.labels` property

Type: `["map", "string"]`. Computed.

Specifies labels optional, and can be any key-value pair for use with the PSO 'fleet' provisioner.

Additional upstream details:

The labels are optional, and can be any key-value pair for use with the PSO "fleet" provisioner.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-0200001223210200-3012103103032021-1001302330322103-2310112332122211-0213121022001203-0131313323330230-1023131312323120-1130002320320110"></a>

<a id="canonical-1233131233301300-0220133131022230-2223011301212210-1013102320100101-1033023320332021-1120132322031230-3312331213021210-0003013301131103"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_dns_name` property

Type: `"string"`. Computed.

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1002222202122210-1022200331213131-0330331202330031-2101113133230232-3002120313020201-1022032203033121-3210010302310020-2012130033302321"></a>

<a id="canonical-3123332322003322-0302301201022221-0211310303013332-1320020203132220-1020010212220202-1203101100132220-3132000122031202-0302302013320102"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_ip` property

Type: `"string"`. Computed.

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3321023201313323-3303221121321200-2200021032301011-3013011202132111-3220321023021033-2022001332013210-2231012012223321-3222320330103321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-003.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-003.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--reference--group-003.md#canonical-3100333230102320-2201210300121102-1332331302002131-1233110231203031-3013323101032112-2023313222111010-1131233313003323-0221211230200313)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--fleet--reference--group-003.md#canonical-3302231123231130-1200133220000322-1120123000331211-2120021120300333-0312210101021321-3311211303232112-3012230333112203-3302201320332102)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token

<a id="canonical-0312220003010031-3202200103111110-2313231221013322-0223033121020023-2122311002211321-1000000122320110-1331102130132220-3302030133210122"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-3311301022320301-1130023202332110-2102120111033001-2200202322311303-1221231213200021-3000231110303321-0120103111322313-1330333222311201"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token`

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-0101333221021310-3111321213132322-1332102111332101-2000231131021001-3230121032302101-0131122211220312-2021023333201121-2310302303030320): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-3330220222201321-1021033010023002-1000120201020120-2012312101103123-0200202233223112-1322120311312223-2332003311111013-0113001001311222): complete subsection reference.

<a id="canonical-0101333221021310-3111321213132322-1332102111332101-2000231131021001-3230121032302101-0131122211220312-2021023333201121-2310302303030320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-003.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-003.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--reference--group-003.md#canonical-3100333230102320-2201210300121102-1332331302002131-1233110231203031-3013323101032112-2023313222111010-1131233313003323-0221211230200313)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--fleet--reference--group-003.md#canonical-3302231123231130-1200133220000322-1120123000331211-2120021120300333-0312210101021321-3311211303232112-3012230333112203-3302201320332102)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--fleet--reference--group-003.md#canonical-3321023201313323-3303221121321200-2200021032301011-3013011202132111-3220321023021033-2022001332013210-2231012012223321-3222320330103321)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info

<a id="canonical-0303332331030102-3213230100301230-2102130323020100-1223212223021120-3213212310022210-0133322103031123-1230212210233232-1102330200231000"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1233320310002301-2010300131203322-0300003330222203-2120100132301002-0121223331233031-3330302020230311-1103010310222201-2231301301323023"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info`

<a id="canonical-1301213221022011-2231221221103202-3000112200213012-2231000223301013-1021212033020000-0233132122102103-3112003130032300-1031203031001113"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1120002232121121-0020133330312102-1233002123022023-2222100012031210-3333012000012201-2312103322322322-2001001013303301-0320031321323322"></a>

<a id="canonical-1310220002012122-2033102103332030-1231112231231011-1332122013320310-2321001333220212-3103102131022201-1221122211202303-0102131211131030"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1112231322311101-0130133230303320-1202231133221001-0131233031230031-3112310032112301-0322201122320222-3112130212301223-2021001220103100"></a>

<a id="canonical-0313303012323030-0312103333220020-0202132032013021-2232333030211323-2331120112212122-2133202222221332-1132332012012121-0023233032102001"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3330220222201321-1021033010023002-1000120201020120-2012312101103123-0200202233223112-1322120311312223-2332003311111013-0113001001311222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-003.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-003.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--reference--group-003.md#canonical-3100333230102320-2201210300121102-1332331302002131-1233110231203031-3013323101032112-2023313222111010-1131233313003323-0221211230200313)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--fleet--reference--group-003.md#canonical-3302231123231130-1200133220000322-1120123000331211-2120021120300333-0312210101021321-3311211303232112-3012230333112203-3302201320332102)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--fleet--reference--group-003.md#canonical-3321023201313323-3303221121321200-2200021032301011-3013011202132111-3220321023021033-2022001332013210-2231012012223321-3222320330103321)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info

<a id="canonical-0021013322031303-1233222001012201-1212301210302101-2132101033133303-2030233032122331-1202122300122130-3001002022322231-1131000320223323"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1213303001212122-2033013233300030-1130221020213230-0210213011001331-2333112110001133-1212131323331131-1122113011011231-1020321333103003"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info`

<a id="canonical-0233021131130012-1003022133313212-2323213030011331-0103220221112132-2101301100023102-0300300203023310-0012332223131302-2031201303321010"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3231022313133212-2310231300232001-3103202220310300-3110013003310323-1122322232300132-2000030202322112-0312013133130130-1022013302120000"></a>

<a id="canonical-3303222333233310-3232032030331111-0013011013202233-2222330332100213-2110201222010123-1313020113231100-2322132221102311-3230211211310110"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2032231100232113-1121332300210032-2300122032301201-1022223222230332-0320022001320313-0322210021221233-1130002200302111-0110333133311303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-003.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-003.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade

<a id="canonical-3301212300021000-3022111311233332-3222331022202022-3011032333032301-1301213320333212-3331122330001310-3311232031032221-1313100121032322"></a>

Type: `"single"`. Computed.

Specify what storage flash blades should be managed the plugin.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3100332111313031-3230033321232130-3302020322210023-3310030032330012-2120332212002302-1331223222010210-0230120010101011-1021022013013203"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade`

<a id="canonical-0023111222020120-2130321133020000-1031123032323023-2302302220113121-2102102203132201-1202302320212330-1020122100230011-0030123122231300"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.enable_snapshot_directory` property

Type: `"bool"`. Computed.

Enable Snapshot Directory. Enable/Disable FlashBlade snapshots.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0233020311321021-3333233212201203-0110201322310133-1232233133231221-1013131132302333-0333213300033002-0101200203201323-0231123121011010"></a>

<a id="canonical-3323013131030030-1201331313300312-0100332100203333-0013010100121002-1032031121113312-2112333030310302-0301210300110301-0010001230231313"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.export_rules` property

Type: `"string"`. Computed.

NFS Export Rules. NFS Export rules.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 250,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 250,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "250",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "250",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [flash_blades](data-sources--fleet--reference--group-003.md#canonical-3330123332213122-0311333303020203-2100301133303321-2131232010222233-1100023113000033-3113311212320100-1102102332012100-0222103102231330): complete subsection reference.

<a id="canonical-3330123332213122-0311333303020203-2100301133303321-2131232010222233-1100023113000033-3113311212320100-1102102332012100-0222103102231330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-003.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-003.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--reference--group-003.md#canonical-2032231100232113-1121332300210032-2300122032301201-1022223222230332-0320022001320313-0322210021221233-1130002200302111-0110333133311303)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades

<a id="canonical-3113322203312112-2320103102230022-3101033023320312-2123121313032122-2320123021333000-2221033111110100-2303310202201022-2321230333121021"></a>

Type: `"list"`. Computed.

For FlashBlades you must set the 'mgmt\_endpoint', 'api\_token' and nfs\_endpoint.

Additional upstream details:

For FlashBlades you must set the "mgmt\_endpoint", "api\_token" and nfs\_endpoint.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1021013331320010-3111333332121013-1112320131223031-2123012320030331-3032100123130031-3120011012130230-2122001103320222-1102101033102032"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades`

- [api_token](data-sources--fleet--reference--group-003.md#canonical-0120031312202210-1213230232032302-3221130111010021-1123100131110111-2100101020033100-0310021321221020-0012013030020321-2020100103212223): complete subsection reference.

<a id="canonical-2333020233221223-1203231313233301-0011030002121232-1321030012230003-0212131322032202-3203121132213300-3333002320200322-0111332122332113"></a>

<a id="canonical-3221223202021003-3330311031310312-1001110311103230-0023221210311310-2133020202333110-1301300102301203-0121230321011302-2020202133030021"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.labels` property

Type: `["map", "string"]`. Computed.

Specifies labels optional, and can be any key-value pair for use with the PSO 'fleet' provisioner.

Additional upstream details:

The labels are optional, and can be any key-value pair for use with the PSO "fleet" provisioner.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-0122232110112012-3300332203201123-3010320001220022-2132312332032022-3203111132202033-0113021022002101-2230213123302100-3110131023000101"></a>

<a id="canonical-3223123232323132-2222020220301001-0300111221022221-1230200103022132-2330032222333220-0123223203000323-1003230101131111-1133331021132333"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_dns_name` property

Type: `"string"`. Computed.

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3101133123322320-3112101220012312-0100301313203222-0313121231233223-1132123032233133-2130031233203210-0120301230333002-3013022301113032"></a>

<a id="canonical-0320030311021031-2021133113323231-2232310103220033-2210301102312131-1233002212322033-0023300211212010-3021101213022001-1222000230131133"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_ip` property

Type: `"string"`. Computed.

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3131003020012111-0320223201311102-2313012303001211-2322001033221203-3311332031133023-3010221210303013-2020220111103120-2221022020202032"></a>

<a id="canonical-1302000022123003-2120302230020230-0322103123301112-1330200200211103-3120033213031222-3131211003320321-2201312132222030-1133232101122010"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_dns_name` property

Type: `"string"`. Computed.

Exclusive with \[nfs\_endpoint\_ip\] Endpoint's IP address is discovered using DNS name resolution.
The name given here is fully qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1232001313211322-3320331330033020-2121120223221232-0133312101320332-2033032303030333-0023311101013021-0300012310331231-1023033220223132"></a>

<a id="canonical-2133033220023022-1123222110113130-1013103233122211-2223302113011302-3332101003220001-1102232333131002-1113123332032231-1220101300110021"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_ip` property

Type: `"string"`. Computed.

Exclusive with \[nfs\_endpoint\_dns\_name\] Endpoint is reachable at the given IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-0120031312202210-1213230232032302-3221130111010021-1123100131110111-2100101020033100-0310021321221020-0012013030020321-2020100103212223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-003.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-003.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--reference--group-003.md#canonical-2032231100232113-1121332300210032-2300122032301201-1022223222230332-0320022001320313-0322210021221233-1130002200302111-0110333133311303)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--fleet--reference--group-003.md#canonical-3330123332213122-0311333303020203-2100301133303321-2131232010222233-1100023113000033-3113311212320100-1102102332012100-0222103102231330)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token

<a id="canonical-0021033303010333-2212110030332110-1021101310013202-3130011001030221-1230101032033020-2010231023223021-1331300022303313-2002202112102001"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-2101310111211122-2232121101112312-0330130110130003-0202131130023112-0033323123033012-0331201330310003-2010120303213103-0012311313221010"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token`

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-2333102332103103-1110001013122020-1102332123133211-2133121132100030-1100202030011331-3323300101312320-3013011312302131-1020010311121322): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-0322013121102330-2231012313300203-1112311002132201-3321010010320213-1221030023112200-0021012100023330-0130023031203211-2322330031223302): complete subsection reference.

<a id="canonical-2333102332103103-1110001013122020-1102332123133211-2133121132100030-1100202030011331-3323300101312320-3013011312302131-1020010311121322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-003.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-003.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--reference--group-003.md#canonical-2032231100232113-1121332300210032-2300122032301201-1022223222230332-0320022001320313-0322210021221233-1130002200302111-0110333133311303)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--fleet--reference--group-003.md#canonical-3330123332213122-0311333303020203-2100301133303321-2131232010222233-1100023113000033-3113311212320100-1102102332012100-0222103102231330)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--fleet--reference--group-003.md#canonical-0120031312202210-1213230232032302-3221130111010021-1123100131110111-2100101020033100-0310021321221020-0012013030020321-2020100103212223)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info

<a id="canonical-1112133021203233-0120323301013121-2310320323113212-1213322002201232-1201210133330322-3031022123330112-3200321232300102-1322202130011321"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0131120121313122-3111101012331230-0030112000323212-0230103222312323-1222223211010323-2320132032120202-2312033021013230-3132100021002313"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info`

<a id="canonical-3200121032302100-1223033021102330-0212023000201132-2003102322122100-0121002101021123-0121131122210223-2000001100131110-3020130000303330"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2033022230311112-3332001132303112-0011201221103031-1212103232021201-1031313223211322-0231311213203213-1010110221330213-2213131120022331"></a>

<a id="canonical-1200122200202001-2022312301023310-0112310021103312-3012302310200112-1032021001200022-3001312023003321-1300302223132022-3111231022013312"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3333203001133002-1213102111310013-1302100000233223-2103221320222321-1201322011130331-1132032220331213-2012211200100232-0031323033000032"></a>

<a id="canonical-1312213332331300-0030322313021211-2131331313131332-3312231311211211-1132200103311032-1020231023003023-3002223031120013-3212033300213012"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0322013121102330-2231012313300203-1112311002132201-3321010010320213-1221030023112200-0021012100023330-0130023031203211-2322330031223302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-003.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-003.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--reference--group-003.md#canonical-2032231100232113-1121332300210032-2300122032301201-1022223222230332-0320022001320313-0322210021221233-1130002200302111-0110333133311303)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--fleet--reference--group-003.md#canonical-3330123332213122-0311333303020203-2100301133303321-2131232010222233-1100023113000033-3113311212320100-1102102332012100-0222103102231330)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--fleet--reference--group-003.md#canonical-0120031312202210-1213230232032302-3221130111010021-1123100131110111-2100101020033100-0310021321221020-0012013030020321-2020100103212223)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info

<a id="canonical-3003220022210311-0331322102200122-2321001333331321-2301120010110122-0021231013302303-3230323001313012-3210201221000130-2201031121321213"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1200332203033220-2220302322013012-3320223011310231-2130200120120012-1232020001120200-0121231332123213-0020320023311212-2001113311213020"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info`

<a id="canonical-0121230202312130-2223021112300133-1102032031212112-0322322113100211-1001030211122201-2011321022333332-2003312022331102-3330021111013212"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0230221020030020-3311113313312211-0101000331332133-3223033303212102-1311303002131123-2111313230332302-0132312032231113-2203102321221102"></a>

<a id="canonical-0132002002331111-0322322031112312-2332001022013100-2311023023312033-0310000020223112-3133320312030133-3330232303133311-0001221012321221"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1332201313101100-2300203232300003-3203012102133211-1111001020322331-3011023323013302-0102200332002030-0003133323233202-2131113031033200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_interface_list` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- storage_interface_list

<a id="canonical-2121301312201313-2003303021323200-1113203311322120-1012103002103230-2013220110320003-0220331131210223-0321010211132303-2320301002023322"></a>

Type: `"single"`. Computed.

Add all interfaces belonging to this fleet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3020030020103323-2021113002321103-0230130232123121-2130203122102320-0203211221113100-2102301301121223-1202002023303131-2032032310002233"></a>

### Direct properties for `storage_interface_list`

- [interfaces](data-sources--fleet--reference--group-003.md#canonical-2033322200112100-2130213320101310-2230111211133032-2103331223111023-0200333312020201-1020302212030332-0231123132030110-2203122202333330): complete subsection reference.

<a id="canonical-2033322200112100-2130213320101310-2230111211133032-2103331223111023-0200333312020201-1020302212030332-0231123132030110-2203122202333330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_interface_list.interfaces` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_interface_list](data-sources--fleet--reference--group-003.md#canonical-1332201313101100-2300203232300003-3203012102133211-1111001020322331-3011023323013302-0102200332002030-0003133323233202-2131113031033200)
- storage_interface_list.interfaces

<a id="canonical-1122201013213312-1213321213321112-2333130011001003-0003030300101031-2312332033001313-3003313032230033-1202330033301311-2111302313321313"></a>

Type: `"list"`. Computed.

Add all interfaces belonging to this fleet.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3203111313210223-2213332322103331-2200323020111133-0320332111212212-3210032223232301-2302213223023223-0200130131330111-1221121012220020"></a>

### Direct properties for `storage_interface_list.interfaces`

<a id="canonical-0212332102111200-3320032000321203-1202332311211102-0201332310320001-2123322132033202-0311313020113102-2231302303222121-1211123330131031"></a>

#### `storage_interface_list.interfaces.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3301020201020100-2010220120303110-0100133332112211-2322233313332201-2310030111130300-2030312131130013-2202002101112113-0302100130332202"></a>

<a id="canonical-3010031132122100-0130310130000010-3011022231321322-1330122032012312-1103203100312223-0133321220322023-0331303303321321-1230122231211111"></a>

#### `storage_interface_list.interfaces.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3101313223200101-0132311203200313-1322123133110022-3230123003300203-2312213320223020-0123013210323210-3233330203023100-3222222230220121"></a>

<a id="canonical-3133321030211233-1012112302110312-1123023332202321-3310222101223010-3310123203313301-3103121022213103-1003230203101202-0002103021133002"></a>

#### `storage_interface_list.interfaces.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- storage_static_routes

<a id="canonical-3012331233311111-0123111321131231-3213032311022211-3000231010022111-1001031021200012-0323022200021110-3333332210321120-3330100233233210"></a>

Type: `"single"`. Computed.

Configuration parameter for storage static routes.

Additional upstream details:

List of storage static routes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2010301101233302-3011020210223321-3013132323001000-3313321130303220-3213131020232232-2132301223210102-1202020111101132-3221101121000131"></a>

### Direct properties for `storage_static_routes`

- [storage_routes](data-sources--fleet--reference--group-003.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110): complete subsection reference.

<a id="canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-003.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- storage_static_routes.storage_routes

<a id="canonical-0002121312211212-2131010023231311-2303033332000203-0203113331203322-3113121222031102-0211010112002032-0330023132100311-1113010230012303"></a>

Type: `"list"`. Computed.

List of Static Routes. List of storage static routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1330222132012233-0313302200032001-1321332010223330-2222100013001112-2020301100213020-0133103120012020-1020231212120302-0330200323330211"></a>

### Direct properties for `storage_static_routes.storage_routes`

<a id="canonical-3310113203102321-0220331313021223-1210113113222021-0310011211133303-0111110313211010-0032313110102321-3010010021213202-2333111201101201"></a>

#### `storage_static_routes.storage_routes.attrs` property

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](data-sources--fleet--reference--group-003.md#canonical-2333322330302012-3111121113301032-1013103232322112-0231022000222301-2313120302310130-2020010130032321-0112300013211020-1302031132120021): complete subsection reference.

- [nexthop](data-sources--fleet--reference--group-003.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000): complete subsection reference.

- [subnets](data-sources--fleet--reference--group-004.md#canonical-3221013101022303-0031020003333031-1030012032322102-2001322333132322-0323001032121312-3112313131203323-0220202103032320-2230222202110303): complete subsection reference.

<a id="canonical-2333322330302012-3111121113301032-1013103232322112-0231022000222301-2313120302310130-2020010130032321-0112300013211020-1302031132120021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.labels` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-003.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-003.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- storage_static_routes.storage_routes.labels

<a id="canonical-3103023220123112-1331103110032301-1100200232101013-3201303231110003-3013312332030102-3302312322303003-1232210221131112-0212323010101333"></a>

Type: `"single"`. Computed.

Add Labels for this Static Route, these labels can be used in network policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.nexthop` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-003.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-003.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- storage_static_routes.storage_routes.nexthop

<a id="canonical-1332000133212021-1023230000201303-0003013230221303-0220302202132323-3232023313031201-2332132010133323-3301122210110313-3101120330300132"></a>

Type: `"single"`. Computed.

Nexthop. Identifies the next-hop for a route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3222203102100133-2303233202321200-3010121131100212-3111021310320120-2130121302112021-2332123302110000-0011013332222130-3113011201301131"></a>

### Direct properties for `storage_static_routes.storage_routes.nexthop`

- [interface](data-sources--fleet--reference--group-003.md#canonical-3000023121020110-0012320102123202-2331230130030331-1223220011332102-2133301101233130-3332220022023312-2032222003102112-3123121021332302): complete subsection reference.

- [nexthop_address](data-sources--fleet--reference--group-003.md#canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312): complete subsection reference.

<a id="canonical-1130201332320032-2211130120201322-1313001230111021-1213111203222323-3030233222333221-3213301320002103-1002302121200331-2111221210121322"></a>

<a id="canonical-1103213022311333-1312200331202303-0321220020300030-0001200331110012-3120003320103113-3111130010133333-0033313131212211-3313010111033222"></a>

#### `storage_static_routes.storage_routes.nexthop.type` property

Type: `"string"`. Computed.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Additional upstream details:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Use the specified address as
nexthop Use the network interface as nexthop Discard nexthop, used when attr type is Advertise Used
in VoltADN private virtual network.

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3000023121020110-0012320102123202-2331230130030331-1223220011332102-2133301101233130-3332220022023312-2032222003102112-3123121021332302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.nexthop.interface` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-003.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-003.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-003.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000)
- storage_static_routes.storage_routes.nexthop.interface

<a id="canonical-1101320030003212-2103123221332333-0010002001331313-1232300103011123-1223231022210033-2010030223001133-3203133330300112-1112003030322011"></a>

Type: `"list"`. Computed.

Nexthop is network interface when type is 'Network-Interface'.

Additional upstream details:

Nexthop is network interface when type is "Network-Interface"

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1303231313011203-1132102101010223-1032131132110232-3320223031131020-1331321033202121-3300130102233303-1033013011123331-2013010102132332"></a>

### Direct properties for `storage_static_routes.storage_routes.nexthop.interface`

<a id="canonical-0221330321302132-3103132202133002-2123302200102120-1201032132113300-2230333012002331-1033033010200322-0021210303103330-2011101200020230"></a>

#### `storage_static_routes.storage_routes.nexthop.interface.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2212023012122320-3321321103022033-1201010103312111-0331321313230123-1012322300122323-0302112321100020-0002130110323230-2012022121320113"></a>

<a id="canonical-0202210033323031-3301033030103102-3122103321100033-1220132111223031-0021331200320330-3121221020201002-1021331112320100-0201320122012112"></a>

#### `storage_static_routes.storage_routes.nexthop.interface.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1030310330232201-2233001002222303-2103102022310111-1311113102221330-0031330210101032-3223300030102220-2133220131110133-2211032323122200"></a>

<a id="canonical-1103210302220220-2211111013230121-0311132000130120-1322301232000013-2013301121032201-0330002133331211-2310303001110302-1130122113233301"></a>

#### `storage_static_routes.storage_routes.nexthop.interface.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1021030213203231-2301210022200121-3320002231001133-2101212100303013-0201113232130320-0023220100223220-1133002213000322-2022132303302202"></a>

<a id="canonical-1212113210300031-1001321132311132-3313210333133020-0211122000122211-0003122311101311-3300321312301000-2223110201120010-3203020232102010"></a>

#### `storage_static_routes.storage_routes.nexthop.interface.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1001311231020131-1013001312320310-0013210013303131-0001300123233331-3023023102332011-1223101103222230-1010102020011321-2101323011120121"></a>

<a id="canonical-1012301001021223-2203001122312112-3330113212113231-2013133201312323-3213002032200232-1231300012021020-2223112222232001-0022003111001011"></a>

#### `storage_static_routes.storage_routes.nexthop.interface.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.nexthop.nexthop_address` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-003.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-003.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-003.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000)
- storage_static_routes.storage_routes.nexthop.nexthop_address

<a id="canonical-0323212123132300-1112331232012132-3220213323231021-3003031110230301-0232111220311022-3123030133110010-2112312012122202-3121103331200021"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

<a id="canonical-3322100033331133-3121031220002323-2101131332222332-3203213031010312-1103012133030310-3300331201303033-2302011332230031-3223111223021113"></a>

### Direct properties for `storage_static_routes.storage_routes.nexthop.nexthop_address`

- [dual_stack](data-sources--fleet--reference--group-003.md#canonical-2301223133112311-2000202212230202-1202322323023311-0210223000310120-3222233322231210-1311332211322211-2310111023212312-3102211221223033): complete subsection reference.

- [IPv4](data-sources--fleet--reference--group-003.md#canonical-1203230330233330-0320331023002231-0012122030311100-3112313322120232-1203113223020201-3033112121033312-1031200202331201-2321200330101200): complete subsection reference.

- [IPv6](data-sources--fleet--reference--group-003.md#canonical-0300332322220010-0132002021012213-1323101311003021-1232221222232332-3223212333122100-3312321023311121-3323201230133223-2201330333332321): complete subsection reference.

<a id="canonical-2301223133112311-2000202212230202-1202322323023311-0210223000310120-3222233322231210-1311332211322211-2310111023212312-3102211221223033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-003.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-003.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-003.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-003.md#canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312)
- storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack

<a id="canonical-1103200233003002-1221133113200013-2203032010230202-1301330010003110-1130233112112303-3311033311031021-1223022331201021-2210133303232312"></a>

Type: `"single"`. Computed.

DualStackAddressType represents both IPv4 and IPv6 together.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1032131302021221-2200021221313020-3302100312000100-2021332100112023-0021222010011011-2332313021213302-2011232310311202-3133103123210012"></a>

### Direct properties for `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack`

- [IPv4](data-sources--fleet--reference--group-003.md#canonical-1211023231012201-1301320231100000-1321021223020323-3332012322121331-1302122322213022-2212230133103320-0320110133121011-2223212112002302): complete subsection reference.

- [IPv6](data-sources--fleet--reference--group-003.md#canonical-2011021111102013-3023110120202033-0331312021311220-1123332323023030-1220322032013232-2210203012032320-1210230133222310-0120020131233230): complete subsection reference.

<a id="canonical-1211023231012201-1301320231100000-1321021223020323-3332012322121331-1302122322213022-2212230133103320-0320110133121011-2223212112002302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-003.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-003.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-003.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-003.md#canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](data-sources--fleet--reference--group-003.md#canonical-2301223133112311-2000202212230202-1202322323023311-0210223000310120-3222233322231210-1311332211322211-2310111023212312-3102211221223033)
- storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-3011310122320321-2331020130230233-2322331013201032-2302311201113132-3030132023113010-1032003010212323-3210313113330220-0302321003223023"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Additional upstream details:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3030031111033213-1332312212200001-0200210201313230-1321200112322222-0132102211000102-1021003011321313-0102230100010310-3020312332123310"></a>

### Direct properties for `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4`

<a id="canonical-3030312222002332-2211112333101022-0322122013122221-3102032202132012-0120010001132030-0200313331123031-0200023233311101-2302132032023210"></a>

#### `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4.addr` property

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2011021111102013-3023110120202033-0331312021311220-1123332323023030-1220322032013232-2210203012032320-1210230133222310-0120020131233230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-003.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-003.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-003.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-003.md#canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](data-sources--fleet--reference--group-003.md#canonical-2301223133112311-2000202212230202-1202322323023311-0210223000310120-3222233322231210-1311332211322211-2310111023212312-3102211221223033)
- storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-2110212021130032-1202110122000102-2022332332133233-1021113013101033-0003310011220110-0300000112312130-1131012212223231-2010220331210113"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3000211000332223-0211021130310012-1203002021202332-3321013110210002-0221023231002103-1202233010100303-3211000012210101-2100232110202003"></a>

### Direct properties for `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6`

<a id="canonical-0031332122230123-2320102032122203-2003313022030331-1120323013330312-2023023211113331-1121100332012200-3000230100211031-1011302210221012"></a>

#### `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6.addr` property

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-1203230330233330-0320331023002231-0012122030311100-3112313322120232-1203113223020201-3033112121033312-1031200202331201-2321200330101200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-003.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-003.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-003.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-003.md#canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312)
- storage_static_routes.storage_routes.nexthop.nexthop_address.IPv4

<a id="canonical-0110331321002230-2010332312131131-2130132323111022-0211020101231312-1310131222211333-2131211122133231-1030010113211303-1111302000302131"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Additional upstream details:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3003200210300123-0332023300112331-0333310310011323-0000103232111132-3012131331321223-2002013000233320-3013333231133123-3310112002011331"></a>

### Direct properties for `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4`

<a id="canonical-3013233232200232-3131103301210110-3200030331021322-0210323303213331-0201202013202110-1212001331011012-3300010111111203-3131223130120313"></a>

#### `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4.addr` property

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0300332322220010-0132002021012213-1323101311003021-1232221222232332-3223212333122100-3312321023311121-3323201230133223-2201330333332321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-003.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-003.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-003.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-003.md#canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312)
- storage_static_routes.storage_routes.nexthop.nexthop_address.IPv6

<a id="canonical-2220130211322320-0302002013002220-1322021313011120-0103321002032111-1021101300231010-1302203202301020-3222203303312210-2031333123230212"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2220202332232211-2332201301301023-2012123011123221-3322101331303013-1323003320303323-3001023231113112-2023013130322110-3032210111312320"></a>

### Direct properties for `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6`

<a id="canonical-2300133101021302-1202013020231010-0113210321032102-2003313301230211-2320211110230220-3202012010223001-2220033111010303-3233131032032232"></a>

#### `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6.addr` property

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```
