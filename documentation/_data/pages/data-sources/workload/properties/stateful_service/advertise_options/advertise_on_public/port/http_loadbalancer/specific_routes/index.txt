---
page_title: "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes"
subcategory: "Container"
description: "This defines various OPTIONS to define a route."
xcsh_docs: {"aliases": ["stateful service advertise options advertise on public port http loadbalancer specific routes"], "body_bytes": 2016, "body_sha256": "sha256:63b587c60ab1d6aa502790715895742f3ea54807e52641e6fd85625efe6dbdb2", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer", "path": "documentation/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/specific_routes/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123", "registry_path": "docs/guides/data-sources--workload--reference--group-025.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "specific_routes"], "schema_version": 1, "sections": [{"aliases": ["stateful service advertise options advertise on public port http loadbalancer specific routes routes"], "anchor": "section", "description": "Routes for this loadbalancer.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "specific_routes", "routes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/specific_routes/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This defines various OPTIONS to define a route.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["workloadCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/)
- [stateful_service.advertise_options.advertise_on_public.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes

<a id="section"></a>

Type: `"single"`. Computed.

This defines various OPTIONS to define a route.

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

- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/specific_routes/routes/): complete subsection reference.
