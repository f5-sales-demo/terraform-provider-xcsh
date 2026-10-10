---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes"
subcategory: "Container"
description: "Routes for this loadbalancer."
xcsh_docs: {"aliases": ["service advertise options advertise custom ports http loadbalancer specific routes routes"], "body_bytes": 3460, "body_sha256": "sha256:f88ffd6cf74fb5c31ee9dc95ab941a621ae9c87ed75fa15ad199e096ad6fdb4d", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:custom_route_object", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:direct_response_route", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:simple_route"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311", "registry_path": "docs/guides/resources--workload--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes"], "schema_version": 1, "sections": [{"aliases": ["service advertise options advertise custom ports http loadbalancer specific routes routes custom route object"], "anchor": "section", "description": "A custom route uses a route object created outside of this view.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:custom_route_object", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "custom_route_object"], "syntax": "block", "type": "object"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer specific routes routes direct response route"], "anchor": "section", "description": "A direct response route matches on path, incoming header, incoming port and/or HTTP method and responds directly to the matching traffic.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:direct_response_route", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "direct_response_route"], "syntax": "block", "type": "object"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer specific routes routes redirect route"], "anchor": "section", "description": "A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the matching traffic to a different URL.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "redirect_route"], "syntax": "block", "type": "object"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer specific routes routes simple route"], "anchor": "section", "description": "A simple route matches on path and/or HTTP method and forwards the matching traffic to the default origin pool specified outside.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:simple_route", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "simple_route"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Routes for this loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["workloadCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/)
- [service.advertise_options.advertise_custom.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Routes. Routes for this loadbalancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
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

- [custom_route_object](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/custom_route_object/): complete subsection reference.

- [direct_response_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/direct_response_route/): complete subsection reference.

- [redirect_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/redirect_route/): complete subsection reference.

- [simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/simple_route/): complete subsection reference.
