---
page_title: "xcsh_workload_flavor reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_workload_flavor reference."
---

# xcsh_workload_flavor reference

<a id="canonical-2101210323120012-1332103220321021-0330203102320020-3302111001233322-1031112300210030-1300131022030232-2232013130123221-0310120231323221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131100022000120-2020010012303313-2103312030112323-3300002200032300-1121102331000213-1223201300313322-2002330303100230-0220223121121221"></a>

## Property reference — Property reference / 233131331121 / 2

Breadcrumbs:

- [xcsh_workload_flavor](../data-sources/workload_flavor.md#canonical-2030203000323030-0232123232021133-3020001332010013-1013132311011132-1021330332030103-1003300002310323-0333232112231113-1231333120223030)
- Property reference

<a id="canonical-0222331120210000-0330233101201233-1111133000300020-3313202001320121-2133330212202201-3022300023323033-1333122201212002-1102001021102032"></a>

## Direct properties — Property reference / 233131331121 / 3

<a id="canonical-2001021012032033-1321102310031130-2221320010133111-2221302200031133-3232133021133102-1222102232223020-2211232211231213-1213230101023221"></a>

<a id="canonical-0232301213201110-3311110212113221-0113121212110201-3100012032032012-1220213200033003-2023021003232012-1201233100323130-3001130332021201"></a>

## annotations property — Property reference / 233131331121 / 4

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

<a id="canonical-0110102111120130-3020022102113130-1011222121021331-0010132222113031-1021010222211213-2232210113132212-1203313033123333-2333031231020221"></a>

<a id="canonical-3032212211231022-0222303132011131-3003201100113111-0210210102311301-2220213321200321-2102020132131332-3332311321132112-0030113313032300"></a>

## description property — Property reference / 233131331121 / 5

Type: `"string"`. Computed.

Description of the WorkloadFlavor.

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

<a id="canonical-3011021032231322-3230300030232313-3330032222123233-3312233221223212-3302101300333102-3331232033031301-3203012231111211-1020213222013323"></a>

<a id="canonical-0330102320330030-2331102213032001-1120030320301321-0331001201101122-1312302021020103-0313320221003323-1120010131120223-1211013200031310"></a>

## ephemeral_storage property — Property reference / 233131331121 / 6

Type: `"string"`. Computed.

Ephemeral storage in MiB (mebibyte) allocated for the workload\_flavor.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.uint32.ranges": "1-6000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "1-6000"
  }
}
```

<a id="canonical-1232211210333331-1210123323332311-2123322122212011-2010330023212123-0323032002033023-1021203232301320-3221111333312102-1312123032231301"></a>

<a id="canonical-0223201032121232-2233133221033110-2133210101221303-0323301331232031-3213222233131000-0301211322330320-2300022212103232-3322310210030003"></a>

## ID property — Property reference / 233131331121 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3120013221311023-2212213021023213-1321121202303131-2030113211133002-2221103033102313-1210231330030133-2130021002333031-2120003313101120"></a>

<a id="canonical-2133310020231303-1032321021333113-0023030230110223-1322311223301101-1231110002103121-1333210023213101-3212100223331312-3230231211012332"></a>

## labels property — Property reference / 233131331121 / 8

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

<a id="canonical-0010300002200031-2203312323233012-0202003002000222-1202301331132130-2130132033132001-2232223301323303-2231331120301031-0131332102023311"></a>

<a id="canonical-2232221211100011-3021021202012031-0212322010120122-3013201131212001-1101030310011133-2103230321213122-2200200101111131-2232111021121211"></a>

## memory property — Property reference / 233131331121 / 9

Type: `"string"`. Computed.

Memory in MiB (mebibyte) allocated for the workload\_flavor.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.uint32.ranges": "1-32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "1-32768"
  }
}
```

<a id="canonical-3221320221312212-3323130201312102-2030022022310203-1320013213202011-1020100133311103-2101323302323223-3300013311133201-3132023300003202"></a>

<a id="canonical-1211030333302313-1132311111321102-1003032122002212-2033023233320001-2130202000020311-3301011031232212-2321132310213113-1102122233003013"></a>

## name property — Property reference / 233131331121 / 10

Type: `"string"`. Required.

Name of the WorkloadFlavor.

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

<a id="canonical-3331211012231213-3002032113203123-2121012302313000-2103203022312203-3233013133321033-1032313323201303-0302103233030013-3020203312020202"></a>

<a id="canonical-0122120113021333-0330031133102112-1031131300132230-1231112112123230-3201013233101102-3331103313231221-1323302211320300-1102003313230220"></a>

## namespace property — Property reference / 233131331121 / 11

Type: `"string"`. Optional, Computed.

Namespace where the WorkloadFlavor exists.

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

<a id="canonical-0001111221112000-1301322322200030-1112030223132210-1321310233320110-2302311031102131-3211003303232223-0322120013010203-2322303101020333"></a>

<a id="canonical-0223312202001112-3313130310000223-1031212200201101-2133233003003301-3331333031132020-0220310132011100-3311230123012321-3233100000202102"></a>

## vcpus property — Property reference / 233131331121 / 12

Type: `"number"`. Computed.

Number of vCPUs allocated for the workload\_flavor. Each vCPU is a thread on a CPU core.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "8.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "8.0"
  }
}
```

<a id="canonical-0030132011320300-2212332123213023-3003222110211230-1223302303103121-0311131123033322-1030230103300300-2101102223102333-3310032220312231"></a>

## All schema paths — Property reference / 233131331121 / 13

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--workload_flavor--reference--group-001.md#canonical-2001021012032033-1321102310031130-2221320010133111-2221302200031133-3232133021133102-1222102232223020-2211232211231213-1213230101023221) |
| `description` | [description](data-sources--workload_flavor--reference--group-001.md#canonical-0110102111120130-3020022102113130-1011222121021331-0010132222113031-1021010222211213-2232210113132212-1203313033123333-2333031231020221) |
| `ephemeral_storage` | [ephemeral_storage](data-sources--workload_flavor--reference--group-001.md#canonical-3011021032231322-3230300030232313-3330032222123233-3312233221223212-3302101300333102-3331232033031301-3203012231111211-1020213222013323) |
| `id` | [id](data-sources--workload_flavor--reference--group-001.md#canonical-1232211210333331-1210123323332311-2123322122212011-2010330023212123-0323032002033023-1021203232301320-3221111333312102-1312123032231301) |
| `labels` | [labels](data-sources--workload_flavor--reference--group-001.md#canonical-3120013221311023-2212213021023213-1321121202303131-2030113211133002-2221103033102313-1210231330030133-2130021002333031-2120003313101120) |
| `memory` | [memory](data-sources--workload_flavor--reference--group-001.md#canonical-0010300002200031-2203312323233012-0202003002000222-1202301331132130-2130132033132001-2232223301323303-2231331120301031-0131332102023311) |
| `name` | [name](data-sources--workload_flavor--reference--group-001.md#canonical-3221320221312212-3323130201312102-2030022022310203-1320013213202011-1020100133311103-2101323302323223-3300013311133201-3132023300003202) |
| `namespace` | [namespace](data-sources--workload_flavor--reference--group-001.md#canonical-3331211012231213-3002032113203123-2121012302313000-2103203022312203-3233013133321033-1032313323201303-0302103233030013-3020203312020202) |
| `vcpus` | [vcpus](data-sources--workload_flavor--reference--group-001.md#canonical-0001111221112000-1301322322200030-1112030223132210-1321310233320110-2302311031102131-3211003303232223-0322120013010203-2322303101020333) |

<a id="canonical-3101332322123223-3332202130122132-1200111320232110-3220231213131233-3102101210233230-0232201322112311-2300013010131123-0011300101013111"></a>

## Next pages — Property reference / 233131331121 / 14

- [xcsh_workload_flavor](../data-sources/workload_flavor.md#canonical-2030203000323030-0232123232021133-3020001332010013-1013132311011132-1021330332030103-1003300002310323-0333232112231113-1231333120223030)
