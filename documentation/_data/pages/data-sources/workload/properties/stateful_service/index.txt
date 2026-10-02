---
page_title: "stateful_service"
subcategory: "Container"
description: "StatefulService maintains per replica state and each replica has its own persistent storage. Each replica has a unique network identity and stable storage. Stateful service are used for distributed stateful applications like cassandra, mongodb, redis, etc."
xcsh_docs: {"aliases": ["stateful service"], "body_bytes": 4760, "body_sha256": "sha256:768c8d6c4b0c693c68ece3c1fa372a4231b5c7e5fdbb0ee23ffb8eb797d5d2c5", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options", "xcsh-docs:data-sources:workload:properties:stateful_service:configuration", "xcsh-docs:data-sources:workload:properties:stateful_service:containers", "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options", "xcsh-docs:data-sources:workload:properties:stateful_service:persistent_volumes", "xcsh-docs:data-sources:workload:properties:stateful_service:scale_to_zero", "xcsh-docs:data-sources:workload:properties:stateful_service:volumes"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service", "parent_id": "xcsh-docs:data-sources:workload:reference", "path": "documentation/data-sources/workload/properties/stateful_service/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113", "registry_path": "docs/guides/data-sources--workload--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service"], "schema_version": 1, "sections": [{"aliases": ["advertise options"], "anchor": "section", "description": "Advertise OPTIONS are used to configure how and where to advertise the workload using load balancers.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "advertise_options"], "syntax": "attribute", "type": "object"}, {"aliases": ["configuration"], "anchor": "section", "description": "Configuration parameters of the workload.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:configuration", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "configuration"], "syntax": "attribute", "type": "object"}, {"aliases": ["containers"], "anchor": "section", "description": "Containers to use for service.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["stateful_service", "containers"], "syntax": "attribute", "type": "object"}, {"aliases": ["deploy options"], "anchor": "section", "description": "Deploy OPTIONS are used to configure the workload deployment OPTIONS.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "deploy_options"], "syntax": "attribute", "type": "object"}, {"aliases": ["num replicas"], "anchor": "schema-stateful_service--num_replicas", "description": "Exclusive with Number of replicas of service to spawn per site.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "num_replicas"], "syntax": "attribute", "type": "number"}, {"aliases": ["persistent volumes"], "anchor": "section", "description": "Persistent storage configuration for the service.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:persistent_volumes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["stateful_service", "persistent_volumes"], "syntax": "attribute", "type": "object"}, {"aliases": ["scale to zero"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:scale_to_zero", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "scale_to_zero"], "syntax": "attribute", "type": "object"}, {"aliases": ["volumes"], "anchor": "section", "description": "Ephemeral volumes for the service.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:volumes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["stateful_service", "volumes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "StatefulService maintains per replica state and each replica has its own persistent storage. Each replica has a unique network identity and stable storage. Stateful service are used for distributed stateful applications like cassandra, mongodb, redis, etc.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
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

- [advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/): complete subsection reference.

- [configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/configuration/): complete subsection reference.

- [containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/containers/): complete subsection reference.

- [deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/): complete subsection reference.

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

- [persistent_volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/persistent_volumes/): complete subsection reference.

- [scale_to_zero](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/scale_to_zero/): complete subsection reference.

- [volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/volumes/): complete subsection reference.

## Next pages

- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/configuration/)
- [stateful_service.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/containers/)
- [stateful_service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/)
- [stateful_service.persistent_volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/persistent_volumes/)
- [stateful_service.scale_to_zero](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/scale_to_zero/)
- [stateful_service.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/volumes/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
