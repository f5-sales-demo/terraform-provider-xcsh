---
page_title: "ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4"
subcategory: "Infrastructure"
description: "ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4 for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 4282, "body_sha256": "sha256:4f5285b37b8aa4c81b6fe6775bb66c184eb0dee3a783103ff0fb70363a3da9dc", "child_ids": [], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_static_routes:static_route_list:custom_static_route:subnets", "path": "documentation/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/index.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["ingress_egress_gw", "outside_static_routes", "static_route_list", "custom_static_route", "subnets", "ipv4"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4 for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/)
- [ingress_egress_gw.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

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

<a id="schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen"></a>

### plen property

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix"></a>

### prefix property

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
