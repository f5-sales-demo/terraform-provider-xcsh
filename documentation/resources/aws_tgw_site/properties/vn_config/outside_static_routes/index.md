---
page_title: "vn_config.outside_static_routes"
subcategory: ""
description: "List of static routes."
xcsh_docs: {"aliases": ["vn config outside static routes"], "body_bytes": 1383, "body_sha256": "sha256:70bec2eb319fd47955ffff675fb0b1b5ebe5029a8ec48cd03213f19f9d9b4411", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config", "path": "documentation/resources/aws_tgw_site/properties/vn_config/outside_static_routes/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.outside_static_routes:RequiredObjectAttributes:static_route_list", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["vn_config", "outside_static_routes"], "schema_version": 1, "sections": [{"aliases": ["vn config outside static routes static route list"], "anchor": "section", "description": "List of Static routes.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-vn_config--outside_static_routes--static_route_list--simple_static_route", "enforcement": "provider-schema", "group": "vn_config.outside_static_routes.static_route_list:ConflictingListObjectAttributes:custom_static_route,simple_static_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.outside_static_routes.static_route_list:ConflictingListObjectAttributes:custom_static_route,simple_static_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route", "type": "conflicts"}], "schema_path": ["vn_config", "outside_static_routes", "static_route_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/outside_static_routes/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.outside_static_routes

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/)
- vn_config.outside_static_routes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside static routes.

Additional upstream details:

List of static routes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/): complete subsection reference.
