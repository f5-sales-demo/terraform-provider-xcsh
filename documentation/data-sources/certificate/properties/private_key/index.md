---
page_title: "private_key"
subcategory: "Security"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["private key"], "body_bytes": 1148, "body_sha256": "sha256:104c93b696d571bc26457088b9a3aabaa6f57727fea61b6b964c02b6896bcaa4", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:certificate:properties:private_key:blindfold_secret_info", "xcsh-docs:data-sources:certificate:properties:private_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:certificate:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certificate:properties:private_key", "parent_id": "xcsh-docs:data-sources:certificate:reference", "path": "documentation/data-sources/certificate/properties/private_key/index.md", "product": "distributed-cloud", "provider_name": "certificate", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0303332301111032-1121312202010230-2010123001002110-1000200100332031-2022332113132133-1001003300221032-0013103033011303-1322002121022121", "registry_path": "docs/guides/data-sources--certificate--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["private_key"], "schema_version": 1, "sections": [{"aliases": ["private key blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:certificate:properties:private_key:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["private_key", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["private key clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:certificate:properties:private_key:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["private_key", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certificate/properties/private_key/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["certificateCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# private_key

Breadcrumbs:

- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/)
- private_key

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/private_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/private_key/clear_secret_info/): complete subsection reference.
