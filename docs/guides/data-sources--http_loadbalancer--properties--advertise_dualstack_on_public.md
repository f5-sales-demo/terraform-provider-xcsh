---
page_title: "advertise_dualstack_on_public"
subcategory: "Load Balancing"
description: "advertise_dualstack_on_public for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1343, "body_sha256": "sha256:f30390989d646d7dbc1e01eec5e84eae44d8a2b115bc78ba97076ee730a5ab7a", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:advertise_dualstack_on_public", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:advertise_dualstack_on_public:public_ip"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:advertise_dualstack_on_public", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "docs/guides/data-sources--http_loadbalancer--properties--advertise_dualstack_on_public.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advertise_dualstack_on_public"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/advertise_dualstack_on_public/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advertise_dualstack_on_public for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_dualstack_on_public

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- advertise_dualstack_on_public

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

- [public_ip](data-sources--http_loadbalancer--properties--advertise_dualstack_on_public--public_ip.md): complete subsection reference.

## Next pages

- [advertise_dualstack_on_public.public_ip](data-sources--http_loadbalancer--properties--advertise_dualstack_on_public--public_ip.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
