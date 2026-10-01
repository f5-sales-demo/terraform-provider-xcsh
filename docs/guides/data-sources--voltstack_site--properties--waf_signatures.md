---
page_title: "waf_signatures"
subcategory: ""
description: "waf_signatures for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1564, "body_sha256": "sha256:26a337821ad84e23e2abd2b38311e9726c085e296dd14243038425713a70df54", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:waf_signatures", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:waf_signatures:automatic", "xcsh-docs:data-sources:voltstack_site:properties:waf_signatures:manual"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:waf_signatures", "parent_id": "xcsh-docs:data-sources:voltstack_site:reference", "path": "docs/guides/data-sources--voltstack_site--properties--waf_signatures.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_signatures"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/waf_signatures/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_signatures for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_signatures

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- waf_signatures

<a id="section"></a>

Type: `"single"`. Computed.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Upstream description:

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

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

## Direct properties

- [automatic](data-sources--voltstack_site--properties--waf_signatures--automatic.md): complete subsection reference.

- [manual](data-sources--voltstack_site--properties--waf_signatures--manual.md): complete subsection reference.

## Next pages

- [waf_signatures.automatic](data-sources--voltstack_site--properties--waf_signatures--automatic.md)
- [waf_signatures.manual](data-sources--voltstack_site--properties--waf_signatures--manual.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
