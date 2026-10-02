---
page_title: "disable_ha"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable ha"], "body_bytes": 1534, "body_sha256": "sha256:4b265e6f859d17bd1be41b8bcfc0eb99c1c57e2e2273fb07708eaca901a795b8", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:disable_ha", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/disable_ha/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3231002302323332-2013010223311332-0021120301210112-2232331120122110-0301210303200310-2002220333321202-2022103131122023-3101003100301002", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_ha"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/disable_ha/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_ha

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- disable_ha

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_ha, enable\_ha; Default: disable\_ha\] Enable this option

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

- [disable_ha](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/disable_ha/#section)
- [enable_ha](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/enable_ha/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_ha = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
