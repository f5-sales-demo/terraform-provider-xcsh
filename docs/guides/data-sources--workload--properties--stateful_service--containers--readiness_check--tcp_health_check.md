---
page_title: "stateful_service.containers.readiness_check.tcp_health_check"
subcategory: "Container"
description: "stateful_service.containers.readiness_check.tcp_health_check for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1518, "body_sha256": "sha256:b67aee01e281093f4980b591b2c50414e724bbed6e612a26f5934497cb567922", "canonical_id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers:readiness_check:tcp_health_check", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:containers:readiness_check:tcp_health_check:port"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers:readiness_check:tcp_health_check", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers:readiness_check", "path": "docs/guides/data-sources--workload--properties--stateful_service--containers--readiness_check--tcp_health_check.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "containers", "readiness_check", "tcp_health_check"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/containers/readiness_check/tcp_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.containers.readiness_check.tcp_health_check for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.containers.readiness_check.tcp_health_check

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [stateful_service](data-sources--workload--properties--stateful_service.md)
- [stateful_service.containers](data-sources--workload--properties--stateful_service--containers.md)
- [stateful_service.containers.readiness_check](data-sources--workload--properties--stateful_service--containers--readiness_check.md)
- stateful_service.containers.readiness_check.tcp_health_check

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

- [port](data-sources--workload--properties--stateful_service--containers--readiness_check--tcp_health_check--port.md): complete subsection reference.

## Next pages

- [stateful_service.containers.readiness_check.tcp_health_check.port](data-sources--workload--properties--stateful_service--containers--readiness_check--tcp_health_check--port.md)
- [stateful_service.containers.readiness_check](data-sources--workload--properties--stateful_service--containers--readiness_check.md)
- [xcsh_workload](../data-sources/workload.md)
