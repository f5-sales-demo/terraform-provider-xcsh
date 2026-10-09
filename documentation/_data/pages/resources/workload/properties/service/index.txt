---
page_title: "service"
subcategory: "Container"
description: "Service does not maintain per replica state, however it can be configured to use persistent storage that is shared amongst all the replicas. Replicas of a service are fungible and do not have a stable network identity or storage. Common examples of services are web servers, application servers, traditional SQL"
xcsh_docs: {"aliases": ["service"], "body_bytes": 2941, "body_sha256": "sha256:03b94e9a02fd9ed5c03da2b80e48ad8c1586985082129ebca59b626345c82f7e", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options", "xcsh-docs:resources:workload:properties:service:configuration", "xcsh-docs:resources:workload:properties:service:containers", "xcsh-docs:resources:workload:properties:service:deploy_options", "xcsh-docs:resources:workload:properties:service:scale_to_zero", "xcsh-docs:resources:workload:properties:service:volumes"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service", "parent_id": "xcsh-docs:resources:workload:reference", "path": "documentation/resources/workload/properties/service/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011", "registry_path": "docs/guides/resources--workload--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service"], "schema_version": 1, "sections": [{"aliases": ["service advertise options"], "anchor": "section", "description": "Advertise OPTIONS are used to configure how and where to advertise the workload using load balancers.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options"], "syntax": "block", "type": "object"}, {"aliases": ["service configuration"], "anchor": "section", "description": "Configuration parameters of the workload.", "document_id": "xcsh-docs:resources:workload:properties:service:configuration", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "configuration"], "syntax": "block", "type": "object"}, {"aliases": ["service containers"], "anchor": "section", "description": "Containers to use for service.", "document_id": "xcsh-docs:resources:workload:properties:service:containers", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["service", "containers"], "syntax": "block", "type": "object"}, {"aliases": ["service deploy options"], "anchor": "section", "description": "Deploy OPTIONS are used to configure the workload deployment OPTIONS.", "document_id": "xcsh-docs:resources:workload:properties:service:deploy_options", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "deploy_options"], "syntax": "block", "type": "object"}, {"aliases": ["service num replicas"], "anchor": "schema-service--num_replicas", "description": "Exclusive with Number of replicas of service to spawn per site.", "document_id": "xcsh-docs:resources:workload:properties:service", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "num_replicas"], "syntax": "attribute", "type": "number"}, {"aliases": ["service scale to zero"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:workload:properties:service:scale_to_zero", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "scale_to_zero"], "syntax": "attribute", "type": "object"}, {"aliases": ["service volumes"], "anchor": "section", "description": "Volumes for the service.", "document_id": "xcsh-docs:resources:workload:properties:service:volumes", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["service", "volumes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Service does not maintain per replica state, however it can be configured to use persistent storage that is shared amongst all the replicas. Replicas of a service are fungible and do not have a stable network identity or storage. Common examples of services are web servers, application servers, traditional SQL", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["workloadCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- service

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
service {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/): complete subsection reference.

- [configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/configuration/): complete subsection reference.

- [containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/): complete subsection reference.

- [deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/deploy_options/): complete subsection reference.

<a id="schema-service--num_replicas"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [scale_to_zero](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/scale_to_zero/): complete subsection reference.

- [volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/): complete subsection reference.
