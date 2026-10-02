---
page_title: "vn_config.outside_static_routes.static_route_list.custom_static_route"
subcategory: ""
description: "Defines a static route, configuring a list of prefixes and a next-hop to be used for them."
xcsh_docs: {"aliases": ["vn config outside static routes static route list custom static route"], "body_bytes": 4678, "body_sha256": "sha256:baaa7bcfe58c7f8ab748100bf8f96ffdff00e39073007ef17a284f56756cb06f", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:labels", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:subnets"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list", "path": "documentation/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.outside_static_routes.static_route_list.custom_static_route:RequiredObjectAttributes:subnets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:subnets", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["vn_config", "outside_static_routes", "static_route_list", "custom_static_route"], "schema_version": 1, "sections": [{"aliases": ["attrs"], "anchor": "schema-vn_config--outside_static_routes--static_route_list--custom_static_route--attrs", "description": "List of route attributes associated with the static route.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "outside_static_routes", "static_route_list", "custom_static_route", "attrs"], "syntax": "attribute", "type": "list"}, {"aliases": ["labels"], "anchor": "section", "description": "Add Labels for this Static Route, these labels can be used in network policy.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:labels", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "outside_static_routes", "static_route_list", "custom_static_route", "labels"], "syntax": "block", "type": "object"}, {"aliases": ["nexthop"], "anchor": "section", "description": "Identifies the next-hop for a route.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop"], "syntax": "block", "type": "object"}, {"aliases": ["subnets"], "anchor": "section", "description": "List of route prefixes.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:subnets", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.outside_static_routes.static_route_list.custom_static_route.subnets:ConflictingListObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.outside_static_routes.static_route_list.custom_static_route.subnets:ConflictingListObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:subnets:ipv6", "type": "conflicts"}], "schema_path": ["vn_config", "outside_static_routes", "static_route_list", "custom_static_route", "subnets"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Defines a static route, configuring a list of prefixes and a next-hop to be used for them.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.outside_static_routes.static_route_list.custom_static_route

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/)
- [vn_config.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/)
- [vn_config.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/)
- vn_config.outside_static_routes.static_route_list.custom_static_route

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnets")}
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
custom_static_route {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-vn_config--outside_static_routes--static_route_list--custom_static_route--attrs"></a>

### attrs property

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/labels/): complete subsection reference.

- [nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/): complete subsection reference.

- [subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/subnets/): complete subsection reference.

## Next pages

- [vn_config.outside_static_routes.static_route_list.custom_static_route.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/labels/)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/subnets/)
- [vn_config.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
