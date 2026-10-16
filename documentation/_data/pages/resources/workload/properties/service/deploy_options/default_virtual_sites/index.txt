---
page_title: "service.deploy_options.default_virtual_sites"
subcategory: "Container"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["service deploy options default virtual sites"], "body_bytes": 1137, "body_sha256": "sha256:255629ab3efd0065fc038ed22858dd3a5fc1604e17de470ad69e6bc7d4fdd824", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:deploy_options:default_virtual_sites", "parent_id": "xcsh-docs:resources:workload:properties:service:deploy_options", "path": "documentation/resources/workload/properties/service/deploy_options/default_virtual_sites/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0123021220301230-0122112311122001-3213103300231202-3110232110313113-2030320231100102-1111033100101031-3212301103121201-1001132210103332", "registry_path": "docs/guides/resources--workload--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "deploy_options", "default_virtual_sites"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/deploy_options/default_virtual_sites/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["workloadCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.deploy_options.default_virtual_sites

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/deploy_options/)
- service.deploy_options.default_virtual_sites

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
default_virtual_sites = {}
```

This is an empty object or choice marker. It has no direct properties.
