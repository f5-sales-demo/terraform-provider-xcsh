---
page_title: "api_protection_rules.api_groups_rules.action"
subcategory: "Load Balancing"
description: "api_protection_rules.api_groups_rules.action for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1744, "body_sha256": "sha256:06fd81712b94c38d626186f84b4b8d800d99027a8260880412235c439f162cf5", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action:allow", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action:deny"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--action.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_protection_rules", "api_groups_rules", "action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_protection_rules.api_groups_rules.action for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_groups_rules.action

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_protection_rules](data-sources--http_loadbalancer--properties--api_protection_rules.md)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules.md)
- api_protection_rules.api_groups_rules.action

<a id="section"></a>

Type: `"single"`. Computed.

The action to take if the input request matches the rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action": "[\"allow\",\"deny\"]"
}
```

## Direct properties

- [allow](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--action--allow.md): complete subsection reference.

- [deny](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--action--deny.md): complete subsection reference.

## Next pages

- [api_protection_rules.api_groups_rules.action.allow](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--action--allow.md)
- [api_protection_rules.api_groups_rules.action.deny](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--action--deny.md)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
