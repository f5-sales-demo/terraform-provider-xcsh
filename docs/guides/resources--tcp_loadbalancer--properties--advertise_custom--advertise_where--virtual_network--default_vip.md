---
page_title: "advertise_custom.advertise_where.virtual_network.default_vip"
subcategory: "Load Balancing"
description: "advertise_custom.advertise_where.virtual_network.default_vip for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1421, "body_sha256": "sha256:aa7a94b126cc7b9ee6e4d18dfd1ea668068d05f77431fbaebc3d36d2114e18ef", "canonical_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:default_vip", "child_ids": [], "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:default_vip", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "path": "docs/guides/resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--default_vip.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advertise_custom", "advertise_where", "virtual_network", "default_vip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/virtual_network/default_vip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advertise_custom.advertise_where.virtual_network.default_vip for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom.advertise_where.virtual_network.default_vip

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
- [Property reference](resources--tcp_loadbalancer--reference.md)
- [advertise_custom](resources--tcp_loadbalancer--properties--advertise_custom.md)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where.md)
- [advertise_custom.advertise_where.virtual_network](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network.md)
- advertise_custom.advertise_where.virtual_network.default_vip

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
default_vip = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [advertise_custom.advertise_where.virtual_network](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network.md)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
