---
page_title: "job.deploy_options"
subcategory: "Container"
description: "Deploy OPTIONS are used to configure the workload deployment OPTIONS."
xcsh_docs: {"aliases": ["job deploy options"], "body_bytes": 2068, "body_sha256": "sha256:0fe278491138ecce25c740481ffcc984fc9a7edc82d5007b933afe044576b4cf", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:deploy_options:all_res", "xcsh-docs:data-sources:workload:properties:job:deploy_options:default_virtual_sites", "xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_ce_sites", "xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_ce_virtual_sites", "xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_re_sites", "xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_re_virtual_sites"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:deploy_options", "parent_id": "xcsh-docs:data-sources:workload:properties:job", "path": "documentation/data-sources/workload/properties/job/deploy_options/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1021013221100131-2310031003132222-1301022330312301-0200301110312230-2121132222333110-2302011012310013-1021020103223211-3032100212011100", "registry_path": "docs/guides/data-sources--workload--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "deploy_options"], "schema_version": 1, "sections": [{"aliases": ["job deploy options all res"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:job:deploy_options:all_res", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "deploy_options", "all_res"], "syntax": "attribute", "type": "object"}, {"aliases": ["job deploy options default virtual sites"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:job:deploy_options:default_virtual_sites", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "deploy_options", "default_virtual_sites"], "syntax": "attribute", "type": "object"}, {"aliases": ["job deploy options deploy ce sites"], "anchor": "section", "description": "This defines a way to deploy a workload on specific Customer sites.", "document_id": "xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_ce_sites", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "deploy_options", "deploy_ce_sites"], "syntax": "attribute", "type": "object"}, {"aliases": ["job deploy options deploy ce virtual sites"], "anchor": "section", "description": "This defines a way to deploy a workload on specific Customer virtual sites.", "document_id": "xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_ce_virtual_sites", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "deploy_options", "deploy_ce_virtual_sites"], "syntax": "attribute", "type": "object"}, {"aliases": ["job deploy options deploy re sites"], "anchor": "section", "description": "This defines a way to deploy a workload on specific Regional Edge sites.", "document_id": "xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_re_sites", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "deploy_options", "deploy_re_sites"], "syntax": "attribute", "type": "object"}, {"aliases": ["job deploy options deploy re virtual sites"], "anchor": "section", "description": "This defines a way to deploy a workload on specific Regional Edge virtual sites.", "document_id": "xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_re_virtual_sites", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "deploy_options", "deploy_re_virtual_sites"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/deploy_options/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Deploy OPTIONS are used to configure the workload deployment OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["workloadCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.deploy_options

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/)
- job.deploy_options

<a id="section"></a>

Type: `"single"`. Computed.

Deploy OPTIONS are used to configure the workload deployment OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-deploy_choice": "[\"all_res\",\"default_virtual_sites\",\"deploy_ce_sites\",\"deploy_ce_virtual_sites\",\"deploy_re_sites\",\"deploy_re_virtual_sites\"]"
}
```

## Direct properties

- [all_res](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/deploy_options/all_res/): complete subsection reference.

- [default_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/deploy_options/default_virtual_sites/): complete subsection reference.

- [deploy_ce_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/deploy_options/deploy_ce_sites/): complete subsection reference.

- [deploy_ce_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/deploy_options/deploy_ce_virtual_sites/): complete subsection reference.

- [deploy_re_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/deploy_options/deploy_re_sites/): complete subsection reference.

- [deploy_re_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/deploy_options/deploy_re_virtual_sites/): complete subsection reference.
