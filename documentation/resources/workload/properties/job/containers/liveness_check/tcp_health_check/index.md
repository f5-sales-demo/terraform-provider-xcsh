---
page_title: "job.containers.liveness_check.tcp_health_check"
subcategory: "Container"
description: "TCPHealthCheckType describes a health check based on opening a TCP connection."
xcsh_docs: {"aliases": ["job containers liveness check tcp health check"], "body_bytes": 1427, "body_sha256": "sha256:bac9efe1a7fbf2550b67d03378bd2cd032433685fc39531c273433c51785f5cf", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:job:containers:liveness_check:tcp_health_check:port"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check:tcp_health_check", "parent_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check", "path": "documentation/resources/workload/properties/job/containers/liveness_check/tcp_health_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1223231110101332-2322020122131330-1123010210221202-3020222322221232-1201100102333333-3011310033110012-1312221031213301-1100110320100300", "registry_path": "docs/guides/resources--workload--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "containers", "liveness_check", "tcp_health_check"], "schema_version": 1, "sections": [{"aliases": ["job containers liveness check tcp health check port"], "anchor": "section", "description": "Port", "document_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check:tcp_health_check:port", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "containers", "liveness_check", "tcp_health_check", "port"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/containers/liveness_check/tcp_health_check/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "TCPHealthCheckType describes a health check based on opening a TCP connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["workloadCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.containers.liveness_check.tcp_health_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/)
- [job.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/)
- [job.containers.liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/liveness_check/)
- job.containers.liveness_check.tcp_health_check

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

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/liveness_check/tcp_health_check/port/): complete subsection reference.
