---
page_title: "service.deploy_options.deploy_re_virtual_sites"
subcategory: "Container"
description: "This defines a way to deploy a workload on specific Regional Edge virtual sites."
xcsh_docs: {"aliases": ["service deploy options deploy re virtual sites"], "body_bytes": 1777, "body_sha256": "sha256:da79df1c741ef1ba308b623c57ca92bf661b30ba8664dc58c97e8cb00a9cb6d1", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:deploy_options:deploy_re_virtual_sites:virtual_site"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:deploy_options:deploy_re_virtual_sites", "parent_id": "xcsh-docs:data-sources:workload:properties:service:deploy_options", "path": "documentation/data-sources/workload/properties/service/deploy_options/deploy_re_virtual_sites/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1032333221231212-0012100102110032-2232330111130330-0002233211230220-3011231233013001-2331331003213001-3221223033110011-3211122303201310", "registry_path": "docs/guides/data-sources--workload--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "deploy_options", "deploy_re_virtual_sites"], "schema_version": 1, "sections": [{"aliases": ["virtual site"], "anchor": "section", "description": "Which regional edge virtual sites should this workload be deployed.", "document_id": "xcsh-docs:data-sources:workload:properties:service:deploy_options:deploy_re_virtual_sites:virtual_site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["service", "deploy_options", "deploy_re_virtual_sites", "virtual_site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/deploy_options/deploy_re_virtual_sites/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines a way to deploy a workload on specific Regional Edge virtual sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Defines a way to deploy a workload on specific Regional Edge virtual sites.

Upstream description:

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

## Next pages

- [service.deploy_options.deploy_re_virtual_sites.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/deploy_options/deploy_re_virtual_sites/virtual_site/)
- [service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/deploy_options/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
