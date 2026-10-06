---
page_title: "xcsh_protected_application reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application reference."
---

# xcsh_protected_application reference

<a id="canonical-1323310101232211-3022030330331220-3221213030103111-1210123300102021-3000231032122231-1011332222123312-0303323200120223-1203201223210122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules.rules` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200)
- CloudFront.js_insertion_rules.rules

<a id="canonical-3112211323200012-1223322323223121-1221102013102200-2231133101331021-2300001032233333-0231120330103210-1212210332300220-1012311313022321"></a>

Type: `"list"`. Computed.

Required list of pages to insert Bot Defense client JavaScript.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0323130012001011-3330023213311223-0131100031230021-0233323310130301-0000320321210211-2212312310031202-3300012010000132-0121301323232302"></a>

### Direct properties for `cloudfront.js_insertion_rules.rules`

- [any_domain](data-sources--protected_application--reference--group-003.md#canonical-3302002122001011-3302130021132031-0001333011033320-0300322020323210-2313001110120312-2113031232021210-1112023131331222-3112330121301233): complete subsection reference.

- [domain](data-sources--protected_application--reference--group-003.md#canonical-1210011000220311-1313103323102110-0102230012221021-0202221001203120-0030001013011111-3312221132223000-1310210311202010-1011131003202032): complete subsection reference.

<a id="canonical-2330022230132211-3221002200113200-1100131123202310-0210100330031012-0221112021331332-2011101230213322-2121232033121311-0332111332222022"></a>

<a id="canonical-0031300302002113-0222303132212322-2110010031200310-0123101103301230-2121233201013311-0000201221012103-0001310231203023-0111210033011121"></a>

#### `cloudfront.js_insertion_rules.rules.exact_path` property

Type: `"string"`. Computed.

Exclusive with \[glob prefix\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0213003102103301-0231133303321310-0010223000220100-0031033033231303-0310101301133333-0111300200213321-3331312121120012-3231012303221121"></a>

<a id="canonical-3222303212231322-1300221330222302-3312031222013312-3202212031022132-2032331303331302-3013003101021220-3212303032210233-3021200022010011"></a>

#### `cloudfront.js_insertion_rules.rules.glob` property

Type: `"string"`. Computed.

Exclusive with \[exact\_path prefix\] Accepts wildcards \* to match multiple characters or ? To
match a single character.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,256}$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,256}$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,256}$"
  }
}
```

- [metadata](data-sources--protected_application--reference--group-003.md#canonical-3113000122001003-0110000112301311-0321320301102330-1123221321102312-3112011333030130-2220002310032301-3012230130300020-2212230302131021): complete subsection reference.

<a id="canonical-1102112322211030-1110013213130020-1203320202302331-0111311100132202-2011023302003030-2110102210033222-3022310112030302-0120230202210230"></a>

<a id="canonical-0321202301023223-2110030222000333-1103222012303202-2001130121002200-1102210020132323-2313301132002130-2120301102303223-2231221013330222"></a>

#### `cloudfront.js_insertion_rules.rules.prefix` property

Type: `"string"`. Computed.

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3302002122001011-3302130021132031-0001333011033320-0300322020323210-2313001110120312-2113031232021210-1112023131331222-3112330121301233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules.rules.any_domain` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200)
- [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-003.md#canonical-1323310101232211-3022030330331220-3221213030103111-1210123300102021-3000231032122231-1011332222123312-0303323200120223-1203201223210122)
- CloudFront.js_insertion_rules.rules.any_domain

<a id="canonical-0222003222210323-3202111102203220-1102211320120213-0333111312123211-0013112101330113-2223033210103231-3112222122223112-2010012010101013"></a>

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

<a id="canonical-1210011000220311-1313103323102110-0102230012221021-0202221001203120-0030001013011111-3312221132223000-1310210311202010-1011131003202032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules.rules.domain` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200)
- [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-003.md#canonical-1323310101232211-3022030330331220-3221213030103111-1210123300102021-3000231032122231-1011332222123312-0303323200120223-1203201223210122)
- CloudFront.js_insertion_rules.rules.domain

<a id="canonical-3331221212300133-1223211303033303-2010130322021122-0102122012131310-2011202003022123-0013311223221032-0102102233320230-2132212220311310"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-3321112102323001-2023121011221302-3020010102113221-1331022200313322-0012022103321131-2013330111213213-3203003032323111-2312033013311302"></a>

### Direct properties for `cloudfront.js_insertion_rules.rules.domain`

<a id="canonical-2123230330211300-0002023101232120-2012122123131320-0213333223012112-2331100330233130-0213213130033010-0122310001122012-0232300002112131"></a>

#### `cloudfront.js_insertion_rules.rules.domain.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1311132312311103-1133212022210313-2102113332101133-1122331100203012-3131133212210323-0212122022220222-2303200310303020-3221003033000202"></a>

<a id="canonical-3132100311102230-1033221303032220-3233103020202032-1310320301311303-3211100020231331-1031103313331002-2021022122301321-1232123210121001"></a>

#### `cloudfront.js_insertion_rules.rules.domain.regex_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3213310122031123-3320313130031030-1133132220220232-1023201321003121-0311032300010220-2232031033203332-1212300320100230-2123332311302032"></a>

<a id="canonical-2321032123221232-3131022202033203-0202310000122012-3111323011031333-2203002211223313-1232323303213023-2112312203030100-3103330120321013"></a>

#### `cloudfront.js_insertion_rules.rules.domain.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3113000122001003-0110000112301311-0321320301102330-1123221321102312-3112011333030130-2220002310032301-3012230130300020-2212230302131021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules.rules.metadata` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200)
- [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-003.md#canonical-1323310101232211-3022030330331220-3221213030103111-1210123300102021-3000231032122231-1011332222123312-0303323200120223-1203201223210122)
- CloudFront.js_insertion_rules.rules.metadata

<a id="canonical-0022310302103031-0331231102030333-3120321222311312-0020131003331111-2212033323301313-1302223032031212-2331130022200020-1322331013311101"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-1101231312130121-3030101132122012-2120330130212313-0203332300232130-3232130311303332-0133200223333221-1323311100323212-3311201210202023"></a>

### Direct properties for `cloudfront.js_insertion_rules.rules.metadata`

<a id="canonical-3001231312311121-1312231001320021-0003313103123212-0330312112110303-1231003321101131-3301230303111333-1231030120122132-2131032220323020"></a>

#### `cloudfront.js_insertion_rules.rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2123133012212133-2310202220311233-0213230023302321-3201300321300013-3210203312211110-3112113210112021-1303030212113023-2030110203023211"></a>

<a id="canonical-1223130322020133-2333011113232201-2013310323312101-2303213100101103-1202213111102320-3133230021122031-1212323302213201-2130333312231003"></a>

#### `cloudfront.js_insertion_rules.rules.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-3123213201022200-1010003020110031-2133001121200133-1103021210123222-3223203013201202-3103020013112213-0232333101021031-3332032331023310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.manual_js_insert` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- CloudFront.manual_js_insert

<a id="canonical-1231011133010223-3122230322022221-3021222003233331-2101311111113001-2313033033311320-2101020223111333-2213123133202022-0212231202212132"></a>

Type: `"single"`. Computed.

Insert JavaScript Manually. Insert JavaScript manually.

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

<a id="canonical-0101300002012313-0033121211300032-2202123331002213-0100300311012321-3032131122310020-0213023220313320-3133313102333131-1223233231203300"></a>

### Direct properties for `cloudfront.manual_js_insert`

<a id="canonical-2031223130012320-1002023220033121-1123303221123200-2103321031032013-3223103101130120-3230211232313203-2030202220003312-2011020133131232"></a>

#### `cloudfront.manual_js_insert.javascript_mode` property

Type: `"string"`. Computed.

\[Enum: ASYNC\_JS\_NO\_CACHING|ASYNC\_JS\_CACHING|SYNC\_JS\_NO\_CACHING|SYNC\_JS\_CACHING\] Web
Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is cacheable Bot Defense JavaScript for telemetry collection is requested.. Possible values
are \`ASYNC\_JS\_NO\_CACHING\`, \`ASYNC\_JS\_CACHING\`, \`SYNC\_JS\_NO\_CACHING\`,
\`SYNC\_JS\_CACHING\`. Defaults to \`ASYNC\_JS\_NO\_CACHING\`.

Additional upstream details:

Web Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested
asynchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested
asynchronously, and it is cacheable Bot Defense JavaScript for telemetry collection is requested
synchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested
synchronously, and it is cacheable.

Receipt-pinned upstream constraints:

```json
{
  "default": "ASYNC_JS_NO_CACHING",
  "enum": [
    "ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0313331030133212-1032111323032303-0100021032110220-3232111112031033-0131011203020002-0330132231012123-3230212022110120-0201303001002023"></a>

<a id="canonical-3030212332221130-0011022121110133-1223122221333002-1210100133312022-2123301101211302-2010303111021133-2331230111102100-2133212331300033"></a>

#### `cloudfront.manual_js_insert.js_download_path` property

Type: `"string"`. Computed.

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths. If not specified, default to ‘/CommonJS’.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

<a id="canonical-2323133023333211-1301331031223130-0323301221013333-3021333133112332-2002133121021322-1321033012103203-2231303102303230-3321111230000113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.mobile_sdk_config` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- CloudFront.mobile_sdk_config

<a id="canonical-1312021210122012-0331300021020123-2230203331133203-1221133112213030-2301102122213221-0222131110123110-3321211320210302-3312002113320132"></a>

Type: `"single"`. Computed.

Mobile SDK Configuration. Mobile SDK configuration.

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

<a id="canonical-2101210222212033-1310131022121100-0300201032122323-3203320310100110-3100033222011121-2002212001000323-3211202132023303-0100310303303102"></a>

### Direct properties for `cloudfront.mobile_sdk_config`

- [mobile_identifier](data-sources--protected_application--reference--group-003.md#canonical-3111022200003202-3213033223231322-3001030302002231-1021330202220111-0022300303131033-3121101331221213-0031123031121320-3020332023002031): complete subsection reference.

<a id="canonical-3111022200003202-3213033223231322-3001030302002231-1021330202220111-0022300303131033-3121101331221213-0031123031121320-3020332023002031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.mobile_sdk_config.mobile_identifier` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.mobile_sdk_config](data-sources--protected_application--reference--group-003.md#canonical-2323133023333211-1301331031223130-0323301221013333-3021333133112332-2002133121021322-1321033012103203-2231303102303230-3321111230000113)
- CloudFront.mobile_sdk_config.mobile_identifier

<a id="canonical-1330032323232100-1320130301232230-3123220222202301-2303203231131111-1230023201202201-3120001300002310-3121012232233113-1220110213011102"></a>

Type: `"single"`. Computed.

Mobile Traffic Identifier. Mobile traffic identifier type.

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

<a id="canonical-3122201120303332-3322333012312323-2120002122121222-1110313331212221-1200313102220103-3122302020320112-2200230000101123-0203000302210330"></a>

### Direct properties for `cloudfront.mobile_sdk_config.mobile_identifier`

- [headers](data-sources--protected_application--reference--group-003.md#canonical-3222012320102033-2322223111110322-2320222210322101-0212011322302131-3020132003323230-1330132321001033-3303333032202113-0331323313020033): complete subsection reference.

<a id="canonical-3222012320102033-2322223111110322-2320222210322101-0212011322302131-3020132003323230-1330132321001033-3303333032202113-0331323313020033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.mobile_sdk_config.mobile_identifier.headers` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.mobile_sdk_config](data-sources--protected_application--reference--group-003.md#canonical-2323133023333211-1301331031223130-0323301221013333-3021333133112332-2002133121021322-1321033012103203-2231303102303230-3321111230000113)
- [cloudfront.mobile_sdk_config.mobile_identifier](data-sources--protected_application--reference--group-003.md#canonical-3111022200003202-3213033223231322-3001030302002231-1021330202220111-0022300303131033-3121101331221213-0031123031121320-3020332023002031)
- CloudFront.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-1322312110312033-2320011333131330-2123130100300003-0333021023212202-2333312330312103-3120323300202121-2023102131330130-2300112230302233"></a>

Type: `"list"`. Computed.

A list of headers that can be used to identify mobile traffic.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2133003031002333-0133310000212000-0121201100211320-3303202221131130-1120322320220200-2121230110120132-2013213133130132-1230221301130222"></a>

### Direct properties for `cloudfront.mobile_sdk_config.mobile_identifier.headers`

<a id="canonical-2322223030021221-0033131321320231-0331223011022122-3102122102020112-2120010302122123-1021200210122323-3032003103331012-0113323213102132"></a>

#### `cloudfront.mobile_sdk_config.mobile_identifier.headers.exact` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\] Header value to match exactly.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-0102231110012330-0103310313211102-2013233022022133-0120310111110032-2330013120001022-3132122312130133-1210211300111203-3121103301110112"></a>

<a id="canonical-0212122013202101-2200213310313001-2233201213320320-2131020031203031-3230210111230100-2302133133323212-3003311100102133-0220022021222122"></a>

#### `cloudfront.mobile_sdk_config.mobile_identifier.headers.name` property

Type: `"string"`. Computed.

Name. Name of the header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0332331133103302-3012133233310233-0332121101132122-1201001030222030-0012031130123333-2030011232113130-3131232201021223-3220310003032102"></a>

<a id="canonical-3011023303033021-0311012210013000-2311103223020312-0232002310202010-0123120122312132-1031003010210200-2012323213003222-0113010123000031"></a>

#### `cloudfront.mobile_sdk_config.mobile_identifier.headers.regex` property

Type: `"string"`. Computed.

Exclusive with \[exact\] regular expression match of the header value in re2 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- CloudFront.protected_endpoints

<a id="canonical-0232213020212323-1020030122103132-2231022031131020-2201113331300303-0111301321330201-2122020311332330-3303032103301223-2001300220202110"></a>

Type: `"list"`. Computed.

List of protected endpoints (max 128 items).

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2322001211320000-0222220121201132-3300030222322103-0301022030331132-0322002320310220-1122130202003332-0101111302203210-2032213301103321"></a>

### Direct properties for `cloudfront.protected_endpoints`

- [any_domain](data-sources--protected_application--reference--group-003.md#canonical-0212131231100333-2110011000010121-2031311012203203-1223132000302322-0333333123301001-2130310113101232-0130301333301102-1232032121111231): complete subsection reference.

- [domain](data-sources--protected_application--reference--group-003.md#canonical-1231021222223112-0213002101313033-1233123023100112-2301210113023100-2132323003033223-3310130002230021-3302312332023010-2300022030133103): complete subsection reference.

- [flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300): complete subsection reference.

<a id="canonical-3303203312203031-0120303221223312-0331030213320330-1330310122213111-0232213031112330-2102000122301330-1321101100210332-3000011322133232"></a>

<a id="canonical-1322111001132313-2003202232112000-0122202110231321-3202123133031220-3020310100023013-2321232212000003-2323322113220030-1322012201211322"></a>

#### `cloudfront.protected_endpoints.http_methods` property

Type: `["list", "string"]`. Computed.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[1,3,4]",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[1,3,4]",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](data-sources--protected_application--reference--group-004.md#canonical-2021230303322000-2030220332003021-3120230122221302-3030222030333130-2213211332121312-1222122333301110-3220301200123102-0201311323123022): complete subsection reference.

- [mobile_client](data-sources--protected_application--reference--group-004.md#canonical-0300100133111022-1320133032023113-0200312122311310-2330312002011033-0123110113021233-1111122331121201-0322031011211223-1323112133031322): complete subsection reference.

<a id="canonical-2122012101213113-2221201200303102-2122011020233202-0221332133220331-0210231203321001-2123103300221131-3123322013331210-0213111111132332"></a>

<a id="canonical-1123212220230333-3032300030333220-0033221033132210-1132022213133000-0031030333102000-3020321031123223-1312233223111301-2223303332311322"></a>

#### `cloudfront.protected_endpoints.path` property

Type: `"string"`. Computed.

Accepts wildcards \* to match multiple characters or ? To match a single character.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  }
}
```

<a id="canonical-3032233123200112-1222002132032020-0331330222131300-1331113200212013-3133301233321120-0030323330310210-1112032002302112-0231003232201101"></a>

<a id="canonical-1132012021001313-0012113120020310-2033121033130210-3212101220333111-1313123202022300-3213111013032112-0021133213112110-0302331232313123"></a>

#### `cloudfront.protected_endpoints.query` property

Type: `"string"`. Computed.

Enter a regular expression to match your query parameters of interest.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

- [undefined_flow_label](data-sources--protected_application--reference--group-004.md#canonical-2130323322302022-0212320201322103-2202001122001133-3110232001330021-2300113223103221-0321313111323130-3330323312330233-3013133120303122): complete subsection reference.

- [web_client](data-sources--protected_application--reference--group-004.md#canonical-0021131221101020-3323020220213013-3221300121212033-3022302311332020-0101211313203023-0133110222010302-0221020001111012-0011131003021220): complete subsection reference.

- [web_mobile_client](data-sources--protected_application--reference--group-004.md#canonical-1221222210233203-0121321330221033-3221101000112032-3233233303213100-2313001001111110-3313022302233212-3003311322313220-1200102012332120): complete subsection reference.

<a id="canonical-0212131231100333-2110011000010121-2031311012203203-1223132000302322-0333333123301001-2130310113101232-0130301333301102-1232032121111231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.any_domain` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- CloudFront.protected_endpoints.any_domain

<a id="canonical-0302122110010130-2031103103101000-0302002022210020-0022230202103223-3313033133303101-1313013020203010-1202213023111123-1323003011003211"></a>

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

<a id="canonical-1231021222223112-0213002101313033-1233123023100112-2301210113023100-2132323003033223-3310130002230021-3302312332023010-2300022030133103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.domain` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- CloudFront.protected_endpoints.domain

<a id="canonical-2302231201023310-2313302100000222-3123132201202211-0313113312232311-3333220322013203-0011200020212100-0210213201332220-0202230123332103"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-2101122133013123-2212130121123032-0222222202311100-0010321330031123-1313002122310332-0032322312112120-1112233031102003-2032231012112233"></a>

### Direct properties for `cloudfront.protected_endpoints.domain`

<a id="canonical-2232133333212102-1110113103100322-3223321213132000-2110230230210100-1113302133010222-1223303130133001-2132011213223011-3011122111003330"></a>

#### `cloudfront.protected_endpoints.domain.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2102020000130020-1230311222231003-2103223202013111-2021311330313121-3033021101322130-0230023011311030-2111113312202332-2000203312323320"></a>

<a id="canonical-3233103113312130-3102012223200113-0223213223203220-1132113311032332-1311111311020013-0222313003320302-0123030120031121-1323313201130212"></a>

#### `cloudfront.protected_endpoints.domain.regex_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3100331203013213-2111233123322013-1323121230320310-3130113203100222-3032203021313231-2030330201023023-3120231113012002-0213033130221103"></a>

<a id="canonical-2203313231112102-2313123223322332-3300303030123232-0201132302123221-0332302211012223-0111012003300223-1010320131022132-2332310101023203"></a>

#### `cloudfront.protected_endpoints.domain.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- CloudFront.protected_endpoints.flow_label

<a id="canonical-2020033321111033-2202222133123020-2131133212213323-0222100102031133-2033032012121202-2313202322130212-3310021331311022-3013222210000203"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-flow_label_choice": "[\"account_management\",\"authentication\",\"financial_services\",\"flight\",\"profile_management\",\"search\",\"shopping_gift_cards\"]"
}
```

<a id="canonical-0301233230313331-3230013102322213-0210011231220212-0220013203000230-3032230302201332-2000121221030220-3022023302130013-3303010221020013"></a>

### Direct properties for `cloudfront.protected_endpoints.flow_label`

- [account_management](data-sources--protected_application--reference--group-003.md#canonical-1201000230102002-3032303001221220-1011311033111201-0122121320212232-1310132100233311-0231201102223133-3031233021232123-3001002122032031): complete subsection reference.

- [authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201): complete subsection reference.

- [financial_services](data-sources--protected_application--reference--group-003.md#canonical-0011220233111020-1330012103322331-2202112020122103-0303112231301000-2010112002103002-1332331220303123-0303301320202331-2030123332030203): complete subsection reference.

- [flight](data-sources--protected_application--reference--group-003.md#canonical-3022031310300013-1303031212122003-1101220112133123-2031112213031201-2311331321103330-2230221333213230-0213031113013220-0222011332013103): complete subsection reference.

- [profile_management](data-sources--protected_application--reference--group-003.md#canonical-2101313131303013-1011212123321211-1000323123112313-0033102233312313-3202223332111233-2301310032010101-2021313232221011-2001032013103303): complete subsection reference.

- [search](data-sources--protected_application--reference--group-003.md#canonical-1313223000232301-2211001332221312-1323112313113301-2113202013233212-1110323031322032-3101232010200320-3111103203333020-3311313011210003): complete subsection reference.

- [shopping_gift_cards](data-sources--protected_application--reference--group-004.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321): complete subsection reference.

<a id="canonical-1201000230102002-3032303001221220-1011311033111201-0122121320212232-1310132100233311-0231201102223133-3031233021232123-3001002122032031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.account_management` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- CloudFront.protected_endpoints.flow_label.account_management

<a id="canonical-3200011001300121-1012011300020310-1220233013302021-3201222303301312-3211000322133331-3201312001031220-1222300101012023-2220323221201110"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Account Management Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"create\",\"password_reset\"]"
}
```

<a id="canonical-0232002020310010-3133322030212111-3102123100323020-1200321301320211-0122230311210001-2003120120322221-3331201222230213-2123102311202120"></a>

### Direct properties for `cloudfront.protected_endpoints.flow_label.account_management`

- [create](data-sources--protected_application--reference--group-003.md#canonical-3210110300012033-1011300211230312-1131311032211333-0002200301302201-1300022300122112-1103112130210202-2032110033322133-3331233030023101): complete subsection reference.

- [password_reset](data-sources--protected_application--reference--group-003.md#canonical-0101310030221112-2222321012133122-1200303032322302-2212230000301113-0312010320102123-1130320133202333-2120310213213133-0122213121333010): complete subsection reference.

<a id="canonical-3210110300012033-1011300211230312-1131311032211333-0002200301302201-1300022300122112-1103112130210202-2032110033322133-3331233030023101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.account_management.create` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.account_management](data-sources--protected_application--reference--group-003.md#canonical-1201000230102002-3032303001221220-1011311033111201-0122121320212232-1310132100233311-0231201102223133-3031233021232123-3001002122032031)
- CloudFront.protected_endpoints.flow_label.account_management.create

<a id="canonical-1201211312121211-1211310230220101-2220221303212022-2331111113122211-1220302231220100-0023210303121113-2212103013222033-3202101121310120"></a>

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

<a id="canonical-0101310030221112-2222321012133122-1200303032322302-2212230000301113-0312010320102123-1130320133202333-2120310213213133-0122213121333010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.account_management.password_reset` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.account_management](data-sources--protected_application--reference--group-003.md#canonical-1201000230102002-3032303001221220-1011311033111201-0122121320212232-1310132100233311-0231201102223133-3031233021232123-3001002122032031)
- CloudFront.protected_endpoints.flow_label.account_management.password_reset

<a id="canonical-3021002030312000-1311121111111201-2210211023013021-3133033102312302-3113131211312223-0323023003100120-2202320032213223-1000020231130232"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for password reset.

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

<a id="canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.authentication` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- CloudFront.protected_endpoints.flow_label.authentication

<a id="canonical-0133021313100221-0203311121003231-0123310101323110-3033333332123233-2021031323112332-0132232031002103-2103131331030100-0020021033001330"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Authentication Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"login\",\"login_mfa\",\"login_partner\",\"logout\",\"token_refresh\"]"
}
```

<a id="canonical-2001111013103100-2232221332003303-1331232323222103-0302132200000013-0313300313122033-2331203320313022-1213131313320021-2101210322201131"></a>

### Direct properties for `cloudfront.protected_endpoints.flow_label.authentication`

- [login](data-sources--protected_application--reference--group-003.md#canonical-3002133111222112-1231032002112122-3211131323230030-3013222033101021-0010013323132210-0231021130112001-2123321110123031-3323211011003222): complete subsection reference.

- [login_mfa](data-sources--protected_application--reference--group-003.md#canonical-0220321331123103-1002301300010023-2230032002223001-2221132032223202-2023100032222223-2301033221033123-3130101003320130-1003201201213100): complete subsection reference.

- [login_partner](data-sources--protected_application--reference--group-003.md#canonical-3031021001023310-1001303023303121-1333220222000312-3200311231032211-3023113112132201-3310303303233002-2002033331110212-1331321123030101): complete subsection reference.

- [logout](data-sources--protected_application--reference--group-003.md#canonical-1012212123203032-3332021312221320-2321023222231133-0101332313111032-2310012102321122-0310023322332100-1303310131212322-0131312113132230): complete subsection reference.

- [token_refresh](data-sources--protected_application--reference--group-003.md#canonical-2222011203231312-1332202010300230-2310220202122201-3231210200021000-1011313301220101-3100111223102101-0310312111022023-1330011320301212): complete subsection reference.

<a id="canonical-3002133111222112-1231032002112122-3211131323230030-3013222033101021-0010013323132210-0231021130112001-2123321110123031-3323211011003222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.authentication.login` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- CloudFront.protected_endpoints.flow_label.authentication.login

<a id="canonical-1012103133031001-2312020033233231-0020320003313002-3020021322310220-0202223030120333-0201212223222101-0212211012133222-2022113330310122"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-transaction_result_choice": "[\"disable_transaction_result\",\"transaction_result\"]"
}
```

<a id="canonical-2302003202331120-1222200302103023-1122221323221203-2101203313111211-0231322301321100-3323303212200331-3332022203233303-2120110323330000"></a>

### Direct properties for `cloudfront.protected_endpoints.flow_label.authentication.login`

- [disable_transaction_result](data-sources--protected_application--reference--group-003.md#canonical-3122233301311321-0301320121032303-1300203222021111-1232303023101211-2132111311312200-3220121032232103-2221301133323323-0103022220203210): complete subsection reference.

- [transaction_result](data-sources--protected_application--reference--group-003.md#canonical-0013020233033310-3302013320031033-1303312000233302-1213302001321311-3221032102001213-3222212322133223-0002103132310330-0232001010120022): complete subsection reference.

<a id="canonical-3122233301311321-0301320121032303-1300203222021111-1232303023101211-2132111311312200-3220121032232103-2221301133323323-0103022220203210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-3002133111222112-1231032002112122-3211131323230030-3013222033101021-0010013323132210-0231021130112001-2123321110123031-3323211011003222)
- CloudFront.protected_endpoints.flow_label.authentication.login.disable_transaction_result

<a id="canonical-0122200211003002-0130212322323301-1102122331111333-1022123311312002-3333100312023032-2121102301313122-2110203331120233-0200020230232231"></a>

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

<a id="canonical-0013020233033310-3302013320031033-1303312000233302-1213302001321311-3221032102001213-3222212322133223-0002103132310330-0232001010120022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-3002133111222112-1231032002112122-3211131323230030-3013222033101021-0010013323132210-0231021130112001-2123321110123031-3323211011003222)
- CloudFront.protected_endpoints.flow_label.authentication.login.transaction_result

<a id="canonical-2212111022132332-2202111122203111-3101231323033320-1020002223321303-3111111021122133-3303310301331013-3013112101123012-2310000021322132"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result Type. Bot Defense Transaction ResultType.

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

<a id="canonical-0220000200220011-1302021300303333-0121021200301102-1320213220023333-3321133111111123-0130320213011200-1102011312001220-1031110023302201"></a>

### Direct properties for `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result`

- [failure_conditions](data-sources--protected_application--reference--group-003.md#canonical-3130100221131213-2230321323333020-0320311132123322-2131333203032003-2303313031022313-1231112323232023-0030121210320003-2232000102332203): complete subsection reference.

- [success_conditions](data-sources--protected_application--reference--group-003.md#canonical-3230002002200023-1301221332312103-0112220022121110-1112223210312001-3220313313233103-3132033120232220-0203320100103002-3122321212032121): complete subsection reference.

<a id="canonical-3130100221131213-2230321323333020-0320311132123322-2131333203032003-2303313031022313-1231112323232023-0030121210320003-2232000102332203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-3002133111222112-1231032002112122-3211131323230030-3013222033101021-0010013323132210-0231021130112001-2123321110123031-3323211011003222)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](data-sources--protected_application--reference--group-003.md#canonical-0013020233033310-3302013320031033-1303312000233302-1213302001321311-3221032102001213-3222212322133223-0002103132310330-0232001010120022)
- CloudFront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions

<a id="canonical-3310220233022301-2120232333033111-2223213113122231-0133320130220230-3333333122031123-1222300203213213-2022013001223211-0213100330200111"></a>

Type: `"list"`. Computed.

Failure Conditions. Failure Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0020310003001032-1311220301110213-1022100131133101-1133230220001310-3010323311003330-1133313222122122-3220122200300002-3121120031233032"></a>

### Direct properties for `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions`

<a id="canonical-0320230031002303-2222132230212120-2320030213131023-1320022222222023-2301202223300311-2220210231110221-3010023030120322-3112112210131202"></a>

#### `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.name` property

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1203203012103232-1030222322300120-3200233103020223-2120203231233323-3123202221003323-0232312200130221-2210201122230110-0313313001100030"></a>

<a id="canonical-1032233102312003-1022213222103321-1130231103001333-2331312330200213-0021130111000301-2122132100113001-2221003300122321-1102230031002020"></a>

#### `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2223312011002310-1002321110003230-3313303301231320-1003102202313112-3313130310322121-3303223313303303-0331302230233012-2311011331222120"></a>

<a id="canonical-0110123222110200-2213221101032331-1322013211321222-2223112301231310-0310303202330121-3331323311011220-2310231301201332-0110112002231012"></a>

#### `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.status` property

Type: `"string"`. Computed.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Additional upstream details:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3230002002200023-1301221332312103-0112220022121110-1112223210312001-3220313313233103-3132033120232220-0203320100103002-3122321212032121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-3002133111222112-1231032002112122-3211131323230030-3013222033101021-0010013323132210-0231021130112001-2123321110123031-3323211011003222)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](data-sources--protected_application--reference--group-003.md#canonical-0013020233033310-3302013320031033-1303312000233302-1213302001321311-3221032102001213-3222212322133223-0002103132310330-0232001010120022)
- CloudFront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions

<a id="canonical-3012221033231200-0121302301101011-1033122332123023-0011011020011121-0030030132332003-0330201132202312-2002330030231131-2313312232120302"></a>

Type: `"list"`. Computed.

Success Conditions. Success Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1133001331310312-2200223221010213-1232323011203110-1323320022121122-3011303222222330-3230111211300232-2332020011030211-3123033223001020"></a>

### Direct properties for `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions`

<a id="canonical-3003122003321110-2031233012202203-3021001023022032-0300220330000103-2022201123303110-1221232102123010-2331213333021013-3201122013303311"></a>

#### `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.name` property

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0011213333313120-2230221211101311-0130111211213103-3331022332001203-2322021122111301-2131221013112032-1211210223132221-2321033333331320"></a>

<a id="canonical-1033310322213000-1121112300111212-3201230103223113-0002302332003032-2000131221330013-3123333133012311-2003033312103200-3313030320001213"></a>

#### `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2003212000301303-2113020321331212-1211311320301102-2332303030121022-2001302112233300-2121020303201000-2201033012102021-0323133011332232"></a>

<a id="canonical-3303111303300310-0030032202220132-1001110122123210-1111311021032232-2101210322111032-0202003000232020-0333223321301010-1001303103023310"></a>

#### `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.status` property

Type: `"string"`. Computed.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Additional upstream details:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0220321331123103-1002301300010023-2230032002223001-2221132032223202-2023100032222223-2301033221033123-3130101003320130-1003201201213100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.authentication.login_mfa` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- CloudFront.protected_endpoints.flow_label.authentication.login_mfa

<a id="canonical-3220033113120103-2101233210303221-1002232301123312-0321023123232323-0121012110130333-0223130032323223-2230023112202232-3222033332211331"></a>

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

<a id="canonical-3031021001023310-1001303023303121-1333220222000312-3200311231032211-3023113112132201-3310303303233002-2002033331110212-1331321123030101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.authentication.login_partner` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- CloudFront.protected_endpoints.flow_label.authentication.login_partner

<a id="canonical-0302023012233013-3221000023200010-2031010330020232-1002320220210023-1003200130132023-1131132111021120-2113330032021110-0320331232321001"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for login partner.

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

<a id="canonical-1012212123203032-3332021312221320-2321023222231133-0101332313111032-2310012102321122-0310023322332100-1303310131212322-0131312113132230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.authentication.logout` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- CloudFront.protected_endpoints.flow_label.authentication.logout

<a id="canonical-1002313002311101-0331131111301323-0133232313003312-2222310222101123-0023101303203211-0133103013010102-1330011002323230-1102111211303301"></a>

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

<a id="canonical-2222011203231312-1332202010300230-2310220202122201-3231210200021000-1011313301220101-3100111223102101-0310312111022023-1330011320301212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.authentication.token_refresh` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- CloudFront.protected_endpoints.flow_label.authentication.token_refresh

<a id="canonical-2301231112210332-3333321122233312-0221330323331200-3233330003013000-2132333300331222-3320331231013100-2121203323120331-0220220300031002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for token refresh.

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

<a id="canonical-0011220233111020-1330012103322331-2202112020122103-0303112231301000-2010112002103002-1332331220303123-0303301320202331-2030123332030203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.financial_services` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- CloudFront.protected_endpoints.flow_label.financial_services

<a id="canonical-2222023211012313-1012110111322331-1201302320020311-1020001302312203-0230132020032331-0231021211110203-2133200311112303-1100233031130313"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Financial Services Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"apply\",\"money_transfer\"]"
}
```

<a id="canonical-1332021230101011-2333333103231200-2222120130322122-0301031332130310-0032312200132030-2221122132111031-3031110000303121-0332003133203010"></a>

### Direct properties for `cloudfront.protected_endpoints.flow_label.financial_services`

- [apply](data-sources--protected_application--reference--group-003.md#canonical-1332030300300100-2001130100010303-1321300101122113-0311303112010101-2232331302131321-1300211202320132-1101002131201122-0002232320123022): complete subsection reference.

- [money_transfer](data-sources--protected_application--reference--group-003.md#canonical-2232301131101210-0223321230301213-0112322230101032-2223033302333313-3221033230200213-0323102231333133-0301132200201000-0310332112300230): complete subsection reference.

<a id="canonical-1332030300300100-2001130100010303-1321300101122113-0311303112010101-2232331302131321-1300211202320132-1101002131201122-0002232320123022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.financial_services.apply` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.financial_services](data-sources--protected_application--reference--group-003.md#canonical-0011220233111020-1330012103322331-2202112020122103-0303112231301000-2010112002103002-1332331220303123-0303301320202331-2030123332030203)
- CloudFront.protected_endpoints.flow_label.financial_services.apply

<a id="canonical-1132102021302213-0122321030023122-1330021103020030-3302210322132330-0333200312300333-0012302301021030-3303312321330010-1123320001302102"></a>

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

<a id="canonical-2232301131101210-0223321230301213-0112322230101032-2223033302333313-3221033230200213-0323102231333133-0301132200201000-0310332112300230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.financial_services.money_transfer` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.financial_services](data-sources--protected_application--reference--group-003.md#canonical-0011220233111020-1330012103322331-2202112020122103-0303112231301000-2010112002103002-1332331220303123-0303301320202331-2030123332030203)
- CloudFront.protected_endpoints.flow_label.financial_services.money_transfer

<a id="canonical-3302201311301221-1330002120222203-1313200313230013-3030023120220001-0002203001221011-0100100021200030-1211100222111002-2111232321310021"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for money transfer.

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

<a id="canonical-3022031310300013-1303031212122003-1101220112133123-2031112213031201-2311331321103330-2230221333213230-0213031113013220-0222011332013103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.flight` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- CloudFront.protected_endpoints.flow_label.flight

<a id="canonical-1232122302233330-0213332221113131-2320101130100000-1112203011031213-1222121212102130-0330113001231220-3132010210130132-1101310120300233"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Flight Category. Bot Defense Flow Label Flight Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"checkin\"]"
}
```

<a id="canonical-1220123220210031-1102331332213103-2323310222203330-3022103302132211-2303131010001013-0032100133120310-0303203002012323-1031302131212121"></a>

### Direct properties for `cloudfront.protected_endpoints.flow_label.flight`

- [checkin](data-sources--protected_application--reference--group-003.md#canonical-3023330122313330-2320320312330330-3312003310123233-2021330131000111-2302303131312313-3222131232323001-0220011213122221-0333101033212013): complete subsection reference.

<a id="canonical-3023330122313330-2320320312330330-3312003310123233-2021330131000111-2302303131312313-3222131232323001-0220011213122221-0333101033212013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.flight.checkin` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.flight](data-sources--protected_application--reference--group-003.md#canonical-3022031310300013-1303031212122003-1101220112133123-2031112213031201-2311331321103330-2230221333213230-0213031113013220-0222011332013103)
- CloudFront.protected_endpoints.flow_label.flight.checkin

<a id="canonical-3202003000010110-1333100201101313-0201201131310222-0033123002200113-2223002211020331-2010103312103302-0111312221103301-3102220011320322"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2101313131303013-1011212123321211-1000323123112313-0033102233312313-3202223332111233-2301310032010101-2021313232221011-2001032013103303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.profile_management` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- CloudFront.protected_endpoints.flow_label.profile_management

<a id="canonical-2303110110213010-0232001133101302-3333101000112220-0120202302203331-2300031022332211-2333332101033002-0131203101001320-2320311002011232"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Profile Management Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"create\",\"update\",\"view\"]"
}
```

<a id="canonical-0131220113230320-3120210030302131-2302232013132201-3332000133002220-2300012102210001-1310131000312210-2302220113130121-0020031230132012"></a>

### Direct properties for `cloudfront.protected_endpoints.flow_label.profile_management`

- [create](data-sources--protected_application--reference--group-003.md#canonical-2232302021010323-0301323323203311-2200330211003110-3221301223102230-3231113312102122-2031123321101113-1222102113210223-2213313230111321): complete subsection reference.

- [update](data-sources--protected_application--reference--group-003.md#canonical-1112121212322033-0313122000310130-1331020302112002-3011122213232100-1202213230302302-1020130102223133-2323113300121210-1301023320032233): complete subsection reference.

- [view](data-sources--protected_application--reference--group-003.md#canonical-0331031132000322-3212033022100310-2200111333110122-1213012103003213-1130233000213311-2211322131301013-3322303101002003-0213330200132221): complete subsection reference.

<a id="canonical-2232302021010323-0301323323203311-2200330211003110-3221301223102230-3231113312102122-2031123321101113-1222102113210223-2213313230111321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.profile_management.create` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-2101313131303013-1011212123321211-1000323123112313-0033102233312313-3202223332111233-2301310032010101-2021313232221011-2001032013103303)
- CloudFront.protected_endpoints.flow_label.profile_management.create

<a id="canonical-3033232101003130-0130121022102313-3003233210333122-3322323230200121-3022303311130021-3223022113223210-3332313020132301-1121121011203332"></a>

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

<a id="canonical-1112121212322033-0313122000310130-1331020302112002-3011122213232100-1202213230302302-1020130102223133-2323113300121210-1301023320032233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.profile_management.update` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-2101313131303013-1011212123321211-1000323123112313-0033102233312313-3202223332111233-2301310032010101-2021313232221011-2001032013103303)
- CloudFront.protected_endpoints.flow_label.profile_management.update

<a id="canonical-2023002023120122-0321323102011210-0021310330010132-1321001130212333-1323111030222233-0232213310111123-0320323213331010-2113210223003102"></a>

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

<a id="canonical-0331031132000322-3212033022100310-2200111333110122-1213012103003213-1130233000213311-2211322131301013-3322303101002003-0213330200132221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.profile_management.view` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-2101313131303013-1011212123321211-1000323123112313-0033102233312313-3202223332111233-2301310032010101-2021313232221011-2001032013103303)
- CloudFront.protected_endpoints.flow_label.profile_management.view

<a id="canonical-0011211333313021-3300023230323311-1323231023201311-3101010312130033-1030231332211103-1212120320103111-1212033200232001-3003232213002103"></a>

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

<a id="canonical-1313223000232301-2211001332221312-1323112313113301-2113202013233212-1110323031322032-3101232010200320-3111103203333020-3311313011210003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.search` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- CloudFront.protected_endpoints.flow_label.search

<a id="canonical-2230122110213122-1110120131032001-3233210112103300-2013030213210220-2010002313212221-2311233132223310-2030032121230020-3001013202212023"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Search Category. Bot Defense Flow Label Search Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"flight_search\",\"product_search\",\"reservation_search\",\"room_search\"]"
}
```
