---
page_title: "software_settings"
subcategory: ""
description: "Select OS and Software version for the site. All nodes in the site will run the same OS and Software version. These settings cannot be changed after the site is created."
xcsh_docs: {"aliases": ["software settings"], "body_bytes": 1594, "body_sha256": "sha256:91db323b85060d27ab8f5d61d8f6902e82463012135b238d08b24a0e94ad35fd", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os", "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:sw", "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/software_settings/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["software_settings"], "schema_version": 1, "sections": [{"aliases": ["software settings os"], "anchor": "section", "description": "Select the F5XC Operating System Version for the site. By default, latest available OS Version will be used. Refer to release notes to find required released OS versions.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-software_settings--os--operating_system_version", "enforcement": "provider-schema", "group": "software_settings.os:ConflictingObjectAttributes:default_os_version,operating_system_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "software_settings.os:ConflictingObjectAttributes:default_os_version,operating_system_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os:default_os_version", "type": "conflicts"}], "schema_path": ["software_settings", "os"], "syntax": "block", "type": "object"}, {"aliases": ["software settings sw"], "anchor": "section", "description": "Select the F5XC Software Version for the site. By default, latest available F5XC Software Version will be used. Refer to release notes to find required released SW versions.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:sw", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-software_settings--sw--volterra_software_version", "enforcement": "provider-schema", "group": "software_settings.sw:ConflictingObjectAttributes:default_sw_version,volterra_software_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:sw", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "software_settings.sw:ConflictingObjectAttributes:default_sw_version,volterra_software_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:sw:default_sw_version", "type": "conflicts"}], "schema_path": ["software_settings", "sw"], "syntax": "block", "type": "object"}, {"aliases": ["software settings waf signatures"], "anchor": "section", "description": "Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied manually. Refer to release notes for details about available Signatures update modes.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "software_settings.waf_signatures:ConflictingObjectAttributes:automatic,manual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures:automatic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "software_settings.waf_signatures:ConflictingObjectAttributes:automatic,manual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures:manual", "type": "conflicts"}], "schema_path": ["software_settings", "waf_signatures"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/software_settings/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Select OS and Software version for the site. All nodes in the site will run the same OS and Software version. These settings cannot be changed after the site is created.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# software_settings

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- software_settings

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select OS and Software version for the site. All nodes in the site will run the same OS and Software
version. These settings cannot be changed after the site is created. This block is a create-only,
write-only input; changing it replaces the resource, and refresh preserves the configured value
without claiming XC observed it.

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

- [os](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/os/): complete subsection reference.

- [sw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/sw/): complete subsection reference.

- [waf_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/waf_signatures/): complete subsection reference.
