---
page_title: "disable_advanced_delivery"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable advanced delivery"], "body_bytes": 1723, "body_sha256": "sha256:9a3d26891c30749a7276ad619f32f1f69e9dbc72ff4205d0417c171a1c9d3457", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:disable_advanced_delivery", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/disable_advanced_delivery/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1112203010312210-0132130311102202-0022222202032000-2131322332010323-3210102110223000-3200001233331101-1223311301211212-3200100111113210", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_advanced_delivery"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/disable_advanced_delivery/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_advanced_delivery

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- disable_advanced_delivery

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_advanced\_delivery, enable\_advanced\_delivery; Default:
disable\_advanced\_delivery\] Configuration parameter for disable advanced delivery.

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

- [disable_advanced_delivery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/disable_advanced_delivery/#section)
- [enable_advanced_delivery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/enable_advanced_delivery/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_advanced_delivery = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
