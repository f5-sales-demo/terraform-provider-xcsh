---
page_title: "voltstack_cluster.outside_static_routes"
subcategory: "Infrastructure"
description: "List of static routes."
xcsh_docs: {"aliases": ["voltstack cluster outside static routes"], "body_bytes": 1849, "body_sha256": "sha256:c35a7e39a44b6c506f9d6251efbd12cd2acbbf221a69a787fed45510d2c407c6", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster", "path": "documentation/resources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.outside_static_routes:RequiredObjectAttributes:static_route_list", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "outside_static_routes"], "schema_version": 1, "sections": [{"aliases": ["static route list"], "anchor": "section", "description": "List of Static routes.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-voltstack_cluster--outside_static_routes--static_route_list--simple_static_route", "enforcement": "provider-schema", "group": "voltstack_cluster.outside_static_routes.static_route_list:ConflictingListObjectAttributes:custom_static_route,simple_static_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.outside_static_routes.static_route_list:ConflictingListObjectAttributes:custom_static_route,simple_static_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route", "type": "conflicts"}], "schema_path": ["voltstack_cluster", "outside_static_routes", "static_route_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.outside_static_routes

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/)
- voltstack_cluster.outside_static_routes

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

- [static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/): complete subsection reference.

## Next pages

- [voltstack_cluster.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
