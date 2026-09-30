---
page_title: "where"
subcategory: "Networking"
description: "where for xcsh_endpoint."
xcsh_docs: {"aliases": [], "body_bytes": 1857, "body_sha256": "sha256:d3990a2deaffc1d217edf4f5a596293e4d0fe7368eba766b16ff6aa49878b203", "canonical_id": "xcsh-docs:data-sources:endpoint:properties:where", "child_ids": ["xcsh-docs:data-sources:endpoint:properties:where:site", "xcsh-docs:data-sources:endpoint:properties:where:virtual_network", "xcsh-docs:data-sources:endpoint:properties:where:virtual_site"], "collection_id": "xcsh-docs:data-sources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:endpoint:properties:where", "parent_id": "xcsh-docs:data-sources:endpoint:reference", "path": "docs/guides/data-sources--endpoint--properties--where.md", "provider_name": "endpoint", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["where"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/endpoint/properties/where/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "where for xcsh_endpoint.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["endpointCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# where

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md)
- [Property reference](data-sources--endpoint--reference.md)
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

- [site](data-sources--endpoint--properties--where--site.md): complete subsection reference.

- [virtual_network](data-sources--endpoint--properties--where--virtual_network.md): complete subsection reference.

- [virtual_site](data-sources--endpoint--properties--where--virtual_site.md): complete subsection reference.

## Next pages

- [where.site](data-sources--endpoint--properties--where--site.md)
- [where.virtual_network](data-sources--endpoint--properties--where--virtual_network.md)
- [where.virtual_site](data-sources--endpoint--properties--where--virtual_site.md)
- [Property reference](data-sources--endpoint--reference.md)
- [xcsh_endpoint](../data-sources/endpoint.md)
