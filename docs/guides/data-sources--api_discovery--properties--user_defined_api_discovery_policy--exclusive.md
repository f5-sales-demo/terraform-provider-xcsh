---
page_title: "user_defined_api_discovery_policy.exclusive"
subcategory: ""
description: "user_defined_api_discovery_policy.exclusive for xcsh_api_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1566, "body_sha256": "sha256:771f830b565e77e30c680a39125d9477443376f2817e86a575ac86d16689c001", "canonical_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:exclusive", "child_ids": ["xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:archive", "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:ignore"], "collection_id": "xcsh-docs:data-sources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:exclusive", "parent_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy", "path": "docs/guides/data-sources--api_discovery--properties--user_defined_api_discovery_policy--exclusive.md", "provider_name": "api_discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["user_defined_api_discovery_policy", "exclusive"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "user_defined_api_discovery_policy.exclusive for xcsh_api_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# user_defined_api_discovery_policy.exclusive

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md)
- [Property reference](data-sources--api_discovery--reference.md)
- [user_defined_api_discovery_policy](data-sources--api_discovery--properties--user_defined_api_discovery_policy.md)
- user_defined_api_discovery_policy.exclusive

<a id="section"></a>

Type: `"single"`. Computed.

Exclusion Configuration. Configuration for exclusion action.

Upstream description:

Configuration for exclusion action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"archive\",\"ignore\"]"
}
```

## Direct properties

- [archive](data-sources--api_discovery--properties--user_defined_api_discovery_policy--exclusive--archive.md): complete subsection reference.

- [ignore](data-sources--api_discovery--properties--user_defined_api_discovery_policy--exclusive--ignore.md): complete subsection reference.

## Next pages

- [user_defined_api_discovery_policy.exclusive.archive](data-sources--api_discovery--properties--user_defined_api_discovery_policy--exclusive--archive.md)
- [user_defined_api_discovery_policy.exclusive.ignore](data-sources--api_discovery--properties--user_defined_api_discovery_policy--exclusive--ignore.md)
- [user_defined_api_discovery_policy](data-sources--api_discovery--properties--user_defined_api_discovery_policy.md)
- [xcsh_api_discovery](../data-sources/api_discovery.md)
