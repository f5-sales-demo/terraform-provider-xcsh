---
page_title: "tgw_security.active_enhanced_firewall_policies"
subcategory: ""
description: "tgw_security.active_enhanced_firewall_policies for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1568, "body_sha256": "sha256:945136bfdf3213a6294748dad4ba3f4b761bca50ac742089d3fbc4202b4a9133", "canonical_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_enhanced_firewall_policies", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_enhanced_firewall_policies:enhanced_firewall_policies"], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_enhanced_firewall_policies", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security", "path": "docs/guides/data-sources--aws_tgw_site--properties--tgw_security--active_enhanced_firewall_policies.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tgw_security", "active_enhanced_firewall_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/tgw_security/active_enhanced_firewall_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tgw_security.active_enhanced_firewall_policies for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
