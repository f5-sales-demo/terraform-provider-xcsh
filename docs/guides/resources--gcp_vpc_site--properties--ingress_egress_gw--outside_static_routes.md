---
page_title: "ingress_egress_gw.outside_static_routes"
subcategory: "Infrastructure"
description: "ingress_egress_gw.outside_static_routes for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1494, "body_sha256": "sha256:e3a29c0030a7fb0f3fe4d95ca7f033be12057a4682e55e0558666b26b8f7ce54", "canonical_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_static_routes", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_static_routes:static_route_list"], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_static_routes", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw", "path": "docs/guides/resources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "outside_static_routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.outside_static_routes for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.outside_static_routes

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
- [ingress_egress_gw](resources--gcp_vpc_site--properties--ingress_egress_gw.md)
- ingress_egress_gw.outside_static_routes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_route_list")}
```

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
outside_static_routes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [static_route_list](resources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.outside_static_routes.static_route_list](resources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list.md)
- [ingress_egress_gw](resources--gcp_vpc_site--properties--ingress_egress_gw.md)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
