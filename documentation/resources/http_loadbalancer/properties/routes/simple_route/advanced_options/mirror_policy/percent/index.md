---
page_title: "routes.simple_route.advanced_options.mirror_policy.percent"
subcategory: "Load Balancing"
description: "Fraction used where sampling percentages are needed. Example sampled requests."
xcsh_docs: {"aliases": ["routes simple route advanced options mirror policy percent"], "body_bytes": 3715, "body_sha256": "sha256:6031edaaa3dc004195d84568f1705b3701a05aafdf553559aef3707026b00e6e", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy:percent", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy", "path": "documentation/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/mirror_policy/percent/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2313121201002301-1112103031031030-2213123011002321-0330220012031022-3302303212211102-3121203313102011-2121121230313002-3102200032303031", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-025.md", "relationships": [{"anchor": "schema-routes--simple_route--advanced_options--mirror_policy--percent--numerator", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.mirror_policy.percent:RequiredObjectAttributes:numerator", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy:percent", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "mirror_policy", "percent"], "schema_version": 1, "sections": [{"aliases": ["routes simple route advanced options mirror policy percent denominator"], "anchor": "schema-routes--simple_route--advanced_options--mirror_policy--percent--denominator", "description": "Denominator used in fraction where sampling percentages are needed. Example sampled requests Use hundred as denominator Use ten thousand as denominator Use million as denominator.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy:percent", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "mirror_policy", "percent", "denominator"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes simple route advanced options mirror policy percent numerator"], "anchor": "schema-routes--simple_route--advanced_options--mirror_policy--percent--numerator", "description": "Sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy:percent", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "mirror_policy", "percent", "numerator"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/mirror_policy/percent/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Fraction used where sampling percentages are needed. Example sampled requests.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options.mirror_policy.percent

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/)
- [routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/)
- [routes.simple_route.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/)
- [routes.simple_route.advanced_options.mirror_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/mirror_policy/)
- routes.simple_route.advanced_options.mirror_policy.percent

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Fraction used where sampling percentages are needed. Example sampled requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("numerator")}
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
percent {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--simple_route--advanced_options--mirror_policy--percent--denominator"></a>

### denominator property

Type: `"string"`. Optional.

\[Enum: HUNDRED|TEN\_THOUSAND|MILLION\] Denominator used in fraction where sampling percentages are
needed. Example sampled requests Use hundred as denominator Use ten thousand as denominator Use
million as denominator. Possible values are \`HUNDRED\`, \`TEN\_THOUSAND\`, \`MILLION\`. Defaults to
\`HUNDRED\`.

Upstream description:

Denominator used in fraction where sampling percentages are needed. Example sampled requests

Use hundred as denominator Use ten thousand as denominator Use million as denominator.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("HUNDRED",
    "TEN_THOUSAND",
    "MILLION"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "HUNDRED",
  "enum": [
    "HUNDRED",
    "TEN_THOUSAND",
    "MILLION"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-routes--simple_route--advanced_options--mirror_policy--percent--numerator"></a>

### numerator property

Type: `"number"`. Optional.

Sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

## Next pages

- [routes.simple_route.advanced_options.mirror_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/mirror_policy/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
