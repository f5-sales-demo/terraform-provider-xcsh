---
page_title: "enable_ai_enhancements"
subcategory: "Security"
description: "enable_ai_enhancements for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1519, "body_sha256": "sha256:97bf00cc5c329dc401095dc3c4ea45ac391646216fe091076679cb76c9331f68", "canonical_id": "xcsh-docs:data-sources:app_firewall:properties:enable_ai_enhancements", "child_ids": ["xcsh-docs:data-sources:app_firewall:properties:enable_ai_enhancements:mitigate_high_medium_risk_action", "xcsh-docs:data-sources:app_firewall:properties:enable_ai_enhancements:mitigate_high_risk_action"], "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:enable_ai_enhancements", "parent_id": "xcsh-docs:data-sources:app_firewall:reference", "path": "docs/guides/data-sources--app_firewall--properties--enable_ai_enhancements.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_ai_enhancements"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/enable_ai_enhancements/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_ai_enhancements for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_ai_enhancements

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md)
- [Property reference](data-sources--app_firewall--reference.md)
- enable_ai_enhancements

<a id="section"></a>

Type: `"single"`. Computed.

Actions complimented by the additional intelligence of the F5 AI Powered Risk-based analysis.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-risk_score_action_choice": "[\"mitigate_high_medium_risk_action\",\"mitigate_high_risk_action\"]"
}
```

## Direct properties

- [mitigate_high_medium_risk_action](data-sources--app_firewall--properties--enable_ai_enhancements--mitigate_high_medium_risk_action.md): complete subsection reference.

- [mitigate_high_risk_action](data-sources--app_firewall--properties--enable_ai_enhancements--mitigate_high_risk_action.md): complete subsection reference.

## Next pages

- [enable_ai_enhancements.mitigate_high_medium_risk_action](data-sources--app_firewall--properties--enable_ai_enhancements--mitigate_high_medium_risk_action.md)
- [enable_ai_enhancements.mitigate_high_risk_action](data-sources--app_firewall--properties--enable_ai_enhancements--mitigate_high_risk_action.md)
- [Property reference](data-sources--app_firewall--reference.md)
- [xcsh_app_firewall](../data-sources/app_firewall.md)
