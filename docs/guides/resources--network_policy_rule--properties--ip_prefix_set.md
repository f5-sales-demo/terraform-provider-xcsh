---
page_title: "ip_prefix_set"
subcategory: ""
description: "ip_prefix_set for xcsh_network_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1533, "body_sha256": "sha256:920652f2c3fc7dc214930ba0c371d5c43b321a0d65ce2accb883e360d61fd05e", "canonical_id": "xcsh-docs:resources:network_policy_rule:properties:ip_prefix_set", "child_ids": ["xcsh-docs:resources:network_policy_rule:properties:ip_prefix_set:ref"], "collection_id": "xcsh-docs:resources:network_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_rule:properties:ip_prefix_set", "parent_id": "xcsh-docs:resources:network_policy_rule:reference", "path": "docs/guides/resources--network_policy_rule--properties--ip_prefix_set.md", "provider_name": "network_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ip_prefix_set"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_rule/properties/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ip_prefix_set for xcsh_network_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ip_prefix_set

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md)
- [Property reference](resources--network_policy_rule--reference.md)
- ip_prefix_set

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ip\_prefix\_set, prefix, prefix\_selector\] List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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

OneOf alternatives in this subsection:

- [ip_prefix_set](resources--network_policy_rule--properties--ip_prefix_set.md#section)
- [prefix](resources--network_policy_rule--properties--prefix.md#section)
- [prefix_selector](resources--network_policy_rule--properties--prefix_selector.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ip_prefix_set {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ref](resources--network_policy_rule--properties--ip_prefix_set--ref.md): complete subsection reference.

## Next pages

- [ip_prefix_set.ref](resources--network_policy_rule--properties--ip_prefix_set--ref.md)
- [Property reference](resources--network_policy_rule--reference.md)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md)
