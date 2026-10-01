---
page_title: "use_tls.tls_config.default_security"
subcategory: "Load Balancing"
description: "use_tls.tls_config.default_security for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1080, "body_sha256": "sha256:ede0fc6ed275b3fb19901b9093f69142f16513bc9cc17d4e311dba9bbdd021a9", "canonical_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:default_security", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:default_security", "parent_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config", "path": "docs/guides/resources--origin_pool--properties--use_tls--tls_config--default_security.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_tls", "tls_config", "default_security"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/use_tls/tls_config/default_security/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_tls.tls_config.default_security for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_tls.tls_config.default_security

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [use_tls](resources--origin_pool--properties--use_tls.md)
- [use_tls.tls_config](resources--origin_pool--properties--use_tls--tls_config.md)
- use_tls.tls_config.default_security

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
default_security = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [use_tls.tls_config](resources--origin_pool--properties--use_tls--tls_config.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
