---
page_title: "where"
subcategory: "Networking"
description: "where for xcsh_endpoint."
xcsh_docs: {"aliases": [], "body_bytes": 2360, "body_sha256": "sha256:549319f1b3a7066c09b576a98a58a75ba4f7d3fa7287eff931b4618fa26f69bf", "canonical_id": "xcsh-docs:resources:endpoint:properties:where", "child_ids": ["xcsh-docs:resources:endpoint:properties:where:site", "xcsh-docs:resources:endpoint:properties:where:virtual_network", "xcsh-docs:resources:endpoint:properties:where:virtual_site"], "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:resources:endpoint:properties:where", "parent_id": "xcsh-docs:resources:endpoint:reference", "path": "docs/guides/resources--endpoint--properties--where.md", "provider_name": "endpoint", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["where"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/properties/where/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "where for xcsh_endpoint.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["endpointCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md)
- [Property reference](resources--endpoint--reference.md)
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

- [site](resources--endpoint--properties--where--site.md): complete subsection reference.

- [virtual_network](resources--endpoint--properties--where--virtual_network.md): complete subsection reference.

- [virtual_site](resources--endpoint--properties--where--virtual_site.md): complete subsection reference.

## Next pages

- [where.site](resources--endpoint--properties--where--site.md)
- [where.virtual_network](resources--endpoint--properties--where--virtual_network.md)
- [where.virtual_site](resources--endpoint--properties--where--virtual_site.md)
- [Property reference](resources--endpoint--reference.md)
- [xcsh_endpoint](../resources/endpoint.md)
