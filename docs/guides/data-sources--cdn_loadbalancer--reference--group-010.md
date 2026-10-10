---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-2232103020322200-2100001030213030-3002323122001031-0301323110222023-3211013012322123-3310313233001322-0012323001233012-3023322013210122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- client_side_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-2211130223020030-1231012110303222-3122333312301310-3202221030113322-3332222101330012-1011133001131113-2121100211111301-1332222321311133"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3311031320022013-0303320333001310-0132301133113033-1110000220330032-0213120213233323-3222103111111331-0321133130200223-3220331230221210"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list`

- [any_domain](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-3012223032110033-1130120112300121-1233121222312313-0233222022101202-2200100020332031-1302133220221200-1102212111211210-2321303321322123): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-3131121320121202-3221213223220000-0331022122223230-2223110103130113-1101231130221123-1323101201232101-0023211212000120-2213321211220323): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1332223322023101-3233103313120113-1012010301230212-1113312012303223-2210322002311311-1301222222220112-3002220231022011-1231023231223321): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2332022103111110-1222123202220231-3023030032100301-3101123031211231-0233101132200313-2121013323222301-0111130201132120-1311231313032201): complete subsection reference.

<a id="canonical-3012223032110033-1130120112300121-1233121222312313-0233222022101202-2200100020332031-1302133220221200-1102212111211210-2321303321322123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2232103020322200-2100001030213030-3002323122001031-0301323110222023-3211013012322123-3310313233001322-0012323001233012-3023322013210122)
- client_side_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-1302023022322231-0003222003231102-3322123201022211-1123200202211113-3231221201113202-0131130312001302-0301211203212123-2123331030231023"></a>

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

<a id="canonical-3131121320121202-3221213223220000-0331022122223230-2223110103130113-1101231130221123-1323101201232101-0023211212000120-2213321211220323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2232103020322200-2100001030213030-3002323122001031-0301323110222023-3211013012322123-3310313233001322-0012323001233012-3023322013210122)
- client_side_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-2023311133231002-2302301330003101-3203331131123012-0113310202330102-1203101111200333-2020232103023023-1121202303202103-1120130313302330"></a>

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

<a id="canonical-0202332100330222-1133202203111201-1201000303303330-1030103113213101-3130323211211310-0102011130001112-2120123330303001-1132221222111312"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list.domain`

<a id="canonical-3222101221221000-0300221130033210-2311030303320032-2013032203012300-1122220233221001-3220111030222022-0312020330122132-3233001100021300"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.domain.exact_value` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0231201332102203-1001232110322202-1002303003011131-3211123301333011-0310203223313033-2103232113132201-1001323020233011-1112322203133320"></a>

<a id="canonical-2300023112100200-1302221212323111-3033013300032311-0331300212232030-2132220132320230-1030331111230133-0210202233130312-1012303031230202"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.domain.regex_value` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3321113011321302-3301223230101131-3032102213311133-1223132020201320-2111113330220311-0210000232122211-2232320200321200-2303123221233131"></a>

<a id="canonical-0013013132123101-2003011020333311-3331333312331202-3310111312112130-3031322221030111-3112002023123100-0023302200032201-1212011132210002"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.domain.suffix_value` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1332223322023101-3233103313120113-1012010301230212-1113312012303223-2210322002311311-1301222222220112-3002220231022011-1231023231223321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2232103020322200-2100001030213030-3002323122001031-0301323110222023-3211013012322123-3310313233001322-0012323001233012-3023322013210122)
- client_side_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-2001222320322020-2203332221001323-3212311122000001-3200301333133012-2312331003301100-3112131130231300-2123200101210012-2031210022000332"></a>

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

<a id="canonical-0321012210002123-0101313233312022-2220100222232030-2201331212123311-2120311121012102-3030011201110312-0133003103212312-0023030010332130"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list.metadata`

<a id="canonical-2310030131330320-1010020321021300-2021030012203023-3101020100203022-0021112013223312-1213110321202030-2220203122301131-2333120032201001"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0321023223233132-3123223332233130-1303101032231303-2011210112332021-0202213333230322-0032102120312130-1133121310020020-3211102102122121"></a>

<a id="canonical-2111230102223131-2330122211203113-2013232000012223-3121330121310233-3222002213010120-2230300100203110-3123102101100011-2321000110130322"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.metadata.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2332022103111110-1222123202220231-3023030032100301-3101123031211231-0233101132200313-2121013323222301-0111130201132120-1311231313032201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.path` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2232103020322200-2100001030213030-3002323122001031-0301323110222023-3211013012322123-3310313233001322-0012323001233012-3023322013210122)
- client_side_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-1322132110031212-1101110031210123-3013211332103112-2002333001213131-3331222030032310-2033232330032112-2212102002100013-1200200023112023"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3032300112122210-3113130001301001-0333031101203202-1301222101203130-1332222212113132-3033230320032322-1211231011233020-3012130300203001"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list.path`

<a id="canonical-1302212100000100-0232000220310301-3321130331213111-0131301121213130-0012203020311332-3213200000201222-2122330120032100-3122012122322121"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.path.path` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2113210221301102-1213120132301333-0002103121112323-3221112222232032-2031023000120003-0121003203321113-3231033322130131-3011312133322033"></a>

<a id="canonical-3010132133233100-0303022200133133-3021211133202303-3233023022232310-1030231010012311-3100102030112123-2301333022200113-2201202230221103"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.path.prefix` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2003120121330122-1232121312001302-0201330030100332-3313201010002000-1211231023222131-0123103001113332-1210032002231013-1303333122213211"></a>

<a id="canonical-0233312302312223-3222020212103223-0001220202312133-2011333003122321-3203131012320002-2103120132122000-1203200221210032-2322011211302200"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.path.regex` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2230201130322102-2313103001002212-3103231131202123-3312132213332213-0302000332310333-3110032231331023-1210321321210233-3321012200203101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- client_side_defense.policy.js_insertion_rules.rules

<a id="canonical-3200230122200032-3300301330231113-0031111221301210-0303203331322003-3022320123000303-2122011032030122-2222101330123232-1301021102210323"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0311023312001123-1122310230220331-2303100213323101-1320133030310320-3311301011302103-0031300213230210-3132000212100110-2110220122013301"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.rules`

- [any_domain](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1032233310300203-0121031021221100-0200130010110113-0020131202333001-0111322332320222-1103201100110022-0010332101121003-3022120310323222): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0331232032213020-3320203122130001-0201311133300031-2321200333311022-0222110302110002-2013020031210203-2213330012122123-2311030210200222): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0322112110103312-2132031001323311-0100200330331100-2003310002322221-2030030202121032-0232221323230110-3323313020203102-1113013322133011): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-3211221110133300-0010101132222313-1303220022201110-0323013113201030-0033121131120201-0002231112001311-3111012100020011-0002121300101200): complete subsection reference.

<a id="canonical-1032233310300203-0121031021221100-0200130010110113-0020131202333001-0111322332320222-1103201100110022-0010332101121003-3022120310323222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2230201130322102-2313103001002212-3103231131202123-3312132213332213-0302000332310333-3110032231331023-1210321321210233-3321012200203101)
- client_side_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-1021102132132221-1013333031222012-3003100010313001-2301010222102001-2022302112130213-0033312020301103-3222011223203310-3023011220032120"></a>

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

<a id="canonical-0331232032213020-3320203122130001-0201311133300031-2321200333311022-0222110302110002-2013020031210203-2213330012122123-2311030210200222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules.domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2230201130322102-2313103001002212-3103231131202123-3312132213332213-0302000332310333-3110032231331023-1210321321210233-3321012200203101)
- client_side_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-3323133330301202-0201232102131232-0233332303300230-2311323213033323-3313233201213303-0302020212331201-1023311222230010-2020331133300133"></a>

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

<a id="canonical-1013231210010233-3302123220002233-0300001101030303-2000200331230121-2132312003312230-1330322220112220-0302211320032321-2211010103122032"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.rules.domain`

<a id="canonical-1330100022023110-2302320310103211-1022102312200202-3210333033001312-3001200210130030-0323121023310333-3033000310120113-3021131130022021"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.domain.exact_value` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0220223112133321-3331321201310120-1011330010233111-2203110131031300-1003321100031132-1123301330331000-1233030221132003-0312310012130133"></a>

<a id="canonical-3012033310122302-0000122211123113-2122323200330013-2212232130210211-1033003001111121-0233330100232302-0201213113103113-0213322211323103"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.domain.regex_value` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2123010300031203-2222230320221002-0232033102133321-3002032132211202-2312111121330212-3333333122112003-1331033311311130-2222231333020322"></a>

<a id="canonical-0003013303023231-0222321312123132-3110122032332222-0021320213002113-3222202222321112-2303131232332133-1312121113111020-0212011233311013"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.domain.suffix_value` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0322112110103312-2132031001323311-0100200330331100-2003310002322221-2030030202121032-0232221323230110-3323313020203102-1113013322133011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2230201130322102-2313103001002212-3103231131202123-3312132213332213-0302000332310333-3110032231331023-1210321321210233-3321012200203101)
- client_side_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-0010021131003311-1313020211232133-0102112311302320-1330013030330231-1223210323032303-3012322121320233-2330123312101100-3333033331113222"></a>

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

<a id="canonical-0302311230131110-0321220320311122-2020032202022123-3200120000212232-2333010032332121-0220112332313011-3031310301110122-1121102130321332"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.rules.metadata`

<a id="canonical-0130003130032311-1020011221312201-0123312112031300-0301013001203003-2122232031120003-2221011100010023-2021031330111212-0210011103013322"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0010212330321201-2322023232300011-2003133121001221-1003103330310301-0022102210130201-3111300230113210-0302310112023021-3133311120012111"></a>

<a id="canonical-2233130021302230-2301211210233010-2320203011100203-0103030000110030-0002233203022131-2203312132213300-3110100221321301-1001100103210321"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.metadata.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3211221110133300-0010101132222313-1303220022201110-0323013113201030-0033121131120201-0002231112001311-3111012100020011-0002121300101200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules.path` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2230201130322102-2313103001002212-3103231131202123-3312132213332213-0302000332310333-3110032231331023-1210321321210233-3321012200203101)
- client_side_defense.policy.js_insertion_rules.rules.path

<a id="canonical-3110302001310002-0210231233203100-1013320212021002-1112210130102120-2210123023102231-2211032132321122-2311230323100120-1100021023300023"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2320322213323033-2013133200132030-1132002011222333-2132313200203112-0011322313220111-3122010231033113-1321002023121020-2011003121011002"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.rules.path`

<a id="canonical-3133330133311133-0000001003330122-1120133301320331-3313030010030110-2303112313101030-3333103110131212-1223112223102021-2101213033232023"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.path.path` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3011233113231001-2103202313130102-1222333202320222-1130330101033130-1222023301000002-3213212003302101-1000221101331012-3001133230002300"></a>

<a id="canonical-1311332012113120-1000011303032323-0021121311120130-3110322200200233-3312203220311130-3030223201001311-1123000001022231-0220331310220232"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.path.prefix` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2200323010020202-3200332212121222-1303000102102123-0033201120020110-0011110333101021-1300130312013123-2030302111311003-3203310211330000"></a>

<a id="canonical-0332202020113332-0002330333101230-2131223022113310-0031122211301302-3331323202233323-0233333202223232-1123211020122323-2313333222213220"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.path.regex` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3330100113233220-0200103010232013-0220221020032201-2012031211002200-0321000202211233-0202201031003232-2322233111100113-1212133131020203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cors_policy` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- cors_policy

<a id="canonical-1021113222023013-2112213003123130-2310110301101012-2230123300011202-3131230211333112-0332202330320122-1233123000232220-3302330022130303"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3130231012230100-0310001331133100-3201332131122221-3331221002321123-0123212320032103-1113102131310102-0030203231111010-1302311030011321"></a>

### Direct properties for `cors_policy`

<a id="canonical-0223330203231202-2021321032120022-2221320312012322-1031032312212213-3311103312331133-2000230132332131-2101131202101032-1130200330130121"></a>

#### `cors_policy.allow_credentials` property

Type: `"bool"`. Computed.

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

<a id="canonical-3031003211322112-0032211212213031-2011003122012123-0310233032222203-0102123202210212-1021320013313300-0220121230031023-1131100313212330"></a>

<a id="canonical-1021132023222033-1131122031202003-1221220103322010-2110232211311023-2332220112211222-3022232023031020-3223123231033301-0002203301021201"></a>

#### `cors_policy.allow_headers` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2132113121131012-1321312202120131-3203021313201010-0313001030100131-0332020223032320-0212103302120310-0223213230332330-2333331011121123"></a>

<a id="canonical-1231012232010010-0103122033111133-3200310331001310-1231122112111020-1223232320003233-2122320113211322-3230021203332332-1013122101233230"></a>

#### `cors_policy.allow_methods` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2133303301212112-2232030300112021-2310012122220203-3003311300123323-2131122121233113-2302002310101330-3100120001100323-1002320322301030"></a>

<a id="canonical-1332030123320033-0033011100001013-1310200222223032-1003302112212112-1111113203131311-2110022333313220-3033232112201332-1132023201213023"></a>

#### `cors_policy.allow_origin` property

Type: `["list", "string"]`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2023230203022130-0231000102103122-2202311211123102-3123212322301023-2121230332032302-3032122032011210-1003220122200210-3131113302203211"></a>

<a id="canonical-3020312323002200-3312213322332200-2101222203031102-0130330113100010-0223021231103030-0212133130032232-1233132323100103-3003220221322223"></a>

#### `cors_policy.allow_origin_regex` property

Type: `["list", "string"]`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2023310201233322-2033122032202023-0033011131221022-2211020110323332-3213102223310220-3031201303311322-3032313310103133-2111331311033332"></a>

<a id="canonical-1313021200201220-3012031331322023-0222222033032131-1231103220001012-0100313232213101-1321132033011231-2300003232232132-0203133332013121"></a>

#### `cors_policy.disabled` property

Type: `"bool"`. Computed.

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

<a id="canonical-2121020200230122-1100310020220123-2001103213103133-1123203021331303-1010102123233233-0131121311012032-0132202201112103-0120232300102030"></a>

<a id="canonical-3303222210233222-0020313220232312-3032210033311102-1200030100013223-2103313311201322-0132331321013202-0033101302120123-0120032233113220"></a>

#### `cors_policy.expose_headers` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0333020330011110-2203132300323331-3332113121031020-0121333101321220-1000000202211202-1323000111133333-1211301031012120-2021221033231222"></a>

<a id="canonical-0312231012001013-1111223301302300-2100312200320302-0333112312323033-1021311032203220-2331002232202311-1012303121230032-1311223021223012"></a>

#### `cors_policy.maximum_age` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1311323223301232-0300332102000101-2321001332311130-2131032223022213-0212021121211332-3030332100231313-0332301223110011-3303032003211230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `csrf_policy` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- csrf_policy

<a id="canonical-2032002311312013-0000112220310102-0133332013012133-3223011322000201-3332000003333122-2201020233330002-1021230101200311-2023231002033131"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2130022231202333-2302120123320333-1221321023320011-0013100310023120-1012003220131221-0100322203120210-2012002200113323-3033002301320301"></a>

### Direct properties for `csrf_policy`

- [all_load_balancer_domains](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0230332033033301-2302102231330312-1013001111131131-1220000000133201-0310232002102301-3021123321002131-3012132133211211-0132032112003312): complete subsection reference.

- [custom_domain_list](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0231310222213201-1322012223002113-2022030123321212-2022110110303222-3023302302320311-3101102100333021-2202201101301302-2122110111112222): complete subsection reference.

- [disabled](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-3211103022323100-2122030211321332-1301332110022133-2221223112330113-1012232203020023-3121320223023320-0020020122320321-3232033311201012): complete subsection reference.

<a id="canonical-0230332033033301-2302102231330312-1013001111131131-1220000000133201-0310232002102301-3021123321002131-3012132133211211-0132032112003312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `csrf_policy.all_load_balancer_domains` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [csrf_policy](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1311323223301232-0300332102000101-2321001332311130-2131032223022213-0212021121211332-3030332100231313-0332301223110011-3303032003211230)
- csrf_policy.all_load_balancer_domains

<a id="canonical-1101231213132123-1303123101020131-1100132110133101-2230010302230323-1321200333121300-1220131023302010-1113131212321033-2012233011322223"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231310222213201-1322012223002113-2022030123321212-2022110110303222-3023302302320311-3101102100333021-2202201101301302-2122110111112222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `csrf_policy.custom_domain_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [csrf_policy](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1311323223301232-0300332102000101-2321001332311130-2131032223022213-0212021121211332-3030332100231313-0332301223110011-3303032003211230)
- csrf_policy.custom_domain_list

<a id="canonical-3320330002320003-3022223223211231-3130233032022313-2211021300201333-2220121002002302-1023030230113133-2303022023112332-3220311023223231"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2301302032103212-3001300221120101-1200323113110231-1031011231333213-0311203012331310-3320213111331001-2102302111203213-2213331201011310"></a>

### Direct properties for `csrf_policy.custom_domain_list`

<a id="canonical-1313210322132123-1213310333320321-2312030223111333-1233223202201331-1212131303011313-2233233313020322-0303221130323010-1031323021211020"></a>

#### `csrf_policy.custom_domain_list.domains` property

Type: `["list", "string"]`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3211103022323100-2122030211321332-1301332110022133-2221223112330113-1012232203020023-3121320223023320-0020020122320321-3232033311201012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `csrf_policy.disabled` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [csrf_policy](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1311323223301232-0300332102000101-2321001332311130-2131032223022213-0212021121211332-3030332100231313-0332301223110011-3303032003211230)
- csrf_policy.disabled

<a id="canonical-2230101132002233-3022231022320032-1313031010233131-0123020221233113-3110332303102323-3201233000213231-3202131312212033-3112032211230031"></a>

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

<a id="canonical-0012310100301121-0211013102110231-1333303221223123-0323131031032113-2122320213230021-2211113110211001-3231311331000132-0313231022331320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_cache_rule` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- custom_cache_rule

<a id="canonical-2002233220302302-0011232132023010-3313233103213113-3232311202100200-1100103302200231-0121010123331303-0330021213321332-3230021313333231"></a>

Type: `"single"`. Computed.

Custom Cache Rules. Caching policies for CDN.

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

<a id="canonical-1112213321123103-1211131331221300-3120302110332320-3301112121320122-2130322022231201-3201011322211330-1131301321333220-1101213133330110"></a>

### Direct properties for `custom_cache_rule`

- [cdn_cache_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2021332313100111-3332313313321312-2321131102001122-1001302322303112-2101020131320022-3331233102032232-3113232311101203-1012021331003203): complete subsection reference.

<a id="canonical-2021332313100111-3332313313321312-2321131102001122-1001302322303112-2101020131320022-3331233102032232-3113232311101203-1012021331003203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_cache_rule.cdn_cache_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [custom_cache_rule](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0012310100301121-0211013102110231-1333303221223123-0323131031032113-2122320213230021-2211113110211001-3231311331000132-0313231022331320)
- custom_cache_rule.cdn_cache_rules

<a id="canonical-2231330000201120-0333233000033323-2101303000130312-3113212111101000-1312221232020113-2320132320333223-0000203120001021-1202331011113232"></a>

Type: `"list"`. Computed.

Reference to CDN Cache Rule configuration object.

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

<a id="canonical-3232121310203103-0032101132302231-2311321210110112-3303211003312031-0003020003001302-0103111231003110-1121032120211201-3022122200102230"></a>

### Direct properties for `custom_cache_rule.cdn_cache_rules`

<a id="canonical-3012031112002102-2321200210221101-3232200212321330-0201013001233301-0311202130313113-1110310032322100-2210212203132330-3112311111221120"></a>

#### `custom_cache_rule.cdn_cache_rules.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1131200202230202-0100303303003331-3322121221320202-0132103330322030-1310120012013012-2221132121030201-1201202220031101-3100200131311102"></a>

<a id="canonical-2222100011110000-0300131002111300-3101333232212302-3220111011131130-0310000022113022-1311200110033212-2333222112123202-2112002013032003"></a>

#### `custom_cache_rule.cdn_cache_rules.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3131111232303033-1322230232200330-3321212113311202-3233310100203201-3110230000003202-2300231202311023-3110323101302122-0310001313331332"></a>

<a id="canonical-3321010100233320-3011330110000110-3131011313302030-2031010201122031-3313031231313310-2231221301011222-0111102101311133-1202211332120013"></a>

#### `custom_cache_rule.cdn_cache_rules.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `data_guard_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- data_guard_rules

<a id="canonical-2210301100023211-2133121113131030-2131333130020032-3223120103202302-3120123321221020-1211110320330323-3333013100121102-2131333113033333"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0000202221123133-2132021323122021-3331111211011200-1330102113233300-0100310113012231-2231110303303022-3020111020031302-1121012212001313"></a>

### Direct properties for `data_guard_rules`

- [any_domain](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-3232321121012031-2231320120122001-3030021103212311-2032210323100132-1111011312203313-2310300031301000-1331012210022312-2123330020301020): complete subsection reference.

- [apply_data_guard](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0231200112332010-3011103033131301-2210013233011300-3000321113103121-2331221221301202-3331323321321222-1210111231110233-0301113300110221): complete subsection reference.

<a id="canonical-2131002010112210-3312232002330103-2123120033320103-2100300320133203-0132111213202112-3103100000102010-3201203021110221-2032010110021212"></a>

<a id="canonical-3303233122321300-3221321301030303-3031301133231012-3312302322001212-2113111322320322-2010103221033313-1230321110310012-0012223003001120"></a>

#### `data_guard_rules.exact_value` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [metadata](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1231031032112033-0132301003311100-1203321332303302-1111112322000023-3020022203332112-2313322030310210-3300021100230202-0000221300200101): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0300121020333212-3003033232111130-3313010331132123-0012103220020020-2213032013310100-3021232333000320-2222331003032323-1233120322311302): complete subsection reference.

- [skip_data_guard](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0032133120001001-3220200103000200-1332011021110003-3003013102321231-1203331201010031-3012130022101031-0211222113023112-2002221331313203): complete subsection reference.

<a id="canonical-3310102131110230-3103332021130221-0213130001133322-3001300333333320-0130001032020203-1323312102012000-1123113322301323-3211010221230111"></a>

<a id="canonical-1132230331203103-3223213001320323-1211121322101203-0232333211023300-1032233300112201-3303123221032313-3202222212213020-0223323113031031"></a>

#### `data_guard_rules.suffix_value` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3232321121012031-2231320120122001-3030021103212311-2032210323100132-1111011312203313-2310300031301000-1331012210022312-2123330020301020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `data_guard_rules.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211)
- data_guard_rules.any_domain

<a id="canonical-1021201221301331-3101302301113301-0233210001220231-1010112322202003-1010122211132333-3332310212230222-2313332232032213-1000202301321132"></a>

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

<a id="canonical-0231200112332010-3011103033131301-2210013233011300-3000321113103121-2331221221301202-3331323321321222-1210111231110233-0301113300110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `data_guard_rules.apply_data_guard` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211)
- data_guard_rules.apply_data_guard

<a id="canonical-3322210121033322-0133102310331223-2003323123320313-1111001133031121-1130330231232013-3223000103022203-2022102200233233-1111001033023132"></a>

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

<a id="canonical-1231031032112033-0132301003311100-1203321332303302-1111112322000023-3020022203332112-2313322030310210-3300021100230202-0000221300200101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `data_guard_rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211)
- data_guard_rules.metadata

<a id="canonical-1201022100012001-1312113011023202-0123221031103230-0000312311123202-1030133023303012-3333311221210200-1332212211023220-2230331313032310"></a>

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

<a id="canonical-2212210020012012-2313312213203021-0011231112112202-2011223222200023-0011021030021033-0210330330231330-2222000012101020-2300331002211102"></a>

### Direct properties for `data_guard_rules.metadata`

<a id="canonical-0300133223200023-0003121030132033-2133101012030132-0311013231232121-1321032102202302-2030011230232013-0200213233230232-0011212002213203"></a>

#### `data_guard_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3103302300130121-0213132133321330-2332311021031310-1121213122130012-0021022023031213-0003033222131102-3033310301111123-2230330102102230"></a>

<a id="canonical-0013213300303301-3320212300233131-3200222001131110-1010322132013001-2230313201311013-3322232000210202-3130001130031113-3102332131203211"></a>

#### `data_guard_rules.metadata.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0300121020333212-3003033232111130-3313010331132123-0012103220020020-2213032013310100-3021232333000320-2222331003032323-1233120322311302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `data_guard_rules.path` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211)
- data_guard_rules.path

<a id="canonical-2000230321020020-0101102023022022-3222012333230120-0300030233002202-2033123211001103-0123202113032003-2013100130311210-2300221100311220"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1301210023222202-0133120122113201-3133031132300023-0201311212022233-1001331233231201-0202232221003233-3313203223203211-0122031121123320"></a>

### Direct properties for `data_guard_rules.path`

<a id="canonical-3100330331202100-1013111222330012-1121011021102231-3302201310311330-1232002203130211-0003003221021212-2222122113322121-0213302023030120"></a>

#### `data_guard_rules.path.path` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3121133113013112-0032000302102033-3120330223011303-1130201302003333-0321021101301032-3003011313013032-1312221221202212-3233313231201120"></a>

<a id="canonical-2311020301212303-3213333030200303-1322300002330203-1011032030112210-2033201322213312-1301000223223233-1022321033321201-1120101202101321"></a>

#### `data_guard_rules.path.prefix` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0133002112202231-0201011030001121-3021321013301332-3030123233000020-2212233321323021-2001132030230330-1200130010301323-0320003321220101"></a>

<a id="canonical-1031110003032032-0122000122110113-2021200223001321-2130231201021332-2323313322211030-3100232312233123-0022333110210202-1332322310312031"></a>

#### `data_guard_rules.path.regex` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0032133120001001-3220200103000200-1332011021110003-3003013102321231-1203331201010031-3012130022101031-0211222113023112-2002221331313203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `data_guard_rules.skip_data_guard` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211)
- data_guard_rules.skip_data_guard

<a id="canonical-1000132113110023-0020112102211122-1332013010220002-1023222230321101-0131013331001331-0331022002333200-0120133020213031-0200302011220310"></a>

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

<a id="canonical-0203332321031033-3130203001012232-3010332203030033-2133312223332032-0321102121221133-1312222023101103-1322331131210123-2310301103323130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- ddos_mitigation_rules

<a id="canonical-0103311221320210-0333301313222313-0112333011330112-3233332011303303-3230120233312321-3102020213020203-0120223103312123-0313133113233022"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0322333101001033-0233301002200330-1321303001311123-1100213300221110-2312102011230001-0123230330033120-3132000221001123-3211201113313331"></a>

### Direct properties for `ddos_mitigation_rules`

- [block](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1110203213331010-0211033320130001-1123132032023132-0123310200112200-1100203002231013-0131310231302113-0113333013330021-3331001323230210): complete subsection reference.

- [ddos_client_source](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0231322231110013-0203001033311331-1033103211213020-3310322033003023-2300013323303032-1031133323221002-2132310332322021-1201201201201311): complete subsection reference.

<a id="canonical-1110301233212010-2232323110231201-2310312321102321-0203102310330012-3012333002030011-1311300200010213-0121230233220332-1210131111123111"></a>

<a id="canonical-1313131113120003-2003032011022213-1212010113132002-2330322133101211-2202210331120323-1303111120031303-1301321202300311-1100021001302222"></a>

#### `ddos_mitigation_rules.expiration_timestamp` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [ip_prefix_list](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1023131012231111-2223333211011103-0230132032323301-2011103010300003-3300112112110213-0210102133231210-1003321200100123-3202201301123130): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2100220232201032-0220103110012110-0210030130100202-3031003111313030-0111111011112011-2100122010300210-2331230311310301-0313121133211311): complete subsection reference.

<a id="canonical-1110203213331010-0211033320130001-1123132032023132-0123310200112200-1100203002231013-0131310231302113-0113333013330021-3331001323230210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0203332321031033-3130203001012232-3010332203030033-2133312223332032-0321102121221133-1312222023101103-1322331131210123-2310301103323130)
- ddos_mitigation_rules.block

<a id="canonical-1110003002100100-0313313113131312-3031332303001312-0000312323232110-1210110313222101-3110322201303322-2211123033020032-0232110101210123"></a>

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

<a id="canonical-0231322231110013-0203001033311331-1033103211213020-3310322033003023-2300013323303032-1031133323221002-2132310332322021-1201201201201311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.ddos_client_source` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0203332321031033-3130203001012232-3010332203030033-2133312223332032-0321102121221133-1312222023101103-1322331131210123-2310301103323130)
- ddos_mitigation_rules.ddos_client_source

<a id="canonical-3013102123312033-1232230212221311-3302312213302203-1321002323333212-1120022223300212-1113213321101230-0003100223133120-3112331102213231"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3113013010121222-2111033121302101-1313322303332313-3312113013112201-2321300330003103-1113222121202231-3032131101330310-2030220132000003"></a>

### Direct properties for `ddos_mitigation_rules.ddos_client_source`

- [asn_list](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0312133021030231-1010312022321220-1101133201222222-0230303000211031-2201120030223202-2303310220132322-3231321133011103-2200221210213003): complete subsection reference.

<a id="canonical-1301002312312020-0002111011231001-1133230230302022-2020323132103032-1303323201211203-0310132001330321-3201103220122223-1110221001123320"></a>

<a id="canonical-3233123110110121-1212303001331301-3032311200112230-3210313030102332-0232210122010301-1311030020130200-1212000311232200-2022303022322301"></a>

#### `ddos_mitigation_rules.ddos_client_source.country_list` property

Type: `["list", "string"]`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [ja4_tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1010020312100303-2110330230121312-2302333202321113-3003032012103131-2011020302013313-2321131022313002-0021311001101200-3122032201311210): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1000233103010133-3310202120331122-1213031103300101-0302012100131232-2000021132101101-1233032122320200-3013202323001112-3301220130130200): complete subsection reference.

<a id="canonical-0312133021030231-1010312022321220-1101133201222222-0230303000211031-2201120030223202-2303310220132322-3231321133011103-2200221210213003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.ddos_client_source.asn_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0203332321031033-3130203001012232-3010332203030033-2133312223332032-0321102121221133-1312222023101103-1322331131210123-2310301103323130)
- [ddos_mitigation_rules.ddos_client_source](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0231322231110013-0203001033311331-1033103211213020-3310322033003023-2300013323303032-1031133323221002-2132310332322021-1201201201201311)
- ddos_mitigation_rules.ddos_client_source.asn_list

<a id="canonical-1121003331302123-1020111102313211-2001102323133001-1211313011232300-2322223032201200-3122011103210033-2313301020022022-3012230223120310"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1011033213201211-0220232111301000-1011231012032030-3211011333223023-2232313211313002-3131313011123112-1303230003033223-1202300003200002"></a>

### Direct properties for `ddos_mitigation_rules.ddos_client_source.asn_list`

<a id="canonical-3221123111112110-1020023010022020-1110111023121311-2211101323003120-1223102333331310-1133101310333222-3330310003031020-2332133311132101"></a>

#### `ddos_mitigation_rules.ddos_client_source.asn_list.as_numbers` property

Type: `["list", "number"]`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1010020312100303-2110330230121312-2302333202321113-3003032012103131-2011020302013313-2321131022313002-0021311001101200-3122032201311210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0203332321031033-3130203001012232-3010332203030033-2133312223332032-0321102121221133-1312222023101103-1322331131210123-2310301103323130)
- [ddos_mitigation_rules.ddos_client_source](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0231322231110013-0203001033311331-1033103211213020-3310322033003023-2300013323303032-1031133323221002-2132310332322021-1201201201201311)
- ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher

<a id="canonical-3011120122000003-2022332233203321-3011000101032003-1211111002213102-0210220201123302-1201113121210232-1230012031103020-0100231323322020"></a>

Type: `"single"`. Computed.

An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

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

<a id="canonical-3223232112232222-1002320132031032-0000112223022220-2022223133203232-1310211021031021-2130030222320120-1133123201310232-3030320331003332"></a>

### Direct properties for `ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher`

<a id="canonical-3031300201103321-3121300332103222-2233112200133133-1010002100202101-2333012221321210-1202303321021331-3321021031010213-0210010211232103"></a>

#### `ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1000233103010133-3310202120331122-1213031103300101-0302012100131232-2000021132101101-1233032122320200-3013202323001112-3301220130130200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0203332321031033-3130203001012232-3010332203030033-2133312223332032-0321102121221133-1312222023101103-1322331131210123-2310301103323130)
- [ddos_mitigation_rules.ddos_client_source](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0231322231110013-0203001033311331-1033103211213020-3310322033003023-2300013323303032-1031133323221002-2132310332322021-1201201201201311)
- ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher

<a id="canonical-2133300310112230-1323212032221012-2102301023300021-1103031322233300-2220233210013121-2021301331130120-1132311113323303-0100132021231122"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2312202332332332-3033123232123311-2202123033311103-1201122122113211-2330132233322321-0203223310213123-0130323222131312-3222113002200322"></a>

### Direct properties for `ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher`

<a id="canonical-3320102033230331-0131010112330112-3033003112112002-1032323333120121-3201020132213111-3122211033221230-2103312313223120-0102223011212032"></a>

#### `ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher.classes` property

Type: `["list", "string"]`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0230231011233032-2131013230013233-1320231103312230-0000301230213330-1131023202011202-3213302233320121-1212303222000132-2122201030022133"></a>

<a id="canonical-3213101223002301-3032012022200203-3311321121233301-0023310003111313-1032103003201000-1122031121323103-3310013032022210-2103320310323211"></a>

#### `ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2130131230121120-0220323030131002-2213100323021011-0312030210012203-0212212132031211-3202012211221000-0211002330331332-0323130113111330"></a>

<a id="canonical-1302231100022210-1031032301211112-0032103221021012-3000332200211221-3301122022103212-2011010013030000-3003332312323200-1233302312121302"></a>

#### `ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1023131012231111-2223333211011103-0230132032323301-2011103010300003-3300112112110213-0210102133231210-1003321200100123-3202201301123130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0203332321031033-3130203001012232-3010332203030033-2133312223332032-0321102121221133-1312222023101103-1322331131210123-2310301103323130)
- ddos_mitigation_rules.ip_prefix_list

<a id="canonical-1331011201023012-1023123023203233-2232112011211300-0110110333111313-3001102001231221-1223213133121002-2210331200133030-2212332101212131"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1222201110303313-3232100002001222-1320301232031202-0333133020311001-3110010303202131-1223302300202113-2110301131133232-3222010203033001"></a>

### Direct properties for `ddos_mitigation_rules.ip_prefix_list`

<a id="canonical-2203300302230020-3010021113022133-2023012001131132-2121311302113111-0213031312300211-0122110332010100-1321311303121001-0303120122020102"></a>

#### `ddos_mitigation_rules.ip_prefix_list.invert_match` property

Type: `"bool"`. Computed.

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

<a id="canonical-3232130232110222-0111311100302211-1323301032200013-1003331012021032-0223220020022101-1320111311210211-2323002023133110-0120222133310001"></a>
