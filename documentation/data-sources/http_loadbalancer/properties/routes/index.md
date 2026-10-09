---
page_title: "routes"
subcategory: "Load Balancing"
description: "Routes allow users to define match condition on a path and/or HTTP method to either forward matching traffic to origin pool or redirect matching traffic to a different URL or respond directly to matching traffic."
xcsh_docs: {"aliases": ["routes"], "body_bytes": 2500, "body_sha256": "sha256:7ffe7ff558d9f83aee059757b46a9326fc7ba7312381445f5e0ad6f63f07cc1c", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:routes:custom_route_object", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:route_state_disabled", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:route_state_enabled", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "documentation/data-sources/http_loadbalancer/properties/routes/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-024.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes"], "schema_version": 1, "sections": [{"aliases": ["routes custom route object"], "anchor": "section", "description": "A custom route uses a route object created outside of this view.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:custom_route_object", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "custom_route_object"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes direct response route"], "anchor": "section", "description": "A direct response route matches on path, incoming header, incoming port and/or HTTP method and responds directly to the matching traffic.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "direct_response_route"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes redirect route"], "anchor": "section", "description": "A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the matching traffic to a different URL.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "redirect_route"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route state disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:route_state_disabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_state_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route state enabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:route_state_enabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_state_enabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes simple route"], "anchor": "section", "description": "A simple route matches on path, incoming header, incoming port and/or HTTP method and forwards the matching traffic to the associated pools.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "simple_route"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Routes allow users to define match condition on a path and/or HTTP method to either forward matching traffic to origin pool or redirect matching traffic to a different URL or respond directly to matching traffic.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- routes

<a id="section"></a>

Type: `"list"`. Computed.

Routes allow users to define match condition on a path and/or HTTP method to either forward matching
traffic to origin pool or redirect matching traffic to a different URL or respond directly to
matching traffic.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [custom_route_object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/custom_route_object/): complete subsection reference.

- [direct_response_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/direct_response_route/): complete subsection reference.

- [redirect_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/redirect_route/): complete subsection reference.

- [route_state_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/route_state_disabled/): complete subsection reference.

- [route_state_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/route_state_enabled/): complete subsection reference.

- [simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/): complete subsection reference.
