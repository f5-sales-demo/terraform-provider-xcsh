---
page_title: "ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4"
subcategory: "Infrastructure"
description: "ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 3830, "body_sha256": "sha256:1d30ff2224657635eca68c19878176a990cfcaefa9062683e954b6cd1ce8ad1a", "child_ids": [], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address", "path": "documentation/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/index.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["ingress_egress_gw", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "ipv4"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/)
- [ingress_egress_gw.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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
ipv4 {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr"></a>

### addr property

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

## Next pages

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
