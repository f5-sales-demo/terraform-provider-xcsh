---
page_title: "simple_service.container.liveness_check.tcp_health_check"
subcategory: "Container"
description: "TCPHealthCheckType describes a health check based on opening a TCP connection."
xcsh_docs: {"aliases": ["simple service container liveness check tcp health check"], "body_bytes": 1412, "body_sha256": "sha256:02fedb3e544564a4662ee733580809b2206264e9bfd8f94980c68b7a1e81e54e", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:simple_service:container:liveness_check:tcp_health_check:port"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:simple_service:container:liveness_check:tcp_health_check", "parent_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:liveness_check", "path": "documentation/data-sources/workload/properties/simple_service/container/liveness_check/tcp_health_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2032231103032010-3220203132033331-3010332021232021-0030203121323211-2320331001100232-2213121310301211-0101010201120103-1212023213330120", "registry_path": "docs/guides/data-sources--workload--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service", "container", "liveness_check", "tcp_health_check"], "schema_version": 1, "sections": [{"aliases": ["simple service container liveness check tcp health check port"], "anchor": "section", "description": "Port", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:liveness_check:tcp_health_check:port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["simple_service", "container", "liveness_check", "tcp_health_check", "port"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/simple_service/container/liveness_check/tcp_health_check/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "TCPHealthCheckType describes a health check based on opening a TCP connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.container.liveness_check.tcp_health_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/)
- [simple_service.container](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/)
- [simple_service.container.liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/liveness_check/)
- simple_service.container.liveness_check.tcp_health_check

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/liveness_check/tcp_health_check/port/): complete subsection reference.
