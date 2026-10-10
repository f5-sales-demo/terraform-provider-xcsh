---
page_title: "custom_hash_algorithms"
subcategory: "Security"
description: "Specifies the hash algorithms to be used."
xcsh_docs: {"aliases": ["custom hash algorithms"], "body_bytes": 2788, "body_sha256": "sha256:2a25dc0bd90c8e596da03dea92b31bd87f2242ed228ea69efe289de7944f3250", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:certificate:collection", "completeness": "complete", "id": "xcsh-docs:resources:certificate:properties:custom_hash_algorithms", "parent_id": "xcsh-docs:resources:certificate:reference", "path": "documentation/resources/certificate/properties/custom_hash_algorithms/index.md", "product": "distributed-cloud", "provider_name": "certificate", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3131031131220113-3330221330321201-1023222220022220-0312222333120011-2110302312132303-2122013133211112-3110020100032323-3230021122333211", "registry_path": "docs/guides/resources--certificate--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_hash_algorithms"], "schema_version": 1, "sections": [{"aliases": ["custom hash algorithms hash algorithms"], "anchor": "schema-custom_hash_algorithms--hash_algorithms", "description": "Ordered list of hash algorithms to be used.", "document_id": "xcsh-docs:resources:certificate:properties:custom_hash_algorithms", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_hash_algorithms", "hash_algorithms"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate/properties/custom_hash_algorithms/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Specifies the hash algorithms to be used.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["certificateCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
