---
page_title: "use_tls.default_session_key_caching"
subcategory: "Load Balancing"
description: "use_tls.default_session_key_caching for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1095, "body_sha256": "sha256:531bafb7b7f513b6647a2937cc709824881d8638d5118d0192e8e8096d6aba26", "canonical_id": "xcsh-docs:resources:origin_pool:properties:use_tls:default_session_key_caching", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:use_tls:default_session_key_caching", "parent_id": "xcsh-docs:resources:origin_pool:properties:use_tls", "path": "docs/guides/resources--origin_pool--properties--use_tls--default_session_key_caching.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_tls", "default_session_key_caching"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/use_tls/default_session_key_caching/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_tls.default_session_key_caching for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_tls.default_session_key_caching

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [use_tls](resources--origin_pool--properties--use_tls.md)
- use_tls.default_session_key_caching

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for default session key caching. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

Terraform syntax:

```terraform
default_session_key_caching = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [use_tls](resources--origin_pool--properties--use_tls.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
