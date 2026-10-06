---
page_title: "waf_exclusion"
subcategory: "Load Balancing"
description: "Configuration parameter for waf exclusion."
xcsh_docs: {"aliases": ["waf exclusion"], "body_bytes": 1481, "body_sha256": "sha256:bc8aba45dbdf11b0c84e16d520fec799594c0d9d8f510a2d1c7057a99b583018", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules", "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_policy"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/waf_exclusion/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-015.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion:ConflictingObjectAttributes:waf_exclusion_inline_rules,waf_exclusion_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion:ConflictingObjectAttributes:waf_exclusion_inline_rules,waf_exclusion_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_policy", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_exclusion"], "schema_version": 1, "sections": [{"aliases": ["waf exclusion waf exclusion inline rules"], "anchor": "section", "description": "A list of WAF exclusion rules that will be applied inline.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules"], "syntax": "block", "type": "object"}, {"aliases": ["waf exclusion waf exclusion policy"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-waf_exclusion--waf_exclusion_policy--name", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_policy:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_policy", "type": "requires"}], "schema_path": ["waf_exclusion", "waf_exclusion_policy"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/waf_exclusion/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configuration parameter for waf exclusion.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- waf_exclusion

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for waf exclusion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("waf_exclusion_inline_rules",
    "waf_exclusion_policy")}
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
  "x-ves-oneof-field-waf_exclusion_choice": "[\"waf_exclusion_inline_rules\",\"waf_exclusion_policy\"]"
}
```

Terraform syntax:

```terraform
waf_exclusion {
  # Configure direct properties listed below.
}
```

## Direct properties

- [waf_exclusion_inline_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/): complete subsection reference.

- [waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_policy/): complete subsection reference.
