---
page_title: "peers.external.interface_list"
subcategory: ""
description: "peers.external.interface_list for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 1111, "body_sha256": "sha256:f6c321f0e74fdbdd754e6a9f48fb67790f09891d7ebf586a3a008756dc77e960", "canonical_id": "xcsh-docs:data-sources:bgp:properties:peers:external:interface_list", "child_ids": ["xcsh-docs:data-sources:bgp:properties:peers:external:interface_list:interfaces"], "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:external:interface_list", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers:external", "path": "docs/guides/data-sources--bgp--properties--peers--external--interface_list.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "external", "interface_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/external/interface_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.external.interface_list for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.interface_list

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md)
- [Property reference](data-sources--bgp--reference.md)
- [peers](data-sources--bgp--properties--peers.md)
- [peers.external](data-sources--bgp--properties--peers--external.md)
- peers.external.interface_list

<a id="section"></a>

Type: `"single"`. Computed.

Interface List. List of network interfaces.

Upstream description:

List of network interfaces.

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

- [interfaces](data-sources--bgp--properties--peers--external--interface_list--interfaces.md): complete subsection reference.

## Next pages

- [peers.external.interface_list.interfaces](data-sources--bgp--properties--peers--external--interface_list--interfaces.md)
- [peers.external](data-sources--bgp--properties--peers--external.md)
- [xcsh_bgp](../data-sources/bgp.md)
