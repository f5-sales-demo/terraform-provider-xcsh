---
page_title: "stateful_service.advertise_options.advertise_in_cluster.port.info.same_as_port"
subcategory: "Container"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["stateful service advertise options advertise in cluster port info same as port"], "body_bytes": 1880, "body_sha256": "sha256:570b18a827eebf339f962a784678238780d8e89628e7ca742a24e846e36883c5", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:port:info:same_as_port", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:port:info", "path": "documentation/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/port/info/same_as_port/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3210020200110330-3333021331320222-0230001130001233-2303222221011122-2131231021120313-1203132023111112-0320302000031110-2020300310221113", "registry_path": "docs/guides/resources--workload--reference--group-021.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_in_cluster", "port", "info", "same_as_port"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/port/info/same_as_port/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["workloadCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_in_cluster.port.info.same_as_port

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.advertise_options.advertise_in_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/)
- [stateful_service.advertise_options.advertise_in_cluster.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/port/)
- [stateful_service.advertise_options.advertise_in_cluster.port.info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/port/info/)
- stateful_service.advertise_options.advertise_in_cluster.port.info.same_as_port

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

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
same_as_port = {}
```

This is an empty object or choice marker. It has no direct properties.
