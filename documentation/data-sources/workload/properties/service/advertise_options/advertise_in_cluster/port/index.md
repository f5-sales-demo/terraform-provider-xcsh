---
page_title: "service.advertise_options.advertise_in_cluster.port"
subcategory: "Container"
description: "Single port."
xcsh_docs: {"aliases": ["service advertise options advertise in cluster port"], "body_bytes": 1339, "body_sha256": "sha256:0df16902676caff99c96ce2bd3af716f759995b6f7c07e53cc9206b23a7bc982", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:port:info"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:port", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster", "path": "documentation/data-sources/workload/properties/service/advertise_options/advertise_in_cluster/port/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2223131103313322-0013111330301232-2313200302013312-2300312202002333-1313130002133111-1321232321000201-1003101301002233-3322012010032123", "registry_path": "docs/guides/data-sources--workload--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_in_cluster", "port"], "schema_version": 1, "sections": [{"aliases": ["service advertise options advertise in cluster port info"], "anchor": "section", "description": "Port information.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:port:info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_in_cluster", "port", "info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_in_cluster/port/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Single port.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["workloadCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_in_cluster.port

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_in_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_in_cluster/)
- service.advertise_options.advertise_in_cluster.port

<a id="section"></a>

Type: `"single"`. Computed.

Port. Single port.

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

- [info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_in_cluster/port/info/): complete subsection reference.
