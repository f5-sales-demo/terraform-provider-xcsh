---
page_title: "peers.external.family_inet"
subcategory: ""
description: "peers.external.family_inet for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 1293, "body_sha256": "sha256:8b0c1279f802aa93e5969ff553376508400b57932d3ddfffeaf2b1897b9928fa", "canonical_id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet", "child_ids": ["xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:disable_spec", "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable"], "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers:external", "path": "docs/guides/data-sources--bgp--properties--peers--external--family_inet.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "external", "family_inet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/external/family_inet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.external.family_inet for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# peers.external.family_inet

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md)
- [Property reference](data-sources--bgp--reference.md)
- [peers](data-sources--bgp--properties--peers.md)
- [peers.external](data-sources--bgp--properties--peers--external.md)
- peers.external.family_inet

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for family inet.

Upstream description:

Parameters for inet family.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-enable_choice": "[\"disable\",\"enable\"]"
}
```

## Direct properties

- [disable_spec](data-sources--bgp--properties--peers--external--family_inet--disable_spec.md): complete subsection reference.

- [enable](data-sources--bgp--properties--peers--external--family_inet--enable.md): complete subsection reference.

## Next pages

- [peers.external.family_inet.disable_spec](data-sources--bgp--properties--peers--external--family_inet--disable_spec.md)
- [peers.external.family_inet.enable](data-sources--bgp--properties--peers--external--family_inet--enable.md)
- [peers.external](data-sources--bgp--properties--peers--external.md)
- [xcsh_bgp](../data-sources/bgp.md)
