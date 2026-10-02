---
page_title: "job.containers.readiness_check.tcp_health_check"
subcategory: "Container"
description: "TCPHealthCheckType describes a health check based on opening a TCP connection."
xcsh_docs: {"aliases": ["job containers readiness check tcp health check"], "body_bytes": 1895, "body_sha256": "sha256:87455b0ef65bdd49fbcfa6804da0e79b6ce7d1c79b617b89be0856a744cd13a6", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:job:containers:readiness_check:tcp_health_check:port"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:containers:readiness_check:tcp_health_check", "parent_id": "xcsh-docs:resources:workload:properties:job:containers:readiness_check", "path": "documentation/resources/workload/properties/job/containers/readiness_check/tcp_health_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1331303311323011-1231220102032100-0033122102322121-0323121321110233-3230223220133313-1001001130132012-1012102132021112-3000233323110222", "registry_path": "docs/guides/resources--workload--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "containers", "readiness_check", "tcp_health_check"], "schema_version": 1, "sections": [{"aliases": ["port"], "anchor": "section", "description": "Port", "document_id": "xcsh-docs:resources:workload:properties:job:containers:readiness_check:tcp_health_check:port", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-job--containers--readiness_check--tcp_health_check--port--name", "enforcement": "provider-schema", "group": "job.containers.readiness_check.tcp_health_check.port:ConflictingObjectAttributes:name,num", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:readiness_check:tcp_health_check:port", "type": "conflicts"}, {"anchor": "schema-job--containers--readiness_check--tcp_health_check--port--num", "enforcement": "provider-schema", "group": "job.containers.readiness_check.tcp_health_check.port:ConflictingObjectAttributes:name,num", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:readiness_check:tcp_health_check:port", "type": "conflicts"}], "schema_path": ["job", "containers", "readiness_check", "tcp_health_check", "port"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/containers/readiness_check/tcp_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "TCPHealthCheckType describes a health check based on opening a TCP connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.containers.readiness_check.tcp_health_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/)
- [job.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/)
- [job.containers.readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/readiness_check/)
- job.containers.readiness_check.tcp_health_check

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

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/readiness_check/tcp_health_check/port/): complete subsection reference.

## Next pages

- [job.containers.readiness_check.tcp_health_check.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/readiness_check/tcp_health_check/port/)
- [job.containers.readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/readiness_check/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
