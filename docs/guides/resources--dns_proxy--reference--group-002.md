---
page_title: "xcsh_dns_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_proxy reference."
---

# xcsh_dns_proxy reference

<a id="canonical-0312123332210011-2123321111303033-2211013112322122-1200003210132303-1031103133030000-1112011121120211-2330022301320032-3201030310320212"></a>

## Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_network`

- [default_v6_vip](resources--dns_proxy--reference--group-002.md#canonical-1220122210013202-3110010111110212-1322100322101333-0033320123132330-2301100013223132-0313233301201313-2001210202121310-0211123203323300): complete subsection reference.

- [default_vip](resources--dns_proxy--reference--group-002.md#canonical-2203223331301230-3330320130220010-1011311223200330-0230333310220212-1123301033122322-1311022221221232-2302221010123022-3321230233223303): complete subsection reference.

<a id="canonical-3110212102132200-3120232310321203-2322303302230101-3130230123100321-1000222030302202-3201130033020301-2003010211213011-0003112111133110"></a>

<a id="canonical-2110010012123031-0101000102210021-1133320300100222-1023222301321211-0001332211211031-2302233011310000-3100022231133033-0121032213313120"></a>

### `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip` property

Type: `"string"`. Optional.

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-1230000013311031-2131303130201300-1003313310130233-2121111000210002-0102300220101132-2320122211322330-1231310112321213-2001310200000012"></a>

<a id="canonical-0031201321211112-2213002000120122-3032012130022233-0223203101113310-0003002332123332-1222211213000323-0032011002300022-1200113033221110"></a>

### `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip` property

Type: `"string"`. Optional.

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [virtual_network](resources--dns_proxy--reference--group-002.md#canonical-1110121310200022-1331323031201213-0110022323020313-2000103120103200-2003221011311312-3231202332120111-1312320311132020-2321212031032320): complete subsection reference.

<a id="canonical-1220122210013202-3110010111110212-1322100322101333-0033320123132330-2301100013223132-0313233301201313-2001210202121310-0211123203323300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-001.md#canonical-2310102012101212-1313110121233032-0222012221233332-0332003311302302-3001031232200013-1030223210030220-3330302311103220-0202103201323322)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-2203220020003122-1223331223330131-1213210021002211-1331300220113020-2320333000310102-3131200201120111-2213013021320012-2310133311103010"></a>

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
default_v6_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203223331301230-3330320130220010-1011311223200330-0230333310220212-1123301033122322-1311022221221232-2302221010123022-3321230233223303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-001.md#canonical-2310102012101212-1313110121233032-0222012221233332-0332003311302302-3001031232200013-1030223210030220-3330302311103220-0202103201323322)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip

<a id="canonical-3002013123313002-2002221302033222-2120010212131001-1113202112223310-1223122033121211-2003231332122221-2010302100230332-2020102010213031"></a>

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
default_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110121310200022-1331323031201213-0110022323020313-2000103120103200-2003221011311312-3231202332120111-1312320311132020-2321212031032320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-001.md#canonical-2310102012101212-1313110121233032-0222012221233332-0332003311302302-3001031232200013-1030223210030220-3330302311103220-0202103201323322)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network

<a id="canonical-0222213132202020-1121330003202230-2031323101023201-0012200222211231-2201122200130000-2031211101010131-2121002312020001-2010333323321223"></a>

Type: `"object"`. single nested block, Optional.

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
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-1120021031031203-3221033331230022-1003032112231222-1131221320302233-2103013301301203-2012212023023200-2201123032120113-1021300030332332"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network`

<a id="canonical-3000030120100023-2221123130313031-2233311332003232-0002031013010222-3312213033231011-3011100011222330-2101100112231230-1230121003332212"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name` property

Type: `"string"`. Optional.

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

<a id="canonical-1213201320010032-2323020212020103-0021021333011123-0011330010320222-1331222122213122-1110132200221330-1303303033121313-2201310210311213"></a>

<a id="canonical-1011203132203230-2002212233213023-3111003312202203-3130031313203133-3211231030101202-2310311020132333-2332220121100133-0131312301030113"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-2202003331322310-2223003132110232-3000010131230310-1332011033030300-1223232303013021-1030010201221032-3223001101011003-2131113113313200"></a>

<a id="canonical-0331013220202331-0213111233011002-0021301121002212-3322200330031201-3121023122003211-3333133212301021-0202213331321112-1212300231221331"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-2121010310201023-0012033031302111-2200021322330221-2310000013030030-0102122021131301-1220332330110022-3322322113031011-0121112303002232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_site` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site

<a id="canonical-2000331302020031-0312102230303301-0303023121001132-0321022210120003-3313010333322232-0213101331032233-1332020130010321-2122023120111232"></a>

Type: `"object"`. single nested block, Optional.

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333002001323311-2120101300003321-3220221232102210-3300103132213213-0122221321020202-2303203110310131-3223032033323310-1333023211020201"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_site`

<a id="canonical-1113203110221003-1131202132321021-0233311312113320-3101003021013022-2002201311231113-3323232113113300-3203323133303200-2003333221311120"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site.network` property

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on site

All inside and outside networks. All outside networks. All outside networks with internet VIP
support. VK8s service network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the
site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SITE_NETWORK_INSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_IP_FABRIC","SITE_NETWORK_OUTSIDE","SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_SERVICE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](resources--dns_proxy--reference--group-002.md#canonical-2030312330101200-1000002022223202-2111300132122030-3332310200013100-0313000330011321-2023221300331202-2230322112122203-3111000001231200): complete subsection reference.

<a id="canonical-2030312330101200-1000002022223202-2111300132122030-3332310200013100-0313000330011321-2023221300331202-2230322112122203-3111000001231200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-2121010310201023-0012033031302111-2200021322330221-2310000013030030-0102122021131301-1220332330110022-3322322113031011-0121112303002232)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-0021100002030003-2312230001232001-1122222010311010-0330101233331201-2212002313313212-1022301200032232-0321210320221323-0200300030302020"></a>

Type: `"object"`. single nested block, Optional.

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2100100001303013-1010311313222211-1103311123012202-3031022202001220-0323200001111023-0133222110232203-1210213032133022-1011302133011212"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site`

<a id="canonical-3020311221030122-1202003213021230-3131032021021123-1023202122213202-0112122213201330-1200010010221332-0211121231213322-3111110322011133"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name` property

Type: `"string"`. Optional.

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

<a id="canonical-0013220031130313-3013001031010213-2221032300333201-2012030233223110-2301003302233122-3333203233303012-1322210120302322-1132001132021111"></a>

<a id="canonical-3300123321002111-2023002300110302-1300210200331300-1313131320320320-0322111120222222-1213231122131103-2123023213320302-3323323302012000"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-0103203210211310-2021030133010203-1121110111130102-0333113032013112-0232113332232031-0303302210312110-0312013303012330-0302203133303123"></a>

<a id="canonical-3311113113112300-3113030001220121-2110003323112231-1202320111120110-2222130102002000-3000212113001101-2022133331011303-2112313301012012"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-3120210322311313-1031233122013201-3010103131230133-1220113202032033-0022231001022221-1121230121113102-2203301120120222-1130010230032210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-2322000121112300-0011031322121120-3021203331300001-3022021012123321-3132222132112302-0123301231010300-3031032201200212-0333213323230201"></a>

Type: `"object"`. single nested block, Optional.

This defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
virtual_site_with_vip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2120231210231310-1302030132202223-1312333300131023-1132122021123123-1312323101000200-0121120233233133-1233133300032310-3311130020022131"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip`

<a id="canonical-2132210203323022-2002010100023111-3312332103331320-1131203323222103-2312132320020230-0111211230011121-2230222232010302-1001133313323231"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip` property

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2200023110002111-3233330110110202-0310101130032002-1223121223231232-3012003001121201-0231330012333000-1101000133210332-3332221232103223"></a>

<a id="canonical-3031003200212101-1002210012031031-0101033133323223-3130031331332202-3301013321300320-3023302301303102-1330111310302113-1200223231021121"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network` property

Type: `"string"`. Optional.

\[Enum: SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE|SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\] Defines
network types to be used on virtual-site with specified VIP All outside networks. All inside
networks. Possible values are \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`,
\`SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\`. Defaults to \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on virtual-site with specified VIP

All outside networks.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SITE_NETWORK_SPECIFIED_VIP_INSIDE","SITE_NETWORK_SPECIFIED_VIP_OUTSIDE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
  "enum": [
    "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](resources--dns_proxy--reference--group-002.md#canonical-0312003011223001-1303020320021031-3321123013210122-0210122130020001-1123202233010322-3010132221110111-3120202331321130-1223310232330323): complete subsection reference.

<a id="canonical-0312003011223001-1303020320021031-3321123013210122-0210122130020001-1123202233010322-3010132221110111-3120202331321130-1223310232330323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--dns_proxy--reference--group-002.md#canonical-3120210322311313-1031233122013201-3010103131230133-1220113202032033-0022231001022221-1121230121113102-2203301120120222-1130010230032210)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-3133103110122322-3103021010203231-1100232131203130-0030320020132023-1101302110231321-3311101130112220-0113111310313232-1332312100121212"></a>

Type: `"object"`. single nested block, Optional.

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011031223010111-2321031112203212-3120333310123212-2010030323210311-3323132013121200-1101012032123211-3012220101201311-0013120222323221"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site`

<a id="canonical-2320122331023032-1102021312221201-1022013013303200-3030120130123203-0330203120331201-1102103331203103-2003001032031310-1320202001333232"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` property

Type: `"string"`. Optional.

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

<a id="canonical-1011122203221030-1123113212013023-3322120131223133-2123123031031331-0122203112213131-3301123121211030-2223133212310030-2120202120310203"></a>

<a id="canonical-3320302133201032-0230003213003221-2000013032133103-2213132213321132-1111221000231010-1323200330212202-0033120222030020-1022211123003131"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-1021133303223230-2201001230213313-1112001120332232-1321330111301111-3202320310003300-3023223333232321-0310200333311020-0012303210323222"></a>

<a id="canonical-3101113212101031-0210002201210023-0021200313100302-1310021300212333-3220312032022202-2123201220003003-1233303123033222-1320030031103222"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-3331020230023233-2321211121222102-3023033122212123-2101133133002010-1023311023330020-2220002321331011-0202230202222210-2111222102020333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.vk8s_service` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service

<a id="canonical-0130112030213102-0210233031303222-3021132321111132-0112031102232131-3010302133123322-0202211020200131-2330222113123311-0032222201313303"></a>

Type: `"object"`. single nested block, Optional.

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
vk8s_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-0331221331131300-0233311021021333-1033003210312233-0211130102023103-1113311011103233-0301132111130031-0312103010330322-2311133101131003"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.vk8s_service`

- [site](resources--dns_proxy--reference--group-002.md#canonical-3320333020331332-2201010102332312-0320303013333121-1132002321120201-2300021202011330-3010011302033130-0213133123001033-2020303200012230): complete subsection reference.

- [virtual_site](resources--dns_proxy--reference--group-002.md#canonical-3103101011303203-2022103310012021-3001313311211213-3322321020321310-0220133030010133-1112333000112311-3301302213102023-0231221312201301): complete subsection reference.

<a id="canonical-3320333020331332-2201010102332312-0320303013333121-1132002321120201-2300021202011330-3010011302033130-0213133123001033-2020303200012230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--dns_proxy--reference--group-002.md#canonical-3331020230023233-2321211121222102-3023033122212123-2101133133002010-1023311023330020-2220002321331011-0202230202222210-2111222102020333)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-2001123130003223-3102010022300233-1000113102130232-3002030030222122-2333113202212310-2233230001012130-1131333012200223-3011013031031102"></a>

Type: `"object"`. single nested block, Optional.

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1021223112032133-0302230102303310-3100030032320321-1103011011332131-3013130300210120-0323330022132333-0202313330330211-2102212011000102"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site`

<a id="canonical-0001221132022202-1222203001212323-3211203320130021-2131031222033322-1212101201101313-0012112101013012-1032032122330303-0302313133212012"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name` property

Type: `"string"`. Optional.

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

<a id="canonical-1030111032030311-1010221012023113-3010030331100301-3110133120232013-3332000133321222-3213321122302233-1003311131301231-0320010323102123"></a>

<a id="canonical-2032311113132200-2230111103110201-3120203220110123-2220023102302230-2101321300220113-3231211020101201-2333322021132211-0313233031232202"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-1330200132333031-0302212303001322-1303022012212031-0313203120323233-3210211020321100-2300031121332202-2113331302021211-2301122231333021"></a>

<a id="canonical-1000211130232130-1121332200301103-1100121021100223-0100221323200211-1210111021202301-3112112313211210-1212112330102312-3001230112020212"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-3103101011303203-2022103310012021-3001313311211213-3322321020321310-0220133030010133-1112333000112311-3301302213102023-0231221312201301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--dns_proxy--reference--group-002.md#canonical-3331020230023233-2321211121222102-3023033122212123-2101133133002010-1023311023330020-2220002321331011-0202230202222210-2111222102020333)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-0220022230213123-0211200020100231-3333333112201112-3000110013002000-2203230312011032-0311013030310231-2212331121000202-1330233300102112"></a>

Type: `"object"`. single nested block, Optional.

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2222003223332312-1022331203312321-0322113231310000-0212221113121200-0123331330111331-1033013002003000-2003313202221320-1201211200013130"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site`

<a id="canonical-0020112322220221-3202032303112203-0202110203223033-2032020031200100-0120202232310132-1122110130303202-2031212203011100-2300112220022021"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name` property

Type: `"string"`. Optional.

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

<a id="canonical-0232011111001203-2020010031301132-1212112131222022-2313313221321230-0331111331103203-3003203131313330-2033212132120232-3021223233113030"></a>

<a id="canonical-3033010011300331-2123322022123001-1221001320112233-3123023000030111-3232333113021032-1132232213301121-1012122331030233-0132311232310330"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-0003331312302013-3110210300130033-2303030030031110-1210200012231303-3013110103323223-1003020133322202-1123101122331232-3213101110333123"></a>

<a id="canonical-2012302310303302-3100011031300130-3231202302120002-1322133333320222-1221101203010110-1210312100122300-2302033322032233-3213320023133131"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-0312321001033003-2000102210132110-2013301302220000-1320023122223111-1132300310003200-1203011203201332-2211111321321133-0310222130312302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_dualstack_on_public` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- proxy_advertisement.advertise_dualstack_on_public

<a id="canonical-3020020013121233-3031311033231313-0030211030221130-1123003331111033-0111302233303300-1221131132300211-2232120103231112-3122030121232213"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
advertise_dualstack_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-3331230202133022-0010212013011332-2001110222010200-1120021020023222-2000222022011201-1212200320001113-0033112000010323-0310023300033032"></a>

### Direct properties for `proxy_advertisement.advertise_dualstack_on_public`

- [public_ip](resources--dns_proxy--reference--group-002.md#canonical-2232002212321132-3033031231032131-3032131102301100-2222201012313302-1013332023321100-0010330303230101-0100001000300303-2231202322201223): complete subsection reference.

<a id="canonical-2232002212321132-3033031231032131-3032131102301100-2222201012313302-1013332023321100-0010330303230101-0100001000300303-2231202322201223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_dualstack_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-0312321001033003-2000102210132110-2013301302220000-1320023122223111-1132300310003200-1203011203201332-2211111321321133-0310222130312302)
- proxy_advertisement.advertise_dualstack_on_public.public_ip

<a id="canonical-2203212130210320-2323202330103120-0202202120201120-3320313010330321-1301111200321033-1010121211013230-3210212103332222-2002311303202022"></a>

Type: `"object"`. single nested block, Optional.

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2012010132112233-0222030112010330-0203013331211301-1022010123001021-1331123113130223-1101101223022013-1233032020223303-2333120331221100"></a>

### Direct properties for `proxy_advertisement.advertise_dualstack_on_public.public_ip`

<a id="canonical-2221313111301231-0111330101312110-3211320220030333-2000131202113000-1031330103121303-3321320231010021-0230310222220113-1220131023312000"></a>

#### `proxy_advertisement.advertise_dualstack_on_public.public_ip.name` property

Type: `"string"`. Optional.

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

<a id="canonical-2300101312112133-3031003200211203-2100223222111003-2133032111210030-3120201122220132-1230102223320023-1113203231300323-1302301113211133"></a>

<a id="canonical-3333322301020231-0011333112323210-2322101213032031-0301030310220010-2320300111101302-1233321220232323-3212133331120110-1132231021001312"></a>

#### `proxy_advertisement.advertise_dualstack_on_public.public_ip.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-3022011221103222-3320122312311232-2222231213323312-1023320232231113-0212213020022200-3300200113133311-1120320021023333-0000030333122013"></a>

<a id="canonical-3001231111322111-2323023312223230-2102133113211103-2023112302131112-2323130031022310-0300310002113021-0312023213100200-3003112132120111"></a>

#### `proxy_advertisement.advertise_dualstack_on_public.public_ip.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-2220303323213103-2110110021231302-1232003120023021-3003122123002120-0021020320000023-3232200110132031-0222010121032220-2103301122232110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_on_public` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- proxy_advertisement.advertise_on_public

<a id="canonical-1110322033101122-3003022311212312-0022132232200202-1221003100230212-2223131233220202-0031012212233122-3122230130002102-0322200120311221"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
advertise_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102110221103131-0033101323212213-2111111232230231-2123130221121103-0332320132110033-3221112003203131-3230122200232003-0121210012032000"></a>

### Direct properties for `proxy_advertisement.advertise_on_public`

- [public_ip](resources--dns_proxy--reference--group-002.md#canonical-1211001132230302-3112320133313021-3310213102223200-3233220301232120-0223012021000110-2023102323120002-3220210033113312-1112323130102201): complete subsection reference.

<a id="canonical-1211001132230302-3112320133313021-3310213102223200-3233220301232120-0223012021000110-2023102323120002-3220210033113312-1112323130102201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-2220303323213103-2110110021231302-1232003120023021-3003122123002120-0021020320000023-3232200110132031-0222010121032220-2103301122232110)
- proxy_advertisement.advertise_on_public.public_ip

<a id="canonical-3001021030310113-2101133213311130-3331323200230030-1121232033133013-3220300001300301-3321313313311032-3132220030033003-1333201122131330"></a>

Type: `"object"`. single nested block, Optional.

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-1001002122012120-3221020100012311-1211211113223110-3311230311210233-3020112000210121-2033021023322012-2023132001003122-3122221113113011"></a>

### Direct properties for `proxy_advertisement.advertise_on_public.public_ip`

<a id="canonical-2313202322003112-1102301011313203-2311320032022130-2221103210232303-3023232120302031-3223220122231203-1123113301121000-2230020203320313"></a>

#### `proxy_advertisement.advertise_on_public.public_ip.name` property

Type: `"string"`. Optional.

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

<a id="canonical-3100121110222121-0102320103011120-0003323132001220-0321302022202230-1320111201310311-2023332123332223-2330313023331301-3211001112033302"></a>

<a id="canonical-2012102202101201-0331003020010120-1301210331320010-1123001120200110-0013212322321213-2231021012303303-3110011022023303-0120131321313301"></a>

#### `proxy_advertisement.advertise_on_public.public_ip.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-1102230123021113-3230233331221020-1313003133033123-0033330311122023-0230211001222132-3032131130013123-3133023020123313-3111320003231111"></a>

<a id="canonical-2131330111112002-3130330130112001-1023122200110121-2321230022032310-3000233132103133-2020122332301013-3332002213313030-3221002033110310"></a>

#### `proxy_advertisement.advertise_on_public.public_ip.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-3121000120202100-0310113221300320-2102133012003031-0222233000032201-3001032020133122-2001303133003202-3101011113003301-0002310313213122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_on_public_default_dualstack_vip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- proxy_advertisement.advertise_on_public_default_dualstack_vip

<a id="canonical-2010310231010331-1113120123100230-1130133211101302-1323300011022312-1313011332032302-2231232320233211-0231121223022222-1212133303212022"></a>

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
advertise_on_public_default_dualstack_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220312223313100-0021103312020302-1331000023323330-1112111302001112-1233121002000332-3311110001302010-1233221013011120-2221310113023300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_on_public_default_ipv6_vip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- proxy_advertisement.advertise_on_public_default_ipv6_vip

<a id="canonical-1332102031200011-1133321222023301-3310333101100033-2011031112313130-1001012102311000-0022332003011103-1313122030303011-3210312303100033"></a>

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
advertise_on_public_default_ipv6_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222203232301300-1301320120330311-0033333301311202-3100200223211001-3002302013120130-1132130033200011-3033130130111121-3021122101310213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_on_public_default_vip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- proxy_advertisement.advertise_on_public_default_vip

<a id="canonical-2302022130021210-0120233122130103-0221120000331100-0112232110212303-2120231312000302-1323232301203220-2030203122001030-2130001213223123"></a>

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
advertise_on_public_default_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010212203201301-0332200213312212-1220010333030212-2201021203301200-0030321311132322-0001222133122102-1212100023322133-0113000303320123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_v6_on_public` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- proxy_advertisement.advertise_v6_on_public

<a id="canonical-1102121313013002-0112123033211102-0303312200120302-0221001121202211-3103001303023120-1123311230231321-2011302033231102-0320313103333302"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
advertise_v6_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-3231313103131012-0203101212002321-0031200031133131-0011123111310101-0330302030221232-0012232100002333-3202012332110210-2032023321101121"></a>

### Direct properties for `proxy_advertisement.advertise_v6_on_public`

- [public_ip](resources--dns_proxy--reference--group-002.md#canonical-3220333302012002-3032210132331233-1313301000022101-2121121121101312-2121023102010221-2222230202210002-2232301122301321-3232121311223223): complete subsection reference.

<a id="canonical-3220333302012002-3032210132331233-1313301000022101-2121121121101312-2121023102010221-2222230202210002-2232301122301321-3232121311223223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_v6_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-0010212203201301-0332200213312212-1220010333030212-2201021203301200-0030321311132322-0001222133122102-1212100023322133-0113000303320123)
- proxy_advertisement.advertise_v6_on_public.public_ip

<a id="canonical-3300021010030330-3100133331003023-0121323203322311-1312201312012223-1301023321100203-0330230103120023-1103011002032102-2001032110001221"></a>

Type: `"object"`. single nested block, Optional.

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-1232023101000223-0102301312313330-1231300332331123-3201332120031021-1023230133130120-2322001000233330-2321120323312103-3302122132000120"></a>

### Direct properties for `proxy_advertisement.advertise_v6_on_public.public_ip`

<a id="canonical-0133333023031101-0000101021012011-0111120102110222-0202003232121211-0200200300303112-3332020210223220-1113031231320200-1111331330003330"></a>

#### `proxy_advertisement.advertise_v6_on_public.public_ip.name` property

Type: `"string"`. Optional.

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

<a id="canonical-3203131020011022-1121133111110032-3000123031131033-0013111301132233-1001022103323120-2112020112030122-1303311302322230-0333122021200121"></a>

<a id="canonical-2202113330223133-3232303223101212-2221020130300202-0110220010030001-2132120132202310-3220121100121310-0123123330331101-2223203201130311"></a>

#### `proxy_advertisement.advertise_v6_on_public.public_ip.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-3100002311122231-3102221233003330-1110030121121230-3122111111200211-1102313031220033-0032102122101122-3001100320020231-1123030231103222"></a>

<a id="canonical-2202202332233323-2010320310220211-2120302000322211-2202210120112111-1232021023210201-3101232221010222-1020300002320300-2103100311203030"></a>

#### `proxy_advertisement.advertise_v6_on_public.public_ip.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-0323231032132032-3103023010222030-3122002213313020-2200330021120102-0010211000032013-0320332021031131-0020201322001221-1333331311212313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.do_not_advertise` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- proxy_advertisement.do_not_advertise

<a id="canonical-3011330200112223-3321223312121211-1330323001012213-3130112113101221-0301013010310310-2213000202032202-0211201223130122-2310021302333213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for do not advertise.

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
do_not_advertise = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323333303000202-2222033001303102-1032113131211303-1311033002122012-1002203323202112-3112220230111220-3311101211320012-3313320002112211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- timeouts

<a id="canonical-2121133002213012-3323330102321310-3111120012100112-0102120022101130-3121022013002130-1103013032222110-1232313023220331-0300331230201033"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2110232102323221-1022101332310313-3020103111310033-0122000311303222-1313223013120103-1203122101113023-1130201312000131-0000313103130003"></a>

### Direct properties for `timeouts`

<a id="canonical-0200202011322323-3301012301121211-0332101003323232-3020211300301023-2002321300210213-0313220111310002-2301323013213202-3110021122030001"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1022130011203113-2100212123031332-2322101203002233-3211030031322230-2232023030102023-3303322301132131-0230201111033022-3010310001203220"></a>

<a id="canonical-0003031302020022-2012203212220033-2200201311113223-3132022212200011-3111020012113302-0130033202231201-0332020031200312-3101231132110102"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2011132030031233-3212321333011301-0031131131013303-1022211311020122-1312022132030312-1020120131003010-2203321131321232-1020100130020332"></a>

<a id="canonical-1310321331120221-1330010331130011-2331031012212322-3133033230102333-3021200203200331-3320212210232012-3101213133231023-3200231221320123"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3231113033232332-3300021010010210-2011120231102311-1202222300233212-2123111000230301-3220131110112002-2301021133203033-0222020322033210"></a>

<a id="canonical-0333220101322001-3233303231222031-1230103132101203-3020230220212120-1320232223132213-2231113322101201-1313011020233211-2132230000022222"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
