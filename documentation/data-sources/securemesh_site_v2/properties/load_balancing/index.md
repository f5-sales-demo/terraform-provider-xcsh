---
page_title: "load_balancing"
subcategory: ""
description: "load_balancing for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1717, "body_sha256": "sha256:33c2acd1d3deee057800999a2f992cb60f76191d291a6625dadac1b505e902ac", "child_ids": [], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:load_balancing", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "documentation/data-sources/securemesh_site_v2/properties/load_balancing/index.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["load_balancing"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/load_balancing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "load_balancing for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# load_balancing

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- load_balancing

<a id="section"></a>

Type: `"single"`. Computed.

Section contains settings on the site that relate to Load Balancing functionality.

Upstream description:

This section contains settings on the site that relate to Load Balancing functionality.

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

<a id="schema-load_balancing--vip_vrrp_mode"></a>

### vip_vrrp_mode property

Type: `"string"`. Computed.

\[Enum: VIP\_VRRP\_INVALID|VIP\_VRRP\_ENABLE|VIP\_VRRP\_DISABLE\] VRRP advertisement mode for VIP
Invalid VRRP mode. Possible values are \`VIP\_VRRP\_INVALID\`, \`VIP\_VRRP\_ENABLE\`,
\`VIP\_VRRP\_DISABLE\`. Defaults to \`VIP\_VRRP\_INVALID\`.

Upstream description:

VRRP advertisement mode for VIP

Invalid VRRP mode.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIP_VRRP_INVALID",
  "enum": [
    "VIP_VRRP_INVALID",
    "VIP_VRRP_ENABLE",
    "VIP_VRRP_DISABLE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
