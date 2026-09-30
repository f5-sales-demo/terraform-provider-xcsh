---
page_title: "rules.match.ip_prefixes.prefixes"
subcategory: ""
description: "rules.match.ip_prefixes.prefixes for xcsh_bgp_routing_policy."
xcsh_docs: {"aliases": [], "body_bytes": 4690, "body_sha256": "sha256:3fe65248893f28b2b519ed4435e35a55aa2ef251a5853a27fbd6459665ff5181", "child_ids": ["xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:equal_or_longer_than", "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:exact_match", "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:longer_than"], "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes", "parent_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes", "path": "documentation/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/index.md", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["rules", "match", "ip_prefixes", "prefixes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.match.ip_prefixes.prefixes for xcsh_bgp_routing_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
