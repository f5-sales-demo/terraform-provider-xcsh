---
page_title: "active_network_policies"
subcategory: "Security"
description: "active_network_policies for xcsh_network_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1351, "body_sha256": "sha256:4ae866a9cfd222fd73429130535af06d2c8e4b8ef9b0e241943ef8244ba04848", "canonical_id": "xcsh-docs:resources:network_firewall:properties:active_network_policies", "child_ids": ["xcsh-docs:resources:network_firewall:properties:active_network_policies:network_policies"], "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_firewall:properties:active_network_policies", "parent_id": "xcsh-docs:resources:network_firewall:reference", "path": "docs/guides/resources--network_firewall--properties--active_network_policies.md", "provider_name": "network_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["active_network_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/properties/active_network_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "active_network_policies for xcsh_network_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_network_policies

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md)
- [Property reference](resources--network_firewall--reference.md)
- active_network_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
```

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

Terraform syntax:

```terraform
active_network_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [network_policies](resources--network_firewall--properties--active_network_policies--network_policies.md): complete subsection reference.

## Next pages

- [active_network_policies.network_policies](resources--network_firewall--properties--active_network_policies--network_policies.md)
- [Property reference](resources--network_firewall--reference.md)
- [xcsh_network_firewall](../resources/network_firewall.md)
