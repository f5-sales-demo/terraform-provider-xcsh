---
page_title: "job.deploy_options.deploy_re_sites"
subcategory: "Container"
description: "This defines a way to deploy a workload on specific Regional Edge sites."
xcsh_docs: {"aliases": ["job deploy options deploy re sites"], "body_bytes": 1645, "body_sha256": "sha256:e804ce639cb01c1eb0908cbb5555e4ca33aafc0af5e4d911737cbda12d48ac19", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_re_sites:site"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_re_sites", "parent_id": "xcsh-docs:data-sources:workload:properties:job:deploy_options", "path": "documentation/data-sources/workload/properties/job/deploy_options/deploy_re_sites/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1202320102200012-1131132202020311-1033032131123113-1220132013322033-1122003012002200-0100312333012111-1211201333200120-1301030321103033", "registry_path": "docs/guides/data-sources--workload--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "deploy_options", "deploy_re_sites"], "schema_version": 1, "sections": [{"aliases": ["site"], "anchor": "section", "description": "Which regional edge sites should this workload be deployed.", "document_id": "xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_re_sites:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["job", "deploy_options", "deploy_re_sites", "site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/deploy_options/deploy_re_sites/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines a way to deploy a workload on specific Regional Edge sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.deploy_options.deploy_re_sites

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/)
- [job.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/deploy_options/)
- job.deploy_options.deploy_re_sites

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

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/deploy_options/deploy_re_sites/site/): complete subsection reference.

## Next pages

- [job.deploy_options.deploy_re_sites.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/deploy_options/deploy_re_sites/site/)
- [job.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/deploy_options/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
