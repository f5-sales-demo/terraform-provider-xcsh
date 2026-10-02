---
page_title: "ingress_egress_gw.outside_static_routes"
subcategory: "Infrastructure"
description: "List of static routes."
xcsh_docs: {"aliases": ["ingress egress gw outside static routes"], "body_bytes": 1876, "body_sha256": "sha256:9377663cfdd57ded5284494f2420e921352d39498073e53406093824e6569275", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:outside_static_routes:static_route_list"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:outside_static_routes", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw", "path": "documentation/resources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2302303312213120-1201320022111332-2201301100330021-0221323202231013-1230113203322002-2002310033132221-0222221231011032-1323123322221222", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.outside_static_routes:RequiredObjectAttributes:static_route_list", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:outside_static_routes:static_route_list", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "outside_static_routes"], "schema_version": 1, "sections": [{"aliases": ["static route list"], "anchor": "section", "description": "List of Static routes.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:outside_static_routes:static_route_list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-ingress_egress_gw--outside_static_routes--static_route_list--simple_static_route", "enforcement": "provider-schema", "group": "ingress_egress_gw.outside_static_routes.static_route_list:ConflictingListObjectAttributes:custom_static_route,simple_static_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:outside_static_routes:static_route_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.outside_static_routes.static_route_list:ConflictingListObjectAttributes:custom_static_route,simple_static_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:outside_static_routes:static_route_list:custom_static_route", "type": "conflicts"}], "schema_path": ["ingress_egress_gw", "outside_static_routes", "static_route_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.outside_static_routes

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/)
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

- [static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/): complete subsection reference.

## Next pages

- [ingress_egress_gw.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
