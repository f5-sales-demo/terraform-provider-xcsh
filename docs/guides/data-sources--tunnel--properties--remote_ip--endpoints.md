---
page_title: "remote_ip.endpoints"
subcategory: ""
description: "remote_ip.endpoints for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1189, "body_sha256": "sha256:7a6cbe2165055e964493c575f1c2d824a3c8a1eb22308585c3e72e7398aaeefb", "canonical_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:endpoints", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:remote_ip:endpoints:endpoints"], "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:endpoints", "parent_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip", "path": "docs/guides/data-sources--tunnel--properties--remote_ip--endpoints.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["remote_ip", "endpoints"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/remote_ip/endpoints/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "remote_ip.endpoints for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# remote_ip.endpoints

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md)
- [Property reference](data-sources--tunnel--reference.md)
- [remote_ip](data-sources--tunnel--properties--remote_ip.md)
- remote_ip.endpoints

<a id="section"></a>

Type: `"single"`. Computed.

Provides a map of ver node name to remote node attributes Ver node should use these attributes to
configure as remote tunnel.

Upstream description:

Provides a map of ver node name to remote node attributes Ver node should use these attributes to
configure as remote tunnel.

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

- [endpoints](data-sources--tunnel--properties--remote_ip--endpoints--endpoints.md): complete subsection reference.

## Next pages

- [remote_ip.endpoints.endpoints](data-sources--tunnel--properties--remote_ip--endpoints--endpoints.md)
- [remote_ip](data-sources--tunnel--properties--remote_ip.md)
- [xcsh_tunnel](../data-sources/tunnel.md)
