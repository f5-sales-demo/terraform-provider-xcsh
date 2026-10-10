---
page_title: "cname_pool"
subcategory: ""
description: "Pool for CNAME Record."
xcsh_docs: {"aliases": ["cname pool"], "body_bytes": 1233, "body_sha256": "sha256:e74adec5b8ee6f975e92f2492089a3f8d7e57ecab5431301d40c3482ca62123b", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:disable_health_check", "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:health_check", "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:members"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool", "parent_id": "xcsh-docs:data-sources:dns_lb_pool:reference", "path": "documentation/data-sources/dns_lb_pool/properties/cname_pool/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1221101101101203-1232231130303113-3233111001012331-0223200211330133-0000031100033023-1121001003332323-0000203310233120-0300332003111211", "registry_path": "docs/guides/data-sources--dns_lb_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cname_pool"], "schema_version": 1, "sections": [{"aliases": ["cname pool disable health check"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:disable_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cname_pool", "disable_health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["cname pool health check"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cname_pool", "health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["cname pool members"], "anchor": "section", "description": "Configuration parameter for members", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:members", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cname_pool", "members"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_pool/properties/cname_pool/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Pool for CNAME Record.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cname_pool

Breadcrumbs:

- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/)
- cname_pool

<a id="section"></a>

Type: `"single"`. Computed.

Pool for CNAME Record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"disable_health_check\",\"health_check\"]"
}
```

## Direct properties

- [disable_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/cname_pool/disable_health_check/): complete subsection reference.

- [health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/cname_pool/health_check/): complete subsection reference.

- [members](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/cname_pool/members/): complete subsection reference.
