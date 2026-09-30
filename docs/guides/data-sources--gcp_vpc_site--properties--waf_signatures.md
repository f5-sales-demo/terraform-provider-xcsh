---
page_title: "waf_signatures"
subcategory: "Infrastructure"
description: "waf_signatures for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1445, "body_sha256": "sha256:0ee2ed3f6c6534a821cd64d7b089358bf14632e92e79f41c97a30b388d808e85", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:waf_signatures", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:properties:waf_signatures:automatic", "xcsh-docs:data-sources:gcp_vpc_site:properties:waf_signatures:manual"], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:waf_signatures", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:reference", "path": "docs/guides/data-sources--gcp_vpc_site--properties--waf_signatures.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_signatures"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/waf_signatures/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_signatures for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# waf_signatures

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
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

- [automatic](data-sources--gcp_vpc_site--properties--waf_signatures--automatic.md): complete subsection reference.

- [manual](data-sources--gcp_vpc_site--properties--waf_signatures--manual.md): complete subsection reference.

## Next pages

- [waf_signatures.automatic](data-sources--gcp_vpc_site--properties--waf_signatures--automatic.md)
- [waf_signatures.manual](data-sources--gcp_vpc_site--properties--waf_signatures--manual.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
