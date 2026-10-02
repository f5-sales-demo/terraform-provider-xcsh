---
page_title: "job.configuration"
subcategory: "Container"
description: "Configuration parameters of the workload."
xcsh_docs: {"aliases": ["job configuration"], "body_bytes": 1298, "body_sha256": "sha256:0318c8117cafce28bd2ed03a28161a73a3df01a956aca72f6185f230b8b19983", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:configuration:parameters"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:configuration", "parent_id": "xcsh-docs:data-sources:workload:properties:job", "path": "documentation/data-sources/workload/properties/job/configuration/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1013210033232002-3213323020033222-3100130012131022-1031113232032001-0113002330332013-1133031102002200-1003222022321133-2110222331232211", "registry_path": "docs/guides/data-sources--workload--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "configuration"], "schema_version": 1, "sections": [{"aliases": ["parameters"], "anchor": "section", "description": "Parameters for the workload.", "document_id": "xcsh-docs:data-sources:workload:properties:job:configuration:parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["job", "configuration", "parameters"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/configuration/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration parameters of the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.configuration

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/)
- job.configuration

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameters of the workload.

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

- [parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/configuration/parameters/): complete subsection reference.

## Next pages

- [job.configuration.parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/configuration/parameters/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
