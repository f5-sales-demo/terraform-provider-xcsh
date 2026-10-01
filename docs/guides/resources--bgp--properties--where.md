---
page_title: "where"
subcategory: ""
description: "where for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 1591, "body_sha256": "sha256:0c0f863884501f8e7f97bd8904c7221d34b949fe92fc8e786fd6d93adf9b3a7a", "canonical_id": "xcsh-docs:resources:bgp:properties:where", "child_ids": ["xcsh-docs:resources:bgp:properties:where:site", "xcsh-docs:resources:bgp:properties:where:virtual_site"], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:where", "parent_id": "xcsh-docs:resources:bgp:reference", "path": "docs/guides/resources--bgp--properties--where.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["where"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/where/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "where for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md)
- [Property reference](resources--bgp--reference.md)
- where

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual\_site It
used to refer site or a group of sites indicated by virtual site.

Upstream description:

VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual\_site It
used to refer site or a group of sites indicated by virtual site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
```

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

Terraform syntax:

```terraform
where {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](resources--bgp--properties--where--site.md): complete subsection reference.

- [virtual_site](resources--bgp--properties--where--virtual_site.md): complete subsection reference.

## Next pages

- [where.site](resources--bgp--properties--where--site.md)
- [where.virtual_site](resources--bgp--properties--where--virtual_site.md)
- [Property reference](resources--bgp--reference.md)
- [xcsh_bgp](../resources/bgp.md)
