---
page_title: "user_defined_api_discovery_policy"
subcategory: ""
description: "user_defined_api_discovery_policy for xcsh_api_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 2093, "body_sha256": "sha256:7af64f2e003be0d460e6616910321a92df54d97e7d5d5c360f292d4c9605fcb9", "canonical_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy", "child_ids": ["xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules", "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive", "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:inclusive"], "collection_id": "xcsh-docs:resources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy", "parent_id": "xcsh-docs:resources:api_discovery:reference", "path": "docs/guides/resources--api_discovery--properties--user_defined_api_discovery_policy.md", "provider_name": "api_discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["user_defined_api_discovery_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_discovery/properties/user_defined_api_discovery_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "user_defined_api_discovery_policy for xcsh_api_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_defined_api_discovery_policy

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md)
- [Property reference](resources--api_discovery--reference.md)
- user_defined_api_discovery_policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Rules are evaluated sequentially, top to bottom. If no rules are added, all traffic will be
discovered or ignored based on the selection in the 'Default Behaviour of the Rule Set' field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exclusive",
    "inclusive")}
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
  "x-ves-oneof-field-default_behavior_choice": "[\"exclusive\",\"inclusive\"]"
}
```

Terraform syntax:

```terraform
user_defined_api_discovery_policy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [discovery_rules](resources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules.md): complete subsection reference.

- [exclusive](resources--api_discovery--properties--user_defined_api_discovery_policy--exclusive.md): complete subsection reference.

- [inclusive](resources--api_discovery--properties--user_defined_api_discovery_policy--inclusive.md): complete subsection reference.

## Next pages

- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules.md)
- [user_defined_api_discovery_policy.exclusive](resources--api_discovery--properties--user_defined_api_discovery_policy--exclusive.md)
- [user_defined_api_discovery_policy.inclusive](resources--api_discovery--properties--user_defined_api_discovery_policy--inclusive.md)
- [Property reference](resources--api_discovery--reference.md)
- [xcsh_api_discovery](../resources/api_discovery.md)
