---
page_title: "service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes"
subcategory: "Container"
description: "This defines various OPTIONS to define a route."
xcsh_docs: {"aliases": ["service advertise options advertise on public port http loadbalancer specific routes"], "body_bytes": 1999, "body_sha256": "sha256:77fd40ecc29e211431003f2890f645b23a9fe89fde17e49d61e306c5286200b6", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/specific_routes/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311", "registry_path": "docs/guides/resources--workload--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "specific_routes"], "schema_version": 1, "sections": [{"aliases": ["service advertise options advertise on public port http loadbalancer specific routes routes"], "anchor": "section", "description": "Routes for this loadbalancer.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "specific_routes", "routes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/specific_routes/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This defines various OPTIONS to define a route.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["workloadCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/)
- [service.advertise_options.advertise_on_public.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/port/)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
specific_routes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/specific_routes/routes/): complete subsection reference.
