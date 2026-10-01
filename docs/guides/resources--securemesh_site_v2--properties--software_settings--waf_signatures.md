---
page_title: "software_settings.waf_signatures"
subcategory: ""
description: "software_settings.waf_signatures for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2122, "body_sha256": "sha256:7c4e7b36f5ff7a51a87451364f439e0c64defdc31be4555c77a8f8a5677e681b", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures:automatic", "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures:manual"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings", "path": "docs/guides/resources--securemesh_site_v2--properties--software_settings--waf_signatures.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["software_settings", "waf_signatures"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/software_settings/waf_signatures/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "software_settings.waf_signatures for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# software_settings.waf_signatures

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [software_settings](resources--securemesh_site_v2--properties--software_settings.md)
- software_settings.waf_signatures

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Upstream description:

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("automatic",
    "manual")}
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
  "x-ves-oneof-field-signatures_update_mode_choice": "[\"automatic\",\"manual\"]"
}
```

Terraform syntax:

```terraform
waf_signatures {
  # Configure direct properties listed below.
}
```

## Direct properties

- [automatic](resources--securemesh_site_v2--properties--software_settings--waf_signatures--automatic.md): complete subsection reference.

- [manual](resources--securemesh_site_v2--properties--software_settings--waf_signatures--manual.md): complete subsection reference.

## Next pages

- [software_settings.waf_signatures.automatic](resources--securemesh_site_v2--properties--software_settings--waf_signatures--automatic.md)
- [software_settings.waf_signatures.manual](resources--securemesh_site_v2--properties--software_settings--waf_signatures--manual.md)
- [software_settings](resources--securemesh_site_v2--properties--software_settings.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
