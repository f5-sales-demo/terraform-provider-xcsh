---
page_title: "stateful_service.containers.readiness_check.tcp_health_check"
subcategory: "Container"
description: "stateful_service.containers.readiness_check.tcp_health_check for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1965, "body_sha256": "sha256:288f91cdcf5bd0c182242da068a4342bfa4fb0b343a888481afac80cc23a49eb", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:tcp_health_check:port"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:tcp_health_check", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check", "path": "documentation/resources/workload/properties/stateful_service/containers/readiness_check/tcp_health_check/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["stateful_service", "containers", "readiness_check", "tcp_health_check"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/containers/readiness_check/tcp_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.containers.readiness_check.tcp_health_check for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.containers.readiness_check.tcp_health_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/)
- [stateful_service.containers.readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/readiness_check/)
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

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/readiness_check/tcp_health_check/port/): complete subsection reference.

## Next pages

- [stateful_service.containers.readiness_check.tcp_health_check.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/readiness_check/tcp_health_check/port/)
- [stateful_service.containers.readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/readiness_check/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
