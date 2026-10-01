---
page_title: "software_settings"
subcategory: ""
description: "software_settings for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1988, "body_sha256": "sha256:dffe42d22ba983f4a1d495ffc63785175737171a58bfab580df65cc5e937f761", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os", "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:sw", "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "docs/guides/resources--securemesh_site_v2--properties--software_settings.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["software_settings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/software_settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "software_settings for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# software_settings

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- software_settings

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select OS and Software version for the site. All nodes in the site will run the same OS and Software
version. These settings cannot be changed after the site is created. This block is a create-only,
write-only input; changing it replaces the resource, and refresh preserves the configured value
without claiming XC observed it.

Upstream description:

Select OS and Software version for the site. All nodes in the site will run the same OS and Software
version. These settings cannot be changed after the site is created.

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
software_settings {
  # Configure direct properties listed below.
}
```

## Direct properties

- [os](resources--securemesh_site_v2--properties--software_settings--os.md): complete subsection reference.

- [sw](resources--securemesh_site_v2--properties--software_settings--sw.md): complete subsection reference.

- [waf_signatures](resources--securemesh_site_v2--properties--software_settings--waf_signatures.md): complete subsection reference.

## Next pages

- [software_settings.os](resources--securemesh_site_v2--properties--software_settings--os.md)
- [software_settings.sw](resources--securemesh_site_v2--properties--software_settings--sw.md)
- [software_settings.waf_signatures](resources--securemesh_site_v2--properties--software_settings--waf_signatures.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
