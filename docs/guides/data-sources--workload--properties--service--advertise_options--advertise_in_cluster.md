---
page_title: "service.advertise_options.advertise_in_cluster"
subcategory: "Container"
description: "service.advertise_options.advertise_in_cluster for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1598, "body_sha256": "sha256:a39efcc22fd961ed9c0359015a1cf5efdd99fea5f35df94d269e547d939558e2", "canonical_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:port"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options", "path": "docs/guides/data-sources--workload--properties--service--advertise_options--advertise_in_cluster.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_in_cluster"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_in_cluster/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_in_cluster for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_in_cluster

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [service](data-sources--workload--properties--service.md)
- [service.advertise_options](data-sources--workload--properties--service--advertise_options.md)
- service.advertise_options.advertise_in_cluster

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

- [multi_ports](data-sources--workload--properties--service--advertise_options--advertise_in_cluster--multi_ports.md): complete subsection reference.

- [port](data-sources--workload--properties--service--advertise_options--advertise_in_cluster--port.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_in_cluster.multi_ports](data-sources--workload--properties--service--advertise_options--advertise_in_cluster--multi_ports.md)
- [service.advertise_options.advertise_in_cluster.port](data-sources--workload--properties--service--advertise_options--advertise_in_cluster--port.md)
- [service.advertise_options](data-sources--workload--properties--service--advertise_options.md)
- [xcsh_workload](../data-sources/workload.md)
