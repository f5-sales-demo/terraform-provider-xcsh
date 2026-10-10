---
page_title: "routes"
subcategory: "Load Balancing"
description: "Routes allow users to define match condition on a path and/or HTTP method to either forward matching traffic to origin pool or redirect matching traffic to a different URL or respond directly to matching traffic."
xcsh_docs: {"aliases": ["routes"], "body_bytes": 2591, "body_sha256": "sha256:852899b2aeee9064d736ec10911ad3e739f3857348b4d096242e378bee8a0314", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object", "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route", "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route", "xcsh-docs:resources:http_loadbalancer:properties:routes:route_state_disabled", "xcsh-docs:resources:http_loadbalancer:properties:routes:route_state_enabled", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/routes/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-024.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes"], "schema_version": 1, "sections": [{"aliases": ["routes custom route object"], "anchor": "section", "description": "A custom route uses a route object created outside of this view.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "custom_route_object"], "syntax": "block", "type": "object"}, {"aliases": ["routes direct response route"], "anchor": "section", "description": "A direct response route matches on path, incoming header, incoming port and/or HTTP method and responds directly to the matching traffic.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "direct_response_route"], "syntax": "block", "type": "object"}, {"aliases": ["routes redirect route"], "anchor": "section", "description": "A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the matching traffic to a different URL.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "redirect_route"], "syntax": "block", "type": "object"}, {"aliases": ["routes route state disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:route_state_disabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_state_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route state enabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:route_state_enabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_state_enabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes simple route"], "anchor": "section", "description": "A simple route matches on path, incoming header, incoming port and/or HTTP method and forwards the matching traffic to the associated pools.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "simple_route"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Routes allow users to define match condition on a path and/or HTTP method to either forward matching traffic to origin pool or redirect matching traffic to a different URL or respond directly to matching traffic.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- routes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Terraform syntax:

```terraform
routes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_route_object](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/custom_route_object/): complete subsection reference.

- [direct_response_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/direct_response_route/): complete subsection reference.

- [redirect_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/redirect_route/): complete subsection reference.

- [route_state_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/route_state_disabled/): complete subsection reference.

- [route_state_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/route_state_enabled/): complete subsection reference.

- [simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/): complete subsection reference.
