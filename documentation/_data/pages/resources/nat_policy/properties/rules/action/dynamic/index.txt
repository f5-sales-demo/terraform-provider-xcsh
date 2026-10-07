---
page_title: "rules.action.dynamic"
subcategory: ""
description: "Dynamic Pool Configuration."
xcsh_docs: {"aliases": ["rules action dynamic"], "body_bytes": 1599, "body_sha256": "sha256:4b65918d8fdc3d830c01e20bd90bf8190757c86f618490b43b4dc10ab85c500d", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:action:dynamic:elastic_ips", "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic:pools"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules:action", "path": "documentation/resources/nat_policy/properties/rules/action/dynamic/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0221213332001013-1211231020320202-2110123331210011-3311233300323212-0323332201123333-3023103321300300-1013003221001113-1100103221120012", "registry_path": "docs/guides/resources--nat_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.action.dynamic:ConflictingObjectAttributes:elastic_ips,pools", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic:elastic_ips", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action.dynamic:ConflictingObjectAttributes:elastic_ips,pools", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic:pools", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "action", "dynamic"], "schema_version": 1, "sections": [{"aliases": ["rules action dynamic elastic ips"], "anchor": "section", "description": "List of references to Cloud Elastic IP Object.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic:elastic_ips", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.action.dynamic.elastic_ips:RequiredObjectAttributes:refs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic:elastic_ips:refs", "type": "requires"}], "schema_path": ["rules", "action", "dynamic", "elastic_ips"], "syntax": "block", "type": "object"}, {"aliases": ["rules action dynamic pools"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic:pools", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "action", "dynamic", "pools"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/action/dynamic/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Dynamic Pool Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["nat_policyCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.action.dynamic

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/)
- [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/)
- rules.action.dynamic

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Dynamic Pool. Dynamic Pool Configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("elastic_ips",
    "pools")}
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
  "x-ves-oneof-field-pool_choice": "[\"elastic_ips\",\"pools\"]"
}
```

Terraform syntax:

```terraform
dynamic {
  # Configure direct properties listed below.
}
```

## Direct properties

- [elastic_ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/): complete subsection reference.

- [pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/pools/): complete subsection reference.
