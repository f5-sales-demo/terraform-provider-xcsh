---
page_title: "cache_profile"
subcategory: ""
description: "DNS Cache specifies cache configuration."
xcsh_docs: {"aliases": ["cache profile"], "body_bytes": 2691, "body_sha256": "sha256:e7206b51cb27a0b57cbbb15420c2e6bb6e7157bdcf25a84b59d6d0ad014756f4", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:cache_profile:disable_cache_profile"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:cache_profile", "parent_id": "xcsh-docs:resources:dns_proxy:reference", "path": "documentation/resources/dns_proxy/properties/cache_profile/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3223201131001301-3012133233101211-1220300001232100-3132313032022201-0320021233121203-3212310110212220-1013310330311333-2031313023321333", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [{"anchor": "schema-cache_profile--cache_size", "enforcement": "provider-schema", "group": "cache_profile:ConflictingObjectAttributes:cache_size,disable_cache_profile", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:cache_profile", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cache_profile:ConflictingObjectAttributes:cache_size,disable_cache_profile", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:cache_profile:disable_cache_profile", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cache_profile"], "schema_version": 1, "sections": [{"aliases": ["cache profile cache size"], "anchor": "schema-cache_profile--cache_size", "description": "Exclusive with cache size.", "document_id": "xcsh-docs:resources:dns_proxy:properties:cache_profile", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_profile", "cache_size"], "syntax": "attribute", "type": "number"}, {"aliases": ["cache profile disable cache profile"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:dns_proxy:properties:cache_profile:disable_cache_profile", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_profile", "disable_cache_profile"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/cache_profile/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "DNS Cache specifies cache configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_profile

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- cache_profile

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

DNS Cache specifies cache configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cache_size",
    "disable_cache_profile")}
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
  "x-ves-oneof-field-cache_profile_choice": "[\"cache_size\",\"disable_cache_profile\"]"
}
```

Terraform syntax:

```terraform
cache_profile {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-cache_profile--cache_size"></a>

### cache_size property

Type: `"number"`. Optional.

Exclusive with \[disable\_cache\_profile\] cache size.

Upstream description:

Exclusive with \[disable\_cache\_profile\] cache size.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 10240),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10240,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "10240"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "10240"
  }
}
```

- [disable_cache_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/cache_profile/disable_cache_profile/): complete subsection reference.

## Next pages

- [cache_profile.disable_cache_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/cache_profile/disable_cache_profile/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
