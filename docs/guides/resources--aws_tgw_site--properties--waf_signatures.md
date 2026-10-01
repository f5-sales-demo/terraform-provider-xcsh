---
page_title: "waf_signatures"
subcategory: ""
description: "waf_signatures for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1808, "body_sha256": "sha256:58cfe3552f74a92a2f312669951775d8a1c34d811eff7faa0e1fa05f02dbb2a3", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:waf_signatures", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:waf_signatures:automatic", "xcsh-docs:resources:aws_tgw_site:properties:waf_signatures:manual"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:waf_signatures", "parent_id": "xcsh-docs:resources:aws_tgw_site:reference", "path": "docs/guides/resources--aws_tgw_site--properties--waf_signatures.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_signatures"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/waf_signatures/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_signatures for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_signatures

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
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

- [automatic](resources--aws_tgw_site--properties--waf_signatures--automatic.md): complete subsection reference.

- [manual](resources--aws_tgw_site--properties--waf_signatures--manual.md): complete subsection reference.

## Next pages

- [waf_signatures.automatic](resources--aws_tgw_site--properties--waf_signatures--automatic.md)
- [waf_signatures.manual](resources--aws_tgw_site--properties--waf_signatures--manual.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
