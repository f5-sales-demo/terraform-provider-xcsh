---
page_title: "origin_servers.consul_service.outside_network"
subcategory: "Load Balancing"
description: "origin_servers.consul_service.outside_network for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1183, "body_sha256": "sha256:4bca736bde00f57ab4b531b38ab9082948358cdf5a323e3573c49b97cd53d2be", "canonical_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:outside_network", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:outside_network", "parent_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service", "path": "docs/guides/resources--origin_pool--properties--origin_servers--consul_service--outside_network.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "consul_service", "outside_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/origin_servers/consul_service/outside_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.consul_service.outside_network for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.consul_service.outside_network

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [origin_servers](resources--origin_pool--properties--origin_servers.md)
- [origin_servers.consul_service](resources--origin_pool--properties--origin_servers--consul_service.md)
- origin_servers.consul_service.outside_network

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [origin_servers.consul_service](resources--origin_pool--properties--origin_servers--consul_service.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
