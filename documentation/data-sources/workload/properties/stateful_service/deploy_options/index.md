---
page_title: "stateful_service.deploy_options"
subcategory: "Container"
description: "Deploy OPTIONS are used to configure the workload deployment OPTIONS."
xcsh_docs: {"aliases": ["stateful service deploy options"], "body_bytes": 3615, "body_sha256": "sha256:d790e908ed9af46709c9e94397306320ce11684aa909391e92168d420bfae4c7", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options:all_res", "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options:default_virtual_sites", "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options:deploy_ce_sites", "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options:deploy_ce_virtual_sites", "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options:deploy_re_sites", "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options:deploy_re_virtual_sites"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service", "path": "documentation/data-sources/workload/properties/stateful_service/deploy_options/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1013001210313000-3001012211030201-1210322313302221-3132213121000320-0012111130102110-3332131123011111-1230211023121203-1131013202130133", "registry_path": "docs/guides/data-sources--workload--reference--group-028.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "deploy_options"], "schema_version": 1, "sections": [{"aliases": ["all res"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options:all_res", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "deploy_options", "all_res"], "syntax": "attribute", "type": "object"}, {"aliases": ["default virtual sites"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options:default_virtual_sites", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "deploy_options", "default_virtual_sites"], "syntax": "attribute", "type": "object"}, {"aliases": ["deploy ce sites"], "anchor": "section", "description": "This defines a way to deploy a workload on specific Customer sites.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options:deploy_ce_sites", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "deploy_options", "deploy_ce_sites"], "syntax": "attribute", "type": "object"}, {"aliases": ["deploy ce virtual sites"], "anchor": "section", "description": "This defines a way to deploy a workload on specific Customer virtual sites.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options:deploy_ce_virtual_sites", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "deploy_options", "deploy_ce_virtual_sites"], "syntax": "attribute", "type": "object"}, {"aliases": ["deploy re sites"], "anchor": "section", "description": "This defines a way to deploy a workload on specific Regional Edge sites.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options:deploy_re_sites", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "deploy_options", "deploy_re_sites"], "syntax": "attribute", "type": "object"}, {"aliases": ["deploy re virtual sites"], "anchor": "section", "description": "This defines a way to deploy a workload on specific Regional Edge virtual sites.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options:deploy_re_virtual_sites", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "deploy_options", "deploy_re_virtual_sites"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/deploy_options/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Deploy OPTIONS are used to configure the workload deployment OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.deploy_options

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- stateful_service.deploy_options

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

- [all_res](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/all_res/): complete subsection reference.

- [default_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/default_virtual_sites/): complete subsection reference.

- [deploy_ce_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/deploy_ce_sites/): complete subsection reference.

- [deploy_ce_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/deploy_ce_virtual_sites/): complete subsection reference.

- [deploy_re_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/deploy_re_sites/): complete subsection reference.

- [deploy_re_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/deploy_re_virtual_sites/): complete subsection reference.

## Next pages

- [stateful_service.deploy_options.all_res](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/all_res/)
- [stateful_service.deploy_options.default_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/default_virtual_sites/)
- [stateful_service.deploy_options.deploy_ce_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/deploy_ce_sites/)
- [stateful_service.deploy_options.deploy_ce_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/deploy_ce_virtual_sites/)
- [stateful_service.deploy_options.deploy_re_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/deploy_re_sites/)
- [stateful_service.deploy_options.deploy_re_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/deploy_re_virtual_sites/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
