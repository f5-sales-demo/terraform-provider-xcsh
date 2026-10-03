---
page_title: "xcsh_ike1 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike1 reference."
---

# xcsh_ike1 reference

<a id="canonical-2113121021131213-0322330121000333-0232103301210212-3101112300210032-3102213101322310-0311031112112121-1311322002110322-0313010312330013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203330223211210-3020202230201311-1122130112112003-3031300210113221-2102222130021111-0120211011230313-0233131133013111-1232003223201012"></a>

## Property reference — Property reference / 112111320022 / 2

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md#canonical-2112011122002111-1121301111312122-3021231213103111-3131302213110331-0301300202011011-1200320222013102-3321012020321333-1000030102103313)
- Property reference

<a id="canonical-2231221102331111-2232201213113233-1000033221233022-1233102202310202-2103231031113131-1302000230013020-0231220121002311-2212312010232311"></a>

## Direct properties — Property reference / 112111320022 / 3

<a id="canonical-3202333202222222-2011010011212103-1213321321213323-0313202202002010-3333321003131322-3131121221121223-2122203313033012-2120010210130303"></a>

<a id="canonical-0031213323321003-0010131301223002-2112103012300222-0003031010121311-3122132123231002-3302010030303022-2111012113312311-1323303210111302"></a>

## annotations property — Property reference / 112111320022 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

<a id="canonical-3210012303220033-2303320330032111-1202213010011120-3033101010321021-0100102122021112-1001111000002210-0313103300332122-2220130303200200"></a>

<a id="canonical-2312220030130222-0031331232301213-0013002223012130-3223310113332110-0033313213103303-2213332313133232-2321032112300033-1103302113113231"></a>

## description property — Property reference / 112111320022 / 5

Type: `"string"`. Computed.

Description of the Ike1.

Upstream description:

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

<a id="canonical-0032132220103220-0331020100331023-3112111100011130-1300000320222033-3333203200203021-0121323112132211-2303010103203202-2132132231210103"></a>

<a id="canonical-0333330130332121-0311300011202311-1222101311131110-2331122301030200-1020211322003032-3321023210030311-2031311300331001-2003311310320121"></a>

## ID property — Property reference / 112111320022 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](data-sources--ike1--reference--group-001.md#canonical-0120031312130031-1132223331220200-2133100030100121-0103222320230321-2330021313322301-1133110020212231-1210023022223030-0330011111032200): complete subsection reference.

- [ike_keylifetime_minutes](data-sources--ike1--reference--group-001.md#canonical-1232302301313123-3223222122302102-1230332113313122-1100330332201202-0220101233231221-0310100100101201-0122221302202320-2112001111213211): complete subsection reference.

<a id="canonical-0012013233310111-0210031111131313-1023202120122220-2201030202222200-2102122332311211-2201010233233210-2113303132323031-2310321210202021"></a>

<a id="canonical-2030011230321212-1322330300333221-1312122303302132-0120120133301030-1322103121030131-2101312101302120-3122302030333123-0110223032311202"></a>

## labels property — Property reference / 112111320022 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-1033002301123310-2303301102203303-0021121012322121-3002230322111002-1011201101101010-2133300123012332-3113301203122311-3021001202313031"></a>

<a id="canonical-0220231111233231-3021320011203122-2200331123311023-1130302022102302-3012121301033313-0101131100330332-0202330210220112-3121303333222331"></a>

## name property — Property reference / 112111320022 / 8

Type: `"string"`. Required.

Name of the Ike1.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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

<a id="canonical-1101032320232001-0230112100202211-0311120300002310-3032110011103212-0011321310033100-3002132322133311-3030010003103211-0300010310111121"></a>

<a id="canonical-2122012002121030-1303120202032012-0330310033101321-3221023130120021-1210221231101023-1203313123133230-3013201020213022-2013331201320330"></a>

## namespace property — Property reference / 112111320022 / 9

Type: `"string"`. Required.

Namespace where the Ike1 exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

- [reauth_disabled](data-sources--ike1--reference--group-001.md#canonical-3033312231301021-1312332110313310-0022003320203320-3311113320322230-2303310002023031-0303231333300103-2003132232103312-3130230321003312): complete subsection reference.

- [reauth_timeout_days](data-sources--ike1--reference--group-001.md#canonical-2133111320003023-0122010212102113-2210003202133112-0212222121011111-1232101211020212-3120030011222322-1002232301311221-2221301000112022): complete subsection reference.

- [reauth_timeout_hours](data-sources--ike1--reference--group-001.md#canonical-0102013100100030-1300321120331131-2201302003323321-1202312123020013-2030320331100303-3113331330013120-3010003022023213-1032000312002110): complete subsection reference.

- [use_default_keylifetime](data-sources--ike1--reference--group-001.md#canonical-0002111212321313-2201102133212333-2133110002001112-1330111101010232-0003303210302032-2233112300001020-0221021111213102-2211102133202210): complete subsection reference.

<a id="canonical-0130312003230103-2333223221130313-2032132201333221-2101133001300020-3321112310101022-3100210130121010-1013233201002330-2223122122332022"></a>

## All schema paths — Property reference / 112111320022 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--ike1--reference--group-001.md#canonical-3202333202222222-2011010011212103-1213321321213323-0313202202002010-3333321003131322-3131121221121223-2122203313033012-2120010210130303) |
| `description` | [description](data-sources--ike1--reference--group-001.md#canonical-3210012303220033-2303320330032111-1202213010011120-3033101010321021-0100102122021112-1001111000002210-0313103300332122-2220130303200200) |
| `id` | [ID](data-sources--ike1--reference--group-001.md#canonical-0032132220103220-0331020100331023-3112111100011130-1300000320222033-3333203200203021-0121323112132211-2303010103203202-2132132231210103) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](data-sources--ike1--reference--group-001.md#canonical-1302030222022312-1030000313121311-1032200023302330-3303012131223313-2023010223332012-3010110122120022-0332132113301131-2210231021020020) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](data-sources--ike1--reference--group-001.md#canonical-2001112012303301-3110211232331112-2012220330102120-3333133011100013-2213130300320213-0203222111100302-1023221221200111-1001311133221021) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](data-sources--ike1--reference--group-001.md#canonical-1330211202122302-1131232202331332-2132222030313111-2311130230330112-3302031103302223-0132111301120002-0121000212103203-1210302331212113) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](data-sources--ike1--reference--group-001.md#canonical-1210332112002230-3132333100003112-3302113102230331-3003133001010132-2110032121301322-0133222323322102-3133221310301330-1210320112130233) |
| `labels` | [labels](data-sources--ike1--reference--group-001.md#canonical-0012013233310111-0210031111131313-1023202120122220-2201030202222200-2102122332311211-2201010233233210-2113303132323031-2310321210202021) |
| `name` | [name](data-sources--ike1--reference--group-001.md#canonical-1033002301123310-2303301102203303-0021121012322121-3002230322111002-1011201101101010-2133300123012332-3113301203122311-3021001202313031) |
| `namespace` | [namespace](data-sources--ike1--reference--group-001.md#canonical-1101032320232001-0230112100202211-0311120300002310-3032110011103212-0011321310033100-3002132322133311-3030010003103211-0300010310111121) |
| `reauth_disabled` | [reauth_disabled](data-sources--ike1--reference--group-001.md#canonical-3322222212322310-1011202013332331-0312001223031111-3023011213132120-3330223132002013-0301213321121032-2211230332132202-3202210112203320) |
| `reauth_timeout_days` | [reauth_timeout_days](data-sources--ike1--reference--group-001.md#canonical-2221131223121121-0130013100130301-0322300101231003-1333130323231100-3123013010131132-1102311201003333-1320310313223200-1012221123110300) |
| `reauth_timeout_days.duration` | [reauth_timeout_days.duration](data-sources--ike1--reference--group-001.md#canonical-2100102000202101-0010033222213021-1221122103101003-0323313222133033-1330103100231001-0003200023312111-1031101212322120-3101020100130313) |
| `reauth_timeout_hours` | [reauth_timeout_hours](data-sources--ike1--reference--group-001.md#canonical-3103000230322220-2200332222101233-2321220203232312-3212330210202322-2021222101212212-0213203300011122-3210233223013232-0130223011320230) |
| `reauth_timeout_hours.duration` | [reauth_timeout_hours.duration](data-sources--ike1--reference--group-001.md#canonical-1302313211321203-2230023120122332-1203233330031021-2333003012001202-3203030031321320-2123223333100132-3331313113133121-0330200111011211) |
| `use_default_keylifetime` | [use_default_keylifetime](data-sources--ike1--reference--group-001.md#canonical-0022313110121122-3112031003021102-1210201220101011-0220020101301211-3203121213033002-3213002331233001-0323121102333013-0133211201223111) |

<a id="canonical-2220011010020233-2001310200022101-3002222310032303-1301100020212332-3102320000301013-2110010110022131-1223322211330122-1320003332200031"></a>

## Next pages — Property reference / 112111320022 / 11

- [ike_keylifetime_hours](data-sources--ike1--reference--group-001.md#canonical-0120031312130031-1132223331220200-2133100030100121-0103222320230321-2330021313322301-1133110020212231-1210023022223030-0330011111032200)
- [ike_keylifetime_minutes](data-sources--ike1--reference--group-001.md#canonical-1232302301313123-3223222122302102-1230332113313122-1100330332201202-0220101233231221-0310100100101201-0122221302202320-2112001111213211)
- [reauth_disabled](data-sources--ike1--reference--group-001.md#canonical-3033312231301021-1312332110313310-0022003320203320-3311113320322230-2303310002023031-0303231333300103-2003132232103312-3130230321003312)
- [reauth_timeout_days](data-sources--ike1--reference--group-001.md#canonical-2133111320003023-0122010212102113-2210003202133112-0212222121011111-1232101211020212-3120030011222322-1002232301311221-2221301000112022)
- [reauth_timeout_hours](data-sources--ike1--reference--group-001.md#canonical-0102013100100030-1300321120331131-2201302003323321-1202312123020013-2030320331100303-3113331330013120-3010003022023213-1032000312002110)
- [use_default_keylifetime](data-sources--ike1--reference--group-001.md#canonical-0002111212321313-2201102133212333-2133110002001112-1330111101010232-0003303210302032-2233112300001020-0221021111213102-2211102133202210)
- [xcsh_ike1](../data-sources/ike1.md#canonical-2112011122002111-1121301111312122-3021231213103111-3131302213110331-0301300202011011-1200320222013102-3321012020321333-1000030102103313)

<a id="canonical-0120031312130031-1132223331220200-2133100030100121-0103222320230321-2330021313322301-1133110020212231-1210023022223030-0330011111032200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020300331301130-1100201113031123-1030213321113312-0221112232110301-1303311133110121-1332200323310300-2001202201020301-1002212333120320"></a>

## ike_keylifetime_hours — ike_keylifetime_hours / 222232312212 / 2

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md#canonical-2112011122002111-1121301111312122-3021231213103111-3131302213110331-0301300202011011-1200320222013102-3321012020321333-1000030102103313)
- [Property reference](data-sources--ike1--reference--group-001.md#canonical-2113121021131213-0322330121000333-0232103301210212-3101112300210032-3102213101322310-0311031112112121-1311322002110322-0313010312330013)
- ike_keylifetime_hours

<a id="canonical-1302030222022312-1030000313121311-1032200023302330-3303012131223313-2023010223332012-3010110122120022-0332132113301131-2210231021020020"></a>

Type: `"single"`. Computed.

\[OneOf: ike\_keylifetime\_hours, ike\_keylifetime\_minutes, use\_default\_keylifetime; Default:
use\_default\_keylifetime\] Configuration parameter for ike keylifetime hours.

Upstream description:

Input Hours.

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

- [ike_keylifetime_hours](data-sources--ike1--reference--group-001.md#canonical-1302030222022312-1030000313121311-1032200023302330-3303012131223313-2023010223332012-3010110122120022-0332132113301131-2210231021020020)
- [ike_keylifetime_minutes](data-sources--ike1--reference--group-001.md#canonical-1330211202122302-1131232202331332-2132222030313111-2311130230330112-3302031103302223-0132111301120002-0121000212103203-1210302331212113)
- [use_default_keylifetime](data-sources--ike1--reference--group-001.md#canonical-0022313110121122-3112031003021102-1210201220101011-0220020101301211-3203121213033002-3213002331233001-0323121102333013-0133211201223111)

Select alternatives according to the provider validators above.

<a id="canonical-2002022301033023-2121233111201210-1003232332330323-1320223020122111-0312102311102112-3000200101311002-2222130001122221-1312000120120321"></a>

## Direct properties — ike_keylifetime_hours / 222232312212 / 3

<a id="canonical-2001112012303301-3110211232331112-2012220330102120-3333133011100013-2213130300320213-0203222111100302-1023221221200111-1001311133221021"></a>

<a id="canonical-0000001202213331-0031322100111212-2310311330003221-3320020020022033-2100112131313330-0301322133220013-1322000333321132-1210301113032312"></a>

## duration property — ike_keylifetime_hours / 222232312212 / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

<a id="canonical-3212330021231133-0210301030332102-3112001122300311-0123220013103301-1132113303033202-1231123232323311-1322210111122102-1021010222103330"></a>

## Next pages — ike_keylifetime_hours / 222232312212 / 5

- [Property reference](data-sources--ike1--reference--group-001.md#canonical-2113121021131213-0322330121000333-0232103301210212-3101112300210032-3102213101322310-0311031112112121-1311322002110322-0313010312330013)
- [xcsh_ike1](../data-sources/ike1.md#canonical-2112011122002111-1121301111312122-3021231213103111-3131302213110331-0301300202011011-1200320222013102-3321012020321333-1000030102103313)

<a id="canonical-1232302301313123-3223222122302102-1230332113313122-1100330332201202-0220101233231221-0310100100101201-0122221302202320-2112001111213211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310233312002130-3002220230212022-3032213132321011-0303113100322130-2110232003232322-1121102321333330-1031333213312131-0020013221331300"></a>

## ike_keylifetime_minutes — ike_keylifetime_minutes / 101110300101 / 2

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md#canonical-2112011122002111-1121301111312122-3021231213103111-3131302213110331-0301300202011011-1200320222013102-3321012020321333-1000030102103313)
- [Property reference](data-sources--ike1--reference--group-001.md#canonical-2113121021131213-0322330121000333-0232103301210212-3101112300210032-3102213101322310-0311031112112121-1311322002110322-0313010312330013)
- ike_keylifetime_minutes

<a id="canonical-1330211202122302-1131232202331332-2132222030313111-2311130230330112-3302031103302223-0132111301120002-0121000212103203-1210302331212113"></a>

Type: `"single"`. Computed.

Configuration parameter for ike keylifetime minutes.

Upstream description:

Set IKE Key Lifetime in minutes.

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

<a id="canonical-2102110130202103-1120222001200013-1200103112223203-2311222010320101-1311310102122101-2011123333100021-1003020103213232-3021023003330200"></a>

## Direct properties — ike_keylifetime_minutes / 101110300101 / 3

<a id="canonical-1210332112002230-3132333100003112-3302113102230331-3003133001010132-2110032121301322-0133222323322102-3133221310301330-1210320112130233"></a>

<a id="canonical-1211232202321021-2001322013323321-3332320011233321-1310100303231332-2111231100031332-3133012000213110-2222133012110310-1233310111211323"></a>

## duration property — ike_keylifetime_minutes / 101110300101 / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 10
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-2021121112010112-1103223112020200-2002000120000221-1100000000332001-1001002100210213-1220221322133113-3303231012000202-3300302201000221"></a>

## Next pages — ike_keylifetime_minutes / 101110300101 / 5

- [Property reference](data-sources--ike1--reference--group-001.md#canonical-2113121021131213-0322330121000333-0232103301210212-3101112300210032-3102213101322310-0311031112112121-1311322002110322-0313010312330013)
- [xcsh_ike1](../data-sources/ike1.md#canonical-2112011122002111-1121301111312122-3021231213103111-3131302213110331-0301300202011011-1200320222013102-3321012020321333-1000030102103313)

<a id="canonical-3033312231301021-1312332110313310-0022003320203320-3311113320322230-2303310002023031-0303231333300103-2003132232103312-3130230321003312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320200303301012-0222203020100232-3202023202020131-2121331122130212-3221032302132003-1201032133031331-2202333203101200-3003301133232020"></a>

## reauth_disabled — reauth_disabled / 103021102012 / 2

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md#canonical-2112011122002111-1121301111312122-3021231213103111-3131302213110331-0301300202011011-1200320222013102-3321012020321333-1000030102103313)
- [Property reference](data-sources--ike1--reference--group-001.md#canonical-2113121021131213-0322330121000333-0232103301210212-3101112300210032-3102213101322310-0311031112112121-1311322002110322-0313010312330013)
- reauth_disabled

<a id="canonical-3322222212322310-1011202013332331-0312001223031111-3023011213132120-3330223132002013-0301213321121032-2211230332132202-3202210112203320"></a>

Type: `["object", {}]`. Computed.

\[OneOf: reauth\_disabled, reauth\_timeout\_days, reauth\_timeout\_hours\] Enable this option

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

- [reauth_disabled](data-sources--ike1--reference--group-001.md#canonical-3322222212322310-1011202013332331-0312001223031111-3023011213132120-3330223132002013-0301213321121032-2211230332132202-3202210112203320)
- [reauth_timeout_days](data-sources--ike1--reference--group-001.md#canonical-2221131223121121-0130013100130301-0322300101231003-1333130323231100-3123013010131132-1102311201003333-1320310313223200-1012221123110300)
- [reauth_timeout_hours](data-sources--ike1--reference--group-001.md#canonical-3103000230322220-2200332222101233-2321220203232312-3212330210202322-2021222101212212-0213203300011122-3210233223013232-0130223011320230)

Select alternatives according to the provider validators above.

<a id="canonical-0310012330321323-1310332320032203-2333111323222302-3322200112013100-1032132111303233-0033003323132102-0323233132110123-3230213301131230"></a>

## Direct properties — reauth_disabled / 103021102012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303013121101133-2321312102213212-3203032311003101-1001113012033310-1203322320102300-3300333132211030-3100222131310031-1301112332220110"></a>

## Next pages — reauth_disabled / 103021102012 / 4

- [Property reference](data-sources--ike1--reference--group-001.md#canonical-2113121021131213-0322330121000333-0232103301210212-3101112300210032-3102213101322310-0311031112112121-1311322002110322-0313010312330013)
- [xcsh_ike1](../data-sources/ike1.md#canonical-2112011122002111-1121301111312122-3021231213103111-3131302213110331-0301300202011011-1200320222013102-3321012020321333-1000030102103313)

<a id="canonical-2133111320003023-0122010212102113-2210003202133112-0212222121011111-1232101211020212-3120030011222322-1002232301311221-2221301000112022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003120313132013-0131130332102313-2023033313203113-0301030330230313-2012121313002102-3021101212202123-0102223011030120-1230223022133130"></a>

## reauth_timeout_days — reauth_timeout_days / 302321333102 / 2

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md#canonical-2112011122002111-1121301111312122-3021231213103111-3131302213110331-0301300202011011-1200320222013102-3321012020321333-1000030102103313)
- [Property reference](data-sources--ike1--reference--group-001.md#canonical-2113121021131213-0322330121000333-0232103301210212-3101112300210032-3102213101322310-0311031112112121-1311322002110322-0313010312330013)
- reauth_timeout_days

<a id="canonical-2221131223121121-0130013100130301-0322300101231003-1333130323231100-3123013010131132-1102311201003333-1320310313223200-1012221123110300"></a>

Type: `"single"`. Computed.

Configuration parameter for reauth timeout days.

Upstream description:

Set Duration in days.

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

<a id="canonical-1202313131011323-3212331121001310-1100000102212120-2123100022222012-1033221101202233-0032303323023120-1303330131130021-0313301010101222"></a>

## Direct properties — reauth_timeout_days / 302321333102 / 3

<a id="canonical-2100102000202101-0010033222213021-1221122103101003-0323313222133033-1330103100231001-0003200023312111-1031101212322120-3101020100130313"></a>

<a id="canonical-3012101002030032-0032302013203120-0000331223131330-1011002311300300-2323310030233211-1230333132232201-0300003112001102-1131101210100021"></a>

## duration property — reauth_timeout_days / 302321333102 / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-1231030131332012-2300120023002332-1210322020230201-1321231303023100-3211333322232200-0101222001020210-3000230122012012-0130223100302002"></a>

## Next pages — reauth_timeout_days / 302321333102 / 5

- [Property reference](data-sources--ike1--reference--group-001.md#canonical-2113121021131213-0322330121000333-0232103301210212-3101112300210032-3102213101322310-0311031112112121-1311322002110322-0313010312330013)
- [xcsh_ike1](../data-sources/ike1.md#canonical-2112011122002111-1121301111312122-3021231213103111-3131302213110331-0301300202011011-1200320222013102-3321012020321333-1000030102103313)

<a id="canonical-0102013100100030-1300321120331131-2201302003323321-1202312123020013-2030320331100303-3113331330013120-3010003022023213-1032000312002110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223320213202201-2013223212301301-3212022010210131-2232321031003101-1123123033213111-1211020221220003-3212021013221322-3333023103202011"></a>

## reauth_timeout_hours — reauth_timeout_hours / 122020303010 / 2

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md#canonical-2112011122002111-1121301111312122-3021231213103111-3131302213110331-0301300202011011-1200320222013102-3321012020321333-1000030102103313)
- [Property reference](data-sources--ike1--reference--group-001.md#canonical-2113121021131213-0322330121000333-0232103301210212-3101112300210032-3102213101322310-0311031112112121-1311322002110322-0313010312330013)
- reauth_timeout_hours

<a id="canonical-3103000230322220-2200332222101233-2321220203232312-3212330210202322-2021222101212212-0213203300011122-3210233223013232-0130223011320230"></a>

Type: `"single"`. Computed.

Configuration parameter for reauth timeout hours.

Upstream description:

Input Hours.

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

<a id="canonical-1211120322302232-3231032333200301-0121011232011112-2312101300312203-3310321032313311-2211232102002333-0120222230022210-2131210010102121"></a>

## Direct properties — reauth_timeout_hours / 122020303010 / 3

<a id="canonical-1302313211321203-2230023120122332-1203233330031021-2333003012001202-3203030031321320-2123223333100132-3331313113133121-0330200111011211"></a>

<a id="canonical-1311111300130030-0120121332131133-1001013330322033-0222003120332303-3132322201230103-3103121113232333-1121122201331313-3210100303012202"></a>

## duration property — reauth_timeout_hours / 122020303010 / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

<a id="canonical-1030023330322211-0310110333033020-0022321222322323-3001302330322121-3131323103012320-2110131111123323-1002311022103021-3212320120032023"></a>

## Next pages — reauth_timeout_hours / 122020303010 / 5

- [Property reference](data-sources--ike1--reference--group-001.md#canonical-2113121021131213-0322330121000333-0232103301210212-3101112300210032-3102213101322310-0311031112112121-1311322002110322-0313010312330013)
- [xcsh_ike1](../data-sources/ike1.md#canonical-2112011122002111-1121301111312122-3021231213103111-3131302213110331-0301300202011011-1200320222013102-3321012020321333-1000030102103313)

<a id="canonical-0002111212321313-2201102133212333-2133110002001112-1330111101010232-0003303210302032-2233112300001020-0221021111213102-2211102133202210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122200222103032-3023230222101230-1221312112003300-0230133220302322-2132221030322120-2112022112022120-0320210000013031-0202203131120120"></a>

## use_default_keylifetime — use_default_keylifetime / 103200000202 / 2

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md#canonical-2112011122002111-1121301111312122-3021231213103111-3131302213110331-0301300202011011-1200320222013102-3321012020321333-1000030102103313)
- [Property reference](data-sources--ike1--reference--group-001.md#canonical-2113121021131213-0322330121000333-0232103301210212-3101112300210032-3102213101322310-0311031112112121-1311322002110322-0313010312330013)
- use_default_keylifetime

<a id="canonical-0022313110121122-3112031003021102-1210201220101011-0220020101301211-3203121213033002-3213002331233001-0323121102333013-0133211201223111"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use default keylifetime.

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

<a id="canonical-3203310111100113-3111203322011111-0311303133103121-0020101220330013-2320110020002211-0113122223230110-0321203030013012-2213220211201113"></a>

## Direct properties — use_default_keylifetime / 103200000202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210310102330310-2012003012303010-3023200001123122-2332231300123202-0301220121303121-1032011032030022-3223203131303202-3303130003213100"></a>

## Next pages — use_default_keylifetime / 103200000202 / 4

- [Property reference](data-sources--ike1--reference--group-001.md#canonical-2113121021131213-0322330121000333-0232103301210212-3101112300210032-3102213101322310-0311031112112121-1311322002110322-0313010312330013)
- [xcsh_ike1](../data-sources/ike1.md#canonical-2112011122002111-1121301111312122-3021231213103111-3131302213110331-0301300202011011-1200320222013102-3321012020321333-1000030102103313)
