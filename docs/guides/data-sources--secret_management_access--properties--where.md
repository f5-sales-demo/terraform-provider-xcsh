---
page_title: "where"
subcategory: ""
description: "where for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 2049, "body_sha256": "sha256:2e95b3f10e7e81d9fc2464c25ee07216f182f8beeb878c13cd6f98b2f5d527c6", "canonical_id": "xcsh-docs:data-sources:secret_management_access:properties:where", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:where:site", "xcsh-docs:data-sources:secret_management_access:properties:where:virtual_network", "xcsh-docs:data-sources:secret_management_access:properties:where:virtual_site"], "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:where", "parent_id": "xcsh-docs:data-sources:secret_management_access:reference", "path": "docs/guides/data-sources--secret_management_access--properties--where.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["where"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/where/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "where for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# where

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md)
- [Property reference](data-sources--secret_management_access--reference.md)
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

- [site](data-sources--secret_management_access--properties--where--site.md): complete subsection reference.

- [virtual_network](data-sources--secret_management_access--properties--where--virtual_network.md): complete subsection reference.

- [virtual_site](data-sources--secret_management_access--properties--where--virtual_site.md): complete subsection reference.

## Next pages

- [where.site](data-sources--secret_management_access--properties--where--site.md)
- [where.virtual_network](data-sources--secret_management_access--properties--where--virtual_network.md)
- [where.virtual_site](data-sources--secret_management_access--properties--where--virtual_site.md)
- [Property reference](data-sources--secret_management_access--reference.md)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md)
