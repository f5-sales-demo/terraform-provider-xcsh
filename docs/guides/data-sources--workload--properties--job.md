---
page_title: "job"
subcategory: "Container"
description: "job for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2837, "body_sha256": "sha256:a6103162511f1147c78971a95e4f89533ae85d11c296bc625d31b6b7182ed99e", "canonical_id": "xcsh-docs:data-sources:workload:properties:job", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:configuration", "xcsh-docs:data-sources:workload:properties:job:containers", "xcsh-docs:data-sources:workload:properties:job:deploy_options", "xcsh-docs:data-sources:workload:properties:job:volumes"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job", "parent_id": "xcsh-docs:data-sources:workload:reference", "path": "docs/guides/data-sources--workload--properties--job.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["job"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "job for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# job

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- job

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: job, service, simple\_service, stateful\_service\] Jobs are used for running batch
processing tasks and run to completion. Jobs are generally used for tasks like report generation,
billing, parallel data processing, ETL processing, etc.

Upstream description:

Jobs are used for running batch processing tasks and run to completion. Jobs are generally used for
tasks like report generation, billing, parallel data processing, ETL processing, etc.

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

OneOf alternatives in this subsection:

- [job](data-sources--workload--properties--job.md#section)
- [service](data-sources--workload--properties--service.md#section)
- [simple_service](data-sources--workload--properties--simple_service.md#section)
- [stateful_service](data-sources--workload--properties--stateful_service.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [configuration](data-sources--workload--properties--job--configuration.md): complete subsection reference.

- [containers](data-sources--workload--properties--job--containers.md): complete subsection reference.

- [deploy_options](data-sources--workload--properties--job--deploy_options.md): complete subsection reference.

<a id="schema-job--num_replicas"></a>

### num_replicas property

Type: `"number"`. Computed.

Number of replicas of the batch job to spawn per site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.lte": "5"
  }
}
```

- [volumes](data-sources--workload--properties--job--volumes.md): complete subsection reference.

## Next pages

- [job.configuration](data-sources--workload--properties--job--configuration.md)
- [job.containers](data-sources--workload--properties--job--containers.md)
- [job.deploy_options](data-sources--workload--properties--job--deploy_options.md)
- [job.volumes](data-sources--workload--properties--job--volumes.md)
- [Property reference](data-sources--workload--reference.md)
- [xcsh_workload](../data-sources/workload.md)
