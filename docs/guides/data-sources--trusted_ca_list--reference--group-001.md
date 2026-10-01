---
page_title: "xcsh_trusted_ca_list reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_trusted_ca_list reference."
---

# xcsh_trusted_ca_list reference

<a id="canonical-3130232031111032-3303330311222223-3022110011301203-1122313200231131-2032233133020322-0030131110322322-1101000221212310-3022300311330213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323330311010012-3233123132131211-1013133030031032-1133232233323233-3200213122230103-2231003012321021-2122200110001102-0033321200132013"></a>

## Property reference — Property reference / 211111100301 / 2

Breadcrumbs:

- [xcsh_trusted_ca_list](../data-sources/trusted_ca_list.md#canonical-3231112200302210-0000022221223030-0200201111213211-3101031133123233-3133122221310202-1301130023031133-3210320322301322-1320111311102122)
- Property reference

<a id="canonical-3033331311113212-0201302310301320-3002101301112011-3323311001113021-0222031301201001-2112101103132330-2303313121000110-2321013233222113"></a>

## Direct properties — Property reference / 211111100301 / 3

<a id="canonical-0001333313303312-2303330010312023-1330223123310000-1323131200010203-1132312023102023-2303212221311003-0223113203113023-3233232323120231"></a>

<a id="canonical-3302012001023223-3002312201232332-3210111221010011-1313211313220123-3300202202030110-1322013330201003-3030022012100113-1221033200202103"></a>

## annotations property — Property reference / 211111100301 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

<a id="canonical-0330120023122311-0122232301022222-2203033200313003-3221123020221032-1310303132102001-2223310023132232-3112013221110122-3003332233022120"></a>

<a id="canonical-2132022001210310-0130201220212113-1000332211032112-0320333003230302-0001220111102112-1122211021310303-0022120111023000-0022330321110220"></a>

## description property — Property reference / 211111100301 / 5

Type: `"string"`. Computed.

Description of the TrustedCAList.

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

<a id="canonical-1102230212213003-3233330203321031-1033232311220121-2330321230323230-2231323220321002-0020222102010231-0322312122221210-0002101210302330"></a>

<a id="canonical-2220233133200210-1011012210013130-0021311100221322-1001302230302310-0312210323213010-0110223203333202-1032130301211102-2330211232323100"></a>

## ID property — Property reference / 211111100301 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0133200331100302-0231211312200331-2132131132313112-1220021022310221-1332100212101221-2011003231223003-2103031121030120-3100133300203320"></a>

<a id="canonical-2103120133222033-2110000220110033-3112003003033101-0332312333232300-3133002130122021-3001112331231033-1132230221032223-3011302021303003"></a>

## labels property — Property reference / 211111100301 / 7

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

<a id="canonical-3312313131231023-3013301103210211-3202101022230320-3021333021322122-1313021323332133-0133211203023200-3330332223211021-2312030330303302"></a>

<a id="canonical-3230212123300030-0100330211311102-1033330321000320-3310312211332101-0330323011021223-2202131110330311-3203312032112030-3311211302203320"></a>

## name property — Property reference / 211111100301 / 8

Type: `"string"`. Required.

Name of the TrustedCAList.

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

<a id="canonical-2002132121112131-3223132212320212-1023221302300033-0002020113202332-1302331021123110-2230013130103331-2013110213222122-0020203200301030"></a>

<a id="canonical-3332021333220313-2032203213311321-1311111220101110-1331121133010130-1223301303203120-1322321101210222-1333131011022323-2322213201231213"></a>

## namespace property — Property reference / 211111100301 / 9

Type: `"string"`. Required.

Namespace where the TrustedCAList exists.

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

<a id="canonical-0202110011010332-1110122023123201-0021300333330220-1310130122231211-1031102222000233-0201013030312303-1223130321300221-1323232202212302"></a>

<a id="canonical-2110013223301013-2321320032000133-2011233000201001-1320020331311301-0223021320032101-0102020313011311-0222211022022220-2222031323102102"></a>

## trusted_ca_url property — Property reference / 211111100301 / 10

Type: `"string"`. Computed.

Trusted CA certificates for validating certificates.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512000",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512000",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-1020233023100120-0133101000101013-1312102032113223-2103101321313032-0111121233223221-0101310311231223-3333233231003230-1031011310030220"></a>

## All schema paths — Property reference / 211111100301 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--trusted_ca_list--reference--group-001.md#canonical-0001333313303312-2303330010312023-1330223123310000-1323131200010203-1132312023102023-2303212221311003-0223113203113023-3233232323120231) |
| `description` | [description](data-sources--trusted_ca_list--reference--group-001.md#canonical-0330120023122311-0122232301022222-2203033200313003-3221123020221032-1310303132102001-2223310023132232-3112013221110122-3003332233022120) |
| `id` | [id](data-sources--trusted_ca_list--reference--group-001.md#canonical-1102230212213003-3233330203321031-1033232311220121-2330321230323230-2231323220321002-0020222102010231-0322312122221210-0002101210302330) |
| `labels` | [labels](data-sources--trusted_ca_list--reference--group-001.md#canonical-0133200331100302-0231211312200331-2132131132313112-1220021022310221-1332100212101221-2011003231223003-2103031121030120-3100133300203320) |
| `name` | [name](data-sources--trusted_ca_list--reference--group-001.md#canonical-3312313131231023-3013301103210211-3202101022230320-3021333021322122-1313021323332133-0133211203023200-3330332223211021-2312030330303302) |
| `namespace` | [namespace](data-sources--trusted_ca_list--reference--group-001.md#canonical-2002132121112131-3223132212320212-1023221302300033-0002020113202332-1302331021123110-2230013130103331-2013110213222122-0020203200301030) |
| `trusted_ca_url` | [trusted_ca_url](data-sources--trusted_ca_list--reference--group-001.md#canonical-0202110011010332-1110122023123201-0021300333330220-1310130122231211-1031102222000233-0201013030312303-1223130321300221-1323232202212302) |

<a id="canonical-0130331032100031-3201232110223021-0313331120200312-0033230022313121-2000012311220212-2033330212031201-3103200201221123-3210300300210331"></a>

## Next pages — Property reference / 211111100301 / 12

- [xcsh_trusted_ca_list](../data-sources/trusted_ca_list.md#canonical-3231112200302210-0000022221223030-0200201111213211-3101031133123233-3133122221310202-1301130023031133-3210320322301322-1320111311102122)
