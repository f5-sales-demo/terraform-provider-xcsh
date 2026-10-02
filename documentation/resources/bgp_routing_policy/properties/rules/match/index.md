---
page_title: "rules.match"
subcategory: ""
description: "Predicates which have to match information in route for action to be applied."
xcsh_docs: {"aliases": ["rules match"], "body_bytes": 2982, "body_sha256": "sha256:d93df837f4df61591b2738740a5256478758ca6f2e948b5d12dc80ee3babfbf1", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp_routing_policy:properties:rules:match:community", "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match", "parent_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules", "path": "documentation/resources/bgp_routing_policy/properties/rules/match/index.md", "product": "distributed-cloud", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3322310213101321-2011223230312312-2012023230122033-0123100010221033-0021213003311303-2133000223233132-3222001330131011-1211313320231111", "registry_path": "docs/guides/resources--bgp_routing_policy--reference--group-001.md", "relationships": [{"anchor": "schema-rules--match--as_path", "enforcement": "provider-schema", "group": "rules.match:ConflictingObjectAttributes:as_path,community", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match", "type": "conflicts"}, {"anchor": "schema-rules--match--as_path", "enforcement": "provider-schema", "group": "rules.match:ConflictingObjectAttributes:as_path,ip_prefixes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match:ConflictingObjectAttributes:as_path,community", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:community", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match:ConflictingObjectAttributes:community,ip_prefixes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:community", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match:ConflictingObjectAttributes:as_path,ip_prefixes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match:ConflictingObjectAttributes:community,ip_prefixes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "match"], "schema_version": 1, "sections": [{"aliases": ["as path"], "anchor": "schema-rules--match--as_path", "description": "Exclusive with AS path can also be a regex, which will be matched against route information.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "match", "as_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["community"], "anchor": "section", "description": "List of BGP communities.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:community", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rules--match--community--community", "enforcement": "provider-schema", "group": "rules.match.community:RequiredObjectAttributes:community", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:community", "type": "requires"}], "schema_path": ["rules", "match", "community"], "syntax": "block", "type": "object"}, {"aliases": ["ip prefixes"], "anchor": "section", "description": "List of IP prefix and prefix length range match condition.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes:RequiredObjectAttributes:prefixes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes", "type": "requires"}], "schema_path": ["rules", "match", "ip_prefixes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/properties/rules/match/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Predicates which have to match information in route for action to be applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.match

Breadcrumbs:

- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/)
- rules.match

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Predicates which have to match information in route for action to be applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("as_path",
    "community"),
  validators.ConflictingObjectAttributes("as_path",
    "ip_prefixes"),
  validators.ConflictingObjectAttributes("community",
    "ip_prefixes")}
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
  "x-ves-oneof-field-type_of_match": "[\"as_path\",\"community\",\"ip_prefixes\"]"
}
```

Terraform syntax:

```terraform
match {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rules--match--as_path"></a>

### as_path property

Type: `"string"`. Optional.

Exclusive with \[community ip\_prefixes\] AS path can also be a regex, which will be matched against
route information.

Upstream description:

Exclusive with \[community ip\_prefixes\] AS path can also be a regex, which will be matched against
route information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [community](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/community/): complete subsection reference.

- [ip_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/): complete subsection reference.

## Next pages

- [rules.match.community](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/community/)
- [rules.match.ip_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/)
- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
