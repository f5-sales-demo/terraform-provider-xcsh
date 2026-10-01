---
page_title: "params.ipsec"
subcategory: ""
description: "params.ipsec for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1162, "body_sha256": "sha256:20c73a2a4905deecbca3b267877cdd6984a1d72937baae34a0e0c07cd42a494b", "canonical_id": "xcsh-docs:resources:tunnel:properties:params:ipsec", "child_ids": ["xcsh-docs:resources:tunnel:properties:params:ipsec:ipsec_psk"], "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:params:ipsec", "parent_id": "xcsh-docs:resources:tunnel:properties:params", "path": "docs/guides/resources--tunnel--properties--params--ipsec.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["params", "ipsec"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/params/ipsec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "params.ipsec for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# params.ipsec

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md)
- [Property reference](resources--tunnel--reference.md)
- [params](resources--tunnel--properties--params.md)
- params.ipsec

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ipsec {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ipsec_psk](resources--tunnel--properties--params--ipsec--ipsec_psk.md): complete subsection reference.

## Next pages

- [params.ipsec.ipsec_psk](resources--tunnel--properties--params--ipsec--ipsec_psk.md)
- [params](resources--tunnel--properties--params.md)
- [xcsh_tunnel](../resources/tunnel.md)
