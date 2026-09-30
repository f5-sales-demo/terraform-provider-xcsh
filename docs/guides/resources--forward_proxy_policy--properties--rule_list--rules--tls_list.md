---
page_title: "rule_list.rules.tls_list"
subcategory: "Security"
description: "rule_list.rules.tls_list for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1174, "body_sha256": "sha256:62df32f2a9b993fb5fca91e095af9aa4b7e9da5a51a44f5f407cff0c69c9f959", "canonical_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:tls_list", "child_ids": ["xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:tls_list:tls_list"], "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:tls_list", "parent_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules", "path": "docs/guides/resources--forward_proxy_policy--properties--rule_list--rules--tls_list.md", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "tls_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/properties/rule_list/rules/tls_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.tls_list for xcsh_forward_proxy_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rule_list.rules.tls_list

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md)
- [Property reference](resources--forward_proxy_policy--reference.md)
- [rule_list](resources--forward_proxy_policy--properties--rule_list.md)
- [rule_list.rules](resources--forward_proxy_policy--properties--rule_list--rules.md)
- rule_list.rules.tls_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

DomainListType.

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
tls_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [tls_list](resources--forward_proxy_policy--properties--rule_list--rules--tls_list--tls_list.md): complete subsection reference.

## Next pages

- [rule_list.rules.tls_list.tls_list](resources--forward_proxy_policy--properties--rule_list--rules--tls_list--tls_list.md)
- [rule_list.rules](resources--forward_proxy_policy--properties--rule_list--rules.md)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md)
