---
page_title: "service"
subcategory: "Container"
description: "Service does not maintain per replica state, however it can be configured to use persistent storage that is shared amongst all the replicas. Replicas of a service are fungible and do not have a stable network identity or storage. Common examples of services are web servers, application servers, traditional SQL"
xcsh_docs: {"aliases": ["service"], "body_bytes": 2849, "body_sha256": "sha256:8b92b756f674b19ad74c85de9fab98ae13335e13983b5564d3ce91c7cea87c72", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options", "xcsh-docs:data-sources:workload:properties:service:configuration", "xcsh-docs:data-sources:workload:properties:service:containers", "xcsh-docs:data-sources:workload:properties:service:deploy_options", "xcsh-docs:data-sources:workload:properties:service:scale_to_zero", "xcsh-docs:data-sources:workload:properties:service:volumes"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service", "parent_id": "xcsh-docs:data-sources:workload:reference", "path": "documentation/data-sources/workload/properties/service/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211", "registry_path": "docs/guides/data-sources--workload--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service"], "schema_version": 1, "sections": [{"aliases": ["service advertise options"], "anchor": "section", "description": "Advertise OPTIONS are used to configure how and where to advertise the workload using load balancers.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options"], "syntax": "attribute", "type": "object"}, {"aliases": ["service configuration"], "anchor": "section", "description": "Configuration parameters of the workload.", "document_id": "xcsh-docs:data-sources:workload:properties:service:configuration", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "configuration"], "syntax": "attribute", "type": "object"}, {"aliases": ["service containers"], "anchor": "section", "description": "Containers to use for service.", "document_id": "xcsh-docs:data-sources:workload:properties:service:containers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["service", "containers"], "syntax": "attribute", "type": "object"}, {"aliases": ["service deploy options"], "anchor": "section", "description": "Deploy OPTIONS are used to configure the workload deployment OPTIONS.", "document_id": "xcsh-docs:data-sources:workload:properties:service:deploy_options", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "deploy_options"], "syntax": "attribute", "type": "object"}, {"aliases": ["service num replicas"], "anchor": "schema-service--num_replicas", "description": "Exclusive with Number of replicas of service to spawn per site.", "document_id": "xcsh-docs:data-sources:workload:properties:service", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "num_replicas"], "syntax": "attribute", "type": "number"}, {"aliases": ["service scale to zero"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:service:scale_to_zero", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "scale_to_zero"], "syntax": "attribute", "type": "object"}, {"aliases": ["service volumes"], "anchor": "section", "description": "Volumes for the service.", "document_id": "xcsh-docs:data-sources:workload:properties:service:volumes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["service", "volumes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Service does not maintain per replica state, however it can be configured to use persistent storage that is shared amongst all the replicas. Replicas of a service are fungible and do not have a stable network identity or storage. Common examples of services are web servers, application servers, traditional SQL", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- service

<a id="section"></a>

Type: `"single"`. Computed.

Service does not maintain per replica state, however it can be configured to use persistent storage
that is shared amongst all the replicas. Replicas of a service are fungible and do not have a stable
network identity or storage. Common examples of services are web servers, application servers,
traditional SQL databases, etc.

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

- [advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/): complete subsection reference.

- [configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/configuration/): complete subsection reference.

- [containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/containers/): complete subsection reference.

- [deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/deploy_options/): complete subsection reference.

<a id="schema-service--num_replicas"></a>

### num_replicas property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [scale_to_zero](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/scale_to_zero/): complete subsection reference.

- [volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/volumes/): complete subsection reference.
