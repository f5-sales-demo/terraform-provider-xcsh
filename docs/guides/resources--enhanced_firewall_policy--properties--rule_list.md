---
page_title: "rule_list"
subcategory: ""
description: "rule_list for xcsh_enhanced_firewall_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1210, "body_sha256": "sha256:dc080e3c4bc829a3a2189174b37bc75512875dc45af76ec2dcae2d20ba5c52f5", "canonical_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list", "child_ids": ["xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules"], "collection_id": "xcsh-docs:resources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list", "parent_id": "xcsh-docs:resources:enhanced_firewall_policy:reference", "path": "docs/guides/resources--enhanced_firewall_policy--properties--rule_list.md", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/enhanced_firewall_policy/properties/rule_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list for xcsh_enhanced_firewall_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rule_list

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md)
- [Property reference](resources--enhanced_firewall_policy--reference.md)
- rule_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Custom Enhanced Firewall Policy Rules. Custom Enhanced Firewall Policy Rules.

Upstream description:

Custom Enhanced Firewall Policy Rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
```

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
rule_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [rules](resources--enhanced_firewall_policy--properties--rule_list--rules.md): complete subsection reference.

## Next pages

- [rule_list.rules](resources--enhanced_firewall_policy--properties--rule_list--rules.md)
- [Property reference](resources--enhanced_firewall_policy--reference.md)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md)
