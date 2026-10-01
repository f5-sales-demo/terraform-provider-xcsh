---
page_title: "xcsh_namespace reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_namespace reference."
---

# xcsh_namespace reference

<a id="canonical-0021301131013303-1333120321233220-3113022223110333-0100200232103001-3111031002202300-3231113102202033-3022100133111131-1310313032202230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301320230011311-3310331323302111-1120303211201122-2310112132203333-0320121122301113-3012110012333002-0011301110223012-2331120123222301"></a>

## Property reference — Property reference / 023201113022 / 2

Breadcrumbs:

- [xcsh_namespace](../resources/namespace.md#canonical-1131103113212120-1020211231300120-1030023112132023-3013023011120122-0313023322131212-0001202033301111-3031320130100032-1113113021202310)
- Property reference

<a id="canonical-1111131231132231-0100231302323003-2012111030122000-1323221300302201-0021101202231320-1220120001223303-3112021301022110-3122010333310202"></a>

## Direct properties — Property reference / 023201113022 / 3

<a id="canonical-2320302100010130-0021200323300002-0123332301003000-2010132101011222-0101323311330010-3210213100002313-1233201312223110-0013100333112321"></a>

<a id="canonical-2233300200320130-3103122321132123-3012323100133213-3331231222112132-2233320200130100-1300113111010012-0210103023301203-0321201012332231"></a>

## annotations property — Property reference / 023201113022 / 4

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

<a id="canonical-2210321211001310-3022002220220000-2032031012112212-3312120113230101-3212022302112223-1033020230221231-1220322213213023-0233300010200012"></a>

<a id="canonical-1210032202002303-0022102110300232-1231230310230132-1031220002210233-0330131200303020-3332322121310332-3132321320231223-3001101133310232"></a>

## description property — Property reference / 023201113022 / 5

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

<a id="canonical-2121300331230131-1021230212232120-1131131301211102-0012033213020221-0021001332133203-3201122332210012-2323033302113002-1110203233230003"></a>

<a id="canonical-1232211301000110-1212113000303320-1010213131032001-2012121321021302-1220212123130101-1233230332323232-2313321332223020-1121211101003121"></a>

## disable property — Property reference / 023201113022 / 6

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

<a id="canonical-0223122000313000-3223300133233013-0013312302012010-1010300011311001-1203311330101231-0331321031131000-0112333231313102-3021311133331103"></a>

<a id="canonical-0101211322033233-3302032202132131-2113123111002332-2121331300002311-2120113310233032-0133033000033112-0023313200112210-3201000112203310"></a>

## ID property — Property reference / 023201113022 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2322012103331323-3131102311311232-3122203121011123-2113311131002123-2122110111022311-1110030332232312-2301030203003120-1323021013113310"></a>

<a id="canonical-2312103312022303-2231210222113020-0310133200313131-1200202203013022-0313233322012223-0113101222111102-0320212023202022-2230013031313221"></a>

## labels property — Property reference / 023201113022 / 8

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

<a id="canonical-2033222012301121-0201333213221021-3320031021301130-0000112311131111-1012213123111200-0321210130221130-0000232131320310-3010020110232110"></a>

<a id="canonical-0311201230122221-0122212311323313-1111121201013022-2020321103322330-0121022112311123-1312122301023023-0010202221021212-2002231330233130"></a>

## name property — Property reference / 023201113022 / 9

Type: `"string"`. Required.

Name of the Namespace. Must be unique within the namespace.

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

<a id="canonical-0230323010213111-0230131101033321-0130230010300123-1322220110131310-3132311201332203-3301211030303302-2110311323331223-0221012121000003"></a>

<a id="canonical-0120131030230111-1113102221003232-2320130002001333-3311033010111130-1223313110203300-1301033300310330-1102201022132222-0232031120103121"></a>

## namespace property — Property reference / 023201113022 / 10

Type: `"string"`. Optional, Computed.

Namespace for the Namespace. For this resource type, namespace should be empty or omitted.

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

- [timeouts](resources--namespace--reference--group-001.md#canonical-0203233120222131-2232012210111320-3130211030311013-1023303031311323-3001220332231330-2231030022323030-2013223000003000-0222133020031200): complete subsection reference.

<a id="canonical-1111323122213120-3201101111000312-3312321013300223-1103002121010312-0022112122202232-2320323112222323-2210120132211312-1231012213203000"></a>

## All schema paths — Property reference / 023201113022 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--namespace--reference--group-001.md#canonical-2320302100010130-0021200323300002-0123332301003000-2010132101011222-0101323311330010-3210213100002313-1233201312223110-0013100333112321) |
| `description` | [description](resources--namespace--reference--group-001.md#canonical-2210321211001310-3022002220220000-2032031012112212-3312120113230101-3212022302112223-1033020230221231-1220322213213023-0233300010200012) |
| `disable` | [disable](resources--namespace--reference--group-001.md#canonical-2121300331230131-1021230212232120-1131131301211102-0012033213020221-0021001332133203-3201122332210012-2323033302113002-1110203233230003) |
| `id` | [id](resources--namespace--reference--group-001.md#canonical-0223122000313000-3223300133233013-0013312302012010-1010300011311001-1203311330101231-0331321031131000-0112333231313102-3021311133331103) |
| `labels` | [labels](resources--namespace--reference--group-001.md#canonical-2322012103331323-3131102311311232-3122203121011123-2113311131002123-2122110111022311-1110030332232312-2301030203003120-1323021013113310) |
| `name` | [name](resources--namespace--reference--group-001.md#canonical-2033222012301121-0201333213221021-3320031021301130-0000112311131111-1012213123111200-0321210130221130-0000232131320310-3010020110232110) |
| `namespace` | [namespace](resources--namespace--reference--group-001.md#canonical-0230323010213111-0230131101033321-0130230010300123-1322220110131310-3132311201332203-3301211030303302-2110311323331223-0221012121000003) |
| `timeouts` | [timeouts](resources--namespace--reference--group-001.md#canonical-1323302002031023-1003022330332212-2230001220011201-0313033012122132-0311012313213022-2301200020021010-3230023131121212-2330303110021210) |
| `timeouts.create` | [timeouts.create](resources--namespace--reference--group-001.md#canonical-2210301132301001-0123232333132302-1233033010110320-0102320130101232-3122303103013022-0203202132201112-0120323310302221-0203102213031301) |
| `timeouts.delete` | [timeouts.delete](resources--namespace--reference--group-001.md#canonical-2221013132202232-0220112031220132-2002002203202213-2113310202011122-3223021112101122-2101223303312313-1100100301331233-3230002011131003) |
| `timeouts.read` | [timeouts.read](resources--namespace--reference--group-001.md#canonical-0302012332220201-2211200021331123-3022030010010103-0100232311230012-3300130312120001-0210331312000012-3312113101003032-3310212301103233) |
| `timeouts.update` | [timeouts.update](resources--namespace--reference--group-001.md#canonical-3323100332100230-0032303203013300-1230203222121103-0310302323020002-3102220112303233-0011202000112323-2121220300121311-1023112320333013) |

<a id="canonical-1312122222220322-1020210021020000-1002110210200032-2133231023013132-0112223101012022-0022010003131000-0003313220313233-1231232301021230"></a>

## Next pages — Property reference / 023201113022 / 12

- [timeouts](resources--namespace--reference--group-001.md#canonical-0203233120222131-2232012210111320-3130211030311013-1023303031311323-3001220332231330-2231030022323030-2013223000003000-0222133020031200)
- [xcsh_namespace](../resources/namespace.md#canonical-1131103113212120-1020211231300120-1030023112132023-3013023011120122-0313023322131212-0001202033301111-3031320130100032-1113113021202310)

<a id="canonical-0203233120222131-2232012210111320-3130211030311013-1023303031311323-3001220332231330-2231030022323030-2013223000003000-0222133020031200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001230102120311-0011202202302203-0111133333101110-0331120211130113-1300323132031211-0010303012313001-0201122202022302-2120123021001202"></a>

## timeouts — timeouts / 331033230122 / 2

Breadcrumbs:

- [xcsh_namespace](../resources/namespace.md#canonical-1131103113212120-1020211231300120-1030023112132023-3013023011120122-0313023322131212-0001202033301111-3031320130100032-1113113021202310)
- [Property reference](resources--namespace--reference--group-001.md#canonical-0021301131013303-1333120321233220-3113022223110333-0100200232103001-3111031002202300-3231113102202033-3022100133111131-1310313032202230)
- timeouts

<a id="canonical-1323302002031023-1003022330332212-2230001220011201-0313033012122132-0311012313213022-2301200020021010-3230023131121212-2330303110021210"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003122200122330-3011023321330320-2322112122312313-1311213231130212-0300321032033130-1122001202031320-0312331011001001-2311223010321022"></a>

## Direct properties — timeouts / 331033230122 / 3

<a id="canonical-2210301132301001-0123232333132302-1233033010110320-0102320130101232-3122303103013022-0203202132201112-0120323310302221-0203102213031301"></a>

<a id="canonical-2030312313320220-3210111112211101-0023001230003222-3233133023101232-2103313031133221-3231220001311322-0132020111301300-0123010323003320"></a>

## create property — timeouts / 331033230122 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2221013132202232-0220112031220132-2002002203202213-2113310202011122-3223021112101122-2101223303312313-1100100301331233-3230002011131003"></a>

<a id="canonical-2013103030233132-3202022033333130-3302301331302100-0201121333133220-3120000323303101-1213001110131202-1332312031123123-0333232303133230"></a>

## delete property — timeouts / 331033230122 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0302012332220201-2211200021331123-3022030010010103-0100232311230012-3300130312120001-0210331312000012-3312113101003032-3310212301103233"></a>

<a id="canonical-2323032001210130-1203103311001022-1331221200131133-2031231132132233-1300230112000110-2201022232330122-2330121003332103-2000130012131213"></a>

## read property — timeouts / 331033230122 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3323100332100230-0032303203013300-1230203222121103-0310302323020002-3102220112303233-0011202000112323-2121220300121311-1023112320333013"></a>

<a id="canonical-0120311310331310-3313033112122322-2212330133201132-1330121130003310-1032130100021023-1120301301022322-0202201303012311-3301133211201021"></a>

## update property — timeouts / 331033230122 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2231220101132233-1013133222011003-1023012120023103-0330223323303300-2323121313301303-2022221132120211-3112333031000033-3103133332001320"></a>

## Next pages — timeouts / 331033230122 / 8

- [Property reference](resources--namespace--reference--group-001.md#canonical-0021301131013303-1333120321233220-3113022223110333-0100200232103001-3111031002202300-3231113102202033-3022100133111131-1310313032202230)
- [xcsh_namespace](../resources/namespace.md#canonical-1131103113212120-1020211231300120-1030023112132023-3013023011120122-0313023322131212-0001202033301111-3031320130100032-1113113021202310)
