---
page_title: "where"
subcategory: ""
description: "where for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 1335, "body_sha256": "sha256:dd3afc2a9dbd71127d6d0aede1415728564a3632293d8c132d95b0f6779ce055", "canonical_id": "xcsh-docs:data-sources:bgp:properties:where", "child_ids": ["xcsh-docs:data-sources:bgp:properties:where:site", "xcsh-docs:data-sources:bgp:properties:where:virtual_site"], "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:where", "parent_id": "xcsh-docs:data-sources:bgp:reference", "path": "docs/guides/data-sources--bgp--properties--where.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["where"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/where/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "where for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md)
- [Property reference](data-sources--bgp--reference.md)
- where

<a id="section"></a>

Type: `"single"`. Computed.

VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual\_site It
used to refer site or a group of sites indicated by virtual site.

Upstream description:

VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual\_site It
used to refer site or a group of sites indicated by virtual site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_site\"]"
}
```

## Direct properties

- [site](data-sources--bgp--properties--where--site.md): complete subsection reference.

- [virtual_site](data-sources--bgp--properties--where--virtual_site.md): complete subsection reference.

## Next pages

- [where.site](data-sources--bgp--properties--where--site.md)
- [where.virtual_site](data-sources--bgp--properties--where--virtual_site.md)
- [Property reference](data-sources--bgp--reference.md)
- [xcsh_bgp](../data-sources/bgp.md)
