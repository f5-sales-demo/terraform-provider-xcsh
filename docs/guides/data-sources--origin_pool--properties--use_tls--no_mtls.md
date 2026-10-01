---
page_title: "use_tls.no_mtls"
subcategory: "Load Balancing"
description: "use_tls.no_mtls for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 953, "body_sha256": "sha256:fb16d98fdca3270bbe6594644e18065686fad286caf6c1ef0d91f8efe41dc7ee", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:no_mtls", "child_ids": [], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:no_mtls", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls", "path": "docs/guides/data-sources--origin_pool--properties--use_tls--no_mtls.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_tls", "no_mtls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/use_tls/no_mtls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_tls.no_mtls for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_tls.no_mtls

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- [use_tls](data-sources--origin_pool--properties--use_tls.md)
- use_tls.no_mtls

<a id="section"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [use_tls](data-sources--origin_pool--properties--use_tls.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
