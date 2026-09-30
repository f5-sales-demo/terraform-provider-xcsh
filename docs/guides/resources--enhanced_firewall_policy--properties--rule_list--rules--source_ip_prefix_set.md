---
page_title: "rule_list.rules.source_ip_prefix_set"
subcategory: ""
description: "rule_list.rules.source_ip_prefix_set for xcsh_enhanced_firewall_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1370, "body_sha256": "sha256:443225f9253b2d67585de2ef1e0caa79d61028e024de2acbef1fd07ce0513b11", "canonical_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:source_ip_prefix_set", "child_ids": ["xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:source_ip_prefix_set:ref"], "collection_id": "xcsh-docs:resources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:source_ip_prefix_set", "parent_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules", "path": "docs/guides/resources--enhanced_firewall_policy--properties--rule_list--rules--source_ip_prefix_set.md", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "source_ip_prefix_set"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/enhanced_firewall_policy/properties/rule_list/rules/source_ip_prefix_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.source_ip_prefix_set for xcsh_enhanced_firewall_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rule_list.rules.source_ip_prefix_set

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md)
- [Property reference](resources--enhanced_firewall_policy--reference.md)
- [rule_list](resources--enhanced_firewall_policy--properties--rule_list.md)
- [rule_list.rules](resources--enhanced_firewall_policy--properties--rule_list--rules.md)
- rule_list.rules.source_ip_prefix_set

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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

Terraform syntax:

```terraform
source_ip_prefix_set {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ref](resources--enhanced_firewall_policy--properties--rule_list--rules--source_ip_prefix_set--ref.md): complete subsection reference.

## Next pages

- [rule_list.rules.source_ip_prefix_set.ref](resources--enhanced_firewall_policy--properties--rule_list--rules--source_ip_prefix_set--ref.md)
- [rule_list.rules](resources--enhanced_firewall_policy--properties--rule_list--rules.md)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md)
