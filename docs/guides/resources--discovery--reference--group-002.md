---
page_title: "xcsh_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_discovery reference."
---

# xcsh_discovery reference

<a id="canonical-3130131030233030-2300332133130303-1311003313120133-1100303031322023-0023201010022321-3003020130300102-2101032002323232-0212210012120013"></a>

## where.site.enable_internet_vip — enable_internet_vip / 020211003020 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-001.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- [where.site](resources--discovery--reference--group-001.md#canonical-2111032303022223-2300132122211213-0100203302230222-1011230233321231-1103310111023200-0013013033313032-1103302322310010-1210312011210200)
- where.site.enable_internet_vip

<a id="canonical-2013232001120033-0302203321031212-1133020011332302-1333012221313120-2000331022223130-1101021100312211-0222032331230000-0230320102112321"></a>

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
enable_internet_vip = {}
```

<a id="canonical-3020020323231310-2313312212011230-0300200001213320-2100010103300213-1222003301110031-1310221101012332-2332222302331222-3020310301013303"></a>

## Direct properties — enable_internet_vip / 020211003020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223032032123300-0330213033323010-3020313230210031-2113033332010320-1312021331231101-3022020203111002-2212022303030212-0202230332103210"></a>

## Next pages — enable_internet_vip / 020211003020 / 4

- [where.site](resources--discovery--reference--group-001.md#canonical-2111032303022223-2300132122211213-0100203302230222-1011230233321231-1103310111023200-0013013033313032-1103302322310010-1210312011210200)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-2032210203231300-1003003000132213-3101202330201122-0001100213311020-0131210232302123-3012213102000232-2331302321030002-1313330003132301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002200330012003-3131013313111130-3231203222012133-0122002002320133-1210301220231323-3311120311001222-2100213101333210-3120021101121203"></a>

## where.site.ref — ref / 132311013000 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-001.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- [where.site](resources--discovery--reference--group-001.md#canonical-2111032303022223-2300132122211213-0100203302230222-1011230233321231-1103310111023200-0013013033313032-1103302322310010-1210312011210200)
- where.site.ref

<a id="canonical-2033133220303202-3013313000222302-2301200020221123-1122313121332100-3130122201311311-0000112123213332-1303231310133313-1202031022130122"></a>

Type: `"object"`. list nested block, Optional.

Reference. A site direct reference.

Upstream description:

A site direct reference.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-3032030203201003-3311221223223133-1022320331330313-1332212132002113-1131220021100010-1222112333111001-0303320021302313-0101033031030222"></a>

## Direct properties — ref / 132311013000 / 3

<a id="canonical-1113130301320303-0102201023122131-3030032223203001-1230013232303200-0011222101202102-0123132303213222-3202320010210311-3121021213212102"></a>

<a id="canonical-3320210003221030-3331122330333020-1122321031001002-1031231011130103-0221301020223203-1331200022232010-3331200330210032-1032122111120101"></a>

## kind property — ref / 132311013000 / 4

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

<a id="canonical-3031320213210302-3220122121323100-3210233133033321-1323022011022011-1200100211232201-0330332302310320-3033221000311312-1101131220220030"></a>

<a id="canonical-2232320303211301-0021221133330001-2111013030103113-1023313300221013-1122031322102221-2020302120320102-2313101300312023-1312110321233223"></a>

## name property — ref / 132311013000 / 5

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

<a id="canonical-2200112313310013-2202311201323300-3333100133332133-0002320320303203-2123202220032331-2330031322020203-2012123023330220-3100130010330302"></a>

<a id="canonical-1310333031022200-3300012323011311-0110312030013020-3310232301013210-3210233100020302-1220130132110110-3023021331023233-3022000330202020"></a>

## namespace property — ref / 132311013000 / 6

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

<a id="canonical-2330123230100021-3011022000123301-3010211231231122-2233001123111213-0111131030321332-2012031031200121-3220312312031010-3200110301112200"></a>

<a id="canonical-0012313331222030-1323200103332320-1131110203322301-2121202320012303-2232030221021213-0300200131320032-0011123321331332-2322213221011131"></a>

## tenant property — ref / 132311013000 / 7

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

<a id="canonical-3302231323230211-1230022212113010-2132212230313210-1112211220021102-2201112103222210-2121122200113332-1232212331001002-3133130123122231"></a>

<a id="canonical-1111231300120323-0331031232211210-2022322110013323-3031311021320221-3232013223021322-1221123100230201-1331123003321121-0103212200030123"></a>

## uid property — ref / 132311013000 / 8

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

<a id="canonical-2330122302033332-3122101101312121-3100013301223232-0313302013231011-2133201321332001-0011211112103233-2323231310210032-1023321321310333"></a>

## Next pages — ref / 132311013000 / 9

- [where.site](resources--discovery--reference--group-001.md#canonical-2111032303022223-2300132122211213-0100203302230222-1011230233321231-1103310111023200-0013013033313032-1103302322310010-1210312011210200)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-1110120031123232-1103133033221300-2333022012020332-2112320331303032-0100001133321112-2231130320002221-2033320232220130-0031211123210313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002110222330132-2331301213110320-0200201123030012-3333111202102111-1112003211201023-2233113111232112-0023012111202121-2111331000120111"></a>

## where.virtual_network — virtual_network / 112113321313 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-001.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- where.virtual_network

<a id="canonical-0232303211103312-3331210310230232-2233302213331130-0103210120311211-0000201320122332-3000300231333221-2002100223333131-2311110230300022"></a>

Type: `"object"`. single nested block, Optional.

Specifies a direct reference to a network configuration object.

Upstream description:

This specifies a direct reference to a network configuration object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ref")}
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

<a id="canonical-2002320232011120-0021303211210220-2100323300312000-2211301312023203-3310111221022110-3002120013033213-2333011033123021-2120120120103201"></a>

## Direct properties — virtual_network / 112113321313 / 3

- [ref](resources--discovery--reference--group-002.md#canonical-3312231132303023-0313321002210120-1100032032133210-2222122233111210-0001122010132331-0103313320302022-3012202230320312-2112101111311001): complete subsection reference.

<a id="canonical-2003100322123010-2122330303013312-3211012232212122-3031323301012023-0202001300133032-2311221212103000-1102313023322021-0333300022321003"></a>

## Next pages — virtual_network / 112113321313 / 4

- [where.virtual_network.ref](resources--discovery--reference--group-002.md#canonical-3312231132303023-0313321002210120-1100032032133210-2222122233111210-0001122010132331-0103313320302022-3012202230320312-2112101111311001)
- [where](resources--discovery--reference--group-001.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-3312231132303023-0313321002210120-1100032032133210-2222122233111210-0001122010132331-0103313320302022-3012202230320312-2112101111311001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011201123211022-2122013223313033-0202232100330333-3331032131333021-2312313100010020-1232101211022020-3011021200300111-2332211223113102"></a>

## where.virtual_network.ref — ref / 112100122223 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-001.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- [where.virtual_network](resources--discovery--reference--group-002.md#canonical-1110120031123232-1103133033221300-2333022012020332-2112320331303032-0100001133321112-2231130320002221-2033320232220130-0031211123210313)
- where.virtual_network.ref

<a id="canonical-0313030022322130-0201012102111211-1323203223021023-3221032231313023-2111033001103001-1132232322001222-1111311222123012-0332300301000303"></a>

Type: `"object"`. list nested block, Optional.

Reference. A virtual network direct reference.

Upstream description:

A virtual network direct reference.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223023010133302-2003201332310312-0032112131120311-0303213021132223-2232220113133021-0300131203013302-0132010132121330-2312303330222013"></a>

## Direct properties — ref / 112100122223 / 3

<a id="canonical-2310303223212320-0321110030100012-0212032031110030-1130130311202121-1230312032102123-2131033233032120-1233201003302012-2213110133030220"></a>

<a id="canonical-0130301121230310-2121220322300022-0113122111011233-1012123130330220-3202223031202230-1320333232112002-0220032112113103-1013132132020030"></a>

## kind property — ref / 112100122223 / 4

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

<a id="canonical-1101213121031030-1010311321101310-2001033232001130-0101113320121100-1011000021311202-0020323232033133-3100213301032230-3332132120111020"></a>

<a id="canonical-3102220232223202-2210233012002001-1002332011132003-3103212313331230-1311313212231031-2223102000233032-2011001202102223-2233332231102302"></a>

## name property — ref / 112100122223 / 5

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

<a id="canonical-2230332102131303-2302330121300100-2031333221222002-3221200012111112-2203001021012013-2112222202300313-2100033221301220-3322312033320301"></a>

<a id="canonical-1210233302303003-2323112012331120-1213223233103232-3311022330031131-1032132321203331-0001301310022122-3120203201201220-0122233102013203"></a>

## namespace property — ref / 112100122223 / 6

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

<a id="canonical-2303110203312302-2123030210132103-3330002012000123-0313233002221332-2100021101111202-2113323111313320-0311032233113213-0133322332013222"></a>

<a id="canonical-0021003111020012-3030110132112032-3131023313311230-1221032130100302-3200300213310022-3210121131213133-1021221211231203-1211130031023200"></a>

## tenant property — ref / 112100122223 / 7

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

<a id="canonical-1330031302023030-1020030301031332-3120221321302000-2120021120311320-1033000332231101-1131022222011001-1133321202003102-1201031121313323"></a>

<a id="canonical-0313011121110323-3211303201113301-1322011322310313-0231210030003020-2100333303030100-2100313302021333-3331202203102130-2112310022221331"></a>

## uid property — ref / 112100122223 / 8

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

<a id="canonical-2323110230112212-3021002231112110-1233200020002003-3301303022203011-3101323233313030-1133200031330323-0323022312111231-1322333130133201"></a>

## Next pages — ref / 112100122223 / 9

- [where.virtual_network](resources--discovery--reference--group-002.md#canonical-1110120031123232-1103133033221300-2333022012020332-2112320331303032-0100001133321112-2231130320002221-2033320232220130-0031211123210313)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-2223322232211101-1111031212021233-2331110221020110-1113201232123120-1032003111222123-1221012223203322-0130200312133221-2222033300110330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311132013333132-3103023120120022-0331032332221013-1003332111032211-2102310102313011-3031002330113012-2212120110103302-0310133310333310"></a>

## where.virtual_site — virtual_site / 323202311220 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-001.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- where.virtual_site

<a id="canonical-0322133000130021-2333312212222111-0020101213231023-2030310233221013-0020031031000102-3020113021303003-1120101000110032-0110002030131332"></a>

Type: `"object"`. single nested block, Optional.

Virtual Site. A reference to virtual\_site object.

Upstream description:

A reference to virtual\_site object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ref"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip")}
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
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303030313021123-2330030223120001-0233122030132333-1210330223021333-1001201101003102-3323101212230033-2132133112120110-2022110120032230"></a>

## Direct properties — virtual_site / 323202311220 / 3

- [disable_internet_vip](resources--discovery--reference--group-002.md#canonical-0012000002320212-1110003030333201-2200132323301111-0031003322022332-2131310131100122-1020102030213221-1322011213212030-0212233130001000): complete subsection reference.

- [enable_internet_vip](resources--discovery--reference--group-002.md#canonical-0300132310111200-0202012132102100-3100103102212123-0313102300221333-3102331310321103-2121102033002200-2111330113031323-1011211010002010): complete subsection reference.

<a id="canonical-2000131222330332-3332323003332032-2221323000032331-0203112013033230-3122310103120230-2102103122211223-3222321131122303-0133011000320313"></a>

<a id="canonical-1020103313302131-1102211213110230-0223330202303230-1020103102232111-0212232223213113-0330201102222130-0130213103221111-0113223311111012"></a>

## network_type property — virtual_site / 323202311220 / 4

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VIRTUAL_NETWORK_GLOBAL","VIRTUAL_NETWORK_IP_AUTO","VIRTUAL_NETWORK_IP_FABRIC","VIRTUAL_NETWORK_MANAGEMENT","VIRTUAL_NETWORK_PER_SITE","VIRTUAL_NETWORK_PUBLIC","VIRTUAL_NETWORK_SEGMENT","VIRTUAL_NETWORK_SITE_LOCAL","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE","VIRTUAL_NETWORK_SITE_SERVICE","VIRTUAL_NETWORK_SRV6_NETWORK","VIRTUAL_NETWORK_VER_INTERNAL","VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](resources--discovery--reference--group-002.md#canonical-2313102201211221-1123332010222132-3130002111113333-1310230310302021-1132233010120030-2130230020331133-3331301032013003-3310010212002322): complete subsection reference.

<a id="canonical-3231101323211113-2301212131130110-1033233313311321-3010100130002110-1000022120131022-1313202112010021-1000233312313300-0131002023203313"></a>

## Next pages — virtual_site / 323202311220 / 5

- [where.virtual_site.disable_internet_vip](resources--discovery--reference--group-002.md#canonical-0012000002320212-1110003030333201-2200132323301111-0031003322022332-2131310131100122-1020102030213221-1322011213212030-0212233130001000)
- [where.virtual_site.enable_internet_vip](resources--discovery--reference--group-002.md#canonical-0300132310111200-0202012132102100-3100103102212123-0313102300221333-3102331310321103-2121102033002200-2111330113031323-1011211010002010)
- [where.virtual_site.ref](resources--discovery--reference--group-002.md#canonical-2313102201211221-1123332010222132-3130002111113333-1310230310302021-1132233010120030-2130230020331133-3331301032013003-3310010212002322)
- [where](resources--discovery--reference--group-001.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-0012000002320212-1110003030333201-2200132323301111-0031003322022332-2131310131100122-1020102030213221-1322011213212030-0212233130001000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033000012310111-2031112032330312-2330201121232301-2103031200022030-0130203232132013-0101123132323013-2301002212133121-0222003322120013"></a>

## where.virtual_site.disable_internet_vip — disable_internet_vip / 130233200211 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-001.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- [where.virtual_site](resources--discovery--reference--group-002.md#canonical-2223322232211101-1111031212021233-2331110221020110-1113201232123120-1032003111222123-1221012223203322-0130200312133221-2222033300110330)
- where.virtual_site.disable_internet_vip

<a id="canonical-2312101123222101-2011011201100032-2221012332123301-3002130230322301-1332302020000333-1302213311003311-3230331320112310-2330320003131110"></a>

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
disable_internet_vip = {}
```

<a id="canonical-2320102233311230-1011012301210210-0101223332221221-3003220130332320-1000310312303322-3000002123202220-3230302302020030-0333212202013121"></a>

## Direct properties — disable_internet_vip / 130233200211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120002101310200-0200231322313133-1311110133132332-1001221232000220-0122023110220030-0200223230102320-2330332011212021-1212211111220113"></a>

## Next pages — disable_internet_vip / 130233200211 / 4

- [where.virtual_site](resources--discovery--reference--group-002.md#canonical-2223322232211101-1111031212021233-2331110221020110-1113201232123120-1032003111222123-1221012223203322-0130200312133221-2222033300110330)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-0300132310111200-0202012132102100-3100103102212123-0313102300221333-3102331310321103-2121102033002200-2111330113031323-1011211010002010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201311031121210-0213003322131102-0232231112113303-2331022001003310-1203030332332012-2112033212110301-2111230130232230-3302121210113331"></a>

## where.virtual_site.enable_internet_vip — enable_internet_vip / 023332023230 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-001.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- [where.virtual_site](resources--discovery--reference--group-002.md#canonical-2223322232211101-1111031212021233-2331110221020110-1113201232123120-1032003111222123-1221012223203322-0130200312133221-2222033300110330)
- where.virtual_site.enable_internet_vip

<a id="canonical-3021002022120123-1022010032122021-1110031333111120-2230322030033022-2320033033023333-2330031110211102-3331023012233013-1113232210103211"></a>

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
enable_internet_vip = {}
```

<a id="canonical-2332221212312011-1222010013323103-3000111223022303-1212233211311110-0331101012013130-1201231231210033-3201000113312100-0110300311301030"></a>

## Direct properties — enable_internet_vip / 023332023230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100302112013012-3000212031133303-3332103120230203-0123311120111001-0213330302212010-1320323132001112-2111333322232332-2210310311211331"></a>

## Next pages — enable_internet_vip / 023332023230 / 4

- [where.virtual_site](resources--discovery--reference--group-002.md#canonical-2223322232211101-1111031212021233-2331110221020110-1113201232123120-1032003111222123-1221012223203322-0130200312133221-2222033300110330)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-2313102201211221-1123332010222132-3130002111113333-1310230310302021-1132233010120030-2130230020331133-3331301032013003-3310010212002322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231010303212211-0023120113230030-3333231321002131-2130101012130133-0111231331301120-0112302120111210-0202220021002122-1300201121111211"></a>

## where.virtual_site.ref — ref / 101102132321 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-001.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- [where.virtual_site](resources--discovery--reference--group-002.md#canonical-2223322232211101-1111031212021233-2331110221020110-1113201232123120-1032003111222123-1221012223203322-0130200312133221-2222033300110330)
- where.virtual_site.ref

<a id="canonical-2223002323111232-1312332333333231-1100100132112221-0012333303020122-3102332001111121-1232303103331221-2230332023123001-3031110313130223"></a>

Type: `"object"`. list nested block, Optional.

Reference. A virtual\_site direct reference.

Upstream description:

A virtual\_site direct reference.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-2011012132322233-3321001123303203-1100010203330303-1133103121123120-3121101211012321-2313323122023330-0320112212210223-0331002221322223"></a>

## Direct properties — ref / 101102132321 / 3

<a id="canonical-2310233223102032-3200231202102012-0332102020222202-2321213203231033-0131300210132221-2122310230200102-3213023230113112-2123333320120000"></a>

<a id="canonical-3133131120333011-1320232223002011-3321131031133021-1311020320003011-2311130132030330-1302003231003220-3030112323033123-0230102220123013"></a>

## kind property — ref / 101102132321 / 4

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

<a id="canonical-2321321030101011-1330320131023013-0333321210013022-1000011201113301-1300033212122000-3213323232321100-0210203301323011-0222000031001100"></a>

<a id="canonical-2231203131101323-0002113030323202-2100320303010001-2021222230230123-0003202313031020-3202310201021032-1010213113122212-3102232312122121"></a>

## name property — ref / 101102132321 / 5

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

<a id="canonical-1310133233122030-3311122223101211-1231133000123210-1233003002231310-0232031212202101-0000111202002232-0013132332033210-2002201110021133"></a>

<a id="canonical-3213333300111022-0130221231310203-3320202201230102-3213313203222230-1202023231333120-3121322303302000-2322322110231120-1203002022013033"></a>

## namespace property — ref / 101102132321 / 6

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

<a id="canonical-0130012332031021-3010032323230100-1022321222223111-3033330120201333-1313333323313330-0131100032320203-2032021301312030-0011231233211132"></a>

<a id="canonical-0333211333100302-3213311112101033-0021310101331333-2012212322133300-2202033200130212-0231113232213022-0101331231310123-3230000111312202"></a>

## tenant property — ref / 101102132321 / 7

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

<a id="canonical-3211223011011211-1221223322021300-1303303103320033-1103022230301222-3320010130100231-0101312233120020-2030010012112132-1120233103310133"></a>

<a id="canonical-3230130103003123-2020010230323313-1201313133123111-1310022112210311-1230232321000220-2031323002330003-1012001123113023-0213023220010123"></a>

## uid property — ref / 101102132321 / 8

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

<a id="canonical-0302323233303222-2302212211012330-2123020210122021-0301300002012021-0331033121122331-1130200220013123-0110211220212321-3113010002002331"></a>

## Next pages — ref / 101102132321 / 9

- [where.virtual_site](resources--discovery--reference--group-002.md#canonical-2223322232211101-1111031212021233-2331110221020110-1113201232123120-1032003111222123-1221012223203322-0130200312133221-2222033300110330)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
