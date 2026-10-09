---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-3202323330000322-1033101010210302-3210323011122131-1212133001021001-2332112012233312-3212021112013121-1132322010100211-2020220310311121"></a>

## `client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata.name` property

Type: `"string"`. Optional.

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

<a id="canonical-0313113200033031-0213201113202133-0212000000101012-1220101330223002-0233201331122321-2012200102121121-0232212210302121-2023222022330020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-3203113232311103-1203301023312113-1103112313002121-2300301213131000-2021012213013110-0200231011312133-0033320012121330-3322223131210231)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0033230223132100-0222132221220121-0013133202323110-3330203023221110-1223312330000301-2201203133133303-1200331111012323-0331103310011313)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-2221000100332222-2303223321131020-0111013311213213-3310323233333123-3000031113100001-2233302121020221-0021123012001131-3212012021021311"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-2013013022300321-1010002231031230-3332111021202231-3320210302101230-1330130130132300-3303201122002001-0311301130213223-3303331101011113"></a>

### Direct properties for `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path`

<a id="canonical-1111212032200320-3011301031023323-3221330302210210-0000311030221221-3101013120030130-0131130212220210-2003302020012202-0011111131132110"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path.path` property

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

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

<a id="canonical-0333303301330320-0312103102302023-0331033022033101-0333203001001100-3131333013111302-2203201101121311-1221102022233120-3101023111201021"></a>

<a id="canonical-0221310312133230-2203133213302202-0020223031132123-0231133100132111-1201120300302223-2202230012333232-1101130013022312-2111221223021233"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path.prefix` property

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2100031031030130-2012301112332302-0223131222120022-0302012310202000-2333023001003312-2300202331010122-3103322232202213-0221210311001322"></a>

<a id="canonical-3021031230112032-2013110101300330-1331023221232211-2002131021310023-1320011032311203-1313111213012323-2013232103211230-3032001113321031"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path.regex` property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- client_side_defense.policy.js_insertion_rules

<a id="canonical-2223330120010033-3330220112301033-0301003123001311-3320220321313233-1132332333110202-3003123131132332-1130231203233122-0213223013101130"></a>

Type: `"object"`. single nested block, Optional.

This defines custom JavaScript insertion rules for Client-Side Defense Policy.

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
js_insertion_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213031210023023-3000011212202323-2211002013132210-3303212311102221-2211101103303023-0100112202221020-3300301203030310-1232122001033000"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules`

- [exclude_list](resources--http_loadbalancer--reference--group-015.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222): complete subsection reference.

- [rules](resources--http_loadbalancer--reference--group-015.md#canonical-2003130231131001-2330111303132012-2002331330312302-1320031033110122-3001110220033102-2320021302020322-3323132133020121-0210202032110321): complete subsection reference.

<a id="canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- client_side_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-0302200211013330-3120012120312110-1023212310321322-1101203233113201-1201102222121132-1220100321322210-3122330033303312-2232322300130320"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203022302023012-2230320022001313-1023130211300120-1133002013231222-3222103000111031-0013311233103333-1221030113213231-0113332321000311"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list`

- [any_domain](resources--http_loadbalancer--reference--group-015.md#canonical-1211032113213031-1312033220303030-2000131032203030-0210101023300331-1302000301020223-3021320231203002-3032322113012022-3023223132032202): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-015.md#canonical-1302133302131301-0302203010332010-3032202302312331-1322312023023320-0121210331002113-3331322031031103-3212132120101113-3210001000000231): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-015.md#canonical-0323321301322131-3031120031101112-2203233330200212-1222233131100003-0310002222021030-1211312130102001-2320312203103213-1012210013102312): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-015.md#canonical-3211010123012220-1110123231023113-1023333033223033-3122203312012010-3321230002200121-0033210201023331-2132301132120200-3021313130322321): complete subsection reference.

<a id="canonical-1211032113213031-1312033220303030-2000131032203030-0210101023300331-1302000301020223-3021320231203002-3032322113012022-3023223132032202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-015.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222)
- client_side_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-2303001012113301-0023200320010131-3220013113232303-1133220220310320-1100231331323113-2322033121231203-0333110011112320-1003012332200223"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302133302131301-0302203010332010-3032202302312331-1322312023023320-0121210331002113-3331322031031103-3212132120101113-3210001000000231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-015.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222)
- client_side_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-1221121033332311-2012000301122123-2113031000012302-3012301133313221-1222101310211300-1002231102321103-0332023012300313-1300132120312113"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-2331210213101210-0233313220303113-1312100230102302-1333331111112322-1323331011232333-1222230220312330-1101113220213220-2210233131333303"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list.domain`

<a id="canonical-1122100220003232-1022313203011322-1120031301310230-0200223222310232-0313110232203130-3330022323330233-0022210020231223-3220001021202211"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.domain.exact_value` property

Type: `"string"`. Optional.

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

<a id="canonical-0031021330022122-3320020301030213-1210022001202101-1201123202331132-2300123000232203-2310232003013201-0313101321213111-1022210011321013"></a>

<a id="canonical-1221013130132331-3210012323033230-0130021000201002-0300310311300230-2330321230023201-3013322211111202-2213310002113103-0103301032223331"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.domain.regex_value` property

Type: `"string"`. Optional.

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

<a id="canonical-3212121011203331-1120101130313031-3333310303102103-1212023331231311-0330021011331013-2123220030212130-3332100020230211-2020120312022212"></a>

<a id="canonical-2311131310311013-2122030102002211-1012313100101003-3312122013103102-2201322132131320-0022130200231100-1112112332322302-3221032100201013"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.domain.suffix_value` property

Type: `"string"`. Optional.

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

<a id="canonical-0323321301322131-3031120031101112-2203233330200212-1222233131100003-0310002222021030-1211312130102001-2320312203103213-1012210013102312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-015.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222)
- client_side_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-0231121201131001-1310103223323133-2210210103100121-1030332233103330-2000032120122311-0100322132313332-0010323223301031-2200323033223010"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031013303003322-0230323310131112-2001112220310212-2320033213003102-0301232103132312-0301121033231230-3232030010200023-3233312112222321"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list.metadata`

<a id="canonical-1331203223102003-2212021133333302-0230023021300001-1002231133223310-1133101023113320-3121010031120313-1303202302033313-0100200130002003"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-3123123003033212-0003203322121032-0201303233321321-0312022301223301-1113011003002333-1003302330123101-1313333201231231-3311000312231121"></a>

<a id="canonical-0313233020301113-2001110200222000-2101113031131232-2322002321210321-0010010220311130-1330302220211332-2113110311123303-3303100210112323"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.metadata.name` property

Type: `"string"`. Optional.

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

<a id="canonical-3211010123012220-1110123231023113-1023333033223033-3122203312012010-3321230002200121-0033210201023331-2132301132120200-3021313130322321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-015.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222)
- client_side_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-0213001220012101-0132033101322101-3331002303112111-2033122100033133-1122023211110301-0211031010131320-0213233130013023-0321231231020222"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321032121132301-0233333031223330-0021233020133222-0030233330201011-3133121213011033-2311211013031233-3231330101202033-2021022022332300"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list.path`

<a id="canonical-3001213122213233-0302112032230113-0010120333220232-0320330321112320-3112023321132120-1112031030202213-2333300332023000-0013331211122322"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.path.path` property

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

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

<a id="canonical-0220323202003033-3100302023111320-1003020021320121-0033121013322003-0001222002002010-1120300023210333-1011321033310100-0111300031320233"></a>

<a id="canonical-1201203013131010-0133233210121020-0011010032312300-3230201331230000-0330332203200121-2103020231132210-1222010212211122-0201231130223022"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.path.prefix` property

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0211332330211200-2300310012010133-3311012200330220-1213210131320213-1331001012213033-1030113001111020-2331010103230113-2320330323223130"></a>

<a id="canonical-2332213022301011-2101022322121012-1222200200113332-1121333003221301-1030210101313110-3212230310101002-0133030222331210-3032112202123212"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.path.regex` property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2003130231131001-2330111303132012-2002331330312302-1320031033110122-3001110220033102-2320021302020322-3323132133020121-0210202032110321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- client_side_defense.policy.js_insertion_rules.rules

<a id="canonical-0230001303330022-1302211122321222-1112003231233221-1122000331220021-3200220102120021-1333320211222200-3002313112332223-3321013210213102"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Client-Side Defense client JavaScript.

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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210110201131120-0133121011121321-3012333100023013-0323233200323300-2200003201010223-2230133020323002-3300312002033012-0212201222033300"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.rules`

- [any_domain](resources--http_loadbalancer--reference--group-015.md#canonical-2013312033032311-1323201123212003-2012231301022113-1203300010212201-0331123303001112-2033030330233322-2301030130002010-2010202221221020): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-015.md#canonical-3121220330311231-0131113030213332-0230023132200210-2331300032230331-2201313210312311-0001100110221303-0032221321322221-1201001201330313): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-015.md#canonical-0100330120231232-3120222230110101-1221321002233010-0210010201301232-2332301333122111-2113133211302110-1213223000220300-0311111321313300): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-015.md#canonical-1100333130000013-0220213311330002-2122311302110313-0202212000023320-2113200210001211-1231003031022221-3103333213321232-2233130132132210): complete subsection reference.

<a id="canonical-2013312033032311-1323201123212003-2012231301022113-1203300010212201-0331123303001112-2033030330233322-2301030130002010-2010202221221020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-015.md#canonical-2003130231131001-2330111303132012-2002331330312302-1320031033110122-3001110220033102-2320021302020322-3323132133020121-0210202032110321)
- client_side_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-2230023003113022-3221211033121102-0320213102020330-1221233103130331-2132112103331033-3232323113230213-2313223132100231-3000232331321001"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121220330311231-0131113030213332-0230023132200210-2331300032230331-2201313210312311-0001100110221303-0032221321322221-1201001201330313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-015.md#canonical-2003130231131001-2330111303132012-2002331330312302-1320031033110122-3001110220033102-2320021302020322-3323132133020121-0210202032110321)
- client_side_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-1302212002330313-1120020221313113-1121312313033030-0322112113000100-0010230201011011-2202300200233210-0120230031333330-2222233323100032"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-2312323202100131-1012113121222200-1320201130212332-3112103312112301-0001003000031230-2202213211103300-1121312322230130-2021210310003013"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.rules.domain`

<a id="canonical-3222103102232010-1320202002130310-0222000302121200-3233321020031122-2211001030230133-0311311222013003-0030100320030201-0022311122113332"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.domain.exact_value` property

Type: `"string"`. Optional.

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

<a id="canonical-1331102032322321-3212113103101103-3220100203110211-1233203002223323-0223032000231222-1111202220230232-0031320131130121-2022222300220012"></a>

<a id="canonical-1100133223133301-1101303202303322-1320002020010321-1012010323020212-1321302210302220-3331022331303021-0101221000101330-1322121002230121"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.domain.regex_value` property

Type: `"string"`. Optional.

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

<a id="canonical-1131302131302213-0203311232322331-1100313010313000-1333311322213232-2031223303013213-3000313130000102-3031222311030132-2321033202102220"></a>

<a id="canonical-3023100032103223-3012012311030300-0111012323110031-2323331113021122-1123023023011312-3321201231321202-0211211032213233-0100100312023312"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.domain.suffix_value` property

Type: `"string"`. Optional.

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

<a id="canonical-0100330120231232-3120222230110101-1221321002233010-0210010201301232-2332301333122111-2113133211302110-1213223000220300-0311111321313300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-015.md#canonical-2003130231131001-2330111303132012-2002331330312302-1320031033110122-3001110220033102-2320021302020322-3323132133020121-0210202032110321)
- client_side_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-1202321312121132-1100031023112131-2031023123212231-3323203100202013-0002220002000113-0332333101031302-3001231213113111-1201221022323203"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200111112121000-3023112223320010-0311322333000101-0101132023222031-0301231001033311-3031310301210211-1111323220012231-0002301223203221"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.rules.metadata`

<a id="canonical-2013321300013111-0202013012122210-0220333002332123-2020022011221133-0012132023213133-0122133213111013-0020322010021222-0121133120331230"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-1030003022321102-3222202031210101-2133102233310013-2010230012101303-2220022301102223-3201123220231101-2311112003102210-3110023311113030"></a>

<a id="canonical-3010220031332133-0103023120221023-2000113232320033-1212233101323002-2322321202222103-3120011312102320-3100110013211122-2300033100331011"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.metadata.name` property

Type: `"string"`. Optional.

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

<a id="canonical-1100333130000013-0220213311330002-2122311302110313-0202212000023320-2113200210001211-1231003031022221-3103333213321232-2233130132132210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-015.md#canonical-2003130231131001-2330111303132012-2002331330312302-1320031033110122-3001110220033102-2320021302020322-3323132133020121-0210202032110321)
- client_side_defense.policy.js_insertion_rules.rules.path

<a id="canonical-3011323303023030-0311011033223133-2111331110133132-2310120031120210-0020300201303311-2232303110200120-2211332132230200-2332101201301231"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010300333121330-1322020313232000-0033013010312122-3023302130002121-3021110210112230-0031303220120333-2330200331110232-0330330100131201"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.rules.path`

<a id="canonical-3020322312232310-1333032303233130-3313010121002311-2312311330110231-3033220032122312-1232223131020123-0321220121211232-0302001222131233"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.path.path` property

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

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

<a id="canonical-0321311011311212-2112133000100011-2331313333021322-1333031213100023-2331102210130211-2222313320021021-3323023100020300-1100311230310012"></a>

<a id="canonical-1031022012032003-0030130032120110-0131320112121222-0011230022301233-1222210012323122-0233201022330132-2321120211303331-1103122123011122"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.path.prefix` property

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0130011033330103-1130321110120032-3113133111033330-1100212331230003-1233210022022303-3030230130103220-0020103031010021-3133000110023113"></a>

<a id="canonical-2312210311133300-0103023123113120-1333220321330011-2122321021230222-1330212222122232-3122100022333301-2012111230200223-1313230001233310"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.path.regex` property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_stickiness` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- cookie_stickiness

<a id="canonical-3022100130130323-2211020000021220-3132110021321012-3201303232001232-0113320030023322-3202112300132331-2333201101110023-1132101301213203"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: cookie\_stickiness, least\_active, random, ring\_hash, round\_robin,
source\_ip\_stickiness; Default: round\_robin\] Two types of cookie affinity: 1. Passive. Takes a
cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and
sets a cookie with an expiration (TTL) on the first request from the client in its response to the
client, based on the endpoint the request gets..

Additional upstream details:

Two types of cookie affinity:

&#8203;1. Generates and sets a cookie with an expiration (TTL) on the first request from the client
in its response to the client, based on the endpoint the request gets sent to. The client then
presents this on the next and all subsequent requests. The hash of this is sufficient to ensure
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

OneOf alternatives in this subsection:

- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3022100130130323-2211020000021220-3132110021321012-3201303232001232-0113320030023322-3202112300132331-2333201101110023-1132101301213203)
- [least_active](resources--http_loadbalancer--reference--group-021.md#canonical-0012031220101203-1331010113102002-1112333232131023-3022102022010211-3120131001301000-1311230213201110-2003321233300023-2003020103022211)
- [random](resources--http_loadbalancer--reference--group-024.md#canonical-1021313003120223-3101211023223000-1333330100323102-3033202113022313-2111113220003123-1112002011233323-1312030112001012-1111020111100321)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3013210122112121-0122221322331320-0013232112123331-2012132320331222-1132231302310032-2101233212003331-1003333001210203-3312200031323332)
- [round_robin](resources--http_loadbalancer--reference--group-024.md#canonical-1120231012232110-3001211103003332-2110332213132110-2233201023123010-3331031002232323-2201313122330202-3321233220103133-3330032203222203)
- [source_ip_stickiness](resources--http_loadbalancer--reference--group-028.md#canonical-0202231310310202-3102203321100120-0022003303331033-1312113312302011-2331102102230311-0003123123020233-2303210033023103-2322033033310111)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
cookie_stickiness {
  # Configure direct properties listed below.
}
```

<a id="canonical-1331103002210032-3331332220333133-1021201203322123-3001112010303020-1130132202303123-1003011103102201-1200332322130122-2300000003023311"></a>

### Direct properties for `cookie_stickiness`

- [add_httponly](resources--http_loadbalancer--reference--group-015.md#canonical-2030131120202023-2110332201010210-0223310001320000-2313012310033112-3132121323021000-3323032103012233-3203112210212112-2330330020210333): complete subsection reference.

- [add_secure](resources--http_loadbalancer--reference--group-015.md#canonical-2010011221312311-2201313103330320-0310313033132213-3230012221222320-0122233030233133-2213133110022103-2201003011023032-1310213320202133): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--reference--group-015.md#canonical-0332032201122022-2110113231101103-3030120121210231-2103232220102020-2003111103302322-2030020210312123-0023303001212200-2202330201201210): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--reference--group-015.md#canonical-0200002032300002-1032312120300023-2120012123133112-1332132122220013-1100120331301011-2012021333313032-1011120132021231-0112103013101330): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--reference--group-015.md#canonical-0033230303023301-1201320022202112-0010301331012220-0003300031000231-0223113312220220-1333122221222310-2022010220330230-3320002200012301): complete subsection reference.

<a id="canonical-3300101222320023-0123011223031301-0002233130003122-0011222000202120-0003322200232131-1112033012313123-1210220032131121-0332321200002211"></a>

<a id="canonical-3223121112012013-1010123312333132-0003103022323031-1131200230312221-3333221212111232-2000022030010021-1312331123033000-1302132231021311"></a>

#### `cookie_stickiness.name` property

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

<a id="canonical-2023003201133013-0031312232113122-3330022102003133-2210131210300003-1002322310330030-0221021013330221-1100203010313012-2300022203330233"></a>

<a id="canonical-3011133120212212-3023312330032301-2031320030200011-2321121210110113-2323320112002010-0203010103200312-0021101031133000-0221231210013223"></a>

#### `cookie_stickiness.path` property

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

- [samesite_lax](resources--http_loadbalancer--reference--group-015.md#canonical-1321001203233233-2321232003021022-0203011131012320-3203022123300223-3002101231020330-3302123321300303-2223213202333202-3330223113122132): complete subsection reference.

- [samesite_none](resources--http_loadbalancer--reference--group-015.md#canonical-1132131112302000-1122301323303133-2010003132020223-3012212122132030-1202021222123020-0223313313123132-2120032300330312-2300013111133200): complete subsection reference.

- [samesite_strict](resources--http_loadbalancer--reference--group-015.md#canonical-1011201020201132-0002101330103300-0003131200223133-2231001133012310-1230202310111221-1033332302022223-1232112022311122-1211320022123012): complete subsection reference.

<a id="canonical-0232231113020322-3022012112033322-3333212021312220-0013203100132113-0311233212103001-0332010212300020-0103212203121210-2131113133022101"></a>

<a id="canonical-1111012110032212-1311230132203002-0023000313320221-0002322322300111-3130301211033033-3211112310133302-0133002233310323-3020022111022023"></a>

#### `cookie_stickiness.ttl` property

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

<a id="canonical-2030131120202023-2110332201010210-0223310001320000-2313012310033112-3132121323021000-3323032103012233-3203112210212112-2330330020210333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_stickiness.add_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- cookie_stickiness.add_httponly

<a id="canonical-3332020222303023-3101122213113110-3101302022101200-3211313023212131-0101322223001031-1302103311102112-1111010231011000-0130102102031213"></a>

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

<a id="canonical-2010011221312311-2201313103330320-0310313033132213-3230012221222320-0122233030233133-2213133110022103-2201003011023032-1310213320202133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_stickiness.add_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- cookie_stickiness.add_secure

<a id="canonical-1021023322210122-2212100023331020-3320003213201312-3313333301030312-2321300213002230-0320221100320002-1131210332113111-2130132021033231"></a>

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

<a id="canonical-0332032201122022-2110113231101103-3030120121210231-2103232220102020-2003111103302322-2030020210312123-0023303001212200-2202330201201210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_stickiness.ignore_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- cookie_stickiness.ignore_httponly

<a id="canonical-2310100203100202-3312233101332033-2220302131000131-3001030231310333-2030133330010330-0133001012211021-2002333221300230-3103113331111301"></a>

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

<a id="canonical-0200002032300002-1032312120300023-2120012123133112-1332132122220013-1100120331301011-2012021333313032-1011120132021231-0112103013101330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_stickiness.ignore_samesite` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- cookie_stickiness.ignore_samesite

<a id="canonical-0203323210102231-3013121301101200-0111011032000101-2111331112302133-0120133022123312-0322013030003212-1101101013203001-2323032122302010"></a>

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

<a id="canonical-0033230303023301-1201320022202112-0010301331012220-0003300031000231-0223113312220220-1333122221222310-2022010220330230-3320002200012301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_stickiness.ignore_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- cookie_stickiness.ignore_secure

<a id="canonical-3031113220030213-0001030110210312-2303220332221002-0310220312302120-1321202010012202-0212221311222021-3021133232100211-0232033101122322"></a>

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

<a id="canonical-1321001203233233-2321232003021022-0203011131012320-3203022123300223-3002101231020330-3302123321300303-2223213202333202-3330223113122132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_stickiness.samesite_lax` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- cookie_stickiness.samesite_lax

<a id="canonical-2202230212121310-2231222211213332-2013231311021110-0011231231122120-1031320330112021-2032120232300010-2322013133231320-0312312213210100"></a>

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

<a id="canonical-1132131112302000-1122301323303133-2010003132020223-3012212122132030-1202021222123020-0223313313123132-2120032300330312-2300013111133200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_stickiness.samesite_none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- cookie_stickiness.samesite_none

<a id="canonical-0211121232222013-0212303220333020-2133223230103001-0003210230221032-3203201023302230-3232131323030121-3302320132200111-0310222031110112"></a>

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

<a id="canonical-1011201020201132-0002101330103300-0003131200223133-2231001133012310-1230202310111221-1033332302022223-1232112022311122-1211320022123012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_stickiness.samesite_strict` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- cookie_stickiness.samesite_strict

<a id="canonical-1100233232032301-0220302303202113-1301122133022011-1312311010301012-1302212013211111-2211221032202203-0102011032110313-1122131220220033"></a>

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

<a id="canonical-2120230123301020-2223032123333110-0222303131313113-1302023032311223-0231200001201300-0311101003010332-1110132103312122-1021322003001231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cors_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- cors_policy

<a id="canonical-2122220210330203-0220012133012121-0103130220122000-3300103303020003-3301132102100322-0231320012332001-0023320010110100-0013033001130130"></a>

Type: `"object"`. single nested block, Optional.

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence.

An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other
User-Agent: Mozilla/5.0 (Macintosh; U; Intel macOS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130
Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8
Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset:
ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive Referrer:
http&#58;//foo.example/examples/access-control/simplexsinvocation.HTML Origin:
http&#58;//foo.example

HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61
Access-Control-Allow-Origin: \* Keep-Alive: timeout=2, max=100 Connection: Keep-Alive
Transfer-Encoding: chunked Content-Type: application/XML

An example for cross origin HTTP OPTIONS request with Access-Control-Request-\* header

OPTIONS /resources/POST-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel
macOS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept:
text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8 Accept-Language: en-us,en;q=0.5
Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive
Origin: http&#58;//foo.example Access-Control-Request-Method: POST Access-Control-Request-Headers:
X-PINGOTHER, Content-Type

HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix)
Access-Control-Allow-Origin: http&#58;//foo.example Access-Control-Allow-Methods: POST, GET, OPTIONS
Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary:
Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive.

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
cors_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-3332203130031002-2130211100312200-3332133122021002-0021233001323112-0303302302231330-2110012032300211-3330112232231322-3312010203123010"></a>

### Direct properties for `cors_policy`

<a id="canonical-2103312020020203-3333220012221320-1011221213101322-1031320323001031-2221001211313331-2103112130002123-3220212330231331-0323110331000103"></a>

#### `cors_policy.allow_credentials` property

Type: `"bool"`. Optional.

Specifies whether the resource allows credentials.

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

<a id="canonical-0312132123213120-0333030012111231-0230130232110131-1110112220000200-2133330303130311-2032332100222001-1210012121323132-3212102312201220"></a>

<a id="canonical-1333121210323323-2221221101122101-2001200001100333-3300123211002112-1232023311223020-1223200121301221-1023131131032301-2232010133321023"></a>

#### `cors_policy.allow_headers` property

Type: `"string"`. Optional.

Specifies the content for the access-control-allow-headers header.

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

<a id="canonical-1213231131101311-3111131223122220-1031103301321112-0230031132001023-0302002220332310-3120332313023132-2310132322112230-2111021330110232"></a>

<a id="canonical-2101123202023011-2032330113012300-2200011031233332-3031033311020013-1201210210111333-3130232131011303-1300111232333022-1121000121013232"></a>

#### `cors_policy.allow_methods` property

Type: `"string"`. Optional.

Specifies the content for the access-control-allow-methods header.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  }
}
```

<a id="canonical-0100212323112021-0331311000310311-3003221003110031-2222320322213030-1111332322103002-3022112113223101-3203112031100030-3003203131301023"></a>

<a id="canonical-2212330011323101-2221131213100110-2320100123110000-0031323302210003-2303001021232130-1023300032101310-3322211111011012-0331320003320330"></a>

#### `cors_policy.allow_origin` property

Type: `["list", "string"]`. Optional.

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0133020312302320-2321013322230203-3123121030121201-1101302200133321-3202013121012332-1002132231202331-1222011301202012-3313131333030133"></a>

<a id="canonical-2131212212233212-3302213020321010-2020323132331233-0032130232331002-2111023113131231-1221213113311133-3012033011131333-3022032233110230"></a>

#### `cors_policy.allow_origin_regex` property

Type: `["list", "string"]`. Optional.

Specifies regular expression patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2022101023010210-0120222010303222-1021223031121010-2111223212112312-1312330101312012-0301230032231210-3302003133020310-3233312100313023"></a>

<a id="canonical-3231110110230303-3033022032230023-0312222011120021-1133023212110302-0121303202021121-1120030332210203-1222030333230123-2303210202302323"></a>

#### `cors_policy.disabled` property

Type: `"bool"`. Optional.

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

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

<a id="canonical-0331203301123313-2223123102030323-3233312033123013-2312210311330011-3220223131301112-1100123323212121-1130130323133331-0131002111002321"></a>

<a id="canonical-3303201011133020-1212003330221023-1221321030310312-3003313130323123-1303110110310201-1330010220113011-0322222011000232-1221130003331233"></a>

#### `cors_policy.expose_headers` property

Type: `"string"`. Optional.

Specifies the content for the access-control-expose-headers header.

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

<a id="canonical-1233333301122231-2102032220333021-2331230323103131-2223031130101201-2033112320233330-0301011233031231-0113201022102203-0213111110032200"></a>

<a id="canonical-3131013022130113-0101003101310112-1222103310011333-3022211221112031-1203301233333222-0203001012322203-0232220300321121-3133322013010003"></a>

#### `cors_policy.maximum_age` property

Type: `"number"`. Optional.

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": -1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  }
}
```

<a id="canonical-3021222013003320-3230102012010221-2203233323210232-0111132233210023-1332310233100013-1310012310002212-3333112112231001-0212212331211311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `csrf_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- csrf_policy

<a id="canonical-3110111223222131-2221222112302020-0022332210201221-2133220221231233-0223031112211323-1013213130211301-1213001220101100-2323003032300302"></a>

Type: `"object"`. single nested block, Optional.

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host.

&#8203;1. The origin that caused the user agent to issue the request (source origin). &#8203;2. The
origin that the request is going to (target origin). When the policy evaluating a request, it
ensures both pieces of information are present and compare their values. If the source origin is
missing or origins do not match the request is rejected. The exception to this being if the
source-origin has been added to they policy as valid. Because CSRF attacks specifically target
state-changing requests, the policy only acts on the HTTP requests that have state-changing method
(PUT,POST, etc.).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allowed_domains": "[\"all_load_balancer_domains\",\"custom_domain_list\",\"disabled\"]"
}
```

Terraform syntax:

```terraform
csrf_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202123002032021-3301221123313032-3021131232112021-3201201121030332-3221232120132013-2130022021300021-2132131102112333-0020303020030311"></a>

### Direct properties for `csrf_policy`

- [all_load_balancer_domains](resources--http_loadbalancer--reference--group-015.md#canonical-1100011121020001-1001321330300312-0203202022123003-2132001123322332-3232332300111011-2033132102233311-2033221020111023-0203100123121112): complete subsection reference.

- [custom_domain_list](resources--http_loadbalancer--reference--group-015.md#canonical-0222321213110001-0022132132020121-1030321303331231-2211122103230121-0011220321233023-2001132331233002-3331332312033013-3333011110233202): complete subsection reference.

- [disabled](resources--http_loadbalancer--reference--group-015.md#canonical-2302111201210301-0223031011201123-1110101312012031-3111331233023213-3200312032033203-1130130233230303-0303112233300100-2300331101013131): complete subsection reference.

<a id="canonical-1100011121020001-1001321330300312-0203202022123003-2132001123322332-3232332300111011-2033132102233311-2033221020111023-0203100123121112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `csrf_policy.all_load_balancer_domains` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [csrf_policy](resources--http_loadbalancer--reference--group-015.md#canonical-3021222013003320-3230102012010221-2203233323210232-0111132233210023-1332310233100013-1310012310002212-3333112112231001-0212212331211311)
- csrf_policy.all_load_balancer_domains

<a id="canonical-0101131132222003-1210021220000330-0211022303000021-2311310333230110-0220122301110312-1110122320130322-1001002002030011-2003102030220023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all load balancer domains.

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
all_load_balancer_domains = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222321213110001-0022132132020121-1030321303331231-2211122103230121-0011220321233023-2001132331233002-3331332312033013-3333011110233202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `csrf_policy.custom_domain_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [csrf_policy](resources--http_loadbalancer--reference--group-015.md#canonical-3021222013003320-3230102012010221-2203233323210232-0111132233210023-1332310233100013-1310012310002212-3333112112231001-0212212331211311)
- csrf_policy.custom_domain_list

<a id="canonical-0123221322012310-2133221301301120-2030133002030332-2321011120233111-0102211222310212-2121102102303221-1323012332330300-2113133132212130"></a>

Type: `"object"`. single nested block, Optional.

List of domain names used for Host header matching.

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
custom_domain_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333020013323211-3103100311323310-0332001021032021-3321323312022333-2201202021331013-2023231002313032-0230032030113130-1231121321000312"></a>

### Direct properties for `csrf_policy.custom_domain_list`

<a id="canonical-0333211032010300-2003230311122201-0302230221123211-1122322033333333-2212333130120321-2203121232230320-1201233322302330-2110201102031101"></a>

#### `csrf_policy.custom_domain_list.domains` property

Type: `["list", "string"]`. Optional.

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2302111201210301-0223031011201123-1110101312012031-3111331233023213-3200312032033203-1130130233230303-0303112233300100-2300331101013131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `csrf_policy.disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [csrf_policy](resources--http_loadbalancer--reference--group-015.md#canonical-3021222013003320-3230102012010221-2203233323210232-0111132233210023-1332310233100013-1310012310002212-3333112112231001-0212212331211311)
- csrf_policy.disabled

<a id="canonical-2023310232303231-0332121311230033-0220311322311233-2021130230233211-1003221002113303-3231330033121133-3100120331123131-2013231300212102"></a>

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

<a id="canonical-0310100223013020-2111222133111333-2201233312231132-0102103013122132-0111220303113031-1122123223301021-1030212023023130-0331013321302221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `data_guard_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- data_guard_rules

<a id="canonical-2022010320103333-3032231310103322-1300320033111030-3203121310020023-2001031202123103-1030032000023310-2212303022223020-2211202010031032"></a>

Type: `"object"`. list nested block, Optional.

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*). Note: App Firewall should be enabled, to use Data
Guard feature.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
data_guard_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233333013303032-3003212320330031-0221013312102010-1301333011112103-0312232010113133-1000221321021011-2102003320210022-2320310302121112"></a>

### Direct properties for `data_guard_rules`

- [any_domain](resources--http_loadbalancer--reference--group-015.md#canonical-1100223213201200-3013211301203301-2232221301023211-0220003020233130-2322231030203220-2211032131231230-0123301020332223-0231223120303232): complete subsection reference.

- [apply_data_guard](resources--http_loadbalancer--reference--group-015.md#canonical-3001102122030200-3301002312100211-1002301213223202-3232311010130213-0101130222013123-2003210012130030-3032001333131231-0330002120111321): complete subsection reference.

<a id="canonical-1030113002301031-3330130301130223-1110023023300330-2232212331233223-3100031222110313-2022103322003031-1133032021113133-0200122311121112"></a>

<a id="canonical-3120113130333330-3130020033000330-3200013233223010-0123221300300002-2232210231310031-3210220202332322-2300221332111013-3032110113223211"></a>

#### `data_guard_rules.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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

- [metadata](resources--http_loadbalancer--reference--group-015.md#canonical-0323221032321021-2030300311221110-0222222132031132-1122210101230112-1323302312123031-0320301102300213-0002110331132031-3102011031231300): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-015.md#canonical-0123302113001130-0111220321333012-1301120122101023-0022201302232111-3220031110222113-1112013302111132-3121322030113223-3213132330021320): complete subsection reference.

- [skip_data_guard](resources--http_loadbalancer--reference--group-015.md#canonical-2223223002232112-0031332231112220-0313321012300221-3130321302211113-3121021320121010-1303023132001130-1332002003021030-2033221213313322): complete subsection reference.

<a id="canonical-1212022010023031-1020003122033012-1220201211111332-1012332103332002-0122123333123002-3211020133033013-2112221320003201-2323012313003003"></a>

<a id="canonical-0202111102112201-2303110001033121-2222020222232323-0232201203230322-2313022230122003-3303000011103320-1201021332133020-1020111300300303"></a>

#### `data_guard_rules.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
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

<a id="canonical-1100223213201200-3013211301203301-2232221301023211-0220003020233130-2322231030203220-2211032131231230-0123301020332223-0231223120303232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `data_guard_rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [data_guard_rules](resources--http_loadbalancer--reference--group-015.md#canonical-0310100223013020-2111222133111333-2201233312231132-0102103013122132-0111220303113031-1122123223301021-1030212023023130-0331013321302221)
- data_guard_rules.any_domain

<a id="canonical-3122033332022100-3220201321022123-2312230120002032-0332312320203222-2330110030020301-1313210121332210-2233000100312213-1002221310300110"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001102122030200-3301002312100211-1002301213223202-3232311010130213-0101130222013123-2003210012130030-3032001333131231-0330002120111321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `data_guard_rules.apply_data_guard` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [data_guard_rules](resources--http_loadbalancer--reference--group-015.md#canonical-0310100223013020-2111222133111333-2201233312231132-0102103013122132-0111220303113031-1122123223301021-1030212023023130-0331013321302221)
- data_guard_rules.apply_data_guard

<a id="canonical-3220121322031232-2121032000121113-1232323022310310-2310200322213012-1313121203133210-2123031213302111-3213211203103200-1233131033203211"></a>

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
apply_data_guard = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323221032321021-2030300311221110-0222222132031132-1122210101230112-1323302312123031-0320301102300213-0002110331132031-3102011031231300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `data_guard_rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [data_guard_rules](resources--http_loadbalancer--reference--group-015.md#canonical-0310100223013020-2111222133111333-2201233312231132-0102103013122132-0111220303113031-1122123223301021-1030212023023130-0331013321302221)
- data_guard_rules.metadata

<a id="canonical-2320020022211320-0330312322123232-0303022211001130-0011331131130132-0030122012322233-2103100120130202-2300113020323203-2300303122122013"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303202131133201-1011000310002201-2002313023132323-0000221200010120-3100212131002033-3300000212120100-3032201212110220-2222300002112301"></a>

### Direct properties for `data_guard_rules.metadata`

<a id="canonical-0311112333212210-0030233110132002-2220330103032021-3202110123102323-2111333023132031-1132023121210301-1123230032221031-3000101302203302"></a>

#### `data_guard_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-3220103033020102-0300211313301322-0023133030121312-3013130233103311-0233011330013210-3313000311302123-2321212223102122-3201110320112310"></a>

<a id="canonical-0003313110033010-1013023113231233-2031300200030022-0002012213300212-2331212323323032-1301012002230100-2230003003212201-1200033233031120"></a>

#### `data_guard_rules.metadata.name` property

Type: `"string"`. Optional.

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

<a id="canonical-0123302113001130-0111220321333012-1301120122101023-0022201302232111-3220031110222113-1112013302111132-3121322030113223-3213132330021320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `data_guard_rules.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [data_guard_rules](resources--http_loadbalancer--reference--group-015.md#canonical-0310100223013020-2111222133111333-2201233312231132-0102103013122132-0111220303113031-1122123223301021-1030212023023130-0331013321302221)
- data_guard_rules.path

<a id="canonical-1021020212101000-1123031231223302-1220233203310030-1231233223231230-2120210323320312-3102001301011200-1103120202102333-2300122031101330"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-3000303222223331-2320013133021011-0132010312331322-0112300201202223-2030123131313232-1201201032020033-3103001013203222-0101021212233320"></a>

### Direct properties for `data_guard_rules.path`

<a id="canonical-2200112303000130-2123103121311323-2321000030002100-0223012200022303-1233021230322100-2221202312202313-1111020203122001-0000112213012130"></a>

#### `data_guard_rules.path.path` property

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

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

<a id="canonical-1202323022333232-0310130013023121-1202101230221100-0213110103330103-3232323210303332-0202130303011002-0212110233223013-3221133330132213"></a>

<a id="canonical-2020302332033112-1011310333033132-1311102231333331-2231313023233331-0331320212121202-0201310203303333-0101003303001122-0232003123310022"></a>

#### `data_guard_rules.path.prefix` property

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0322201030313310-0320302302302300-2120321223220032-2012203223300123-3222033000331333-0220123131101120-3012122021001222-3003100230233323"></a>

<a id="canonical-0121121332300023-2232101123211122-0223203120123030-2203013200122332-0210213210301032-2020023021010213-0012130300211332-3322001122213211"></a>

#### `data_guard_rules.path.regex` property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2223223002232112-0031332231112220-0313321012300221-3130321302211113-3121021320121010-1303023132001130-1332002003021030-2033221213313322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `data_guard_rules.skip_data_guard` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [data_guard_rules](resources--http_loadbalancer--reference--group-015.md#canonical-0310100223013020-2111222133111333-2201233312231132-0102103013122132-0111220303113031-1122123223301021-1030212023023130-0331013321302221)
- data_guard_rules.skip_data_guard

<a id="canonical-3233222210111123-1013222330323001-0200132032310312-1101310023333110-0212121333311320-1230121223202330-1021222302033333-3223301232133331"></a>

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
skip_data_guard = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113021303031131-1312012222010231-0331133200303132-3332023021101222-0113133213312010-1321130002311311-0012301201200102-1130020133302102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- ddos_mitigation_rules

<a id="canonical-2310100011112312-2222030220022330-3320033003021233-2022100203132203-1201120200323100-3330020021130132-0200021230032331-3111102132020133"></a>

Type: `"object"`. list nested block, Optional.

Define manual mitigation rules to block L7 DDoS attacks.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
ddos_mitigation_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0300301232101231-3132222012000123-0020311120110213-2102111212103212-0133111022030313-1031111230002020-1111112210021030-2013121331212033"></a>

### Direct properties for `ddos_mitigation_rules`

- [block](resources--http_loadbalancer--reference--group-015.md#canonical-3313103130322312-1312312213233301-2230220011033311-3313232113322333-2030110002201231-2102332030332003-3022223120213302-0300031033230022): complete subsection reference.

- [ddos_client_source](resources--http_loadbalancer--reference--group-015.md#canonical-2011213210300031-0330202201103102-0322200312222011-3232303332012220-2100231300233131-1100312130200112-2233220130221320-3011233231000021): complete subsection reference.

<a id="canonical-1221230000220131-2330202012021201-2200103312100233-1210121210221131-3333320132112031-2000120123133230-3233110030301000-2112333201220112"></a>

<a id="canonical-1002103232102000-0001121123121201-3133323012130212-2213203112310012-3302101023112211-3032322111100032-0333320303210200-2213021300311322"></a>

#### `ddos_mitigation_rules.expiration_timestamp` property

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Additional upstream details:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [ip_prefix_list](resources--http_loadbalancer--reference--group-016.md#canonical-1321132311131313-2033121331313303-3130133112123221-3130131220230331-1321011002332303-0231002203033031-0101022331300202-0033031003200210): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-016.md#canonical-1312030223203120-3311011003320303-0233321311230000-0010111210311333-3030331001001331-1311333111331332-1032200133313120-3030313233113233): complete subsection reference.

<a id="canonical-3313103130322312-1312312213233301-2230220011033311-3313232113322333-2030110002201231-2102332030332003-3022223120213302-0300031033230022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ddos_mitigation_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3113021303031131-1312012222010231-0331133200303132-3332023021101222-0113133213312010-1321130002311311-0012301201200102-1130020133302102)
- ddos_mitigation_rules.block

<a id="canonical-1100230323010312-1023121121001200-0122300331113102-3313331333120102-0032111111300102-3123210200221201-0311331233020210-2303001131103331"></a>

Type: `"object"`. single nested block, Optional.

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
block {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011213210300031-0330202201103102-0322200312222011-3232303332012220-2100231300233131-1100312130200112-2233220130221320-3011233231000021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.ddos_client_source` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ddos_mitigation_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3113021303031131-1312012222010231-0331133200303132-3332023021101222-0113133213312010-1321130002311311-0012301201200102-1130020133302102)
- ddos_mitigation_rules.ddos_client_source

<a id="canonical-1113120003032003-3220231102031110-0001102002031220-3131223331211003-1333113122001022-2030232100023302-3331023322003310-3231121312323313"></a>

Type: `"object"`. single nested block, Optional.

DDoS Client Source Choice. DDoS Mitigation sources to be blocked.

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
ddos_client_source {
  # Configure direct properties listed below.
}
```

<a id="canonical-2030233203120230-1312202310323031-0223111103110213-1131312010021103-2013210210332213-2332330100030132-3100003222120120-3322123120221122"></a>

### Direct properties for `ddos_mitigation_rules.ddos_client_source`

- [asn_list](resources--http_loadbalancer--reference--group-016.md#canonical-1301330111333132-0110110311312222-3123210123222332-1230131003020231-3003133001301200-2020233211222110-0333112122121121-0030000200322323): complete subsection reference.

<a id="canonical-1210033210203131-2130312012230202-3023102111111310-2202111200220332-1130122023211331-1112020111103223-0131320300232321-3023010103032230"></a>

<a id="canonical-0022231012112210-0120220322303302-3012321320213003-3302033231120331-0131123331012320-1102311211202133-0313120023300002-0113033020303120"></a>

#### `ddos_mitigation_rules.ddos_client_source.country_list` property

Type: `["list", "string"]`. Optional.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Sources that are located in one of the countries in the given list. Possible values are
\`COUNTRY\_NONE\`, \`COUNTRY\_AD\`, \`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`,
\`COUNTRY\_AI\`, \`COUNTRY\_AL\`, \`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`,
\`COUNTRY\_AQ\`, \`COUNTRY\_AR\`, \`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`,
\`COUNTRY\_AW\`, \`COUNTRY\_AX\`, \`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`,
\`COUNTRY\_BD\`, \`COUNTRY\_BE\`, \`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`,
\`COUNTRY\_BI\`, \`COUNTRY\_BJ\`, \`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`,
\`COUNTRY\_BO\`, \`COUNTRY\_BQ\`, \`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`,
\`COUNTRY\_BV\`, \`COUNTRY\_BW\`, \`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`,
\`COUNTRY\_CC\`, \`COUNTRY\_CD\`, \`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`,
\`COUNTRY\_CI\`, \`COUNTRY\_CK\`, \`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`,
\`COUNTRY\_CO\`, \`COUNTRY\_CR\`, \`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`,
\`COUNTRY\_CW\`, \`COUNTRY\_CX\`, \`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`,
\`COUNTRY\_DJ\`, \`COUNTRY\_DK\`, \`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`,
\`COUNTRY\_EC\`, \`COUNTRY\_EE\`, \`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`,
\`COUNTRY\_ES\`, \`COUNTRY\_ET\`, \`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`,
\`COUNTRY\_FM\`, \`COUNTRY\_FO\`, \`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`,
\`COUNTRY\_GD\`, \`COUNTRY\_GE\`, \`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`,
\`COUNTRY\_GI\`, \`COUNTRY\_GL\`, \`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`,
\`COUNTRY\_GQ\`, \`COUNTRY\_GR\`, \`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`,
\`COUNTRY\_GW\`, \`COUNTRY\_GY\`, \`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`,
\`COUNTRY\_HR\`, \`COUNTRY\_HT\`, \`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`,
\`COUNTRY\_IL\`, \`COUNTRY\_IM\`, \`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`,
\`COUNTRY\_IR\`, \`COUNTRY\_IS\`, \`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`,
\`COUNTRY\_JO\`, \`COUNTRY\_JP\`, \`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`,
\`COUNTRY\_KI\`, \`COUNTRY\_KM\`, \`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`,
\`COUNTRY\_KW\`, \`COUNTRY\_KY\`, \`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`,
\`COUNTRY\_LC\`, \`COUNTRY\_LI\`, \`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`,
\`COUNTRY\_LT\`, \`COUNTRY\_LU\`, \`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`,
\`COUNTRY\_MC\`, \`COUNTRY\_MD\`, \`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`,
\`COUNTRY\_MH\`, \`COUNTRY\_MK\`, \`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`,
\`COUNTRY\_MO\`, \`COUNTRY\_MP\`, \`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`,
\`COUNTRY\_MT\`, \`COUNTRY\_MU\`, \`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`,
\`COUNTRY\_MY\`, \`COUNTRY\_MZ\`, \`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`,
\`COUNTRY\_NF\`, \`COUNTRY\_NG\`, \`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`,
\`COUNTRY\_NP\`, \`COUNTRY\_NR\`, \`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`,
\`COUNTRY\_PA\`, \`COUNTRY\_PE\`, \`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`,
\`COUNTRY\_PK\`, \`COUNTRY\_PL\`, \`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`,
\`COUNTRY\_PS\`, \`COUNTRY\_PT\`, \`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`,
\`COUNTRY\_RE\`, \`COUNTRY\_RO\`, \`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`,
\`COUNTRY\_SA\`, \`COUNTRY\_SB\`, \`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`,
\`COUNTRY\_SG\`, \`COUNTRY\_SH\`, \`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`,
\`COUNTRY\_SL\`, \`COUNTRY\_SM\`, \`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`,
\`COUNTRY\_SS\`, \`COUNTRY\_ST\`, \`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`,
\`COUNTRY\_SZ\`, \`COUNTRY\_TC\`, \`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`,
\`COUNTRY\_TH\`, \`COUNTRY\_TJ\`, \`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`,
\`COUNTRY\_TN\`, \`COUNTRY\_TO\`, \`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`,
\`COUNTRY\_TW\`, \`COUNTRY\_TZ\`, \`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`,
\`COUNTRY\_US\`, \`COUNTRY\_UY\`, \`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`,
\`COUNTRY\_VE\`, \`COUNTRY\_VG\`, \`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`,
\`COUNTRY\_WF\`, \`COUNTRY\_WS\`, \`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`,
\`COUNTRY\_YT\`, \`COUNTRY\_ZA\`, \`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [ja4_tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-016.md#canonical-2210313030223112-0221132223013102-2133320023222302-2223212323022331-0111110010032032-2231320102021233-2132022013210322-2231020122012333): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-016.md#canonical-0031133313011232-1222230120123210-0233333120231321-3133333023031031-2020213302230211-1112301011201130-2033120322011123-0220211013233023): complete subsection reference.
