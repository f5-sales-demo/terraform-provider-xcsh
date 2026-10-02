---
page_title: "waf_signatures"
subcategory: "Infrastructure"
description: "Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied manually. Refer to release notes for details about available Signatures update modes."
xcsh_docs: {"aliases": ["waf signatures"], "body_bytes": 2216, "body_sha256": "sha256:84d0d23c6ad736badcb843cf7c64c51b6b8e83739adab7122428e4099e555b1c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:waf_signatures:automatic", "xcsh-docs:resources:aws_vpc_site:properties:waf_signatures:manual"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:waf_signatures", "parent_id": "xcsh-docs:resources:aws_vpc_site:reference", "path": "documentation/resources/aws_vpc_site/properties/waf_signatures/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2110022231300120-3130210123131302-1331212201020032-2333303013221203-0003211002203112-2311312323322011-1330000301331202-1312022031112120", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "waf_signatures:ConflictingObjectAttributes:automatic,manual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:waf_signatures:automatic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_signatures:ConflictingObjectAttributes:automatic,manual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:waf_signatures:manual", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_signatures"], "schema_version": 1, "sections": [{"aliases": ["automatic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:waf_signatures:automatic", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_signatures", "automatic"], "syntax": "attribute", "type": "object"}, {"aliases": ["manual"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:waf_signatures:manual", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_signatures", "manual"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/waf_signatures/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied manually. Refer to release notes for details about available Signatures update modes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_signatures

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
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

- [automatic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/waf_signatures/automatic/): complete subsection reference.

- [manual](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/waf_signatures/manual/): complete subsection reference.

## Next pages

- [waf_signatures.automatic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/waf_signatures/automatic/)
- [waf_signatures.manual](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/waf_signatures/manual/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
