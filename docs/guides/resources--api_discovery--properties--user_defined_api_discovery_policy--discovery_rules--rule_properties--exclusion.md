---
page_title: "user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion"
subcategory: ""
description: "user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion for xcsh_api_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 2576, "body_sha256": "sha256:9c4a34b0661b2d50fec4de8583fabcfef5e874382fbdf9d70a2a99a3bc1c172f", "canonical_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:exclusion", "child_ids": ["xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:exclusion:archive", "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:exclusion:ignore"], "collection_id": "xcsh-docs:resources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:exclusion", "parent_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties", "path": "docs/guides/resources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--rule_properties--exclusion.md", "provider_name": "api_discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties", "exclusion"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/exclusion/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion for xcsh_api_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md)
- [Property reference](resources--api_discovery--reference.md)
- [user_defined_api_discovery_policy](resources--api_discovery--properties--user_defined_api_discovery_policy.md)
- [user_defined_api_discovery_policy.discovery_rules](resources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules.md)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--rule_properties.md)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion

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
exclusion {
  # Configure direct properties listed below.
}
```

## Direct properties

- [archive](resources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--rule_properties--exclusion--archive.md): complete subsection reference.

- [ignore](resources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--rule_properties--exclusion--ignore.md): complete subsection reference.

## Next pages

- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive](resources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--rule_properties--exclusion--archive.md)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore](resources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--rule_properties--exclusion--ignore.md)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--rule_properties.md)
- [xcsh_api_discovery](../resources/api_discovery.md)
