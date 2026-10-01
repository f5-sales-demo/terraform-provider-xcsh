---
page_title: "voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6"
subcategory: "Infrastructure"
description: "voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6 for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 4425, "body_sha256": "sha256:7b84b666d6da1fc839edde1fcf951cac0d4172bf634768bbe0b890e1327b91c9", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:subnets:ipv6", "child_ids": [], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:subnets:ipv6", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:subnets", "path": "docs/guides/resources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster_ar", "outside_static_routes", "static_route_list", "custom_static_route", "subnets", "ipv6"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6 for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [voltstack_cluster_ar](resources--azure_vnet_site--properties--voltstack_cluster_ar.md)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes.md)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list.md)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route.md)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets.md)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

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
ipv6 {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen"></a>

### plen property

Type: `"number"`. Optional.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix"></a>

### prefix property

Type: `"string"`. Optional.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

## Next pages

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
