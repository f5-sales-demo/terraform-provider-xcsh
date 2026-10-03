---
page_title: "api_specification.validation_custom_list"
subcategory: "Load Balancing"
description: "Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other API-endpoint not listed will act according to \"Fall Through Mode\"."
xcsh_docs: {"aliases": ["api specification validation custom list"], "body_bytes": 2801, "body_sha256": "sha256:1bc5eae41832c9d251426c513bd9d71a6d67c92e1f8cf5d00e18614ff82dbf01", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:settings"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification", "path": "documentation/data-sources/http_loadbalancer/properties/api_specification/validation_custom_list/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list"], "schema_version": 1, "sections": [{"aliases": ["api specification validation custom list fall through mode"], "anchor": "section", "description": "Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. Swagger) or doesn't have a specific rule in custom rules)", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation custom list open api validation rules"], "anchor": "section", "description": "Rule or policy definition", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation custom list settings"], "anchor": "section", "description": "OpenAPI specification validation settings relevant for \"API Inventory\" enforcement and for \"Custom list\" enforcement.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:settings", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "settings"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_specification/validation_custom_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other API-endpoint not listed will act according to \"Fall Through Mode\".", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/)
- api_specification.validation_custom_list

<a id="section"></a>

Type: `"single"`. Computed.

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to 'Fall Through Mode'.

Upstream description:

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to "Fall Through Mode".

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-oversized_body_choice": "[]"
}
```

## Direct properties

- [fall_through_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/): complete subsection reference.

- [open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/): complete subsection reference.

- [settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.fall_through_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/)
- [api_specification.validation_custom_list.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/)
- [api_specification.validation_custom_list.settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
