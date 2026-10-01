---
page_title: "where"
subcategory: ""
description: "where for xcsh_advertise_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2052, "body_sha256": "sha256:b66dcf91f013bd7140291c7801f55eb9a94dcd5aed2377d9cae17d64bcde8005", "canonical_id": "xcsh-docs:data-sources:advertise_policy:properties:where", "child_ids": ["xcsh-docs:data-sources:advertise_policy:properties:where:site", "xcsh-docs:data-sources:advertise_policy:properties:where:virtual_network", "xcsh-docs:data-sources:advertise_policy:properties:where:virtual_site"], "collection_id": "xcsh-docs:data-sources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:advertise_policy:properties:where", "parent_id": "xcsh-docs:data-sources:advertise_policy:reference", "path": "docs/guides/data-sources--advertise_policy--properties--where.md", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["where"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/advertise_policy/properties/where/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "where for xcsh_advertise_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md)
- [Property reference](data-sources--advertise_policy--reference.md)
- where

<a id="section"></a>

Type: `"single"`. Computed.

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local..

Upstream description:

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local networks for sites selected by referring to virtual\_site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_network\",\"virtual_site\"]"
}
```

## Direct properties

- [site](data-sources--advertise_policy--properties--where--site.md): complete subsection reference.

- [virtual_network](data-sources--advertise_policy--properties--where--virtual_network.md): complete subsection reference.

- [virtual_site](data-sources--advertise_policy--properties--where--virtual_site.md): complete subsection reference.

## Next pages

- [where.site](data-sources--advertise_policy--properties--where--site.md)
- [where.virtual_network](data-sources--advertise_policy--properties--where--virtual_network.md)
- [where.virtual_site](data-sources--advertise_policy--properties--where--virtual_site.md)
- [Property reference](data-sources--advertise_policy--reference.md)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md)
