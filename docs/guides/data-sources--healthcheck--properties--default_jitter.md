---
page_title: "default_jitter"
subcategory: "Monitoring"
description: "default_jitter for xcsh_healthcheck."
xcsh_docs: {"aliases": [], "body_bytes": 1194, "body_sha256": "sha256:34dfd3a7a0642a4cbe1af9aa992cd1334929c7c848aac33f52b043187cf93b79", "canonical_id": "xcsh-docs:data-sources:healthcheck:properties:default_jitter", "child_ids": [], "collection_id": "xcsh-docs:data-sources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:healthcheck:properties:default_jitter", "parent_id": "xcsh-docs:data-sources:healthcheck:reference", "path": "docs/guides/data-sources--healthcheck--properties--default_jitter.md", "provider_name": "healthcheck", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_jitter"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/healthcheck/properties/default_jitter/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_jitter for xcsh_healthcheck.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_jitter

Breadcrumbs:

- [xcsh_healthcheck](../data-sources/healthcheck.md)
- [Property reference](data-sources--healthcheck--reference.md)
- default_jitter

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_jitter, jitter\_percent; Default: default\_jitter\] Configuration parameter for
default jitter.

Upstream description:

This can be used for messages where no values are needed.

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

- [default_jitter](data-sources--healthcheck--properties--default_jitter.md#section)
- [jitter_percent](data-sources--healthcheck--reference.md#schema-jitter_percent)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--healthcheck--reference.md)
- [xcsh_healthcheck](../data-sources/healthcheck.md)
