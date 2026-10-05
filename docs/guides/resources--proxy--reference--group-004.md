---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-0021111020232021-1212000123123133-2031010112220123-0013211021013220-2311331021131003-1312101030312020-2111030013100033-2333000322332112"></a>

## trusted_ca_url property — use_mtls / 003012323023 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](resources--proxy--reference--group-004.md#canonical-3301301013003012-1010130000202121-0230031303003100-2111102033131312-0200320003300232-1300031313132231-1320223300221001-0012021223131322): complete subsection reference.

- [xfcc_options](resources--proxy--reference--group-004.md#canonical-0012033023300202-1223030021113200-1023321133332103-1233011321201233-1231122211102010-1031000133330202-2021121032011301-0233203202103022): complete subsection reference.

<a id="canonical-0313313022132003-1000110300132132-2000111023120001-2201230131000021-1103213010212213-1000101220221113-3331321310012013-2132201303220012"></a>

## Next pages — use_mtls / 003012323023 / 6

- [dynamic_proxy.https_proxy.tls_params.use_mtls.crl](resources--proxy--reference--group-004.md#canonical-0001322300020012-2202212022120010-1131232201213100-0000301003323111-1123330313013320-1110121300013301-1020000312223323-2200211201330030)
- [dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl](resources--proxy--reference--group-004.md#canonical-2212123203301223-2303123302332301-1033020013321211-0212311111123033-2100000200020300-1120330133022321-2030023121323133-1302003230302120)
- [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca](resources--proxy--reference--group-004.md#canonical-2230303001100132-2100001122031031-1021322122310121-2330022311303113-3230213033230130-2302031221001030-2120331111023010-0120103221232032)
- [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled](resources--proxy--reference--group-004.md#canonical-3301301013003012-1010130000202121-0230031303003100-2111102033131312-0200320003300232-1300031313132231-1320223300221001-0012021223131322)
- [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options](resources--proxy--reference--group-004.md#canonical-0012033023300202-1223030021113200-1023321133332103-1233011321201233-1231122211102010-1031000133330202-2021121032011301-0233203202103022)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0001322300020012-2202212022120010-1131232201213100-0000301003323111-1123330313013320-1110121300013301-1020000312223323-2200211201330030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203121220212033-1003133133033211-1122310110013020-3301012312011200-0233303103020300-3302122212310321-1002321022312001-1201132012231222"></a>

## dynamic_proxy.https_proxy.tls_params.use_mtls.crl — crl / 223021233103 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-1120302120201130-2021020112211312-2012121003331033-1202131121023223-2102330022132231-0123033302022200-3110300012022001-3113320030210032)
- dynamic_proxy.https_proxy.tls_params.use_mtls.crl

<a id="canonical-3000233230301000-1310120333101011-3222232112301300-3330010330030100-0122231311200121-0323121320102013-1003111131323002-1121200321020130"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-0311012012130330-2230202122110332-2202333010311222-3330213233010221-2003312111232333-0300010300321332-0311130302102223-0002130211300011"></a>

## Direct properties — crl / 223021233103 / 3

<a id="canonical-2100012213311002-3112111331232000-1232231331102003-1332310010303213-1313112232112310-2102322101210121-1203110302320212-1123312222032112"></a>

<a id="canonical-1311310223232113-3230232121203202-0001120301332113-2311320200032022-3233002123313001-2013212012302002-1232210332201102-2020013022012020"></a>

## name property — crl / 223021233103 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2011013033332322-2012013022331210-2232130323320033-0030312310210131-3021133003101032-0003013231221030-0122010103103023-0130330003111213"></a>

<a id="canonical-1332120030103310-2203120110202311-1111212231201222-0031020010023102-3100022003100322-0303211331213301-0120111330102010-0023020100131122"></a>

## namespace property — crl / 223021233103 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2301200221321002-2213201030300011-0230112203033202-2310320032022121-0210332320310012-0233223032223002-1331133331312313-0330332122330022"></a>

<a id="canonical-2122100002113021-2332300212332120-3331202110031331-2003301231222321-0232332230032231-3313330220300030-0333232032032121-1232300122103212"></a>

## tenant property — crl / 223021233103 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0220301331301222-2312330300331030-1230300100011323-1032131222120320-1301231012303212-0232333333022230-0300010332032203-0230023221302021"></a>

## Next pages — crl / 223021233103 / 7

- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-1120302120201130-2021020112211312-2012121003331033-1202131121023223-2102330022132231-0123033302022200-3110300012022001-3113320030210032)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2212123203301223-2303123302332301-1033020013321211-0212311111123033-2100000200020300-1120330133022321-2030023121323133-1302003230302120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101232110012103-3133112210303133-2330223333330230-3021001022223010-3011323203022322-2322123023231320-0133111331033303-2311322010123231"></a>

## dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl — no_crl / 313013133021 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-1120302120201130-2021020112211312-2012121003331033-1202131121023223-2102330022132231-0123033302022200-3110300012022001-3113320030210032)
- dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl

<a id="canonical-0012332021113300-0101132122322010-1233111203122131-3233321011020222-2220213101211133-2022013323301130-0203100030230012-3020202233320013"></a>

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
no_crl = {}
```

<a id="canonical-2013012132123110-0323311301022302-2223021330122310-1333313232111333-2321330313202133-1100311232110021-3333213310133130-1101033032330220"></a>

## Direct properties — no_crl / 313013133021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122300202311332-3022000131211011-3221311211030230-2301230121012333-2112020312323312-2203100131022200-0312112022321121-2222103000220002"></a>

## Next pages — no_crl / 313013133021 / 4

- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-1120302120201130-2021020112211312-2012121003331033-1202131121023223-2102330022132231-0123033302022200-3110300012022001-3113320030210032)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2230303001100132-2100001122031031-1021322122310121-2330022311303113-3230213033230130-2302031221001030-2120331111023010-0120103221232032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320310330001130-2320203231303210-1112312210300222-3220200111322032-2331303113202220-1313220212132012-2022333011230200-1100123033030031"></a>

## dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca — trusted_ca / 220012100331 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-1120302120201130-2021020112211312-2012121003331033-1202131121023223-2102330022132231-0123033302022200-3110300012022001-3113320030210032)
- dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca

<a id="canonical-2101133330131000-1210300310201203-2120330333320222-3220313320012132-0033101031132320-3020302310233202-0333113201311211-3011310001320203"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-0130113313312100-2203232112222130-0022020311003302-1130132110221003-0330113300300223-3233221233331233-0103202310031303-1121220000020223"></a>

## Direct properties — trusted_ca / 220012100331 / 3

<a id="canonical-1113223012213120-1312200010332033-1303302220020303-1303320023202333-1333132130311333-0310321130220013-3200303122210012-3102331130023212"></a>

<a id="canonical-1033303213212302-3001133322030112-3202322002000202-3011101000113221-1300012033302303-2113101100311122-3211021302011323-1212000300232211"></a>

## name property — trusted_ca / 220012100331 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2020320002332323-3121111023011020-0120202321222220-3022130211300131-3311302020132212-3122311021011222-3321311230212123-3022203003201100"></a>

<a id="canonical-0231220101220100-1322033230202211-1210121202122322-1222323001011223-2100131032301023-0011302032032201-1102123133120202-3223330121102023"></a>

## namespace property — trusted_ca / 220012100331 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2131203132123300-1331311113102133-0030321222032120-0010212302033301-1002230022222033-3230221211331332-3013223131101321-1030321000111212"></a>

<a id="canonical-1312301333231020-1312120101121300-2001000010011023-1223122133332131-2211302331031020-2203230203033300-3002133311212332-3213321120312303"></a>

## tenant property — trusted_ca / 220012100331 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0123310330312231-3012120000102221-1021310022120220-1310101301331003-0301310031121223-0133133011323120-0203301220131221-3230023303122313"></a>

## Next pages — trusted_ca / 220012100331 / 7

- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-1120302120201130-2021020112211312-2012121003331033-1202131121023223-2102330022132231-0123033302022200-3110300012022001-3113320030210032)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3301301013003012-1010130000202121-0230031303003100-2111102033131312-0200320003300232-1300031313132231-1320223300221001-0012021223131322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131322120201123-2030011301022033-1313000303000301-3101331311020311-3111331023221023-3302221232032003-3311013311033330-0330003130330313"></a>

## dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled — xfcc_disabled / 032012300310 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-1120302120201130-2021020112211312-2012121003331033-1202131121023223-2102330022132231-0123033302022200-3110300012022001-3113320030210032)
- dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled

<a id="canonical-3103231300321003-2332320233132101-0112303220110113-0013222023313233-1022010303023100-2222310231010132-1113121132121211-1333311111230012"></a>

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
xfcc_disabled = {}
```

<a id="canonical-0221012122132333-0133003303320333-3303202231310010-2223000200302013-3013032131302203-1323320132221110-0002111031302002-0032321102322121"></a>

## Direct properties — xfcc_disabled / 032012300310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302311120030300-2220223300123010-1332321320233212-1333010012200311-3302033212011122-3023320033312102-3011003302230211-1011232212201102"></a>

## Next pages — xfcc_disabled / 032012300310 / 4

- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-1120302120201130-2021020112211312-2012121003331033-1202131121023223-2102330022132231-0123033302022200-3110300012022001-3113320030210032)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0012033023300202-1223030021113200-1023321133332103-1233011321201233-1231122211102010-1031000133330202-2021121032011301-0233203202103022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100331303212022-1310333031113300-0200213002132020-3121223011212310-1223312033311212-3002203303320103-3222020232022222-1302000120031012"></a>

## dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options — xfcc_options / 123003002201 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-1120302120201130-2021020112211312-2012121003331033-1202131121023223-2102330022132231-0123033302022200-3110300012022001-3113320030210032)
- dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options

<a id="canonical-3211002013213020-0021331212032032-2110133202312122-3130330121020202-1020020322133333-1022130000232201-0301301131200330-2203002331301331"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0010013220101031-1230312122300121-2301121332122023-0100211302323132-3100320230201212-3220131312101233-1012333233313002-0110100011002021"></a>

## Direct properties — xfcc_options / 123003002201 / 3

<a id="canonical-0303132300303032-3323122220133003-2113000112303033-2123202220313310-0231133312333303-0123223021330130-2301310210230033-1003331303013100"></a>

<a id="canonical-0021120113221333-0003011330131330-0321100213120001-3132111020223100-0213100102110002-2112332333332300-1213301231020303-3323020021221130"></a>

## xfcc_header_elements property — xfcc_options / 123003002201 / 4

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-3130023110002022-1310231232130121-1121130122020320-2102323212113112-2022321010012203-3222223230113230-0113120330032030-1121232131111122"></a>

## Next pages — xfcc_options / 123003002201 / 5

- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-1120302120201130-2021020112211312-2012121003331033-1202131121023223-2102330022132231-0123033302022200-3110300012022001-3113320030210032)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3112020001021313-3001013121113231-2312300311233300-0101002320313231-2321103212111010-1200232211133121-0211133000013012-3022312230220132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011320011030003-0120120130103220-0210332213012331-1211210210312212-3300302131200120-3323230212311202-0001212132003011-3213211121300123"></a>

## dynamic_proxy.sni_proxy — sni_proxy / 201013300211 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- dynamic_proxy.sni_proxy

<a id="canonical-1200030100100302-3313101203023312-3310213031113310-3011223200220232-2103312102230333-1313221202010332-0111200212103102-0112023213123033"></a>

Type: `"object"`. single nested block, Optional.

Dynamic SNI Proxy Type. Parameters for dynamic SNI proxy.

Upstream description:

Parameters for dynamic SNI proxy.

Receipt-pinned upstream constraints:

```json
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
sni_proxy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0012101233223000-2212222202020311-3220312031121120-3222212101113330-1213222203323233-3023220230332030-1032210111020113-3020003120131213"></a>

## Direct properties — sni_proxy / 201013300211 / 3

<a id="canonical-3102213023212303-0330313003231333-3320111333210230-1103302013332103-1112333310132031-2121023231131130-0232232223121321-3301102012232102"></a>

<a id="canonical-2132022310222133-2030120303301213-3000212200001033-2022231200313113-0211100211302203-3313133021012112-3110001121300102-0100220010313212"></a>

## idle_timeout property — sni_proxy / 201013300211 / 4

Type: `"number"`. Optional.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(86400000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400000,
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
    "ves.io.schema.rules.uint32.lte": "86400000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400000"
  }
}
```

<a id="canonical-2121010230010022-3201032133020122-3100201113320012-2132213031101032-2120131032200110-3021221322002031-2223203102301012-1122331310030002"></a>

## Next pages — sni_proxy / 201013300211 / 5

- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001130321111012-0123330210031333-1200230203212123-0200333323001333-0030320110230222-2120111032223232-3221232200322222-2120313320211032"></a>

## http_proxy — http_proxy / 201330323120 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- http_proxy

<a id="canonical-0023030222010311-3300322230002013-2001102311323223-1122022322221322-2020120013210033-0032330013323132-3222112033103212-3223122333232132"></a>

Type: `"object"`. single nested block, Optional.

HTTP Connect Proxy. Parameters for HTTP Connect Proxy.

Upstream description:

Parameters for HTTP Connect Proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_https_choice": "[\"enable_http\"]"
}
```

Terraform syntax:

```terraform
http_proxy {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320022103003102-1331311000200033-1101320102202223-3330120211110100-3022010032133022-2213111311002320-0013002030303031-3113221021303111"></a>

## Direct properties — http_proxy / 201330323120 / 3

- [enable_http](resources--proxy--reference--group-004.md#canonical-2002121210003201-0021203110203212-0231123222023331-1322322121312213-2112302100301200-1311323303120011-0332131100230302-1321302012212010): complete subsection reference.

- [more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003): complete subsection reference.

<a id="canonical-1303111221122222-2233103211331120-0332212303213133-1220012330030033-2220011103103210-1210033301123301-0100110303101100-1132101333332001"></a>

## Next pages — http_proxy / 201330323120 / 4

- [http_proxy.enable_http](resources--proxy--reference--group-004.md#canonical-2002121210003201-0021203110203212-0231123222023331-1322322121312213-2112302100301200-1311323303120011-0332131100230302-1321302012212010)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2002121210003201-0021203110203212-0231123222023331-1322322121312213-2112302100301200-1311323303120011-0332131100230302-1321302012212010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020131101310211-0312331113122221-0033321030110330-2220311300321323-1330022112012210-2202030001113103-1023210112030231-2312333233012031"></a>

## http_proxy.enable_http — enable_http / 221312132032 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- http_proxy.enable_http

<a id="canonical-2200301101203123-0013221023302022-1232013120121021-3011201200132231-1230013223110322-3123022010130320-0222232030103110-2221202002230222"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for enable http.

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
enable_http {}
```

<a id="canonical-0211302223200122-0113031210132311-2231310003320213-3300233033203332-0310232311021010-1300232011033132-1320001220102120-1223220300031100"></a>

## Direct properties — enable_http / 221312132032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223301302300123-3103301330211222-3003203303213130-0113302332131121-1102030122230013-2323203111110233-3220022320312101-0121210103213333"></a>

## Next pages — enable_http / 221312132032 / 4

- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020232023301010-0233323002131131-2120010131020320-3120003132102001-3332222110312330-3022300300110012-1213123031021000-1101021210002230"></a>

## http_proxy.more_option — more_option / 312100232331 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- http_proxy.more_option

<a id="canonical-2221001223210112-1100323203103002-1221130002010130-3330211303220210-3321303213313233-3133102003321312-1003210223123213-2133113122313131"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to define a route.

Upstream description:

This defines various OPTIONS to define a route.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("max_requests_per_connection",
    "no_request_limit_per_connection")}
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
  "x-ves-oneof-field-max_requests_per_connection_choice": "[\"max_requests_per_connection\",\"no_request_limit_per_connection\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-strict_sni_host_header_check_choice": "[]"
}
```

Terraform syntax:

```terraform
more_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-0311123313000323-3032100313330001-0132233323100122-3333122000010220-3211001000312030-3211220003330100-3012103002332101-0201013130220301"></a>

## Direct properties — more_option / 312100232331 / 3

- [buffer_policy](resources--proxy--reference--group-004.md#canonical-0300111113000322-0301122303123033-2331031131312202-2232013311303303-2130110102131033-3332231021130301-2213011301023020-3333033311323303): complete subsection reference.

- [compression_params](resources--proxy--reference--group-004.md#canonical-3320023111221201-1200202123231322-0232100223203033-2300211013131021-2111033013110111-1031222122321032-3331001111222321-1223223332010113): complete subsection reference.

<a id="canonical-0313022031120112-0330201322330102-2202030133331203-0310301212223021-0011100232010303-2121001310020313-3100201212011233-0003031130000313"></a>

<a id="canonical-2312122322202212-2010001030012110-1122201233312313-0031210302303223-2013320113122023-2332212032310331-1122122020320112-3110130113203311"></a>

## custom_errors property — more_option / 312100232331 / 4

Type: `["map", "string"]`. Optional.

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx..

Upstream description:

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response
code class 5 -- for 5xx response code class Value of the map is string which represents custom HTTP
responses. Specific response code takes preference when both response code and response code class
matches for a request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"ranges\":[[3,3],[4,4],[5,5],[300,599]],\"type\":\"uint32-string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.uint32.ranges\":\"3,4,5,300-599\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"65536\",\"ves.io.schema.rules.map.values.string.uri_ref\":\"true\"},\"values\":{\"format\":\"uri-reference\",\"maxLength\":65536,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "ranges": [
        [
          3,
          3
        ],
        [
          4,
          4
        ],
        [
          5,
          5
        ],
        [
          300,
          599
        ]
      ],
      "type": "uint32-string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "65536",
      "ves.io.schema.rules.map.values.string.uri_ref": "true"
    },
    "values": {
      "format": "uri-reference",
      "maxLength": 65536,
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
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  }
}
```

<a id="canonical-3231320000303110-1012233021103302-2102122331023303-0231021323111200-3100110000332321-3223322102131121-3110020330303030-3210220023021330"></a>

<a id="canonical-1002022303133302-2201102200333221-0033233313010001-3133312131103010-1101302103030303-3102001313123021-3212230121123203-0302003301000003"></a>

## disable_default_error_pages property — more_option / 312100232331 / 5

Type: `"bool"`. Optional.

Disable the use of default F5XC error pages.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [disable_path_normalize](resources--proxy--reference--group-004.md#canonical-0233013301333121-0003002030332233-0221001332302133-0113132323023221-0313010200131210-2102233002103312-3131300033223113-1102011222311112): complete subsection reference.

- [enable_path_normalize](resources--proxy--reference--group-004.md#canonical-2301321023223101-2012123133321310-3120133122312330-3322110201211220-2120110022201321-0230303210303012-3111320112022320-2300212111030323): complete subsection reference.

<a id="canonical-0112113002130301-2002113131330321-0132330230130300-3133311110212013-3200102133232212-3103213231203200-2121130310333101-3332133011113012"></a>

<a id="canonical-3202302321000332-0330030232033300-2033031013330210-1002323221110200-2332331330221103-3130301112203130-3103323101033120-2301321313212121"></a>

## idle_timeout property — more_option / 312100232331 / 6

Type: `"number"`. Optional.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Upstream description:

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(3600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3600000,
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
    "ves.io.schema.rules.uint32.lte": "3600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "3600000"
  }
}
```

<a id="canonical-1310120112213001-2201330230223221-0020012310332112-2012013230030010-3312203223100101-3030122030323330-3130333110003021-1321322201333201"></a>

<a id="canonical-0021022200103000-0233111202012332-2213032011231112-2322223003233322-3321231002133300-0301200213202320-1122011133022022-1233123103003013"></a>

## max_request_header_size property — more_option / 312100232331 / 7

Type: `"number"`. Optional.

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size. If multiple load balancers
share the same advertise\_policy, the highest value configured across all such load balancers is
used..

Upstream description:

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size.

If multiple load balancers share the same advertise\_policy, the highest value configured across all
such load balancers is used for all the load balancers in question.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(96),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 96,
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
    "ves.io.schema.rules.uint32.lte": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "96"
  }
}
```

<a id="canonical-1002130211130201-2101032201222021-1120321121103302-1300220103010130-3100122213330011-2133111200121022-1121310331111122-2202110311000302"></a>

<a id="canonical-3320300021320033-0223230010301033-3320102220233211-1320313110112313-1023233220221132-3302320121133213-0220302022112301-0032312022303332"></a>

## max_requests_per_connection property — more_option / 312100232331 / 8

Type: `"number"`. Optional.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [no_request_limit_per_connection](resources--proxy--reference--group-004.md#canonical-1003122322301321-1333110123011112-3221222213203212-1023220213321033-2023002003031031-3010120133222233-1001030333011203-3301303002212231): complete subsection reference.

- [request_cookies_to_add](resources--proxy--reference--group-004.md#canonical-1311033121201011-3223001021303303-1330100313013302-1300333130322122-3121002120320021-0010212303011003-1032011123031100-3230213212313312): complete subsection reference.

<a id="canonical-0120232230030030-2321001331333320-2021302031113212-3132332013031220-3332323000313301-1212230310331123-3013320321310111-3312000112021333"></a>

<a id="canonical-2020210032101331-3030222131020211-1032112201130312-3123203022221213-2112302000322013-0133112320131332-0333202030333303-3122011203330103"></a>

## request_cookies_to_remove property — more_option / 312100232331 / 9

Type: `["list", "string"]`. Optional.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](resources--proxy--reference--group-004.md#canonical-2310002312312021-3131313220231323-3103000310102012-1232012231303310-1022012331001022-1220023212200233-2312330020030233-3132033101221100): complete subsection reference.

<a id="canonical-0123101101130310-1130221002002033-2220012002122330-3223202111300120-0333333200033323-3003032122220102-0221101233303313-0210333023012213"></a>

<a id="canonical-1131203301130313-0002201130103010-3123331213303001-3213011003223223-0300110123131032-0131221210130032-1133123231000301-1222103110023003"></a>

## request_headers_to_remove property — more_option / 312100232331 / 10

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311): complete subsection reference.

<a id="canonical-2100130031301331-1211022003131011-2331122102111202-1033032130122033-0002322110223220-1221031322133213-0311213110022020-3030322110233211"></a>

<a id="canonical-2113033321113331-1331200212230331-2223220222331033-3231330232301212-3121130100331100-1113213310023031-1011113133213203-3033320100213102"></a>

## response_cookies_to_remove property — more_option / 312100232331 / 11

Type: `["list", "string"]`. Optional.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](resources--proxy--reference--group-004.md#canonical-0321302212120332-1003310003320101-2331211322200330-1231022123102030-0001322133103022-1202202120331013-0312112103131001-3233331200001032): complete subsection reference.

<a id="canonical-2002112100110131-0103300002133100-0010332111230100-1233301232121321-2232103211230233-3002313103012231-1011131310023122-0323031120030323"></a>

<a id="canonical-0332021333211102-1011133233002231-1033111210123222-0221303313113330-2103011330312213-3232333321310023-0212200100203111-0223202321100000"></a>

## response_headers_to_remove property — more_option / 312100232331 / 12

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0310211000032021-3003022032120233-2222212111123113-2211022110130123-1301300201021233-0110113033030333-2110220113012123-1232012321331231"></a>

## Next pages — more_option / 312100232331 / 13

- [http_proxy.more_option.buffer_policy](resources--proxy--reference--group-004.md#canonical-0300111113000322-0301122303123033-2331031131312202-2232013311303303-2130110102131033-3332231021130301-2213011301023020-3333033311323303)
- [http_proxy.more_option.compression_params](resources--proxy--reference--group-004.md#canonical-3320023111221201-1200202123231322-0232100223203033-2300211013131021-2111033013110111-1031222122321032-3331001111222321-1223223332010113)
- [http_proxy.more_option.disable_path_normalize](resources--proxy--reference--group-004.md#canonical-0233013301333121-0003002030332233-0221001332302133-0113132323023221-0313010200131210-2102233002103312-3131300033223113-1102011222311112)
- [http_proxy.more_option.enable_path_normalize](resources--proxy--reference--group-004.md#canonical-2301321023223101-2012123133321310-3120133122312330-3322110201211220-2120110022201321-0230303210303012-3111320112022320-2300212111030323)
- [http_proxy.more_option.no_request_limit_per_connection](resources--proxy--reference--group-004.md#canonical-1003122322301321-1333110123011112-3221222213203212-1023220213321033-2023002003031031-3010120133222233-1001030333011203-3301303002212231)
- [http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-004.md#canonical-1311033121201011-3223001021303303-1330100313013302-1300333130322122-3121002120320021-0010212303011003-1032011123031100-3230213212313312)
- [http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-004.md#canonical-2310002312312021-3131313220231323-3103000310102012-1232012231303310-1022012331001022-1220023212200233-2312330020030233-3132033101221100)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-004.md#canonical-0321302212120332-1003310003320101-2331211322200330-1231022123102030-0001322133103022-1202202120331013-0312112103131001-3233331200001032)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0300111113000322-0301122303123033-2331031131312202-2232013311303303-2130110102131033-3332231021130301-2213011301023020-3333033311323303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112303211310331-3220110013112010-1100220323203020-1330203003033100-3321313222013032-2231002130103211-1101002321022120-1310021013320012"></a>

## http_proxy.more_option.buffer_policy — buffer_policy / 312011323302 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- http_proxy.more_option.buffer_policy

<a id="canonical-2030331223113202-1301222121300032-1333023200302230-2223210310123030-2103212211031020-3020122131311010-3212011120111103-0222130023203220"></a>

Type: `"object"`. single nested block, Optional.

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Upstream description:

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer
configuration takes precedence.

Receipt-pinned upstream constraints:

```json
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
buffer_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2100101132131302-2030303133100103-0111200131013131-2101112003223213-2210203313231331-0031211031000103-1112333301212103-0103022212322332"></a>

## Direct properties — buffer_policy / 312011323302 / 3

<a id="canonical-3030213210301303-3333030201220321-0230101011202023-2003022131230110-2101302003303110-3001120103222012-3333331030333213-3200333203112001"></a>

<a id="canonical-1123303312220200-3123121223130131-1023220012323230-3001230000213323-0220220313233210-0201123033330103-1203030120112010-2220231313033313"></a>

## disabled property — buffer_policy / 312011323302 / 4

Type: `"bool"`. Optional.

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3022321331221031-3101331230333301-0021011112311301-0221312013023013-2132031333301013-2331310300322220-0031213220302100-1133310220022010"></a>

<a id="canonical-0010023110113001-3013012312223112-2031210330002323-3302012321213211-2231003322203310-2331210320013102-1131311323121022-2221011021021312"></a>

## max_request_bytes property — buffer_policy / 312011323302 / 5

Type: `"number"`. Optional.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Upstream description:

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(10485760),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
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
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

<a id="canonical-2300131303012132-0200030330010213-1310200130321231-3313221110113002-3300100211312223-0222111300132002-2023230033120103-0030102131333322"></a>

## Next pages — buffer_policy / 312011323302 / 6

- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3320023111221201-1200202123231322-0232100223203033-2300211013131021-2111033013110111-1031222122321032-3331001111222321-1223223332010113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313313113021002-0003002123021332-2011123130032132-0231313031022302-2130210103110220-3010121001301031-2321311001011333-3200132331111111"></a>

## http_proxy.more_option.compression_params — compression_params / 130211311000 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- http_proxy.more_option.compression_params

<a id="canonical-0320033203010212-1211021031113132-2000000320222000-3110300022022233-1020031200101233-0020321320102303-1032023200311321-1231210233003101"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to compress dispatched data from an upstream service upon client request. The
content is compressed and then sent to the client with the appropriate headers if either response
and request allow. Only GZIP compression is supported.

Upstream description:

Enables loadbalancer to compress dispatched data from an upstream service upon client request. The
content is compressed and then sent to the client with the appropriate headers if either response
and request allow. Only GZIP compression is supported.

By default compression will be skipped when:

A request does NOT contain accept-encoding header. A request includes accept-encoding header, but it
does not contain “gzip” or “\*”. A request includes accept-encoding with “gzip” or “\*” with the
weight “q=0”. Note that the “gzip” will have a higher weight then “\*”. For example, if
accept-encoding is “gzip;q=0,\*;q=1”, the filter will not compress. But if the header is set to
“\*;q=0,gzip;q=1”, the filter will compress. A request whose accept-encoding header includes
“identity”. A response contains a content-encoding header. A response contains a cache-control
header whose value includes “no-transform”. A response contains a transfer-encoding header whose
value includes “gzip”. A response does not contain a content-type value that matches one of the
selected mime-types, which default to application/JavaScript, application/JSON,
application/xhtml+XML, image/svg+XML, text/CSS, text/HTML, text/plain, text/XML. Neither
content-length nor transfer-encoding headers are present in the response. Response size is smaller
than 30 bytes (only applicable when transfer-encoding is not chunked).

When compression is applied:

The content-length is removed from response headers. Response headers contain “transfer-encoding:
chunked” and do not contain “content-encoding” header. The “vary: accept-encoding” header is
inserted on every response.

GZIP Compression Level:

A value which is optimal balance between speed of compression and amount of compression is chosen.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("content_length")}
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
compression_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003132031320301-0122101322220022-2211313013002133-3100120331320301-0110310203301210-1102302203210200-0020102102111201-2320010322103031"></a>

## Direct properties — compression_params / 130211311000 / 3

<a id="canonical-2231331100321200-3031101132033022-3311333302312010-0302120312200032-0303300133131332-1202231322023211-0320033210230112-1230101110200303"></a>

<a id="canonical-2002000113021303-2033023101000000-3112300011221101-0231221211001101-2213311121223321-1220112003310202-3330112332031002-2331320112300023"></a>

## content_length property — compression_params / 130211311000 / 4

Type: `"number"`. Optional.

Minimum response length, in bytes, which will trigger compression. The. Defaults to \`30\`.

Upstream description:

Minimum response length, in bytes, which will trigger compression. The default value is 30.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtLeast(30),
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
    "minimum": 30
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "30"
  }
}
```

<a id="canonical-0132030320221310-0320321031021113-0212032212130010-1101103000110030-0001123021221321-3003332333033103-3001333000010201-1132113220310322"></a>

<a id="canonical-3220033112112102-1231102320303002-2012103322133020-3133133213002120-1101301001102010-0312003233001203-0201322323113001-3100001221333232"></a>

## content_type property — compression_params / 130211311000 / 5

Type: `["list", "string"]`. Optional.

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: 'application/JavaScript'
'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain'
'text/XML'.

Upstream description:

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: "application/JavaScript"
"application/JSON", "application/xhtml+XML" "image/svg+XML" "text/CSS" "text/HTML" "text/plain"
"text/XML"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(50),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
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
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1011211133321111-1303000003000102-1123222321310302-1200221131213310-3120023031123200-0120200020000312-3031232212120311-3212222213031032"></a>

<a id="canonical-0300022322332112-0212002100123331-0220233130210221-3030000213312103-1032131322111222-0030223213012101-0313131003111323-3010102000023113"></a>

## disable_on_etag_header property — compression_params / 130211311000 / 6

Type: `"bool"`. Optional.

If true, disables compression when the response contains an etag header. When it is false, weak
etags will be preserved and the ones that require strong validation will be removed.

Upstream description:

If true, disables compression when the response contains an etag header. When it is false, weak
etags will be preserved and the ones that require strong validation will be removed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2132023021032303-3122203000311123-0313311221113132-1120321210130112-1200101120313313-1232323222000222-0321031221010231-2230333022231131"></a>

<a id="canonical-0123110231110220-1212003222220320-3113002110132213-2023201102321031-0110232001131033-1222232233310223-0000321232030130-2300010313230211"></a>

## remove_accept_encoding_header property — compression_params / 130211311000 / 7

Type: `"bool"`. Optional.

If true, removes accept-encoding from the request headers before dispatching it to the upstream so
that responses do not GET compressed before reaching the filter.

Upstream description:

If true, removes accept-encoding from the request headers before dispatching it to the upstream so
that responses do not GET compressed before reaching the filter.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2103000001222131-0203322232132300-1332233121210221-2311102200301322-3022323113001222-1201021122331300-0002012003232003-1131130301320101"></a>

## Next pages — compression_params / 130211311000 / 8

- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0233013301333121-0003002030332233-0221001332302133-0113132323023221-0313010200131210-2102233002103312-3131300033223113-1102011222311112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313212001013032-3110110330102213-2313032022200003-2111013312103310-3200110222122101-0100232000200121-0101201213001311-0322121220233101"></a>

## http_proxy.more_option.disable_path_normalize — disable_path_normalize / 320001023013 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- http_proxy.more_option.disable_path_normalize

<a id="canonical-0112221233100331-0113012122102211-3011200303303011-1031032220303131-1100323202123211-2233030223201213-1331122233020302-2000333322130033"></a>

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
disable_path_normalize = {}
```

<a id="canonical-3311300213201320-2121333320312303-0201032100213313-3200100130203012-3213302121030320-0003132112231112-1303203201332230-2231100230111212"></a>

## Direct properties — disable_path_normalize / 320001023013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000320123020033-2230223112021130-2013132013202230-2013301210122320-0201110333230322-3201110101121322-3112303031003221-0333100233122321"></a>

## Next pages — disable_path_normalize / 320001023013 / 4

- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2301321023223101-2012123133321310-3120133122312330-3322110201211220-2120110022201321-0230303210303012-3111320112022320-2300212111030323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330332230221213-2113301033013201-0222320112301323-1200332111311012-1211031112331032-3321120203210033-3120021131120132-3020332220301223"></a>

## http_proxy.more_option.enable_path_normalize — enable_path_normalize / 311013132111 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- http_proxy.more_option.enable_path_normalize

<a id="canonical-0203101203312302-2312202110311310-2220302001001003-0320023221131300-3001112322010233-1212200111231300-3222203231201223-2301302213220111"></a>

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
enable_path_normalize = {}
```

<a id="canonical-0033220330201101-0311002303211002-2123001001033031-2132220100302120-0301032012221301-3102001001010323-1331210320211100-0110300030031113"></a>

## Direct properties — enable_path_normalize / 311013132111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232111230033100-1113221033121313-3121133032021200-1123001211302130-0303331131003121-3022102110313100-1032222102230032-2103031012110031"></a>

## Next pages — enable_path_normalize / 311013132111 / 4

- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1003122322301321-1333110123011112-3221222213203212-1023220213321033-2023002003031031-3010120133222233-1001030333011203-3301303002212231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201000220021220-1012211001323203-3211222031132312-2102230312211031-2003221212102102-3000330102101102-1111110032110321-2333203122011120"></a>

## http_proxy.more_option.no_request_limit_per_connection — no_request_limit_per_connection / 002032333010 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- http_proxy.more_option.no_request_limit_per_connection

<a id="canonical-2013003000303112-2201123200223002-3133201302102000-3033010231020010-0210033222213323-0301020130232020-0101322331132030-1301132223330003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no request limit per connection.

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
no_request_limit_per_connection = {}
```

<a id="canonical-2211011020000031-0220131102111300-2323233331033010-1320112323131002-0331010333001110-0001320003023002-2120210020021323-2212101310123130"></a>

## Direct properties — no_request_limit_per_connection / 002032333010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300031303300303-3012310020001031-0310221110333101-1100012021311221-3030202312330133-2023122211321233-1032322112023120-2111132221111011"></a>

## Next pages — no_request_limit_per_connection / 002032333010 / 4

- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1311033121201011-3223001021303303-1330100313013302-1300333130322122-3121002120320021-0010212303011003-1032011123031100-3230213212313312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011112113222330-1332331123121201-0330121302230122-3330201023122220-2302322030223300-3300333213220112-3330000313130233-0102022110213233"></a>

## http_proxy.more_option.request_cookies_to_add — request_cookies_to_add / 232210322112 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- http_proxy.more_option.request_cookies_to_add

<a id="canonical-1210000331112233-3112003330210012-0230201313032113-0123320223110231-2301003312222303-3032133310110022-2302212333333231-3033003101121103"></a>

Type: `"object"`. list nested block, Optional.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Upstream description:

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
request_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313302220122213-3130321100102202-3301133033231231-2032200031210112-3203311111131231-2023001101003233-1330030330031202-0203102121223330"></a>

## Direct properties — request_cookies_to_add / 232210322112 / 3

<a id="canonical-3102202110110102-0012310312222312-3333030331110311-3302131230122302-0211312030300201-3313231310101332-0313231013210101-1211101233123310"></a>

<a id="canonical-2333322323123331-0033320032221210-3123203023032201-1200322021202030-3220321100220220-3200300222011023-2203232313233200-2321023113020120"></a>

## name property — request_cookies_to_add / 232210322112 / 4

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0203322032102113-3202110212331020-2201301320303113-3203320100333101-2121131123132332-0221202001110130-0222030111200311-3113022000101310"></a>

<a id="canonical-3332302000220113-3013213012020312-2122022121221112-3003122112212120-0213001301212330-3130320001101210-3032013022210203-0003323012011220"></a>

## overwrite property — request_cookies_to_add / 232210322112 / 5

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [secret_value](resources--proxy--reference--group-004.md#canonical-3010332010112330-1233131123133131-3000223203333231-2330012221023020-2212031000021203-0123020021231110-0201332132110201-0113212212212331): complete subsection reference.

<a id="canonical-2111203323003033-0030313313112322-2201132003320111-3201033321221330-3023101110133332-1002300320131011-0210301200013301-2002113232202012"></a>

<a id="canonical-0100223122100232-1133200020010021-0320030132233332-0302202122112232-1202220331033321-1130010312313030-1102112122233311-0021012000031301"></a>

## value property — request_cookies_to_add / 232210322112 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-1103301102002123-1313221012112333-0312002133133333-1010310023130013-1300220221102023-3113320012110203-0031011232133030-1021303100320302"></a>

## Next pages — request_cookies_to_add / 232210322112 / 7

- [http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-3010332010112330-1233131123133131-3000223203333231-2330012221023020-2212031000021203-0123020021231110-0201332132110201-0113212212212331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3010332010112330-1233131123133131-3000223203333231-2330012221023020-2212031000021203-0123020021231110-0201332132110201-0113212212212331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321331231312333-0233323000302113-2220312222320332-3201033131120313-0132321330021001-0100303130112303-2123001233011102-1210310012022022"></a>

## http_proxy.more_option.request_cookies_to_add.secret_value — secret_value / 203211103002 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-004.md#canonical-1311033121201011-3223001021303303-1330100313013302-1300333130322122-3121002120320021-0010212303011003-1032011123031100-3230213212313312)
- http_proxy.more_option.request_cookies_to_add.secret_value

<a id="canonical-0032302311112022-3220210031202120-2100222202123221-3110333132333112-1110031030233223-3233100312232031-0032130020101300-1000002031133020"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320231333303302-2111103012332232-2133230013003013-0130132123223033-0332231311132013-0132001311200300-1001330202101020-0230110013203230"></a>

## Direct properties — secret_value / 203211103002 / 3

- [blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-2300202030130121-3302312131303030-1320112132301301-0023333100030123-0123113131313031-1220202201322003-3022213331231121-3122002202132222): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-004.md#canonical-3032322112310300-0330132331133331-3112033012013211-0003331023313132-0320021013213232-1012022031023230-2033220310022330-2322301222330131): complete subsection reference.

<a id="canonical-2001322312223222-2311123103033220-1203310203212002-0021221011030120-0022312022113130-1133312100233032-0203300301222221-3133230113321111"></a>

## Next pages — secret_value / 203211103002 / 4

- [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-2300202030130121-3302312131303030-1320112132301301-0023333100030123-0123113131313031-1220202201322003-3022213331231121-3122002202132222)
- [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-004.md#canonical-3032322112310300-0330132331133331-3112033012013211-0003331023313132-0320021013213232-1012022031023230-2033220310022330-2322301222330131)
- [http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-004.md#canonical-1311033121201011-3223001021303303-1330100313013302-1300333130322122-3121002120320021-0010212303011003-1032011123031100-3230213212313312)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2300202030130121-3302312131303030-1320112132301301-0023333100030123-0123113131313031-1220202201322003-3022213331231121-3122002202132222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111013003010020-2113232111103323-3230332102023022-3323220113000201-2111012331313130-1311313301113013-1121321231111230-3230113232221112"></a>

## http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 002212330003 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-004.md#canonical-1311033121201011-3223001021303303-1330100313013302-1300333130322122-3121002120320021-0010212303011003-1032011123031100-3230213212313312)
- [http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-3010332010112330-1233131123133131-3000223203333231-2330012221023020-2212031000021203-0123020021231110-0201332132110201-0113212212212331)
- http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-1010010223100223-0203100100122031-2300201230110212-0332010211032022-2111001132323003-2122122113130210-1112021023210010-1331200210011110"></a>

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

<a id="canonical-3110233012302110-2011322012220311-2231020212312231-3333030333330030-1223130102011023-0003310233230331-2122130223122112-2221120332220200"></a>

## Direct properties — blindfold_secret_info / 002212330003 / 3

<a id="canonical-3331331233222231-0021031113232222-2111002202000012-2112002322323321-2320000303231120-2323112312132033-3112210201332202-1001011312302103"></a>

<a id="canonical-3021131013123112-2000002023201033-2301313110100331-2011213213130313-2130322011312330-2100331220003211-0131021121201321-2310321311030302"></a>

## decryption_provider property — blindfold_secret_info / 002212330003 / 4

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

<a id="canonical-1010032322301332-3211332133013000-1303332212023012-3320331120001233-1003020302001033-1221120001200103-1000012133013212-0112011030212002"></a>

<a id="canonical-2232011330102313-1200333031000103-0000210030021320-0231331211032021-1331100122220003-3312110220102220-1133133302200002-3301210023302100"></a>

## location property — blindfold_secret_info / 002212330003 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3310200010310031-0330112202113002-3203031210003220-0233121203200023-1013132132323131-0323001231331010-1200332010123130-1010000210300011"></a>

<a id="canonical-2303202322031202-3112221120000112-3330301311112010-0120013220223001-2011033202131100-3023123321201233-0321102313211030-3202230102121122"></a>

## store_provider property — blindfold_secret_info / 002212330003 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-2010301103123201-2303230131230322-1323002112123121-2011322131123323-2101233131302222-0231002213203301-3303322323101030-1322002012202221"></a>

## Next pages — blindfold_secret_info / 002212330003 / 7

- [http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-3010332010112330-1233131123133131-3000223203333231-2330012221023020-2212031000021203-0123020021231110-0201332132110201-0113212212212331)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3032322112310300-0330132331133331-3112033012013211-0003331023313132-0320021013213232-1012022031023230-2033220310022330-2322301222330131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130131220121130-1300322333013130-2303331130201033-2333120231331032-0020132302302131-3312022231112120-0121202000132301-0333012013213103"></a>

## http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info — clear_secret_info / 202313100132 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-004.md#canonical-1311033121201011-3223001021303303-1330100313013302-1300333130322122-3121002120320021-0010212303011003-1032011123031100-3230213212313312)
- [http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-3010332010112330-1233131123133131-3000223203333231-2330012221023020-2212031000021203-0123020021231110-0201332132110201-0113212212212331)
- http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-3021131101030110-2220002212111122-2111220200213300-1201102330123030-0112213023022100-3232202333202031-3211222223310122-2013100302113113"></a>

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

<a id="canonical-2302303002302000-0200311232030103-0112322200130130-1213013033302232-2023300302333110-2233012213201003-3332033033011122-2100000112233022"></a>

## Direct properties — clear_secret_info / 202313100132 / 3

<a id="canonical-3220111300202111-1331000010032330-3032021200011123-0303133110102211-1200123231213011-0012011202002103-1333302102320303-1313321003132013"></a>

<a id="canonical-1310101322013101-1312030222223011-1031232210213030-0210033231103102-0011103300320032-0333111213131031-3020033011032033-3032022012201223"></a>

## provider_ref property — clear_secret_info / 202313100132 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2211233212012220-2333132133213302-0321220200333222-3302300010320200-2133103300132103-1222023032333011-0330311130002031-3322022202002202"></a>

<a id="canonical-1112323100313223-0022220132112010-0213130123301221-1323113112300022-2321101021033131-2322133122012002-3311311223012210-2121321131103110"></a>

## URL property — clear_secret_info / 202313100132 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0132001100300021-3322113300210321-1200112230113103-3313031303122033-3101223010021230-2022303023120331-1232301020133303-0221001211122303"></a>

## Next pages — clear_secret_info / 202313100132 / 6

- [http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-3010332010112330-1233131123133131-3000223203333231-2330012221023020-2212031000021203-0123020021231110-0201332132110201-0113212212212331)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2310002312312021-3131313220231323-3103000310102012-1232012231303310-1022012331001022-1220023212200233-2312330020030233-3132033101221100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202100030023123-2202311110013300-2232220211203113-2002302310220230-2220331331323120-2032002000020233-3120102122210013-3000022231323102"></a>

## http_proxy.more_option.request_headers_to_add — request_headers_to_add / 332322123100 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- http_proxy.more_option.request_headers_to_add

<a id="canonical-1230012112232311-2322032002321033-2203021003032312-3332130230101301-3110110320203223-1023021203021101-0233010213021222-1302103230000110"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
request_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000230013332011-1202100302130211-3000233003331313-1000221231112132-2223021010003300-3002331010130223-2303131323011321-0113102300312122"></a>

## Direct properties — request_headers_to_add / 332322123100 / 3

<a id="canonical-3013220321331202-1030322313102032-2203300013232001-0112212112332122-0223101002302202-2221313003312030-1302223201031223-0212323003301231"></a>

<a id="canonical-2223011012010320-1133212003101321-2103221011321030-2021212312231022-3300301203100203-0323133021033332-3121312020001322-2331020133311303"></a>

## append property — request_headers_to_add / 332322123100 / 4

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1203012110003233-1113331323000333-3130231333031022-0102322122211300-2020331011200310-0333301012303131-2120333330313030-3102321001231111"></a>

<a id="canonical-2120301102322320-2312010331123201-3113031113221332-3103012223213122-0331323022030331-3323110110113131-0303333312130301-1302132222211213"></a>

## name property — request_headers_to_add / 332322123100 / 5

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](resources--proxy--reference--group-004.md#canonical-1221020210303123-2323332201130110-1022223200200331-3021132112233032-3232210111232213-0313113330212103-3321130301022111-3333302221133320): complete subsection reference.

<a id="canonical-2003232011312130-3313211102303323-0310323312310323-0230112133100213-2033323010231020-2131132121133023-0331010232001333-3111100201002021"></a>

<a id="canonical-3200203312113232-2012112002132203-2220203021220233-2213033030332210-2031323211203011-1011010210132003-0011320130320103-0223030031022021"></a>

## value property — request_headers_to_add / 332322123100 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-3012132000112123-2003300330003011-2202110232322221-3103331323000012-2231112003322211-0331233113030021-3311213001332020-2001112000022120"></a>

## Next pages — request_headers_to_add / 332322123100 / 7

- [http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-1221020210303123-2323332201130110-1022223200200331-3021132112233032-3232210111232213-0313113330212103-3321130301022111-3333302221133320)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1221020210303123-2323332201130110-1022223200200331-3021132112233032-3232210111232213-0313113330212103-3321130301022111-3333302221133320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211331033303320-3000102002002020-2101100012232311-3031301030000203-1202021223230111-2010101121322212-1223001103133232-3220110001113212"></a>

## http_proxy.more_option.request_headers_to_add.secret_value — secret_value / 210101101110 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-004.md#canonical-2310002312312021-3131313220231323-3103000310102012-1232012231303310-1022012331001022-1220023212200233-2312330020030233-3132033101221100)
- http_proxy.more_option.request_headers_to_add.secret_value

<a id="canonical-2300231120000222-2210203002022322-3020322312201133-2232223001000320-1202201021113203-2021320210113032-2130000233333201-3232300313103313"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-1302220311320030-3313023000322321-1220012201020022-0000230200112230-2212133221123222-2002123123313112-3001122010322132-1300221010120302"></a>

## Direct properties — secret_value / 210101101110 / 3

- [blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-2110202002210301-0033222003032113-0131101202201201-3313202031333333-2312133023131102-1210110303231232-2333321320233002-3130112013012302): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-004.md#canonical-1200031203213032-3231313111013212-0203302011121021-2310312111230312-0121113312123120-1010000110101111-2132311300210013-2002030213131330): complete subsection reference.

<a id="canonical-0322012122103322-1220220203223010-1033222222003011-0011020113202111-1220000101131223-3020333222101332-1012222133100303-1200333121022300"></a>

## Next pages — secret_value / 210101101110 / 4

- [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-2110202002210301-0033222003032113-0131101202201201-3313202031333333-2312133023131102-1210110303231232-2333321320233002-3130112013012302)
- [http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-004.md#canonical-1200031203213032-3231313111013212-0203302011121021-2310312111230312-0121113312123120-1010000110101111-2132311300210013-2002030213131330)
- [http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-004.md#canonical-2310002312312021-3131313220231323-3103000310102012-1232012231303310-1022012331001022-1220023212200233-2312330020030233-3132033101221100)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2110202002210301-0033222003032113-0131101202201201-3313202031333333-2312133023131102-1210110303231232-2333321320233002-3130112013012302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023112213210312-2322131321101301-0131033213103311-3323323133213300-2023332032120101-3310022033333322-1301112300023021-3102131013332022"></a>

## http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 111221220113 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-004.md#canonical-2310002312312021-3131313220231323-3103000310102012-1232012231303310-1022012331001022-1220023212200233-2312330020030233-3132033101221100)
- [http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-1221020210303123-2323332201130110-1022223200200331-3021132112233032-3232210111232213-0313113330212103-3321130301022111-3333302221133320)
- http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-1031322130332031-2313202200012232-3312032121110312-1213002022212332-1112332000003313-1031100100223332-0201222221111003-2320233232230333"></a>

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

<a id="canonical-0310121222113312-2000100121232323-1010102302103103-3201232033200303-2301303322122323-0221021203032201-0203201132002003-1210313320131310"></a>

## Direct properties — blindfold_secret_info / 111221220113 / 3

<a id="canonical-3113013301200113-2021031312203012-0022322200302100-0020312111223002-1031120031200211-1232020302301121-2012302332031111-2220030203211113"></a>

<a id="canonical-0100323202221330-0310321123002011-2201001213000122-0232300111013321-3002121200003021-2022111110300002-0131222100103321-2020333113033310"></a>

## decryption_provider property — blindfold_secret_info / 111221220113 / 4

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

<a id="canonical-3231111013321022-1000303002003201-2100312113112020-0211213311002030-3330223130233302-3231112232331213-3222331202001300-2120311122110320"></a>

<a id="canonical-0301030212002002-3301303323332102-2020212130203222-2123100120003011-1012303013000222-3213101120132031-1031020123211132-2121202023120212"></a>

## location property — blindfold_secret_info / 111221220113 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2013303323012021-2131220223212323-3120221321111313-2313020222111130-0221003320201211-1032000122202230-1230312022310210-3022000233320020"></a>

<a id="canonical-0303202333021200-2321002303212203-1320000011003131-1220132122303201-0210300132322011-0310013212122113-1101213233101003-1200001023010302"></a>

## store_provider property — blindfold_secret_info / 111221220113 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-3021222201331201-3030011133113333-1323010121121112-0031202330001231-3021001020130301-0230023122301003-3102001213031332-0100120120321331"></a>

## Next pages — blindfold_secret_info / 111221220113 / 7

- [http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-1221020210303123-2323332201130110-1022223200200331-3021132112233032-3232210111232213-0313113330212103-3321130301022111-3333302221133320)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1200031203213032-3231313111013212-0203302011121021-2310312111230312-0121113312123120-1010000110101111-2132311300210013-2002030213131330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311030031322230-2232311322200231-2031223201320000-2323113300032112-1113030203010320-3103113033333122-3010103201200330-2201231202113301"></a>

## http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 232230223030 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-004.md#canonical-2310002312312021-3131313220231323-3103000310102012-1232012231303310-1022012331001022-1220023212200233-2312330020030233-3132033101221100)
- [http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-1221020210303123-2323332201130110-1022223200200331-3021132112233032-3232210111232213-0313113330212103-3321130301022111-3333302221133320)
- http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-2311311000202303-3023321133211030-0112113010311230-1332312133230231-2130111202112000-3123031200113022-3122200313203010-2001000200300322"></a>

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

<a id="canonical-2222301312213231-0330111301021312-3100012121103132-2130013113230313-2002111311102213-0303223233023231-0212003313130313-0021330311221012"></a>

## Direct properties — clear_secret_info / 232230223030 / 3

<a id="canonical-0012113100031233-3332000210230320-2303202321030130-1223201322011211-1330121030311232-1222103331130032-2313211013211210-2020021013101101"></a>

<a id="canonical-2110111011030011-0120322101012133-2013223300300332-1230133210310302-1231221031200121-0001010213100103-0323102301021232-2311213332031311"></a>

## provider_ref property — clear_secret_info / 232230223030 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0310201103103212-1101311330010203-3332331312231012-0013322012220211-0213232212022301-0322200213103303-0003330120312123-1010132020302110"></a>

<a id="canonical-2221001030123201-2100323311130030-2012010120331133-0332212213030111-2110333023110211-3323232233023311-1203330121222120-3032211132032201"></a>

## URL property — clear_secret_info / 232230223030 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2020201123310331-1101230231300132-2201230013233213-1020332131130011-3130102233032001-2223111133310300-1221012131122021-1201301111330233"></a>

## Next pages — clear_secret_info / 232230223030 / 6

- [http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-1221020210303123-2323332201130110-1022223200200331-3021132112233032-3232210111232213-0313113330212103-3321130301022111-3333302221133320)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201011313210012-1131320231100330-3223323301200010-0322121330030130-3220131311101012-0010023021212021-1332321223002212-1121003320302330"></a>

## http_proxy.more_option.response_cookies_to_add — response_cookies_to_add / 030203020330 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- http_proxy.more_option.response_cookies_to_add

<a id="canonical-1110113233031013-0303232013113321-1111012232011023-0231212322222302-2311002212132112-3102031331333331-1020030303221002-3023103221320021"></a>

Type: `"object"`. list nested block, Optional.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Upstream description:

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("add_domain",
    "ignore_domain"),
  validators.ConflictingListObjectAttributes("add_expiry",
    "ignore_expiry"),
  validators.ConflictingListObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingListObjectAttributes("add_partitioned",
    "ignore_partitioned"),
  validators.ConflictingListObjectAttributes("add_path",
    "ignore_path"),
  validators.ConflictingListObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingListObjectAttributes("ignore_max_age",
    "max_age_value"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "secret_value"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "value"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_none",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
response_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-1213023112022122-2110212011101003-2121033303323021-2320201031201221-3112031211102322-3023231000221203-1323223311132031-1310130212200033"></a>

## Direct properties — response_cookies_to_add / 030203020330 / 3

<a id="canonical-3013001312222112-1301311032110131-3120013310333030-2121033123320011-2023201201220213-3320201321221022-1113122020132030-2022003033110113"></a>

<a id="canonical-2111233303213332-0210013031010102-2001021110330333-1123013121002321-1123332323330213-0101121330231121-2001203202101200-0030331033212212"></a>

## add_domain property — response_cookies_to_add / 030203020330 / 4

Type: `"string"`. Optional.

Exclusive with \[ignore\_domain\] Add domain attribute.

Upstream description:

Exclusive with \[ignore\_domain\] Add domain attribute.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3232101222113000-0201002030012012-2131211312020101-0123201313133123-0130311313032102-2230001130101220-3122310131012033-3013210000333232"></a>

<a id="canonical-0323302121323131-3332203201220321-0210022201123310-1212112311223332-0031221021112022-1122220323022000-1130322112313112-3233331030030121"></a>

## add_expiry property — response_cookies_to_add / 030203020330 / 5

Type: `"string"`. Optional.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Upstream description:

Exclusive with \[ignore\_expiry\] Add expiry attribute.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_httponly](resources--proxy--reference--group-004.md#canonical-0322013202332010-2030231103122231-3231303232203303-2032203302113112-2031030000332031-3132330213131202-0013233332130021-1131221202111120): complete subsection reference.

- [add_partitioned](resources--proxy--reference--group-004.md#canonical-3120333300200020-3220011211122220-0210232103101022-0111313302133333-2033110320012003-3203013202211011-3003222221201211-1031032310222321): complete subsection reference.

<a id="canonical-0023311121131322-3210023031020001-1011201132313100-0312212132010010-2123131221122001-3120323223130033-1332310300323022-2000232301231002"></a>

<a id="canonical-2213321020012310-2202320131201032-0110123231300121-1303321113310021-1010033232011100-3000100332301231-3201323012303131-3300031012002121"></a>

## add_path property — response_cookies_to_add / 030203020330 / 6

Type: `"string"`. Optional.

Exclusive with \[ignore\_path\] Add path attribute.

Upstream description:

Exclusive with \[ignore\_path\] Add path attribute.

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

- [add_secure](resources--proxy--reference--group-004.md#canonical-1120113031102033-3311223300132202-0003111130112133-1031302321113210-2001313331103213-3022322310000020-3321122301112321-3112111301012012): complete subsection reference.

- [ignore_domain](resources--proxy--reference--group-004.md#canonical-0230111310231220-0022230232320033-2001112110322132-0021320021301300-3012120001021120-3231132113231133-2000102033301312-0220221221123323): complete subsection reference.

- [ignore_expiry](resources--proxy--reference--group-004.md#canonical-2033222002030110-3233202113213202-2113331203022231-0320121110012031-2332232223012113-0203021121211312-3210003103210011-1032211210022333): complete subsection reference.

- [ignore_httponly](resources--proxy--reference--group-004.md#canonical-0320001110232130-3312320130031211-2122113003123303-2212001322310122-3201312232320101-2302221233320111-1023331312003220-0013033130300121): complete subsection reference.

- [ignore_max_age](resources--proxy--reference--group-004.md#canonical-3020333013311201-2021311030003000-1033033230122330-2011120333313213-0003132311113332-1033032010010130-3210101132310300-1100233001021330): complete subsection reference.

- [ignore_partitioned](resources--proxy--reference--group-004.md#canonical-0123111011200201-0302223113101031-3212002000133333-0111011212322021-1103212133121322-0212013322333033-0230213123202330-0121300001003122): complete subsection reference.

- [ignore_path](resources--proxy--reference--group-004.md#canonical-0321323213312032-0011221302101101-3201100322103203-2132110001122333-1103210003202103-2321100301321001-3300312221201033-3322021330302231): complete subsection reference.

- [ignore_samesite](resources--proxy--reference--group-004.md#canonical-2111200001020121-3110202000221230-0023303100300132-1122030032011310-1021120313102101-2102302300232131-3333011132032220-3212120331330203): complete subsection reference.

- [ignore_secure](resources--proxy--reference--group-004.md#canonical-0032120130021302-0201013033000220-0202332223110330-2031013200331123-1100100021332110-3133030202102031-2213210331220022-3030222301221230): complete subsection reference.

- [ignore_value](resources--proxy--reference--group-004.md#canonical-1031222023100232-3133320301231330-0312312111112012-1010303031233111-3003200000002033-0020112101021130-1221203331012312-1120103113021320): complete subsection reference.

<a id="canonical-2310032332322001-3220200203003332-1022010100011112-3220231322321131-1212130200012222-1000130121302203-3332313100123320-1002111212310323"></a>

<a id="canonical-3132223110202331-0203123120003020-3000231313122320-2312120121001200-0020313231301310-0312303112021122-3311233022213022-1110201100022000"></a>

## max_age_value property — response_cookies_to_add / 030203020330 / 7

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(34560000),
}
```

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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-2032122312200131-2111332013212203-1021331031300313-2313311211103332-3302230322232122-3022012230002121-0331330113000212-1233110320112012"></a>

<a id="canonical-2230021032221100-2121022330332122-2333103101100003-1110110020231110-3311203033302323-2032221100320320-1121130203232021-3021122032303033"></a>

## name property — response_cookies_to_add / 030203020330 / 8

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3123302332311112-0231212123321032-3300030101300110-0210112033110321-1130233333233013-3210230323301131-3320313300313120-2012330330210030"></a>

<a id="canonical-0220102322010200-2312123220022223-1000121011121231-2021111103102231-3312003221333312-2133212101333322-1313023102111123-2322300321202222"></a>

## overwrite property — response_cookies_to_add / 030203020330 / 9

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [samesite_lax](resources--proxy--reference--group-004.md#canonical-1212103011120333-0223001021231133-1321303222211111-2100302322210312-0221333311012000-3223203131002122-2332032200121331-0211031130111321): complete subsection reference.

- [samesite_none](resources--proxy--reference--group-004.md#canonical-3132332212330322-0202310233313222-2003022020022302-3113230213300121-0200202011212113-1020310103120000-1323101010233002-3312222223212020): complete subsection reference.

- [samesite_strict](resources--proxy--reference--group-004.md#canonical-2222032212220010-0332200103321102-2131011023001102-3310033321132213-0103322202210010-0220312231211301-3310330222200123-3320210000213321): complete subsection reference.

- [secret_value](resources--proxy--reference--group-004.md#canonical-0032223220331321-2122103333322322-0112302321320233-1222120311123120-3122320100100323-2000101331022231-0320013223311032-0110033301203222): complete subsection reference.

<a id="canonical-2313220302022013-3032113230332112-3111022103330310-1100300002332301-1102033220111211-1020122231321022-0100032223213231-0133023300113320"></a>

<a id="canonical-1212103321002022-3131300111102321-1302211302210033-3220313200030020-1031331122213021-2011211110312123-2303333313032202-3300133033210131"></a>

## value property — response_cookies_to_add / 030203020330 / 10

Type: `"string"`. Optional.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-0213201332322001-0100100023003233-3011123212312201-0210311323122100-3303320323000121-3323220023230320-1130300010002333-1131110132033330"></a>

## Next pages — response_cookies_to_add / 030203020330 / 11

- [http_proxy.more_option.response_cookies_to_add.add_httponly](resources--proxy--reference--group-004.md#canonical-0322013202332010-2030231103122231-3231303232203303-2032203302113112-2031030000332031-3132330213131202-0013233332130021-1131221202111120)
- [http_proxy.more_option.response_cookies_to_add.add_partitioned](resources--proxy--reference--group-004.md#canonical-3120333300200020-3220011211122220-0210232103101022-0111313302133333-2033110320012003-3203013202211011-3003222221201211-1031032310222321)
- [http_proxy.more_option.response_cookies_to_add.add_secure](resources--proxy--reference--group-004.md#canonical-1120113031102033-3311223300132202-0003111130112133-1031302321113210-2001313331103213-3022322310000020-3321122301112321-3112111301012012)
- [http_proxy.more_option.response_cookies_to_add.ignore_domain](resources--proxy--reference--group-004.md#canonical-0230111310231220-0022230232320033-2001112110322132-0021320021301300-3012120001021120-3231132113231133-2000102033301312-0220221221123323)
- [http_proxy.more_option.response_cookies_to_add.ignore_expiry](resources--proxy--reference--group-004.md#canonical-2033222002030110-3233202113213202-2113331203022231-0320121110012031-2332232223012113-0203021121211312-3210003103210011-1032211210022333)
- [http_proxy.more_option.response_cookies_to_add.ignore_httponly](resources--proxy--reference--group-004.md#canonical-0320001110232130-3312320130031211-2122113003123303-2212001322310122-3201312232320101-2302221233320111-1023331312003220-0013033130300121)
- [http_proxy.more_option.response_cookies_to_add.ignore_max_age](resources--proxy--reference--group-004.md#canonical-3020333013311201-2021311030003000-1033033230122330-2011120333313213-0003132311113332-1033032010010130-3210101132310300-1100233001021330)
- [http_proxy.more_option.response_cookies_to_add.ignore_partitioned](resources--proxy--reference--group-004.md#canonical-0123111011200201-0302223113101031-3212002000133333-0111011212322021-1103212133121322-0212013322333033-0230213123202330-0121300001003122)
- [http_proxy.more_option.response_cookies_to_add.ignore_path](resources--proxy--reference--group-004.md#canonical-0321323213312032-0011221302101101-3201100322103203-2132110001122333-1103210003202103-2321100301321001-3300312221201033-3322021330302231)
- [http_proxy.more_option.response_cookies_to_add.ignore_samesite](resources--proxy--reference--group-004.md#canonical-2111200001020121-3110202000221230-0023303100300132-1122030032011310-1021120313102101-2102302300232131-3333011132032220-3212120331330203)
- [http_proxy.more_option.response_cookies_to_add.ignore_secure](resources--proxy--reference--group-004.md#canonical-0032120130021302-0201013033000220-0202332223110330-2031013200331123-1100100021332110-3133030202102031-2213210331220022-3030222301221230)
- [http_proxy.more_option.response_cookies_to_add.ignore_value](resources--proxy--reference--group-004.md#canonical-1031222023100232-3133320301231330-0312312111112012-1010303031233111-3003200000002033-0020112101021130-1221203331012312-1120103113021320)
- [http_proxy.more_option.response_cookies_to_add.samesite_lax](resources--proxy--reference--group-004.md#canonical-1212103011120333-0223001021231133-1321303222211111-2100302322210312-0221333311012000-3223203131002122-2332032200121331-0211031130111321)
- [http_proxy.more_option.response_cookies_to_add.samesite_none](resources--proxy--reference--group-004.md#canonical-3132332212330322-0202310233313222-2003022020022302-3113230213300121-0200202011212113-1020310103120000-1323101010233002-3312222223212020)
- [http_proxy.more_option.response_cookies_to_add.samesite_strict](resources--proxy--reference--group-004.md#canonical-2222032212220010-0332200103321102-2131011023001102-3310033321132213-0103322202210010-0220312231211301-3310330222200123-3320210000213321)
- [http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-0032223220331321-2122103333322322-0112302321320233-1222120311123120-3122320100100323-2000101331022231-0320013223311032-0110033301203222)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0322013202332010-2030231103122231-3231303232203303-2032203302113112-2031030000332031-3132330213131202-0013233332130021-1131221202111120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203313232210103-2300212320002313-1302133020112001-1321133013200010-3120321003002022-3231303113222120-3230322333000202-2011123033322312"></a>

## http_proxy.more_option.response_cookies_to_add.add_httponly — add_httponly / 200330200011 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.add_httponly

<a id="canonical-2300221132220032-1212303011022123-0102011133303233-2212221022123222-0003001022000033-2000130100210121-1022013010231130-2213221221301330"></a>

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

<a id="canonical-0100033022033000-3300330023300033-3221323111233223-0303033133211113-3122130112322321-2301110232231020-2322131113303331-3232211311322023"></a>

## Direct properties — add_httponly / 200330200011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322322312103320-3103103021001322-2010101322003133-1331110110122211-2201122303232021-2000132200132300-2221130102003013-3112232103301013"></a>

## Next pages — add_httponly / 200330200011 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3120333300200020-3220011211122220-0210232103101022-0111313302133333-2033110320012003-3203013202211011-3003222221201211-1031032310222321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133331123233031-1100033220232323-1210330330110133-2312200230013201-1030333101010333-0132023232020113-1301121022131232-3221313001211302"></a>

## http_proxy.more_option.response_cookies_to_add.add_partitioned — add_partitioned / 310322332112 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.add_partitioned

<a id="canonical-0000220000131020-0112002310222230-1312332010332001-3220230022310100-2332132203133132-1012330313310301-0311322023311130-0303301332131221"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add partitioned.

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
add_partitioned = {}
```

<a id="canonical-3300311210030332-0203203223333000-3211132112011201-1222011011011013-1011022320012321-1230332131100112-2312331323102201-3113232132133021"></a>

## Direct properties — add_partitioned / 310322332112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111113011001220-2201320002020122-2303020103101111-3000223131200200-2021210231200011-3221011212101210-2030312301013100-0322123312330101"></a>

## Next pages — add_partitioned / 310322332112 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1120113031102033-3311223300132202-0003111130112133-1031302321113210-2001313331103213-3022322310000020-3321122301112321-3112111301012012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130111333032332-0130311033331312-1300000210120202-2312112121102110-2111123110022230-1310233120223312-0330011000201220-2301031323203312"></a>

## http_proxy.more_option.response_cookies_to_add.add_secure — add_secure / 012112021302 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.add_secure

<a id="canonical-1333133032032232-2020023020031330-2002000023320230-1303000230212332-0120301210200303-1121000321113200-1311023211220102-1231131312230222"></a>

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

<a id="canonical-1201323213101111-3031033222222223-2310233312000132-3123333213132300-0220012001022133-3102233000103103-3111123021323022-1110020332010332"></a>

## Direct properties — add_secure / 012112021302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110103301211220-1030211001321132-2303113303021112-0201033331123233-1130323302120011-0131113011102323-0230120222120222-0331201131020120"></a>

## Next pages — add_secure / 012112021302 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0230111310231220-0022230232320033-2001112110322132-0021320021301300-3012120001021120-3231132113231133-2000102033301312-0220221221123323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131101323330102-3300300212211301-1013012223003103-1000220021303212-3212322031303021-2320022320132211-1011122132320122-3232103313033310"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_domain — ignore_domain / 103200033010 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.ignore_domain

<a id="canonical-0221201230023312-0221323211011323-3332120332233110-2332022022123031-0122021132331132-2111021020131301-1102230013332100-0311030103022130"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore domain.

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
ignore_domain = {}
```

<a id="canonical-3210323021303110-3213310232100332-0022201302002033-3000130311200120-3222003012022230-0330123003012301-0322330212031321-2123201023013323"></a>

## Direct properties — ignore_domain / 103200033010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021101230133332-2022111130320031-3131322010330001-1003301332002333-1301121321031131-2032233201023222-2122100312011312-1012200021030110"></a>

## Next pages — ignore_domain / 103200033010 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2033222002030110-3233202113213202-2113331203022231-0320121110012031-2332232223012113-0203021121211312-3210003103210011-1032211210022333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023103103001111-1230330031211332-3323300133301030-1113212221022101-1013113000211332-3031130222123120-0323213011013232-3121222310013013"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_expiry — ignore_expiry / 320031102331 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.ignore_expiry

<a id="canonical-3301221122213231-1331033131130203-1323002020111003-3031230221203231-0311102313120022-2001212033133011-2000121011022112-1133021311323230"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore expiry.

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
ignore_expiry = {}
```

<a id="canonical-3220132203012001-3331010011311031-0201333303221213-2221123322030022-1322322002321231-2111300203003123-2210002212232123-1023122220000203"></a>

## Direct properties — ignore_expiry / 320031102331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210301033232301-3213031001210112-2012203123021133-1330102302230233-2332211011033102-1012331030020221-1232002130333012-1232032223112001"></a>

## Next pages — ignore_expiry / 320031102331 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0320001110232130-3312320130031211-2122113003123303-2212001322310122-3201312232320101-2302221233320111-1023331312003220-0013033130300121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312031100302113-2302012102113032-1213100111030300-0131200130322330-1210201120113131-1230332201010010-1111100332301330-1332121201101203"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_httponly — ignore_httponly / 101022003223 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.ignore_httponly

<a id="canonical-3021201202113032-3001303132200023-1000220310322101-0031023123032102-3231233121232202-2333023031120023-3120213203232020-2303132202332332"></a>

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

<a id="canonical-0003312121120032-2101120013023213-0030232220331000-3330230112201333-2030331313213020-0210022303211220-3013132221012012-2130000233001233"></a>

## Direct properties — ignore_httponly / 101022003223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110102013313123-3332012120121131-1220221000311320-3230133301001212-0032223301221223-3131123213232322-2001003001202010-2132120101033103"></a>

## Next pages — ignore_httponly / 101022003223 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3020333013311201-2021311030003000-1033033230122330-2011120333313213-0003132311113332-1033032010010130-3210101132310300-1100233001021330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212203222331002-1030031033221121-0023301210130213-3120232310103120-2032112102230132-2212213301330333-2211232331321011-0001203222003302"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_max_age — ignore_max_age / 200032023230 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.ignore_max_age

<a id="canonical-1220101022102300-1003133022313203-2003121220233210-0212233101212032-1221100331213001-1131010230211301-3033132101002103-3022323303022301"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore max age.

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
ignore_max_age = {}
```

<a id="canonical-0313021323233120-0332123100202312-0031330111102123-2202123000010313-2003032112023131-3120222211130202-3213201032333133-2112101103213102"></a>

## Direct properties — ignore_max_age / 200032023230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312233023022320-3212210133101121-3222002313031323-1320000122220012-1201030202223211-2230003132123010-1303310131333213-2231020013011322"></a>

## Next pages — ignore_max_age / 200032023230 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0123111011200201-0302223113101031-3212002000133333-0111011212322021-1103212133121322-0212013322333033-0230213123202330-0121300001003122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102103301313312-0121101132010210-2101213133333221-3033201100303021-1013003330231230-0011213301201210-3020110030012132-3330110333203321"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_partitioned — ignore_partitioned / 110221100011 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.ignore_partitioned

<a id="canonical-2100101233323032-1130213211301030-0001300120122332-1310030330221022-2220012210311133-1330210133022130-0231030201313003-3000211013303303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore partitioned.

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
ignore_partitioned = {}
```

<a id="canonical-3223111221222230-2211001201212323-0033302021133121-1120102222020022-0121222112112203-3032213001303223-0021213333330312-1130011011212312"></a>

## Direct properties — ignore_partitioned / 110221100011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021110300111003-3120130300030331-1222330231111232-2233020230021013-3212120210222320-1032001201002230-2222233303331303-0203213312332312"></a>

## Next pages — ignore_partitioned / 110221100011 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0321323213312032-0011221302101101-3201100322103203-2132110001122333-1103210003202103-2321100301321001-3300312221201033-3322021330302231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021313233230010-2101132221203022-3210032301322322-2022313312032323-2002120021120220-2310020333002330-1103031320323113-0030033022332022"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_path — ignore_path / 333110223101 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.ignore_path

<a id="canonical-3130232232002203-0101102011220030-0223203312101321-2133223023012313-1201330320002233-3231221133000232-2231230233021000-0220222132200313"></a>

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
ignore_path = {}
```

<a id="canonical-3231100131123113-3031120121220330-1331013201223301-2331023131311131-3103312310122022-2131200313331222-3202313303303201-2221200322231013"></a>

## Direct properties — ignore_path / 333110223101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303031112202023-1032100201302003-3302100013332312-2102310301002022-2101111111131101-0023221111332102-2120023032100300-0123001111220023"></a>

## Next pages — ignore_path / 333110223101 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2111200001020121-3110202000221230-0023303100300132-1122030032011310-1021120313102101-2102302300232131-3333011132032220-3212120331330203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100211332111212-3232231211101013-2012033010133023-1300031032211121-0113330310132020-2301332113222211-3100311020301112-2100210031223303"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_samesite — ignore_samesite / 303213101310 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.ignore_samesite

<a id="canonical-3133123012001302-3003231110111001-2013202131123130-2110021011013110-0131123300023112-1131230123222203-2013312012310022-3223111003033310"></a>

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

<a id="canonical-3303113331121030-0311303212000021-2033212311032113-1022222000023131-1011222130303022-1230133202321021-2301113202032033-3032310023200012"></a>

## Direct properties — ignore_samesite / 303213101310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233303200021220-3330010221220212-2223201232302311-1301132221102111-2301012032130202-3322133121202332-2113320023303020-3031302102013130"></a>

## Next pages — ignore_samesite / 303213101310 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0032120130021302-0201013033000220-0202332223110330-2031013200331123-1100100021332110-3133030202102031-2213210331220022-3030222301221230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102002132110203-0202201212212300-0303301100322211-3323130211030332-0010003313123313-0033202221331312-1001122222002030-0133312231330033"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_secure — ignore_secure / 012001332312 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.ignore_secure

<a id="canonical-1233333310313010-1110313202123232-3022121322303321-1321330103210013-2200223232312120-1000103133110022-3123132130001131-3210030332122332"></a>

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

<a id="canonical-1032323111312303-3110203020130131-0011232231202223-3223103202002011-0200021322101000-0111330011031010-0311112112320310-3022202033223201"></a>

## Direct properties — ignore_secure / 012001332312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321010333001133-2110211333133211-0300322011210322-0323330222030222-3312032012112033-1031201110300210-0020211133320303-0120302012122100"></a>

## Next pages — ignore_secure / 012001332312 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1031222023100232-3133320301231330-0312312111112012-1010303031233111-3003200000002033-0020112101021130-1221203331012312-1120103113021320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131002130122313-0331301323200320-1332322331103112-1201003231301120-1320303030202120-3332220013002021-1102131231020331-1322032201332120"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_value — ignore_value / 023313123213 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.ignore_value

<a id="canonical-2000011001011323-1300131221132110-3303122103213333-1011011021303021-2012212223232201-0003232211233331-1132133013322310-0020230200220121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore value.

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
ignore_value = {}
```

<a id="canonical-2212320030100301-3203211221122200-3110333311232322-1223323102210313-0022211010101103-2121020000222222-1020232120202211-2203130203201001"></a>

## Direct properties — ignore_value / 023313123213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020213103332110-1202002200131123-1230131031211102-0201120022330332-2033313011020212-3320233313223102-0013212312200102-2232123110111100"></a>

## Next pages — ignore_value / 023313123213 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1212103011120333-0223001021231133-1321303222211111-2100302322210312-0221333311012000-3223203131002122-2332032200121331-0211031130111321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332310301311213-3033133332132321-1231201223310133-2112030013213021-0313322201213131-3202313033301310-2022323130112132-0211310330202310"></a>

## http_proxy.more_option.response_cookies_to_add.samesite_lax — samesite_lax / 023310121022 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.samesite_lax

<a id="canonical-0323000230312231-2003211031102031-1103001100233000-1202311101333212-3020232303303111-2003230311033002-2012113111100201-2122332202210211"></a>

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

<a id="canonical-2031120231232022-1132010131333332-1122233032220121-2011102101322120-2032222303321130-1323202201031233-1311122103032022-1003201323103203"></a>

## Direct properties — samesite_lax / 023310121022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311321330333123-3322120022221320-1020131313301010-0002332303131313-1221223012123313-3033120231032003-2131012302130013-0100031032100301"></a>

## Next pages — samesite_lax / 023310121022 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3132332212330322-0202310233313222-2003022020022302-3113230213300121-0200202011212113-1020310103120000-1323101010233002-3312222223212020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112103021300220-2303123023022103-1021221111231022-3120323023201220-2012221131223231-2332012333102120-3103300201120320-2030111322022022"></a>

## http_proxy.more_option.response_cookies_to_add.samesite_none — samesite_none / 012210012302 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.samesite_none

<a id="canonical-0333030220112210-2021023111232103-0122120132213233-2033202133101020-1110312213101233-3331210333132111-3133213211111101-0130332211011103"></a>

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

<a id="canonical-3232102030300120-2301222210122133-2321102131113003-2002232031222202-0120232323100103-2120212232202123-2331101132133103-0022311011103220"></a>

## Direct properties — samesite_none / 012210012302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310023003222031-3012220223021001-0021030001323312-1310012303310131-0000201102210202-3321230210202032-2222130032203333-2323022332103023"></a>

## Next pages — samesite_none / 012210012302 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2222032212220010-0332200103321102-2131011023001102-3310033321132213-0103322202210010-0220312231211301-3310330222200123-3320210000213321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320303300320222-2322201221210130-1131200133330113-1301310000003111-1001033322030112-3333201302112010-1210300312110030-2121201312010220"></a>

## http_proxy.more_option.response_cookies_to_add.samesite_strict — samesite_strict / 120001133332 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.samesite_strict

<a id="canonical-1111222130212203-0131223230132311-3010311232233220-3111231010232322-3301221010120130-3211003012131223-0303220300123312-0302102210001222"></a>

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

<a id="canonical-3221212000330302-0131310221110212-1321313203333301-2112111321021102-2103002223320201-3231000012113000-0331232302312123-3220212110311203"></a>

## Direct properties — samesite_strict / 120001133332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321020022231330-3132120112211223-0220313330010331-2022003302101303-3331333032102203-0021311201323233-2321210033230230-2020201102010020"></a>

## Next pages — samesite_strict / 120001133332 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0032223220331321-2122103333322322-0112302321320233-1222120311123120-3122320100100323-2000101331022231-0320013223311032-0110033301203222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001100202131202-2322300211233123-2332223313213312-3202323030203030-3132213010023203-0102002233302011-3201102123210101-0223233200002201"></a>

## http_proxy.more_option.response_cookies_to_add.secret_value — secret_value / 200322220211 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- http_proxy.more_option.response_cookies_to_add.secret_value

<a id="canonical-1300211320131301-0203023131301210-1303101312013001-1112303003212202-0120331100100003-2001013102011323-2100133111233203-3101011302130220"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000103100332221-2233201221323302-2231112001331230-0233020211323002-3002133031112323-0020300210122021-1213333012313122-1200123200023111"></a>

## Direct properties — secret_value / 200322220211 / 3

- [blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-1210113233221230-1203323122300112-2012110322330310-2123210112123301-2110120032232111-0200223332121310-3122221101333302-1132231101001031): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-004.md#canonical-2111320201010202-1001012323322201-0233312100110032-2123030212002221-3332123331003122-0122100302330333-2233101000313130-2011200001003302): complete subsection reference.

<a id="canonical-3312113233331331-1102113110220323-1332123011231230-3111222121010022-1000201101230121-3103112231010112-0030313021101322-1020002130323013"></a>

## Next pages — secret_value / 200322220211 / 4

- [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-1210113233221230-1203323122300112-2012110322330310-2123210112123301-2110120032232111-0200223332121310-3122221101333302-1132231101001031)
- [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-004.md#canonical-2111320201010202-1001012323322201-0233312100110032-2123030212002221-3332123331003122-0122100302330333-2233101000313130-2011200001003302)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1210113233221230-1203323122300112-2012110322330310-2123210112123301-2110120032232111-0200223332121310-3122221101333302-1132231101001031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100330312301200-0222330323230300-1110223111232103-0232200313002022-1033230022202312-0322333130130121-2131201103022023-2200211122303322"></a>

## http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 101012010031 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-0032223220331321-2122103333322322-0112302321320233-1222120311123120-3122320100100323-2000101331022231-0320013223311032-0110033301203222)
- http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-3333123130100000-2100030320313133-1023322130300210-2110122230231023-1003110310131130-2231022121323302-0201202132100000-0313013132223032"></a>

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

<a id="canonical-1220000310211323-1030302032122301-2212200211001012-0213310302210002-2012131200003203-3330331232012013-0301102033013220-3013122103100210"></a>

## Direct properties — blindfold_secret_info / 101012010031 / 3

<a id="canonical-1322101213233323-0021003222021021-0313103233202203-0312133311320113-0131331130102022-0121103111122102-0001233021330003-3030033331111011"></a>

<a id="canonical-0032321033231031-0110000000012123-3311123220203111-1302312023013230-3111130330233102-2102302230122133-0203000200003122-1232313000232232"></a>

## decryption_provider property — blindfold_secret_info / 101012010031 / 4

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

<a id="canonical-3310112200003233-0330310212312012-1112233102021320-3111222012311013-0301010202320300-1023021312220311-0032032112202232-0312221232123101"></a>

<a id="canonical-3001313112123332-0011103123312131-0030332323001303-2131022130030133-2220313323202230-1110003303033312-1033210113321013-3333203311122301"></a>

## location property — blindfold_secret_info / 101012010031 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0121031200311302-0221100023112130-1130220012202102-0122110212010003-3322131220212233-1210200213131113-1113030212001213-2303100300012211"></a>

<a id="canonical-0101102110000230-3333231003333133-2131132223312012-1230301011330133-0130021020111112-3223330020220002-2002012031113302-0311103300101230"></a>

## store_provider property — blindfold_secret_info / 101012010031 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-1032232233130211-0323103130110031-0321231301121301-0110313100301321-3132030231231110-0002003121311312-3111200100032002-3102323121232322"></a>

## Next pages — blindfold_secret_info / 101012010031 / 7

- [http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-0032223220331321-2122103333322322-0112302321320233-1222120311123120-3122320100100323-2000101331022231-0320013223311032-0110033301203222)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2111320201010202-1001012323322201-0233312100110032-2123030212002221-3332123331003122-0122100302330333-2233101000313130-2011200001003302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003111303301111-1313123220300123-3301121123120001-3302121220313130-2031333002302002-0121223022321111-3310130312221033-1301121101303203"></a>

## http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info — clear_secret_info / 202103003211 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-0032223220331321-2122103333322322-0112302321320233-1222120311123120-3122320100100323-2000101331022231-0320013223311032-0110033301203222)
- http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-2100103200111022-1232311232303131-3302200200002321-0030031313122311-0130110020122031-3223311213201232-1103222003022021-3112313133223303"></a>

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

<a id="canonical-3202130313032021-1031000111001131-3130112300111121-3102131121012303-1010131331121310-0300230100030303-0302010023213212-2321322200212213"></a>

## Direct properties — clear_secret_info / 202103003211 / 3

<a id="canonical-0332120021321111-2101201130210302-1221133321311313-1120312111011021-0302231010213031-0211111202131032-2210310133311312-1322333013301002"></a>

<a id="canonical-2333311130202001-3102220113023023-2212123010032233-0001333133310121-2300113033212120-3302331230100030-1222100212220213-1001023222332123"></a>

## provider_ref property — clear_secret_info / 202103003211 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1232330320220220-0010100010231021-1003301110333302-2103220211230100-2203023302020221-0020320032110313-2301113230320311-1110331111003023"></a>

<a id="canonical-2303002323003320-1230021110003000-1220221220312031-3230101010013303-2002210231201311-0123032203020200-3231022202330331-2003331121002122"></a>

## URL property — clear_secret_info / 202103003211 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0101013121211323-1122311331231032-2123023301003303-0200302100013232-3011232201302000-2112030222312210-2310001223031012-0213031132110102"></a>

## Next pages — clear_secret_info / 202103003211 / 6

- [http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-0032223220331321-2122103333322322-0112302321320233-1222120311123120-3122320100100323-2000101331022231-0320013223311032-0110033301203222)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0321302212120332-1003310003320101-2331211322200330-1231022123102030-0001322133103022-1202202120331013-0312112103131001-3233331200001032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310323122233211-0000102322010201-1332032101323203-2210313231021022-1120013021012001-1322032102202333-1130333210312230-1312020110300230"></a>

## http_proxy.more_option.response_headers_to_add — response_headers_to_add / 020033323001 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- http_proxy.more_option.response_headers_to_add

<a id="canonical-1333331213213220-1313002212301213-0311101201203133-0102130131011013-3101203013301303-1021301033120000-3213210013232103-1221322223212331"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
response_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-3102213121020230-3323311203011121-1331302033012112-1123102011022112-3013331021123202-3102221212312001-2113033112002111-0013220213031211"></a>

## Direct properties — response_headers_to_add / 020033323001 / 3

<a id="canonical-3303320120322021-1301322213313021-1233212312311013-0323322032320113-1212323001011330-0302021030210220-2201132011212212-1312031111030303"></a>

<a id="canonical-2030100032301000-2203311221310321-3233321333312223-0203132112132132-2220120020310100-2122032002132232-1213021132121221-2030322333112000"></a>

## append property — response_headers_to_add / 020033323001 / 4

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3313221033320223-3302221133110232-3022030031120130-3320111101111302-3122332312323210-1122212303010201-0120121100301133-2101033112331321"></a>

<a id="canonical-3332003121122100-0132000103022020-3223123022120211-2231303022301113-0330302112033203-3111300023020003-3010313331212222-2221113231303012"></a>

## name property — response_headers_to_add / 020033323001 / 5

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](resources--proxy--reference--group-004.md#canonical-2300011033210230-3200023212122302-3133032202031123-3031312001002303-2211020320303111-3210011300022320-2331101223122002-1130322131232013): complete subsection reference.

<a id="canonical-2230222201003303-1330320021123112-0020112223121310-0231110111313132-3022211301313222-3131033301313020-0031121301023313-2013321012033022"></a>

<a id="canonical-1002303320130003-3001201112123102-0020013101022022-2322110332321230-0022331120112133-2033230032103303-0211202022022312-2003031330102111"></a>

## value property — response_headers_to_add / 020033323001 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-1230211121022202-3211013333120000-3013000220020031-3033002233113011-2030310103223233-2031132010113132-3111332330211013-0022022011013203"></a>

## Next pages — response_headers_to_add / 020033323001 / 7

- [http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-2300011033210230-3200023212122302-3133032202031123-3031312001002303-2211020320303111-3210011300022320-2331101223122002-1130322131232013)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2300011033210230-3200023212122302-3133032202031123-3031312001002303-2211020320303111-3210011300022320-2331101223122002-1130322131232013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033322311233332-0313002332110220-2123110123202121-0000212031302330-1003333312131010-3212023313210001-0033101222102210-2023212000202122"></a>

## http_proxy.more_option.response_headers_to_add.secret_value — secret_value / 201022121230 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-004.md#canonical-0321302212120332-1003310003320101-2331211322200330-1231022123102030-0001322133103022-1202202120331013-0312112103131001-3233331200001032)
- http_proxy.more_option.response_headers_to_add.secret_value

<a id="canonical-3101123131023332-1232313122120131-2221032013303333-0122202011000103-3213012213321111-1320331103202313-1202132103012100-2111002322232122"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-3332230332022003-0321001333131023-1210311002113100-3210021321101331-1203230102110321-1322020300331220-3333322232012012-3013331301311310"></a>

## Direct properties — secret_value / 201022121230 / 3

- [blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-1001121112013303-2001222331111110-0131023301312020-0200231122202010-1223312200222032-3230330332320230-2311000203321210-2132122001323023): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-004.md#canonical-1010013020132002-1033212202221011-2332220021033303-0313002000023320-3113323013011221-1103111030031202-3233202312232001-3330221032020030): complete subsection reference.

<a id="canonical-2023112133311323-3133223020131033-1000002112121311-0310201303011233-2202111330131122-0300133012030222-2022203320131012-2130120301323201"></a>

## Next pages — secret_value / 201022121230 / 4

- [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-1001121112013303-2001222331111110-0131023301312020-0200231122202010-1223312200222032-3230330332320230-2311000203321210-2132122001323023)
- [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-004.md#canonical-1010013020132002-1033212202221011-2332220021033303-0313002000023320-3113323013011221-1103111030031202-3233202312232001-3330221032020030)
- [http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-004.md#canonical-0321302212120332-1003310003320101-2331211322200330-1231022123102030-0001322133103022-1202202120331013-0312112103131001-3233331200001032)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1001121112013303-2001222331111110-0131023301312020-0200231122202010-1223312200222032-3230330332320230-2311000203321210-2132122001323023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312233110300301-3311223000012321-2321010121032022-1132131021212130-1331132123133303-0130133221031303-1012330231030033-3033011002023210"></a>

## http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 133330201003 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-004.md#canonical-0321302212120332-1003310003320101-2331211322200330-1231022123102030-0001322133103022-1202202120331013-0312112103131001-3233331200001032)
- [http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-2300011033210230-3200023212122302-3133032202031123-3031312001002303-2211020320303111-3210011300022320-2331101223122002-1130322131232013)
- http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-0121133002111022-1002000221032031-2000113021022120-3102302022222122-1120232323303023-1333312301201303-1230031121122202-1210320112231313"></a>

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

<a id="canonical-3120200300013223-2013131113000110-3130022000201210-3123011311320113-2003321123232033-0123121011300020-0033333221013201-1003220001122113"></a>

## Direct properties — blindfold_secret_info / 133330201003 / 3

<a id="canonical-2110212123001032-1101302133002011-2102223322230132-0323011323112000-3312012012213031-2230213230111003-3221111023332321-0310132232232210"></a>

<a id="canonical-0113321321330232-3311011332001131-3233332232232133-2302000030123122-3230211113120312-0330010132231330-1110223223131210-2130222100102322"></a>

## decryption_provider property — blindfold_secret_info / 133330201003 / 4

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

<a id="canonical-3231122103112003-3313302033100331-1321100220000233-1013220000002203-1331121230000032-0110231310013311-0002211203010200-3310312003113222"></a>

<a id="canonical-2320121200321332-3010321031121212-1212310231103312-1222113132032233-0001200113323030-3201302232232001-2231032013100003-1010002312132231"></a>

## location property — blindfold_secret_info / 133330201003 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3013103000030102-0023123021032023-1321310333023132-3122032233111331-2202023032013132-3212233000312313-2230002200220123-0223012100001311"></a>

<a id="canonical-2222330013122120-2013301000022032-3202303331221112-0311301221121230-1303200320000322-2322311330021230-2323301211302310-0110230102231013"></a>

## store_provider property — blindfold_secret_info / 133330201003 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-2011311002331333-3211201122000232-1013231312202211-3303300010202033-1311330131003031-2332133023200220-2013110303113132-0320222110100102"></a>

## Next pages — blindfold_secret_info / 133330201003 / 7

- [http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-2300011033210230-3200023212122302-3133032202031123-3031312001002303-2211020320303111-3210011300022320-2331101223122002-1130322131232013)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1010013020132002-1033212202221011-2332220021033303-0313002000023320-3113323013011221-1103111030031202-3233202312232001-3330221032020030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221033320231023-3010201310033232-3201130312233133-0031010233301012-0220312312121230-0102200031003121-2001001130021312-1133220032201020"></a>

## http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 131110131033 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-004.md#canonical-0321302212120332-1003310003320101-2331211322200330-1231022123102030-0001322133103022-1202202120331013-0312112103131001-3233331200001032)
- [http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-2300011033210230-3200023212122302-3133032202031123-3031312001002303-2211020320303111-3210011300022320-2331101223122002-1130322131232013)
- http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-2200213112232300-3313331032330221-1202320001001221-0020102133011231-1300131333022001-0220201012233223-1332132221012203-3003301310020330"></a>

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

<a id="canonical-0101113223010001-3021102101313232-2123212021030310-0112313120222100-2231011003023211-3020031022331312-0223011001131232-2033033031331001"></a>

## Direct properties — clear_secret_info / 131110131033 / 3

<a id="canonical-1020230230001113-2222202233133031-3223103120332120-3002322321003020-2102310212230102-3213032033033320-3232300030202000-0122013112300020"></a>

<a id="canonical-2103221332012223-2221021232210220-3302230222331211-3122102213312003-1201003022312212-2201303123022003-3030000323103130-1201300303023010"></a>

## provider_ref property — clear_secret_info / 131110131033 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0330010312031012-1031211200003032-1113101120121000-0103322101000003-1123231100201303-1330122212321002-2013223333021300-3311313330112122"></a>

<a id="canonical-2000121030112232-0013123310320232-3213212100100013-1301012211130300-2020301233203011-1002232013010202-2110201312131132-0013331120303113"></a>

## URL property — clear_secret_info / 131110131033 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1300133101232132-0202312321132011-0030112302330001-0113200231332310-2022131231012313-2211233000110321-1213201330231123-0022223010300030"></a>

## Next pages — clear_secret_info / 131110131033 / 6

- [http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-2300011033210230-3200023212122302-3133032202031123-3031312001002303-2211020320303111-3210011300022320-2331101223122002-1130322131232013)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3333131321121110-3221311322032300-1111103312333233-0321012201211000-1102333201111020-0020303330301101-2313211321302213-0221203120310222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031201322121110-0333101120123232-0202103230321003-2123300303102123-2020203020133130-0013013022323031-2013002100311132-3133211020121022"></a>

## no_forward_proxy_policy — no_forward_proxy_policy / 111012302101 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- no_forward_proxy_policy

<a id="canonical-1302311121130102-2131203200103320-1303323313202210-0113023213112321-0121211321110201-2311112312303120-1201103232201010-1112001213313132"></a>

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
no_forward_proxy_policy = {}
```

<a id="canonical-1123221112101230-0322213232021123-0322211302222130-2030000110130102-2130010131010202-2102001103313233-3001000111223233-3223031300123110"></a>

## Direct properties — no_forward_proxy_policy / 111012302101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212121310012202-0313022303311001-2232113031013103-2122022002020210-0023032233021102-0011121302103010-3333012313310002-2000313120033223"></a>

## Next pages — no_forward_proxy_policy / 111012302101 / 4

- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3001311021210113-1203322120301211-2003112333102203-1131010301220310-1212123211110331-3201330202033133-1001002102301300-0023323331213323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132020323132132-3122302023133003-2220101202203300-3303103112103202-3330032122030310-2110102121111120-1013222133313001-0131100301230111"></a>

## no_interception — no_interception / 310302201300 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- no_interception

<a id="canonical-3332023313312012-3012310221032321-3101200211301200-2222112211033303-1033311031230203-0102333033323201-3101321210231120-3233331221000130"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_interception, tls\_intercept; Default: no\_interception\] Configuration parameter for
no interception.

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

- [no_interception](resources--proxy--reference--group-004.md#canonical-3332023313312012-3012310221032321-3101200211301200-2222112211033303-1033311031230203-0102333033323201-3101321210231120-3233331221000130)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-2300212110133011-2303221131021333-1221330301300022-0231220012001111-0213123123321222-0000030203023101-1203222021230313-0321231013011230)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_interception = {}
```

<a id="canonical-2102221313303323-1312301303232330-2222213301123002-2102301320233023-0001223123302331-0211100132201131-0001203331203202-0120210203211022"></a>

## Direct properties — no_interception / 310302201300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000132100010100-3320200133221202-1103132233133201-1000111221210211-3101032101011110-2232310202103113-2122213011030121-1033010221022033"></a>

## Next pages — no_interception / 310302201300 / 4

- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2300220130222220-0321012100023110-3312120100301031-2332310123110321-0222103331130303-0122320111223322-2100122213201320-3222032202331000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
