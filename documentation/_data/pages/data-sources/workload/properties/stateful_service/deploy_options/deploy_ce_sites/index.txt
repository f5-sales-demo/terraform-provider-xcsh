---
page_title: "stateful_service.deploy_options.deploy_ce_sites"
subcategory: "Container"
description: "This defines a way to deploy a workload on specific Customer sites."
xcsh_docs: {"aliases": ["stateful service deploy options deploy ce sites"], "body_bytes": 1778, "body_sha256": "sha256:562337da35a9f2724e9907d3269efb19cb6dc8a5194b84b1487e74c456a69649", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options:deploy_ce_sites:site"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options:deploy_ce_sites", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options", "path": "documentation/data-sources/workload/properties/stateful_service/deploy_options/deploy_ce_sites/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0331120123230233-2122000011230021-0222222332013221-3231012212320100-1230021300310130-0202332212333100-1010010101020110-0000323233332233", "registry_path": "docs/guides/data-sources--workload--reference--group-028.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "deploy_options", "deploy_ce_sites"], "schema_version": 1, "sections": [{"aliases": ["site"], "anchor": "section", "description": "Which customer sites should this workload be deployed.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:deploy_options:deploy_ce_sites:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["stateful_service", "deploy_options", "deploy_ce_sites", "site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/deploy_options/deploy_ce_sites/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines a way to deploy a workload on specific Customer sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.deploy_options.deploy_ce_sites

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- [stateful_service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/)
- stateful_service.deploy_options.deploy_ce_sites

<a id="section"></a>

Type: `"single"`. Computed.

Defines a way to deploy a workload on specific Customer sites.

Upstream description:

This defines a way to deploy a workload on specific Customer sites.

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

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/deploy_ce_sites/site/): complete subsection reference.

## Next pages

- [stateful_service.deploy_options.deploy_ce_sites.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/deploy_ce_sites/site/)
- [stateful_service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/deploy_options/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
