---
page_title: "local_control_plane"
subcategory: ""
description: "local_control_plane for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1996, "body_sha256": "sha256:2ad122db09331330084f621a4f4f2216d1f10d5000e3036ea7c109bc0cbee7eb", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:inside_vn", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:outside_vn"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane", "parent_id": "xcsh-docs:data-sources:voltstack_site:reference", "path": "docs/guides/data-sources--voltstack_site--properties--local_control_plane.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_control_plane"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/local_control_plane/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_control_plane for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
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

- [local_control_plane](data-sources--voltstack_site--properties--local_control_plane.md#section)
- [no_local_control_plane](data-sources--voltstack_site--properties--no_local_control_plane.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [bgp_config](data-sources--voltstack_site--properties--local_control_plane--bgp_config.md): complete subsection reference.

- [inside_vn](data-sources--voltstack_site--properties--local_control_plane--inside_vn.md): complete subsection reference.

- [outside_vn](data-sources--voltstack_site--properties--local_control_plane--outside_vn.md): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config](data-sources--voltstack_site--properties--local_control_plane--bgp_config.md)
- [local_control_plane.inside_vn](data-sources--voltstack_site--properties--local_control_plane--inside_vn.md)
- [local_control_plane.outside_vn](data-sources--voltstack_site--properties--local_control_plane--outside_vn.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
