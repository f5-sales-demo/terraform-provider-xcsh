---
page_title: "cname_pool"
subcategory: ""
description: "Pool for CNAME Record."
xcsh_docs: {"aliases": ["cname pool"], "body_bytes": 1914, "body_sha256": "sha256:c4e2e078d5f79a5684034f22c8a0bbe0a2c98bf760b39e624f232078ca4fb0bb", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:disable_health_check", "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:health_check", "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:members"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool", "parent_id": "xcsh-docs:data-sources:dns_lb_pool:reference", "path": "documentation/data-sources/dns_lb_pool/properties/cname_pool/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1221101101101203-1232231130303113-3233111001012331-0223200211330133-0000031100033023-1121001003332323-0000203310233120-0300332003111211", "registry_path": "docs/guides/data-sources--dns_lb_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cname_pool"], "schema_version": 1, "sections": [{"aliases": ["disable health check"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:disable_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cname_pool", "disable_health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["health check"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cname_pool", "health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["members"], "anchor": "section", "description": "Configuration parameter for members", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:members", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cname_pool", "members"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_pool/properties/cname_pool/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Pool for CNAME Record.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [cname_pool.disable_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/cname_pool/disable_health_check/)
- [cname_pool.health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/cname_pool/health_check/)
- [cname_pool.members](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/cname_pool/members/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/)
- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/)
