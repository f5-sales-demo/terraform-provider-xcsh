---
page_title: "xcsh_api_testing reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_testing reference."
---

# xcsh_api_testing reference

<a id="canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- Property reference

<a id="canonical-2223332223131033-2310130033330002-2222320333011001-0032312120023201-1100021301211323-1102023301202302-1103121313323030-3223112223121012"></a>

### Direct properties for `xcsh_api_testing`

<a id="canonical-2302311100111320-1032001131210332-3331310231013330-0023211201332130-2333003323033303-0122002132300212-3220132211303122-2132323032323323"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

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

<a id="canonical-2013123111010232-1321211310120312-2320310030300212-0030013202210022-2121221232112033-3200200200201303-0131202323031131-2203122311223301"></a>

<a id="canonical-2133022301101020-3022310112220132-1031023300033010-0331231102132033-0332030311230102-3302220220301231-3321213211011003-0302321312031133"></a>

#### `custom_header_value` property

Type: `"string"`. Optional, Computed.

Add x-F5-API-testing-identifier header value to prevent security flags on API testing traffic.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0233001221031100-1133122201112111-1201203211123102-0000132033311101-3333132331200201-2221203231110130-0113212020020120-3112012030300101"></a>

<a id="canonical-0200013021013210-2210130313010220-2003313001232123-2202303121332122-2011010111002120-1321202013011002-0122022110100200-2231023002333021"></a>

#### `description` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3003102203321013-0332100312102310-2303022013000033-1203200313230213-1211221211333300-1301230332330030-2333222033300101-2210230330003230"></a>

<a id="canonical-1110201212011232-1230021020303233-0011222013203130-3301313212001311-3210323311220112-1211021321002231-1200002130001003-1133020133302011"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

A value of true will administratively disable the object.

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

- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303): complete subsection reference.

- [every_day](resources--api_testing--reference--group-001.md#canonical-0110133330203132-0300210311112100-1220023032121103-3000313102030220-3123012332002301-1030201222002031-3231121132130330-1120323230220100): complete subsection reference.

- [every_month](resources--api_testing--reference--group-001.md#canonical-1323120202200200-3120122313023333-2112000223300001-2210110033112032-1033031302013222-0202010110312210-3111003132122233-2102132121222333): complete subsection reference.

- [every_week](resources--api_testing--reference--group-001.md#canonical-3003302130133212-2030223321333013-0223331221002310-0112202233231032-2022332023223330-0202032113013331-2112123012003322-0233330122012011): complete subsection reference.

<a id="canonical-2231020223003112-1221021111021223-3321013033012123-1122311213013021-3212332001201203-3133302111330131-1230112011233302-3113231023210131"></a>

<a id="canonical-3210032222021201-2131312331223023-3220030332332121-2131222011313030-0312332103130311-0322000133222020-3200131221301311-3233321011010102"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0212331333123201-3132021210010300-3002203030221202-0232100302120121-2000013131321122-0022101000010131-1111012112211210-1130031021220121"></a>

<a id="canonical-3212112113320132-3203133301310332-1313022031211021-0000300103110031-2200303201121201-3120102033123111-2323322022311032-1120303210201331"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

Additional upstream details:

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

<a id="canonical-2301002323302122-0323030233000101-2103113110220230-2203220012102200-2130230310012132-3303313233033220-0130302022001030-2022031220001232"></a>

<a id="canonical-3110331223302233-2320023033210130-0111321100121122-1003300003003333-2200320333233203-0200103230312211-0122032310032020-0322003021012331"></a>

#### `name` property

Type: `"string"`. Required.

Name of the API Testing. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NameValidator(),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1011023133110223-2100020022033330-1312011020302012-3020122200313332-0023122301102121-2120212222313201-3333032310323311-2313133321131303"></a>

<a id="canonical-0213102021020203-2111012011331333-1233100331131211-0322130012033331-1100110201102000-1221010122032030-2200113002330033-3102330310330331"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the API Testing is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [timeouts](resources--api_testing--reference--group-001.md#canonical-2220221222303221-1321122301102101-3213032112213022-2113023312203202-0313020120112103-0030330120312111-1212213100132010-0212222120301330): complete subsection reference.

<a id="canonical-1033200220201222-1131200323330110-1110202032032331-0031010221302013-2023002223321331-1331201030133133-0013001333223100-0133233312323332"></a>

### All schema paths for `xcsh_api_testing`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--api_testing--reference--group-001.md#canonical-2302311100111320-1032001131210332-3331310231013330-0023211201332130-2333003323033303-0122002132300212-3220132211303122-2132323032323323) |
| `custom_header_value` | [custom_header_value](resources--api_testing--reference--group-001.md#canonical-2013123111010232-1321211310120312-2320310030300212-0030013202210022-2121221232112033-3200200200201303-0131202323031131-2203122311223301) |
| `description` | [description](resources--api_testing--reference--group-001.md#canonical-0233001221031100-1133122201112111-1201203211123102-0000132033311101-3333132331200201-2221203231110130-0113212020020120-3112012030300101) |
| `disable` | [disable](resources--api_testing--reference--group-001.md#canonical-3003102203321013-0332100312102310-2303022013000033-1203200313230213-1211221211333300-1301230332330030-2333222033300101-2210230330003230) |
| `domains` | [domains](resources--api_testing--reference--group-001.md#canonical-3000110130032222-0313221110233203-1303200210100200-0231012120310000-3213033233233120-1222133223013020-0123311303013200-2023200021113202) |
| `domains.allow_destructive_methods` | [domains.allow_destructive_methods](resources--api_testing--reference--group-001.md#canonical-1030032120310200-2112320220001031-2011333132301001-0212212021210220-2203221311103001-1322312210000333-2332132032123030-0200130203112113) |
| `domains.credentials` | [domains.credentials](resources--api_testing--reference--group-001.md#canonical-3311333013222310-0300111023033313-3330331212211301-1001322003233102-1320020330233322-2230022130303123-3020132102323001-3021030020112123) |
| `domains.credentials.admin` | [domains.credentials.admin](resources--api_testing--reference--group-001.md#canonical-1220312000223131-2301201313301310-3332023302312120-2020012101313223-2120221122323132-1011121033023302-2013001000323131-3113313122102003) |
| `domains.credentials.api_key` | [domains.credentials.api_key](resources--api_testing--reference--group-001.md#canonical-2111233032210310-1220000103021203-3131111131102010-1102221032232101-0100011000120113-2233202311033220-2321302232212011-0011131102200111) |
| `domains.credentials.api_key.key` | [domains.credentials.api_key.key](resources--api_testing--reference--group-001.md#canonical-2310213331321232-3202211101130311-1113010112131032-2020332031221133-1322112332132302-1203131332002323-1202011232103301-0312330312103212) |
| `domains.credentials.api_key.value` | [domains.credentials.api_key.value](resources--api_testing--reference--group-001.md#canonical-3132201000122012-0301303132300010-0110302201223210-1123103030232023-0220013301023102-1222223030303000-0131012203231230-2223233131211112) |
| `domains.credentials.api_key.value.blindfold_secret_info` | [domains.credentials.api_key.value.blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-0231333101312333-0222002002103312-3123322001111030-3010212202200221-2302000132122232-1110330110011230-0112223312102322-0302103302003213) |
| `domains.credentials.api_key.value.blindfold_secret_info.decryption_provider` | [domains.credentials.api_key.value.blindfold_secret_info.decryption_provider](resources--api_testing--reference--group-001.md#canonical-0312021312020003-1031320122003202-3313030132311200-1113221103310120-2321013330033210-3110110121211220-0001322010301000-3223332300230010) |
| `domains.credentials.api_key.value.blindfold_secret_info.location` | [domains.credentials.api_key.value.blindfold_secret_info.location](resources--api_testing--reference--group-001.md#canonical-2013323221100000-2220023101131132-2023212032110011-3213302213012311-0232230001212022-3320311003030133-1020302200211023-0330030132310002) |
| `domains.credentials.api_key.value.blindfold_secret_info.store_provider` | [domains.credentials.api_key.value.blindfold_secret_info.store_provider](resources--api_testing--reference--group-001.md#canonical-0332301122333133-0010210130100021-2320233103000210-1322231031211222-3231111201211211-0003022311130100-0220210321203101-2022030231100232) |
| `domains.credentials.api_key.value.clear_secret_info` | [domains.credentials.api_key.value.clear_secret_info](resources--api_testing--reference--group-001.md#canonical-2323010030302211-0110300112312201-0030231122002230-3211022110322100-1321322201332220-0212311201320220-0201222111031022-3103233231213321) |
| `domains.credentials.api_key.value.clear_secret_info.provider_ref` | [domains.credentials.api_key.value.clear_secret_info.provider_ref](resources--api_testing--reference--group-001.md#canonical-2233003112112221-2310111230021331-3031110031010013-0231322300233221-1313232221222223-0211223222310011-1113032311031021-0003022010321303) |
| `domains.credentials.api_key.value.clear_secret_info.url` | [domains.credentials.api_key.value.clear_secret_info.url](resources--api_testing--reference--group-001.md#canonical-1201120230130310-2132103333330110-1313323330023321-3221101123322123-3121200001313111-2102230011113332-0123110313301113-2313032222002310) |
| `domains.credentials.basic_auth` | [domains.credentials.basic_auth](resources--api_testing--reference--group-001.md#canonical-3122303112202130-0102311332101203-2232200131330032-1003002231001330-0313103112112033-1111100022320132-1013003203331021-3223300210021232) |
| `domains.credentials.basic_auth.password` | [domains.credentials.basic_auth.password](resources--api_testing--reference--group-001.md#canonical-2311300030110223-0233211303231331-0030223302132022-0120212030333110-2200100130021122-0202200033021200-0000320021213020-1111200203313022) |
| `domains.credentials.basic_auth.password.blindfold_secret_info` | [domains.credentials.basic_auth.password.blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-3010021321300131-2133222121303301-3310133111220102-2101023033111223-3211011112012313-1003023312120121-0102230230221030-1002133203111001) |
| `domains.credentials.basic_auth.password.blindfold_secret_info.decryption_provider` | [domains.credentials.basic_auth.password.blindfold_secret_info.decryption_provider](resources--api_testing--reference--group-001.md#canonical-1310022111302313-0200103123011332-2220003212221323-1323321323031303-0011112310113303-3023030231111003-3012022133212130-2111321131110321) |
| `domains.credentials.basic_auth.password.blindfold_secret_info.location` | [domains.credentials.basic_auth.password.blindfold_secret_info.location](resources--api_testing--reference--group-001.md#canonical-1110303202121322-0030233333133302-2011312230021332-2130013220031323-0320020102021303-0111221332110100-2121030110231201-2110020003022132) |
| `domains.credentials.basic_auth.password.blindfold_secret_info.store_provider` | [domains.credentials.basic_auth.password.blindfold_secret_info.store_provider](resources--api_testing--reference--group-001.md#canonical-2201102100223202-0132331002121320-3113222003030300-3122030031332231-3131012332112211-0210113303213310-2321213110010310-3212003123013013) |
| `domains.credentials.basic_auth.password.clear_secret_info` | [domains.credentials.basic_auth.password.clear_secret_info](resources--api_testing--reference--group-001.md#canonical-1000233013300302-1211012333001001-1202223303202122-3133201312303232-3312100220200031-0302323000233010-0231222211101212-1031312333221120) |
| `domains.credentials.basic_auth.password.clear_secret_info.provider_ref` | [domains.credentials.basic_auth.password.clear_secret_info.provider_ref](resources--api_testing--reference--group-001.md#canonical-3231021021311021-3220220131012212-1033031320122320-0312122323312200-1112220213233332-3213320130310021-3031302311212010-3020303321121303) |
| `domains.credentials.basic_auth.password.clear_secret_info.url` | [domains.credentials.basic_auth.password.clear_secret_info.url](resources--api_testing--reference--group-001.md#canonical-1011320132000121-3130021032032320-2012131130111110-0221001202310002-2220011032300001-3010023111212222-3131101133100013-3331330232311333) |
| `domains.credentials.basic_auth.user` | [domains.credentials.basic_auth.user](resources--api_testing--reference--group-001.md#canonical-2011331313211001-0313101210220211-3212120330230031-2231102023100303-0102323103103232-2310122003221213-1122113331023321-3302011201330320) |
| `domains.credentials.bearer_token` | [domains.credentials.bearer_token](resources--api_testing--reference--group-001.md#canonical-2200312201130130-3022212231302121-2110033020202131-1000003303122330-3100021222103021-3230021221221001-2301010321302023-2113101330101101) |
| `domains.credentials.bearer_token.token` | [domains.credentials.bearer_token.token](resources--api_testing--reference--group-001.md#canonical-3132333000302201-2302230231232102-3032202001022232-3100330232033212-2313022123002212-3220200331320231-3033311113121313-0033123201322131) |
| `domains.credentials.bearer_token.token.blindfold_secret_info` | [domains.credentials.bearer_token.token.blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-2233002030222021-1222133332301232-3212113303120010-0103220102133032-0221111122121132-0302322203012212-2123030323232113-1010023313101103) |
| `domains.credentials.bearer_token.token.blindfold_secret_info.decryption_provider` | [domains.credentials.bearer_token.token.blindfold_secret_info.decryption_provider](resources--api_testing--reference--group-001.md#canonical-1112031022022120-2332030002122022-0111312002132232-1013002312112303-3120230022133210-3312202010011312-2202121210120103-1000113311032133) |
| `domains.credentials.bearer_token.token.blindfold_secret_info.location` | [domains.credentials.bearer_token.token.blindfold_secret_info.location](resources--api_testing--reference--group-001.md#canonical-1313031130331111-1322221332201321-1101223323112032-3231032212122130-2202230110020010-2333220100233301-1113110202300031-2213103001013113) |
| `domains.credentials.bearer_token.token.blindfold_secret_info.store_provider` | [domains.credentials.bearer_token.token.blindfold_secret_info.store_provider](resources--api_testing--reference--group-001.md#canonical-2321323320003011-2032223000210220-2132213331111010-0013133100200322-0131133302310112-1112121200120032-3133001130112322-1223102030333311) |
| `domains.credentials.bearer_token.token.clear_secret_info` | [domains.credentials.bearer_token.token.clear_secret_info](resources--api_testing--reference--group-001.md#canonical-3120300131112200-3001123312302020-3011101001333232-3223033020031230-0030323312102211-1300101001231233-3013122101012221-1010210210330031) |
| `domains.credentials.bearer_token.token.clear_secret_info.provider_ref` | [domains.credentials.bearer_token.token.clear_secret_info.provider_ref](resources--api_testing--reference--group-001.md#canonical-3303033012002130-3013221100101133-0031032031320133-2201022203213221-2133002023211001-3212003012000011-3221311202121332-2013210332033202) |
| `domains.credentials.bearer_token.token.clear_secret_info.url` | [domains.credentials.bearer_token.token.clear_secret_info.url](resources--api_testing--reference--group-001.md#canonical-3221212313202310-2103212102102101-1100011133300212-2110233013100331-3311002030131112-3203312002133010-3132123132021322-2302221131230332) |
| `domains.credentials.credential_name` | [domains.credentials.credential_name](resources--api_testing--reference--group-001.md#canonical-1102223000330210-2100312111302310-0210333102013033-1031220113300031-1131221101212100-2323202323221231-0011331130330322-3111323132102032) |
| `domains.credentials.login_endpoint` | [domains.credentials.login_endpoint](resources--api_testing--reference--group-001.md#canonical-0023021021120100-2110130121213311-1102301201200011-2012101220002312-3230303331030323-0003302130001231-0022020202021101-1231101000212223) |
| `domains.credentials.login_endpoint.json_payload` | [domains.credentials.login_endpoint.json_payload](resources--api_testing--reference--group-001.md#canonical-3322221202121210-1230011331220102-3130101311201330-3302112001231110-0212021111320230-0013233230032201-3132222021013230-1100113320213113) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-0022301033221322-1331001231330002-0201000031032332-1100030332333220-1000020321213020-3013331033300332-3000030030223131-1123002300102312) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.decryption_provider` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info.decryption_provider](resources--api_testing--reference--group-001.md#canonical-3111333000102333-3100133322331133-3333230233313200-1001030333021003-1103323010111333-1033210230002301-1002120321222210-2103210223002123) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.location` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info.location](resources--api_testing--reference--group-001.md#canonical-3002021113312213-2322322121221231-2002313000211002-1020303330120310-0122010133120202-0020120223233221-3111220021211131-2011031022212130) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.store_provider` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info.store_provider](resources--api_testing--reference--group-001.md#canonical-3110323133333322-3200213313033010-2312102311033011-0013002202100211-3122102120002032-3221223332020103-2221231130222011-2203131100201032) |
| `domains.credentials.login_endpoint.json_payload.clear_secret_info` | [domains.credentials.login_endpoint.json_payload.clear_secret_info](resources--api_testing--reference--group-001.md#canonical-3012000301330231-2003220130210212-2112301110122323-0102032033023112-3133301223112003-2323302313030211-3002333111222210-1001311101320220) |
| `domains.credentials.login_endpoint.json_payload.clear_secret_info.provider_ref` | [domains.credentials.login_endpoint.json_payload.clear_secret_info.provider_ref](resources--api_testing--reference--group-001.md#canonical-2132310033330301-0332033120103232-1011200002233312-1320010210211233-2223100112030123-3233321020233320-0230131201001210-0222312031302232) |
| `domains.credentials.login_endpoint.json_payload.clear_secret_info.url` | [domains.credentials.login_endpoint.json_payload.clear_secret_info.url](resources--api_testing--reference--group-001.md#canonical-3010212111313123-1000131012112013-0220202000312111-0220132301203111-1303201301311233-0322210120003100-3130133323321310-1213130202222302) |
| `domains.credentials.login_endpoint.method` | [domains.credentials.login_endpoint.method](resources--api_testing--reference--group-001.md#canonical-2322000012311221-0333310302123313-2300033020212301-0331232120103221-3230310310002012-1013122011312223-1003200210003232-3020021312011033) |
| `domains.credentials.login_endpoint.path` | [domains.credentials.login_endpoint.path](resources--api_testing--reference--group-001.md#canonical-3122102233310321-0002311133303112-1313031200330330-1023113210002023-1012131220032121-2123032111223000-2221001200303213-3311221232210330) |
| `domains.credentials.login_endpoint.token_response_key` | [domains.credentials.login_endpoint.token_response_key](resources--api_testing--reference--group-001.md#canonical-3021202021010310-3013212013213103-1003330030303301-3001222023023223-1101210113111210-2103112102223103-1303122322301012-3131201310323101) |
| `domains.credentials.standard` | [domains.credentials.standard](resources--api_testing--reference--group-001.md#canonical-1220302023112021-3021330122302211-2033231201232200-3223111320101321-1330330211230320-3332213002211321-3020330032011100-2010200002202223) |
| `domains.domain` | [domains.domain](resources--api_testing--reference--group-001.md#canonical-3021301113100320-2332211312021033-3032003132120010-1203013113232323-1123110013103030-3333222030033211-1201020213033300-1020023330130333) |
| `every_day` | [every_day](resources--api_testing--reference--group-001.md#canonical-1121033022331320-3113200120231112-0200210231013322-1001130230313322-0101131202323100-1010102013333312-3200001121112001-3321123230220023) |
| `every_month` | [every_month](resources--api_testing--reference--group-001.md#canonical-1002110121321233-0110021120131230-3122220200001211-1010231200012330-0132230332113111-3333011132102112-1211301332122332-1310221223321122) |
| `every_week` | [every_week](resources--api_testing--reference--group-001.md#canonical-0331001330211323-3330113021323212-1212213000000300-0322101000233222-1110122200222011-1202123110101113-1310001322111122-1231101130301111) |
| `id` | [ID](resources--api_testing--reference--group-001.md#canonical-2231020223003112-1221021111021223-3321013033012123-1122311213013021-3212332001201203-3133302111330131-1230112011233302-3113231023210131) |
| `labels` | [labels](resources--api_testing--reference--group-001.md#canonical-0212331333123201-3132021210010300-3002203030221202-0232100302120121-2000013131321122-0022101000010131-1111012112211210-1130031021220121) |
| `name` | [name](resources--api_testing--reference--group-001.md#canonical-2301002323302122-0323030233000101-2103113110220230-2203220012102200-2130230310012132-3303313233033220-0130302022001030-2022031220001232) |
| `namespace` | [namespace](resources--api_testing--reference--group-001.md#canonical-1011023133110223-2100020022033330-1312011020302012-3020122200313332-0023122301102121-2120212222313201-3333032310323311-2313133321131303) |
| `timeouts` | [timeouts](resources--api_testing--reference--group-001.md#canonical-2311001321033301-3302013302233100-2122001203300033-0231332233200013-2102102121213121-1323033121121100-1231033211033313-1021112011320303) |
| `timeouts.create` | [timeouts.create](resources--api_testing--reference--group-001.md#canonical-2131132300301123-1031322011120000-2333023122010130-2122122031013313-3333011310303123-1003000323200311-1113321002031120-2102303322003021) |
| `timeouts.delete` | [timeouts.delete](resources--api_testing--reference--group-001.md#canonical-3301303330220310-3300322230233022-3103123111210223-2312110221011131-1013230220331200-1013321313120002-2312121102302303-1111212011023322) |
| `timeouts.read` | [timeouts.read](resources--api_testing--reference--group-001.md#canonical-2032331201301301-1200233000130231-0010130320122010-3210200133121001-3332233021331002-2100212000011300-1211212221312021-3210130231113010) |
| `timeouts.update` | [timeouts.update](resources--api_testing--reference--group-001.md#canonical-2213330211000001-0303331122112302-2112213021320201-1102100220201023-0113022223333131-2320213200202003-3003233013000132-1222122131322330) |

<a id="canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- domains

<a id="canonical-3000110130032222-0313221110233203-1303200210100200-0231012120310000-3213033233233120-1222133223013020-0123311303013200-2023200021113202"></a>

Type: `"object"`. list nested block, Optional.

Add and configure testing domains and credentials.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("credentials",
    "domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
domains {
  # Configure direct properties listed below.
}
```

<a id="canonical-1302122122300113-0230113223320331-3032032203021232-3013331100302103-1122200031033130-0203322331011303-3110002201013211-2030320130332322"></a>

### Direct properties for `domains`

<a id="canonical-1030032120310200-2112320220001031-2011333132301001-0212212021210220-2203221311103001-1322312210000333-2332132032123030-0200130203112113"></a>

#### `domains.allow_destructive_methods` property

Type: `"bool"`. Optional.

Enable to allow API Testing to execute against destructive methods. Use with caution as these may
modify or DELETE data.

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

- [credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301): complete subsection reference.

<a id="canonical-3021301113100320-2332211312021033-3032003132120010-1203013113232323-1123110013103030-3333222030033211-1201020213033300-1020023330130333"></a>

<a id="canonical-2232102300131003-3222101202031233-2121102112330231-3213121331111223-2202131221323113-2210211333221102-0010312003300300-2032013210133311"></a>

#### `domains.domain` property

Type: `"string"`. Optional.

Add your testing environment domain. Be aware that running tests on a production domain can impact
live applications, as API testing cannot distinguish between production and testing environments.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- domains.credentials

<a id="canonical-3311333013222310-0300111023033313-3330331212211301-1001322003233102-1320020330233322-2230022130303123-3020132102323001-3021030020112123"></a>

Type: `"object"`. list nested block, Optional.

Add credentials for API testing to use in the selected environment.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("credential_name"),
  validators.ConflictingListObjectAttributes("admin",
    "standard"),
  validators.ConflictingListObjectAttributes("api_key",
    "basic_auth"),
  validators.ConflictingListObjectAttributes("api_key",
    "bearer_token"),
  validators.ConflictingListObjectAttributes("api_key",
    "login_endpoint"),
  validators.ConflictingListObjectAttributes("basic_auth",
    "bearer_token"),
  validators.ConflictingListObjectAttributes("basic_auth",
    "login_endpoint"),
  validators.ConflictingListObjectAttributes("bearer_token",
    "login_endpoint")}
```

Receipt-pinned upstream constraints:

```json
{
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

Terraform syntax:

```terraform
credentials {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132021133331023-1231220333313033-3201001320113131-3022233321200301-3021223011211012-3220202203121221-2201113103013210-1002200233111013"></a>

### Direct properties for `domains.credentials`

- [admin](resources--api_testing--reference--group-001.md#canonical-3122213202111132-2220200131230222-0133133211302110-1003333230010112-1301022313033321-3313323120100023-0323021032320223-2301233300101212): complete subsection reference.

- [api_key](resources--api_testing--reference--group-001.md#canonical-1310322331313200-2120112223003333-0312311302230121-2213021111301031-2322202103201013-1111001112101213-0113333213002202-2021030011001022): complete subsection reference.

- [basic_auth](resources--api_testing--reference--group-001.md#canonical-2033211223323113-2331012220022033-1231231000223221-2233221003202021-1033332122310111-3220032320212333-1202210233011231-1200331122120232): complete subsection reference.

- [bearer_token](resources--api_testing--reference--group-001.md#canonical-2102000003130030-2110103012123130-2321313213030133-0300232112121133-0122113201230000-1100213033111103-1323112220033201-2121323232122323): complete subsection reference.

<a id="canonical-1102223000330210-2100312111302310-0210333102013033-1031220113300031-1131221101212100-2323202323221231-0011331130330322-3111323132102032"></a>

<a id="canonical-0020121333033013-0201230333302132-0321122123013331-0132201120120030-0333023001120122-2313302223020113-1230201301230213-2302022021323301"></a>

#### `domains.credentials.credential_name` property

Type: `"string"`. Optional.

Enter a unique name for the credentials used in API testing.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [login_endpoint](resources--api_testing--reference--group-001.md#canonical-1312310220102302-0103103210322222-2130003122012103-0203103010212333-3203303011133113-2023122313232220-2020331001121111-1332103212030311): complete subsection reference.

- [standard](resources--api_testing--reference--group-001.md#canonical-0321022203131213-3031011300013101-3002210120000322-1112003033022101-2020232002313102-2321202203310213-1230012333020222-2233200021203210): complete subsection reference.

<a id="canonical-3122213202111132-2220200131230222-0133133211302110-1003333230010112-1301022313033321-3313323120100023-0323021032320223-2301233300101212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.admin` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301)
- domains.credentials.admin

<a id="canonical-1220312000223131-2301201313301310-3332023302312120-2020012101313223-2120221122323132-1011121033023302-2013001000323131-3113313122102003"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
admin = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310322331313200-2120112223003333-0312311302230121-2213021111301031-2322202103201013-1111001112101213-0113333213002202-2021030011001022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.api_key` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301)
- domains.credentials.api_key

<a id="canonical-2111233032210310-1220000103021203-3131111131102010-1102221032232101-0100011000120113-2233202311033220-2321302232212011-0011131102200111"></a>

Type: `"object"`. single nested block, Optional.

API Key

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("key")}
```

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

Terraform syntax:

```terraform
api_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-1311310120023130-2312011312020101-0121303010322013-0310131111300122-0212213010131102-0032110213231302-3232000312212233-2113301231000211"></a>

### Direct properties for `domains.credentials.api_key`

<a id="canonical-2310213331321232-3202211101130311-1113010112131032-2020332031221133-1322112332132302-1203131332002323-1202011232103301-0312330312103212"></a>

#### `domains.credentials.api_key.key` property

Type: `"string"`. Optional.

Key. Cryptographic key material

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [value](resources--api_testing--reference--group-001.md#canonical-3112312300210202-3311001121123231-0122332323322133-3231220321010003-0113121030320031-2220103203012320-3030121332213200-3312201110021123): complete subsection reference.

<a id="canonical-3112312300210202-3311001121123231-0122332323322133-3231220321010003-0113121030320031-2220103203012320-3030121332213200-3312201110021123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.api_key.value` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301)
- [domains.credentials.api_key](resources--api_testing--reference--group-001.md#canonical-1310322331313200-2120112223003333-0312311302230121-2213021111301031-2322202103201013-1111001112101213-0113333213002202-2021030011001022)
- domains.credentials.api_key.value

<a id="canonical-3132201000122012-0301303132300010-0110302201223210-1123103030232023-0220013301023102-1222223030303000-0131012203231230-2223233131211112"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
value {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303112303331331-0323231202011233-0013132322032211-3022121312013310-2100120203001210-0000123020000021-2023010320100222-1220020011011100"></a>

### Direct properties for `domains.credentials.api_key.value`

- [blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-3331223013111213-0020133123123113-3132023101001010-3131110011120123-1010212301302003-0133300202030211-3312223111200111-1333321131233102): complete subsection reference.

- [clear_secret_info](resources--api_testing--reference--group-001.md#canonical-3001133012120300-1300303320332233-3102000022233033-2330210212201023-3032213011331222-0123332203013322-2000000000132111-2003233300001211): complete subsection reference.

<a id="canonical-3331223013111213-0020133123123113-3132023101001010-3131110011120123-1010212301302003-0133300202030211-3312223111200111-1333321131233102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.api_key.value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301)
- [domains.credentials.api_key](resources--api_testing--reference--group-001.md#canonical-1310322331313200-2120112223003333-0312311302230121-2213021111301031-2322202103201013-1111001112101213-0113333213002202-2021030011001022)
- [domains.credentials.api_key.value](resources--api_testing--reference--group-001.md#canonical-3112312300210202-3311001121123231-0122332323322133-3231220321010003-0113121030320031-2220103203012320-3030121332213200-3312201110021123)
- domains.credentials.api_key.value.blindfold_secret_info

<a id="canonical-0231333101312333-0222002002103312-3123322001111030-3010212202200221-2302000132122232-1110330110011230-0112223312102322-0302103302003213"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232333130021003-2031121223130333-3321310002210213-2202233231212301-2322233323002101-0322201311222003-2331002132312202-0113333221223312"></a>

### Direct properties for `domains.credentials.api_key.value.blindfold_secret_info`

<a id="canonical-0312021312020003-1031320122003202-3313030132311200-1113221103310120-2321013330033210-3110110121211220-0001322010301000-3223332300230010"></a>

#### `domains.credentials.api_key.value.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2013323221100000-2220023101131132-2023212032110011-3213302213012311-0232230001212022-3320311003030133-1020302200211023-0330030132310002"></a>

<a id="canonical-2130103332130113-0133222212022012-1332221012220203-1323320002230232-3203100002222202-2020311020122303-2101023200323000-0122123120223122"></a>

#### `domains.credentials.api_key.value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0332301122333133-0010210130100021-2320233103000210-1322231031211222-3231111201211211-0003022311130100-0220210321203101-2022030231100232"></a>

<a id="canonical-1033332100112322-1033022221131313-0301210122023212-2332333100211302-0223211103330023-0320220110031330-2002001222103202-2001200132023213"></a>

#### `domains.credentials.api_key.value.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3001133012120300-1300303320332233-3102000022233033-2330210212201023-3032213011331222-0123332203013322-2000000000132111-2003233300001211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.api_key.value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301)
- [domains.credentials.api_key](resources--api_testing--reference--group-001.md#canonical-1310322331313200-2120112223003333-0312311302230121-2213021111301031-2322202103201013-1111001112101213-0113333213002202-2021030011001022)
- [domains.credentials.api_key.value](resources--api_testing--reference--group-001.md#canonical-3112312300210202-3311001121123231-0122332323322133-3231220321010003-0113121030320031-2220103203012320-3030121332213200-3312201110021123)
- domains.credentials.api_key.value.clear_secret_info

<a id="canonical-2323010030302211-0110300112312201-0030231122002230-3211022110322100-1321322201332220-0212311201320220-0201222111031022-3103233231213321"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222012013131121-1201213233213232-2201203030312221-1023033233123203-0120030321133112-2202022213323310-0113201221301211-0233000202013202"></a>

### Direct properties for `domains.credentials.api_key.value.clear_secret_info`

<a id="canonical-2233003112112221-2310111230021331-3031110031010013-0231322300233221-1313232221222223-0211223222310011-1113032311031021-0003022010321303"></a>

#### `domains.credentials.api_key.value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1201120230130310-2132103333330110-1313323330023321-3221101123322123-3121200001313111-2102230011113332-0123110313301113-2313032222002310"></a>

<a id="canonical-1301203333212330-1011211231203111-1233222311202002-2333121101102013-2002021010212000-3030231010300223-2320021112232313-2310010221102111"></a>

#### `domains.credentials.api_key.value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2033211223323113-2331012220022033-1231231000223221-2233221003202021-1033332122310111-3220032320212333-1202210233011231-1200331122120232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.basic_auth` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301)
- domains.credentials.basic_auth

<a id="canonical-3122303112202130-0102311332101203-2232200131330032-1003002231001330-0313103112112033-1111100022320132-1013003203331021-3223300210021232"></a>

Type: `"object"`. single nested block, Optional.

Basic Authentication.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("user")}
```

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

Terraform syntax:

```terraform
basic_auth {
  # Configure direct properties listed below.
}
```

<a id="canonical-1012203310302302-3222001133321331-0230122133030111-2333120233332322-0301233221103201-3113320230001102-0033011233221100-0021312032231131"></a>

### Direct properties for `domains.credentials.basic_auth`

- [password](resources--api_testing--reference--group-001.md#canonical-0331001113312230-0103212312121003-3301131202232210-0030032301331230-2000023223320332-2223310231011101-0321013212233130-0331111000302321): complete subsection reference.

<a id="canonical-2011331313211001-0313101210220211-3212120330230031-2231102023100303-0102323103103232-2310122003221213-1122113331023321-3302011201330320"></a>

<a id="canonical-0322223212223310-0111223302211233-0111333232121122-0131303233202111-2030230012013022-2200012003103000-1022122210201112-0021211023113012"></a>

#### `domains.credentials.basic_auth.user` property

Type: `"string"`. Optional.

User. Configuration parameter for user

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-0331001113312230-0103212312121003-3301131202232210-0030032301331230-2000023223320332-2223310231011101-0321013212233130-0331111000302321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.basic_auth.password` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301)
- [domains.credentials.basic_auth](resources--api_testing--reference--group-001.md#canonical-2033211223323113-2331012220022033-1231231000223221-2233221003202021-1033332122310111-3220032320212333-1202210233011231-1200331122120232)
- domains.credentials.basic_auth.password

<a id="canonical-2311300030110223-0233211303231331-0030223302132022-0120212030333110-2200100130021122-0202200033021200-0000320021213020-1111200203313022"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330312110003102-2232012020131321-1231113102031220-3222102123202230-0223312233121022-2332203133322300-2311321202322313-0003030331221332"></a>

### Direct properties for `domains.credentials.basic_auth.password`

- [blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-1221122211110303-2002333300112000-2020122103130301-2332003200330003-0112322010300121-3221100211322001-3112023100212002-0210332110321233): complete subsection reference.

- [clear_secret_info](resources--api_testing--reference--group-001.md#canonical-3212220103120302-0030133000303200-2331131233300221-0322031110200102-3031123333300231-2121300211112002-3123300200212311-2112333311000220): complete subsection reference.

<a id="canonical-1221122211110303-2002333300112000-2020122103130301-2332003200330003-0112322010300121-3221100211322001-3112023100212002-0210332110321233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.basic_auth.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301)
- [domains.credentials.basic_auth](resources--api_testing--reference--group-001.md#canonical-2033211223323113-2331012220022033-1231231000223221-2233221003202021-1033332122310111-3220032320212333-1202210233011231-1200331122120232)
- [domains.credentials.basic_auth.password](resources--api_testing--reference--group-001.md#canonical-0331001113312230-0103212312121003-3301131202232210-0030032301331230-2000023223320332-2223310231011101-0321013212233130-0331111000302321)
- domains.credentials.basic_auth.password.blindfold_secret_info

<a id="canonical-3010021321300131-2133222121303301-3310133111220102-2101023033111223-3211011112012313-1003023312120121-0102230230221030-1002133203111001"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1201230020120100-0031222113212230-2221011321322100-3110310020233010-1113121322123320-0131122133210111-0103220001030030-1230230012110123"></a>

### Direct properties for `domains.credentials.basic_auth.password.blindfold_secret_info`

<a id="canonical-1310022111302313-0200103123011332-2220003212221323-1323321323031303-0011112310113303-3023030231111003-3012022133212130-2111321131110321"></a>

#### `domains.credentials.basic_auth.password.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1110303202121322-0030233333133302-2011312230021332-2130013220031323-0320020102021303-0111221332110100-2121030110231201-2110020003022132"></a>

<a id="canonical-1111303210003302-1310333201213203-1021313000010211-1220011323003312-2332113122332013-1230131200333233-1100021311232032-2122222003121221"></a>

#### `domains.credentials.basic_auth.password.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2201102100223202-0132331002121320-3113222003030300-3122030031332231-3131012332112211-0210113303213310-2321213110010310-3212003123013013"></a>

<a id="canonical-1221010110223013-1100110123123120-1232222222000322-1210020012220202-2122331022130312-2313022202011112-1200021103301210-1213103322123202"></a>

#### `domains.credentials.basic_auth.password.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3212220103120302-0030133000303200-2331131233300221-0322031110200102-3031123333300231-2121300211112002-3123300200212311-2112333311000220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.basic_auth.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301)
- [domains.credentials.basic_auth](resources--api_testing--reference--group-001.md#canonical-2033211223323113-2331012220022033-1231231000223221-2233221003202021-1033332122310111-3220032320212333-1202210233011231-1200331122120232)
- [domains.credentials.basic_auth.password](resources--api_testing--reference--group-001.md#canonical-0331001113312230-0103212312121003-3301131202232210-0030032301331230-2000023223320332-2223310231011101-0321013212233130-0331111000302321)
- domains.credentials.basic_auth.password.clear_secret_info

<a id="canonical-1000233013300302-1211012333001001-1202223303202122-3133201312303232-3312100220200031-0302323000233010-0231222211101212-1031312333221120"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1111022101122233-0210213112210313-3003213201320210-3322103111030221-1030201322210223-0302120102030232-1323333313302211-2013111313211021"></a>

### Direct properties for `domains.credentials.basic_auth.password.clear_secret_info`

<a id="canonical-3231021021311021-3220220131012212-1033031320122320-0312122323312200-1112220213233332-3213320130310021-3031302311212010-3020303321121303"></a>

#### `domains.credentials.basic_auth.password.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1011320132000121-3130021032032320-2012131130111110-0221001202310002-2220011032300001-3010023111212222-3131101133100013-3331330232311333"></a>

<a id="canonical-2120020012200002-2203013333322333-0300010100222033-1012301232302120-3211132200223010-0110002121020123-2201032123101302-3302100211113331"></a>

#### `domains.credentials.basic_auth.password.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2102000003130030-2110103012123130-2321313213030133-0300232112121133-0122113201230000-1100213033111103-1323112220033201-2121323232122323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.bearer_token` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301)
- domains.credentials.bearer_token

<a id="canonical-2200312201130130-3022212231302121-2110033020202131-1000003303122330-3100021222103021-3230021221221001-2301010321302023-2113101330101101"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bearer token.

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

Terraform syntax:

```terraform
bearer_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311030300000303-2132300110003032-0323130012022330-3001303233301033-3022011321011212-0211210202110331-2333331011222033-1313300133210122"></a>

### Direct properties for `domains.credentials.bearer_token`

- [token](resources--api_testing--reference--group-001.md#canonical-0221310330223200-0113211011210010-2312122011223200-1211230323232203-3032330002023120-0332002120320212-3323002210333322-1100010202103100): complete subsection reference.

<a id="canonical-0221310330223200-0113211011210010-2312122011223200-1211230323232203-3032330002023120-0332002120320212-3323002210333322-1100010202103100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.bearer_token.token` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301)
- [domains.credentials.bearer_token](resources--api_testing--reference--group-001.md#canonical-2102000003130030-2110103012123130-2321313213030133-0300232112121133-0122113201230000-1100213033111103-1323112220033201-2121323232122323)
- domains.credentials.bearer_token.token

<a id="canonical-3132333000302201-2302230231232102-3032202001022232-3100330232033212-2313022123002212-3220200331320231-3033311113121313-0033123201322131"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
token {
  # Configure direct properties listed below.
}
```

<a id="canonical-2030332202123232-1010230313222111-2323310102221130-1211031320311330-0212123120310202-0002110030332002-3332112030312103-1232310033213323"></a>

### Direct properties for `domains.credentials.bearer_token.token`

- [blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-1032130023210230-3320320002021221-2031122211110123-3021021333101011-1303013111230213-0230132013113210-0330023312220333-1013312103003223): complete subsection reference.

- [clear_secret_info](resources--api_testing--reference--group-001.md#canonical-3130120303123330-1023021233102223-0120333013210120-2122200120031100-1221203330302230-2333113003331000-3023221203130103-3200322201132003): complete subsection reference.

<a id="canonical-1032130023210230-3320320002021221-2031122211110123-3021021333101011-1303013111230213-0230132013113210-0330023312220333-1013312103003223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.bearer_token.token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301)
- [domains.credentials.bearer_token](resources--api_testing--reference--group-001.md#canonical-2102000003130030-2110103012123130-2321313213030133-0300232112121133-0122113201230000-1100213033111103-1323112220033201-2121323232122323)
- [domains.credentials.bearer_token.token](resources--api_testing--reference--group-001.md#canonical-0221310330223200-0113211011210010-2312122011223200-1211230323232203-3032330002023120-0332002120320212-3323002210333322-1100010202103100)
- domains.credentials.bearer_token.token.blindfold_secret_info

<a id="canonical-2233002030222021-1222133332301232-3212113303120010-0103220102133032-0221111122121132-0302322203012212-2123030323232113-1010023313101103"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3120223230122013-1232000113012311-2312022223201222-0130213103010313-3012021031010111-1223101013002010-3010033111103023-1012332213013210"></a>

### Direct properties for `domains.credentials.bearer_token.token.blindfold_secret_info`

<a id="canonical-1112031022022120-2332030002122022-0111312002132232-1013002312112303-3120230022133210-3312202010011312-2202121210120103-1000113311032133"></a>

#### `domains.credentials.bearer_token.token.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1313031130331111-1322221332201321-1101223323112032-3231032212122130-2202230110020010-2333220100233301-1113110202300031-2213103001013113"></a>

<a id="canonical-1032231223122300-1120231112031321-1320010133203021-0101003122032212-3132113111213021-1132302203113012-2222220021123121-2301233102332032"></a>

#### `domains.credentials.bearer_token.token.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2321323320003011-2032223000210220-2132213331111010-0013133100200322-0131133302310112-1112121200120032-3133001130112322-1223102030333311"></a>

<a id="canonical-1311020113310112-2111211231121300-3020120033103312-0001121033212030-3110000110121113-1121232013111101-1102133101032302-1330201300021130"></a>

#### `domains.credentials.bearer_token.token.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3130120303123330-1023021233102223-0120333013210120-2122200120031100-1221203330302230-2333113003331000-3023221203130103-3200322201132003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.bearer_token.token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301)
- [domains.credentials.bearer_token](resources--api_testing--reference--group-001.md#canonical-2102000003130030-2110103012123130-2321313213030133-0300232112121133-0122113201230000-1100213033111103-1323112220033201-2121323232122323)
- [domains.credentials.bearer_token.token](resources--api_testing--reference--group-001.md#canonical-0221310330223200-0113211011210010-2312122011223200-1211230323232203-3032330002023120-0332002120320212-3323002210333322-1100010202103100)
- domains.credentials.bearer_token.token.clear_secret_info

<a id="canonical-3120300131112200-3001123312302020-3011101001333232-3223033020031230-0030323312102211-1300101001231233-3013122101012221-1010210210330031"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330330301112131-1331330011331223-0313120111230233-2211100033200111-2221112333221011-1121132330210210-1233223023112320-0023121023032330"></a>

### Direct properties for `domains.credentials.bearer_token.token.clear_secret_info`

<a id="canonical-3303033012002130-3013221100101133-0031032031320133-2201022203213221-2133002023211001-3212003012000011-3221311202121332-2013210332033202"></a>

#### `domains.credentials.bearer_token.token.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3221212313202310-2103212102102101-1100011133300212-2110233013100331-3311002030131112-3203312002133010-3132123132021322-2302221131230332"></a>

<a id="canonical-0103332200000320-0333230312213111-0321131002200123-3200110101310221-2133110031102311-2210320130131012-1222103010003120-2133023132122111"></a>

#### `domains.credentials.bearer_token.token.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1312310220102302-0103103210322222-2130003122012103-0203103010212333-3203303011133113-2023122313232220-2020331001121111-1332103212030311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.login_endpoint` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301)
- domains.credentials.login_endpoint

<a id="canonical-0023021021120100-2110130121213311-1102301201200011-2012101220002312-3230303331030323-0003302130001231-0022020202021101-1231101000212223"></a>

Type: `"object"`. single nested block, Optional.

Login Endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("path",
    "token_response_key")}
```

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

Terraform syntax:

```terraform
login_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011110323203023-2311312010231223-0112033010002012-3012232313100100-2001111103121032-0311203230033023-3001111232333023-1300330113311201"></a>

### Direct properties for `domains.credentials.login_endpoint`

- [json_payload](resources--api_testing--reference--group-001.md#canonical-2022021213220201-2233331101332133-1312233301111020-3102231111200321-2303231331113321-0120220311133230-2201100332012212-2011001001012123): complete subsection reference.

<a id="canonical-2322000012311221-0333310302123313-2300033020212301-0331232120103221-3230310310002012-1013122011312223-1003200210003232-3020021312011033"></a>

<a id="canonical-3311212200113011-1220223103113223-1011020023121312-2013212001131312-2133322110231123-2312120120022203-2230011221321310-1120202202031212"></a>

#### `domains.credentials.login_endpoint.method` property

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ANY","CONNECT","COPY","DELETE","GET","HEAD","OPTIONS","PATCH","POST","PUT","TRACE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3122102233310321-0002311133303112-1313031200330330-1023113210002023-1012131220032121-2123032111223000-2221001200303213-3311221232210330"></a>

<a id="canonical-1301031130301100-3211131010302201-1323110333022120-2002121200233021-1112112033002322-0322131232300320-0220003012112010-2033233332112303"></a>

#### `domains.credentials.login_endpoint.path` property

Type: `"string"`. Optional.

Path. URL path for the endpoint

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "1024"
  }
}
```

<a id="canonical-3021202021010310-3013212013213103-1003330030303301-3001222023023223-1101210113111210-2103112102223103-1303122322301012-3131201310323101"></a>

<a id="canonical-1201103321122231-3010131202020222-3212311001222023-2311030313232011-0133222320321322-1223230200330210-1022303223213130-3221100113210121"></a>

#### `domains.credentials.login_endpoint.token_response_key` property

Type: `"string"`. Optional.

Configuration parameter for token response key.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2022021213220201-2233331101332133-1312233301111020-3102231111200321-2303231331113321-0120220311133230-2201100332012212-2011001001012123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.login_endpoint.json_payload` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301)
- [domains.credentials.login_endpoint](resources--api_testing--reference--group-001.md#canonical-1312310220102302-0103103210322222-2130003122012103-0203103010212333-3203303011133113-2023122313232220-2020331001121111-1332103212030311)
- domains.credentials.login_endpoint.json_payload

<a id="canonical-3322221202121210-1230011331220102-3130101311201330-3302112001231110-0212021111320230-0013233230032201-3132222021013230-1100113320213113"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
json_payload {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022002233021023-0133102321123111-1200200321130003-1021210330111032-1221310332102303-3200120021201333-2000201110111321-0023313123322320"></a>

### Direct properties for `domains.credentials.login_endpoint.json_payload`

- [blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-1031130201203203-2130331023231213-1231230001100111-3120230121211220-2213103223200011-0030111213202301-3200020223232211-0130121030313212): complete subsection reference.

- [clear_secret_info](resources--api_testing--reference--group-001.md#canonical-1100021321303321-0213032221113301-2231021322323100-0032313001032123-1121133030321033-3321033323031220-3131221021132121-3031211331332103): complete subsection reference.

<a id="canonical-1031130201203203-2130331023231213-1231230001100111-3120230121211220-2213103223200011-0030111213202301-3200020223232211-0130121030313212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.login_endpoint.json_payload.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301)
- [domains.credentials.login_endpoint](resources--api_testing--reference--group-001.md#canonical-1312310220102302-0103103210322222-2130003122012103-0203103010212333-3203303011133113-2023122313232220-2020331001121111-1332103212030311)
- [domains.credentials.login_endpoint.json_payload](resources--api_testing--reference--group-001.md#canonical-2022021213220201-2233331101332133-1312233301111020-3102231111200321-2303231331113321-0120220311133230-2201100332012212-2011001001012123)
- domains.credentials.login_endpoint.json_payload.blindfold_secret_info

<a id="canonical-0022301033221322-1331001231330002-0201000031032332-1100030332333220-1000020321213020-3013331033300332-3000030030223131-1123002300102312"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220210033302230-0321300322323113-3210211103230030-1210220001132032-3122121003010033-3003123002312111-3210132333332010-3233313200322000"></a>

### Direct properties for `domains.credentials.login_endpoint.json_payload.blindfold_secret_info`

<a id="canonical-3111333000102333-3100133322331133-3333230233313200-1001030333021003-1103323010111333-1033210230002301-1002120321222210-2103210223002123"></a>

#### `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3002021113312213-2322322121221231-2002313000211002-1020303330120310-0122010133120202-0020120223233221-3111220021211131-2011031022212130"></a>

<a id="canonical-1312130220231010-1223032020301323-2231131303113020-1122112131013013-3113133021100133-3210322312300110-3221303103303033-1100331323320320"></a>

#### `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3110323133333322-3200213313033010-2312102311033011-0013002202100211-3122102120002032-3221223332020103-2221231130222011-2203131100201032"></a>

<a id="canonical-3102002312232121-3122232231111021-2222222312231212-3031002210020110-1321102131223221-1203222123013210-0130300303213321-3300021102130201"></a>

#### `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1100021321303321-0213032221113301-2231021322323100-0032313001032123-1121133030321033-3321033323031220-3131221021132121-3031211331332103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.login_endpoint.json_payload.clear_secret_info` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301)
- [domains.credentials.login_endpoint](resources--api_testing--reference--group-001.md#canonical-1312310220102302-0103103210322222-2130003122012103-0203103010212333-3203303011133113-2023122313232220-2020331001121111-1332103212030311)
- [domains.credentials.login_endpoint.json_payload](resources--api_testing--reference--group-001.md#canonical-2022021213220201-2233331101332133-1312233301111020-3102231111200321-2303231331113321-0120220311133230-2201100332012212-2011001001012123)
- domains.credentials.login_endpoint.json_payload.clear_secret_info

<a id="canonical-3012000301330231-2003220130210212-2112301110122323-0102032033023112-3133301223112003-2323302313030211-3002333111222210-1001311101320220"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202031101130022-3230002320120123-2301131112000321-2223030211212212-1312231321303101-3313233123130221-3030002213311130-3213121200110200"></a>

### Direct properties for `domains.credentials.login_endpoint.json_payload.clear_secret_info`

<a id="canonical-2132310033330301-0332033120103232-1011200002233312-1320010210211233-2223100112030123-3233321020233320-0230131201001210-0222312031302232"></a>

#### `domains.credentials.login_endpoint.json_payload.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3010212111313123-1000131012112013-0220202000312111-0220132301203111-1303201301311233-0322210120003100-3130133323321310-1213130202222302"></a>

<a id="canonical-3033001131010132-2001221010230010-0012202333001302-1320211312211020-2310022101122013-3001233113310331-0223122133331102-3311112222233033"></a>

#### `domains.credentials.login_endpoint.json_payload.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0321022203131213-3031011300013101-3002210120000322-1112003033022101-2020232002313102-2321202203310213-1230012333020222-2233200021203210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.standard` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [domains](resources--api_testing--reference--group-001.md#canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301)
- domains.credentials.standard

<a id="canonical-1220302023112021-3021330122302211-2033231201232200-3223111320101321-1330330211230320-3332213002211321-3020330032011100-2010200002202223"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
standard = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110133330203132-0300210311112100-1220023032121103-3000313102030220-3123012332002301-1030201222002031-3231121132130330-1120323230220100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `every_day` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- every_day

<a id="canonical-1121033022331320-3113200120231112-0200210231013322-1001130230313322-0101131202323100-1010102013333312-3200001121112001-3321123230220023"></a>

Type: `["object", {}]`. Optional.

\[OneOf: every\_day, every\_month, every\_week\] Enable this option

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

OneOf alternatives in this subsection:

- [every_day](resources--api_testing--reference--group-001.md#canonical-1121033022331320-3113200120231112-0200210231013322-1001130230313322-0101131202323100-1010102013333312-3200001121112001-3321123230220023)
- [every_month](resources--api_testing--reference--group-001.md#canonical-1002110121321233-0110021120131230-3122220200001211-1010231200012330-0132230332113111-3333011132102112-1211301332122332-1310221223321122)
- [every_week](resources--api_testing--reference--group-001.md#canonical-0331001330211323-3330113021323212-1212213000000300-0322101000233222-1110122200222011-1202123110101113-1310001322111122-1231101130301111)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
every_day = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323120202200200-3120122313023333-2112000223300001-2210110033112032-1033031302013222-0202010110312210-3111003132122233-2102132121222333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `every_month` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- every_month

<a id="canonical-1002110121321233-0110021120131230-3122220200001211-1010231200012330-0132230332113111-3333011132102112-1211301332122332-1310221223321122"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for every month.

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

Terraform syntax:

```terraform
every_month = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003302130133212-2030223321333013-0223331221002310-0112202233231032-2022332023223330-0202032113013331-2112123012003322-0233330122012011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `every_week` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- every_week

<a id="canonical-0331001330211323-3330113021323212-1212213000000300-0322101000233222-1110122200222011-1202123110101113-1310001322111122-1231101130301111"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
every_week = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220221222303221-1321122301102101-3213032112213022-2113023312203202-0313020120112103-0030330120312111-1212213100132010-0212222120301330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- timeouts

<a id="canonical-2311001321033301-3302013302233100-2122001203300033-0231332233200013-2102102121213121-1323033121121100-1231033211033313-1021112011320303"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3123010212210201-0020233300113301-2223110122010112-3111222022302310-1023313132021112-0233203013010221-0302031222330113-2332032021200131"></a>

### Direct properties for `timeouts`

<a id="canonical-2131132300301123-1031322011120000-2333023122010130-2122122031013313-3333011310303123-1003000323200311-1113321002031120-2102303322003021"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3301303330220310-3300322230233022-3103123111210223-2312110221011131-1013230220331200-1013321313120002-2312121102302303-1111212011023322"></a>

<a id="canonical-3020130332130002-2303301300202132-2202120310123210-1323012230301122-1312021302032100-3033230002330323-0002302000000200-0030123102301330"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2032331201301301-1200233000130231-0010130320122010-3210200133121001-3332233021331002-2100212000011300-1211212221312021-3210130231113010"></a>

<a id="canonical-3222132223222132-2123012111202001-3220301120033020-2001120130000031-1311131031101101-3313121322203300-1003301301100100-0330132021020203"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2213330211000001-0303331122112302-2112213021320201-1102100220201023-0113022223333131-2320213200202003-3003233013000132-1222122131322330"></a>

<a id="canonical-0010221131013311-2111102123201123-2330101213111330-3331122212322231-0002213113200301-3002312112023001-1031211202113323-2032323011002021"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
