---
page_title: "rules.client_country"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["rules client country"], "body_bytes": 1247, "body_sha256": "sha256:8841b359ef21b183ef839f80d3077c63deda0252dbe7af5fe46e053ffe4f0f32", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "id": "xcsh-docs:resources:user_identification:properties:rules:client_country", "parent_id": "xcsh-docs:resources:user_identification:properties:rules", "path": "documentation/resources/user_identification/properties/rules/client_country/index.md", "product": "distributed-cloud", "provider_name": "user_identification", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1212201131133223-2002120013132312-1311200302110022-2311131133021132-3301130022112302-2221032212230301-0302231231213312-2012011232100331", "registry_path": "docs/guides/resources--user_identification--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "client_country"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/properties/rules/client_country/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["user_identificationCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.client_country

Breadcrumbs:

- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/rules/)
- rules.client_country

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
client_country = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/rules/)
- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/)
