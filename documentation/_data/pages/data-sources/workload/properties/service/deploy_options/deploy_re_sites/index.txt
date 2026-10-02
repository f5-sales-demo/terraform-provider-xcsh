---
page_title: "service.deploy_options.deploy_re_sites"
subcategory: "Container"
description: "This defines a way to deploy a workload on specific Regional Edge sites."
xcsh_docs: {"aliases": ["service deploy options deploy re sites"], "body_bytes": 1689, "body_sha256": "sha256:596f29f17860e93bf5754f8df1b1c8b8a14dbf4737d0bd29445435cea5f53c39", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:deploy_options:deploy_re_sites:site"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:deploy_options:deploy_re_sites", "parent_id": "xcsh-docs:data-sources:workload:properties:service:deploy_options", "path": "documentation/data-sources/workload/properties/service/deploy_options/deploy_re_sites/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1303232022011320-1200101023133223-1211210313033103-2030201303303322-0231212030222200-3113013231131010-0012330122032110-2201123321132232", "registry_path": "docs/guides/data-sources--workload--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "deploy_options", "deploy_re_sites"], "schema_version": 1, "sections": [{"aliases": ["site"], "anchor": "section", "description": "Which regional edge sites should this workload be deployed.", "document_id": "xcsh-docs:data-sources:workload:properties:service:deploy_options:deploy_re_sites:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["service", "deploy_options", "deploy_re_sites", "site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/deploy_options/deploy_re_sites/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines a way to deploy a workload on specific Regional Edge sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.deploy_options.deploy_re_sites

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
- [service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/deploy_options/)
- service.deploy_options.deploy_re_sites

<a id="section"></a>

Type: `"single"`. Computed.

Defines a way to deploy a workload on specific Regional Edge sites.

Upstream description:

This defines a way to deploy a workload on specific Regional Edge sites.

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

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/deploy_options/deploy_re_sites/site/): complete subsection reference.

## Next pages

- [service.deploy_options.deploy_re_sites.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/deploy_options/deploy_re_sites/site/)
- [service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/deploy_options/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
