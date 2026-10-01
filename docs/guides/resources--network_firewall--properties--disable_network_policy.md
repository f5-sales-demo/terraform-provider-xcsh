---
page_title: "disable_network_policy"
subcategory: "Security"
description: "disable_network_policy for xcsh_network_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1017, "body_sha256": "sha256:3f36d3b10078c4ee0a3a899a19740f9bf0cfc600edf2d15cc0a441235aba1bc3", "canonical_id": "xcsh-docs:resources:network_firewall:properties:disable_network_policy", "child_ids": [], "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_firewall:properties:disable_network_policy", "parent_id": "xcsh-docs:resources:network_firewall:reference", "path": "docs/guides/resources--network_firewall--properties--disable_network_policy.md", "provider_name": "network_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_network_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/properties/disable_network_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_network_policy for xcsh_network_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_network_policy

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md)
- [Property reference](resources--network_firewall--reference.md)
- disable_network_policy

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Policy configuration for this feature. Defaults to \`map\[\]\`. Server applies default when omitted.

Upstream description:

This can be used for messages where no values are needed.

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
disable_network_policy = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--network_firewall--reference.md)
- [xcsh_network_firewall](../resources/network_firewall.md)
