---
page_title: "xcsh_forward_proxy_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_forward_proxy_policy reference."
---

# xcsh_forward_proxy_policy reference

<a id="canonical-3201330013320233-3223133011203333-1331020110112323-3020203211003321-0001311333032023-0322321300132231-2322113222302212-0311032223032301"></a>

## rule_list.rules.port_matcher — port_matcher / 312223331202 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.port_matcher

<a id="canonical-2223210130031330-3322111110333000-0232120132010310-2310330231232132-3323020012121232-2033322201200330-0302302232022332-1031020232131121"></a>

Type: `"object"`. single nested block, Optional.

Port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true.

Upstream description:

A port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ports")}
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
port_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222202123030022-3222211120211132-2030000122211113-2030021300322300-2103331223331200-2021221222122312-0112212223330022-2002323122130101"></a>

## Direct properties — port_matcher / 312223331202 / 3

<a id="canonical-2031230022033200-1220133000313003-3310210211031010-2312003221303330-0313231111130210-2032003311133131-0102223231332201-0322012233103232"></a>

<a id="canonical-0312220001113113-3312110001123132-2113212102112321-2130130000010313-0303223211213101-0300031133021132-2302022033131222-0333030023133122"></a>

## invert_matcher property — port_matcher / 312223331202 / 4

Type: `"bool"`. Optional.

Invert Port Matcher. Invert the match result.

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

<a id="canonical-3301203210020311-3231130311123200-1321231101330031-1101233123123203-1002322130031303-1010101013211122-0102111221101023-1322013330103312"></a>

<a id="canonical-1101322213130031-2102211333020001-0121302003231303-3322000030022333-3123221030130112-2211013100223120-2121030313122011-3232022203032310"></a>

## ports property — port_matcher / 312223331202 / 5

Type: `["list", "string"]`. Optional.

List of strings, each of which is a single port value or a tuple of start and end port values
separated by '-'. The start and end values are considered to be part of the range.

Upstream description:

A list of strings, each of which is a single port value or a tuple of start and end port values
separated by "-". The start and end values are considered to be part of the range.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2111022221232200-0030330122321000-0203133231211002-3232133112223010-2122123020232311-0022302310010203-1022033323200211-2233213113231302"></a>

## Next pages — port_matcher / 312223331202 / 6

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-0312131012032312-0131110323211333-0123311001012011-2103130111123002-2020111100012001-3211231333323102-3111311122001102-1233223101130020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010020032233031-3032330131231012-0310231132323010-2130200112221320-3032301110103222-3032200132322111-1110010122120231-0332311000113211"></a>

## rule_list.rules.prefix_list — prefix_list / 021210202211 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.prefix_list

<a id="canonical-1211122323021202-0133301231103221-0220121222330021-2011313231220002-1302231303213021-2132320101120110-1120033302011123-0303330333031021"></a>

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
prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1300322222232211-2322000213221112-1122102013022102-3130102221101310-0313122310230313-1312013133020011-1200012031331111-0030221032031111"></a>

## Direct properties — prefix_list / 021210202211 / 3

<a id="canonical-3202103130210002-0303202322222201-3131323233020000-1032321012120302-3303222032300010-1310112310113321-1223132120323121-0021212203132213"></a>

<a id="canonical-2213311012111002-3211132112103031-3323022130202102-1021331123223000-0010010103030132-3301221013001033-1021033111330132-2012203102122232"></a>

## prefixes property — prefix_list / 021210202211 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1323303203223301-0220113133103030-1210112333231111-1321100103233213-1302120001112200-3112332130202210-1222101103032001-2321102130123020"></a>

## Next pages — prefix_list / 021210202211 / 5

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-0321220033202311-3123011210302013-3200210310012232-1230303120121032-2003133313330030-3321320211122231-2200231330033213-2211001300113003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101012300101023-2122312310212121-2022131123221103-1220020122201223-0002000232233012-1122210213130301-3033003002100233-3122232230322211"></a>

## rule_list.rules.tls_list — tls_list / 321302232001 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.tls_list

<a id="canonical-1210221123213132-0203202220231001-0221211321313010-3132220223103201-3330133110311231-3313102001230210-0012003131101132-3201023102012202"></a>

Type: `"object"`. single nested block, Optional.

DomainListType.

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
tls_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1002123321321322-3300302101100300-2222031112111122-0201210002312020-0302212200222223-2212103321333330-1010212201133121-3103133200121131"></a>

## Direct properties — tls_list / 321302232001 / 3

- [tls_list](resources--forward_proxy_policy--reference--group-002.md#canonical-1133323022102133-0031321013011302-2333300131111211-1232030103102131-3021311322113232-2331003013012020-3122100212320020-2220113100303010): complete subsection reference.

<a id="canonical-1101030311100030-3220033323202130-1021112301102333-1002030022310131-1010121112300230-1220023110133233-1230303212121232-1103010303211110"></a>

## Next pages — tls_list / 321302232001 / 4

- [rule_list.rules.tls_list.tls_list](resources--forward_proxy_policy--reference--group-002.md#canonical-1133323022102133-0031321013011302-2333300131111211-1232030103102131-3021311322113232-2331003013012020-3122100212320020-2220113100303010)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-1133323022102133-0031321013011302-2333300131111211-1232030103102131-3021311322113232-2331003013012020-3122100212320020-2220113100303010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020233102203123-1331131101100321-2302002203321213-2113223210121103-0301000310301332-3322020013033000-2133320021221212-1102033130302220"></a>

## rule_list.rules.tls_list.tls_list — tls_list / 003311131322 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [rule_list.rules.tls_list](resources--forward_proxy_policy--reference--group-002.md#canonical-0321220033202311-3123011210302013-3200210310012232-1230303120121032-2003133313330030-3321320211122231-2200231330033213-2211001300113003)
- rule_list.rules.tls_list.tls_list

<a id="canonical-0320220031133313-0130201231002103-1231201120112001-1331023212011310-2033322223202032-0301311201130303-1232213000333200-0320323303110202"></a>

Type: `"object"`. list nested block, Optional.

TLS Domains. Domains in SNI for TLS connections.

Upstream description:

Domains in SNI for TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("regex_value",
    "suffix_value")}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
tls_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0030203232110213-1323010232321110-2232200000011032-0023311101212330-0110010120301013-1300211123131220-3110002003100122-2121001223113332"></a>

## Direct properties — tls_list / 003311131322 / 3

<a id="canonical-2201020223300212-1232232212110133-3203021000333022-1301110310010322-1210022220123121-2222121002120100-3020223032103200-1303231320212113"></a>

<a id="canonical-0233100021313032-1303001221311132-2332210103121012-2120101011013322-2003110111303323-3031012223100003-0312100112200001-1120120010002122"></a>

## exact_value property — tls_list / 003311131322 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2333312232301233-2123321110303321-2300230212320103-1020203302000201-1000200322201031-1221131001330232-2221200101032223-1222132303101233"></a>

<a id="canonical-0321322300133113-3003123220231131-2101232032201100-0233301201312301-1111203311020103-3321111022231323-1110221133103011-1300222203100222"></a>

## regex_value property — tls_list / 003311131322 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2131322113012020-3133313202132232-2122222001332100-3103212103122010-1103322003212310-2212321111333220-2311013330300120-2002222202001222"></a>

<a id="canonical-0010301232302103-3001300222302300-3120313122220211-2231300133331013-3103002122202222-3301333223112133-0020002131112032-3113120210131331"></a>

## suffix_value property — tls_list / 003311131322 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0102003303333311-3210302013013120-3101030210002310-2000122022311033-3113223331021010-3033101203211321-2203023301000302-0320213030130133"></a>

## Next pages — tls_list / 003311131322 / 7

- [rule_list.rules.tls_list](resources--forward_proxy_policy--reference--group-002.md#canonical-0321220033202311-3123011210302013-3200210310012232-1230303120121032-2003133313330030-3321320211122231-2200231330033213-2211001300113003)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-3133120013122023-1010000233103233-3100211201111331-1132333033310100-1001321120000211-0013313110101113-0312320322323313-0031022220302232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122210130010012-1213210132311320-2201000133010213-0110223333233221-2333013211312303-0221130301030101-0013210110330102-1031231221331013"></a>

## rule_list.rules.url_category_list — url_category_list / 311203221022 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.url_category_list

<a id="canonical-1101212112133130-0012123030323210-2001102332220101-0320131122131103-2002132301022320-1233000110003110-0210213223123000-0312130202013230"></a>

Type: `"object"`. single nested block, Optional.

URL Category List Type. List of URL categories.

Upstream description:

List of URL categories.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url_categories")}
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
url_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1230322111323201-1221100133023021-1301213002200031-0302202021211033-3023302231022233-2321011301323220-2001310333233302-3230320133021232"></a>

## Direct properties — url_category_list / 311203221022 / 3

<a id="canonical-2210032102001010-0111200021312310-2023030213021110-0033210311202033-2201330003033213-1100110230031303-3101132231122103-3330303323110030"></a>

<a id="canonical-1231302301203303-2223012310300300-3001022103212113-2032310020302222-2130220123010102-1221203212300331-3133330023130222-1320002031101221"></a>

## url_categories property — url_category_list / 311203221022 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
UNCATEGORIZED|REAL\_ESTATE|COMPUTER\_AND\_INTERNET\_SECURITY|FINANCIAL\_SERVICES|BUSINESS\_AND\_ECONOMY|COMPUTER\_AND\_INTERNET\_INFO|AUCTIONS|SHOPPING|CULT\_AND\_OCCULT|TRAVEL|ABUSED\_DRUGS|ADULT\_AND\_PORNOGRAPHY|HOME\_AND\_GARDEN|MILITARY|SOCIAL\_NETWORKING|DEAD\_SITES|INDIVIDUAL\_STOCK\_ADVICE\_AND\_TOOLS|TRAINING\_AND\_TOOLS|DATING|SEX\_EDUCATION|RELIGION|ENTERTAINMENT\_AND\_ARTS|PERSONAL\_SITES\_AND\_BLOGS|LEGAL|LOCAL\_INFORMATION|STREAMING\_MEDIA|JOB\_SEARCH|GAMBLING|TRANSLATION|REFERENCE\_AND\_RESEARCH|SHAREWARE\_AND\_FREEWARE|PEER\_TO\_PEER|MARIJUANA|HACKING|GAMES|PHILOSOPHY\_AND\_POLITICAL\_ADVOCACY|WEAPONS|PAY\_TO\_SURF|HUNTING\_AND\_FISHING|SOCIETY|EDUCATIONAL\_INSTITUTIONS|ONLINE\_GREETING\_CARDS|SPORTS|SWIMSUITS\_AND\_INTIMATE\_APPAREL|QUESTIONABLE|KIDS|HATE\_AND\_RACISM|PERSONAL\_STORAGE|VIOLENCE|KEYLOGGERS\_AND\_MONITORING|SEARCH\_ENGINES|INTERNET\_PORTALS|WEB\_ADVERTISEMENTS|CHEATING|GROSS|WEB\_BASED\_EMAIL|MALWARE\_SITES|PHISHING\_AND\_OTHER\_FRAUDS|PROXY\_AVOIDANCE\_AND\_ANONYMIZERS|SPYWARE\_AND\_ADWARE|MUSIC|GOVERNMENT|NUDITY|NEWS\_AND\_MEDIA|ILLEGAL|CONTENT\_DELIVERY\_NETWORKS|INTERNET\_COMMUNICATIONS|BOT\_NETS|ABORTION|HEALTH\_AND\_MEDICINE|CONFIRMED\_SPAM\_SOURCES|SPAM\_URLS|UNCONFIRMED\_SPAM\_SOURCES|OPEN\_HTTP\_PROXIES|DYNAMICALLY\_GENERATED\_CONTENT|PARKED\_DOMAINS|ALCOHOL\_AND\_TOBACCO|PRIVATE\_IP\_ADDRESSES|IMAGE\_AND\_VIDEO\_SEARCH|FASHION\_AND\_BEAUTY|RECREATION\_AND\_HOBBIES|MOTOR\_VEHICLES|WEB\_HOSTING\]
URL Categories. List of URL categories to be selected. Possible values are \`UNCATEGORIZED\`,
\`REAL\_ESTATE\`, \`COMPUTER\_AND\_INTERNET\_SECURITY\`, \`FINANCIAL\_SERVICES\`,
\`BUSINESS\_AND\_ECONOMY\`, \`COMPUTER\_AND\_INTERNET\_INFO\`, \`AUCTIONS\`, \`SHOPPING\`,
\`CULT\_AND\_OCCULT\`, \`TRAVEL\`, \`ABUSED\_DRUGS\`, \`ADULT\_AND\_PORNOGRAPHY\`,
\`HOME\_AND\_GARDEN\`, \`MILITARY\`, \`SOCIAL\_NETWORKING\`, \`DEAD\_SITES\`,
\`INDIVIDUAL\_STOCK\_ADVICE\_AND\_TOOLS\`, \`TRAINING\_AND\_TOOLS\`, \`DATING\`, \`SEX\_EDUCATION\`,
\`RELIGION\`, \`ENTERTAINMENT\_AND\_ARTS\`, \`PERSONAL\_SITES\_AND\_BLOGS\`, \`LEGAL\`,
\`LOCAL\_INFORMATION\`, \`STREAMING\_MEDIA\`, \`JOB\_SEARCH\`, \`GAMBLING\`, \`TRANSLATION\`,
\`REFERENCE\_AND\_RESEARCH\`, \`SHAREWARE\_AND\_FREEWARE\`, \`PEER\_TO\_PEER\`, \`MARIJUANA\`,
\`HACKING\`, \`GAMES\`, \`PHILOSOPHY\_AND\_POLITICAL\_ADVOCACY\`, \`WEAPONS\`, \`PAY\_TO\_SURF\`,
\`HUNTING\_AND\_FISHING\`, \`SOCIETY\`, \`EDUCATIONAL\_INSTITUTIONS\`, \`ONLINE\_GREETING\_CARDS\`,
\`SPORTS\`, \`SWIMSUITS\_AND\_INTIMATE\_APPAREL\`, \`QUESTIONABLE\`, \`KIDS\`,
\`HATE\_AND\_RACISM\`, \`PERSONAL\_STORAGE\`, \`VIOLENCE\`, \`KEYLOGGERS\_AND\_MONITORING\`,
\`SEARCH\_ENGINES\`, \`INTERNET\_PORTALS\`, \`WEB\_ADVERTISEMENTS\`, \`CHEATING\`, \`GROSS\`,
\`WEB\_BASED\_EMAIL\`, \`MALWARE\_SITES\`, \`PHISHING\_AND\_OTHER\_FRAUDS\`,
\`PROXY\_AVOIDANCE\_AND\_ANONYMIZERS\`, \`SPYWARE\_AND\_ADWARE\`, \`MUSIC\`, \`GOVERNMENT\`,
\`NUDITY\`, \`NEWS\_AND\_MEDIA\`, \`ILLEGAL\`, \`CONTENT\_DELIVERY\_NETWORKS\`,
\`INTERNET\_COMMUNICATIONS\`, \`BOT\_NETS\`, \`ABORTION\`, \`HEALTH\_AND\_MEDICINE\`,
\`CONFIRMED\_SPAM\_SOURCES\`, \`SPAM\_URLS\`, \`UNCONFIRMED\_SPAM\_SOURCES\`,
\`OPEN\_HTTP\_PROXIES\`, \`DYNAMICALLY\_GENERATED\_CONTENT\`, \`PARKED\_DOMAINS\`,
\`ALCOHOL\_AND\_TOBACCO\`, \`PRIVATE\_IP\_ADDRESSES\`, \`IMAGE\_AND\_VIDEO\_SEARCH\`,
\`FASHION\_AND\_BEAUTY\`, \`RECREATION\_AND\_HOBBIES\`, \`MOTOR\_VEHICLES\`, \`WEB\_HOSTING\`.
Defaults to \`UNCATEGORIZED\`.

Upstream description:

List of URL categories to be selected.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1233321010321100-1123000112000223-3013031233130322-3021210111233110-1123012301022300-3303311110111200-0010211100330002-3213102323311000"></a>

## Next pages — url_category_list / 311203221022 / 5

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-0022013123220103-0231302233112311-3331201312000023-3031022310310321-1300302233001110-2213110210212201-2331202121302302-0031230221002103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112102332233122-3323131313012212-0220312310112012-2132313103010031-1300222202211313-0310301022312033-2212031312331332-0133111133012322"></a>

## timeouts — timeouts / 011110300322 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- timeouts

<a id="canonical-0312333203320031-2001323121301301-2113320111022333-1231122110212013-2112302103320203-2103010000220022-2200022121212202-0311001302331110"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3221230221333110-2002221133233211-3202021321132233-2202102323100110-3132203203103100-3230022222013110-3123030112332112-2032201311210322"></a>

## Direct properties — timeouts / 011110300322 / 3

<a id="canonical-2213123131121111-2003023001310022-3233110121302133-3332021310311333-1223232223222031-2220332203230002-1030331232011310-2210102102320223"></a>

<a id="canonical-1320002123133032-3031320000333023-0300020113103120-2230300012210031-3031210113323112-0103013312233312-3332213212102023-3313222120323233"></a>

## create property — timeouts / 011110300322 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0300322100323133-0201121113223010-2110130301231110-1203030033211001-1311113313112223-2330232130231302-3310211203102123-3031013202333322"></a>

<a id="canonical-2232002321230020-1320323313123130-1320312013202130-1223202122212233-3200302330230120-1321330201022311-1321203120002211-3122012103111032"></a>

## delete property — timeouts / 011110300322 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1303112313021102-3120111033121101-3203312232201202-3312120101300222-0022111202312021-2102001213232013-3022001023320321-3213031122111211"></a>

<a id="canonical-2030322331312101-2133233020332323-1330131203332301-0013322123012332-2132100110111123-0121102221033031-1212110013020120-3200031300131221"></a>

## read property — timeouts / 011110300322 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1133330220013103-1010331003022300-2013122001033321-3033022332310323-0010301233300202-0112133323331020-1231212230120213-3301131233111100"></a>

<a id="canonical-1110200221220000-2333010120203310-1221033332001010-2310011320211103-0212213000001232-1013212213131122-2032202331331110-1133213220131211"></a>

## update property — timeouts / 011110300322 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2031023313223313-3231033102321232-2322230220332332-0101331230013303-0321122212223313-3333011200033102-1333100001332113-0222102210233103"></a>

## Next pages — timeouts / 011110300322 / 8

- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
