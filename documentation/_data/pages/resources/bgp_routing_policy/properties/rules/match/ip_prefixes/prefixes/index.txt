---
page_title: "rules.match.ip_prefixes.prefixes"
subcategory: ""
description: "List of IP prefix."
xcsh_docs: {"aliases": ["rules match ip prefixes prefixes"], "body_bytes": 4789, "body_sha256": "sha256:fac83819f7fe5c78e474b484fa8e54b802d8da0db8e28eee2090bc695132e9e1", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:equal_or_longer_than", "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:exact_match", "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:longer_than"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes", "parent_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes", "path": "documentation/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/index.md", "product": "distributed-cloud", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0132120330013330-1212032122110301-2220310332132012-3013212311221331-2231233031001113-0103320013313102-0131302100230132-3301221331021023", "registry_path": "docs/guides/resources--bgp_routing_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes.prefixes:ConflictingListObjectAttributes:equal_or_longer_than,exact_match", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:equal_or_longer_than", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes.prefixes:ConflictingListObjectAttributes:equal_or_longer_than,longer_than", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:equal_or_longer_than", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes.prefixes:ConflictingListObjectAttributes:equal_or_longer_than,exact_match", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:exact_match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes.prefixes:ConflictingListObjectAttributes:exact_match,longer_than", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:exact_match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes.prefixes:ConflictingListObjectAttributes:equal_or_longer_than,longer_than", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:longer_than", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes.prefixes:ConflictingListObjectAttributes:exact_match,longer_than", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:longer_than", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "match", "ip_prefixes", "prefixes"], "schema_version": 1, "sections": [{"aliases": ["equal or longer than"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:equal_or_longer_than", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "match", "ip_prefixes", "prefixes", "equal_or_longer_than"], "syntax": "attribute", "type": "object"}, {"aliases": ["exact match"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:exact_match", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "match", "ip_prefixes", "prefixes", "exact_match"], "syntax": "attribute", "type": "object"}, {"aliases": ["ip prefixes"], "anchor": "schema-rules--match--ip_prefixes--prefixes--ip_prefixes", "description": "IP prefix to match on BGP route.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "match", "ip_prefixes", "prefixes", "ip_prefixes"], "syntax": "attribute", "type": "string"}, {"aliases": ["longer than"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:longer_than", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "match", "ip_prefixes", "prefixes", "longer_than"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of IP prefix.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.match.ip_prefixes.prefixes

Breadcrumbs:

- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/)
- [rules.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/)
- [rules.match.ip_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/)
- rules.match.ip_prefixes.prefixes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Prefix list. List of IP prefix.

Upstream description:

List of IP prefix.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("equal_or_longer_than",
    "exact_match"),
  validators.ConflictingListObjectAttributes("equal_or_longer_than",
    "longer_than"),
  validators.ConflictingListObjectAttributes("exact_match",
    "longer_than")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
prefixes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [equal_or_longer_than](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/equal_or_longer_than/): complete subsection reference.

- [exact_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/exact_match/): complete subsection reference.

<a id="schema-rules--match--ip_prefixes--prefixes--ip_prefixes"></a>

### ip_prefixes property

Type: `"string"`. Optional.

IP Prefix. IP prefix to match on BGP route.

Upstream description:

IP prefix to match on BGP route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [longer_than](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/longer_than/): complete subsection reference.

## Next pages

- [rules.match.ip_prefixes.prefixes.equal_or_longer_than](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/equal_or_longer_than/)
- [rules.match.ip_prefixes.prefixes.exact_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/exact_match/)
- [rules.match.ip_prefixes.prefixes.longer_than](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/longer_than/)
- [rules.match.ip_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/)
- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
