---
page_title: "rules.match.ip_prefixes"
subcategory: ""
description: "List of IP prefix and prefix length range match condition."
xcsh_docs: {"aliases": ["rules match ip prefixes"], "body_bytes": 1819, "body_sha256": "sha256:f3b3f635517dfbc771ff9b32ea23d7fca3f36c7f0d6d8b02b17302f034511c93", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes", "parent_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match", "path": "documentation/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/index.md", "product": "distributed-cloud", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0220100211010101-1303213010003311-1312332131123303-0320011333112112-1010120101020113-1011320021022312-0032031022002221-1033332320233020", "registry_path": "docs/guides/resources--bgp_routing_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes:RequiredObjectAttributes:prefixes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "match", "ip_prefixes"], "schema_version": 1, "sections": [{"aliases": ["prefixes"], "anchor": "section", "description": "List of IP prefix.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes.prefixes:ConflictingListObjectAttributes:equal_or_longer_than,exact_match", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:equal_or_longer_than", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes.prefixes:ConflictingListObjectAttributes:equal_or_longer_than,longer_than", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:equal_or_longer_than", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes.prefixes:ConflictingListObjectAttributes:equal_or_longer_than,exact_match", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:exact_match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes.prefixes:ConflictingListObjectAttributes:exact_match,longer_than", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:exact_match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes.prefixes:ConflictingListObjectAttributes:equal_or_longer_than,longer_than", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:longer_than", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match.ip_prefixes.prefixes:ConflictingListObjectAttributes:exact_match,longer_than", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:longer_than", "type": "conflicts"}], "schema_path": ["rules", "match", "ip_prefixes", "prefixes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of IP prefix and prefix length range match condition.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.match.ip_prefixes

Breadcrumbs:

- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/)
- [rules.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/)
- rules.match.ip_prefixes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of IP prefix and prefix length range match condition.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefixes")}
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
ip_prefixes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/): complete subsection reference.

## Next pages

- [rules.match.ip_prefixes.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/)
- [rules.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/)
- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
