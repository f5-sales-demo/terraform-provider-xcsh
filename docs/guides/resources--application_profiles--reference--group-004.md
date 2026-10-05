---
page_title: "xcsh_application_profiles reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles reference."
---

# xcsh_application_profiles reference

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

<a id="canonical-2021010203222213-3011012132021013-0310231311333133-3023230232031021-0012313123202010-3121223013223110-1012101333302112-3002220000201333"></a>

## Next pages — udp / 132301012102 / 4

- [virtual_server.udp.client_ssl_profile](resources--application_profiles--reference--group-004.md#canonical-1122130210022223-2231220222323121-1323211202123333-3030233201003223-3323032210220331-1201311331312321-1032111223231330-2230322133113123)
- [virtual_server.udp.server_ssl_profile](resources--application_profiles--reference--group-004.md#canonical-2021322022310220-0032331120333010-1323213313312110-2023101111222200-2330113101321231-2322220013230220-2100331121121300-2000302300221003)
- [virtual_server.udp.udp_client_profile](resources--application_profiles--reference--group-004.md#canonical-0303321111021332-3021213130021133-2210022131031212-1110100113102131-1321300011032303-0202032201130003-1000100312002222-0103020212123120)
- [virtual_server.udp.udp_server_profile](resources--application_profiles--reference--group-004.md#canonical-1132032320032120-0030311302123110-3033222311022123-2300300010000331-0320011113122103-2012231120102222-1010023023221231-1103333021102012)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1122130210022223-2231220222323121-1323211202123333-3030233201003223-3323032210220331-1201311331312321-1032111223231330-2230322133113123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300201201133210-2231322203211331-2102222212120220-1002310012111301-2303300001302023-1132003333230112-3300000023131023-0030002013232113"></a>

## virtual_server.udp.client_ssl_profile — client_ssl_profile / 213020332103 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-0100212231103000-2031131131020010-3003110221032012-0212301320022111-2133230201233310-2233210103212202-3223201320002311-1322101113222313)
- virtual_server.udp.client_ssl_profile

<a id="canonical-3033031223130232-3113223130200330-2021133322032131-2121022022233233-0110300131213332-1221213010021110-2033230033101033-1301003103210032"></a>

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

<a id="canonical-1323133131023003-3200122012310202-0003012331012202-0022212322111300-2221133200310132-2212233313203223-1130230010132322-2031032222222031"></a>

## Direct properties — client_ssl_profile / 213020332103 / 3

<a id="canonical-0112221330330033-0123100222330113-0000131223033231-3122001033223013-2313131213121103-3032233332110211-0112002133302311-0012332122121212"></a>

<a id="canonical-0121220313123332-0202322122203202-0223010003321202-2112032101212002-2101112313022213-3200232233203121-3002132213023030-0331313111012321"></a>

## kind property — client_ssl_profile / 213020332103 / 4

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

<a id="canonical-3310003111312011-0110110302200230-1310313220211111-1320110221232332-1232133312013202-1032321212101131-3233002001333230-0003032200000230"></a>

<a id="canonical-3110231011123211-3131021123233121-1230113303131300-1220001220303012-2133300013011223-0112001101211131-2031212031312211-0301313023112313"></a>

## name property — client_ssl_profile / 213020332103 / 5

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

<a id="canonical-2203021033323210-3131000030120021-3332320122220013-1032233323200031-1123322132130213-3232103210212020-0312111222223001-3103011111101010"></a>

<a id="canonical-0113132233003133-0331000220222321-1033123000211020-0330211130201212-1010112001320023-1300000023023010-0211310311221123-1210100210032133"></a>

## namespace property — client_ssl_profile / 213020332103 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3303230300320023-2312211010220130-3120021233220312-2311103213322001-0320231313302113-3200110203101102-1230220101132132-3111200323330002"></a>

<a id="canonical-2020112213100201-2110013300121221-3310233311200021-1233021130213303-2230331323210203-1010113102121220-0233320032330132-2231102313133332"></a>

## tenant property — client_ssl_profile / 213020332103 / 7

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

<a id="canonical-1021230312010001-3323222330013223-0312201330133212-0310300310330323-1130122012220120-3013113132311122-2320000302220020-1022320232202003"></a>

<a id="canonical-1120200102312323-2322313100232113-1333221000103330-2013320130321123-0200220220200132-1332313010320311-2133110321001202-2031231220021100"></a>

## uid property — client_ssl_profile / 213020332103 / 8

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

<a id="canonical-1021001011023313-3230000303201310-2133302003211023-2220231002230032-2312131322130331-1321311213103201-1333323331121022-3212300000300121"></a>

## Next pages — client_ssl_profile / 213020332103 / 9

- [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-0100212231103000-2031131131020010-3003110221032012-0212301320022111-2133230201233310-2233210103212202-3223201320002311-1322101113222313)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2021322022310220-0032331120333010-1323213313312110-2023101111222200-2330113101321231-2322220013230220-2100331121121300-2000302300221003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131023011023011-0303113103220300-2012001322021300-3110113131233303-0003321311312321-2013133003230231-0121131323302013-2321330303321022"></a>

## virtual_server.udp.server_ssl_profile — server_ssl_profile / 211102123013 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-0100212231103000-2031131131020010-3003110221032012-0212301320022111-2133230201233310-2233210103212202-3223201320002311-1322101113222313)
- virtual_server.udp.server_ssl_profile

<a id="canonical-3003003201122202-0223010232323030-0120011011231320-2210202332111112-1123030111021102-0231120121330001-1203103020331123-0003003303322302"></a>

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

<a id="canonical-3023033321031201-0331121001221031-3031031131230030-3330030201220300-1112112310301232-0131302000202320-3332022223011110-0133312320330101"></a>

## Direct properties — server_ssl_profile / 211102123013 / 3

<a id="canonical-3330011101321333-2102122213332322-3331123203310012-1320120133210330-2131113210033133-0021313322130030-2021220000200022-2211021220032103"></a>

<a id="canonical-2320122112212011-2203221020100032-3123212211132021-3300203310332312-2120011020123100-1310210132211103-1312312332023212-2233020111022231"></a>

## kind property — server_ssl_profile / 211102123013 / 4

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

<a id="canonical-1313322120200101-0001100011203130-2222023031332012-1313211011110230-2100322022323232-2300311032020022-3131301121032133-3302213120130011"></a>

<a id="canonical-2330313200132023-0320113003202013-3313020003311120-1303112101222110-1012002232332332-3312133322011311-1320201103320331-0111232202013121"></a>

## name property — server_ssl_profile / 211102123013 / 5

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

<a id="canonical-1221003120222303-2113123323330112-1120232022033123-2211003222320221-1200120022301111-2023021012211121-1322023011013201-0022313301010113"></a>

<a id="canonical-1110103012322130-2030221303232232-1020231032301221-1231320102023333-1332231333132013-2320310133311122-3133202223021312-3012110131121220"></a>

## namespace property — server_ssl_profile / 211102123013 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0222121030223213-1110131020220332-3323322211320002-3131033333022111-0132230222133020-2103303023322220-3022021002231122-3302011210201131"></a>

<a id="canonical-2330023002030113-0122110002032130-1321113201232112-2002132300022131-2112222213300310-0002211300101331-1022011303202202-2013101322230101"></a>

## tenant property — server_ssl_profile / 211102123013 / 7

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

<a id="canonical-3013133102112210-1320121112233011-1133130310232133-1032320211021023-3300021022330023-3221103221010123-1102013113013303-3003200012300220"></a>

<a id="canonical-1010220012320210-0311320112022121-3232231230213132-2202031323000032-0210203133132133-1113103030000322-2313000331313213-1211002111210322"></a>

## uid property — server_ssl_profile / 211102123013 / 8

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

<a id="canonical-2110120022010003-0101321221203010-0303030121012210-0230201133222110-0201110202322313-2032202133022211-0133032323321303-2120231101103022"></a>

## Next pages — server_ssl_profile / 211102123013 / 9

- [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-0100212231103000-2031131131020010-3003110221032012-0212301320022111-2133230201233310-2233210103212202-3223201320002311-1322101113222313)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0303321111021332-3021213130021133-2210022131031212-1110100113102131-1321300011032303-0202032201130003-1000100312002222-0103020212123120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110100332203322-3033030010213320-1331333312331223-1110222231333020-1121022221101330-1022100032320323-1310223311212303-2232321111120231"></a>

## virtual_server.udp.udp_client_profile — udp_client_profile / 221120321103 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-0100212231103000-2031131131020010-3003110221032012-0212301320022111-2133230201233310-2233210103212202-3223201320002311-1322101113222313)
- virtual_server.udp.udp_client_profile

<a id="canonical-3102232222323031-2323123101221120-2230023032121103-1011120012111311-1103321001112222-0301133311233110-0213100312222021-0122013311111330"></a>

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

<a id="canonical-1200023213010122-2210012231123302-3200131312020212-0310323122033322-3213332103020302-3201131110223311-0222231111223312-2122130231201202"></a>

## Direct properties — udp_client_profile / 221120321103 / 3

<a id="canonical-2030030302310122-3211100003112200-1223203112013030-1113312033320103-2123330331330303-0211100120030331-0231211210212131-1113230110001003"></a>

<a id="canonical-0210200333210312-3120023132033233-0101013332332221-0013200301132310-1120211100112021-1122223223320301-3131212333113232-1122210031313332"></a>

## kind property — udp_client_profile / 221120321103 / 4

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

<a id="canonical-3102132213113330-1100000322332332-0213231020133021-0013313220102203-3333123202303023-0003131110020201-1103123231120101-3012111333313211"></a>

<a id="canonical-2102333113301110-0132010310210201-3210032313102032-2122113333010212-1200033001010220-0231132331332033-3103131202132200-1233101013010303"></a>

## name property — udp_client_profile / 221120321103 / 5

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

<a id="canonical-2321111021331113-1232122033013100-3333200333220210-0001322102210200-1331333111123130-2033012002030031-1320102202202203-0323020011312012"></a>

<a id="canonical-2231113320110122-1331332333313020-1111021111302303-2020323301313023-2322130233230210-0021132333030323-1011133011122001-3321023210122001"></a>

## namespace property — udp_client_profile / 221120321103 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0032100333201103-0220001121020133-2221320033230312-2210200323112023-2001120130030213-0302222113003011-0023310301120230-3321030330121022"></a>

<a id="canonical-1333001033321013-1113033300011301-1022333300000023-3220023323323111-3122101033220112-0113013332023131-2333100122122133-3003131331330212"></a>

## tenant property — udp_client_profile / 221120321103 / 7

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

<a id="canonical-1300012133023202-0031103230103113-2020122322002101-3100100020132020-2233112012330303-0201332200033202-2130111122123301-0301023110220103"></a>

<a id="canonical-0033030212310313-2322121033223010-3110031330203133-2013301230030210-0313103132111133-0202111320030301-3012212130022012-2322111330002212"></a>

## uid property — udp_client_profile / 221120321103 / 8

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

<a id="canonical-3003002302223001-2012003110300113-1032130330113201-0010223330231112-3122233230312010-2211300231021231-3101103231201200-3131132021031233"></a>

## Next pages — udp_client_profile / 221120321103 / 9

- [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-0100212231103000-2031131131020010-3003110221032012-0212301320022111-2133230201233310-2233210103212202-3223201320002311-1322101113222313)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1132032320032120-0030311302123110-3033222311022123-2300300010000331-0320011113122103-2012231120102222-1010023023221231-1103333021102012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213030232223212-0031331213023232-1312100320113210-3330320332232200-0230032300313013-0213313331131112-3032212123322330-0103210001023032"></a>

## virtual_server.udp.udp_server_profile — udp_server_profile / 121203103020 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-0100212231103000-2031131131020010-3003110221032012-0212301320022111-2133230201233310-2233210103212202-3223201320002311-1322101113222313)
- virtual_server.udp.udp_server_profile

<a id="canonical-0230123203221322-0330012221231312-0321110103233220-1311131131303031-0212322030321232-1001002021323123-1333323130323301-0303213012323011"></a>

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

<a id="canonical-0021123122101330-2330203112023022-0023020101230111-2303313011330202-3102212202021300-0013220222200222-1003203212102010-1123211210021120"></a>

## Direct properties — udp_server_profile / 121203103020 / 3

<a id="canonical-3333223030203310-3022013212003103-1220121302230210-2023313013012312-0112010202020300-0302002003213321-2311120013310122-3023312002011220"></a>

<a id="canonical-1002310232223102-0230003322331303-2110131120200231-2332132131031001-3133331230001022-2012131221120231-3223203203022102-3132332333122222"></a>

## kind property — udp_server_profile / 121203103020 / 4

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

<a id="canonical-2120003111113001-3231302300033013-2112210312301322-3030223033023233-2100013120213022-3111230331012331-2022020002012113-0033331232310103"></a>

<a id="canonical-1312113023120331-2311011303121022-1112023203132333-2320111303113121-1221001121211023-0320012302010323-0302112200323023-3020320210030111"></a>

## name property — udp_server_profile / 121203103020 / 5

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

<a id="canonical-0001023012113110-0133300100233311-2203122230012223-0000220021112210-3002112020313031-2322210333210311-2301113100333313-1203301001031003"></a>

<a id="canonical-3223033011202312-0131201020330121-2310023033000112-2310322013222323-0100003003002030-0233232222111320-1103100232212312-3021001223210211"></a>

## namespace property — udp_server_profile / 121203103020 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0231311131102033-1200223220033132-2301230000211130-1321232200332202-0030230310020323-2012031321313002-3113310002201220-2022013320310022"></a>

<a id="canonical-2332210310133231-3233333003331031-3313201103011001-0200102103312212-1010022000012213-0232202021100311-3200120322312112-1231012331220221"></a>

## tenant property — udp_server_profile / 121203103020 / 7

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

<a id="canonical-0100303333012021-2200032023233320-2121312121001322-0031313023212230-1202011231110003-3301033311302232-3320333220313223-2012322201011332"></a>

<a id="canonical-1231332101320121-0301221220011020-3031321100210213-1212311231033322-1001030211123202-3223123301332000-1201212213303101-2000211203332331"></a>

## uid property — udp_server_profile / 121203103020 / 8

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

<a id="canonical-3020022313213010-2121201132101300-2333203331322312-1123030113212122-0121221321322321-0221321132113321-0202132323230200-1013122013332311"></a>

## Next pages — udp_server_profile / 121203103020 / 9

- [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-0100212231103000-2031131131020010-3003110221032012-0212301320022111-2133230201233310-2233210103212202-3223201320002311-1322101113222313)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0323021232302203-3330212203221312-0021222111110121-2230020130302203-1311230313200231-3310002022201103-2331300211333032-3021322223300011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230011010102021-0031120032011202-0321223001323110-3230202322333122-3133233203121302-3310323331232310-2130332232132033-0221223322310221"></a>

## virtual_server.virtual_server_state — virtual_server_state / 303112210121 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.virtual_server_state

<a id="canonical-0122303122010231-1321031133131130-1202302030103331-3111033233131023-3101020302031303-3022112323200233-2003020011121311-2213000201203322"></a>

Type: `"object"`. single nested block, Optional.

Displays the current state on the object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("state_disabled",
    "state_enabled")}
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
  "x-ves-oneof-field-state_choice": "[\"state_disabled\",\"state_enabled\"]"
}
```

Terraform syntax:

```terraform
virtual_server_state {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000101010031123-0322230333103011-2021333122220323-3130003013101100-2011333303131031-3333212320022122-2332302002112332-2020322230221203"></a>

## Direct properties — virtual_server_state / 303112210121 / 3

- [state_disabled](resources--application_profiles--reference--group-004.md#canonical-0121331000132220-1313223010021111-1300122021210133-2121010231300102-3330300311303212-0031200013322333-1123123322332330-2010212310333330): complete subsection reference.

- [state_enabled](resources--application_profiles--reference--group-004.md#canonical-2223230303100333-0102311020102121-3202121131020232-2012301223110233-3111211321210113-1033030133000223-2033131020033201-1222211122320011): complete subsection reference.

<a id="canonical-1322121210011022-2200122132011303-2122233102313012-2121332203212103-0003101130100330-2021122130122112-0223312300302303-3001011111332230"></a>

## Next pages — virtual_server_state / 303112210121 / 4

- [virtual_server.virtual_server_state.state_disabled](resources--application_profiles--reference--group-004.md#canonical-0121331000132220-1313223010021111-1300122021210133-2121010231300102-3330300311303212-0031200013322333-1123123322332330-2010212310333330)
- [virtual_server.virtual_server_state.state_enabled](resources--application_profiles--reference--group-004.md#canonical-2223230303100333-0102311020102121-3202121131020232-2012301223110233-3111211321210113-1033030133000223-2033131020033201-1222211122320011)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0121331000132220-1313223010021111-1300122021210133-2121010231300102-3330300311303212-0031200013322333-1123123322332330-2010212310333330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002321120013232-2230120320332232-2131100112003330-1303223003232032-0021332233020301-3102002233203001-3122021302031130-3023011003200023"></a>

## virtual_server.virtual_server_state.state_disabled — state_disabled / 012301030013 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.virtual_server_state](resources--application_profiles--reference--group-004.md#canonical-0323021232302203-3330212203221312-0021222111110121-2230020130302203-1311230313200231-3310002022201103-2331300211333032-3021322223300011)
- virtual_server.virtual_server_state.state_disabled

<a id="canonical-0120301130000121-0203210122211133-1021132101231303-0101332022233021-2133030330213310-3301330232222331-0202031022230300-1121322301311323"></a>

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
state_disabled = {}
```

<a id="canonical-0233030113000330-2332233100232201-2303020023210212-0013132010322103-3313221113313020-2131022033311103-1330202010211101-1011032010001013"></a>

## Direct properties — state_disabled / 012301030013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310302232130021-1301133131201302-0130100210230312-0033010131332311-1012011221233020-3103200123231103-2310133112210133-0202330202213202"></a>

## Next pages — state_disabled / 012301030013 / 4

- [virtual_server.virtual_server_state](resources--application_profiles--reference--group-004.md#canonical-0323021232302203-3330212203221312-0021222111110121-2230020130302203-1311230313200231-3310002022201103-2331300211333032-3021322223300011)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2223230303100333-0102311020102121-3202121131020232-2012301223110233-3111211321210113-1033030133000223-2033131020033201-1222211122320011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121300211010010-1300031212122000-0322133123111322-1021331113023201-2111332231000030-1212330322130211-3021213310101201-1220223112032301"></a>

## virtual_server.virtual_server_state.state_enabled — state_enabled / 133210010300 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.virtual_server_state](resources--application_profiles--reference--group-004.md#canonical-0323021232302203-3330212203221312-0021222111110121-2230020130302203-1311230313200231-3310002022201103-2331300211333032-3021322223300011)
- virtual_server.virtual_server_state.state_enabled

<a id="canonical-2030200221303101-1201013133000300-1132311030000301-3132012300330200-2133023103032200-2032232030013012-1332111122103212-3201113121302101"></a>

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
state_enabled = {}
```

<a id="canonical-0332232331000220-0101200300211100-3002012012101132-1022231331203033-0223103323333321-1202123013133013-3013301311123131-1320110213030331"></a>

## Direct properties — state_enabled / 133210010300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0331112331312113-2323031100003003-0303333111200223-3130022011100330-2012130302010313-3321030122020101-1320301202201313-3331013312232200"></a>

## Next pages — state_enabled / 133210010300 / 4

- [virtual_server.virtual_server_state](resources--application_profiles--reference--group-004.md#canonical-0323021232302203-3330212203221312-0021222111110121-2230020130302203-1311230313200231-3310002022201103-2331300211333032-3021322223300011)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
