---
page_title: "remote_ip.ip.dual_stack.ipv4"
subcategory: ""
description: "remote_ip.ip.dual_stack.ipv4 for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1810, "body_sha256": "sha256:c80c5f7c11e144cdc19fce97cb6ae81a4ea748f62b3d1b56fce8b91c1ae1303a", "canonical_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip:dual_stack:ipv4", "child_ids": [], "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip:dual_stack:ipv4", "parent_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip:dual_stack", "path": "docs/guides/data-sources--tunnel--properties--remote_ip--ip--dual_stack--ipv4.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["remote_ip", "ip", "dual_stack", "ipv4"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/remote_ip/ip/dual_stack/ipv4/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "remote_ip.ip.dual_stack.ipv4 for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# remote_ip.ip.dual_stack.ipv4

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md)
- [Property reference](data-sources--tunnel--reference.md)
- [remote_ip](data-sources--tunnel--properties--remote_ip.md)
- [remote_ip.ip](data-sources--tunnel--properties--remote_ip--ip.md)
- [remote_ip.ip.dual_stack](data-sources--tunnel--properties--remote_ip--ip--dual_stack.md)
- remote_ip.ip.dual_stack.ipv4

<a id="section"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

## Direct properties

<a id="schema-remote_ip--ip--dual_stack--ipv4--addr"></a>

### addr property

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

## Next pages

- [remote_ip.ip.dual_stack](data-sources--tunnel--properties--remote_ip--ip--dual_stack.md)
- [xcsh_tunnel](../data-sources/tunnel.md)
