---
page_title: "xcsh_mitigated_domain reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_mitigated_domain reference."
---

# xcsh_mitigated_domain reference

<a id="canonical-0201213223313332-2132200331130222-1202220001302010-1022310120223013-3020032223131202-2123112211203031-2213120030001221-2330123113023202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222312200031202-0130223102011011-3000020002023110-3022012213003300-0311200030332023-0123012303121231-1301220232311120-1330002202322101"></a>

## Property reference — Property reference / 101212202230 / 2

Breadcrumbs:

- [xcsh_mitigated_domain](../data-sources/mitigated_domain.md#canonical-3233303323320020-3201003222322030-3103103212101100-1321230323023131-3310220303032323-0211212333301333-1330120302113020-2012202211332310)
- Property reference

<a id="canonical-0222233020233210-2320212312200030-0332112210210333-0232111120013311-1111231310131003-3020121223201013-3323312010233023-3030302302000001"></a>

## Direct properties — Property reference / 101212202230 / 3

<a id="canonical-1022013012300330-1202232210111330-1132102313100301-1000310213121001-0332320113302310-3301132310211131-2302011121213203-2001121200210220"></a>

<a id="canonical-3200210130332010-3000221022121213-0230020013332010-0113202030002312-0132030322132131-1100013130130123-2302132020310121-1303313001203033"></a>

## annotations property — Property reference / 101212202230 / 4

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

<a id="canonical-0131232231022000-0020020110121102-1300011231101230-1310000320230001-0001210310132300-2313020112300323-3100222333331200-2233222021211113"></a>

<a id="canonical-3011322023033201-2312220232313123-0021211223122030-2131230000322231-2323312023122120-3020301020303103-0103123223310122-0333222031010032"></a>

## description property — Property reference / 101212202230 / 5

Type: `"string"`. Computed.

Description of the MitigatedDomain.

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

<a id="canonical-0313113100011131-0322232233020330-1321203020003330-1220202013310223-2321303003023001-0010002211332220-1031003213113100-2231202133002122"></a>

<a id="canonical-3333020022010222-0020231110122013-3203212033203120-2001032230010020-3120110100021031-0300312021223032-0321001122112023-3330023333233020"></a>

## ID property — Property reference / 101212202230 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1303021133310030-3022132102103212-3200222200303103-1333103330311112-1221323222021121-1230220302003002-0023211333311313-2212213232300120"></a>

<a id="canonical-1202333101002033-2322213020323130-0320233133222030-3013310231310121-1332203122133200-1322222122213322-3223211212302101-2021113201333030"></a>

## labels property — Property reference / 101212202230 / 7

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

<a id="canonical-0201221231012031-1003031112302123-0000212301230122-1321022112221212-2210131322320030-2212123302130012-3020021122221322-0121013133123221"></a>

<a id="canonical-1323122232110232-1232122210020003-3032121201010112-2103233011110210-1302001112331221-0002220131302112-0220031030232222-1300201032210201"></a>

## mitigated_domain property — Property reference / 101212202230 / 8

Type: `"string"`. Computed.

Enter root domain or domain to be entered to mitigated list below. Domains can be entered only one
at a time. In case of conflicting entries, the domain entry takes precedence over the root domain
entry.

Upstream description:

Enter root domain or domain to be entered to mitigated list below. Domains can be entered only one
at a time. In case of conflicting entries, the domain entry takes precedence over the root domain
entry. Example: if you are adding Client-Side Defense JS on checkout.example.com, you should enter
example.com here.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-0023222331333011-2210312121232122-1011112200320210-3032102131031023-2121111100111133-0010300322012200-2310032320212321-1101120231013001"></a>

<a id="canonical-2323332222222030-2112101012131301-3002132113313132-0112013300220320-3233003132310013-1122003331300013-3312031031312203-1111112110010010"></a>

## name property — Property reference / 101212202230 / 9

Type: `"string"`. Required.

Name of the MitigatedDomain.

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

<a id="canonical-0110113023121203-3122022002300110-2011000111000322-0021220003002323-1320121001013203-3000203032023110-2121032112231211-2320130323132312"></a>

<a id="canonical-2222212230132000-3013030133202001-1300001012211012-1323132303222331-3001033012221303-1111222310223212-0232000021030323-0301023122133101"></a>

## namespace property — Property reference / 101212202230 / 10

Type: `"string"`. Required.

Namespace where the MitigatedDomain exists.

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

<a id="canonical-1300332110122332-1021101313112103-2010111332133210-2223131232131233-0312201002320112-0130032021320102-1023011331003213-3330031012301330"></a>

## All schema paths — Property reference / 101212202230 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--mitigated_domain--reference--group-001.md#canonical-1022013012300330-1202232210111330-1132102313100301-1000310213121001-0332320113302310-3301132310211131-2302011121213203-2001121200210220) |
| `description` | [description](data-sources--mitigated_domain--reference--group-001.md#canonical-0131232231022000-0020020110121102-1300011231101230-1310000320230001-0001210310132300-2313020112300323-3100222333331200-2233222021211113) |
| `id` | [ID](data-sources--mitigated_domain--reference--group-001.md#canonical-0313113100011131-0322232233020330-1321203020003330-1220202013310223-2321303003023001-0010002211332220-1031003213113100-2231202133002122) |
| `labels` | [labels](data-sources--mitigated_domain--reference--group-001.md#canonical-1303021133310030-3022132102103212-3200222200303103-1333103330311112-1221323222021121-1230220302003002-0023211333311313-2212213232300120) |
| `mitigated_domain` | [mitigated_domain](data-sources--mitigated_domain--reference--group-001.md#canonical-0201221231012031-1003031112302123-0000212301230122-1321022112221212-2210131322320030-2212123302130012-3020021122221322-0121013133123221) |
| `name` | [name](data-sources--mitigated_domain--reference--group-001.md#canonical-0023222331333011-2210312121232122-1011112200320210-3032102131031023-2121111100111133-0010300322012200-2310032320212321-1101120231013001) |
| `namespace` | [namespace](data-sources--mitigated_domain--reference--group-001.md#canonical-0110113023121203-3122022002300110-2011000111000322-0021220003002323-1320121001013203-3000203032023110-2121032112231211-2320130323132312) |

<a id="canonical-2320330233200032-1101333020213100-2102322203121200-0231103130133121-0201301303131033-3010302113010022-1020031120230120-1123131200220012"></a>

## Next pages — Property reference / 101212202230 / 12

- [xcsh_mitigated_domain](../data-sources/mitigated_domain.md#canonical-3233303323320020-3201003222322030-3103103212101100-1321230323023131-3310220303032323-0211212333301333-1330120302113020-2012202211332310)
