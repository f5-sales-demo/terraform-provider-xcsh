---
page_title: "origin_servers.custom_endpoint_object"
subcategory: "Load Balancing"
description: "origin_servers.custom_endpoint_object for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1233, "body_sha256": "sha256:4ab4d09d7752b176332f5c5bf32ed0e2957febbc567db54d70b0419d7122a540", "canonical_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:custom_endpoint_object", "child_ids": ["xcsh-docs:resources:origin_pool:properties:origin_servers:custom_endpoint_object:endpoint"], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:origin_servers:custom_endpoint_object", "parent_id": "xcsh-docs:resources:origin_pool:properties:origin_servers", "path": "docs/guides/resources--origin_pool--properties--origin_servers--custom_endpoint_object.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "custom_endpoint_object"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/origin_servers/custom_endpoint_object/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.custom_endpoint_object for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.custom_endpoint_object

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [origin_servers](resources--origin_pool--properties--origin_servers.md)
- origin_servers.custom_endpoint_object

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom_endpoint_object {
  # Configure direct properties listed below.
}
```

## Direct properties

- [endpoint](resources--origin_pool--properties--origin_servers--custom_endpoint_object--endpoint.md): complete subsection reference.

## Next pages

- [origin_servers.custom_endpoint_object.endpoint](resources--origin_pool--properties--origin_servers--custom_endpoint_object--endpoint.md)
- [origin_servers](resources--origin_pool--properties--origin_servers.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
