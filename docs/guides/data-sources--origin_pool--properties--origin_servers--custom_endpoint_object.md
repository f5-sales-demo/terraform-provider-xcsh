---
page_title: "origin_servers.custom_endpoint_object"
subcategory: "Load Balancing"
description: "origin_servers.custom_endpoint_object for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1123, "body_sha256": "sha256:5f0bf3c4a46a55121dee9de8c759affbdd970ba746b80d4ba31cff6cd2c2d7ab", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:custom_endpoint_object", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:origin_servers:custom_endpoint_object:endpoint"], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:custom_endpoint_object", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers", "path": "docs/guides/data-sources--origin_pool--properties--origin_servers--custom_endpoint_object.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "custom_endpoint_object"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/origin_servers/custom_endpoint_object/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.custom_endpoint_object for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.custom_endpoint_object

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- [origin_servers](data-sources--origin_pool--properties--origin_servers.md)
- origin_servers.custom_endpoint_object

<a id="section"></a>

Type: `"single"`. Computed.

Specify origin server with a reference to endpoint object.

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

- [endpoint](data-sources--origin_pool--properties--origin_servers--custom_endpoint_object--endpoint.md): complete subsection reference.

## Next pages

- [origin_servers.custom_endpoint_object.endpoint](data-sources--origin_pool--properties--origin_servers--custom_endpoint_object--endpoint.md)
- [origin_servers](data-sources--origin_pool--properties--origin_servers.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
