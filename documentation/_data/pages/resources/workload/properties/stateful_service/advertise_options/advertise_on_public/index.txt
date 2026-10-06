---
page_title: "stateful_service.advertise_options.advertise_on_public"
subcategory: "Container"
description: "Advertise this workload via loadbalancer on Internet with default VIP."
xcsh_docs: {"aliases": ["stateful service advertise options advertise on public"], "body_bytes": 1799, "body_sha256": "sha256:bdb058196ca422f0210c33b06974d6a868f382053cab97af1b1802d1608447d5", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options", "path": "documentation/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202", "registry_path": "docs/guides/resources--workload--reference--group-021.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_on_public:ConflictingObjectAttributes:multi_ports,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_on_public:ConflictingObjectAttributes:multi_ports,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public"], "schema_version": 1, "sections": [{"aliases": ["stateful service advertise options advertise on public multi ports"], "anchor": "section", "description": "Advertise multiple ports.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_on_public.multi_ports:RequiredObjectAttributes:ports", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports", "type": "requires"}], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "multi_ports"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service advertise options advertise on public port"], "anchor": "section", "description": "Advertise single port.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_on_public.port:ConflictingObjectAttributes:http_loadbalancer,tcp_loadbalancer", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_on_public.port:ConflictingObjectAttributes:http_loadbalancer,tcp_loadbalancer", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:tcp_loadbalancer", "type": "conflicts"}], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Advertise this workload via loadbalancer on Internet with default VIP.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_on_public

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/)
- stateful_service.advertise_options.advertise_on_public

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Advertise this workload via loadbalancer on Internet with default VIP.

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
  "x-ves-oneof-field-advertise_choice": "[\"multi_ports\",\"port\"]"
}
```

Terraform syntax:

```terraform
advertise_on_public {
  # Configure direct properties listed below.
}
```

## Direct properties

- [multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/): complete subsection reference.

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/): complete subsection reference.
