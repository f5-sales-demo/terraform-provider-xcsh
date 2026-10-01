---
page_title: "advertise_on_public"
subcategory: ""
description: "advertise_on_public for xcsh_udp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1285, "body_sha256": "sha256:1b6538bb22cee43deeb44fd75420dbc1ed558c3b84e9ff7f2e7488f0fe992866", "canonical_id": "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_on_public", "child_ids": ["xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_on_public:public_ip"], "collection_id": "xcsh-docs:data-sources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_on_public", "parent_id": "xcsh-docs:data-sources:udp_loadbalancer:reference", "path": "docs/guides/data-sources--udp_loadbalancer--properties--advertise_on_public.md", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advertise_on_public"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/udp_loadbalancer/properties/advertise_on_public/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advertise_on_public for xcsh_udp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_on_public

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md)
- [Property reference](data-sources--udp_loadbalancer--reference.md)
- advertise_on_public

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

- [public_ip](data-sources--udp_loadbalancer--properties--advertise_on_public--public_ip.md): complete subsection reference.

## Next pages

- [advertise_on_public.public_ip](data-sources--udp_loadbalancer--properties--advertise_on_public--public_ip.md)
- [Property reference](data-sources--udp_loadbalancer--reference.md)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md)
