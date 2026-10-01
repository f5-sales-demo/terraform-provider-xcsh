---
page_title: "automatic_port"
subcategory: "Load Balancing"
description: "automatic_port for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1229, "body_sha256": "sha256:2b8f9f1bac9aeec320430ca382fca9fa9e8eb601b1b03170e468c758e0d9dea2", "canonical_id": "xcsh-docs:resources:origin_pool:properties:automatic_port", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:automatic_port", "parent_id": "xcsh-docs:resources:origin_pool:reference", "path": "docs/guides/resources--origin_pool--properties--automatic_port.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["automatic_port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/automatic_port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "automatic_port for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# automatic_port

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- automatic_port

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: automatic\_port, lb\_port, port\] Enable this option

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

- [automatic_port](resources--origin_pool--properties--automatic_port.md#section)
- [lb_port](resources--origin_pool--properties--lb_port.md#section)
- [port](resources--origin_pool--reference.md#schema-port)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
automatic_port = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--origin_pool--reference.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
