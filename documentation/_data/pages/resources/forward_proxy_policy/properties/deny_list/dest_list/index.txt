---
page_title: "deny_list.dest_list"
subcategory: "Security"
description: "L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI."
xcsh_docs: {"aliases": ["deny list dest list"], "body_bytes": 5952, "body_sha256": "sha256:a9f475157c8b750ee5684b7dc6b9d4a04831398be4290a8c33bc0b5e0476479a", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:dest_list", "parent_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list", "path": "documentation/resources/forward_proxy_policy/properties/deny_list/dest_list/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2021023221213120-2220230112212201-0313100222120123-1100030123122122-1000300112300300-1031311201022221-0220222331203131-2221231130321032", "registry_path": "docs/guides/resources--forward_proxy_policy--reference--group-001.md", "relationships": [{"anchor": "schema-deny_list--dest_list--port_ranges", "enforcement": "provider-schema", "group": "deny_list.dest_list:RequiredListObjectAttributes:port_ranges", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:dest_list", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["deny_list", "dest_list"], "schema_version": 1, "sections": [{"aliases": ["ipv6 prefixes"], "anchor": "schema-deny_list--dest_list--ipv6_prefixes", "description": "Destination IPv6 prefixes.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:dest_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["deny_list", "dest_list", "ipv6_prefixes"], "syntax": "attribute", "type": "list"}, {"aliases": ["port ranges"], "anchor": "schema-deny_list--dest_list--port_ranges", "description": "A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by \"-\".", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:dest_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["deny_list", "dest_list", "port_ranges"], "syntax": "attribute", "type": "string"}, {"aliases": ["prefixes"], "anchor": "schema-deny_list--dest_list--prefixes", "description": "Destination IPv4 prefixes.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:dest_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["deny_list", "dest_list", "prefixes"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/properties/deny_list/dest_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

Destination IPv6 prefixes.

Provider validators and defaults (from schema source):

```go
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

Upstream description:

A string containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Upstream description:

Destination IPv4 prefixes.

Provider validators and defaults (from schema source):

```go
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

## Next pages

- [deny_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
