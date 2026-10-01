---
page_title: "use_tls.tls_config"
subcategory: "Load Balancing"
description: "use_tls.tls_config for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 2338, "body_sha256": "sha256:93cd446c22a3f6f94edce5e5f588ccd5e6c13664c71c210d0cb8c4fc186e6f9a", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:tls_config", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:use_tls:tls_config:custom_security", "xcsh-docs:data-sources:origin_pool:properties:use_tls:tls_config:default_security", "xcsh-docs:data-sources:origin_pool:properties:use_tls:tls_config:low_security", "xcsh-docs:data-sources:origin_pool:properties:use_tls:tls_config:medium_security"], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:tls_config", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls", "path": "docs/guides/data-sources--origin_pool--properties--use_tls--tls_config.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_tls", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/use_tls/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_tls.tls_config for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_tls.tls_config

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- [use_tls](data-sources--origin_pool--properties--use_tls.md)
- use_tls.tls_config

<a id="section"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "tls",
    "constraintType": "object",
    "deterministic": true,
    "metadata": {
      "category": "tls",
      "confidence": 0.99,
      "note": "Required when use_tls is selected — API returns 400 if nil",
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

## Direct properties

- [custom_security](data-sources--origin_pool--properties--use_tls--tls_config--custom_security.md): complete subsection reference.

- [default_security](data-sources--origin_pool--properties--use_tls--tls_config--default_security.md): complete subsection reference.

- [low_security](data-sources--origin_pool--properties--use_tls--tls_config--low_security.md): complete subsection reference.

- [medium_security](data-sources--origin_pool--properties--use_tls--tls_config--medium_security.md): complete subsection reference.

## Next pages

- [use_tls.tls_config.custom_security](data-sources--origin_pool--properties--use_tls--tls_config--custom_security.md)
- [use_tls.tls_config.default_security](data-sources--origin_pool--properties--use_tls--tls_config--default_security.md)
- [use_tls.tls_config.low_security](data-sources--origin_pool--properties--use_tls--tls_config--low_security.md)
- [use_tls.tls_config.medium_security](data-sources--origin_pool--properties--use_tls--tls_config--medium_security.md)
- [use_tls](data-sources--origin_pool--properties--use_tls.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
