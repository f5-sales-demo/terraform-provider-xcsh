---
page_title: "routes.simple_route.advanced_options.mirror_policy.percent"
subcategory: "Load Balancing"
description: "Fraction used where sampling percentages are needed. Example sampled requests."
xcsh_docs: {"aliases": ["routes simple route advanced options mirror policy percent"], "body_bytes": 3303, "body_sha256": "sha256:9bb2238b942a56d005a1f71f53b365e02f80630a1fc07ada7b48767a8451dfb0", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy:percent", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy", "path": "documentation/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/mirror_policy/percent/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1012222121333123-3110111322300322-2102332111130002-3310322123012232-1323031012103313-0003010210312232-0010110101031311-2122131332033103", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-024.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "mirror_policy", "percent"], "schema_version": 1, "sections": [{"aliases": ["routes simple route advanced options mirror policy percent denominator"], "anchor": "schema-routes--simple_route--advanced_options--mirror_policy--percent--denominator", "description": "Denominator used in fraction where sampling percentages are needed. Example sampled requests Use hundred as denominator Use ten thousand as denominator Use million as denominator.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy:percent", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "mirror_policy", "percent", "denominator"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes simple route advanced options mirror policy percent numerator"], "anchor": "schema-routes--simple_route--advanced_options--mirror_policy--percent--numerator", "description": "Sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy:percent", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "mirror_policy", "percent", "numerator"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/mirror_policy/percent/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Fraction used where sampling percentages are needed. Example sampled requests.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options.mirror_policy.percent

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/)
- [routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/)
- [routes.simple_route.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/)
- [routes.simple_route.advanced_options.mirror_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/mirror_policy/)
- routes.simple_route.advanced_options.mirror_policy.percent

<a id="section"></a>

Type: `"single"`. Computed.

Fraction used where sampling percentages are needed. Example sampled requests.

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

<a id="schema-routes--simple_route--advanced_options--mirror_policy--percent--denominator"></a>

### denominator property

Type: `"string"`. Computed.

\[Enum: HUNDRED|TEN\_THOUSAND|MILLION\] Denominator used in fraction where sampling percentages are
needed. Example sampled requests Use hundred as denominator Use ten thousand as denominator Use
million as denominator. Possible values are \`HUNDRED\`, \`TEN\_THOUSAND\`, \`MILLION\`. Defaults to
\`HUNDRED\`.

Upstream description:

Denominator used in fraction where sampling percentages are needed. Example sampled requests

Use hundred as denominator Use ten thousand as denominator Use million as denominator.

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

Type: `"number"`. Computed.

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

- [routes.simple_route.advanced_options.mirror_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/mirror_policy/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
