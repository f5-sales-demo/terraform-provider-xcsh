---
page_title: "api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom"
subcategory: "Load Balancing"
description: "Define the fall through settings."
xcsh_docs: {"aliases": ["api specification validation all spec endpoints fall through mode fall through mode custom"], "body_bytes": 1871, "body_sha256": "sha256:eb676a119362f8f6caf973d196cc0c51988602c36edeacf14a152708a41a5da3", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom:open_api_validation_rules"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode", "path": "documentation/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "fall_through_mode", "fall_through_mode_custom"], "schema_version": 1, "sections": [{"aliases": ["api specification validation all spec endpoints fall through mode fall through mode custom open api validation rules"], "anchor": "section", "description": "Rule or policy definition", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom:open_api_validation_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "fall_through_mode", "fall_through_mode_custom", "open_api_validation_rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Define the fall through settings.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for fall through mode custom.

Additional upstream details:

Define the fall through settings.

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
fall_through_mode_custom {
  # Configure direct properties listed below.
}
```

## Direct properties

- [open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/): complete subsection reference.
