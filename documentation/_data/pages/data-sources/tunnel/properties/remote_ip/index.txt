---
page_title: "remote_ip"
subcategory: ""
description: "remote_ip for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1824, "body_sha256": "sha256:2ec62ac907eef221024c9d7b96daf8ec404cef9aa9f6cebe0f6668357ce2ae17", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:remote_ip:endpoints", "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip"], "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:remote_ip", "parent_id": "xcsh-docs:data-sources:tunnel:reference", "path": "documentation/data-sources/tunnel/properties/remote_ip/index.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["remote_ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/remote_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "remote_ip for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

Upstream description:

Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - &#8203;1.
IP Address - Specifies the remote IP to which tunnel has to be connected &#8203;2. Remote endpoint -
Is a map of IP address on per ver node basis.

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

## Next pages

- [remote_ip.endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/endpoints/)
- [remote_ip.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/ip/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
