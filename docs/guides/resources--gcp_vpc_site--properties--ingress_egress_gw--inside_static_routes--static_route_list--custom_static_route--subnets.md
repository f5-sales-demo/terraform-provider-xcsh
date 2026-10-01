---
page_title: "ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets"
subcategory: "Infrastructure"
description: "ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 3218, "body_sha256": "sha256:2922abe66632f6ee9109a1d0788d79a05e032a8e0829c88913c78287a6cbcb75", "canonical_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:subnets", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:subnets:ipv6"], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:subnets", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route", "path": "docs/guides/resources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list", "custom_static_route", "subnets"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
- [ingress_egress_gw](resources--gcp_vpc_site--properties--ingress_egress_gw.md)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes.md)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list.md)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route.md)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ipv4",
    "ipv6")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ipv4](resources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md): complete subsection reference.

- [ipv6](resources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route.md)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
