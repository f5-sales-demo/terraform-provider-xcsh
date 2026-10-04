---
page_title: "custom_hash_algorithms"
subcategory: "Security"
description: "Specifies the hash algorithms to be used."
xcsh_docs: {"aliases": ["custom hash algorithms"], "body_bytes": 3441, "body_sha256": "sha256:b7bdc30c0358931f84efed235d1c6bd037f66e74177d3df21ef9c0c2fc5b5f24", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:certificate:collection", "completeness": "complete", "id": "xcsh-docs:resources:certificate:properties:custom_hash_algorithms", "parent_id": "xcsh-docs:resources:certificate:reference", "path": "documentation/resources/certificate/properties/custom_hash_algorithms/index.md", "product": "distributed-cloud", "provider_name": "certificate", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3131031131220113-3330221330321201-1023222220022220-0312222333120011-2110302312132303-2122013133211112-3110020100032323-3230021122333211", "registry_path": "docs/guides/resources--certificate--reference--group-001.md", "relationships": [{"anchor": "schema-custom_hash_algorithms--hash_algorithms", "enforcement": "provider-schema", "group": "custom_hash_algorithms:RequiredObjectAttributes:hash_algorithms", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:certificate:properties:custom_hash_algorithms", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_hash_algorithms"], "schema_version": 1, "sections": [{"aliases": ["custom hash algorithms hash algorithms"], "anchor": "schema-custom_hash_algorithms--hash_algorithms", "description": "Ordered list of hash algorithms to be used.", "document_id": "xcsh-docs:resources:certificate:properties:custom_hash_algorithms", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_hash_algorithms", "hash_algorithms"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate/properties/custom_hash_algorithms/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Specifies the hash algorithms to be used.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["certificateCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_hash_algorithms

Breadcrumbs:

- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/)
- custom_hash_algorithms

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_hash\_algorithms, disable\_ocsp\_stapling, use\_system\_defaults; Default:
use\_system\_defaults\] Specifies the hash algorithms to be used.

Upstream description:

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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

OneOf alternatives in this subsection:

- [custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/custom_hash_algorithms/#section)
- [disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/disable_ocsp_stapling/#section)
- [use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/use_system_defaults/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-custom_hash_algorithms--hash_algorithms"></a>

### hash_algorithms property

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/)
- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/)
