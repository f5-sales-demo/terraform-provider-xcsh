---
page_title: "xcsh_application_profiles reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles reference."
---

# xcsh_application_profiles reference

<a id="canonical-0200132200320020-2100302022210220-0130100121303213-0023111323121203-3300103211030111-2032330112020302-2212320320101122-3301003121002013"></a>

## namespace property — server_ssl_profile / 332102001310 / 6

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

<a id="canonical-2330313112311101-2323331322211132-3032313321111113-2133000113202002-2300220332223110-2230130012233031-2010002010020222-0003033320100202"></a>

<a id="canonical-1102332330020011-0032021303123131-2003332230223302-1011321333200221-3110213111332221-0111313012002323-2020230000300212-1201231311033230"></a>

## tenant property — server_ssl_profile / 332102001310 / 7

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

<a id="canonical-0203003212200112-3310300301310211-2323330302211331-3322110132233122-1003021030333111-1012130322132301-0003003311211121-0001103031022320"></a>

<a id="canonical-0200321120103322-0311002010312231-1131121031220211-1123310033023031-2220321311311013-2331033132231200-3213121323303320-0302023211212033"></a>

## uid property — server_ssl_profile / 332102001310 / 8

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

<a id="canonical-2012200312021123-3231100121210101-1223221213303203-1021003021210332-2120212123023122-1033210200311011-0323033313210233-2030203023201321"></a>

## Next pages — server_ssl_profile / 332102001310 / 9

- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1001310203000122-0311102113233113-3302310000201022-0200230311200102-1233203113320233-0132232313301222-3211300233232121-0102200101100012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102202103102213-1121101130031331-0100221313223312-1112200213313123-1020011301331012-3331032102111203-1331100011010132-1301101123323021"></a>

## virtual_server.http3.tcp_server_profile — tcp_server_profile / 213201333320 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- virtual_server.http3.tcp_server_profile

<a id="canonical-2112022202030023-0100133022230321-0003322233201012-1233302133322131-1221012120113011-1300301322102201-3220232232012300-2313211223230131"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for tcp server profile.

Upstream description:

Configuration parameter for tcp server profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tcp_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3330333011323111-2110301122132231-3120300312022011-2010210102122031-3113032031300120-2130210201321312-3221013220023021-2023112020320202"></a>

## Direct properties — tcp_server_profile / 213201333320 / 3

<a id="canonical-3233212222020113-2013030122200212-1312021020332233-2213131313303312-2022033212012200-1332231110032222-2330112220031320-1203213110201331"></a>

<a id="canonical-3112122113332302-1023203210203232-3232000123013121-3331223122002301-3333121301002312-3000333130100012-1121313200012013-2211202233332302"></a>

## kind property — tcp_server_profile / 213201333320 / 4

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

<a id="canonical-1313331100013012-1030031223311021-3012301331111302-3102010301010222-2331303203022011-0123303021133010-1311032111222033-1321301300000111"></a>

<a id="canonical-1030232320110000-1301002331131311-0023120301322332-1201013301130030-1331003222030001-2113201010313223-0011121232233001-1110330303101111"></a>

## name property — tcp_server_profile / 213201333320 / 5

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

<a id="canonical-1222310012213002-3123101310310103-1303100221030210-3010102203123330-3223033230301311-3312210132221300-0223301222313211-0320132001222301"></a>

<a id="canonical-1220100303321023-0221032211031330-0132313200122313-0011213101120233-3301302103213113-0023003012320022-0232231311313222-3231310012220333"></a>

## namespace property — tcp_server_profile / 213201333320 / 6

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

<a id="canonical-3302321003121013-1303320221101203-3330120030100003-3322220212132311-2131210331022232-2123000202110111-1112211223002332-3120001103303301"></a>

<a id="canonical-2301302012032203-1020101112111222-0011223200122201-1123331002200110-2310313301003101-3210222221302313-2113010232002110-1102200100233201"></a>

## tenant property — tcp_server_profile / 213201333320 / 7

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

<a id="canonical-2011231123300211-1110331203033123-3133331233211012-3302200221020313-0202100131330100-3311102111013213-2021210000213321-1302330013220212"></a>

<a id="canonical-2222301301122230-1000012120030330-2202101121020322-0100121103203221-3231100121011200-1002121123113003-1013333331303230-2313002223320303"></a>

## uid property — tcp_server_profile / 213201333320 / 8

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

<a id="canonical-0133012032210022-2220221013211230-0103102033222021-3020131000100011-2003113313203110-2021010031021032-1223121231302021-2332122210221330"></a>

## Next pages — tcp_server_profile / 213201333320 / 9

- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-3022120130332131-3313210302001322-3231201032200020-1312101002032320-1200300302023322-3130201220303033-0122013032031230-2112001220310130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030120101122012-3312203010212010-2201320033002312-2220002113231021-1033123113103101-0001323102320120-3010103321120313-3121131033233122"></a>

## virtual_server.http3.udp_client_profile — udp_client_profile / 011133033211 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- virtual_server.http3.udp_client_profile

<a id="canonical-1301232231102003-3331013000233123-3210231321202022-1323203122210233-1102103310230133-3132113123220232-2201022010231232-0120001122103011"></a>

Type: `"object"`. list nested block, Optional.

Protocol Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
udp_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221332313000002-0102212310111011-2313013010203032-1230031303230032-1301030132033012-0120000031212311-2112211010333113-2123012000231002"></a>

## Direct properties — udp_client_profile / 011133033211 / 3

<a id="canonical-3013223112100102-3013001111311323-1231002102311223-0011023031023330-3213100330021322-0112130103212321-3230013311130223-1203212231210101"></a>

<a id="canonical-3222213110332200-1022113203102023-2002221333330312-0003123302020321-3333120310330003-2203230321122113-2200332231313301-0133122301311303"></a>

## kind property — udp_client_profile / 011133033211 / 4

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

<a id="canonical-1121320020312332-1110231313033021-3303210021322300-2102203323220310-1021020102232111-3231001101021221-1133231201032011-1322313032310000"></a>

<a id="canonical-0012321123322033-0001321311312311-3202221232200230-3013232002122312-3002321132030013-3030211023021211-1332011203103200-1212123002112100"></a>

## name property — udp_client_profile / 011133033211 / 5

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

<a id="canonical-2322102203111312-1033312323301130-2003312130120211-0130002323321213-2010113120023333-3330323300321233-0112113203200133-1030200000111310"></a>

<a id="canonical-1232021301210313-2132131233111200-2121020103022333-0003310113313231-1023100021123330-1212011100202032-0311131201021330-1201002012200231"></a>

## namespace property — udp_client_profile / 011133033211 / 6

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

<a id="canonical-1103011102221312-1320032331123122-3213233021233110-0223223312220023-2120020230011321-2031131321131120-0222301222320002-3210233123322330"></a>

<a id="canonical-0321330311202320-0032220131133100-1123301110233312-1312002233102022-0123003221101010-1233202101330112-3003333201321020-1201313103303230"></a>

## tenant property — udp_client_profile / 011133033211 / 7

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

<a id="canonical-1011023113133001-0310312311131200-1032221021310230-0213203012203220-0232131313113202-0030201032312003-3311211120310032-3202202133302212"></a>

<a id="canonical-0302212221123331-0203213012102021-1232233110123130-2301223130022220-2312023010021120-0122020203110122-1113003122231201-1200201331121313"></a>

## uid property — udp_client_profile / 011133033211 / 8

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

<a id="canonical-0110120330032011-2203332310003320-0211133000112322-2100123121122031-3100310312121223-0103001201232211-1132100203112312-2003313213222213"></a>

## Next pages — udp_client_profile / 011133033211 / 9

- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1001201132330211-3022323222212131-0312122213332223-0022033331023200-1322321300101231-3303303132220203-0323330033312310-0211000013331213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030023331101110-3112133133031022-0230312012112201-2312310102311310-3232223312002310-0220230101221131-1120211023022223-2032011212001223"></a>

## virtual_server.http3.udp_server_profile — udp_server_profile / 203330302112 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- virtual_server.http3.udp_server_profile

<a id="canonical-2101303221030311-3333130023032321-2133221203131113-0122010021103001-3012322030021323-1021120122000203-0313111000212222-0032022121330320"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for udp server profile.

Upstream description:

Configuration parameter for udp server profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
udp_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202011320330111-3331211210323001-1211123301010003-3133320012201011-2030320312313233-2020222133132311-2301303120300311-1103302231300311"></a>

## Direct properties — udp_server_profile / 203330302112 / 3

<a id="canonical-0000131221231020-1310212223312213-0331013132010110-1311233302130323-3013200313320131-3223120132012132-1120120113023012-1221233222201021"></a>

<a id="canonical-0202020332300322-2233113131113203-0013111102310030-2010221323112201-2201130031103121-1031210201113111-2002112300102312-0321221022120203"></a>

## kind property — udp_server_profile / 203330302112 / 4

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

<a id="canonical-2300230332231111-1013120233103012-0203213202332000-1232000321230111-2212310102023203-2001023000301211-3120203330302323-0330122212330033"></a>

<a id="canonical-1111333233112111-3231331103103231-0123033300010103-0222031331020231-2201113311321120-0223213101330310-1023211131202221-1133322213120121"></a>

## name property — udp_server_profile / 203330302112 / 5

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

<a id="canonical-0003000132010202-3210313322011303-2310323320002310-0322300310101032-2030311333313110-2021322120301300-3203311300000122-3021203332112323"></a>

<a id="canonical-2033301331023021-1322232132132202-3010033210320220-2303032022223121-1021122030332220-3231333030113112-0231323201322232-1330023033123322"></a>

## namespace property — udp_server_profile / 203330302112 / 6

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

<a id="canonical-2222313230003331-2033032001033011-0320333200333303-1130122013322021-0020211312201313-2000301122120130-3122220220333010-2012302012302320"></a>

<a id="canonical-1132013222300202-1202223323210011-2203013103121120-3000313112320203-0022120330131232-0211031221322220-0331311001333331-2221332212121223"></a>

## tenant property — udp_server_profile / 203330302112 / 7

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

<a id="canonical-0203011010032210-2131222022303211-3221221203322223-3230300330201001-1010302121021133-0110123110333002-2232000111331000-0012232023221200"></a>

<a id="canonical-1120220222000310-0013130303233223-1003120220122030-0120132113200130-1130222002112213-0130303111011032-3232022231001200-2032023110231111"></a>

## uid property — udp_server_profile / 203330302112 / 8

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

<a id="canonical-0220331310212022-2102300310311113-1123213200013311-1221332222230120-0130332132132103-2322011133101330-1213300121032123-0131132313032312"></a>

## Next pages — udp_server_profile / 203330302112 / 9

- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222112201011130-0020101021012113-3102201212030313-1333233111132311-0201003113222210-2300011322202221-0221111020233001-3221122303221003"></a>

## virtual_server.https — https / 222032203131 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.https

<a id="canonical-3001023313121321-0102311133101033-0023030120103121-2330130013111333-2123113113011130-0202212203232022-1123300210111113-1023120132333222"></a>

Type: `"object"`. single nested block, Optional.

HTTP profiles.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
https {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110111120323022-0223210123223300-2233130213001023-2220312323202330-3213330232020323-1313300022133003-0131022211201112-2312230110022010"></a>

## Direct properties — https / 222032203131 / 3

- [client_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-0122200003233301-3311211133001102-3103233332021321-0232000221020331-1000213223212230-3220313123122130-2332212200120320-0313001001103301): complete subsection reference.

- [http2_client_profile](resources--application_profiles--reference--group-003.md#canonical-0203301111203233-3202103203121002-0010112113330300-3100300003203113-3210120202012123-3100113221120002-2200223020221303-0211310322130123): complete subsection reference.

- [http2_server_profile](resources--application_profiles--reference--group-003.md#canonical-1132311103330130-1222222201100031-2313002121311120-3101210200020120-3002120332022011-2300231020301001-0312203101301233-2303010131110231): complete subsection reference.

- [http_client_profile](resources--application_profiles--reference--group-003.md#canonical-3212010031101123-1131101022132221-2211033221011012-0033101123321221-0020131203013333-3333213132110323-0221100301132011-1321322002030113): complete subsection reference.

- [http_server_profile](resources--application_profiles--reference--group-003.md#canonical-1122030322310031-0231311121030033-1302200331211212-0333220201121310-1210223321101021-0222202330113220-0221202102111003-1233230322202223): complete subsection reference.

- [ocsp_profile](resources--application_profiles--reference--group-003.md#canonical-3232210112131231-2320022120332210-3322123003301021-2111110123131030-2101301231333021-2232331231321203-3321222020010312-2320211131133203): complete subsection reference.

- [server_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-3210100200202213-3031331313010323-0132303233310220-2032230221331212-1102102133301311-2123021010322233-1231122120102031-3013320331102233): complete subsection reference.

- [stream_profile](resources--application_profiles--reference--group-003.md#canonical-3212022302301301-3103022201102322-3321120203323221-2023331133030122-0110010301200323-3321303211010033-0202220000313202-2322031102333330): complete subsection reference.

- [tcp_client_profile](resources--application_profiles--reference--group-003.md#canonical-3331122100212103-2103303202311112-2323020330102212-3020211133121232-2333002231201330-3111113212023302-1330211031211310-2112013332011103): complete subsection reference.

- [tcp_server_profile](resources--application_profiles--reference--group-003.md#canonical-0022210002333001-3112222323110013-3020033033022233-1111120310313331-1110213213313203-1331130321003201-3100013203212102-1123123203203132): complete subsection reference.

- [websocket_client_profile](resources--application_profiles--reference--group-003.md#canonical-1222221330201310-3101213130310000-2333211331100210-2102023100223033-0030012112211323-2023233213123221-1220103312330003-1013000000133131): complete subsection reference.

- [websocket_server_profile](resources--application_profiles--reference--group-003.md#canonical-0211103332301002-3311220121001320-2202321013331102-0113223312313112-0102220010300023-0213310332210133-3211320101132301-0213110231333311): complete subsection reference.

<a id="canonical-2113103311212313-0002031233002223-0133121021102331-3123330302003000-3131323231213231-0012331331231010-0121210120032311-0002012332311123"></a>

## Next pages — https / 222032203131 / 4

- [virtual_server.https.client_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-0122200003233301-3311211133001102-3103233332021321-0232000221020331-1000213223212230-3220313123122130-2332212200120320-0313001001103301)
- [virtual_server.https.http2_client_profile](resources--application_profiles--reference--group-003.md#canonical-0203301111203233-3202103203121002-0010112113330300-3100300003203113-3210120202012123-3100113221120002-2200223020221303-0211310322130123)
- [virtual_server.https.http2_server_profile](resources--application_profiles--reference--group-003.md#canonical-1132311103330130-1222222201100031-2313002121311120-3101210200020120-3002120332022011-2300231020301001-0312203101301233-2303010131110231)
- [virtual_server.https.http_client_profile](resources--application_profiles--reference--group-003.md#canonical-3212010031101123-1131101022132221-2211033221011012-0033101123321221-0020131203013333-3333213132110323-0221100301132011-1321322002030113)
- [virtual_server.https.http_server_profile](resources--application_profiles--reference--group-003.md#canonical-1122030322310031-0231311121030033-1302200331211212-0333220201121310-1210223321101021-0222202330113220-0221202102111003-1233230322202223)
- [virtual_server.https.ocsp_profile](resources--application_profiles--reference--group-003.md#canonical-3232210112131231-2320022120332210-3322123003301021-2111110123131030-2101301231333021-2232331231321203-3321222020010312-2320211131133203)
- [virtual_server.https.server_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-3210100200202213-3031331313010323-0132303233310220-2032230221331212-1102102133301311-2123021010322233-1231122120102031-3013320331102233)
- [virtual_server.https.stream_profile](resources--application_profiles--reference--group-003.md#canonical-3212022302301301-3103022201102322-3321120203323221-2023331133030122-0110010301200323-3321303211010033-0202220000313202-2322031102333330)
- [virtual_server.https.tcp_client_profile](resources--application_profiles--reference--group-003.md#canonical-3331122100212103-2103303202311112-2323020330102212-3020211133121232-2333002231201330-3111113212023302-1330211031211310-2112013332011103)
- [virtual_server.https.tcp_server_profile](resources--application_profiles--reference--group-003.md#canonical-0022210002333001-3112222323110013-3020033033022233-1111120310313331-1110213213313203-1331130321003201-3100013203212102-1123123203203132)
- [virtual_server.https.websocket_client_profile](resources--application_profiles--reference--group-003.md#canonical-1222221330201310-3101213130310000-2333211331100210-2102023100223033-0030012112211323-2023233213123221-1220103312330003-1013000000133131)
- [virtual_server.https.websocket_server_profile](resources--application_profiles--reference--group-003.md#canonical-0211103332301002-3311220121001320-2202321013331102-0113223312313112-0102220010300023-0213310332210133-3211320101132301-0213110231333311)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0122200003233301-3311211133001102-3103233332021321-0232000221020331-1000213223212230-3220313123122130-2332212200120320-0313001001103301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020020112330332-1113021322333010-0201032120300010-0212320000232210-1332210213213301-1033022210110230-0030222220202020-0333003310012321"></a>

## virtual_server.https.client_ssl_profile — client_ssl_profile / 123123130103 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.client_ssl_profile

<a id="canonical-0310121121023300-0202320010121122-2230211332011001-3202032212312211-1102322000011001-0120031112230010-0003210133022323-3330320012220231"></a>

Type: `"object"`. list nested block, Optional.

Client SSL Profile. Client-side configuration

Upstream description:

Client-side configuration

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
client_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2233222220310103-2110010301333300-2201020021123120-3021110313111220-2203010222031013-1130110030201012-0230213202200131-0303333302232230"></a>

## Direct properties — client_ssl_profile / 123123130103 / 3

<a id="canonical-3032231110110323-1033321021030322-0123301303001130-3223330130020011-3221120132020132-1223221002031223-2210010320200032-3313013112301010"></a>

<a id="canonical-0313023231300030-0101023033320230-0031331210320021-1313123311330032-3112022022332231-0022012131231011-0032223321233312-2333021220212312"></a>

## kind property — client_ssl_profile / 123123130103 / 4

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

<a id="canonical-0003003301000222-0301100210030302-3132310302032101-0220320330301001-3021220010212333-3131332111311002-0220322002310223-3032220023200211"></a>

<a id="canonical-1313120233003232-3111322022111100-3011331001301312-2202121130232122-2300012120111310-3002033011001301-2112311113231311-0231303031311013"></a>

## name property — client_ssl_profile / 123123130103 / 5

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

<a id="canonical-1311323220110203-2133123021312233-2023032031232113-3203001331011010-0132130330102233-1101300101220331-1031203301132233-3201010223100210"></a>

<a id="canonical-3131002031210313-2202330323200032-2010221132320121-1000323301231201-0102300020301011-0120132132111120-0331333112031310-2103221310111020"></a>

## namespace property — client_ssl_profile / 123123130103 / 6

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

<a id="canonical-2022202320302123-1133322223320312-1230332000322123-3022112012223322-2212321330310320-1100322000232323-2221111330332110-0021110023012030"></a>

<a id="canonical-1210301111330132-3022132100303111-2220303213213112-3300230323321223-3010003320003312-2001123302011312-2321220003313120-3132332310123102"></a>

## tenant property — client_ssl_profile / 123123130103 / 7

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

<a id="canonical-2223310102001111-0022200133320100-0231300013011313-0311010133030002-3211032220131323-1313213312231102-3331320220023212-2210302232123231"></a>

<a id="canonical-0033232101012232-1312301220102120-3122330101013303-2022330220230110-2110331130303012-0113203122223131-2303321201011202-0232231212201121"></a>

## uid property — client_ssl_profile / 123123130103 / 8

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

<a id="canonical-0332200131030023-1303201131232200-2122322033122230-3210103310021211-2032212121110210-1323102030302020-0302023333333021-0320121322103102"></a>

## Next pages — client_ssl_profile / 123123130103 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0203301111203233-3202103203121002-0010112113330300-3100300003203113-3210120202012123-3100113221120002-2200223020221303-0211310322130123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203221322100312-1203320200122331-0032331123000012-3133023020033321-1100320301132332-1122011111011000-0312130022033120-1322000212322011"></a>

## virtual_server.https.http2_client_profile — http2_client_profile / 013120012001 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.http2_client_profile

<a id="canonical-1003123300312211-1200121022121330-1013021323113111-0210100223202000-2320003131022321-2223221010130012-2100302320320122-2133311112211120"></a>

Type: `"object"`. list nested block, Optional.

HTTP/2 Profile Client. Client-side configuration

Upstream description:

Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http2_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313003010320201-2230320110221331-2200311312010010-3311021021331001-0100011212313200-1332111200322000-2331101130232303-2122020122201020"></a>

## Direct properties — http2_client_profile / 013120012001 / 3

<a id="canonical-0020210321321021-2030112332221211-2120011022223230-0311333010303030-2330123101122300-1220123112032232-1123022003131133-1331113200330021"></a>

<a id="canonical-1331131212020232-0301231110103012-0323131110111330-1222122033132311-0212300113212003-3313333021231303-2021201122031220-3013213001031102"></a>

## kind property — http2_client_profile / 013120012001 / 4

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

<a id="canonical-3122002103220133-0131023110022301-0321122302302113-0003311233310333-3332032033033013-1032010113313330-0322012321230310-0300222310200331"></a>

<a id="canonical-3323000210231311-1200113221021112-0322111100011020-2230020203123331-3322132301312220-2131112131220332-2023230313321301-0110202102100010"></a>

## name property — http2_client_profile / 013120012001 / 5

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

<a id="canonical-1211331331323102-0210322300233311-3033002012102113-1121231111200120-2102220120111131-1221011020303012-2102131130211330-0123321132001200"></a>

<a id="canonical-0331120220212311-3011322321131032-2332011232033012-0122213012113131-1313133201131133-1321302113333300-3201220202022001-0112011101300320"></a>

## namespace property — http2_client_profile / 013120012001 / 6

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

<a id="canonical-1322300310122122-0003132023032102-1033301113123323-2021133331010012-1131013212021122-2333023300110121-1203302301210210-1302330200320033"></a>

<a id="canonical-2313022222100011-0011130230003120-0220110333111203-3330003031332211-2002201331112223-3322203223021013-1020230111032102-0222332331101021"></a>

## tenant property — http2_client_profile / 013120012001 / 7

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

<a id="canonical-1103130313031020-3223210221132310-2031233210210213-2200232120203323-3001232212322123-3123211010323320-1302003030312312-1001020323122201"></a>

<a id="canonical-3013112203212112-0210223331320333-2323111231133123-0311122013010322-2110333233130203-0322313023130322-1000333110132221-3102300132302110"></a>

## uid property — http2_client_profile / 013120012001 / 8

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

<a id="canonical-3201202111022330-1001221110230121-0231221000301233-2320113321033301-1201102112231221-2210012302123000-3200020320013200-0123123301230202"></a>

## Next pages — http2_client_profile / 013120012001 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1132311103330130-1222222201100031-2313002121311120-3101210200020120-3002120332022011-2300231020301001-0312203101301233-2303010131110231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131311230132033-0030223233120330-1000020123131121-3113310100323231-0223130103230201-3310221031102033-0302221330130131-3000022013122120"></a>

## virtual_server.https.http2_server_profile — http2_server_profile / 313131103112 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.http2_server_profile

<a id="canonical-1301311030022121-3200103330220220-3210233112201320-0010303121231223-0313222113100020-3112130201100122-3322123010223020-1002303231002203"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for http2 server profile.

Upstream description:

Configuration parameter for http2 server profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http2_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0203300221200331-1100130033110313-0012230032321032-1121032321023222-1030331210320302-0333002020230332-2032313130030122-2220031323011323"></a>

## Direct properties — http2_server_profile / 313131103112 / 3

<a id="canonical-1113211231233032-0222313101010330-2331111002231120-0232200310102012-3230200321113121-1320033113303311-1211123020303130-0123033131000123"></a>

<a id="canonical-3302211233203322-1222100030102223-3113320321223110-0012222211313021-1211311112002311-2212133301012120-2133220312300122-1022102012112033"></a>

## kind property — http2_server_profile / 313131103112 / 4

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

<a id="canonical-3023233033020131-1321311130111021-2331012103212023-0013301133200031-2211210012330022-3322300200021012-0021231010122232-3021231002013233"></a>

<a id="canonical-3101310121322123-1000222231312312-3331131322123203-1101300120310320-3210333123232103-0303020001003030-0322030123022222-0032201130223230"></a>

## name property — http2_server_profile / 313131103112 / 5

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

<a id="canonical-2301320332123011-1301313120311322-0000211013023311-0223020320313313-1002031022231321-2133313022023200-2103102102320121-1110311123032110"></a>

<a id="canonical-2102020321301010-2233322120130210-3301101223211201-0133010002200020-2131302133012031-3332313311302022-1032331123303320-0101321320220220"></a>

## namespace property — http2_server_profile / 313131103112 / 6

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

<a id="canonical-0320221013202203-2010221331110003-3203132200302111-3020303301233130-3202231323130001-2023012112122322-0020320110203112-1020232301110003"></a>

<a id="canonical-3313331032102300-1222121312330101-1211301302123300-2013202001100310-2020101131010003-0003210321020301-0001133303123120-3323300320320113"></a>

## tenant property — http2_server_profile / 313131103112 / 7

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

<a id="canonical-0103301103002230-0031310120131100-0331022101323301-3211200320000320-0110100321023120-2333130113200321-1310221113101303-3300232120132233"></a>

<a id="canonical-2232031233302322-3021301303103222-2102202023132321-0101011021032220-3331020303230231-2132230101210003-3332003102211302-2320310010210203"></a>

## uid property — http2_server_profile / 313131103112 / 8

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

<a id="canonical-1101100212023030-2112210031100303-3201312301120122-2303303322021100-1300133333123001-0110320201210103-2001111121201321-3000110033331113"></a>

## Next pages — http2_server_profile / 313131103112 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-3212010031101123-1131101022132221-2211033221011012-0033101123321221-0020131203013333-3333213132110323-0221100301132011-1321322002030113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132321303310022-0100012322233021-1112122013130313-1322203230311001-1320113312331101-0030111000121232-3022222133210020-0321232221303120"></a>

## virtual_server.https.http_client_profile — http_client_profile / 003000331023 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.http_client_profile

<a id="canonical-0032323313122011-1113013231233012-1211201132213211-1120003003232021-1113331031023332-2003213113121200-0303213030113103-0221312101011203"></a>

Type: `"object"`. list nested block, Optional.

HTTP Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100010221223002-3202330232023011-3323302201320130-3123202133113033-2001122230002332-2030101222300133-1100012313033132-2300230213331233"></a>

## Direct properties — http_client_profile / 003000331023 / 3

<a id="canonical-2023113210200000-0200112222213111-0231203130300021-0131103012301303-1233202032321121-2222322012120130-1300000323021303-2103213221211221"></a>

<a id="canonical-3100231323123321-3022112023030132-1120213313112212-3320330212103002-3011223103102010-3112212003123122-3202333231210111-2221231100331111"></a>

## kind property — http_client_profile / 003000331023 / 4

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

<a id="canonical-0303323220213020-2000222113103230-0223320012112210-3230021003123010-3131020110011331-3001220030012321-1131211301030010-1331231133232233"></a>

<a id="canonical-3202303112300003-1003023323201133-3031222032202110-0202200021103102-2221121100221013-2110120120121131-1031122232201020-0002313100320332"></a>

## name property — http_client_profile / 003000331023 / 5

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

<a id="canonical-0022233331201131-2312323210310333-1233301230322012-3231321322022312-3120100003010302-3002200323200103-0220321012113221-0301132013330101"></a>

<a id="canonical-0203330103213220-1311113320123211-0132101032123300-1001312011232001-3111000123301103-1122013310203212-0112231220013201-3213310300311021"></a>

## namespace property — http_client_profile / 003000331023 / 6

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

<a id="canonical-2021313123201301-3203133211212100-1320023310123102-0211302230010300-0323012020300220-2130133103331020-0313321003222233-1103022003202320"></a>

<a id="canonical-0012211300012222-3022212012132030-2300313232231100-1001023133121301-1102033031210201-2300303021310230-3322120131323010-2022323111030120"></a>

## tenant property — http_client_profile / 003000331023 / 7

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

<a id="canonical-2233033110310331-2303202013321113-3112300221103022-0223232133221001-0130323302312003-2013100003333210-1103321033320232-0322230010012321"></a>

<a id="canonical-2122233012031311-2010003203032012-3022220130020002-3132112231220302-2221220130013331-2123302131031332-0022200023303231-0313000011300030"></a>

## uid property — http_client_profile / 003000331023 / 8

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

<a id="canonical-3120220202332220-1313311011113332-3220201221033322-2231333100222322-2331223331231310-1222313113212112-0300333301130030-1232132030100032"></a>

## Next pages — http_client_profile / 003000331023 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1122030322310031-0231311121030033-1302200331211212-0333220201121310-1210223321101021-0222202330113220-0221202102111003-1233230322202223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220213223111123-0020133333103220-3133032201212311-0323310220310132-1213213230022132-1223333223122211-3020010210320002-1210213121023322"></a>

## virtual_server.https.http_server_profile — http_server_profile / 213301323320 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.http_server_profile

<a id="canonical-0022113000123221-1310311301020202-3033300230312220-3112103313120203-3220123203100231-2233231201330130-2013212103023230-2303102003312012"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for http server profile.

Upstream description:

Configuration parameter for http server profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322231300100322-0203233113210331-3213133323312202-0112300310012230-1202332202310202-0233100131310000-0331233201123002-3330331101321101"></a>

## Direct properties — http_server_profile / 213301323320 / 3

<a id="canonical-3203212101112002-2302130332100001-2312322011031023-3123133031311232-0031333131031233-1102130110320233-0002102301013232-2202131022300202"></a>

<a id="canonical-1011103033130102-2303032133010230-2010323030020002-0022023223330123-1221013000123232-2300333130302323-3121000312033110-2111120200223103"></a>

## kind property — http_server_profile / 213301323320 / 4

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

<a id="canonical-0203320031122111-0012323323321330-0232021203130212-1333013233130302-3221201132300102-2112331122111320-1320322122120031-1303201303101220"></a>

<a id="canonical-1030033011102312-1230011310331213-2202223133311321-2123003333211230-2222301221330201-0023230222311201-0230332021120203-3233211230030221"></a>

## name property — http_server_profile / 213301323320 / 5

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

<a id="canonical-0301311010130332-2013223110331003-2302223131330010-1000133101312133-3033012102320203-1313133132032131-2300021120322213-3333203323131311"></a>

<a id="canonical-1221200322001112-2023300130001220-0330033112000221-0320322221013000-2202201323311121-0330032021012213-0202102001100212-1033131322220230"></a>

## namespace property — http_server_profile / 213301323320 / 6

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

<a id="canonical-1033322130333333-3310212312100302-1222121103220102-3011321303001233-2010212113130103-0210303033332102-2001213320123001-0120101101032032"></a>

<a id="canonical-0101230323022300-1332331211000303-1210010330011110-2113011231020130-2301011321330331-0123013210230130-3203332212213313-3210231203120123"></a>

## tenant property — http_server_profile / 213301323320 / 7

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

<a id="canonical-0332102231332221-1010303130330101-1123001332332033-2302311132213211-2233102332022310-3332221230333310-1033323303233022-0030111223231310"></a>

<a id="canonical-3322013313231213-2230231330331301-0102123231201332-2312231320221000-1131320001333203-0031130033330311-2002022030030302-3033030321322022"></a>

## uid property — http_server_profile / 213301323320 / 8

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

<a id="canonical-2100220101030002-3331031021331002-3132320120002330-2021022122321013-0130212013013132-3333111031133312-0212003231221301-3030033010300311"></a>

## Next pages — http_server_profile / 213301323320 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-3232210112131231-2320022120332210-3322123003301021-2111110123131030-2101301231333021-2232331231321203-3321222020010312-2320211131133203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331023303222110-2213211022012122-2300312313221210-2322230333231131-0100012022032313-2200300212211033-2010313233122002-3030100031333110"></a>

## virtual_server.https.ocsp_profile — ocsp_profile / 211013011122 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.ocsp_profile

<a id="canonical-1132012023312201-2221332332121311-0112130333130300-0020032030200233-3313013312303212-1330022213030112-2313113303201211-0103121230002203"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for ocsp profile.

Upstream description:

Configuration parameter for ocsp profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
ocsp_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102013200201200-0030133102100100-0203302222031113-3011000301121202-2311102301331212-2221200330010213-2132332010333332-1321231103031132"></a>

## Direct properties — ocsp_profile / 211013011122 / 3

<a id="canonical-0331303310021132-3120032200223020-3333010133320103-0333233232201121-2322310013000031-2110223020321130-2201332003010223-2321120231031112"></a>

<a id="canonical-3022330332323032-1012330023323013-1332113320003111-1303233120032221-3233120301330321-3320013200220311-2320112222321122-3001302030012130"></a>

## kind property — ocsp_profile / 211013011122 / 4

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

<a id="canonical-3230223322021133-0322233321110023-2103131011132232-0230121003000232-3033332313333333-1023332232331022-2311212200220013-3003223211220211"></a>

<a id="canonical-3133331112131230-1012001031211103-3313223211313122-3001030302321112-1001301013230302-0000202220000112-0111300220233013-0012032023230112"></a>

## name property — ocsp_profile / 211013011122 / 5

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

<a id="canonical-0200000111002232-0131311002102013-1013002000322012-3301003301302030-2210110001102331-0232022112113211-1320020222210102-0323100232321312"></a>

<a id="canonical-1311121011132013-2000323122202120-0210002132221021-2222103213311030-0213110021202203-1110220100131311-1121302020201022-1210222130100200"></a>

## namespace property — ocsp_profile / 211013011122 / 6

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

<a id="canonical-2022323331132330-0120331103211023-0233322101011010-0311120122030322-2223033322002110-2002132003022002-2033223321312220-2011311222000203"></a>

<a id="canonical-0300120021332030-1301321312310322-1301023022213200-0303220130313322-3220202232210331-0302330103313320-1330132020021002-3322330220013320"></a>

## tenant property — ocsp_profile / 211013011122 / 7

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

<a id="canonical-0201011233111201-3003223310323100-0301020223011200-0223333112123322-1120321301331013-0222011321001200-2310330030023103-0321301122013302"></a>

<a id="canonical-3220023313021112-1301121100203123-3031120220311002-1312233132130302-2020000030110221-2013310131011313-2020101113200012-1301003031002202"></a>

## uid property — ocsp_profile / 211013011122 / 8

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

<a id="canonical-1012023033202220-0333032232131331-3302112003120130-0033320223232301-2130310331103330-3300113031301113-3010103033231213-1332032323313310"></a>

## Next pages — ocsp_profile / 211013011122 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-3210100200202213-3031331313010323-0132303233310220-2032230221331212-1102102133301311-2123021010322233-1231122120102031-3013320331102233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033301201123011-1212103310231101-3031311103212023-3310013010120313-2331320331110300-3231020220212333-3203311201022111-1103303020011120"></a>

## virtual_server.https.server_ssl_profile — server_ssl_profile / 003300110333 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.server_ssl_profile

<a id="canonical-1302131030322013-3023030122211333-1303321201211310-3023221200331023-3301001001112232-3322321222032130-3131022020013032-2011300332200312"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for server SSL profile.

Upstream description:

Configuration parameter for server SSL profile

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
server_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1223320122201211-3221132210331332-0132233011222312-1021213100303133-1002320120112230-0202031021332302-3000222320021021-2013130331003100"></a>

## Direct properties — server_ssl_profile / 003300110333 / 3

<a id="canonical-0121330131230122-2132203220311032-2203212221212202-1231331211233321-1232022031133022-2010321122311320-3301002322212330-2113103130122030"></a>

<a id="canonical-3112102122300033-0103222311213002-0103211132311201-2310102033030312-2330121112223030-1211230020233110-2233033300220311-0011113002303210"></a>

## kind property — server_ssl_profile / 003300110333 / 4

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

<a id="canonical-3011030022032230-1232001313203022-1020123021201331-2022011220011003-2003220100103320-1222302222100331-3302101302312002-0020300312000130"></a>

<a id="canonical-2000311303222113-1032222300022311-2303002300032322-1033201212230013-2230130032200201-0301112103113310-2331330321321300-1111023332311231"></a>

## name property — server_ssl_profile / 003300110333 / 5

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

<a id="canonical-0011221132001012-1112311330330132-0322031302020303-2002211120002132-0012323003303223-2323003000001203-1211210300223012-0030333133311120"></a>

<a id="canonical-3313303021203321-1211310030220003-0032113333301022-3222313001131103-1001033110102111-1031221022133301-2122003230011311-0032332222011133"></a>

## namespace property — server_ssl_profile / 003300110333 / 6

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

<a id="canonical-0310330122030031-0101222311320002-1030110132331323-2303313130222120-3111101011303331-0121313301233201-1320011100133302-1013032111302132"></a>

<a id="canonical-1001012101211000-1212121130320131-0123203210230103-1120300301100001-0302112223313323-0301303220202001-3210232133120233-3032100220312230"></a>

## tenant property — server_ssl_profile / 003300110333 / 7

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

<a id="canonical-0302101201032331-3220323320312022-0010120103003331-0312000331223203-3101231030132310-3113322310313322-3310222102210333-1220222303022232"></a>

<a id="canonical-0232300330013121-1323013213311320-0322333321320303-1012333223013233-0122313210020030-3011132303332100-1120230023232331-1330032232130202"></a>

## uid property — server_ssl_profile / 003300110333 / 8

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

<a id="canonical-1333211221212121-3101120012221031-1203331231323232-3220232020113210-1221003002231331-1033231301011030-0320000302310013-2233102003313130"></a>

## Next pages — server_ssl_profile / 003300110333 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-3212022302301301-3103022201102322-3321120203323221-2023331133030122-0110010301200323-3321303211010033-0202220000313202-2322031102333330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322312322333301-0110222023211211-2310101210202203-2221130223201121-2200101311101012-0303222122033022-3020121303023123-1131122113213020"></a>

## virtual_server.https.stream_profile — stream_profile / 213200213212 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.stream_profile

<a id="canonical-1323133221230101-1211220021102331-0311111233130122-3331120312301222-3323200113222330-1002010112200121-2320201121033111-1031323222012022"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for stream profile.

Upstream description:

Configuration parameter for stream profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
stream_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0113023032203122-1020100130323020-3210312030002213-1010232202301011-0023312300302310-1011212333122211-3112222220223010-2301201213002103"></a>

## Direct properties — stream_profile / 213200213212 / 3

<a id="canonical-1100212030322331-1233021212131001-3203103033222110-0312311010323000-3003033133101213-0032013230210332-2011103220222221-3300000123311330"></a>

<a id="canonical-0122323021123100-3103331321012020-2020133023331312-2012221111223100-1101021100200022-1013203210033113-2132230313211012-2121101130221213"></a>

## kind property — stream_profile / 213200213212 / 4

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

<a id="canonical-3211331103210323-2221332333220231-3201101220133002-3003330330201213-3332332132020110-0003231230031121-2011321101120303-3032200311112121"></a>

<a id="canonical-2011223132030222-1331131020322212-0012033033123303-1313003221302323-3013210213221002-1313322332113022-2121032301211330-0301321032312333"></a>

## name property — stream_profile / 213200213212 / 5

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

<a id="canonical-3101131301333110-2002313032232303-3321100022111132-1222010332102301-0102222223130221-0330330012032010-0210332100223200-1030130331021130"></a>

<a id="canonical-3322203102210202-2221003133023220-3332203013122211-0013333123220312-3133023321203320-3100110213130120-0211131010113111-0222023211223303"></a>

## namespace property — stream_profile / 213200213212 / 6

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

<a id="canonical-2303132231321321-2323232030330122-0033333030001002-3231312310010231-3120002201121233-0313033030003223-3303322230230330-0111332030130232"></a>

<a id="canonical-3322012220123323-0310221222112220-1332303221121212-2023233311330200-1223033033232013-3223033003003112-1020010222001002-3333213301021030"></a>

## tenant property — stream_profile / 213200213212 / 7

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

<a id="canonical-2300211022301120-1311323011231211-2332020310032031-2023203333203001-2101302313213330-1222121123103302-2310010011232003-2301123322331020"></a>

<a id="canonical-0030020301023231-0200303113323102-0322002120131133-3032333333002133-3010223110212122-3100310010020302-0332130102312120-1200200323121130"></a>

## uid property — stream_profile / 213200213212 / 8

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

<a id="canonical-0132033223310123-2301101022121201-0032233302112130-2001332100010122-0230101300030122-1301232212032213-2132120230121220-3131011213320202"></a>

## Next pages — stream_profile / 213200213212 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-3331122100212103-2103303202311112-2323020330102212-3020211133121232-2333002231201330-3111113212023302-1330211031211310-2112013332011103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133302302201112-1211223213203001-2302110130221203-0331023311203333-2302030120320000-1020311113021111-1321303121233012-2210221210102001"></a>

## virtual_server.https.tcp_client_profile — tcp_client_profile / 032323200120 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.tcp_client_profile

<a id="canonical-3303020301320113-0013112032021200-1322222023000302-2120131003333003-0321131323223231-3111231031023213-3310322101031110-3311133000000013"></a>

Type: `"object"`. list nested block, Optional.

Protocol Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tcp_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212331021202130-3222202121001102-0131302111231230-2331123302032221-2132121030302131-0112113130201100-3110120302101332-2313333102003230"></a>

## Direct properties — tcp_client_profile / 032323200120 / 3

<a id="canonical-3032300212330101-2020003032001222-2321233223030001-0031220233120212-1033211301113110-0130301220303322-0002023030331023-3232012223031132"></a>

<a id="canonical-1202023110113202-3030001302313333-1032110001200102-2132232332013200-0320022230130033-1101012212301233-2001131020110210-2211332202323312"></a>

## kind property — tcp_client_profile / 032323200120 / 4

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

<a id="canonical-0201133010110031-2110033022201102-3032130132113332-0321133230123002-1211121022122113-2230033223023021-2032311302311230-3010331321330103"></a>

<a id="canonical-0010113310002032-2022002230220133-0302033110033210-1032030013312113-1123021122223000-1100132113201211-1323311202010132-2223121201132322"></a>

## name property — tcp_client_profile / 032323200120 / 5

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

<a id="canonical-2003221130112030-2123320021131023-1001330331110302-2311013101330310-2033020220213132-1102200102312230-0101121120323213-2300121302010132"></a>

<a id="canonical-2023310102210332-2121002131222102-0132032220011132-3321300232102323-3021211222310313-3332322222230300-0022003010233331-3103020012223123"></a>

## namespace property — tcp_client_profile / 032323200120 / 6

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

<a id="canonical-3022200102020212-2102312321101213-3332122213120300-3030201033213310-1303231031001323-3123211112200022-0103303321333321-3331211000121300"></a>

<a id="canonical-1132123232013311-1033322300102323-2003303112332300-2100230300000202-0213231210123301-2220210022223003-3313113111021210-1033001100032111"></a>

## tenant property — tcp_client_profile / 032323200120 / 7

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

<a id="canonical-1132330013210000-2001200002021010-2203022121320113-0321020220212100-0230320312221200-3020202033303133-3220122133111023-0031301233003033"></a>

<a id="canonical-2233102133302110-3333011000110211-0303031033110101-0212230003023000-1331132033232122-2330030331321121-2333312113121233-2200030203323211"></a>

## uid property — tcp_client_profile / 032323200120 / 8

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

<a id="canonical-1232103023101301-1213221011033031-2202330331212100-3230330321032312-2210210123013122-2012313230222010-2211202031222303-0302123101013330"></a>

## Next pages — tcp_client_profile / 032323200120 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0022210002333001-3112222323110013-3020033033022233-1111120310313331-1110213213313203-1331130321003201-3100013203212102-1123123203203132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303010100100322-3320310030321112-2102203222201121-0320103222330323-2121120322320010-2012211130211321-0301110213203020-0331232122112130"></a>

## virtual_server.https.tcp_server_profile — tcp_server_profile / 113123121103 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.tcp_server_profile

<a id="canonical-2210332031001101-3032113311222310-0111231003010232-0333222211331022-3130220330300311-0323123333023213-3122230031130333-0322301330020202"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for tcp server profile.

Upstream description:

Configuration parameter for tcp server profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tcp_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112301101310031-0301102233313021-2220331310220031-3320122233303223-2002313303203323-2120010100330212-1330112121230312-0300333013331323"></a>

## Direct properties — tcp_server_profile / 113123121103 / 3

<a id="canonical-2123101320202333-1230313111120003-3321213300233003-0333021020221313-0122301022221021-2011302311110330-2310110030312331-1202330211330231"></a>

<a id="canonical-2112020122102002-1231023012202001-0232103033111123-1232010112110233-1301332112302121-2002333031303322-0301033220102322-1212220212310303"></a>

## kind property — tcp_server_profile / 113123121103 / 4

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

<a id="canonical-1211213332131223-3321032031122311-1302230121123021-1110132313232011-3021320321330033-0230000301011221-3023032210133220-0033132122300102"></a>

<a id="canonical-2311131013103311-3032303032230222-2321300002232312-1232223322011211-1223033002030010-0113322001023231-1131220333001213-2111010122213222"></a>

## name property — tcp_server_profile / 113123121103 / 5

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

<a id="canonical-0223010133022003-0003310213232130-2003103333303231-0330000211011332-1331110033200231-2113210322321010-2011001012221320-2200330310211322"></a>

<a id="canonical-3003013123030200-0012113031033202-1033313001200033-2210301230030031-2323220202022313-0103222130201133-2032200203211013-0231121102133032"></a>

## namespace property — tcp_server_profile / 113123121103 / 6

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

<a id="canonical-3332301021033311-2001013131001110-1310310331323133-0333313333111103-1020230100232310-3213023121121103-1233001321330130-2122023113332033"></a>

<a id="canonical-2211202303130221-2021221122012212-3012303000103102-0213030033131313-3022233332130013-2130123303113101-0210330033233110-0020203231132300"></a>

## tenant property — tcp_server_profile / 113123121103 / 7

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

<a id="canonical-2303310310321332-2011003112300000-2223123033322230-0021233113230202-1111111213121203-1000020113011011-1300203323021233-0101033021023022"></a>

<a id="canonical-1211003222001033-2002320023011002-0123123221312320-3210110230223010-2322201001132100-2032223121010231-1203101033120033-0100133331323013"></a>

## uid property — tcp_server_profile / 113123121103 / 8

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

<a id="canonical-3211110201023301-1323123223023321-2130320103300133-0122111332033202-1311123211100102-3010311213003302-1102201301211203-3311330033022001"></a>

## Next pages — tcp_server_profile / 113123121103 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1222221330201310-3101213130310000-2333211331100210-2102023100223033-0030012112211323-2023233213123221-1220103312330003-1013000000133131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100100101112123-3001211322013031-2012322323122002-0102303301332223-1332022312313022-3312032233312203-1300230233222031-1103131002201102"></a>

## virtual_server.https.websocket_client_profile — websocket_client_profile / 010223131103 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.websocket_client_profile

<a id="canonical-0321132123300232-1313031133033202-3120303020320213-2103110113311202-3101300112300221-2302231110003323-3111332212002330-0213011031011312"></a>

Type: `"object"`. list nested block, Optional.

WebSocket Profile Client. Web-related configuration

Upstream description:

Web-related configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
websocket_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020103032232202-1103222311122021-0103300113123322-3032233200113012-0101023022203223-0130210300032331-2220313230131132-3112032321313303"></a>

## Direct properties — websocket_client_profile / 010223131103 / 3

<a id="canonical-1230101103122213-2220322121321111-3110001101012202-0213003332023021-0001321201121302-0231013122310220-3332231112010002-2203130122311302"></a>

<a id="canonical-1322122220300213-3101311103231103-2213200232222221-3103123133111103-3013200101232202-0103023130003003-3211100301001320-0201002221322211"></a>

## kind property — websocket_client_profile / 010223131103 / 4

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

<a id="canonical-3033102210031010-2220102133122020-0131311020202300-1032310000002131-2000132121113122-1331110023120303-2132220233022210-0312331110020210"></a>

<a id="canonical-2312300102012222-3222231030233202-0202300112102331-2023100211333122-0020030102200113-1022230010201202-3120130030301222-1100113102132332"></a>

## name property — websocket_client_profile / 010223131103 / 5

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

<a id="canonical-1223302120121302-2000212220312303-1013001023021311-1330223230130130-1300112122300003-1020220232230130-3110333110113001-0210310321100201"></a>

<a id="canonical-2203312001201102-1112132103013231-1010032120030122-0332111130322321-2303001101103332-1220301003031120-0301113120012230-1320022012002332"></a>

## namespace property — websocket_client_profile / 010223131103 / 6

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

<a id="canonical-0031233030311313-3233012110320121-2133020221203023-0321220300100330-3000010113121113-1231133223320003-3131320303322210-3133300213302100"></a>

<a id="canonical-2322220110020023-2130033201030233-2133333033121210-1323323222102131-0113210331303011-2230220330330023-2111020333100212-0022323110101033"></a>

## tenant property — websocket_client_profile / 010223131103 / 7

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

<a id="canonical-0202100213232022-2112002130110032-1233320311121300-3300312012203332-2322110313201220-0232102103220211-2120313022103033-0302001030300111"></a>

<a id="canonical-0102130133231323-0202300031220022-0211002311311300-3100300022302103-1002030211113202-3130330011130332-0331230133320110-0330032122121030"></a>

## uid property — websocket_client_profile / 010223131103 / 8

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

<a id="canonical-0323122313122302-1331210312012110-3233122222302311-1023000222222103-1201121322330313-0203322112003011-1232020013322211-2211200130102123"></a>

## Next pages — websocket_client_profile / 010223131103 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0211103332301002-3311220121001320-2202321013331102-0113223312313112-0102220010300023-0213310332210133-3211320101132301-0213110231333311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131032303323202-1331101231123212-0101222032233311-3132003121132213-0311321220321320-0103320113213213-3232201131221130-3303322000000203"></a>

## virtual_server.https.websocket_server_profile — websocket_server_profile / 201132221233 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.websocket_server_profile

<a id="canonical-1211103303022013-3011120031202312-3312010310110313-0331213122130111-0103113120321002-0100023210232022-3213113010130310-3303210301033233"></a>

Type: `"object"`. list nested block, Optional.

WebSocket Profile Server. Web-related configuration

Upstream description:

Web-related configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
websocket_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102130133302133-2103222022302233-3023321202203301-3333100332220020-2201013300333202-3021010331313330-2321321321130222-0112023000013332"></a>

## Direct properties — websocket_server_profile / 201132221233 / 3

<a id="canonical-2002132132300230-0201310132212033-1322112023211312-3023130003100103-0032230023013302-3202130322133201-0032303112130113-3220211000320232"></a>

<a id="canonical-3022312302000321-2311200131123121-2101111132222313-2230331330322202-0100100112130000-2030312230310031-0002133023202003-3033200103213322"></a>

## kind property — websocket_server_profile / 201132221233 / 4

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

<a id="canonical-3103101222123011-0220212113302331-0203032132132012-3213123211011220-2011233033232100-2131022112212000-1123201031011012-2303232003210310"></a>

<a id="canonical-0320301223233300-0233300213323023-1100132030033021-0320111231210022-2210130323122313-1132300312330221-3121212231333312-2222332310002232"></a>

## name property — websocket_server_profile / 201132221233 / 5

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

<a id="canonical-0213002233302131-0000211332221202-3301223000223001-3332223012330000-3312002103321233-3033122021300301-0201032310203311-0332030202202133"></a>

<a id="canonical-1233002210103212-3301212310121222-3313131000233002-1112100130221330-1330023033313231-1213312111011023-2212332013203213-2101333302220222"></a>

## namespace property — websocket_server_profile / 201132221233 / 6

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

<a id="canonical-1023131000113303-0330032322331212-3210211211332231-2003202330121003-3301000110113230-2220122100102220-2330201220000131-3011121012130012"></a>

<a id="canonical-1202202131000001-0203110203131322-0113120201113221-0121133333131003-1213132030021200-0301102232121333-1022322232023013-3320010020121223"></a>

## tenant property — websocket_server_profile / 201132221233 / 7

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

<a id="canonical-3131223303321333-0012221311011021-3031000011303211-0111313020323322-1030020110220032-3133220121310033-2231220333020022-2312010310032002"></a>

<a id="canonical-1110202332320321-1221311120200212-2022323221313320-1012212302332121-1212211003110300-2231031103213211-3333231310231003-1003213003300030"></a>

## uid property — websocket_server_profile / 201132221233 / 8

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

<a id="canonical-2010003120032122-3132032131223002-2300320302013130-1112130102311132-2000032003331322-3000221330320013-3301331133322201-3210313212131100"></a>

## Next pages — websocket_server_profile / 201132221233 / 9

- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2120313011001333-0332132133112211-2032113202213200-1133333320210120-0120003121121103-1110233000210030-2313321130012321-0301012230001201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232133220020001-3001001333212220-0331211120033222-0210001311100202-2110213220302011-0321323020021202-2003020020233233-3033202322103301"></a>

## virtual_server.immediate_action_on_service_down — immediate_action_on_service_down / 233200012303 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.immediate_action_on_service_down

<a id="canonical-3230213032221222-0013333133210033-1222132013313112-2322100010110313-1330313131112110-0230121231012010-0232033001322133-0310202132120231"></a>

Type: `"object"`. single nested block, Optional.

Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial
client's SYN packet, if the availability status of the virtual server is Offline or Unavailable.
This is supported for the virtual server of Standard type and TCP protocol. The default is None.

Upstream description:

Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial
client's SYN packet, if the availability status of the virtual server is Offline or Unavailable.
This is supported for the virtual server of Standard type and TCP protocol. The default is None.
None: Specifies that the system takes no immediate action if the virtual server is reported Offline
or Unavailable. Reset: Specifies that the system resets the connections when the virtual server is
reported Offline or Unavailable. Drop: Specifies that the system drops the connections when the
virtual server is reported Offline or Unavailable.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("immediate_action_on_service_down_drop",
    "immediate_action_on_service_down_none"),
  validators.ConflictingObjectAttributes("immediate_action_on_service_down_drop",
    "immediate_action_on_service_down_reset"),
  validators.ConflictingObjectAttributes("immediate_action_on_service_down_none",
    "immediate_action_on_service_down_reset")}
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
  "x-ves-oneof-field-immediate_action_on_service_down_choice": "[\"immediate_action_on_service_down_drop\",\"immediate_action_on_service_down_none\",\"immediate_action_on_service_down_reset\"]"
}
```

Terraform syntax:

```terraform
immediate_action_on_service_down {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202332021202200-3121122312000310-2220022220312132-3303323001212331-2122110210011332-3302211332122011-0312302233232231-1120121310212000"></a>

## Direct properties — immediate_action_on_service_down / 233200012303 / 3

- [immediate_action_on_service_down_drop](resources--application_profiles--reference--group-003.md#canonical-3331020103122320-2130223203321331-2022122221212103-2201303102023313-2011113333303021-1010110302202213-3330201123212123-2321310322212122): complete subsection reference.

- [immediate_action_on_service_down_none](resources--application_profiles--reference--group-003.md#canonical-1211320131313223-0300030133102303-1120303202100033-0302033120121202-0211200222132132-2221120100211201-0332212332001102-2320302000100222): complete subsection reference.

- [immediate_action_on_service_down_reset](resources--application_profiles--reference--group-003.md#canonical-1020323002022302-2031200012130221-0130123012221301-0303002023310122-1011303313130223-0213220032010232-0220130012302202-3020330121013201): complete subsection reference.

<a id="canonical-3122122323101020-2003001200113111-0302313210113231-0111213032103133-2031012121300332-3302232110323212-1220132200123231-3333210201002123"></a>

## Next pages — immediate_action_on_service_down / 233200012303 / 4

- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop](resources--application_profiles--reference--group-003.md#canonical-3331020103122320-2130223203321331-2022122221212103-2201303102023313-2011113333303021-1010110302202213-3330201123212123-2321310322212122)
- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none](resources--application_profiles--reference--group-003.md#canonical-1211320131313223-0300030133102303-1120303202100033-0302033120121202-0211200222132132-2221120100211201-0332212332001102-2320302000100222)
- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset](resources--application_profiles--reference--group-003.md#canonical-1020323002022302-2031200012130221-0130123012221301-0303002023310122-1011303313130223-0213220032010232-0220130012302202-3020330121013201)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-3331020103122320-2130223203321331-2022122221212103-2201303102023313-2011113333303021-1010110302202213-3330201123212123-2321310322212122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213012131012303-3031020223303012-0332010322001132-0012213303223201-0201032100210030-0111002010332221-1121011030310322-1110310312131232"></a>

## virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop — immediate_action_on_service_down_drop / 331300020210 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-003.md#canonical-2120313011001333-0332132133112211-2032113202213200-1133333320210120-0120003121121103-1110233000210030-2313321130012321-0301012230001201)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop

<a id="canonical-1221013321201031-0111102300222301-3311010131213030-0001200212301332-0200022011103130-1230030322301001-1210102033020203-0310010232233012"></a>

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
immediate_action_on_service_down_drop = {}
```

<a id="canonical-3130203213020022-3332102210123322-2313223123000233-3123220020223210-1300331333222332-3320002100201130-2030231010230111-1002030203310020"></a>

## Direct properties — immediate_action_on_service_down_drop / 331300020210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322021101323011-3223023211203132-1122210330113100-2112223303000213-0232303110300023-2230303000211223-3212233310122332-3203021220300133"></a>

## Next pages — immediate_action_on_service_down_drop / 331300020210 / 4

- [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-003.md#canonical-2120313011001333-0332132133112211-2032113202213200-1133333320210120-0120003121121103-1110233000210030-2313321130012321-0301012230001201)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1211320131313223-0300030133102303-1120303202100033-0302033120121202-0211200222132132-2221120100211201-0332212332001102-2320302000100222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203231321320223-2333103201333203-2323132313111322-3303033230201031-3223113201121312-2032133132021311-1321033333323230-0003303313221001"></a>

## virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none — immediate_action_on_service_down_none / 110000302130 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-003.md#canonical-2120313011001333-0332132133112211-2032113202213200-1133333320210120-0120003121121103-1110233000210030-2313321130012321-0301012230001201)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none

<a id="canonical-2003131000131233-3100223321020233-1201100121133302-3200203320011111-0112132020230001-2020331313201112-0001223210133112-2031122322303203"></a>

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
immediate_action_on_service_down_none = {}
```

<a id="canonical-3120112133231220-0313130233033203-3132111002203210-2131030312000303-3320312003033132-0032330300331321-3301322321200212-3101222031302313"></a>

## Direct properties — immediate_action_on_service_down_none / 110000302130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200332122130002-0021000301222110-1010300313303203-0110101033023300-1221130110100220-0213013310312011-3201012001032333-0013111300001212"></a>

## Next pages — immediate_action_on_service_down_none / 110000302130 / 4

- [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-003.md#canonical-2120313011001333-0332132133112211-2032113202213200-1133333320210120-0120003121121103-1110233000210030-2313321130012321-0301012230001201)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1020323002022302-2031200012130221-0130123012221301-0303002023310122-1011303313130223-0213220032010232-0220130012302202-3020330121013201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203012103233131-2213121013201221-2300210300212321-3111121313002222-0101302323123230-1033212022330331-3231220301102213-1030023131011311"></a>

## virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset — immediate_action_on_service_down_reset / 120132200011 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-003.md#canonical-2120313011001333-0332132133112211-2032113202213200-1133333320210120-0120003121121103-1110233000210030-2313321130012321-0301012230001201)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset

<a id="canonical-3302211020211201-1110322130010101-0223133020003013-1123322322203021-3312323132031330-3200020132302312-0000211003333112-3033230302112203"></a>

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
immediate_action_on_service_down_reset = {}
```

<a id="canonical-0311102301023103-3022330033312033-1010310212023301-1113202231211302-1203122110022230-2212102101203013-1100002201313313-1001021012110221"></a>

## Direct properties — immediate_action_on_service_down_reset / 120132200011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1303301303011322-1203200110303002-3201211131030133-2001203213031001-1311101312331202-2130011102311323-1323002202212031-0102032110123213"></a>

## Next pages — immediate_action_on_service_down_reset / 120132200011 / 4

- [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-003.md#canonical-2120313011001333-0332132133112211-2032113202213200-1133333320210120-0120003121121103-1110233000210030-2313321130012321-0301012230001201)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1133301311230011-1310300321322210-2110201322213000-0102311121311103-1133120133100303-3022030212232231-3000032201020203-0023300101231100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011023201113313-2123301031202131-1223321221210133-3003210100323300-1111221110011101-2133331303021333-1311311022300313-2031133113320201"></a>

## virtual_server.last_hop_pool — last_hop_pool / 311113321121 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.last_hop_pool

<a id="canonical-3130100312210212-0020201222232323-0102223330023013-2103121210010212-2120033023232203-2312233331101003-1200120122133211-3232032330003110"></a>

Type: `"object"`. list nested block, Optional.

Directs reply traffic to the last hop router using the specified pool.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
last_hop_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-3000033032201312-0003011320211132-3311032003322032-1023122121220132-0210233131233232-2122121003112310-0201121102010112-1002100313122001"></a>

## Direct properties — last_hop_pool / 311113321121 / 3

<a id="canonical-0133220332112203-2120203331023303-0103103201001100-0020310302031313-1322121321222110-2223230321020200-2323310102331130-2110322010101133"></a>

<a id="canonical-1010030312111122-3320032033210110-1012132020032202-3320110132223122-1020220102130123-0303000132201201-3223110221312232-3332202111330113"></a>

## kind property — last_hop_pool / 311113321121 / 4

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

<a id="canonical-0131123000130202-2001333301303310-3310133113312212-2211331301120031-3021020201333212-1023003310010321-3102323310010113-1303302321031230"></a>

<a id="canonical-0030313132323001-3220311020012031-2012322210020211-2002221021200010-3312011032120123-2330031012122113-1123230121212233-3223021202013011"></a>

## name property — last_hop_pool / 311113321121 / 5

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

<a id="canonical-3110220010333333-0132110232200320-2333301323012210-1012002310320022-1210302020303131-0312011232331102-3203201301113230-2213010313021030"></a>

<a id="canonical-0021110133332301-3221311201132113-2301300311320033-2123022222233220-2321031331123331-0302022030100131-0121222323230220-3313201021023131"></a>

## namespace property — last_hop_pool / 311113321121 / 6

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

<a id="canonical-1020121201321111-0321222303220311-3031220120003133-1023233230022012-2021133322232211-3002130130133210-0313013101231031-1310120223223031"></a>

<a id="canonical-0201221221300320-0113203133132322-0223303023021113-1321121201320231-1130113213120132-3300233201001003-0202302220203333-1302011100323320"></a>

## tenant property — last_hop_pool / 311113321121 / 7

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

<a id="canonical-2003030013011001-2001200312132330-3012223122303112-2103222123001133-3212331021031331-1222010300102330-0130102200233130-0100223322332213"></a>

<a id="canonical-0122111022111020-1211012312100301-2000203220113113-2311222331303111-0001312210311322-0311210201033000-0322210113132011-0001100011122131"></a>

## uid property — last_hop_pool / 311113321121 / 8

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

<a id="canonical-3300132120231321-1202201130012201-0023310210323313-1021230333300300-1013312233322332-1222313313202233-3103323002200231-1133023221203110"></a>

## Next pages — last_hop_pool / 311113321121 / 9

- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2002031331123221-1331221331321023-3111232122011101-0320122210120313-0212322110001221-0230222031021312-0203110233131210-3013333012331002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220333102101031-3000200212111212-2211100221012022-1210302133133021-3211123032332112-2103110232302302-0302010303311000-3202033313320233"></a>

## virtual_server.nat64 — nat64 / 021313020333 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.nat64

<a id="canonical-1003101233121130-3210201232023112-3310130202030221-2000220312333131-1103301120112202-2232110110103211-3021002033300032-3230222222320030"></a>

Type: `"object"`. single nested block, Optional.

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if
the..

Upstream description:

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if the
system does not have a default route configured and the client is located on a remote network. This
setting is also useful when the system is load balancing transparent devices that do not modify the
source IP address of the packet. Without the last hop option enabled, the system could return
connections to a different transparent node, resulting in asymmetric routing. You can configure this
setting globally and on an object level. You set the global Auto Last Hop value on the System ::
Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the
Default setting. When you configure Auto Last Hop with a value other than Default at the object
level, its setting takes precedence over the global setting. This enables you to configure auto last
hop on a per-virtual server basis. The default is Default, meaning that the system uses the global
auto-lasthop setting to send back the request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("nat64_disable",
    "nat64_enable")}
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
  "x-ves-oneof-field-nat64_choice": "[\"nat64_disable\",\"nat64_enable\"]"
}
```

Terraform syntax:

```terraform
nat64 {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031212110312232-3232030101022103-0120132230010201-0111011121312303-0102233330320213-2020312033013311-2123112030233002-2110010213013031"></a>

## Direct properties — nat64 / 021313020333 / 3

- [nat64_disable](resources--application_profiles--reference--group-003.md#canonical-1101131332320313-1030303021303121-3201103010330112-0200130321003203-3233323021222323-0021303201111130-0000313032310333-0112133213301313): complete subsection reference.

- [nat64_enable](resources--application_profiles--reference--group-003.md#canonical-0132100231000220-3233110220132320-0202230232001221-0301000023030303-3231320123022201-2323010213123020-0132002101033031-3120303123210130): complete subsection reference.

<a id="canonical-3010300313300021-3331031101120302-2101332032203330-2302131222213033-2101210201300020-1020131023233012-2103033223221231-1313223330230332"></a>

## Next pages — nat64 / 021313020333 / 4

- [virtual_server.nat64.nat64_disable](resources--application_profiles--reference--group-003.md#canonical-1101131332320313-1030303021303121-3201103010330112-0200130321003203-3233323021222323-0021303201111130-0000313032310333-0112133213301313)
- [virtual_server.nat64.nat64_enable](resources--application_profiles--reference--group-003.md#canonical-0132100231000220-3233110220132320-0202230232001221-0301000023030303-3231320123022201-2323010213123020-0132002101033031-3120303123210130)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1101131332320313-1030303021303121-3201103010330112-0200130321003203-3233323021222323-0021303201111130-0000313032310333-0112133213301313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211223022013200-0000102213322002-1033321113121333-1011132300200030-3022323223233210-2021203022320003-2200010023002000-3013111001222001"></a>

## virtual_server.nat64.nat64_disable — nat64_disable / 130113303322 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.nat64](resources--application_profiles--reference--group-003.md#canonical-2002031331123221-1331221331321023-3111232122011101-0320122210120313-0212322110001221-0230222031021312-0203110233131210-3013333012331002)
- virtual_server.nat64.nat64_disable

<a id="canonical-0032200222113123-3223310301322110-3102323310301321-0323320301120132-3213301222232333-1303323203000301-1230002321301113-0200223100132103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for nat64 disable.

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
nat64_disable = {}
```

<a id="canonical-1212012332023313-3322113031000331-2211221323010333-1312323313111202-2033121302310002-0113133232031203-3102320230131103-1020203332213303"></a>

## Direct properties — nat64_disable / 130113303322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2212032223313101-1331323231210011-2223312300311230-1331102021132200-2200122321310130-2211310132303211-0113311023300123-0311222230102012"></a>

## Next pages — nat64_disable / 130113303322 / 4

- [virtual_server.nat64](resources--application_profiles--reference--group-003.md#canonical-2002031331123221-1331221331321023-3111232122011101-0320122210120313-0212322110001221-0230222031021312-0203110233131210-3013333012331002)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0132100231000220-3233110220132320-0202230232001221-0301000023030303-3231320123022201-2323010213123020-0132002101033031-3120303123210130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311233013023200-0002032122130011-3030321323023003-2021311211203330-1323311002322002-1100000023223111-2330011221130122-1020230003032013"></a>

## virtual_server.nat64.nat64_enable — nat64_enable / 331013021332 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.nat64](resources--application_profiles--reference--group-003.md#canonical-2002031331123221-1331221331321023-3111232122011101-0320122210120313-0212322110001221-0230222031021312-0203110233131210-3013333012331002)
- virtual_server.nat64.nat64_enable

<a id="canonical-3213312001302303-0123300102331001-2011233202113212-3023132011020311-0222123030123111-1023322311011003-1003301320323332-1231200122323330"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for nat64 enable.

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
nat64_enable = {}
```

<a id="canonical-2130100231232021-0103101301213200-2002322123233003-2332022003310033-1032003011012031-1301012001003113-2310233002121020-2231110011303233"></a>

## Direct properties — nat64_enable / 331013021332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3230020102020333-1212102020111313-0201222212333320-1110332102111020-2332321111011120-0233321232031111-3122113331223332-0133202233201130"></a>

## Next pages — nat64_enable / 331013021332 / 4

- [virtual_server.nat64](resources--application_profiles--reference--group-003.md#canonical-2002031331123221-1331221331321023-3111232122011101-0320122210120313-0212322110001221-0230222031021312-0203110233131210-3013333012331002)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-3121121231001031-3330300000133002-1031230222111031-0110330122312033-2031202031313323-3102103200300201-2230013301021233-3121002202231030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110110031022023-2213322211011013-2112230023222011-0230111112100113-0111322312202330-2121313223321011-3031301211232210-3202210202113330"></a>

## virtual_server.port_translation — port_translation / 110302131220 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.port_translation

<a id="canonical-3300211131303022-0213330231302003-0323211023223112-0221132220220312-1121333121213210-2133331232122332-0010300301232312-0122121000212121"></a>

Type: `"object"`. single nested block, Optional.

Specifies, when checked (enabled), that the system translates the port of the virtual server. When
cleared (disabled), specifies that the system uses the port without translation. Turning off port
translation for a virtual server is useful if you want to use the virtual server to load balance..

Upstream description:

Specifies, when checked (enabled), that the system translates the port of the virtual server. When
cleared (disabled), specifies that the system uses the port without translation. Turning off port
translation for a virtual server is useful if you want to use the virtual server to load balance
connections to any service. The default is enabled.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("port_translation_disable",
    "port_translation_enable")}
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
  "x-ves-oneof-field-port_translation_choice": "[\"port_translation_disable\",\"port_translation_enable\"]"
}
```

Terraform syntax:

```terraform
port_translation {
  # Configure direct properties listed below.
}
```

<a id="canonical-3211131130331022-1333230001033213-1131103030212030-0300010300012120-0012233131201000-2103003202232301-2302332012130031-0000210002003233"></a>

## Direct properties — port_translation / 110302131220 / 3

- [port_translation_disable](resources--application_profiles--reference--group-003.md#canonical-2322122032001202-0320333323113103-1320111022100231-1122220021233203-0132331132331011-0222313003233312-0311310233221122-0300233002300010): complete subsection reference.

- [port_translation_enable](resources--application_profiles--reference--group-003.md#canonical-2333311333021133-0102003021323220-2320330011130202-0302232011300011-0133021103300303-1030131302202301-2121022231132022-0333012033121023): complete subsection reference.

<a id="canonical-0202312230133232-0010112013222300-1333023213122201-0133331121002011-0331210220233332-1221332103130112-2003022102332120-2322032230012311"></a>

## Next pages — port_translation / 110302131220 / 4

- [virtual_server.port_translation.port_translation_disable](resources--application_profiles--reference--group-003.md#canonical-2322122032001202-0320333323113103-1320111022100231-1122220021233203-0132331132331011-0222313003233312-0311310233221122-0300233002300010)
- [virtual_server.port_translation.port_translation_enable](resources--application_profiles--reference--group-003.md#canonical-2333311333021133-0102003021323220-2320330011130202-0302232011300011-0133021103300303-1030131302202301-2121022231132022-0333012033121023)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2322122032001202-0320333323113103-1320111022100231-1122220021233203-0132331132331011-0222313003233312-0311310233221122-0300233002300010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213031021002212-2030310300033231-0322210200213221-1310221023203100-2322222312103300-0201230302320321-0310131331033221-2112023322201200"></a>

## virtual_server.port_translation.port_translation_disable — port_translation_disable / 212213312323 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.port_translation](resources--application_profiles--reference--group-003.md#canonical-3121121231001031-3330300000133002-1031230222111031-0110330122312033-2031202031313323-3102103200300201-2230013301021233-3121002202231030)
- virtual_server.port_translation.port_translation_disable

<a id="canonical-0331201330112220-3322000201302223-3302323211013310-1011022233300321-0102203301310010-0301110102203200-0222131221201313-0130021112300230"></a>

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
port_translation_disable = {}
```

<a id="canonical-0032122120013100-1211202232331311-2122301101311313-2333102102212213-0322233120320221-1121231220010321-1322110112211302-0113213311012321"></a>

## Direct properties — port_translation_disable / 212213312323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213311111003302-0322000113231110-0221101103223130-0331203210102222-0023223313020220-0022020013121120-2322312311221203-3223023330300321"></a>

## Next pages — port_translation_disable / 212213312323 / 4

- [virtual_server.port_translation](resources--application_profiles--reference--group-003.md#canonical-3121121231001031-3330300000133002-1031230222111031-0110330122312033-2031202031313323-3102103200300201-2230013301021233-3121002202231030)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2333311333021133-0102003021323220-2320330011130202-0302232011300011-0133021103300303-1030131302202301-2121022231132022-0333012033121023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032123033311012-2000231333002003-1211033310211023-2103332123120301-0310010103012223-3200031012002221-0222021301103032-0322003123101312"></a>

## virtual_server.port_translation.port_translation_enable — port_translation_enable / 212101322022 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.port_translation](resources--application_profiles--reference--group-003.md#canonical-3121121231001031-3330300000133002-1031230222111031-0110330122312033-2031202031313323-3102103200300201-2230013301021233-3121002202231030)
- virtual_server.port_translation.port_translation_enable

<a id="canonical-2110210322123011-1100230132320211-0101312110330201-0010313202120310-2302300301132231-3303200010330110-1020120100133223-0330011233231021"></a>

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
port_translation_enable = {}
```

<a id="canonical-2310203213203103-3220331031023310-3320230032102113-1012123110220320-1201210001022233-2212301012021101-3300121031232300-3033201332321222"></a>

## Direct properties — port_translation_enable / 212101322022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002031003000200-2032030211331310-0030302012330310-1313300232232222-0321331031131202-0321002332231011-1021323022200330-2021031320233000"></a>

## Next pages — port_translation_enable / 212101322022 / 4

- [virtual_server.port_translation](resources--application_profiles--reference--group-003.md#canonical-3121121231001031-3330300000133002-1031230222111031-0110330122312033-2031202031313323-3102103200300201-2230013301021233-3121002202231030)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-3210202300313301-0310203232211033-1321103003021001-0212002312311332-1112331303033331-3323000030313321-0300000231230113-3220231130103020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220123331312013-1233210103033002-0200210010131212-1030102122312233-3313000030023322-1303203310021301-3033033010322301-0232111203321221"></a>

## virtual_server.request_logging_profile — request_logging_profile / 333320212101 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.request_logging_profile

<a id="canonical-3021100321112301-1233301012311200-0312330023111200-2001301020303311-2233302130132031-2113023101013111-0133021233213323-3232312330213023"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for request logging profile.

Upstream description:

Configuration parameter for request logging profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
request_logging_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322333210020223-2120230110312331-1112110332100120-1331221210011010-1233012033321120-1033320220212112-2203021200000120-0123103231200231"></a>

## Direct properties — request_logging_profile / 333320212101 / 3

<a id="canonical-2233232030023313-3133202323110131-3131322232122022-1020123321200333-2321211132333330-0330210122202000-2232233121020312-1121133312313321"></a>

<a id="canonical-3122110221220000-1200001012321230-3012131020313301-0112300320223203-2132303222112321-3312130313222112-0212022200231033-3112010312003011"></a>

## kind property — request_logging_profile / 333320212101 / 4

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

<a id="canonical-2330113122030033-3333310220311332-0223021101213303-1323103321221121-3201211312121333-0033032101131310-2332022230133023-3322223132003220"></a>

<a id="canonical-3310030321013232-0321313131032033-3321220231011032-3023100321202022-3233333330213031-3312213020131021-0220003131103101-1220030031321012"></a>

## name property — request_logging_profile / 333320212101 / 5

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

<a id="canonical-0233120103100302-1311201110110101-1120331232001210-0000121221300121-2120013031003101-0213331233212322-0010322230122333-3123123332012210"></a>

<a id="canonical-1230033223300210-3233310213203303-0003003213032312-2232300333300232-1230122101210311-3001201021030333-1213132211320002-2022231222231222"></a>

## namespace property — request_logging_profile / 333320212101 / 6

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

<a id="canonical-0203102211223131-3103102013132000-3130123210303311-2100310013322023-2003321300202302-0101001321232302-0232203321100000-2310013110330001"></a>

<a id="canonical-0221131221320331-0010113223303032-2303331132213001-1312102003113232-0122122213222322-3202301112023110-0202330100332213-1012103023123321"></a>

## tenant property — request_logging_profile / 333320212101 / 7

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

<a id="canonical-1031020323010220-1310111021303131-1211130301120110-0120303302330312-0301202230332110-0122211111020020-2323102113133330-0220213033212231"></a>

<a id="canonical-1113213323331013-3233020200113231-0101210033322111-2301003311203103-2331310022131010-0332232112212221-3231231210231232-2102031212233010"></a>

## uid property — request_logging_profile / 333320212101 / 8

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

<a id="canonical-3203020313021200-3331203031231110-0121001103003311-3131013033223233-2330333312112301-1332210231221022-3201131012102311-1100223301311120"></a>

## Next pages — request_logging_profile / 333320212101 / 9

- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1232230330103113-1011323112020110-2112322333333013-3302112210103221-2002022032123220-0023200230233130-2102030102320103-0013123112003122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233131132332312-0200321001011311-3000310222012331-2232102121012220-0223031020202201-3231000001232202-0230103332330300-2130212300001000"></a>

## virtual_server.source_port — source_port / 000233212233 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.source_port

<a id="canonical-2212010332223131-2111323213003221-3122213303023310-2020333131322323-0103201100101002-1020010232021313-0111203202021022-2223132001320213"></a>

Type: `"object"`. single nested block, Optional.

Specifies whether the system preserves the source port of the connection. The default is Preserve.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("source_port_change",
    "source_port_preserve"),
  validators.ConflictingObjectAttributes("source_port_change",
    "source_port_preserve_strict"),
  validators.ConflictingObjectAttributes("source_port_preserve",
    "source_port_preserve_strict")}
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
  "x-ves-oneof-field-source_port_choice": "[\"source_port_change\",\"source_port_preserve\",\"source_port_preserve_strict\"]"
}
```

Terraform syntax:

```terraform
source_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213203131013212-3001033120130202-2303013310321122-1130300332130311-3201111230303203-3123323312021112-0110231320010311-0021023023200200"></a>

## Direct properties — source_port / 000233212233 / 3

- [source_port_change](resources--application_profiles--reference--group-003.md#canonical-0001232132032021-0203012123211133-0320111311333103-3333303322101113-3130310123301120-3310300232031132-0113103322101012-3021301033310022): complete subsection reference.

- [source_port_preserve](resources--application_profiles--reference--group-003.md#canonical-0103110100200031-2120012230132120-3221310123003222-0102011031233010-3323321000111202-3210032023300110-1210012201200122-2012133333321303): complete subsection reference.

- [source_port_preserve_strict](resources--application_profiles--reference--group-003.md#canonical-2110002333222312-2100332330303222-0000003122110232-0110132011311230-1123023003303233-2120121032203202-3211123212022100-3030233302310031): complete subsection reference.

<a id="canonical-0232110210023022-2330023221320002-0301123302030122-3202203223331002-1003101120203333-1231131020211232-0232101222130312-1032131321321001"></a>

## Next pages — source_port / 000233212233 / 4

- [virtual_server.source_port.source_port_change](resources--application_profiles--reference--group-003.md#canonical-0001232132032021-0203012123211133-0320111311333103-3333303322101113-3130310123301120-3310300232031132-0113103322101012-3021301033310022)
- [virtual_server.source_port.source_port_preserve](resources--application_profiles--reference--group-003.md#canonical-0103110100200031-2120012230132120-3221310123003222-0102011031233010-3323321000111202-3210032023300110-1210012201200122-2012133333321303)
- [virtual_server.source_port.source_port_preserve_strict](resources--application_profiles--reference--group-003.md#canonical-2110002333222312-2100332330303222-0000003122110232-0110132011311230-1123023003303233-2120121032203202-3211123212022100-3030233302310031)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0001232132032021-0203012123211133-0320111311333103-3333303322101113-3130310123301120-3310300232031132-0113103322101012-3021301033310022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301312301322122-1330023301102310-3122220331201111-2320130013033231-0102020313033033-0123000033211223-1210313011023221-3210223121323230"></a>

## virtual_server.source_port.source_port_change — source_port_change / 102030022112 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.source_port](resources--application_profiles--reference--group-003.md#canonical-1232230330103113-1011323112020110-2112322333333013-3302112210103221-2002022032123220-0023200230233130-2102030102320103-0013123112003122)
- virtual_server.source_port.source_port_change

<a id="canonical-1311330323021032-2210013333333302-1331301003130203-0003333220201000-3131331012012000-2013032302131030-2303202113122003-3323113201112320"></a>

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
source_port_change = {}
```

<a id="canonical-3002103311013030-2210030312231133-2100220001003021-1322011322201213-2120212112233200-1122013121033323-2010301103312011-1003311030230311"></a>

## Direct properties — source_port_change / 102030022112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201310311021202-3103210101010132-1003121000212221-1301333313031003-1212322112001022-0230012023301123-2022120201303130-3133331310200011"></a>

## Next pages — source_port_change / 102030022112 / 4

- [virtual_server.source_port](resources--application_profiles--reference--group-003.md#canonical-1232230330103113-1011323112020110-2112322333333013-3302112210103221-2002022032123220-0023200230233130-2102030102320103-0013123112003122)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0103110100200031-2120012230132120-3221310123003222-0102011031233010-3323321000111202-3210032023300110-1210012201200122-2012133333321303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331100021132331-2201012031030121-2203211301110221-2112231302133112-3120332112320123-3103313331312031-1022010202232003-1203103312223000"></a>

## virtual_server.source_port.source_port_preserve — source_port_preserve / 130010023103 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.source_port](resources--application_profiles--reference--group-003.md#canonical-1232230330103113-1011323112020110-2112322333333013-3302112210103221-2002022032123220-0023200230233130-2102030102320103-0013123112003122)
- virtual_server.source_port.source_port_preserve

<a id="canonical-1022313012110112-0320200311020023-1202203230011030-1020100221232300-1322001023000120-3123313223012212-0303223030221123-1102121132013322"></a>

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
source_port_preserve = {}
```

<a id="canonical-1020101023132012-3313100302111122-0133301121122230-2012232230231323-2031312221312202-2032300101202132-2101203220311220-1222301233333103"></a>

## Direct properties — source_port_preserve / 130010023103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231021022202033-2112322331200231-2222032230010003-1313200031211223-2121013121220200-2333020231312201-3113222103233311-2221211331331000"></a>

## Next pages — source_port_preserve / 130010023103 / 4

- [virtual_server.source_port](resources--application_profiles--reference--group-003.md#canonical-1232230330103113-1011323112020110-2112322333333013-3302112210103221-2002022032123220-0023200230233130-2102030102320103-0013123112003122)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2110002333222312-2100332330303222-0000003122110232-0110132011311230-1123023003303233-2120121032203202-3211123212022100-3030233302310031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300312221332132-0301210003212123-0221011002231333-2323023023002120-1310012121003312-0032220110331331-1121003301113311-3020111210100021"></a>

## virtual_server.source_port.source_port_preserve_strict — source_port_preserve_strict / 301310231020 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.source_port](resources--application_profiles--reference--group-003.md#canonical-1232230330103113-1011323112020110-2112322333333013-3302112210103221-2002022032123220-0023200230233130-2102030102320103-0013123112003122)
- virtual_server.source_port.source_port_preserve_strict

<a id="canonical-2212111003100111-1020312302333103-1300101311131310-0311112203132102-2032131012231213-1123102122310031-2310031230010301-0130002031012300"></a>

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
source_port_preserve_strict = {}
```

<a id="canonical-3320022002031030-3131210122011102-0020132102031332-0312031103020022-0312202113023101-2101020100203031-3222130002011123-0122031202003300"></a>

## Direct properties — source_port_preserve_strict / 301310231020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113112030033030-3322311111311132-2002302232033103-2321223323001001-2203131303302232-3233211211001112-3032333210230030-1132020300223200"></a>

## Next pages — source_port_preserve_strict / 301310231020 / 4

- [virtual_server.source_port](resources--application_profiles--reference--group-003.md#canonical-1232230330103113-1011323112020110-2112322333333013-3302112210103221-2002022032123220-0023200230233130-2102030102320103-0013123112003122)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2230332010120210-3032201031320120-1320321332001001-1230210320213223-3313130100303010-3030221032130120-0211333311033312-3013103322102113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230121030222002-0033031003112020-1222012022220223-2023010123202333-1120310203220321-0213303332301122-0311000112200202-3132212212222123"></a>

## virtual_server.statistics_profile — statistics_profile / 011012333102 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.statistics_profile

<a id="canonical-2033031113031132-0023012103300133-1113311132332021-2001201030102232-3111110013031123-1322303301011332-3021010132203122-3121230300333000"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for statistics profile.

Upstream description:

Configuration parameter for statistics profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
statistics_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3123312201013233-3012033020121103-2031020102000321-0311023031122023-3111000332220011-0123121020213320-3213102330021031-2310332202131220"></a>

## Direct properties — statistics_profile / 011012333102 / 3

<a id="canonical-1003022020103133-2131111133003331-0330033120200303-2201030200030322-2303321002002211-0202203103201322-2310331120023300-2022202210203031"></a>

<a id="canonical-0102030000302012-3101333231110332-2221303100201220-0312032002223123-2333133220131123-1232122132311030-2210130112332233-2111130010301213"></a>

## kind property — statistics_profile / 011012333102 / 4

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

<a id="canonical-1300331213131100-1210231223321210-0020103121333312-3101130222010232-3132202211311101-3010011311133122-0031101310031213-1302031013313323"></a>

<a id="canonical-2323221120022300-3331121302022011-2000112310323332-2322210131333121-2221022211300321-1101103012331233-0123223300013120-1102302233001111"></a>

## name property — statistics_profile / 011012333102 / 5

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

<a id="canonical-3223323302013221-2212221320311113-0100312022002101-0330320332101201-0100233201322021-1030312112310232-0011022110313021-3300023211200223"></a>

<a id="canonical-1231020131210033-3321233133300323-2000223022322031-2330000021130100-1130232002033231-2223223313330112-1001013032103320-2023032010111031"></a>

## namespace property — statistics_profile / 011012333102 / 6

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

<a id="canonical-2221133111123231-3110101301023112-3231003013000012-2122230101110303-3212013230220030-2232232033012220-1223233112020031-0002030333311121"></a>

<a id="canonical-3131032012000001-2331110033313310-2033102313021202-1000220331232321-3012220010121000-3012312130121021-3220223010231321-3303320210010231"></a>

## tenant property — statistics_profile / 011012333102 / 7

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

<a id="canonical-0132320122003211-0231231122332331-1003010313310213-2123301122301231-0322332221310022-0203123211123220-1222031302133221-1200331122023110"></a>

<a id="canonical-2013132101130230-1303200311300322-2000321032133312-0030101331301133-2022101103300300-0331132130020023-0203011310133003-3311310221220100"></a>

## uid property — statistics_profile / 011012333102 / 8

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

<a id="canonical-2132211131231023-0020003330030323-0113101002320002-2012132220230032-0312011202010002-3112103122303203-0023131221302001-3232031302032202"></a>

## Next pages — statistics_profile / 011012333102 / 9

- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110020231202222-0303102112022011-0100132321010011-0111033303203232-2122322111020020-2321100121131110-1223303023221233-3230132122201210"></a>

## virtual_server.tcp — tcp / 102220020003 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.tcp

<a id="canonical-1310210103201212-1320301132201233-1210200013231223-1101230001132123-3302033232011132-2332321312221012-2323132221331212-0130323232233301"></a>

Type: `"object"`. single nested block, Optional.

TCP profiles.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
tcp {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202311012120201-2023022031030300-0313233120303320-3313301223010100-1131100131210221-3313223130123020-3023030123221103-1203333321323120"></a>

## Direct properties — tcp / 102220020003 / 3

- [client_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-2132222223303130-1022130203031333-3310131122103110-0212220202301313-3312320232032323-3331223010113331-1322233131031331-3313323020033222): complete subsection reference.

- [ocsp_profile](resources--application_profiles--reference--group-003.md#canonical-1001023112303131-2331323203333210-1232132101202000-2233311002102231-3110312133322122-3322100033331020-3302210231201232-2313103221001212): complete subsection reference.

- [server_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-0121320311113012-3102203220321101-2131030131221121-1312131030132203-2021130233323110-1023233231211102-3231031033113313-3221001212010323): complete subsection reference.

- [tcp_client_profile](resources--application_profiles--reference--group-003.md#canonical-1112301121000320-0032302331230222-3123213322221103-2320310231001320-2232202220023120-0332311010112102-3303010032311201-2001323230021001): complete subsection reference.

- [tcp_server_profile](resources--application_profiles--reference--group-003.md#canonical-0032321113220230-2210221333303233-0123322330001031-0211010312302301-0111111221033222-2310210331211213-3000322022123112-2231210311122112): complete subsection reference.

<a id="canonical-0133130332101130-3313101201210033-0020012000103302-2223230033023232-2222200123122333-3213003302201120-1320031220121212-0103110121223200"></a>

## Next pages — tcp / 102220020003 / 4

- [virtual_server.tcp.client_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-2132222223303130-1022130203031333-3310131122103110-0212220202301313-3312320232032323-3331223010113331-1322233131031331-3313323020033222)
- [virtual_server.tcp.ocsp_profile](resources--application_profiles--reference--group-003.md#canonical-1001023112303131-2331323203333210-1232132101202000-2233311002102231-3110312133322122-3322100033331020-3302210231201232-2313103221001212)
- [virtual_server.tcp.server_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-0121320311113012-3102203220321101-2131030131221121-1312131030132203-2021130233323110-1023233231211102-3231031033113313-3221001212010323)
- [virtual_server.tcp.tcp_client_profile](resources--application_profiles--reference--group-003.md#canonical-1112301121000320-0032302331230222-3123213322221103-2320310231001320-2232202220023120-0332311010112102-3303010032311201-2001323230021001)
- [virtual_server.tcp.tcp_server_profile](resources--application_profiles--reference--group-003.md#canonical-0032321113220230-2210221333303233-0123322330001031-0211010312302301-0111111221033222-2310210331211213-3000322022123112-2231210311122112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2132222223303130-1022130203031333-3310131122103110-0212220202301313-3312320232032323-3331223010113331-1322233131031331-3313323020033222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313120332123020-3002311011333001-3300220310220202-0000232231022213-3232232112133322-3323322212023110-2021213021131012-1100021210230213"></a>

## virtual_server.tcp.client_ssl_profile — client_ssl_profile / 332232013303 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231)
- virtual_server.tcp.client_ssl_profile

<a id="canonical-1012331213102311-2311220320111322-0123323233221201-0303203130002202-2000012130110011-0131133020313202-2111122321103312-1332310121103313"></a>

Type: `"object"`. list nested block, Optional.

Client SSL Profile. Client-side configuration

Upstream description:

Client-side configuration

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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
client_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011322202233302-0303313123002231-1020331302133133-3300333220302320-0110000233011130-1121211202333313-3113032021131103-3200011221001302"></a>

## Direct properties — client_ssl_profile / 332232013303 / 3

<a id="canonical-1233012303331002-1301300223130020-1210102110333010-1201103032033230-1030230201231331-2330230321233222-2123313200310302-3010200130331201"></a>

<a id="canonical-2000030311312200-1232120320110322-2030132121031121-2113301313120031-2113111302333210-1221131223103211-0200303123321002-3121220121132210"></a>

## kind property — client_ssl_profile / 332232013303 / 4

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

<a id="canonical-3312021332221321-2210030022101301-1223102233113103-1310220121202012-0202103103302322-0320203233310212-2322212223322033-1233030323030321"></a>

<a id="canonical-1100003222211122-2223020232101113-2311302013013231-1333320210200301-1030231203213330-3211030002032012-2113201312033310-2111033212303023"></a>

## name property — client_ssl_profile / 332232013303 / 5

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

<a id="canonical-2231021120022302-3031120322012220-3123202230023202-1201233022301012-3112121200011123-3331223131023202-1230310323230232-3220010310211111"></a>

<a id="canonical-0001111321330321-2102331321121231-0000313211302203-2003131022032031-1302130311202121-3331131213002011-1002003303101021-0022010010122321"></a>

## namespace property — client_ssl_profile / 332232013303 / 6

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

<a id="canonical-3120212012111032-1013103233222010-1203000331212110-1133102031232311-3121223232313333-0233221303232133-0323231331110220-1201223021000213"></a>

<a id="canonical-3033321210233112-3112123313132300-1323020302201120-2311333331211232-1120110313023302-1103021001201112-1300011132213300-0100301021300303"></a>

## tenant property — client_ssl_profile / 332232013303 / 7

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

<a id="canonical-1121233010211331-3031222020231121-2312000122013110-1220221302100230-3200321013220312-3331200330132012-3001310002311313-3322122323132033"></a>

<a id="canonical-1121020310332033-0232000301012221-2002011330313100-2100021211231101-0013233123203100-2203010001121322-1030330200313033-0111323131321001"></a>

## uid property — client_ssl_profile / 332232013303 / 8

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

<a id="canonical-3210311232032211-0010203011012033-0003101033231333-1333320203012003-3020112230200001-0032202210203103-2302032032001300-1120021023133322"></a>

## Next pages — client_ssl_profile / 332232013303 / 9

- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1001023112303131-2331323203333210-1232132101202000-2233311002102231-3110312133322122-3322100033331020-3302210231201232-2313103221001212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311222201333030-2130331003231121-3130130020212123-1032311133233232-2130310230331321-2103021112000300-0101003220112111-2301013311232231"></a>

## virtual_server.tcp.ocsp_profile — ocsp_profile / 302230313223 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231)
- virtual_server.tcp.ocsp_profile

<a id="canonical-2201330102321311-3113102120311103-0132000300300123-2232331231112021-0030311232033330-1030333220331133-2233212221203000-2220021013322130"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for ocsp profile.

Upstream description:

Configuration parameter for ocsp profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
ocsp_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112123102212310-3233222111133230-2312203232320202-3223101232021122-0211021322212303-2231303012311300-1321300000303300-3130012200332212"></a>

## Direct properties — ocsp_profile / 302230313223 / 3

<a id="canonical-1201102121120222-1022222303212032-0322203300233201-2321122032331211-1312031123202300-3123131021001101-0031312312010221-3032313300130002"></a>

<a id="canonical-0222211230322302-1232210001301213-2031333302100013-3131312123011211-0021021131313200-1231000201211322-3323031011200110-0133202003122200"></a>

## kind property — ocsp_profile / 302230313223 / 4

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

<a id="canonical-3220000202123301-2133312020312122-2213001131121321-2300122321203311-1333130233201213-1203023101110133-0311130020302200-2133220203102110"></a>

<a id="canonical-2332132330302112-2331221312002221-3200331001130101-2231323310313332-3302321020012323-3030310120023020-1303003233110132-2013131321122221"></a>

## name property — ocsp_profile / 302230313223 / 5

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

<a id="canonical-1030330131020231-3012302212102302-0131222233220221-2313221220330210-1113333230120223-3223121032303210-3022220012221013-3021233212311311"></a>

<a id="canonical-3210322032012302-1322233133021020-3233300002220001-1131330120331331-0113213231223131-2203200220030320-1013311332103110-3113312323100203"></a>

## namespace property — ocsp_profile / 302230313223 / 6

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

<a id="canonical-0032230003100312-1121023211302231-3120031120223200-2212022112333313-2120130201310201-1132102223112330-0303301332002311-0001202023003323"></a>

<a id="canonical-2120211233000322-2202323230202120-3301000020311322-2123033111213330-0131031100221331-1133020232113033-2220230310033111-0320220030201020"></a>

## tenant property — ocsp_profile / 302230313223 / 7

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

<a id="canonical-2221023120123210-3013033120202100-1222203231232011-0210232310230203-3321200030133200-2100022322310200-3100301220121010-1120022013123311"></a>

<a id="canonical-2012321110330131-3301333012201322-2332223012202022-1100322103123011-0223211321100223-3203130012112023-2002033331130211-0102233202132222"></a>

## uid property — ocsp_profile / 302230313223 / 8

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

<a id="canonical-2001302030313220-2131233333111323-3012311103220021-0213111331012321-2012110020231113-0011112311212211-0031210321022310-1203122303111010"></a>

## Next pages — ocsp_profile / 302230313223 / 9

- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0121320311113012-3102203220321101-2131030131221121-1312131030132203-2021130233323110-1023233231211102-3231031033113313-3221001212010323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033111120331301-0321320120303021-3332213021233103-2012213310331202-0023221233130132-0221313122110131-2201113230232132-1310223101130333"></a>

## virtual_server.tcp.server_ssl_profile — server_ssl_profile / 033102122110 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231)
- virtual_server.tcp.server_ssl_profile

<a id="canonical-0313332111330002-2300020230310231-2231001022303321-0233332113220203-3231013103110021-2131313133233102-1133110013032032-3130221031223302"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for server SSL profile.

Upstream description:

Configuration parameter for server SSL profile

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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
server_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011230212110333-2033123220002310-2031213032013313-3013212310023313-1000333101230002-3311011331033022-1300323220210311-3222012221111113"></a>

## Direct properties — server_ssl_profile / 033102122110 / 3

<a id="canonical-3000301213030020-0232223320210223-1100203023111111-2200123330230123-0220113113122020-0033322302122200-1020203113131031-2033323320221123"></a>

<a id="canonical-0233133311220002-3203023001333333-0333320131031200-2323010222001302-2103022220212232-0032033130221220-2331311013130120-3223130322333020"></a>

## kind property — server_ssl_profile / 033102122110 / 4

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

<a id="canonical-0331011311300103-2303020313011023-0323311131001332-2232012123332212-1213101230013110-2203232012211103-0312112021022210-2131202323331112"></a>

<a id="canonical-0220113323031231-3232223313213130-0310130032022102-0012012213210022-3030212111232231-3302003132121131-3212031300120203-1030130101133203"></a>

## name property — server_ssl_profile / 033102122110 / 5

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

<a id="canonical-2211313211222321-1200130133311001-1321223203123312-2322013012003030-2101320031032233-3021131130011110-3320231203110033-2110132130210033"></a>

<a id="canonical-3112213032132323-0331300201233230-0203200110001103-0030120020112302-2122020300322200-2203101210312021-2310310330333310-1332131313332310"></a>

## namespace property — server_ssl_profile / 033102122110 / 6

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

<a id="canonical-0320303333101212-1300001123203212-1120030002323031-3032033320113002-2232031333132003-1033022112033322-1300013131312023-1210002101332203"></a>

<a id="canonical-1032122201131211-3022231211230020-3033012222211012-1102201223213121-2300202012332110-3232221102003022-1022332313020111-2033130221300321"></a>

## tenant property — server_ssl_profile / 033102122110 / 7

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

<a id="canonical-2301220023033330-2023131320213003-2131230213320311-1300222113211131-1112331122121130-2211030220030333-2112302322201101-3201301223333221"></a>

<a id="canonical-3320133313200012-1033130333202312-2220111302313310-0021210003310022-3331010231220011-1323200002322212-0000202200203221-3033130100232101"></a>

## uid property — server_ssl_profile / 033102122110 / 8

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

<a id="canonical-3312003303003110-3122332123112233-3022130002320203-3221101112212032-2112121133222321-1321020131310202-0111303120013221-3230013101303212"></a>

## Next pages — server_ssl_profile / 033102122110 / 9

- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1112301121000320-0032302331230222-3123213322221103-2320310231001320-2232202220023120-0332311010112102-3303010032311201-2001323230021001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023321102201221-3130121202210200-3313213033131103-2302312001113201-1221013130320232-1102332303000030-0233210213101123-0132110212011113"></a>

## virtual_server.tcp.tcp_client_profile — tcp_client_profile / 200210200012 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231)
- virtual_server.tcp.tcp_client_profile

<a id="canonical-2323011223033022-3013123232310303-2200023233330130-2110012032322320-2211000322003323-1132330113030022-1210110101322320-2011132102222300"></a>

Type: `"object"`. list nested block, Optional.

Protocol Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tcp_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0202233311032000-1132302313231013-2011330323132322-3321320303303230-3131023023313222-3303311013233031-0022330030032331-0001002003222322"></a>

## Direct properties — tcp_client_profile / 200210200012 / 3

<a id="canonical-1202333133013223-0321300232132300-3322332332223322-3310100210232000-2313000123313022-0133031012031030-0111033001113023-1333023133112133"></a>

<a id="canonical-3130301233233220-1102322131132211-3020303033000032-0302130222010221-3332303021033322-0100331311123132-0032000131331222-1202031230123231"></a>

## kind property — tcp_client_profile / 200210200012 / 4

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

<a id="canonical-2313321202202003-2211103320031111-1320213310113001-1231201312123112-0133033013113213-3001313330131133-2200001003013203-2233312331323120"></a>

<a id="canonical-3102213111231102-0003310012331131-1301003333010230-1111223020002220-3001313110010302-3123300330312333-2030112021132323-1213010031311330"></a>

## name property — tcp_client_profile / 200210200012 / 5

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

<a id="canonical-2111120120313200-3222101022030010-1301020013112011-2231120323103120-2010320021303330-3022113313332101-0230303203133131-2231323202010312"></a>

<a id="canonical-2003120001132130-0203113001333321-1031233231012103-1000211032132300-1021220311223110-2311130310002103-0303203222100123-1002310000222013"></a>

## namespace property — tcp_client_profile / 200210200012 / 6

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

<a id="canonical-1031111320132203-0200012003211020-3021103333133103-2210212320101013-1000010323023303-2000010111101101-0230111331202123-3231120012033013"></a>

<a id="canonical-0122032011002323-0322332013012030-0302232320130320-1211302333011210-0033231012133100-3311220220300123-2321020003222230-2101221101202021"></a>

## tenant property — tcp_client_profile / 200210200012 / 7

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

<a id="canonical-2031013330200223-0133032220302133-3223210223232202-1110331220121021-0213221001321033-0112020211120203-3232331310200220-1120233212320331"></a>

<a id="canonical-2133131330320000-1230011110101011-0210110002012123-2013000211033312-2020011311003323-3021303130101212-3001200311211120-2222020211331002"></a>

## uid property — tcp_client_profile / 200210200012 / 8

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

<a id="canonical-1320010030203310-1023333220012223-1310021112033123-0113331223101112-0121111111101010-1133123013113223-1120223220213011-2200002202131111"></a>

## Next pages — tcp_client_profile / 200210200012 / 9

- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0032321113220230-2210221333303233-0123322330001031-0211010312302301-0111111221033222-2310210331211213-3000322022123112-2231210311122112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130111120013310-0313130330202112-3211310020200032-0320100133002301-1332022030012320-2111010302220212-1031013313132330-0113032110230200"></a>

## virtual_server.tcp.tcp_server_profile — tcp_server_profile / 023200112201 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231)
- virtual_server.tcp.tcp_server_profile

<a id="canonical-3122102111112221-0210010201323231-2021200021100112-1020010323310032-1011031101220320-1111000012210123-0010330130312223-1231232220122133"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for tcp server profile.

Upstream description:

Configuration parameter for tcp server profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tcp_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233221322132033-2002300132323023-0310021312013323-2023120020021002-1333010222332202-3122103230102303-1311212300320131-0031132332013012"></a>

## Direct properties — tcp_server_profile / 023200112201 / 3

<a id="canonical-1303000301321302-3323023113031001-2131111011202020-1113011322321101-3101312301111011-2100030020031300-2221222311010002-0030123213213302"></a>

<a id="canonical-2123302123301123-0123212220132132-2222110032113223-2110321132301323-3322221331333232-3210322103013121-2321000123313101-1000101130333321"></a>

## kind property — tcp_server_profile / 023200112201 / 4

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

<a id="canonical-1010002221322231-2200211120100222-3223323222032100-0030000033123203-1201012112031311-1201213011010333-0222033300320320-2032120100321330"></a>

<a id="canonical-1221302203311332-2202121112201323-1103231021123231-2203212220331002-3011312213113013-3101032010121231-2032322112311320-3021002332112232"></a>

## name property — tcp_server_profile / 023200112201 / 5

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

<a id="canonical-2120132210000323-2102211130312203-1223202021100201-2131133002303321-3203123113221012-2213313132001133-1102132201232002-1021032020110230"></a>

<a id="canonical-1102222100000201-1101011030310202-2121210311123033-2022310011010223-3123113102310023-0312210032010020-0213313332103022-0333223103131101"></a>

## namespace property — tcp_server_profile / 023200112201 / 6

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

<a id="canonical-2030130300331200-1133103021231303-0312010033233120-1332321301230300-2002110230112212-2010022302000312-0122331020010202-3031032010011133"></a>

<a id="canonical-0103213120232102-3313112221103111-2201130231011103-3323233132033230-1022213123132111-0033221102202000-1001122332222313-2230100101122133"></a>

## tenant property — tcp_server_profile / 023200112201 / 7

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

<a id="canonical-3222222311333301-2121222012111010-1002102211212032-2321032311302023-2312132301021221-2013023103122003-2023330233313221-3202312331301211"></a>

<a id="canonical-2010000012020112-3031001203001230-0233112302322232-1113002203133221-1122011313010102-1011221220002330-3303303303332201-0002201213102112"></a>

## uid property — tcp_server_profile / 023200112201 / 8

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

<a id="canonical-2012132132001100-3113333020100322-2232012121213222-1020301011111132-2123302031033122-0300210200300031-3233032332000321-2300221022020012"></a>

## Next pages — tcp_server_profile / 023200112201 / 9

- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0100212231103000-2031131131020010-3003110221032012-0212301320022111-2133230201233310-2233210103212202-3223201320002311-1322101113222313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100331010201223-0033210310232321-0311112111223232-0102012303303320-0333000203103303-2101220311133200-2111221233001012-0332321233021102"></a>

## virtual_server.udp — udp / 132301012102 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.udp

<a id="canonical-0202310031202010-2013230330033220-3203120030300320-3233311000332012-2333000000230213-3011230113132321-2311033300201330-2320323313023303"></a>

Type: `"object"`. single nested block, Optional.

UDP profiles.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
udp {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003233333200111-3111202132211223-3330202002323213-1300311202110300-3131313312301021-3233012301311020-1233101210001030-0021100003102100"></a>

## Direct properties — udp / 132301012102 / 3

- [client_ssl_profile](resources--application_profiles--reference--group-004.md#canonical-1122130210022223-2231220222323121-1323211202123333-3030233201003223-3323032210220331-1201311331312321-1032111223231330-2230322133113123): complete subsection reference.

- [server_ssl_profile](resources--application_profiles--reference--group-004.md#canonical-2021322022310220-0032331120333010-1323213313312110-2023101111222200-2330113101321231-2322220013230220-2100331121121300-2000302300221003): complete subsection reference.

- [udp_client_profile](resources--application_profiles--reference--group-004.md#canonical-0303321111021332-3021213130021133-2210022131031212-1110100113102131-1321300011032303-0202032201130003-1000100312002222-0103020212123120): complete subsection reference.

- [udp_server_profile](resources--application_profiles--reference--group-004.md#canonical-1132032320032120-0030311302123110-3033222311022123-2300300010000331-0320011113122103-2012231120102222-1010023023221231-1103333021102012): complete subsection reference.
