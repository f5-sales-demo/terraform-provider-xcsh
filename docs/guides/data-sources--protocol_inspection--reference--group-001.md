---
page_title: "xcsh_protocol_inspection reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_inspection reference."
---

# xcsh_protocol_inspection reference

<a id="canonical-2031133223011203-2331031213330110-1111011320132233-1130032201323122-3030112030303121-2231033112022021-0213331300320023-3130013121123023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200000112311332-0031002032120010-2230112231102330-3332021121111033-1301121213103123-2213322201312021-1203002022211020-2311111233012201"></a>

## Property reference — Property reference / 313222101320 / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-1313301302113001-0102313332203012-2202123331012332-2332033323132031-2332220030232012-3021230003332301-2212111100302303-0210333132120310)
- Property reference

<a id="canonical-0320200031311300-0133233200221011-2212130302031312-2332113102021122-1132211313010131-3101033300332201-2122321013223312-3122133021203213"></a>

## Direct properties — Property reference / 313222101320 / 3

<a id="canonical-2103032132012223-3111332102222211-1100130313223213-2232303321030113-2323322200320212-2020311313112030-3010233010120212-0132223030301003"></a>

<a id="canonical-3013031120132021-2131302100200321-0312121232301101-1223232332331112-1201213030331223-1301223103202122-1230123301012213-1123121300211222"></a>

## action property — Property reference / 313222101320 / 4

Type: `"string"`. Computed.

\[Enum: ALLOW|DENY|DROP\] Action after inspection - ALLOW: Allow Allow traffic - DENY: Deny Throw
RST error for TCP and ICMP error for UDP - DROP: DROP Silently drop traffic. Possible values are
\`ALLOW\`, \`DENY\`, \`DROP\`. Defaults to \`ALLOW\`. Server applies default when omitted.

Upstream description:

Action after inspection

&#8203;- ALLOW: Allow

Allow traffic &#8203;- DENY: Deny

Throw RST error for TCP and ICMP error for UDP &#8203;- DROP: DROP

Silently drop traffic.

Receipt-pinned upstream constraints:

```json
{
  "default": "ALLOW",
  "enum": [
    "ALLOW",
    "DENY",
    "DROP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0002131000222110-1331200331320312-3233210023132333-0312102030133310-2011231103033111-2003210022321331-0112312010012313-2113100130301033"></a>

<a id="canonical-2031322221311303-0321010301122031-2322212022100121-0322002102002330-1303010110023033-3211232320013033-1213111333302300-1133220102223003"></a>

## annotations property — Property reference / 313222101320 / 5

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-2101121021102112-1223201012301011-2032313130223122-1032132332200121-2033111311022132-3010132232202203-0301223221200103-0331230220331013"></a>

<a id="canonical-2003203031013230-0222323231002031-2021023321302331-3232232112102121-0313033210212203-1122230303121100-0321121201113012-2210111231200013"></a>

## description property — Property reference / 313222101320 / 6

Type: `"string"`. Computed.

Description of the ProtocolInspection.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

- [enable_disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-1223020321120303-0303012222310120-1230131211300021-2121032111222012-3110102121333010-2132331121321033-0210312231020030-0323222302001133): complete subsection reference.

- [enable_disable_signatures](data-sources--protocol_inspection--reference--group-001.md#canonical-0102233113110032-1023220020302113-3101333100212310-2100211003310220-1111303010230120-2221223002230211-2333033021220230-1231021303030220): complete subsection reference.

<a id="canonical-3023213021221201-2111223103120003-0312200121133211-1120111212230203-0222002123313300-3211101230230103-3212323033113131-0103201320201010"></a>

<a id="canonical-3313300112022200-3211133310302020-0223021001013032-2032113221230010-0311123233102101-2321120103230021-2233212313102210-2233300201313001"></a>

## ID property — Property reference / 313222101320 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3000113013112103-2121013121322033-1212000301230002-3201200300220103-3200310110003133-2330131022223232-1122102200313232-3322310303112032"></a>

<a id="canonical-2023300232000332-0313332003122123-1030333113321012-2323232003121331-0103310020303223-0133120102121202-3210001000202330-3102332321203001"></a>

## labels property — Property reference / 313222101320 / 8

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="canonical-2232130302111323-0222021130013010-1223033110323001-0031020101000032-0030201313210311-3223302303233303-2120133301313022-0311000221102202"></a>

<a id="canonical-3202221233121311-0031302223001100-0011111003323112-1011131013323023-3003133031311333-3221323211021133-2211330232332200-2202021113002111"></a>

## name property — Property reference / 313222101320 / 9

Type: `"string"`. Required.

Name of the ProtocolInspection.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1012312231233213-1332212301123202-3301203132131311-1230320210020303-1303132130300222-2021230100131313-1333332200113213-3101031103203213"></a>

<a id="canonical-1323323222311002-0023013322203001-3020223033233213-2122110133312223-3023133123303031-0001001111231202-1111112233111030-0120200223312303"></a>

## namespace property — Property reference / 313222101320 / 10

Type: `"string"`. Required.

Namespace where the ProtocolInspection exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

<a id="canonical-2222301222030100-3031000102233213-2032223120231103-0032311030210010-0331312301212303-1031331210030110-3013101101112121-2002233202302213"></a>

## All schema paths — Property reference / 313222101320 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](data-sources--protocol_inspection--reference--group-001.md#canonical-2103032132012223-3111332102222211-1100130313223213-2232303321030113-2323322200320212-2020311313112030-3010233010120212-0132223030301003) |
| `annotations` | [annotations](data-sources--protocol_inspection--reference--group-001.md#canonical-0002131000222110-1331200331320312-3233210023132333-0312102030133310-2011231103033111-2003210022321331-0112312010012313-2113100130301033) |
| `description` | [description](data-sources--protocol_inspection--reference--group-001.md#canonical-2101121021102112-1223201012301011-2032313130223122-1032132332200121-2033111311022132-3010132232202203-0301223221200103-0331230220331013) |
| `enable_disable_compliance_checks` | [enable_disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-2113131223123011-2133313110133122-0113232210311122-1003111031311303-1220222010223323-3102021130122002-2023333123013232-3101013012300320) |
| `enable_disable_compliance_checks.disable_compliance_checks` | [enable_disable_compliance_checks.disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-0223032302102301-1133203312210121-0310013003102102-0200020312222132-2003102310000123-3020313011331231-2201233022131312-0133022212223000) |
| `enable_disable_compliance_checks.enable_compliance_checks` | [enable_disable_compliance_checks.enable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-3123222023210101-3113022200032322-3001131212212102-2323230321112033-2210330301311120-0100200322111322-1322203232033331-2222020013332323) |
| `enable_disable_compliance_checks.enable_compliance_checks.name` | [enable_disable_compliance_checks.enable_compliance_checks.name](data-sources--protocol_inspection--reference--group-001.md#canonical-2333211231022200-2331322111103013-2113112021210020-2130123331103312-1301110023002223-0112221323132312-1311001120230231-1111221132220232) |
| `enable_disable_compliance_checks.enable_compliance_checks.namespace` | [enable_disable_compliance_checks.enable_compliance_checks.namespace](data-sources--protocol_inspection--reference--group-001.md#canonical-2230120111333013-0221030000002011-0110021332012003-2100130112101103-0203330111133232-3323303111023123-2120022220333111-1120132232231113) |
| `enable_disable_compliance_checks.enable_compliance_checks.tenant` | [enable_disable_compliance_checks.enable_compliance_checks.tenant](data-sources--protocol_inspection--reference--group-001.md#canonical-2300020220213001-1213212023010013-3302033003133231-3121021303101102-3233032202211220-0111321012331021-0121133100331123-0022030223123112) |
| `enable_disable_signatures` | [enable_disable_signatures](data-sources--protocol_inspection--reference--group-001.md#canonical-0031131323032333-1300201233211220-2313323020233233-1120302131201022-3133333320331032-2201233301113322-1212322200331210-0110003131311001) |
| `enable_disable_signatures.disable_signature` | [enable_disable_signatures.disable_signature](data-sources--protocol_inspection--reference--group-001.md#canonical-0112203000031211-1123211012111200-1301120203301121-0121201231332033-2303330011121201-0122232212212230-2300211123002203-1221101223232031) |
| `enable_disable_signatures.enable_signature` | [enable_disable_signatures.enable_signature](data-sources--protocol_inspection--reference--group-001.md#canonical-2232330112333112-2331303121233111-1021020020023330-3133113212213300-2132310302120021-0233100223320201-0200300001232133-2000000133321102) |
| `id` | [ID](data-sources--protocol_inspection--reference--group-001.md#canonical-3023213021221201-2111223103120003-0312200121133211-1120111212230203-0222002123313300-3211101230230103-3212323033113131-0103201320201010) |
| `labels` | [labels](data-sources--protocol_inspection--reference--group-001.md#canonical-3000113013112103-2121013121322033-1212000301230002-3201200300220103-3200310110003133-2330131022223232-1122102200313232-3322310303112032) |
| `name` | [name](data-sources--protocol_inspection--reference--group-001.md#canonical-2232130302111323-0222021130013010-1223033110323001-0031020101000032-0030201313210311-3223302303233303-2120133301313022-0311000221102202) |
| `namespace` | [namespace](data-sources--protocol_inspection--reference--group-001.md#canonical-1012312231233213-1332212301123202-3301203132131311-1230320210020303-1303132130300222-2021230100131313-1333332200113213-3101031103203213) |

<a id="canonical-3220003313122130-1110330203110020-2111203302332023-2202320023003323-2210300101331102-3320202322012130-1110033211112023-2133133001011202"></a>

## Next pages — Property reference / 313222101320 / 12

- [enable_disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-1223020321120303-0303012222310120-1230131211300021-2121032111222012-3110102121333010-2132331121321033-0210312231020030-0323222302001133)
- [enable_disable_signatures](data-sources--protocol_inspection--reference--group-001.md#canonical-0102233113110032-1023220020302113-3101333100212310-2100211003310220-1111303010230120-2221223002230211-2333033021220230-1231021303030220)
- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-1313301302113001-0102313332203012-2202123331012332-2332033323132031-2332220030232012-3021230003332301-2212111100302303-0210333132120310)

<a id="canonical-1223020321120303-0303012222310120-1230131211300021-2121032111222012-3110102121333010-2132331121321033-0210312231020030-0323222302001133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013121013231232-2110002020320131-0210003121200110-3332020000111003-1113131101030122-3332300313133232-2031031133011323-2223020133002122"></a>

## enable_disable_compliance_checks — enable_disable_compliance_checks / 131011231123 / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-1313301302113001-0102313332203012-2202123331012332-2332033323132031-2332220030232012-3021230003332301-2212111100302303-0210333132120310)
- [Property reference](data-sources--protocol_inspection--reference--group-001.md#canonical-2031133223011203-2331031213330110-1111011320132233-1130032201323122-3030112030303121-2231033112022021-0213331300320023-3130013121123023)
- enable_disable_compliance_checks

<a id="canonical-2113131223123011-2133313110133122-0113232210311122-1003111031311303-1220222010223323-3102021130122002-2023333123013232-3101013012300320"></a>

Type: `"single"`. Computed.

Enable Disable Compliance Checks Choice.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-compliance_check_choice": "[\"disable_compliance_checks\",\"enable_compliance_checks\"]"
}
```

<a id="canonical-3102030031311310-2113102332220120-2101321132321120-2230222110301001-1203131200323003-2132000221120210-3233032133303032-0332201101212011"></a>

## Direct properties — enable_disable_compliance_checks / 131011231123 / 3

- [disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-2223121201201321-0110330133221330-1011320123000001-3032210122202023-0331311301322320-2221030020301103-2210000322000332-1203203331310030): complete subsection reference.

- [enable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-2011013222100102-2111123223313133-0312322310132200-2121210001003311-3231202300303102-2233011021321103-1220320211211222-3232113100331312): complete subsection reference.

<a id="canonical-3003013222330110-3021211131130230-3230133233201202-1103231033001120-0022322312220121-2000220233132113-0333130312303222-3111020211121312"></a>

## Next pages — enable_disable_compliance_checks / 131011231123 / 4

- [enable_disable_compliance_checks.disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-2223121201201321-0110330133221330-1011320123000001-3032210122202023-0331311301322320-2221030020301103-2210000322000332-1203203331310030)
- [enable_disable_compliance_checks.enable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-2011013222100102-2111123223313133-0312322310132200-2121210001003311-3231202300303102-2233011021321103-1220320211211222-3232113100331312)
- [Property reference](data-sources--protocol_inspection--reference--group-001.md#canonical-2031133223011203-2331031213330110-1111011320132233-1130032201323122-3030112030303121-2231033112022021-0213331300320023-3130013121123023)
- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-1313301302113001-0102313332203012-2202123331012332-2332033323132031-2332220030232012-3021230003332301-2212111100302303-0210333132120310)

<a id="canonical-2223121201201321-0110330133221330-1011320123000001-3032210122202023-0331311301322320-2221030020301103-2210000322000332-1203203331310030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022000201220333-3322031211012200-1111012031301203-0032302002011301-3011313022331312-2032212310313232-0020303213211313-1210113303022203"></a>

## enable_disable_compliance_checks.disable_compliance_checks — disable_compliance_checks / 120202232321 / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-1313301302113001-0102313332203012-2202123331012332-2332033323132031-2332220030232012-3021230003332301-2212111100302303-0210333132120310)
- [Property reference](data-sources--protocol_inspection--reference--group-001.md#canonical-2031133223011203-2331031213330110-1111011320132233-1130032201323122-3030112030303121-2231033112022021-0213331300320023-3130013121123023)
- [enable_disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-1223020321120303-0303012222310120-1230131211300021-2121032111222012-3110102121333010-2132331121321033-0210312231020030-0323222302001133)
- enable_disable_compliance_checks.disable_compliance_checks

<a id="canonical-0223032302102301-1133203312210121-0310013003102102-0200020312222132-2003102310000123-3020313011331231-2201233022131312-0133022212223000"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable compliance checks.

Upstream description:

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

<a id="canonical-1310022321030030-0000120113301112-0122033230330033-2002031001311121-1032310332232120-1200112021100210-3302023301331231-3022213122221122"></a>

## Direct properties — disable_compliance_checks / 120202232321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213123100231022-1033133223102002-0120031011103231-1233000010332330-1021203321113333-0200213113210113-2311013112220332-0123322130333330"></a>

## Next pages — disable_compliance_checks / 120202232321 / 4

- [enable_disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-1223020321120303-0303012222310120-1230131211300021-2121032111222012-3110102121333010-2132331121321033-0210312231020030-0323222302001133)
- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-1313301302113001-0102313332203012-2202123331012332-2332033323132031-2332220030232012-3021230003332301-2212111100302303-0210333132120310)

<a id="canonical-2011013222100102-2111123223313133-0312322310132200-2121210001003311-3231202300303102-2233011021321103-1220320211211222-3232113100331312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302031313232120-2002132313002003-2131201230112221-0233312013032112-2323130131332001-1220023003103033-0221231111010111-3122321233231013"></a>

## enable_disable_compliance_checks.enable_compliance_checks — enable_compliance_checks / 200122113033 / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-1313301302113001-0102313332203012-2202123331012332-2332033323132031-2332220030232012-3021230003332301-2212111100302303-0210333132120310)
- [Property reference](data-sources--protocol_inspection--reference--group-001.md#canonical-2031133223011203-2331031213330110-1111011320132233-1130032201323122-3030112030303121-2231033112022021-0213331300320023-3130013121123023)
- [enable_disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-1223020321120303-0303012222310120-1230131211300021-2121032111222012-3110102121333010-2132331121321033-0210312231020030-0323222302001133)
- enable_disable_compliance_checks.enable_compliance_checks

<a id="canonical-3123222023210101-3113022200032322-3001131212212102-2323230321112033-2210330301311120-0100200322111322-1322203232033331-2222020013332323"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-0023320120231323-2212012322221230-0333101122221232-2233123032013010-3231020201222033-0301212122022000-3100011223221312-3210220310120002"></a>

## Direct properties — enable_compliance_checks / 200122113033 / 3

<a id="canonical-2333211231022200-2331322111103013-2113112021210020-2130123331103312-1301110023002223-0112221323132312-1311001120230231-1111221132220232"></a>

<a id="canonical-0012033310321312-3310211203201031-0031333010103000-0020110210101011-0101131330332303-0213121311001301-0013221022110213-0012301232101231"></a>

## name property — enable_compliance_checks / 200122113033 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2230120111333013-0221030000002011-0110021332012003-2100130112101103-0203330111133232-3323303111023123-2120022220333111-1120132232231113"></a>

<a id="canonical-3022012232320103-1022000232223122-3011333332112200-0333332222122223-2232112030013220-0230222223322221-3023311201012023-2002100101213231"></a>

## namespace property — enable_compliance_checks / 200122113033 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-2300020220213001-1213212023010013-3302033003133231-3121021303101102-3233032202211220-0111321012331021-0121133100331123-0022030223123112"></a>

<a id="canonical-3223132301102200-2232113201213031-0030301231300222-2123000011020223-2202033012000331-3121013321220132-3203112022310130-1231202210322122"></a>

## tenant property — enable_compliance_checks / 200122113033 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2320111320001120-0131020330303021-1031303103212112-0322201101110012-2013212102132023-1322002303312320-2120231012132310-2311212211123233"></a>

## Next pages — enable_compliance_checks / 200122113033 / 7

- [enable_disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-1223020321120303-0303012222310120-1230131211300021-2121032111222012-3110102121333010-2132331121321033-0210312231020030-0323222302001133)
- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-1313301302113001-0102313332203012-2202123331012332-2332033323132031-2332220030232012-3021230003332301-2212111100302303-0210333132120310)

<a id="canonical-0102233113110032-1023220020302113-3101333100212310-2100211003310220-1111303010230120-2221223002230211-2333033021220230-1231021303030220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230032112120313-3230121220100032-1023020313303231-3313013002201333-1101210302202121-3002022333313223-3232223112301213-1233031330302122"></a>

## enable_disable_signatures — enable_disable_signatures / 030013213130 / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-1313301302113001-0102313332203012-2202123331012332-2332033323132031-2332220030232012-3021230003332301-2212111100302303-0210333132120310)
- [Property reference](data-sources--protocol_inspection--reference--group-001.md#canonical-2031133223011203-2331031213330110-1111011320132233-1130032201323122-3030112030303121-2231033112022021-0213331300320023-3130013121123023)
- enable_disable_signatures

<a id="canonical-0031131323032333-1300201233211220-2313323020233233-1120302131201022-3133333320331032-2201233301113322-1212322200331210-0110003131311001"></a>

Type: `"single"`. Computed.

Configuration parameter for enable disable signatures.

Upstream description:

Enable Disable Signature Choice.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-signature_choice": "[\"disable_signature\",\"enable_signature\"]"
}
```

<a id="canonical-1312000330300211-2201121300311110-0333000311000121-1130130010111221-0103133011330031-1023221201302031-0011002313133021-0010220100013302"></a>

## Direct properties — enable_disable_signatures / 030013213130 / 3

- [disable_signature](data-sources--protocol_inspection--reference--group-001.md#canonical-1301321210100300-0112022030021332-1221212023303210-0100002023122113-3112022003132131-0031110211122122-0333020101001233-0013102311111200): complete subsection reference.

- [enable_signature](data-sources--protocol_inspection--reference--group-001.md#canonical-3310301221322010-0202003312123020-3301312302320100-2331022203021113-3212123131210323-2030133132330212-1010000100132123-0000130213333321): complete subsection reference.

<a id="canonical-2011121033323323-0000210213133331-2202200112303022-3111100203323100-0332020131210022-2312021120200210-3300220202333011-0012332032302111"></a>

## Next pages — enable_disable_signatures / 030013213130 / 4

- [enable_disable_signatures.disable_signature](data-sources--protocol_inspection--reference--group-001.md#canonical-1301321210100300-0112022030021332-1221212023303210-0100002023122113-3112022003132131-0031110211122122-0333020101001233-0013102311111200)
- [enable_disable_signatures.enable_signature](data-sources--protocol_inspection--reference--group-001.md#canonical-3310301221322010-0202003312123020-3301312302320100-2331022203021113-3212123131210323-2030133132330212-1010000100132123-0000130213333321)
- [Property reference](data-sources--protocol_inspection--reference--group-001.md#canonical-2031133223011203-2331031213330110-1111011320132233-1130032201323122-3030112030303121-2231033112022021-0213331300320023-3130013121123023)
- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-1313301302113001-0102313332203012-2202123331012332-2332033323132031-2332220030232012-3021230003332301-2212111100302303-0210333132120310)

<a id="canonical-1301321210100300-0112022030021332-1221212023303210-0100002023122113-3112022003132131-0031110211122122-0333020101001233-0013102311111200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031322000222213-0202000103121123-0133103111332202-2231200200211232-1303210333303320-0322111011120301-3030030132213130-3333213311001332"></a>

## enable_disable_signatures.disable_signature — disable_signature / 031332333301 / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-1313301302113001-0102313332203012-2202123331012332-2332033323132031-2332220030232012-3021230003332301-2212111100302303-0210333132120310)
- [Property reference](data-sources--protocol_inspection--reference--group-001.md#canonical-2031133223011203-2331031213330110-1111011320132233-1130032201323122-3030112030303121-2231033112022021-0213331300320023-3130013121123023)
- [enable_disable_signatures](data-sources--protocol_inspection--reference--group-001.md#canonical-0102233113110032-1023220020302113-3101333100212310-2100211003310220-1111303010230120-2221223002230211-2333033021220230-1231021303030220)
- enable_disable_signatures.disable_signature

<a id="canonical-0112203000031211-1123211012111200-1301120203301121-0121201231332033-2303330011121201-0122232212212230-2300211123002203-1221101223232031"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable signature.

Upstream description:

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

<a id="canonical-1113000033103313-2302123123223310-2100210320321220-0231122230112020-2300222221121203-3211303110303112-0100202220232033-3200110231212222"></a>

## Direct properties — disable_signature / 031332333301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000321220221112-1220020212133132-0210131231223331-0122301020132300-1130301001000122-2101001123132322-0123322213003313-0120131102022330"></a>

## Next pages — disable_signature / 031332333301 / 4

- [enable_disable_signatures](data-sources--protocol_inspection--reference--group-001.md#canonical-0102233113110032-1023220020302113-3101333100212310-2100211003310220-1111303010230120-2221223002230211-2333033021220230-1231021303030220)
- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-1313301302113001-0102313332203012-2202123331012332-2332033323132031-2332220030232012-3021230003332301-2212111100302303-0210333132120310)

<a id="canonical-3310301221322010-0202003312123020-3301312302320100-2331022203021113-3212123131210323-2030133132330212-1010000100132123-0000130213333321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010030223213323-0221121130312032-2210031113132303-0221211203100220-2032313110322010-1121123331111102-2203120220232330-0321111222300032"></a>

## enable_disable_signatures.enable_signature — enable_signature / 100230013221 / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-1313301302113001-0102313332203012-2202123331012332-2332033323132031-2332220030232012-3021230003332301-2212111100302303-0210333132120310)
- [Property reference](data-sources--protocol_inspection--reference--group-001.md#canonical-2031133223011203-2331031213330110-1111011320132233-1130032201323122-3030112030303121-2231033112022021-0213331300320023-3130013121123023)
- [enable_disable_signatures](data-sources--protocol_inspection--reference--group-001.md#canonical-0102233113110032-1023220020302113-3101333100212310-2100211003310220-1111303010230120-2221223002230211-2333033021220230-1231021303030220)
- enable_disable_signatures.enable_signature

<a id="canonical-2232330112333112-2331303121233111-1021020020023330-3133113212213300-2132310302120021-0233100223320201-0200300001232133-2000000133321102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable signature.

Upstream description:

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

<a id="canonical-3101133302013310-2002011212330301-2312331120012223-3003111103303030-2211122022101003-2022300230103021-2020330233230320-0121320303011312"></a>

## Direct properties — enable_signature / 100230013221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212300133330103-2233332101133220-2332313100302012-1212120000223220-3323212231011023-0220222123030233-1210103212230233-2010021323030321"></a>

## Next pages — enable_signature / 100230013221 / 4

- [enable_disable_signatures](data-sources--protocol_inspection--reference--group-001.md#canonical-0102233113110032-1023220020302113-3101333100212310-2100211003310220-1111303010230120-2221223002230211-2333033021220230-1231021303030220)
- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-1313301302113001-0102313332203012-2202123331012332-2332033323132031-2332220030232012-3021230003332301-2212111100302303-0210333132120310)
