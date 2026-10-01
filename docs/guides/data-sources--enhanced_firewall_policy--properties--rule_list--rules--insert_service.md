---
page_title: "rule_list.rules.insert_service"
subcategory: ""
description: "rule_list.rules.insert_service for xcsh_enhanced_firewall_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1293, "body_sha256": "sha256:45386f8c7ec482ff3de1b7eea006bace0e7b8bbf8e560e9c4762a646b82f5cbe", "canonical_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:insert_service", "child_ids": ["xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:insert_service:nfv_service"], "collection_id": "xcsh-docs:data-sources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:insert_service", "parent_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules", "path": "docs/guides/data-sources--enhanced_firewall_policy--properties--rule_list--rules--insert_service.md", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "insert_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/enhanced_firewall_policy/properties/rule_list/rules/insert_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.insert_service for xcsh_enhanced_firewall_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.insert_service

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md)
- [Property reference](data-sources--enhanced_firewall_policy--reference.md)
- [rule_list](data-sources--enhanced_firewall_policy--properties--rule_list.md)
- [rule_list.rules](data-sources--enhanced_firewall_policy--properties--rule_list--rules.md)
- rule_list.rules.insert_service

<a id="section"></a>

Type: `"single"`. Computed.

Action to forward traffic to external service.

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

- [nfv_service](data-sources--enhanced_firewall_policy--properties--rule_list--rules--insert_service--nfv_service.md): complete subsection reference.

## Next pages

- [rule_list.rules.insert_service.nfv_service](data-sources--enhanced_firewall_policy--properties--rule_list--rules--insert_service--nfv_service.md)
- [rule_list.rules](data-sources--enhanced_firewall_policy--properties--rule_list--rules.md)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md)
