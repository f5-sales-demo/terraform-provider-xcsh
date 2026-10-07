---
page_title: "simple_service.container.readiness_check.tcp_health_check"
subcategory: "Container"
description: "TCPHealthCheckType describes a health check based on opening a TCP connection."
xcsh_docs: {"aliases": ["simple service container readiness check tcp health check"], "body_bytes": 1524, "body_sha256": "sha256:291d7cb16794902fde2991641e161ca8954fe71b0b45fbd0d890fdf8de01050e", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:tcp_health_check:port"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:tcp_health_check", "parent_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check", "path": "documentation/resources/workload/properties/simple_service/container/readiness_check/tcp_health_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2120132203003331-2020222300211130-1230021011312132-1001012000221222-1131332001122301-3000222103312120-1010001332230132-0210313102002011", "registry_path": "docs/guides/resources--workload--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service", "container", "readiness_check", "tcp_health_check"], "schema_version": 1, "sections": [{"aliases": ["simple service container readiness check tcp health check port"], "anchor": "section", "description": "Port", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:tcp_health_check:port", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-simple_service--container--readiness_check--tcp_health_check--port--name", "enforcement": "provider-schema", "group": "simple_service.container.readiness_check.tcp_health_check.port:ConflictingObjectAttributes:name,num", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:tcp_health_check:port", "type": "conflicts"}, {"anchor": "schema-simple_service--container--readiness_check--tcp_health_check--port--num", "enforcement": "provider-schema", "group": "simple_service.container.readiness_check.tcp_health_check.port:ConflictingObjectAttributes:name,num", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:tcp_health_check:port", "type": "conflicts"}], "schema_path": ["simple_service", "container", "readiness_check", "tcp_health_check", "port"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/simple_service/container/readiness_check/tcp_health_check/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "TCPHealthCheckType describes a health check based on opening a TCP connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["workloadCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.container.readiness_check.tcp_health_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/)
- [simple_service.container](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/)
- [simple_service.container.readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/readiness_check/)
- simple_service.container.readiness_check.tcp_health_check

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

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/readiness_check/tcp_health_check/port/): complete subsection reference.
