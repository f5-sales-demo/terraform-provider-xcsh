---
page_title: "allow_all_requests"
subcategory: "Security"
description: "allow_all_requests for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1901, "body_sha256": "sha256:c6fcb1087eb8339066665ca3e42bbbbeec2ad5ad0c2c22ea3ae73720495e799c", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:allow_all_requests", "parent_id": "xcsh-docs:resources:service_policy:reference", "path": "documentation/resources/service_policy/properties/allow_all_requests/index.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["allow_all_requests"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/allow_all_requests/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "allow_all_requests for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# allow_all_requests

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- allow_all_requests

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: allow\_all\_requests, allow\_list, deny\_all\_requests, deny\_list, rule\_list\]
Configuration parameter for allow all requests.

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

- [allow_all_requests](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_all_requests/#section)
- [allow_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/#section)
- [deny_all_requests](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_all_requests/#section)
- [deny_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/#section)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
allow_all_requests = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
