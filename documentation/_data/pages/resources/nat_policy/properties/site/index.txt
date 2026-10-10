---
page_title: "site"
subcategory: ""
description: "Reference to Site Object."
xcsh_docs: {"aliases": ["site"], "body_bytes": 1064, "body_sha256": "sha256:6bd827a50e40b858ab8fbf1c8d41832b68f234673d115a4bb18bab3ce79920ef", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:nat_policy:properties:site:refs"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:site", "parent_id": "xcsh-docs:resources:nat_policy:reference", "path": "documentation/resources/nat_policy/properties/site/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3001223222011220-0003203113332212-3332203303130003-1310223301012333-2223122022201100-2321110020002221-0120100011321320-2020123030023230", "registry_path": "docs/guides/resources--nat_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "site:RequiredObjectAttributes:refs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:site:refs", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["site"], "schema_version": 1, "sections": [{"aliases": ["site refs"], "anchor": "section", "description": "Reference to Site Object.", "document_id": "xcsh-docs:resources:nat_policy:properties:site:refs", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["site", "refs"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/site/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Reference to Site Object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["nat_policyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/)
- site

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Site Reference Type. Reference to Site Object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("refs")}
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
site {
  # Configure direct properties listed below.
}
```

## Direct properties

- [refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/refs/): complete subsection reference.
