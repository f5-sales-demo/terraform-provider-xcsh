---
page_title: "xcsh_ike1 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike1 reference."
---

# xcsh_ike1 reference

<a id="canonical-0231322323021312-2223311002332013-3200002221122100-2200120310330232-3230312122213012-1111030312321102-2331002032223020-1222113231322110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111122100202313-2110131100002132-0002011131113101-3110132021010100-0222333203031302-3212011000120213-1203113321011000-2233320122101212"></a>

## Property reference — Property reference / 012302030031 / 2

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130)
- Property reference

<a id="canonical-0210113233323322-1330102130222123-3022332310313120-2301323231121110-2111023302131032-2232011223013213-2001330031212330-0330101203212122"></a>

## Direct properties — Property reference / 012302030031 / 3

<a id="canonical-2100302200012333-0211302303131023-1020023212232231-1220310202103000-0201030010200210-2132303030311310-0200212023023333-2021333130313220"></a>

<a id="canonical-0030210302210322-3300312010112302-0320203233331023-2203212210302112-3110222203301301-1012002231032213-3232103001311210-2032311131002022"></a>

## annotations property — Property reference / 012302030031 / 4

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

<a id="canonical-3021202201031301-3301012331303230-1110331102330321-2020222012131110-0301113000111100-2222023231112123-0332001220003033-1021021302131130"></a>

<a id="canonical-1330111012013333-1313103202332021-3302102302132201-1322000212223022-1002021033023102-2322103210022031-2020130223012111-2222232130330320"></a>

## description property — Property reference / 012302030031 / 5

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

<a id="canonical-0321001232301021-2222031023230110-3301333212100023-3111223130330111-0133002033222230-0210312132120303-2220112133030322-2112110332103013"></a>

<a id="canonical-3233233121101131-0232313001333031-0331301133002332-3310102312103223-2122320313201002-0332130010301230-1202213123130022-0311231230123331"></a>

## disable property — Property reference / 012302030031 / 6

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

<a id="canonical-2300013102232113-2211230222211231-2123130330131120-0322233113023201-3120001103103112-1111010002200122-0322030132103023-0003211312302231"></a>

<a id="canonical-2220301010302212-0211303231313002-0233230301310121-1002320033320330-2110000323103010-0133233002320330-2220102323313332-3331201331220012"></a>

## ID property — Property reference / 012302030031 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](resources--ike1--reference--group-001.md#canonical-1331103023310210-3303300122110230-3303110012211012-3210310233133331-3031130111311300-2003010332103002-0331110001201232-1202020111313323): complete subsection reference.

- [ike_keylifetime_minutes](resources--ike1--reference--group-001.md#canonical-1331031321203231-2213012032201103-0033033320110332-2023021001032213-0300100220313231-2130011101032101-0002033320121112-2202103122203121): complete subsection reference.

<a id="canonical-3203130200220122-1020021230123001-0121332122131122-1033030322210012-0003033202012033-1323130222330120-2313121220122020-3101012310103020"></a>

<a id="canonical-3202003222303312-1102221312032021-1112302211220111-1013131022033323-0032022313223221-1032111333202203-1212122222133203-1120031102221123"></a>

## labels property — Property reference / 012302030031 / 8

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

<a id="canonical-0120011002333123-3031113113030001-0223032301330111-1013033021320112-2301221002313222-3010231001111111-2313332210033212-3120213001232033"></a>

<a id="canonical-2133102301022101-1321221010020203-1222030333121013-2303213133330031-0120303322200333-3023212303221033-0302200131230031-0220101000123301"></a>

## name property — Property reference / 012302030031 / 9

Type: `"string"`. Required.

Name of the Ike1. Must be unique within the namespace.

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

<a id="canonical-3103121331100031-1233123102200312-2211121211120201-0201220013110103-1200233232223313-1211223330123303-1301201020211012-0303031131300000"></a>

<a id="canonical-3022120012131130-0133001333003330-0013232203213212-3331320230331212-3031031123200202-2031020210020323-0332102210011002-2100301133133330"></a>

## namespace property — Property reference / 012302030031 / 10

Type: `"string"`. Required.

Namespace where the Ike1 is created.

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

- [reauth_disabled](resources--ike1--reference--group-001.md#canonical-0130213003021102-3221102331303201-3120330230322233-0201010222030121-1020031323123322-0332031113231101-3031133203003122-3322322301033323): complete subsection reference.

- [reauth_timeout_days](resources--ike1--reference--group-001.md#canonical-0220323021313113-1033010000130003-0220310211031020-1021200012332310-0001102230113202-1001013232032120-1010122331133030-2212010211021003): complete subsection reference.

- [reauth_timeout_hours](resources--ike1--reference--group-001.md#canonical-1320112311132212-0012023330133233-3000022101220033-1312200010211111-3010000301023232-1310120202333121-3111132200223001-1101023111322223): complete subsection reference.

- [timeouts](resources--ike1--reference--group-001.md#canonical-1322221002002100-2333212003121133-1000333211121131-0332301213121012-1132321003200321-0202113330302331-0203000121212223-2132110121231203): complete subsection reference.

- [use_default_keylifetime](resources--ike1--reference--group-001.md#canonical-1112213223110313-1203322131320102-0310210100213022-1123032232110031-1130131201010113-0102010333011230-0201023200023132-0221312331321122): complete subsection reference.

<a id="canonical-3211101203221203-3003122110322113-1302332221230210-3301101321230111-2300201322232221-2130030113013101-1212300033021220-3123023131302030"></a>

## All schema paths — Property reference / 012302030031 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--ike1--reference--group-001.md#canonical-2100302200012333-0211302303131023-1020023212232231-1220310202103000-0201030010200210-2132303030311310-0200212023023333-2021333130313220) |
| `description` | [description](resources--ike1--reference--group-001.md#canonical-3021202201031301-3301012331303230-1110331102330321-2020222012131110-0301113000111100-2222023231112123-0332001220003033-1021021302131130) |
| `disable` | [disable](resources--ike1--reference--group-001.md#canonical-0321001232301021-2222031023230110-3301333212100023-3111223130330111-0133002033222230-0210312132120303-2220112133030322-2112110332103013) |
| `id` | [id](resources--ike1--reference--group-001.md#canonical-2300013102232113-2211230222211231-2123130330131120-0322233113023201-3120001103103112-1111010002200122-0322030132103023-0003211312302231) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](resources--ike1--reference--group-001.md#canonical-2020032032313011-2203210230220313-0301313131333011-3030013013002020-2202221312030011-0210222002210223-2203213331221102-1111122032333202) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](resources--ike1--reference--group-001.md#canonical-1201132113120212-0230022112233330-1333232210003333-2132211302002002-2110300331312111-0110212313121112-2122003203101023-3210321320200303) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](resources--ike1--reference--group-001.md#canonical-0203212322131000-2032321223331021-1122322322001322-0303001313333002-2013212002332300-3222323133132120-0032000121223120-3302300000322022) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](resources--ike1--reference--group-001.md#canonical-2102100333311302-0311213332123033-1021131103121212-1100230010222123-1022121002011212-2223113223001211-1102202230001112-1330323232031202) |
| `labels` | [labels](resources--ike1--reference--group-001.md#canonical-3203130200220122-1020021230123001-0121332122131122-1033030322210012-0003033202012033-1323130222330120-2313121220122020-3101012310103020) |
| `name` | [name](resources--ike1--reference--group-001.md#canonical-0120011002333123-3031113113030001-0223032301330111-1013033021320112-2301221002313222-3010231001111111-2313332210033212-3120213001232033) |
| `namespace` | [namespace](resources--ike1--reference--group-001.md#canonical-3103121331100031-1233123102200312-2211121211120201-0201220013110103-1200233232223313-1211223330123303-1301201020211012-0303031131300000) |
| `reauth_disabled` | [reauth_disabled](resources--ike1--reference--group-001.md#canonical-1123210002230232-2000313331311120-2322132222113213-0011132323000303-2311110133221321-0123130230203302-0333310102303231-0222322300200333) |
| `reauth_timeout_days` | [reauth_timeout_days](resources--ike1--reference--group-001.md#canonical-1330331020120001-1132111000110020-0200033020312213-0311031231320321-0333012212323311-3203212121212102-0110211312112010-2300001302010103) |
| `reauth_timeout_days.duration` | [reauth_timeout_days.duration](resources--ike1--reference--group-001.md#canonical-1322001001100021-1122110130002303-3011233031311111-0113111303131103-0233230300010333-0120211103302321-3331330123320231-3002332313302101) |
| `reauth_timeout_hours` | [reauth_timeout_hours](resources--ike1--reference--group-001.md#canonical-0001131223131021-2000200220001300-2132030211322321-1110002332311022-1303130300120313-3323022113312323-3012033010023120-0022130232322210) |
| `reauth_timeout_hours.duration` | [reauth_timeout_hours.duration](resources--ike1--reference--group-001.md#canonical-3123012321301310-2202110230001031-2222320130331203-2133031013121333-2231132102220131-2233230113201331-0302222130212012-2222021023332020) |
| `timeouts` | [timeouts](resources--ike1--reference--group-001.md#canonical-3031031132200002-1231000221122233-3133112103123220-0223220222032003-0120102000011202-2012031132022111-2203112232011000-1321032003332133) |
| `timeouts.create` | [timeouts.create](resources--ike1--reference--group-001.md#canonical-0213120033201213-1000120102020020-3031111222131302-3312311032031233-3111210101110021-3111300003001213-2301201131023130-2323030111320101) |
| `timeouts.delete` | [timeouts.delete](resources--ike1--reference--group-001.md#canonical-1032303012223332-2012201103230020-2011223020300301-1310122211013123-2003223112321122-0003211232200132-3131102003133200-2000323203031132) |
| `timeouts.read` | [timeouts.read](resources--ike1--reference--group-001.md#canonical-1221202333230112-3300002211302313-0111101301100203-1113301000103100-0312101213103110-1311023212302113-3321333230202120-3333213010020120) |
| `timeouts.update` | [timeouts.update](resources--ike1--reference--group-001.md#canonical-1123132033120330-3010232122313223-2033312002210011-3220033002020002-3201022111323222-1012233300232111-0201011012003013-2002230032333000) |
| `use_default_keylifetime` | [use_default_keylifetime](resources--ike1--reference--group-001.md#canonical-0023203113132233-1123101203003221-1020022100300333-1302231100332132-2003110111312022-3012021101132133-2230212223212030-3132103012132123) |

<a id="canonical-0100301010121312-0112321002121102-2001300022023000-1201131232002110-0130032111231112-2032022031033220-0310111312021002-3323203111232320"></a>

## Next pages — Property reference / 012302030031 / 12

- [ike_keylifetime_hours](resources--ike1--reference--group-001.md#canonical-1331103023310210-3303300122110230-3303110012211012-3210310233133331-3031130111311300-2003010332103002-0331110001201232-1202020111313323)
- [ike_keylifetime_minutes](resources--ike1--reference--group-001.md#canonical-1331031321203231-2213012032201103-0033033320110332-2023021001032213-0300100220313231-2130011101032101-0002033320121112-2202103122203121)
- [reauth_disabled](resources--ike1--reference--group-001.md#canonical-0130213003021102-3221102331303201-3120330230322233-0201010222030121-1020031323123322-0332031113231101-3031133203003122-3322322301033323)
- [reauth_timeout_days](resources--ike1--reference--group-001.md#canonical-0220323021313113-1033010000130003-0220310211031020-1021200012332310-0001102230113202-1001013232032120-1010122331133030-2212010211021003)
- [reauth_timeout_hours](resources--ike1--reference--group-001.md#canonical-1320112311132212-0012023330133233-3000022101220033-1312200010211111-3010000301023232-1310120202333121-3111132200223001-1101023111322223)
- [timeouts](resources--ike1--reference--group-001.md#canonical-1322221002002100-2333212003121133-1000333211121131-0332301213121012-1132321003200321-0202113330302331-0203000121212223-2132110121231203)
- [use_default_keylifetime](resources--ike1--reference--group-001.md#canonical-1112213223110313-1203322131320102-0310210100213022-1123032232110031-1130131201010113-0102010333011230-0201023200023132-0221312331321122)
- [xcsh_ike1](../resources/ike1.md#canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130)

<a id="canonical-1331103023310210-3303300122110230-3303110012211012-3210310233133331-3031130111311300-2003010332103002-0331110001201232-1202020111313323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321110032310320-2121202123311322-3300222122313213-2200301301222220-2300201302003003-1220131223302233-0312222333310203-3130013031320101"></a>

## ike_keylifetime_hours — ike_keylifetime_hours / 030131230230 / 2

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130)
- [Property reference](resources--ike1--reference--group-001.md#canonical-0231322323021312-2223311002332013-3200002221122100-2200120310330232-3230312122213012-1111030312321102-2331002032223020-1222113231322110)
- ike_keylifetime_hours

<a id="canonical-2020032032313011-2203210230220313-0301313131333011-3030013013002020-2202221312030011-0210222002210223-2203213331221102-1111122032333202"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ike\_keylifetime\_hours, ike\_keylifetime\_minutes, use\_default\_keylifetime; Default:
use\_default\_keylifetime\] Configuration parameter for ike keylifetime hours.

Upstream description:

Input Hours.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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

OneOf alternatives in this subsection:

- [ike_keylifetime_hours](resources--ike1--reference--group-001.md#canonical-2020032032313011-2203210230220313-0301313131333011-3030013013002020-2202221312030011-0210222002210223-2203213331221102-1111122032333202)
- [ike_keylifetime_minutes](resources--ike1--reference--group-001.md#canonical-0203212322131000-2032321223331021-1122322322001322-0303001313333002-2013212002332300-3222323133132120-0032000121223120-3302300000322022)
- [use_default_keylifetime](resources--ike1--reference--group-001.md#canonical-0023203113132233-1123101203003221-1020022100300333-1302231100332132-2003110111312022-3012021101132133-2230212223212030-3132103012132123)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ike_keylifetime_hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-2003332213313103-2110002232311200-2310011213221301-1313110020331113-3321000320101222-3022013310003123-0011000133311101-0212333303212010"></a>

## Direct properties — ike_keylifetime_hours / 030131230230 / 3

<a id="canonical-1201132113120212-0230022112233330-1333232210003333-2132211302002002-2110300331312111-0110212313121112-2122003203101023-3210321320200303"></a>

<a id="canonical-3012302001331123-3203000020002132-0323001023011030-1000002022121120-0131231030100332-3211003020133111-3003031330121201-3132032232201103"></a>

## duration property — ike_keylifetime_hours / 030131230230 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1010112223122333-3232022311100023-0132203210302022-0121000322100311-2221233100030231-3203231123333200-3031230220033103-2113031132033001"></a>

## Next pages — ike_keylifetime_hours / 030131230230 / 5

- [Property reference](resources--ike1--reference--group-001.md#canonical-0231322323021312-2223311002332013-3200002221122100-2200120310330232-3230312122213012-1111030312321102-2331002032223020-1222113231322110)
- [xcsh_ike1](../resources/ike1.md#canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130)

<a id="canonical-1331031321203231-2213012032201103-0033033320110332-2023021001032213-0300100220313231-2130011101032101-0002033320121112-2202103122203121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320113033300102-2201022330102103-1322103310003000-1002031303012022-2121332003030002-3213103222031032-1032233302220213-0011033323123233"></a>

## ike_keylifetime_minutes — ike_keylifetime_minutes / 300101123022 / 2

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130)
- [Property reference](resources--ike1--reference--group-001.md#canonical-0231322323021312-2223311002332013-3200002221122100-2200120310330232-3230312122213012-1111030312321102-2331002032223020-1222113231322110)
- ike_keylifetime_minutes

<a id="canonical-0203212322131000-2032321223331021-1122322322001322-0303001313333002-2013212002332300-3222323133132120-0032000121223120-3302300000322022"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ike keylifetime minutes.

Upstream description:

Set IKE Key Lifetime in minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
ike_keylifetime_minutes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0032011022120031-0123120210101330-2002210133002020-2201333002110022-1313213003310330-0212200221030022-3121123210211201-1110130121221320"></a>

## Direct properties — ike_keylifetime_minutes / 300101123022 / 3

<a id="canonical-2102100333311302-0311213332123033-1021131103121212-1100230010222123-1022121002011212-2223113223001211-1102202230001112-1330323232031202"></a>

<a id="canonical-0202230230333323-1220300330111302-1303221113123000-0223123010101301-3102211101032131-3322233133312323-1321022133233312-2002133231210211"></a>

## duration property — ike_keylifetime_minutes / 300101123022 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(10, 300),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0132133032201001-0032300213232132-0033321233122303-2132312232123111-3222032103313301-0230311032012221-2021201323332230-0103213130310012"></a>

## Next pages — ike_keylifetime_minutes / 300101123022 / 5

- [Property reference](resources--ike1--reference--group-001.md#canonical-0231322323021312-2223311002332013-3200002221122100-2200120310330232-3230312122213012-1111030312321102-2331002032223020-1222113231322110)
- [xcsh_ike1](../resources/ike1.md#canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130)

<a id="canonical-0130213003021102-3221102331303201-3120330230322233-0201010222030121-1020031323123322-0332031113231101-3031133203003122-3322322301033323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110122103312232-3211311301203302-0321201003013313-0011222300222020-1232133311133010-0330022123232322-1213123311020010-0001332330221310"></a>

## reauth_disabled — reauth_disabled / 133132132031 / 2

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130)
- [Property reference](resources--ike1--reference--group-001.md#canonical-0231322323021312-2223311002332013-3200002221122100-2200120310330232-3230312122213012-1111030312321102-2331002032223020-1222113231322110)
- reauth_disabled

<a id="canonical-1123210002230232-2000313331311120-2322132222113213-0011132323000303-2311110133221321-0123130230203302-0333310102303231-0222322300200333"></a>

Type: `["object", {}]`. Optional.

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

- [reauth_disabled](resources--ike1--reference--group-001.md#canonical-1123210002230232-2000313331311120-2322132222113213-0011132323000303-2311110133221321-0123130230203302-0333310102303231-0222322300200333)
- [reauth_timeout_days](resources--ike1--reference--group-001.md#canonical-1330331020120001-1132111000110020-0200033020312213-0311031231320321-0333012212323311-3203212121212102-0110211312112010-2300001302010103)
- [reauth_timeout_hours](resources--ike1--reference--group-001.md#canonical-0001131223131021-2000200220001300-2132030211322321-1110002332311022-1303130300120313-3323022113312323-3012033010023120-0022130232322210)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
reauth_disabled = {}
```

<a id="canonical-1221330123331232-0010333021102313-2332003323000000-0323311103212101-3010102101323322-3031202003132302-0000230012220212-2221113322033330"></a>

## Direct properties — reauth_disabled / 133132132031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103223211322121-0102232232110310-3321030201220320-1012013001223123-1133023321030202-3330201131011121-3032221300203131-3003010233100332"></a>

## Next pages — reauth_disabled / 133132132031 / 4

- [Property reference](resources--ike1--reference--group-001.md#canonical-0231322323021312-2223311002332013-3200002221122100-2200120310330232-3230312122213012-1111030312321102-2331002032223020-1222113231322110)
- [xcsh_ike1](../resources/ike1.md#canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130)

<a id="canonical-0220323021313113-1033010000130003-0220310211031020-1021200012332310-0001102230113202-1001013232032120-1010122331133030-2212010211021003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331232201132332-1231320030111033-2231321100013120-1032333222302131-2202000030121122-0003332113211011-0330230110203032-1330303000010003"></a>

## reauth_timeout_days — reauth_timeout_days / 331213010100 / 2

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130)
- [Property reference](resources--ike1--reference--group-001.md#canonical-0231322323021312-2223311002332013-3200002221122100-2200120310330232-3230312122213012-1111030312321102-2331002032223020-1222113231322110)
- reauth_timeout_days

<a id="canonical-1330331020120001-1132111000110020-0200033020312213-0311031231320321-0333012212323311-3203212121212102-0110211312112010-2300001302010103"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for reauth timeout days.

Upstream description:

Set Duration in days.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
reauth_timeout_days {
  # Configure direct properties listed below.
}
```

<a id="canonical-3230323330021210-3130330302002132-1221201123310022-1322322322131322-1000012000320011-0013033203312033-2103010122021213-0300332302131033"></a>

## Direct properties — reauth_timeout_days / 331213010100 / 3

<a id="canonical-1322001001100021-1122110130002303-3011233031311111-0113111303131103-0233230300010333-0120211103302321-3331330123320231-3002332313302101"></a>

<a id="canonical-0012123302010222-1102300313113031-0233323202301011-0032003330031113-3213303332322000-1302302020111011-2333112203101111-1303312330001232"></a>

## duration property — reauth_timeout_days / 331213010100 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2011200202201312-2000110202322233-3302320112313201-1232211313131321-3121332302101200-1122220333203032-1212313220200322-1310102113013221"></a>

## Next pages — reauth_timeout_days / 331213010100 / 5

- [Property reference](resources--ike1--reference--group-001.md#canonical-0231322323021312-2223311002332013-3200002221122100-2200120310330232-3230312122213012-1111030312321102-2331002032223020-1222113231322110)
- [xcsh_ike1](../resources/ike1.md#canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130)

<a id="canonical-1320112311132212-0012023330133233-3000022101220033-1312200010211111-3010000301023232-1310120202333121-3111132200223001-1101023111322223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112112231331223-3031013101313322-0111030112010132-1101311311133022-3111010233301231-3303023021220111-3103330003113310-1313223211332122"></a>

## reauth_timeout_hours — reauth_timeout_hours / 333032130311 / 2

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130)
- [Property reference](resources--ike1--reference--group-001.md#canonical-0231322323021312-2223311002332013-3200002221122100-2200120310330232-3230312122213012-1111030312321102-2331002032223020-1222113231322110)
- reauth_timeout_hours

<a id="canonical-0001131223131021-2000200220001300-2132030211322321-1110002332311022-1303130300120313-3323022113312323-3012033010023120-0022130232322210"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for reauth timeout hours.

Upstream description:

Input Hours.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
reauth_timeout_hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-1111233210201131-1313010302202132-2211201123133010-0121001010230112-2332200323232300-2100012021203013-1111032103031213-2232022330131321"></a>

## Direct properties — reauth_timeout_hours / 333032130311 / 3

<a id="canonical-3123012321301310-2202110230001031-2222320130331203-2133031013121333-2231132102220131-2233230113201331-0302222130212012-2222021023332020"></a>

<a id="canonical-2120030033032321-0220030302312101-3232121320212022-3212031332133213-0212221332200121-1303200332103011-1202023021320030-2121333100213021"></a>

## duration property — reauth_timeout_hours / 333032130311 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0021101101231133-3232023303103000-1313020113233133-1000321322103213-3123011231303210-1311103130001301-0131312230011232-0002302200321132"></a>

## Next pages — reauth_timeout_hours / 333032130311 / 5

- [Property reference](resources--ike1--reference--group-001.md#canonical-0231322323021312-2223311002332013-3200002221122100-2200120310330232-3230312122213012-1111030312321102-2331002032223020-1222113231322110)
- [xcsh_ike1](../resources/ike1.md#canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130)

<a id="canonical-1322221002002100-2333212003121133-1000333211121131-0332301213121012-1132321003200321-0202113330302331-0203000121212223-2132110121231203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310032232230000-3002213022132213-2020303120231321-3103202313112102-1102223333130100-1002220221013321-1000303101210323-1032331100000230"></a>

## timeouts — timeouts / 210302321110 / 2

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130)
- [Property reference](resources--ike1--reference--group-001.md#canonical-0231322323021312-2223311002332013-3200002221122100-2200120310330232-3230312122213012-1111030312321102-2331002032223020-1222113231322110)
- timeouts

<a id="canonical-3031031132200002-1231000221122233-3133112103123220-0223220222032003-0120102000011202-2012031132022111-2203112232011000-1321032003332133"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313323001123321-3313331221201111-2133122012102123-2220113023322221-0112310311102120-0232332101023100-2122230230311233-0212331200032303"></a>

## Direct properties — timeouts / 210302321110 / 3

<a id="canonical-0213120033201213-1000120102020020-3031111222131302-3312311032031233-3111210101110021-3111300003001213-2301201131023130-2323030111320101"></a>

<a id="canonical-0213223202012020-1302112210122120-2312230320313011-3322320001202301-2312220132102321-0123030200220131-1031031001110003-1203232123011202"></a>

## create property — timeouts / 210302321110 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1032303012223332-2012201103230020-2011223020300301-1310122211013123-2003223112321122-0003211232200132-3131102003133200-2000323203031132"></a>

<a id="canonical-0302220331321200-3022333022101211-3203321020132000-0303230112111033-2010223032000301-0131333321032022-2013100012202332-3313312331202020"></a>

## delete property — timeouts / 210302321110 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1221202333230112-3300002211302313-0111101301100203-1113301000103100-0312101213103110-1311023212302113-3321333230202120-3333213010020120"></a>

<a id="canonical-2100033101102121-1033332000132022-1112110021301130-0222023313223013-0330111033232312-3221223330301013-3110300221123220-0130111100323100"></a>

## read property — timeouts / 210302321110 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1123132033120330-3010232122313223-2033312002210011-3220033002020002-3201022111323222-1012233300232111-0201011012003013-2002230032333000"></a>

<a id="canonical-1331312202033213-1232111010212320-1103101013333032-2220201221200020-2311220103020022-2000222233101033-0313001231023323-3200031131113132"></a>

## update property — timeouts / 210302321110 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1201320130101220-2122013003022103-1223122312011003-3231110101232311-0233211210100110-1131030313311021-2311103032232310-2310300021133333"></a>

## Next pages — timeouts / 210302321110 / 8

- [Property reference](resources--ike1--reference--group-001.md#canonical-0231322323021312-2223311002332013-3200002221122100-2200120310330232-3230312122213012-1111030312321102-2331002032223020-1222113231322110)
- [xcsh_ike1](../resources/ike1.md#canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130)

<a id="canonical-1112213223110313-1203322131320102-0310210100213022-1123032232110031-1130131201010113-0102010333011230-0201023200023132-0221312331321122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200223200110321-0110031201122110-2222030221010101-1113230002021302-3231333301333013-3312311221310101-3310113220221102-2031221113320313"></a>

## use_default_keylifetime — use_default_keylifetime / 003021030231 / 2

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130)
- [Property reference](resources--ike1--reference--group-001.md#canonical-0231322323021312-2223311002332013-3200002221122100-2200120310330232-3230312122213012-1111030312321102-2331002032223020-1222113231322110)
- use_default_keylifetime

<a id="canonical-0023203113132233-1123101203003221-1020022100300333-1302231100332132-2003110111312022-3012021101132133-2230212223212030-3132103012132123"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
use_default_keylifetime = {}
```

<a id="canonical-2312031222122301-3201312321133200-0203103302220320-2000313110203302-3210303131301003-1233201131312130-3231221033122223-1101200122310230"></a>

## Direct properties — use_default_keylifetime / 003021030231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110221000032032-2130020100332013-3100020112321003-2011300233130112-1111122220001301-0320323123213330-0323102331031112-0123123123032320"></a>

## Next pages — use_default_keylifetime / 003021030231 / 4

- [Property reference](resources--ike1--reference--group-001.md#canonical-0231322323021312-2223311002332013-3200002221122100-2200120310330232-3230312122213012-1111030312321102-2331002032223020-1222113231322110)
- [xcsh_ike1](../resources/ike1.md#canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130)
