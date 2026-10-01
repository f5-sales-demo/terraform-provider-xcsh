---
page_title: "stateful_service.containers.liveness_check.tcp_health_check"
subcategory: "Container"
description: "stateful_service.containers.liveness_check.tcp_health_check for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2055, "body_sha256": "sha256:d635aa6b4188c0a66de0375c05b79fbabbf71df48427283311d49a09cbc7ecc0", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:containers:liveness_check:tcp_health_check:port"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:containers:liveness_check:tcp_health_check", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:liveness_check", "path": "documentation/resources/workload/properties/stateful_service/containers/liveness_check/tcp_health_check/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["stateful_service", "containers", "liveness_check", "tcp_health_check"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/containers/liveness_check/tcp_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.containers.liveness_check.tcp_health_check for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.containers.liveness_check.tcp_health_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/)
- [stateful_service.containers.liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/liveness_check/)
- stateful_service.containers.liveness_check.tcp_health_check

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

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/liveness_check/tcp_health_check/port/): complete subsection reference.

## Next pages

- [stateful_service.containers.liveness_check.tcp_health_check.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/liveness_check/tcp_health_check/port/)
- [stateful_service.containers.liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/liveness_check/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
