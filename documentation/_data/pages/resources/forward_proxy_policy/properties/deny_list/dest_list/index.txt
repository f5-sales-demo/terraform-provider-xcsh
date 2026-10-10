---
page_title: "deny_list.dest_list"
subcategory: "Security"
description: "L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI."
xcsh_docs: {"aliases": ["deny list dest list"], "body_bytes": 5717, "body_sha256": "sha256:d008403f9f23a5eb6b623b4b15af326cfcc25cc1876328f2db72ab91172db1d1", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:dest_list", "parent_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list", "path": "documentation/resources/forward_proxy_policy/properties/deny_list/dest_list/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2021023221213120-2220230112212201-0313100222120123-1100030123122122-1000300112300300-1031311201022221-0220222331203131-2221231130321032", "registry_path": "docs/guides/resources--forward_proxy_policy--reference--group-001.md", "relationships": [{"anchor": "schema-deny_list--dest_list--port_ranges", "enforcement": "provider-schema", "group": "deny_list.dest_list:RequiredListObjectAttributes:port_ranges", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:dest_list", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["deny_list", "dest_list"], "schema_version": 1, "sections": [{"aliases": ["deny list dest list ipv6 prefixes"], "anchor": "schema-deny_list--dest_list--ipv6_prefixes", "description": "Destination IPv6 prefixes.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:dest_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["deny_list", "dest_list", "ipv6_prefixes"], "syntax": "attribute", "type": "list"}, {"aliases": ["deny list dest list port ranges"], "anchor": "schema-deny_list--dest_list--port_ranges", "description": "A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by \"-\".", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:dest_list", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["deny_list", "dest_list", "port_ranges"], "syntax": "attribute", "type": "string"}, {"aliases": ["deny list dest list prefixes"], "anchor": "schema-deny_list--dest_list--prefixes", "description": "Destination IPv4 prefixes.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:dest_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["deny_list", "dest_list", "prefixes"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/properties/deny_list/dest_list/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# deny_list.dest_list

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/)
- [deny_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/)
- deny_list.dest_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("port_ranges")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
dest_list {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-deny_list--dest_list--ipv6_prefixes"></a>

### ipv6_prefixes property

Type: `["list", "string"]`. Optional.

IPv6 Prefixes. Destination IPv6 prefixes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-deny_list--dest_list--port_ranges"></a>

### port_ranges property

Type: `"string"`. Optional.

String containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by '-'.

Additional upstream details:

A string containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="schema-deny_list--dest_list--prefixes"></a>

### prefixes property

Type: `["list", "string"]`. Optional.

IPv4 Prefixes. Destination IPv4 prefixes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
