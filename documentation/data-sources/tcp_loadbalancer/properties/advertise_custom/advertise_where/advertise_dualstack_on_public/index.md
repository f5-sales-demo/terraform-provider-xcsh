---
page_title: "advertise_custom.advertise_where.advertise_dualstack_on_public"
subcategory: "Load Balancing"
description: "advertise_custom.advertise_where.advertise_dualstack_on_public for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2062, "body_sha256": "sha256:e1743f8837e4233b6491dc3e10c4b951f81c18c40de1607104239d654bc6ecd8", "child_ids": ["xcsh-docs:data-sources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_dualstack_on_public:public_ip"], "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_dualstack_on_public", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:advertise_custom:advertise_where", "path": "documentation/data-sources/tcp_loadbalancer/properties/advertise_custom/advertise_where/advertise_dualstack_on_public/index.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["advertise_custom", "advertise_where", "advertise_dualstack_on_public"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/properties/advertise_custom/advertise_where/advertise_dualstack_on_public/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advertise_custom.advertise_where.advertise_dualstack_on_public for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# advertise_custom.advertise_where.advertise_dualstack_on_public

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/)
- [advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/advertise_custom/)
- [advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/advertise_custom/advertise_where/)
- advertise_custom.advertise_where.advertise_dualstack_on_public

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

- [public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/advertise_custom/advertise_where/advertise_dualstack_on_public/public_ip/): complete subsection reference.

## Next pages

- [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/advertise_custom/advertise_where/advertise_dualstack_on_public/public_ip/)
- [advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/advertise_custom/advertise_where/)
- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/)
