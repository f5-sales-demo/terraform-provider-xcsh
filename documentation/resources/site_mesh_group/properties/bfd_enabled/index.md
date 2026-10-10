---
page_title: "bfd_enabled"
subcategory: "Infrastructure"
description: "BFD parameters."
xcsh_docs: {"aliases": ["bfd enabled"], "body_bytes": 4559, "body_sha256": "sha256:3e84affa919d4c31648b53b6830127513498b6be4d0057d50a62209dce6727bd", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:site_mesh_group:properties:bfd_enabled", "parent_id": "xcsh-docs:resources:site_mesh_group:reference", "path": "documentation/resources/site_mesh_group/properties/bfd_enabled/index.md", "product": "distributed-cloud", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0030312132232112-2303110333223203-2033302223131130-2123310023033230-1313103013010103-0011023222103320-3330033300112032-0231330313022102", "registry_path": "docs/guides/resources--site_mesh_group--reference--group-001.md", "relationships": [{"anchor": "schema-bfd_enabled--multiplier", "enforcement": "provider-schema", "group": "bfd_enabled:RequiredObjectAttributes:multiplier,receive_interval_milliseconds,transmit_interval_milliseconds", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:site_mesh_group:properties:bfd_enabled", "type": "requires"}, {"anchor": "schema-bfd_enabled--receive_interval_milliseconds", "enforcement": "provider-schema", "group": "bfd_enabled:RequiredObjectAttributes:multiplier,receive_interval_milliseconds,transmit_interval_milliseconds", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:site_mesh_group:properties:bfd_enabled", "type": "requires"}, {"anchor": "schema-bfd_enabled--transmit_interval_milliseconds", "enforcement": "provider-schema", "group": "bfd_enabled:RequiredObjectAttributes:multiplier,receive_interval_milliseconds,transmit_interval_milliseconds", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:site_mesh_group:properties:bfd_enabled", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bfd_enabled"], "schema_version": 1, "sections": [{"aliases": ["bfd enabled multiplier"], "anchor": "schema-bfd_enabled--multiplier", "description": "Specify Number of missed packets to bring session down\"", "document_id": "xcsh-docs:resources:site_mesh_group:properties:bfd_enabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bfd_enabled", "multiplier"], "syntax": "attribute", "type": "number"}, {"aliases": ["bfd enabled receive interval milliseconds"], "anchor": "schema-bfd_enabled--receive_interval_milliseconds", "description": "BFD receive interval timer, in milliseconds.", "document_id": "xcsh-docs:resources:site_mesh_group:properties:bfd_enabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bfd_enabled", "receive_interval_milliseconds"], "syntax": "attribute", "type": "number"}, {"aliases": ["bfd enabled transmit interval milliseconds"], "anchor": "schema-bfd_enabled--transmit_interval_milliseconds", "description": "BFD transmit interval timer, in milliseconds.", "document_id": "xcsh-docs:resources:site_mesh_group:properties:bfd_enabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bfd_enabled", "transmit_interval_milliseconds"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/site_mesh_group/properties/bfd_enabled/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "BFD parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bfd_enabled

Breadcrumbs:

- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/)
- bfd_enabled

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

BFD. BFD parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("multiplier",
    "receive_interval_milliseconds",
    "transmit_interval_milliseconds")}
```

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

Terraform syntax:

```terraform
bfd_enabled {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-bfd_enabled--multiplier"></a>

### multiplier property

Type: `"number"`. Optional.

Specify Number of missed packets to bring session down'.

Additional upstream details:

Specify Number of missed packets to bring session down"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(2, 255),
}
```

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Type: `"number"`. Optional.

BFD receive interval timer, in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(300, 60000),
}
```

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Type: `"number"`. Optional.

BFD transmit interval timer, in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(300, 60000),
}
```

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
