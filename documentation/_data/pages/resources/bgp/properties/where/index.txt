---
page_title: "where"
subcategory: ""
description: "where for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 1900, "body_sha256": "sha256:2ee73143d88f639143451ef1fc192087c2030c9b7aafcc77a53f3f94bec09b8a", "child_ids": ["xcsh-docs:resources:bgp:properties:where:site", "xcsh-docs:resources:bgp:properties:where:virtual_site"], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:where", "parent_id": "xcsh-docs:resources:bgp:reference", "path": "documentation/resources/bgp/properties/where/index.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["where"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/where/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "where for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# where

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
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

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/site/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/virtual_site/): complete subsection reference.

## Next pages

- [where.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/site/)
- [where.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/virtual_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
