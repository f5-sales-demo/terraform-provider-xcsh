---
page_title: "job.containers.liveness_check.tcp_health_check"
subcategory: "Container"
description: "job.containers.liveness_check.tcp_health_check for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1339, "body_sha256": "sha256:bd4d030faa28cf4feacf0a32ad16d90b273baf5fc19b45f18f92651b9dc2bf6e", "canonical_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check:tcp_health_check", "child_ids": ["xcsh-docs:resources:workload:properties:job:containers:liveness_check:tcp_health_check:port"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check:tcp_health_check", "parent_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check", "path": "docs/guides/resources--workload--properties--job--containers--liveness_check--tcp_health_check.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["job", "containers", "liveness_check", "tcp_health_check"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/containers/liveness_check/tcp_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "job.containers.liveness_check.tcp_health_check for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# job.containers.liveness_check.tcp_health_check

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [job](resources--workload--properties--job.md)
- [job.containers](resources--workload--properties--job--containers.md)
- [job.containers.liveness_check](resources--workload--properties--job--containers--liveness_check.md)
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

- [port](resources--workload--properties--job--containers--liveness_check--tcp_health_check--port.md): complete subsection reference.

## Next pages

- [job.containers.liveness_check.tcp_health_check.port](resources--workload--properties--job--containers--liveness_check--tcp_health_check--port.md)
- [job.containers.liveness_check](resources--workload--properties--job--containers--liveness_check.md)
- [xcsh_workload](../resources/workload.md)
