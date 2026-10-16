---
page_title: "advanced_options.enable_lb_source_ip_persistence"
subcategory: "Load Balancing"
description: "IP address configuration"
xcsh_docs: {"aliases": ["advanced options enable lb source ip persistence"], "body_bytes": 949, "body_sha256": "sha256:ff22c63bf17c7af6c5165b497bea12938c6f6277af6e78417acdaf9fa6a4e2bd", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_lb_source_ip_persistence", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options", "path": "documentation/data-sources/origin_pool/properties/advanced_options/enable_lb_source_ip_persistence/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-2031332023330111-0003133313331220-2300212112321120-0232132200120331-1022021123000030-0202010113022010-3333201013330233-0030200302312021", "registry_path": "docs/guides/data-sources--origin_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_options", "enable_lb_source_ip_persistence"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/advanced_options/enable_lb_source_ip_persistence/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "IP address configuration", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["origin_poolCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.enable_lb_source_ip_persistence

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
- [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/)
- advanced_options.enable_lb_source_ip_persistence

<a id="section"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

IP address configuration

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

This is an empty object or choice marker. It has no direct properties.
