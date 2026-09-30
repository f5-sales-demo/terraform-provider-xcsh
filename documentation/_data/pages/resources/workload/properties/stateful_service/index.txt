---
page_title: "stateful_service"
subcategory: "Container"
description: "stateful_service for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 5118, "body_sha256": "sha256:f8b29311076167b0169ca9811a96df2f913f16a7aa4bbf41442726303fac7d55", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options", "xcsh-docs:resources:workload:properties:stateful_service:configuration", "xcsh-docs:resources:workload:properties:stateful_service:containers", "xcsh-docs:resources:workload:properties:stateful_service:deploy_options", "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes", "xcsh-docs:resources:workload:properties:stateful_service:scale_to_zero", "xcsh-docs:resources:workload:properties:stateful_service:volumes"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service", "parent_id": "xcsh-docs:resources:workload:reference", "path": "documentation/resources/workload/properties/stateful_service/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["stateful_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
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

- [advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/): complete subsection reference.

- [configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/configuration/): complete subsection reference.

- [containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/): complete subsection reference.

- [deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/deploy_options/): complete subsection reference.

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

- [persistent_volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/persistent_volumes/): complete subsection reference.

- [scale_to_zero](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/scale_to_zero/): complete subsection reference.

- [volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/volumes/): complete subsection reference.

## Next pages

- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/configuration/)
- [stateful_service.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/)
- [stateful_service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/deploy_options/)
- [stateful_service.persistent_volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/persistent_volumes/)
- [stateful_service.scale_to_zero](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/scale_to_zero/)
- [stateful_service.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/volumes/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
