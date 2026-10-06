---
page_title: "rules.match.ip_prefixes.prefixes"
subcategory: ""
description: "List of IP prefix."
xcsh_docs: {"aliases": ["rules match ip prefixes prefixes"], "body_bytes": 3874, "body_sha256": "sha256:c450ee66a8ab830504995e65ad76bf6cb5826ba88a099a95f488ca2d931b5272", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:equal_or_longer_than", "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:exact_match", "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:longer_than"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes", "parent_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes", "path": "documentation/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/index.md", "product": "distributed-cloud", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0132120330013330-1212032122110301-2220310332132012-3013212311221331-2231233031001113-0103320013313102-0131302100230132-3301221331021023", "registry_path": "docs/guides/resources--bgp_routing_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes.prefixes:ConflictingListObjectAttributes:equal_or_longer_than,exact_match", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:equal_or_longer_than", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes.prefixes:ConflictingListObjectAttributes:equal_or_longer_than,longer_than", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:equal_or_longer_than", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes.prefixes:ConflictingListObjectAttributes:equal_or_longer_than,exact_match", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:exact_match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes.prefixes:ConflictingListObjectAttributes:exact_match,longer_than", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:exact_match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes.prefixes:ConflictingListObjectAttributes:equal_or_longer_than,longer_than", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:longer_than", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes.prefixes:ConflictingListObjectAttributes:exact_match,longer_than", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:longer_than", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "match", "ip_prefixes", "prefixes"], "schema_version": 1, "sections": [{"aliases": ["rules match ip prefixes prefixes equal or longer than"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:equal_or_longer_than", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "match", "ip_prefixes", "prefixes", "equal_or_longer_than"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules match ip prefixes prefixes exact match"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:exact_match", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "match", "ip_prefixes", "prefixes", "exact_match"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules match ip prefixes prefixes ip prefixes"], "anchor": "schema-rules--match--ip_prefixes--prefixes--ip_prefixes", "description": "IP prefix to match on BGP route.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "match", "ip_prefixes", "prefixes", "ip_prefixes"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules match ip prefixes prefixes longer than"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:longer_than", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "match", "ip_prefixes", "prefixes", "longer_than"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of IP prefix.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
