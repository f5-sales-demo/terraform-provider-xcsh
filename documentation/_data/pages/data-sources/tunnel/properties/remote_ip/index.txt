---
page_title: "remote_ip"
subcategory: ""
description: "Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - 1. IP Address - Specifies the remote IP to which tunnel has to be connected 2. Remote endpoint - Is a map of IP address on per ver node basis."
xcsh_docs: {"aliases": ["remote ip"], "body_bytes": 1178, "body_sha256": "sha256:17164549814704b2cbe3119b6aa5f3750ce83f379178641a21c6bb84d79de50d", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:remote_ip:endpoints", "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:remote_ip", "parent_id": "xcsh-docs:data-sources:tunnel:reference", "path": "documentation/data-sources/tunnel/properties/remote_ip/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3132312023230012-3013110002122031-3132111221301201-3022331010323223-1232301300132031-2222202103210231-2231122010211301-2302103301333120", "registry_path": "docs/guides/data-sources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["remote_ip"], "schema_version": 1, "sections": [{"aliases": ["remote ip endpoints"], "anchor": "section", "description": "Provides a map of ver node name to remote node attributes Ver node should use these attributes to configure as remote tunnel.", "document_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:endpoints", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["remote_ip", "endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["remote ip ip"], "anchor": "section", "description": "IP Address used to specify an IPv4 or IPv6 address.", "document_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["remote_ip", "ip"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/remote_ip/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - 1. IP Address - Specifies the remote IP to which tunnel has to be connected 2. Remote endpoint - Is a map of IP address on per ver node basis.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["tunnelCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# remote_ip

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- remote_ip

<a id="section"></a>

Type: `"single"`. Computed.

Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - 1. IP
Address - Specifies the remote IP to which tunnel has to be connected 2. Remote endpoint - Is a map
of IP address on per ver node basis.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type": "[\"endpoints\",\"ip\"]"
}
```

## Direct properties

- [endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/endpoints/): complete subsection reference.

- [ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/ip/): complete subsection reference.
