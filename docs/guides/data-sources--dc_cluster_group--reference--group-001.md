---
page_title: "xcsh_dc_cluster_group reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dc_cluster_group reference."
---

# xcsh_dc_cluster_group reference

<a id="canonical-1010333130230311-1033232300311001-0022220103320030-1332313110202221-3021223120002030-2000202203133231-3233130200102213-2123030302321023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md#canonical-2200021000201211-3130032310312033-3013132131200022-0333300021021222-1220133033212100-0102010030313222-0332101231122310-0203313010012313)
- Property reference

<a id="canonical-1310111302122003-0132033000023233-2031302333031312-2332121321302000-1232200220320200-1332330230003111-2033210332102102-3301032223232010"></a>

### Direct properties for `xcsh_dc_cluster_group`

<a id="canonical-1201311221232131-1330303200131110-2102023123132212-1303301121131102-1101111220102120-0323000311112131-1121322220310311-0023123320330322"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

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

<a id="canonical-3310120022111223-2233231223030010-1123022122222131-1000023330113223-0201010222222321-1323120201233030-2111111312211000-2021201011233133"></a>

<a id="canonical-0022010211132213-3000313310311120-2230002312012311-2020012100323223-2132113130023022-1020012211113333-0322333012300220-0331303221112203"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the DcClusterGroup.

Additional upstream details:

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1332230023230310-3113301221030323-1330103310320111-3213213323330100-2013123232112303-2023330220330012-1020121033203310-0300010220213210"></a>

<a id="canonical-2230103032230113-2011002212012002-3110031312112301-0120032303321002-2022023131021330-1030012123133021-0131022203330322-2032231303221220"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0213211222201222-3033212100312022-1001110003123311-1301331022312233-2222112102133130-0030201310300021-3110303102310123-2103112302112011"></a>

<a id="canonical-1123031131121323-0222320013032032-3310321033210230-2000102020120200-3003303030311210-2223003302322131-2212223123221311-3320112132213100"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-3333203123201311-3310323101111030-3001121102332102-1213302311312332-3221011321032103-1300001212330323-3031331331312030-0222102012113332"></a>

<a id="canonical-1203001322313233-0320123000122301-2203103122332320-2030321131010132-1230123221203010-3030212103102120-0032122110230031-0230332201100020"></a>

#### `name` property

Type: `"string"`. Required.

Name of the DcClusterGroup.

Additional upstream details:

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3010331313010211-0121213300123232-0300232002103323-1231032122313121-3033310012113300-0230300003032211-3330022030301332-1100103112230032"></a>

<a id="canonical-2333022233033032-0311010312333111-1230130032211020-3023033323113122-3201101313301213-0231031010002231-2022312000212030-0212110002120323"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace where the DcClusterGroup exists.

Additional upstream details:

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [type](data-sources--dc_cluster_group--reference--group-001.md#canonical-2121331020332121-1223023231322301-0313101210122203-0120122211011202-2000020120123123-1221321122023132-0133202212221032-0322033003200322): complete subsection reference.

<a id="canonical-2013013130012122-0323011023100130-1002010012012113-2011013310310322-1103211221211023-1333132001333310-3101310201003101-2231111010213233"></a>

### All schema paths for `xcsh_dc_cluster_group`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--dc_cluster_group--reference--group-001.md#canonical-1201311221232131-1330303200131110-2102023123132212-1303301121131102-1101111220102120-0323000311112131-1121322220310311-0023123320330322) |
| `description` | [description](data-sources--dc_cluster_group--reference--group-001.md#canonical-3310120022111223-2233231223030010-1123022122222131-1000023330113223-0201010222222321-1323120201233030-2111111312211000-2021201011233133) |
| `id` | [ID](data-sources--dc_cluster_group--reference--group-001.md#canonical-1332230023230310-3113301221030323-1330103310320111-3213213323330100-2013123232112303-2023330220330012-1020121033203310-0300010220213210) |
| `labels` | [labels](data-sources--dc_cluster_group--reference--group-001.md#canonical-0213211222201222-3033212100312022-1001110003123311-1301331022312233-2222112102133130-0030201310300021-3110303102310123-2103112302112011) |
| `name` | [name](data-sources--dc_cluster_group--reference--group-001.md#canonical-3333203123201311-3310323101111030-3001121102332102-1213302311312332-3221011321032103-1300001212330323-3031331331312030-0222102012113332) |
| `namespace` | [namespace](data-sources--dc_cluster_group--reference--group-001.md#canonical-3010331313010211-0121213300123232-0300232002103323-1231032122313121-3033310012113300-0230300003032211-3330022030301332-1100103112230032) |
| `type` | [type](data-sources--dc_cluster_group--reference--group-001.md#canonical-3111022103230330-1221032201222100-0123130331103330-3200313022022212-3231033322031301-1213112210011003-1132121121101210-3322113003332220) |
| `type.control_and_data_plane_mesh` | [type.control_and_data_plane_mesh](data-sources--dc_cluster_group--reference--group-001.md#canonical-2123011302320013-1020322221302031-0032221210203212-1023111333021013-2013322330032021-0303012223223332-2333132212301123-0011030203120100) |
| `type.data_plane_mesh` | [type.data_plane_mesh](data-sources--dc_cluster_group--reference--group-001.md#canonical-0210002210311221-0331223303123112-1010310233320131-2212120300211010-2323123300033011-0121202011102300-1120020213202011-2133113233122011) |

<a id="canonical-2121331020332121-1223023231322301-0313101210122203-0120122211011202-2000020120123123-1221321122023132-0133202212221032-0322033003200322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `type` properties

Breadcrumbs:

- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md#canonical-2200021000201211-3130032310312033-3013132131200022-0333300021021222-1220133033212100-0102010030313222-0332101231122310-0203313010012313)
- [Property reference](data-sources--dc_cluster_group--reference--group-001.md#canonical-1010333130230311-1033232300311001-0022220103320030-1332313110202221-3021223120002030-2000202203133231-3233130200102213-2123030302321023)
- type

<a id="canonical-3111022103230330-1221032201222100-0123130331103330-3200313022022212-3231033322031301-1213112210011003-1132121121101210-3322113003332220"></a>

Type: `"single"`. Computed.

DC Cluster Group Mesh Type. Details of DC Cluster Group Mesh Type.

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

<a id="canonical-1122020103123211-0030310322012321-1122101102003021-2010022033122301-3201312012031120-0120223111202110-1120203300102102-1003113320010021"></a>

### Direct properties for `type`

- [control_and_data_plane_mesh](data-sources--dc_cluster_group--reference--group-001.md#canonical-1321131100211131-0121313233322321-0231331321021200-1112030121313003-1201023122033031-2322002322231323-3232023122101102-3113123100113330): complete subsection reference.

- [data_plane_mesh](data-sources--dc_cluster_group--reference--group-001.md#canonical-1130231102003213-1232030010221320-2120333122000310-3313322102232211-0110030012121301-1010121310113320-2300110301123113-3021033113231212): complete subsection reference.

<a id="canonical-1321131100211131-0121313233322321-0231331321021200-1112030121313003-1201023122033031-2322002322231323-3232023122101102-3113123100113330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `type.control_and_data_plane_mesh` properties

Breadcrumbs:

- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md#canonical-2200021000201211-3130032310312033-3013132131200022-0333300021021222-1220133033212100-0102010030313222-0332101231122310-0203313010012313)
- [Property reference](data-sources--dc_cluster_group--reference--group-001.md#canonical-1010333130230311-1033232300311001-0022220103320030-1332313110202221-3021223120002030-2000202203133231-3233130200102213-2123030302321023)
- [type](data-sources--dc_cluster_group--reference--group-001.md#canonical-2121331020332121-1223023231322301-0313101210122203-0120122211011202-2000020120123123-1221321122023132-0133202212221032-0322033003200322)
- type.control_and_data_plane_mesh

<a id="canonical-2123011302320013-1020322221302031-0032221210203212-1023111333021013-2013322330032021-0303012223223332-2333132212301123-0011030203120100"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130231102003213-1232030010221320-2120333122000310-3313322102232211-0110030012121301-1010121310113320-2300110301123113-3021033113231212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `type.data_plane_mesh` properties

Breadcrumbs:

- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md#canonical-2200021000201211-3130032310312033-3013132131200022-0333300021021222-1220133033212100-0102010030313222-0332101231122310-0203313010012313)
- [Property reference](data-sources--dc_cluster_group--reference--group-001.md#canonical-1010333130230311-1033232300311001-0022220103320030-1332313110202221-3021223120002030-2000202203133231-3233130200102213-2123030302321023)
- [type](data-sources--dc_cluster_group--reference--group-001.md#canonical-2121331020332121-1223023231322301-0313101210122203-0120122211011202-2000020120123123-1221321122023132-0133202212221032-0322033003200322)
- type.data_plane_mesh

<a id="canonical-0210002210311221-0331223303123112-1010310233320131-2212120300211010-2323123300033011-0121202011102300-1120020213202011-2133113233122011"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.
