---
page_title: "xcsh_token reference"
subcategory: "Identity"
description: "Complete grouped canonical reference for xcsh_token reference."
---

# xcsh_token reference

<a id="canonical-0011220303302322-2103033031012202-3113213332012123-0021203310111012-1122320333101333-3213031012021212-0131120002113213-3022102111320110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320221003233031-0010310231222330-2202331010321122-0123330100333102-3230023321013003-3121011213211112-3013303002221103-3231220300112203"></a>

## Property reference — Property reference / 301003122332 / 2

Breadcrumbs:

- [xcsh_token](../resources/token.md#canonical-1033011013002132-3200233320301222-0133100331111013-2201110031233021-3032032333322002-0033103111032332-0232210133220103-2232103023033020)
- Property reference

<a id="canonical-2100021323003220-1003033102012220-2332130132000021-0331010202310303-3321032202330211-3010101122133211-2212333100011203-3032120101232030"></a>

## Direct properties — Property reference / 301003122332 / 3

<a id="canonical-2000203022200211-0130122213313112-2020111320312321-3313000223023123-0013231010222201-0231203032321010-1022032100122011-1212320302000300"></a>

<a id="canonical-1321000213013111-1212020210312232-0120020320030100-1312313023103301-0002303202012122-2333012333022011-0011321202201303-1113121023013323"></a>

## annotations property — Property reference / 301003122332 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-2120210031323332-1201323210003102-0113302311133203-1211313303210013-0120012010332200-1132023222111033-2323221013102331-1012320331130121"></a>

<a id="canonical-1300030131232123-0133101131102311-2122103023021233-2301113011023103-1200021222220223-2222221121301103-0020211222330230-0131311023301213"></a>

## content property — Property reference / 301003122332 / 5

Type: `"string"`. Computed, Sensitive.

Server-issued JWT registration credential.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-sensitive": true
}
```

<a id="canonical-0123322000230203-0211213023002201-3023101311023030-3010131220321100-1310123033200203-3220313003302302-1312222112301032-0210100031033021"></a>

<a id="canonical-2131133021131310-0000132031131300-3321303231213031-0022313323311132-0200120331013012-1022110031200000-1020201303132202-2221320122033133"></a>

## description property — Property reference / 301003122332 / 6

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-2200111230013231-0310321303223310-2221333003033200-3231121331002110-3123011013111132-3201323120211120-3210020211310123-1302212221232310"></a>

<a id="canonical-2323112223223202-1221023122333120-0111032203022023-3313031013332303-3303201031001021-3010311000103031-3003330123323311-2023120230200202"></a>

## disable property — Property reference / 301003122332 / 7

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

<a id="canonical-3010221302021011-0121121313333120-0130210012330200-1100100110231000-3103012003221310-3001111231121213-3333302300011213-3201323010302220"></a>

<a id="canonical-3131023030003100-0233213102201213-0301202123222022-1202003223110122-3013030022212220-1001010122312011-1213130320123023-2002230202100222"></a>

## ID property — Property reference / 301003122332 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1313313303203120-3030322020320003-0023321133131311-3333033303003011-0132222031123031-3210030002312030-3211131000231021-3230311010200231"></a>

<a id="canonical-1010200013102023-0031310101210121-3110131010102031-2000231003023220-3333220230221231-2022301133202322-3220030300031223-1232010031012023"></a>

## labels property — Property reference / 301003122332 / 9

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="canonical-2331323303332011-1313230301310301-3101223332120333-0012022031330010-1210313103202201-2101202321320223-3302013120300010-2100122112223233"></a>

<a id="canonical-0332313001111232-2311022021222310-1321322011001202-2122123232203103-2222100110303230-0233101013011120-0101213023002120-0030202321200023"></a>

## name property — Property reference / 301003122332 / 10

Type: `"string"`. Required.

Name of the Token. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NameValidator(),
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1110200201033032-3012321313012110-0310110330312113-1130113201110110-3103313013323001-2323312101103201-0331003020302113-0122102032001303"></a>

<a id="canonical-0310312133013303-2130333020210213-1313112001013132-2011230311122313-3200202213230021-3212321231132133-1223100311223023-3020103032222333"></a>

## namespace property — Property reference / 301003122332 / 11

Type: `"string"`. Optional, Computed.

Namespace for the Token. The F5 XC API restricts this resource to the system namespace; it defaults
to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
EnumExtractionComplete: false
EnumValidators: [{"version":1,"validator":"OneOf","values":["system"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

<a id="canonical-0300200203211033-2313210211111301-2030031310011100-3011123233200030-1030113030132221-1110001001230300-0230113222220011-3013011322120303"></a>

<a id="canonical-0212201022332211-2302213111023101-1012031313130301-2033331012222100-1313200302021010-0000122000132301-2012023113131323-3100200333123022"></a>

## site_name property — Property reference / 301003122332 / 12

Type: `"string"`. Optional, Computed.

Secure Mesh Site v2 name bound into a JWT token.

- [timeouts](resources--token--reference--group-001.md#canonical-3121231312030123-0333320233101011-1021110000110333-0023122210230311-2333030012111020-3311200302211321-0203002333021021-1310322211101202): complete subsection reference.

<a id="canonical-2231231220100231-1022303321011001-2321023320012213-0302003120211022-1103013023213000-2012202202220320-2211331100233223-2023103113321212"></a>

<a id="canonical-1231021202102023-2202201313311303-3331101010112130-1323000032313202-1201220000312313-1312120220130111-3220213313311301-3103123333021311"></a>

## type property — Property reference / 301003122332 / 13

Type: `"number"`. Optional, Computed.

\[Enum: 0|1\] Token type, where 0 is NORMAL and 1 is JWT. Possible values are \`0\`, \`1\`.

Upstream description:

Token type, where 0 is NORMAL and 1 is JWT.

Receipt-pinned upstream constraints:

```json
{
  "default": 0,
  "enum": [
    0,
    1
  ]
}
```

<a id="canonical-0300322133030031-2131001101200302-0133212321020132-3222213300320323-3010130222120020-1033100221032111-1111011200001332-3121012303121222"></a>

<a id="canonical-0000021102333100-0122201012310112-2103210032031103-3110322100322011-2203320121123221-0210323022322310-0312021222202131-0313012233212213"></a>

## uid property — Property reference / 301003122332 / 14

Type: `"string"`. Computed, Sensitive.

Effective sensitive CE registration credential. NORMAL tokens use \`system\_metadata.uid\`; JWT
tokens use \`spec.content\`. This value is stored in plain text in the Terraform state file; ensure
your state file is properly secured.

<a id="canonical-2301211030332221-0133331033113230-1030102131231103-0221002200322033-2222012320230123-3232202310320121-2021220031100323-2332233131211012"></a>

## All schema paths — Property reference / 301003122332 / 15

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--token--reference--group-001.md#canonical-2000203022200211-0130122213313112-2020111320312321-3313000223023123-0013231010222201-0231203032321010-1022032100122011-1212320302000300) |
| `content` | [content](resources--token--reference--group-001.md#canonical-2120210031323332-1201323210003102-0113302311133203-1211313303210013-0120012010332200-1132023222111033-2323221013102331-1012320331130121) |
| `description` | [description](resources--token--reference--group-001.md#canonical-0123322000230203-0211213023002201-3023101311023030-3010131220321100-1310123033200203-3220313003302302-1312222112301032-0210100031033021) |
| `disable` | [disable](resources--token--reference--group-001.md#canonical-2200111230013231-0310321303223310-2221333003033200-3231121331002110-3123011013111132-3201323120211120-3210020211310123-1302212221232310) |
| `id` | [ID](resources--token--reference--group-001.md#canonical-3010221302021011-0121121313333120-0130210012330200-1100100110231000-3103012003221310-3001111231121213-3333302300011213-3201323010302220) |
| `labels` | [labels](resources--token--reference--group-001.md#canonical-1313313303203120-3030322020320003-0023321133131311-3333033303003011-0132222031123031-3210030002312030-3211131000231021-3230311010200231) |
| `name` | [name](resources--token--reference--group-001.md#canonical-2331323303332011-1313230301310301-3101223332120333-0012022031330010-1210313103202201-2101202321320223-3302013120300010-2100122112223233) |
| `namespace` | [namespace](resources--token--reference--group-001.md#canonical-1110200201033032-3012321313012110-0310110330312113-1130113201110110-3103313013323001-2323312101103201-0331003020302113-0122102032001303) |
| `site_name` | [site_name](resources--token--reference--group-001.md#canonical-0300200203211033-2313210211111301-2030031310011100-3011123233200030-1030113030132221-1110001001230300-0230113222220011-3013011322120303) |
| `timeouts` | [timeouts](resources--token--reference--group-001.md#canonical-2012322233222230-1033320013120321-0133011003110022-2003213213003233-2022312310031130-2200003322130020-1222102313301230-1303132323012322) |
| `timeouts.create` | [timeouts.create](resources--token--reference--group-001.md#canonical-0233122202021113-3232103100312310-0021222300031211-2011323031230031-0122303330210312-1201212130223112-0333120203232300-1101103330210331) |
| `timeouts.delete` | [timeouts.delete](resources--token--reference--group-001.md#canonical-2211230202221302-0121331333020222-1222331103311032-3312233302022032-2222200300220003-3322210120102002-2131330030120133-0211031112021031) |
| `timeouts.read` | [timeouts.read](resources--token--reference--group-001.md#canonical-2121231032213331-2002330120122132-1323113230032123-1210311303123101-3003212032103212-0302112233223202-1131101221320313-1003100010323231) |
| `timeouts.update` | [timeouts.update](resources--token--reference--group-001.md#canonical-2312130213010333-3223110301322023-3200303001123022-2210301223300032-3023313213000003-1231132210101210-3131201302320213-1322133013113321) |
| `type` | [type](resources--token--reference--group-001.md#canonical-2231231220100231-1022303321011001-2321023320012213-0302003120211022-1103013023213000-2012202202220320-2211331100233223-2023103113321212) |
| `uid` | [uid](resources--token--reference--group-001.md#canonical-0300322133030031-2131001101200302-0133212321020132-3222213300320323-3010130222120020-1033100221032111-1111011200001332-3121012303121222) |

<a id="canonical-0320221000123020-0112010023020321-1323203320123032-2022200321230311-3302312110302102-0101023013220202-0203133202113013-3000211301210223"></a>

## Next pages — Property reference / 301003122332 / 16

- [timeouts](resources--token--reference--group-001.md#canonical-3121231312030123-0333320233101011-1021110000110333-0023122210230311-2333030012111020-3311200302211321-0203002333021021-1310322211101202)
- [xcsh_token](../resources/token.md#canonical-1033011013002132-3200233320301222-0133100331111013-2201110031233021-3032032333322002-0033103111032332-0232210133220103-2232103023033020)

<a id="canonical-3121231312030123-0333320233101011-1021110000110333-0023122210230311-2333030012111020-3311200302211321-0203002333021021-1310322211101202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210301310202022-0002000302223233-0303202201033102-3111231302230111-3022021221122010-2221320020122022-0002112033321102-2133130231013313"></a>

## timeouts — timeouts / 230302331103 / 2

Breadcrumbs:

- [xcsh_token](../resources/token.md#canonical-1033011013002132-3200233320301222-0133100331111013-2201110031233021-3032032333322002-0033103111032332-0232210133220103-2232103023033020)
- [Property reference](resources--token--reference--group-001.md#canonical-0011220303302322-2103033031012202-3113213332012123-0021203310111012-1122320333101333-3213031012021212-0131120002113213-3022102111320110)
- timeouts

<a id="canonical-2012322233222230-1033320013120321-0133011003110022-2003213213003233-2022312310031130-2200003322130020-1222102313301230-1303132323012322"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000232310100110-3313223232101320-3301212102222000-2022313330010113-2020102013211301-1133033111313202-1212330211202111-3221210203122133"></a>

## Direct properties — timeouts / 230302331103 / 3

<a id="canonical-0233122202021113-3232103100312310-0021222300031211-2011323031230031-0122303330210312-1201212130223112-0333120203232300-1101103330210331"></a>

<a id="canonical-3303330232010123-3203022233032121-3213011122311313-2223113123331020-2100313302213310-2101200201112013-1100033020233313-2321203312012110"></a>

## create property — timeouts / 230302331103 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2211230202221302-0121331333020222-1222331103311032-3312233302022032-2222200300220003-3322210120102002-2131330030120133-0211031112021031"></a>

<a id="canonical-3213012323310100-1132312100010011-0202222310313102-1221211132111221-2310223301311110-0322321321212200-2212112303013333-0031212320233201"></a>

## delete property — timeouts / 230302331103 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2121231032213331-2002330120122132-1323113230032123-1210311303123101-3003212032103212-0302112233223202-1131101221320313-1003100010323231"></a>

<a id="canonical-1222032131010220-3301233302011301-2220302231223130-3031100133032303-0031123331013102-3212221233231021-3000010301332013-2021213230122110"></a>

## read property — timeouts / 230302331103 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2312130213010333-3223110301322023-3200303001123022-2210301223300032-3023313213000003-1231132210101210-3131201302320213-1322133013113321"></a>

<a id="canonical-1303232320020101-0112311320122332-3223021103302023-0000121000330010-1313030202300003-0030122223230111-3320300100112010-3111011113212302"></a>

## update property — timeouts / 230302331103 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3111002221310322-0123313312013103-0033312022202123-0013011003123103-2203322302220312-3113321222033133-1113330000210011-1203000022023313"></a>

## Next pages — timeouts / 230302331103 / 8

- [Property reference](resources--token--reference--group-001.md#canonical-0011220303302322-2103033031012202-3113213332012123-0021203310111012-1122320333101333-3213031012021212-0131120002113213-3022102111320110)
- [xcsh_token](../resources/token.md#canonical-1033011013002132-3200233320301222-0133100331111013-2201110031233021-3032032333322002-0033103111032332-0232210133220103-2232103023033020)
