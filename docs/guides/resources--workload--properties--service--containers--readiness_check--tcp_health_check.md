---
page_title: "service.containers.readiness_check.tcp_health_check"
subcategory: "Container"
description: "service.containers.readiness_check.tcp_health_check for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1400, "body_sha256": "sha256:2ae234a64e5dcd84926e0598d6f4d220c6a5e9fc68b9fe32572508c8252c7889", "canonical_id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check:tcp_health_check", "child_ids": ["xcsh-docs:resources:workload:properties:service:containers:readiness_check:tcp_health_check:port"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check:tcp_health_check", "parent_id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check", "path": "docs/guides/resources--workload--properties--service--containers--readiness_check--tcp_health_check.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "containers", "readiness_check", "tcp_health_check"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/containers/readiness_check/tcp_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.containers.readiness_check.tcp_health_check for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# service.containers.readiness_check.tcp_health_check

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.containers](resources--workload--properties--service--containers.md)
- [service.containers.readiness_check](resources--workload--properties--service--containers--readiness_check.md)
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

- [port](resources--workload--properties--service--containers--readiness_check--tcp_health_check--port.md): complete subsection reference.

## Next pages

- [service.containers.readiness_check.tcp_health_check.port](resources--workload--properties--service--containers--readiness_check--tcp_health_check--port.md)
- [service.containers.readiness_check](resources--workload--properties--service--containers--readiness_check.md)
- [xcsh_workload](../resources/workload.md)
