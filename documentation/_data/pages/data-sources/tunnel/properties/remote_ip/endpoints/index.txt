---
page_title: "remote_ip.endpoints"
subcategory: ""
description: "Provides a map of ver node name to remote node attributes Ver node should use these attributes to configure as remote tunnel."
xcsh_docs: {"aliases": ["remote ip endpoints"], "body_bytes": 1018, "body_sha256": "sha256:4b4a268d6807ffbb3bb55a8c7b417e450bf21b0d7f69989dda558d66cb5f6a91", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:remote_ip:endpoints:endpoints"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:endpoints", "parent_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip", "path": "documentation/data-sources/tunnel/properties/remote_ip/endpoints/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3003001003012231-2020031310333033-1120113023030332-2131213220100033-3331220330322202-1332221101013030-1121210230102320-0332331033123313", "registry_path": "docs/guides/data-sources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["remote_ip", "endpoints"], "schema_version": 1, "sections": [{"aliases": ["remote ip endpoints endpoints"], "anchor": "section", "description": "Map of remote attributes to which tunnel will be established on per site node basis Every node can have a different attributes and IP address to connect to Key is ver node name and value is Remote node attributes.", "document_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:endpoints:endpoints", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["remote_ip", "endpoints", "endpoints"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/remote_ip/endpoints/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Provides a map of ver node name to remote node attributes Ver node should use these attributes to configure as remote tunnel.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["tunnelCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# remote_ip.endpoints

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- [remote_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/)
- remote_ip.endpoints

<a id="section"></a>

Type: `"single"`. Computed.

Provides a map of ver node name to remote node attributes Ver node should use these attributes to
configure as remote tunnel.

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

- [endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/endpoints/endpoints/): complete subsection reference.
