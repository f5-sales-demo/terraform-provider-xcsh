---
page_title: "job"
subcategory: "Container"
description: "Jobs are used for running batch processing tasks and run to completion. Jobs are generally used for tasks like report generation, billing, parallel data processing, ETL processing, etc."
xcsh_docs: {"aliases": ["job"], "body_bytes": 3748, "body_sha256": "sha256:81c1bed320a9efe96511735abde5533a11ccbde416caf3f80b158e4627b1ac21", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:configuration", "xcsh-docs:data-sources:workload:properties:job:containers", "xcsh-docs:data-sources:workload:properties:job:deploy_options", "xcsh-docs:data-sources:workload:properties:job:volumes"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job", "parent_id": "xcsh-docs:data-sources:workload:reference", "path": "documentation/data-sources/workload/properties/job/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131", "registry_path": "docs/guides/data-sources--workload--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job"], "schema_version": 1, "sections": [{"aliases": ["job configuration"], "anchor": "section", "description": "Configuration parameters of the workload.", "document_id": "xcsh-docs:data-sources:workload:properties:job:configuration", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "configuration"], "syntax": "attribute", "type": "object"}, {"aliases": ["job containers"], "anchor": "section", "description": "Containers to use for the job.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["job", "containers"], "syntax": "attribute", "type": "object"}, {"aliases": ["job deploy options"], "anchor": "section", "description": "Deploy OPTIONS are used to configure the workload deployment OPTIONS.", "document_id": "xcsh-docs:data-sources:workload:properties:job:deploy_options", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "deploy_options"], "syntax": "attribute", "type": "object"}, {"aliases": ["job num replicas"], "anchor": "schema-job--num_replicas", "description": "Number of replicas of the batch job to spawn per site.", "document_id": "xcsh-docs:data-sources:workload:properties:job", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "num_replicas"], "syntax": "attribute", "type": "number"}, {"aliases": ["job volumes"], "anchor": "section", "description": "Volumes for the job.", "document_id": "xcsh-docs:data-sources:workload:properties:job:volumes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["job", "volumes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Jobs are used for running batch processing tasks and run to completion. Jobs are generally used for tasks like report generation, billing, parallel data processing, ETL processing, etc.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
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

- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/#section)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/#section)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/#section)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/configuration/): complete subsection reference.

- [containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/): complete subsection reference.

- [deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/deploy_options/): complete subsection reference.

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

- [volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/volumes/): complete subsection reference.

## Next pages

- [job.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/configuration/)
- [job.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/)
- [job.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/deploy_options/)
- [job.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/volumes/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
