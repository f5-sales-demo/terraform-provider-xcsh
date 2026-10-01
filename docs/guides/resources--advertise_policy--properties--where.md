---
page_title: "where"
subcategory: ""
description: "where for xcsh_advertise_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2456, "body_sha256": "sha256:9f08e2a232251d520b8cf4bd53606f7bce4993b995a1819d84b5a6c88ecd2431", "canonical_id": "xcsh-docs:resources:advertise_policy:properties:where", "child_ids": ["xcsh-docs:resources:advertise_policy:properties:where:site", "xcsh-docs:resources:advertise_policy:properties:where:virtual_network", "xcsh-docs:resources:advertise_policy:properties:where:virtual_site"], "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:advertise_policy:properties:where", "parent_id": "xcsh-docs:resources:advertise_policy:reference", "path": "docs/guides/resources--advertise_policy--properties--where.md", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["where"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/properties/where/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "where for xcsh_advertise_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md)
- [Property reference](resources--advertise_policy--reference.md)
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

- [site](resources--advertise_policy--properties--where--site.md): complete subsection reference.

- [virtual_network](resources--advertise_policy--properties--where--virtual_network.md): complete subsection reference.

- [virtual_site](resources--advertise_policy--properties--where--virtual_site.md): complete subsection reference.

## Next pages

- [where.site](resources--advertise_policy--properties--where--site.md)
- [where.virtual_network](resources--advertise_policy--properties--where--virtual_network.md)
- [where.virtual_site](resources--advertise_policy--properties--where--virtual_site.md)
- [Property reference](resources--advertise_policy--reference.md)
- [xcsh_advertise_policy](../resources/advertise_policy.md)
