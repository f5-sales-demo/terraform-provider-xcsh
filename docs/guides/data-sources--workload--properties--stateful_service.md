---
page_title: "stateful_service"
subcategory: "Container"
description: "stateful_service for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3852, "body_sha256": "sha256:3039a96d09e66db3b70243bc3005b249044fc5f4e8e7536110c946dc3641d06a", "canonical_id": "xcsh-docs:data-sources:workload:properties:stateful_service", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options", "xcsh-docs:data-sources:workload:properties:stateful_service:configuration", "xcsh-docs:data-sources:workload:properties:stateful_service:containers", "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options", "xcsh-docs:data-sources:workload:properties:stateful_service:persistent_volumes", "xcsh-docs:data-sources:workload:properties:stateful_service:scale_to_zero", "xcsh-docs:data-sources:workload:properties:stateful_service:volumes"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service", "parent_id": "xcsh-docs:data-sources:workload:reference", "path": "docs/guides/data-sources--workload--properties--stateful_service.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- stateful_service

<a id="section"></a>

Type: `"single"`. Computed.

StatefulService maintains per replica state and each replica has its own persistent storage. Each
replica has a unique network identity and stable storage. Stateful service are used for distributed
stateful applications like cassandra, mongodb, redis, etc.

Upstream description:

StatefulService maintains per replica state and each replica has its own persistent storage. Each
replica has a unique network identity and stable storage. Stateful service are used for distributed
stateful applications like cassandra, mongodb, redis, etc.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-scaling_choice": "[\"num_replicas\",\"scale_to_zero\"]"
}
```

## Direct properties

- [advertise_options](data-sources--workload--properties--stateful_service--advertise_options.md): complete subsection reference.

- [configuration](data-sources--workload--properties--stateful_service--configuration.md): complete subsection reference.

- [containers](data-sources--workload--properties--stateful_service--containers.md): complete subsection reference.

- [deploy_options](data-sources--workload--properties--stateful_service--deploy_options.md): complete subsection reference.

<a id="schema-stateful_service--num_replicas"></a>

### num_replicas property

Type: `"number"`. Computed.

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Upstream description:

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

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
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  }
}
```

- [persistent_volumes](data-sources--workload--properties--stateful_service--persistent_volumes.md): complete subsection reference.

- [scale_to_zero](data-sources--workload--properties--stateful_service--scale_to_zero.md): complete subsection reference.

- [volumes](data-sources--workload--properties--stateful_service--volumes.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options](data-sources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.configuration](data-sources--workload--properties--stateful_service--configuration.md)
- [stateful_service.containers](data-sources--workload--properties--stateful_service--containers.md)
- [stateful_service.deploy_options](data-sources--workload--properties--stateful_service--deploy_options.md)
- [stateful_service.persistent_volumes](data-sources--workload--properties--stateful_service--persistent_volumes.md)
- [stateful_service.scale_to_zero](data-sources--workload--properties--stateful_service--scale_to_zero.md)
- [stateful_service.volumes](data-sources--workload--properties--stateful_service--volumes.md)
- [Property reference](data-sources--workload--reference.md)
- [xcsh_workload](../data-sources/workload.md)
