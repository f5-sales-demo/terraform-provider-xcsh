---
page_title: "stateful_service.advertise_options.advertise_in_cluster"
subcategory: "Container"
description: "Advertise the workload locally in-cluster."
xcsh_docs: {"aliases": ["stateful service advertise options advertise in cluster"], "body_bytes": 2222, "body_sha256": "sha256:ca25a875dec689ae8c09c73b7de6691e7cc7d6c889fcf67af31994f93d7c500d", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:port"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_in_cluster", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options", "path": "documentation/data-sources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2000012203330031-3013213320003321-1130203203021012-0301212320010201-2013312012132313-2201113011320310-2123010110120120-3030200031122122", "registry_path": "docs/guides/data-sources--workload--reference--group-020.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_in_cluster"], "schema_version": 1, "sections": [{"aliases": ["stateful service advertise options advertise in cluster multi ports"], "anchor": "section", "description": "Multiple ports.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_in_cluster", "multi_ports"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service advertise options advertise in cluster port"], "anchor": "section", "description": "Single port.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_in_cluster", "port"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Advertise the workload locally in-cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_in_cluster

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/)
- stateful_service.advertise_options.advertise_in_cluster

<a id="section"></a>

Type: `"single"`. Computed.

Advertise the workload locally in-cluster.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"multi_ports\",\"port\"]"
}
```

## Direct properties

- [multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/multi_ports/): complete subsection reference.

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/port/): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_in_cluster.multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/multi_ports/)
- [stateful_service.advertise_options.advertise_in_cluster.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/port/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
