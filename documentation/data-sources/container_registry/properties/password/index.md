---
page_title: "password"
subcategory: "Container"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["password"], "body_bytes": 1755, "body_sha256": "sha256:5608e8b92a52f3240521e52684943916bd6e9f6504793e1208368880b0a364b6", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:container_registry:properties:password:blindfold_secret_info", "xcsh-docs:data-sources:container_registry:properties:password:clear_secret_info"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:container_registry:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:container_registry:properties:password", "parent_id": "xcsh-docs:data-sources:container_registry:reference", "path": "documentation/data-sources/container_registry/properties/password/index.md", "product": "distributed-cloud", "provider_name": "container_registry", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1212000100200100-1302221200112101-1300300033113211-1202131302222323-1032221332111232-3310112110033231-3210221320333321-2320303221321233", "registry_path": "docs/guides/data-sources--container_registry--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["password"], "schema_version": 1, "sections": [{"aliases": ["password blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:container_registry:properties:password:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["password", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["password clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:container_registry:properties:password:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["password", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/container_registry/properties/password/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["container_registryCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# password

Breadcrumbs:

- [xcsh_container_registry](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/container_registry/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/container_registry/properties/)
- password

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/container_registry/properties/password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/container_registry/properties/password/clear_secret_info/): complete subsection reference.

## Next pages

- [password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/container_registry/properties/password/blindfold_secret_info/)
- [password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/container_registry/properties/password/clear_secret_info/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/container_registry/properties/)
- [xcsh_container_registry](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/container_registry/)
