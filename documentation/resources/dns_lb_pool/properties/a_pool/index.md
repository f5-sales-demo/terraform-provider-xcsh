---
page_title: "a_pool"
subcategory: ""
description: "Pool for A Record."
xcsh_docs: {"aliases": ["a pool"], "body_bytes": 3503, "body_sha256": "sha256:bc19c4769144febd90d3e9a16da2785f2e7ece6e3afd92662946aea0a8c9e226", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_lb_pool:properties:a_pool:disable_health_check", "xcsh-docs:resources:dns_lb_pool:properties:a_pool:health_check", "xcsh-docs:resources:dns_lb_pool:properties:a_pool:members"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_pool:properties:a_pool", "parent_id": "xcsh-docs:resources:dns_lb_pool:reference", "path": "documentation/resources/dns_lb_pool/properties/a_pool/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0123311331023030-2123210223210232-2020121111022213-0033300301133130-3013123123122303-2220212232330101-0132302013203032-0000203310030023", "registry_path": "docs/guides/resources--dns_lb_pool--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "a_pool:ConflictingObjectAttributes:disable_health_check,health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:a_pool:disable_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "a_pool:ConflictingObjectAttributes:disable_health_check,health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:a_pool:health_check", "type": "conflicts"}, {"anchor": "schema-a_pool--max_answers", "enforcement": "provider-schema", "group": "a_pool:RequiredObjectAttributes:max_answers,members", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:a_pool", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "a_pool:RequiredObjectAttributes:max_answers,members", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:a_pool:members", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["a_pool"], "schema_version": 1, "sections": [{"aliases": ["a pool disable health check", "disable dns health check", "no health check a pool"], "anchor": "section", "description": "Disables health checking for the DNS load balancer A-record pool.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:a_pool:disable_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["a_pool", "disable_health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["a pool health check"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:a_pool:health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-a_pool--health_check--name", "enforcement": "provider-schema", "group": "a_pool.health_check:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:a_pool:health_check", "type": "requires"}], "schema_path": ["a_pool", "health_check"], "syntax": "block", "type": "object"}, {"aliases": ["a pool max answers"], "anchor": "schema-a_pool--max_answers", "description": "Limit on number of Resource Records to be included in the response to query.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:a_pool", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["a_pool", "max_answers"], "syntax": "attribute", "type": "number"}, {"aliases": ["a pool members"], "anchor": "section", "description": "Configuration parameter for members", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:a_pool:members", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-a_pool--members--ip_endpoint", "enforcement": "provider-schema", "group": "a_pool.members:RequiredListObjectAttributes:ip_endpoint", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:a_pool:members", "type": "requires"}], "schema_path": ["a_pool", "members"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_pool/properties/a_pool/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Pool for A Record.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# a_pool

Breadcrumbs:

- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/)
- a_pool

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: a\_pool, aaaa\_pool, cname\_pool, mx\_pool, srv\_pool\] Pool for A Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("max_answers",
    "members"),
  validators.ConflictingObjectAttributes("disable_health_check",
    "health_check")}
```

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

- [a_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/#section)
- [aaaa_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/aaaa_pool/#section)
- [cname_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/#section)
- [mx_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/mx_pool/#section)
- [srv_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/srv_pool/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
a_pool {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/disable_health_check/): complete subsection reference.

- [health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/health_check/): complete subsection reference.

<a id="schema-a_pool--max_answers"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [members](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/members/): complete subsection reference.
