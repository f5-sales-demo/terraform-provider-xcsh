---
page_title: "api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules"
subcategory: "Load Balancing"
description: "api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 7249, "body_sha256": "sha256:2cd6a3d9f519c8f603fabb33e77c39391f8df5bea38140c2b3a50dd49a4f3d54", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:action_block", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:action_report", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:action_skip", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:api_endpoint", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:metadata"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode", "fall_through_mode_custom", "open_api_validation_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_specification](data-sources--http_loadbalancer--properties--api_specification.md)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list.md)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode.md)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom.md)
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

- [action_block](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--action_block.md): complete subsection reference.

- [action_report](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--action_report.md): complete subsection reference.

- [action_skip](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--action_skip.md): complete subsection reference.

- [api_endpoint](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint.md): complete subsection reference.

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

- [metadata](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--metadata.md): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--action_block.md)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--action_report.md)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--action_skip.md)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint.md)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--metadata.md)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
