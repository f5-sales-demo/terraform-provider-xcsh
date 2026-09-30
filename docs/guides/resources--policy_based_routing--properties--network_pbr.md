---
page_title: "network_pbr"
subcategory: ""
description: "network_pbr for xcsh_policy_based_routing."
xcsh_docs: {"aliases": [], "body_bytes": 2108, "body_sha256": "sha256:bfcf9937e3a2be94baea8678d889f07aa4a67b907de40667092293fd3c97eb2f", "canonical_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr", "child_ids": ["xcsh-docs:resources:policy_based_routing:properties:network_pbr:any", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:label_selector", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:prefix_list"], "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr", "parent_id": "xcsh-docs:resources:policy_based_routing:reference", "path": "docs/guides/resources--policy_based_routing--properties--network_pbr.md", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["network_pbr"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/properties/network_pbr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "network_pbr for xcsh_policy_based_routing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# network_pbr

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md)
- [Property reference](resources--policy_based_routing--reference.md)
- network_pbr

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for network pbr.

Upstream description:

Network(L3/L4) routing policy rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any",
    "label_selector"),
  validators.ConflictingObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingObjectAttributes("label_selector",
    "prefix_list")}
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
  "x-ves-oneof-field-source_choice": "[\"any\",\"label_selector\",\"prefix_list\"]"
}
```

Terraform syntax:

```terraform
network_pbr {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any](resources--policy_based_routing--properties--network_pbr--any.md): complete subsection reference.

- [label_selector](resources--policy_based_routing--properties--network_pbr--label_selector.md): complete subsection reference.

- [network_pbr_rules](resources--policy_based_routing--properties--network_pbr--network_pbr_rules.md): complete subsection reference.

- [prefix_list](resources--policy_based_routing--properties--network_pbr--prefix_list.md): complete subsection reference.

## Next pages

- [network_pbr.any](resources--policy_based_routing--properties--network_pbr--any.md)
- [network_pbr.label_selector](resources--policy_based_routing--properties--network_pbr--label_selector.md)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--properties--network_pbr--network_pbr_rules.md)
- [network_pbr.prefix_list](resources--policy_based_routing--properties--network_pbr--prefix_list.md)
- [Property reference](resources--policy_based_routing--reference.md)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md)
