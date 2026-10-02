---
page_title: "software_settings"
subcategory: ""
description: "Select OS and Software version for the site. All nodes in the site will run the same OS and Software version. These settings cannot be changed after the site is created."
xcsh_docs: {"aliases": ["software settings"], "body_bytes": 2496, "body_sha256": "sha256:881901bb1c054a6b5e3015f558ee66fbc486a260b03a0c600fe130786562fadf", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os", "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:sw", "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/software_settings/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["software_settings"], "schema_version": 1, "sections": [{"aliases": ["os"], "anchor": "section", "description": "Select the F5XC Operating System Version for the site. By default, latest available OS Version will be used. Refer to release notes to find required released OS versions.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-software_settings--os--operating_system_version", "enforcement": "provider-schema", "group": "software_settings.os:ConflictingObjectAttributes:default_os_version,operating_system_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "software_settings.os:ConflictingObjectAttributes:default_os_version,operating_system_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os:default_os_version", "type": "conflicts"}], "schema_path": ["software_settings", "os"], "syntax": "block", "type": "object"}, {"aliases": ["sw"], "anchor": "section", "description": "Select the F5XC Software Version for the site. By default, latest available F5XC Software Version will be used. Refer to release notes to find required released SW versions.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:sw", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-software_settings--sw--volterra_software_version", "enforcement": "provider-schema", "group": "software_settings.sw:ConflictingObjectAttributes:default_sw_version,volterra_software_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:sw", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "software_settings.sw:ConflictingObjectAttributes:default_sw_version,volterra_software_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:sw:default_sw_version", "type": "conflicts"}], "schema_path": ["software_settings", "sw"], "syntax": "block", "type": "object"}, {"aliases": ["waf signatures"], "anchor": "section", "description": "Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied manually. Refer to release notes for details about available Signatures update modes.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "software_settings.waf_signatures:ConflictingObjectAttributes:automatic,manual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures:automatic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "software_settings.waf_signatures:ConflictingObjectAttributes:automatic,manual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures:manual", "type": "conflicts"}], "schema_path": ["software_settings", "waf_signatures"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/software_settings/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Select OS and Software version for the site. All nodes in the site will run the same OS and Software version. These settings cannot be changed after the site is created.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

- [os](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/os/): complete subsection reference.

- [sw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/sw/): complete subsection reference.

- [waf_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/waf_signatures/): complete subsection reference.

## Next pages

- [software_settings.os](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/os/)
- [software_settings.sw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/sw/)
- [software_settings.waf_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/waf_signatures/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
