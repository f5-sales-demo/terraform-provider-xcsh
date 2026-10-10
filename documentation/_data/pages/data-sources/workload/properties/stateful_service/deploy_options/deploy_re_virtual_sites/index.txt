---
page_title: "stateful_service.deploy_options.deploy_re_virtual_sites"
subcategory: "Container"
description: "This defines a way to deploy a workload on specific Regional Edge virtual sites."
xcsh_docs: {"aliases": ["stateful service deploy options deploy re virtual sites"], "body_bytes": 1270, "body_sha256": "sha256:eee5cdd0131c278585c305268f5fd0a4853523be05f04878451f6f84719038fe", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options:deploy_re_virtual_sites:virtual_site"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options:deploy_re_virtual_sites", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options", "path": "documentation/data-sources/workload/properties/stateful_service/deploy_options/deploy_re_virtual_sites/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3121001123303100-3221011200111132-3002021330303313-1333120230130323-1031101200021213-1323132121332323-3112123112321231-2032233011011030", "registry_path": "docs/guides/data-sources--workload--reference--group-027.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "deploy_options", "deploy_re_virtual_sites"], "schema_version": 1, "sections": [{"aliases": ["stateful service deploy options deploy re virtual sites virtual site"], "anchor": "section", "description": "Which regional edge virtual sites should this workload be deployed.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options:deploy_re_virtual_sites:virtual_site", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["stateful_service", "deploy_options", "deploy_re_virtual_sites", "virtual_site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/deploy_options/deploy_re_virtual_sites/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This defines a way to deploy a workload on specific Regional Edge virtual sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["workloadCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.deploy_options.deploy_re_virtual_sites

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- [stateful_service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/)
- stateful_service.deploy_options.deploy_re_virtual_sites

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

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/deploy_re_virtual_sites/virtual_site/): complete subsection reference.
