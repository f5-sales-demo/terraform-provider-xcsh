---
page_title: "site_virtual_sites"
subcategory: ""
description: "site_virtual_sites for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1154, "body_sha256": "sha256:f2262b025703c1255f5e54126344429de31b8c91de88f34c9a9b845723b223f9", "canonical_id": "xcsh-docs:resources:proxy:properties:site_virtual_sites", "child_ids": ["xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where"], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:site_virtual_sites", "parent_id": "xcsh-docs:resources:proxy:reference", "path": "docs/guides/resources--proxy--properties--site_virtual_sites.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site_virtual_sites"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/site_virtual_sites/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_virtual_sites for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# site_virtual_sites

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- site_virtual_sites

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a VIP on specific sites.

Upstream description:

This defines a way to advertise a VIP on specific sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
site_virtual_sites {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advertise_where](resources--proxy--properties--site_virtual_sites--advertise_where.md): complete subsection reference.

## Next pages

- [site_virtual_sites.advertise_where](resources--proxy--properties--site_virtual_sites--advertise_where.md)
- [Property reference](resources--proxy--reference.md)
- [xcsh_proxy](../resources/proxy.md)
