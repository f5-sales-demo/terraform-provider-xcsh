---
page_title: "ingress_egress_gw.inside_static_routes"
subcategory: "Infrastructure"
description: "List of static routes."
xcsh_docs: {"aliases": ["ingress egress gw inside static routes"], "body_bytes": 1842, "body_sha256": "sha256:7536e8fe5986b9167e0a741a8ced8ff2740388fba4cae73b6bb01abec268802c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:inside_static_routes", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw", "path": "documentation/resources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_static_routes:RequiredObjectAttributes:static_route_list", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "inside_static_routes"], "schema_version": 1, "sections": [{"aliases": ["ingress egress gw inside static routes static route list"], "anchor": "section", "description": "List of Static routes.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-ingress_egress_gw--inside_static_routes--static_route_list--simple_static_route", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_static_routes.static_route_list:ConflictingListObjectAttributes:custom_static_route,simple_static_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_static_routes.static_route_list:ConflictingListObjectAttributes:custom_static_route,simple_static_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route", "type": "conflicts"}], "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.inside_static_routes

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/)
- ingress_egress_gw.inside_static_routes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inside static routes.

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
inside_static_routes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/): complete subsection reference.

## Next pages

- [ingress_egress_gw.inside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
