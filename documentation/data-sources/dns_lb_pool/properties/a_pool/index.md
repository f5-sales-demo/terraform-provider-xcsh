---
page_title: "a_pool"
subcategory: ""
description: "Pool for A Record."
xcsh_docs: {"aliases": ["a pool"], "body_bytes": 3631, "body_sha256": "sha256:e22cfbc1a4c16975d285287e648a25e3a51a52ee4654fc6a542867eda118cec6", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_lb_pool:properties:a_pool:disable_health_check", "xcsh-docs:data-sources:dns_lb_pool:properties:a_pool:health_check", "xcsh-docs:data-sources:dns_lb_pool:properties:a_pool:members"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_pool:properties:a_pool", "parent_id": "xcsh-docs:data-sources:dns_lb_pool:reference", "path": "documentation/data-sources/dns_lb_pool/properties/a_pool/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3111123132023110-0033301321322121-0320203222211000-2010033200211202-3101101202310110-0310100033023000-1133121301131210-3310320300121300", "registry_path": "docs/guides/data-sources--dns_lb_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["a_pool"], "schema_version": 1, "sections": [{"aliases": ["disable health check"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:a_pool:disable_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["a_pool", "disable_health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["health check"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:a_pool:health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["a_pool", "health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["max answers"], "anchor": "schema-a_pool--max_answers", "description": "Limit on number of Resource Records to be included in the response to query.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:a_pool", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["a_pool", "max_answers"], "syntax": "attribute", "type": "number"}, {"aliases": ["members"], "anchor": "section", "description": "Configuration parameter for members", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:a_pool:members", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["a_pool", "members"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_pool/properties/a_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Pool for A Record.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
