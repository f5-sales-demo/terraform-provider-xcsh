---
page_title: "where"
subcategory: ""
description: "where for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 2372, "body_sha256": "sha256:b399431a6bda03767049f29a965137587e4415d2427f3f27c798f7b1b83aad84", "canonical_id": "xcsh-docs:resources:discovery:properties:where", "child_ids": ["xcsh-docs:resources:discovery:properties:where:site", "xcsh-docs:resources:discovery:properties:where:virtual_network", "xcsh-docs:resources:discovery:properties:where:virtual_site"], "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:where", "parent_id": "xcsh-docs:resources:discovery:reference", "path": "docs/guides/resources--discovery--properties--where.md", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["where"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/where/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "where for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md)
- [Property reference](resources--discovery--reference.md)
- where

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local..

Upstream description:

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local networks for sites selected by referring to virtual\_site object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_network"),
  validators.ConflictingObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingObjectAttributes("virtual_network",
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
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_network\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
where {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](resources--discovery--properties--where--site.md): complete subsection reference.

- [virtual_network](resources--discovery--properties--where--virtual_network.md): complete subsection reference.

- [virtual_site](resources--discovery--properties--where--virtual_site.md): complete subsection reference.

## Next pages

- [where.site](resources--discovery--properties--where--site.md)
- [where.virtual_network](resources--discovery--properties--where--virtual_network.md)
- [where.virtual_site](resources--discovery--properties--where--virtual_site.md)
- [Property reference](resources--discovery--reference.md)
- [xcsh_discovery](../resources/discovery.md)
