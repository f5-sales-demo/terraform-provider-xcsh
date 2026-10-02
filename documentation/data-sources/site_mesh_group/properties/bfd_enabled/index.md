---
page_title: "bfd_enabled"
subcategory: "Infrastructure"
description: "BFD parameters."
xcsh_docs: {"aliases": ["bfd enabled"], "body_bytes": 3973, "body_sha256": "sha256:2fcaadbcde58fcaeeb7e27e3ee537fa576c629afd6a6d7ff8d71148d2c224a96", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_mesh_group:properties:bfd_enabled", "parent_id": "xcsh-docs:data-sources:site_mesh_group:reference", "path": "documentation/data-sources/site_mesh_group/properties/bfd_enabled/index.md", "product": "distributed-cloud", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3213211310202102-0211010112232012-2201332101100002-0000332311333313-2331210033000311-0220021223001321-0113001300201002-0221021220002020", "registry_path": "docs/guides/data-sources--site_mesh_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bfd_enabled"], "schema_version": 1, "sections": [{"aliases": ["multiplier"], "anchor": "schema-bfd_enabled--multiplier", "description": "Specify Number of missed packets to bring session down\"", "document_id": "xcsh-docs:data-sources:site_mesh_group:properties:bfd_enabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bfd_enabled", "multiplier"], "syntax": "attribute", "type": "number"}, {"aliases": ["receive interval milliseconds"], "anchor": "schema-bfd_enabled--receive_interval_milliseconds", "description": "BFD receive interval timer, in milliseconds.", "document_id": "xcsh-docs:data-sources:site_mesh_group:properties:bfd_enabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bfd_enabled", "receive_interval_milliseconds"], "syntax": "attribute", "type": "number"}, {"aliases": ["transmit interval milliseconds"], "anchor": "schema-bfd_enabled--transmit_interval_milliseconds", "description": "BFD transmit interval timer, in milliseconds.", "document_id": "xcsh-docs:data-sources:site_mesh_group:properties:bfd_enabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bfd_enabled", "transmit_interval_milliseconds"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_mesh_group/properties/bfd_enabled/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "BFD parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bfd_enabled

Breadcrumbs:

- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/)
- bfd_enabled

<a id="section"></a>

Type: `"single"`. Computed.

BFD. BFD parameters.

Upstream description:

BFD parameters.

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

## Direct properties

<a id="schema-bfd_enabled--multiplier"></a>

### multiplier property

Type: `"number"`. Computed.

Specify Number of missed packets to bring session down'.

Upstream description:

Specify Number of missed packets to bring session down"

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="schema-bfd_enabled--receive_interval_milliseconds"></a>

### receive_interval_milliseconds property

Type: `"number"`. Computed.

BFD receive interval timer, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="schema-bfd_enabled--transmit_interval_milliseconds"></a>

### transmit_interval_milliseconds property

Type: `"number"`. Computed.

BFD transmit interval timer, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/)
- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/)
