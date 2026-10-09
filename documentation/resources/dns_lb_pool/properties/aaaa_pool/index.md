---
page_title: "aaaa_pool"
subcategory: ""
description: "Pool for AAAA Record."
xcsh_docs: {"aliases": ["aaaa pool"], "body_bytes": 2231, "body_sha256": "sha256:09884003971b714dca3436b6d07de9132ca188df945ed914915fa77406633534", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_lb_pool:properties:aaaa_pool:members"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_pool:properties:aaaa_pool", "parent_id": "xcsh-docs:resources:dns_lb_pool:reference", "path": "documentation/resources/dns_lb_pool/properties/aaaa_pool/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2212212022130300-3213012311313203-2210121222221333-2013311220021222-0022110313202023-1323111030012311-3031202030303213-0113222132323033", "registry_path": "docs/guides/resources--dns_lb_pool--reference--group-001.md", "relationships": [{"anchor": "schema-aaaa_pool--max_answers", "enforcement": "provider-schema", "group": "aaaa_pool:RequiredObjectAttributes:max_answers,members", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:aaaa_pool", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aaaa_pool:RequiredObjectAttributes:max_answers,members", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:aaaa_pool:members", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aaaa_pool"], "schema_version": 1, "sections": [{"aliases": ["aaaa pool max answers"], "anchor": "schema-aaaa_pool--max_answers", "description": "Limit on number of Resource Records to be included in the response to query.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:aaaa_pool", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aaaa_pool", "max_answers"], "syntax": "attribute", "type": "number"}, {"aliases": ["aaaa pool members"], "anchor": "section", "description": "Configuration parameter for members", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:aaaa_pool:members", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-aaaa_pool--members--ip_endpoint", "enforcement": "provider-schema", "group": "aaaa_pool.members:RequiredListObjectAttributes:ip_endpoint", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:aaaa_pool:members", "type": "requires"}], "schema_path": ["aaaa_pool", "members"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_pool/properties/aaaa_pool/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Pool for AAAA Record.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aaaa_pool

Breadcrumbs:

- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/)
- aaaa_pool

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Pool for AAAA Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
aaaa_pool {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-aaaa_pool--max_answers"></a>

### max_answers property

Type: `"number"`. Optional.

Limit on number of Resource Records to be included in the response to query.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [members](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/aaaa_pool/members/): complete subsection reference.
