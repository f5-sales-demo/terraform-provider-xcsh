---
page_title: "waf_signatures"
subcategory: ""
description: "Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied manually. Refer to release notes for details about available Signatures update modes."
xcsh_docs: {"aliases": ["waf signatures"], "body_bytes": 2236, "body_sha256": "sha256:8108535c8c5bdeefee202b21a6ff6b57fce27e2d472e9852715f8ec7816f96c3", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:waf_signatures:automatic", "xcsh-docs:resources:voltstack_site:properties:waf_signatures:manual"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:waf_signatures", "parent_id": "xcsh-docs:resources:voltstack_site:reference", "path": "documentation/resources/voltstack_site/properties/waf_signatures/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3001210003032302-1322031333003130-3131300313220213-0112200202131003-3012101010211301-1323112023133110-3121212221011012-2031210320020132", "registry_path": "docs/guides/resources--voltstack_site--reference--group-010.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "waf_signatures:ConflictingObjectAttributes:automatic,manual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:waf_signatures:automatic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_signatures:ConflictingObjectAttributes:automatic,manual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:waf_signatures:manual", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_signatures"], "schema_version": 1, "sections": [{"aliases": ["automatic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:waf_signatures:automatic", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_signatures", "automatic"], "syntax": "attribute", "type": "object"}, {"aliases": ["manual"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:waf_signatures:manual", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_signatures", "manual"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/waf_signatures/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied manually. Refer to release notes for details about available Signatures update modes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_signatures

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- waf_signatures

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

- [automatic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/waf_signatures/automatic/): complete subsection reference.

- [manual](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/waf_signatures/manual/): complete subsection reference.

## Next pages

- [waf_signatures.automatic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/waf_signatures/automatic/)
- [waf_signatures.manual](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/waf_signatures/manual/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
