---
page_title: "where"
subcategory: ""
description: "where for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 2552, "body_sha256": "sha256:a116036b556d2b3b1e9a56c3faa9e75829536d89a28bd534b9c9cc4171956f23", "canonical_id": "xcsh-docs:resources:secret_management_access:properties:where", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:where:site", "xcsh-docs:resources:secret_management_access:properties:where:virtual_network", "xcsh-docs:resources:secret_management_access:properties:where:virtual_site"], "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:where", "parent_id": "xcsh-docs:resources:secret_management_access:reference", "path": "docs/guides/resources--secret_management_access--properties--where.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["where"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/where/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "where for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md)
- [Property reference](resources--secret_management_access--reference.md)
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

- [site](resources--secret_management_access--properties--where--site.md): complete subsection reference.

- [virtual_network](resources--secret_management_access--properties--where--virtual_network.md): complete subsection reference.

- [virtual_site](resources--secret_management_access--properties--where--virtual_site.md): complete subsection reference.

## Next pages

- [where.site](resources--secret_management_access--properties--where--site.md)
- [where.virtual_network](resources--secret_management_access--properties--where--virtual_network.md)
- [where.virtual_site](resources--secret_management_access--properties--where--virtual_site.md)
- [Property reference](resources--secret_management_access--reference.md)
- [xcsh_secret_management_access](../resources/secret_management_access.md)
