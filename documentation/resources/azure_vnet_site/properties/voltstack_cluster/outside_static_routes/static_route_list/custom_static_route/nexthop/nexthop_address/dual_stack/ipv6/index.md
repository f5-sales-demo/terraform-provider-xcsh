---
page_title: "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6"
subcategory: "Infrastructure"
description: "IPv6 Address specified as hexadecimal numbers separated by ':'"
xcsh_docs: {"aliases": ["voltstack cluster outside static routes static route list custom static route nexthop nexthop address dual stack ipv6"], "body_bytes": 4757, "body_sha256": "sha256:014de91a039bc47d7063dab702bb8729b9865dac325e48c7135ddeb9f7b665c6", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv6", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "path": "documentation/resources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1230213221310110-3023113322233310-3303130033210100-2002132003020101-1130333021332110-0200020301233102-2231313111231310-1300213302320000", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack", "ipv6"], "schema_version": 1, "sections": [{"aliases": ["addr"], "anchor": "schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr", "description": "IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by ':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv6", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack", "ipv6", "addr"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "IPv6 Address specified as hexadecimal numbers separated by ':'", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/)
- [voltstack_cluster.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/)
- [voltstack_cluster.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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

<a id="schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr"></a>

### addr property

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
