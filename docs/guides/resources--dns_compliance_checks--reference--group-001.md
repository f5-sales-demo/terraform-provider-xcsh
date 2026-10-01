---
page_title: "xcsh_dns_compliance_checks reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_compliance_checks reference."
---

# xcsh_dns_compliance_checks reference

<a id="canonical-eee1b4d710d52db719c841deb6e007376a78ef3ab29c2e2a6fc1e96dd6936c22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-531b29a3bb66847000cfe381b06714107e09358a6cb38a761a2936c9d7438f6e"></a>

## Property reference — Property reference / 624d8fc3618d / 2

Breadcrumbs:

- [xcsh_dns_compliance_checks](../resources/dns_compliance_checks.md#canonical-b8f0d3ec4e785ef6fabb81b92471aa91dd33fcda47717d8aa7b007104462e924)
- Property reference

<a id="canonical-77b9e804918db492a0f51fd4531238b9d96f4dd793464ee044ad431181856937"></a>

## Direct properties — Property reference / 624d8fc3618d / 3

<a id="canonical-93d27a5826576e14ec3c5617efaf2f89820c852b408b1214052d12cdd02f0c83"></a>

<a id="canonical-dc868061b708a26117119b10b1c70a01ce2ab2d74ee31da0cf79ea1e19a048cd"></a>

## annotations property — Property reference / 624d8fc3618d / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

<a id="canonical-1d13da16474200882f50cc9181c75e81ff3705474df72b70828cd2663e97d028"></a>

<a id="canonical-4f803fe1ab82a59ede8250abcbef5e2913cd205a13386a3ad185a45c9baf7791"></a>

## description property — Property reference / 624d8fc3618d / 5

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

<a id="canonical-79001ca5d59adb4994117a80f3860df0ef4bce89231e332dd3f3a6e52b254303"></a>

<a id="canonical-9526a99c763eb0d77c924c82c9fa4a095cd421c21b7fc7f5c4138372db5f95ba"></a>

## disable property — Property reference / 624d8fc3618d / 6

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

<a id="canonical-1ecde88e72da056068e9969d8054b688cead8d9f4131a751bf372e5daf1945a7"></a>

<a id="canonical-9940fee869f99904aa139ae0b69c5354fc4c515ac5077c66a8a956ad5bc6778b"></a>

## disallowed_query_type_list property — Property reference / 624d8fc3618d / 7

Type: `["list", "string"]`. Optional.

\[Enum: QUERY|IQUERY|STATUS|NOTIFY|UPDATE\] Disallowed Query Type Values. Disallowed Query Type
Values. Possible values are \`QUERY\`, \`IQUERY\`, \`STATUS\`, \`NOTIFY\`, \`UPDATE\`. Defaults to
\`QUERY\`.

Upstream description:

Disallowed Query Type Values.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(8),
}
```

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

<a id="canonical-5f31a64ec3f22c209554e11443d4265e92b5ae1e5cc303e60ae3f0cd2f66e58b"></a>

<a id="canonical-78d91c1b697efc0ded3460112f287e917c4777367597482a6f8f7f1dbbfe43e1"></a>

## disallowed_resource_record_type_list property — Property reference / 624d8fc3618d / 8

Type: `["list", "string"]`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(8),
}
```

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

<a id="canonical-dab1eeb088a0e0a909b04c7f7c5ee1eb4b46b8aec4c82608ee0ac1f694e605f7"></a>

<a id="canonical-fb0eeb4b661641bdcd3b67c4d696e0dff841578a4354c38c32233ee4e96c57f7"></a>

## domain_denylist property — Property reference / 624d8fc3618d / 9

Type: `["list", "string"]`. Required.

List of domains to be denied by configuration object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(30),
}
```

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

<a id="canonical-70f8dbe57612019d7d03019274b78ee52cf27ff4a6e4a9b2749a9759e78c0ebe"></a>

<a id="canonical-e94df1b9b8c8dff7bd5e8c86a116e2311e583ea705b63f04c600f4238aa9791a"></a>

## id property — Property reference / 624d8fc3618d / 10

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-d75ea6097dffa5d88f17bd311dead8d152bcac6819bac3cc7019b68cdd8e4022"></a>

<a id="canonical-498f2f6f68cf5177a6a89f606bd55f9358215f3df614b2a91392a8afd2f79cb4"></a>

## labels property — Property reference / 624d8fc3618d / 11

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
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

<a id="canonical-4ec67d188842d32a3e133816fe4d988faf16d2a06054e8fb63c2308dd4ce2b7a"></a>

<a id="canonical-87de2a5c62f825929163d14c98d5576fadc10f757ed1d29f9b6a6921d08680e9"></a>

## name property — Property reference / 624d8fc3618d / 12

Type: `"string"`. Required.

Name of the DNS Compliance Checks. Must be unique within the namespace.

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

<a id="canonical-3613f41cae02c317076df61b7f5975acca0775722f05572b9c8b4e822cd6a75d"></a>

<a id="canonical-a17a67f03cf498cfb3ea34009f71bfed071c10b0707841fb9bafc0aa7502d25f"></a>

## namespace property — Property reference / 624d8fc3618d / 13

Type: `"string"`. Required.

Namespace where the DNS Compliance Checks is created.

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

- [timeouts](resources--dns_compliance_checks--reference--group-001.md#canonical-199d8bdf76f2f46e7a37feea668199530aa885f9d9764fdec7c86e3e6068ba1a): complete subsection reference.

<a id="canonical-276604b922223b238f479b47f464cd88b6046640fe03818ce796c2f0c67a32a7"></a>

## All schema paths — Property reference / 624d8fc3618d / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--dns_compliance_checks--reference--group-001.md#canonical-93d27a5826576e14ec3c5617efaf2f89820c852b408b1214052d12cdd02f0c83) |
| `description` | [description](resources--dns_compliance_checks--reference--group-001.md#canonical-1d13da16474200882f50cc9181c75e81ff3705474df72b70828cd2663e97d028) |
| `disable` | [disable](resources--dns_compliance_checks--reference--group-001.md#canonical-79001ca5d59adb4994117a80f3860df0ef4bce89231e332dd3f3a6e52b254303) |
| `disallowed_query_type_list` | [disallowed_query_type_list](resources--dns_compliance_checks--reference--group-001.md#canonical-1ecde88e72da056068e9969d8054b688cead8d9f4131a751bf372e5daf1945a7) |
| `disallowed_resource_record_type_list` | [disallowed_resource_record_type_list](resources--dns_compliance_checks--reference--group-001.md#canonical-5f31a64ec3f22c209554e11443d4265e92b5ae1e5cc303e60ae3f0cd2f66e58b) |
| `domain_denylist` | [domain_denylist](resources--dns_compliance_checks--reference--group-001.md#canonical-dab1eeb088a0e0a909b04c7f7c5ee1eb4b46b8aec4c82608ee0ac1f694e605f7) |
| `id` | [id](resources--dns_compliance_checks--reference--group-001.md#canonical-70f8dbe57612019d7d03019274b78ee52cf27ff4a6e4a9b2749a9759e78c0ebe) |
| `labels` | [labels](resources--dns_compliance_checks--reference--group-001.md#canonical-d75ea6097dffa5d88f17bd311dead8d152bcac6819bac3cc7019b68cdd8e4022) |
| `name` | [name](resources--dns_compliance_checks--reference--group-001.md#canonical-4ec67d188842d32a3e133816fe4d988faf16d2a06054e8fb63c2308dd4ce2b7a) |
| `namespace` | [namespace](resources--dns_compliance_checks--reference--group-001.md#canonical-3613f41cae02c317076df61b7f5975acca0775722f05572b9c8b4e822cd6a75d) |
| `timeouts` | [timeouts](resources--dns_compliance_checks--reference--group-001.md#canonical-98a021adb39ea0078be6e8dfb1b35a8bbe6dd6c8ee9cfe5e384eb7cd17a33879) |
| `timeouts.create` | [timeouts.create](resources--dns_compliance_checks--reference--group-001.md#canonical-441811748920cc4a010f203a1122ff9d6403e834d77225f9d1c4e604fa8bc969) |
| `timeouts.delete` | [timeouts.delete](resources--dns_compliance_checks--reference--group-001.md#canonical-73450572f303b5f940da197a44b76eee35524ca3a27352c2a0bf553a35e96c8a) |
| `timeouts.read` | [timeouts.read](resources--dns_compliance_checks--reference--group-001.md#canonical-fa3b23be0a49b6fa32ee2ac4cb214dfeb57bc449c9b50c6e2437adb684cfc133) |
| `timeouts.update` | [timeouts.update](resources--dns_compliance_checks--reference--group-001.md#canonical-f54f3cb2c3759b569d18b8f791df88a5f1bcf6c01f31b6c204e720a65077a465) |

<a id="canonical-57777904e3e4149c07cb46c1bbe0eb3cdc85d5bcc5bf9253ca7a307748d39ce0"></a>

## Next pages — Property reference / 624d8fc3618d / 15

- [timeouts](resources--dns_compliance_checks--reference--group-001.md#canonical-199d8bdf76f2f46e7a37feea668199530aa885f9d9764fdec7c86e3e6068ba1a)
- [xcsh_dns_compliance_checks](../resources/dns_compliance_checks.md#canonical-b8f0d3ec4e785ef6fabb81b92471aa91dd33fcda47717d8aa7b007104462e924)

<a id="canonical-199d8bdf76f2f46e7a37feea668199530aa885f9d9764fdec7c86e3e6068ba1a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-217ea02ee6b5ebdc4fe88ea19b885ecad86c332e840af8ab86ec53605f35e896"></a>

## timeouts — timeouts / 44fefd68186b / 2

Breadcrumbs:

- [xcsh_dns_compliance_checks](../resources/dns_compliance_checks.md#canonical-b8f0d3ec4e785ef6fabb81b92471aa91dd33fcda47717d8aa7b007104462e924)
- [Property reference](resources--dns_compliance_checks--reference--group-001.md#canonical-eee1b4d710d52db719c841deb6e007376a78ef3ab29c2e2a6fc1e96dd6936c22)
- timeouts

<a id="canonical-98a021adb39ea0078be6e8dfb1b35a8bbe6dd6c8ee9cfe5e384eb7cd17a33879"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-924fcd149dc6dd221cc0a1cd9ece532d538aca6c1a6364d09347584eed1d0dbb"></a>

## Direct properties — timeouts / 44fefd68186b / 3

<a id="canonical-441811748920cc4a010f203a1122ff9d6403e834d77225f9d1c4e604fa8bc969"></a>

<a id="canonical-f9809794616852ae550ad174d34e62329049aa204c7a7fb3bfc84ade685ab9b4"></a>

## create property — timeouts / 44fefd68186b / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-73450572f303b5f940da197a44b76eee35524ca3a27352c2a0bf553a35e96c8a"></a>

<a id="canonical-9d4c4089a9fd4f514a3ff045deeba7caa55744a1fad91dcbbe5fc28bbc0cf46d"></a>

## delete property — timeouts / 44fefd68186b / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-fa3b23be0a49b6fa32ee2ac4cb214dfeb57bc449c9b50c6e2437adb684cfc133"></a>

<a id="canonical-02b5e41163e26b30c43ad0623b103958d8255f4e320c1b9ef04a3fa3769d125e"></a>

## read property — timeouts / 44fefd68186b / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-f54f3cb2c3759b569d18b8f791df88a5f1bcf6c01f31b6c204e720a65077a465"></a>

<a id="canonical-d682246cd3677c8809a137bc69f2ce9ca6f6d74c16d8367254665cd617deb22c"></a>

## update property — timeouts / 44fefd68186b / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-c1b3d8badf6fcea440ddf54f8d3d16d4da1f7878a9f82a4b623ff1daa510c0c1"></a>

## Next pages — timeouts / 44fefd68186b / 8

- [Property reference](resources--dns_compliance_checks--reference--group-001.md#canonical-eee1b4d710d52db719c841deb6e007376a78ef3ab29c2e2a6fc1e96dd6936c22)
- [xcsh_dns_compliance_checks](../resources/dns_compliance_checks.md#canonical-b8f0d3ec4e785ef6fabb81b92471aa91dd33fcda47717d8aa7b007104462e924)
