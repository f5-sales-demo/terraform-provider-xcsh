---
page_title: "service.containers.liveness_check.exec_health_check"
subcategory: "Container"
description: "ExecHealthCheckType describes a health check based on \"run in container\" action. Exit status of 0 is treated as live/healthy and non-zero is unhealthy."
xcsh_docs: {"aliases": ["service containers liveness check exec health check"], "body_bytes": 3042, "body_sha256": "sha256:4092aa27aa54ecd5295a9bc178137d742da0c7f98bf5d87a219af73e90c6e774", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:exec_health_check", "parent_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check", "path": "documentation/resources/workload/properties/service/containers/liveness_check/exec_health_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3012230211213021-1302303300032103-0333331110210113-0212020331213003-2223223101203330-0131322021233300-0330000033011012-0100031330133333", "registry_path": "docs/guides/resources--workload--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "containers", "liveness_check", "exec_health_check"], "schema_version": 1, "sections": [{"aliases": ["service containers liveness check exec health check command"], "anchor": "schema-service--containers--liveness_check--exec_health_check--command", "description": "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that", "document_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:exec_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "containers", "liveness_check", "exec_health_check", "command"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/containers/liveness_check/exec_health_check/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "ExecHealthCheckType describes a health check based on \"run in container\" action. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["workloadCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.containers.liveness_check.exec_health_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/)
- [service.containers.liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/liveness_check/)
- service.containers.liveness_check.exec_health_check

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Additional upstream details:

ExecHealthCheckType describes a health check based on "run in container" action.

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
exec_health_check {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-service--containers--liveness_check--exec_health_check--command"></a>

### command property

Type: `["list", "string"]`. Optional.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
