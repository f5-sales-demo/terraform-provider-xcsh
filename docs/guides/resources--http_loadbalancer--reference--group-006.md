---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1203220120100121-0000201313131211-1311323210012232-3323230320130333-2311230202101032-2002032111030203-2013223332002012-2002020200213320"></a>

## api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list — ip_threat_category_list / 201213121110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list

<a id="canonical-3021103122302301-3223332333330312-1021223133323310-1123221312032002-2120033112013313-0100331332003030-0022311032320303-3330132021301330"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_threat_categories")}
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
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121212303022223-1212233123130001-2123002021221232-1312210333031030-1332231123333233-3132303000321132-2111103202121123-2311130120200123"></a>

## Direct properties — ip_threat_category_list / 201213121110 / 3

<a id="canonical-1023101331302200-1013132331133332-2103020321032031-1322100211021123-0332220020010101-0310201213022003-1333120330101322-2301200132021210"></a>

<a id="canonical-2303023003212311-0323122120333231-3302331010113133-1213230033302223-2300131210321212-0212012202131011-2231210311312111-0222100211112322"></a>

## ip_threat_categories property — ip_threat_category_list / 201213121110 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Upstream description:

The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2323222132200133-2002021223123133-3330013220233133-3133121211121313-2313120202020033-2123300133000003-0033120201300013-0001023221130220"></a>

## Next pages — ip_threat_category_list / 201213121110 / 5

- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0020130311132331-1120230211123201-2203233322212130-2002331231003300-0300110323300011-3231313303001130-2202322012200333-3212013201120332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002000031313012-2032221133131220-3311230310123211-0332321223322312-2032210302311103-3113223333322231-3122330312002312-1201103110102222"></a>

## api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher — tls_fingerprint_matcher / 232111322332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-3002123121211111-3100310330111301-1122010203333001-2133120112230322-2020003201021131-1233321133011302-1222022210311232-1330211330321231"></a>

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

<a id="canonical-2332231033021132-3330203000200120-2321112003033200-0113012323131110-2000111211222120-3102302311231332-0011220111003223-1111021002121123"></a>

## Direct properties — tls_fingerprint_matcher / 232111322332 / 3

<a id="canonical-2321123032230201-2121233131121132-1132020123101311-3301320330031010-0020330113300113-1210101033023111-3200210103033023-1210003000332031"></a>

<a id="canonical-3012131130322002-3202111000011231-2312002232302310-2110202322011303-3332103133111031-1120201312232212-1302302322131300-3313303302311133"></a>

## classes property — tls_fingerprint_matcher / 232111322332 / 4

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

<a id="canonical-1221213202111100-0031001330213111-2100333333103131-2332201113102131-3111020301030310-3211013313103020-3002231100320202-1223100020102032"></a>

<a id="canonical-2023202211113323-0021110012010101-3113012122113211-2122203310023311-3011210302021200-2010303010303002-0331322212100212-1131302203132302"></a>

## exact_values property — tls_fingerprint_matcher / 232111322332 / 5

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

<a id="canonical-0131311020233320-3301011020000333-1320222203233131-1100311023031131-3110031223131211-0230020013212330-0112002202100132-1112311123012321"></a>

<a id="canonical-0002303021130010-3130113300133311-0113021133323202-3233020310233011-0122231223301311-1313111031022111-1331313021011101-1213333333012221"></a>

## excluded_values property — tls_fingerprint_matcher / 232111322332 / 6

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

<a id="canonical-2300102122220210-0213303032201222-0333332230222111-3332001313313230-0302120123120320-1010210230213322-1221300312302212-1321120012232210"></a>

## Next pages — tls_fingerprint_matcher / 232111322332 / 7

- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3122101313113213-0230122323313020-2213032333110001-0122003110111023-2020313332102030-1103032103222301-1211233231220112-0322000110200233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021300102220020-0331030111112031-1103001330130001-3232221223210003-2122011012220233-1033110033133202-2222113110212331-2323031230021032"></a>

## api_protection_rules.api_groups_rules.metadata — metadata / 013111221001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- api_protection_rules.api_groups_rules.metadata

<a id="canonical-3221120032313213-3302302130220213-3002323201001110-3330300103321210-3132321311213011-1232210203331303-3312012301311223-2102012231010012"></a>

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

<a id="canonical-0121302311222332-1132230031033032-3031001102303202-2200032130103131-2313012222112301-1103011302101223-1033233212233001-2233031323302111"></a>

## Direct properties — metadata / 013111221001 / 3

<a id="canonical-1031230011103231-0131231213330220-0211120131013320-3032320230223003-1200301133030000-3321003221203002-3230023011103100-3202312023301032"></a>

<a id="canonical-2001110002112212-3301233303111001-3303331310023011-3010103212233211-3121121311300030-0221303312320210-0110223322322230-0313032010312011"></a>

## description_spec property — metadata / 013111221001 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-0233302333330133-2102010103310301-1133100322031022-2330131332122013-2013230120001213-2223121131033130-2001210210220133-1102231331232130"></a>

<a id="canonical-1232322300101322-0111212020310121-1201100322213103-0132311013311321-1311213303321010-2333333311002031-0130301011003301-1321112202313010"></a>

## name property — metadata / 013111221001 / 5

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

<a id="canonical-1320002301223213-2322001210030231-1301222202301111-2023003113011032-3122013112032013-2302102013223030-1020221333310011-1002120132002131"></a>

## Next pages — metadata / 013111221001 / 6

- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233223022133213-3021020302121320-3100000322121303-0300223012200101-2132332101330121-0130100133312332-2103120231323330-3230320211332201"></a>

## api_protection_rules.api_groups_rules.request_matcher — request_matcher / 111113201031 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- api_protection_rules.api_groups_rules.request_matcher

<a id="canonical-3202112031121121-3132032220312211-3332123222323102-0012021131101310-1312201213230111-2020002302300001-2222223020312001-3203002300232010"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for request matcher.

Upstream description:

Request conditions for matching a rule.

Receipt-pinned upstream constraints:

```json
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
request_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2121022021021213-0231013003201312-3123033330131302-1130320312022120-3023312100222222-0330323302003101-0210002232001033-0201031023023012"></a>

## Direct properties — request_matcher / 111113201031 / 3

- [cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-2330223010032220-2103331312331301-3013322330020123-1101220300222333-1023300323201133-0203302332130032-2200000120333102-1132221332331012): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-006.md#canonical-1102113023001102-3003123110221310-3002231200030103-3302100303220112-2000122003203002-3100301023002011-0000112102303003-0030220020333230): complete subsection reference.

- [jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-1000232202121331-0220321323200111-2211202201332332-0320003000223110-2322010033113302-2032133032000033-1020023231222310-0120011231031310): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-006.md#canonical-3110002111102211-3332012221211002-3323110233102203-2312123130220123-2010223211321200-2013123123333003-2323220020311331-1322220031300221): complete subsection reference.

<a id="canonical-3302232102201103-3212303222023103-0023302200312300-2313203030023322-3111223103023330-1322200101323201-3002312110012210-3003113221311010"></a>

## Next pages — request_matcher / 111113201031 / 4

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-2330223010032220-2103331312331301-3013322330020123-1101220300222333-1023300323201133-0203302332130032-2200000120333102-1132221332331012)
- [api_protection_rules.api_groups_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-006.md#canonical-1102113023001102-3003123110221310-3002231200030103-3302100303220112-2000122003203002-3100301023002011-0000112102303003-0030220020333230)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-1000232202121331-0220321323200111-2211202201332332-0320003000223110-2322010033113302-2032133032000033-1020023231222310-0120011231031310)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-3110002111102211-3332012221211002-3323110233102203-2312123130220123-2010223211321200-2013123123333003-2323220020311331-1322220031300221)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2330223010032220-2103331312331301-3013322330020123-1101220300222333-1023300323201133-0203302332130032-2200000120333102-1132221332331012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221333310301132-0312122121301122-2302131233132133-2103001133203233-0112030133301133-2301222101033030-3320130100321012-2313123320031201"></a>

## api_protection_rules.api_groups_rules.request_matcher.cookie_matchers — cookie_matchers / 022133203012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers

<a id="canonical-3333031233221103-2021321121212022-1323030023001111-0102211033222212-2223311220201000-3030212222000130-3111003302033000-0032100210232221"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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

<a id="canonical-0212332030001222-3002201332310131-0002210023001011-2130203122000031-2303012010113220-0110200230300312-3121122231301203-1212313203002130"></a>

## Direct properties — cookie_matchers / 022133203012 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-0111003230333002-0132232133331110-1233131331121333-0010101301032030-3232112101303302-2000223001323313-1013020000111011-3310322011321003): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-006.md#canonical-1031102332213300-1332223223321002-1100010010221113-3121200103330330-1322300120333322-1222100211300112-0000011022200100-2002101331011220): complete subsection reference.

<a id="canonical-1032110210113020-1213133303020332-3000122311032023-1300230112112210-0321321111012223-0033220200013311-1011021330213212-2121111233322213"></a>

<a id="canonical-1201302200313013-0322112210133330-2122212223301200-0330331012312023-3322220232301002-3012113020233130-3111100033210023-1322020303123030"></a>

## invert_matcher property — cookie_matchers / 022133203012 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--http_loadbalancer--reference--group-006.md#canonical-3233103230131132-3310103322011132-3331212013102300-3133013122223210-3213130111220121-3123301321123213-3033033131202312-1103122132100332): complete subsection reference.

<a id="canonical-2320103133002300-3232003202120320-2102000331101202-2331011003321330-2031202333133220-0101110220133120-3300320311301032-1233320031201211"></a>

<a id="canonical-2312312301231200-3302330022122123-1333203331030110-3113132113022213-1203123110112301-0321300321123110-2200200333013330-2100303112323123"></a>

## name property — cookie_matchers / 022133203012 / 5

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2301031131330332-2312323030103002-1210310333333020-3312122131321010-2202201222213021-2311203333333100-2012130122113012-0311030211223021"></a>

## Next pages — cookie_matchers / 022133203012 / 6

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-0111003230333002-0132232133331110-1233131331121333-0010101301032030-3232112101303302-2000223001323313-1013020000111011-3310322011321003)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_present](resources--http_loadbalancer--reference--group-006.md#canonical-1031102332213300-1332223223321002-1100010010221113-3121200103330330-1322300120333322-1222100211300112-0000011022200100-2002101331011220)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item](resources--http_loadbalancer--reference--group-006.md#canonical-3233103230131132-3310103322011132-3331212013102300-3133013122223210-3213130111220121-3123301321123213-3033033131202312-1103122132100332)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0111003230333002-0132232133331110-1233131331121333-0010101301032030-3232112101303302-2000223001323313-1013020000111011-3310322011321003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013020032200321-0020232112103101-2013013112332001-2210103111031210-0120300331321311-3110113130330032-3231211101033231-3012220020231202"></a>

## api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_present — check_not_present / 223210303133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-2330223010032220-2103331312331301-3013322330020123-1101220300222333-1023300323201133-0203302332130032-2200000120333102-1132221332331012)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-1312010102132022-0313321111231333-1221002001323230-1131312200231101-1301332223322111-1313102213213123-1121223030012321-0022032103320113"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-2202330333122233-3320320321110012-2211120000132031-2010200321113001-2203032210113001-0013331301123201-0232213031033202-0132003120311121"></a>

## Direct properties — check_not_present / 223210303133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201020222220302-0022233322330222-3132233011123010-0111100203322202-3200130033212302-3002200301320110-1030232230121022-2332003122312133"></a>

## Next pages — check_not_present / 223210303133 / 4

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-2330223010032220-2103331312331301-3013322330020123-1101220300222333-1023300323201133-0203302332130032-2200000120333102-1132221332331012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1031102332213300-1332223223321002-1100010010221113-3121200103330330-1322300120333322-1222100211300112-0000011022200100-2002101331011220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231321023122330-3102223130021033-3233223212302333-2300123312333233-1130001121310120-1033302112313331-3223102023223233-3202321132220013"></a>

## api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_present — check_present / 202033313020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-2330223010032220-2103331312331301-3013322330020123-1101220300222333-1023300323201133-0203302332130032-2200000120333102-1132221332331012)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-1012231221120102-3313120131231100-2000113301201213-2311321212311012-3120023313030331-1010023203231032-1000322000022002-0210032100002311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-0233231222230030-0232213212301122-2331021120303030-3221313311013233-3200123120310011-3231313002311211-1113223333233033-1301002001112023"></a>

## Direct properties — check_present / 202033313020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310331302322223-3212230311231312-0133120033030200-0010010330300301-3000232033322321-3022113110130320-2033012220122113-2332332003211232"></a>

## Next pages — check_present / 202033313020 / 4

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-2330223010032220-2103331312331301-3013322330020123-1101220300222333-1023300323201133-0203302332130032-2200000120333102-1132221332331012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3233103230131132-3310103322011132-3331212013102300-3133013122223210-3213130111220121-3123301321123213-3033033131202312-1103122132100332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312102221321021-2123032300101320-1230011013103233-0202122012001103-1210020210023123-3221130321123101-1020210331223100-3221213031103322"></a>

## api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item — item / 233020013303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-2330223010032220-2103331312331301-3013322330020123-1101220300222333-1023300323201133-0203302332130032-2200000120333102-1132221332331012)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item

<a id="canonical-1033213322230012-3000111311012130-0113201110020312-2302130131233021-2102121102022102-1330311313323222-1010310223111301-0302320012003322"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-2310122211332000-0210120330011332-0012301000312203-1011102122233001-1011122231222030-3333200120030033-3122101213131212-0221200320331300"></a>

## Direct properties — item / 233020013303 / 3

<a id="canonical-2231232212231013-3230223020200013-0301200233120221-3021311300133201-2320120221113131-1222033111231013-3333223031131120-0332211103130312"></a>

<a id="canonical-0120220131300331-1311210121301301-2010122021333030-1331201322230020-3110210020000030-1300223200303310-3022132130103021-2131331232320233"></a>

## exact_values property — item / 233020013303 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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

<a id="canonical-1113233223001223-2103130223321220-1231323122231012-0011031013031303-1113331011201333-0211100203202100-1111220330210132-1211111300111130"></a>

<a id="canonical-1031210211130200-1322302001322302-2301223200132100-2122101101003131-2331000223100212-2303132300101121-3320020011200301-0122203212310232"></a>

## regex_values property — item / 233020013303 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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

<a id="canonical-0302010123023133-1130023131211201-3033321133121322-1210220120113230-3013001121201132-3101200102013210-0230310013231313-0112212233303321"></a>

<a id="canonical-0210003222102023-0323120200032110-3100030102110120-2130131030112011-1220112011201311-3131310033210003-2231032203211223-2123222112100200"></a>

## transformers property — item / 233020013303 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1320200232131023-0311203322110222-3132333121322303-0310222011131311-2332121002323121-2113123132313111-2011303230113012-1321023300322220"></a>

## Next pages — item / 233020013303 / 7

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-2330223010032220-2103331312331301-3013322330020123-1101220300222333-1023300323201133-0203302332130032-2200000120333102-1132221332331012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1102113023001102-3003123110221310-3002231200030103-3302100303220112-2000122003203002-3100301023002011-0000112102303003-0030220020333230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012202231000211-2313323100102231-0101121210102012-3322022102231012-3101221332013123-3210022130000113-1322202323310220-1101222231000203"></a>

## api_protection_rules.api_groups_rules.request_matcher.headers — headers / 032112300223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- api_protection_rules.api_groups_rules.request_matcher.headers

<a id="canonical-2321212010023130-2110111210222313-1030122320033202-2300220331101030-2331201310221331-2310223301312321-1203130202010210-0221021022131020"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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

<a id="canonical-3201002011131331-2003330323111020-0000023122023132-3330322130310001-3120221302010223-1033301000320001-1001021022213101-1020213113012132"></a>

## Direct properties — headers / 032112300223 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-0111031002110111-3222121311112200-2220301230201120-0020333233222101-3213030323222201-1101231321333300-1211332011132132-3023112320200120): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-006.md#canonical-3230303120123113-0020202120332113-2121230020230322-0230012020203032-0101130301203023-0010310023130213-2200331031010222-2202310221130030): complete subsection reference.

<a id="canonical-2231111002330320-0220322000302300-1021233230312332-3220030232022230-0301033323311010-0322013311210330-0001230222222103-0322211202013333"></a>

<a id="canonical-0320012312330133-2121302220222102-3330013203231230-2303010021033232-1120310020001101-1031133221211303-0032220232211321-2133033333023203"></a>

## invert_matcher property — headers / 032112300223 / 4

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-006.md#canonical-2112120021310311-3101013232321221-0123222010003011-0302201210223320-0033220030331032-2112220112130022-3221213001330030-3201023033003032): complete subsection reference.

<a id="canonical-0332301132133300-2322011000330331-2031111203113330-2130103101021112-0320333102033011-1313303232300113-0021110121221020-2233223311202102"></a>

<a id="canonical-3013023203232103-0113212320203122-1002210012321002-2023332123212122-3033201303211113-0010233332032202-1121103110030002-0000033232301022"></a>

## name property — headers / 032112300223 / 5

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-2230221313012001-3233003233010302-2010033020303222-0000213230033323-0233333132312302-3230132200121121-3320230131032302-0001022112023122"></a>

## Next pages — headers / 032112300223 / 6

- [api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-0111031002110111-3222121311112200-2220301230201120-0020333233222101-3213030323222201-1101231321333300-1211332011132132-3023112320200120)
- [api_protection_rules.api_groups_rules.request_matcher.headers.check_present](resources--http_loadbalancer--reference--group-006.md#canonical-3230303120123113-0020202120332113-2121230020230322-0230012020203032-0101130301203023-0010310023130213-2200331031010222-2202310221130030)
- [api_protection_rules.api_groups_rules.request_matcher.headers.item](resources--http_loadbalancer--reference--group-006.md#canonical-2112120021310311-3101013232321221-0123222010003011-0302201210223320-0033220030331032-2112220112130022-3221213001330030-3201023033003032)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0111031002110111-3222121311112200-2220301230201120-0020333233222101-3213030323222201-1101231321333300-1211332011132132-3023112320200120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010331110022131-1032211201232112-2201113333101003-0220110002131303-0201020122121223-1130033311211032-0013020021033231-3002220223201012"></a>

## api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present — check_not_present / 031203231220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-006.md#canonical-1102113023001102-3003123110221310-3002231200030103-3302100303220112-2000122003203002-3100301023002011-0000112102303003-0030220020333230)
- api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present

<a id="canonical-2022002231102101-2232202200110002-2232323121222302-1130110311200302-1121223103102023-3132322101303213-2200110023221220-3110012121300101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-1200330300201300-1310203110221220-3122213131112132-0231303202222310-1112102122312103-1000301123120303-3022022032220011-0002110322130031"></a>

## Direct properties — check_not_present / 031203231220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133231001230112-2112312332213322-1011210311213120-1120032031313322-1131121231210121-1301300311003121-3013003001112231-3333323323023001"></a>

## Next pages — check_not_present / 031203231220 / 4

- [api_protection_rules.api_groups_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-006.md#canonical-1102113023001102-3003123110221310-3002231200030103-3302100303220112-2000122003203002-3100301023002011-0000112102303003-0030220020333230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3230303120123113-0020202120332113-2121230020230322-0230012020203032-0101130301203023-0010310023130213-2200331031010222-2202310221130030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200011331103303-3011121211101202-1021032031313301-3112033030131313-0133113310112202-3311012213222011-0201031011110021-0103001323211032"></a>

## api_protection_rules.api_groups_rules.request_matcher.headers.check_present — check_present / 330110330001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-006.md#canonical-1102113023001102-3003123110221310-3002231200030103-3302100303220112-2000122003203002-3100301023002011-0000112102303003-0030220020333230)
- api_protection_rules.api_groups_rules.request_matcher.headers.check_present

<a id="canonical-1102212111033222-1132220212012232-2030233330211003-0230320100001230-1111311303312023-1331301323120313-2313333031202022-3320010020223122"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-0323213112321303-1202202323120102-2303210302313100-3131221323032012-2100332302220203-3313022322102002-1123203113012200-3202322323302223"></a>

## Direct properties — check_present / 330110330001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230121213312222-0000323303023132-0123303132023212-2101213301020000-2103030102023011-1021030003221232-0220020323102003-0131000213130321"></a>

## Next pages — check_present / 330110330001 / 4

- [api_protection_rules.api_groups_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-006.md#canonical-1102113023001102-3003123110221310-3002231200030103-3302100303220112-2000122003203002-3100301023002011-0000112102303003-0030220020333230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2112120021310311-3101013232321221-0123222010003011-0302201210223320-0033220030331032-2112220112130022-3221213001330030-3201023033003032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322030211002310-0223310312200023-2221230102212310-1200012003300101-3201331001231210-1202111322021212-1300210320301120-2122223013210003"></a>

## api_protection_rules.api_groups_rules.request_matcher.headers.item — item / 310030110331 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-006.md#canonical-1102113023001102-3003123110221310-3002231200030103-3302100303220112-2000122003203002-3100301023002011-0000112102303003-0030220020333230)
- api_protection_rules.api_groups_rules.request_matcher.headers.item

<a id="canonical-2303131220031203-1201031120032212-3202133320101033-2012301012211311-2323220331330022-3201213311300130-3230102231232021-1303203312100030"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-2013230320113012-0230212131130230-3303132212230003-0013130002130113-2121313202301011-2030230212213003-0313213010300220-2003231000231023"></a>

## Direct properties — item / 310030110331 / 3

<a id="canonical-2103121322010200-3002030011101211-2100011321211330-0121303310022101-0303010003313120-3102303203002121-0111221333010211-0321021212311010"></a>

<a id="canonical-2100221021211012-3100113213213123-2303312113133301-2230202122120320-2203003023001311-0020232130133200-2201122120301122-1200133113302002"></a>

## exact_values property — item / 310030110331 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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

<a id="canonical-1122101233200020-3032201223120311-1132310220030100-2122011031030300-2002233130002231-2311212231201303-2020202103320032-3122130230020210"></a>

<a id="canonical-3213310333030323-0032113002001322-2023002221311222-1003212323112010-2033332210303003-3321313322010121-3201113032311202-1312023211210331"></a>

## regex_values property — item / 310030110331 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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

<a id="canonical-3013001102233130-3102220111203120-2302130122211213-3010022220033322-0120030221312321-2033223021301312-0212223032120301-0101031100133003"></a>

<a id="canonical-3331102103031210-1122303311332010-1121332123102223-3032331310132133-2302120322102212-0223101003021302-3213132130003321-2223111020230233"></a>

## transformers property — item / 310030110331 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0023302231113323-0211021313121301-2130313221033003-0202313023212133-0123310322211003-2232023223212232-0300212133120320-3131200303012220"></a>

## Next pages — item / 310030110331 / 7

- [api_protection_rules.api_groups_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-006.md#canonical-1102113023001102-3003123110221310-3002231200030103-3302100303220112-2000122003203002-3100301023002011-0000112102303003-0030220020333230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1000232202121331-0220321323200111-2211202201332332-0320003000223110-2322010033113302-2032133032000033-1020023231222310-0120011231031310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330122111302212-2131221203332012-3002211323333110-3103013020133100-1301023321011312-3310302232213301-1020311002112210-3302123123000323"></a>

## api_protection_rules.api_groups_rules.request_matcher.jwt_claims — jwt_claims / 032122200113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims

<a id="canonical-1232013133302300-1300101221213231-3122310133003330-3212120233300022-1203321122133210-0132200233133123-0201201012313002-1231011220022022"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
jwt_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-0113012330312133-1323010001302113-0030101000132003-3331310220010001-1333301133033101-1021120301100111-3010323301002020-1320102230302120"></a>

## Direct properties — jwt_claims / 032122200113 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-1221022321103011-2323212110312021-3301313311210313-2313023113311122-0220130312322000-3020003110301201-1211210030123330-1113023111131212): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-006.md#canonical-0112212123012302-2133330201021303-3203100112120223-2131131102302123-1330230312012001-2032230101300130-2113032230100012-3120131002213302): complete subsection reference.

<a id="canonical-3211320332202322-2122010322301020-0200333313031203-2113001101233021-1121203003102231-3112132222330103-2201302303222012-2002330310113013"></a>

<a id="canonical-2030102212200201-1300102302220131-1113123222012203-2130100220112122-2120301212131320-0303323231031111-2322230030001022-0131302213333103"></a>

## invert_matcher property — jwt_claims / 032122200113 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-006.md#canonical-3221201132223132-0322013020201110-0020312210130003-2033223123010021-3301323211313200-2213110120203221-1033223223301323-3332323222322321): complete subsection reference.

<a id="canonical-3222132112301332-2201102220013330-2110030010211230-0223132213201322-2232301232112113-1213102101222321-1322233021100010-1203212201021021"></a>

<a id="canonical-1132212211230222-2320323202020001-3210133112011010-1121002111333321-2203111130211203-0103000313031020-1232310201333202-0120110202303032"></a>

## name property — jwt_claims / 032122200113 / 5

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1210211030333301-0212122020300112-0312002222231001-2301131230013112-3231313301202001-1101022003222331-1221000030120303-3301120103120303"></a>

## Next pages — jwt_claims / 032122200113 / 6

- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-1221022321103011-2323212110312021-3301313311210313-2313023113311122-0220130312322000-3020003110301201-1211210030123330-1113023111131212)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present](resources--http_loadbalancer--reference--group-006.md#canonical-0112212123012302-2133330201021303-3203100112120223-2131131102302123-1330230312012001-2032230101300130-2113032230100012-3120131002213302)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item](resources--http_loadbalancer--reference--group-006.md#canonical-3221201132223132-0322013020201110-0020312210130003-2033223123010021-3301323211313200-2213110120203221-1033223223301323-3332323222322321)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1221022321103011-2323212110312021-3301313311210313-2313023113311122-0220130312322000-3020003110301201-1211210030123330-1113023111131212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212121012210133-1022230021130111-3130032103201123-3010300302033123-3120312031101322-2333023120001221-2321032100230323-0322001033031132"></a>

## api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present — check_not_present / 202103022101 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-1000232202121331-0220321323200111-2211202201332332-0320003000223110-2322010033113302-2032133032000033-1020023231222310-0120011231031310)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-3032032302123302-0113330330110111-3131110303311223-0231031021321101-1113111112102102-0313012133022022-0310321311101032-2013320103313330"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-2220003322233111-3112312020310323-2033030023122321-1230332222032310-1323223103320020-2123111323302331-1120013010222023-3321211000223102"></a>

## Direct properties — check_not_present / 202103022101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032000230033312-0033330112213023-1220333320213130-2013122120321132-3020010200310223-2001330211112303-3223003022110211-3011220310133000"></a>

## Next pages — check_not_present / 202103022101 / 4

- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-1000232202121331-0220321323200111-2211202201332332-0320003000223110-2322010033113302-2032133032000033-1020023231222310-0120011231031310)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0112212123012302-2133330201021303-3203100112120223-2131131102302123-1330230312012001-2032230101300130-2113032230100012-3120131002213302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213133311130123-1312232321022023-3332112211113120-2231101030012223-2102321230123111-1123310123202302-3332233203213202-0310133021020212"></a>

## api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present — check_present / 312130200321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-1000232202121331-0220321323200111-2211202201332332-0320003000223110-2322010033113302-2032133032000033-1020023231222310-0120011231031310)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present

<a id="canonical-3321322012113230-1222111002110231-0032333313111000-3203313231320313-2133021311213320-2312332132313131-2212030231201000-0121233033103201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-1300211123000000-2303223213332021-0123112220312312-0221120110002111-3112023013233100-2313302313321231-3200131330222222-2110300231121312"></a>

## Direct properties — check_present / 312130200321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211323222320233-2203022000213103-3100011210231133-0213100123021202-1333203131323012-2200033023000100-0223201010303132-0300233103303013"></a>

## Next pages — check_present / 312130200321 / 4

- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-1000232202121331-0220321323200111-2211202201332332-0320003000223110-2322010033113302-2032133032000033-1020023231222310-0120011231031310)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3221201132223132-0322013020201110-0020312210130003-2033223123010021-3301323211313200-2213110120203221-1033223223301323-3332323222322321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012320313322333-2330102100311212-3110032011110232-0110331222313221-2012130230222100-3310101103303332-1112123332233303-1013003211323320"></a>

## api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item — item / 312120110132 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-1000232202121331-0220321323200111-2211202201332332-0320003000223110-2322010033113302-2032133032000033-1020023231222310-0120011231031310)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item

<a id="canonical-2213313330003003-2303333221201323-0120323223313300-3321231331023320-1332202312012023-2031201122211011-1020302223133112-2112211102233020"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-0323220203301213-0312120002111100-0322130000101331-1203220213033333-0221323001322200-2111323201233313-1013222302121220-0003231003131301"></a>

## Direct properties — item / 312120110132 / 3

<a id="canonical-1313332132000210-2123102233210031-2003323122210302-2101301321133030-2001233332003130-2013130203020101-1132311301322123-3133231332220001"></a>

<a id="canonical-1210212233003233-1213121013230132-2221000223330033-1223023220001131-0213013203122010-1110301302133031-0310012233001202-3302322331331313"></a>

## exact_values property — item / 312120110132 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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

<a id="canonical-3111202232103230-2122212001131211-0003303132310123-0103210232333302-1300320202330210-0122230223311031-0030133320333200-0113300223221213"></a>

<a id="canonical-0022203303120210-3020212331103031-0132330331221130-0200320033312233-1012322300010220-3121210233111033-0220103122311131-0302320330102310"></a>

## regex_values property — item / 312120110132 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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

<a id="canonical-3211312132002112-2012201321103111-1013211322330302-0131311330222022-1131323020022013-2233111322102332-2320120311201022-2012001320030220"></a>

<a id="canonical-2333013310122013-3132030102000203-0123002202100200-0030310022023132-3033213133301023-2001111130111310-2021130331002330-0232221331131132"></a>

## transformers property — item / 312120110132 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2031333300232213-1120220222103000-2302213112032102-1031030223311330-1023211220323330-1021120201011311-1213030332330322-2131200001131311"></a>

## Next pages — item / 312120110132 / 7

- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-1000232202121331-0220321323200111-2211202201332332-0320003000223110-2322010033113302-2032133032000033-1020023231222310-0120011231031310)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3110002111102211-3332012221211002-3323110233102203-2312123130220123-2010223211321200-2013123123333003-2323220020311331-1322220031300221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311033321031220-1213031301213231-0320301312103100-3321310302321003-1331132023231220-2111332123202211-3210231223113023-2210020301323020"></a>

## api_protection_rules.api_groups_rules.request_matcher.query_params — query_params / 211021113233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- api_protection_rules.api_groups_rules.request_matcher.query_params

<a id="canonical-1222322313132320-0121201111103001-0310330222302331-3300121002133013-0212320213332223-1010201113301323-3323331210333222-3003102122300323"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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

<a id="canonical-1203203032022221-1303003020130320-2013221120001321-1130313110000013-1131212012110213-2033222222300320-2120110100021022-0122113011303112"></a>

## Direct properties — query_params / 211021113233 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-2301301231130220-3022330120203122-3312122012121223-1011021100300311-0011212233230003-2230130122211123-0210310102112213-2332303323300312): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-006.md#canonical-2302200113210110-3323122133201313-2232122033020121-3210013030221001-3020120301002231-1133211000313033-0033131013321131-0211000020012301): complete subsection reference.

<a id="canonical-1023302032031321-3220320320003102-0210020023313132-0020320303330203-0120103302122321-2210131113011322-3221122222221002-1011031230022000"></a>

<a id="canonical-0120201132213330-1002100311223032-0122330121131002-1222321111332110-2222011331011233-3303323121223310-2120232131011000-3031002302300231"></a>

## invert_matcher property — query_params / 211021113233 / 4

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-006.md#canonical-0112023020121132-1033320220310300-0310233323221212-2021032103333032-0332013022221331-3111033112322011-1211013002322032-2223000113002332): complete subsection reference.

<a id="canonical-2032330003113313-1121003133303113-0110012011000311-0021033110102130-1133011003113213-3010211323001223-3202221212020122-0120223312303110"></a>

<a id="canonical-3302130101101100-0002221231101332-2031131031131102-0211133110333012-0002030110110222-2030200321120020-0011231311113323-1130123002221313"></a>

## key property — query_params / 211021113233 / 5

Type: `"string"`. Optional.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2103021002330232-1131123122321011-1030320303311331-3233233302112112-2331320130132011-2212210321121333-1032223303013222-0012020300112222"></a>

## Next pages — query_params / 211021113233 / 6

- [api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-2301301231130220-3022330120203122-3312122012121223-1011021100300311-0011212233230003-2230130122211123-0210310102112213-2332303323300312)
- [api_protection_rules.api_groups_rules.request_matcher.query_params.check_present](resources--http_loadbalancer--reference--group-006.md#canonical-2302200113210110-3323122133201313-2232122033020121-3210013030221001-3020120301002231-1133211000313033-0033131013321131-0211000020012301)
- [api_protection_rules.api_groups_rules.request_matcher.query_params.item](resources--http_loadbalancer--reference--group-006.md#canonical-0112023020121132-1033320220310300-0310233323221212-2021032103333032-0332013022221331-3111033112322011-1211013002322032-2223000113002332)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2301301231130220-3022330120203122-3312122012121223-1011021100300311-0011212233230003-2230130122211123-0210310102112213-2332303323300312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230223112211011-3320333311201311-3320203221110101-3331103320310013-0112221122100302-1320213232000101-0200020332030221-0111303230003013"></a>

## api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_present — check_not_present / 231000302211 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-3110002111102211-3332012221211002-3323110233102203-2312123130220123-2010223211321200-2013123123333003-2323220020311331-1322220031300221)
- api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_present

<a id="canonical-2010110013031313-2222210322211232-3300323112312312-0123332133223110-1130112221200023-2300101130000212-1033310313033323-2032320122020113"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-1131200200111131-3311310020323301-1310003023210113-2203131122000331-2211133301013303-1131102130113113-0132131230311132-1133110311132211"></a>

## Direct properties — check_not_present / 231000302211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323021203223010-2300022320020202-3033333023321312-1222100223310230-0031010001100221-1322103303232322-3132303213013131-1301103310133331"></a>

## Next pages — check_not_present / 231000302211 / 4

- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-3110002111102211-3332012221211002-3323110233102203-2312123130220123-2010223211321200-2013123123333003-2323220020311331-1322220031300221)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2302200113210110-3323122133201313-2232122033020121-3210013030221001-3020120301002231-1133211000313033-0033131013321131-0211000020012301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033111331023213-3312031011221311-0321113011332032-0313122023203012-3210233221301213-1010230023013030-3232022332323010-3303201210101001"></a>

## api_protection_rules.api_groups_rules.request_matcher.query_params.check_present — check_present / 012100212103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-3110002111102211-3332012221211002-3323110233102203-2312123130220123-2010223211321200-2013123123333003-2323220020311331-1322220031300221)
- api_protection_rules.api_groups_rules.request_matcher.query_params.check_present

<a id="canonical-1111032032201011-0021322011313120-3130001213330213-1322321110131001-3032320210220032-3133012211223313-3313101013010022-1230132301102313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-0130101101001023-2110302002012031-3133000132021011-0013301322023212-3303121033320312-3100232203102220-0330323022010123-1323000223023331"></a>

## Direct properties — check_present / 012100212103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121331001212210-1003302111331102-2330322013111132-3232301113233312-2000102232012123-3331123102233230-0000230211223122-3003300131312120"></a>

## Next pages — check_present / 012100212103 / 4

- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-3110002111102211-3332012221211002-3323110233102203-2312123130220123-2010223211321200-2013123123333003-2323220020311331-1322220031300221)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0112023020121132-1033320220310300-0310233323221212-2021032103333032-0332013022221331-3111033112322011-1211013002322032-2223000113002332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113230003020010-2133203300210203-3030011200010320-0223232100003332-0311131232310331-3323013200020001-3213201131333111-2133112022120011"></a>

## api_protection_rules.api_groups_rules.request_matcher.query_params.item — item / 001210033301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-3110002111102211-3332012221211002-3323110233102203-2312123130220123-2010223211321200-2013123123333003-2323220020311331-1322220031300221)
- api_protection_rules.api_groups_rules.request_matcher.query_params.item

<a id="canonical-3003321001211231-1231132103323022-0110323000030201-3132201321311111-2122203330002203-2100131030011211-3000133210310210-1331131203010213"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-2021320202222203-2223231233003333-2102022302330021-0023331202032311-2121311323100013-3302311202101100-3113301201003211-3213023332222121"></a>

## Direct properties — item / 001210033301 / 3

<a id="canonical-1212320103321131-1322000031212221-2202331200311201-2332011311111231-3313323321010230-1310122300013113-0133320220100031-2223313303023012"></a>

<a id="canonical-1103110321233320-3011003112332312-2321123131303220-1322301133202110-3123020203113033-0023200330100010-2023120130103003-1300303231103120"></a>

## exact_values property — item / 001210033301 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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

<a id="canonical-1221311003102213-3003120000333013-0033233211133003-1033002202212021-1100310010203330-2012020112011222-0002331031321021-0301121100232013"></a>

<a id="canonical-1330211131301221-3130212322323002-2022301002100220-1322121112200232-0030201231133210-1232303303331201-0013023210013010-1011312113300133"></a>

## regex_values property — item / 001210033301 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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

<a id="canonical-3311201032012023-2232321312203001-3010302311010033-3021231102320112-0300013320300022-3302022031233013-1002033000011222-1320300020200301"></a>

<a id="canonical-2333330131312311-2122322233100331-1113230021121022-1122030100020321-1203022102330012-3112231010211213-3101112131013000-2112020131300300"></a>

## transformers property — item / 001210033301 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0302212333110023-1110213122300221-3301000233033310-0123002020001111-3031233122032313-2021023100232310-1302101131103302-2323221133001122"></a>

## Next pages — item / 001210033301 / 7

- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-3110002111102211-3332012221211002-3323110233102203-2312123130220123-2010223211321200-2013123123333003-2323220020311331-1322220031300221)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212101002021330-2311301211310331-0110120033203310-2002330003132102-3013302102213100-1320231030000021-1000213323323000-3301303122001111"></a>

## api_rate_limit — api_rate_limit / 303301002331 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- api_rate_limit

<a id="canonical-0233031301102130-1130230220030210-3300120230310332-2002022322233223-2023100032330031-0223303022120030-3300222102202013-2011031031103200"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: api\_rate\_limit, disable\_rate\_limit, rate\_limit; Default: disable\_rate\_limit\] Path-
or API-group-scoped rate limiting. Define server\_url\_rules or api\_endpoint\_rules and choose
inline\_rate\_limiter for an inline limit, or ref\_rate\_limiter for a stored rate-limiter
reference.

Upstream description:

Path- or API-group-scoped rate limiting. Define server\_url\_rules or api\_endpoint\_rules and
choose inline\_rate\_limiter for an inline limit, or ref\_rate\_limiter for a stored rate-limiter
reference.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("bypass_rate_limiting_rules",
    "custom_ip_allowed_list"),
  validators.ConflictingObjectAttributes("bypass_rate_limiting_rules",
    "ip_allowed_list"),
  validators.ConflictingObjectAttributes("bypass_rate_limiting_rules",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "ip_allowed_list"),
  validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("ip_allowed_list",
    "no_ip_allowed_list")}
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
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"bypass_rate_limiting_rules\",\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]"
}
```

OneOf alternatives in this subsection:

- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0233031301102130-1130230220030210-3300120230310332-2002022322233223-2023100032330031-0223303022120030-3300222102202013-2011031031103200)
- [disable_rate_limit](resources--http_loadbalancer--reference--group-018.md#canonical-1013223133201022-1313332221032232-1331200200123112-3000301121301112-0033001302203101-2312333301123320-2111320003320121-2311310003323022)
- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-1212231232113102-0000210123322230-2321331212231011-1223210321030302-2020032100331231-0213102112123132-3222021323133333-1010333102020321)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_rate_limit {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301312302121121-0013220311102021-3132313030133131-3231111230202000-2213032121032131-3103321031111113-3101022033001133-1200313013103002"></a>

## Direct properties — api_rate_limit / 303301002331 / 3

- [api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220): complete subsection reference.

- [bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001): complete subsection reference.

- [custom_ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-2313230220001103-3212301300010212-3013123133132033-1130203100112210-1121313032010300-0313200120132330-2230131221332112-1011333232233322): complete subsection reference.

- [ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-1201210302313203-2211300232121320-0323321131002323-0331122322021300-2230032123313000-3330020132002300-3030111322130233-3130310012323300): complete subsection reference.

- [no_ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-1030302110202203-3022230330030230-2123232230310122-2013030321120121-3102121110303121-1120312201020213-1322321131223223-1130233033033321): complete subsection reference.

- [server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103): complete subsection reference.

<a id="canonical-0332230132132110-3233201311310021-1322002011120002-2321320201111000-2312230031201112-3200332222321222-0033132133332313-1022100311212300"></a>

## Next pages — api_rate_limit / 303301002331 / 4

- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.custom_ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-2313230220001103-3212301300010212-3013123133132033-1130203100112210-1121313032010300-0313200120132330-2230131221332112-1011333232233322)
- [api_rate_limit.ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-1201210302313203-2211300232121320-0323321131002323-0331122322021300-2230032123313000-3330020132002300-3030111322130233-3130310012323300)
- [api_rate_limit.no_ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-1030302110202203-3022230330030230-2123232230310122-2013030321120121-3102121110303121-1120312201020213-1322321131223223-1130233033033321)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011212333003130-0132202111201102-3110103033003313-0301220102103110-0321112020021202-1202331201201212-1303330232303123-2203313211303021"></a>

## api_rate_limit.api_endpoint_rules — api_endpoint_rules / 011332223000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- api_rate_limit.api_endpoint_rules

<a id="canonical-1132131113333103-2022130102320022-2230110113330212-0132031113100120-1231101222210233-3302112221010212-0100032312311122-1323013233103113"></a>

Type: `"object"`. list nested block, Optional.

Ordered endpoint-specific rate-limit rules. Each rule must choose exactly one rate\_limiter\_choice:
inline\_rate\_limiter or ref\_rate\_limiter.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("api_endpoint_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain"),
  validators.ConflictingListObjectAttributes("inline_rate_limiter",
    "ref_rate_limiter")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

Terraform syntax:

```terraform
api_endpoint_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3211003130301301-0121023210021232-3000222030233133-3211220112022120-2111022101231333-1010003012013000-1113013310331220-2303311012022203"></a>

## Direct properties — api_endpoint_rules / 011332223000 / 3

- [any_domain](resources--http_loadbalancer--reference--group-006.md#canonical-1102221321330123-1112312231203321-2212320130100310-2020322220032221-2332230110121010-3112302312320212-3022321223202320-3021100213113332): complete subsection reference.

- [api_endpoint_method](resources--http_loadbalancer--reference--group-006.md#canonical-3211220233321223-0231100031032022-2012223000121203-1232010310013022-2210231321223033-3000133331210021-1023230032112210-2002333100210103): complete subsection reference.

<a id="canonical-2331022231313010-2113003022132122-2210231312301110-3003323021102313-0012011003003223-2002113203031230-1031120331212231-1312000113102132"></a>

<a id="canonical-0210130313110330-2130103222121211-1001230012003222-0223300221112033-2231113120231323-1322012221123130-0213320101321330-1131301223232001"></a>

## api_endpoint_path property — api_endpoint_rules / 011332223000 / 4

Type: `"string"`. Optional.

API Endpoint. The endpoint (path) of the request.

Upstream description:

The endpoint (path) of the request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

- [client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022): complete subsection reference.

- [inline_rate_limiter](resources--http_loadbalancer--reference--group-006.md#canonical-1121233023131220-1003111220221011-0133310311113000-1021131213123031-3323030010330300-0200200213101021-1233030130032333-0013322100312031): complete subsection reference.

- [ref_rate_limiter](resources--http_loadbalancer--reference--group-006.md#canonical-3031121110211321-0222223200032133-1223103211312023-1012222201100010-0323202231130000-2323201330321110-1222232311021302-3202202311112033): complete subsection reference.

- [request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122): complete subsection reference.

<a id="canonical-3203121311310230-2100333322210103-3231301201012120-2322101111331110-0113333001101211-1133203230213233-3123210021321023-0022012232220120"></a>

<a id="canonical-2002230221310120-3130031313131220-3102222103011130-0110303211113121-0200223213233200-3323012220021123-1011203133322122-2333103310212221"></a>

## specific_domain property — api_endpoint_rules / 011332223000 / 5

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Provider validators and defaults (from schema source):

```go
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
    "format": "fqdn",
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-3121130301101013-2223102020032333-2311013012200020-3003300201131031-3320323200110223-0131203100312113-3110222023310232-2212132022331231"></a>

## Next pages — api_endpoint_rules / 011332223000 / 6

- [api_rate_limit.api_endpoint_rules.any_domain](resources--http_loadbalancer--reference--group-006.md#canonical-1102221321330123-1112312231203321-2212320130100310-2020322220032221-2332230110121010-3112302312320212-3022321223202320-3021100213113332)
- [api_rate_limit.api_endpoint_rules.api_endpoint_method](resources--http_loadbalancer--reference--group-006.md#canonical-3211220233321223-0231100031032022-2012223000121203-1232010310013022-2210231321223033-3000133331210021-1023230032112210-2002333100210103)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-006.md#canonical-1121233023131220-1003111220221011-0133310311113000-1021131213123031-3323030010330300-0200200213101021-1233030130032333-0013322100312031)
- [api_rate_limit.api_endpoint_rules.ref_rate_limiter](resources--http_loadbalancer--reference--group-006.md#canonical-3031121110211321-0222223200032133-1223103211312023-1012222201100010-0323202231130000-2323201330321110-1222232311021302-3202202311112033)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1102221321330123-1112312231203321-2212320130100310-2020322220032221-2332230110121010-3112302312320212-3022321223202320-3021100213113332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233230222011330-2021223112332323-3032200210000330-1012333221031233-2312233221212313-3223202301332232-1323332330230221-1010221110213221"></a>

## api_rate_limit.api_endpoint_rules.any_domain — any_domain / 201322322101 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- api_rate_limit.api_endpoint_rules.any_domain

<a id="canonical-1110303132222231-0210011010211110-1302312302110020-2010010132211001-0000103222331131-1231220010102200-0001112322312021-2123332133301000"></a>

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

<a id="canonical-3200100121202300-2023123031302322-1010332031222321-1230303300212100-1013211131013023-3220032021320211-2110223221011223-3211212100112003"></a>

## Direct properties — any_domain / 201322322101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122121213030211-0020202102112101-3000110303011003-1101021231121330-2013122001302020-3120232121030133-1211301210313103-1231020122120332"></a>

## Next pages — any_domain / 201322322101 / 4

- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3211220233321223-0231100031032022-2012223000121203-1232010310013022-2210231321223033-3000133331210021-1023230032112210-2002333100210103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103122013102103-3032010220333310-3203210322123122-2323132101213103-1302131212100132-3033331211022003-1132021212010300-3110213230310130"></a>

## api_rate_limit.api_endpoint_rules.api_endpoint_method — api_endpoint_method / 211002302221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- api_rate_limit.api_endpoint_rules.api_endpoint_method

<a id="canonical-3133330222100221-3123200201332122-2301210032320032-2220232302020100-0130003323201012-3230201033330132-0010230123200332-1031122132002111"></a>

Type: `"object"`. single nested block, Optional.

HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

Upstream description:

A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
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
api_endpoint_method {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230210332023122-2022023323310312-3031301020300222-1302102323103132-3300213023012301-0013000022132022-3233010332112201-0020132120301003"></a>

## Direct properties — api_endpoint_method / 211002302221 / 3

<a id="canonical-1332313221233021-0302301130011130-2210003302132300-0221303303112220-1333311101202123-1120020012010322-3321332311203013-1103223322130330"></a>

<a id="canonical-0301112133203323-0231312011203332-3003233132311221-0122001312022203-2113303333030000-1233210300212130-2113200203113120-1003131203031101"></a>

## invert_matcher property — api_endpoint_method / 211002302221 / 4

Type: `"bool"`. Optional.

Invert Method Matcher. Invert the match result.

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

<a id="canonical-3003211102231313-2123223301120102-2200131103301203-3122013023013303-2332010312202003-3023002130322022-3001221323121323-2002010023100123"></a>

<a id="canonical-0230021121022130-3331110202230133-2211102133032121-0101212213321301-2020203123100032-1121121200311211-0013023202303302-3323100011003010"></a>

## methods property — api_endpoint_method / 211002302221 / 5

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

List of methods values to match against.

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

<a id="canonical-1302103133002131-3130320220212132-1332221031213202-2020010102103111-2133013210021201-0002010123232202-0002212311232021-2322211323313313"></a>

## Next pages — api_endpoint_method / 211002302221 / 6

- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310121103222333-1321203010112201-1320031000022231-1331230311002231-2013331001000012-1003103212120331-1200301000323110-3223122321220311"></a>

## api_rate_limit.api_endpoint_rules.client_matcher — client_matcher / 101212111210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- api_rate_limit.api_endpoint_rules.client_matcher

<a id="canonical-3102022023020021-3321211122201030-1330012333122230-1002010222001122-2201312311031222-0131232321011013-3232103110111321-2000223122010220"></a>

Type: `"object"`. single nested block, Optional.

Client Matcher. Client conditions for matching a rule.

Upstream description:

Client conditions for matching a rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_client",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("client_selector",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
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
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

Terraform syntax:

```terraform
client_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301221332120022-3223321121321033-3011001012313132-2003202111201022-0313022110102132-1210013333213313-2013223223021001-0011012001120002"></a>

## Direct properties — client_matcher / 101212111210 / 3

- [any_client](resources--http_loadbalancer--reference--group-006.md#canonical-0012221131313133-3123300110213130-3211112200313130-3000312112230301-2022023132233213-1201000232011321-2132222323102032-2211011231100101): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-006.md#canonical-3111231132022032-2012201213130111-3210122323121103-2130110002110211-0003213133322113-0102101110002231-1020231322230313-0033122012310213): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-006.md#canonical-0231131102210101-2133013103213130-0321113113303011-0210321100230110-0212110222110300-1203311112312300-3030332313100211-3001132332112302): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2133101103210301-1133300211033320-2220320131022112-0003022120230021-0102310230123222-0002201330021303-0302221113020123-0331103300033123): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-006.md#canonical-3100320111032122-2112302021131110-2213020113220103-2123310222003200-1320103210121220-2203021031111200-0013233111012101-0233131002202122): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-3200322322220212-1113313312121132-0022020102332110-0003210012123332-0232201320020023-1331003101031110-1133111233110332-1032321212200130): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-006.md#canonical-1312202323020121-2330122112103213-2102331212330331-0311311101121312-0331221200021133-2111102011123223-0111002302103312-2030312132201000): complete subsection reference.

- [ip_threat_category_list](resources--http_loadbalancer--reference--group-006.md#canonical-1231130011002201-0201220210132022-1202011011322320-3222221230011211-1220312323300101-1113033232232300-1300130111323030-0113012032033021): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-3201202030331112-3102102112311122-2223110300113033-2231001222303211-3001320300131233-3202223231022132-1213133331012100-3122020302033123): complete subsection reference.

<a id="canonical-0012110233003230-1021331322012202-2210312231020020-3013300121210010-3222301002003121-1220002310111102-2113011200133111-1130303312200320"></a>

## Next pages — client_matcher / 101212111210 / 4

- [api_rate_limit.api_endpoint_rules.client_matcher.any_client](resources--http_loadbalancer--reference--group-006.md#canonical-0012221131313133-3123300110213130-3211112200313130-3000312112230301-2022023132233213-1201000232011321-2132222323102032-2211011231100101)
- [api_rate_limit.api_endpoint_rules.client_matcher.any_ip](resources--http_loadbalancer--reference--group-006.md#canonical-3111231132022032-2012201213130111-3210122323121103-2130110002110211-0003213133322113-0102101110002231-1020231322230313-0033122012310213)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_list](resources--http_loadbalancer--reference--group-006.md#canonical-0231131102210101-2133013103213130-0321113113303011-0210321100230110-0212110222110300-1203311112312300-3030332313100211-3001132332112302)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2133101103210301-1133300211033320-2220320131022112-0003022120230021-0102310230123222-0002201330021303-0302221113020123-0331103300033123)
- [api_rate_limit.api_endpoint_rules.client_matcher.client_selector](resources--http_loadbalancer--reference--group-006.md#canonical-3100320111032122-2112302021131110-2213020113220103-2123310222003200-1320103210121220-2203021031111200-0013233111012101-0233131002202122)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-3200322322220212-1113313312121132-0022020102332110-0003210012123332-0232201320020023-1331003101031110-1133111233110332-1032321212200130)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list](resources--http_loadbalancer--reference--group-006.md#canonical-1312202323020121-2330122112103213-2102331212330331-0311311101121312-0331221200021133-2111102011123223-0111002302103312-2030312132201000)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list](resources--http_loadbalancer--reference--group-006.md#canonical-1231130011002201-0201220210132022-1202011011322320-3222221230011211-1220312323300101-1113033232232300-1300130111323030-0113012032033021)
- [api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-3201202030331112-3102102112311122-2223110300113033-2231001222303211-3001320300131233-3202223231022132-1213133331012100-3122020302033123)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0012221131313133-3123300110213130-3211112200313130-3000312112230301-2022023132233213-1201000232011321-2132222323102032-2211011231100101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212231323032330-0210131110310113-2032121311223323-1001121133033212-3010023100212120-0203212002310123-0030211220200131-2331220213101213"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.any_client — any_client / 331332031020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.any_client

<a id="canonical-3323010020322300-0313023030332200-0222203212100021-0221102222200111-1113321102022103-1003311313201032-1000103221012213-0011103032313231"></a>

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
any_client = {}
```

<a id="canonical-1113133113011101-0002100001112210-1032110021202011-2201123023012030-0013121013200133-1203313312131320-0131331001103223-3231133121123220"></a>

## Direct properties — any_client / 331332031020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313333030011132-2301311130123003-0033102221232230-2222122100120213-1333300200233323-3231322220100313-2232030223012101-2320211031332130"></a>

## Next pages — any_client / 331332031020 / 4

- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3111231132022032-2012201213130111-3210122323121103-2130110002110211-0003213133322113-0102101110002231-1020231322230313-0033122012310213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333333131203031-2312030222211001-2003320021112031-3021010233322030-0020223213110320-3022331011112132-3133101222331112-1013112223020122"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.any_ip — any_ip / 211000321110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.any_ip

<a id="canonical-1232103001103323-2123100133003223-1310020131111032-1202311103230302-3031220211232332-0220203121021033-2331102021013203-1220031123021212"></a>

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
any_ip = {}
```

<a id="canonical-1302202312020022-0302122120321101-2203303001120001-1121003002302201-0103022003021101-0001012231222210-0333330320221333-2233011012231120"></a>

## Direct properties — any_ip / 211000321110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003201203123333-3202100033123230-0312121333322033-3123002112322001-0103111122222013-2021213321030032-2000301333021302-1310332112033311"></a>

## Next pages — any_ip / 211000321110 / 4

- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0231131102210101-2133013103213130-0321113113303011-0210321100230110-0212110222110300-1203311112312300-3030332313100211-3001132332112302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030113210101323-2313111130210122-3300013220333110-2301213131033001-1130201320212210-0230000302321220-2232011332130123-1320003023010202"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.asn_list — asn_list / 311233001321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_list

<a id="canonical-0003333232231212-1232330320133012-3232102103300113-2113000113003231-1200203303123002-0103031223033233-2231230101310100-0203002312213022"></a>

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

<a id="canonical-2010112023110113-3001010113231223-1020122132222112-1330012221112022-3122033122001020-2310132111302322-2213220231130003-2221222030022133"></a>

## Direct properties — asn_list / 311233001321 / 3

<a id="canonical-1200121330130020-1002313201331323-2112331223232002-0230311131021021-0222333323333301-1211022100310010-3032000202221212-3032312102200032"></a>

<a id="canonical-0320330103322012-0101203120111133-2322233023202333-2102023002021212-2311110320132303-1200120023313110-2133320003312321-2222112032130213"></a>

## as_numbers property — asn_list / 311233001321 / 4

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

<a id="canonical-3011113302220010-3213123113300230-3230103223231020-1013030200212201-2033000232332221-3012011203300121-3333121023110103-1021103021102212"></a>

## Next pages — asn_list / 311233001321 / 5

- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2133101103210301-1133300211033320-2220320131022112-0003022120230021-0102310230123222-0002201330021303-0302221113020123-0331103300033123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311032022131202-0101110000131111-1123003120311233-3200303103033031-0120322302011013-0300013131232033-0301120300222002-2200021221231022"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher — asn_matcher / 311100011303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher

<a id="canonical-2033113121211121-1130333112313221-3322012121000020-2302223201001213-0221013300130330-1333232121320320-0200101303301322-1303321123000000"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2323020213312001-3121023031323231-2030131012023200-3323020232031332-0100121011201312-1201101300311222-2233200110333112-1130221310301220"></a>

## Direct properties — asn_matcher / 311100011303 / 3

- [asn_sets](resources--http_loadbalancer--reference--group-006.md#canonical-1200210320110211-0310131300222120-0012333132101230-3223331322111202-3300201212002110-0201012210223030-0202323130200002-1013322023301323): complete subsection reference.

<a id="canonical-2100202213332113-1233303030211222-0032211001210210-1033012032012221-2023210212131020-2222030210000032-2123301110112020-0201131003130103"></a>

## Next pages — asn_matcher / 311100011303 / 4

- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets](resources--http_loadbalancer--reference--group-006.md#canonical-1200210320110211-0310131300222120-0012333132101230-3223331322111202-3300201212002110-0201012210223030-0202323130200002-1013322023301323)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1200210320110211-0310131300222120-0012333132101230-3223331322111202-3300201212002110-0201012210223030-0202323130200002-1013322023301323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110321203220223-0131320121132103-3231312301033003-1233023311110000-2010331310320310-2100300310011011-3300322010331010-1111131330212021"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets — asn_sets / 130002012330 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2133101103210301-1133300211033320-2220320131022112-0003022120230021-0102310230123222-0002201330021303-0302221113020123-0331103300033123)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-2203322122202032-1310023322220032-3103302022313223-1031311330322120-1221030232321123-2201101132112320-3012023311230213-2103333211021130"></a>

Type: `"object"`. list nested block, Optional.

List of references to bgp\_asn\_set objects.

Upstream description:

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

<a id="canonical-1033203013002031-3202322323212102-0131301210303021-3233231223202103-0303101331323232-2302023310000001-3300201312220321-0133331313203232"></a>

## Direct properties — asn_sets / 130002012330 / 3

<a id="canonical-3010012003333110-0123320332330311-1030330330111023-0133233322100212-0211123322002231-2313100310113110-1033312023213321-3210222101200120"></a>

<a id="canonical-2033302200300202-2303210121110100-1211113002010230-1203110032012202-3320133112230121-3023030113121211-1202212130330013-3302112300130231"></a>

## kind property — asn_sets / 130002012330 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-3033013223013232-1123010333022010-0032312230211230-3301322113212230-3333301220222130-1323102312130330-0002211112313213-3010012301121201"></a>

<a id="canonical-1232200131112302-2312203003022130-2033322132231023-3212031220222320-2031123313011002-3012201131222021-2221311203102300-3102333011100031"></a>

## name property — asn_sets / 130002012330 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3112221202003310-2030202111032232-2023221323223101-0020333131331022-3333022313311103-0021113113313012-0111231010001113-3312313332303013"></a>

<a id="canonical-0213033212133330-1230331321321231-0233201232221001-1020112001120020-0321103312112101-1011330202212000-0022130333210210-2331211121110012"></a>

## namespace property — asn_sets / 130002012330 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-1231202131222002-2130110101321030-0012233222113321-0212213223113301-1110003323021212-1011122121120201-3320000303101120-0022122031113202"></a>

<a id="canonical-2023300233333232-0222001000113020-1023130323022113-0332231131002331-1111320222221302-1013100122310120-0003133010010223-3120120110200100"></a>

## tenant property — asn_sets / 130002012330 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2020123320122213-1233123013332321-3301333201012200-3320011232311230-2032303122010221-3233301230231133-0233023330122102-2322001212233311"></a>

<a id="canonical-0203033023010331-3312321033102130-1311223213201013-2322322133120031-2003210233012030-2201032110211320-2323201232301002-3122102320322303"></a>

## uid property — asn_sets / 130002012330 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-0122320110202231-1033000002113011-0213320302202200-0322122332012203-1110210203121023-1103322211013232-3012320301222000-2333132100233200"></a>

## Next pages — asn_sets / 130002012330 / 9

- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2133101103210301-1133300211033320-2220320131022112-0003022120230021-0102310230123222-0002201330021303-0302221113020123-0331103300033123)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3100320111032122-2112302021131110-2213020113220103-2123310222003200-1320103210121220-2203021031111200-0013233111012101-0233131002202122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102200230130010-2330010033303202-1330103103102131-2031032112231221-2301023130321022-1311111111100100-1331111201110212-3333212223331231"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.client_selector — client_selector / 032223321010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.client_selector

<a id="canonical-3320313301030012-2133022223100122-2312030323133302-1331133321333313-3102020012111312-1321300000311311-0031131100323200-3033332311332133"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-1120220102300030-2321201103102030-2022223031232312-1300332002101120-3131103332233022-3132312213032121-2201003030301332-3001100310000001"></a>

## Direct properties — client_selector / 032223321010 / 3

<a id="canonical-3100120202202223-2202313020231022-0033330123012130-2122003320002323-3230212123223301-3013003230222202-0020100303101130-2121312302122200"></a>

<a id="canonical-3033010102202002-0123220000222031-2220121103022132-2302021210012012-2322123300132123-3121220332033012-0030220113030012-0231222222020221"></a>

## expressions property — client_selector / 032223321010 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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

<a id="canonical-3123303303130233-1213201302103323-2322312223020133-2211232103213031-3311111101021033-2112133223330000-0102003013123233-2021130120111123"></a>

## Next pages — client_selector / 032223321010 / 5

- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3200322322220212-1113313312121132-0022020102332110-0003210012123332-0232201320020023-1331003101031110-1133111233110332-1032321212200130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022123323311010-3312301032212112-2103310302032131-3303013302100132-0012012313231230-0031231023120313-1312110203003201-2001013233131311"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher — ip_matcher / 021220313010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher

<a id="canonical-3321023330020202-0120001131222133-0003103331110231-0310221223223232-3220020303120122-2112200130231330-3301301323201203-0210032120113320"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3130200310330322-3120102222120233-1311300033333021-1020100211332322-2103011121310123-2012201221021310-1122002212003200-1012301321332013"></a>

## Direct properties — ip_matcher / 021220313010 / 3

<a id="canonical-3200001321311111-0113320233303132-3310000200320320-0213020013333103-2302120220230103-1011333031221300-2012011223201123-0131311102323212"></a>

<a id="canonical-2102101310102123-1103130020300130-2133013001212022-3311101100331313-0300031200020001-0011322023031120-0032021332312202-1312023202130231"></a>

## invert_matcher property — ip_matcher / 021220313010 / 4

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

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

- [prefix_sets](resources--http_loadbalancer--reference--group-006.md#canonical-2023121332212302-0021301111000301-3022222231220000-2210212020101232-3233312313100222-2210001112112231-1012213133000313-0310210223320132): complete subsection reference.

<a id="canonical-0303333321310231-2232212201232023-0002212322213131-3202120001223233-3122131033003102-1011103031011232-3310323100330102-2103320312200201"></a>

## Next pages — ip_matcher / 021220313010 / 5

- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets](resources--http_loadbalancer--reference--group-006.md#canonical-2023121332212302-0021301111000301-3022222231220000-2210212020101232-3233312313100222-2210001112112231-1012213133000313-0310210223320132)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2023121332212302-0021301111000301-3022222231220000-2210212020101232-3233312313100222-2210001112112231-1012213133000313-0310210223320132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102220033101110-1202010120033230-2322212331110102-0031331122310331-0311320002002233-0110302312301331-3033132230330003-1311213030022021"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets — prefix_sets / 213010223130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-3200322322220212-1113313312121132-0022020102332110-0003210012123332-0232201320020023-1331003101031110-1133111233110332-1032321212200130)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-3223122113221002-3203110203021002-1120312102330123-2111330011313133-2130210121001302-3320301011012210-1122001031303233-3213120023321111"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

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

<a id="canonical-1302002303133031-1133220133022003-1032013210210301-1211312111100010-2133232133030321-3201221120212103-1211203123133003-1102130130033301"></a>

## Direct properties — prefix_sets / 213010223130 / 3

<a id="canonical-1200223103333112-0111330330011203-3201122121332212-0233103332212231-3002313133032121-0303310210013011-0211020121203120-3032303210331102"></a>

<a id="canonical-0110331323201031-3303112130101230-2021122030232121-0131223132323101-0102333013123323-2231131033022223-3110123110310011-0212202031201032"></a>

## kind property — prefix_sets / 213010223130 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-3231021232232133-1223310010001113-3113030232101212-3332000210003021-0330032110220311-2200321020122101-1031213002021111-0201100111022120"></a>

<a id="canonical-0100111121331132-2310013120113221-2001302021022131-0200023302333010-3120202220202332-2311031322200210-2330323133020200-3300332113022322"></a>

## name property — prefix_sets / 213010223130 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-0221132201233221-2121303130131103-3301023033122020-2210000133311302-3203332312233331-1203320132203211-0321131020103112-0322321023331223"></a>

<a id="canonical-0222001211003212-2021313103100211-0321310301023132-3332131112230233-0023233123020111-1302213132210021-2321312203013012-1123322021031310"></a>

## namespace property — prefix_sets / 213010223130 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-2113213213200302-1303120321221231-0320002220300203-3210100323132033-3201102333132302-1221031332203303-2122121021210331-0311232231312313"></a>

<a id="canonical-3320302113213032-2032023133320332-3002212312203101-3113121130202121-2302033331210031-0332220333230101-3321312023323021-2132233232131232"></a>

## tenant property — prefix_sets / 213010223130 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2133131203130330-3111013001213302-3330102323333023-3020112103333211-1130033310312022-2013233331213222-1300212323100122-3310023133320330"></a>

<a id="canonical-2220122300102321-2320003200310303-2023223331232010-3323121221321100-1013020030322011-1032303022232202-2021213303012110-0100021211213110"></a>

## uid property — prefix_sets / 213010223130 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-0112212300011310-0202212301310323-2102323010211000-1221023013313210-1313211303113032-3033323222302133-1311323102210001-1230130310302331"></a>

## Next pages — prefix_sets / 213010223130 / 9

- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-3200322322220212-1113313312121132-0022020102332110-0003210012123332-0232201320020023-1331003101031110-1133111233110332-1032321212200130)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1312202323020121-2330122112103213-2102331212330331-0311311101121312-0331221200021133-2111102011123223-0111002302103312-2030312132201000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002331231313111-2133331120222011-0122323022322201-0020023132020122-2121301103000200-1220031011033211-0113332023031223-2201023233000000"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list — ip_prefix_list / 032130221202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list

<a id="canonical-1330322301103320-3002021223203303-3011033110031031-1213200322231031-2120210010003003-3323302313112232-1033011113032013-0220223330320030"></a>

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

<a id="canonical-1110310003023010-0230122233223010-0113230100303102-2131231113123030-0300122220203303-0133303102000123-1301100212021300-3232031122323322"></a>

## Direct properties — ip_prefix_list / 032130221202 / 3

<a id="canonical-0103320030100023-1032133201201222-3002302030331233-1021301201230301-3233231302102112-0100111331001311-2323111220023121-0013212211011222"></a>

<a id="canonical-3020202332231231-0211003311102101-0110200033230132-1330231102320220-3003023320111121-3230212203302213-3310300000120321-2013120112202201"></a>

## invert_match property — ip_prefix_list / 032130221202 / 4

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

<a id="canonical-2132131203111130-2122112133232030-2331323221012102-0210330110100132-2213032020210010-1121110033012210-0111233321121222-2301301223123023"></a>

<a id="canonical-0113023222321011-1311321031001233-2113121103311223-3032232211321201-2231222110020011-2130113002233033-2101122002033202-0111303330320011"></a>

## ip_prefixes property — ip_prefix_list / 032130221202 / 5

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

<a id="canonical-1300001212320032-0203212331322211-2323333112332000-3130021021321230-1103202123310310-2232123102332310-1231230022201120-0100330130001320"></a>

## Next pages — ip_prefix_list / 032130221202 / 6

- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1231130011002201-0201220210132022-1202011011322320-3222221230011211-1220312323300101-1113033232232300-1300130111323030-0113012032033021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330110002323031-0022221022321223-3131033012221030-2331110203031100-2000023220131002-0021202301231022-3133213121020203-2300031120013331"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list — ip_threat_category_list / 231121301312 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list

<a id="canonical-0001121323103103-3300130121010021-0201000132233033-2000002010301303-1003323220132001-2121303301200131-0132000231200131-1310222100032211"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_threat_categories")}
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
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2012221133011312-1022030311231311-0020323111323302-2321000212202103-2100301131133222-1111031332010022-1221113303102313-2203022321023022"></a>

## Direct properties — ip_threat_category_list / 231121301312 / 3

<a id="canonical-3112103230121320-1030120301101101-3000030123011332-0203323222330031-3203333031003310-3103102131221202-1302221000310022-3313330003000300"></a>

<a id="canonical-1213010103013302-2103123111212212-2011233210033110-2013110032333023-3301020132101321-1113123311310301-2020312111300312-0302113111200230"></a>

## ip_threat_categories property — ip_threat_category_list / 231121301312 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Upstream description:

The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3301230023010312-3210313013213331-1321122301000213-3332110202322032-1212031212021113-2311111220023230-3130332111012210-3031333322111322"></a>

## Next pages — ip_threat_category_list / 231121301312 / 5

- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3201202030331112-3102102112311122-2223110300113033-2231001222303211-3001320300131233-3202223231022132-1213133331012100-3122020302033123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211302001310313-0202223213210003-1310111020022232-3323212133313313-2111101222213000-1132130220003123-1100330212311202-0221200202102113"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher — tls_fingerprint_matcher / 031022112300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-1020030210001303-0033221023301031-0310321013333311-3213230212231002-0311330133131000-3212331221002112-0311203332333312-2302321033001322"></a>

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

<a id="canonical-0121231020112002-2033323113110133-3000200131201023-3130213102232311-1322312212122010-0023101213223313-2201010222202322-2021013320233313"></a>

## Direct properties — tls_fingerprint_matcher / 031022112300 / 3

<a id="canonical-0132321300203110-1203002233331023-0312211320210030-1311302233320132-0203013113230303-0213210003133131-2013022120013312-2130003200103030"></a>

<a id="canonical-1212110122101001-0300002330311301-0132011120211320-1122110111233000-3222212122200320-1220213032223113-0012212233012121-3312321300311121"></a>

## classes property — tls_fingerprint_matcher / 031022112300 / 4

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

<a id="canonical-2110312103111032-0111300133313231-1122223111131000-2331332102112101-1102321110003123-3022323203213132-2231233120303023-2100313131112032"></a>

<a id="canonical-2211222003131213-1102031100033310-0130021022302123-3021202332130101-2030121032221023-3113302200333200-1213320113331031-1313332001010111"></a>

## exact_values property — tls_fingerprint_matcher / 031022112300 / 5

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

<a id="canonical-1230110222023032-0202123322201032-1231333212321331-1030033123201112-2220320312011211-3102312132230211-2120002321310013-0111121033000203"></a>

<a id="canonical-0332313322023011-2113231222011012-2111211103203302-0001113313202012-0322002203122313-1210310111301030-3023013013122332-2013313333313113"></a>

## excluded_values property — tls_fingerprint_matcher / 031022112300 / 6

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

<a id="canonical-2202320330310330-0322322301220213-1232332112123121-3110001031331133-1232131033230122-3132121223220110-1231320223212020-3021112323020231"></a>

## Next pages — tls_fingerprint_matcher / 031022112300 / 7

- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1121233023131220-1003111220221011-0133310311113000-1021131213123031-3323030010330300-0200200213101021-1233030130032333-0013322100312031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323223100331201-2030110331220232-3100302003012123-1213110000013232-2300230131222131-3120012132201221-0110101310201221-3101233023211300"></a>

## api_rate_limit.api_endpoint_rules.inline_rate_limiter — inline_rate_limiter / 233032312311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter

<a id="canonical-3032133021102233-1331030100031322-3030202220021113-1223322100201011-3110301231002132-0101010211022101-2100000313002220-0202120031122033"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inline rate limiter.

Upstream description:

Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the
required rate\_limiter\_choice when no stored rate-limiter object is used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("threshold"),
  validators.ConflictingObjectAttributes("ref_user_id",
    "use_http_lb_user_id")}
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
  "x-ves-oneof-field-count_by_choice": "[\"ref_user_id\",\"use_http_lb_user_id\"]"
}
```

Terraform syntax:

```terraform
inline_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-3332012121233030-1201333121111211-3313002113033302-2330003323232123-0023213330212102-1110123203322201-2100221132211112-3231023230110220"></a>

## Direct properties — inline_rate_limiter / 233032312311 / 3

- [ref_user_id](resources--http_loadbalancer--reference--group-006.md#canonical-1322000210130123-0113311310103133-3112203131210121-0320302333110311-0113031113200002-3223121312302200-1113233233203222-3000010213202003): complete subsection reference.

<a id="canonical-1311232223200320-3301220321321211-1112320123123300-0113023000002303-1012120320103000-1232021121221211-0002111333301032-3231202000323111"></a>

<a id="canonical-2230211213011333-2203130221132121-2212222221321110-1213312130210133-1031133220013233-2023002210230022-3031233323330323-3313103100031221"></a>

## threshold property — inline_rate_limiter / 233032312311 / 4

Type: `"number"`. Optional.

The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified
period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 8192),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
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

<a id="canonical-2130211202000102-1122131022031000-1302113200201232-0220323001230131-0330332013020312-3032211201003303-1002000131331113-0110000231130312"></a>

<a id="canonical-1232212123223131-2110011231101112-1322302230212200-3201322333121030-3330311133323013-2013223203330020-3310333220103023-1232132231002230"></a>

## unit property — inline_rate_limiter / 233032312311 / 5

Type: `"string"`. Optional.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Upstream description:

Unit for the period per which the rate limit is applied.

&#8203;- SECOND: Second

Rate limit period unit is seconds &#8203;- MINUTE: Minute

Rate limit period unit is minutes &#8203;- HOUR: Hour

Rate limit period unit is hours &#8203;- DAY: Day

Rate limit period unit is days.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SECOND",
    "MINUTE",
    "HOUR"),
}
```

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

- [use_http_lb_user_id](resources--http_loadbalancer--reference--group-006.md#canonical-3330031100020111-3030203233132221-1000321133200203-1321313131010003-3122213211133010-3300211303020132-1200031300300131-2210032020311321): complete subsection reference.

<a id="canonical-1121320231312100-2222300203123322-1131030210112100-2013320323211211-1001312031013212-3112003003111220-2101323003322331-0112013102201203"></a>

## Next pages — inline_rate_limiter / 233032312311 / 6

- [api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id](resources--http_loadbalancer--reference--group-006.md#canonical-1322000210130123-0113311310103133-3112203131210121-0320302333110311-0113031113200002-3223121312302200-1113233233203222-3000010213202003)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id](resources--http_loadbalancer--reference--group-006.md#canonical-3330031100020111-3030203233132221-1000321133200203-1321313131010003-3122213211133010-3300211303020132-1200031300300131-2210032020311321)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1322000210130123-0113311310103133-3112203131210121-0320302333110311-0113031113200002-3223121312302200-1113233233203222-3000010213202003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301110103323102-1232122011123301-0220223023223323-0301021000111232-0103333313111023-2223302300232011-2121310122311101-0002202100223123"></a>

## api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id — ref_user_id / 330113202030 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-006.md#canonical-1121233023131220-1003111220221011-0133310311113000-1021131213123031-3323030010330300-0200200213101021-1233030130032333-0013322100312031)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id

<a id="canonical-0232331112003121-0302311101311331-2331313120203301-1212132313310033-1202003333201312-2300103323213001-0033223300222010-0302110032012222"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
ref_user_id {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222302331203102-0333302210231033-2000133012013122-3203133033212131-3032222230210320-1322230231211222-2310300322232323-1330221032103320"></a>

## Direct properties — ref_user_id / 330113202030 / 3

<a id="canonical-1313010013102033-2130020321321311-1210233100113032-3103323032023110-2333010023121012-3230133132223220-0001201101231200-0110323322302131"></a>

<a id="canonical-0010230132013112-1012111231310300-3303112233302303-3112233223000201-2130022233332231-2000011103331212-3211020000310231-1113110133111012"></a>

## name property — ref_user_id / 330113202030 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-0120213103130302-1201010301322301-1221020302323121-2211000313023332-2021213221122122-2231312103230020-1133021322211001-1311311033100013"></a>

<a id="canonical-2000001101222320-1001320123103310-2100100212022031-1021222111003213-1222200121200300-1320200012011231-0302130313100230-3321001203002300"></a>

## namespace property — ref_user_id / 330113202030 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-1102030003221110-3032012132102022-3202301330132322-0210113301221213-2031110021122230-3221132132321111-2101000132101320-2202301212302032"></a>

<a id="canonical-0330033003023311-1333311310302333-0301122003113320-2311123310010331-2310100233210123-2110333312210030-0011011323222023-0221333111102210"></a>

## tenant property — ref_user_id / 330113202030 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-3020020222233210-1110023221120003-0031102013321321-2030121001211333-1302303000011013-3221100221231111-1200300220331231-2021213101310033"></a>

## Next pages — ref_user_id / 330113202030 / 7

- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-006.md#canonical-1121233023131220-1003111220221011-0133310311113000-1021131213123031-3323030010330300-0200200213101021-1233030130032333-0013322100312031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3330031100020111-3030203233132221-1000321133200203-1321313131010003-3122213211133010-3300211303020132-1200031300300131-2210032020311321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333202211000130-2102311021333033-2111003133223031-1100131303101323-3330321212131122-3313211311112223-1110233110211022-1120312312310220"></a>

## api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id — use_http_lb_user_id / 130333231230 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-006.md#canonical-1121233023131220-1003111220221011-0133310311113000-1021131213123031-3323030010330300-0200200213101021-1233030130032333-0013322100312031)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-2203110110331322-3032012202233323-1113021333010211-1232210210220321-1332202303223133-0001323102023022-0221103020311310-3331000102013303"></a>

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
use_http_lb_user_id = {}
```

<a id="canonical-2103223101032300-0201133310212131-2310202232103330-1130020220203001-1231232213230320-2003110222200021-0301001313222120-2300320213110331"></a>

## Direct properties — use_http_lb_user_id / 130333231230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120120303112032-1333112201201232-3321213010300210-1023122213202123-0221113030300310-1202020001112203-2000220121110333-2113033012312130"></a>

## Next pages — use_http_lb_user_id / 130333231230 / 4

- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-006.md#canonical-1121233023131220-1003111220221011-0133310311113000-1021131213123031-3323030010330300-0200200213101021-1233030130032333-0013322100312031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3031121110211321-0222223200032133-1223103211312023-1012222201100010-0323202231130000-2323201330321110-1222232311021302-3202202311112033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021303010000330-3201110102231212-0231311322320232-2020300200311102-1132232230311013-0321223012302123-1302203001200010-3123030200102300"></a>

## api_rate_limit.api_endpoint_rules.ref_rate_limiter — ref_rate_limiter / 010211123222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- api_rate_limit.api_endpoint_rules.ref_rate_limiter

<a id="canonical-3112213233313300-2121330033100010-0303132133030301-3320002023002222-2232220020303001-3300000201222122-2210013213011210-3033321100331201"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

Reference to a stored rate-limiter object for this scoped rule. Select exactly one of
ref\_rate\_limiter and inline\_rate\_limiter.

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
ref_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131021130030200-1222223001000203-3003132131002010-2022022112111010-0220233022013311-1002122210013303-1210130100333113-2222112111021001"></a>

## Direct properties — ref_rate_limiter / 010211123222 / 3

<a id="canonical-2100123312323022-2130220123200132-3201213220130011-3122103112103300-1032030132220232-3212003001102133-1203232212032221-0310033112223322"></a>

<a id="canonical-1223100302310010-0122030302102321-2232321032220311-2220211100323301-2003131321003230-3011321331020230-0002210230132012-3303001320323312"></a>

## name property — ref_rate_limiter / 010211123222 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-1302033100220222-2131010303211321-1211122221202311-1232112033101330-2012322002323201-3010130332322321-3301232332132111-2122333101223031"></a>

<a id="canonical-1001221320133113-3030023301231220-1031331111021113-0310301223133211-1101022333110020-2212031010023210-0113232032011200-0310132323112123"></a>

## namespace property — ref_rate_limiter / 010211123222 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-1222332210122300-3120310010333223-2032021230211232-3231123131001020-0111331013132320-1213103020100023-0213111231301032-1111312213330321"></a>

<a id="canonical-0010103111130123-3021233111231313-1121201322231120-2212031233020200-1023311302311223-0222232323012330-0311120210302100-1220121332333213"></a>

## tenant property — ref_rate_limiter / 010211123222 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-0210122210113320-2010202330201012-1330213011202101-1222132023212121-1231002120202223-1312002211203032-0330033010302313-3332210032021103"></a>

## Next pages — ref_rate_limiter / 010211123222 / 7

- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012101011333320-2312001320032131-1123020300220002-1112121103001102-3020013131033233-3012130310233311-1320213033110110-0021113230011231"></a>

## api_rate_limit.api_endpoint_rules.request_matcher — request_matcher / 112012131222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- api_rate_limit.api_endpoint_rules.request_matcher

<a id="canonical-1320020013211132-2031013002002013-1203201303313130-1213200320011033-2103201322111302-0133123010012232-3213222101200313-3302231313332121"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for request matcher.

Upstream description:

Request conditions for matching a rule.

Receipt-pinned upstream constraints:

```json
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
request_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303023333303012-1320033211023011-1213033233231030-3002013113100200-2330012221020233-0123323312022233-3233233231000303-1132113302322021"></a>

## Direct properties — request_matcher / 112012131222 / 3

- [cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-007.md#canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221): complete subsection reference.

- [jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-007.md#canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321): complete subsection reference.

<a id="canonical-3330222033301202-0320111233310212-0232021020221102-2230131220031202-3120220303202032-0123002103110323-0312223203033132-2331231011301033"></a>

## Next pages — request_matcher / 112012131222 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-007.md#canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020312320222103-0332033302031213-2313130333022222-3202213001323220-2332022203221102-0103113000321203-1112203311022013-1201001210131123"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers — cookie_matchers / 311033320113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers

<a id="canonical-1312232002030230-3031023203302003-2233331213111113-3230312233323132-0322031211203333-1313120203133101-1023232211001021-1132122212221201"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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

<a id="canonical-3110132300130231-1001323012021233-1322002021013221-2222010013012322-0222103321332112-1030322002210110-1200022230012302-2210330102321322"></a>

## Direct properties — cookie_matchers / 311033320113 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-2131023022232100-1313031230102113-1322331330011233-2003100010212013-1221031002330332-2130233033302100-1301100011232323-1232303001203211): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-007.md#canonical-2122101123213330-3233231220232002-1332013023321002-0223022233200113-3332002102113313-1330000032111110-3200220223030122-2200323201130011): complete subsection reference.

<a id="canonical-2013322200103021-2332012002310130-3302020130113310-1310032220221231-2133100321132131-0322310111003223-2333013332212003-2113103000213000"></a>

<a id="canonical-3001101202131231-0110023320103220-2220020210200023-1330031100200021-2111303332203312-0113312133301301-3221312320310023-3131130210013001"></a>

## invert_matcher property — cookie_matchers / 311033320113 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--http_loadbalancer--reference--group-007.md#canonical-1012000023300311-2100310300202101-0330200133301120-2132232320103013-2222300100123211-2232201001100122-1233230020202012-1012030213200011): complete subsection reference.

<a id="canonical-0132130122112023-0001303101133310-3313223232121300-2023203133000321-1002030020230220-1130121023321113-1331102323332203-2023323231312321"></a>

<a id="canonical-2012110031022323-1000013013122301-2310313132021330-2210102301222322-2131123123033320-3302322232221202-0302312023003223-0012211203230121"></a>

## name property — cookie_matchers / 311033320113 / 5

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```
