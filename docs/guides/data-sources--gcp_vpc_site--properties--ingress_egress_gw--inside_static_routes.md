---
page_title: "ingress_egress_gw.inside_static_routes"
subcategory: "Infrastructure"
description: "ingress_egress_gw.inside_static_routes for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1124, "body_sha256": "sha256:c825c957c03b612fba3ab49b43d26194e83d2a12bf57b753bd4599128614779e", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list"], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw", "path": "docs/guides/data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "inside_static_routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.inside_static_routes for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw.inside_static_routes

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
- [ingress_egress_gw](data-sources--gcp_vpc_site--properties--ingress_egress_gw.md)
- ingress_egress_gw.inside_static_routes

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for inside static routes.

Upstream description:

List of static routes.

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

- [static_route_list](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list.md)
- [ingress_egress_gw](data-sources--gcp_vpc_site--properties--ingress_egress_gw.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
