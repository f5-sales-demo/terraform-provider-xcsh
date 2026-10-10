---
page_title: "stateful_service.advertise_options.advertise_on_public.port"
subcategory: "Container"
description: "Advertise single port."
xcsh_docs: {"aliases": ["stateful service advertise options advertise on public port"], "body_bytes": 2058, "body_sha256": "sha256:a9a7458908766272f5ac178cf3e938558258b7c40a35e76ff9c3ab35569f72e5", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:port", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:tcp_loadbalancer"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public", "path": "documentation/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011", "registry_path": "docs/guides/resources--workload--reference--group-025.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port"], "schema_version": 1, "sections": [{"aliases": ["stateful service advertise options advertise on public port http loadbalancer"], "anchor": "section", "description": "HTTP/HTTPS Load balancer.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service advertise options advertise on public port port"], "anchor": "section", "description": "Single port.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:port", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "port"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service advertise options advertise on public port tcp loadbalancer"], "anchor": "section", "description": "TCP loadbalancer.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:tcp_loadbalancer", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "tcp_loadbalancer"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Advertise single port.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["workloadCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_on_public.port

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/)
- stateful_service.advertise_options.advertise_on_public.port

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Advertise Port. Advertise single port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"http_loadbalancer\",\"tcp_loadbalancer\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

## Direct properties

- [http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/): complete subsection reference.

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/port/): complete subsection reference.

- [tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/tcp_loadbalancer/): complete subsection reference.
