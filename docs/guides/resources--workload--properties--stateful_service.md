---
page_title: "stateful_service"
subcategory: "Container"
description: "stateful_service for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 4309, "body_sha256": "sha256:983f02364ad875026c6c792518d20670956bb6b4db39421051fa515bb6189d66", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options", "xcsh-docs:resources:workload:properties:stateful_service:configuration", "xcsh-docs:resources:workload:properties:stateful_service:containers", "xcsh-docs:resources:workload:properties:stateful_service:deploy_options", "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes", "xcsh-docs:resources:workload:properties:stateful_service:scale_to_zero", "xcsh-docs:resources:workload:properties:stateful_service:volumes"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service", "parent_id": "xcsh-docs:resources:workload:reference", "path": "docs/guides/resources--workload--properties--stateful_service.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- stateful_service

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

StatefulService maintains per replica state and each replica has its own persistent storage. Each
replica has a unique network identity and stable storage. Stateful service are used for distributed
stateful applications like cassandra, mongodb, redis, etc.

Upstream description:

StatefulService maintains per replica state and each replica has its own persistent storage. Each
replica has a unique network identity and stable storage. Stateful service are used for distributed
stateful applications like cassandra, mongodb, redis, etc.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("containers",
    "persistent_volumes"),
  validators.ConflictingObjectAttributes("num_replicas",
    "scale_to_zero")}
```

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

Terraform syntax:

```terraform
stateful_service {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advertise_options](resources--workload--properties--stateful_service--advertise_options.md): complete subsection reference.

- [configuration](resources--workload--properties--stateful_service--configuration.md): complete subsection reference.

- [containers](resources--workload--properties--stateful_service--containers.md): complete subsection reference.

- [deploy_options](resources--workload--properties--stateful_service--deploy_options.md): complete subsection reference.

<a id="schema-stateful_service--num_replicas"></a>

### num_replicas property

Type: `"number"`. Optional.

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Upstream description:

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5),
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

- [persistent_volumes](resources--workload--properties--stateful_service--persistent_volumes.md): complete subsection reference.

- [scale_to_zero](resources--workload--properties--stateful_service--scale_to_zero.md): complete subsection reference.

- [volumes](resources--workload--properties--stateful_service--volumes.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options](resources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.configuration](resources--workload--properties--stateful_service--configuration.md)
- [stateful_service.containers](resources--workload--properties--stateful_service--containers.md)
- [stateful_service.deploy_options](resources--workload--properties--stateful_service--deploy_options.md)
- [stateful_service.persistent_volumes](resources--workload--properties--stateful_service--persistent_volumes.md)
- [stateful_service.scale_to_zero](resources--workload--properties--stateful_service--scale_to_zero.md)
- [stateful_service.volumes](resources--workload--properties--stateful_service--volumes.md)
- [Property reference](resources--workload--reference.md)
- [xcsh_workload](../resources/workload.md)
