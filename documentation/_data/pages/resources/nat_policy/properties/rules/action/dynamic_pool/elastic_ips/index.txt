---
page_title: "rules.action.dynamic_pool.elastic_ips"
subcategory: ""
description: "List of references to Cloud Elastic IP Object."
xcsh_docs: {"aliases": ["rules action dynamic pool elastic ips"], "body_bytes": 1367, "body_sha256": "sha256:c2cd6c2399e9f6d8af303c721fbf6d7a8b1ed664e1b5545f987b8f30b7a9f43b", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:action:dynamic_pool:elastic_ips:refs"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic_pool:elastic_ips", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic_pool", "path": "documentation/resources/nat_policy/properties/rules/action/dynamic_pool/elastic_ips/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1321030321131233-1211133330201130-1030020022332202-3210001332110133-2301303300323100-0022033302313303-0023333301233000-1113332222223113", "registry_path": "docs/guides/resources--nat_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.action.dynamic_pool.elastic_ips:RequiredObjectAttributes:refs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic_pool:elastic_ips:refs", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "action", "dynamic_pool", "elastic_ips"], "schema_version": 1, "sections": [{"aliases": ["rules action dynamic pool elastic ips refs"], "anchor": "section", "description": "Reference to one or more cloud elastic IP objects.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic_pool:elastic_ips:refs", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules", "action", "dynamic_pool", "elastic_ips", "refs"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/action/dynamic_pool/elastic_ips/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "List of references to Cloud Elastic IP Object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["nat_policyCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.action.dynamic_pool.elastic_ips

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/)
- [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/)
- [rules.action.dynamic_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic_pool/)
- rules.action.dynamic_pool.elastic_ips

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of references to Cloud Elastic IP Object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("refs")}
```

Terraform syntax:

```terraform
elastic_ips {
  # Configure direct properties listed below.
}
```

## Direct properties

- [refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic_pool/elastic_ips/refs/): complete subsection reference.
