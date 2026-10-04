---
page_title: "default_cache_action"
subcategory: "Load Balancing"
description: "This defines a Default Cache Action."
xcsh_docs: {"aliases": ["default cache action"], "body_bytes": 3943, "body_sha256": "sha256:9a6b7f47030d6bb962191465c78c522777ae452f1a7628c1ac5671a9c90352b3", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:default_cache_action:cache_disabled"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:default_cache_action", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/default_cache_action/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2300021122222120-3103201200301030-2101110210210123-3021222103220303-1221111131320223-2202110300233100-2330133310101313-0202223013312313", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-010.md", "relationships": [{"anchor": "schema-default_cache_action--cache_ttl_default", "enforcement": "provider-schema", "group": "default_cache_action:ConflictingObjectAttributes:cache_disabled,cache_ttl_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:default_cache_action", "type": "conflicts"}, {"anchor": "schema-default_cache_action--cache_ttl_default", "enforcement": "provider-schema", "group": "default_cache_action:ConflictingObjectAttributes:cache_ttl_default,cache_ttl_override", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:default_cache_action", "type": "conflicts"}, {"anchor": "schema-default_cache_action--cache_ttl_override", "enforcement": "provider-schema", "group": "default_cache_action:ConflictingObjectAttributes:cache_disabled,cache_ttl_override", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:default_cache_action", "type": "conflicts"}, {"anchor": "schema-default_cache_action--cache_ttl_override", "enforcement": "provider-schema", "group": "default_cache_action:ConflictingObjectAttributes:cache_ttl_default,cache_ttl_override", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:default_cache_action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_cache_action:ConflictingObjectAttributes:cache_disabled,cache_ttl_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:default_cache_action:cache_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_cache_action:ConflictingObjectAttributes:cache_disabled,cache_ttl_override", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:default_cache_action:cache_disabled", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["default_cache_action"], "schema_version": 1, "sections": [{"aliases": ["default cache action cache disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:default_cache_action:cache_disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_cache_action", "cache_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["default cache action cache ttl default"], "anchor": "schema-default_cache_action--cache_ttl_default", "description": "Exclusive with Use Cache TTL Provided by Origin, and set a contigency TTL value in case one is not provided.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:default_cache_action", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_cache_action", "cache_ttl_default"], "syntax": "attribute", "type": "string"}, {"aliases": ["default cache action cache ttl override"], "anchor": "schema-default_cache_action--cache_ttl_override", "description": "Exclusive with Always override the Cache TTL provided by Origin.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:default_cache_action", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_cache_action", "cache_ttl_override"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/default_cache_action/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "This defines a Default Cache Action.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_cache_action

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- default_cache_action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Default Cache Behaviour. This defines a Default Cache Action.

Upstream description:

This defines a Default Cache Action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cache_disabled",
    "cache_ttl_default"),
  validators.ConflictingObjectAttributes("cache_disabled",
    "cache_ttl_override"),
  validators.ConflictingObjectAttributes("cache_ttl_default",
    "cache_ttl_override")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cache_actions": "[\"cache_disabled\",\"cache_ttl_default\",\"cache_ttl_override\"]"
}
```

Terraform syntax:

```terraform
default_cache_action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cache_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/default_cache_action/cache_disabled/): complete subsection reference.

<a id="schema-default_cache_action--cache_ttl_default"></a>

### cache_ttl_default property

Type: `"string"`. Optional.

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

Upstream description:

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="schema-default_cache_action--cache_ttl_override"></a>

### cache_ttl_override property

Type: `"string"`. Optional.

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

Upstream description:

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

## Next pages

- [default_cache_action.cache_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/default_cache_action/cache_disabled/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
