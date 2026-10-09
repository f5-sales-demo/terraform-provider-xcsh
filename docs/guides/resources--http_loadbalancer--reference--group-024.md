---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0003301213131101-1312033131333213-2101303001113021-1023212221120031-1331213111320113-3131102321210212-0323000011233200-3231112332221032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.asn_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.asn_list

<a id="canonical-3313120122012310-1312303301112131-0330231303213310-2123203000010211-1300220010012002-1202311122313213-0122023202310323-1323200031333331"></a>

Type: `"object"`. single nested block, Optional.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122232220022111-2313330200023223-2002212301312012-0233110300122110-1301102030012312-3332321332320002-3130223121102311-2100013200220030"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.asn_list`

<a id="canonical-2011021001231323-2000131220333220-1213030231001021-2222211233332011-3200203011230231-1032303202201010-3033311321013300-3033212211221031"></a>

#### `policy_based_challenge.rule_list.rules.spec.asn_list.as_numbers` property

Type: `["list", "number"]`. Optional.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2330232220313013-2120000211302223-3012212202222313-1113232301013021-3231202321121233-0332211110312023-1013001211321310-2012322132110120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.asn_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.asn_matcher

<a id="canonical-3332020131111132-3123312111131233-0221120323023112-0032020302001230-1123232131010332-2012320311231101-1302113001110312-3223032100201333"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131123003101103-1001130313231203-0302303201002010-3022100130331111-2133311213231233-1033310021312202-3013012121222013-3332222022130311"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.asn_matcher`

- [asn_sets](resources--http_loadbalancer--reference--group-024.md#canonical-0331010323003123-2200312103010033-1210131010130000-2103022000112201-3321121120330130-1022200100002313-0113122201303213-0203132322011021): complete subsection reference.

<a id="canonical-0331010323003123-2200312103010033-1210131010130000-2103022000112201-3321121120330130-1022200100002313-0113122201303213-0203132322011021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](resources--http_loadbalancer--reference--group-024.md#canonical-2330232220313013-2120000211302223-3012212202222313-1113232301013021-3231202321121233-0332211110312023-1013001211321310-2012322132110120)
- policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets

<a id="canonical-0013021303033210-1310033201100322-3122233210301123-1330330222000221-3003222212323023-1000102221322202-0121222022133110-0330112311003203"></a>

Type: `"object"`. list nested block, Optional.

A list of references to bgp\_asn\_set objects.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332102123323020-3310120000112032-2113123312312323-2101132113131211-0033212110333232-0231330122031101-1111023020011202-2203333010232331"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets`

<a id="canonical-0320221303031113-1201232302222013-2220013000231333-0231101013122131-1232213201030302-3013222323303001-0021302230012310-0332321322121202"></a>

#### `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0230311112202302-2301012301101323-1311131202301211-2200212332113200-1313003010320113-0013223202321213-1112332101330033-0032222300101331"></a>

<a id="canonical-3302113112310220-3121121123332011-3311210022200113-2001102001133302-2233123313303221-1212300101301000-1200200123011322-1222302310220213"></a>

#### `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3212100010123103-2201132023213020-1211212303303201-2123031302211020-2013030122021222-1201023120030321-3000022303233000-3323331201010001"></a>

<a id="canonical-3320122012320132-0202320012122131-0222311000000121-3012332300300322-3010221020002101-3121232001032132-1023113201302313-1321301020212313"></a>

#### `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1220333021333210-3132101033202303-3103321130213330-3233000200202023-0123113103103333-1200202320212323-2130233111210101-2330331333021011"></a>

<a id="canonical-0300213200120212-0123302231201203-3233312031002311-0211102003111120-1311030013032310-2202000211131023-0303111203212021-0300131330131301"></a>

#### `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0212030212030313-1320332202232022-1320023311300203-1131000132123221-1222323212112202-0001200211302323-0233231320133300-0300232321131022"></a>

<a id="canonical-2312012303230012-1311213303330030-0211111103303233-1220321213132132-0133023121122111-1211011010311111-1223022030120032-2302032203313023"></a>

#### `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1322031002303310-1111010110113010-3010000120211223-0112031202323033-3213001011000120-0032032230230101-3220223203111233-0023030020301031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.body_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.body_matcher

<a id="canonical-1120202020002120-0033002221031111-1101231111313331-0021301303031310-2020011020322002-3003230313003111-2203330013030122-0031120312000232"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
body_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300031320131112-3002112103211213-1312300331303122-0333230213013133-3323201332002202-3023103313131333-2330122111012121-2131113320221102"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.body_matcher`

<a id="canonical-1222110331201200-0100012130110120-0311010032103212-2031202223133001-3231120223330201-0232332132131331-0103113100303210-3212301233203210"></a>

#### `policy_based_challenge.rule_list.rules.spec.body_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2101100123321313-3131330133110131-1011331311203301-2211111320222320-3130322230311101-1210230132203313-3132331112012122-2220132303230232"></a>

<a id="canonical-1223231322033311-1030001011232201-2121132101000131-2333313100323020-3020002333233102-1302123222012323-0022323023101333-2002300303030131"></a>

#### `policy_based_challenge.rule_list.rules.spec.body_matcher.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1212100332130320-1012101311113323-1300030002112003-2133222013231000-3333013213332003-3033311011131030-0211010201021202-3333323210232100"></a>

<a id="canonical-0312111021133203-2210301111312123-0101311333122201-0203010133211100-0320301222332031-2302301312223110-3222301232330101-0330212010103302"></a>

#### `policy_based_challenge.rule_list.rules.spec.body_matcher.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3301001033202231-2133133323213110-2321233202103301-1310201321100023-1302201121201002-3200232221023023-0110031303001223-0212313222223133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.client_selector` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.client_selector

<a id="canonical-0121031210120212-2200023300010320-1303230332213032-2021223200323330-3021322303321000-3022022012331003-0202202311020323-2011303233000013"></a>

Type: `"object"`. single nested block, Optional.

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003111203100113-2100031022032131-2303230131013220-2223113232220002-2321003020230210-3003101202312112-2122203331202031-2331101203001132"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.client_selector`

<a id="canonical-0302312133102010-0010310201312322-2202231331120321-0223332301122022-1230212203132110-0022023102313323-3201133310001231-1223002021111220"></a>

#### `policy_based_challenge.rule_list.rules.spec.client_selector.expressions` property

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers

<a id="canonical-2311302301030200-3113330322112312-3320022100223123-0111322103012232-3103312312332003-3310211202231110-2203000200033132-2003313132233310"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
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

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212230230323221-2130110112023221-0120331300331232-1001103221301133-3010102222223120-3232211311030100-0222312002030101-0111220332200213"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.cookie_matchers`

- [check_not_present](resources--http_loadbalancer--reference--group-024.md#canonical-0103313330202032-0333323222300211-3231010121102231-1022300331232331-1311002322210132-3210201002132031-2200001313310231-1132120002030213): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-024.md#canonical-1011011121103030-1023321311310333-2210310200001020-2031223022321101-0211112103002331-0113300012011220-0213102312210012-3331322211331322): complete subsection reference.

<a id="canonical-2320031132013223-2200212003011211-1122112213300011-3222201101010102-1323020003210220-3100132203131102-2000300332231321-3032220311311221"></a>

<a id="canonical-3002210022200332-2222312020011021-0332020211113222-3212033201111133-3213101231230110-3222031333222202-3312333232032233-0022013233230133"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.invert_matcher` property

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

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

- [item](resources--http_loadbalancer--reference--group-024.md#canonical-2301120303000033-1120110232011011-0311033303003031-3013101323203110-3103223000001020-2012122203222232-3020031332212020-3132031221321231): complete subsection reference.

<a id="canonical-2320220323220232-0222321220003111-1032113000013020-1302313312110003-3012331330201301-3333220103101322-3232200232133112-1121332003000322"></a>

<a id="canonical-3230203020331111-3113323233103233-0311121130001332-3320010001133222-3331111331133323-0320012020130333-1100210211221310-2313101122303120"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.name` property

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0103313330202032-0333323222300211-3231010121102231-1022300331232331-1311002322210132-3210201002132031-2200001313310231-1132120002030213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-024.md#canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present

<a id="canonical-1301310221031133-3232011221111032-2201210210121031-1210020202020322-1203132333021121-0032111003223002-1203222123220112-3111022203232100"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011011121103030-1023321311310333-2210310200001020-2031223022321101-0211112103002331-0113300012011220-0213102312210012-3331322211331322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-024.md#canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present

<a id="canonical-2122002112213001-1220221103323200-3311212230131022-3010330110220012-2300122233332000-1300013310123211-1000320321023032-1223110132111310"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301120303000033-1120110232011011-0311033303003031-3013101323203110-3103223000001020-2012122203222232-3020031332212020-3132031221321231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-024.md#canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.item

<a id="canonical-0233310200121001-1010031310210020-1002330202312030-0112101200233223-2100130121231000-0033203200213032-0002230032123233-0002301313312010"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310230223303231-3123232210103201-2330121313033000-0133213022202002-1023022311100332-1313003031312333-2000231111203212-0203231223131111"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item`

<a id="canonical-0211321132033130-1330031123321102-1102111320112213-3222023003021013-2333321132233100-3123320101220031-1100321100123210-3210203310010021"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3132130322300322-2211222320230323-0232220200002330-2132033130013202-3320003301322210-2122030300203032-1130233112103121-0133012122111020"></a>

<a id="canonical-1033301233101110-3223213110013002-2103113302301020-0212212223230303-1333202300223123-2130211133302030-2223313020102131-1210011020233233"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1033022011011132-0030003332130331-2231313220102023-1302201010311232-3121220103132032-2211210132311132-2332233333322033-0321003333232312"></a>

<a id="canonical-1022120132220033-2023302023023331-3011331233033112-2213102232021201-3312223211220020-2031330221131101-1130111201201020-0013021200310033"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0202302323030031-0332332132022130-1301100023003233-2302201010011300-1300320313311232-3332303233323013-2021332220011213-0032313121021322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.disable_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.disable_challenge

<a id="canonical-2303302200301023-0230123210232312-3110113233311112-2211230302131132-3222201213233310-3211011001032212-2100113110000000-1021122133032231"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable challenge.

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
disable_challenge = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231322321300130-2211003200213110-0230120220232310-0330031202332122-1202330333221203-2022232030123301-0132023201301310-1122331113003030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.domain_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.domain_matcher

<a id="canonical-1232323002333033-0112032300221013-1133320131331310-0103302201222000-3001320023012200-2112013212022221-0032101200212311-2212321012220203"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
domain_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301200211320232-1120113123120033-3011002003133223-1311033112021233-0312010303020113-1002233031132023-3323002333313301-2001310032111021"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.domain_matcher`

<a id="canonical-0013211133032103-3232013113002302-3230113022221032-1132211331203001-3113023233023301-0103120223333213-1020212133232220-3010130102010320"></a>

#### `policy_based_challenge.rule_list.rules.spec.domain_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2202110300322313-2013321002200231-1311231313202100-0132003012202131-1122113020331011-0300000223013202-3213010220210303-3213212130020003"></a>

<a id="canonical-2130000102131321-2203131013303102-2330213000003203-2121011030102202-2312310211311121-0013021303123231-0120023113312203-2011203233211112"></a>

#### `policy_based_challenge.rule_list.rules.spec.domain_matcher.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2331103200120201-0301001021002322-0132233020221032-2203102002320022-3001110313102332-0311012113002330-0202002031220211-1023110131120121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge

<a id="canonical-2132331203310320-1222200200202033-0011330110111031-0312132201102202-0013320113013122-3102310210001133-1120212312030103-0100322013013022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable captcha challenge.

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
enable_captcha_challenge = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221112113201102-3022103132211101-1320220303332222-2303330112030301-1020203320020311-2100330302222210-1012001232303023-3130312322023023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge

<a id="canonical-1130202110233103-2131102003313220-3030212130333030-2001123211330000-1012230010300023-2231123011102202-3222120333303323-1012211201202020"></a>

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
enable_javascript_challenge = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313201031011033-0203021101211130-2233102100210202-2231200332022130-0212312030111003-0030330021121330-0011012123121011-2011131311103111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.headers

<a id="canonical-3213121202022121-0002321330203123-0233123032313100-2330323312031012-3330313002033230-0330302002020002-2201311111210003-0302110231212101"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1332133313121102-0030220013123001-3301211213312110-0012131033113322-0101323000121030-1313332032100231-0301213222200002-0223130120112132"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.headers`

- [check_not_present](resources--http_loadbalancer--reference--group-024.md#canonical-2323020220021030-1221122111120010-1000002310111301-3103132212310231-3222000202211323-2223022323033322-3200012012111322-2010121000223113): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-024.md#canonical-1211303211131310-1223121301221223-0331003023331300-3013201222033021-3213232122201131-1201322000101313-3231203021010120-2301213312132013): complete subsection reference.

<a id="canonical-1022001022301210-0301201121032130-1210031220100332-0233100233200020-0212020022312231-2111301031301200-3110000333023133-0130122100020212"></a>

<a id="canonical-1000021210212322-2221202301011212-3323301220003110-3023311102021310-2210033303112100-0210202222132233-3200121201230330-0101030031012210"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.invert_matcher` property

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-024.md#canonical-3000212303030230-2323101000310303-2133302302323120-1302012010122311-3230232013302223-2221311321310131-0313233101301123-0123020121000020): complete subsection reference.

<a id="canonical-0223111223131203-3012133001321010-0203303210331200-0332022302323220-2010233130030112-3211000312013102-1331130102132110-2202310222100323"></a>

<a id="canonical-2111211221322210-0023111231130122-3032203032030000-1012231210030220-0001011003332301-1112202233312322-1212302021033100-0030311222130323"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2323020220021030-1221122111120010-1000002310111301-3103132212310231-3222000202211323-2223022323033322-3200012012111322-2010121000223113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-024.md#canonical-1313201031011033-0203021101211130-2233102100210202-2231200332022130-0212312030111003-0030330021121330-0011012123121011-2011131311103111)
- policy_based_challenge.rule_list.rules.spec.headers.check_not_present

<a id="canonical-0002300313103112-3301102110122300-3301202122012112-1033233202110120-0000213133220003-3202030031123301-2123031230210031-3211322121300300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211303211131310-1223121301221223-0331003023331300-3013201222033021-3213232122201131-1201322000101313-3231203021010120-2301213312132013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.headers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-024.md#canonical-1313201031011033-0203021101211130-2233102100210202-2231200332022130-0212312030111003-0030330021121330-0011012123121011-2011131311103111)
- policy_based_challenge.rule_list.rules.spec.headers.check_present

<a id="canonical-1321330332113011-2101112110100033-3111113123330110-1121210322220332-3120030323110122-3002022122200332-3323113222112222-0331332222211122"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000212303030230-2323101000310303-2133302302323120-1302012010122311-3230232013302223-2221311321310131-0313233101301123-0123020121000020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.headers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-024.md#canonical-1313201031011033-0203021101211130-2233102100210202-2231200332022130-0212312030111003-0030330021121330-0011012123121011-2011131311103111)
- policy_based_challenge.rule_list.rules.spec.headers.item

<a id="canonical-0111123330021001-3132230101323211-0013032210303001-1321303020130103-3312300320201001-0213013021220112-1003300220322122-0231003310002013"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310101323121001-0211110331100322-1312012100130231-2010100321320303-0110122012120201-2230301111233330-0330121003121120-2223212011222032"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.headers.item`

<a id="canonical-0333223002223333-1030033220212320-0230212212132010-0113330322320013-3211112120311110-2300121330023322-2213332311201021-3221110110100001"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1121113112100010-1230120030331313-1122012133110102-1033030322321112-2332310033300322-1230323031111201-2123301113133002-2333321220010311"></a>

<a id="canonical-0303111321112133-0211133322230330-0103333121333212-0130011231113221-0012032233212201-3320300220220032-0332002102113230-1033021332202233"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1211321000113030-1230333002310302-1211233023011231-3322213033113020-0231021023202210-0010303132011122-2021323311133122-2130123112012231"></a>

<a id="canonical-0300100100011322-3323230113201313-2120133313021112-1110201303331312-3013213111332323-3330312010121031-2321113003001333-2022321330121331"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2330011102111203-3200210100131120-2112200113232323-3120021212132232-3203312330132100-2321330223303101-1023321200013301-2303010303011323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.http_method` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.http_method

<a id="canonical-0320200311032022-0313303221002101-1020100111313001-1023123212203203-2012111201110121-2320030130233021-2202110322203113-3321111323232030"></a>

Type: `"object"`. single nested block, Optional.

An HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

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
http_method {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302310112031223-3010121020010212-1100133110333102-1033122102232203-0331203313322030-3102331103212303-0323320313002300-2112213220132102"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.http_method`

<a id="canonical-2322231000110332-1013032003311300-2233231232011320-0112133100311110-3121110121213321-3120203021201033-3301000023001232-3123311011303113"></a>

#### `policy_based_challenge.rule_list.rules.spec.http_method.invert_matcher` property

Type: `"bool"`. Optional.

Invert Method Matcher. Invert the match result.

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

<a id="canonical-0233332131131232-1200032030013311-3130032222102322-1300103210133013-2200101000121211-0020033303302232-3203123111323322-3001030103020221"></a>

<a id="canonical-1003231001311033-1302003133220021-1302010012003120-1033033010210322-0012131211210001-1112111100302113-0311301313311000-3133023001323010"></a>

#### `policy_based_challenge.rule_list.rules.spec.http_method.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3131232312332313-2331122113110111-0131010211100102-0310213132121112-3111200310110021-1010121212033013-3232311000121330-1220231122112110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.ip_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.ip_matcher

<a id="canonical-0023210031221301-0320301010030320-2003132110122223-3200103210200311-1321320111103111-2020021032132221-2030003211130130-1331022000121320"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

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
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311232001321213-0102202033130230-3310323102031032-3003233211221223-0200323302030220-3313202302231302-3012232110230012-0331311130301020"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.ip_matcher`

<a id="canonical-3132120210121311-3032132212210022-0301222301010110-3322331131000322-0302330031022333-1232013331133202-3003211233330032-2222122322110333"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.invert_matcher` property

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

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

- [prefix_sets](resources--http_loadbalancer--reference--group-024.md#canonical-0212332012223100-0032200302201123-2123130003231102-1023020003022121-1133332212200021-0132021010113201-0202231131023203-2100111123101220): complete subsection reference.

<a id="canonical-0212332012223100-0032200302201123-2123130003231102-1023020003022121-1133332212200021-0132021010113201-0202231131023203-2100111123101220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](resources--http_loadbalancer--reference--group-024.md#canonical-3131232312332313-2331122113110111-0131010211100102-0310213132121112-3111200310110021-1010121212033013-3232311000121330-1220231122112110)
- policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets

<a id="canonical-0330020020230210-0021231320002110-0010301220333003-3311213313022022-2332322121022201-0002110000300321-1310322230132333-1111103201211230"></a>

Type: `"object"`. list nested block, Optional.

A list of references to ip\_prefix\_set objects.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0333112213200013-2030313001330212-2131201001101331-0003130210230232-1210331010303300-3331103113222202-1130302102133133-2333201003310213"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets`

<a id="canonical-3013310222301300-1303103300311211-0110131313201031-2120010231023233-3133310030321131-0021102231212322-2220021133001110-0213221300022131"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0133303333302231-0011331323122221-1330200113023302-2003111302011220-2032230022132102-2332320102202322-2130133210111130-0103311111203010"></a>

<a id="canonical-3011000033323231-3203002033123133-2333333110102202-0100011120123203-2303302203311201-3320301010111030-0030022313211123-0211002210031030"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1202123301203013-3302310302113032-3031301202212123-0201123331033010-0312211221203321-3022113130100123-0303321313302113-0010010011133123"></a>

<a id="canonical-3202112310212033-1300000123333020-2203313020100120-3313030332221011-2102312320221222-0133320202101300-2032013010023101-0130033210101023"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2003201002111230-0301232123032323-0220212331012131-1330230313032132-0121223032031130-1201120231210300-2001102010222222-2100311113012312"></a>

<a id="canonical-1232212210323111-2212011212002321-0122201331000220-1113113212002302-0303230213303312-1221023113222230-3211130123101331-0220312202300221"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0332300213330031-3130122101212030-1013222320002312-0313110030003131-3033103122010133-0011311223023230-2113101223322201-1122313320133003"></a>

<a id="canonical-3121303220012122-1101003131010301-3101210121313001-2100003011222222-1213031033110101-3113133110112221-0003122131302002-0002100133133021"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1102303013213221-2101121232233030-3032103211220121-2031323223113213-2210001011111012-1001212110123010-0120200232130210-3120021232133212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.ip_prefix_list

<a id="canonical-3210033200001103-2012311001110030-3021022230112021-3121211323010123-1101313300033203-1101232313320203-1022013101101330-1100220210102312"></a>

Type: `"object"`. single nested block, Optional.

List of IP Prefix strings to match against.

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
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1113211320232312-2133312123123223-3201101113230213-3031131321020002-0023001232100112-1110330130202322-0310323211323313-1210232202120311"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.ip_prefix_list`

<a id="canonical-0023300312023321-1222112232000303-3312020333323100-3010130101213223-2102133130320000-0111122002211030-2021230230212113-0211012302331033"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_prefix_list.invert_match` property

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

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

<a id="canonical-0133032110130220-2310121223020020-0123310010210101-3003212322113100-0332332333110222-3011012202111312-0122310132302330-3303201322113222"></a>

<a id="canonical-1112202021011323-1110001102022230-3111011223220212-0321122321203031-3130121012301020-0111313322102010-3302233200213011-0303120130212311"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2110130302132222-1333323202323202-2101132211132333-3120123331020002-2321313212221101-0010000002320121-1033331331113211-1113310312332201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.path

<a id="canonical-1013203123330210-1012013200102220-0003230000130221-0032221333233303-0232201031022121-1113220202200323-1130210333213332-3201203112101200"></a>

Type: `"object"`. single nested block, Optional.

A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

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
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-2133220012331120-2122122233102121-3023302110001311-2010200200333100-3010013210303211-1033211101100330-2112133230221113-0323113121313221"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.path`

<a id="canonical-2322213223333133-2230111223213113-1101030332232301-2312111231210202-1001000210222102-1102333303220203-2232023020012330-1302120230032331"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.encoded_path_matcher` property

Type: `"bool"`. Optional.

Match against the encoded, escaped path.

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

<a id="canonical-2032022121322200-1132302300113032-1330302213013232-1212021303311003-3221103100300013-2321302312322201-0023000111101321-3101321303201233"></a>

<a id="canonical-2000210211012110-2111230122302121-1132332111210312-3033210212030213-1231121323202202-0111202231121310-3312310232022322-2112011232022002"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact path values to match the input HTTP path against.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0103011303301111-0020211031023001-0230102012233333-2211311222232032-2133103012030033-1300302330111333-3022200121020112-0013111211022231"></a>

<a id="canonical-1200231313123222-3013023230022331-2001023133233103-3331031011132121-2221130302320213-0223213323312022-0312332222113303-2221203110133210"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.invert_matcher` property

Type: `"bool"`. Optional.

Invert Path Matcher. Invert the match result.

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

<a id="canonical-1303031331230310-2001233201022112-0202233230120120-3101303133200212-1032201333023321-2120002230031330-1130220023203132-0023203311310311"></a>

<a id="canonical-1211100300033300-3002132113221111-3303001011132323-2221331013122012-2222002332302033-2010101213123301-2021003012322121-1020110211230030"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.prefix_values` property

Type: `["list", "string"]`. Optional.

A list of path prefix values to match the input HTTP path against.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1130102211330003-2222230333300210-3302213111012320-1003323322312130-0222303132322103-0331231303320230-1100120303100111-2321113312020121"></a>

<a id="canonical-3303123032221013-2211030022200201-2200303012133233-0012021213001013-1000310122000311-0132021110212202-1311200002130121-2201302013311111"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input HTTP path against.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0221232012101112-1313011111020110-2231103332101331-0310313113013220-2000010303033011-3221201010321131-3300132103331132-3222212030302220"></a>

<a id="canonical-3300131311312331-3212331322233102-0312222203003021-2121121222203312-3212013033230102-3000310000203310-2132112112002121-3120220213030102"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.suffix_values` property

Type: `["list", "string"]`. Optional.

A list of path suffix values to match the input HTTP path against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1312322031021220-0320122303112122-0320230103103300-3102031321330000-3110132213132230-0310331232221303-2132301001233230-0311302031231000"></a>

<a id="canonical-2032202003321212-3232011202101102-3001200023110003-2022212211313232-2230213300331013-1011333303202131-1031320133323102-0130112100113010"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.query_params

<a id="canonical-2203120210113211-2112000210103232-3102220020002112-2020330112232021-2101223003312230-0312303201312023-3031122000103101-3013000312102301"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
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

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102113321203001-1331211222203230-3223232032320001-3120100331313010-2312011112310032-3003003231020022-3212331203001023-3001110023212320"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.query_params`

- [check_not_present](resources--http_loadbalancer--reference--group-024.md#canonical-3103300010123120-3200220011010132-3033202200030111-2303133000000131-0311023323212320-0013122300031122-0222001002323220-2332111330223102): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-024.md#canonical-2112003030120300-3231323312231200-0322203333323223-3322103103222131-3012202202123232-0312103310010201-0032132013131023-0001222211300312): complete subsection reference.

<a id="canonical-3021231120131001-1310320220111012-2023102213231133-1100132023023123-1332132102033100-1212122232202300-1123111031233320-3332030310102010"></a>

<a id="canonical-2010301201110232-3113120110021303-0302321233223032-0101020230212321-1122200111012032-1122300121102333-1103321331000102-2210310213003031"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.invert_matcher` property

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-024.md#canonical-3221303202213312-3211202232332112-3101121302023021-3001021033200203-3301011202001122-1221212210100122-1300213130231002-3330311103023100): complete subsection reference.

<a id="canonical-2213321300002311-2112331211233132-3210113330113321-0101011031122232-3022013312103303-3103031011021011-2121230131303312-3223130213312202"></a>

<a id="canonical-3210231322322222-1003000110123001-2301202033312322-2120033133231212-3230231210230122-1123313001230332-0113220220211021-2001311023131023"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.key` property

Type: `"string"`. Optional.

A case-sensitive HTTP query parameter name.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3103300010123120-3200220011010132-3033202200030111-2303133000000131-0311023323212320-0013122300031122-0222001002323220-2332111330223102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-024.md#canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303)
- policy_based_challenge.rule_list.rules.spec.query_params.check_not_present

<a id="canonical-1101303100132312-0332013330032313-0322313012033222-1203311103123331-3301300321311102-1310132011213220-3200210002113130-1333201120100332"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112003030120300-3231323312231200-0322203333323223-3322103103222131-3012202202123232-0312103310010201-0032132013131023-0001222211300312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-024.md#canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303)
- policy_based_challenge.rule_list.rules.spec.query_params.check_present

<a id="canonical-2012232230312011-2000002102012122-2312132210110121-2112113203321232-0200231201101010-2323330333301321-2201301222133133-0023212023130201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221303202213312-3211202232332112-3101121302023021-3001021033200203-3301011202001122-1221212210100122-1300213130231002-3330311103023100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-024.md#canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303)
- policy_based_challenge.rule_list.rules.spec.query_params.item

<a id="canonical-2021111131031122-1233033121101103-0232020210133021-3231312303321331-0332101000121310-3203301010032330-3000011302111023-0310102312313100"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031103030220132-2331300003310232-1312013111133223-3202032033210111-0332212132212101-0031101013312301-0103000011203012-2310311001030032"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.query_params.item`

<a id="canonical-2000211123332030-2312331230233130-0231333310131100-0233231231311302-3300203021110132-0232132211223001-0021331100330232-0230102120110212"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2231211210223302-3221220321110013-2222131131220002-0002300123031321-2230211202033333-0012031032022033-3013221012101301-3220312010312122"></a>

<a id="canonical-1131201012300113-0133232213202002-0032112122133311-3121111212323221-1300310103303020-2221003202333120-1010132011321113-1100223233232103"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1130330003220001-2130120332033120-0302103320323222-2323001321301111-3211023222033022-2030212331202010-2222121312012321-0133233211000333"></a>

<a id="canonical-2101011333123332-2023222030321302-3332230321032220-3202211212221111-1333222111030301-0302323001102300-2002021202031201-1310111211201011"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1131033020202020-3130333330012030-3002001123121030-1323223321202032-1221131132031001-3323021113022011-1203131303033130-3010031021310313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-023.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-023.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-023.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher

<a id="canonical-0023121113310101-3000100110211120-0321313013123100-2223023123312331-0021220321232200-0320221203123231-3231231132101211-1112110300100212"></a>

Type: `"object"`. single nested block, Optional.

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

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
tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013031323323213-3313121012212012-0112303321001030-1100023031023311-0210001200203111-3132002132022111-2010033301222100-0102210102133120"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher`

<a id="canonical-1021111321203203-3301132133030002-2010322033311100-1303332302003321-3213112020223121-0220221202133210-0303002333032202-1302001231301203"></a>

#### `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher.classes` property

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Additional upstream details:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2200130213322322-1202121203133021-2311320332102331-2320112203331011-1132223232220312-1301022313313100-0322311023001212-1332110222103030"></a>

<a id="canonical-2010021303323201-3021331310213001-0121022023311100-3000002131303021-3213113020311211-3210100120233332-1131202302322303-1301020103220021"></a>

#### `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1201130212030203-2312220320013322-0320311001231222-2232020111320120-1110313223210301-3210113301312033-2121201200222023-0112222120120130"></a>

<a id="canonical-0112323210101301-1331121030333023-3033301212222231-2103032032321022-0110210132003200-2301130123200220-0111301112033112-2133233202130310"></a>

#### `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Optional.

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0332322110110101-3123223013302023-0321200302320201-2312111021001220-0032301331120201-2310003212230313-2113130213200312-2022222232311001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.temporary_user_blocking` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.temporary_user_blocking

<a id="canonical-1213330020020110-0201003231121130-0322032122200333-1222232301221110-1231122130323233-0213300101032113-3211212012313013-1313011303000120"></a>

Type: `"object"`. single nested block, Optional.

Specifies configuration for temporary user blocking resulting from user behavior analysis.

When Malicious User Mitigation is enabled from service policy rules, users' accessing the
application will be analyzed for malicious activity and the configured mitigation actions will be
taken on identified malicious users. These mitigation actions include setting up temporary blocking
on that user. This configuration specifies settings on how that blocking should be done by the
loadbalancer.

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
temporary_user_blocking {
  # Configure direct properties listed below.
}
```

<a id="canonical-1001233223103102-2321002300333321-1010002332023330-2130131120101022-3320211110300131-0000103020320331-1111331221322023-3112130013330213"></a>

### Direct properties for `policy_based_challenge.temporary_user_blocking`

<a id="canonical-3120200303022222-1320212320330112-2211020312113103-1012301200100332-0131032211233132-2100000001010121-3122131220320112-0210110121030023"></a>

#### `policy_based_challenge.temporary_user_blocking.custom_page` property

Type: `"string"`. Optional.

Custom message is of type . Currently supported URL schemes is . For scheme, message needs to be
encoded in base64 format. You can specify this message as base64 encoded plain text message e.g.
'Blocked.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '&lt;p&gt;
Blocked..

Additional upstream details:

Custom message is of type \`uri\_ref\`. Currently supported URL schemes is \`string:///\`. For
\`string:///\` scheme, message needs to be encoded in base64 format. You can specify this message as
base64 encoded plain text message e.g. "Blocked.." or it can be HTML paragraph or a body string
encoded as base64 string E.g. "&lt;p&gt; Blocked &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- protected_cookies

<a id="canonical-0012220312310111-2303132013122322-1110302013133310-3121322300132331-3211323303223113-1121220330010001-3331130300302213-1211121030222300"></a>

Type: `"object"`. list nested block, Optional.

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF
(monitoring or blocking) will be enforced on the request when cookie tampering is identified. Note:
We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
protected_cookies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313133323322130-1111101311101100-1311201113313021-2211332110031211-0312113310011223-1213211301102330-0220030111133311-1002023200222010"></a>

### Direct properties for `protected_cookies`

- [add_httponly](resources--http_loadbalancer--reference--group-024.md#canonical-3300300313122023-0131101133121122-0111302321111221-1320231013003202-0133111311232132-2000031230020032-2303313311132310-0320021312132211): complete subsection reference.

- [add_secure](resources--http_loadbalancer--reference--group-024.md#canonical-2231033302001101-0132231313110213-2031112210223222-3333133100110021-1233300112010200-2023001101013012-2002112312201120-1102200221332013): complete subsection reference.

- [disable_tampering_protection](resources--http_loadbalancer--reference--group-024.md#canonical-3213201012220223-2211332120011102-3023133313022103-0320131213311121-2310000233223030-2021102301111110-1000311211111311-1220303110323223): complete subsection reference.

- [enable_tampering_protection](resources--http_loadbalancer--reference--group-024.md#canonical-2120321222300232-3201213120231321-3002320121130331-3033120022323313-2001001201032333-3213123201002332-1312021122123131-1103100231332211): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--reference--group-024.md#canonical-0203212110203021-2200220231120210-2230311113031102-3221022331201212-2022331233333222-1023330101000213-2212231131222231-2200000112100101): complete subsection reference.

- [ignore_max_age](resources--http_loadbalancer--reference--group-024.md#canonical-0113232103232213-1220110302131233-2220201300012301-1223330101120121-3102023230122321-2021001312001331-3212303132300032-1121003002111111): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--reference--group-024.md#canonical-1222213222130010-1030032131320230-1313212131202032-0230100032002121-2200132212221021-2032011033232331-1222130131101301-3033111203030322): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--reference--group-024.md#canonical-2003122021101102-2103020000110303-1022023202020222-2212321100200322-1113103131312223-1311113032323233-2222101320213033-3231231121303232): complete subsection reference.

<a id="canonical-1323120312300302-3320100011310320-0011202033330112-3301121130120012-0301301111231001-0102330220212010-3133000111121113-1020102133110003"></a>

<a id="canonical-3320233203100101-0123331113233032-2013002133222330-3323310301303231-1011212011011202-0333210110322202-3220212211231022-3013220001211032"></a>

#### `protected_cookies.max_age_value` property

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-1321133322311221-1000131311233203-0132313321131303-0222222031301101-0233031223010201-0331020201130000-1203312001333210-3333222232301112"></a>

<a id="canonical-1023311022121202-2100322213130030-0000113020103113-3123001221123122-3013211203111212-3212211322322301-0023203201122112-1011021231131202"></a>

#### `protected_cookies.name` property

Type: `"string"`. Optional.

Cookie Name. Name of the Cookie.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [samesite_lax](resources--http_loadbalancer--reference--group-024.md#canonical-2000130001101232-3100031031201010-3211203321113321-1000321203033310-2232002312022102-1223200123130111-2321322310212321-2000102201320130): complete subsection reference.

- [samesite_none](resources--http_loadbalancer--reference--group-024.md#canonical-1010202312223010-3101022001131300-0120102210132103-1310330112122232-3020103212011201-0132121012321221-2320211000122301-2023300212031311): complete subsection reference.

- [samesite_strict](resources--http_loadbalancer--reference--group-024.md#canonical-3202002233301113-3300122322310232-3222101130120113-3130000312231310-2013103001102003-0122111300211202-0030121200112032-3020100220132200): complete subsection reference.

<a id="canonical-3300300313122023-0131101133121122-0111302321111221-1320231013003202-0133111311232132-2000031230020032-2303313311132310-0320021312132211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.add_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-024.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.add_httponly

<a id="canonical-3211233200132230-1320201203133210-3222121231102200-3101310200102331-1131033211103203-2210123131100133-0230202031023310-2213013200001220"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add httponly.

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
add_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231033302001101-0132231313110213-2031112210223222-3333133100110021-1233300112010200-2023001101013012-2002112312201120-1102200221332013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.add_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-024.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.add_secure

<a id="canonical-0020002302213013-0321120233322011-2303112122210301-0220003000230021-0212331113021030-2020021121211131-3020302100222213-3111320232030233"></a>

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
add_secure = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213201012220223-2211332120011102-3023133313022103-0320131213311121-2310000233223030-2021102301111110-1000311211111311-1220303110323223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.disable_tampering_protection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-024.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.disable_tampering_protection

<a id="canonical-1311001213222122-0311202222022223-0120102331233220-3130333230233133-1110013230032203-3223310200332130-2130333212112111-3111321302113321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable tampering protection.

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
disable_tampering_protection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120321222300232-3201213120231321-3002320121130331-3033120022323313-2001001201032333-3213123201002332-1312021122123131-1103100231332211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.enable_tampering_protection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-024.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.enable_tampering_protection

<a id="canonical-0002202213011112-0311010233013220-1302123011033320-0012132113121123-2210221023223110-2311311303013232-2010300133122211-3032030323010300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable tampering protection.

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
enable_tampering_protection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203212110203021-2200220231120210-2230311113031102-3221022331201212-2022331233333222-1023330101000213-2212231131222231-2200000112100101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.ignore_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-024.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.ignore_httponly

<a id="canonical-2021211012122313-0031031302100232-0001013320130133-0131033130312033-3231021213132201-1230300332232332-0233220101310323-3103110023013223"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113232103232213-1220110302131233-2220201300012301-1223330101120121-3102023230122321-2021001312001331-3212303132300032-1121003002111111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.ignore_max_age` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-024.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.ignore_max_age

<a id="canonical-3102101331001031-1020221120320220-3023120133211031-3232311303122110-2220012122230101-1131002101220123-1201320113013011-3003112300031121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore max age.

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
ignore_max_age = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222213222130010-1030032131320230-1313212131202032-0230100032002121-2200132212221021-2032011033232331-1222130131101301-3033111203030322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.ignore_samesite` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-024.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.ignore_samesite

<a id="canonical-1003303130031103-3220313222211303-1313122113333312-0130130022102001-0312331133011002-2011220003023222-2033310012230110-3132322330103132"></a>

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
ignore_samesite = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003122021101102-2103020000110303-1022023202020222-2212321100200322-1113103131312223-1311113032323233-2222101320213033-3231231121303232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.ignore_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-024.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.ignore_secure

<a id="canonical-3021010300301100-3321220031113313-3033233222300202-0223123021330123-1120113120322110-0132203013000112-1011020203231313-0203311310020133"></a>

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
ignore_secure = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000130001101232-3100031031201010-3211203321113321-1000321203033310-2232002312022102-1223200123130111-2321322310212321-2000102201320130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.samesite_lax` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-024.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.samesite_lax

<a id="canonical-2231000213132033-1031012230003310-1001222221110022-3000311232220212-2310010100220213-3022011033303122-3020122001101120-0322321103313103"></a>

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
samesite_lax = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010202312223010-3101022001131300-0120102210132103-1310330112122232-3020103212011201-0132121012321221-2320211000122301-2023300212031311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.samesite_none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-024.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.samesite_none

<a id="canonical-3012022311323223-2331300201331233-0103310031123132-0031221331122013-0203200322000332-2033113110210131-1111332221232013-1130113033101123"></a>

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
samesite_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202002233301113-3300122322310232-3222101130120113-3130000312231310-2013103001102003-0122111300211202-0030121200112032-3020100220132200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.samesite_strict` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-024.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.samesite_strict

<a id="canonical-2121330213211000-2200123230023111-1120021233023311-2222101302323212-2122012322123302-0010323322232231-1011201201130112-0332321022200302"></a>

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
samesite_strict = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312120313213323-2103003010102231-1132113202300010-0103032212233002-0220031323220302-3231002030321130-0010020301002111-1312011221320111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `random` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- random

<a id="canonical-1021313003120223-3101211023223000-1333330100323102-3033202113022313-2111113220003123-1112002011233323-1312030112001012-1111020111100321"></a>

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
random = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- rate_limit

<a id="canonical-1212231232113102-0000210123322230-2321331212231011-1223210321030302-2020032100331231-0213102112123132-3222021323133333-1010333102020321"></a>

Type: `"object"`. single nested block, Optional.

Load-balancer-wide per-client rate limiting. The counter applies across every path; use
api\_rate\_limit rules when only selected paths such as /login should be limited.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]",
  "x-ves-oneof-field-policy_choice": "[\"no_policies\",\"policies\"]"
}
```

Terraform syntax:

```terraform
rate_limit {
  # Configure direct properties listed below.
}
```

<a id="canonical-2321223131033012-2032331230200132-1302111022313203-1000133020113002-2231301330103221-0222223212320201-2322012230101001-3002000100302201"></a>

### Direct properties for `rate_limit`

- [custom_ip_allowed_list](resources--http_loadbalancer--reference--group-024.md#canonical-3111012121110331-0220030113112232-0310311011223312-3011311300331132-3031132310013331-2312333030201102-2032320032220102-3023013332301231): complete subsection reference.

- [ip_allowed_list](resources--http_loadbalancer--reference--group-024.md#canonical-0232210100232323-3030202021133132-1202300110321011-0302032013001321-3232023111021033-1321311022302323-0310203233111223-3020100011333002): complete subsection reference.

- [no_ip_allowed_list](resources--http_loadbalancer--reference--group-024.md#canonical-3010321303322033-1212332122021131-3013312103113211-2131102032023233-3123302121323033-2320112001033113-3312022230012033-3102230311220002): complete subsection reference.

- [no_policies](resources--http_loadbalancer--reference--group-024.md#canonical-1012222010200213-1013130113011130-1203230321210101-0313310220312312-0311230001310213-1113313130100101-2020020003003112-3312012120123223): complete subsection reference.

- [policies](resources--http_loadbalancer--reference--group-024.md#canonical-1201121223001130-0102301012320112-0332103302021323-3120320310211121-2331102322201132-0130330012330030-3122013001202010-1320120013120101): complete subsection reference.

- [rate_limiter](resources--http_loadbalancer--reference--group-024.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031): complete subsection reference.

<a id="canonical-3111012121110331-0220030113112232-0310311011223312-3011311300331132-3031132310013331-2312333030201102-2032320032220102-3023013332301231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.custom_ip_allowed_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- rate_limit.custom_ip_allowed_list

<a id="canonical-3023030301000222-2020202210131000-2230300203230310-2010122011332100-0212133223233100-0322001102312102-3132333221333020-0012303220013220"></a>

Type: `"object"`. single nested block, Optional.

IP Allowed list using existing ip\_prefix\_set objects.

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
custom_ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132300111322010-0211021003112333-1101000211023210-3133030310113020-3232110100200221-2213032332133220-0010223220111031-2313233311030332"></a>

### Direct properties for `rate_limit.custom_ip_allowed_list`

- [rate_limiter_allowed_prefixes](resources--http_loadbalancer--reference--group-024.md#canonical-0220012023113211-1122013121013033-3230313310101110-1331120330133301-1033312101103121-0131131033321010-0231222121022201-2313210322200033): complete subsection reference.

<a id="canonical-0220012023113211-1122013121013033-3230313310101110-1331120330133301-1033312101103121-0131131033321010-0231222121022201-2313210322200033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.custom_ip_allowed_list](resources--http_loadbalancer--reference--group-024.md#canonical-3111012121110331-0220030113112232-0310311011223312-3011311300331132-3031132310013331-2312333030201102-2032320032220102-3023013332301231)
- rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes

<a id="canonical-2000123031021033-3223132012102132-3231013132332332-0130332221003321-1122312113200300-0023001323113301-0133122032332333-1122230203103312"></a>

Type: `"object"`. list nested block, Optional.

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
rate_limiter_allowed_prefixes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3133301330101023-0100300220130200-0200013020101221-1300013131303220-1131311221223023-1300210323030321-2133212222122201-0330012120013303"></a>

### Direct properties for `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes`

<a id="canonical-0130033032121322-2110201010312330-2031012013112332-3203310322133323-2320212210211101-3302112112033131-2331123100110300-3323113302320110"></a>

#### `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3021320010033301-1203300230030301-3310320321130000-1232331113100101-3321311201212001-1232202002000323-2323223222211221-1321201300211120"></a>

<a id="canonical-3101023303220002-3212131000301113-0001023033100322-0131310233111010-0303322323102012-3222213132033323-2002212031330211-2232030222110013"></a>

#### `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1230103032221011-1112302112121002-1032213222300133-3121303132302231-0230313112322331-3333332021221320-2110331131320322-2030203223301020"></a>

<a id="canonical-3230021331331230-3331012101002212-2011100000223331-0031131132300111-2202221030103221-0100003033033301-0213203231201331-0203011011321320"></a>

#### `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0232210100232323-3030202021133132-1202300110321011-0302032013001321-3232023111021033-1321311022302323-0310203233111223-3020100011333002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.ip_allowed_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- rate_limit.ip_allowed_list

<a id="canonical-3103033003311002-1200130211131302-2213201303103331-3031320123113330-2021311303122332-1231311010022303-3230112212230320-1002133121102033"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232211132101301-1030100301020331-3233313300013210-2213300100120033-0211101203311312-1021222200303320-1031132103313101-3113023301212200"></a>

### Direct properties for `rate_limit.ip_allowed_list`

<a id="canonical-2131112302013213-0233210323130101-3310001232032331-3003322021222112-1333002311223101-0130300103011221-2200312331110100-2132302231213031"></a>

#### `rate_limit.ip_allowed_list.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3010321303322033-1212332122021131-3013312103113211-2131102032023233-3123302121323033-2320112001033113-3312022230012033-3102230311220002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.no_ip_allowed_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- rate_limit.no_ip_allowed_list

<a id="canonical-0122110212212020-1100201100100210-0122123233122013-0113231212102312-3022132010230013-2210233332010021-2121231233211120-2201203212303220"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
no_ip_allowed_list = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012222010200213-1013130113011130-1203230321210101-0313310220312312-0311230001310213-1113313130100101-2020020003003112-3312012120123223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.no_policies` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- rate_limit.no_policies

<a id="canonical-2311111332132033-3203120133203120-3223003203201302-0303213122001132-0320202211030100-0110123003333110-1313212021323313-2201333130013200"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no policies. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
no_policies = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201121223001130-0102301012320112-0332103302021323-3120320310211121-2331102322201132-0130330012330030-3122013001202010-1320120013120101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.policies` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- rate_limit.policies

<a id="canonical-1311212123020300-1202120201123332-1120101331132123-1310031310211112-3323003212013030-3102221103030333-1321333122323331-2122022321132103"></a>

Type: `"object"`. single nested block, Optional.

List of rate limiter policies to be applied.

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
policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2101003312111121-0030323102020111-2223210103110010-2032320122322332-1013021110103233-1030221133320233-1001220333031121-0221222313111012"></a>

### Direct properties for `rate_limit.policies`

- [policies](resources--http_loadbalancer--reference--group-024.md#canonical-3101001310133003-3001131113313331-0130031131012332-3322332320111210-0220323211221110-0202113030011302-0120211202101330-2132032321310302): complete subsection reference.

<a id="canonical-3101001310133003-3001131113313331-0130031131012332-3322332320111210-0220323211221110-0202113030011302-0120211202101330-2132032321310302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.policies.policies` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.policies](resources--http_loadbalancer--reference--group-024.md#canonical-1201121223001130-0102301012320112-0332103302021323-3120320310211121-2331102322201132-0130330012330030-3122013001202010-1320120013120101)
- rate_limit.policies.policies

<a id="canonical-1132311333313033-3331110210011121-0223123000200333-3131030320223011-0133301132213021-2311122120312201-2033301003302032-2133001010133022"></a>

Type: `"object"`. list nested block, Optional.

Rate Limiter Policies. Ordered list of rate limiter policies.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303122020201103-0123320031123230-3302222022311100-3213131002331312-1322223033101313-1120030203222033-0232120120023100-2232123221113032"></a>

### Direct properties for `rate_limit.policies.policies`

<a id="canonical-1013132032310330-2010231113122122-1201110122301100-1232130121033210-3122233031300013-1033112103031230-0320233120203001-1211002012202131"></a>

#### `rate_limit.policies.policies.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1033103333211331-3112201311023132-0210310112011103-2030202100012021-1023201033300231-3203213012210211-3230211221302121-3013023022131300"></a>

<a id="canonical-3130002201210121-0201200011033220-0100203302213103-1211303202001300-1220211122303201-1312313330311211-1123030110000203-0220332233033331"></a>

#### `rate_limit.policies.policies.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0001232312122333-0010233211302322-1132102320310121-3203011231120012-1003132111011220-0211112001302103-3020030111031212-1113020211033331"></a>

<a id="canonical-0123321323002021-2010033221301223-0323213311132313-1011122102201230-3311031020233322-2012201112123130-2313023230323301-1233020233000001"></a>

#### `rate_limit.policies.policies.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- rate_limit.rate_limiter

<a id="canonical-1103110010233130-2233133213021322-0202121221220302-2233033212231103-1112302201122210-0211200201021331-3210320212303013-3222333310221232"></a>

Type: `"object"`. single nested block, Optional.

A tuple consisting of a rate limit period unit and the total number of allowed requests for that
period.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"action_block\",\"disabled\"]",
  "x-ves-oneof-field-algorithm": "[\"leaky_bucket\",\"token_bucket\"]"
}
```

Terraform syntax:

```terraform
rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122333212200233-1201111333110123-3133302130013032-0101303131130320-0301221113333230-1311113033020000-3030323202101020-3223211020313113"></a>

### Direct properties for `rate_limit.rate_limiter`

- [action_block](resources--http_loadbalancer--reference--group-024.md#canonical-3103202121112302-1221322032222100-1212030133113120-2323120022200233-3321232113311332-1033310032011022-1300011312223023-0333001032310333): complete subsection reference.

<a id="canonical-2000021212311200-1321200012003130-1102300000323302-3313133120001200-0103013203021230-2320202103212221-0110200223122210-2232132310121300"></a>

<a id="canonical-0232210112202310-1211203313030203-1303310232231110-1203311111301313-1233320100222020-2202233332220311-2013112201103020-3020230322202031"></a>

#### `rate_limit.rate_limiter.burst_multiplier` property

Type: `"number"`. Optional.

The maximum burst of requests to accommodate, expressed as a multiple of the rate.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [disabled](resources--http_loadbalancer--reference--group-024.md#canonical-3232033201121333-3023321102300202-0203022310220000-3131212232330322-3113303131212013-0200123021011021-1233230122113213-3103023100232011): complete subsection reference.

- [leaky_bucket](resources--http_loadbalancer--reference--group-024.md#canonical-0123123131000310-3001121302331023-2333103110313033-2231001002030010-1032020221100330-0020022233032020-3333212003233230-1203231010321320): complete subsection reference.

<a id="canonical-0210223003020100-1021313110031303-3123303211222220-1211332132003220-3113111030021210-0120223103210311-2201002003013033-1033232302023021"></a>

<a id="canonical-0111222112300022-1000301010022023-3200021102231302-3211100213222303-3000031001130022-3202203113011111-0132031032031322-1133033011111121"></a>

#### `rate_limit.rate_limiter.period_multiplier` property

Type: `"number"`. Optional, Computed.

Setting, combined with Per Period units, provides a duration. Server applies default when omitted.

Additional upstream details:

This setting, combined with Per Period units, provides a duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.uint32.gte": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  }
}
```

- [token_bucket](resources--http_loadbalancer--reference--group-024.md#canonical-0303221200330202-3123321021000110-2332300002133112-2220222003033013-2013103231112210-3132200122112030-0221300021323023-2013233123003032): complete subsection reference.

<a id="canonical-0112031132301223-3130323311131010-3122232123200131-3213003010331131-3213220202011223-3130210231121332-3112123001103201-1113232031202102"></a>

<a id="canonical-2333103312100110-1130302212221221-2033332132023121-3121313102103211-1202203103200113-0121223023210132-2230303000013012-1111213020232133"></a>

#### `rate_limit.rate_limiter.total_number` property

Type: `"number"`. Optional.

The total number of allowed requests per rate-limiting period.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="canonical-2132010130102001-2232231331221223-0021200330212032-2001320020100323-2312113233211103-0000013320231002-0322302012021003-2022022102011331"></a>

<a id="canonical-3231120210311213-0023212231011322-0323313123123313-0103321010221203-1001010131110332-0021211233222330-1321133101102021-3211212313231201"></a>

#### `rate_limit.rate_limiter.unit` property

Type: `"string"`. Optional.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3103202121112302-1221322032222100-1212030133113120-2323120022200233-3321232113311332-1033310032011022-1300011312223023-0333001032310333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-024.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- rate_limit.rate_limiter.action_block

<a id="canonical-0213300333112130-3323221222133220-2311120201200003-1210111132202002-2122332221303211-0221233220200102-1322213132231300-2010310331100321"></a>

Type: `"object"`. single nested block, Optional.

Action where a user is blocked from making further requests after exceeding rate limit threshold.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-block_duration_choice": "[\"hours\",\"minutes\",\"seconds\"]"
}
```

Terraform syntax:

```terraform
action_block {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323233230232103-1013000311321002-0220201022212222-1003231331133232-1123202201222102-1022203001032212-2311322203013221-3003300222333331"></a>

### Direct properties for `rate_limit.rate_limiter.action_block`

- [hours](resources--http_loadbalancer--reference--group-024.md#canonical-2121310033333031-0103220231231331-1032030013313221-3122010023000002-1312012020333332-3210031232232311-3022101233033212-2213012122220320): complete subsection reference.

- [minutes](resources--http_loadbalancer--reference--group-024.md#canonical-1123220322033011-2022131300031302-0331123002023223-0013333220111023-2103032022232203-1331212310031113-3222233201131302-1132332303222222): complete subsection reference.

- [seconds](resources--http_loadbalancer--reference--group-024.md#canonical-1100321133312232-1010213301300133-1221233020303230-0320102001333003-3133303321011103-3133132232103001-0303113322032210-0221022013233021): complete subsection reference.

<a id="canonical-2121310033333031-0103220231231331-1032030013313221-3122010023000002-1312012020333332-3210031232232311-3022101233033212-2213012122220320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block.hours` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-024.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- [rate_limit.rate_limiter.action_block](resources--http_loadbalancer--reference--group-024.md#canonical-3103202121112302-1221322032222100-1212030133113120-2323120022200233-3321232113311332-1033310032011022-1300011312223023-0333001032310333)
- rate_limit.rate_limiter.action_block.hours

<a id="canonical-1212121023112330-3230313123332212-2033332012333202-2200111302031332-3232122222121330-2200323100032002-1231233131220101-1223221211301331"></a>

Type: `"object"`. single nested block, Optional.

Hours. Input Duration Hours.

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
hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103213001130112-3002133223311123-0101200303322210-2233002313030101-1101220010120323-0010223202011113-0030003113303110-1030321120210033"></a>

### Direct properties for `rate_limit.rate_limiter.action_block.hours`

<a id="canonical-0133313110003300-2131003100120020-1310103120111211-0200010220002100-0231120322013233-0131300200211130-0311211312323100-1001023233313333"></a>

#### `rate_limit.rate_limiter.action_block.hours.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 48,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  }
}
```

<a id="canonical-1123220322033011-2022131300031302-0331123002023223-0013333220111023-2103032022232203-1331212310031113-3222233201131302-1132332303222222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block.minutes` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-024.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- [rate_limit.rate_limiter.action_block](resources--http_loadbalancer--reference--group-024.md#canonical-3103202121112302-1221322032222100-1212030133113120-2323120022200233-3321232113311332-1033310032011022-1300011312223023-0333001032310333)
- rate_limit.rate_limiter.action_block.minutes

<a id="canonical-2203203013223101-2200113021201011-1100330032330331-0103212031203031-2232031111002220-1110121103321222-1131312301030131-3313132313113021"></a>

Type: `"object"`. single nested block, Optional.

Minutes. Input Duration Minutes.

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
minutes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303220032301032-3333133000213011-0102220231223113-3031331312101321-0331112202012123-1012202313320133-2320333103331331-0130332022132211"></a>

### Direct properties for `rate_limit.rate_limiter.action_block.minutes`

<a id="canonical-0101133120101020-1322213102101030-2303033022031013-3102021121101130-2230330302113210-3301003023203132-0230111310213212-2310230212012133"></a>

#### `rate_limit.rate_limiter.action_block.minutes.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  }
}
```

<a id="canonical-1100321133312232-1010213301300133-1221233020303230-0320102001333003-3133303321011103-3133132232103001-0303113322032210-0221022013233021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block.seconds` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-024.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- [rate_limit.rate_limiter.action_block](resources--http_loadbalancer--reference--group-024.md#canonical-3103202121112302-1221322032222100-1212030133113120-2323120022200233-3321232113311332-1033310032011022-1300011312223023-0333001032310333)
- rate_limit.rate_limiter.action_block.seconds

<a id="canonical-2102331331230211-3312013003332002-0332203201212003-2310331130213023-1121022213321120-3223001011213330-0331312030203110-3032321301032322"></a>

Type: `"object"`. single nested block, Optional.

Seconds. Input Duration Seconds.

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
seconds {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102012013100311-2103313203303011-0212002033022233-0033203223320022-3000121332032202-0101101000132100-0230210300311131-1312222133010011"></a>

### Direct properties for `rate_limit.rate_limiter.action_block.seconds`

<a id="canonical-2033332002002300-2312210033323220-1311320033133103-2012022131322322-3303032331103121-2230301133121322-2312302012033331-0310030120210221"></a>

#### `rate_limit.rate_limiter.action_block.seconds.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-3232033201121333-3023321102300202-0203022310220000-3131212232330322-3113303131212013-0200123021011021-1233230122113213-3103023100232011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-024.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- rate_limit.rate_limiter.disabled

<a id="canonical-0010013332033122-0111110320233332-2112113320311311-1230012102202313-0201013312020323-3030120221312133-3120300303002213-2113132201032112"></a>

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
disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123123131000310-3001121302331023-2333103110313033-2231001002030010-1032020221100330-0020022233032020-3333212003233230-1203231010321320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.leaky_bucket` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-024.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- rate_limit.rate_limiter.leaky_bucket

<a id="canonical-2131030311121003-1030133220332212-1011212013002131-3231113221020320-2321223231122300-3301330201310330-0202201110222010-3123222021301022"></a>

Type: `["object", {}]`. Optional.

Leaky-Bucket is the default rate limiter algorithm for F5.

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
leaky_bucket = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303221200330202-3123321021000110-2332300002133112-2220222003033013-2013103231112210-3132200122112030-0221300021323023-2013233123003032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.token_bucket` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-024.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- rate_limit.rate_limiter.token_bucket

<a id="canonical-0100203310120213-3200101121020211-2202323212003001-1011323011030332-1100301103012221-0112021202023321-3112223113023112-1111200023011110"></a>

Type: `["object", {}]`. Optional.

Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.

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
token_bucket = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- ring_hash

<a id="canonical-3013210122112121-0122221322331320-0013232112123331-2012132320331222-1132231302310032-2101233212003331-1003333001210203-3312200031323332"></a>

Type: `"object"`. single nested block, Optional.

Hash Policy List. List of hash policy rules.

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
ring_hash {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302212102131332-2013333122123320-0332022333133022-3311100202303201-1001113302300230-2322333313321133-0233300311131302-1003021301303210"></a>

### Direct properties for `ring_hash`

- [hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313): complete subsection reference.

<a id="canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- ring_hash.hash_policy

<a id="canonical-1001233103230101-1000300000232320-2033102230313003-1310230311210113-1012230230203331-3223231202203232-0011132203210031-0130133230011322"></a>

Type: `"object"`. list nested block, Optional.

Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated
individually and the combined result is used to route the request.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
hash_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223223003010302-0230300113113120-0011000323022103-3213302113210330-1313302022021321-0211203100323110-0233032313220301-3100022003013133"></a>

### Direct properties for `ring_hash.hash_policy`

- [cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022): complete subsection reference.

<a id="canonical-3130221302202133-2230001003300020-0010021133302231-1232033202332221-1010101011032233-3013000312011331-0301123102131101-0221313130020111"></a>

<a id="canonical-0113203213122220-0010200033110212-1010212120100230-3221032333233012-0000000200202220-0210303302221111-0010332231203232-3012112021002231"></a>

#### `ring_hash.hash_policy.header_name` property

Type: `"string"`. Optional.

Exclusive with \[cookie source\_ip\] The name or key of the request header that will be used to
obtain the hash key.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3022230202300222-1103201130121133-2020232311113203-0010032322013233-2231112201012302-2020302313220222-3122111300011000-0033301111001310"></a>

<a id="canonical-3213233031012232-2102220323312001-3212332120223031-2200323312131331-1303101131100323-1323013020012100-1132222132322213-1213130211032333"></a>

#### `ring_hash.hash_policy.source_ip` property

Type: `"bool"`. Optional.

Exclusive with \[cookie header\_name\] Hash based on source IP address.

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

<a id="canonical-3100100313301221-2003203103011131-0210322222110321-3310323303021000-0013331320110102-3021311233100030-3213102021100013-1211321323300001"></a>

<a id="canonical-0112010311012000-2302310021211301-0232012103313010-0322213033302202-0112321212010002-1122022202030023-3300101230032012-1000301013130213"></a>

#### `ring_hash.hash_policy.terminal` property

Type: `"bool"`. Optional.

Terminal. Specify if its a terminal policy.

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

<a id="canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- ring_hash.hash_policy.cookie

<a id="canonical-1123221300031301-0233102300131322-0231223220233120-3303232200331313-0212311210210200-3311321223022320-0020111333303303-1231013112310213"></a>

Type: `"object"`. single nested block, Optional.

Two types of cookie affinity:

&#8203;1. Passive. Takes a cookie that's present in the cookies header and hashes on its value.

&#8203;2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from
the client in its response to the client, based on the endpoint the request gets sent to. The client
then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure
these requests GET sent to the same endpoint. The cookie is generated by hashing the source and
destination ports and addresses so that multiple independent HTTP2 streams on the same connection
will independently receive the same cookie, even if they arrive simultaneously.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-httponly": "[\"add_httponly\",\"ignore_httponly\"]",
  "x-ves-oneof-field-samesite": "[\"ignore_samesite\",\"samesite_lax\",\"samesite_none\",\"samesite_strict\"]",
  "x-ves-oneof-field-secure": "[\"add_secure\",\"ignore_secure\"]"
}
```

Terraform syntax:

```terraform
cookie {
  # Configure direct properties listed below.
}
```

<a id="canonical-1101123332113122-2333210302230022-0321130102231013-2103030331101000-1121113031002210-3213031120132201-2301202031223233-3302011331031201"></a>

### Direct properties for `ring_hash.hash_policy.cookie`

- [add_httponly](resources--http_loadbalancer--reference--group-024.md#canonical-0233023321022032-3211013303101231-2222233200031322-2133030222102003-2302321230212320-0333010103001021-0120230300323003-3022321112010031): complete subsection reference.

- [add_secure](resources--http_loadbalancer--reference--group-024.md#canonical-3211223231000113-0222213231133230-3212131130101231-1001311011131331-0133032202322231-2032323232301021-0111130223012020-2300121033030212): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--reference--group-024.md#canonical-3300213232032231-2313023333310112-1121011103222021-0302330321101021-2013003213013331-0213002111110013-3101131111310201-0200002112231203): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--reference--group-024.md#canonical-2332333230022233-0132022121301303-0103202022312020-2313313001110002-0010110100130212-2310031310302213-2102022001003100-3103303321201213): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--reference--group-024.md#canonical-1120222021213231-3003220120213132-1322210223123032-2300111330030321-2133103202003233-0211133213032211-1020320331332121-2112122222011100): complete subsection reference.

<a id="canonical-2112003132221332-1102011332030131-3301333332120300-3001301302323033-1211010120332021-2323323021010222-3023033010103221-3222112232231223"></a>

<a id="canonical-1313000002211333-3032122211233312-2112223020120322-2100331102013131-3122231110012232-2013100230111220-2020013323333023-0012332030212322"></a>

#### `ring_hash.hash_policy.cookie.name` property

Type: `"string"`. Optional.

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0312330232332012-0031000233000213-1120331101023313-2302110201213031-1313012103322132-3201222312320331-3012323110133301-1333332032130123"></a>

<a id="canonical-1200322331310201-0213111302331000-1123030330221101-0310222012032113-2032033121022012-3212200011133020-1031111023211120-0033201021121201"></a>

#### `ring_hash.hash_policy.cookie.path` property

Type: `"string"`. Optional.

The name of the path for the cookie. If no path is specified here, no path will be set for the
cookie.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [samesite_lax](resources--http_loadbalancer--reference--group-024.md#canonical-1000320300010321-3113212121101033-2123003330212330-0323233012330112-0110012230030333-1210122230222230-2303323122231003-2120223222320230): complete subsection reference.

- [samesite_none](resources--http_loadbalancer--reference--group-024.md#canonical-1102322301312031-3201231211122112-2101022200020331-1021321333212213-3103022322333203-3313100032331102-3033023233111210-1130311231303313): complete subsection reference.

- [samesite_strict](resources--http_loadbalancer--reference--group-024.md#canonical-2332013231031311-3111321203331230-1210323300201203-0103111103210012-1133012210301011-1203213110013110-1012302001330233-3330321002132301): complete subsection reference.

<a id="canonical-1132130012331100-0310322232232230-2022212013030023-1232113320022300-0100303231132003-1322201301120203-1201022110222211-0023001001103103"></a>

<a id="canonical-0323003211220232-1330033011112133-2133121310131020-2102000331023031-3221021231132230-0023021202300300-0131300301323102-1102210220210221"></a>

#### `ring_hash.hash_policy.cookie.ttl` property

Type: `"number"`. Optional.

If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is
present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.

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

<a id="canonical-0233023321022032-3211013303101231-2222233200031322-2133030222102003-2302321230212320-0333010103001021-0120230300323003-3022321112010031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.add_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.add_httponly

<a id="canonical-3133310313000111-2031021211113022-2333121132001200-0002000320113033-1032321122023211-1203222021320021-3030013320002321-0023332000013022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add httponly.

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
add_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211223231000113-0222213231133230-3212131130101231-1001311011131331-0133032202322231-2032323232301021-0111130223012020-2300121033030212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.add_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.add_secure

<a id="canonical-2223202333002113-3121121013230322-3100033102002300-0303013132133101-2002130031223011-3223113000330122-2330300202320002-3302130321002233"></a>

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
add_secure = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300213232032231-2313023333310112-1121011103222021-0302330321101021-2013003213013331-0213002111110013-3101131111310201-0200002112231203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.ignore_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.ignore_httponly

<a id="canonical-0102320010131233-1133121101111033-2311002213233001-1103001232310101-0220101200020010-0220321110212200-3301121111230021-2323122121003102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332333230022233-0132022121301303-0103202022312020-2313313001110002-0010110100130212-2310031310302213-2102022001003100-3103303321201213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.ignore_samesite` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.ignore_samesite

<a id="canonical-3331020003223201-0012120213232031-1131223211002332-0112201012023132-0310021003100100-3113013012100203-0132032311221020-2130322123013200"></a>

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
ignore_samesite = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120222021213231-3003220120213132-1322210223123032-2300111330030321-2133103202003233-0211133213032211-1020320331332121-2112122222011100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.ignore_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.ignore_secure

<a id="canonical-0333113130311222-0210223021200112-1213131031223333-1020320212103020-2133133332201312-1032021330020310-3322100020311310-1223302032102331"></a>

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
ignore_secure = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000320300010321-3113212121101033-2123003330212330-0323233012330112-0110012230030333-1210122230222230-2303323122231003-2120223222320230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.samesite_lax` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.samesite_lax

<a id="canonical-3322132213203201-1231030122102312-3313323023333302-1103130202122202-2321000011313322-1023332010301112-0003012211311132-2021101230311332"></a>

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
samesite_lax = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1102322301312031-3201231211122112-2101022200020331-1021321333212213-3103022322333203-3313100032331102-3033023233111210-1130311231303313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.samesite_none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.samesite_none

<a id="canonical-3202023323212002-3032033213100112-1123312131202001-2131312331303001-0113230000100210-3001210211011302-0111113312301003-2322003301000330"></a>

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
samesite_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332013231031311-3111321203331230-1210323300201203-0103111103210012-1133012210301011-1203213110013110-1012302001330233-3330321002132301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.samesite_strict` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.samesite_strict

<a id="canonical-2221232320232210-3102000210122201-0013322231302022-1213121230230323-0210210232213131-3323121133012011-3022332100130213-3032122111300113"></a>

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
samesite_strict = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120201302000120-1101032023332332-1203301030103211-2301112131023100-3323333313021213-2102103120123013-3211232030113002-1310102202033321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `round_robin` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- round_robin

<a id="canonical-1120231012232110-3001211103003332-2110332213132110-2233201023123010-3331031002232323-2201313122330202-3321233220103133-3330032203222203"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for round robin. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
round_robin = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- routes

<a id="canonical-1022112302210201-1010011312111323-0033100111313112-0212233102202102-1021022230231123-1121103333323120-1231212033301110-1123111133332122"></a>

Type: `"object"`. list nested block, Optional.

Routes allow users to define match condition on a path and/or HTTP method to either forward matching
traffic to origin pool or redirect matching traffic to a different URL or respond directly to
matching traffic.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1302333330103332-3000231123012200-1223313032013033-3110312130233300-3201000202121113-1221011200232010-0313220022331010-2221323130111313"></a>

### Direct properties for `routes`

- [custom_route_object](resources--http_loadbalancer--reference--group-024.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130): complete subsection reference.

- [direct_response_route](resources--http_loadbalancer--reference--group-025.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311): complete subsection reference.

- [redirect_route](resources--http_loadbalancer--reference--group-025.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223): complete subsection reference.

- [route_state_disabled](resources--http_loadbalancer--reference--group-025.md#canonical-1103212330200010-1002200200101111-3203321323032332-2322002023022130-2011310020111101-3122113102023032-0332312121303202-3320010033212303): complete subsection reference.

- [route_state_enabled](resources--http_loadbalancer--reference--group-025.md#canonical-0233030011110210-1112111102103333-2313130301030100-2022033113121211-0031211131223103-2133000303130212-3221221230133232-2113131210231203): complete subsection reference.

- [simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232): complete subsection reference.

<a id="canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom_route_object` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- routes.custom_route_object

<a id="canonical-2132120132332033-0111230301202203-1332310120131330-3101230211330320-1302310110310033-2231032202200231-0302213013300201-0020203132023302"></a>

Type: `"object"`. single nested block, Optional.

A custom route uses a route object created outside of this view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]"
}
```

Terraform syntax:

```terraform
custom_route_object {
  # Configure direct properties listed below.
}
```

<a id="canonical-1310333203312222-2032203322200333-3103123323230211-1320102023312023-0013322211322022-0302203320221221-1122002313223023-1000303311113320"></a>

### Direct properties for `routes.custom_route_object`

- [caching_disable](resources--http_loadbalancer--reference--group-024.md#canonical-1311233302110112-2123122312121123-3301300201233322-3000022223123013-2103011011100200-0220322320333122-3102212130300133-0111330221321312): complete subsection reference.

- [caching_inherit](resources--http_loadbalancer--reference--group-024.md#canonical-1212201113232021-1213023333100221-3121132020000323-0333203233000210-2112121231111212-0112211131311133-3323220002021033-3201103230031110): complete subsection reference.

- [route_ref](resources--http_loadbalancer--reference--group-024.md#canonical-3032301330003202-1101233001033313-2120303011100012-0322132022031210-2133201012222110-0102201103210232-1001021232132210-0010021002202223): complete subsection reference.

<a id="canonical-1311233302110112-2123122312121123-3301300201233322-3000022223123013-2103011011100200-0220322320333122-3102212130300133-0111330221321312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom_route_object.caching_disable` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.custom_route_object](resources--http_loadbalancer--reference--group-024.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130)
- routes.custom_route_object.caching_disable

<a id="canonical-1300132233003321-2310012321003102-3313000230313012-0113022222000330-3212023132023010-0023031201321130-3322202333303202-1100023203132133"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching disable.

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
caching_disable = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212201113232021-1213023333100221-3121132020000323-0333203233000210-2112121231111212-0112211131311133-3323220002021033-3201103230031110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom_route_object.caching_inherit` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.custom_route_object](resources--http_loadbalancer--reference--group-024.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130)
- routes.custom_route_object.caching_inherit

<a id="canonical-0023203021031101-3102012330230032-2322300303011202-1102031110221332-2122231020331130-0013003023132321-3233331212131000-3111121233333103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching inherit.

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
caching_inherit = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032301330003202-1101233001033313-2120303011100012-0322132022031210-2133201012222110-0102201103210232-1001021232132210-0010021002202223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom_route_object.route_ref` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.custom_route_object](resources--http_loadbalancer--reference--group-024.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130)
- routes.custom_route_object.route_ref

<a id="canonical-1123021100312002-1020112010103001-2313033300010013-1303122122012223-0322203113312222-3001003121332000-0101330122123222-0033201313203300"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
route_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-0102303303101220-2220313232201222-2311212111013010-1023203313233111-0333303322101003-3011203321211000-0022233012120313-3313232320032320"></a>

### Direct properties for `routes.custom_route_object.route_ref`

<a id="canonical-2323000200333123-0132121222212333-2210330112033001-0112112303232223-1213013021101003-1213132222113321-2330222001330012-2320213231123322"></a>

#### `routes.custom_route_object.route_ref.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2201320121001121-0103233133130321-2101212203012032-2033101021032220-1230300223120133-1111112032301030-3312332122301100-0220322033203122"></a>
