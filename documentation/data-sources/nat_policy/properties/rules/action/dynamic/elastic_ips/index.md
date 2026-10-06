---
page_title: "rules.action.dynamic.elastic_ips"
subcategory: ""
description: "List of references to Cloud Elastic IP Object."
xcsh_docs: {"aliases": ["rules action dynamic elastic ips"], "body_bytes": 1242, "body_sha256": "sha256:d9612ed4e9a4139a77c4a03f834809f6afef30b341e2a6a338c9a85c58c954dc", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic:elastic_ips:refs"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic:elastic_ips", "parent_id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic", "path": "documentation/data-sources/nat_policy/properties/rules/action/dynamic/elastic_ips/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0230003220303133-1230022313322321-2000012100313200-3120031001030001-2222023000333022-2331113020212101-0323310011033322-2013300011312320", "registry_path": "docs/guides/data-sources--nat_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "action", "dynamic", "elastic_ips"], "schema_version": 1, "sections": [{"aliases": ["rules action dynamic elastic ips refs"], "anchor": "section", "description": "Reference to one or more cloud elastic IP objects.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic:elastic_ips:refs", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules", "action", "dynamic", "elastic_ips", "refs"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/action/dynamic/elastic_ips/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of references to Cloud Elastic IP Object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.action.dynamic.elastic_ips

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/)
- [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/)
- [rules.action.dynamic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/dynamic/)
- rules.action.dynamic.elastic_ips

<a id="section"></a>

Type: `"single"`. Computed.

List of references to Cloud Elastic IP Object.

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

## Direct properties

- [refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/dynamic/elastic_ips/refs/): complete subsection reference.
