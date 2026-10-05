---
page_title: "software_settings.waf_signatures"
subcategory: ""
description: "Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied manually. Refer to release notes for details about available Signatures update modes."
xcsh_docs: {"aliases": ["software settings waf signatures"], "body_bytes": 2575, "body_sha256": "sha256:78fcb5f62b901d91285f74b9e18cbbcb07c024cd30f03028fab0ff4dcdc845c7", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures:automatic", "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures:manual"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings", "path": "documentation/resources/securemesh_site_v2/properties/software_settings/waf_signatures/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0010013231031330-1121200000030022-2121223331000232-0001101321233133-3012232310002232-1123032033001322-2001130231000322-2233301113322031", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-017.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "software_settings.waf_signatures:ConflictingObjectAttributes:automatic,manual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures:automatic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "software_settings.waf_signatures:ConflictingObjectAttributes:automatic,manual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures:manual", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["software_settings", "waf_signatures"], "schema_version": 1, "sections": [{"aliases": ["software settings waf signatures automatic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures:automatic", "flags": ["optional", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["software_settings", "waf_signatures", "automatic"], "syntax": "attribute", "type": "object"}, {"aliases": ["software settings waf signatures manual"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:waf_signatures:manual", "flags": ["optional", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["software_settings", "waf_signatures", "manual"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/software_settings/waf_signatures/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied manually. Refer to release notes for details about available Signatures update modes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# software_settings.waf_signatures

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [software_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/)
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

- [automatic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/waf_signatures/automatic/): complete subsection reference.

- [manual](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/waf_signatures/manual/): complete subsection reference.

## Next pages

- [software_settings.waf_signatures.automatic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/waf_signatures/automatic/)
- [software_settings.waf_signatures.manual](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/waf_signatures/manual/)
- [software_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
