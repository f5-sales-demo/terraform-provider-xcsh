---
page_title: "mx_pool"
subcategory: ""
description: "Pool for MX Record."
xcsh_docs: {"aliases": ["mx pool"], "body_bytes": 2517, "body_sha256": "sha256:0f31b09ac711c5b6ba7a785db9c3874c29d5c3a4ca281859d0f53db74f3b54ac", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_lb_pool:properties:mx_pool:members"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_pool:properties:mx_pool", "parent_id": "xcsh-docs:resources:dns_lb_pool:reference", "path": "documentation/resources/dns_lb_pool/properties/mx_pool/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2212032032130133-2301313011222230-0123131011123021-0200203312302320-1232223301102320-1012121332111002-1123122031330210-2100311112312320", "registry_path": "docs/guides/resources--dns_lb_pool--reference--group-001.md", "relationships": [{"anchor": "schema-mx_pool--max_answers", "enforcement": "provider-schema", "group": "mx_pool:RequiredObjectAttributes:max_answers,members", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:mx_pool", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mx_pool:RequiredObjectAttributes:max_answers,members", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:mx_pool:members", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["mx_pool"], "schema_version": 1, "sections": [{"aliases": ["max answers"], "anchor": "schema-mx_pool--max_answers", "description": "Limit on number of Resource Records to be included in the response to query.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:mx_pool", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mx_pool", "max_answers"], "syntax": "attribute", "type": "number"}, {"aliases": ["members"], "anchor": "section", "description": "Configuration parameter for members", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:mx_pool:members", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-mx_pool--members--domain", "enforcement": "provider-schema", "group": "mx_pool.members:RequiredListObjectAttributes:domain", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:mx_pool:members", "type": "requires"}], "schema_path": ["mx_pool", "members"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_pool/properties/mx_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Pool for MX Record.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mx_pool

Breadcrumbs:

- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/)
- mx_pool

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Pool for MX Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("max_answers",
    "members")}
```

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
mx_pool {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-mx_pool--max_answers"></a>

### max_answers property

Type: `"number"`. Optional.

Limit on number of Resource Records to be included in the response to query.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 32),
}
```

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

- [members](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/mx_pool/members/): complete subsection reference.

## Next pages

- [mx_pool.members](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/mx_pool/members/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/)
- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
