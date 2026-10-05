---
page_title: "simple_service.container.liveness_check.tcp_health_check"
subcategory: "Container"
description: "TCPHealthCheckType describes a health check based on opening a TCP connection."
xcsh_docs: {"aliases": ["simple service container liveness check tcp health check"], "body_bytes": 1920, "body_sha256": "sha256:59fbe2a011d62969cdf2df390a6893ae62138b324537fabb5974b03dd51bb0dc", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:simple_service:container:liveness_check:tcp_health_check:port"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:simple_service:container:liveness_check:tcp_health_check", "parent_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:liveness_check", "path": "documentation/data-sources/workload/properties/simple_service/container/liveness_check/tcp_health_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2032231103032010-3220203132033331-3010332021232021-0030203121323211-2320331001100232-2213121310301211-0101010201120103-1212023213330120", "registry_path": "docs/guides/data-sources--workload--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service", "container", "liveness_check", "tcp_health_check"], "schema_version": 1, "sections": [{"aliases": ["simple service container liveness check tcp health check port"], "anchor": "section", "description": "Port", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:liveness_check:tcp_health_check:port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["simple_service", "container", "liveness_check", "tcp_health_check", "port"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/simple_service/container/liveness_check/tcp_health_check/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "TCPHealthCheckType describes a health check based on opening a TCP connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

## Next pages

- [simple_service.container.liveness_check.tcp_health_check.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/liveness_check/tcp_health_check/port/)
- [simple_service.container.liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/liveness_check/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
