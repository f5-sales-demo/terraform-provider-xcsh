---
page_title: "local_control_plane"
subcategory: ""
description: "local_control_plane for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 2507, "body_sha256": "sha256:c16dc90f008139c6569dd9b64df5eb7418a2caaddfb3be1a0f7e3ad1e5d7ee08", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:inside_vn", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:outside_vn"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane", "parent_id": "xcsh-docs:data-sources:voltstack_site:reference", "path": "documentation/data-sources/voltstack_site/properties/local_control_plane/index.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["local_control_plane"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/local_control_plane/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_control_plane for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# local_control_plane

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- local_control_plane

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: local\_control\_plane, no\_local\_control\_plane; Default: no\_local\_control\_plane\]
Enable local control plane for L3VPN, SRV6, EVPN etc.

Upstream description:

Enable local control plane for L3VPN, SRV6, EVPN etc.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_vn\",\"outside_vn\"]"
}
```

OneOf alternatives in this subsection:

- [local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/#section)
- [no_local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/no_local_control_plane/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [bgp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/): complete subsection reference.

- [inside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/inside_vn/): complete subsection reference.

- [outside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/outside_vn/): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/)
- [local_control_plane.inside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/inside_vn/)
- [local_control_plane.outside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/outside_vn/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
