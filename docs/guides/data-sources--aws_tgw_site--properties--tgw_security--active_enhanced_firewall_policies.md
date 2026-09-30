---
page_title: "tgw_security.active_enhanced_firewall_policies"
subcategory: ""
description: "tgw_security.active_enhanced_firewall_policies for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1469, "body_sha256": "sha256:281d51c9bc25b3d4bb31d539ee85f4576feb7b5170485d6328bed85a6f9b634c", "canonical_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_enhanced_firewall_policies", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_enhanced_firewall_policies:enhanced_firewall_policies"], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_enhanced_firewall_policies", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security", "path": "docs/guides/data-sources--aws_tgw_site--properties--tgw_security--active_enhanced_firewall_policies.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tgw_security", "active_enhanced_firewall_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/tgw_security/active_enhanced_firewall_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tgw_security.active_enhanced_firewall_policies for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# tgw_security.active_enhanced_firewall_policies

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
- [Property reference](data-sources--aws_tgw_site--reference.md)
- [tgw_security](data-sources--aws_tgw_site--properties--tgw_security.md)
- tgw_security.active_enhanced_firewall_policies

<a id="section"></a>

Type: `"single"`. Computed.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

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

## Direct properties

- [enhanced_firewall_policies](data-sources--aws_tgw_site--properties--tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies.md): complete subsection reference.

## Next pages

- [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--aws_tgw_site--properties--tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies.md)
- [tgw_security](data-sources--aws_tgw_site--properties--tgw_security.md)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
