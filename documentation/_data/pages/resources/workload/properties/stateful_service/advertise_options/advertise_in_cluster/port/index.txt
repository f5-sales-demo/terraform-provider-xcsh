---
page_title: "stateful_service.advertise_options.advertise_in_cluster.port"
subcategory: "Container"
description: "Single port."
xcsh_docs: {"aliases": ["stateful service advertise options advertise in cluster port"], "body_bytes": 2091, "body_sha256": "sha256:3cf1b546ac50b1d5e69623e9c966884d2bb75b0ef4d0eaa407633f0fa64608eb", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:port:info"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:port", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster", "path": "documentation/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/port/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0321122020003212-1202311332110210-0323031000133111-3021322002001202-3302313330222033-1233211302112320-3101312321102200-3230320123113312", "registry_path": "docs/guides/resources--workload--reference--group-021.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_in_cluster", "port"], "schema_version": 1, "sections": [{"aliases": ["info"], "anchor": "section", "description": "Port information.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:port:info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-stateful_service--advertise_options--advertise_in_cluster--port--info--target_port", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_in_cluster.port.info:ConflictingObjectAttributes:same_as_port,target_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:port:info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_in_cluster.port.info:ConflictingObjectAttributes:same_as_port,target_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:port:info:same_as_port", "type": "conflicts"}, {"anchor": "schema-stateful_service--advertise_options--advertise_in_cluster--port--info--port", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_in_cluster.port.info:RequiredObjectAttributes:port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:port:info", "type": "requires"}], "schema_path": ["stateful_service", "advertise_options", "advertise_in_cluster", "port", "info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Single port.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_in_cluster.port

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.advertise_options.advertise_in_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/)
- stateful_service.advertise_options.advertise_in_cluster.port

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

- [info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/port/info/): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_in_cluster.port.info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/port/info/)
- [stateful_service.advertise_options.advertise_in_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
