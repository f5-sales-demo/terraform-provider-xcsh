---
page_title: "xcsh_dns_compliance_checks reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_compliance_checks reference."
---

# xcsh_dns_compliance_checks reference

<a id="canonical-3c742fd4796a81209dac248974569ee528b2f8284b59ae8fbf6e8838bda5854f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09f521b10ca1ebb089e4a3eed5430a8583e289ce715dec56bf5b0ffe776824b9"></a>

## Property reference — Property reference / 64dadf66a298 / 2

Breadcrumbs:

- [xcsh_dns_compliance_checks](../data-sources/dns_compliance_checks.md#canonical-0df8fdb3734c554a4a4bef589f19acd0c3a88fa45c35acb7d5923f8e5221913a)
- Property reference

<a id="canonical-af88d67060a354f5081e3b51df4d7b3e3976c4f9e9e378857705b037a3dfd5e0"></a>

## Direct properties — Property reference / 64dadf66a298 / 3

<a id="canonical-3d9513e241127cb7d4cbdb4c6ee6c8a3bf3c7885cd1c22a46e62d2ecdb69b393"></a>

<a id="canonical-80bf7e7995d19d8a0a831e0eb17c94e5c85a77e2837b3f985e929f6f6f003f28"></a>

## annotations property — Property reference / 64dadf66a298 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
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

<a id="canonical-ec99604106b35d75ecf347d20c5c0ebd255d96138c12f84d16f43f2727f4c70a"></a>

<a id="canonical-eda2b481a7f33c26c968cedf1c912c2f9f733f9269dc6f3fb649cef2478174bc"></a>

## description property — Property reference / 64dadf66a298 / 5

Type: `"string"`. Computed.

Description of the DNSComplianceChecks.

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

<a id="canonical-e426fcb3265c2c0e2295d72cf862a566eeb57bd0de7621385a807cd9f2f44986"></a>

<a id="canonical-362f4ad81acc2aecd9a34fea43bd7b64c0c5d29d24c653c74bc9e8d1713c1ac4"></a>

## disallowed_query_type_list property — Property reference / 64dadf66a298 / 6

Type: `["list", "string"]`. Computed.

\[Enum: QUERY|IQUERY|STATUS|NOTIFY|UPDATE\] Disallowed Query Type Values. Disallowed Query Type
Values. Possible values are \`QUERY\`, \`IQUERY\`, \`STATUS\`, \`NOTIFY\`, \`UPDATE\`. Defaults to
\`QUERY\`.

Upstream description:

Disallowed Query Type Values.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3c5d7a1c7d9b130eeefc7c99100ae65ecb5feef8147516516f02dbd59fe26bb6"></a>

<a id="canonical-5e051e65bcb688d57a698fd6abedeca4c254776157912b1370627dbc20e6bfba"></a>

## disallowed_resource_record_type_list property — Property reference / 64dadf66a298 / 7

Type: `["list", "string"]`. Computed.

\[Enum:
T|A|NS|MD|MF|CNAME|SOA|MB|MG|MR|NULL|WKS|PTR|HINFO|MINFO|MX|TXT|RP|AFSDB|X25|ISDN|RT|NSAP|NSAP\_PTR|SIG|KEY|PX|GPOS|AAAA|LOC|NXT|EID|NIMLOC|SRV|ATMA|NAPTR|KX|CERT|A6|DNAME|SINK|OPT|APL|DS|SSHFP|IPSECKEY|RRSIG|NSEC|DNSKEY|DHCID|NSEC3|NSEC3PARAM|TLSA|SMIMEA|HIP|NINFO|RKEY|TALINK|CDS|CDNSKEY|OPENPGPKEY|CSYNC|SPF|UINFO|UID|GID|UNSPEC|NID|L32|L64|LP|EUI48|EUI64|TKEY|TSIG|IXFR|AXFR|MAILB|MAILA|URI|CAA|TA|DLV\]
Disallowed Resource Record Types. Disallowed Resource Record Type List. Possible values are \`T\`,
\`A\`, \`NS\`, \`MD\`, \`MF\`, \`CNAME\`, \`SOA\`, \`MB\`, \`MG\`, \`MR\`, \`NULL\`, \`WKS\`,
\`PTR\`, \`HINFO\`, \`MINFO\`, \`MX\`, \`TXT\`, \`RP\`, \`AFSDB\`, \`X25\`, \`ISDN\`, \`RT\`,
\`NSAP\`, \`NSAP\_PTR\`, \`SIG\`, \`KEY\`, \`PX\`, \`GPOS\`, \`AAAA\`, \`LOC\`, \`NXT\`, \`EID\`,
\`NIMLOC\`, \`SRV\`, \`ATMA\`, \`NAPTR\`, \`KX\`, \`CERT\`, \`A6\`, \`DNAME\`, \`SINK\`, \`OPT\`,
\`APL\`, \`DS\`, \`SSHFP\`, \`IPSECKEY\`, \`RRSIG\`, \`NSEC\`, \`DNSKEY\`, \`DHCID\`, \`NSEC3\`,
\`NSEC3PARAM\`, \`TLSA\`, \`SMIMEA\`, \`HIP\`, \`NINFO\`, \`RKEY\`, \`TALINK\`, \`CDS\`,
\`CDNSKEY\`, \`OPENPGPKEY\`, \`CSYNC\`, \`SPF\`, \`UINFO\`, \`UID\`, \`GID\`, \`UNSPEC\`, \`NID\`,
\`L32\`, \`L64\`, \`LP\`, \`EUI48\`, \`EUI64\`, \`TKEY\`, \`TSIG\`, \`IXFR\`, \`AXFR\`, \`MAILB\`,
\`MAILA\`, \`URI\`, \`CAA\`, \`TA\`, \`DLV\`. Defaults to \`T\`.

Upstream description:

Disallowed Resource Record Type List.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-29365bb7ac1e4f1caf02502a18c4b33fe0d646e0e3054ab1938067c076305a48"></a>

<a id="canonical-d05821929bcc5e8644819f4c0dc80906de666ba1810361b0549af01f98381d00"></a>

## domain_denylist property — Property reference / 64dadf66a298 / 8

Type: `["list", "string"]`. Computed.

List of domains to be denied by configuration object.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 30,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.etld_plus_one": "true",
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.etld_plus_one": "true",
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1cd84cb4c004c4ef2e8db2ca8364355f50ad961ed659640d4be6d32ee88c10ea"></a>

<a id="canonical-96b7e91e9ffbc85bcdf0ca5ec915c2904f9fd110430772cb3da988765f42611f"></a>

## id property — Property reference / 64dadf66a298 / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-206d8e5e4b8083d7aab50a479d68ec3a1cc07646cabd7edbb0cc21a9b83c4d81"></a>

<a id="canonical-41564bf0b19adea3b96b3b3a1b5e1ef0317fd14bfc1bdb91e78273ae7ea46380"></a>

## labels property — Property reference / 64dadf66a298 / 10

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

<a id="canonical-8b21fb55468ad2a479dad79a944958ed99e0b210ef25f3c7ba135f38116ea590"></a>

<a id="canonical-b11e14247b33fdbb71c5a0d106511521f21bf656db17b15f1b5a606967a3b825"></a>

## name property — Property reference / 64dadf66a298 / 11

Type: `"string"`. Required.

Name of the DNSComplianceChecks.

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

<a id="canonical-99074b24031f81f85907b059b3f00c82932ca808e33df234df5c8e8d8deece7b"></a>

<a id="canonical-9322534731b16069a4b45a40f63c8ec9baaf97bc313eddcd958835151577488c"></a>

## namespace property — Property reference / 64dadf66a298 / 12

Type: `"string"`. Required.

Namespace where the DNSComplianceChecks exists.

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

<a id="canonical-6535cace43a02dd5d4b1841619b49dd4ae1c06a36956133f00f8c1e4d1e580d0"></a>

## All schema paths — Property reference / 64dadf66a298 / 13

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--dns_compliance_checks--reference--group-001.md#canonical-3d9513e241127cb7d4cbdb4c6ee6c8a3bf3c7885cd1c22a46e62d2ecdb69b393) |
| `description` | [description](data-sources--dns_compliance_checks--reference--group-001.md#canonical-ec99604106b35d75ecf347d20c5c0ebd255d96138c12f84d16f43f2727f4c70a) |
| `disallowed_query_type_list` | [disallowed_query_type_list](data-sources--dns_compliance_checks--reference--group-001.md#canonical-e426fcb3265c2c0e2295d72cf862a566eeb57bd0de7621385a807cd9f2f44986) |
| `disallowed_resource_record_type_list` | [disallowed_resource_record_type_list](data-sources--dns_compliance_checks--reference--group-001.md#canonical-3c5d7a1c7d9b130eeefc7c99100ae65ecb5feef8147516516f02dbd59fe26bb6) |
| `domain_denylist` | [domain_denylist](data-sources--dns_compliance_checks--reference--group-001.md#canonical-29365bb7ac1e4f1caf02502a18c4b33fe0d646e0e3054ab1938067c076305a48) |
| `id` | [id](data-sources--dns_compliance_checks--reference--group-001.md#canonical-1cd84cb4c004c4ef2e8db2ca8364355f50ad961ed659640d4be6d32ee88c10ea) |
| `labels` | [labels](data-sources--dns_compliance_checks--reference--group-001.md#canonical-206d8e5e4b8083d7aab50a479d68ec3a1cc07646cabd7edbb0cc21a9b83c4d81) |
| `name` | [name](data-sources--dns_compliance_checks--reference--group-001.md#canonical-8b21fb55468ad2a479dad79a944958ed99e0b210ef25f3c7ba135f38116ea590) |
| `namespace` | [namespace](data-sources--dns_compliance_checks--reference--group-001.md#canonical-99074b24031f81f85907b059b3f00c82932ca808e33df234df5c8e8d8deece7b) |

<a id="canonical-1da64e8427a195d93fe215d15aa1dbdc54cd32fe9cc52f1e747afece20763d28"></a>

## Next pages — Property reference / 64dadf66a298 / 14

- [xcsh_dns_compliance_checks](../data-sources/dns_compliance_checks.md#canonical-0df8fdb3734c554a4a4bef589f19acd0c3a88fa45c35acb7d5923f8e5221913a)
