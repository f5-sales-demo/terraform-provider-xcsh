---
page_title: "xcsh_bot_infrastructure reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_infrastructure reference."
---

# xcsh_bot_infrastructure reference

<a id="canonical-3211120122133210-0122110022200232-1211012113212211-2130003230230030-1201022221000022-0012221313230312-0000102002302203-0201002321021101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233223300211110-2331122112311130-0102110230320010-3210003100233032-2231333103123231-1101003013132103-3303023203033011-2220233303231203"></a>

## Property reference — Property reference / 303101312133 / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-2222321211302222-3000301022222311-2112020111020303-0211321333200200-0330000200113123-1222031102030300-3031133300033023-1330233231022033)
- Property reference

<a id="canonical-3233210121231203-3220100011013030-0211103322102300-3101212011201032-3122001032311231-0202023320031120-2003223203100311-3112321111300001"></a>

## Direct properties — Property reference / 303101312133 / 3

<a id="canonical-0133010111223212-0232130301231303-3220223313332011-1122201211133320-3223132220303203-3330031313303311-1102111300213023-1122300133230312"></a>

<a id="canonical-3013032120120310-2332001022213021-1112330330232022-3300203320113220-1332230221230002-2011020323021211-3231121121120232-2302002211010103"></a>

## annotations property — Property reference / 303101312133 / 4

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

- [create_cloud_hosted](resources--bot_infrastructure--reference--group-001.md#canonical-2330310010102100-2003222100001321-3123220220113030-3220331102211303-0232000201002002-2321233202221132-3101030133100233-1003211122120022): complete subsection reference.

<a id="canonical-2000322133323332-3323321303130333-1311020113313210-0201230003102312-3323212221210222-1331220001033201-0212323023100223-0031232130023232"></a>

<a id="canonical-2330000022032331-0231303031230233-3320321301103310-0132111032312303-0132121011001012-0013203002033001-1310302300230000-1130230110302331"></a>

## description property — Property reference / 303101312133 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0320232210230312-2120210022203121-2022110203031210-0333111212122313-3033033300123033-2303203122233022-1210332313212123-2220222032010200"></a>

<a id="canonical-1201312133130101-2003112013301011-2113312221212123-2202003000303000-0212202132120222-3010330320303030-0001022033031110-1021221333330021"></a>

## disable property — Property reference / 303101312133 / 6

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

<a id="canonical-2012310122210120-2110201301021121-3121310332332201-3210213330300220-1211110113033030-1022132313001313-3230020220222031-0021312010012102"></a>

<a id="canonical-0301121233110300-0113111301200300-1231031231232223-1211132201332233-3113000301303123-0121311130030002-0022111222121020-1323030031330212"></a>

## ID property — Property reference / 303101312133 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3113102111211033-3100311313020131-0102230233322332-0211032031321210-3310323321220201-0102230223303110-2220101320212120-2220220223313230"></a>

<a id="canonical-1220131033031000-3021132303212223-3101313203132311-2311330000033133-2031210320233013-1033020313100321-2112231100220233-1222130002131032"></a>

## labels property — Property reference / 303101312133 / 8

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

<a id="canonical-1231321230000211-1231123100011202-0301012221323022-3010021102131203-0130222221300212-0311222333310212-1203020332022322-1122331122013111"></a>

<a id="canonical-1311202021330131-0322111302013123-2000300012022020-0221130022202131-0031333130133122-2023201000213221-3201220001222100-3030203301121011"></a>

## name property — Property reference / 303101312133 / 9

Type: `"string"`. Required.

Name of the Bot Infrastructure. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1301131220130123-0030011231220303-2312200330210232-0012030102301122-0302133010033201-2122032232312110-1213332010200111-1033302032100123"></a>

<a id="canonical-3203233230102300-1111303232122112-1303210021333230-2203333121333213-1112220332321030-2230332032001012-0002113230012320-1120121031203132"></a>

## namespace property — Property reference / 303101312133 / 10

Type: `"string"`. Required.

Namespace where the Bot Infrastructure is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
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

- [timeouts](resources--bot_infrastructure--reference--group-001.md#canonical-0222221022023101-0123223130031121-1111133013131322-3202210001321011-1321112231120010-3202002333113122-1300203222130212-1203232030232212): complete subsection reference.

<a id="canonical-2013333030221123-2000220033332223-3121130213011001-2300031013122333-0102020012133102-1013211300023202-0010003311031320-1033131230132002"></a>

<a id="canonical-2311102310200020-1300313310033233-1220203303303322-0133021202133111-3002313000021311-0023313222010101-2201211202100210-3300222022332203"></a>

## traffic_type property — Property reference / 303101312133 / 11

Type: `"string"`. Optional, Computed.

\[Enum: WEB|MOBILE\] The type of traffic that is routed to and processed by this infrastructure (Web
or Mobile). Only web traffic, including browser-based traffic from mobile devices, is routed through
this Bot Defense infrastructure. Only mobile traffic from native mobile apps with the Bot Defense
SDK are routed.. Possible values are \`WEB\`, \`MOBILE\`. Defaults to \`WEB\`.

Upstream description:

The type of traffic that is routed to and processed by this infrastructure (Web or Mobile).

Only web traffic, including browser-based traffic from mobile devices, is routed through this Bot
Defense infrastructure. Only mobile traffic from native mobile apps with the Bot Defense SDK are
routed through this Bot Defense infrastructure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("WEB",
    "MOBILE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "WEB",
  "enum": [
    "WEB",
    "MOBILE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0101020301301033-0022113100131111-1233122200133202-2012302300032223-0133100100012133-0303233000210331-3002101112020321-1312202210020001"></a>

## All schema paths — Property reference / 303101312133 / 12

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--bot_infrastructure--reference--group-001.md#canonical-0133010111223212-0232130301231303-3220223313332011-1122201211133320-3223132220303203-3330031313303311-1102111300213023-1122300133230312) |
| `create_cloud_hosted` | [create_cloud_hosted](resources--bot_infrastructure--reference--group-001.md#canonical-2012010121010323-3102133010030103-1132302121222311-1203133320012210-1103203101013323-2212331012110221-1113130103232202-0003023202011332) |
| `create_cloud_hosted.ip_addresses` | [create_cloud_hosted.ip_addresses](resources--bot_infrastructure--reference--group-001.md#canonical-0303313123133310-1020102223332311-3301111102221301-2203013102303120-0132121022121231-0003221310332032-2101311023011301-1231120003121230) |
| `create_cloud_hosted.production` | [create_cloud_hosted.production](resources--bot_infrastructure--reference--group-001.md#canonical-1131000011221310-1202220320323003-2033112232222122-2121211213303321-1301033000131203-0130321313233230-3323230233331000-3033210201132203) |
| `create_cloud_hosted.production.region_1` | [create_cloud_hosted.production.region_1](resources--bot_infrastructure--reference--group-001.md#canonical-0301311002210302-1221033132130111-1112230212213213-2220210212113221-2210101332122210-2131133102122303-0222223130110232-0110313010213232) |
| `create_cloud_hosted.production.region_2` | [create_cloud_hosted.production.region_2](resources--bot_infrastructure--reference--group-001.md#canonical-3003313031320313-3222211111110132-2121112201232010-2023230310212320-2222321313122100-0100300211213131-2032222211001032-1101321111302102) |
| `create_cloud_hosted.testing` | [create_cloud_hosted.testing](resources--bot_infrastructure--reference--group-001.md#canonical-1212021100232223-0000210211233111-0002222320322031-3200130303033223-1233131130033113-0013020122003300-1223133112131300-2031101123323320) |
| `create_cloud_hosted.testing.region_1` | [create_cloud_hosted.testing.region_1](resources--bot_infrastructure--reference--group-001.md#canonical-1202310003223132-0103023132323012-1331030010233301-2323303033333321-0322111323203030-2303311231301313-3023330201111333-0313011221320310) |
| `description` | [description](resources--bot_infrastructure--reference--group-001.md#canonical-2000322133323332-3323321303130333-1311020113313210-0201230003102312-3323212221210222-1331220001033201-0212323023100223-0031232130023232) |
| `disable` | [disable](resources--bot_infrastructure--reference--group-001.md#canonical-0320232210230312-2120210022203121-2022110203031210-0333111212122313-3033033300123033-2303203122233022-1210332313212123-2220222032010200) |
| `id` | [ID](resources--bot_infrastructure--reference--group-001.md#canonical-2012310122210120-2110201301021121-3121310332332201-3210213330300220-1211110113033030-1022132313001313-3230020220222031-0021312010012102) |
| `labels` | [labels](resources--bot_infrastructure--reference--group-001.md#canonical-3113102111211033-3100311313020131-0102230233322332-0211032031321210-3310323321220201-0102230223303110-2220101320212120-2220220223313230) |
| `name` | [name](resources--bot_infrastructure--reference--group-001.md#canonical-1231321230000211-1231123100011202-0301012221323022-3010021102131203-0130222221300212-0311222333310212-1203020332022322-1122331122013111) |
| `namespace` | [namespace](resources--bot_infrastructure--reference--group-001.md#canonical-1301131220130123-0030011231220303-2312200330210232-0012030102301122-0302133010033201-2122032232312110-1213332010200111-1033302032100123) |
| `timeouts` | [timeouts](resources--bot_infrastructure--reference--group-001.md#canonical-0020131030021331-2111202102132320-1311230001003002-0333112331300002-2121231121100221-2013120200233232-1200211120222031-1003020202021100) |
| `timeouts.create` | [timeouts.create](resources--bot_infrastructure--reference--group-001.md#canonical-3313020032312131-2120023331233032-3322332003231211-0020102330001023-0313320221121213-1001101333121131-3022332212102310-3031211333132022) |
| `timeouts.delete` | [timeouts.delete](resources--bot_infrastructure--reference--group-001.md#canonical-1001012020132103-0202103323302000-3020202023202322-1233222131101323-0223311331113321-0332010112233013-0322033013111203-1010101130213002) |
| `timeouts.read` | [timeouts.read](resources--bot_infrastructure--reference--group-001.md#canonical-3001022110312301-3320002022130031-0201112012103112-0022031331123030-0320200002012201-1213302112103220-0230000333022313-2101223101330313) |
| `timeouts.update` | [timeouts.update](resources--bot_infrastructure--reference--group-001.md#canonical-3100131102323021-1021220123122131-0103020031123220-3002131223320220-2031223010010113-0013120032032220-0001020130001331-1230032201330211) |
| `traffic_type` | [traffic_type](resources--bot_infrastructure--reference--group-001.md#canonical-2013333030221123-2000220033332223-3121130213011001-2300031013122333-0102020012133102-1013211300023202-0010003311031320-1033131230132002) |

<a id="canonical-1230100211321300-3023333010311033-1302122001102333-2133022000102302-0121103100120111-1120112013231200-2302030200030021-3101010331223213"></a>

## Next pages — Property reference / 303101312133 / 13

- [create_cloud_hosted](resources--bot_infrastructure--reference--group-001.md#canonical-2330310010102100-2003222100001321-3123220220113030-3220331102211303-0232000201002002-2321233202221132-3101030133100233-1003211122120022)
- [timeouts](resources--bot_infrastructure--reference--group-001.md#canonical-0222221022023101-0123223130031121-1111133013131322-3202210001321011-1321112231120010-3202002333113122-1300203222130212-1203232030232212)
- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-2222321211302222-3000301022222311-2112020111020303-0211321333200200-0330000200113123-1222031102030300-3031133300033023-1330233231022033)

<a id="canonical-2330310010102100-2003222100001321-3123220220113030-3220331102211303-0232000201002002-2321233202221132-3101030133100233-1003211122120022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200003232012011-2021232002211002-1221031313203122-2313300320320303-2011013113210300-1312122030333312-3320201100002022-0311313122031213"></a>

## create_cloud_hosted — create_cloud_hosted / 300230213302 / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-2222321211302222-3000301022222311-2112020111020303-0211321333200200-0330000200113123-1222031102030300-3031133300033023-1330233231022033)
- [Property reference](resources--bot_infrastructure--reference--group-001.md#canonical-3211120122133210-0122110022200232-1211012113212211-2130003230230030-1201022221000022-0012221313230312-0000102002302203-0201002321021101)
- create_cloud_hosted

<a id="canonical-2012010121010323-3102133010030103-1132302121222311-1203133320012210-1103203101013323-2212331012110221-1113130103232202-0003023202011332"></a>

Type: `"object"`. single nested block, Optional.

F5 Cloud Hosted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("production",
    "testing")}
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
  "x-ves-oneof-field-type_choice": "[\"production\",\"testing\"]"
}
```

Terraform syntax:

```terraform
create_cloud_hosted {
  # Configure direct properties listed below.
}
```

<a id="canonical-0110001320212031-2302003332210130-0231022012002102-0130321101223030-1230230000131230-0102203020002311-1200002303120130-2132113303210233"></a>

## Direct properties — create_cloud_hosted / 300230213302 / 3

<a id="canonical-0303313123133310-1020102223332311-3301111102221301-2203013102303120-0132121022121231-0003221310332032-2101311023011301-1231120003121230"></a>

<a id="canonical-2003013013213133-3311012312012032-1113011000321122-1002032323301023-2033113020330322-0102201330202100-1233113313323131-1202223020022131"></a>

## ip_addresses property — create_cloud_hosted / 300230213302 / 4

Type: `["list", "string"]`. Optional.

Only traffic from these IP addresses is allowed to access this Bot Defense infrastructure.

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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

- [production](resources--bot_infrastructure--reference--group-001.md#canonical-1331230011332113-1210000300323221-1201310220310031-0223300213320130-3133003023121202-1223223010120302-2302001333030011-1023203021313003): complete subsection reference.

- [testing](resources--bot_infrastructure--reference--group-001.md#canonical-3033321133303122-0331120203303201-2002100100022211-3211122031301311-3321333331333303-0210211002333012-3221003021312121-1220022220231213): complete subsection reference.

<a id="canonical-1021332233123201-0323110002310301-2310231113010111-3231120232333032-0232120323223013-0211211332312032-1013021322303033-3033200113321012"></a>

## Next pages — create_cloud_hosted / 300230213302 / 5

- [create_cloud_hosted.production](resources--bot_infrastructure--reference--group-001.md#canonical-1331230011332113-1210000300323221-1201310220310031-0223300213320130-3133003023121202-1223223010120302-2302001333030011-1023203021313003)
- [create_cloud_hosted.testing](resources--bot_infrastructure--reference--group-001.md#canonical-3033321133303122-0331120203303201-2002100100022211-3211122031301311-3321333331333303-0210211002333012-3221003021312121-1220022220231213)
- [Property reference](resources--bot_infrastructure--reference--group-001.md#canonical-3211120122133210-0122110022200232-1211012113212211-2130003230230030-1201022221000022-0012221313230312-0000102002302203-0201002321021101)
- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-2222321211302222-3000301022222311-2112020111020303-0211321333200200-0330000200113123-1222031102030300-3031133300033023-1330233231022033)

<a id="canonical-1331230011332113-1210000300323221-1201310220310031-0223300213320130-3133003023121202-1223223010120302-2302001333030011-1023203021313003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212111200332132-1330322102222033-2321133211123100-3101310313002330-1102001333101211-3310120321102221-3201132121221032-0321013331003232"></a>

## create_cloud_hosted.production — production / 302220202313 / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-2222321211302222-3000301022222311-2112020111020303-0211321333200200-0330000200113123-1222031102030300-3031133300033023-1330233231022033)
- [Property reference](resources--bot_infrastructure--reference--group-001.md#canonical-3211120122133210-0122110022200232-1211012113212211-2130003230230030-1201022221000022-0012221313230312-0000102002302203-0201002321021101)
- [create_cloud_hosted](resources--bot_infrastructure--reference--group-001.md#canonical-2330310010102100-2003222100001321-3123220220113030-3220331102211303-0232000201002002-2321233202221132-3101030133100233-1003211122120022)
- create_cloud_hosted.production

<a id="canonical-1131000011221310-1202220320323003-2033112232222122-2121211213303321-1301033000131203-0130321313233230-3323230233331000-3033210201132203"></a>

Type: `"object"`. single nested block, Optional.

Production.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("region_1",
    "region_2")}
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
production {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213021210030113-3021213010031113-3321011033200203-1112113332212223-3003303123230220-0133232111001232-2322302000000101-3221103201033220"></a>

## Direct properties — production / 302220202313 / 3

<a id="canonical-0301311002210302-1221033132130111-1112230212213213-2220210212113221-2210101332122210-2131133102122303-0222223130110232-0110313010213232"></a>

<a id="canonical-2000022331330310-3233030030200333-2113012111203222-1323032110233202-3130002233022121-0312213033013231-2022330202112001-0121311120233021"></a>

## region_1 property — production / 302220202313 / 4

Type: `"string"`. Optional.

Active-Active Infrastructure configuration where traffic is routed equally between the two regions.

Upstream description:

This is an Active-Active Infrastructure configuration where traffic is routed equally between the
two regions.

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

<a id="canonical-3003313031320313-3222211111110132-2121112201232010-2023230310212320-2222321313122100-0100300211213131-2032222211001032-1101321111302102"></a>

<a id="canonical-2110111021231212-1232302310223233-3020000213230300-1123031233332321-2302112311111032-2130130110013112-3120213322100121-1013201021200222"></a>

## region_2 property — production / 302220202313 / 5

Type: `"string"`. Optional.

Active-Active Infrastructure configuration where traffic is routed equally between the two regions.

Upstream description:

This is an Active-Active Infrastructure configuration where traffic is routed equally between the
two regions.

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

<a id="canonical-2320133130132202-1230012122133122-0211303300211302-2203013203033312-0111101212310112-0212131301021001-1333122102021111-2033111320032102"></a>

## Next pages — production / 302220202313 / 6

- [create_cloud_hosted](resources--bot_infrastructure--reference--group-001.md#canonical-2330310010102100-2003222100001321-3123220220113030-3220331102211303-0232000201002002-2321233202221132-3101030133100233-1003211122120022)
- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-2222321211302222-3000301022222311-2112020111020303-0211321333200200-0330000200113123-1222031102030300-3031133300033023-1330233231022033)

<a id="canonical-3033321133303122-0331120203303201-2002100100022211-3211122031301311-3321333331333303-0210211002333012-3221003021312121-1220022220231213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121130223202230-2120111230101112-1320010103130031-2203112131230111-0112102310212002-3030333300113322-0023110210110113-0133302021032202"></a>

## create_cloud_hosted.testing — testing / 020331223100 / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-2222321211302222-3000301022222311-2112020111020303-0211321333200200-0330000200113123-1222031102030300-3031133300033023-1330233231022033)
- [Property reference](resources--bot_infrastructure--reference--group-001.md#canonical-3211120122133210-0122110022200232-1211012113212211-2130003230230030-1201022221000022-0012221313230312-0000102002302203-0201002321021101)
- [create_cloud_hosted](resources--bot_infrastructure--reference--group-001.md#canonical-2330310010102100-2003222100001321-3123220220113030-3220331102211303-0232000201002002-2321233202221132-3101030133100233-1003211122120022)
- create_cloud_hosted.testing

<a id="canonical-1212021100232223-0000210211233111-0002222320322031-3200130303033223-1233131130033113-0013020122003300-1223133112131300-2031101123323320"></a>

Type: `"object"`. single nested block, Optional.

Testing

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("region_1")}
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
testing {
  # Configure direct properties listed below.
}
```

<a id="canonical-1023131211102213-0100121320102121-2332120002203301-0201232103210130-3210033031303301-1033321003020223-0212110202011222-3132012222023203"></a>

## Direct properties — testing / 020331223100 / 3

<a id="canonical-1202310003223132-0103023132323012-1331030010233301-2323303033333321-0322111323203030-2303311231301313-3023330201111333-0313011221320310"></a>

<a id="canonical-0023231313100113-1022311223021311-0130002112300023-3323331131131112-3210122313003101-0333201011031130-1312331033001220-2320022220101001"></a>

## region_1 property — testing / 020331223100 / 4

Type: `"string"`. Optional.

Active-Passive Infrastructure configuration where traffic is routed to a single region.

Upstream description:

This is an Active-Passive Infrastructure configuration where traffic is routed to a single region.

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

<a id="canonical-1221110232201020-2120203110200331-0220222101022022-1312301222313113-0031333212320031-1113310323110220-0222312133303122-2001233110331133"></a>

## Next pages — testing / 020331223100 / 5

- [create_cloud_hosted](resources--bot_infrastructure--reference--group-001.md#canonical-2330310010102100-2003222100001321-3123220220113030-3220331102211303-0232000201002002-2321233202221132-3101030133100233-1003211122120022)
- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-2222321211302222-3000301022222311-2112020111020303-0211321333200200-0330000200113123-1222031102030300-3031133300033023-1330233231022033)

<a id="canonical-0222221022023101-0123223130031121-1111133013131322-3202210001321011-1321112231120010-3202002333113122-1300203222130212-1203232030232212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020232110103200-2300001232000230-2222303101333113-3200231110111320-0021030132122133-0203300231022032-3220121111210021-2310111322030020"></a>

## timeouts — timeouts / 300201222030 / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-2222321211302222-3000301022222311-2112020111020303-0211321333200200-0330000200113123-1222031102030300-3031133300033023-1330233231022033)
- [Property reference](resources--bot_infrastructure--reference--group-001.md#canonical-3211120122133210-0122110022200232-1211012113212211-2130003230230030-1201022221000022-0012221313230312-0000102002302203-0201002321021101)
- timeouts

<a id="canonical-0020131030021331-2111202102132320-1311230001003002-0333112331300002-2121231121100221-2013120200233232-1200211120222031-1003020202021100"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2011033120210030-0222322332031231-0032030102001100-0121112233102333-3011221222322032-0003231121230231-2301331112333323-1131323333231323"></a>

## Direct properties — timeouts / 300201222030 / 3

<a id="canonical-3313020032312131-2120023331233032-3322332003231211-0020102330001023-0313320221121213-1001101333121131-3022332212102310-3031211333132022"></a>

<a id="canonical-0312310220301233-1003110331012020-1322333311111220-3023202221210122-1321211121202230-0203030032121110-0103021322120333-0030332013223233"></a>

## create property — timeouts / 300201222030 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1001012020132103-0202103323302000-3020202023202322-1233222131101323-0223311331113321-0332010112233013-0322033013111203-1010101130213002"></a>

<a id="canonical-0100201311230012-2011310121033313-3002010231212320-3202300310323222-0322003021211113-1201002023113203-1110222101230121-3133222123200303"></a>

## delete property — timeouts / 300201222030 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3001022110312301-3320002022130031-0201112012103112-0022031331123030-0320200002012201-1213302112103220-0230000333022313-2101223101330313"></a>

<a id="canonical-2301103203211031-1000031332131311-1102303202022201-3223030032320211-1330110112320001-0013212003300033-2021223312320210-2203030013000323"></a>

## read property — timeouts / 300201222030 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3100131102323021-1021220123122131-0103020031123220-3002131223320220-2031223010010113-0013120032032220-0001020130001331-1230032201330211"></a>

<a id="canonical-0200030002230331-2212102022222020-1333012213232233-2223330231322033-2010000032012011-3133102101033223-0200322022102312-3103210102130321"></a>

## update property — timeouts / 300201222030 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3112330333203313-3320030300031203-1221013322213213-2200232112120131-0303220211112211-3203121200333211-3323220023222213-1303011102230203"></a>

## Next pages — timeouts / 300201222030 / 8

- [Property reference](resources--bot_infrastructure--reference--group-001.md#canonical-3211120122133210-0122110022200232-1211012113212211-2130003230230030-1201022221000022-0012221313230312-0000102002302203-0201002321021101)
- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-2222321211302222-3000301022222311-2112020111020303-0211321333200200-0330000200113123-1222031102030300-3031133300033023-1330233231022033)
