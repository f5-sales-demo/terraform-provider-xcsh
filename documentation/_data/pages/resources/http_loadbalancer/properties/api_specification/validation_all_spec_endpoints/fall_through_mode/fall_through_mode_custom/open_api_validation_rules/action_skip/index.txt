---
page_title: "api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["api specification validation all spec endpoints fall through mode fall through mode custom open api validation rules action skip"], "body_bytes": 2261, "body_sha256": "sha256:34fff9a4dfa0094a1983c04bb8a4a32b9356473ba8897fc2fb6403fbe879e660", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:action_skip", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom:open_api_validation_rules", "path": "documentation/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/action_skip/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0131212003302220-2200201323110330-0023221301033121-1320020311320112-3010232220333303-0230210002032310-2231001220023232-1302021021230033", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "fall_through_mode", "fall_through_mode_custom", "open_api_validation_rules", "action_skip"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/action_skip/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

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
action_skip = {}
```

This is an empty object or choice marker. It has no direct properties.
