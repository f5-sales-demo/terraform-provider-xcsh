---
page_title: "params.ipsec"
subcategory: ""
description: "params.ipsec for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1069, "body_sha256": "sha256:e0c4eee0ee0728e134643107f17170e69a1f946d7a9d3d2f5bb452de086302c9", "canonical_id": "xcsh-docs:data-sources:tunnel:properties:params:ipsec", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:params:ipsec:ipsec_psk"], "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:params:ipsec", "parent_id": "xcsh-docs:data-sources:tunnel:properties:params", "path": "docs/guides/data-sources--tunnel--properties--params--ipsec.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["params", "ipsec"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/params/ipsec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "params.ipsec for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# params.ipsec

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md)
- [Property reference](data-sources--tunnel--reference.md)
- [params](data-sources--tunnel--properties--params.md)
- params.ipsec

<a id="section"></a>

Type: `"single"`. Computed.

Configuration for IPsec encapsulation are: 1. PSK - pre shared key to be used by IKE.

Upstream description:

Configuration for IPsec encapsulation are: &#8203;1. PSK - pre shared key to be used by IKE.

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

- [ipsec_psk](data-sources--tunnel--properties--params--ipsec--ipsec_psk.md): complete subsection reference.

## Next pages

- [params.ipsec.ipsec_psk](data-sources--tunnel--properties--params--ipsec--ipsec_psk.md)
- [params](data-sources--tunnel--properties--params.md)
- [xcsh_tunnel](../data-sources/tunnel.md)
