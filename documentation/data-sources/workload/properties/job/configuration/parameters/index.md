---
page_title: "job.configuration.parameters"
subcategory: "Container"
description: "Parameters for the workload."
xcsh_docs: {"aliases": ["job configuration parameters"], "body_bytes": 2459, "body_sha256": "sha256:cbc3ec336dab8e0db304a4e6989db36a5ee1b39d197769d0383de4b861d6a1b3", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:configuration:parameters:env_var", "xcsh-docs:data-sources:workload:properties:job:configuration:parameters:file"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:configuration:parameters", "parent_id": "xcsh-docs:data-sources:workload:properties:job:configuration", "path": "documentation/data-sources/workload/properties/job/configuration/parameters/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0123121112013033-2300320122030031-0233030030122131-2133002033200123-0033021202333220-1001330303032322-2123012310000033-1133320321233312", "registry_path": "docs/guides/data-sources--workload--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "configuration", "parameters"], "schema_version": 1, "sections": [{"aliases": ["env var"], "anchor": "section", "description": "Environment Variable.", "document_id": "xcsh-docs:data-sources:workload:properties:job:configuration:parameters:env_var", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "configuration", "parameters", "env_var"], "syntax": "attribute", "type": "object"}, {"aliases": ["file"], "anchor": "section", "description": "Configuration File for the workload.", "document_id": "xcsh-docs:data-sources:workload:properties:job:configuration:parameters:file", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "configuration", "parameters", "file"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/configuration/parameters/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Parameters for the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.configuration.parameters

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/)
- [job.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/configuration/)
- job.configuration.parameters

<a id="section"></a>

Type: `"list"`. Computed.

Parameters. Parameters for the workload.

Upstream description:

Parameters for the workload.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [env_var](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/configuration/parameters/env_var/): complete subsection reference.

- [file](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/configuration/parameters/file/): complete subsection reference.

## Next pages

- [job.configuration.parameters.env_var](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/configuration/parameters/env_var/)
- [job.configuration.parameters.file](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/configuration/parameters/file/)
- [job.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/configuration/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
