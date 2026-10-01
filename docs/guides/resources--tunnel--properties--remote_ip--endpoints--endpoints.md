---
page_title: "remote_ip.endpoints.endpoints"
subcategory: ""
description: "remote_ip.endpoints.endpoints for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1781, "body_sha256": "sha256:a58d7a00c1f6d396258f177bb35878b72a60bc107906bad5da2b7fc8a017b57f", "canonical_id": "xcsh-docs:resources:tunnel:properties:remote_ip:endpoints:endpoints", "child_ids": [], "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:remote_ip:endpoints:endpoints", "parent_id": "xcsh-docs:resources:tunnel:properties:remote_ip:endpoints", "path": "docs/guides/resources--tunnel--properties--remote_ip--endpoints--endpoints.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["remote_ip", "endpoints", "endpoints"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/remote_ip/endpoints/endpoints/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "remote_ip.endpoints.endpoints for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# remote_ip.endpoints.endpoints

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md)
- [Property reference](resources--tunnel--reference.md)
- [remote_ip](resources--tunnel--properties--remote_ip.md)
- [remote_ip.endpoints](resources--tunnel--properties--remote_ip--endpoints.md)
- remote_ip.endpoints.endpoints

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Map of remote attributes to which tunnel will be established on per site node basis Every node can
have a different attributes and IP address to connect to Key is ver node name and value is Remote
node attributes.

Upstream description:

Map of remote attributes to which tunnel will be established on per site node basis Every node can
have a different attributes and IP address to connect to Key is ver node name and value is Remote
node attributes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

Terraform syntax:

```terraform
endpoints {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [remote_ip.endpoints](resources--tunnel--properties--remote_ip--endpoints.md)
- [xcsh_tunnel](../resources/tunnel.md)
