---
page_title: "xcsh_gcp_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site reference."
---

# xcsh_gcp_vpc_site reference

<a id="canonical-1200311111133100-0020033221223212-3021002202113332-2301121021212132-0232212133303310-2231010113102221-0033221003131012-2022210321113201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021231213230110-0330001311023301-3321122321023012-2322010132032220-3200320133020122-2331020323203130-1131032331312323-0331312113202031"></a>

## private_connect_disabled — private_connect_disabled / 333321023122 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- private_connect_disabled

<a id="canonical-2233301233101133-1320123210011220-2031030310331002-1213231021002033-3120333011113032-0311131033130112-3131132121200013-1113011101123023"></a>

Type: `["object", {}]`. Optional.

\[OneOf: private\_connect\_disabled, private\_connectivity\] Enable this option

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

OneOf alternatives in this subsection:

- [private_connect_disabled](resources--gcp_vpc_site--reference--group-004.md#canonical-2233301233101133-1320123210011220-2031030310331002-1213231021002033-3120333011113032-0311131033130112-3131132121200013-1113011101123023)
- [private_connectivity](resources--gcp_vpc_site--reference--group-004.md#canonical-0012302121203202-1121213313020210-2231002021013133-3303022320311132-2003013103011130-2102110312212203-2330210031022200-1121300300012103)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
private_connect_disabled = {}
```

<a id="canonical-0222123331010330-1211003201203033-3312131320302212-2031200132011001-1133012023121103-3130033121011203-2212131020320113-1223211012232120"></a>

## Direct properties — private_connect_disabled / 333321023122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330231333002321-2001123011112301-1312300112021300-1330302301231312-2033110100030333-0223022003102033-1233201103133001-2200320223202222"></a>

## Next pages — private_connect_disabled / 333321023122 / 4

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0223011112231012-1302132230002202-0010012111321332-1313123023032111-0101211323012012-2102120320023221-2221022121130120-3120021033132133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103320132102103-2012330322032032-2120021313202033-1020022011121111-3113110213230020-3000201033133003-0321202202232100-3002132130220221"></a>

## private_connectivity — private_connectivity / 230323302111 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- private_connectivity

<a id="canonical-0012302121203202-1121213313020210-2231002021013133-3303022320311132-2003013103011130-2102110312212203-2330210031022200-1121300300012103"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for private connectivity.

Upstream description:

Private Connect Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside",
    "outside")}
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
  "x-ves-oneof-field-network_options": "[\"inside\",\"outside\"]"
}
```

Terraform syntax:

```terraform
private_connectivity {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213201322211212-3021020303303003-3213130212021300-2032003322221122-0213321320103013-2030002121132111-3122300222130013-3013102312111033"></a>

## Direct properties — private_connectivity / 230323302111 / 3

- [cloud_link](resources--gcp_vpc_site--reference--group-004.md#canonical-0133221210221313-2312310223201210-1100302320022031-0113013131333221-3032312113213332-1211120112012202-3013321212212323-2311332001020033): complete subsection reference.

- [inside](resources--gcp_vpc_site--reference--group-004.md#canonical-1001002332103232-3330302311332223-3200210312000331-3301000202221331-0303202032301110-1321320212233301-2322001022210310-3311020201022301): complete subsection reference.

- [outside](resources--gcp_vpc_site--reference--group-004.md#canonical-0223210311120100-0330311122030023-3201023300202003-1003300003231001-2022232212231111-3031123302022332-2233031300012012-0320122000022223): complete subsection reference.

<a id="canonical-0110023330301123-3120223222313000-0331001011121013-0223232103211031-1220113110322333-1233230320130312-0030313212133112-3331321231323311"></a>

## Next pages — private_connectivity / 230323302111 / 4

- [private_connectivity.cloud_link](resources--gcp_vpc_site--reference--group-004.md#canonical-0133221210221313-2312310223201210-1100302320022031-0113013131333221-3032312113213332-1211120112012202-3013321212212323-2311332001020033)
- [private_connectivity.inside](resources--gcp_vpc_site--reference--group-004.md#canonical-1001002332103232-3330302311332223-3200210312000331-3301000202221331-0303202032301110-1321320212233301-2322001022210310-3311020201022301)
- [private_connectivity.outside](resources--gcp_vpc_site--reference--group-004.md#canonical-0223210311120100-0330311122030023-3201023300202003-1003300003231001-2022232212231111-3031123302022332-2233031300012012-0320122000022223)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0133221210221313-2312310223201210-1100302320022031-0113013131333221-3032312113213332-1211120112012202-3013321212212323-2311332001020033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000231003321020-1122201310030231-0002202301101022-1300030132222123-3301010212110122-3232310332331002-0123102213122323-3112112000320101"></a>

## private_connectivity.cloud_link — cloud_link / 301120211212 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [private_connectivity](resources--gcp_vpc_site--reference--group-004.md#canonical-0223011112231012-1302132230002202-0010012111321332-1313123023032111-0101211323012012-2102120320023221-2221022121130120-3120021033132133)
- private_connectivity.cloud_link

<a id="canonical-0210000023220202-0033032230123002-2020021031131221-2032101233131303-2013103113010221-3231100313213003-2210113002220211-1022022303333032"></a>

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
cloud_link {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201322213012011-1233022200232012-1301123102221010-0320221011221222-2020200333010111-0333320020310103-2322102121033223-3323100110023101"></a>

## Direct properties — cloud_link / 301120211212 / 3

<a id="canonical-3323232130122201-0311200223131222-3110030001033331-3032023131020330-0201003123322113-1323230020230332-1102232111232212-2302323011232202"></a>

<a id="canonical-0312321120333221-0213000300311003-2201122230001202-0221311100100103-0200030302023200-0200331210111232-1120310221302200-1011210101111022"></a>

## name property — cloud_link / 301120211212 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1022031331201123-3200311022312321-0120200232332201-0231201331331003-2233023120202310-1012320012300200-3322111311113013-0321310233023230"></a>

<a id="canonical-2312121023120031-1123203103131233-3023001011220030-0322103103300112-2223221122330301-3113030130131032-2011112020110010-1023302321301230"></a>

## namespace property — cloud_link / 301120211212 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0210321321210230-0200000322131012-0221132101112320-1110033313200231-0111322000330021-1123002212223231-1033020101120100-1321030011230133"></a>

<a id="canonical-0130003300302100-2331103201031230-2131200320032322-1112302211330330-2203301233330220-3011133002023220-0233302301220200-2022320120300320"></a>

## tenant property — cloud_link / 301120211212 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0111121003202323-3130232030000201-1113332310223323-2110030232301213-3303310132300030-1000031332130032-0200023110321211-0310313011102213"></a>

## Next pages — cloud_link / 301120211212 / 7

- [private_connectivity](resources--gcp_vpc_site--reference--group-004.md#canonical-0223011112231012-1302132230002202-0010012111321332-1313123023032111-0101211323012012-2102120320023221-2221022121130120-3120021033132133)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1001002332103232-3330302311332223-3200210312000331-3301000202221331-0303202032301110-1321320212233301-2322001022210310-3311020201022301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210322120222123-2202001311301013-3322211033030211-3030302001111101-2031110121101001-0010202323011013-2200101133303233-3213133021332123"></a>

## private_connectivity.inside — inside / 333113112321 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [private_connectivity](resources--gcp_vpc_site--reference--group-004.md#canonical-0223011112231012-1302132230002202-0010012111321332-1313123023032111-0101211323012012-2102120320023221-2221022121130120-3120021033132133)
- private_connectivity.inside

<a id="canonical-3023020310313033-1133101321220232-0002332000020001-2322313121100320-3130212313130331-3021003130003310-1111000022132222-2123310211310022"></a>

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
inside = {}
```

<a id="canonical-2133123321303103-1110103131200222-2020130213112121-0021330310113310-3212202332331113-0130121201223323-0031132230123002-2311301103010303"></a>

## Direct properties — inside / 333113112321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023331303200133-3223311131100022-2002302112301003-2333112103321303-0030321213200321-1203323021010031-2203323102030203-3110310301132203"></a>

## Next pages — inside / 333113112321 / 4

- [private_connectivity](resources--gcp_vpc_site--reference--group-004.md#canonical-0223011112231012-1302132230002202-0010012111321332-1313123023032111-0101211323012012-2102120320023221-2221022121130120-3120021033132133)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0223210311120100-0330311122030023-3201023300202003-1003300003231001-2022232212231111-3031123302022332-2233031300012012-0320122000022223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131123202032202-2002001210202212-2311220201133220-1022022032222012-2233330202033122-3012230303200223-2001003000201002-2001020312023133"></a>

## private_connectivity.outside — outside / 322311033222 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [private_connectivity](resources--gcp_vpc_site--reference--group-004.md#canonical-0223011112231012-1302132230002202-0010012111321332-1313123023032111-0101211323012012-2102120320023221-2221022121130120-3120021033132133)
- private_connectivity.outside

<a id="canonical-2301130220031333-1032302232212020-2000130322333312-2000000210223022-1132220030211001-1030320301212230-2101123221122010-0212102311312302"></a>

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
outside = {}
```

<a id="canonical-2230100123130033-2200223202033112-0333211333100001-3102331320222101-1333001232123002-1312101212302013-2211223313003210-0100031000311102"></a>

## Direct properties — outside / 322311033222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232003212101110-0232213310233022-0231320332231222-2301331032012333-0130010023022333-2222303212013023-1111130123000102-2211203333321130"></a>

## Next pages — outside / 322311033222 / 4

- [private_connectivity](resources--gcp_vpc_site--reference--group-004.md#canonical-0223011112231012-1302132230002202-0010012111321332-1313123023032111-0101211323012012-2102120320023221-2221022121130120-3120021033132133)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0212133321121330-1301131103023331-2123132213311103-0031232100002231-1233033231110220-1102201033230030-3020012013122301-2110012031323002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320023130022010-0112100321000033-0301030233133231-3221203101013031-2013112330231030-2130202221211332-2121000310010232-3320331230031003"></a>

## sw — sw / 100303020123 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- sw

<a id="canonical-2230022121132012-3223032230323023-1320010311112313-3023032010031002-1023002302330022-3231111312200112-2200100031111003-3210133123231221"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Upstream description:

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_sw_version",
    "volterra_software_version")}
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
  "x-ves-oneof-field-volterra_sw_version_choice": "[\"default_sw_version\",\"volterra_software_version\"]"
}
```

Terraform syntax:

```terraform
sw {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302001103221112-3231222323131131-0131301131100313-2031032211122321-0033131221110220-0000122233011310-0323333330000132-0302333102231002"></a>

## Direct properties — sw / 100303020123 / 3

- [default_sw_version](resources--gcp_vpc_site--reference--group-004.md#canonical-2010231121221120-2013203331121100-3213101323300211-3210110213023031-3112112132002312-3111032321310132-1031121133231011-0101203323331102): complete subsection reference.

<a id="canonical-3123111223133002-3012011121232301-1333301202213200-0302031301221220-1200030110222023-3202130212023012-1133221022000132-3231002223021122"></a>

<a id="canonical-2220321022311312-1231011330012021-1320111023123200-0231133201220102-2303223133300021-1110102310033322-2312113222313311-1311233032021032"></a>

## volterra_software_version property — sw / 100303020123 / 4

Type: `"string"`. Optional.

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Upstream description:

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-3332230011112232-1331312100112023-2303223330130123-3300313122221211-1230320002202232-1230112002123321-2112232310000100-0201212113321221"></a>

## Next pages — sw / 100303020123 / 5

- [sw.default_sw_version](resources--gcp_vpc_site--reference--group-004.md#canonical-2010231121221120-2013203331121100-3213101323300211-3210110213023031-3112112132002312-3111032321310132-1031121133231011-0101203323331102)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2010231121221120-2013203331121100-3213101323300211-3210110213023031-3112112132002312-3111032321310132-1031121133231011-0101203323331102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303322222100130-1210332033330020-3210132323210010-1202131332331130-2010101313200213-3330130120010310-2203211013132110-3312323232132130"></a>

## sw.default_sw_version — default_sw_version / 333311000300 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [sw](resources--gcp_vpc_site--reference--group-004.md#canonical-0212133321121330-1301131103023331-2123132213311103-0031232100002231-1233033231110220-1102201033230030-3020012013122301-2110012031323002)
- sw.default_sw_version

<a id="canonical-1022302300321021-1312023001231313-3313201023130130-0201122333001310-1332113111233313-0130213323102132-2000023301102210-2313212211223202"></a>

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
default_sw_version = {}
```

<a id="canonical-3311223131310123-0320203201111331-2331121211201212-1330033332303130-3123103131231021-3202020330332210-2232101301213210-3223110300033131"></a>

## Direct properties — default_sw_version / 333311000300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010211012013321-3132333123322122-3230012230013300-0000101323012331-0130232201100133-1130331133130002-1133233202201212-0131202212300331"></a>

## Next pages — default_sw_version / 333311000300 / 4

- [sw](resources--gcp_vpc_site--reference--group-004.md#canonical-0212133321121330-1301131103023331-2123132213311103-0031232100002231-1233033231110220-1102201033230030-3020012013122301-2110012031323002)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0202212023111322-0302222202030003-3121101231202200-2213031121020101-0121233300121223-1032213023303312-0220022023213022-0231230022223121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313123003100222-3303130030303001-2110303031320212-2303311002101331-3012011102130012-0112110312313230-0333003320010212-1131111112302010"></a>

## timeouts — timeouts / 230231231220 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- timeouts

<a id="canonical-2032223202203233-0230223113122212-0211131303203223-1211220330323100-0233000102221303-1010131310223011-3112011211100110-3210300013032132"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2323220112131022-2223100322010011-1130120313033333-2033021032031221-1333022111201301-3213131321032321-0033021113023323-0003111211102111"></a>

## Direct properties — timeouts / 230231231220 / 3

<a id="canonical-0320201010322333-2231210101121230-0022130122331330-2230131332003111-0120212302030021-0331333022020321-1312201000131233-3323212131323033"></a>

<a id="canonical-3121032210333102-2232333223023302-3221121013133113-2311003003130220-2030230031111133-2123023232031222-3212003203010111-1200211232003213"></a>

## create property — timeouts / 230231231220 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3020113212011120-3001313311333111-0321223123123312-0233013212302221-0213211002001013-2012100230021120-0023122022233310-2332231203200023"></a>

<a id="canonical-1100133333302023-1031210230230213-3012332321000330-3310200003031001-3231131301033322-3033011032103302-1230121120123132-3330012211010113"></a>

## delete property — timeouts / 230231231220 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0312331023101330-2133031121203110-3330010302133111-1321112020013322-1132323002201213-2112123211213311-3000020300102123-0201332221303310"></a>

<a id="canonical-1110233302121120-3232121133300111-1312010312332123-2322123030011333-3120222031311000-1312120131011031-0232132223032312-2121221002122013"></a>

## read property — timeouts / 230231231220 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0133213020320121-2110123122032222-1230311110330232-2332022231233101-2023332111322113-2210332012130112-0002113311103210-2332113302001212"></a>

<a id="canonical-3231130023100000-1320121203100002-1103300302323012-2001002220021303-0122301021331232-1021313333010133-0011303221032303-3031200030213120"></a>

## update property — timeouts / 230231231220 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2212101213010031-2232220030223023-1032022131202331-3202113232220300-0202321003031122-3111031031121133-3023213211200210-0000330211303102"></a>

## Next pages — timeouts / 230231231220 / 8

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002033303311213-3100221222120331-2233113201332133-0300232211210131-1003302333211020-0121112200300330-0032111300202301-0213103303021132"></a>

## voltstack_cluster — voltstack_cluster / 213213010113 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- voltstack_cluster

<a id="canonical-1312003312232212-2011113310110131-0010131322113332-3131330000201333-3320312120210223-3032202123102121-1200233120130210-3103303212322322"></a>

Type: `"object"`. single nested block, Optional.

App Stack cluster of single interface GCP site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("gcp_certified_hw",
    "gcp_zone_names"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "active_network_policies"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "forward_proxy_allow_all"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("active_network_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("dc_cluster_group",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("default_storage",
    "storage_class_list"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("k8s_cluster",
    "no_k8s_cluster"),
  validators.ConflictingObjectAttributes("no_outside_static_routes",
    "outside_static_routes"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-k8s_cluster_choice": "[\"k8s_cluster\",\"no_k8s_cluster\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]",
  "x-ves-oneof-field-storage_class_choice": "[\"default_storage\",\"storage_class_list\"]"
}
```

Terraform syntax:

```terraform
voltstack_cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-3010113313021020-3003120323203312-0211311100210333-0202123123032223-0032123011020202-1111113302321321-1310200213011123-0101201320021111"></a>

## Direct properties — voltstack_cluster / 213213010113 / 3

- [active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-0013032210231210-3130232233133123-2030120322033300-1111200101331012-3111132210102213-2000000003310111-3113230013210211-0332200121031300): complete subsection reference.

- [active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-3312111010301003-0130023032013332-2022321323101201-1121230133310020-1200001301312323-1320322210222110-0311310110321013-2212131111322023): complete subsection reference.

- [active_network_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-0121131301010003-2322312003102100-0230100310312000-1223202211131010-2031101011102012-1231322201013023-2101020023031123-0301132031301310): complete subsection reference.

- [dc_cluster_group](resources--gcp_vpc_site--reference--group-004.md#canonical-1130111002232032-1120202321002331-2003200330312011-0003133202133132-1221210333322301-1002030011203132-0132310110222110-3013101300023101): complete subsection reference.

- [default_storage](resources--gcp_vpc_site--reference--group-004.md#canonical-1032233013101222-0120133202121031-3222110133200101-0201212113311212-0022312013322011-3211100333221110-1103010130210302-3331221303102023): complete subsection reference.

- [forward_proxy_allow_all](resources--gcp_vpc_site--reference--group-004.md#canonical-2311102031021112-0302330203020300-1010230201123312-2131223030033213-2012231222131310-2311310031112221-2200312223322013-2223211321120130): complete subsection reference.

<a id="canonical-1230032001131031-1211111322120131-1103200133101021-3000233300301303-3122330030102302-2100202320120230-3103230020133000-1220300331220200"></a>

<a id="canonical-2002210120203323-3003012202232001-2232330231122322-3001232120330020-3321232300122232-2311132312030002-3022022330010301-3111300200122222"></a>

## gcp_certified_hw property — voltstack_cluster / 213213010113 / 4

Type: `"string"`. Optional.

\[Enum: gcp-byol-voltstack-combo\] GCP Certified Hardware. Name for GCP certified hardware. The only
possible value is \`gcp-byol-voltstack-combo\`.

Upstream description:

Name for GCP certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("gcp-byol-voltstack-combo"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "gcp-byol-voltstack-combo"
  ],
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-2032023210311213-2321323333200032-3223210002113301-2310212320232133-0023000000313203-3111200321332111-0032003220231122-1122302020321001"></a>

<a id="canonical-1322030101230210-1123111032220133-0123033102331110-1213030230333331-1330331131013302-3323020012000213-2233110010232112-3123302232210300"></a>

## gcp_zone_names property — voltstack_cluster / 213213010113 / 5

Type: `["list", "string"]`. Optional.

X-required List of zones when instances will be created, needs to match with region selected.

Upstream description:

X-required List of zones when instances will be created, needs to match with region selected.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(3),
}
```

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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [global_network_list](resources--gcp_vpc_site--reference--group-004.md#canonical-1333132321010123-1313031331321231-1212331120001123-0100212001231131-2213013320113322-3210103133122120-3313133330033113-2130303130210020): complete subsection reference.

- [k8s_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-1222231131333230-0113000121132120-1000132321133333-3212021230133132-2222002111312120-0020012232033023-2322111300303223-0012200200102232): complete subsection reference.

- [no_dc_cluster_group](resources--gcp_vpc_site--reference--group-004.md#canonical-1001131233200020-3032331103202032-3001200301021233-0101231212103123-3300231033202332-2112201210110031-3322113300220030-1133213311301222): complete subsection reference.

- [no_forward_proxy](resources--gcp_vpc_site--reference--group-004.md#canonical-3100322021212032-1020223333302222-1122300021021102-2110033203300300-2210010312230213-0233330330112311-3101332030332101-3201000100303332): complete subsection reference.

- [no_global_network](resources--gcp_vpc_site--reference--group-004.md#canonical-2311111001231211-2000113311333023-1332000013323020-2211113333001021-0101313330022200-0223222133201123-0122220300120131-0331111322221120): complete subsection reference.

- [no_k8s_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-3220011312322231-3322211113300132-2102320122223311-2121120221032120-2132333033002120-3132203030312211-2021120130123012-2200121020020220): complete subsection reference.

- [no_network_policy](resources--gcp_vpc_site--reference--group-004.md#canonical-0121032232121200-3202013002313021-0001101332201023-3023030321122233-1232023113311330-0201210210012230-1332013213021110-2231111123103313): complete subsection reference.

- [no_outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-3221122321322132-0123030132333130-1100030013030230-1211310212210132-0322113011330122-0331022011121130-2021301132200011-2203120313103121): complete subsection reference.

<a id="canonical-0030103222301332-3301301102132020-3303110220010031-0313133030233010-1201321330021002-2300202110332000-3110012012203000-2000100232030101"></a>

<a id="canonical-0032001220031001-2000131312212113-2100323112202231-1123113123113023-1302221032000200-1211230030311110-3323130010313323-0001001021120213"></a>

## node_number property — voltstack_cluster / 213213010113 / 6

Type: `"number"`. Optional.

Number of main nodes to create, either 1 or 3.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.in": "[1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.in": "[1,3]"
  }
}
```

- [outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-1130033011120230-1313001331200331-1003100020122121-0222023332233321-3112121012033033-1001310210132203-2112312311112212-1033200030012022): complete subsection reference.

- [site_local_network](resources--gcp_vpc_site--reference--group-004.md#canonical-0031010113331203-0332021211122002-2133232221022033-1323212302220010-1332220012121111-1233120030320131-0003300032001033-1013330221332200): complete subsection reference.

- [site_local_subnet](resources--gcp_vpc_site--reference--group-005.md#canonical-1010131003002313-2213332312132020-0100123302320021-3220331133311300-2021223122230031-3212103010132011-0232201330012223-3030111231322233): complete subsection reference.

- [sm_connection_public_ip](resources--gcp_vpc_site--reference--group-005.md#canonical-2110313330032303-3110333222121201-2220303222032313-0202012020100132-0013113323310002-2221121220013110-3121003203331012-0010232203321022): complete subsection reference.

- [sm_connection_pvt_ip](resources--gcp_vpc_site--reference--group-005.md#canonical-0331101100010033-0222011110310013-2032302133123312-0220133021132300-2101231320333020-0020000111202113-1303322132113030-2132013223233320): complete subsection reference.

- [storage_class_list](resources--gcp_vpc_site--reference--group-005.md#canonical-1112333322110321-0200101320023311-3310121222132330-0033000133132300-1232002332132331-2111300311231333-1211022132123111-1313233202133121): complete subsection reference.

<a id="canonical-3022131203300321-0123002323021111-0033032110003222-1303323221111312-3210010030113111-0123010133300012-2003120232201300-2131013131130011"></a>

## Next pages — voltstack_cluster / 213213010113 / 7

- [voltstack_cluster.active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-0013032210231210-3130232233133123-2030120322033300-1111200101331012-3111132210102213-2000000003310111-3113230013210211-0332200121031300)
- [voltstack_cluster.active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-3312111010301003-0130023032013332-2022321323101201-1121230133310020-1200001301312323-1320322210222110-0311310110321013-2212131111322023)
- [voltstack_cluster.active_network_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-0121131301010003-2322312003102100-0230100310312000-1223202211131010-2031101011102012-1231322201013023-2101020023031123-0301132031301310)
- [voltstack_cluster.dc_cluster_group](resources--gcp_vpc_site--reference--group-004.md#canonical-1130111002232032-1120202321002331-2003200330312011-0003133202133132-1221210333322301-1002030011203132-0132310110222110-3013101300023101)
- [voltstack_cluster.default_storage](resources--gcp_vpc_site--reference--group-004.md#canonical-1032233013101222-0120133202121031-3222110133200101-0201212113311212-0022312013322011-3211100333221110-1103010130210302-3331221303102023)
- [voltstack_cluster.forward_proxy_allow_all](resources--gcp_vpc_site--reference--group-004.md#canonical-2311102031021112-0302330203020300-1010230201123312-2131223030033213-2012231222131310-2311310031112221-2200312223322013-2223211321120130)
- [voltstack_cluster.global_network_list](resources--gcp_vpc_site--reference--group-004.md#canonical-1333132321010123-1313031331321231-1212331120001123-0100212001231131-2213013320113322-3210103133122120-3313133330033113-2130303130210020)
- [voltstack_cluster.k8s_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-1222231131333230-0113000121132120-1000132321133333-3212021230133132-2222002111312120-0020012232033023-2322111300303223-0012200200102232)
- [voltstack_cluster.no_dc_cluster_group](resources--gcp_vpc_site--reference--group-004.md#canonical-1001131233200020-3032331103202032-3001200301021233-0101231212103123-3300231033202332-2112201210110031-3322113300220030-1133213311301222)
- [voltstack_cluster.no_forward_proxy](resources--gcp_vpc_site--reference--group-004.md#canonical-3100322021212032-1020223333302222-1122300021021102-2110033203300300-2210010312230213-0233330330112311-3101332030332101-3201000100303332)
- [voltstack_cluster.no_global_network](resources--gcp_vpc_site--reference--group-004.md#canonical-2311111001231211-2000113311333023-1332000013323020-2211113333001021-0101313330022200-0223222133201123-0122220300120131-0331111322221120)
- [voltstack_cluster.no_k8s_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-3220011312322231-3322211113300132-2102320122223311-2121120221032120-2132333033002120-3132203030312211-2021120130123012-2200121020020220)
- [voltstack_cluster.no_network_policy](resources--gcp_vpc_site--reference--group-004.md#canonical-0121032232121200-3202013002313021-0001101332201023-3023030321122233-1232023113311330-0201210210012230-1332013213021110-2231111123103313)
- [voltstack_cluster.no_outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-3221122321322132-0123030132333130-1100030013030230-1211310212210132-0322113011330122-0331022011121130-2021301132200011-2203120313103121)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-1130033011120230-1313001331200331-1003100020122121-0222023332233321-3112121012033033-1001310210132203-2112312311112212-1033200030012022)
- [voltstack_cluster.site_local_network](resources--gcp_vpc_site--reference--group-004.md#canonical-0031010113331203-0332021211122002-2133232221022033-1323212302220010-1332220012121111-1233120030320131-0003300032001033-1013330221332200)
- [voltstack_cluster.site_local_subnet](resources--gcp_vpc_site--reference--group-005.md#canonical-1010131003002313-2213332312132020-0100123302320021-3220331133311300-2021223122230031-3212103010132011-0232201330012223-3030111231322233)
- [voltstack_cluster.sm_connection_public_ip](resources--gcp_vpc_site--reference--group-005.md#canonical-2110313330032303-3110333222121201-2220303222032313-0202012020100132-0013113323310002-2221121220013110-3121003203331012-0010232203321022)
- [voltstack_cluster.sm_connection_pvt_ip](resources--gcp_vpc_site--reference--group-005.md#canonical-0331101100010033-0222011110310013-2032302133123312-0220133021132300-2101231320333020-0020000111202113-1303322132113030-2132013223233320)
- [voltstack_cluster.storage_class_list](resources--gcp_vpc_site--reference--group-005.md#canonical-1112333322110321-0200101320023311-3310121222132330-0033000133132300-1232002332132331-2111300311231333-1211022132123111-1313233202133121)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0013032210231210-3130232233133123-2030120322033300-1111200101331012-3111132210102213-2000000003310111-3113230013210211-0332200121031300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100211032123333-0301103222233102-0200313202012232-1200022313131332-3102220022320012-2001131001022331-1232112332021002-3300120220213001"></a>

## voltstack_cluster.active_enhanced_firewall_policies — active_enhanced_firewall_policies / 000102232120 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.active_enhanced_firewall_policies

<a id="canonical-0323220302231320-0232121331330322-3320303232313213-0202233102310123-1323332032313110-2220323202020133-3113012302012012-1212020202333333"></a>

Type: `"object"`. single nested block, Optional.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
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
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223203222330023-3120233212233102-1331201222031302-3023010033101221-2200310232230010-0010220211001213-1200221302023101-3112112110123022"></a>

## Direct properties — active_enhanced_firewall_policies / 000102232120 / 3

- [enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-0331111112220221-1311300312212030-3311332300031020-2220011321132111-1102021130203033-1100300020311321-3331012103023021-1033110010123312): complete subsection reference.

<a id="canonical-0333233020322201-1031300102330133-1301311311332330-1321023203200010-0210101030123300-1021110133233310-1002331132003301-1032132201010130"></a>

## Next pages — active_enhanced_firewall_policies / 000102232120 / 4

- [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-0331111112220221-1311300312212030-3311332300031020-2220011321132111-1102021130203033-1100300020311321-3331012103023021-1033110010123312)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0331111112220221-1311300312212030-3311332300031020-2220011321132111-1102021130203033-1100300020311321-3331012103023021-1033110010123312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103231302001220-1023310231012110-1201301031022320-0132002130310122-1311102223102013-1111012310223211-3303220020201033-2203030110133201"></a>

## voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies — enhanced_firewall_policies / 031101030031 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-0013032210231210-3130232233133123-2030120322033300-1111200101331012-3111132210102213-2000000003310111-3113230013210211-0332200121031300)
- voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-0122121330302212-2112001103023230-1131102223123110-2023333210321233-2010230133102320-2313030021213233-1011213012210322-0200322333032313"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222012113202110-1300023231011111-3012000301310330-0013120110321111-1022311320321332-2232210103032100-0133030133320102-2301331310221030"></a>

## Direct properties — enhanced_firewall_policies / 031101030031 / 3

<a id="canonical-3032232303130002-3233231133013133-1212010220221220-3031121202032100-3003201332233103-1333133131311223-1230321012010221-1013201123323100"></a>

<a id="canonical-3121130123100223-3320331212111102-0301223303301313-1323111013323033-2013231222002001-2310003131120320-3320321303132021-0122121201331300"></a>

## name property — enhanced_firewall_policies / 031101030031 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1121020223121322-1331021002022321-0023013330031133-0121013330121313-1313210322220123-3223103121031033-0330313002123120-2130123011032003"></a>

<a id="canonical-3311302302033312-2011311102122320-0132033212323221-0101320223230101-3322323221111321-3031112122202310-3131321313222321-2201011211200120"></a>

## namespace property — enhanced_firewall_policies / 031101030031 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0222211103322213-0331130333323301-2021220300113101-0331121011310030-0112212010332120-3023230101000031-0021320103020302-2310312322010123"></a>

<a id="canonical-0020012223213223-1003232202203022-3130223313122302-3210300031003320-0111211121211103-2022222003131003-2212302212113320-1123001333023103"></a>

## tenant property — enhanced_firewall_policies / 031101030031 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3012103332230302-0001332302021100-0212211212011022-3002302321021331-2320212200021033-0320123032322302-2032212312322322-0122121000112102"></a>

## Next pages — enhanced_firewall_policies / 031101030031 / 7

- [voltstack_cluster.active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-0013032210231210-3130232233133123-2030120322033300-1111200101331012-3111132210102213-2000000003310111-3113230013210211-0332200121031300)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3312111010301003-0130023032013332-2022321323101201-1121230133310020-1200001301312323-1320322210222110-0311310110321013-2212131111322023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310331303032012-2020212013023102-3201320102013331-1123332220311110-2231213220333130-1113023330020210-0130211303232111-2011201023113013"></a>

## voltstack_cluster.active_forward_proxy_policies — active_forward_proxy_policies / 320211212102 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.active_forward_proxy_policies

<a id="canonical-0311011300332003-3302003020211322-2321323021121001-2321311220112122-0223333001031220-0133332332110312-2022020323330003-2212001322312331"></a>

Type: `"object"`. single nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1331012101031101-0103101131020012-0232312112120232-2320320023221303-0000203123313311-0120013233110123-2001301331131000-0213023232323100"></a>

## Direct properties — active_forward_proxy_policies / 320211212102 / 3

- [forward_proxy_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-2332132212032322-2013300322213000-2113133211030110-1231301203030211-0321011220011010-3011000133102203-1313220023200232-1310231222100012): complete subsection reference.

<a id="canonical-1021220023112310-0303223303303111-1013132121221213-2301220300101013-2032321121330203-0021000103010313-2230111101231312-0333122100103103"></a>

## Next pages — active_forward_proxy_policies / 320211212102 / 4

- [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-2332132212032322-2013300322213000-2113133211030110-1231301203030211-0321011220011010-3011000133102203-1313220023200232-1310231222100012)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2332132212032322-2013300322213000-2113133211030110-1231301203030211-0321011220011010-3011000133102203-1313220023200232-1310231222100012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232001203120322-3332201312123133-0112320301232322-3011033123231003-0231320231200102-3323203311320031-0012232132103002-2130111122223322"></a>

## voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies — forward_proxy_policies / 300011100113 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-3312111010301003-0130023032013332-2022321323101201-1121230133310020-1200001301312323-1320322210222110-0311310110321013-2212131111322023)
- voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-1213012333101200-2003333223031010-2001323131021000-0312011102013000-2220330020033121-2303010202232230-1223020103301300-2220131332033201"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1113310013311120-1313202103011213-2321300102123302-2320203003033013-0313122201232321-0230322103000201-3120022130032120-2120202300002021"></a>

## Direct properties — forward_proxy_policies / 300011100113 / 3

<a id="canonical-0113103220223013-0111231033333130-1330222312020302-2131331002021223-2222320232302313-0330011320011101-2120122032313233-2130311021300231"></a>

<a id="canonical-2331001121333102-3200220023102130-1003031312002033-3102030210213111-1321113131212003-2013103233122330-3303100210321023-1133221221202202"></a>

## name property — forward_proxy_policies / 300011100113 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3021003223131231-2220032310033113-3133202111221220-1120203111230112-3203030003101030-3122011133302111-2301020303232322-1331121312201333"></a>

<a id="canonical-3131311330103220-1322103301200211-3131312010110010-0221001110123321-2110133200001323-3223100311121101-3121101101233003-1023001221321113"></a>

## namespace property — forward_proxy_policies / 300011100113 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1332021102110202-1230001033223202-0200020210201220-3230003303123131-1121302030221021-1021102322023202-2311332331220121-0330323000110011"></a>

<a id="canonical-3002121303212331-1200200021111022-1001220203322322-1203032110023013-2230132132311110-0221033100022011-0213213302130030-1121202013123010"></a>

## tenant property — forward_proxy_policies / 300011100113 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3202112202201033-2201020131130332-3220202023111122-0002100313020201-1212123320031201-0100111302103113-1003323330310021-3230100223212333"></a>

## Next pages — forward_proxy_policies / 300011100113 / 7

- [voltstack_cluster.active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-3312111010301003-0130023032013332-2022321323101201-1121230133310020-1200001301312323-1320322210222110-0311310110321013-2212131111322023)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0121131301010003-2322312003102100-0230100310312000-1223202211131010-2031101011102012-1231322201013023-2101020023031123-0301132031301310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300213220121002-1122023022112333-3310112123333233-1203002023210032-1102231333113303-3000021233210230-0330201021030210-3332021012220200"></a>

## voltstack_cluster.active_network_policies — active_network_policies / 232121112321 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.active_network_policies

<a id="canonical-2022130321323021-2102210321321230-2030310212001022-3300332330032231-0201112031333301-3102000111132020-2233132103303032-2010012033313021"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2022323103310232-0312223321103101-3121022031102303-0032003032300133-1111333211303132-2200022312221301-2212302323311301-1211030210021233"></a>

## Direct properties — active_network_policies / 232121112321 / 3

- [network_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-3021213213201002-1313231123210031-0202002333013232-1211332103001131-1030333231203130-3100330320121000-3112312110010123-1011223110031222): complete subsection reference.

<a id="canonical-3013223000123203-3331310330103011-3303210100300301-2101300212020210-3002202002330201-1111010223130222-0332212300310022-0113120300131021"></a>

## Next pages — active_network_policies / 232121112321 / 4

- [voltstack_cluster.active_network_policies.network_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-3021213213201002-1313231123210031-0202002333013232-1211332103001131-1030333231203130-3100330320121000-3112312110010123-1011223110031222)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3021213213201002-1313231123210031-0202002333013232-1211332103001131-1030333231203130-3100330320121000-3112312110010123-1011223110031222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303123121022012-3133002323200313-1320010021033102-0231113003320333-3310330202331333-2331023310231220-2122300002001130-2331011003201210"></a>

## voltstack_cluster.active_network_policies.network_policies — network_policies / 001011020312 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.active_network_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-0121131301010003-2322312003102100-0230100310312000-1223202211131010-2031101011102012-1231322201013023-2101020023031123-0301132031301310)
- voltstack_cluster.active_network_policies.network_policies

<a id="canonical-2121121000131320-1003212221220303-0333332120202223-1101113230031311-3311131231211231-0003300202001121-1222021232330101-3201202002300112"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Firewall Policies active for this network firewall.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-0313033311122333-1300102201203301-0132212120031002-1222122022133233-1231223303013012-2002312032023013-3111211331003310-0123312011212322"></a>

## Direct properties — network_policies / 001011020312 / 3

<a id="canonical-1201230223321221-0033131311132102-3310111301301331-1122031121011330-2332313101113023-3011323232201100-3022303203013310-2312211001210231"></a>

<a id="canonical-2323030132230012-3323323131201223-0203200123133301-1003330312232210-3010033013320203-0002221132113212-1230030131112312-2003311133020313"></a>

## name property — network_policies / 001011020312 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0011222030213233-0030023021120331-3113212031321221-0130131022133111-0103221121331200-2322131000011311-3322202330133033-1301232202223302"></a>

<a id="canonical-3122122011330001-1202113103112031-2331021103311001-0002221022232221-2233212231010203-1203022022033033-0120211322131002-1321121123332233"></a>

## namespace property — network_policies / 001011020312 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1313200332231031-2233021131211132-0020212223123202-0320011012100100-3210230301103001-2232302030311033-2230001321220221-2200122000020231"></a>

<a id="canonical-2311131311222131-3303113321303301-2203313033100320-1232022330120101-2101101220202313-2013010223001332-0101201222221212-1320201230000112"></a>

## tenant property — network_policies / 001011020312 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2300203030222133-0111230231011330-1032020331121123-1111313321012323-0003301332220301-3001100310231023-0203101200211102-3333020330023012"></a>

## Next pages — network_policies / 001011020312 / 7

- [voltstack_cluster.active_network_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-0121131301010003-2322312003102100-0230100310312000-1223202211131010-2031101011102012-1231322201013023-2101020023031123-0301132031301310)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1130111002232032-1120202321002331-2003200330312011-0003133202133132-1221210333322301-1002030011203132-0132310110222110-3013101300023101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122313002321020-3202010223111213-1100113320310002-1210033300323001-0030221000200300-0203031310001313-0022023132233113-1201310223033221"></a>

## voltstack_cluster.dc_cluster_group — dc_cluster_group / 312302311010 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.dc_cluster_group

<a id="canonical-1011210220301302-3213010231210001-3011201022223230-2103011330131201-3302330222023010-0021111231102223-2303001132131020-3102011203221030"></a>

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
dc_cluster_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-0222122122112120-3323013120112022-0302010122230221-0032221212101230-2312122021002131-2230010111312132-2311233101110120-2331202300133302"></a>

## Direct properties — dc_cluster_group / 312302311010 / 3

<a id="canonical-1333210110211132-1330233030111012-1023100320203231-3000031201201302-2103032201332311-3103111031011023-0131123131000101-1023322211013330"></a>

<a id="canonical-0313020231202312-0130312331331122-1333001303212301-3033213322101021-0312300220211323-0330101303100011-1020221133333120-1200113213103020"></a>

## name property — dc_cluster_group / 312302311010 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1101132001130021-1313322101030332-1032313023010313-3302330131111200-3011122220220321-1123133122333000-3233220003023303-3221023020011000"></a>

<a id="canonical-2303000232333000-2221011212130020-3313311013032323-3200111201331112-0331031112011032-0021013313121322-3102212130321212-1313200301102013"></a>

## namespace property — dc_cluster_group / 312302311010 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1303021133231322-2201113001330223-3223203330223330-0001031020303030-2130122330031332-1000200232020212-2022103030310210-0232311212012203"></a>

<a id="canonical-1202332122231001-2322111002220321-1000202100120132-0122023031331001-3033022032302133-1233131112122203-0032123203032331-2333013103330213"></a>

## tenant property — dc_cluster_group / 312302311010 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1121102101123022-2100102331232023-1220212030011103-3132221310333012-3103320321303000-1333100132223102-2211033011210220-2032203213032201"></a>

## Next pages — dc_cluster_group / 312302311010 / 7

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1032233013101222-0120133202121031-3222110133200101-0201212113311212-0022312013322011-3211100333221110-1103010130210302-3331221303102023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033102302022133-1121022221221201-1133100103210012-2200213101011022-2121303310210222-1223102020011030-1211112132030102-0031321010031111"></a>

## voltstack_cluster.default_storage — default_storage / 320000130102 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.default_storage

<a id="canonical-3210210310130012-3131230000101321-3223100200303111-3211000132123211-0013133300323303-1230023330030001-3011220232131102-3303310001310113"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default storage.

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
default_storage = {}
```

<a id="canonical-2100200031032122-0011223101212323-3231232131032002-3202201331331132-0033132032230111-3322303002102302-1112111220020031-2002310111131022"></a>

## Direct properties — default_storage / 320000130102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022101013120023-0123010011300001-0310211210302013-2012202020321222-2203221333102233-0323302210111302-2220030231311301-3111330010320331"></a>

## Next pages — default_storage / 320000130102 / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2311102031021112-0302330203020300-1010230201123312-2131223030033213-2012231222131310-2311310031112221-2200312223322013-2223211321120130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132032030131121-2221212231122131-0213203223002130-3231010230122200-0333113201002230-2101111030100212-3123031032330033-1003003313313031"></a>

## voltstack_cluster.forward_proxy_allow_all — forward_proxy_allow_all / 331001220032 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.forward_proxy_allow_all

<a id="canonical-2302130023333312-1131022320130133-3111001022331233-1222310232112122-2033120323220102-2111212331011023-0000220012100323-2222113320323302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for forward proxy allow all.

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
forward_proxy_allow_all = {}
```

<a id="canonical-0130321220202000-2330132223120233-3000111012130233-3300111011031132-0123333011120300-2011021121302331-3103113332212320-3111232220001100"></a>

## Direct properties — forward_proxy_allow_all / 331001220032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111201220323120-3002021003031230-2300022203000212-1122220113332032-0302102321103010-3031301130121322-0120331013332022-3221001103321332"></a>

## Next pages — forward_proxy_allow_all / 331001220032 / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1333132321010123-1313031331321231-1212331120001123-0100212001231131-2213013320113322-3210103133122120-3313133330033113-2130303130210020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210223321113223-3012230310323130-1111000001323221-2312323321003313-1302103222331202-2021121201023121-2122000330200320-0120121033000323"></a>

## voltstack_cluster.global_network_list — global_network_list / 203301113111 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.global_network_list

<a id="canonical-0211002121113202-2211002211223231-3030300202113003-0221001022123333-1103210312332321-2323321331200131-0221331310031331-0023101131120110"></a>

Type: `"object"`. single nested block, Optional.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("global_network_connections")}
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
global_network_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3231221231213312-2320021031303210-1100211332021230-0333030102110130-3111033320320201-0131230203320301-1102121200030010-2333310232220112"></a>

## Direct properties — global_network_list / 203301113111 / 3

- [global_network_connections](resources--gcp_vpc_site--reference--group-004.md#canonical-0333321200210120-2132332211222122-3132103332312330-0333001102023223-2221310301111111-0103031323230330-0000001120233311-1123033011203223): complete subsection reference.

<a id="canonical-1223330011123110-3012202300032311-1313330111122002-0223212120010222-2033311120022132-3310001330001202-1110011330001102-3301020102030030"></a>

## Next pages — global_network_list / 203301113111 / 4

- [voltstack_cluster.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-004.md#canonical-0333321200210120-2132332211222122-3132103332312330-0333001102023223-2221310301111111-0103031323230330-0000001120233311-1123033011203223)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0333321200210120-2132332211222122-3132103332312330-0333001102023223-2221310301111111-0103031323230330-0000001120233311-1123033011203223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032222102332011-0110313123223030-1003332331203033-3212012321232301-0310102332210312-1020000232011130-1002312002312303-0123011022001303"></a>

## voltstack_cluster.global_network_list.global_network_connections — global_network_connections / 203201032032 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.global_network_list](resources--gcp_vpc_site--reference--group-004.md#canonical-1333132321010123-1313031331321231-1212331120001123-0100212001231131-2213013320113322-3210103133122120-3313133330033113-2130303130210020)
- voltstack_cluster.global_network_list.global_network_connections

<a id="canonical-1131312233313233-3131120211100310-0300223313213031-0331331323022120-1021321330133130-3103233222003111-1212030222121003-3222022020031103"></a>

Type: `"object"`. list nested block, Optional.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("sli_to_global_dr",
    "slo_to_global_dr")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
global_network_connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020301023301233-0231110022212122-0201213030021301-0222331321310321-3122323011122131-2133323313312221-0221130000200113-2231031013023113"></a>

## Direct properties — global_network_connections / 203201032032 / 3

- [sli_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-1223021200110313-3032022112320302-3132320102320320-1100033103003330-1123230122103223-0233010300122312-3002131022303033-3100211113112320): complete subsection reference.

- [slo_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-2122030001130130-0220230130323003-3232223111300300-2302302030233002-2231311330302321-0113000331012110-0220123332111310-3130023023031131): complete subsection reference.

<a id="canonical-1120022300013200-0103312221000003-2310003002131020-1333120212030200-2021221202213201-2222210113332132-2231233031221122-3222321212011331"></a>

## Next pages — global_network_connections / 203201032032 / 4

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-1223021200110313-3032022112320302-3132320102320320-1100033103003330-1123230122103223-0233010300122312-3002131022303033-3100211113112320)
- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-2122030001130130-0220230130323003-3232223111300300-2302302030233002-2231311330302321-0113000331012110-0220123332111310-3130023023031131)
- [voltstack_cluster.global_network_list](resources--gcp_vpc_site--reference--group-004.md#canonical-1333132321010123-1313031331321231-1212331120001123-0100212001231131-2213013320113322-3210103133122120-3313133330033113-2130303130210020)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1223021200110313-3032022112320302-3132320102320320-1100033103003330-1123230122103223-0233010300122312-3002131022303033-3100211113112320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023211020311002-2212323231203031-3031000323110103-2131201020112310-2000020233113030-3211213121012013-3003000210011331-0212300322112300"></a>

## voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr — sli_to_global_dr / 311203203133 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.global_network_list](resources--gcp_vpc_site--reference--group-004.md#canonical-1333132321010123-1313031331321231-1212331120001123-0100212001231131-2213013320113322-3210103133122120-3313133330033113-2130303130210020)
- [voltstack_cluster.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-004.md#canonical-0333321200210120-2132332211222122-3132103332312330-0333001102023223-2221310301111111-0103031323230330-0000001120233311-1123033011203223)
- voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-0221012313020012-0230333232210333-3211201322332301-0220201212323313-3322201330001020-1331321203313000-1000001213001323-1022031001010223"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
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
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-1001223320003000-0202121130023302-0200223120210113-3010323310330231-2313212302230303-3223001122320032-2021202021203220-2222300103320100"></a>

## Direct properties — sli_to_global_dr / 311203203133 / 3

- [global_vn](resources--gcp_vpc_site--reference--group-004.md#canonical-2113220212012222-0120031223310112-2210113010132102-3200131330030323-0031330112113320-1013123210031221-0121001303220020-3310110122001030): complete subsection reference.

<a id="canonical-0203030220022101-0012221102223310-3032311122120131-2003221113122330-2022020312110113-2022333000300112-3033222331012310-1011130213022022"></a>

## Next pages — sli_to_global_dr / 311203203133 / 4

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--gcp_vpc_site--reference--group-004.md#canonical-2113220212012222-0120031223310112-2210113010132102-3200131330030323-0031330112113320-1013123210031221-0121001303220020-3310110122001030)
- [voltstack_cluster.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-004.md#canonical-0333321200210120-2132332211222122-3132103332312330-0333001102023223-2221310301111111-0103031323230330-0000001120233311-1123033011203223)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2113220212012222-0120031223310112-2210113010132102-3200131330030323-0031330112113320-1013123210031221-0121001303220020-3310110122001030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302221201000000-2021000130213303-2000121133032130-0123332033302202-1010123130202202-1121223110303201-3222020010211123-2001331320323100"></a>

## voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn — global_vn / 021121312001 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.global_network_list](resources--gcp_vpc_site--reference--group-004.md#canonical-1333132321010123-1313031331321231-1212331120001123-0100212001231131-2213013320113322-3210103133122120-3313133330033113-2130303130210020)
- [voltstack_cluster.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-004.md#canonical-0333321200210120-2132332211222122-3132103332312330-0333001102023223-2221310301111111-0103031323230330-0000001120233311-1123033011203223)
- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-1223021200110313-3032022112320302-3132320102320320-1100033103003330-1123230122103223-0233010300122312-3002131022303033-3100211113112320)
- voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-3010002033200322-2201200101303221-0213230322300033-2222033323001211-0021021010121311-0110113022211221-3031133131033121-2220110323212312"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100303332112311-1123103213013220-1013031030330001-1332011333201230-1212121201222022-2032212200330130-1233223122221113-0113111220011120"></a>

## Direct properties — global_vn / 021121312001 / 3

<a id="canonical-0002311201021301-0122021323123213-1232211012331330-1121332323001020-3302200313002002-1113112320233130-1332333230222212-0112112212211223"></a>

<a id="canonical-3122223222332130-3223130111203102-1233330132220222-0300322103123212-0211011101121303-2233233300012032-3133112332031322-2002220330103031"></a>

## name property — global_vn / 021121312001 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3001101132320132-2220330330033020-0330130030001023-2021032331132111-0033003331230130-0003310321310221-2223233221102133-1131322121212203"></a>

<a id="canonical-3003302012031202-3302203203212320-2233021021130233-1232221013102103-0303110101102131-3211030210320121-0131123122012322-0100133223230223"></a>

## namespace property — global_vn / 021121312001 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1112030222023331-3211101220032100-2031330202210020-0323231333212203-3230120213300230-2112123033210031-3322012201021113-0122210031201121"></a>

<a id="canonical-2232330312111021-1303021122122121-0133131031032213-1331303211011211-3112100133032220-0311203331100110-1003023002002333-2111021311221012"></a>

## tenant property — global_vn / 021121312001 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0133000112331330-3231103200133211-0220223102313220-0113320321313222-2223002213010111-1331101131013101-2313010203200132-1000010110023031"></a>

## Next pages — global_vn / 021121312001 / 7

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-1223021200110313-3032022112320302-3132320102320320-1100033103003330-1123230122103223-0233010300122312-3002131022303033-3100211113112320)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2122030001130130-0220230130323003-3232223111300300-2302302030233002-2231311330302321-0113000331012110-0220123332111310-3130023023031131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112333212121121-0001322231032302-1213133230203032-1031003123203222-3232232202201221-1331031303033221-3023320103332101-3110012103232023"></a>

## voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr — slo_to_global_dr / 211220201101 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.global_network_list](resources--gcp_vpc_site--reference--group-004.md#canonical-1333132321010123-1313031331321231-1212331120001123-0100212001231131-2213013320113322-3210103133122120-3313133330033113-2130303130210020)
- [voltstack_cluster.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-004.md#canonical-0333321200210120-2132332211222122-3132103332312330-0333001102023223-2221310301111111-0103031323230330-0000001120233311-1123033011203223)
- voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-3100001031330120-3201033101033122-2001011132322031-2021031200200020-0033002121222221-3102333203130232-3133231200301030-3202332310331210"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
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
slo_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-0021201102022233-3302100011301003-0033112301000213-0223302220303031-1231010023103103-0232222330301300-0112120331221211-3222022310311020"></a>

## Direct properties — slo_to_global_dr / 211220201101 / 3

- [global_vn](resources--gcp_vpc_site--reference--group-004.md#canonical-2003010032312231-1030222203002000-2120212201221301-1002111311010033-1232011331123311-0323310320231130-0002130320110202-0302032102202130): complete subsection reference.

<a id="canonical-1213121311122022-3000101310010113-2022212121133333-0012101331131001-0030311013311003-2002313121230313-3231202122021311-2122201031321313"></a>

## Next pages — slo_to_global_dr / 211220201101 / 4

- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--gcp_vpc_site--reference--group-004.md#canonical-2003010032312231-1030222203002000-2120212201221301-1002111311010033-1232011331123311-0323310320231130-0002130320110202-0302032102202130)
- [voltstack_cluster.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-004.md#canonical-0333321200210120-2132332211222122-3132103332312330-0333001102023223-2221310301111111-0103031323230330-0000001120233311-1123033011203223)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2003010032312231-1030222203002000-2120212201221301-1002111311010033-1232011331123311-0323310320231130-0002130320110202-0302032102202130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131223110033310-0230220213313122-0132111132200313-2232321130100210-0332111303301210-3230003102033003-3300311203302103-1122022103120022"></a>

## voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn — global_vn / 032030003313 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.global_network_list](resources--gcp_vpc_site--reference--group-004.md#canonical-1333132321010123-1313031331321231-1212331120001123-0100212001231131-2213013320113322-3210103133122120-3313133330033113-2130303130210020)
- [voltstack_cluster.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-004.md#canonical-0333321200210120-2132332211222122-3132103332312330-0333001102023223-2221310301111111-0103031323230330-0000001120233311-1123033011203223)
- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-2122030001130130-0220230130323003-3232223111300300-2302302030233002-2231311330302321-0113000331012110-0220123332111310-3130023023031131)
- voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-1303102012101312-3313102330012020-3002310303123303-1211010220222212-1232000323013322-3212302110001333-0210031333311302-0000011022032031"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020321230101230-0133021210321010-1133323321130102-0010011202102002-0323333103221302-1202211322030021-3010103130131033-3223023011000013"></a>

## Direct properties — global_vn / 032030003313 / 3

<a id="canonical-1231023201212121-1030120002120322-0031202131102203-3032303012111111-1211312013212113-0312103231323101-0220002123110313-1310232231212130"></a>

<a id="canonical-0212112133232110-3331121232233332-1131130120321123-0132122010101010-2002011001010031-1122323101332310-0312003310223121-2202020023213203"></a>

## name property — global_vn / 032030003313 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1313223202322311-1020202131332120-1313133130322002-2130330212302201-2131203022100002-3000101122100120-2320112021231322-0101221030322202"></a>

<a id="canonical-3021023000212202-1013122231031213-2222330021310133-3232003021230230-1313200231020103-3111101230033120-2121301311101030-1132213203230221"></a>

## namespace property — global_vn / 032030003313 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1223230113233022-2231322031000002-0003102013311011-0131031111030000-1303112220300302-2130303311322002-3032003120100210-0022313213030130"></a>

<a id="canonical-0003201203203333-0202301321123002-0133330220022011-3010200202312313-2332303000003021-0230320033132030-0030302202323321-0133021033310230"></a>

## tenant property — global_vn / 032030003313 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1122113013011111-1012110230111133-1001101022233103-3111210302121113-1002111232021000-3101021031321012-3102113021302330-0330311111233202"></a>

## Next pages — global_vn / 032030003313 / 7

- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-2122030001130130-0220230130323003-3232223111300300-2302302030233002-2231311330302321-0113000331012110-0220123332111310-3130023023031131)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1222231131333230-0113000121132120-1000132321133333-3212021230133132-2222002111312120-0020012232033023-2322111300303223-0012200200102232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012313232310111-3010021320133302-3030122131322322-0323033200210230-2300212002330022-0321330110322200-3330123113101133-3032330130132022"></a>

## voltstack_cluster.k8s_cluster — k8s_cluster / 213013123212 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.k8s_cluster

<a id="canonical-0223331013200000-3123101202112122-3111022230211032-2103012130110220-0200332103130310-1030001330131212-0310023032123110-1113330203222112"></a>

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
k8s_cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-0212132330320131-0102201130213032-0230221112030123-0201013231220132-0122002300113101-1201102010113330-3331133200221310-1320321201110210"></a>

## Direct properties — k8s_cluster / 213013123212 / 3

<a id="canonical-3133302130131000-0121230133122000-0221201103130132-3222010231130000-0001321323112121-3132113031111201-2313102232021031-1123213301123123"></a>

<a id="canonical-0312010203332203-0310021001132023-3220021131323021-2100320213013121-2233301030131030-0202022221232302-2020303030211210-1200320321333202"></a>

## name property — k8s_cluster / 213013123212 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0221210111323232-0221123303101100-1000311233312002-3113121023320131-1032030121001033-3303323301113202-3121021010023310-2032322110010120"></a>

<a id="canonical-3011031303012013-2210012210203332-2303313001123231-1112002312030231-0310213321103021-2301000223001211-1012111200101310-2001321203110321"></a>

## namespace property — k8s_cluster / 213013123212 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1102312111120301-0312332031320002-1222201331132310-3323202003032001-1200112331331031-0111030000330233-2321101332012130-2031233203020123"></a>

<a id="canonical-0121223033021031-3002233303330111-0222221102302313-2320030022131033-3332313201032131-2331213013333212-0200011301323133-1333131213303211"></a>

## tenant property — k8s_cluster / 213013123212 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1121233333030322-0220220032132113-1330021010112121-0110123103223320-3023231132101330-1322220313022321-1210023332310320-0320222320003121"></a>

## Next pages — k8s_cluster / 213013123212 / 7

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1001131233200020-3032331103202032-3001200301021233-0101231212103123-3300231033202332-2112201210110031-3322113300220030-1133213311301222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331121003123102-3112221220113233-2212333310332213-2330302112111101-2300132210000222-3011331231103301-3002133220210230-0002132200011303"></a>

## voltstack_cluster.no_dc_cluster_group — no_dc_cluster_group / 122212311221 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.no_dc_cluster_group

<a id="canonical-3220201320120000-1011012212201120-0310330213001223-0210320030203313-3211231132013303-2211210002203233-3010322020031222-0321110112101112"></a>

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
no_dc_cluster_group = {}
```

<a id="canonical-3101323221331103-2131122230023032-0102102210331203-0133200102222310-2122021210133003-1101101221030221-0003120021230113-1121200003321103"></a>

## Direct properties — no_dc_cluster_group / 122212311221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101022110232223-0232220221220220-3332331131200231-0002211111203233-0310232031121101-3232121311331230-3101202203233012-1213233330010321"></a>

## Next pages — no_dc_cluster_group / 122212311221 / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3100322021212032-1020223333302222-1122300021021102-2110033203300300-2210010312230213-0233330330112311-3101332030332101-3201000100303332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231110222110000-2320222003003320-2101010130121210-2221130231332323-1212111233202300-2113221311012301-3133311103301322-3022320233020202"></a>

## voltstack_cluster.no_forward_proxy — no_forward_proxy / 203120103003 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.no_forward_proxy

<a id="canonical-2120011113003203-1101121111012011-0113133322300021-1022331003233000-1202110301031300-1211321302032012-1330310111133101-3311100211330022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no forward proxy.

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
no_forward_proxy = {}
```

<a id="canonical-0120022202301113-1201132213112023-2132122322122320-0232301020300102-3211123322220122-3101101012233223-3201000320102103-3320101103330210"></a>

## Direct properties — no_forward_proxy / 203120103003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323132001011332-0112022010033310-1011331321302330-2132323031130011-1212133003313103-1322113112012222-2022320100300122-3030133031302222"></a>

## Next pages — no_forward_proxy / 203120103003 / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2311111001231211-2000113311333023-1332000013323020-2211113333001021-0101313330022200-0223222133201123-0122220300120131-0331111322221120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330300220213300-3223001331022321-3020213321320133-2320330001313222-2133030312012210-1201233112133100-3031133331111133-1313222012001213"></a>

## voltstack_cluster.no_global_network — no_global_network / 223203212111 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.no_global_network

<a id="canonical-0031301202330333-3002113133000202-0313100033022232-3133200011301112-1302220032203030-1211231211212300-3232022003123020-1332122211012313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no global network.

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
no_global_network = {}
```

<a id="canonical-2002013201032032-2022303000033210-3102203221031120-0201301303211310-1011132130000121-3031302120011111-0101310312120303-2322320311230001"></a>

## Direct properties — no_global_network / 223203212111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0021201321010012-3200200123212011-0133122221201231-2022122032332030-0200001021311122-0101120203301012-2120202022303310-2313233303123320"></a>

## Next pages — no_global_network / 223203212111 / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3220011312322231-3322211113300132-2102320122223311-2121120221032120-2132333033002120-3132203030312211-2021120130123012-2200121020020220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033320201200010-2313313101301031-3110233002000200-1110113203132303-0223122112103013-0213030203032331-0020323122122222-0121203223101230"></a>

## voltstack_cluster.no_k8s_cluster — no_k8s_cluster / 233203310221 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.no_k8s_cluster

<a id="canonical-1122212012233000-2223100212222031-3021033010301210-0333123130211002-0312301132120222-3201101303002110-2032331121221310-2100122031131102"></a>

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
no_k8s_cluster = {}
```

<a id="canonical-1323232110002032-2302301132203320-1132021221132123-1321323131322002-0223312323312311-2313121212221303-0031232223231323-3232011023000113"></a>

## Direct properties — no_k8s_cluster / 233203310221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301320032212102-0032203020301110-1211102211033331-3323103012031002-0023333130331103-1022321100132301-3220001211210000-2312202110022112"></a>

## Next pages — no_k8s_cluster / 233203310221 / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0121032232121200-3202013002313021-0001101332201023-3023030321122233-1232023113311330-0201210210012230-1332013213021110-2231111123103313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331100001203030-2201211210013313-0233201000021333-1331300102013202-1331333311022310-2311001100230322-3323101110110003-1311211023020022"></a>

## voltstack_cluster.no_network_policy — no_network_policy / 001012320332 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.no_network_policy

<a id="canonical-2320212130100013-1030210112323321-0112102033333200-0301003023021003-0230122210312201-2221000100211301-2300313322030233-2101223123021130"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
no_network_policy = {}
```

<a id="canonical-2020132103222132-2013223011031201-0212311331103230-0322000211103021-3112023111313333-0102312010000211-0002212233203220-1000231023222112"></a>

## Direct properties — no_network_policy / 001012320332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302133222102210-3032201011323131-1131133310323223-1333322201013020-3013012312231221-3130203301113231-3332202132322213-3120333100310101"></a>

## Next pages — no_network_policy / 001012320332 / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3221122321322132-0123030132333130-1100030013030230-1211310212210132-0322113011330122-0331022011121130-2021301132200011-2203120313103121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133223210012010-3111232112203220-1201112112223200-3310322012120113-0012113121212102-0123003313032312-1002222233033122-1302310320110132"></a>

## voltstack_cluster.no_outside_static_routes — no_outside_static_routes / 202320233232 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.no_outside_static_routes

<a id="canonical-0103131130011222-3100230002202330-1311223202303203-2112031022122002-3033013313030021-2200201201232302-0202120113201123-1232033333313012"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no outside static routes.

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
no_outside_static_routes = {}
```

<a id="canonical-2011331210031312-3131320223102232-1012221023233202-3230013221332022-2332120200130031-0322301111031023-3030303203301213-3122110130323221"></a>

## Direct properties — no_outside_static_routes / 202320233232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102000222313102-2302303313210130-1033302300231310-3222300233110210-1331312220231033-0232011011200312-1032231211010000-1020202320200020"></a>

## Next pages — no_outside_static_routes / 202320233232 / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1130033011120230-1313001331200331-1003100020122121-0222023332233321-3112121012033033-1001310210132203-2112312311112212-1033200030012022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211302201311210-2212021223212332-1322002120032120-2010101032102333-3013113331120223-0313000023202200-1132121212033212-3211332111310200"></a>

## voltstack_cluster.outside_static_routes — outside_static_routes / 002012322203 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.outside_static_routes

<a id="canonical-3111233033111112-2103121333111022-2111001110303210-1132112013331012-1011030332211001-0200213031001200-3301111002301222-0021120011123301"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_route_list")}
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
outside_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2121100300322223-0202010122221312-1213112321102312-2023233311302022-1113110301331212-0133131331320102-3331111223130310-3232021021230002"></a>

## Direct properties — outside_static_routes / 002012322203 / 3

- [static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-2323300301222323-2121132313300323-1122222132001022-3210020300112222-1031203223230203-2030132233201302-0321110321203202-0122133200130111): complete subsection reference.

<a id="canonical-0001130303012122-2020210102200002-3210001121030212-3022010210103210-2103031230233111-1100032030130013-0001020100101022-1132102130330310"></a>

## Next pages — outside_static_routes / 002012322203 / 4

- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-2323300301222323-2121132313300323-1122222132001022-3210020300112222-1031203223230203-2030132233201302-0321110321203202-0122133200130111)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2323300301222323-2121132313300323-1122222132001022-3210020300112222-1031203223230203-2030132233201302-0321110321203202-0122133200130111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120203321201311-3010331212211133-1013133122032123-0321111233130011-0103310233300103-0322201122302313-0111302033031022-2101030103012203"></a>

## voltstack_cluster.outside_static_routes.static_route_list — static_route_list / 313201200201 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-1130033011120230-1313001331200331-1003100020122121-0222023332233321-3112121012033033-1001310210132203-2112312311112212-1033200030012022)
- voltstack_cluster.outside_static_routes.static_route_list

<a id="canonical-3002133220021122-2232023003113121-3322333222302220-2231233310330022-1103321233001111-0102102300200030-2211233303330010-2130310231303023"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_static_route",
    "simple_static_route")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
static_route_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0110002111221123-0110201110023230-0300320011100300-2013320102130111-3103203210203102-1230022312113101-3332000130213121-3101020002222202"></a>

## Direct properties — static_route_list / 313201200201 / 3

- [custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-0303323311112102-2312200213101013-3021213031100000-1330100032322330-2030133210101101-2223133101313022-2012220020030032-1230002002303321): complete subsection reference.

<a id="canonical-0330102112000330-1030023011222100-0310320100211300-1201320202321323-1032112100220011-2021310102212121-1203030132212333-1120201210232301"></a>

<a id="canonical-3013100120031233-1321030103330331-3101103003303030-2101001203330303-1232023301110202-3023211013211032-3323232300030303-0311322300101211"></a>

## simple_static_route property — static_route_list / 313201200201 / 4

Type: `"string"`. Optional.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-2103013003030331-2220120323331333-2031222322323211-0201211220113021-0203222300102031-3020230231321120-1012121312201000-2333003002332313"></a>

## Next pages — static_route_list / 313201200201 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-0303323311112102-2312200213101013-3021213031100000-1330100032322330-2030133210101101-2223133101313022-2012220020030032-1230002002303321)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-1130033011120230-1313001331200331-1003100020122121-0222023332233321-3112121012033033-1001310210132203-2112312311112212-1033200030012022)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0303323311112102-2312200213101013-3021213031100000-1330100032322330-2030133210101101-2223133101313022-2012220020030032-1230002002303321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232012101130113-0002202301301311-1210200031322102-0212232102311321-0112022030211102-2001122030033030-2002122113001210-0110120312232021"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route — custom_static_route / 021213232032 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-1130033011120230-1313001331200331-1003100020122121-0222023332233321-3112121012033033-1001310210132203-2112312311112212-1033200030012022)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-2323300301222323-2121132313300323-1122222132001022-3210020300112222-1031203223230203-2030132233201302-0321110321203202-0122133200130111)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-1121310230231302-1303033113022033-2113202202320303-1103011331332213-3103123130030100-0020223121320322-3133020113212320-3021331323032311"></a>

Type: `"object"`. single nested block, Optional.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnets")}
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
custom_static_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-1030200202002323-1312330003231033-0321033110233233-1212132111211012-2222332012000011-2222321200010011-3210333003312210-2012303311030111"></a>

## Direct properties — custom_static_route / 021213232032 / 3

<a id="canonical-1321002003333012-0010002320333322-2011312332332311-3122132202221001-3212120201213010-1321131233313231-1030221121220233-1322011110011321"></a>

<a id="canonical-0233321022321001-1100121021001100-0320102103121301-2220233203000330-3311002133302210-0100212032222010-1021002011133322-2231223200302122"></a>

## attrs property — custom_static_route / 021213232032 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [labels](resources--gcp_vpc_site--reference--group-004.md#canonical-0301331112230211-1010313110212100-1032031122323212-0002133013013331-3200120221203311-3113030103002130-2321130020232121-2132020330031110): complete subsection reference.

- [nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-3321011303101300-1322103221313021-2312001132221121-0220013122110221-0020102231100331-1200101232013232-3203121132230030-3202112131310221): complete subsection reference.

- [subnets](resources--gcp_vpc_site--reference--group-004.md#canonical-2113123120301030-1222222232103322-1321133111323110-1122013323201113-3331213001320233-2221113322120213-2033002322122323-2233300133221111): complete subsection reference.

<a id="canonical-1132200303033023-0210110323311032-1121123013311230-1133301211031102-2301002313231022-3202030211032200-2313110132012223-0333020320333100"></a>

## Next pages — custom_static_route / 021213232032 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](resources--gcp_vpc_site--reference--group-004.md#canonical-0301331112230211-1010313110212100-1032031122323212-0002133013013331-3200120221203311-3113030103002130-2321130020232121-2132020330031110)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-3321011303101300-1322103221313021-2312001132221121-0220013122110221-0020102231100331-1200101232013232-3203121132230030-3202112131310221)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-004.md#canonical-2113123120301030-1222222232103322-1321133111323110-1122013323201113-3331213001320233-2221113322120213-2033002322122323-2233300133221111)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-2323300301222323-2121132313300323-1122222132001022-3210020300112222-1031203223230203-2030132233201302-0321110321203202-0122133200130111)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0301331112230211-1010313110212100-1032031122323212-0002133013013331-3200120221203311-3113030103002130-2321130020232121-2132020330031110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101320112200023-1130101020023103-1023123200300230-1303000321012010-0220111212133323-2033220223201011-1210032330022301-3101122201330121"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels — labels / 332100320232 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-1130033011120230-1313001331200331-1003100020122121-0222023332233321-3112121012033033-1001310210132203-2112312311112212-1033200030012022)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-2323300301222323-2121132313300323-1122222132001022-3210020300112222-1031203223230203-2030132233201302-0321110321203202-0122133200130111)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-0303323311112102-2312200213101013-3021213031100000-1330100032322330-2030133210101101-2223133101313022-2012220020030032-1230002002303321)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-3101111101022131-1111200112303320-1311303001200023-3021200301211001-2020133320020120-1300030122100303-3232332100300322-3022213113231110"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
labels {}
```

<a id="canonical-1332112131213200-0121012131013311-0023023213223222-1112111022321120-0013321331323213-2120102033322110-1331301031322100-2010131221033330"></a>

## Direct properties — labels / 332100320232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011132332120303-3201133133020312-1131232303120012-0103002330321223-0213110131013212-2223211020211010-0213312211101000-1323311320011233"></a>

## Next pages — labels / 332100320232 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-0303323311112102-2312200213101013-3021213031100000-1330100032322330-2030133210101101-2223133101313022-2012220020030032-1230002002303321)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3321011303101300-1322103221313021-2312001132221121-0220013122110221-0020102231100331-1200101232013232-3203121132230030-3202112131310221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120123233103002-0213321223111123-2301230300230331-1302212310011132-2212231233031113-1233122111101210-0231313113322112-2313130310103303"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop — nexthop / 323302030202 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-1130033011120230-1313001331200331-1003100020122121-0222023332233321-3112121012033033-1001310210132203-2112312311112212-1033200030012022)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-2323300301222323-2121132313300323-1122222132001022-3210020300112222-1031203223230203-2030132233201302-0321110321203202-0122133200130111)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-0303323311112102-2312200213101013-3021213031100000-1330100032322330-2030133210101101-2223133101313022-2012220020030032-1230002002303321)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-2232231233200232-3100310330201323-1003131223003320-1313222222020333-2203010203020011-3223022120313121-3011100020220221-2013002020223133"></a>

Type: `"object"`. single nested block, Optional.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

Receipt-pinned upstream constraints:

```json
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
nexthop {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031100113102320-1021322313113220-3021131210003323-3130031233010300-0312021230302030-1011202230120003-1001000300022111-3130001320003101"></a>

## Direct properties — nexthop / 323302030202 / 3

- [interface](resources--gcp_vpc_site--reference--group-004.md#canonical-3131110000330121-3330221203010230-0201330210121032-3303011020230200-0210022223133211-3121223023320102-2333100230232113-2123321020331303): complete subsection reference.

- [nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-2000220311021312-0330123320202021-0211133102121202-3112313211122111-2232303332231022-0022011113203103-1113110300111023-0002032032200121): complete subsection reference.

<a id="canonical-0303112110013220-0311032211310332-2011131111010111-3110121130310010-3230333122231121-2310103030113222-3011221311033302-0030210222302110"></a>

<a id="canonical-0322320030023233-0332012123311311-1310203210100122-0310020133132330-1123023032121022-2112102212113112-1112321212320220-3220122332322233"></a>

## type property — nexthop / 323302030202 / 4

Type: `"string"`. Optional.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"),
}
```

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

<a id="canonical-3101312202230330-1033330011002123-1223133120320231-2311230200123120-2223333220201033-0322232022212231-1033323300110222-3333121110310223"></a>

## Next pages — nexthop / 323302030202 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--gcp_vpc_site--reference--group-004.md#canonical-3131110000330121-3330221203010230-0201330210121032-3303011020230200-0210022223133211-3121223023320102-2333100230232113-2123321020331303)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-2000220311021312-0330123320202021-0211133102121202-3112313211122111-2232303332231022-0022011113203103-1113110300111023-0002032032200121)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-0303323311112102-2312200213101013-3021213031100000-1330100032322330-2030133210101101-2223133101313022-2012220020030032-1230002002303321)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3131110000330121-3330221203010230-0201330210121032-3303011020230200-0210022223133211-3121223023320102-2333100230232113-2123321020331303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010331003103321-3020022023102210-2103222233313223-2323331131021002-2133201300101131-3001031022102132-3101312320033132-3113322310231333"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — interface / 233220313131 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-1130033011120230-1313001331200331-1003100020122121-0222023332233321-3112121012033033-1001310210132203-2112312311112212-1033200030012022)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-2323300301222323-2121132313300323-1122222132001022-3210020300112222-1031203223230203-2030132233201302-0321110321203202-0122133200130111)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-0303323311112102-2312200213101013-3021213031100000-1330100032322330-2030133210101101-2223133101313022-2012220020030032-1230002002303321)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-3321011303101300-1322103221313021-2312001132221121-0220013122110221-0020102231100331-1200101232013232-3203121132230030-3202112131310221)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-0302101032201101-1302022233021030-1012012122031020-2112312201132110-3231103013323001-3130111122122121-1132001101131313-3100233220010111"></a>

Type: `"object"`. list nested block, Optional.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-2133313223333213-1331013112022310-2200101032223331-0231322012101002-2113113200013203-0200110202101211-1002311110120310-2101320232023330"></a>

## Direct properties — interface / 233220313131 / 3

<a id="canonical-3222101001203201-3131211303032022-2301123122313031-1201233100121322-2102223333012131-0021210303323333-0230200200301311-0002220103001010"></a>

<a id="canonical-3233301032210310-0033032231010021-0031000331110223-1300312031121102-3223333000101211-2113120020222000-3111033232002310-3311012213220021"></a>

## kind property — interface / 233220313131 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3010212233232231-3013033113003123-1203302032231301-1211300120110333-2323300102031032-3231231033020320-0013222100030132-1233012301122312"></a>

<a id="canonical-2300222113032233-0120022010301101-2003103001333123-3201102030212312-1132033233130321-0121113331302132-1322332213303010-0011311102003023"></a>

## name property — interface / 233220313131 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0110310300111123-0132301311011212-2110322031132312-0220212322010131-1330013112301023-3211020310123312-1121301120022110-3333210131311130"></a>

<a id="canonical-0213021302133031-0203211303312001-2132010302311033-2023333320122321-1331002200010210-1110111212313110-3332120010200203-3230223322022203"></a>

## namespace property — interface / 233220313131 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0021211213223331-1131313201023311-1233133201323313-3130303321320210-0031120302212000-3300233030332022-1311202312232010-3223331312121302"></a>

<a id="canonical-1032021212021202-3320222123002210-0121033322023022-2111000302201100-0022012201223032-2330101000202311-2010101300020031-1322003321133203"></a>

## tenant property — interface / 233220313131 / 7

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3020223103022331-2312033002231032-3122311232212133-3001330033023003-0013320102111010-3220331331303223-3131321211012220-2311112202221033"></a>

<a id="canonical-2011333301133211-2333210023131232-0102120101100210-1221022023232301-1232103110120122-3000130220030130-2233130312202221-3020033213200000"></a>

## uid property — interface / 233220313131 / 8

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3000320010131211-1320323120321223-1110000002310030-3320033202233231-1202023023123100-3033130232123201-3213130231123210-2330332302101010"></a>

## Next pages — interface / 233220313131 / 9

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-3321011303101300-1322103221313021-2312001132221121-0220013122110221-0020102231100331-1200101232013232-3203121132230030-3202112131310221)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2000220311021312-0330123320202021-0211133102121202-3112313211122111-2232303332231022-0022011113203103-1113110300111023-0002032032200121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320032121201313-3023121030223003-0322323321131030-2330232321220112-0001311122002112-1030232203123300-3120321312103333-0322303013320312"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — nexthop_address / 313211021123 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-1130033011120230-1313001331200331-1003100020122121-0222023332233321-3112121012033033-1001310210132203-2112312311112212-1033200030012022)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-2323300301222323-2121132313300323-1122222132001022-3210020300112222-1031203223230203-2030132233201302-0321110321203202-0122133200130111)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-0303323311112102-2312200213101013-3021213031100000-1330100032322330-2030133210101101-2223133101313022-2012220020030032-1230002002303321)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-3321011303101300-1322103221313021-2312001132221121-0220013122110221-0020102231100331-1200101232013232-3203121132230030-3202112131310221)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-1121311211102200-2020311221021133-2113313203321020-3322231333113123-2022212301303310-0233310111123002-1302101001212323-1030201022123303"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
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
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
nexthop_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131310112123110-3023103122023103-0232310003030002-3231301120023100-1011112110011000-2013311101210230-3012130030323210-2003013313110232"></a>

## Direct properties — nexthop_address / 313211021123 / 3

- [dual_stack](resources--gcp_vpc_site--reference--group-004.md#canonical-3023232120032331-2310332303001330-2312313332203222-3332300312331313-3300231302113012-1232131202032000-2322101303231030-2201022300032132): complete subsection reference.

- [ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-0021120120222020-1233013323101003-0121130312020303-1200200002030223-2102002311123232-0302212320323101-1223133110103001-1033201331210232): complete subsection reference.

- [ipv6](resources--gcp_vpc_site--reference--group-004.md#canonical-0333220321002223-1330233111123032-3222002332001102-2030300113320330-3321133002102212-1220203322212331-3123222111112213-1011121001133213): complete subsection reference.

<a id="canonical-2313002102002200-0101101330213221-1133311123302303-2122032323020211-1201013123103003-2223333021200210-3102231231111010-1111233112022103"></a>

## Next pages — nexthop_address / 313211021123 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-004.md#canonical-3023232120032331-2310332303001330-2312313332203222-3332300312331313-3300231302113012-1232131202032000-2322101303231030-2201022300032132)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-0021120120222020-1233013323101003-0121130312020303-1200200002030223-2102002311123232-0302212320323101-1223133110103001-1033201331210232)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--gcp_vpc_site--reference--group-004.md#canonical-0333220321002223-1330233111123032-3222002332001102-2030300113320330-3321133002102212-1220203322212331-3123222111112213-1011121001133213)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-3321011303101300-1322103221313021-2312001132221121-0220013122110221-0020102231100331-1200101232013232-3203121132230030-3202112131310221)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3023232120032331-2310332303001330-2312313332203222-3332300312331313-3300231302113012-1232131202032000-2322101303231030-2201022300032132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130321023322300-2111322300303223-1103212000021222-1100122023220130-1023233001332210-3332111002000031-3001222122310232-1111310123110302"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — dual_stack / 113230320222 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-1130033011120230-1313001331200331-1003100020122121-0222023332233321-3112121012033033-1001310210132203-2112312311112212-1033200030012022)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-2323300301222323-2121132313300323-1122222132001022-3210020300112222-1031203223230203-2030132233201302-0321110321203202-0122133200130111)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-0303323311112102-2312200213101013-3021213031100000-1330100032322330-2030133210101101-2223133101313022-2012220020030032-1230002002303321)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-3321011303101300-1322103221313021-2312001132221121-0220013122110221-0020102231100331-1200101232013232-3203121132230030-3202112131310221)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-2000220311021312-0330123320202021-0211133102121202-3112313211122111-2232303332231022-0022011113203103-1113110300111023-0002032032200121)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-0222021023021233-1023112033000200-0002101112022331-0013322002033010-2102231032002302-2113331322211211-0022332010023133-0323201123331332"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
dual_stack {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310220101201212-0002330100230213-3320331302003202-2121331020030210-0031101012121100-2202222131221233-3010121201131012-1013230110331203"></a>

## Direct properties — dual_stack / 113230320222 / 3

- [ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-2011100112023312-3032332310303211-3211121300133112-3223113213332333-3203100312202210-0020021013210002-0230332300202222-0131303202200010): complete subsection reference.

- [ipv6](resources--gcp_vpc_site--reference--group-004.md#canonical-0231131003220132-1030212221311123-1211103120123030-3301231222133230-1020002013113020-0022003233130332-0120212333121101-3202221311213021): complete subsection reference.

<a id="canonical-1331223331001212-1312330231222102-0200101321212111-1332021003221022-2220233121033221-3333312210113311-2312103303022011-3332001010303331"></a>

## Next pages — dual_stack / 113230320222 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-2011100112023312-3032332310303211-3211121300133112-3223113213332333-3203100312202210-0020021013210002-0230332300202222-0131303202200010)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--gcp_vpc_site--reference--group-004.md#canonical-0231131003220132-1030212221311123-1211103120123030-3301231222133230-1020002013113020-0022003233130332-0120212333121101-3202221311213021)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-2000220311021312-0330123320202021-0211133102121202-3112313211122111-2232303332231022-0022011113203103-1113110300111023-0002032032200121)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2011100112023312-3032332310303211-3211121300133112-3223113213332333-3203100312202210-0020021013210002-0230332300202222-0131303202200010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301220033211110-0010300231320232-0030322310233120-2310232231131121-3210302023032123-3101303110232030-1202310110133332-0201032333030300"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 311111311212 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-1130033011120230-1313001331200331-1003100020122121-0222023332233321-3112121012033033-1001310210132203-2112312311112212-1033200030012022)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-2323300301222323-2121132313300323-1122222132001022-3210020300112222-1031203223230203-2030132233201302-0321110321203202-0122133200130111)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-0303323311112102-2312200213101013-3021213031100000-1330100032322330-2030133210101101-2223133101313022-2012220020030032-1230002002303321)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-3321011303101300-1322103221313021-2312001132221121-0220013122110221-0020102231100331-1200101232013232-3203121132230030-3202112131310221)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-2000220311021312-0330123320202021-0211133102121202-3112313211122111-2232303332231022-0022011113203103-1113110300111023-0002032032200121)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-004.md#canonical-3023232120032331-2310332303001330-2312313332203222-3332300312331313-3300231302113012-1232131202032000-2322101303231030-2201022300032132)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-3033120030123232-0002233000313233-1111222110330301-2121231333203002-0202132132233110-3110132133020222-0001213331012010-1322133321220110"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

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

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132031103301301-3331023202223113-1122031312003122-3230211001011233-0131123220201103-1111320011112331-2130312020030301-0011022222133013"></a>

## Direct properties — IPv4 / 311111311212 / 3

<a id="canonical-2200222213312102-1032003201302310-0103302233312300-0133120001122222-0223133130333120-1201302101303132-2232023330230000-3022100202032012"></a>

<a id="canonical-0230313021013000-3011002313220233-3020321220302302-1003123003212331-1323000210011013-1121330321220001-0331110312102130-2302233301211122"></a>

## addr property — IPv4 / 311111311212 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3302302103103200-3033213003302320-3103100322122330-0221302330320132-3122033123100133-0213312311321311-2131312230203123-0231311201333030"></a>

## Next pages — IPv4 / 311111311212 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-004.md#canonical-3023232120032331-2310332303001330-2312313332203222-3332300312331313-3300231302113012-1232131202032000-2322101303231030-2201022300032132)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0231131003220132-1030212221311123-1211103120123030-3301231222133230-1020002013113020-0022003233130332-0120212333121101-3202221311213021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101132230012322-2031002201220223-3030111233212313-0210032020013300-2203223033023021-0200022203132231-2300330032203232-2121212222120123"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 002230302333 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-1130033011120230-1313001331200331-1003100020122121-0222023332233321-3112121012033033-1001310210132203-2112312311112212-1033200030012022)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-2323300301222323-2121132313300323-1122222132001022-3210020300112222-1031203223230203-2030132233201302-0321110321203202-0122133200130111)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-0303323311112102-2312200213101013-3021213031100000-1330100032322330-2030133210101101-2223133101313022-2012220020030032-1230002002303321)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-3321011303101300-1322103221313021-2312001132221121-0220013122110221-0020102231100331-1200101232013232-3203121132230030-3202112131310221)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-2000220311021312-0330123320202021-0211133102121202-3112313211122111-2232303332231022-0022011113203103-1113110300111023-0002032032200121)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-004.md#canonical-3023232120032331-2310332303001330-2312313332203222-3332300312331313-3300231302113012-1232131202032000-2322101303231030-2201022300032132)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-0221133130132033-3012120122302320-2121300212022202-3211130110311120-3332121300023333-1010201230112121-1202230220110122-0231101233113130"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302331002331112-0302111022010323-3002201001020232-2103303213331030-1002210212203222-0202123030310031-0221212221102331-0230030100112133"></a>

## Direct properties — IPv6 / 002230302333 / 3

<a id="canonical-1020103130332111-0303220101020320-1333221100102232-1310212032022311-2302112222231220-2213131111012023-3310322010211211-3121122323021132"></a>

<a id="canonical-2202233320130321-0133020110213132-0201011333100321-0303030102312030-0320132300330213-0130012230130220-1313312220201303-3101132011322200"></a>

## addr property — IPv6 / 002230302333 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3221311122123232-3010032300302003-2321201313233323-0102002311000002-1001131330122132-2320010030211133-0110313223210121-1202010010132233"></a>

## Next pages — IPv6 / 002230302333 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-004.md#canonical-3023232120032331-2310332303001330-2312313332203222-3332300312331313-3300231302113012-1232131202032000-2322101303231030-2201022300032132)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0021120120222020-1233013323101003-0121130312020303-1200200002030223-2102002311123232-0302212320323101-1223133110103001-1033201331210232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021022032321233-2113023132213111-2102212332322203-0231110032032230-2031333023013020-0223112111121223-1110132310010313-1131133030020211"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4 — IPv4 / 022322030321 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-1130033011120230-1313001331200331-1003100020122121-0222023332233321-3112121012033033-1001310210132203-2112312311112212-1033200030012022)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-2323300301222323-2121132313300323-1122222132001022-3210020300112222-1031203223230203-2030132233201302-0321110321203202-0122133200130111)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-0303323311112102-2312200213101013-3021213031100000-1330100032322330-2030133210101101-2223133101313022-2012220020030032-1230002002303321)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-3321011303101300-1322103221313021-2312001132221121-0220013122110221-0020102231100331-1200101232013232-3203121132230030-3202112131310221)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-2000220311021312-0330123320202021-0211133102121202-3112313211122111-2232303332231022-0022011113203103-1113110300111023-0002032032200121)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-0101133221333000-2220313031102310-0300030222131212-1310321101330102-3101101001112212-0312111132020210-3301300133320303-2012301111212323"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

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

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311132123023321-0012221230322330-1221032113202202-2031103223231210-3012233321233220-1122133033230131-1223312030222332-2101320101120012"></a>

## Direct properties — IPv4 / 022322030321 / 3

<a id="canonical-2230230022133221-3113330230301232-1020320322230032-3003030122020212-0323201330231320-0233113033113133-3300222312132112-3001302320222310"></a>

<a id="canonical-0121312003122100-2121000323122132-1212103103223233-0220033303021311-1332332302221210-0223300333020121-1322210212002011-2131321200300120"></a>

## addr property — IPv4 / 022322030321 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0303113203121313-3121111112111020-1123220123211211-2212032330231332-1031012223032112-2232132113223321-2301322131330033-3311200233132012"></a>

## Next pages — IPv4 / 022322030321 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-2000220311021312-0330123320202021-0211133102121202-3112313211122111-2232303332231022-0022011113203103-1113110300111023-0002032032200121)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0333220321002223-1330233111123032-3222002332001102-2030300113320330-3321133002102212-1220203322212331-3123222111112213-1011121001133213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001031110102131-1300003330313100-0131031130030112-3201031320320120-3100211030030311-2023133122303231-1012133300100221-3232232302313023"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 230300211013 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-1130033011120230-1313001331200331-1003100020122121-0222023332233321-3112121012033033-1001310210132203-2112312311112212-1033200030012022)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-2323300301222323-2121132313300323-1122222132001022-3210020300112222-1031203223230203-2030132233201302-0321110321203202-0122133200130111)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-0303323311112102-2312200213101013-3021213031100000-1330100032322330-2030133210101101-2223133101313022-2012220020030032-1230002002303321)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-3321011303101300-1322103221313021-2312001132221121-0220013122110221-0020102231100331-1200101232013232-3203121132230030-3202112131310221)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-2000220311021312-0330123320202021-0211133102121202-3112313211122111-2232303332231022-0022011113203103-1113110300111023-0002032032200121)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-0122013301300302-0310321332303333-2323002030332131-3021112110130201-2131302213131210-3120202110302010-3322102112032003-3121121233120030"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201012003011002-2031320100201323-3010323210230103-1130323311210313-0211101032212310-1011331202133103-2021033301200301-2022203220000302"></a>

## Direct properties — IPv6 / 230300211013 / 3

<a id="canonical-1033321000122221-3113021022332013-3212022100232111-1032232302331123-3031232131010322-3330330123223012-3020020211123330-3121301313132023"></a>

<a id="canonical-2113332131221311-0003300212223301-2033131302002121-3022032230122222-0322102231003101-3312111320212213-1111023310123103-2011130200000213"></a>

## addr property — IPv6 / 230300211013 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3310102120201003-3002313211102132-2300213300132113-3131031322111220-2123012302113031-0233000302310131-2102211202313102-2000212110311010"></a>

## Next pages — IPv6 / 230300211013 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-2000220311021312-0330123320202021-0211133102121202-3112313211122111-2232303332231022-0022011113203103-1113110300111023-0002032032200121)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2113123120301030-1222222232103322-1321133111323110-1122013323201113-3331213001320233-2221113322120213-2033002322122323-2233300133221111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203012000200233-0303001231123011-0032032201112113-1333301220101031-1012020320011031-0033123210302223-1131213203121203-2102231131210222"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets — subnets / 021320232002 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-1130033011120230-1313001331200331-1003100020122121-0222023332233321-3112121012033033-1001310210132203-2112312311112212-1033200030012022)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-2323300301222323-2121132313300323-1122222132001022-3210020300112222-1031203223230203-2030132233201302-0321110321203202-0122133200130111)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-0303323311112102-2312200213101013-3021213031100000-1330100032322330-2030133210101101-2223133101313022-2012220020030032-1230002002303321)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-0313100012010200-0120320212213231-0233110212321331-3010312300011033-2113122200310223-3013030103203320-3122230102033122-3232030132220201"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ipv4",
    "ipv6")}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302320221121123-2132202011310213-0220323320233232-3032131231003233-1100033223101211-1322132303032210-1120000300310033-1130203301111223"></a>

## Direct properties — subnets / 021320232002 / 3

- [ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-2230212211122210-0032110221000131-2130131122312332-3333130102131201-1021110301300001-2203121022111103-1302013211301030-2312001123232212): complete subsection reference.

- [ipv6](resources--gcp_vpc_site--reference--group-004.md#canonical-0332020301301111-0003331003103200-0221011023300310-2321232112123000-3330011203022110-1023322221202020-0300203001112230-3333333212103201): complete subsection reference.

<a id="canonical-1130323133210101-2012213001303132-3211313322121033-0210302032031113-1030022321121303-2203133102313122-1013312012132333-0130231223203222"></a>

## Next pages — subnets / 021320232002 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-2230212211122210-0032110221000131-2130131122312332-3333130102131201-1021110301300001-2203121022111103-1302013211301030-2312001123232212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--gcp_vpc_site--reference--group-004.md#canonical-0332020301301111-0003331003103200-0221011023300310-2321232112123000-3330011203022110-1023322221202020-0300203001112230-3333333212103201)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-0303323311112102-2312200213101013-3021213031100000-1330100032322330-2030133210101101-2223133101313022-2012220020030032-1230002002303321)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2230212211122210-0032110221000131-2130131122312332-3333130102131201-1021110301300001-2203121022111103-1302013211301030-2312001123232212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212203312300011-2001133320031212-1200003010312113-1001023120120211-1211330220321212-3021132100313002-1213023322212012-2311123121330111"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 101102103323 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-1130033011120230-1313001331200331-1003100020122121-0222023332233321-3112121012033033-1001310210132203-2112312311112212-1033200030012022)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-2323300301222323-2121132313300323-1122222132001022-3210020300112222-1031203223230203-2030132233201302-0321110321203202-0122133200130111)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-0303323311112102-2312200213101013-3021213031100000-1330100032322330-2030133210101101-2223133101313022-2012220020030032-1230002002303321)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-004.md#canonical-2113123120301030-1222222232103322-1321133111323110-1122013323201113-3331213001320233-2221113322120213-2033002322122323-2233300133221111)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-1231210021323323-2310120233230322-1130030110123333-0012111313333030-3121212201003201-2110122022102311-3013122331302320-0320120322320102"></a>

Type: `"object"`. single nested block, Optional.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230231231121210-0321211220331013-2002330202112231-1022132111003101-2321303001210203-3203133000223233-1120210203010311-1123313000103100"></a>

## Direct properties — IPv4 / 101102103323 / 3

<a id="canonical-2132101120321210-3232111120230021-2101013112020220-0020112022000102-0022223210022022-3232031122200202-1113301302231110-0320211203302312"></a>

<a id="canonical-2102012112110112-0201332113311221-3112210301111313-3323133233312220-2012130030020202-0001230312321200-1302231223330221-0201020133110223"></a>

## plen property — IPv4 / 101102103323 / 4

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-2121102320210200-2023123101312300-1123330022021312-0222223130103311-3131121020112303-3200001100132323-1320321212211012-2220322323133203"></a>

<a id="canonical-2020011103022113-2020313230123131-1213322232120222-2231213300031301-3212002310010200-3300213201200022-2130112213313310-3013211321033230"></a>

## prefix property — IPv4 / 101102103323 / 5

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1110322330102113-1302313113001030-1211120030132321-3320231233233312-1100011133113032-3331202322311310-1313222013121122-2123000202310210"></a>

## Next pages — IPv4 / 101102103323 / 6

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-004.md#canonical-2113123120301030-1222222232103322-1321133111323110-1122013323201113-3331213001320233-2221113322120213-2033002322122323-2233300133221111)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0332020301301111-0003331003103200-0221011023300310-2321232112123000-3330011203022110-1023322221202020-0300203001112230-3333333212103201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113013130221213-3232333331222030-1123220210020303-2022321321100222-0111310302321020-3202211133030232-1333213333012132-2223033333000020"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 021213210103 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-1130033011120230-1313001331200331-1003100020122121-0222023332233321-3112121012033033-1001310210132203-2112312311112212-1033200030012022)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-2323300301222323-2121132313300323-1122222132001022-3210020300112222-1031203223230203-2030132233201302-0321110321203202-0122133200130111)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-0303323311112102-2312200213101013-3021213031100000-1330100032322330-2030133210101101-2223133101313022-2012220020030032-1230002002303321)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-004.md#canonical-2113123120301030-1222222232103322-1321133111323110-1122013323201113-3331213001320233-2221113322120213-2033002322122323-2233300133221111)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-1123013323112120-1101200003032122-3121200200213203-1221232303201320-2200111323101223-3033121220313102-2302021123130330-2030203112233220"></a>

Type: `"object"`. single nested block, Optional.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130322023122130-0010122331210033-3213302220131102-2023020023220200-1100320232203030-1223231210301130-2122230003200121-1022020211132321"></a>

## Direct properties — IPv6 / 021213210103 / 3

<a id="canonical-3120211330030011-2102113122220312-0100332133230112-0222133321103130-1001023101200113-2222102023111210-2303213032223312-0333003303021020"></a>

<a id="canonical-2203203231021302-2033312011030010-1103123120223220-1111123322210011-1310031200303331-2333313031030130-0012210120133233-1030130122002111"></a>

## plen property — IPv6 / 021213210103 / 4

Type: `"number"`. Optional.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-3111223132123130-2213001312220122-3331031333133221-3312330032132310-2132021231032132-2333000112021200-3002132122102302-0313222010310203"></a>

<a id="canonical-1303120010233221-2333113131303333-3123031011202233-1130030333013132-1202002123011133-3212302113030220-3320203113023101-1123311003003001"></a>

## prefix property — IPv6 / 021213210103 / 5

Type: `"string"`. Optional.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1130222300032010-0302223203023010-3201330331122323-1012103021132321-1133122133200013-1002030311230313-1323231102333303-2133130310331301"></a>

## Next pages — IPv6 / 021213210103 / 6

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-004.md#canonical-2113123120301030-1222222232103322-1321133111323110-1122013323201113-3331213001320233-2221113322120213-2033002322122323-2233300133221111)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0031010113331203-0332021211122002-2133232221022033-1323212302220010-1332220012121111-1233120030320131-0003300032001033-1013330221332200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323322131333022-3002302010010003-3112232101311320-1011210210330210-0031100211331122-2032023130002102-2233113011230113-2102331022212010"></a>

## voltstack_cluster.site_local_network — site_local_network / 011231233022 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- voltstack_cluster.site_local_network

<a id="canonical-3312112033010321-3313133200100220-1312113103010000-3012203331320010-3033332300011210-1303010331223000-3310003200132201-0210013203202012"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_network",
    "new_network"),
  validators.ConflictingObjectAttributes("existing_network",
    "new_network_autogenerate"),
  validators.ConflictingObjectAttributes("new_network",
    "new_network_autogenerate")}
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
  "x-ves-oneof-field-choice": "[\"existing_network\",\"new_network\",\"new_network_autogenerate\"]"
}
```

Terraform syntax:

```terraform
site_local_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-2222222311023021-1232032003012233-0121013110121301-2130111203221311-1312011032321103-3200332130211231-0002130310113001-3033233223310123"></a>

## Direct properties — site_local_network / 011231233022 / 3

- [existing_network](resources--gcp_vpc_site--reference--group-004.md#canonical-2102113313210211-3133210000013230-2123313231112103-3102201330223133-3231233130001120-0330020131132100-2233111003112321-1111233210302302): complete subsection reference.

- [new_network](resources--gcp_vpc_site--reference--group-004.md#canonical-3021221131001010-3121130032103200-0003110021232323-3312000312131323-1003231013331212-3121120031100301-3222303223300301-0010013201322123): complete subsection reference.

- [new_network_autogenerate](resources--gcp_vpc_site--reference--group-005.md#canonical-2221021031302100-2310231033112111-2333333110213200-3212311123133201-2011221231020333-3123122102322211-1000330301223202-0320111111311131): complete subsection reference.

<a id="canonical-1101132333201012-1001123301001021-0230131230313320-0202333020100112-3022302120111020-2201001312031132-0033031212010012-0023203200021332"></a>

## Next pages — site_local_network / 011231233022 / 4

- [voltstack_cluster.site_local_network.existing_network](resources--gcp_vpc_site--reference--group-004.md#canonical-2102113313210211-3133210000013230-2123313231112103-3102201330223133-3231233130001120-0330020131132100-2233111003112321-1111233210302302)
- [voltstack_cluster.site_local_network.new_network](resources--gcp_vpc_site--reference--group-004.md#canonical-3021221131001010-3121130032103200-0003110021232323-3312000312131323-1003231013331212-3121120031100301-3222303223300301-0010013201322123)
- [voltstack_cluster.site_local_network.new_network_autogenerate](resources--gcp_vpc_site--reference--group-005.md#canonical-2221021031302100-2310231033112111-2333333110213200-3212311123133201-2011221231020333-3123122102322211-1000330301223202-0320111111311131)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2102113313210211-3133210000013230-2123313231112103-3102201330223133-3231233130001120-0330020131132100-2233111003112321-1111233210302302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122300030023203-0021232100022210-2202323023103213-0301032121113233-1330200213102002-0031003031002121-0302032122130330-2011000201200120"></a>

## voltstack_cluster.site_local_network.existing_network — existing_network / 332333012113 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113)
- [voltstack_cluster.site_local_network](resources--gcp_vpc_site--reference--group-004.md#canonical-0031010113331203-0332021211122002-2133232221022033-1323212302220010-1332220012121111-1233120030320131-0003300032001033-1013330221332200)
- voltstack_cluster.site_local_network.existing_network

<a id="canonical-1303001131232032-3302112102322133-3311132233000001-2031232010012022-1213202312232100-2030312101210130-0200311212032100-0331102203020002"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for existing network.

Upstream description:

Name of existing VPC network.

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
  },
  "x-ves-oneof-field-routing_type": "[]"
}
```

Terraform syntax:

```terraform
existing_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-0123122222122200-1130320000001323-1032330312220111-1323232201103211-2323102230121103-2230023302122312-2300121203100021-1013201131023121"></a>

## Direct properties — existing_network / 332333012113 / 3

<a id="canonical-3010302131010323-0123310022311122-3311203101221123-3312320022202003-1131313203233000-0132013233121323-0002031021200122-3113020311130112"></a>

<a id="canonical-0110210310100300-1213033220000013-1113122313023000-1322211002201212-1323002221023213-1313131311201021-2021111302223030-3332222210002133"></a>

## name property — existing_network / 332333012113 / 4

Type: `"string"`. Optional.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2322102021102131-0131211020120002-1113023132310230-2122102333110031-1102021302022221-3131113032011000-0123002020310330-0211003130023032"></a>

## Next pages — existing_network / 332333012113 / 5

- [voltstack_cluster.site_local_network](resources--gcp_vpc_site--reference--group-004.md#canonical-0031010113331203-0332021211122002-2133232221022033-1323212302220010-1332220012121111-1233120030320131-0003300032001033-1013330221332200)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3021221131001010-3121130032103200-0003110021232323-3312000312131323-1003231013331212-3121120031100301-3222303223300301-0010013201322123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
