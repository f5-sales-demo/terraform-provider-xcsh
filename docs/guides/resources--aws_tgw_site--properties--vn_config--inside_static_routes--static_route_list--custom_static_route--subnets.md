---
page_title: "vn_config.inside_static_routes.static_route_list.custom_static_route.subnets"
subcategory: ""
description: "vn_config.inside_static_routes.static_route_list.custom_static_route.subnets for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 3074, "body_sha256": "sha256:87a2dbbceee4713ff30134eb4eee1f42c6fdf0a746579032454d055c46a042a7", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:subnets", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:subnets:ipv6"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:subnets", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route", "path": "docs/guides/resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vn_config", "inside_static_routes", "static_route_list", "custom_static_route", "subnets"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/subnets/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vn_config.inside_static_routes.static_route_list.custom_static_route.subnets for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.inside_static_routes.static_route_list.custom_static_route.subnets

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [vn_config](resources--aws_tgw_site--properties--vn_config.md)
- [vn_config.inside_static_routes](resources--aws_tgw_site--properties--vn_config--inside_static_routes.md)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list.md)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route.md)
- vn_config.inside_static_routes.static_route_list.custom_static_route.subnets

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

- [ipv4](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md): complete subsection reference.

- [ipv6](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md): complete subsection reference.

## Next pages

- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
