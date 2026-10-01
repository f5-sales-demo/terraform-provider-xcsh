---
page_title: "simple_service.container.liveness_check.tcp_health_check"
subcategory: "Container"
description: "simple_service.container.liveness_check.tcp_health_check for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1570, "body_sha256": "sha256:54e4733e08365f4ae6662008ea8b683fb6d0b6674971db64650a28057f15761b", "canonical_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:tcp_health_check", "child_ids": ["xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:tcp_health_check:port"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:tcp_health_check", "parent_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check", "path": "docs/guides/resources--workload--properties--simple_service--container--liveness_check--tcp_health_check.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["simple_service", "container", "liveness_check", "tcp_health_check"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/simple_service/container/liveness_check/tcp_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "simple_service.container.liveness_check.tcp_health_check for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.container.liveness_check.tcp_health_check

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [simple_service](resources--workload--properties--simple_service.md)
- [simple_service.container](resources--workload--properties--simple_service--container.md)
- [simple_service.container.liveness_check](resources--workload--properties--simple_service--container--liveness_check.md)
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

- [port](resources--workload--properties--simple_service--container--liveness_check--tcp_health_check--port.md): complete subsection reference.

## Next pages

- [simple_service.container.liveness_check.tcp_health_check.port](resources--workload--properties--simple_service--container--liveness_check--tcp_health_check--port.md)
- [simple_service.container.liveness_check](resources--workload--properties--simple_service--container--liveness_check.md)
- [xcsh_workload](../resources/workload.md)
