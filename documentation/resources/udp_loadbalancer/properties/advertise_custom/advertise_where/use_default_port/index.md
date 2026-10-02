---
page_title: "advertise_custom.advertise_where.use_default_port"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["advertise custom advertise where use default port"], "body_bytes": 1528, "body_sha256": "sha256:8330289c83fa2b85065512ea27caee9ce11a9f4199bb172031f3d42e98324779", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:use_default_port", "parent_id": "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where", "path": "documentation/resources/udp_loadbalancer/properties/advertise_custom/advertise_where/use_default_port/index.md", "product": "distributed-cloud", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3223033300212032-1333311221011013-0020020020322222-0013201132113031-2012200201021121-1202212313033222-0102233110002223-3012210133233132", "registry_path": "docs/guides/resources--udp_loadbalancer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advertise_custom", "advertise_where", "use_default_port"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/udp_loadbalancer/properties/advertise_custom/advertise_where/use_default_port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom.advertise_where.use_default_port

Breadcrumbs:

- [xcsh_udp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/)
- [advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/advertise_custom/)
- [advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/advertise_custom/advertise_where/)
- advertise_custom.advertise_where.use_default_port

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
use_default_port = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/advertise_custom/advertise_where/)
- [xcsh_udp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/)
