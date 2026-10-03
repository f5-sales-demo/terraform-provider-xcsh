---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_dns_compliance_checks."
xcsh_docs: {"aliases": ["dns compliance checks"], "body_bytes": 12186, "body_sha256": "sha256:3a68c12a68ac793c8a36552c5ffbca7a75943776669c6f38465237cdb1d819a1", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_compliance_checks:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_compliance_checks:reference", "parent_id": "xcsh-docs:data-sources:dns_compliance_checks:fundamentals", "path": "documentation/data-sources/dns_compliance_checks/properties/index.md", "product": "distributed-cloud", "provider_name": "dns_compliance_checks", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0330131002333110-1321122220010200-2131223002102021-1310111221323211-0220230233200220-1023112122322033-2333123220200320-2331221120111033", "registry_path": "docs/guides/data-sources--dns_compliance_checks--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:dns_compliance_checks:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:dns_compliance_checks:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disallowed query type list"], "anchor": "schema-disallowed_query_type_list", "description": "Disallowed Query Type Values.", "document_id": "xcsh-docs:data-sources:dns_compliance_checks:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disallowed_query_type_list"], "syntax": "attribute", "type": "list"}, {"aliases": ["disallowed resource record type list"], "anchor": "schema-disallowed_resource_record_type_list", "description": "Disallowed Resource Record Type List.", "document_id": "xcsh-docs:data-sources:dns_compliance_checks:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disallowed_resource_record_type_list"], "syntax": "attribute", "type": "list"}, {"aliases": ["domain denylist"], "anchor": "schema-domain_denylist", "description": "List of domains to be denied by configuration object.", "document_id": "xcsh-docs:data-sources:dns_compliance_checks:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domain_denylist"], "syntax": "attribute", "type": "list"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:dns_compliance_checks:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:dns_compliance_checks:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:dns_compliance_checks:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:dns_compliance_checks:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_compliance_checks/properties/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "Property reference for xcsh_dns_compliance_checks.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["dns_compliance_checksCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_dns_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

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

<a id="schema-description"></a>

### description property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-disallowed_query_type_list"></a>

### disallowed_query_type_list property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-disallowed_resource_record_type_list"></a>

### disallowed_resource_record_type_list property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-domain_denylist"></a>

### domain_denylist property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

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

<a id="schema-name"></a>

### name property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-namespace"></a>

### namespace property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/properties/#schema-description) |
| `disallowed_query_type_list` | [disallowed_query_type_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/properties/#schema-disallowed_query_type_list) |
| `disallowed_resource_record_type_list` | [disallowed_resource_record_type_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/properties/#schema-disallowed_resource_record_type_list) |
| `domain_denylist` | [domain_denylist](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/properties/#schema-domain_denylist) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/properties/#schema-namespace) |

## Next pages

- [xcsh_dns_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/)
