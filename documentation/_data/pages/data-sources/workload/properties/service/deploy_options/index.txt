---
page_title: "service.deploy_options"
subcategory: "Container"
description: "Deploy OPTIONS are used to configure the workload deployment OPTIONS."
xcsh_docs: {"aliases": ["service deploy options"], "body_bytes": 2108, "body_sha256": "sha256:41744c3d1254c823505892c0be7101473ba79f97e8b81fa923ecb1f889939625", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:deploy_options:all_res", "xcsh-docs:data-sources:workload:properties:service:deploy_options:default_virtual_sites", "xcsh-docs:data-sources:workload:properties:service:deploy_options:deploy_ce_sites", "xcsh-docs:data-sources:workload:properties:service:deploy_options:deploy_ce_virtual_sites", "xcsh-docs:data-sources:workload:properties:service:deploy_options:deploy_re_sites", "xcsh-docs:data-sources:workload:properties:service:deploy_options:deploy_re_virtual_sites"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:deploy_options", "parent_id": "xcsh-docs:data-sources:workload:properties:service", "path": "documentation/data-sources/workload/properties/service/deploy_options/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0121032001011031-2321102332010202-3221321233030232-3111022330331330-3220203323221010-1002033122322321-2321230330230311-1011232330212320", "registry_path": "docs/guides/data-sources--workload--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "deploy_options"], "schema_version": 1, "sections": [{"aliases": ["service deploy options all res"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:service:deploy_options:all_res", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "deploy_options", "all_res"], "syntax": "attribute", "type": "object"}, {"aliases": ["service deploy options default virtual sites"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:service:deploy_options:default_virtual_sites", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "deploy_options", "default_virtual_sites"], "syntax": "attribute", "type": "object"}, {"aliases": ["service deploy options deploy ce sites"], "anchor": "section", "description": "This defines a way to deploy a workload on specific Customer sites.", "document_id": "xcsh-docs:data-sources:workload:properties:service:deploy_options:deploy_ce_sites", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "deploy_options", "deploy_ce_sites"], "syntax": "attribute", "type": "object"}, {"aliases": ["service deploy options deploy ce virtual sites"], "anchor": "section", "description": "This defines a way to deploy a workload on specific Customer virtual sites.", "document_id": "xcsh-docs:data-sources:workload:properties:service:deploy_options:deploy_ce_virtual_sites", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "deploy_options", "deploy_ce_virtual_sites"], "syntax": "attribute", "type": "object"}, {"aliases": ["service deploy options deploy re sites"], "anchor": "section", "description": "This defines a way to deploy a workload on specific Regional Edge sites.", "document_id": "xcsh-docs:data-sources:workload:properties:service:deploy_options:deploy_re_sites", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "deploy_options", "deploy_re_sites"], "syntax": "attribute", "type": "object"}, {"aliases": ["service deploy options deploy re virtual sites"], "anchor": "section", "description": "This defines a way to deploy a workload on specific Regional Edge virtual sites.", "document_id": "xcsh-docs:data-sources:workload:properties:service:deploy_options:deploy_re_virtual_sites", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "deploy_options", "deploy_re_virtual_sites"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/deploy_options/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Deploy OPTIONS are used to configure the workload deployment OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.deploy_options

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
- service.deploy_options

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

- [all_res](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/deploy_options/all_res/): complete subsection reference.

- [default_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/deploy_options/default_virtual_sites/): complete subsection reference.

- [deploy_ce_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/deploy_options/deploy_ce_sites/): complete subsection reference.

- [deploy_ce_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/deploy_options/deploy_ce_virtual_sites/): complete subsection reference.

- [deploy_re_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/deploy_options/deploy_re_sites/): complete subsection reference.

- [deploy_re_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/deploy_options/deploy_re_virtual_sites/): complete subsection reference.
