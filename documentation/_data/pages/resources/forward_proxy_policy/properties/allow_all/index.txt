---
page_title: "allow_all"
subcategory: "Security"
description: "allow_all for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1714, "body_sha256": "sha256:71b31faa798359c772ce97246778027bdf168dcf3c857b20ad806ab378564efc", "child_ids": [], "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_all", "parent_id": "xcsh-docs:resources:forward_proxy_policy:reference", "path": "documentation/resources/forward_proxy_policy/properties/allow_all/index.md", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["allow_all"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/properties/allow_all/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "allow_all for xcsh_forward_proxy_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# allow_all

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/)
- allow_all

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: allow\_all, allow\_list, deny\_list, rule\_list\] Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_all/#section)
- [allow_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/#section)
- [deny_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/#section)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
allow_all = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
