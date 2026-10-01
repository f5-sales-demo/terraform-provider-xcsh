---
page_title: "job"
subcategory: "Container"
description: "job for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3277, "body_sha256": "sha256:ad632a5e86bbc5ad55bcd8b3827d80bb399663834f8559de5d86c22c137b13dd", "canonical_id": "xcsh-docs:resources:workload:properties:job", "child_ids": ["xcsh-docs:resources:workload:properties:job:configuration", "xcsh-docs:resources:workload:properties:job:containers", "xcsh-docs:resources:workload:properties:job:deploy_options", "xcsh-docs:resources:workload:properties:job:volumes"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job", "parent_id": "xcsh-docs:resources:workload:reference", "path": "docs/guides/resources--workload--properties--job.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["job"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "job for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- job

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: job, service, simple\_service, stateful\_service\] Jobs are used for running batch
processing tasks and run to completion. Jobs are generally used for tasks like report generation,
billing, parallel data processing, ETL processing, etc.

Upstream description:

Jobs are used for running batch processing tasks and run to completion. Jobs are generally used for
tasks like report generation, billing, parallel data processing, ETL processing, etc.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("containers")}
```

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

- [job](resources--workload--properties--job.md#section)
- [service](resources--workload--properties--service.md#section)
- [simple_service](resources--workload--properties--simple_service.md#section)
- [stateful_service](resources--workload--properties--stateful_service.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
job {
  # Configure direct properties listed below.
}
```

## Direct properties

- [configuration](resources--workload--properties--job--configuration.md): complete subsection reference.

- [containers](resources--workload--properties--job--containers.md): complete subsection reference.

- [deploy_options](resources--workload--properties--job--deploy_options.md): complete subsection reference.

<a id="schema-job--num_replicas"></a>

### num_replicas property

Type: `"number"`. Optional.

Number of replicas of the batch job to spawn per site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(5),
}
```

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

- [volumes](resources--workload--properties--job--volumes.md): complete subsection reference.

## Next pages

- [job.configuration](resources--workload--properties--job--configuration.md)
- [job.containers](resources--workload--properties--job--containers.md)
- [job.deploy_options](resources--workload--properties--job--deploy_options.md)
- [job.volumes](resources--workload--properties--job--volumes.md)
- [Property reference](resources--workload--reference.md)
- [xcsh_workload](../resources/workload.md)
