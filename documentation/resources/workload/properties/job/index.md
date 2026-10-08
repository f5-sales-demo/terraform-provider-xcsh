---
page_title: "job"
subcategory: "Container"
description: "Jobs are used for running batch processing tasks and run to completion. Jobs are generally used for tasks like report generation, billing, parallel data processing, ETL processing, etc."
xcsh_docs: {"aliases": ["job"], "body_bytes": 3221, "body_sha256": "sha256:ad18e29a6fd31e36d139ac685d1444dd640b18c5ce4cbe95cabe86625dbbe2ad", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:job:configuration", "xcsh-docs:resources:workload:properties:job:containers", "xcsh-docs:resources:workload:properties:job:deploy_options", "xcsh-docs:resources:workload:properties:job:volumes"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job", "parent_id": "xcsh-docs:resources:workload:reference", "path": "documentation/resources/workload/properties/job/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311", "registry_path": "docs/guides/resources--workload--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "job:RequiredObjectAttributes:containers", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["job"], "schema_version": 1, "sections": [{"aliases": ["job configuration"], "anchor": "section", "description": "Configuration parameters of the workload.", "document_id": "xcsh-docs:resources:workload:properties:job:configuration", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "configuration"], "syntax": "block", "type": "object"}, {"aliases": ["job containers"], "anchor": "section", "description": "Containers to use for the job.", "document_id": "xcsh-docs:resources:workload:properties:job:containers", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-job--containers--flavor", "enforcement": "provider-schema", "group": "job.containers:ConflictingListObjectAttributes:custom_flavor,flavor", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers", "type": "conflicts"}, {"anchor": "schema-job--containers--flavor", "enforcement": "provider-schema", "group": "job.containers:ConflictingListObjectAttributes:default_flavor,flavor", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.containers:ConflictingListObjectAttributes:custom_flavor,default_flavor", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:custom_flavor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.containers:ConflictingListObjectAttributes:custom_flavor,flavor", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:custom_flavor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.containers:ConflictingListObjectAttributes:custom_flavor,default_flavor", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:default_flavor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.containers:ConflictingListObjectAttributes:default_flavor,flavor", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:default_flavor", "type": "conflicts"}, {"anchor": "schema-job--containers--name", "enforcement": "provider-schema", "group": "job.containers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers", "type": "requires"}], "schema_path": ["job", "containers"], "syntax": "block", "type": "object"}, {"aliases": ["job deploy options"], "anchor": "section", "description": "Deploy OPTIONS are used to configure the workload deployment OPTIONS.", "document_id": "xcsh-docs:resources:workload:properties:job:deploy_options", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:all_res,default_virtual_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:all_res", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:all_res,deploy_ce_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:all_res", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:all_res,deploy_ce_virtual_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:all_res", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:all_res,deploy_re_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:all_res", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:all_res,deploy_re_virtual_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:all_res", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:all_res,default_virtual_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:default_virtual_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:default_virtual_sites,deploy_ce_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:default_virtual_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:default_virtual_sites,deploy_ce_virtual_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:default_virtual_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:default_virtual_sites,deploy_re_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:default_virtual_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:default_virtual_sites,deploy_re_virtual_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:default_virtual_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:all_res,deploy_ce_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:default_virtual_sites,deploy_ce_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:deploy_ce_sites,deploy_ce_virtual_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:deploy_ce_sites,deploy_re_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:deploy_ce_sites,deploy_re_virtual_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:all_res,deploy_ce_virtual_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_virtual_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:default_virtual_sites,deploy_ce_virtual_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_virtual_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:deploy_ce_sites,deploy_ce_virtual_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_virtual_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:deploy_ce_virtual_sites,deploy_re_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_virtual_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:deploy_ce_virtual_sites,deploy_re_virtual_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_virtual_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:all_res,deploy_re_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:default_virtual_sites,deploy_re_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:deploy_ce_sites,deploy_re_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:deploy_ce_virtual_sites,deploy_re_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:deploy_re_sites,deploy_re_virtual_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:all_res,deploy_re_virtual_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_virtual_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:default_virtual_sites,deploy_re_virtual_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_virtual_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:deploy_ce_sites,deploy_re_virtual_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_virtual_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:deploy_ce_virtual_sites,deploy_re_virtual_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_virtual_sites", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options:ConflictingObjectAttributes:deploy_re_sites,deploy_re_virtual_sites", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_virtual_sites", "type": "conflicts"}], "schema_path": ["job", "deploy_options"], "syntax": "block", "type": "object"}, {"aliases": ["job num replicas"], "anchor": "schema-job--num_replicas", "description": "Number of replicas of the batch job to spawn per site.", "document_id": "xcsh-docs:resources:workload:properties:job", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "num_replicas"], "syntax": "attribute", "type": "number"}, {"aliases": ["job volumes"], "anchor": "section", "description": "Volumes for the job.", "document_id": "xcsh-docs:resources:workload:properties:job:volumes", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "job.volumes:ConflictingListObjectAttributes:empty_dir,host_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:volumes:empty_dir", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.volumes:ConflictingListObjectAttributes:empty_dir,persistent_volume", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:volumes:empty_dir", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.volumes:ConflictingListObjectAttributes:empty_dir,host_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:volumes:host_path", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.volumes:ConflictingListObjectAttributes:host_path,persistent_volume", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:volumes:host_path", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.volumes:ConflictingListObjectAttributes:empty_dir,persistent_volume", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.volumes:ConflictingListObjectAttributes:host_path,persistent_volume", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume", "type": "conflicts"}], "schema_path": ["job", "volumes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Jobs are used for running batch processing tasks and run to completion. Jobs are generally used for tasks like report generation, billing, parallel data processing, ETL processing, etc.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["workloadCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- job

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: job, service, simple\_service, stateful\_service\] Jobs are used for running batch
processing tasks and run to completion. Jobs are generally used for tasks like report generation,
billing, parallel data processing, ETL processing, etc.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/#section)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/#section)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/#section)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
job {
  # Configure direct properties listed below.
}
```

## Direct properties

- [configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/configuration/): complete subsection reference.

- [containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/): complete subsection reference.

- [deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/): complete subsection reference.

<a id="schema-job--num_replicas"></a>

### num_replicas property

Type: `"number"`. Optional.

Number of replicas of the batch job to spawn per site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/): complete subsection reference.
