---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0331130132232332-1221332210122001-0202022230130032-3213300102111220-3121131032312213-3130320023112211-2330001132200120-0232322021202122"></a>

## Direct properties — rules / 031200002200 / 3

- [any_domain](resources--http_loadbalancer--reference--group-015.md#canonical-2013312033032311-1323201123212003-2012231301022113-1203300010212201-0331123303001112-2033030330233322-2301030130002010-2010202221221020): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-015.md#canonical-3121220330311231-0131113030213332-0230023132200210-2331300032230331-2201313210312311-0001100110221303-0032221321322221-1201001201330313): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-015.md#canonical-0100330120231232-3120222230110101-1221321002233010-0210010201301232-2332301333122111-2113133211302110-1213223000220300-0311111321313300): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-015.md#canonical-1100333130000013-0220213311330002-2122311302110313-0202212000023320-2113200210001211-1231003031022221-3103333213321232-2233130132132210): complete subsection reference.

<a id="canonical-1001321111210201-2220130232301120-1331131301300203-0122103301002330-3302102332310132-1002233032330302-1101011022021312-3012322133230102"></a>

## Next pages — rules / 031200002200 / 4

- [client_side_defense.policy.js_insertion_rules.rules.any_domain](resources--http_loadbalancer--reference--group-015.md#canonical-2013312033032311-1323201123212003-2012231301022113-1203300010212201-0331123303001112-2033030330233322-2301030130002010-2010202221221020)
- [client_side_defense.policy.js_insertion_rules.rules.domain](resources--http_loadbalancer--reference--group-015.md#canonical-3121220330311231-0131113030213332-0230023132200210-2331300032230331-2201313210312311-0001100110221303-0032221321322221-1201001201330313)
- [client_side_defense.policy.js_insertion_rules.rules.metadata](resources--http_loadbalancer--reference--group-015.md#canonical-0100330120231232-3120222230110101-1221321002233010-0210010201301232-2332301333122111-2113133211302110-1213223000220300-0311111321313300)
- [client_side_defense.policy.js_insertion_rules.rules.path](resources--http_loadbalancer--reference--group-015.md#canonical-1100333130000013-0220213311330002-2122311302110313-0202212000023320-2113200210001211-1231003031022221-3103333213321232-2233130132132210)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2013312033032311-1323201123212003-2012231301022113-1203300010212201-0331123303001112-2033030330233322-2301030130002010-2010202221221020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033020333232031-3101022303110232-1131131220133301-0102101212213313-2312112333301000-2220321320321123-3112023323032231-3133322320001302"></a>

## client_side_defense.policy.js_insertion_rules.rules.any_domain — any_domain / 203231213201 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-2003130231131001-2330111303132012-2002331330312302-1320031033110122-3001110220033102-2320021302020322-3323132133020121-0210202032110321)
- client_side_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-2230023003113022-3221211033121102-0320213102020330-1221233103130331-2132112103331033-3232323113230213-2313223132100231-3000232331321001"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
any_domain = {}
```

<a id="canonical-2313100010013120-2320022130100311-0023011302221203-3332000001222013-1121003222321203-2221111232322100-1333210012013321-3333311002230210"></a>

## Direct properties — any_domain / 203231213201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013121330222303-2221231220213012-3203103120220312-1102120122320000-3231201112300012-3033203020330102-0312022000330220-2011011032113032"></a>

## Next pages — any_domain / 203231213201 / 4

- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-2003130231131001-2330111303132012-2002331330312302-1320031033110122-3001110220033102-2320021302020322-3323132133020121-0210202032110321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3121220330311231-0131113030213332-0230023132200210-2331300032230331-2201313210312311-0001100110221303-0032221321322221-1201001201330313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312323202100131-1012113121222200-1320201130212332-3112103312112301-0001003000031230-2202213211103300-1121312322230130-2021210310003013"></a>

## client_side_defense.policy.js_insertion_rules.rules.domain — domain / 102232030321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-2003130231131001-2330111303132012-2002331330312302-1320031033110122-3001110220033102-2320021302020322-3323132133020121-0210202032110321)
- client_side_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-1302212002330313-1120020221313113-1121312313033030-0322112113000100-0010230201011011-2202300200233210-0120230031333330-2222233323100032"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-1100133223133301-1101303202303322-1320002020010321-1012010323020212-1321302210302220-3331022331303021-0101221000101330-1322121002230121"></a>

## Direct properties — domain / 102232030321 / 3

<a id="canonical-3222103102232010-1320202002130310-0222000302121200-3233321020031122-2211001030230133-0311311222013003-0030100320030201-0022311122113332"></a>

<a id="canonical-3023100032103223-3012012311030300-0111012323110031-2323331113021122-1123023023011312-3321201231321202-0211211032213233-0100100312023312"></a>

## exact_value property — domain / 102232030321 / 4

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-3302330132232201-2312200123000101-3231210333101310-0300302332100310-1303331203030001-3202313003331311-3130202301322233-1030220003201013"></a>

## regex_value property — domain / 102232030321 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-1200123220003031-1013323103031200-2320102210102103-3320133121222310-0331020132101322-2123322022012102-3112313322221000-0112213333310020"></a>

## suffix_value property — domain / 102232030321 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-3200213030211112-2033303031020200-1000100311003123-0201121022332311-3121222031203213-0212300320233230-2130031312033232-3202202203003121"></a>

## Next pages — domain / 102232030321 / 7

- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-2003130231131001-2330111303132012-2002331330312302-1320031033110122-3001110220033102-2320021302020322-3323132133020121-0210202032110321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0100330120231232-3120222230110101-1221321002233010-0210010201301232-2332301333122111-2113133211302110-1213223000220300-0311111321313300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200111112121000-3023112223320010-0311322333000101-0101132023222031-0301231001033311-3031310301210211-1111323220012231-0002301223203221"></a>

## client_side_defense.policy.js_insertion_rules.rules.metadata — metadata / 133202011122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-2003130231131001-2330111303132012-2002331330312302-1320031033110122-3001110220033102-2320021302020322-3323132133020121-0210202032110321)
- client_side_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-1202321312121132-1100031023112131-2031023123212231-3323203100202013-0002220002000113-0332333101031302-3001231213113111-1201221022323203"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-3010220031332133-0103023120221023-2000113232320033-1212233101323002-2322321202222103-3120011312102320-3100110013211122-2300033100331011"></a>

## Direct properties — metadata / 133202011122 / 3

<a id="canonical-2013321300013111-0202013012122210-0220333002332123-2020022011221133-0012132023213133-0122133213111013-0020322010021222-0121133120331230"></a>

<a id="canonical-1133303213000021-0211110121002122-2333100010320012-2212100030231202-3302130003313321-2201132012202011-0122030203300121-1201230321133210"></a>

## description_spec property — metadata / 133202011122 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1030003022321102-3222202031210101-2133102233310013-2010230012101303-2220022301102223-3201123220231101-2311112003102210-3110023311113030"></a>

<a id="canonical-0123012132320201-0023302211333120-3223231112313322-1133010320311022-2012221212320022-1302021101222123-0033321121233031-0031310031012300"></a>

## name property — metadata / 133202011122 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-3120011330210301-1212031030231113-2101322100131313-0313001110331111-1001101032330301-1200120201101332-1011222022331030-3132230020231131"></a>

## Next pages — metadata / 133202011122 / 6

- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-2003130231131001-2330111303132012-2002331330312302-1320031033110122-3001110220033102-2320021302020322-3323132133020121-0210202032110321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1100333130000013-0220213311330002-2122311302110313-0202212000023320-2113200210001211-1231003031022221-3103333213321232-2233130132132210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010300333121330-1322020313232000-0033013010312122-3023302130002121-3021110210112230-0031303220120333-2330200331110232-0330330100131201"></a>

## client_side_defense.policy.js_insertion_rules.rules.path — path / 321231331301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-2003130231131001-2330111303132012-2002331330312302-1320031033110122-3001110220033102-2320021302020322-3323132133020121-0210202032110321)
- client_side_defense.policy.js_insertion_rules.rules.path

<a id="canonical-3011323303023030-0311011033223133-2111331110133132-2310120031120210-0020300201303311-2232303110200120-2211332132230200-2332101201301231"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-1031022012032003-0030130032120110-0131320112121222-0011230022301233-1222210012323122-0233201022330132-2321120211303331-1103122123011122"></a>

## Direct properties — path / 321231331301 / 3

<a id="canonical-3020322312232310-1333032303233130-3313010121002311-2312311330110231-3033220032122312-1232223131020123-0321220121211232-0302001222131233"></a>

<a id="canonical-2312210311133300-0103023123113120-1333220321330011-2122321021230222-1330212222122232-3122100022333301-2012111230200223-1313230001233310"></a>

## path property — path / 321231331301 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1030230020312202-3303031122033310-1203012010132010-0023010021102311-0100330111121213-0330130012032203-2332332213321010-0112212123322032"></a>

## prefix property — path / 321231331301 / 5

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1101233323333102-2321110002110332-2230201030232331-3032311200233112-3320202322231331-1030201312230302-1130332320130033-1131023320222030"></a>

## regular expression property — path / 321231331301 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-3100212232120000-3332021123211230-3123123312303003-2210301132303312-1103012112310003-0223303010001123-2020122213103001-3132123013333220"></a>

## Next pages — path / 321231331301 / 7

- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-2003130231131001-2330111303132012-2002331330312302-1320031033110122-3001110220033102-2320021302020322-3323132133020121-0210202032110321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331103002210032-3331332220333133-1021201203322123-3001112010303020-1130132202303123-1003011103102201-1200332322130122-2300000003023311"></a>

## cookie_stickiness — cookie_stickiness / 202233021010 / 2

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

Upstream description:

Two types of cookie affinity:

&#8203;1. Passive. Takes a cookie that's present in the cookies header and hashes on its value.

&#8203;2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from
the client in its response to the client, based on the endpoint the request gets sent to. The client
then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure
these requests GET sent to the same endpoint. The cookie is generated by hashing the source and
destination ports and addresses so that multiple independent HTTP2 streams on the same connection
will independently receive the same cookie, even if they arrive simultaneously.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name"),
  validators.ConflictingObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingObjectAttributes("samesite_none",
    "samesite_strict")}
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
  "x-ves-oneof-field-httponly": "[\"add_httponly\",\"ignore_httponly\"]",
  "x-ves-oneof-field-samesite": "[\"ignore_samesite\",\"samesite_lax\",\"samesite_none\",\"samesite_strict\"]",
  "x-ves-oneof-field-secure": "[\"add_secure\",\"ignore_secure\"]"
}
```

OneOf alternatives in this subsection:

- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3022100130130323-2211020000021220-3132110021321012-3201303232001232-0113320030023322-3202112300132331-2333201101110023-1132101301213203)
- [least_active](resources--http_loadbalancer--reference--group-021.md#canonical-0012031220101203-1331010113102002-1112333232131023-3022102022010211-3120131001301000-1311230213201110-2003321233300023-2003020103022211)
- [random](resources--http_loadbalancer--reference--group-023.md#canonical-1021313003120223-3101211023223000-1333330100323102-3033202113022313-2111113220003123-1112002011233323-1312030112001012-1111020111100321)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3013210122112121-0122221322331320-0013232112123331-2012132320331222-1132231302310032-2101233212003331-1003333001210203-3312200031323332)
- [round_robin](resources--http_loadbalancer--reference--group-024.md#canonical-1120231012232110-3001211103003332-2110332213132110-2233201023123010-3331031002232323-2201313122330202-3321233220103133-3330032203222203)
- [source_ip_stickiness](resources--http_loadbalancer--reference--group-027.md#canonical-0202231310310202-3102203321100120-0022003303331033-1312113312302011-2331102102230311-0003123123020233-2303210033023103-2322033033310111)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
cookie_stickiness {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223121112012013-1010123312333132-0003103022323031-1131200230312221-3333221212111232-2000022030010021-1312331123033000-1302132231021311"></a>

## Direct properties — cookie_stickiness / 202233021010 / 3

- [add_httponly](resources--http_loadbalancer--reference--group-015.md#canonical-2030131120202023-2110332201010210-0223310001320000-2313012310033112-3132121323021000-3323032103012233-3203112210212112-2330330020210333): complete subsection reference.

- [add_secure](resources--http_loadbalancer--reference--group-015.md#canonical-2010011221312311-2201313103330320-0310313033132213-3230012221222320-0122233030233133-2213133110022103-2201003011023032-1310213320202133): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--reference--group-015.md#canonical-0332032201122022-2110113231101103-3030120121210231-2103232220102020-2003111103302322-2030020210312123-0023303001212200-2202330201201210): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--reference--group-015.md#canonical-0200002032300002-1032312120300023-2120012123133112-1332132122220013-1100120331301011-2012021333313032-1011120132021231-0112103013101330): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--reference--group-015.md#canonical-0033230303023301-1201320022202112-0010301331012220-0003300031000231-0223113312220220-1333122221222310-2022010220330230-3320002200012301): complete subsection reference.

<a id="canonical-3300101222320023-0123011223031301-0002233130003122-0011222000202120-0003322200232131-1112033012313123-1210220032131121-0332321200002211"></a>

<a id="canonical-3011133120212212-3023312330032301-2031320030200011-2321121210110113-2323320112002010-0203010103200312-0021101031133000-0221231210013223"></a>

## name property — cookie_stickiness / 202233021010 / 4

Type: `"string"`. Optional.

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

Upstream description:

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-1111012110032212-1311230132203002-0023000313320221-0002322322300111-3130301211033033-3211112310133302-0133002233310323-3020022111022023"></a>

## path property — cookie_stickiness / 202233021010 / 5

Type: `"string"`. Optional.

The name of the path for the cookie. If no path is specified here, no path will be set for the
cookie.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0133133201333223-3222321131130201-1111322232323310-2220023120032000-0200010032231110-3223120113121202-1222330221031123-1032332331321131"></a>

## TTL property — cookie_stickiness / 202233021010 / 6

Type: `"number"`. Optional.

If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is
present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.

Upstream description:

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

<a id="canonical-0320213122003312-1031231002122122-3332110223000230-3001322322111013-3022002020222201-1031002220102233-3312213112101330-1221010121333330"></a>

## Next pages — cookie_stickiness / 202233021010 / 7

- [cookie_stickiness.add_httponly](resources--http_loadbalancer--reference--group-015.md#canonical-2030131120202023-2110332201010210-0223310001320000-2313012310033112-3132121323021000-3323032103012233-3203112210212112-2330330020210333)
- [cookie_stickiness.add_secure](resources--http_loadbalancer--reference--group-015.md#canonical-2010011221312311-2201313103330320-0310313033132213-3230012221222320-0122233030233133-2213133110022103-2201003011023032-1310213320202133)
- [cookie_stickiness.ignore_httponly](resources--http_loadbalancer--reference--group-015.md#canonical-0332032201122022-2110113231101103-3030120121210231-2103232220102020-2003111103302322-2030020210312123-0023303001212200-2202330201201210)
- [cookie_stickiness.ignore_samesite](resources--http_loadbalancer--reference--group-015.md#canonical-0200002032300002-1032312120300023-2120012123133112-1332132122220013-1100120331301011-2012021333313032-1011120132021231-0112103013101330)
- [cookie_stickiness.ignore_secure](resources--http_loadbalancer--reference--group-015.md#canonical-0033230303023301-1201320022202112-0010301331012220-0003300031000231-0223113312220220-1333122221222310-2022010220330230-3320002200012301)
- [cookie_stickiness.samesite_lax](resources--http_loadbalancer--reference--group-015.md#canonical-1321001203233233-2321232003021022-0203011131012320-3203022123300223-3002101231020330-3302123321300303-2223213202333202-3330223113122132)
- [cookie_stickiness.samesite_none](resources--http_loadbalancer--reference--group-015.md#canonical-1132131112302000-1122301323303133-2010003132020223-3012212122132030-1202021222123020-0223313313123132-2120032300330312-2300013111133200)
- [cookie_stickiness.samesite_strict](resources--http_loadbalancer--reference--group-015.md#canonical-1011201020201132-0002101330103300-0003131200223133-2231001133012310-1230202310111221-1033332302022223-1232112022311122-1211320022123012)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2030131120202023-2110332201010210-0223310001320000-2313012310033112-3132121323021000-3323032103012233-3203112210212112-2330330020210333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012023003302011-2023103220211103-1210020111213013-2113031020320122-3322210310230220-2302110222032210-1102013330123221-0110113203223312"></a>

## cookie_stickiness.add_httponly — add_httponly / 320110112203 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- cookie_stickiness.add_httponly

<a id="canonical-3332020222303023-3101122213113110-3101302022101200-3211313023212131-0101322223001031-1302103311102112-1111010231011000-0130102102031213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add httponly.

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

Terraform syntax:

```terraform
add_httponly = {}
```

<a id="canonical-2030023000201100-1121220230131322-2321310213310110-3032103030003232-3212331323122302-0322123320211113-2020311103113031-2310301321030201"></a>

## Direct properties — add_httponly / 320110112203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103320200331302-0031333230013002-2013012321032220-1112022110130321-0111133223022201-1300002303231031-2220202313212111-1303121313122221"></a>

## Next pages — add_httponly / 320110112203 / 4

- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2010011221312311-2201313103330320-0310313033132213-3230012221222320-0122233030233133-2213133110022103-2201003011023032-1310213320202133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312212022211321-1222120221330023-0132011033123332-3133132003330212-2021030030101201-0132100330233112-2111023123012102-3331312333231001"></a>

## cookie_stickiness.add_secure — add_secure / 122233030332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- cookie_stickiness.add_secure

<a id="canonical-1021023322210122-2212100023331020-3320003213201312-3313333301030312-2321300213002230-0320221100320002-1131210332113111-2130132021033231"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
add_secure = {}
```

<a id="canonical-3021210211000100-2312111201130033-2122021003320301-0120203001112221-0031211130032310-2303331203302321-1021231323123120-1201210120012312"></a>

## Direct properties — add_secure / 122233030332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001101202132010-2121000101030300-3333032301133331-1021300103123112-1300123212032311-3000110133020123-1303113113222021-2230312133022200"></a>

## Next pages — add_secure / 122233030332 / 4

- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0332032201122022-2110113231101103-3030120121210231-2103232220102020-2003111103302322-2030020210312123-0023303001212200-2202330201201210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322021221321123-1321213023211101-2322221000303110-3203023023201333-3012223033111232-1331012323031011-0130223000011302-0001321231032112"></a>

## cookie_stickiness.ignore_httponly — ignore_httponly / 302121301230 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- cookie_stickiness.ignore_httponly

<a id="canonical-2310100203100202-3312233101332033-2220302131000131-3001030231310333-2030133330010330-0133001012211021-2002333221300230-3103113331111301"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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

Terraform syntax:

```terraform
ignore_httponly = {}
```

<a id="canonical-3003223223332111-3113020203000210-1033302003220101-0121113121331312-1310003222121211-3133032023332201-1000330012010311-3312021030133302"></a>

## Direct properties — ignore_httponly / 302121301230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311303000313233-0311001133003131-0013302013100001-2120211012123111-1313122111103310-1110320033121102-1031113102000313-2232002330212132"></a>

## Next pages — ignore_httponly / 302121301230 / 4

- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0200002032300002-1032312120300023-2120012123133112-1332132122220013-1100120331301011-2012021333313032-1011120132021231-0112103013101330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330010121230021-3231113203313013-3300101111021103-3131212121213021-2233103322331302-0000021133202330-1133120210021130-3102210300000200"></a>

## cookie_stickiness.ignore_samesite — ignore_samesite / 130230302031 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- cookie_stickiness.ignore_samesite

<a id="canonical-0203323210102231-3013121301101200-0111011032000101-2111331112302133-0120133022123312-0322013030003212-1101101013203001-2323032122302010"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
ignore_samesite = {}
```

<a id="canonical-3233010320201313-2001332310211130-0331100020221001-3211312223101312-3130333031032030-3131332011001231-3032302223321011-1222311301011230"></a>

## Direct properties — ignore_samesite / 130230302031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333121332303213-2023131223211320-2031220113311230-0132031002213011-2022321002003032-2321103132233000-2211030333031002-0332222112132001"></a>

## Next pages — ignore_samesite / 130230302031 / 4

- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0033230303023301-1201320022202112-0010301331012220-0003300031000231-0223113312220220-1333122221222310-2022010220330230-3320002200012301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133003202320230-1312010322112100-2112010010032203-1103232001231023-2330120323223321-3231012233032030-0212231130123112-2313301131033211"></a>

## cookie_stickiness.ignore_secure — ignore_secure / 203003000320 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- cookie_stickiness.ignore_secure

<a id="canonical-3031113220030213-0001030110210312-2303220332221002-0310220312302120-1321202010012202-0212221311222021-3021133232100211-0232033101122322"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
ignore_secure = {}
```

<a id="canonical-1223200021210321-2230333111201020-2002302223113333-2301311103022321-1302122103322121-1220222223322302-0120221322021132-3220023110023300"></a>

## Direct properties — ignore_secure / 203003000320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312323212303203-0230222131332103-1000003331110301-0133110001233310-3331033213113301-0032101101200333-3013230330230131-1121133332121201"></a>

## Next pages — ignore_secure / 203003000320 / 4

- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1321001203233233-2321232003021022-0203011131012320-3203022123300223-3002101231020330-3302123321300303-2223213202333202-3330223113122132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113121101222131-0330300111330131-3021022303022103-1113311212010220-1321303202313110-2130130221231313-0133212003301122-1321210333231210"></a>

## cookie_stickiness.samesite_lax — samesite_lax / 101031000211 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- cookie_stickiness.samesite_lax

<a id="canonical-2202230212121310-2231222211213332-2013231311021110-0011231231122120-1031320330112021-2032120232300010-2322013133231320-0312312213210100"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
samesite_lax = {}
```

<a id="canonical-3120222311002322-3333223030033121-3132003330130220-1210231333303111-1321030022232100-1301113122023132-3001222112322201-3220131132021320"></a>

## Direct properties — samesite_lax / 101031000211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2132232313332233-2131233322100232-3110213020103202-1130133032323303-2211033333101110-2211220211210211-0100230002301131-2031013021213120"></a>

## Next pages — samesite_lax / 101031000211 / 4

- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1132131112302000-1122301323303133-2010003132020223-3012212122132030-1202021222123020-0223313313123132-2120032300330312-2300013111133200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203212211320113-2313222202000121-0030020103110132-0032320321330322-2200121311012120-3133121331112203-2101023020133100-2122012131021202"></a>

## cookie_stickiness.samesite_none — samesite_none / 013213000313 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- cookie_stickiness.samesite_none

<a id="canonical-0211121232222013-0212303220333020-2133223230103001-0003210230221032-3203201023302230-3232131323030121-3302320132200111-0310222031110112"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
samesite_none = {}
```

<a id="canonical-3211030010133313-3331131020323002-2221103030230200-1213223002110213-1232102130301331-3013301302211132-2021013322103130-0023112331002122"></a>

## Direct properties — samesite_none / 013213000313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201111130320013-1211021333020113-1203310333102003-0211111233332133-0002130332023302-3003122203221023-0211210312132313-2311002111320120"></a>

## Next pages — samesite_none / 013213000313 / 4

- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1011201020201132-0002101330103300-0003131200223133-2231001133012310-1230202310111221-1033332302022223-1232112022311122-1211320022123012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322212012122313-3131020131331222-0023003303101310-1032300000121100-2120332323320031-1102000121323110-3101202331130220-3310200023100013"></a>

## cookie_stickiness.samesite_strict — samesite_strict / 133112100010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- cookie_stickiness.samesite_strict

<a id="canonical-1100233232032301-0220302303202113-1301122133022011-1312311010301012-1302212013211111-2211221032202203-0102011032110313-1122131220220033"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
samesite_strict = {}
```

<a id="canonical-3203123320313001-1333313201320203-2220022010322332-2101033000233303-3301310221112231-2112130130330032-1311020033102320-2313010203030000"></a>

## Direct properties — samesite_strict / 133112100010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0202223001012302-2202023021211310-3203212300033020-3103110100310113-3010122230021123-0013103131221010-2220131000211011-2023310200303003"></a>

## Next pages — samesite_strict / 133112100010 / 4

- [cookie_stickiness](resources--http_loadbalancer--reference--group-015.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2120230123301020-2223032123333110-0222303131313113-1302023032311223-0231200001201300-0311101003010332-1110132103312122-1021322003001231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332203130031002-2130211100312200-3332133122021002-0021233001323112-0303302302231330-2110012032300211-3330112232231322-3312010203123010"></a>

## cors_policy — cors_policy / 003320201003 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- cors_policy

<a id="canonical-2122220210330203-0220012133012121-0103130220122000-3300103303020003-3301132102100322-0231320012332001-0023320010110100-0013033001130130"></a>

Type: `"object"`. single nested block, Optional.

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence. An example of an Cross origin HTTP request GET
/resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel macOS
X 10.5..

Upstream description:

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

<a id="canonical-1333121210323323-2221221101122101-2001200001100333-3300123211002112-1232023311223020-1223200121301221-1023131131032301-2232010133321023"></a>

## Direct properties — cors_policy / 003320201003 / 3

<a id="canonical-2103312020020203-3333220012221320-1011221213101322-1031320323001031-2221001211313331-2103112130002123-3220212330231331-0323110331000103"></a>

<a id="canonical-2101123202023011-2032330113012300-2200011031233332-3031033311020013-1201210210111333-3130232131011303-1300111232333022-1121000121013232"></a>

## allow_credentials property — cors_policy / 003320201003 / 4

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

<a id="canonical-2212330011323101-2221131213100110-2320100123110000-0031323302210003-2303001021232130-1023300032101310-3322211111011012-0331320003320330"></a>

## allow_headers property — cors_policy / 003320201003 / 5

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

<a id="canonical-1213231131101311-3111131223122220-1031103301321112-0230031132001023-0302002220332310-3120332313023132-2310132322112230-2111021330110232"></a>

<a id="canonical-2131212212233212-3302213020321010-2020323132331233-0032130232331002-2111023113131231-1221213113311133-3012033011131333-3022032233110230"></a>

## allow_methods property — cors_policy / 003320201003 / 6

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
    "ves.io.schema.rules.string.http_valid_methods": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  }
}
```

<a id="canonical-0100212323112021-0331311000310311-3003221003110031-2222320322213030-1111332322103002-3022112113223101-3203112031100030-3003203131301023"></a>

<a id="canonical-3231110110230303-3033022032230023-0312222011120021-1133023212110302-0121303202021121-1120030332210203-1222030333230123-2303210202302323"></a>

## allow_origin property — cors_policy / 003320201003 / 7

Type: `["list", "string"]`. Optional.

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

Upstream description:

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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

<a id="canonical-3303201011133020-1212003330221023-1221321030310312-3003313130323123-1303110110310201-1330010220113011-0322222011000232-1221130003331233"></a>

## allow_origin_regex property — cors_policy / 003320201003 / 8

Type: `["list", "string"]`. Optional.

Specifies regular expression patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

Upstream description:

Specifies regular expression patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-3131013022130113-0101003101310112-1222103310011333-3022211221112031-1203301233333222-0203001012322203-0232220300321121-3133322013010003"></a>

## disabled property — cors_policy / 003320201003 / 9

Type: `"bool"`. Optional.

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

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

<a id="canonical-0020000031220113-3303332012321230-3002322123132331-1000213003232001-2222001213003001-2220212223020311-2113323210233011-2021320303011322"></a>

## expose_headers property — cors_policy / 003320201003 / 10

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

<a id="canonical-1233333301122231-2102032220333021-2331230323103131-2223031130101201-2033112320233330-0301011233031231-0113201022102203-0213111110032200"></a>

<a id="canonical-3002012200101031-3300321011231220-3212132000213130-3122301022330220-3231113211311201-0022200123103100-0112121331013302-0133200310123020"></a>

## maximum_age property — cors_policy / 003320201003 / 11

Type: `"number"`. Optional.

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours).

Upstream description:

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(-1, 86400),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1303302003330333-3313232223131323-0210013323101331-0131111223022220-3330002303113300-3121020100121332-3311133323010231-0111221303310103"></a>

## Next pages — cors_policy / 003320201003 / 12

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3021222013003320-3230102012010221-2203233323210232-0111132233210023-1332310233100013-1310012310002212-3333112112231001-0212212331211311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202123002032021-3301221123313032-3021131232112021-3201201121030332-3221232120132013-2130022021300021-2132131102112333-0020303020030311"></a>

## csrf_policy — csrf_policy / 330101200112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- csrf_policy

<a id="canonical-3110111223222131-2221222112302020-0022332210201221-2133220221231233-0223031112211323-1013213130211301-1213001220101100-2323003032300302"></a>

Type: `"object"`. single nested block, Optional.

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host. 1. The origin that caused the user
agent..

Upstream description:

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_load_balancer_domains",
    "custom_domain_list"),
  validators.ConflictingObjectAttributes("all_load_balancer_domains",
    "disabled"),
  validators.ConflictingObjectAttributes("custom_domain_list",
    "disabled")}
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
  "x-ves-oneof-field-allowed_domains": "[\"all_load_balancer_domains\",\"custom_domain_list\",\"disabled\"]"
}
```

Terraform syntax:

```terraform
csrf_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203002313010222-1323221222003111-1211013011321030-2000021322310133-2310012013330233-2223012112333211-3011333230012012-2210201330010313"></a>

## Direct properties — csrf_policy / 330101200112 / 3

- [all_load_balancer_domains](resources--http_loadbalancer--reference--group-015.md#canonical-1100011121020001-1001321330300312-0203202022123003-2132001123322332-3232332300111011-2033132102233311-2033221020111023-0203100123121112): complete subsection reference.

- [custom_domain_list](resources--http_loadbalancer--reference--group-015.md#canonical-0222321213110001-0022132132020121-1030321303331231-2211122103230121-0011220321233023-2001132331233002-3331332312033013-3333011110233202): complete subsection reference.

- [disabled](resources--http_loadbalancer--reference--group-015.md#canonical-2302111201210301-0223031011201123-1110101312012031-3111331233023213-3200312032033203-1130130233230303-0303112233300100-2300331101013131): complete subsection reference.

<a id="canonical-1320330020120201-0101103220121133-0131123131001030-3022221010310001-1201203000333131-3221031000201213-3213021001023302-1202002120020310"></a>

## Next pages — csrf_policy / 330101200112 / 4

- [csrf_policy.all_load_balancer_domains](resources--http_loadbalancer--reference--group-015.md#canonical-1100011121020001-1001321330300312-0203202022123003-2132001123322332-3232332300111011-2033132102233311-2033221020111023-0203100123121112)
- [csrf_policy.custom_domain_list](resources--http_loadbalancer--reference--group-015.md#canonical-0222321213110001-0022132132020121-1030321303331231-2211122103230121-0011220321233023-2001132331233002-3331332312033013-3333011110233202)
- [csrf_policy.disabled](resources--http_loadbalancer--reference--group-015.md#canonical-2302111201210301-0223031011201123-1110101312012031-3111331233023213-3200312032033203-1130130233230303-0303112233300100-2300331101013131)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1100011121020001-1001321330300312-0203202022123003-2132001123322332-3232332300111011-2033132102233311-2033221020111023-0203100123121112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321231130123013-3122232123100021-3013030030020220-1121201021231002-1210212023303232-0032210021023201-2131320133120100-1231130032333033"></a>

## csrf_policy.all_load_balancer_domains — all_load_balancer_domains / 021001102323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [csrf_policy](resources--http_loadbalancer--reference--group-015.md#canonical-3021222013003320-3230102012010221-2203233323210232-0111132233210023-1332310233100013-1310012310002212-3333112112231001-0212212331211311)
- csrf_policy.all_load_balancer_domains

<a id="canonical-0101131132222003-1210021220000330-0211022303000021-2311310333230110-0220122301110312-1110122320130322-1001002002030011-2003102030220023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all load balancer domains.

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

Terraform syntax:

```terraform
all_load_balancer_domains = {}
```

<a id="canonical-1013311210221222-0221003132313012-1032003033322202-0022030111032222-1131113323132211-0121331232230101-2203222020020012-3202020123122120"></a>

## Direct properties — all_load_balancer_domains / 021001102323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2013221332021302-0131330023332233-1200212203222232-0320300212330302-0132212130110202-3000032131210133-3321233222030211-1013101012300131"></a>

## Next pages — all_load_balancer_domains / 021001102323 / 4

- [csrf_policy](resources--http_loadbalancer--reference--group-015.md#canonical-3021222013003320-3230102012010221-2203233323210232-0111132233210023-1332310233100013-1310012310002212-3333112112231001-0212212331211311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0222321213110001-0022132132020121-1030321303331231-2211122103230121-0011220321233023-2001132331233002-3331332312033013-3333011110233202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333020013323211-3103100311323310-0332001021032021-3321323312022333-2201202021331013-2023231002313032-0230032030113130-1231121321000312"></a>

## csrf_policy.custom_domain_list — custom_domain_list / 322233103112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [csrf_policy](resources--http_loadbalancer--reference--group-015.md#canonical-3021222013003320-3230102012010221-2203233323210232-0111132233210023-1332310233100013-1310012310002212-3333112112231001-0212212331211311)
- csrf_policy.custom_domain_list

<a id="canonical-0123221322012310-2133221301301120-2030133002030332-2321011120233111-0102211222310212-2121102102303221-1323012332330300-2113133132212130"></a>

Type: `"object"`. single nested block, Optional.

List of domain names used for Host header matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains")}
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
custom_domain_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2131122302212023-3020012202212133-3333200210120330-2212020011010213-1330130013033202-0113023323022122-2320112001221233-2301313201033010"></a>

## Direct properties — custom_domain_list / 322233103112 / 3

<a id="canonical-0333211032010300-2003230311122201-0302230221123211-1122322033333333-2212333130120321-2203121232230320-1201233322302330-2110201102031101"></a>

<a id="canonical-0032320301223100-1121023311311322-1132212220311300-0303000322321231-3202113112011020-3123321313122131-2131031221130120-0000003322103002"></a>

## domains property — custom_domain_list / 322233103112 / 4

Type: `["list", "string"]`. Optional.

List of domain names that will be matched to loadbalancer. These domains are not used for SNI match.
Wildcard names are supported in the suffix or prefix form.

Upstream description:

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

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

<a id="canonical-2002312310210322-3202223101131321-1020122102022230-0103112020200322-2303212021333221-1323311102300110-1011230000330010-2300231231011310"></a>

## Next pages — custom_domain_list / 322233103112 / 5

- [csrf_policy](resources--http_loadbalancer--reference--group-015.md#canonical-3021222013003320-3230102012010221-2203233323210232-0111132233210023-1332310233100013-1310012310002212-3333112112231001-0212212331211311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2302111201210301-0223031011201123-1110101312012031-3111331233023213-3200312032033203-1130130233230303-0303112233300100-2300331101013131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231321132001312-0333220011330002-3220101130301203-3302003310002023-2120110032220112-0113102001310011-1011323020302000-0021230211130001"></a>

## csrf_policy.disabled — disabled / 023133103323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [csrf_policy](resources--http_loadbalancer--reference--group-015.md#canonical-3021222013003320-3230102012010221-2203233323210232-0111132233210023-1332310233100013-1310012310002212-3333112112231001-0212212331211311)
- csrf_policy.disabled

<a id="canonical-2023310232303231-0332121311230033-0220311322311233-2021130230233211-1003221002113303-3231330033121133-3100120331123131-2013231300212102"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
disabled = {}
```

<a id="canonical-3133121132130132-0020011301221000-3010321103301011-2133000002133230-3130021101001301-2313110301010323-2300101100010201-3120133123103223"></a>

## Direct properties — disabled / 023133103323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000022203333310-0212101010212021-0102223333130322-2120111301121213-2301130320200320-1212312232013201-1121113020133230-3322233303001330"></a>

## Next pages — disabled / 023133103323 / 4

- [csrf_policy](resources--http_loadbalancer--reference--group-015.md#canonical-3021222013003320-3230102012010221-2203233323210232-0111132233210023-1332310233100013-1310012310002212-3333112112231001-0212212331211311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0310100223013020-2111222133111333-2201233312231132-0102103013122132-0111220303113031-1122123223301021-1030212023023130-0331013321302221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233333013303032-3003212320330031-0221013312102010-1301333011112103-0312232010113133-1000221321021011-2102003320210022-2320310302121112"></a>

## data_guard_rules — data_guard_rules / 001230221202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- data_guard_rules

<a id="canonical-2022010320103333-3032231310103322-1300320033111030-3203121310020023-2001031202123103-1030032000023310-2212303022223020-2211202010031032"></a>

Type: `"object"`. list nested block, Optional.

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*).

Upstream description:

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*). Note: App Firewall should be enabled, to use Data
Guard feature.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "exact_value"),
  validators.ConflictingListObjectAttributes("any_domain",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("apply_data_guard",
    "skip_data_guard"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value")}
```

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

<a id="canonical-3120113130333330-3130020033000330-3200013233223010-0123221300300002-2232210231310031-3210220202332322-2300221332111013-3032110113223211"></a>

## Direct properties — data_guard_rules / 001230221202 / 3

- [any_domain](resources--http_loadbalancer--reference--group-015.md#canonical-1100223213201200-3013211301203301-2232221301023211-0220003020233130-2322231030203220-2211032131231230-0123301020332223-0231223120303232): complete subsection reference.

- [apply_data_guard](resources--http_loadbalancer--reference--group-015.md#canonical-3001102122030200-3301002312100211-1002301213223202-3232311010130213-0101130222013123-2003210012130030-3032001333131231-0330002120111321): complete subsection reference.

<a id="canonical-1030113002301031-3330130301130223-1110023023300330-2232212331233223-3100031222110313-2022103322003031-1133032021113133-0200122311121112"></a>

<a id="canonical-0202111102112201-2303110001033121-2222020222232323-0232201203230322-2313022230122003-3303000011103320-1201021332133020-1020111300300303"></a>

## exact_value property — data_guard_rules / 001230221202 / 4

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-3231222132212132-0032012311231100-2121312022021023-3321232232233303-0101130100100123-2023003000221133-0300330313312232-0222213301123330"></a>

## suffix_value property — data_guard_rules / 001230221202 / 5

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-2230002032032000-0200231310300201-0203023230303030-2131101203233113-2210323001233203-3003000221322321-3232101111111103-2202223301132230"></a>

## Next pages — data_guard_rules / 001230221202 / 6

- [data_guard_rules.any_domain](resources--http_loadbalancer--reference--group-015.md#canonical-1100223213201200-3013211301203301-2232221301023211-0220003020233130-2322231030203220-2211032131231230-0123301020332223-0231223120303232)
- [data_guard_rules.apply_data_guard](resources--http_loadbalancer--reference--group-015.md#canonical-3001102122030200-3301002312100211-1002301213223202-3232311010130213-0101130222013123-2003210012130030-3032001333131231-0330002120111321)
- [data_guard_rules.metadata](resources--http_loadbalancer--reference--group-015.md#canonical-0323221032321021-2030300311221110-0222222132031132-1122210101230112-1323302312123031-0320301102300213-0002110331132031-3102011031231300)
- [data_guard_rules.path](resources--http_loadbalancer--reference--group-015.md#canonical-0123302113001130-0111220321333012-1301120122101023-0022201302232111-3220031110222113-1112013302111132-3121322030113223-3213132330021320)
- [data_guard_rules.skip_data_guard](resources--http_loadbalancer--reference--group-015.md#canonical-2223223002232112-0031332231112220-0313321012300221-3130321302211113-3121021320121010-1303023132001130-1332002003021030-2033221213313322)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1100223213201200-3013211301203301-2232221301023211-0220003020233130-2322231030203220-2211032131231230-0123301020332223-0231223120303232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101211323331023-0320033213133010-3102002233012301-1132210112300210-3313110330133321-3022303123230221-3213210123021012-0110313233202320"></a>

## data_guard_rules.any_domain — any_domain / 211313201000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [data_guard_rules](resources--http_loadbalancer--reference--group-015.md#canonical-0310100223013020-2111222133111333-2201233312231132-0102103013122132-0111220303113031-1122123223301021-1030212023023130-0331013321302221)
- data_guard_rules.any_domain

<a id="canonical-3122033332022100-3220201321022123-2312230120002032-0332312320203222-2330110030020301-1313210121332210-2233000100312213-1002221310300110"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
any_domain = {}
```

<a id="canonical-1332102211320302-0012112103332103-0221032310011133-0233122123301233-3223132102113133-2300210100220203-2101120031310113-2213000112100022"></a>

## Direct properties — any_domain / 211313201000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211022133311231-0223013212200112-2232212132133312-0211112301320113-1012233033202032-0223131120101302-1323023021123110-1212102332032213"></a>

## Next pages — any_domain / 211313201000 / 4

- [data_guard_rules](resources--http_loadbalancer--reference--group-015.md#canonical-0310100223013020-2111222133111333-2201233312231132-0102103013122132-0111220303113031-1122123223301021-1030212023023130-0331013321302221)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3001102122030200-3301002312100211-1002301213223202-3232311010130213-0101130222013123-2003210012130030-3032001333131231-0330002120111321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121012120323120-3120212013021021-2320330303322300-1003200213302313-0232211110323321-3212312131033011-1111120033221313-3000020210113333"></a>

## data_guard_rules.apply_data_guard — apply_data_guard / 003100313132 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [data_guard_rules](resources--http_loadbalancer--reference--group-015.md#canonical-0310100223013020-2111222133111333-2201233312231132-0102103013122132-0111220303113031-1122123223301021-1030212023023130-0331013321302221)
- data_guard_rules.apply_data_guard

<a id="canonical-3220121322031232-2121032000121113-1232323022310310-2310200322213012-1313121203133210-2123031213302111-3213211203103200-1233131033203211"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
apply_data_guard = {}
```

<a id="canonical-2203310202131200-2232123022322120-3311131203031133-0303220002031220-1101302311320202-2311233230110103-3231223112221211-2321001011302101"></a>

## Direct properties — apply_data_guard / 003100313132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211232201321112-1011223320211030-3122131021010100-1213132112003123-0032320113211133-0301322330232302-0002003000230312-2333333312102232"></a>

## Next pages — apply_data_guard / 003100313132 / 4

- [data_guard_rules](resources--http_loadbalancer--reference--group-015.md#canonical-0310100223013020-2111222133111333-2201233312231132-0102103013122132-0111220303113031-1122123223301021-1030212023023130-0331013321302221)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0323221032321021-2030300311221110-0222222132031132-1122210101230112-1323302312123031-0320301102300213-0002110331132031-3102011031231300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303202131133201-1011000310002201-2002313023132323-0000221200010120-3100212131002033-3300000212120100-3032201212110220-2222300002112301"></a>

## data_guard_rules.metadata — metadata / 023202330222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [data_guard_rules](resources--http_loadbalancer--reference--group-015.md#canonical-0310100223013020-2111222133111333-2201233312231132-0102103013122132-0111220303113031-1122123223301021-1030212023023130-0331013321302221)
- data_guard_rules.metadata

<a id="canonical-2320020022211320-0330312322123232-0303022211001130-0011331131130132-0030122012322233-2103100120130202-2300113020323203-2300303122122013"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003313110033010-1013023113231233-2031300200030022-0002012213300212-2331212323323032-1301012002230100-2230003003212201-1200033233031120"></a>

## Direct properties — metadata / 023202330222 / 3

<a id="canonical-0311112333212210-0030233110132002-2220330103032021-3202110123102323-2111333023132031-1132023121210301-1123230032221031-3000101302203302"></a>

<a id="canonical-3322030010123112-3021132111201230-2101012032010121-1323302101203232-0003231130123131-3131002321132032-1000322132332000-0112202031011030"></a>

## description_spec property — metadata / 023202330222 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3220103033020102-0300211313301322-0023133030121312-3013130233103311-0233011330013210-3313000311302123-2321212223102122-3201110320112310"></a>

<a id="canonical-1100001133012322-3123011102013003-3020202111310220-3120210201322203-0232202100312123-1220002100230110-0103022100312133-1000010011203030"></a>

## name property — metadata / 023202330222 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-2133233201001033-1021122211232030-2211301300101303-1332123031320111-1331332232113120-0212101122302303-2123211121102010-3300213033133221"></a>

## Next pages — metadata / 023202330222 / 6

- [data_guard_rules](resources--http_loadbalancer--reference--group-015.md#canonical-0310100223013020-2111222133111333-2201233312231132-0102103013122132-0111220303113031-1122123223301021-1030212023023130-0331013321302221)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0123302113001130-0111220321333012-1301120122101023-0022201302232111-3220031110222113-1112013302111132-3121322030113223-3213132330021320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000303222223331-2320013133021011-0132010312331322-0112300201202223-2030123131313232-1201201032020033-3103001013203222-0101021212233320"></a>

## data_guard_rules.path — path / 022013231213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [data_guard_rules](resources--http_loadbalancer--reference--group-015.md#canonical-0310100223013020-2111222133111333-2201233312231132-0102103013122132-0111220303113031-1122123223301021-1030212023023130-0331013321302221)
- data_guard_rules.path

<a id="canonical-1021020212101000-1123031231223302-1220233203310030-1231233223231230-2120210323320312-3102001301011200-1103120202102333-2300122031101330"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020302332033112-1011310333033132-1311102231333331-2231313023233331-0331320212121202-0201310203303333-0101003303001122-0232003123310022"></a>

## Direct properties — path / 022013231213 / 3

<a id="canonical-2200112303000130-2123103121311323-2321000030002100-0223012200022303-1233021230322100-2221202312202313-1111020203122001-0000112213012130"></a>

<a id="canonical-0121121332300023-2232101123211122-0223203120123030-2203013200122332-0210213210301032-2020023021010213-0012130300211332-3322001122213211"></a>

## path property — path / 022013231213 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3223013332130311-2032010322102031-3333320330121230-1321213300133233-3223203131200010-2202100223303201-2210222123230230-2120222213220221"></a>

## prefix property — path / 022013231213 / 5

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0102111311002311-2321213122131103-2021213123002111-0303303231223322-2232201011121010-1032201100023001-3002002320330200-1203213200303200"></a>

## regular expression property — path / 022013231213 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-0122020310200011-1001013313233201-3103321132321233-2103130322331211-1130332313212123-3233301233301202-0032231203322101-3011200031131323"></a>

## Next pages — path / 022013231213 / 7

- [data_guard_rules](resources--http_loadbalancer--reference--group-015.md#canonical-0310100223013020-2111222133111333-2201233312231132-0102103013122132-0111220303113031-1122123223301021-1030212023023130-0331013321302221)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2223223002232112-0031332231112220-0313321012300221-3130321302211113-3121021320121010-1303023132001130-1332002003021030-2033221213313322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002111123101110-1112302120032120-1110230032231331-0221010202323112-0322133023001013-1032020223300233-2101000223121311-3030100202113111"></a>

## data_guard_rules.skip_data_guard — skip_data_guard / 210103222011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [data_guard_rules](resources--http_loadbalancer--reference--group-015.md#canonical-0310100223013020-2111222133111333-2201233312231132-0102103013122132-0111220303113031-1122123223301021-1030212023023130-0331013321302221)
- data_guard_rules.skip_data_guard

<a id="canonical-3233222210111123-1013222330323001-0200132032310312-1101310023333110-0212121333311320-1230121223202330-1021222302033333-3223301232133331"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
skip_data_guard = {}
```

<a id="canonical-1121033303122311-2101202022102010-3113010230300100-0210311131023313-0230222101313100-0113233021212313-1022202321011301-2112233033030220"></a>

## Direct properties — skip_data_guard / 210103222011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232333213301030-0000202020323331-2002201002101230-3331010200310310-2111101301021323-2202131000213302-2320202031223200-0212131311222111"></a>

## Next pages — skip_data_guard / 210103222011 / 4

- [data_guard_rules](resources--http_loadbalancer--reference--group-015.md#canonical-0310100223013020-2111222133111333-2201233312231132-0102103013122132-0111220303113031-1122123223301021-1030212023023130-0331013321302221)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3113021303031131-1312012222010231-0331133200303132-3332023021101222-0113133213312010-1321130002311311-0012301201200102-1130020133302102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300301232101231-3132222012000123-0020311120110213-2102111212103212-0133111022030313-1031111230002020-1111112210021030-2013121331212033"></a>

## ddos_mitigation_rules — ddos_mitigation_rules / 111111220301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- ddos_mitigation_rules

<a id="canonical-2310100011112312-2222030220022330-3320033003021233-2022100203132203-1201120200323100-3330020021130132-0200021230032331-3111102132020133"></a>

Type: `"object"`. list nested block, Optional.

Define manual mitigation rules to block L7 DDoS attacks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ddos_client_source",
    "ip_prefix_list")}
```

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

<a id="canonical-1002103232102000-0001121123121201-3133323012130212-2213203112310012-3302101023112211-3032322111100032-0333320303210200-2213021300311322"></a>

## Direct properties — ddos_mitigation_rules / 111111220301 / 3

- [block](resources--http_loadbalancer--reference--group-015.md#canonical-3313103130322312-1312312213233301-2230220011033311-3313232113322333-2030110002201231-2102332030332003-3022223120213302-0300031033230022): complete subsection reference.

- [ddos_client_source](resources--http_loadbalancer--reference--group-015.md#canonical-2011213210300031-0330202201103102-0322200312222011-3232303332012220-2100231300233131-1100312130200112-2233220130221320-3011233231000021): complete subsection reference.

<a id="canonical-1221230000220131-2330202012021201-2200103312100233-1210121210221131-3333320132112031-2000120123133230-3233110030301000-2112333201220112"></a>

<a id="canonical-3133212112320021-3002021011112030-1011223122102232-1033022002021202-3320103312213002-0220101113323133-2103101023031310-3333212130003213"></a>

## expiration_timestamp property — ddos_mitigation_rules / 111111220301 / 4

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

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
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [ip_prefix_list](resources--http_loadbalancer--reference--group-015.md#canonical-1321132311131313-2033121331313303-3130133112123221-3130131220230331-1321011002332303-0231002203033031-0101022331300202-0033031003200210): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-015.md#canonical-1312030223203120-3311011003320303-0233321311230000-0010111210311333-3030331001001331-1311333111331332-1032200133313120-3030313233113233): complete subsection reference.

<a id="canonical-3202302003223333-3011010122032111-3222221020203113-1212212101002012-3001022013211013-2223023333103132-0332312011021100-1012011110011331"></a>

## Next pages — ddos_mitigation_rules / 111111220301 / 5

- [ddos_mitigation_rules.block](resources--http_loadbalancer--reference--group-015.md#canonical-3313103130322312-1312312213233301-2230220011033311-3313232113322333-2030110002201231-2102332030332003-3022223120213302-0300031033230022)
- [ddos_mitigation_rules.ddos_client_source](resources--http_loadbalancer--reference--group-015.md#canonical-2011213210300031-0330202201103102-0322200312222011-3232303332012220-2100231300233131-1100312130200112-2233220130221320-3011233231000021)
- [ddos_mitigation_rules.ip_prefix_list](resources--http_loadbalancer--reference--group-015.md#canonical-1321132311131313-2033121331313303-3130133112123221-3130131220230331-1321011002332303-0231002203033031-0101022331300202-0033031003200210)
- [ddos_mitigation_rules.metadata](resources--http_loadbalancer--reference--group-015.md#canonical-1312030223203120-3311011003320303-0233321311230000-0010111210311333-3030331001001331-1311333111331332-1032200133313120-3030313233113233)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3313103130322312-1312312213233301-2230220011033311-3313232113322333-2030110002201231-2102332030332003-3022223120213302-0300031033230022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033321322101112-0323212122103300-0003223320122120-3113203230021012-2321330030101231-0231220030002302-1132203130213302-0022133103232110"></a>

## ddos_mitigation_rules.block — block / 100113121230 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ddos_mitigation_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3113021303031131-1312012222010231-0331133200303132-3332023021101222-0113133213312010-1321130002311311-0012301201200102-1130020133302102)
- ddos_mitigation_rules.block

<a id="canonical-1100230323010312-1023121121001200-0122300331113102-3313331333120102-0032111111300102-3123210200221201-0311331233020210-2303001131103331"></a>

Type: `"object"`. single nested block, Optional.

Enable this option

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

Terraform syntax:

```terraform
block {}
```

<a id="canonical-2211233022211123-3103133120133311-2022232200201212-3232110330231110-1233101132000012-3313222111221323-3313321023200222-1303231311000301"></a>

## Direct properties — block / 100113121230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0330032102202323-0032222031130101-1031222021223332-3211101331133223-1011002003100330-1010131223003133-0102012212121330-1320010013210100"></a>

## Next pages — block / 100113121230 / 4

- [ddos_mitigation_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3113021303031131-1312012222010231-0331133200303132-3332023021101222-0113133213312010-1321130002311311-0012301201200102-1130020133302102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2011213210300031-0330202201103102-0322200312222011-3232303332012220-2100231300233131-1100312130200112-2233220130221320-3011233231000021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030233203120230-1312202310323031-0223111103110213-1131312010021103-2013210210332213-2332330100030132-3100003222120120-3322123120221122"></a>

## ddos_mitigation_rules.ddos_client_source — ddos_client_source / 002332030100 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ddos_mitigation_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3113021303031131-1312012222010231-0331133200303132-3332023021101222-0113133213312010-1321130002311311-0012301201200102-1130020133302102)
- ddos_mitigation_rules.ddos_client_source

<a id="canonical-1113120003032003-3220231102031110-0001102002031220-3131223331211003-1333113122001022-2030232100023302-3331023322003310-3231121312323313"></a>

Type: `"object"`. single nested block, Optional.

DDoS Client Source Choice. DDoS Mitigation sources to be blocked.

Upstream description:

DDoS Mitigation sources to be blocked.

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

<a id="canonical-0022231012112210-0120220322303302-3012321320213003-3302033231120331-0131123331012320-1102311211202133-0313120023300002-0113033020303120"></a>

## Direct properties — ddos_client_source / 002332030100 / 3

- [asn_list](resources--http_loadbalancer--reference--group-015.md#canonical-1301330111333132-0110110311312222-3123210123222332-1230131003020231-3003133001301200-2020233211222110-0333112122121121-0030000200322323): complete subsection reference.

<a id="canonical-1210033210203131-2130312012230202-3023102111111310-2202111200220332-1130122023211331-1112020111103223-0131320300232321-3023010103032230"></a>

<a id="canonical-1030222220222012-2130301330202220-0132130233220030-3302331021230202-1132213112231221-2112310130000022-1232302010020033-1231030232123013"></a>

## country_list property — ddos_client_source / 002332030100 / 4

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

Upstream description:

Sources that are located in one of the countries in the given list.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

- [ja4_tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-015.md#canonical-2210313030223112-0221132223013102-2133320023222302-2223212323022331-0111110010032032-2231320102021233-2132022013210322-2231020122012333): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-015.md#canonical-0031133313011232-1222230120123210-0233333120231321-3133333023031031-2020213302230211-1112301011201130-2033120322011123-0220211013233023): complete subsection reference.

<a id="canonical-0302120220312012-0130001023012032-3330221220311333-2110221120123132-3201010102303120-1132333121112210-1231113130113131-1320230211123003"></a>

## Next pages — ddos_client_source / 002332030100 / 5

- [ddos_mitigation_rules.ddos_client_source.asn_list](resources--http_loadbalancer--reference--group-015.md#canonical-1301330111333132-0110110311312222-3123210123222332-1230131003020231-3003133001301200-2020233211222110-0333112122121121-0030000200322323)
- [ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-015.md#canonical-2210313030223112-0221132223013102-2133320023222302-2223212323022331-0111110010032032-2231320102021233-2132022013210322-2231020122012333)
- [ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-015.md#canonical-0031133313011232-1222230120123210-0233333120231321-3133333023031031-2020213302230211-1112301011201130-2033120322011123-0220211013233023)
- [ddos_mitigation_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3113021303031131-1312012222010231-0331133200303132-3332023021101222-0113133213312010-1321130002311311-0012301201200102-1130020133302102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1301330111333132-0110110311312222-3123210123222332-1230131003020231-3003133001301200-2020233211222110-0333112122121121-0030000200322323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023100031010321-2230323230022113-0201230332200220-1012020330230102-2321232220013301-0002111312033210-0021203323122313-1022012232201303"></a>

## ddos_mitigation_rules.ddos_client_source.asn_list — asn_list / 330302010210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ddos_mitigation_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3113021303031131-1312012222010231-0331133200303132-3332023021101222-0113133213312010-1321130002311311-0012301201200102-1130020133302102)
- [ddos_mitigation_rules.ddos_client_source](resources--http_loadbalancer--reference--group-015.md#canonical-2011213210300031-0330202201103102-0322200312222011-3232303332012220-2100231300233131-1100312130200112-2233220130221320-3011233231000021)
- ddos_mitigation_rules.ddos_client_source.asn_list

<a id="canonical-0023301300101310-2311113001302100-1132200012111330-3211231213233113-3312300203030303-2021102311102320-1232323200012131-3300330312302033"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221003133132333-3101021221332123-2212103302101333-0203231333201330-2023102011120211-1101020223201011-0020320333100230-3033101313001013"></a>

## Direct properties — asn_list / 330302010210 / 3

<a id="canonical-1310021322230211-2100203311121330-1000203313010020-1213032331310221-1212103111032301-2023013100033313-3223013010120012-1302021020110113"></a>

<a id="canonical-3230100030303020-2312023301312003-0211131010101101-2301103210212323-2323223020322011-2113031312201021-2032313112212030-3012001001302031"></a>

## as_numbers property — asn_list / 330302010210 / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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

<a id="canonical-0132310222101031-1330031031130001-0031231012232213-2320223123213212-3202321110330330-3033210333330033-2121102012311013-2133131103000321"></a>

## Next pages — asn_list / 330302010210 / 5

- [ddos_mitigation_rules.ddos_client_source](resources--http_loadbalancer--reference--group-015.md#canonical-2011213210300031-0330202201103102-0322200312222011-3232303332012220-2100231300233131-1100312130200112-2233220130221320-3011233231000021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2210313030223112-0221132223013102-2133320023222302-2223212323022331-0111110010032032-2231320102021233-2132022013210322-2231020122012333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211223021120211-1021303023130112-3001010013313000-0302010032112321-1231033113111132-2211000032121030-0103101122211132-1323200233123030"></a>

## ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher — ja4_tls_fingerprint_matcher / 032020312132 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ddos_mitigation_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3113021303031131-1312012222010231-0331133200303132-3332023021101222-0113133213312010-1321130002311311-0012301201200102-1130020133302102)
- [ddos_mitigation_rules.ddos_client_source](resources--http_loadbalancer--reference--group-015.md#canonical-2011213210300031-0330202201103102-0322200312222011-3232303332012220-2100231300233131-1100312130200112-2233220130221320-3011233231000021)
- ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher

<a id="canonical-0213333131203211-1231130303002010-1330103020332011-1223111102230020-2231020013032131-2121001322220223-1221100011113023-2333310211102312"></a>

Type: `"object"`. single nested block, Optional.

Extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

Upstream description:

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

Terraform syntax:

```terraform
ja4_tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0010301010200021-1320321011130132-3232021001223310-3303121202113101-2000020013320032-3132222300210322-0322320101322010-0010123123322033"></a>

## Direct properties — ja4_tls_fingerprint_matcher / 032020312132 / 3

<a id="canonical-0231111002311222-1012030211113310-0032231201013302-2003301003110102-3012032200331301-2310230331210333-1202231200303220-2003330122032130"></a>

<a id="canonical-0221102311303103-2312312112300102-2010330132120330-3332211120131102-3123221112112330-2211230132323202-0122023332012211-0032301103320222"></a>

## exact_values property — ja4_tls_fingerprint_matcher / 032020312132 / 4

Type: `["list", "string"]`. Optional.

List of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

Upstream description:

A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-3001030111103113-1220311103111231-1110023221220232-3331213213321310-0230022220001112-2233222233223122-3023312202032022-3313203010131302"></a>

## Next pages — ja4_tls_fingerprint_matcher / 032020312132 / 5

- [ddos_mitigation_rules.ddos_client_source](resources--http_loadbalancer--reference--group-015.md#canonical-2011213210300031-0330202201103102-0322200312222011-3232303332012220-2100231300233131-1100312130200112-2233220130221320-3011233231000021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0031133313011232-1222230120123210-0233333120231321-3133333023031031-2020213302230211-1112301011201130-2033120322011123-0220211013233023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001023011312230-2313220022010301-2322220130033232-2110112211131002-3302120231103323-0123221121301021-3111302303220013-3311300111001013"></a>

## ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher — tls_fingerprint_matcher / 213033002321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ddos_mitigation_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3113021303031131-1312012222010231-0331133200303132-3332023021101222-0113133213312010-1321130002311311-0012301201200102-1130020133302102)
- [ddos_mitigation_rules.ddos_client_source](resources--http_loadbalancer--reference--group-015.md#canonical-2011213210300031-0330202201103102-0322200312222011-3232303332012220-2100231300233131-1100312130200112-2233220130221320-3011233231000021)
- ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher

<a id="canonical-1301231012032030-2312000130223012-3102210331321121-2230031332121121-2211300320120213-1021121003330003-0023333123013202-0330323013020231"></a>

Type: `"object"`. single nested block, Optional.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

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

<a id="canonical-2103323110321230-3100321311323231-3331323320203332-0122323332121101-1212201010013130-2321101211230021-1023111332110103-1320321231132103"></a>

## Direct properties — tls_fingerprint_matcher / 213033002321 / 3

<a id="canonical-1111313303201232-2300302012130103-0101113013310222-1331203102200222-1201121013221023-0032020112220031-3013132133223020-2212321203032201"></a>

<a id="canonical-2132310121013230-3010003232100101-3103103120313303-2012211212113121-3100012013202110-2333021332113321-1301122001222011-2110130212330122"></a>

## classes property — tls_fingerprint_matcher / 213033002321 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1331021101111330-3202032100301003-1212333001302121-2210313321223032-0021322212220333-1113303223230130-0001013301221330-0021301113122320"></a>

<a id="canonical-0132221121022312-3000210320203330-0301211131212213-3001002112003010-3031221231203331-0002313123123230-0113202001123013-1313013122211032"></a>

## exact_values property — tls_fingerprint_matcher / 213033002321 / 5

Type: `["list", "string"]`. Optional.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-1303121322110022-1003020201202003-3113132202311011-3000203020233022-1311211032122303-0202210332232123-1321123122323113-2203033321000303"></a>

<a id="canonical-2022333333122311-0200033121112321-1220322301111131-1321231103201123-1331302021110212-0100011230231200-2322210012213232-3010033300322212"></a>

## excluded_values property — tls_fingerprint_matcher / 213033002321 / 6

Type: `["list", "string"]`. Optional.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
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

<a id="canonical-2122300032221111-2222111210330023-0130000232300013-1321221010103233-0231223233002100-1021310310321121-1113030212311021-1022321112102003"></a>

## Next pages — tls_fingerprint_matcher / 213033002321 / 7

- [ddos_mitigation_rules.ddos_client_source](resources--http_loadbalancer--reference--group-015.md#canonical-2011213210300031-0330202201103102-0322200312222011-3232303332012220-2100231300233131-1100312130200112-2233220130221320-3011233231000021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1321132311131313-2033121331313303-3130133112123221-3130131220230331-1321011002332303-0231002203033031-0101022331300202-0033031003200210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332211223002000-3222113212012002-1330232201303310-0120222101333210-2001230122323011-3331333302202003-3232132112331333-1232003110230031"></a>

## ddos_mitigation_rules.ip_prefix_list — ip_prefix_list / 321030110211 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ddos_mitigation_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3113021303031131-1312012222010231-0331133200303132-3332023021101222-0113133213312010-1321130002311311-0012301201200102-1130020133302102)
- ddos_mitigation_rules.ip_prefix_list

<a id="canonical-2220332223002033-2301310230130011-0231222203030332-0222312221201323-3313333220012333-2313313120103002-2313100232221231-2000020132123132"></a>

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

<a id="canonical-0203300202200032-3320333123102323-2021332231103111-2123121003201220-0002021323112310-2011220211332200-3010112212303101-2111031312030031"></a>

## Direct properties — ip_prefix_list / 321030110211 / 3

<a id="canonical-0010202201120001-0012013110212130-0011110322113012-0230120301003322-1222031322110003-0230022303312012-0313033123222200-0003211300031310"></a>

<a id="canonical-2003002003111213-1030320233133200-0001123022303132-1223202133212322-1013033110201110-3102201033213221-2223312023203121-2230320300230003"></a>

## invert_match property — ip_prefix_list / 321030110211 / 4

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-1022302111212103-2000133201322023-0030233201113331-1211131023132211-1223122110003101-0112213300101332-3001310333312330-1012113020123302"></a>

<a id="canonical-1222230032220210-1110002233112012-3033020101211101-0022131201021311-3222321100210010-3232000131313313-0023233022000010-0300333001013321"></a>

## ip_prefixes property — ip_prefix_list / 321030110211 / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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

<a id="canonical-2133232011210101-1232010101130113-2221332333100031-1302303203332200-1122102311101113-1013020023130221-3202113213033333-0102100230222311"></a>

## Next pages — ip_prefix_list / 321030110211 / 6

- [ddos_mitigation_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3113021303031131-1312012222010231-0331133200303132-3332023021101222-0113133213312010-1321130002311311-0012301201200102-1130020133302102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1312030223203120-3311011003320303-0233321311230000-0010111210311333-3030331001001331-1311333111331332-1032200133313120-3030313233113233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231031330221232-2031232210303133-1302112323321003-2200101132011022-3120101131030023-0133011332033212-2133201121120230-0212113320203003"></a>

## ddos_mitigation_rules.metadata — metadata / 210103212013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ddos_mitigation_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3113021303031131-1312012222010231-0331133200303132-3332023021101222-0113133213312010-1321130002311311-0012301201200102-1130020133302102)
- ddos_mitigation_rules.metadata

<a id="canonical-3112213013320312-1313033300233232-2220101320021223-0303102202111102-0212222101300231-0202233012120033-3223311213223313-0212203131311221"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212021011102011-1003220310203030-1203200220232232-2001330231011113-2303212012031031-2101100012001130-2122233232002202-1012110130311313"></a>

## Direct properties — metadata / 210103212013 / 3

<a id="canonical-3013222301231323-3100200120312220-2003233321003012-0133230010233002-0211332132203301-3322303331131032-3020120312223223-3211022023202233"></a>

<a id="canonical-3133002012020200-2330023002213311-0322321321202031-3213001101230123-3320300132111012-3110102233332220-3222011023121223-1330312132102112"></a>

## description_spec property — metadata / 210103212013 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1031231033320012-2232210210221331-0212031203123112-1121131301301101-0113331002030123-1032300323321321-2232201102222022-0030323123212101"></a>

<a id="canonical-1233233330322311-1332132312022000-1212233120010203-1300232220021110-1110031111331001-2321100011203322-2313213200211002-0120213233112121"></a>

## name property — metadata / 210103212013 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-2032200311123331-0120031200300021-0212210120013303-2021020130232021-2130003223203233-0000102211030101-2331112320120012-1211221221022103"></a>

## Next pages — metadata / 210103212013 / 6

- [ddos_mitigation_rules](resources--http_loadbalancer--reference--group-015.md#canonical-3113021303031131-1312012222010231-0331133200303132-3332023021101222-0113133213312010-1321130002311311-0012301201200102-1130020133302102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313311312330121-3100113323301313-2012102320010002-0321311210303111-2210332210131331-0312212202013111-0031132222113121-1000130220311302"></a>

## default_pool — default_pool / 220310202111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- default_pool

<a id="canonical-1213330012310303-3223033000020112-1132322020000213-3203211100210032-3203202100010003-0323120100333103-3310310021000230-2332331033130013"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: default\_pool, default\_pool\_list; Default: default\_pool\] Configuration parameter for
default pool.

Upstream description:

Shape of the origin pool specification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("origin_servers"),
  validators.ConflictingObjectAttributes("automatic_port",
    "lb_port"),
  validators.ConflictingObjectAttributes("automatic_port",
    "port"),
  validators.ConflictingObjectAttributes("health_check_port",
    "same_as_endpoint_port"),
  validators.ConflictingObjectAttributes("lb_port",
    "port"),
  validators.ConflictingObjectAttributes("no_tls",
    "use_tls")}
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
  "x-ves-oneof-field-health_check_port_choice": "[\"health_check_port\",\"same_as_endpoint_port\"]",
  "x-ves-oneof-field-port_choice": "[\"automatic_port\",\"lb_port\",\"port\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

OneOf alternatives in this subsection:

- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-1213330012310303-3223033000020112-1132322020000213-3203211100210032-3203202100010003-0323120100333103-3310310021000230-2332331033130013)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-0130032320003112-3013231322132011-2010113111012220-3312013033110232-0201300301333010-2200230132012313-2122033102033333-3101112132302330)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200033130233022-2111031030233113-2033123120032200-0301320231120233-0010303201100120-0331202202132202-3010221022100131-1023133002221332"></a>

## Direct properties — default_pool / 220310202111 / 3

- [advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100): complete subsection reference.

- [automatic_port](resources--http_loadbalancer--reference--group-016.md#canonical-1322002312102330-3201331032321312-3300300133011331-3120132210110301-3210102232133203-0002132201013013-2023203103322300-1312001000200312): complete subsection reference.

<a id="canonical-1220013021102311-1210300212021323-3030313021111112-2331311200001322-1301030222003002-2332032220130322-3233323111021113-2321012103002111"></a>

<a id="canonical-0122133123300312-3202000120222003-1030023200201002-2233203202220001-1112212130113013-2333100113022003-0321013222222331-0001031003002212"></a>

## endpoint_selection property — default_pool / 220310202111 / 4

Type: `"string"`. Optional, Computed.

\[Enum: DISTRIBUTED|LOCAL\_ONLY|LOCAL\_PREFERRED\] Policy for selection of endpoints from local
site/remote site/both Consider both remote and local endpoints for load balancing LOCAL\_ONLY:
Consider only local endpoints for load balancing Enable this policy to load balance ONLY among
locally discovered endpoints Prefer the local endpoints for.. Possible values are \`DISTRIBUTED\`,
\`LOCAL\_ONLY\`, \`LOCAL\_PREFERRED\`. Defaults to \`DISTRIBUTED\`. Server applies default when
omitted.

Upstream description:

Policy for selection of endpoints from local site/remote site/both

Consider both remote and local endpoints for load balancing LOCAL\_ONLY: Consider only local
endpoints for load balancing Enable this policy to load balance ONLY among locally discovered
endpoints Prefer the local endpoints for load balancing. If local endpoints are not present remote
endpoints will be considered.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DISTRIBUTED",
    "LOCAL_ONLY",
    "LOCAL_PREFERRED"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DISTRIBUTED",
  "enum": [
    "DISTRIBUTED",
    "LOCAL_ONLY",
    "LOCAL_PREFERRED"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2333031322131122-1333231311010031-0023200220303113-2223032011333030-1203001000013012-0333300323010101-2103102303202333-2212032221013220"></a>

<a id="canonical-1030000100233130-1011130013132331-1301121230320021-3331112313131312-0320130023233230-1301222002330310-2133313332132123-1302031122033231"></a>

## health_check_port property — default_pool / 220310202111 / 5

Type: `"number"`. Optional.

Exclusive with \[same\_as\_endpoint\_port\] Port used for performing health check.

Upstream description:

Exclusive with \[same\_as\_endpoint\_port\] Port used for performing health check.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "category": "networking",
      "confidence": 0.99,
      "note": "Asymmetry: port enforces [1,65535], health_check_port allows 0",
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [healthcheck](resources--http_loadbalancer--reference--group-016.md#canonical-1221110002132303-0012322203331031-0200320230330313-0111101330333233-2122220212333113-0311121023131102-2021213213213203-2231002020013210): complete subsection reference.

- [lb_port](resources--http_loadbalancer--reference--group-016.md#canonical-1023332221023002-1120001320231221-2230003000223121-1203332032211030-3112300112131312-2203210121000020-1110310100200211-0012031020100313): complete subsection reference.

<a id="canonical-2001303000020210-3130102121230130-1031001003223123-1300321230302002-0023213122001213-2022000210011132-0033113233123233-3202100031231222"></a>

<a id="canonical-0202221012133213-3120301033212222-3303130032003223-2121332010112030-2221232112000020-0300321131331000-3131222033311032-1010010221121300"></a>

## loadbalancer_algorithm property — default_pool / 220310202111 / 6

Type: `"string"`. Optional, Computed.

\[Enum: ROUND\_ROBIN|LEAST\_REQUEST|RING\_HASH|RANDOM|LB\_OVERRIDE\] Different load balancing
algorithms supported When a connection to a endpoint in an upstream cluster is required, the load
balancer uses loadbalancer\_algorithm to determine which host is selected. - ROUND\_ROBIN:
ROUND\_ROBIN Policy in which each healthy/available upstream endpoint is selected in.. Possible
values are \`ROUND\_ROBIN\`, \`LEAST\_REQUEST\`, \`RING\_HASH\`, \`RANDOM\`, \`LB\_OVERRIDE\`.
Defaults to \`ROUND\_ROBIN\`. Server applies default when omitted.

Upstream description:

Different load balancing algorithms supported When a connection to a endpoint in an upstream cluster
is required, the load balancer uses loadbalancer\_algorithm to determine which host is selected.

&#8203;- ROUND\_ROBIN: ROUND\_ROBIN

Policy in which each healthy/available upstream endpoint is selected in round robin order. &#8203;-
LEAST\_REQUEST: LEAST\_REQUEST

Policy in which loadbalancer picks the upstream endpoint which has the fewest active requests
&#8203;- RING\_HASH: RING\_HASH

Policy implements consistent hashing to upstream endpoints using ring hash of endpoint names Hash of
the incoming request is calculated using request hash policy. The ring/modulo hash load balancer
implements consistent hashing to upstream hosts. The algorithm is based on mapping all hosts onto a
circle such that the addition or removal of a host from the host set changes only affect 1/N
requests. This technique is also commonly known as “ketama” hashing. A consistent hashing load
balancer is only effective when protocol routing is used that specifies a value to hash on. The
minimum ring size governs the replication factor for each host in the ring. For example, if the
minimum ring size is 1024 and there are 16 hosts, each host will be replicated 64 times. &#8203;-
RANDOM: RANDOM

Policy in which each available upstream endpoint is selected in random order. The random load
balancer selects a random healthy host. The random load balancer generally performs better than
round robin if no health checking policy is configured. Random selection avoids bias towards the
host in the set that comes after a failed host. &#8203;- LB\_OVERRIDE: Load Balancer Override

Hash policy is taken from from the load balancer which is using this origin pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ROUND_ROBIN",
    "LEAST_REQUEST",
    "RING_HASH",
    "RANDOM",
    "LB_OVERRIDE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ROUND_ROBIN",
  "enum": [
    "ROUND_ROBIN",
    "LEAST_REQUEST",
    "RING_HASH",
    "RANDOM",
    "LB_OVERRIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_tls](resources--http_loadbalancer--reference--group-016.md#canonical-3222232123033110-0311020222300102-2311031131230111-0311001112220210-0122011130233121-0322223310331112-0222133023233112-0031210201010010): complete subsection reference.

- [origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031): complete subsection reference.

<a id="canonical-1110332222020032-1030112012011212-0120103200222313-2331302310323003-2322213313231001-1300233203310123-0212302001322020-3231203111022120"></a>

<a id="canonical-2131312003011133-3211130320313223-2101103130113013-2330311301033303-3323112123030021-3003300300121011-3213320211321333-0030133103312310"></a>

## port property — default_pool / 220310202111 / 7

Type: `"number"`. Optional.

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port. Recommended:
\`443\`.

Upstream description:

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [same_as_endpoint_port](resources--http_loadbalancer--reference--group-017.md#canonical-1232100333033333-0221010121032103-2010033203301010-1223212021203320-2001102032013210-0333301222330003-1231200013113223-2010211132110221): complete subsection reference.

- [upstream_conn_pool_reuse_type](resources--http_loadbalancer--reference--group-017.md#canonical-1112003010300102-0001210211031011-0021210310211031-2120013323123301-1203112022230211-2311310013202331-2323031211322202-1321003211023230): complete subsection reference.

- [use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022): complete subsection reference.

- [view_internal](resources--http_loadbalancer--reference--group-017.md#canonical-0210031322021130-2010122303312010-3232211331331322-0132231002001122-1111001310002033-1200111313232223-3023311310020010-2112111310030332): complete subsection reference.

<a id="canonical-1221031332202330-3133210233123203-1201323322232020-1322033312001330-0202203233103121-3330301312032022-2001313031223230-1203103320212020"></a>

## Next pages — default_pool / 220310202111 / 8

- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [default_pool.automatic_port](resources--http_loadbalancer--reference--group-016.md#canonical-1322002312102330-3201331032321312-3300300133011331-3120132210110301-3210102232133203-0002132201013013-2023203103322300-1312001000200312)
- [default_pool.healthcheck](resources--http_loadbalancer--reference--group-016.md#canonical-1221110002132303-0012322203331031-0200320230330313-0111101330333233-2122220212333113-0311121023131102-2021213213213203-2231002020013210)
- [default_pool.lb_port](resources--http_loadbalancer--reference--group-016.md#canonical-1023332221023002-1120001320231221-2230003000223121-1203332032211030-3112300112131312-2203210121000020-1110310100200211-0012031020100313)
- [default_pool.no_tls](resources--http_loadbalancer--reference--group-016.md#canonical-3222232123033110-0311020222300102-2311031131230111-0311001112220210-0122011130233121-0322223310331112-0222133023233112-0031210201010010)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.same_as_endpoint_port](resources--http_loadbalancer--reference--group-017.md#canonical-1232100333033333-0221010121032103-2010033203301010-1223212021203320-2001102032013210-0333301222330003-1231200013113223-2010211132110221)
- [default_pool.upstream_conn_pool_reuse_type](resources--http_loadbalancer--reference--group-017.md#canonical-1112003010300102-0001210211031011-0021210310211031-2120013323123301-1203112022230211-2311310013202331-2323031211322202-1321003211023230)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.view_internal](resources--http_loadbalancer--reference--group-017.md#canonical-0210031322021130-2010122303312010-3232211331331322-0132231002001122-1111001310002033-1200111313232223-3023311310020010-2112111310030332)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203020001311110-2120133110033003-3322002213123033-1121113031003113-1113300011201231-3310331120322203-0210320330010003-2212300222021300"></a>

## default_pool.advanced_options — advanced_options / 323213322131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.advanced_options

<a id="canonical-2201323003203213-3213010002212030-3101030122020221-1021231203300201-1221330132003001-2201133313112133-2323222210220222-1133020302333331"></a>

Type: `"object"`. single nested block, Optional.

Configure Advanced OPTIONS for origin pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_http_config",
    "http1_config"),
  validators.ConflictingObjectAttributes("auto_http_config",
    "http2_options"),
  validators.ConflictingObjectAttributes("circuit_breaker",
    "default_circuit_breaker"),
  validators.ConflictingObjectAttributes("circuit_breaker",
    "disable_circuit_breaker"),
  validators.ConflictingObjectAttributes("default_circuit_breaker",
    "disable_circuit_breaker"),
  validators.ConflictingObjectAttributes("disable_lb_source_ip_persistence",
    "enable_lb_source_ip_persistence"),
  validators.ConflictingObjectAttributes("disable_outlier_detection",
    "outlier_detection"),
  validators.ConflictingObjectAttributes("disable_proxy_protocol",
    "proxy_protocol_v1"),
  validators.ConflictingObjectAttributes("disable_proxy_protocol",
    "proxy_protocol_v2"),
  validators.ConflictingObjectAttributes("disable_subsets",
    "enable_subsets"),
  validators.ConflictingObjectAttributes("http1_config",
    "http2_options"),
  validators.ConflictingObjectAttributes("max_requests_per_connection",
    "no_request_limit_per_connection"),
  validators.ConflictingObjectAttributes("no_panic_threshold",
    "panic_threshold"),
  validators.ConflictingObjectAttributes("proxy_protocol_v1",
    "proxy_protocol_v2")}
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
  "x-ves-oneof-field-circuit_breaker_choice": "[\"circuit_breaker\",\"default_circuit_breaker\",\"disable_circuit_breaker\"]",
  "x-ves-oneof-field-http_protocol_type": "[\"auto_http_config\",\"http1_config\",\"http2_options\"]",
  "x-ves-oneof-field-lb_source_ip_persistence_choice": "[\"disable_lb_source_ip_persistence\",\"enable_lb_source_ip_persistence\"]",
  "x-ves-oneof-field-max_requests_per_connection_choice": "[\"max_requests_per_connection\",\"no_request_limit_per_connection\"]",
  "x-ves-oneof-field-outlier_detection_choice": "[\"disable_outlier_detection\",\"outlier_detection\"]",
  "x-ves-oneof-field-panic_threshold_type": "[\"no_panic_threshold\",\"panic_threshold\"]",
  "x-ves-oneof-field-proxy_protocol_choice": "[\"disable_proxy_protocol\",\"proxy_protocol_v1\",\"proxy_protocol_v2\"]",
  "x-ves-oneof-field-subset_choice": "[\"disable_subsets\",\"enable_subsets\"]"
}
```

Terraform syntax:

```terraform
advanced_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0222110112333101-0303201202033301-0233121231323020-0200130130320022-2110130320310210-3110320021221211-2022301312210221-2233223102323113"></a>

## Direct properties — advanced_options / 323213322131 / 3

- [auto_http_config](resources--http_loadbalancer--reference--group-015.md#canonical-3112103233111313-3303220101001113-3022220201103300-2310102331122013-0202002302022030-1310031130301100-1231313211312123-1131333023120221): complete subsection reference.

- [circuit_breaker](resources--http_loadbalancer--reference--group-015.md#canonical-1031212032101202-3130331233121230-1022333233112301-0011311210120111-0130221013310320-2201133030023101-1021133320211103-0302331222320211): complete subsection reference.

<a id="canonical-0002230021000320-2212101133003322-0233132212110213-2310101011321021-0311211000003020-1102212120121103-0231202302003033-1101311110300322"></a>

<a id="canonical-3222110212332003-3202333312232012-2300333121233311-0120300002232323-2032331100220223-3021000100322033-2032221222122032-1211333331233101"></a>

## connection_timeout property — advanced_options / 323213322131 / 4

Type: `"number"`. Optional, Computed.

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds. Server applies default when omitted. Recommended:
\`2000\`.

Upstream description:

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(1800000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1800000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "1800000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  }
}
```

- [default_circuit_breaker](resources--http_loadbalancer--reference--group-015.md#canonical-0220231330001103-1312111311013211-2301111133301230-0123001132310332-0121313113212333-0001211212112332-2201120003102123-2223202111113133): complete subsection reference.

- [disable_circuit_breaker](resources--http_loadbalancer--reference--group-015.md#canonical-3030013113131131-3103010310302203-2123231330311322-1231031222331233-1321202223300003-3001203202013202-3000310231301223-3101032101210223): complete subsection reference.

- [disable_lb_source_ip_persistence](resources--http_loadbalancer--reference--group-015.md#canonical-3113100222020202-3133203310100220-0330332303302301-2031303313121332-0323121303222111-3223132221322032-0310310300022112-3030312000202003): complete subsection reference.

- [disable_outlier_detection](resources--http_loadbalancer--reference--group-015.md#canonical-1300111032032202-0223330012333200-0003032323033300-1220012210221033-1221031213013231-0112210331000110-3103001312310113-2332322021231202): complete subsection reference.

- [disable_proxy_protocol](resources--http_loadbalancer--reference--group-015.md#canonical-1323211300100232-0123211012202031-0232200001130201-1010020100302322-1122010113130203-0321323222203113-0330110330301230-2203311200223113): complete subsection reference.

- [disable_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-0023103001332003-0303111121333030-1302202301203023-2300202001223132-3000303201112022-3312000113212101-3222132031232321-1110111201302101): complete subsection reference.

- [enable_lb_source_ip_persistence](resources--http_loadbalancer--reference--group-015.md#canonical-1030023020110210-2012011100303020-3003123221320323-3003020312113033-1230031030231320-0130321233030101-1032013131300003-2201032222030012): complete subsection reference.

- [enable_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-3122011012203130-2032023233101333-0322112222003112-2311302231102332-2113303332212233-1102000111031122-0303332031001202-2221302030020230): complete subsection reference.

- [http1_config](resources--http_loadbalancer--reference--group-016.md#canonical-0302331112230021-3012012233210300-0212201311313112-1010303210232213-1332033331013232-3023212022323232-1321110310002103-2032000231020000): complete subsection reference.

- [http2_options](resources--http_loadbalancer--reference--group-016.md#canonical-3113102323303220-1222012231233210-3132230022230202-2102202133132231-2231111233120333-0131210111011331-3220111022033123-3110000201103023): complete subsection reference.

<a id="canonical-3101121222220101-1333020313101022-3010101311133101-2010131210230212-0102100210212100-3213233303001020-0111120223101123-1130022001221311"></a>

<a id="canonical-1111231000003202-1230310223130123-0223121133222011-2001330003230122-2100321100120230-1031021223221232-2130022013113310-2210032002333121"></a>

## http_idle_timeout property — advanced_options / 323213322131 / 5

Type: `"number"`. Optional, Computed.

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Server applies default when omitted. Recommended: \`300000\`.

Upstream description:

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.
This is specified in milliseconds. The default value is 5 minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-3312313120101212-3310113103331102-2030021330123121-1202121211210203-2100131131303331-3103211310312030-3230031121311332-2032002201300033"></a>

<a id="canonical-0123033020210300-2220132011233120-2013221330103221-1012333212003302-0131310302033212-3211002312322102-1322211100131213-3320011121013120"></a>

## max_requests_per_connection property — advanced_options / 323213322131 / 6

Type: `"number"`. Optional.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [no_panic_threshold](resources--http_loadbalancer--reference--group-016.md#canonical-3010321022203332-1133303230302303-2331101002312232-0110221102133000-2100020203120200-3333120023011003-2312121111020320-1223232102333011): complete subsection reference.

- [no_request_limit_per_connection](resources--http_loadbalancer--reference--group-016.md#canonical-3202011012323122-2310202220132112-1203221321333321-3232310230130231-0223011220002020-3131220013013302-1112021212030131-2212302110230313): complete subsection reference.

- [outlier_detection](resources--http_loadbalancer--reference--group-016.md#canonical-3122120002230323-1331302213332300-0221022000101211-0233001322022311-3200121333120233-0120103223222330-0001121011111302-0100313133112132): complete subsection reference.

<a id="canonical-1011102130302100-3311233002102131-1210231013102012-3301233233202330-0331002201332113-2102312333332312-2200130212312232-1311032321222202"></a>

<a id="canonical-1110321202323020-3013212113202230-3202121012131020-1011201012333110-2132310233112012-3211012030113223-3331301213120131-2322312330101103"></a>

## panic_threshold property — advanced_options / 323213322131 / 7

Type: `"number"`. Optional.

Exclusive with \[no\_panic\_threshold\] Configure a threshold (percentage of unhealthy endpoints)
below which all endpoints will be considered for load balancing ignoring its health status.

Upstream description:

Exclusive with \[no\_panic\_threshold\]

Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be
considered for load balancing ignoring its health status.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(100),
}
```

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [proxy_protocol_v1](resources--http_loadbalancer--reference--group-016.md#canonical-0320332213030303-2110221221131312-1112220010132111-3111231331132203-1012323131311113-3203232033330032-3300331111032003-0132310330010030): complete subsection reference.

- [proxy_protocol_v2](resources--http_loadbalancer--reference--group-016.md#canonical-2002203333000021-3223210132201331-1111301110300313-0211110110221123-3233022321110020-0320011110202330-3210010203121323-3010312231222231): complete subsection reference.

<a id="canonical-1023133023320232-3132303113010023-3132113022301301-0110002102200232-0230320122202110-3000323130222022-1322122121331303-0002120002103013"></a>

## Next pages — advanced_options / 323213322131 / 8

- [default_pool.advanced_options.auto_http_config](resources--http_loadbalancer--reference--group-015.md#canonical-3112103233111313-3303220101001113-3022220201103300-2310102331122013-0202002302022030-1310031130301100-1231313211312123-1131333023120221)
- [default_pool.advanced_options.circuit_breaker](resources--http_loadbalancer--reference--group-015.md#canonical-1031212032101202-3130331233121230-1022333233112301-0011311210120111-0130221013310320-2201133030023101-1021133320211103-0302331222320211)
- [default_pool.advanced_options.default_circuit_breaker](resources--http_loadbalancer--reference--group-015.md#canonical-0220231330001103-1312111311013211-2301111133301230-0123001132310332-0121313113212333-0001211212112332-2201120003102123-2223202111113133)
- [default_pool.advanced_options.disable_circuit_breaker](resources--http_loadbalancer--reference--group-015.md#canonical-3030013113131131-3103010310302203-2123231330311322-1231031222331233-1321202223300003-3001203202013202-3000310231301223-3101032101210223)
- [default_pool.advanced_options.disable_lb_source_ip_persistence](resources--http_loadbalancer--reference--group-015.md#canonical-3113100222020202-3133203310100220-0330332303302301-2031303313121332-0323121303222111-3223132221322032-0310310300022112-3030312000202003)
- [default_pool.advanced_options.disable_outlier_detection](resources--http_loadbalancer--reference--group-015.md#canonical-1300111032032202-0223330012333200-0003032323033300-1220012210221033-1221031213013231-0112210331000110-3103001312310113-2332322021231202)
- [default_pool.advanced_options.disable_proxy_protocol](resources--http_loadbalancer--reference--group-015.md#canonical-1323211300100232-0123211012202031-0232200001130201-1010020100302322-1122010113130203-0321323222203113-0330110330301230-2203311200223113)
- [default_pool.advanced_options.disable_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-0023103001332003-0303111121333030-1302202301203023-2300202001223132-3000303201112022-3312000113212101-3222132031232321-1110111201302101)
- [default_pool.advanced_options.enable_lb_source_ip_persistence](resources--http_loadbalancer--reference--group-015.md#canonical-1030023020110210-2012011100303020-3003123221320323-3003020312113033-1230031030231320-0130321233030101-1032013131300003-2201032222030012)
- [default_pool.advanced_options.enable_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-3122011012203130-2032023233101333-0322112222003112-2311302231102332-2113303332212233-1102000111031122-0303332031001202-2221302030020230)
- [default_pool.advanced_options.http1_config](resources--http_loadbalancer--reference--group-016.md#canonical-0302331112230021-3012012233210300-0212201311313112-1010303210232213-1332033331013232-3023212022323232-1321110310002103-2032000231020000)
- [default_pool.advanced_options.http2_options](resources--http_loadbalancer--reference--group-016.md#canonical-3113102323303220-1222012231233210-3132230022230202-2102202133132231-2231111233120333-0131210111011331-3220111022033123-3110000201103023)
- [default_pool.advanced_options.no_panic_threshold](resources--http_loadbalancer--reference--group-016.md#canonical-3010321022203332-1133303230302303-2331101002312232-0110221102133000-2100020203120200-3333120023011003-2312121111020320-1223232102333011)
- [default_pool.advanced_options.no_request_limit_per_connection](resources--http_loadbalancer--reference--group-016.md#canonical-3202011012323122-2310202220132112-1203221321333321-3232310230130231-0223011220002020-3131220013013302-1112021212030131-2212302110230313)
- [default_pool.advanced_options.outlier_detection](resources--http_loadbalancer--reference--group-016.md#canonical-3122120002230323-1331302213332300-0221022000101211-0233001322022311-3200121333120233-0120103223222330-0001121011111302-0100313133112132)
- [default_pool.advanced_options.proxy_protocol_v1](resources--http_loadbalancer--reference--group-016.md#canonical-0320332213030303-2110221221131312-1112220010132111-3111231331132203-1012323131311113-3203232033330032-3300331111032003-0132310330010030)
- [default_pool.advanced_options.proxy_protocol_v2](resources--http_loadbalancer--reference--group-016.md#canonical-2002203333000021-3223210132201331-1111301110300313-0211110110221123-3233022321110020-0320011110202330-3210010203121323-3010312231222231)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3112103233111313-3303220101001113-3022220201103300-2310102331122013-0202002302022030-1310031130301100-1231313211312123-1131333023120221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022133133311213-1010100302211133-2010312321200220-1323303331100201-0300023011222120-2331332300300103-0001033233101301-3210112013211133"></a>

## default_pool.advanced_options.auto_http_config — auto_http_config / 222303100233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.auto_http_config

<a id="canonical-2222110010010322-0311011300222101-1321130230021303-0012031022002130-1303301110220031-0232210132221202-3333222033322221-0011010213221202"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

Terraform syntax:

```terraform
auto_http_config = {}
```

<a id="canonical-0310033100103333-0213101110131310-2031233032030103-0011201310211002-1022012021130331-1230332331211103-2032030233101321-0321333301311021"></a>

## Direct properties — auto_http_config / 222303100233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100121010221332-1102232023301133-2231211332312000-0012000020203222-0222010023121313-1321100323312102-0132133330233310-1121000232022030"></a>

## Next pages — auto_http_config / 222303100233 / 4

- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1031212032101202-3130331233121230-1022333233112301-0011311210120111-0130221013310320-2201133030023101-1021133320211103-0302331222320211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221103120111131-2221033002320213-1232223002011222-1212322133031003-1011213203310033-3113233022103201-2323002010222032-3001120000202111"></a>

## default_pool.advanced_options.circuit_breaker — circuit_breaker / 233020020010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.circuit_breaker

<a id="canonical-3233111032021020-3230121313312001-1120032120103330-1002003010201123-0022010201321001-2032101002202110-2203010031122031-1212133000010020"></a>

Type: `"object"`. single nested block, Optional.

CircuitBreaker provides a mechanism for watching failures in upstream connections or requests and if
the failures reach a certain threshold, automatically fail subsequent requests which allows to apply
back pressure on downstream quickly.

Upstream description:

CircuitBreaker provides a mechanism for watching failures in upstream connections or requests and if
the failures reach a certain threshold, automatically fail subsequent requests which allows to apply
back pressure on downstream quickly.

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
circuit_breaker {
  # Configure direct properties listed below.
}
```

<a id="canonical-2101213313303200-2323310131233003-1010130021213010-0212022331021201-1023013220112301-3100000210131220-1212311322303010-1232101103133300"></a>

## Direct properties — circuit_breaker / 233020020010 / 3

<a id="canonical-0133112222330102-1032011100210302-0031311111011000-1212112133011122-0121002331103103-1222220330020120-1323113323301021-1120113021000211"></a>

<a id="canonical-3302121332111030-3311211112032323-2132130301331000-1331003210131111-0132211222131112-3012233101213111-2033100012110210-3320211122002033"></a>

## connection_limit property — circuit_breaker / 233020020010 / 4

Type: `"number"`. Optional.

The maximum number of connections that loadbalancer will establish to all hosts in an upstream
cluster. In practice this is only applicable to TCP and HTTP/1.1 clusters since HTTP/2 uses a single
connection to each host. Remove endpoint out of load balancing decision, if number of connections..

Upstream description:

The maximum number of connections that loadbalancer will establish to all hosts in an upstream
cluster. In practice this is only applicable to TCP and HTTP/1.1 clusters since HTTP/2 uses a single
connection to each host. Remove endpoint out of load balancing decision, if number of connections
reach connection limit.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32768),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-0333100220301322-1203012213120213-0310200333002312-2111110213201020-2003333032011303-0000102130122213-0323211021030211-0103130301313232"></a>

<a id="canonical-1313312300101110-1133011322223203-2212123132030010-0020003030021230-3001133121121230-0312033301322131-0111202231200122-0313311020231000"></a>

## max_requests property — circuit_breaker / 233020020010 / 5

Type: `"number"`. Optional.

The maximum number of requests that can be outstanding to all hosts in a cluster at any given time.
In practice this is applicable to HTTP/2 clusters since HTTP/1.1 clusters are governed by the
maximum connections (connection\_limit). Remove endpoint out of load balancing decision, if
requests..

Upstream description:

The maximum number of requests that can be outstanding to all hosts in a cluster at any given time.
In practice this is applicable to HTTP/2 clusters since HTTP/1.1 clusters are governed by the
maximum connections (connection\_limit). Remove endpoint out of load balancing decision, if requests
exceed this count.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32768),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-2111101031232003-3220223212221012-0223220231013213-2103022321333021-3230020330302000-3311303020313031-0120013023210221-3020132333002110"></a>

<a id="canonical-1230313122110020-3133123221122111-3212311311131022-3302310231012113-2303322211312003-0210313203212201-3313111333112133-0133213231312220"></a>

## pending_requests property — circuit_breaker / 233020020010 / 6

Type: `"number"`. Optional.

The maximum number of requests that will be queued while waiting for a ready connection pool
connection. Since HTTP/2 requests are sent over a single connection, this circuit breaker only comes
into play as the initial connection is created, as requests will be multiplexed immediately..

Upstream description:

The maximum number of requests that will be queued while waiting for a ready connection pool
connection. Since HTTP/2 requests are sent over a single connection, this circuit breaker only comes
into play as the initial connection is created, as requests will be multiplexed immediately
afterwards. For HTTP/1.1, requests are added to the list of pending requests whenever there aren’t
enough upstream connections available to immediately dispatch the request, so this circuit breaker
will remain in play for the lifetime of the process. Remove endpoint out of load balancing decision,
if pending request reach pending\_request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32768),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-2031113110313111-1030203300320212-2120323032103331-1303331112120122-3300132313212001-0320113200212013-0212313131121332-1013200031133311"></a>

<a id="canonical-3010002302130333-3030213121020212-3031313120030310-1300223303000211-2100213211113013-3231232021231233-1313203112012012-2002110131131133"></a>

## priority property — circuit_breaker / 233020020010 / 7

Type: `"string"`. Optional.

\[Enum: DEFAULT|HIGH\] Priority routing for each request. Different connection pools are used based
on the priority selected for the request. Also, circuit-breaker configuration at destination cluster
is chosen based on selected priority. Possible values are \`DEFAULT\`, \`HIGH\`. Defaults to
\`DEFAULT\`.

Upstream description:

Priority routing for each request. Different connection pools are used based on the priority
selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based
on selected priority.

Default routing mechanism High-Priority routing mechanism.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DEFAULT",
    "HIGH"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DEFAULT",
  "enum": [
    "DEFAULT",
    "HIGH"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3022133031223010-0321131220313232-0000323221221321-2100300112331301-0113112000323012-3123323321001211-2111130031313231-1122311122012123"></a>

<a id="canonical-3323230112100321-0303130121122130-0310201002332302-2312122300013303-2330232122203022-0233100201332302-3223310323222312-0331322031223220"></a>

## retries property — circuit_breaker / 233020020010 / 8

Type: `"number"`. Optional.

The maximum number of retries that can be outstanding to all hosts in a cluster at any given time.
Remove endpoint out of load balancing decision, if retries for request exceed this count.

Upstream description:

The maximum number of retries that can be outstanding to all hosts in a cluster at any given time.
Remove endpoint out of load balancing decision, if retries for request exceed this count.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32768),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-0310300102102130-1233210033022321-1021001311330212-2130200101110220-1003020210301132-0222012000021031-0213023111323123-1110113033021113"></a>

## Next pages — circuit_breaker / 233020020010 / 9

- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0220231330001103-1312111311013211-2301111133301230-0123001132310332-0121313113212333-0001211212112332-2201120003102123-2223202111113133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000232311312223-2203120013330211-0231111212311202-1303013322223310-0321223001233211-2311213332312130-0002231002003111-2011103033022203"></a>

## default_pool.advanced_options.default_circuit_breaker — default_circuit_breaker / 223203113230 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.default_circuit_breaker

<a id="canonical-2311223321130021-2322220301201322-3103101100201223-0120313123301311-2321110002000220-2312122131102120-0312000013010231-1231030012332031"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for default circuit breaker. Defaults to \`map\[\]\`. Server applies default
when omitted.

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

Terraform syntax:

```terraform
default_circuit_breaker = {}
```

<a id="canonical-2101321132111310-1320313232223033-3003021010222012-1310103131030030-2200023033321130-1223330032012210-1301231111121311-3233202133020022"></a>

## Direct properties — default_circuit_breaker / 223203113230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031132330200302-1311030031203231-1112220010300022-0201002232220010-0210213020332311-1311103101331012-1321313002232202-1122311011333022"></a>

## Next pages — default_circuit_breaker / 223203113230 / 4

- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3030013113131131-3103010310302203-2123231330311322-1231031222331233-1321202223300003-3001203202013202-3000310231301223-3101032101210223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221103133010012-0330133213101331-3331303022022110-1201103102032233-0021321123112330-1232230121312101-2223003113332031-3031221102201110"></a>

## default_pool.advanced_options.disable_circuit_breaker — disable_circuit_breaker / 221221011213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.disable_circuit_breaker

<a id="canonical-2012333203011332-2210122221231131-2320332013312011-2020133202331001-1001232330302121-1220010312001323-2323222202230333-3213103103132000"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable circuit breaker.

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

Terraform syntax:

```terraform
disable_circuit_breaker = {}
```

<a id="canonical-2132101030231233-3031001111010303-1201313032330012-3130220103023222-3101100100122323-1221322133210001-1031222030020030-3323021301100013"></a>

## Direct properties — disable_circuit_breaker / 221221011213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323132222010223-2201311132303123-3003331021321001-2312203032231302-3031303023302122-3220332031101211-2321332130133201-3032130131112001"></a>

## Next pages — disable_circuit_breaker / 221221011213 / 4

- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3113100222020202-3133203310100220-0330332303302301-2031303313121332-0323121303222111-3223132221322032-0310310300022112-3030312000202003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002232012210100-3231133132200110-1330011220213321-3231231321221323-1202023122321332-3122022311122211-0312130220301333-1110311120333310"></a>

## default_pool.advanced_options.disable_lb_source_ip_persistence — disable_lb_source_ip_persistence / 322111131333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.disable_lb_source_ip_persistence

<a id="canonical-3311300012320322-0121211301313010-1001312211033110-2112333221022112-2000221032332123-3333222120211231-3021032220312201-1010313013033110"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

IP address configuration

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
disable_lb_source_ip_persistence = {}
```

<a id="canonical-0002202122022302-2023131212133020-3032023313233121-2102313323012020-1232120330131130-0033213210223121-1333311001322010-2003013230321032"></a>

## Direct properties — disable_lb_source_ip_persistence / 322111131333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113321010210313-2200220323230200-2221000230330230-0021123000122333-1321122113231331-0220120102000123-3111301123001122-1023222022121112"></a>

## Next pages — disable_lb_source_ip_persistence / 322111131333 / 4

- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1300111032032202-0223330012333200-0003032323033300-1220012210221033-1221031213013231-0112210331000110-3103001312310113-2332322021231202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331212333331330-0303330200311221-1300322021222303-1132011100321201-1231220302303320-2213102330132303-3130300321220103-0112200033110323"></a>

## default_pool.advanced_options.disable_outlier_detection — disable_outlier_detection / 213332202332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.disable_outlier_detection

<a id="canonical-1323102303233300-0000323231203132-3203201112023100-2200332202320100-3123011331310031-1202232022310313-3112323231002230-3212232131030022"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable outlier detection. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

Terraform syntax:

```terraform
disable_outlier_detection = {}
```

<a id="canonical-3203221121223010-3033222013231323-0122122333322301-1000300321302003-1200220132321333-0022113020022210-0010130333021223-3131201111323010"></a>

## Direct properties — disable_outlier_detection / 213332202332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1112131300031313-3202101103333201-1011111322001122-0010310303112320-0333313300012113-3300201011003113-0123011203113302-1002132110130303"></a>

## Next pages — disable_outlier_detection / 213332202332 / 4

- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1323211300100232-0123211012202031-0232200001130201-1010020100302322-1122010113130203-0321323222203113-0330110330301230-2203311200223113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012313333222023-3011120130133111-2200213233232012-3303021010001302-0310332321120101-1213001131002013-1233230011233312-0310110133313102"></a>

## default_pool.advanced_options.disable_proxy_protocol — disable_proxy_protocol / 203301103201 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.disable_proxy_protocol

<a id="canonical-1023212212311013-1211022333031122-3310032221123012-3233100303313210-3223223220022321-2120110130031302-0323111013200320-1132031110022002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable proxy protocol.

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

Terraform syntax:

```terraform
disable_proxy_protocol = {}
```

<a id="canonical-3300133021222230-2312210132321302-0112003131120103-1232303213320311-3321112001032320-0030333211113200-2223332313010123-1000220313021102"></a>

## Direct properties — disable_proxy_protocol / 203301103201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1230033023113003-2130122203020301-1111011112211011-0101201330222313-1333122202003000-3221130102112203-3032311232113213-1212100332310212"></a>

## Next pages — disable_proxy_protocol / 203301103201 / 4

- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0023103001332003-0303111121333030-1302202301203023-2300202001223132-3000303201112022-3312000113212101-3222132031232321-1110111201302101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132012222000010-1313132023002102-0023320231033010-3321221322302021-1021112311002311-2000010322330023-3131120132032022-2022022231022000"></a>

## default_pool.advanced_options.disable_subsets — disable_subsets / 223333313131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.disable_subsets

<a id="canonical-0102132000213013-3322002221102033-3100320030031203-0311232200232330-2012212020100231-1022012212232312-1031101312203130-3022202110311302"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable subsets. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

Terraform syntax:

```terraform
disable_subsets = {}
```

<a id="canonical-0032303013311110-0221102232202323-0303012230230211-1121010233012223-2002202001301230-1103320130122220-3301113232121200-3321232120121233"></a>

## Direct properties — disable_subsets / 223333313131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323232333100133-0030211001230133-3211303020332310-3100320020133210-1113220021131101-2002313111033301-2311322330112102-3321303111123002"></a>

## Next pages — disable_subsets / 223333313131 / 4

- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1030023020110210-2012011100303020-3003123221320323-3003020312113033-1230031030231320-0130321233030101-1032013131300003-2201032222030012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311200122011123-2212003313022130-1321312010001201-1232203101133132-2320022101012003-0021332132021010-1213330210203002-1233031013233333"></a>

## default_pool.advanced_options.enable_lb_source_ip_persistence — enable_lb_source_ip_persistence / 233201031302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.enable_lb_source_ip_persistence

<a id="canonical-0202213131211030-2213121123310313-1232230113113203-1130231122332210-2311232210013010-3330101020133100-3210222033313003-0003033310313113"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

IP address configuration

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
enable_lb_source_ip_persistence = {}
```

<a id="canonical-3132010123333010-1223000012012102-0203232200111103-0011031002302110-3321211010210031-0212000203020032-0221121210302102-0130300333303220"></a>

## Direct properties — enable_lb_source_ip_persistence / 233201031302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321111300232322-1100023133233131-1203231012233121-0200323322301331-1321320301312301-2201303233333310-0211322010131102-1010233000023312"></a>

## Next pages — enable_lb_source_ip_persistence / 233201031302 / 4

- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3122011012203130-2032023233101333-0322112222003112-2311302231102332-2113303332212233-1102000111031122-0303332031001202-2221302030020230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310113220101322-3111031123201002-1233010321102100-2323002200302210-2031231320101113-0212210313232022-0131320032033001-1321310000123020"></a>

## default_pool.advanced_options.enable_subsets — enable_subsets / 330023333232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- default_pool.advanced_options.enable_subsets

<a id="canonical-2230231312121130-1032301133102111-3111133223133033-0203011113020101-0313311202133022-2203333022200212-1022103132100131-0013001010332212"></a>

Type: `"object"`. single nested block, Optional.

Configure subset OPTIONS for origin pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("endpoint_subsets"),
  validators.ConflictingObjectAttributes("any_endpoint",
    "default_subset"),
  validators.ConflictingObjectAttributes("any_endpoint",
    "fail_request"),
  validators.ConflictingObjectAttributes("default_subset",
    "fail_request")}
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
  "x-ves-oneof-field-fallback_policy_choice": "[\"any_endpoint\",\"default_subset\",\"fail_request\"]"
}
```

Terraform syntax:

```terraform
enable_subsets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0213132012033202-3210122100020302-0021212210231313-1323203012221011-3232203132112202-2000100013011023-1320001033021313-1122330333131201"></a>

## Direct properties — enable_subsets / 330023333232 / 3

- [any_endpoint](resources--http_loadbalancer--reference--group-015.md#canonical-3233312033110200-1111010030213011-3320211103320013-3233012111330131-3103031321022103-1332102101221111-0033311313130002-0332111310210222): complete subsection reference.

- [default_subset](resources--http_loadbalancer--reference--group-015.md#canonical-0311121031003212-1023312103112111-1223133333320310-2000132222222210-0111223230023131-1033323031033322-0303122132222203-2121101033013301): complete subsection reference.

- [endpoint_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-1323012021010023-2303102102321011-0331330121102101-3230020120302033-1222230310221220-3003022122203230-0012010311033122-2233222113222233): complete subsection reference.

- [fail_request](resources--http_loadbalancer--reference--group-015.md#canonical-1333011222302002-0211100111302313-0332200031311133-1223021130221131-2020220212132020-2332130122030302-0320320000130321-1012020200232321): complete subsection reference.

<a id="canonical-1121011231220102-0111131110320201-0300222101212212-1220120020031120-0310003202112200-1222200221210033-0023120202233033-0122310011303231"></a>

## Next pages — enable_subsets / 330023333232 / 4

- [default_pool.advanced_options.enable_subsets.any_endpoint](resources--http_loadbalancer--reference--group-015.md#canonical-3233312033110200-1111010030213011-3320211103320013-3233012111330131-3103031321022103-1332102101221111-0033311313130002-0332111310210222)
- [default_pool.advanced_options.enable_subsets.default_subset](resources--http_loadbalancer--reference--group-015.md#canonical-0311121031003212-1023312103112111-1223133333320310-2000132222222210-0111223230023131-1033323031033322-0303122132222203-2121101033013301)
- [default_pool.advanced_options.enable_subsets.endpoint_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-1323012021010023-2303102102321011-0331330121102101-3230020120302033-1222230310221220-3003022122203230-0012010311033122-2233222113222233)
- [default_pool.advanced_options.enable_subsets.fail_request](resources--http_loadbalancer--reference--group-015.md#canonical-1333011222302002-0211100111302313-0332200031311133-1223021130221131-2020220212132020-2332130122030302-0320320000130321-1012020200232321)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3233312033110200-1111010030213011-3320211103320013-3233012111330131-3103031321022103-1332102101221111-0033311313130002-0332111310210222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012011000000011-1211101212100230-3102000220333001-3031023210313032-1102132320132001-0011023030330331-3020213121203323-2000113003020321"></a>

## default_pool.advanced_options.enable_subsets.any_endpoint — any_endpoint / 001121030120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [default_pool.advanced_options.enable_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-3122011012203130-2032023233101333-0322112222003112-2311302231102332-2113303332212233-1102000111031122-0303332031001202-2221302030020230)
- default_pool.advanced_options.enable_subsets.any_endpoint

<a id="canonical-1000303030013302-0323313313321130-0112221323000311-0001222213133300-3232131301100302-1233322320321222-3212021133220311-0130101322100013"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
any_endpoint = {}
```

<a id="canonical-1302000303012222-0000332321003322-2021233200230300-0202101301103033-2031301013333212-1132030203131200-0321312031102233-0212101100133112"></a>

## Direct properties — any_endpoint / 001121030120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013002202201131-1203302203032001-3101220323302330-3120312121002330-2113221222122002-1113220031031301-0021202303003130-2001132021220310"></a>

## Next pages — any_endpoint / 001121030120 / 4

- [default_pool.advanced_options.enable_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-3122011012203130-2032023233101333-0322112222003112-2311302231102332-2113303332212233-1102000111031122-0303332031001202-2221302030020230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0311121031003212-1023312103112111-1223133333320310-2000132222222210-0111223230023131-1033323031033322-0303122132222203-2121101033013301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212323300030301-2101230323032332-1332323102200012-3013120303031313-0220020310302313-0123213122022323-0230200210322021-3021011022110232"></a>

## default_pool.advanced_options.enable_subsets.default_subset — default_subset / 213302233133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [default_pool.advanced_options.enable_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-3122011012203130-2032023233101333-0322112222003112-2311302231102332-2113303332212233-1102000111031122-0303332031001202-2221302030020230)
- default_pool.advanced_options.enable_subsets.default_subset

<a id="canonical-2210003123013121-2200321322330231-1321031021100120-3303020130133000-3022121011322111-3220312321220133-3110011210223020-0303013221212000"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default subset.

Upstream description:

Default Subset definition.

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
default_subset {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311313123331310-1102013012230320-3211300231323030-1110120332011303-2003011310212123-1332021112132230-0333323100321020-0303103132300202"></a>

## Direct properties — default_subset / 213302233133 / 3

- [default_subset](resources--http_loadbalancer--reference--group-015.md#canonical-3100312012123303-1223333121231133-0202011211212120-2012000330200330-1212201010233222-2202021010100311-0322010101113001-0002110020100032): complete subsection reference.

<a id="canonical-1123212101301300-2103001023213121-2031320020110013-1223203330012201-3112213100320110-1312011032130212-0002233332300221-1300121201010311"></a>

## Next pages — default_subset / 213302233133 / 4

- [default_pool.advanced_options.enable_subsets.default_subset.default_subset](resources--http_loadbalancer--reference--group-015.md#canonical-3100312012123303-1223333121231133-0202011211212120-2012000330200330-1212201010233222-2202021010100311-0322010101113001-0002110020100032)
- [default_pool.advanced_options.enable_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-3122011012203130-2032023233101333-0322112222003112-2311302231102332-2113303332212233-1102000111031122-0303332031001202-2221302030020230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3100312012123303-1223333121231133-0202011211212120-2012000330200330-1212201010233222-2202021010100311-0322010101113001-0002110020100032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131220322130332-3111301232232013-0020113120323013-3132311303111222-3020112012001312-0011233032100233-1003303221013311-3130230322133130"></a>

## default_pool.advanced_options.enable_subsets.default_subset.default_subset — default_subset / 332111021202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [default_pool.advanced_options.enable_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-3122011012203130-2032023233101333-0322112222003112-2311302231102332-2113303332212233-1102000111031122-0303332031001202-2221302030020230)
- [default_pool.advanced_options.enable_subsets.default_subset](resources--http_loadbalancer--reference--group-015.md#canonical-0311121031003212-1023312103112111-1223133333320310-2000132222222210-0111223230023131-1033323031033322-0303122132222203-2121101033013301)
- default_pool.advanced_options.enable_subsets.default_subset.default_subset

<a id="canonical-1111223122221220-2231311230203031-1310320333000310-1332302120233002-3131222012030031-0320111110202232-0233111213101133-2322211110312302"></a>

Type: `"object"`. single nested block, Optional.

List of key-value pairs that define default subset. Which gets used when route specifies no metadata
or no subset matching the metadata exists.

Upstream description:

List of key-value pairs that define default subset. Which gets used when route specifies no metadata
or no subset matching the metadata exists.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 32
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "32"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "32"
  }
}
```

Terraform syntax:

```terraform
default_subset {}
```

<a id="canonical-0103130101213021-3200221301113221-3213100230012030-1323331312323132-1011022322202002-0002102123112012-2033210011201323-1210120232003012"></a>

## Direct properties — default_subset / 332111021202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232022222213233-0233320210000232-2301112100033103-3332122200020121-1213331012011230-1223331002002230-2322010021222133-1003213321011211"></a>

## Next pages — default_subset / 332111021202 / 4

- [default_pool.advanced_options.enable_subsets.default_subset](resources--http_loadbalancer--reference--group-015.md#canonical-0311121031003212-1023312103112111-1223133333320310-2000132222222210-0111223230023131-1033323031033322-0303122132222203-2121101033013301)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1323012021010023-2303102102321011-0331330121102101-3230020120302033-1222230310221220-3003022122203230-0012010311033122-2233222113222233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101302323312131-1330310013001031-0332303233021231-3020221221211133-2101332303320330-2311022212121133-1023113132222233-3023010030233330"></a>

## default_pool.advanced_options.enable_subsets.endpoint_subsets — endpoint_subsets / 110132311001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.advanced_options](resources--http_loadbalancer--reference--group-015.md#canonical-2222122111331210-2233310023212212-1032010200303111-0132303220321311-0121033221020120-0233002303021120-1100000123201321-0212203223022100)
- [default_pool.advanced_options.enable_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-3122011012203130-2032023233101333-0322112222003112-2311302231102332-2113303332212233-1102000111031122-0303332031001202-2221302030020230)
- default_pool.advanced_options.enable_subsets.endpoint_subsets

<a id="canonical-1310130021001011-1033003231003212-2010212211212232-2033303132320232-3131022211323031-3030201220333200-1112320002232101-2331300232211331"></a>

Type: `"object"`. list nested block, Optional.

List of subset class. Subsets class is defined using list of keys. Every unique combination of
values of these keys form a subset within the class.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("keys")}
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
endpoint_subsets {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003300333101333-0010201133230333-0201111110232032-3031321221303130-2233221110311213-1201010113103001-2321013002300333-2223223001323032"></a>

## Direct properties — endpoint_subsets / 110132311001 / 3

<a id="canonical-0010310311122032-1331031001123323-0320021331122000-0312200133331120-0323102213112301-0333113231032111-0122130132200220-0300312330311310"></a>

<a id="canonical-2101223010310023-0313331230110020-2132111011310301-1122202202302003-0002212000001020-3130311001211133-0011130122221233-1332223102000211"></a>

## keys property — endpoint_subsets / 110132311001 / 4

Type: `["list", "string"]`. Optional.

List of keys that define a cluster subset class.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2213023322003032-2332123002330321-3103033032123131-2330101220210010-1130302321201131-1011310210110013-2333102120333221-0211220011000033"></a>

## Next pages — endpoint_subsets / 110132311001 / 5

- [default_pool.advanced_options.enable_subsets](resources--http_loadbalancer--reference--group-015.md#canonical-3122011012203130-2032023233101333-0322112222003112-2311302231102332-2113303332212233-1102000111031122-0303332031001202-2221302030020230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1333011222302002-0211100111302313-0332200031311133-1223021130221131-2020220212132020-2332130122030302-0320320000130321-1012020200232321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
