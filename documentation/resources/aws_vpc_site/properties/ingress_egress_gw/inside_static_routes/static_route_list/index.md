---
page_title: "ingress_egress_gw.inside_static_routes.static_route_list"
subcategory: "Infrastructure"
description: "List of Static routes."
xcsh_docs: {"aliases": ["ingress egress gw inside static routes static route list"], "body_bytes": 3891, "body_sha256": "sha256:a0c95d114e26f50d732c18b9d818767e85c1110ccebd8fec5a4f742af941dcf4", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:inside_static_routes", "path": "documentation/resources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2122022330331131-1130130320201110-2233030222101031-0310012323130223-1102312212303231-0231300210100211-3233112222023132-3223211102222331", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-003.md", "relationships": [{"anchor": "schema-ingress_egress_gw--inside_static_routes--static_route_list--simple_static_route", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_static_routes.static_route_list:ConflictingListObjectAttributes:custom_static_route,simple_static_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_static_routes.static_route_list:ConflictingListObjectAttributes:custom_static_route,simple_static_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list"], "schema_version": 1, "sections": [{"aliases": ["custom static route"], "anchor": "section", "description": "Defines a static route, configuring a list of prefixes and a next-hop to be used for them.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route:RequiredObjectAttributes:subnets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:subnets", "type": "requires"}], "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list", "custom_static_route"], "syntax": "block", "type": "object"}, {"aliases": ["simple static route"], "anchor": "schema-ingress_egress_gw--inside_static_routes--static_route_list--simple_static_route", "description": "Exclusive with Use simple static route for prefix pointing to single interface in the network.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list", "simple_static_route"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of Static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.inside_static_routes.static_route_list

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/)
- ingress_egress_gw.inside_static_routes.static_route_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_static_route",
    "simple_static_route")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
static_route_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/): complete subsection reference.

<a id="schema-ingress_egress_gw--inside_static_routes--static_route_list--simple_static_route"></a>

### simple_static_route property

Type: `"string"`. Optional.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

## Next pages

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/)
- [ingress_egress_gw.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
