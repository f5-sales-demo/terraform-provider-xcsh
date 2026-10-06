---
page_title: "simple_service.container.liveness_check.tcp_health_check"
subcategory: "Container"
description: "TCPHealthCheckType describes a health check based on opening a TCP connection."
xcsh_docs: {"aliases": ["simple service container liveness check tcp health check"], "body_bytes": 1519, "body_sha256": "sha256:e1b5d1853b7bcff49c7cfbac20217928c92a77c1f85d2c84c1ab0bb5c6920793", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:tcp_health_check:port"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:tcp_health_check", "parent_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check", "path": "documentation/resources/workload/properties/simple_service/container/liveness_check/tcp_health_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1321011121310301-2221231110312101-2203132301320030-3030201110120231-3002002032122030-1130100002311331-3130212221120312-1230122223333111", "registry_path": "docs/guides/resources--workload--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service", "container", "liveness_check", "tcp_health_check"], "schema_version": 1, "sections": [{"aliases": ["simple service container liveness check tcp health check port"], "anchor": "section", "description": "Port", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:tcp_health_check:port", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-simple_service--container--liveness_check--tcp_health_check--port--name", "enforcement": "provider-schema", "group": "simple_service.container.liveness_check.tcp_health_check.port:ConflictingObjectAttributes:name,num", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:tcp_health_check:port", "type": "conflicts"}, {"anchor": "schema-simple_service--container--liveness_check--tcp_health_check--port--num", "enforcement": "provider-schema", "group": "simple_service.container.liveness_check.tcp_health_check.port:ConflictingObjectAttributes:name,num", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:tcp_health_check:port", "type": "conflicts"}], "schema_path": ["simple_service", "container", "liveness_check", "tcp_health_check", "port"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/simple_service/container/liveness_check/tcp_health_check/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "TCPHealthCheckType describes a health check based on opening a TCP connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.container.liveness_check.tcp_health_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/)
- [simple_service.container](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/)
- [simple_service.container.liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/liveness_check/)
- simple_service.container.liveness_check.tcp_health_check

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

TCPHealthCheckType describes a health check based on opening a TCP connection.

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
tcp_health_check {
  # Configure direct properties listed below.
}
```

## Direct properties

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/liveness_check/tcp_health_check/port/): complete subsection reference.
