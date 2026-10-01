---
page_title: "stateful_service.advertise_options.advertise_in_cluster"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_in_cluster for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1988, "body_sha256": "sha256:dcc5e02caceaab10e4fe34eaa89715d677bd3324698dcf4eba726f036de6b32a", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:port"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options", "path": "docs/guides/resources--workload--properties--stateful_service--advertise_options--advertise_in_cluster.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_in_cluster"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_in_cluster for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_in_cluster

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](resources--workload--properties--stateful_service--advertise_options.md)
- stateful_service.advertise_options.advertise_in_cluster

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Advertise the workload locally in-cluster.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("multi_ports",
    "port")}
```

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

Terraform syntax:

```terraform
advertise_in_cluster {
  # Configure direct properties listed below.
}
```

## Direct properties

- [multi_ports](resources--workload--properties--stateful_service--advertise_options--advertise_in_cluster--multi_ports.md): complete subsection reference.

- [port](resources--workload--properties--stateful_service--advertise_options--advertise_in_cluster--port.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--properties--stateful_service--advertise_options--advertise_in_cluster--multi_ports.md)
- [stateful_service.advertise_options.advertise_in_cluster.port](resources--workload--properties--stateful_service--advertise_options--advertise_in_cluster--port.md)
- [stateful_service.advertise_options](resources--workload--properties--stateful_service--advertise_options.md)
- [xcsh_workload](../resources/workload.md)
