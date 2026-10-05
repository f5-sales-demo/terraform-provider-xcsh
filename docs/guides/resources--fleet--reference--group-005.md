---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-0212010000100332-3003320021133202-0012010123313120-2331202032103001-3000203320321112-3321331121022120-2112032203221112-3220010000113102"></a>

## update property — timeouts / 303003032320 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2301111321133110-3130231223111012-2200231110210030-1120121022032322-3032113333003303-3021133110011200-1013021121032321-3221322120113203"></a>

## Next pages — timeouts / 303003032320 / 8

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1313133033021303-0322312200331123-0101000033230330-1300123032023012-1022323323332131-3002113220222232-0302220132222213-1022313213313302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011313030230232-2021332310120103-3222000030330132-2013220010301320-1011300102021023-2233333103103101-0002230000103023-0130010233213232"></a>

## usb_policy — usb_policy / 222203321332 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- usb_policy

<a id="canonical-1122110212330113-3101131220333203-2013203030011212-2310131001313330-1321230331331320-0133123320203112-1102233320013230-2122202121023032"></a>

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
usb_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011300300110001-1300113313023110-2020013031213110-0333302033320131-2131320201320113-0213223331210331-2301302132120223-3133033013300130"></a>

## Direct properties — usb_policy / 222203321332 / 3

<a id="canonical-0210031313123110-1013201323233220-3321022023222200-2131332103232323-2203032111130102-3313032333332030-0000332002311233-0320333133001220"></a>

<a id="canonical-2321002121312331-0100023330212123-2301033021003333-2020203003022013-2102022320133001-0120332212033002-1222332203300300-1310013300022332"></a>

## name property — usb_policy / 222203321332 / 4

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

<a id="canonical-2000203111211212-1122201033323001-3310033331330103-3320330022113132-0300023133131111-0131201313122323-2113011103200123-1002331020122212"></a>

<a id="canonical-2013321200011212-3303201203312000-3020301302212320-2330113223132203-3113303231200121-0330210001110101-2111122102221332-3110231123031323"></a>

## namespace property — usb_policy / 222203321332 / 5

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

<a id="canonical-1322013111303300-3212302013200320-3303113031311132-2220211312101300-3130101033211130-2010232231023103-3230111312312322-2223111210320202"></a>

<a id="canonical-1313223330032231-1232230302121022-3012122033312101-1312112130103012-3302200021310302-0211200103211012-0300033310322332-1313322103132111"></a>

## tenant property — usb_policy / 222203321332 / 6

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

<a id="canonical-1002020313123210-1330022233133010-1301130113331231-0221003332220202-1121000012132122-0200320033230212-1320322113013130-2322111211333102"></a>

## Next pages — usb_policy / 222203321332 / 7

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
