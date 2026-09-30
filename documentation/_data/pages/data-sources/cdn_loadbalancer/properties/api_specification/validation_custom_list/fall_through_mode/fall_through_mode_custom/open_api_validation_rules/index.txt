---
page_title: "api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules"
subcategory: "Load Balancing"
description: "api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 7991, "body_sha256": "sha256:d8fb062d93218cc1a250063d96135826eaf588956ebd9e4dc01ede9406b32a75", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:action_block", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:action_report", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:action_skip", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:api_endpoint", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:metadata"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom", "path": "documentation/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/index.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode", "fall_through_mode_custom", "open_api_validation_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/)
- [api_specification.validation_custom_list.fall_through_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="section"></a>

Type: `"list"`. Computed.

Custom Fall Through Rule List. Rule or policy definition

Upstream description:

Rule or policy definition

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [action_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/action_block/): complete subsection reference.

- [action_report](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/action_report/): complete subsection reference.

- [action_skip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/action_skip/): complete subsection reference.

- [api_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/api_endpoint/): complete subsection reference.

<a id="schema-api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_group"></a>

### api_group property

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="schema-api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--base_path"></a>

### base_path property

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/metadata/): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/action_block/)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/action_report/)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/action_skip/)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/api_endpoint/)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/metadata/)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
