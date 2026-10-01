---
page_title: "params"
subcategory: ""
description: "params for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1087, "body_sha256": "sha256:ee30c020928c66940b87051f6db07cfb9f8b6ed53e3f5ef00b9c6c11c9f92a64", "canonical_id": "xcsh-docs:data-sources:tunnel:properties:params", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:params:ipsec"], "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:params", "parent_id": "xcsh-docs:data-sources:tunnel:reference", "path": "docs/guides/data-sources--tunnel--properties--params.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "params for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# params

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md)
- [Property reference](data-sources--tunnel--reference.md)
- params

<a id="section"></a>

Type: `"single"`. Computed.

Tunnel configuration parameters for supported encapsulation 1. IPsec is supported with PSK for which
PSK can be configured.

Upstream description:

Tunnel configuration parameters for supported encapsulation &#8203;1. IPsec is supported with PSK
for which PSK can be configured.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type": "[\"ipsec\"]"
}
```

## Direct properties

- [ipsec](data-sources--tunnel--properties--params--ipsec.md): complete subsection reference.

## Next pages

- [params.ipsec](data-sources--tunnel--properties--params--ipsec.md)
- [Property reference](data-sources--tunnel--reference.md)
- [xcsh_tunnel](../data-sources/tunnel.md)
