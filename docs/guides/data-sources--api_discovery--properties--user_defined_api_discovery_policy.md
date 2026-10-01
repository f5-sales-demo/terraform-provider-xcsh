---
page_title: "user_defined_api_discovery_policy"
subcategory: ""
description: "user_defined_api_discovery_policy for xcsh_api_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1813, "body_sha256": "sha256:4ccdf18d53cdea9ee1737a0cbf2eed03ff9a7aa15671786dbf694cd9808a1850", "canonical_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy", "child_ids": ["xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules", "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:exclusive", "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:inclusive"], "collection_id": "xcsh-docs:data-sources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy", "parent_id": "xcsh-docs:data-sources:api_discovery:reference", "path": "docs/guides/data-sources--api_discovery--properties--user_defined_api_discovery_policy.md", "provider_name": "api_discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["user_defined_api_discovery_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_discovery/properties/user_defined_api_discovery_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "user_defined_api_discovery_policy for xcsh_api_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_defined_api_discovery_policy

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md)
- [Property reference](data-sources--api_discovery--reference.md)
- user_defined_api_discovery_policy

<a id="section"></a>

Type: `"single"`. Computed.

Rules are evaluated sequentially, top to bottom. If no rules are added, all traffic will be
discovered or ignored based on the selection in the 'Default Behaviour of the Rule Set' field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_behavior_choice": "[\"exclusive\",\"inclusive\"]"
}
```

## Direct properties

- [discovery_rules](data-sources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules.md): complete subsection reference.

- [exclusive](data-sources--api_discovery--properties--user_defined_api_discovery_policy--exclusive.md): complete subsection reference.

- [inclusive](data-sources--api_discovery--properties--user_defined_api_discovery_policy--inclusive.md): complete subsection reference.

## Next pages

- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules.md)
- [user_defined_api_discovery_policy.exclusive](data-sources--api_discovery--properties--user_defined_api_discovery_policy--exclusive.md)
- [user_defined_api_discovery_policy.inclusive](data-sources--api_discovery--properties--user_defined_api_discovery_policy--inclusive.md)
- [Property reference](data-sources--api_discovery--reference.md)
- [xcsh_api_discovery](../data-sources/api_discovery.md)
