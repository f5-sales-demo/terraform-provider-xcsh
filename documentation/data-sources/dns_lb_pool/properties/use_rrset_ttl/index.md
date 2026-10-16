---
page_title: "use_rrset_ttl"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["use rrset ttl"], "body_bytes": 804, "body_sha256": "sha256:7febacf354967458cd588b8f73d57d41064b628b4f78dcf20c6119822cf7d46d", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_pool:properties:use_rrset_ttl", "parent_id": "xcsh-docs:data-sources:dns_lb_pool:reference", "path": "documentation/data-sources/dns_lb_pool/properties/use_rrset_ttl/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-3103203103023110-1213233013101233-0020122011003001-2311022001103201-0020331202131321-1010221101300002-1311010322110100-3230130000033321", "registry_path": "docs/guides/data-sources--dns_lb_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["use_rrset_ttl"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_pool/properties/use_rrset_ttl/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_rrset_ttl

Breadcrumbs:

- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/)
- use_rrset_ttl

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use rrset ttl.

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

This is an empty object or choice marker. It has no direct properties.
