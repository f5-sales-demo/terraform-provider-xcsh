---
page_title: "api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["api specification validation custom list fall through mode fall through mode custom open api validation rules action skip"], "body_bytes": 2191, "body_sha256": "sha256:58f5d4a67e77914226d4f289cc483582437572b77576ab7d84e96d5ebd00df97", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:action_skip", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules", "path": "documentation/resources/http_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/action_skip/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2131001201203002-2121121220233310-1123330102303132-1203220222231322-3201112231023213-0123312030001310-3133320323023220-2221000213020000", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode", "fall_through_mode_custom", "open_api_validation_rules", "action_skip"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/action_skip/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/)
- [api_specification.validation_custom_list.fall_through_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

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
