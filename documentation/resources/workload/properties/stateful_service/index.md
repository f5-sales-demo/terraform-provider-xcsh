---
page_title: "stateful_service"
subcategory: "Container"
description: "StatefulService maintains per replica state and each replica has its own persistent storage. Each replica has a unique network identity and stable storage. Stateful service are used for distributed stateful applications like cassandra, mongodb, redis, etc."
xcsh_docs: {"aliases": ["stateful service"], "body_bytes": 3140, "body_sha256": "sha256:a20e65d205d35ba42b3dd670153be51a4ba32c1455c0ea620939d10fd16e4d8d", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options", "xcsh-docs:resources:workload:properties:stateful_service:configuration", "xcsh-docs:resources:workload:properties:stateful_service:containers", "xcsh-docs:resources:workload:properties:stateful_service:deploy_options", "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes", "xcsh-docs:resources:workload:properties:stateful_service:scale_to_zero", "xcsh-docs:resources:workload:properties:stateful_service:volumes"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service", "parent_id": "xcsh-docs:resources:workload:reference", "path": "documentation/resources/workload/properties/stateful_service/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333", "registry_path": "docs/guides/resources--workload--reference--group-018.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service"], "schema_version": 1, "sections": [{"aliases": ["stateful service advertise options"], "anchor": "section", "description": "Advertise OPTIONS are used to configure how and where to advertise the workload using load balancers.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "advertise_options"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service configuration"], "anchor": "section", "description": "Configuration parameters of the workload.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:configuration", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "configuration"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service containers"], "anchor": "section", "description": "Containers to use for service.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["stateful_service", "containers"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service deploy options"], "anchor": "section", "description": "Deploy OPTIONS are used to configure the workload deployment OPTIONS.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "deploy_options"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service num replicas"], "anchor": "schema-stateful_service--num_replicas", "description": "Exclusive with Number of replicas of service to spawn per site.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "num_replicas"], "syntax": "attribute", "type": "number"}, {"aliases": ["stateful service persistent volumes"], "anchor": "section", "description": "Persistent storage configuration for the service.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["stateful_service", "persistent_volumes"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service scale to zero"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:scale_to_zero", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "scale_to_zero"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service volumes"], "anchor": "section", "description": "Ephemeral volumes for the service.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:volumes", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["stateful_service", "volumes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "StatefulService maintains per replica state and each replica has its own persistent storage. Each replica has a unique network identity and stable storage. Stateful service are used for distributed stateful applications like cassandra, mongodb, redis, etc.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["workloadCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
