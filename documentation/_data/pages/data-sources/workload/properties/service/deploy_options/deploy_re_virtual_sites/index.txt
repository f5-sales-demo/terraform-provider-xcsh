---
page_title: "service.deploy_options.deploy_re_virtual_sites"
subcategory: "Container"
description: "This defines a way to deploy a workload on specific Regional Edge virtual sites."
xcsh_docs: {"aliases": ["service deploy options deploy re virtual sites"], "body_bytes": 1207, "body_sha256": "sha256:890430498e3d939565c442df4b59c656bc9d5df0d5414984df7881d7ec26cd7e", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:deploy_options:deploy_re_virtual_sites:virtual_site"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:deploy_options:deploy_re_virtual_sites", "parent_id": "xcsh-docs:data-sources:workload:properties:service:deploy_options", "path": "documentation/data-sources/workload/properties/service/deploy_options/deploy_re_virtual_sites/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1032333221231212-0012100102110032-2232330111130330-0002233211230220-3011231233013001-2331331003213001-3221223033110011-3211122303201310", "registry_path": "docs/guides/data-sources--workload--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "deploy_options", "deploy_re_virtual_sites"], "schema_version": 1, "sections": [{"aliases": ["service deploy options deploy re virtual sites virtual site"], "anchor": "section", "description": "Which regional edge virtual sites should this workload be deployed.", "document_id": "xcsh-docs:data-sources:workload:properties:service:deploy_options:deploy_re_virtual_sites:virtual_site", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["service", "deploy_options", "deploy_re_virtual_sites", "virtual_site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/deploy_options/deploy_re_virtual_sites/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This defines a way to deploy a workload on specific Regional Edge virtual sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["workloadCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.deploy_options.deploy_re_virtual_sites

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
- [service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/deploy_options/)
- service.deploy_options.deploy_re_virtual_sites

<a id="section"></a>

Type: `"single"`. Computed.

This defines a way to deploy a workload on specific Regional Edge virtual sites.

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

## Direct properties

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/deploy_options/deploy_re_virtual_sites/virtual_site/): complete subsection reference.
