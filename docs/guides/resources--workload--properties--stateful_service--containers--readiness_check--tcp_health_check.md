---
page_title: "stateful_service.containers.readiness_check.tcp_health_check"
subcategory: "Container"
description: "stateful_service.containers.readiness_check.tcp_health_check for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1517, "body_sha256": "sha256:0194ad6a3d8c73722dad97e9185cff9e66fd470b64fd5e403cd02bbf367b5642", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:tcp_health_check", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:tcp_health_check:port"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:tcp_health_check", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check", "path": "docs/guides/resources--workload--properties--stateful_service--containers--readiness_check--tcp_health_check.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "containers", "readiness_check", "tcp_health_check"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/containers/readiness_check/tcp_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.containers.readiness_check.tcp_health_check for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.containers.readiness_check.tcp_health_check

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.containers](resources--workload--properties--stateful_service--containers.md)
- [stateful_service.containers.readiness_check](resources--workload--properties--stateful_service--containers--readiness_check.md)
- stateful_service.containers.readiness_check.tcp_health_check

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

- [port](resources--workload--properties--stateful_service--containers--readiness_check--tcp_health_check--port.md): complete subsection reference.

## Next pages

- [stateful_service.containers.readiness_check.tcp_health_check.port](resources--workload--properties--stateful_service--containers--readiness_check--tcp_health_check--port.md)
- [stateful_service.containers.readiness_check](resources--workload--properties--stateful_service--containers--readiness_check.md)
- [xcsh_workload](../resources/workload.md)
