---
page_title: "stateful_service.advertise_options.advertise_on_public"
subcategory: "Container"
description: "Advertise this workload via loadbalancer on Internet with default VIP."
xcsh_docs: {"aliases": ["stateful service advertise options advertise on public"], "body_bytes": 1524, "body_sha256": "sha256:322c136df2c441b1138667bdcdbd984c22590fec22c8aff90bb55c3b64fbedd6", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options", "path": "documentation/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220", "registry_path": "docs/guides/data-sources--workload--reference--group-020.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public"], "schema_version": 1, "sections": [{"aliases": ["stateful service advertise options advertise on public multi ports"], "anchor": "section", "description": "Advertise multiple ports.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "multi_ports"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service advertise options advertise on public port"], "anchor": "section", "description": "Advertise single port.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Advertise this workload via loadbalancer on Internet with default VIP.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["workloadCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_on_public

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/)
- stateful_service.advertise_options.advertise_on_public

<a id="section"></a>

Type: `"single"`. Computed.

Advertise this workload via loadbalancer on Internet with default VIP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"multi_ports\",\"port\"]"
}
```

## Direct properties

- [multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/): complete subsection reference.

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/): complete subsection reference.
