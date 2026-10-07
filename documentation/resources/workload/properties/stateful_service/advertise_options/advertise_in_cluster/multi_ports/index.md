---
page_title: "stateful_service.advertise_options.advertise_in_cluster.multi_ports"
subcategory: "Container"
description: "Multiple ports."
xcsh_docs: {"aliases": ["stateful service advertise options advertise in cluster multi ports"], "body_bytes": 1732, "body_sha256": "sha256:bb131381e764ec10bd01f2b56a1034c81dd2302f36fba65ed7cac7481f483bb6", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports:ports"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster", "path": "documentation/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/multi_ports/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1011312021130103-1132201221323011-3121333020103321-3032102100312103-3112201131010023-2123121010313000-2113022332010030-0031032321111332", "registry_path": "docs/guides/resources--workload--reference--group-021.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_in_cluster.multi_ports:RequiredObjectAttributes:ports", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports:ports", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_in_cluster", "multi_ports"], "schema_version": 1, "sections": [{"aliases": ["stateful service advertise options advertise in cluster multi ports ports"], "anchor": "section", "description": "Ports to advertise.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports:ports", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-stateful_service--advertise_options--advertise_in_cluster--multi_ports--ports--name", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports:ports", "type": "requires"}], "schema_path": ["stateful_service", "advertise_options", "advertise_in_cluster", "multi_ports", "ports"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/multi_ports/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Multiple ports.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["workloadCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_in_cluster.multi_ports

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.advertise_options.advertise_in_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/)
- stateful_service.advertise_options.advertise_in_cluster.multi_ports

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Multiple Ports. Multiple ports.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/multi_ports/ports/): complete subsection reference.
