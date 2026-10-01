---
page_title: "xcsh_nfv_service reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service reference."
---

# xcsh_nfv_service reference

<a id="canonical-1101303313103303-3332101110003303-2303203120112330-3212201201133313-0001120022003112-3302013202022021-0302021120022112-2221212131102131"></a>

## namespace property — crl / 220122111011 / 5

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

<a id="canonical-1102002020022313-3002213222313030-2122200132033223-2020322321000131-2302231201313313-1221233032013320-3103322231322132-2110312132131301"></a>

<a id="canonical-3102101100003323-2001220033123312-1130122100222322-2022230103110103-3312313331232101-2001033120113323-3003112312320322-3311111313221231"></a>

## tenant property — crl / 220122111011 / 6

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

<a id="canonical-2110232220121101-2131233213230302-3321001131201330-2023001010232131-2313010323020020-0222131213130230-0100311322111210-3212301011132031"></a>

## Next pages — crl / 220122111011 / 7

- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-2220303111010322-0112002023201232-0133012112231002-3103111320133100-1120331002010022-2320111201201332-2112303320000210-2130023033333112)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-3003031320320332-0321333001323110-1022203213110311-1323101231313031-3013220113030013-2111132110031203-0102112031312132-0230332133312303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311233121232322-1112200031021113-1020112300313203-3121221322333132-0023101201131332-3200101213322112-2202130331233130-2231113130310110"></a>

## https_management.advertise_on_slo_internet_vip.use_mtls.no_crl — no_crl / 101101113013 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-3031132200011022-1313221310121032-3200012331302121-0120313212301311-1101211132233033-2113333100000023-0121110203223213-0303202010232130)
- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-2220303111010322-0112002023201232-0133012112231002-3103111320133100-1120331002010022-2320111201201332-2112303320000210-2130023033333112)
- https_management.advertise_on_slo_internet_vip.use_mtls.no_crl

<a id="canonical-0131033311220010-0111221133010001-3013200000202122-2103123033323133-3200032303132331-3102220030203012-3020231213211231-2111122120210021"></a>

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

<a id="canonical-1100010213311031-2233201203303333-0222032121233220-0103201120301121-1133222311001012-2230301302222303-2200313200210310-2031113130211221"></a>

## Direct properties — no_crl / 101101113013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320213013113033-1312202323331021-0332312223210032-1333220103131312-3000133121301133-1312331112133033-3112303332130331-0332121121031112"></a>

## Next pages — no_crl / 101101113013 / 4

- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-2220303111010322-0112002023201232-0133012112231002-3103111320133100-1120331002010022-2320111201201332-2112303320000210-2130023033333112)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-0300010231301102-1300013232001021-3130003100132310-0320033103133301-2103113201012311-0100313003232233-0300003132301301-3022202303302030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321011223101321-2100001313012030-0012033131113031-1332121212333001-1103222220022210-0130213113320031-2003122032031210-3201231102130201"></a>

## https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca — trusted_ca / 312131111221 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-3031132200011022-1313221310121032-3200012331302121-0120313212301311-1101211132233033-2113333100000023-0121110203223213-0303202010232130)
- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-2220303111010322-0112002023201232-0133012112231002-3103111320133100-1120331002010022-2320111201201332-2112303320000210-2130023033333112)
- https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca

<a id="canonical-3212200110222312-2321131210222313-0321011000130032-0132302100212101-2332033032103311-1311011031113112-1220032102033323-2130201311123212"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200223222023002-2003101032330030-3020013301302300-3210313121303223-0132032111231012-0301210100320200-1232222203221103-1320023332220033"></a>

## Direct properties — trusted_ca / 312131111221 / 3

<a id="canonical-3212120022223222-3323031213022133-2330220203100002-0332121310012230-3122133121100003-3003222331313311-2331321300310332-1132031122320200"></a>

<a id="canonical-2313220133110301-2113133232033203-1200110310302112-0013310213033310-2312331022311030-3330012101323330-0111132302001302-2032300332000223"></a>

## name property — trusted_ca / 312131111221 / 4

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

<a id="canonical-2030202333213210-1003313111010211-3322033331322103-1330321133302020-1101112213111001-0113001121213312-1211100212211121-2131030200121321"></a>

<a id="canonical-1301303021031323-2010123130101122-0213130232031023-2102013230222011-2022101010130203-0201202001332233-2231223221032220-1212333100231202"></a>

## namespace property — trusted_ca / 312131111221 / 5

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

<a id="canonical-3030112230010031-2102221330322332-1331103202310132-1202210200111200-3112222202203120-3310100031220003-1331221212102212-1200331033122301"></a>

<a id="canonical-2220111201003220-1021020312131021-0121011112111333-0223122120021030-3031010000211023-0221120213200231-0102301111312030-2123322001112320"></a>

## tenant property — trusted_ca / 312131111221 / 6

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

<a id="canonical-2100023221013122-1200220331003110-2113333303122330-3332232220103313-1231113303330332-3322120230013022-1121102002321011-2211232111313322"></a>

## Next pages — trusted_ca / 312131111221 / 7

- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-2220303111010322-0112002023201232-0133012112231002-3103111320133100-1120331002010022-2320111201201332-2112303320000210-2130023033333112)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-3110102211112330-1123023120010002-1033312223202210-0123212101233111-2300030310312033-2030132333310031-2100100100010222-0213013313320231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123223211022012-2003231120020120-0203230003332322-1200123123003303-0011023031300011-0121030211130211-3312012132031030-0333112001211132"></a>

## https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled — xfcc_disabled / 213313313222 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-3031132200011022-1313221310121032-3200012331302121-0120313212301311-1101211132233033-2113333100000023-0121110203223213-0303202010232130)
- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-2220303111010322-0112002023201232-0133012112231002-3103111320133100-1120331002010022-2320111201201332-2112303320000210-2130023033333112)
- https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled

<a id="canonical-3132303122120001-2322032211030112-2133320303120223-1113203331002120-2033212003130223-2210110300121130-0330022321112210-2122012223331113"></a>

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

<a id="canonical-3111233121310110-1323331322232203-3002001200121221-0310002032033130-2032012111223131-1103122313220223-2022103121101332-3211122100032203"></a>

## Direct properties — xfcc_disabled / 213313313222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033023103301221-2130102311203113-0313303332120310-3112331221021220-0200020132130203-0201012222020101-0012333110302030-2330321100331331"></a>

## Next pages — xfcc_disabled / 213313313222 / 4

- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-2220303111010322-0112002023201232-0133012112231002-3103111320133100-1120331002010022-2320111201201332-2112303320000210-2130023033333112)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-2212213112313311-3122110131300103-3021131321320332-2131122231132020-2111013001221220-1310212122100202-1200033003222330-1203302030311303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233021033312020-1001311131303012-3231001032213123-2301110200222333-2202220032012000-3003331012222020-3112202323203213-2022013211202020"></a>

## https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options — xfcc_options / 110302023120 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-3031132200011022-1313221310121032-3200012331302121-0120313212301311-1101211132233033-2113333100000023-0121110203223213-0303202010232130)
- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-2220303111010322-0112002023201232-0133012112231002-3103111320133100-1120331002010022-2320111201201332-2112303320000210-2130023033333112)
- https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options

<a id="canonical-3212211231012230-0230223010133130-2030210000303021-2300303312011203-0132123010231213-0031113032111122-1013031112222301-0210023022120133"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2330231122011223-3222320100231010-2232211322000122-3331321132112021-1012321031213333-1130333321133222-1010222133333310-2013030202020130"></a>

## Direct properties — xfcc_options / 110302023120 / 3

<a id="canonical-2110300210103203-1200312202131301-1030000130232011-1130002233201310-2210033223021203-1313032132011213-3003010333022333-1032102101221011"></a>

<a id="canonical-0121132130333312-2113133101013233-0211213023213131-0311203110201023-3102112020332311-0323230311010321-1031320332120033-0312130011110132"></a>

## xfcc_header_elements property — xfcc_options / 110302023120 / 4

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

<a id="canonical-0202213212220230-2220030220323132-1313110203300100-0313233010333222-3323002332103001-3310123220033233-2321233032301212-3223203002221300"></a>

## Next pages — xfcc_options / 110302023120 / 5

- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-2220303111010322-0112002023201232-0133012112231002-3103111320133100-1120331002010022-2320111201201332-2112303320000210-2130023033333112)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122030120301322-0000133033222003-0303220131332212-2320322331233211-3133331012312320-3022221313110130-1003303031111033-2101123003230020"></a>

## https_management.advertise_on_slo_sli — advertise_on_slo_sli / 012131022331 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- https_management.advertise_on_slo_sli

<a id="canonical-1222310230103003-2030221101312331-3000330233111031-2200303020312322-0110001110321021-2333201123200123-1332223311201110-2330030101210230"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for advertise on slo sli.

Upstream description:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
advertise_on_slo_sli {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033322131220030-2220103133322302-0311112023101133-2102031130103201-3200303012012033-3223210000321002-3311210130001023-0021031021232130"></a>

## Direct properties — advertise_on_slo_sli / 012131022331 / 3

- [no_mtls](resources--nfv_service--reference--group-003.md#canonical-1130313331121202-0111232220003011-0103033303320223-2120021203213211-0121112223201032-3110331102023222-3212212010031302-0310211322013020): complete subsection reference.

- [tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113): complete subsection reference.

- [tls_config](resources--nfv_service--reference--group-003.md#canonical-3220330322200213-1131332232110011-2033122101021203-3133303310201321-1100333102323010-2210233023331021-2221332001303332-0331211032201031): complete subsection reference.

- [use_mtls](resources--nfv_service--reference--group-003.md#canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310): complete subsection reference.

<a id="canonical-3022333220011111-2113200213321130-0330303310010222-2023313332122030-3302131101032111-0132212232321302-3122012112112322-0232101021111132"></a>

## Next pages — advertise_on_slo_sli / 012131022331 / 4

- [https_management.advertise_on_slo_sli.no_mtls](resources--nfv_service--reference--group-003.md#canonical-1130313331121202-0111232220003011-0103033303320223-2120021203213211-0121112223201032-3110331102023222-3212212010031302-0310211322013020)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113)
- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-3220330322200213-1131332232110011-2033122101021203-3133303310201321-1100333102323010-2210233023331021-2221332001303332-0331211032201031)
- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1130313331121202-0111232220003011-0103033303320223-2120021203213211-0121112223201032-3110331102023222-3212212010031302-0310211322013020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132321331233220-1130120122233032-1332333300322232-2202323123222030-3320223320332232-3013230033300231-2210003110213001-1230132111012110"></a>

## https_management.advertise_on_slo_sli.no_mtls — no_mtls / 201212200103 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- https_management.advertise_on_slo_sli.no_mtls

<a id="canonical-0322033010311130-0022320003121201-3113011020223313-3232231003321320-0300230030123202-2203122112030232-0323133000022110-0022203221323023"></a>

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
no_mtls = {}
```

<a id="canonical-3333221230201331-0211000312011200-2023110310201223-2311133203110213-3300013331030320-1210102330323231-3122031323313002-2230221122210123"></a>

## Direct properties — no_mtls / 201212200103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323331003200023-0231333013013213-0120212301111001-3221100223123120-3302320330121212-0223230232313203-2211321122123210-2101123023000021"></a>

## Next pages — no_mtls / 201212200103 / 4

- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223033330121120-0012332331212110-1110312011122322-2032320301002103-1321233013323212-2132113312333321-2110123113120003-1100310120102223"></a>

## https_management.advertise_on_slo_sli.tls_certificates — tls_certificates / 131021002002 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- https_management.advertise_on_slo_sli.tls_certificates

<a id="canonical-2222132112223302-0332312122221233-0320110121020203-1031001022030002-3123021201112213-3020300223003103-0003103121332201-1223202131101201"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130113221003031-2333212331202131-3031322300001011-3113121223231313-0102100130021101-3102020202312310-1320032023122300-3102203202301211"></a>

## Direct properties — tls_certificates / 131021002002 / 3

<a id="canonical-0110213100123121-1000322120212120-0321033020211011-2203113110023200-3003321002200333-3012330102212333-1000121101130321-0330221303001100"></a>

<a id="canonical-2121203200230100-2203011120331131-3102021103233212-1100011120311312-0013012302112302-3232132033002303-3031320322113223-1101232310220010"></a>

## certificate_url property — tls_certificates / 131021002002 / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](resources--nfv_service--reference--group-003.md#canonical-2303321112330003-2223311011030233-2332321211031112-1201311231223013-3222201133030120-3230230230032332-1233230330110210-3213012313331132): complete subsection reference.

<a id="canonical-1010202211022311-0211131011331123-1002311102113312-2001110233223000-0201022233222112-3310013103000230-1333211321312202-0110111232332013"></a>

<a id="canonical-3312233230212211-1200031321033012-1232210301302011-1310112200101130-2310020331332312-1100213021321133-3012001110131312-2113312323213320"></a>

## description_spec property — tls_certificates / 131021002002 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--nfv_service--reference--group-003.md#canonical-1230130333103303-1031122031320031-0120200122220002-0022303132211113-2032232123131120-3302222223003330-3112110311213311-2301321032011110): complete subsection reference.

- [private_key](resources--nfv_service--reference--group-003.md#canonical-2220121131203232-3202100013123232-2221330330202323-3131100122221013-3221230232121021-3133003111020322-0103003210222311-3230101011023000): complete subsection reference.

- [use_system_defaults](resources--nfv_service--reference--group-003.md#canonical-1032013113221221-2330201313131201-2012232102321120-2021021033220122-0023000111021213-0131113010312113-1231110201003300-1333012213221332): complete subsection reference.

<a id="canonical-2003032222202132-0213331303023311-2333122103332100-0101110122012103-1012330023213022-1303010120012323-1330010012302231-3312311113001211"></a>

## Next pages — tls_certificates / 131021002002 / 6

- [https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms](resources--nfv_service--reference--group-003.md#canonical-2303321112330003-2223311011030233-2332321211031112-1201311231223013-3222201133030120-3230230230032332-1233230330110210-3213012313331132)
- [https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling](resources--nfv_service--reference--group-003.md#canonical-1230130333103303-1031122031320031-0120200122220002-0022303132211113-2032232123131120-3302222223003330-3112110311213311-2301321032011110)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-2220121131203232-3202100013123232-2221330330202323-3131100122221013-3221230232121021-3133003111020322-0103003210222311-3230101011023000)
- [https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults](resources--nfv_service--reference--group-003.md#canonical-1032013113221221-2330201313131201-2012232102321120-2021021033220122-0023000111021213-0131113010312113-1231110201003300-1333012213221332)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-2303321112330003-2223311011030233-2332321211031112-1201311231223013-3222201133030120-3230230230032332-1233230330110210-3213012313331132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122333100330212-0123020031233202-1223013313101103-1133001221223121-2100221132212210-2000121330102201-1003202212312211-1211120202020031"></a>

## https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 113302201221 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113)
- https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms

<a id="canonical-0023013102032200-0233131301222211-1200232321210020-1320032131133020-2133332103012210-1200220200200212-0122333030002123-2030200312110010"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333122310121020-3123313010311011-2313330203002101-1223132311131202-1020300000122330-1132100223213202-0313101300331203-1201102220310132"></a>

## Direct properties — custom_hash_algorithms / 113302201221 / 3

<a id="canonical-1003032001010322-2110330013122202-2113312222103000-1321220133331130-3013311000212331-1020122231221201-0113000212023313-3131301132321233"></a>

<a id="canonical-2113221323101233-2320302201122303-1120321211003121-0133301301100110-3320300031121222-0032323300033133-3122232322113133-0003231320020003"></a>

## hash_algorithms property — custom_hash_algorithms / 113302201221 / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1003000223101200-1223231000210223-1120330320131102-3023200321233230-2222120313233310-3112321202002230-3211110322303223-1122010002121301"></a>

## Next pages — custom_hash_algorithms / 113302201221 / 5

- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1230130333103303-1031122031320031-0120200122220002-0022303132211113-2032232123131120-3302222223003330-3112110311213311-2301321032011110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100123320212331-0120210200122121-2000102003222200-1001032211223303-1221030030233302-3332221012333311-2121332211300230-0200112022311111"></a>

## https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 022020020200 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113)
- https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling

<a id="canonical-3011233100130321-3013222003232231-0230100121203322-1110102211221320-0033322213102203-2333210010333202-1012231213200221-3121103121133003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

<a id="canonical-3210332132202330-2002303103130123-2221013102313003-3200132113003022-3123212113323201-3103222110211011-3102012330330310-3001032311030331"></a>

## Direct properties — disable_ocsp_stapling / 022020020200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231101230313000-3313032013111313-0330321010033330-0313322203111311-0202102320323331-1022230302033213-0210303010322223-1321202100310201"></a>

## Next pages — disable_ocsp_stapling / 022020020200 / 4

- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-2220121131203232-3202100013123232-2221330330202323-3131100122221013-3221230232121021-3133003111020322-0103003210222311-3230101011023000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130132110221020-2020321230011111-0223120131233102-3221301000303210-1032300333010013-3013123020121312-3212110300112122-1101001113031311"></a>

## https_management.advertise_on_slo_sli.tls_certificates.private_key — private_key / 132300331130 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113)
- https_management.advertise_on_slo_sli.tls_certificates.private_key

<a id="canonical-0211221232320211-2013103032113323-3103010030102020-2220102022233032-3200110030000211-3113311111231130-1203213300323010-1202312333113123"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
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
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-2231223031133231-0021220123131230-1010111212020102-2211101102322310-0123310021233102-2200012132311331-0002213023302001-2333120333002323"></a>

## Direct properties — private_key / 132300331130 / 3

- [blindfold_secret_info](resources--nfv_service--reference--group-003.md#canonical-2100111301202103-0330110312010320-3330323203130132-1012211113222112-2131203211031121-0102023122232121-0210211010102122-2110133113203230): complete subsection reference.

- [clear_secret_info](resources--nfv_service--reference--group-003.md#canonical-0231010001133320-3230032322031013-1321332103323213-0120120300231233-1221100220101331-2011202010330002-1220120223303312-0132300132110233): complete subsection reference.

<a id="canonical-1222221310131321-0103211321310012-2303233232201121-3222032131233222-0132231121031223-1023032110223110-2133212000100022-0022102210101230"></a>

## Next pages — private_key / 132300331130 / 4

- [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--reference--group-003.md#canonical-2100111301202103-0330110312010320-3330323203130132-1012211113222112-2131203211031121-0102023122232121-0210211010102122-2110133113203230)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info](resources--nfv_service--reference--group-003.md#canonical-0231010001133320-3230032322031013-1321332103323213-0120120300231233-1221100220101331-2011202010330002-1220120223303312-0132300132110233)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-2100111301202103-0330110312010320-3330323203130132-1012211113222112-2131203211031121-0102023122232121-0210211010102122-2110133113203230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233131221332220-3201120231123221-2230130121201001-2123212233321321-2302202131103101-0101330310313222-2101023211320233-1130001300201331"></a>

## https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 320133131322 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-2220121131203232-3202100013123232-2221330330202323-3131100122221013-3221230232121021-3133003111020322-0103003210222311-3230101011023000)
- https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-0122301330013003-3232321333221120-1303310131300323-0233113032301133-0013112010302101-3222110012001311-0333020021210111-3220003330122300"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2321210313210230-0002312120113313-3102313222023131-2331010000131133-2231230212111313-0303021203230323-3023331222131100-3102133101113202"></a>

## Direct properties — blindfold_secret_info / 320133131322 / 3

<a id="canonical-3222003112330000-1031331012330322-3122120131022013-0323322032120322-1102222200231101-0033123130322010-1100332322322301-1202200320002020"></a>

<a id="canonical-0230020320132222-0200332232102322-0220102013033101-2133110103213212-1333030221211322-1212212131123112-3031333133113033-0021201312002320"></a>

## decryption_provider property — blindfold_secret_info / 320133131322 / 4

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

<a id="canonical-2232312211233331-3100112233111212-0111103012123122-3332012230001002-2020001032202102-0010122320002030-1001301223102123-1322233211313210"></a>

<a id="canonical-2012001032230001-0210220003121020-1012231003013012-0300102320310111-0122021131302130-0112322023033130-0220212201201022-3002133333202013"></a>

## location property — blindfold_secret_info / 320133131322 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1023100103112010-0003032120211000-0113100122223320-1222121132113200-0103221232233022-0110322100211102-0113022000131302-2023212010112312"></a>

<a id="canonical-0022313211022010-0203000303120322-2030223112113130-0212032331131012-1222313112122031-2031313330001200-1223211101022030-1101320001213311"></a>

## store_provider property — blindfold_secret_info / 320133131322 / 6

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

<a id="canonical-1320112110311012-3332223323212301-2331233101000102-3021132110203120-0330012022333010-2022223001300013-3300303131101122-3221012310122132"></a>

## Next pages — blindfold_secret_info / 320133131322 / 7

- [https_management.advertise_on_slo_sli.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-2220121131203232-3202100013123232-2221330330202323-3131100122221013-3221230232121021-3133003111020322-0103003210222311-3230101011023000)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-0231010001133320-3230032322031013-1321332103323213-0120120300231233-1221100220101331-2011202010330002-1220120223303312-0132300132110233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002012001111003-0232233023303310-2033331131212120-1132022130332020-1323220323101322-0023002123210133-2103230333233011-2212221022113313"></a>

## https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info — clear_secret_info / 122302120032 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-2220121131203232-3202100013123232-2221330330202323-3131100122221013-3221230232121021-3133003111020322-0103003210222311-3230101011023000)
- https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info

<a id="canonical-0223110131031313-3121233223320320-1310233230301303-2212301110220013-2120201122333033-2322323210022020-2220100202302223-1122200130231300"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2100202222322010-2233000231130310-2033012330131031-0310133111313003-1330021232023211-2231103201121322-1013131131310221-2311000202221312"></a>

## Direct properties — clear_secret_info / 122302120032 / 3

<a id="canonical-3132220103031220-3020022000132030-2212330212120103-1321113013023323-1332002320122100-0201302300201110-1310130301323231-0123101300033121"></a>

<a id="canonical-3300102222112300-3130020313302202-3130333231230133-1012130312320000-0310131200203312-1123110310111320-2111011201002331-1132330321100310"></a>

## provider_ref property — clear_secret_info / 122302120032 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0212022103102203-1321011132010213-3201211200330313-3223303100001031-1303001322303232-1012303303002020-3321321202031131-2111113002033333"></a>

<a id="canonical-3300111121220211-0202010202312210-3123101121001210-3313112001310333-1133100211223220-3313223230000111-2021210033103202-1110033330110211"></a>

## URL property — clear_secret_info / 122302120032 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3003101110301001-1101121112201222-1102233331031122-2310011331010021-2230103310113200-3333202210022212-0301111220111330-1100333013330232"></a>

## Next pages — clear_secret_info / 122302120032 / 6

- [https_management.advertise_on_slo_sli.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-2220121131203232-3202100013123232-2221330330202323-3131100122221013-3221230232121021-3133003111020322-0103003210222311-3230101011023000)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1032013113221221-2330201313131201-2012232102321120-2021021033220122-0023000111021213-0131113010312113-1231110201003300-1333012213221332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230001122231232-2130213100003011-2210000230222231-3030111110230110-0332022231301103-0013230231233000-3303232312102333-1023123130302300"></a>

## https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults — use_system_defaults / 300233021302 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113)
- https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults

<a id="canonical-2200301201111133-2313022022013102-3113110020000000-2131032112121023-2222223210213302-1311210202322013-1130322332032001-0033303321223130"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

<a id="canonical-1022001221122021-3002031111232210-3332122321230332-3222230201003323-3320130211112012-3202110123132003-1121023312010112-3301323301233310"></a>

## Direct properties — use_system_defaults / 300233021302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332103333133023-2332113202232021-3021020110222113-2210112302201301-1013332021101332-1121200103300321-1000000311310022-3221132121221202"></a>

## Next pages — use_system_defaults / 300233021302 / 4

- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-3220330322200213-1131332232110011-2033122101021203-3133303310201321-1100333102323010-2210233023331021-2221332001303332-0331211032201031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000310100021320-2120233003033322-0100312100212320-2232000223300113-0332202233111013-3020230000013121-1132202303220010-1233120023030021"></a>

## https_management.advertise_on_slo_sli.tls_config — tls_config / 030003300203 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- https_management.advertise_on_slo_sli.tls_config

<a id="canonical-3031201110103220-2332033131020233-3100133000113300-0220032103110233-3333333321111132-0201113321230023-0232030123131331-3113333120110222"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312111002223303-3330002200330132-1202022011301222-1011200132003210-1001122022201031-3021132012221000-1121113331310210-3101230201300103"></a>

## Direct properties — tls_config / 030003300203 / 3

- [custom_security](resources--nfv_service--reference--group-003.md#canonical-3032330203113133-0300213012202320-2003232020002100-0203330112213222-2201131230311213-1021301332011022-3202111330233322-1003031130032302): complete subsection reference.

- [default_security](resources--nfv_service--reference--group-003.md#canonical-1002021221233202-3130021310012220-1310110003230101-3003323320213322-3310110222322330-3130121121120100-1201313002332131-3131200121123002): complete subsection reference.

- [low_security](resources--nfv_service--reference--group-003.md#canonical-1313030333021133-1013033023021223-0033130302313320-2212101202312033-0311010232220003-0113012233000221-2302032020012303-2033302131321302): complete subsection reference.

- [medium_security](resources--nfv_service--reference--group-003.md#canonical-2112302102010123-1033302111322023-1121001202132030-0321211301021213-1302021203010030-2111330211332132-2210102210033100-1011303231030233): complete subsection reference.

<a id="canonical-3002112233030023-1120302021203222-3102203320102231-0313232132003111-0212222312231110-1222211212331222-0333331232123213-3213320001013213"></a>

## Next pages — tls_config / 030003300203 / 4

- [https_management.advertise_on_slo_sli.tls_config.custom_security](resources--nfv_service--reference--group-003.md#canonical-3032330203113133-0300213012202320-2003232020002100-0203330112213222-2201131230311213-1021301332011022-3202111330233322-1003031130032302)
- [https_management.advertise_on_slo_sli.tls_config.default_security](resources--nfv_service--reference--group-003.md#canonical-1002021221233202-3130021310012220-1310110003230101-3003323320213322-3310110222322330-3130121121120100-1201313002332131-3131200121123002)
- [https_management.advertise_on_slo_sli.tls_config.low_security](resources--nfv_service--reference--group-003.md#canonical-1313030333021133-1013033023021223-0033130302313320-2212101202312033-0311010232220003-0113012233000221-2302032020012303-2033302131321302)
- [https_management.advertise_on_slo_sli.tls_config.medium_security](resources--nfv_service--reference--group-003.md#canonical-2112302102010123-1033302111322023-1121001202132030-0321211301021213-1302021203010030-2111330211332132-2210102210033100-1011303231030233)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-3032330203113133-0300213012202320-2003232020002100-0203330112213222-2201131230311213-1021301332011022-3202111330233322-1003031130032302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110223301201231-3220110210313212-0210011300023213-0100200113120023-2021033311223320-0130000130300303-2210302310202332-3023023311220230"></a>

## https_management.advertise_on_slo_sli.tls_config.custom_security — custom_security / 012111301330 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-3220330322200213-1131332232110011-2033122101021203-3133303310201321-1100333102323010-2210233023331021-2221332001303332-0331211032201031)
- https_management.advertise_on_slo_sli.tls_config.custom_security

<a id="canonical-3000120202122123-3002200103132220-3221200331221031-3213122123223111-0201231213103211-2223130310123132-2002210210302001-2300033213311313"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-1103122000331201-2333100323230000-2312033220121103-1220233122131123-1312202301202132-2220330233003013-2100301102112000-0300332311331012"></a>

## Direct properties — custom_security / 012111301330 / 3

<a id="canonical-1302021011020110-0121132011323321-0203323033133333-0102003022301201-2231021221111012-1030210210300110-3121213131013310-0320332220322223"></a>

<a id="canonical-2313130132033123-1112031111102031-1223212210311102-2130001131323301-3131211123300222-2322332131121112-1310000100002011-2101033212002231"></a>

## cipher_suites property — custom_security / 012111301330 / 4

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2023000203003103-0230132300331311-3001013112300212-3110031132030132-2223123101301031-3221220223300210-1210013303311131-3131121331103132"></a>

<a id="canonical-3013022033331123-1121222302202110-0030220231330210-2313023201223202-0001213210233020-0132331303023033-2021210231100323-0203322321331213"></a>

## max_version property — custom_security / 012111301330 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2331130001230231-0001212313112222-0302001110322100-1200121003210133-1313133012113233-0303313010133222-3122221313111200-3000320023131321"></a>

<a id="canonical-1223103021331322-1301201000131230-0001131101113212-3013112201302303-1011123333203223-1221323311213310-1300121113100000-1211311002121330"></a>

## min_version property — custom_security / 012111301330 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0323001330023221-0210222103232103-3012112122031322-2003002101110202-1131203321322300-1333323021003312-3122221003121032-3010101313001322"></a>

## Next pages — custom_security / 012111301330 / 7

- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-3220330322200213-1131332232110011-2033122101021203-3133303310201321-1100333102323010-2210233023331021-2221332001303332-0331211032201031)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1002021221233202-3130021310012220-1310110003230101-3003323320213322-3310110222322330-3130121121120100-1201313002332131-3131200121123002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012301013000101-1322323113310332-2201302333311201-3110022201232122-2012030311103230-2220112213233310-3131020233033233-3003100101132003"></a>

## https_management.advertise_on_slo_sli.tls_config.default_security — default_security / 310213323330 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-3220330322200213-1131332232110011-2033122101021203-3133303310201321-1100333102323010-2210233023331021-2221332001303332-0331211032201031)
- https_management.advertise_on_slo_sli.tls_config.default_security

<a id="canonical-2211111111233321-0331013132321302-3320211030003020-1223022020000203-1100120100000120-0123112203222110-0332211033110130-2103220123030321"></a>

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
default_security = {}
```

<a id="canonical-1001202323002010-2323031131233312-2233033300220103-2011000031013221-0123330223211333-0130233310330233-2100210212111203-0122203220123101"></a>

## Direct properties — default_security / 310213323330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2202022231223033-1000222221310022-2221330212112122-0301232222311323-3220321333200333-0121120232010233-3302122122012330-2313310032011323"></a>

## Next pages — default_security / 310213323330 / 4

- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-3220330322200213-1131332232110011-2033122101021203-3133303310201321-1100333102323010-2210233023331021-2221332001303332-0331211032201031)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1313030333021133-1013033023021223-0033130302313320-2212101202312033-0311010232220003-0113012233000221-2302032020012303-2033302131321302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331110201313311-2320030233112202-0023101133010230-3303110120112103-0021130230302123-0232331302313012-2131133202300220-2201233013101221"></a>

## https_management.advertise_on_slo_sli.tls_config.low_security — low_security / 222320010030 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-3220330322200213-1131332232110011-2033122101021203-3133303310201321-1100333102323010-2210233023331021-2221332001303332-0331211032201031)
- https_management.advertise_on_slo_sli.tls_config.low_security

<a id="canonical-3011310322301103-3000300013103011-3101221220033103-3223112202023023-3322113113133100-0211312112220302-0131210110130212-1033311012332032"></a>

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
low_security = {}
```

<a id="canonical-3300012121110312-0021033122330313-2020031010313023-3302102222121000-1212112302011213-3120311210023112-2303121311300012-1031122202120103"></a>

## Direct properties — low_security / 222320010030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313130313330222-0313132100202203-2212021100321130-1201311320120300-1103131022130313-1023321101012203-3331323030121310-0120013121201310"></a>

## Next pages — low_security / 222320010030 / 4

- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-3220330322200213-1131332232110011-2033122101021203-3133303310201321-1100333102323010-2210233023331021-2221332001303332-0331211032201031)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-2112302102010123-1033302111322023-1121001202132030-0321211301021213-1302021203010030-2111330211332132-2210102210033100-1011303231030233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303220311102201-2111001032302121-2103022100212000-0200313232120201-3130203330331302-1230101002313030-2212302300210130-0320300112220013"></a>

## https_management.advertise_on_slo_sli.tls_config.medium_security — medium_security / 011312203203 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-3220330322200213-1131332232110011-2033122101021203-3133303310201321-1100333102323010-2210233023331021-2221332001303332-0331211032201031)
- https_management.advertise_on_slo_sli.tls_config.medium_security

<a id="canonical-2022100312330102-2103131100103111-2202030202121132-2303013301020101-1311101102203012-2122011310111232-1013212131110203-3332011332302023"></a>

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
medium_security = {}
```

<a id="canonical-0132003013201331-2021301101110320-3123023212310320-0010203012231120-3120313320031220-0320321323210213-3001312322120011-3103111122321032"></a>

## Direct properties — medium_security / 011312203203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212020100210013-1213331200122311-0020203102310220-1201223233311332-2010221233333200-2133022320211311-0101313213333033-3311122000003221"></a>

## Next pages — medium_security / 011312203203 / 4

- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-3220330322200213-1131332232110011-2033122101021203-3133303310201321-1100333102323010-2210233023331021-2221332001303332-0331211032201031)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313021011221113-2000020100101001-1120102013020120-0113031332303233-0221311303031321-3110213113012031-2311323130102333-1101312012111233"></a>

## https_management.advertise_on_slo_sli.use_mtls — use_mtls / 100300012110 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- https_management.advertise_on_slo_sli.use_mtls

<a id="canonical-0101002012210331-3130133103321132-0122211223123002-3131001203213102-0123232111233223-0032233021132033-1101330321232312-3321311203121303"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322112323020322-1122200000132000-2022330000310332-1111210030203122-3201300310133233-0302112201221003-3001133020312211-3233321233321200"></a>

## Direct properties — use_mtls / 100300012110 / 3

<a id="canonical-1202210002310331-1110132200031013-2021320321222332-1300301232123000-0022022022230212-0311032010312010-0003130130122203-3231223002012032"></a>

<a id="canonical-1132233032312223-2212321323002021-2212121332030222-3023230023123132-2022331201323220-3022312033232220-2032232213202012-1001002102031120"></a>

## client_certificate_optional property — use_mtls / 100300012110 / 4

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](resources--nfv_service--reference--group-003.md#canonical-2122323020322313-3112221322132323-3230213310332002-2122231223123100-1213323123202002-2122210023333030-3210132311133003-1010300231330103): complete subsection reference.

- [no_crl](resources--nfv_service--reference--group-003.md#canonical-1203023000000003-2302301231310103-2023100211121011-0201303020031000-1022023201333321-1220102231102200-3131222002220233-1133322300102301): complete subsection reference.

- [trusted_ca](resources--nfv_service--reference--group-003.md#canonical-0300022102301312-1131213301311333-1313011200032130-1033221120132133-1331001133010231-2223002012321110-2200122122311223-1131233322333203): complete subsection reference.

<a id="canonical-0300102112310320-0012011321030133-2103130122102211-1002202332202301-2121321330312133-3323213030132013-0221200312122120-3010221100103013"></a>

<a id="canonical-1010303301333120-1010023200200023-3221112323032201-1221131223003131-1102123002330111-2012321302031320-1220313231331031-2101133123031110"></a>

## trusted_ca_url property — use_mtls / 100300012110 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Provider validators and defaults (from schema source):

```go
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

- [xfcc_disabled](resources--nfv_service--reference--group-003.md#canonical-3303321210232100-3033322310101121-1131030131110032-2003321300131230-3320222310111123-1320012022120011-0212313223103013-3130123033102210): complete subsection reference.

- [xfcc_options](resources--nfv_service--reference--group-003.md#canonical-2202121020010233-0110202113211132-2332301321323120-0021021112223301-3210202210232031-3233002032120033-0222202210211133-1010223320333201): complete subsection reference.

<a id="canonical-1013311332233123-0302023111133001-1022211022313301-3030120332112201-2022013130131121-1023033300312113-0011231133302220-3102233122203231"></a>

## Next pages — use_mtls / 100300012110 / 6

- [https_management.advertise_on_slo_sli.use_mtls.crl](resources--nfv_service--reference--group-003.md#canonical-2122323020322313-3112221322132323-3230213310332002-2122231223123100-1213323123202002-2122210023333030-3210132311133003-1010300231330103)
- [https_management.advertise_on_slo_sli.use_mtls.no_crl](resources--nfv_service--reference--group-003.md#canonical-1203023000000003-2302301231310103-2023100211121011-0201303020031000-1022023201333321-1220102231102200-3131222002220233-1133322300102301)
- [https_management.advertise_on_slo_sli.use_mtls.trusted_ca](resources--nfv_service--reference--group-003.md#canonical-0300022102301312-1131213301311333-1313011200032130-1033221120132133-1331001133010231-2223002012321110-2200122122311223-1131233322333203)
- [https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled](resources--nfv_service--reference--group-003.md#canonical-3303321210232100-3033322310101121-1131030131110032-2003321300131230-3320222310111123-1320012022120011-0212313223103013-3130123033102210)
- [https_management.advertise_on_slo_sli.use_mtls.xfcc_options](resources--nfv_service--reference--group-003.md#canonical-2202121020010233-0110202113211132-2332301321323120-0021021112223301-3210202210232031-3233002032120033-0222202210211133-1010223320333201)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-2122323020322313-3112221322132323-3230213310332002-2122231223123100-1213323123202002-2122210023333030-3210132311133003-1010300231330103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110001302132310-0012110113200212-1130121202210320-0100222011210230-0031111113111213-0302111333012003-0212032000020222-3221001022002021"></a>

## https_management.advertise_on_slo_sli.use_mtls.crl — crl / 103022301002 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310)
- https_management.advertise_on_slo_sli.use_mtls.crl

<a id="canonical-2102223303121033-3003323311233023-1310300103003310-2311001233230132-0333231323232111-1102111212221010-1130033012313322-2222002311022231"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300103032130313-1202021333310112-2120330330111231-3000230012333101-0232301112123300-3321200321001231-3113331303203122-1133313203201033"></a>

## Direct properties — crl / 103022301002 / 3

<a id="canonical-3011022203222132-1132111132033011-1130022201303223-2331312302203300-2313023311231233-2320012223203120-1331203130111020-2300131302121003"></a>

<a id="canonical-3022022010201310-1332111333330233-0023113312311000-1101030112311321-2001113102030233-3123213103120230-2031002131022032-2020033100311011"></a>

## name property — crl / 103022301002 / 4

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

<a id="canonical-2230033001232033-2120033231223231-3332201020213230-3332311002030110-3002310202032301-3332100332320202-0133302222320203-3330031322301103"></a>

<a id="canonical-1332122112111201-2302221132321131-0103132110203311-2131322103230302-1122110003021322-1323013100231020-2012202202313123-2132300100100313"></a>

## namespace property — crl / 103022301002 / 5

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

<a id="canonical-0332303302220003-0332222112201002-2310131133010221-2302103302302223-2301223300013003-3233013122002010-2221310212032123-2103302300030311"></a>

<a id="canonical-0320121333233013-0131100221320132-1331320201013020-3121102130013133-3212103303100322-2033231103200232-3133333123002012-1101112323102010"></a>

## tenant property — crl / 103022301002 / 6

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

<a id="canonical-0032311031301122-2222313031222101-1103301111200300-1200022220111100-0121001213212310-0202210133013302-1102013301230232-1012031032323330"></a>

## Next pages — crl / 103022301002 / 7

- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1203023000000003-2302301231310103-2023100211121011-0201303020031000-1022023201333321-1220102231102200-3131222002220233-1133322300102301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220002330003033-2113122233323103-3302000013013330-3202132030022301-1322001033222313-0102310200131030-2002200130111211-3132310122231321"></a>

## https_management.advertise_on_slo_sli.use_mtls.no_crl — no_crl / 232022103331 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310)
- https_management.advertise_on_slo_sli.use_mtls.no_crl

<a id="canonical-0030012230222331-2002311130032222-0011012220130033-1030323331011220-0223000113223232-2231303130322210-3201213101333312-2220132130033011"></a>

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

<a id="canonical-2023022111133201-3002223210332201-1131232333133211-0003301230302321-1021333001113111-1231313110011012-2022211000003310-3230013322220003"></a>

## Direct properties — no_crl / 232022103331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323011322112313-1313022012032201-2202300122230101-1320221112311312-3202321332023311-2213310302313221-2333310101303313-0020011203323002"></a>

## Next pages — no_crl / 232022103331 / 4

- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-0300022102301312-1131213301311333-1313011200032130-1033221120132133-1331001133010231-2223002012321110-2200122122311223-1131233322333203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020231112213130-0200033230010112-3331213012010213-1321031022230023-0120200230011123-3121222002113200-3103210200020001-1132322212013213"></a>

## https_management.advertise_on_slo_sli.use_mtls.trusted_ca — trusted_ca / 130130101031 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310)
- https_management.advertise_on_slo_sli.use_mtls.trusted_ca

<a id="canonical-3131023010230002-0231002203322301-1132012311333303-1310010312023212-0320131120220230-2221212300110100-3011231323110122-1312223033212110"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-2132112333030201-1302301021213313-2221321200212130-0100021001123031-1022102130332100-1333313032013233-2001222200323000-3310222330001112"></a>

## Direct properties — trusted_ca / 130130101031 / 3

<a id="canonical-1301301030001303-0020323202023312-1313333113033231-2110323101021310-3022100012321201-1322201303111133-1231030012311131-0123301111030233"></a>

<a id="canonical-3021323002022333-2332320311133200-3230311121030322-0012203202120211-0310321212222000-3100301310203003-3131122220000211-1133022031131211"></a>

## name property — trusted_ca / 130130101031 / 4

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

<a id="canonical-1302130320313333-2012313032001211-1200221321210311-2102031331103321-0213303110013013-1203310110112031-3220120122320332-3312200113132232"></a>

<a id="canonical-0331010010031130-0123111110101330-1132113122123222-1322022103023123-1021222103032311-2132031033302300-1012102313103130-2121201030023000"></a>

## namespace property — trusted_ca / 130130101031 / 5

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

<a id="canonical-0221210212212310-2103332213201110-3032201023222313-3030120103322020-2110210300201123-0013123203310002-0003223203322030-0111101023331002"></a>

<a id="canonical-2320020321032322-3300000321103102-0120001222311022-1232312123103103-1323302302310313-3010231121121333-2231231011131121-0212101022330133"></a>

## tenant property — trusted_ca / 130130101031 / 6

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

<a id="canonical-1011022302113202-0333321210033133-0222300201000023-2022223330013102-3222013201301022-0231131201223301-1033102002213001-3023001311022002"></a>

## Next pages — trusted_ca / 130130101031 / 7

- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-3303321210232100-3033322310101121-1131030131110032-2003321300131230-3320222310111123-1320012022120011-0212313223103013-3130123033102210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011230303132110-0213003301121113-2200133333010213-0111233331301221-0201011020133303-2323133123330022-1211122032321220-3322110311200130"></a>

## https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled — xfcc_disabled / 303312301233 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310)
- https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled

<a id="canonical-0230021031001003-3311031210100201-2100011213310322-2032231013301012-3033303313211333-0113110111312300-1030312301333101-1132202020211222"></a>

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

<a id="canonical-2300100020320212-0100231231212202-3133310111120332-1121010323120310-2122301220013301-2102331113032032-3211132313212113-2111210133102311"></a>

## Direct properties — xfcc_disabled / 303312301233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2310103323330321-2021112010210131-1122131310120221-3221020213013132-3312023103023023-2212123102313030-3003323132200103-1313022132311032"></a>

## Next pages — xfcc_disabled / 303312301233 / 4

- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-2202121020010233-0110202113211132-2332301321323120-0021021112223301-3210202210232031-3233002032120033-0222202210211133-1010223320333201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023131202322322-1301002131021113-2102033203011133-1020103123130332-3212200013233130-1112101212223210-3133302322222123-2001122333321020"></a>

## https_management.advertise_on_slo_sli.use_mtls.xfcc_options — xfcc_options / 110002323200 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310)
- https_management.advertise_on_slo_sli.use_mtls.xfcc_options

<a id="canonical-0320223201130322-1301030333320222-2120313130300112-2112210121030100-0321020002131301-1311112010311130-3323330122211021-2111211212202230"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3230312212330213-0023010211223202-1332223322211212-0310131133030030-3123311010000020-2203012100333100-2223300112300303-2120110033220311"></a>

## Direct properties — xfcc_options / 110002323200 / 3

<a id="canonical-1232002303002201-3220113203023223-3122213121331031-3021320312302003-3121131223011212-1222310221312021-1103200021011313-1201031011203021"></a>

<a id="canonical-2012313113320210-2123132131200310-2111221312121210-1013130210300101-3303211000202302-1311200011102020-0203113122021232-0023201010330113"></a>

## xfcc_header_elements property — xfcc_options / 110002323200 / 4

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

<a id="canonical-3300212113000022-3102120302203120-3330322323010222-1300203320130023-1022300120213122-3332313201123110-2330333223001112-2010030331323032"></a>

## Next pages — xfcc_options / 110002323200 / 5

- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222331132031103-2131232030022312-1231313020022313-2103023310300201-2320011203101302-0222032022222002-0221310333003020-3111230300030221"></a>

## https_management.advertise_on_slo_vip — advertise_on_slo_vip / 332203011323 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- https_management.advertise_on_slo_vip

<a id="canonical-0311030123113020-0310323023030223-2132032033103233-3000210331321312-1131200021222000-1232022213100001-1221331332332222-3313121002001232"></a>

Type: `"object"`. single nested block, Optional.

Inline TLS Parameters. Inline TLS parameters.

Upstream description:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
advertise_on_slo_vip {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220222203321331-3113000101230231-2331333133210033-1110033201212232-2130222130012331-3302003322330113-3300130300320211-1003230213233102"></a>

## Direct properties — advertise_on_slo_vip / 332203011323 / 3

- [no_mtls](resources--nfv_service--reference--group-003.md#canonical-0211010101302002-1033333201312123-0303321221111102-0233212330023110-1221033301032313-3130023100303333-1320012331023231-3010321120110301): complete subsection reference.

- [tls_certificates](resources--nfv_service--reference--group-003.md#canonical-0201100000120203-2021020301330032-0310123232303233-0102332130311132-1233200003321230-1112232313301301-3210030321330101-0101132333333033): complete subsection reference.

- [tls_config](resources--nfv_service--reference--group-003.md#canonical-1220000301313030-3210123130033303-0000003031222101-0230222200021002-2212330020213020-3221320323031031-1203232101302221-3230202011332313): complete subsection reference.

- [use_mtls](resources--nfv_service--reference--group-003.md#canonical-0013001011312211-0302330103310232-2330013030103010-2111232121111121-1132202200220233-3303310322202002-2110202013220100-2223302323332223): complete subsection reference.

<a id="canonical-1231132211310031-2330302132332102-1011213023213213-2213213002222301-2133302203102113-0300101323320213-1000333233131122-2331210121022130"></a>

## Next pages — advertise_on_slo_vip / 332203011323 / 4

- [https_management.advertise_on_slo_vip.no_mtls](resources--nfv_service--reference--group-003.md#canonical-0211010101302002-1033333201312123-0303321221111102-0233212330023110-1221033301032313-3130023100303333-1320012331023231-3010321120110301)
- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-0201100000120203-2021020301330032-0310123232303233-0102332130311132-1233200003321230-1112232313301301-3210030321330101-0101132333333033)
- [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-1220000301313030-3210123130033303-0000003031222101-0230222200021002-2212330020213020-3221320323031031-1203232101302221-3230202011332313)
- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0013001011312211-0302330103310232-2330013030103010-2111232121111121-1132202200220233-3303310322202002-2110202013220100-2223302323332223)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-0211010101302002-1033333201312123-0303321221111102-0233212330023110-1221033301032313-3130023100303333-1320012331023231-3010321120110301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130113001221023-2321000122222220-3330111331122010-1112321320111232-1213322001322010-0133331321110000-3233231120222213-0221310300102311"></a>

## https_management.advertise_on_slo_vip.no_mtls — no_mtls / 220101200222 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- https_management.advertise_on_slo_vip.no_mtls

<a id="canonical-0223102102213323-3120131112012331-2011112320231311-1021011120131321-0223330102011211-1303132001130100-2120020021323301-0033120103310330"></a>

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
no_mtls = {}
```

<a id="canonical-3332320303133330-2020300120312331-0012323102221110-3330221331230221-0322210321313031-2131002302330112-1111122233023021-3323012032103021"></a>

## Direct properties — no_mtls / 220101200222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323110122320200-0130030102313201-3310212333121332-1213103012332212-1203321032331200-2313212130023313-0333220012201223-2121213131122311"></a>

## Next pages — no_mtls / 220101200222 / 4

- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-0201100000120203-2021020301330032-0310123232303233-0102332130311132-1233200003321230-1112232313301301-3210030321330101-0101132333333033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333231212312301-3201130222232021-1032001021111200-0020030203103113-2302212120300311-2302032011010202-1121203313330312-1033333310030102"></a>

## https_management.advertise_on_slo_vip.tls_certificates — tls_certificates / 000132102201 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- https_management.advertise_on_slo_vip.tls_certificates

<a id="canonical-2033311030122032-3110102122003230-0011002203212233-1032211121000322-2122131133333230-0200313213020211-1202331022100100-3301023220233000"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323310303312200-3103213120332101-3230233013330132-3311020302222133-2210112032032211-3213311003101120-0102211012123032-2132020103133311"></a>

## Direct properties — tls_certificates / 000132102201 / 3

<a id="canonical-1120032033333030-3203101020110111-2300121111222022-0213213231333133-3222331220330300-2303112003300321-1310021131201312-1200123211122223"></a>

<a id="canonical-0213130210012320-0011010121301002-0313312202222232-3230021320100103-2022311001111120-0323001021301112-2103013010312023-2013123321302101"></a>

## certificate_url property — tls_certificates / 000132102201 / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](resources--nfv_service--reference--group-003.md#canonical-3332021021113313-0310032003132330-0222331210123333-3322003030331020-0210101130020113-1321321012202011-1101101010333033-1032011210212011): complete subsection reference.

<a id="canonical-1301110230122310-3000311233031332-3121111031002000-3010230001301101-1101203131321121-0213110012203131-2323323311220321-3023322223020220"></a>

<a id="canonical-1103112031200213-2002031220311322-1121103301112231-1033122112110200-1030333200020220-1011012001002323-1332131303310232-3223101030310133"></a>

## description_spec property — tls_certificates / 000132102201 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--nfv_service--reference--group-003.md#canonical-1223312232131130-3323102031233232-3103230230232313-0023320033320202-0210000233222211-3012010210210210-2201000102130223-2220232000002001): complete subsection reference.

- [private_key](resources--nfv_service--reference--group-003.md#canonical-0022012123311000-1212203230321020-1330231021210332-3201020233133002-2323100303020232-0312203302233123-3000300312202110-2223011033132013): complete subsection reference.

- [use_system_defaults](resources--nfv_service--reference--group-003.md#canonical-1011101211010132-3312322032312331-3113201301221030-0122312102211300-0221000121010231-3102333110002230-2312323213001223-1111002232331221): complete subsection reference.

<a id="canonical-1003221021103033-3023312122122230-0123303131003010-2102330231210020-1210233321020230-2023321110111103-2002220213311202-0323233302322311"></a>

## Next pages — tls_certificates / 000132102201 / 6

- [https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms](resources--nfv_service--reference--group-003.md#canonical-3332021021113313-0310032003132330-0222331210123333-3322003030331020-0210101130020113-1321321012202011-1101101010333033-1032011210212011)
- [https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling](resources--nfv_service--reference--group-003.md#canonical-1223312232131130-3323102031233232-3103230230232313-0023320033320202-0210000233222211-3012010210210210-2201000102130223-2220232000002001)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-0022012123311000-1212203230321020-1330231021210332-3201020233133002-2323100303020232-0312203302233123-3000300312202110-2223011033132013)
- [https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults](resources--nfv_service--reference--group-003.md#canonical-1011101211010132-3312322032312331-3113201301221030-0122312102211300-0221000121010231-3102333110002230-2312323213001223-1111002232331221)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-3332021021113313-0310032003132330-0222331210123333-3322003030331020-0210101130020113-1321321012202011-1101101010333033-1032011210212011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002010233311010-3321233120122021-2332301231010320-2101233211212033-2301012330033133-3002022212020023-0213122020023203-0312012021333210"></a>

## https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 123211002233 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-0201100000120203-2021020301330032-0310123232303233-0102332130311132-1233200003321230-1112232313301301-3210030321330101-0101132333333033)
- https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms

<a id="canonical-1123112301132312-1301222223010313-1013321112310331-2130133001312332-2332222000032303-1301231010301310-3120203211332011-2231222100321310"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-2221303300131032-0133223001103301-0333310012122231-0000330223333203-2321100022312331-0323103301001001-3000323231220330-1333131001202133"></a>

## Direct properties — custom_hash_algorithms / 123211002233 / 3

<a id="canonical-3213230122223201-0300123032111130-1022332131222302-0101131123311331-1302131210333213-2031331003312132-1220010222210302-1032110020322333"></a>

<a id="canonical-0131322002131212-0312030000100012-3113111113111010-1210020200132233-0232221032221313-0202033311132031-2023132032230200-2002133333223030"></a>

## hash_algorithms property — custom_hash_algorithms / 123211002233 / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2310110200112033-1232123102310212-1221223132321001-2200122132022020-0013223121022101-2313023030231030-3021013212230313-2123102030120103"></a>

## Next pages — custom_hash_algorithms / 123211002233 / 5

- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-0201100000120203-2021020301330032-0310123232303233-0102332130311132-1233200003321230-1112232313301301-3210030321330101-0101132333333033)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1223312232131130-3323102031233232-3103230230232313-0023320033320202-0210000233222211-3012010210210210-2201000102130223-2220232000002001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321110210203101-3322012120020031-1023000032210201-1311213213003122-1301111120131131-1110323312310111-1323331312030001-2232123023202113"></a>

## https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 121113101110 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-0201100000120203-2021020301330032-0310123232303233-0102332130311132-1233200003321230-1112232313301301-3210030321330101-0101132333333033)
- https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling

<a id="canonical-0221121130213003-0221302100231011-1332330232032013-0000321233021313-2212103103133210-1013212211333213-1330112012111233-0323300330223112"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

<a id="canonical-0301201220203231-0313012300112220-0330003220131320-3233130213300203-0322221313111301-1323121012210333-2221113001333111-1203212212301303"></a>

## Direct properties — disable_ocsp_stapling / 121113101110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2311301213010130-2302301222020130-3123110210132322-0331303303021321-1120013132321313-1012102333310302-3001223320333221-1312223311311021"></a>

## Next pages — disable_ocsp_stapling / 121113101110 / 4

- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-0201100000120203-2021020301330032-0310123232303233-0102332130311132-1233200003321230-1112232313301301-3210030321330101-0101132333333033)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-0022012123311000-1212203230321020-1330231021210332-3201020233133002-2323100303020232-0312203302233123-3000300312202110-2223011033132013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010123322220121-0000032222010111-2213122230231332-1131032110103212-1231320011210302-3020011210023211-0323313200320123-0033332301023002"></a>

## https_management.advertise_on_slo_vip.tls_certificates.private_key — private_key / 102202032223 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-0201100000120203-2021020301330032-0310123232303233-0102332130311132-1233200003321230-1112232313301301-3210030321330101-0101132333333033)
- https_management.advertise_on_slo_vip.tls_certificates.private_key

<a id="canonical-2123010121130221-0131130220001200-3011012221222210-1103000020112311-0111030031330223-0332201202131330-1311331301232212-1201221111331100"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
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
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023022111321112-2333312201222321-3102330033000322-2033222223203313-3031313103123212-3110022132002031-2003003210300023-2021302110010102"></a>

## Direct properties — private_key / 102202032223 / 3

- [blindfold_secret_info](resources--nfv_service--reference--group-003.md#canonical-0303233320112200-3220033130320200-3223032232212000-2132032213123332-1012210220020200-3102211222300113-0212030203120101-3331102210022210): complete subsection reference.

- [clear_secret_info](resources--nfv_service--reference--group-003.md#canonical-2112031210010210-1111112120300332-0010203003220232-1033332002002210-1311220031111310-0113102010002113-0000010300102300-1001023330320103): complete subsection reference.

<a id="canonical-3012331010030302-1030231222221123-1103010132330101-0131300130313033-2320232020213322-1211212223320211-3131202322200312-3130120132303321"></a>

## Next pages — private_key / 102202032223 / 4

- [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--reference--group-003.md#canonical-0303233320112200-3220033130320200-3223032232212000-2132032213123332-1012210220020200-3102211222300113-0212030203120101-3331102210022210)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info](resources--nfv_service--reference--group-003.md#canonical-2112031210010210-1111112120300332-0010203003220232-1033332002002210-1311220031111310-0113102010002113-0000010300102300-1001023330320103)
- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-0201100000120203-2021020301330032-0310123232303233-0102332130311132-1233200003321230-1112232313301301-3210030321330101-0101132333333033)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-0303233320112200-3220033130320200-3223032232212000-2132032213123332-1012210220020200-3102211222300113-0212030203120101-3331102210022210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133003230130220-0201323000000112-0311331322132322-3122220200100221-3231111322302330-1012331311202333-3122231331113311-3222202022033000"></a>

## https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 313221132020 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-0201100000120203-2021020301330032-0310123232303233-0102332130311132-1233200003321230-1112232313301301-3210030321330101-0101132333333033)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-0022012123311000-1212203230321020-1330231021210332-3201020233133002-2323100303020232-0312203302233123-3000300312202110-2223011033132013)
- https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-1331111321322033-3212131212031310-2231023021332322-3333003101202020-1330132302130321-1120321003203012-2230023111111213-3203033111022323"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0320121221000202-1100232230231113-3303121220030000-2213223321103231-1311032313131202-0002103011030023-1113111332131031-0302120021302222"></a>

## Direct properties — blindfold_secret_info / 313221132020 / 3

<a id="canonical-3211033203210233-3021100103321211-0113213303233220-2003022202221211-0131122331222331-0001221303201332-1211021233013131-0331103213300132"></a>

<a id="canonical-3112030131233102-3221220023200222-2031202232322221-1321202201212233-0223202131200112-3121230013232233-0031233122002121-2101033112223232"></a>

## decryption_provider property — blindfold_secret_info / 313221132020 / 4

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

<a id="canonical-1132213331011100-0301000132302203-0023301313312020-1213300103201230-0330321011302011-0132002122101302-2301212231130310-0123001202102303"></a>

<a id="canonical-2033220300130022-2233023031023232-0223322210110322-2331013002123120-1131201321023300-2230131120031012-2212310330320102-3321312121103023"></a>

## location property — blindfold_secret_info / 313221132020 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0330221101110201-1233332133032021-1133101120321202-2210102203220330-2123132220100123-1213332132202002-2013002223331020-3232100021301030"></a>

<a id="canonical-1022100301233023-1000010312103313-2312131233022202-3202330120010013-2001222101023012-0220121030302103-3020232003033112-0200203232300101"></a>

## store_provider property — blindfold_secret_info / 313221132020 / 6

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

<a id="canonical-3333303312222031-1312221331133001-2220133221323221-0123321233203320-2023020123311023-2322300022001133-2001031001033023-2211212101122321"></a>

## Next pages — blindfold_secret_info / 313221132020 / 7

- [https_management.advertise_on_slo_vip.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-0022012123311000-1212203230321020-1330231021210332-3201020233133002-2323100303020232-0312203302233123-3000300312202110-2223011033132013)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-2112031210010210-1111112120300332-0010203003220232-1033332002002210-1311220031111310-0113102010002113-0000010300102300-1001023330320103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120322021200023-3101031211331222-3111233122311033-3111223223000112-0213312222213102-2331322133212122-2001211222112110-1023111231203211"></a>

## https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info — clear_secret_info / 332023132020 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-0201100000120203-2021020301330032-0310123232303233-0102332130311132-1233200003321230-1112232313301301-3210030321330101-0101132333333033)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-0022012123311000-1212203230321020-1330231021210332-3201020233133002-2323100303020232-0312203302233123-3000300312202110-2223011033132013)
- https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info

<a id="canonical-3112202333130320-2010030322332211-3130233201012322-1221302032101201-2202113231331232-1101111313303330-3321210023100121-3320231101000232"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1100222320030220-1311233300313222-1133323220210133-3133221211123001-1203131133232220-3000103030323202-0100102113313202-3223120212220211"></a>

## Direct properties — clear_secret_info / 332023132020 / 3

<a id="canonical-2330310221230030-2033003003301231-3220231311230301-1221311302303332-3010131220222321-3123321302330211-2023320220201312-3330012011300301"></a>

<a id="canonical-3233011232100031-1112102100320111-1030001122031131-2300331112123021-1313311220120012-1310303120223203-3312322030003110-3331331310310120"></a>

## provider_ref property — clear_secret_info / 332023132020 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1330000113003310-1221211232022200-3332313231310121-2133222211303223-1220312220022113-1201230200203233-2022301023033230-1311032121212121"></a>

<a id="canonical-2102223201031212-2323301012112320-2031220001100012-3013020313203033-2001220330010332-1020212130110122-0213101323002012-3231133012232301"></a>

## URL property — clear_secret_info / 332023132020 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3202131311102321-0120333321301332-0010020032302002-3303030233321212-2111312233100312-2002301320212130-3202220102122233-0323130011320303"></a>

## Next pages — clear_secret_info / 332023132020 / 6

- [https_management.advertise_on_slo_vip.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-0022012123311000-1212203230321020-1330231021210332-3201020233133002-2323100303020232-0312203302233123-3000300312202110-2223011033132013)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1011101211010132-3312322032312331-3113201301221030-0122312102211300-0221000121010231-3102333110002230-2312323213001223-1111002232331221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033110331233320-0231231230102022-2112000132000232-2011302223103222-2212231301313003-2121232312103122-0032322313130230-0211230323033023"></a>

## https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults — use_system_defaults / 013222000223 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-0201100000120203-2021020301330032-0310123232303233-0102332130311132-1233200003321230-1112232313301301-3210030321330101-0101132333333033)
- https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults

<a id="canonical-2031012322002012-2010113222110322-2201220223000202-2103121223311322-0010210210332020-0210210031002211-1021230010011103-0021120103201102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

<a id="canonical-3021331333333221-0022332001112312-2033123032112131-3122123230333212-1121101313200132-3230321021003300-0111331130011010-3322102002101020"></a>

## Direct properties — use_system_defaults / 013222000223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3100120230103002-3001330022302210-3313321322323213-3330311001223322-1310123011331313-1320102233022332-1311133200113230-0123013201131223"></a>

## Next pages — use_system_defaults / 013222000223 / 4

- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-0201100000120203-2021020301330032-0310123232303233-0102332130311132-1233200003321230-1112232313301301-3210030321330101-0101132333333033)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1220000301313030-3210123130033303-0000003031222101-0230222200021002-2212330020213020-3221320323031031-1203232101302221-3230202011332313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333300212023012-1131131113222132-2323230211321233-3222301210201133-0332323031310021-3220111023102303-0032133333200032-0020201022303100"></a>

## https_management.advertise_on_slo_vip.tls_config — tls_config / 222313303130 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- https_management.advertise_on_slo_vip.tls_config

<a id="canonical-1313201200001300-0122302203303021-1011222313013121-3222212101023200-2202310100223220-2320332211322110-1030330332233120-0000220030000031"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-2123101301211201-3212202233003021-2000212232303231-1032032112021211-3323101302130010-3102112121210103-3322201202131003-1321332013322332"></a>

## Direct properties — tls_config / 222313303130 / 3

- [custom_security](resources--nfv_service--reference--group-003.md#canonical-3201031332003023-3033323011221223-0102102020000213-1300103033102211-2110032320231012-1303123231300102-0332113101133130-1210232003222333): complete subsection reference.

- [default_security](resources--nfv_service--reference--group-003.md#canonical-1330131011213232-0312111203313122-2030100023332311-2111021013121100-2202330011021130-3333030133113322-1111232220133201-0121203113031103): complete subsection reference.

- [low_security](resources--nfv_service--reference--group-003.md#canonical-0221112003220310-3222311333330300-3303223310301220-1032022013201000-2103003200212201-2133333333032331-0232313212202133-0021123223003122): complete subsection reference.

- [medium_security](resources--nfv_service--reference--group-003.md#canonical-3131221302121031-0301003201212113-0121002010311301-2233331000033101-2332322333202130-1223331020103211-2233033313330003-1212230311012203): complete subsection reference.

<a id="canonical-0323203103230111-3200030103123100-0230022230313322-0322011223232312-0330100211031113-3221311230221230-2230333321230111-2110321030032310"></a>

## Next pages — tls_config / 222313303130 / 4

- [https_management.advertise_on_slo_vip.tls_config.custom_security](resources--nfv_service--reference--group-003.md#canonical-3201031332003023-3033323011221223-0102102020000213-1300103033102211-2110032320231012-1303123231300102-0332113101133130-1210232003222333)
- [https_management.advertise_on_slo_vip.tls_config.default_security](resources--nfv_service--reference--group-003.md#canonical-1330131011213232-0312111203313122-2030100023332311-2111021013121100-2202330011021130-3333030133113322-1111232220133201-0121203113031103)
- [https_management.advertise_on_slo_vip.tls_config.low_security](resources--nfv_service--reference--group-003.md#canonical-0221112003220310-3222311333330300-3303223310301220-1032022013201000-2103003200212201-2133333333032331-0232313212202133-0021123223003122)
- [https_management.advertise_on_slo_vip.tls_config.medium_security](resources--nfv_service--reference--group-003.md#canonical-3131221302121031-0301003201212113-0121002010311301-2233331000033101-2332322333202130-1223331020103211-2233033313330003-1212230311012203)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-3201031332003023-3033323011221223-0102102020000213-1300103033102211-2110032320231012-1303123231300102-0332113101133130-1210232003222333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322030121003231-2201312320302012-0133221020213202-2231013201333232-3002011230111232-1211201331213222-3023333323322003-2331020203110200"></a>

## https_management.advertise_on_slo_vip.tls_config.custom_security — custom_security / 323322220302 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-1220000301313030-3210123130033303-0000003031222101-0230222200021002-2212330020213020-3221320323031031-1203232101302221-3230202011332313)
- https_management.advertise_on_slo_vip.tls_config.custom_security

<a id="canonical-1101323212030332-1010223232300102-2033332330332011-2033111302131100-3233012310013022-3031312022211032-0021311300033323-1212103303312230"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332123022120202-1000202022200030-0233201110320120-3322210103120213-1303333133333302-2131330031131103-2212030002223301-0220021120320000"></a>

## Direct properties — custom_security / 323322220302 / 3

<a id="canonical-0331230020232022-1200003320012120-0033031001113120-1202021230311210-1312302013230323-2320030211223320-0312110303110230-0010022033212131"></a>

<a id="canonical-3122113321023023-3000220310302330-2021301102202113-2203323030320231-3232322203121010-0131010001202221-0100203133230232-1113312331101231"></a>

## cipher_suites property — custom_security / 323322220302 / 4

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3222200121010012-0330230200120233-1230200100000201-3130001333130330-3320212330103012-0330033220330233-2113110113232331-2312100012301030"></a>

<a id="canonical-1102101203213133-3202232210010032-1020021302321202-3313211131202020-1012221302131233-0201021011030202-0330100203010022-2002033001333310"></a>

## max_version property — custom_security / 323322220302 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2321001301201123-2233213212103110-2222201000022123-1202123011102330-2330201003211002-0232030323211122-2122212133110311-1330102023331302"></a>

<a id="canonical-0332210001002032-2022022310102023-0032332201012002-2332200323120303-2302122230230002-2201203231301321-1020321112331231-2113033203130321"></a>

## min_version property — custom_security / 323322220302 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0232201132201332-2011110111030233-2220010333013201-0301230113223322-1220120111032102-1131012120032332-1210310333302212-0212120313322231"></a>

## Next pages — custom_security / 323322220302 / 7

- [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-1220000301313030-3210123130033303-0000003031222101-0230222200021002-2212330020213020-3221320323031031-1203232101302221-3230202011332313)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1330131011213232-0312111203313122-2030100023332311-2111021013121100-2202330011021130-3333030133113322-1111232220133201-0121203113031103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300032131312221-3222312132121030-0131023012303102-1002010111331302-1120230302323130-0021202200220131-0203021100032321-0132322221100202"></a>

## https_management.advertise_on_slo_vip.tls_config.default_security — default_security / 013221312331 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-1220000301313030-3210123130033303-0000003031222101-0230222200021002-2212330020213020-3221320323031031-1203232101302221-3230202011332313)
- https_management.advertise_on_slo_vip.tls_config.default_security

<a id="canonical-3332030021111030-1113101321222033-0201230203232021-2213021300313011-2211121121031031-2223003300211230-1012300212211330-2122330213231131"></a>

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
default_security = {}
```

<a id="canonical-0022101121000302-3232112022112222-1333331000011201-0023320121320301-0200320200100230-2123332231033210-1322323330023111-1110013133303303"></a>

## Direct properties — default_security / 013221312331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323012233001301-0211023233323221-2212323001332132-3200332311011122-0303031013133330-2032330110213021-2330322022333123-1212313013101112"></a>

## Next pages — default_security / 013221312331 / 4

- [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-1220000301313030-3210123130033303-0000003031222101-0230222200021002-2212330020213020-3221320323031031-1203232101302221-3230202011332313)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-0221112003220310-3222311333330300-3303223310301220-1032022013201000-2103003200212201-2133333333032331-0232313212202133-0021123223003122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300002201300003-2103200221100331-1132013121120201-3003122321303021-1210100213300031-2222132211101030-2110120020001231-0001021202202101"></a>

## https_management.advertise_on_slo_vip.tls_config.low_security — low_security / 010321003010 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-1220000301313030-3210123130033303-0000003031222101-0230222200021002-2212330020213020-3221320323031031-1203232101302221-3230202011332313)
- https_management.advertise_on_slo_vip.tls_config.low_security

<a id="canonical-1000233322232110-0231303001213012-3003210203323132-0103012001221131-1023320201112120-0112310313321313-0033200012313231-2132031312203211"></a>

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
low_security = {}
```

<a id="canonical-0301332103222033-0300312113010003-2331212110233323-2322031110123230-1303000130130201-0133212312231220-2100321200313201-2313103121332330"></a>

## Direct properties — low_security / 010321003010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221311111122332-3131101001311331-2003120121220221-2103211301122311-2201111102331002-1332020201331312-2012133301300332-0312030333130330"></a>

## Next pages — low_security / 010321003010 / 4

- [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-1220000301313030-3210123130033303-0000003031222101-0230222200021002-2212330020213020-3221320323031031-1203232101302221-3230202011332313)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-3131221302121031-0301003201212113-0121002010311301-2233331000033101-2332322333202130-1223331020103211-2233033313330003-1212230311012203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022002312330033-2233113013323303-1312311330132330-1212231103031213-3323102320020020-2010103322003110-1032003130113313-2122021332103113"></a>

## https_management.advertise_on_slo_vip.tls_config.medium_security — medium_security / 111220000032 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-1220000301313030-3210123130033303-0000003031222101-0230222200021002-2212330020213020-3221320323031031-1203232101302221-3230202011332313)
- https_management.advertise_on_slo_vip.tls_config.medium_security

<a id="canonical-0023303013333212-1123100113313230-2200121022202303-2320303332321100-2100003333121222-3231121301313320-0130303121200011-3300111302033322"></a>

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
medium_security = {}
```

<a id="canonical-2122333202331033-1001320221200231-1303030211311021-3223213021223131-1102102101001321-3230213230332001-3100012020110131-1312133103301012"></a>

## Direct properties — medium_security / 111220000032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120120122022010-2221203301121101-3021223032101220-1110201102312122-1301110011020112-0221300210201022-3203011123003131-0130303221110130"></a>

## Next pages — medium_security / 111220000032 / 4

- [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-1220000301313030-3210123130033303-0000003031222101-0230222200021002-2212330020213020-3221320323031031-1203232101302221-3230202011332313)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-0013001011312211-0302330103310232-2330013030103010-2111232121111121-1132202200220233-3303310322202002-2110202013220100-2223302323332223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122123032302302-0131110011322133-3132313323202233-3321321331012322-0000232200010013-2012102002221230-0230130321003220-1331232320111302"></a>

## https_management.advertise_on_slo_vip.use_mtls — use_mtls / 112121222232 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- https_management.advertise_on_slo_vip.use_mtls

<a id="canonical-2200112003032112-1010021111303003-1023010233101132-0112120202101103-3020221312010011-0211211303333131-0011021310102020-1321233302210010"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312231022220200-0030213113211221-2210233100303213-1101113331310113-1103210030321213-2133333030312102-1213323200120200-3100331203032103"></a>

## Direct properties — use_mtls / 112121222232 / 3

<a id="canonical-2313330321013002-3202013002330222-3233113013221101-2033331002030032-3230000132200021-2200023131323111-2221032030201302-2223312202000202"></a>

<a id="canonical-0112000320210211-1331003122011033-0303203033033202-2122103033111112-2031202130202232-3132212130023300-0112130130211132-2310313113213232"></a>

## client_certificate_optional property — use_mtls / 112121222232 / 4

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](resources--nfv_service--reference--group-003.md#canonical-1211213133032023-1331331032210202-2232222231230112-1223110221111122-0031133333203122-3211121232003102-0311013132001132-2030130032333110): complete subsection reference.

- [no_crl](resources--nfv_service--reference--group-003.md#canonical-2311332333230021-3322130201230221-3312013100112301-2231212212201121-1210212232111333-1222102200203003-1012133201130323-3110011030102211): complete subsection reference.

- [trusted_ca](resources--nfv_service--reference--group-003.md#canonical-2123010010310132-1121312303020113-1101203321100033-0203101122233000-1331100230312132-0322321103013220-2310323311002011-3222230012313000): complete subsection reference.

<a id="canonical-0133012010033123-2120131301102210-3100311011020003-1313322222022022-2122121310010311-1112131121303110-0200320302330311-1203223223113322"></a>

<a id="canonical-0332120302120333-0200333102130223-1200100132302301-3320301203033003-1331012103221230-3101112230103201-2003030013030313-0010323332100220"></a>

## trusted_ca_url property — use_mtls / 112121222232 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Provider validators and defaults (from schema source):

```go
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

- [xfcc_disabled](resources--nfv_service--reference--group-003.md#canonical-1021123323213212-2110100230220301-2012331211010233-3133233011211323-3212030003231332-0321230200103330-3030011112320332-0222102101011131): complete subsection reference.

- [xfcc_options](resources--nfv_service--reference--group-003.md#canonical-3223112210013103-1311113030331010-3121321332011303-2102323201001131-3220022230313220-2312300022222123-0100011211310032-2231110110020003): complete subsection reference.

<a id="canonical-1222122030213011-1321201210220312-1022201010122331-2320003000022220-0112130310002320-1330301312110020-0321102011321332-2011213223013001"></a>

## Next pages — use_mtls / 112121222232 / 6

- [https_management.advertise_on_slo_vip.use_mtls.crl](resources--nfv_service--reference--group-003.md#canonical-1211213133032023-1331331032210202-2232222231230112-1223110221111122-0031133333203122-3211121232003102-0311013132001132-2030130032333110)
- [https_management.advertise_on_slo_vip.use_mtls.no_crl](resources--nfv_service--reference--group-003.md#canonical-2311332333230021-3322130201230221-3312013100112301-2231212212201121-1210212232111333-1222102200203003-1012133201130323-3110011030102211)
- [https_management.advertise_on_slo_vip.use_mtls.trusted_ca](resources--nfv_service--reference--group-003.md#canonical-2123010010310132-1121312303020113-1101203321100033-0203101122233000-1331100230312132-0322321103013220-2310323311002011-3222230012313000)
- [https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled](resources--nfv_service--reference--group-003.md#canonical-1021123323213212-2110100230220301-2012331211010233-3133233011211323-3212030003231332-0321230200103330-3030011112320332-0222102101011131)
- [https_management.advertise_on_slo_vip.use_mtls.xfcc_options](resources--nfv_service--reference--group-003.md#canonical-3223112210013103-1311113030331010-3121321332011303-2102323201001131-3220022230313220-2312300022222123-0100011211310032-2231110110020003)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1211213133032023-1331331032210202-2232222231230112-1223110221111122-0031133333203122-3211121232003102-0311013132001132-2030130032333110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120303221202220-3211312220012123-1030100133012230-0202021321333111-3112330121301330-3033200322033313-3310320312331221-0132030312011000"></a>

## https_management.advertise_on_slo_vip.use_mtls.crl — crl / 010232113133 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0013001011312211-0302330103310232-2330013030103010-2111232121111121-1132202200220233-3303310322202002-2110202013220100-2223302323332223)
- https_management.advertise_on_slo_vip.use_mtls.crl

<a id="canonical-2130002231333210-1313222133110302-1202303320213031-1120222210010333-2113311301133022-0223110211302332-2001012300332310-1033230133111202"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-2331301012020333-0302003223310121-3011332323213113-1332033330133211-2100230212301001-0121200001302230-3211231033001320-3101032132120033"></a>

## Direct properties — crl / 010232113133 / 3

<a id="canonical-0002321323323132-1010301303030231-2320113002130013-1020301220023133-2300320200323312-1103232022122330-1120113221331012-0210120301132230"></a>

<a id="canonical-1310022003032300-3001000223030031-3313113222222003-2232023210300213-0123200302302223-1031232303201002-0101203132233203-1132201030332303"></a>

## name property — crl / 010232113133 / 4

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

<a id="canonical-3312212220211032-3201123022211330-2033303213322312-3012001011302233-3003213023103321-3032313312200223-2302301132032330-2122321210231301"></a>

<a id="canonical-0013331113130202-1122323103323013-2120303131000112-1020120133011303-1022331103101030-1201202321023302-2122120221221133-2120101330132322"></a>

## namespace property — crl / 010232113133 / 5

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

<a id="canonical-3103200223220101-1112312121012233-0212111223230112-2201232200331020-3211311233032203-0100303010200010-3313313320032223-0311302320002003"></a>

<a id="canonical-3333131233200131-0121011123031331-2332320203120100-1220032312133201-3122323222023113-1123103302110313-1210303300213222-1120021100010211"></a>

## tenant property — crl / 010232113133 / 6

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

<a id="canonical-0133020210131220-3212022203301001-3122300322102133-0230113021200013-2030210110021013-0123310010200220-2131322333212303-0031211200221103"></a>

## Next pages — crl / 010232113133 / 7

- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0013001011312211-0302330103310232-2330013030103010-2111232121111121-1132202200220233-3303310322202002-2110202013220100-2223302323332223)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-2311332333230021-3322130201230221-3312013100112301-2231212212201121-1210212232111333-1222102200203003-1012133201130323-3110011030102211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133201023211221-0203030220330130-3001213110230300-1111020003012012-0003321223111030-3312231021320103-3231213013003132-2300133102020202"></a>

## https_management.advertise_on_slo_vip.use_mtls.no_crl — no_crl / 003312003201 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0013001011312211-0302330103310232-2330013030103010-2111232121111121-1132202200220233-3303310322202002-2110202013220100-2223302323332223)
- https_management.advertise_on_slo_vip.use_mtls.no_crl

<a id="canonical-2132332003300103-3113011002312213-2031132200002102-0302130103313111-3131111110203310-2011221310312122-2021223021121012-0233230210033201"></a>

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

<a id="canonical-3220300320322133-3230330012232131-1021223022323100-0302000322031321-1112321203303102-1202320130203003-3211223101033100-1303302200032132"></a>

## Direct properties — no_crl / 003312003201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113101031210022-3120121100021122-2331023221310013-1123211303213021-3023211131220210-0220002330003332-1030032102211020-0011130112000221"></a>

## Next pages — no_crl / 003312003201 / 4

- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0013001011312211-0302330103310232-2330013030103010-2111232121111121-1132202200220233-3303310322202002-2110202013220100-2223302323332223)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-2123010010310132-1121312303020113-1101203321100033-0203101122233000-1331100230312132-0322321103013220-2310323311002011-3222230012313000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321110020330323-1202020003213220-1023101321232303-1001202021231130-1013123213030312-2010020303212300-1333020323322113-3103113202101001"></a>

## https_management.advertise_on_slo_vip.use_mtls.trusted_ca — trusted_ca / 311203110331 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0013001011312211-0302330103310232-2330013030103010-2111232121111121-1132202200220233-3303310322202002-2110202013220100-2223302323332223)
- https_management.advertise_on_slo_vip.use_mtls.trusted_ca

<a id="canonical-3031030100112003-0123312130332120-0323100223200332-3030021200122113-0301100330022102-2000230112110033-2212000300220000-3103211000022320"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-0010002032223101-1320231312222011-0132003000012100-1323103121121321-0312221331330133-0232120323113023-0321122312222300-1103323211313213"></a>

## Direct properties — trusted_ca / 311203110331 / 3

<a id="canonical-1303103011321322-0323212111022331-0022330123211002-2313223001032311-3011001302302133-2331333300310021-2221313131312021-3030300103002333"></a>

<a id="canonical-1211331110022210-3121333210100133-0312102310102131-1210033211313213-1123211230001310-2211203001201221-3112232011303123-1330213030033023"></a>

## name property — trusted_ca / 311203110331 / 4

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

<a id="canonical-3211013020303230-3011013233210033-2320303320023312-1213303231321100-0131013113021300-2130002332112020-3011333130202031-1003321110012002"></a>

<a id="canonical-1121331212023010-0312210302023001-1103200231322301-0200221221100302-1322213310020303-3320303330110111-2200321111323331-0220221333122012"></a>

## namespace property — trusted_ca / 311203110331 / 5

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

<a id="canonical-1222311001323310-0121022013112311-0013233000120121-2100230230001020-0333223301123233-0221313310332013-2023012013013330-1012222331123021"></a>

<a id="canonical-0032312331002322-1331022130230300-0321122221311103-3033201013331230-2312102330211332-1200000111122301-2322230201031120-0333200023223021"></a>

## tenant property — trusted_ca / 311203110331 / 6

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

<a id="canonical-3001311210323031-1010312303220300-1001333223233121-0031303202030300-1122122131301201-2213022220232311-2331221032232320-1232230210303011"></a>

## Next pages — trusted_ca / 311203110331 / 7

- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0013001011312211-0302330103310232-2330013030103010-2111232121111121-1132202200220233-3303310322202002-2110202013220100-2223302323332223)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1021123323213212-2110100230220301-2012331211010233-3133233011211323-3212030003231332-0321230200103330-3030011112320332-0222102101011131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122102110012321-2311103110100232-1100013121222010-1223303012333103-3313310333120013-1303300003110120-2331213112112220-1121212121132231"></a>

## https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled — xfcc_disabled / 310130302110 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0013001011312211-0302330103310232-2330013030103010-2111232121111121-1132202200220233-3303310322202002-2110202013220100-2223302323332223)
- https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled

<a id="canonical-2032322113310033-0223322130231333-2230312000123322-0031311111333312-2212331330110232-0010301130023102-2331310031032011-0103221310211302"></a>

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

<a id="canonical-1000031313333210-1312201111313113-3032122003120003-3111232110211010-2210330013323021-0111012111113102-3330230101031032-1300121131000122"></a>

## Direct properties — xfcc_disabled / 310130302110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331231033131202-2220300320021332-1212003032010021-0013021313100301-0313110032333001-1110012322032131-2112023131313332-0210123031201030"></a>

## Next pages — xfcc_disabled / 310130302110 / 4

- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0013001011312211-0302330103310232-2330013030103010-2111232121111121-1132202200220233-3303310322202002-2110202013220100-2223302323332223)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-3223112210013103-1311113030331010-3121321332011303-2102323201001131-3220022230313220-2312300022222123-0100011211310032-2231110110020003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220211212210131-0002121230303102-3030223212331231-1030022230200020-1013102002321310-3300203201303020-0011012012203212-1100323012320331"></a>

## https_management.advertise_on_slo_vip.use_mtls.xfcc_options — xfcc_options / 230021321020 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0013001011312211-0302330103310232-2330013030103010-2111232121111121-1132202200220233-3303310322202002-2110202013220100-2223302323332223)
- https_management.advertise_on_slo_vip.use_mtls.xfcc_options

<a id="canonical-2333231231323323-2133022210030203-3003200312100312-3300210313131010-2220111223223331-3213223302002320-1123000021321320-1010110330311303"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2013110203102102-2232310023022310-2213220233332033-0110003200113103-2333330323012130-3212113012003103-3212220203000110-3321002002100322"></a>

## Direct properties — xfcc_options / 230021321020 / 3

<a id="canonical-2233231322021022-3003020210323122-3331331121100030-0211132331301002-0203020201333001-0201010101210222-1110113202032020-2000010000132101"></a>

<a id="canonical-1302331202201322-2210312100121123-2221310102101033-3311313020321121-2300223131010212-0330221101101101-3122113232210323-1130022221312100"></a>

## xfcc_header_elements property — xfcc_options / 230021321020 / 4

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

<a id="canonical-3120320220130002-2001133203301300-1121222020120300-3100011022220110-0202112113122031-0321113303220330-2223311101122332-0032321322031101"></a>

## Next pages — xfcc_options / 230021321020 / 5

- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0013001011312211-0302330103310232-2330013030103010-2111232121111121-1132202200220233-3303310322202002-2110202013220100-2223302323332223)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1031033213022103-1023333001002003-1001122010332122-2120200101013131-3210333031311132-3130212220330310-3120322302022212-0102303131310332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312222012230313-2333203132011111-0220223001232003-2111312311321131-0310101223001333-2211231203031021-0213332010120123-3032112111301200"></a>

## https_management.default_https_port — default_https_port / 130122010311 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- https_management.default_https_port

<a id="canonical-0320211030330020-2221000001121023-2100203330232133-3311230331201001-0322003221233200-2022031033323333-0112023223112111-3022322032020002"></a>

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
default_https_port = {}
```

<a id="canonical-0213211322133113-2101213031002322-1301030331322303-0300011321223130-3313322213330311-2210323200021202-3222131332213322-0311101311103212"></a>

## Direct properties — default_https_port / 130122010311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232303333213213-2210003300121231-3200000101330013-2312110112011303-1310133310311131-3311231321003031-2121200113210332-0003132100203322"></a>

## Next pages — default_https_port / 130122010311 / 4

- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120101301121331-3020233232023021-3012110322321012-2311201301110132-3312323303130331-0030320012223323-3221303030022133-1200111320100322"></a>

## palo_alto_fw_service — palo_alto_fw_service / 123201032120 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- palo_alto_fw_service

<a id="canonical-1000211303110321-2132033102233220-3202330310311301-1011321010001303-2002102233013002-0110203131021011-1312303320302102-2012032301010010"></a>

Type: `"object"`. single nested block, Optional.

Palo Alto Networks VM-Series next-generation firewall configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_setup",
    "ssh_key"),
  validators.ConflictingObjectAttributes("disable_panaroma",
    "panorama_server"),
  validators.ConflictingObjectAttributes("pan_ami_bundle1",
    "pan_ami_bundle2")}
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
  "x-ves-oneof-field-ami_choice": "[\"pan_ami_bundle1\",\"pan_ami_bundle2\"]",
  "x-ves-oneof-field-panaroma_connection": "[\"disable_panaroma\",\"panorama_server\"]",
  "x-ves-oneof-field-setup_options": "[\"auto_setup\",\"ssh_key\"]"
}
```

Terraform syntax:

```terraform
palo_alto_fw_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223021102201021-1033330120223023-1221301333323222-3030013112110002-2301121203003230-0000330303002213-3321231003321031-0122013110231203"></a>

## Direct properties — palo_alto_fw_service / 123201032120 / 3

- [auto_setup](resources--nfv_service--reference--group-003.md#canonical-1323331203203220-3311331231103031-1001233001210020-3313130122302010-2233001013101302-3333333033103312-2000002012201020-3033003110033123): complete subsection reference.

- [aws_tgw_site](resources--nfv_service--reference--group-004.md#canonical-1301111030110103-0221103130313220-0310020330231021-3213300230211120-0300001131331322-1302213200103213-2300100123201131-0022123322310221): complete subsection reference.

- [disable_panaroma](resources--nfv_service--reference--group-004.md#canonical-1123032201111330-0000230312030121-2011301213013022-0012110231212311-2323223023121032-2121213231232023-1132132013020021-3132002201011002): complete subsection reference.

<a id="canonical-2213122313101232-3203212022110121-3113222323130102-1332110123031000-1321203232212032-3010311012231201-3130032232010011-1003302030223131"></a>

<a id="canonical-3102012112330323-3103331312002333-1330021221013202-2110022132123321-3200122111311322-3200223232232222-2311000010220330-1313212210221303"></a>

## instance_type property — palo_alto_fw_service / 123201032120 / 4

Type: `"string"`. Optional.

\[Enum:
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_12XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_8XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_9XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_18XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_9XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_18XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_R5\_2XLARGE\]
&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE: m4.xlarge -
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_2XLARGE: m4.2xlarge -
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_4XLARGE: m4.4xlarge -
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_LARGE: m5.large -
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_XLARGE: m5.xlarge .. Possible values are
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_12XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_8XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_9XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_18XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_9XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_18XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_R5\_2XLARGE\`. Defaults to
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE\`.

Upstream description:

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE: m4.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_2XLARGE: m4.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_4XLARGE: m4.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_LARGE: m5.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_XLARGE: m5.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_2XLARGE: m5.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_4XLARGE: m5.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_12XLARGE: m5.12xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_LARGE: m5n.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_XLARGE: m5n.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_2XLARGE: m5n.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_4XLARGE: m5n.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_LARGE: c4.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_XLARGE: c4.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_2XLARGE: c4.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_4XLARGE: c4.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_8XLARGE: c4.8xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_LARGE: c5.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_XLARGE: c5.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_2XLARGE: c5.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_4XLARGE: c5.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_9XLARGE: c5.9xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_18XLARGE: c5.18xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_LARGE: c5n.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_XLARGE: c5n.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_2XLARGE: c5n.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_4XLARGE: c5n.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_9XLARGE: c5n.9xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_18XLARGE: c5n.18xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_R5\_2XLARGE: r5.2xlarge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_12XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_8XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_9XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_18XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_9XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_18XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_R5_2XLARGE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_XLARGE",
  "enum": [
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_12XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_8XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_9XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_18XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_9XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_18XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_R5_2XLARGE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pan_ami_bundle1](resources--nfv_service--reference--group-004.md#canonical-3033330222311121-1332010330203303-1002302201120120-0021101102310211-3031013112310322-1220320330331121-3212331103011223-3332101111321030): complete subsection reference.

- [pan_ami_bundle2](resources--nfv_service--reference--group-004.md#canonical-1211032002030203-2321110130000111-0100333311120113-3102313030201023-0131001310210103-3201210103110232-2302102303231003-3203101301001000): complete subsection reference.

- [panorama_server](resources--nfv_service--reference--group-004.md#canonical-1011323001103321-2030100221203210-2122232000230032-2223123210312101-0202020002101023-0101121233000113-3313223221110312-1303032221032111): complete subsection reference.

- [service_nodes](resources--nfv_service--reference--group-004.md#canonical-2031200013222301-1113112221323113-0100320310111020-2012332001313133-1212000001301222-1002020103233310-3330332120300223-1122101300020031): complete subsection reference.

<a id="canonical-0012322131130123-0120023232110303-0000302311232212-2101211230330313-0021311112213233-0111102100133301-3320122002203211-0103111331231221"></a>

<a id="canonical-2221133031333201-3033003102012113-2310100231131031-1013020220021111-2302000232102202-0130321231231122-2223102101012032-0112012300132231"></a>

## ssh_key property — palo_alto_fw_service / 123201032120 / 5

Type: `"string"`. Optional.

Exclusive with \[auto\_setup\] Setup Authorized Public SSH key. User will be able to SSH to the
vmseries nodes using its corresponding SSH private key.

Upstream description:

Exclusive with \[auto\_setup\] Setup Authorized Public SSH key. User will be able to SSH to the
vmseries nodes using its corresponding SSH private key.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2222200222102312-2213202200002330-2322120223121012-0211010210303330-1331120313221203-3213032211211020-3002123130003131-1110320202331232"></a>

<a id="canonical-3012032133123332-3111113112100032-3002012100133032-0012231033232130-3023230303002103-3003002210221000-2222110012110233-1230021132001132"></a>

## tags property — palo_alto_fw_service / 123201032120 / 6

Type: `["map", "string"]`. Optional.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Upstream description:

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

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
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

<a id="canonical-3133201303322112-3200103333121012-2123130323122010-0300222021211033-1022111311000133-1210121113303002-1212230303100230-2231000130111310"></a>

<a id="canonical-1233113033310022-2302012211122022-2102011203100122-0332333323021321-3200320131211003-3033032112031222-3101133020011112-1121021012011230"></a>

## version property — palo_alto_fw_service / 123201032120 / 7

Type: `"string"`. Optional.

\[Enum: 11.0.0\] PAN VM-Series version. PAN-OS version. The only possible value is \`11.0.0\`.

Upstream description:

PAN-OS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("11.0.0"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "11.0.0"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"11.0.0\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"11.0.0\\\"]"
  }
}
```

<a id="canonical-0101120131120010-2120013301210231-1023223103010112-2112321302321212-3213001322020200-0303320023212013-0210101130133103-0100233201200110"></a>

## Next pages — palo_alto_fw_service / 123201032120 / 8

- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-1323331203203220-3311331231103031-1001233001210020-3313130122302010-2233001013101302-3333333033103312-2000002012201020-3033003110033123)
- [palo_alto_fw_service.aws_tgw_site](resources--nfv_service--reference--group-004.md#canonical-1301111030110103-0221103130313220-0310020330231021-3213300230211120-0300001131331322-1302213200103213-2300100123201131-0022123322310221)
- [palo_alto_fw_service.disable_panaroma](resources--nfv_service--reference--group-004.md#canonical-1123032201111330-0000230312030121-2011301213013022-0012110231212311-2323223023121032-2121213231232023-1132132013020021-3132002201011002)
- [palo_alto_fw_service.pan_ami_bundle1](resources--nfv_service--reference--group-004.md#canonical-3033330222311121-1332010330203303-1002302201120120-0021101102310211-3031013112310322-1220320330331121-3212331103011223-3332101111321030)
- [palo_alto_fw_service.pan_ami_bundle2](resources--nfv_service--reference--group-004.md#canonical-1211032002030203-2321110130000111-0100333311120113-3102313030201023-0131001310210103-3201210103110232-2302102303231003-3203101301001000)
- [palo_alto_fw_service.panorama_server](resources--nfv_service--reference--group-004.md#canonical-1011323001103321-2030100221203210-2122232000230032-2223123210312101-0202020002101023-0101121233000113-3313223221110312-1303032221032111)
- [palo_alto_fw_service.service_nodes](resources--nfv_service--reference--group-004.md#canonical-2031200013222301-1113112221323113-0100320310111020-2012332001313133-1212000001301222-1002020103233310-3330332120300223-1122101300020031)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1323331203203220-3311331231103031-1001233001210020-3313130122302010-2233001013101302-3333333033103312-2000002012201020-3033003110033123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110031313330101-1123131133003001-0203121323312021-3001330103023220-3123123313120301-1200103131230121-2220031000321030-0030222112332132"></a>

## palo_alto_fw_service.auto_setup — auto_setup / 021133313232 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- palo_alto_fw_service.auto_setup

<a id="canonical-3310302012302032-0132011203023101-1302100100013130-1010111331312131-2232100232203002-1131312233001023-3132321131020132-0022132333300322"></a>

Type: `"object"`. single nested block, Optional.

For auto-setup, SSH public and pvt keys are needed. Using the given config user, SSH and API access
will be configured.

Upstream description:

For auto-setup, SSH public and pvt keys are needed. Using the given config user, SSH and API access
will be configured.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("admin_username")}
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
  "x-ves-oneof-field-ssh_keys_choice": "[\"manual_ssh_keys\"]"
}
```

Terraform syntax:

```terraform
auto_setup {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020331321213203-3301331112201111-0113320003233120-1232113321011022-0332102212001313-1122000330220302-2100112122020212-0221301120013100"></a>

## Direct properties — auto_setup / 021133313232 / 3

- [admin_password](resources--nfv_service--reference--group-004.md#canonical-2021013202121333-0332132013033030-0130230331300102-0033212321212113-3001333312121030-0303203233003111-2123010300222022-0131120132110211): complete subsection reference.

<a id="canonical-2000103103321101-3101320102231022-0310121132320223-2021323230002232-1301322031210000-3013100302223002-2032310312033030-2010023002232300"></a>

<a id="canonical-2013301200211112-1100003032311320-0023112202120310-2332113221123213-0021032130023101-0103210313310210-1003120310012133-0111202121002221"></a>

## admin_username property — auto_setup / 021133313232 / 4

Type: `"string"`. Optional.

Firewall Admin Username. Firewall Admin Username.

Upstream description:

Firewall Admin Username.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [manual_ssh_keys](resources--nfv_service--reference--group-004.md#canonical-0033320030010322-1202102231033331-2100333321203211-1002102220220222-2103103103300010-1013013232210101-3211202330031320-2120100320132201): complete subsection reference.
