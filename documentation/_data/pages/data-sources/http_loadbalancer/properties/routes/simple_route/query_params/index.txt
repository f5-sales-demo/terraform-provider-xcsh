---
page_title: "routes.simple_route.query_params"
subcategory: "Load Balancing"
description: "Handling of incoming query parameters in simple route."
xcsh_docs: {"aliases": ["routes simple route query params"], "body_bytes": 3225, "body_sha256": "sha256:e01c31c7898dab662717c424482abb5c3242cd3d45774645f597a69f3221948a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:query_params:remove_all_params", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:query_params:retain_all_params"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:query_params", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route", "path": "documentation/data-sources/http_loadbalancer/properties/routes/simple_route/query_params/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3013022233310333-3211123033103002-2011301011131203-1221021311100202-1220322312201131-1232212212011223-1330231331301232-3032220031230320", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-025.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "simple_route", "query_params"], "schema_version": 1, "sections": [{"aliases": ["remove all params"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:query_params:remove_all_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "query_params", "remove_all_params"], "syntax": "attribute", "type": "object"}, {"aliases": ["replace params"], "anchor": "schema-routes--simple_route--query_params--replace_params", "description": "Exclusive with", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:query_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "query_params", "replace_params"], "syntax": "attribute", "type": "string"}, {"aliases": ["retain all params"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:query_params:retain_all_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "query_params", "retain_all_params"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/simple_route/query_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Handling of incoming query parameters in simple route.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.query_params

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/)
- [routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/)
- routes.simple_route.query_params

<a id="section"></a>

Type: `"single"`. Computed.

Handling of incoming query parameters in simple route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-query_params": "[\"remove_all_params\",\"replace_params\",\"retain_all_params\"]"
}
```

## Direct properties

- [remove_all_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/query_params/remove_all_params/): complete subsection reference.

<a id="schema-routes--simple_route--query_params--replace_params"></a>

### replace_params property

Type: `"string"`. Computed.

Exclusive with \[remove\_all\_params retain\_all\_params\].

Upstream description:

Exclusive with \[remove\_all\_params retain\_all\_params\]

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [retain_all_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/query_params/retain_all_params/): complete subsection reference.

## Next pages

- [routes.simple_route.query_params.remove_all_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/query_params/remove_all_params/)
- [routes.simple_route.query_params.retain_all_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/query_params/retain_all_params/)
- [routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
