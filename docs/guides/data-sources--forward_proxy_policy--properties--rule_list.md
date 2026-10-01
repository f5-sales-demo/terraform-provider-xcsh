---
page_title: "rule_list"
subcategory: "Security"
description: "rule_list for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": [], "body_bytes": 978, "body_sha256": "sha256:fdd0cc938b8f0ab053f52da419e6a86513d936b054f0f4cfaa4f8704ce940e4c", "canonical_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list", "child_ids": ["xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules"], "collection_id": "xcsh-docs:data-sources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list", "parent_id": "xcsh-docs:data-sources:forward_proxy_policy:reference", "path": "docs/guides/data-sources--forward_proxy_policy--properties--rule_list.md", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forward_proxy_policy/properties/rule_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list for xcsh_forward_proxy_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md)
- [Property reference](data-sources--forward_proxy_policy--reference.md)
- rule_list

<a id="section"></a>

Type: `"single"`. Computed.

Custom Rule List. List of custom rules.

Upstream description:

List of custom rules.

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

- [rules](data-sources--forward_proxy_policy--properties--rule_list--rules.md): complete subsection reference.

## Next pages

- [rule_list.rules](data-sources--forward_proxy_policy--properties--rule_list--rules.md)
- [Property reference](data-sources--forward_proxy_policy--reference.md)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md)
