---
page_title: "allow_list.default_action_deny"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["allow list default action deny"], "body_bytes": 1257, "body_sha256": "sha256:07a16ae5a52043b20c4a7762769a490748f59340224f259c0ca24fac6a6fe37f", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:allow_list:default_action_deny", "parent_id": "xcsh-docs:resources:service_policy:properties:allow_list", "path": "documentation/resources/service_policy/properties/allow_list/default_action_deny/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1223032033013321-1311120122122222-3011310221133212-2310230102322211-1300020320320303-2100203112101221-1132023332200223-2203311203030230", "registry_path": "docs/guides/resources--service_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["allow_list", "default_action_deny"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/allow_list/default_action_deny/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allow_list.default_action_deny

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- [allow_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/)
- allow_list.default_action_deny

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
default_action_deny = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [allow_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
