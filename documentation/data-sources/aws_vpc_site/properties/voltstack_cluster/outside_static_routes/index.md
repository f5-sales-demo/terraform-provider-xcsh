---
page_title: "voltstack_cluster.outside_static_routes"
subcategory: "Infrastructure"
description: "List of static routes."
xcsh_docs: {"aliases": ["voltstack cluster outside static routes"], "body_bytes": 1584, "body_sha256": "sha256:94e1e9e1cd9103ece64137837f18d642b2f3357a7db5ea1b65cda34dcf9ed01d", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster", "path": "documentation/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3330010113230211-2122000120122201-0012001202310221-2100021112121221-3131313101333213-2213112122121111-2003120322301203-1211030310321231", "registry_path": "docs/guides/data-sources--aws_vpc_site--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "outside_static_routes"], "schema_version": 1, "sections": [{"aliases": ["static route list"], "anchor": "section", "description": "List of Static routes.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["voltstack_cluster", "outside_static_routes", "static_route_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.outside_static_routes

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/)
- voltstack_cluster.outside_static_routes

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for outside static routes.

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

- [static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/): complete subsection reference.

## Next pages

- [voltstack_cluster.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
