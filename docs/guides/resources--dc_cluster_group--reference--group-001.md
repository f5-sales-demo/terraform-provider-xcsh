---
page_title: "xcsh_dc_cluster_group reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dc_cluster_group reference."
---

# xcsh_dc_cluster_group reference

<a id="canonical-3020323102031131-2120211103310111-3012300022003323-3023301231000121-3002020033132213-0201300331100120-3122330131102121-1132110021113321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0023121103022030-2302020033112202-0212333320231232-3001211200201011-0122330223232010-0110110123001022-0123201000313123-3333311302021321)
- Property reference

<a id="canonical-0030120202132033-2003100000213323-0231220323201203-2230211311311103-1100222101013102-3301002313000031-1133021320101030-3312020111232223"></a>

### Direct properties for `xcsh_dc_cluster_group`

<a id="canonical-0133223330313312-1020322132330302-3313010313023322-2022010110130201-2110000113313222-0122110322311010-0133300321030310-3312130101311230"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

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

<a id="canonical-0232211012222112-3100332021330220-0221222332103222-3302002320232221-2202001033300212-0132310003222212-3321313311122323-1311332223233231"></a>

<a id="canonical-0301032312220110-0130203321000120-0101122010211231-1113232120232323-0301010022301030-3020320223232001-2123323101001030-3311321103312231"></a>

#### `description` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3333203023300101-3300221133300311-0200202010123300-1110003110100311-0132202002202311-2011132330132020-3013203322211023-1100301030013112"></a>

<a id="canonical-3110300313212133-0300013022203120-1331011320221110-3000102212211202-3221223313012210-0230102223320201-3011101330112110-1230321111133201"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

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

<a id="canonical-1111311121221233-1213010133312300-2120211300210101-2021012033201310-0231112031131322-0013312332103131-2110121302131130-3312032310031000"></a>

<a id="canonical-1030201323331203-2002213323022002-2213100021321123-0033331101313232-0310033330022112-1003301103100211-1330100321022132-2231222001322311"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2031323132012212-3101023210202303-1021231212201223-1020001023121211-0132022001233220-2212030222331013-1102122323301020-2032202302223321"></a>

<a id="canonical-0020323123131113-1333221322131120-0130223111002330-2230210111233002-2132100011301312-1223112101202212-2033133331210021-1101120111102333"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

Additional upstream details:

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

<a id="canonical-1220323333102103-2203113331310112-2120223002332321-3313332322303010-0313233302200122-3001000111201002-1202120020002120-1012301010131220"></a>

<a id="canonical-0332133313332021-0323323233311110-0331022121122130-0113003032101222-0103212022311332-0323122220011102-0332203030301313-2010013110003210"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Dc Cluster Group. Must be unique within the namespace.

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1121031332002212-0032122131303131-2220012301200021-3313001111330210-1231201311220203-3310232013121030-3132210322322123-2100131332201000"></a>

<a id="canonical-0020111100030300-3110333132230132-1111323303320020-2033130230111323-3020032123130220-3200311112003320-0020133013123221-1233203230320031"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace for the Dc Cluster Group. The F5 XC API restricts this resource to the system namespace;
it defaults to that value and may be omitted.

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [timeouts](resources--dc_cluster_group--reference--group-001.md#canonical-3221033032013000-0232031300011123-3111110021220110-0033221012301112-2021211302222132-1321222310321030-0000122132210321-3220310132212012): complete subsection reference.

- [type](resources--dc_cluster_group--reference--group-001.md#canonical-3110132011101032-2303210303323220-0123223130111211-2212123323110202-2013332020111101-3323132110301011-1202311001131331-3100003213112221): complete subsection reference.

<a id="canonical-3331200330302302-0031110033133313-3300030110012333-0320132212030013-0312333222123031-1201211020131312-2103000300320003-0223030232033313"></a>

### All schema paths for `xcsh_dc_cluster_group`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--dc_cluster_group--reference--group-001.md#canonical-0133223330313312-1020322132330302-3313010313023322-2022010110130201-2110000113313222-0122110322311010-0133300321030310-3312130101311230) |
| `description` | [description](resources--dc_cluster_group--reference--group-001.md#canonical-0232211012222112-3100332021330220-0221222332103222-3302002320232221-2202001033300212-0132310003222212-3321313311122323-1311332223233231) |
| `disable` | [disable](resources--dc_cluster_group--reference--group-001.md#canonical-3333203023300101-3300221133300311-0200202010123300-1110003110100311-0132202002202311-2011132330132020-3013203322211023-1100301030013112) |
| `id` | [ID](resources--dc_cluster_group--reference--group-001.md#canonical-1111311121221233-1213010133312300-2120211300210101-2021012033201310-0231112031131322-0013312332103131-2110121302131130-3312032310031000) |
| `labels` | [labels](resources--dc_cluster_group--reference--group-001.md#canonical-2031323132012212-3101023210202303-1021231212201223-1020001023121211-0132022001233220-2212030222331013-1102122323301020-2032202302223321) |
| `name` | [name](resources--dc_cluster_group--reference--group-001.md#canonical-1220323333102103-2203113331310112-2120223002332321-3313332322303010-0313233302200122-3001000111201002-1202120020002120-1012301010131220) |
| `namespace` | [namespace](resources--dc_cluster_group--reference--group-001.md#canonical-1121031332002212-0032122131303131-2220012301200021-3313001111330210-1231201311220203-3310232013121030-3132210322322123-2100131332201000) |
| `timeouts` | [timeouts](resources--dc_cluster_group--reference--group-001.md#canonical-1011213211301223-0120232321310103-2100321001102033-1312331110001331-0121021102211211-2213222203021113-1113311120022203-1333010220001102) |
| `timeouts.create` | [timeouts.create](resources--dc_cluster_group--reference--group-001.md#canonical-1200103310013223-1331230330120312-2000230001310233-3221112001321302-0302032011330213-2032103101301200-0000213103131203-1132031232121132) |
| `timeouts.delete` | [timeouts.delete](resources--dc_cluster_group--reference--group-001.md#canonical-1100310111003133-1133030032322111-0202333202213322-3212203001313111-3310011110011301-0113222110321202-2332031032020322-1211312321002002) |
| `timeouts.read` | [timeouts.read](resources--dc_cluster_group--reference--group-001.md#canonical-2011333322113011-2102120200001213-2120100121323212-2330223313133213-1032320300021132-2122102112300103-2312100033121103-2002302202111212) |
| `timeouts.update` | [timeouts.update](resources--dc_cluster_group--reference--group-001.md#canonical-2200020002330301-0300120023011003-2223221223131221-0320321212030123-3030120112311112-3131331210310213-0030132131321321-0003013011003002) |
| `type` | [type](resources--dc_cluster_group--reference--group-001.md#canonical-3210233311300301-3212333102330113-1210231013010111-3032102003122001-0203022133333232-0003131231003130-0010222113012133-1223220133111232) |
| `type.control_and_data_plane_mesh` | [type.control_and_data_plane_mesh](resources--dc_cluster_group--reference--group-001.md#canonical-1130212320132221-1021200203021211-3301220033211110-1001331302012111-3121300032110332-3113221231023212-1311102320032022-2032001321020131) |
| `type.data_plane_mesh` | [type.data_plane_mesh](resources--dc_cluster_group--reference--group-001.md#canonical-3232011102201321-3333100203022311-2023333232123212-2310003001313122-1012232233101033-1013201320022233-0113333201101000-0130202300102300) |

<a id="canonical-3221033032013000-0232031300011123-3111110021220110-0033221012301112-2021211302222132-1321222310321030-0000122132210321-3220310132212012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0023121103022030-2302020033112202-0212333320231232-3001211200201011-0122330223232010-0110110123001022-0123201000313123-3333311302021321)
- [Property reference](resources--dc_cluster_group--reference--group-001.md#canonical-3020323102031131-2120211103310111-3012300022003323-3023301231000121-3002020033132213-0201300331100120-3122330131102121-1132110021113321)
- timeouts

<a id="canonical-1011213211301223-0120232321310103-2100321001102033-1312331110001331-0121021102211211-2213222203021113-1113311120022203-1333010220001102"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232121332220233-2122121310112032-0132320331011101-2133111213233202-3301310033020300-1233332331021012-3210323200101312-2200113010212010"></a>

### Direct properties for `timeouts`

<a id="canonical-1200103310013223-1331230330120312-2000230001310233-3221112001321302-0302032011330213-2032103101301200-0000213103131203-1132031232121132"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1100310111003133-1133030032322111-0202333202213322-3212203001313111-3310011110011301-0113222110321202-2332031032020322-1211312321002002"></a>

<a id="canonical-0102100302033313-3330332020020011-0100113130032000-2223121010012132-2002111212330032-2230331132033301-1033020101001100-2310320022033303"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2011333322113011-2102120200001213-2120100121323212-2330223313133213-1032320300021132-2122102112300103-2312100033121103-2002302202111212"></a>

<a id="canonical-3033223010003123-1200210103003003-2311033222110113-1123332132320303-1103010020110302-0010212130130103-2121201112023233-1312100112031311"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2200020002330301-0300120023011003-2223221223131221-0320321212030123-3030120112311112-3131331210310213-0030132131321321-0003013011003002"></a>

<a id="canonical-0110001223013302-0120033232321100-1120113232323102-2000100230013130-3103333333203312-2210020012123013-1111300333221030-0123113310231222"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3110132011101032-2303210303323220-0123223130111211-2212123323110202-2013332020111101-3323132110301011-1202311001131331-3100003213112221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `type` properties

Breadcrumbs:

- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0023121103022030-2302020033112202-0212333320231232-3001211200201011-0122330223232010-0110110123001022-0123201000313123-3333311302021321)
- [Property reference](resources--dc_cluster_group--reference--group-001.md#canonical-3020323102031131-2120211103310111-3012300022003323-3023301231000121-3002020033132213-0201300331100120-3122330131102121-1132110021113321)
- type

<a id="canonical-3210233311300301-3212333102330113-1210231013010111-3032102003122001-0203022133333232-0003131231003130-0010222113012133-1223220133111232"></a>

Type: `"object"`. single nested block, Optional.

DC Cluster Group Mesh Type. Details of DC Cluster Group Mesh Type.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("control_and_data_plane_mesh",
    "data_plane_mesh")}
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
  "x-ves-oneof-field-dc_cluster_group_mesh_choice": "[\"control_and_data_plane_mesh\",\"data_plane_mesh\"]"
}
```

Terraform syntax:

```terraform
type {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122033203133330-0102300020233331-2303103331110332-0120230020020320-2103310210221230-3333103323211300-3022012031031202-2302233101012201"></a>

### Direct properties for `type`

- [control_and_data_plane_mesh](resources--dc_cluster_group--reference--group-001.md#canonical-0200311211231101-3231031202003133-3121110031113210-2023321233130123-3002332022110202-2110230332101211-0310100113102021-0122022132222311): complete subsection reference.

- [data_plane_mesh](resources--dc_cluster_group--reference--group-001.md#canonical-2320232132000312-1010113003001200-1311022122203122-2211313203022221-3112122033021320-3313203000022331-1320121233021221-3321202230320213): complete subsection reference.

<a id="canonical-0200311211231101-3231031202003133-3121110031113210-2023321233130123-3002332022110202-2110230332101211-0310100113102021-0122022132222311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `type.control_and_data_plane_mesh` properties

Breadcrumbs:

- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0023121103022030-2302020033112202-0212333320231232-3001211200201011-0122330223232010-0110110123001022-0123201000313123-3333311302021321)
- [Property reference](resources--dc_cluster_group--reference--group-001.md#canonical-3020323102031131-2120211103310111-3012300022003323-3023301231000121-3002020033132213-0201300331100120-3122330131102121-1132110021113321)
- [type](resources--dc_cluster_group--reference--group-001.md#canonical-3110132011101032-2303210303323220-0123223130111211-2212123323110202-2013332020111101-3323132110301011-1202311001131331-3100003213112221)
- type.control_and_data_plane_mesh

<a id="canonical-1130212320132221-1021200203021211-3301220033211110-1001331302012111-3121300032110332-3113221231023212-1311102320032022-2032001321020131"></a>

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
control_and_data_plane_mesh = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320232132000312-1010113003001200-1311022122203122-2211313203022221-3112122033021320-3313203000022331-1320121233021221-3321202230320213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `type.data_plane_mesh` properties

Breadcrumbs:

- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0023121103022030-2302020033112202-0212333320231232-3001211200201011-0122330223232010-0110110123001022-0123201000313123-3333311302021321)
- [Property reference](resources--dc_cluster_group--reference--group-001.md#canonical-3020323102031131-2120211103310111-3012300022003323-3023301231000121-3002020033132213-0201300331100120-3122330131102121-1132110021113321)
- [type](resources--dc_cluster_group--reference--group-001.md#canonical-3110132011101032-2303210303323220-0123223130111211-2212123323110202-2013332020111101-3323132110301011-1202311001131331-3100003213112221)
- type.data_plane_mesh

<a id="canonical-3232011102201321-3333100203022311-2023333232123212-2310003001313122-1012232233101033-1013201320022233-0113333201101000-0130202300102300"></a>

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
data_plane_mesh = {}
```

This is an empty object or choice marker. It has no direct properties.
