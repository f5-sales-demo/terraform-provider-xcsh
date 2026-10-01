---
page_title: "local_control_plane"
subcategory: ""
description: "local_control_plane for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 2606, "body_sha256": "sha256:c12d8aa7315abd17a7f32f7527697f6e46ffd55b817c69b7c9a6931e98aa9806", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:inside_vn", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:outside_vn"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane", "parent_id": "xcsh-docs:data-sources:voltstack_site:reference", "path": "documentation/data-sources/voltstack_site/properties/local_control_plane/index.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["local_control_plane"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/local_control_plane/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_control_plane for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
