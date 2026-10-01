---
page_title: "advertise_custom.advertise_where.advertise_v6_on_public"
subcategory: ""
description: "advertise_custom.advertise_where.advertise_v6_on_public for xcsh_udp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1724, "body_sha256": "sha256:59b2ddf0df5bd79cf8456d3aa37b7de9514cf82fd6aeb89e19b9626606944297", "canonical_id": "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where:advertise_v6_on_public", "child_ids": ["xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where:advertise_v6_on_public:public_ip"], "collection_id": "xcsh-docs:data-sources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where:advertise_v6_on_public", "parent_id": "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where", "path": "docs/guides/data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public.md", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advertise_custom", "advertise_where", "advertise_v6_on_public"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/udp_loadbalancer/properties/advertise_custom/advertise_where/advertise_v6_on_public/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advertise_custom.advertise_where.advertise_v6_on_public for xcsh_udp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom.advertise_where.advertise_v6_on_public

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md)
- [Property reference](data-sources--udp_loadbalancer--reference.md)
- [advertise_custom](data-sources--udp_loadbalancer--properties--advertise_custom.md)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where.md)
- advertise_custom.advertise_where.advertise_v6_on_public

<a id="section"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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

- [public_ip](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md): complete subsection reference.

## Next pages

- [advertise_custom.advertise_where.advertise_v6_on_public.public_ip](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where.md)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md)
