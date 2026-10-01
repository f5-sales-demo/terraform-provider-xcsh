---
page_title: "user_defined_api_discovery_policy.exclusive"
subcategory: ""
description: "user_defined_api_discovery_policy.exclusive for xcsh_api_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1919, "body_sha256": "sha256:2df8bdbfe2f5c57b38b9e9ff75586284dcfdcccc961590ba050bac998a8dcdfd", "canonical_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive", "child_ids": ["xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:archive", "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:ignore"], "collection_id": "xcsh-docs:resources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive", "parent_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy", "path": "docs/guides/resources--api_discovery--properties--user_defined_api_discovery_policy--exclusive.md", "provider_name": "api_discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["user_defined_api_discovery_policy", "exclusive"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "user_defined_api_discovery_policy.exclusive for xcsh_api_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_defined_api_discovery_policy.exclusive

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md)
- [Property reference](resources--api_discovery--reference.md)
- [user_defined_api_discovery_policy](resources--api_discovery--properties--user_defined_api_discovery_policy.md)
- user_defined_api_discovery_policy.exclusive

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Exclusion Configuration. Configuration for exclusion action.

Upstream description:

Configuration for exclusion action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("archive",
    "ignore")}
```

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

Terraform syntax:

```terraform
exclusive {
  # Configure direct properties listed below.
}
```

## Direct properties

- [archive](resources--api_discovery--properties--user_defined_api_discovery_policy--exclusive--archive.md): complete subsection reference.

- [ignore](resources--api_discovery--properties--user_defined_api_discovery_policy--exclusive--ignore.md): complete subsection reference.

## Next pages

- [user_defined_api_discovery_policy.exclusive.archive](resources--api_discovery--properties--user_defined_api_discovery_policy--exclusive--archive.md)
- [user_defined_api_discovery_policy.exclusive.ignore](resources--api_discovery--properties--user_defined_api_discovery_policy--exclusive--ignore.md)
- [user_defined_api_discovery_policy](resources--api_discovery--properties--user_defined_api_discovery_policy.md)
- [xcsh_api_discovery](../resources/api_discovery.md)
