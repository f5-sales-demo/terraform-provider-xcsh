---
page_title: "advertise_custom.advertise_where.advertise_v6_on_public"
subcategory: ""
description: "advertise_custom.advertise_where.advertise_v6_on_public for xcsh_udp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1831, "body_sha256": "sha256:27750db8d27b7d00c37484fa32f2ae73fedd05491f39160762a2cecabf776476", "canonical_id": "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:advertise_v6_on_public", "child_ids": ["xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:advertise_v6_on_public:public_ip"], "collection_id": "xcsh-docs:resources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:advertise_v6_on_public", "parent_id": "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where", "path": "docs/guides/resources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public.md", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advertise_custom", "advertise_where", "advertise_v6_on_public"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/udp_loadbalancer/properties/advertise_custom/advertise_where/advertise_v6_on_public/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advertise_custom.advertise_where.advertise_v6_on_public for xcsh_udp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom.advertise_where.advertise_v6_on_public

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md)
- [Property reference](resources--udp_loadbalancer--reference.md)
- [advertise_custom](resources--udp_loadbalancer--properties--advertise_custom.md)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--properties--advertise_custom--advertise_where.md)
- advertise_custom.advertise_where.advertise_v6_on_public

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
advertise_v6_on_public {
  # Configure direct properties listed below.
}
```

## Direct properties

- [public_ip](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md): complete subsection reference.

## Next pages

- [advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--properties--advertise_custom--advertise_where.md)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md)
