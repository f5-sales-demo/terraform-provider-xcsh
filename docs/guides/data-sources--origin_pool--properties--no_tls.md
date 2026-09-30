---
page_title: "no_tls"
subcategory: "Load Balancing"
description: "no_tls for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1067, "body_sha256": "sha256:93a66cd628c6edf6f0d030710e68f1d0003834b85293580a0167ce01398a3b87", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:no_tls", "child_ids": [], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:no_tls", "parent_id": "xcsh-docs:data-sources:origin_pool:reference", "path": "docs/guides/data-sources--origin_pool--properties--no_tls.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["no_tls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/no_tls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "no_tls for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# no_tls

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- no_tls

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: no\_tls, use\_tls; Default: no\_tls\] Enable this option. Defaults to \`map\[\]\`. Server
applies default when omitted.

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

- [no_tls](data-sources--origin_pool--properties--no_tls.md#section)
- [use_tls](data-sources--origin_pool--properties--use_tls.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--origin_pool--reference.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
