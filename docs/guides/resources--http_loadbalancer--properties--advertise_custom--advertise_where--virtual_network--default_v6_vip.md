---
page_title: "advertise_custom.advertise_where.virtual_network.default_v6_vip"
subcategory: "Load Balancing"
description: "advertise_custom.advertise_where.virtual_network.default_v6_vip for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1439, "body_sha256": "sha256:c8b9f7875969e8fbe118ddf3d4b30a7eef0ea5a1fbb47202b7325d17acd9b5fd", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:default_v6_vip", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:default_v6_vip", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "path": "docs/guides/resources--http_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--default_v6_vip.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advertise_custom", "advertise_where", "virtual_network", "default_v6_vip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/advertise_custom/advertise_where/virtual_network/default_v6_vip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advertise_custom.advertise_where.virtual_network.default_v6_vip for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom.advertise_where.virtual_network.default_v6_vip

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [advertise_custom](resources--http_loadbalancer--properties--advertise_custom.md)
- [advertise_custom.advertise_where](resources--http_loadbalancer--properties--advertise_custom--advertise_where.md)
- [advertise_custom.advertise_where.virtual_network](resources--http_loadbalancer--properties--advertise_custom--advertise_where--virtual_network.md)
- advertise_custom.advertise_where.virtual_network.default_v6_vip

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
default_v6_vip = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [advertise_custom.advertise_where.virtual_network](resources--http_loadbalancer--properties--advertise_custom--advertise_where--virtual_network.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
