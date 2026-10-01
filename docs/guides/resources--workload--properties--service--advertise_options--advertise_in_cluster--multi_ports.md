---
page_title: "service.advertise_options.advertise_in_cluster.multi_ports"
subcategory: "Container"
description: "service.advertise_options.advertise_in_cluster.multi_ports for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1732, "body_sha256": "sha256:4929347fe58941d66bc575ca46b2c6c82724048deac420e4e634deefc8edcfc8", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_in_cluster--multi_ports.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_in_cluster", "multi_ports"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_in_cluster/multi_ports/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_in_cluster.multi_ports for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_in_cluster.multi_ports

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.advertise_options](resources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_in_cluster](resources--workload--properties--service--advertise_options--advertise_in_cluster.md)
- service.advertise_options.advertise_in_cluster.multi_ports

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Multiple Ports. Multiple ports.

Upstream description:

Multiple ports.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ports")}
```

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
multi_ports {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ports](resources--workload--properties--service--advertise_options--advertise_in_cluster--multi_ports--ports.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_in_cluster.multi_ports.ports](resources--workload--properties--service--advertise_options--advertise_in_cluster--multi_ports--ports.md)
- [service.advertise_options.advertise_in_cluster](resources--workload--properties--service--advertise_options--advertise_in_cluster.md)
- [xcsh_workload](../resources/workload.md)
