---
page_title: "job.containers.readiness_check.exec_health_check"
subcategory: "Container"
description: "ExecHealthCheckType describes a health check based on \"run in container\" action. Exit status of 0 is treated as live/healthy and non-zero is unhealthy."
xcsh_docs: {"aliases": ["job containers readiness check exec health check"], "body_bytes": 2900, "body_sha256": "sha256:52374b07ea728f7da3530cbaeb64fdc29f08adc67258047f59a5257cd0f99fcc", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:containers:readiness_check:exec_health_check", "parent_id": "xcsh-docs:data-sources:workload:properties:job:containers:readiness_check", "path": "documentation/data-sources/workload/properties/job/containers/readiness_check/exec_health_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-2130032022011200-2303331333310203-1020003313311111-3112231231331220-0013211220131121-1330333022332012-0030320233303122-0211331222021333", "registry_path": "docs/guides/data-sources--workload--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "containers", "readiness_check", "exec_health_check"], "schema_version": 1, "sections": [{"aliases": ["job containers readiness check exec health check command"], "anchor": "schema-job--containers--readiness_check--exec_health_check--command", "description": "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:readiness_check:exec_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "readiness_check", "exec_health_check", "command"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/containers/readiness_check/exec_health_check/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "ExecHealthCheckType describes a health check based on \"run in container\" action. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["workloadCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.containers.readiness_check.exec_health_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/)
- [job.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/)
- [job.containers.readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/readiness_check/)
- job.containers.readiness_check.exec_health_check

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-job--containers--readiness_check--exec_health_check--command"></a>

### command property

Type: `["list", "string"]`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
