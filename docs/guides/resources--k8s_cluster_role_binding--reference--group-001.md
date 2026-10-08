---
page_title: "xcsh_k8s_cluster_role_binding reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role_binding reference."
---

# xcsh_k8s_cluster_role_binding reference

<a id="canonical-3122222001022132-1111202303122022-2033123222301020-0030321020302103-0030301220220131-0121213310030202-3231213321200003-3112033132100012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-1203120300333222-3203223200312120-0002210222112313-1113120331003323-3102000331030303-1133200312133122-1031101033120033-3233030130121012)
- Property reference

<a id="canonical-2213223301210201-0110020002223303-3033220200121000-2312130121023223-1031031201031032-1312323303030303-3020210131333312-2033211330011231"></a>

### Direct properties for `xcsh_k8s_cluster_role_binding`

<a id="canonical-2230213021122321-1033322110110020-0120313000103302-3203220303323130-2123012201032103-3230203332231021-2003320233132131-1033002321332122"></a>

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

<a id="canonical-0201003003302122-3323013110003110-3023012031310011-3320111313313111-2203312030300132-3312012032100122-0011330033110212-2111302022023220"></a>

<a id="canonical-1012012220032310-3021020130221133-3133331021200201-0001003332312121-0033031113132102-1010302101233102-1221011032100130-1110200133111032"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3333010312323211-0213121001222001-3111223332132323-0220122212001010-1102300333233101-2133000120310300-2130002003131313-1200003102230333"></a>

<a id="canonical-2033112203322313-2201322200123120-2131110232210033-0021223110012002-1120323130030122-2102201021332333-1121203330101202-2302121131113200"></a>

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

<a id="canonical-3121102311310022-1020233013120130-2231121230032021-2110330003231111-2133131320010102-1011322132232123-2030030021110103-2200130210322003"></a>

<a id="canonical-2203313312311133-3113132202003321-0122131030331330-0022011202130330-1333110123113132-0311000132002031-0322021311220312-3100221231223002"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [k8s_cluster_role](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-2000233113113223-1103320223231300-1313321310123311-2010012203313311-3333001030101311-0011233220332220-0003213230030100-1221311033110121): complete subsection reference.

<a id="canonical-2321323001022103-2010231203033033-0011321022333123-1002301112132201-2031321122201232-3131003322330230-1023203131230020-3020332111001333"></a>

<a id="canonical-1001011002330323-2012131030020331-3331201221233030-3131330001033222-0203300110221323-0332020231033313-3110211230122102-1121013003120231"></a>

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

<a id="canonical-3112003220013032-3320331310312113-2202332012203011-0132131321001201-1000201020201032-2203322232003233-2000333222021332-0013011010010031"></a>

<a id="canonical-0201312011330300-0012320033213033-3103021123303130-1010022231333200-1130132212231210-1302331211331212-1310310300321030-1230303010222301"></a>

#### `name` property

Type: `"string"`. Required.

Name of the K8S Cluster Role Binding. Must be unique within the namespace.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2032223323130130-1213303022233122-3033110223103131-2033313020231222-2102313202011031-0100003002220322-1122122213221300-1130030002200130"></a>

<a id="canonical-3333102302331123-3212013002210011-2213102132222333-0222112210203112-3313032203000322-2013221303212010-3333012233121222-2021122231222112"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the K8S Cluster Role Binding is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [subjects](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-3202231323013321-2032223021010103-2021303133012001-3313313131013202-1232010101210021-2302132012221100-0313212310022231-0130230122331332): complete subsection reference.

- [timeouts](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-3302013110231321-1022022233001010-3003220330100331-2213210110000123-0301231200201111-2302330331332300-0221300020011210-2212213201233100): complete subsection reference.

<a id="canonical-2232132032332112-3200031120312021-0233221122100011-0200021130211323-3211211103000110-1121030233021310-2112212313033303-3033301103200023"></a>

### All schema paths for `xcsh_k8s_cluster_role_binding`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-2230213021122321-1033322110110020-0120313000103302-3203220303323130-2123012201032103-3230203332231021-2003320233132131-1033002321332122) |
| `description` | [description](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-0201003003302122-3323013110003110-3023012031310011-3320111313313111-2203312030300132-3312012032100122-0011330033110212-2111302022023220) |
| `disable` | [disable](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-3333010312323211-0213121001222001-3111223332132323-0220122212001010-1102300333233101-2133000120310300-2130002003131313-1200003102230333) |
| `id` | [ID](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-3121102311310022-1020233013120130-2231121230032021-2110330003231111-2133131320010102-1011322132232123-2030030021110103-2200130210322003) |
| `k8s_cluster_role` | [k8s_cluster_role](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-3322113321332111-2021201100301132-1301022032210230-0322212222320111-2332101121122303-2103200230011020-0203100121031102-1110233313120232) |
| `k8s_cluster_role.name` | [k8s_cluster_role.name](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-1202231323132032-3121033033101032-2002300103331003-0132020322101213-3330032012210002-3111311113000110-3220220223002013-1330031210003010) |
| `k8s_cluster_role.namespace` | [k8s_cluster_role.namespace](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-1333123231320120-1131001220132002-3231233001031212-0320102111313111-2012132100232120-1030103123202222-2323023301211313-0313102322001200) |
| `k8s_cluster_role.tenant` | [k8s_cluster_role.tenant](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-2120011031123232-1333001002011011-3123133230013212-3103311013322233-2321100300021210-3330111330133111-1003200221230331-3202101100103310) |
| `labels` | [labels](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-2321323001022103-2010231203033033-0011321022333123-1002301112132201-2031321122201232-3131003322330230-1023203131230020-3020332111001333) |
| `name` | [name](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-3112003220013032-3320331310312113-2202332012203011-0132131321001201-1000201020201032-2203322232003233-2000333222021332-0013011010010031) |
| `namespace` | [namespace](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-2032223323130130-1213303022233122-3033110223103131-2033313020231222-2102313202011031-0100003002220322-1122122213221300-1130030002200130) |
| `subjects` | [subjects](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-1123332223023331-3001202100103301-3110213031301313-2033300113023131-1112012030032230-3310113223100102-2012022002223233-3232233213112112) |
| `subjects.group` | [subjects.group](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-3001320133312121-2032300320003221-0233322002010100-1320103030221022-2023111320110231-2310320302033102-0212300300202320-1311330223333123) |
| `subjects.service_account` | [subjects.service_account](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-3222011010230332-2032011332102303-1222222220221201-1232201331331013-2112023023332120-0203123231310312-0313020032211100-2023331113132112) |
| `subjects.service_account.name` | [subjects.service_account.name](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-0332103311311120-1231330020010002-3021133320002030-2120323200023312-2101222222200332-0133331322022220-1002011022300301-3003312011112322) |
| `subjects.service_account.namespace` | [subjects.service_account.namespace](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-3111020233230021-2323211331210231-0210212310113211-3313123101102002-2130030233122230-0022200133213132-0210311111230303-3100122103103213) |
| `subjects.user` | [subjects.user](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-0000332030321000-3030331221030003-2002303303222233-3232121001011221-1131211300133212-0121302023302102-3232320011121111-2023200011320322) |
| `timeouts` | [timeouts](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-3133310020221303-0022302310312110-1323212012012212-2222113102320123-2231131031312022-3323033023000021-0011313213223111-2010023020303031) |
| `timeouts.create` | [timeouts.create](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-2310020323000220-2331130020133100-3132121011112133-0103120021333232-2031331112102333-3302202212000132-2103213000032011-1100032320121002) |
| `timeouts.delete` | [timeouts.delete](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-2130003220030003-0033110320320222-2021220321203121-0111111022111112-2121021131231132-2111213210032331-1000131323303311-3202113101110231) |
| `timeouts.read` | [timeouts.read](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-1111201023111201-1302111211313201-2032020320032321-0220101221100223-1300111201213023-0201011202000320-3322120102232210-1323312312310323) |
| `timeouts.update` | [timeouts.update](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-3032020102331112-2031332133333001-0010033320302030-3333021320203022-1020132121033113-1301301112112031-3332001112101230-0121230233223310) |

<a id="canonical-2000233113113223-1103320223231300-1313321310123311-2010012203313311-3333001030101311-0011233220332220-0003213230030100-1221311033110121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `k8s_cluster_role` properties

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-1203120300333222-3203223200312120-0002210222112313-1113120331003323-3102000331030303-1133200312133122-1031101033120033-3233030130121012)
- [Property reference](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-3122222001022132-1111202303122022-2033123222301020-0030321020302103-0030301220220131-0121213310030202-3231213321200003-3112033132100012)
- k8s_cluster_role

<a id="canonical-3322113321332111-2021201100301132-1301022032210230-0322212222320111-2332101121122303-2103200230011020-0203100121031102-1110233313120232"></a>

Type: `"object"`. single nested block, Optional.

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
k8s_cluster_role {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102311212133000-1330220313111202-0121003233023133-0010321001000331-2320310002130231-3001312300212000-1322123300203130-3212110101220030"></a>

### Direct properties for `k8s_cluster_role`

<a id="canonical-1202231323132032-3121033033101032-2002300103331003-0132020322101213-3330032012210002-3111311113000110-3220220223002013-1330031210003010"></a>

#### `k8s_cluster_role.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1333123231320120-1131001220132002-3231233001031212-0320102111313111-2012132100232120-1030103123202222-2323023301211313-0313102322001200"></a>

<a id="canonical-3203232023020132-3003200331111221-1022311222030201-3123321210211222-1131310231101020-0012003113313122-3321010201022310-3302113033212011"></a>

#### `k8s_cluster_role.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2120011031123232-1333001002011011-3123133230013212-3103311013322233-2321100300021210-3330111330133111-1003200221230331-3202101100103310"></a>

<a id="canonical-2332332112021031-0312013313303003-1201031330121020-0200301310123102-3211101302301211-3223133202011222-1312202332022010-3111210122230030"></a>

#### `k8s_cluster_role.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3202231323013321-2032223021010103-2021303133012001-3313313131013202-1232010101210021-2302132012221100-0313212310022231-0130230122331332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `subjects` properties

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-1203120300333222-3203223200312120-0002210222112313-1113120331003323-3102000331030303-1133200312133122-1031101033120033-3233030130121012)
- [Property reference](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-3122222001022132-1111202303122022-2033123222301020-0030321020302103-0030301220220131-0121213310030202-3231213321200003-3112033132100012)
- subjects

<a id="canonical-1123332223023331-3001202100103301-3110213031301313-2033300113023131-1112012030032230-3310113223100102-2012022002223233-3232233213112112"></a>

Type: `"object"`. list nested block, Optional.

List of subjects (user, group or service account) to which this role is bound.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("group",
    "service_account"),
  validators.ConflictingListObjectAttributes("group",
    "user"),
  validators.ConflictingListObjectAttributes("service_account",
    "user")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
subjects {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223110003232131-3202231311331233-0332030202011113-2000323120130203-3332232203203203-3013330132032301-0222320322312203-0133023032232301"></a>

### Direct properties for `subjects`

<a id="canonical-3001320133312121-2032300320003221-0233322002010100-1320103030221022-2023111320110231-2310320302033102-0212300300202320-1311330223333123"></a>

#### `subjects.group` property

Type: `"string"`. Optional.

Exclusive with \[service\_account user\] Group ID of the user group.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [service_account](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-1133023320132221-1202231111122312-1013102311103130-2131231013113132-1001313031131130-3201231112123021-1002133103000223-1000120230001210): complete subsection reference.

<a id="canonical-0000332030321000-3030331221030003-2002303303222233-3232121001011221-1131211300133212-0121302023302102-3232320011121111-2023200011320322"></a>

<a id="canonical-2032021103003300-0303032021312203-2203332110022010-0220110021030332-1103133310000113-1000032303130120-3100132103120030-3333200133223010"></a>

#### `subjects.user` property

Type: `"string"`. Optional.

Exclusive with \[group service\_account\] User ID of the user.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1133023320132221-1202231111122312-1013102311103130-2131231013113132-1001313031131130-3201231112123021-1002133103000223-1000120230001210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `subjects.service_account` properties

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-1203120300333222-3203223200312120-0002210222112313-1113120331003323-3102000331030303-1133200312133122-1031101033120033-3233030130121012)
- [Property reference](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-3122222001022132-1111202303122022-2033123222301020-0030321020302103-0030301220220131-0121213310030202-3231213321200003-3112033132100012)
- [subjects](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-3202231323013321-2032223021010103-2021303133012001-3313313131013202-1232010101210021-2302132012221100-0313212310022231-0130230122331332)
- subjects.service_account

<a id="canonical-3222011010230332-2032011332102303-1222222220221201-1232201331331013-2112023023332120-0203123231310312-0313020032211100-2023331113132112"></a>

Type: `"object"`. single nested block, Optional.

ServiceAccountType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name",
    "namespace")}
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
service_account {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302011120312020-3231302022102300-1323333331023122-0120010110301231-1211232110101113-2120012002321203-0210102210302120-3332011120331233"></a>

### Direct properties for `subjects.service_account`

<a id="canonical-0332103311311120-1231330020010002-3021133320002030-2120323200023312-2101222222200332-0133331322022220-1002011022300301-3003312011112322"></a>

#### `subjects.service_account.name` property

Type: `"string"`. Optional.

Name. Name of the service account.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3111020233230021-2323211331210231-0210212310113211-3313123101102002-2130030233122230-0022200133213132-0210311111230303-3100122103103213"></a>

<a id="canonical-3213303323303130-2232300011201203-3230022213120321-1201021100200130-2322120200332121-1023220222033013-1001101120232223-3003010111233120"></a>

#### `subjects.service_account.namespace` property

Type: `"string"`. Optional, Computed.

Namespace. Namespace of the service account.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3302013110231321-1022022233001010-3003220330100331-2213210110000123-0301231200201111-2302330331332300-0221300020011210-2212213201233100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-1203120300333222-3203223200312120-0002210222112313-1113120331003323-3102000331030303-1133200312133122-1031101033120033-3233030130121012)
- [Property reference](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-3122222001022132-1111202303122022-2033123222301020-0030321020302103-0030301220220131-0121213310030202-3231213321200003-3112033132100012)
- timeouts

<a id="canonical-3133310020221303-0022302310312110-1323212012012212-2222113102320123-2231131031312022-3323033023000021-0011313213223111-2010023020303031"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333001330310203-1013322131010221-1001311330120011-1110333220320332-3022313013331201-1200130123001101-0121110223010021-3013300022012312"></a>

### Direct properties for `timeouts`

<a id="canonical-2310020323000220-2331130020133100-3132121011112133-0103120021333232-2031331112102333-3302202212000132-2103213000032011-1100032320121002"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2130003220030003-0033110320320222-2021220321203121-0111111022111112-2121021131231132-2111213210032331-1000131323303311-3202113101110231"></a>

<a id="canonical-1312101121010333-1310330212112122-3202021002330133-2202131020221032-3320133221212112-3112122121001210-2030031202113223-1122221000133303"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1111201023111201-1302111211313201-2032020320032321-0220101221100223-1300111201213023-0201011202000320-3322120102232210-1323312312310323"></a>

<a id="canonical-0030323103233000-2213103211310133-1131302321112203-2011022003123310-2210003310213023-0332312313110021-3121010101130311-0010003300300313"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3032020102331112-2031332133333001-0010033320302030-3333021320203022-1020132121033113-1301301112112031-3332001112101230-0121230233223310"></a>

<a id="canonical-0232013123023320-2200010203200311-3030200322303303-0301330330021331-2033110323213203-0120112313332232-0003230312300300-1300023313211211"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
