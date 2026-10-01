---
page_title: "custom_hash_algorithms"
subcategory: "Security"
description: "custom_hash_algorithms for xcsh_certificate."
xcsh_docs: {"aliases": [], "body_bytes": 2682, "body_sha256": "sha256:772765088552daab85f054d3c622557f78df70c20af206bee5728462410eec54", "canonical_id": "xcsh-docs:data-sources:certificate:properties:custom_hash_algorithms", "child_ids": [], "collection_id": "xcsh-docs:data-sources:certificate:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certificate:properties:custom_hash_algorithms", "parent_id": "xcsh-docs:data-sources:certificate:reference", "path": "docs/guides/data-sources--certificate--properties--custom_hash_algorithms.md", "provider_name": "certificate", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_hash_algorithms"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certificate/properties/custom_hash_algorithms/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_hash_algorithms for xcsh_certificate.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["certificateCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_hash_algorithms

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md)
- [Property reference](data-sources--certificate--reference.md)
- custom_hash_algorithms

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_hash\_algorithms, disable\_ocsp\_stapling, use\_system\_defaults; Default:
use\_system\_defaults\] Specifies the hash algorithms to be used.

Upstream description:

Specifies the hash algorithms to be used.

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

- [custom_hash_algorithms](data-sources--certificate--properties--custom_hash_algorithms.md#section)
- [disable_ocsp_stapling](data-sources--certificate--properties--disable_ocsp_stapling.md#section)
- [use_system_defaults](data-sources--certificate--properties--use_system_defaults.md#section)

Select alternatives according to the provider validators above.

## Direct properties

<a id="schema-custom_hash_algorithms--hash_algorithms"></a>

### hash_algorithms property

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [Property reference](data-sources--certificate--reference.md)
- [xcsh_certificate](../data-sources/certificate.md)
