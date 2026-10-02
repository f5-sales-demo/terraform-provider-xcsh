---
page_title: "remote_ip.endpoints.endpoints"
subcategory: ""
description: "Map of remote attributes to which tunnel will be established on per site node basis Every node can have a different attributes and IP address to connect to Key is ver node name and value is Remote node attributes."
xcsh_docs: {"aliases": ["remote ip endpoints endpoints"], "body_bytes": 2034, "body_sha256": "sha256:2740123406ceca93dc14a30eaeca29eafc2061104b72a2bc7b6e4486bc65d0a8", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:endpoints:endpoints", "parent_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:endpoints", "path": "documentation/data-sources/tunnel/properties/remote_ip/endpoints/endpoints/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0001021301213332-0110321233030033-3003203203122222-1221300300330300-3321031210001310-3303113213022033-1121323231023320-0330333210330033", "registry_path": "docs/guides/data-sources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["remote_ip", "endpoints", "endpoints"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/remote_ip/endpoints/endpoints/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Map of remote attributes to which tunnel will be established on per site node basis Every node can have a different attributes and IP address to connect to Key is ver node name and value is Remote node attributes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# remote_ip.endpoints.endpoints

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- [remote_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/)
- [remote_ip.endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/endpoints/)
- remote_ip.endpoints.endpoints

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [remote_ip.endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/endpoints/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
