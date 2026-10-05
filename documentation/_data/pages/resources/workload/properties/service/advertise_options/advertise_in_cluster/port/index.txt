---
page_title: "service.advertise_options.advertise_in_cluster.port"
subcategory: "Container"
description: "Single port."
xcsh_docs: {"aliases": ["service advertise options advertise in cluster port"], "body_bytes": 1974, "body_sha256": "sha256:d649f42e9ecb1d92602d9978486be466c2cdbf5711eb369732522675456886ad", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster:port:info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster:port", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_in_cluster/port/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2032101121303011-1103222233000330-0203232311001010-2322120323102000-3202033332221123-2310113223210032-0131222201221033-2112013321213130", "registry_path": "docs/guides/resources--workload--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_in_cluster", "port"], "schema_version": 1, "sections": [{"aliases": ["service advertise options advertise in cluster port info"], "anchor": "section", "description": "Port information.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster:port:info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-service--advertise_options--advertise_in_cluster--port--info--target_port", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_in_cluster.port.info:ConflictingObjectAttributes:same_as_port,target_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster:port:info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_in_cluster.port.info:ConflictingObjectAttributes:same_as_port,target_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster:port:info:same_as_port", "type": "conflicts"}, {"anchor": "schema-service--advertise_options--advertise_in_cluster--port--info--port", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_in_cluster.port.info:RequiredObjectAttributes:port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster:port:info", "type": "requires"}], "schema_path": ["service", "advertise_options", "advertise_in_cluster", "port", "info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_in_cluster/port/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Single port.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_in_cluster.port

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_in_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_in_cluster/)
- service.advertise_options.advertise_in_cluster.port

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Port. Single port.

Upstream description:

Single port.

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
port {
  # Configure direct properties listed below.
}
```

## Direct properties

- [info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_in_cluster/port/info/): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_in_cluster.port.info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_in_cluster/port/info/)
- [service.advertise_options.advertise_in_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_in_cluster/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
