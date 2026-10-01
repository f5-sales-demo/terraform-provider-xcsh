---
page_title: "disable_fast_acl"
subcategory: "Security"
description: "disable_fast_acl for xcsh_network_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1006, "body_sha256": "sha256:17a3b2feba78cb1f3d855b71badd784a6e27a778a9e28c11519649d777441f6b", "canonical_id": "xcsh-docs:resources:network_firewall:properties:disable_fast_acl", "child_ids": [], "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_firewall:properties:disable_fast_acl", "parent_id": "xcsh-docs:resources:network_firewall:reference", "path": "docs/guides/resources--network_firewall--properties--disable_fast_acl.md", "provider_name": "network_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_fast_acl"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/properties/disable_fast_acl/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_fast_acl for xcsh_network_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_fast_acl

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md)
- [Property reference](resources--network_firewall--reference.md)
- disable_fast_acl

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable fast acl. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
disable_fast_acl = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--network_firewall--reference.md)
- [xcsh_network_firewall](../resources/network_firewall.md)
