---
page_title: "rule_list"
subcategory: "Security"
description: "rule_list for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1222, "body_sha256": "sha256:d9c43c61e86fc7c373a4930cc8c807f1b880c495046d3babbf2fb79b787889bf", "canonical_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list", "child_ids": ["xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules"], "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list", "parent_id": "xcsh-docs:resources:forward_proxy_policy:reference", "path": "docs/guides/resources--forward_proxy_policy--properties--rule_list.md", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/properties/rule_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list for xcsh_forward_proxy_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md)
- [Property reference](resources--forward_proxy_policy--reference.md)
- rule_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Custom Rule List. List of custom rules.

Upstream description:

List of custom rules.

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

- [rules](resources--forward_proxy_policy--properties--rule_list--rules.md): complete subsection reference.

## Next pages

- [rule_list.rules](resources--forward_proxy_policy--properties--rule_list--rules.md)
- [Property reference](resources--forward_proxy_policy--reference.md)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md)
