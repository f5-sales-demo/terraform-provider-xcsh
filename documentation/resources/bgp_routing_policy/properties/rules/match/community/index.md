---
page_title: "rules.match.community"
subcategory: ""
description: "List of BGP communities."
xcsh_docs: {"aliases": ["rules match community"], "body_bytes": 3157, "body_sha256": "sha256:fe5665713e7608021c485e96dad955758811f80714fbec689555ebaf0fe0a9bb", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:community", "parent_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match", "path": "documentation/resources/bgp_routing_policy/properties/rules/match/community/index.md", "product": "distributed-cloud", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1233011320220011-0332311213110301-3123132110302233-0200233200131200-1320320000122200-3203032322131103-1332203002302102-2022320002030130", "registry_path": "docs/guides/resources--bgp_routing_policy--reference--group-001.md", "relationships": [{"anchor": "schema-rules--match--community--community", "enforcement": "provider-schema", "group": "rules.match.community:RequiredObjectAttributes:community", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:community", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "match", "community"], "schema_version": 1, "sections": [{"aliases": ["community"], "anchor": "schema-rules--match--community--community", "description": "An unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits being value.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:community", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "match", "community", "community"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/properties/rules/match/community/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of BGP communities.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.match.community

Breadcrumbs:

- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/)
- [rules.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/)
- rules.match.community

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

BGP Community list. List of BGP communities.

Upstream description:

List of BGP communities.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("community")}
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
community {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rules--match--community--community"></a>

### community property

Type: `["list", "string"]`. Optional.

Unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits being
value.

Upstream description:

An unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits
being value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9]{1,5}:[0-9]{1,5}$",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9]{1,5}:[0-9]{1,5}$",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [rules.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/)
- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
