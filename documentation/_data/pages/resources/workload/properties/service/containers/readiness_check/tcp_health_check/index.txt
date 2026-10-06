---
page_title: "service.containers.readiness_check.tcp_health_check"
subcategory: "Container"
description: "TCPHealthCheckType describes a health check based on opening a TCP connection."
xcsh_docs: {"aliases": ["service containers readiness check tcp health check"], "body_bytes": 1468, "body_sha256": "sha256:f9af844f8c51a9aafc394282685399d36f6450fb866b7851ab5f6fd54c536c39", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:containers:readiness_check:tcp_health_check:port"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check:tcp_health_check", "parent_id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check", "path": "documentation/resources/workload/properties/service/containers/readiness_check/tcp_health_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2322321332130111-1231333003131211-1221123201022120-2003012033130020-3133021312033322-2231300213201212-1203032222120301-0102222123221323", "registry_path": "docs/guides/resources--workload--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "containers", "readiness_check", "tcp_health_check"], "schema_version": 1, "sections": [{"aliases": ["service containers readiness check tcp health check port"], "anchor": "section", "description": "Port", "document_id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check:tcp_health_check:port", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-service--containers--readiness_check--tcp_health_check--port--name", "enforcement": "provider-schema", "group": "service.containers.readiness_check.tcp_health_check.port:ConflictingObjectAttributes:name,num", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check:tcp_health_check:port", "type": "conflicts"}, {"anchor": "schema-service--containers--readiness_check--tcp_health_check--port--num", "enforcement": "provider-schema", "group": "service.containers.readiness_check.tcp_health_check.port:ConflictingObjectAttributes:name,num", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check:tcp_health_check:port", "type": "conflicts"}], "schema_path": ["service", "containers", "readiness_check", "tcp_health_check", "port"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/containers/readiness_check/tcp_health_check/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "TCPHealthCheckType describes a health check based on opening a TCP connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.containers.readiness_check.tcp_health_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/)
- [service.containers.readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/readiness_check/)
- service.containers.readiness_check.tcp_health_check

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

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/readiness_check/tcp_health_check/port/): complete subsection reference.
