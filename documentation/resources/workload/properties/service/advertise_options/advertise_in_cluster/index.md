---
page_title: "service.advertise_options.advertise_in_cluster"
subcategory: "Container"
description: "service.advertise_options.advertise_in_cluster for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2261, "body_sha256": "sha256:c7391aa99e127284c6082b5fee238c933873aafe13faa8ac2a397d8c9e127032", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster:port"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_in_cluster/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_in_cluster"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_in_cluster/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_in_cluster for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# service.advertise_options.advertise_in_cluster

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- service.advertise_options.advertise_in_cluster

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

- [multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_in_cluster/multi_ports/): complete subsection reference.

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_in_cluster/port/): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_in_cluster.multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_in_cluster/multi_ports/)
- [service.advertise_options.advertise_in_cluster.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_in_cluster/port/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
