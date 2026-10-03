---
page_title: "job.containers.liveness_check.tcp_health_check"
subcategory: "Container"
description: "TCPHealthCheckType describes a health check based on opening a TCP connection."
xcsh_docs: {"aliases": ["job containers liveness check tcp health check"], "body_bytes": 1788, "body_sha256": "sha256:be5fcc5ce37d893bc94cd90640f626c90ce9206c82e92ba4964ca478f1ce6c1c", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:containers:liveness_check:tcp_health_check:port"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:containers:liveness_check:tcp_health_check", "parent_id": "xcsh-docs:data-sources:workload:properties:job:containers:liveness_check", "path": "documentation/data-sources/workload/properties/job/containers/liveness_check/tcp_health_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0323122300100223-2211211302223321-0002033002110203-3232211012022010-2322212223300100-3111230202001021-2301012232231223-2102310302132223", "registry_path": "docs/guides/data-sources--workload--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "containers", "liveness_check", "tcp_health_check"], "schema_version": 1, "sections": [{"aliases": ["job containers liveness check tcp health check port"], "anchor": "section", "description": "Port", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:liveness_check:tcp_health_check:port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "containers", "liveness_check", "tcp_health_check", "port"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/containers/liveness_check/tcp_health_check/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "TCPHealthCheckType describes a health check based on opening a TCP connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.containers.liveness_check.tcp_health_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/)
- [job.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/)
- [job.containers.liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/liveness_check/)
- job.containers.liveness_check.tcp_health_check

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

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/liveness_check/tcp_health_check/port/): complete subsection reference.

## Next pages

- [job.containers.liveness_check.tcp_health_check.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/liveness_check/tcp_health_check/port/)
- [job.containers.liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/liveness_check/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
