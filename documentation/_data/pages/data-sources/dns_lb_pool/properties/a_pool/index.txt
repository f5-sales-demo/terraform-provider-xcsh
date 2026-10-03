---
page_title: "a_pool"
subcategory: ""
description: "Pool for A Record."
xcsh_docs: {"aliases": ["a pool"], "body_bytes": 3631, "body_sha256": "sha256:e22cfbc1a4c16975d285287e648a25e3a51a52ee4654fc6a542867eda118cec6", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_lb_pool:properties:a_pool:disable_health_check", "xcsh-docs:data-sources:dns_lb_pool:properties:a_pool:health_check", "xcsh-docs:data-sources:dns_lb_pool:properties:a_pool:members"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_pool:properties:a_pool", "parent_id": "xcsh-docs:data-sources:dns_lb_pool:reference", "path": "documentation/data-sources/dns_lb_pool/properties/a_pool/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3111123132023110-0033301321322121-0320203222211000-2010033200211202-3101101202310110-0310100033023000-1133121301131210-3310320300121300", "registry_path": "docs/guides/data-sources--dns_lb_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["a_pool"], "schema_version": 1, "sections": [{"aliases": ["a pool disable health check"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:a_pool:disable_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["a_pool", "disable_health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["a pool health check"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:a_pool:health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["a_pool", "health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["a pool max answers"], "anchor": "schema-a_pool--max_answers", "description": "Limit on number of Resource Records to be included in the response to query.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:a_pool", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["a_pool", "max_answers"], "syntax": "attribute", "type": "number"}, {"aliases": ["a pool members"], "anchor": "section", "description": "Configuration parameter for members", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:a_pool:members", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["a_pool", "members"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_pool/properties/a_pool/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Pool for A Record.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# a_pool

Breadcrumbs:

- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/)
- a_pool

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: a\_pool, aaaa\_pool, cname\_pool, mx\_pool, srv\_pool\] Pool for A Record.

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

OneOf alternatives in this subsection:

- [a_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/a_pool/#section)
- [aaaa_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/aaaa_pool/#section)
- [cname_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/cname_pool/#section)
- [mx_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/mx_pool/#section)
- [srv_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/srv_pool/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [disable_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/a_pool/disable_health_check/): complete subsection reference.

- [health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/a_pool/health_check/): complete subsection reference.

<a id="schema-a_pool--max_answers"></a>

### max_answers property

Type: `"number"`. Computed.

Limit on number of Resource Records to be included in the response to query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

- [members](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/a_pool/members/): complete subsection reference.

## Next pages

- [a_pool.disable_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/a_pool/disable_health_check/)
- [a_pool.health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/a_pool/health_check/)
- [a_pool.members](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/a_pool/members/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/)
- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/)
