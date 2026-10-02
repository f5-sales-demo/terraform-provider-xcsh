---
page_title: "job.deploy_options.deploy_ce_virtual_sites"
subcategory: "Container"
description: "This defines a way to deploy a workload on specific Customer virtual sites."
xcsh_docs: {"aliases": ["job deploy options deploy ce virtual sites"], "body_bytes": 1723, "body_sha256": "sha256:f88f901a14ad0a54bb68fd1d237c6dd184c8af2a9736628afbd8c69edfcb1de6", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_ce_virtual_sites:virtual_site"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_ce_virtual_sites", "parent_id": "xcsh-docs:data-sources:workload:properties:job:deploy_options", "path": "documentation/data-sources/workload/properties/job/deploy_options/deploy_ce_virtual_sites/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1211013222123021-1112130112123202-1213020033031330-3330223311120002-1000010231103301-1221032002020001-0100301311003321-0213202223331213", "registry_path": "docs/guides/data-sources--workload--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "deploy_options", "deploy_ce_virtual_sites"], "schema_version": 1, "sections": [{"aliases": ["virtual site"], "anchor": "section", "description": "Which customer virtual sites should this workload be deployed.", "document_id": "xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_ce_virtual_sites:virtual_site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["job", "deploy_options", "deploy_ce_virtual_sites", "virtual_site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/deploy_options/deploy_ce_virtual_sites/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines a way to deploy a workload on specific Customer virtual sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.deploy_options.deploy_ce_virtual_sites

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/)
- [job.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/deploy_options/)
- job.deploy_options.deploy_ce_virtual_sites

<a id="section"></a>

Type: `"single"`. Computed.

Defines a way to deploy a workload on specific Customer virtual sites.

Upstream description:

This defines a way to deploy a workload on specific Customer virtual sites.

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

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/deploy_options/deploy_ce_virtual_sites/virtual_site/): complete subsection reference.

## Next pages

- [job.deploy_options.deploy_ce_virtual_sites.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/deploy_options/deploy_ce_virtual_sites/virtual_site/)
- [job.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/deploy_options/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
