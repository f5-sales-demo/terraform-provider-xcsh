---
page_title: "aws.byoc.connections.auth_key"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["aws byoc connections auth key"], "body_bytes": 2225, "body_sha256": "sha256:990a049f8183f536b774cad3ef2cf4e7b369c9efbdd88f7c22ea1a7104e886ce", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_link:properties:aws:byoc:connections:auth_key:blindfold_secret_info", "xcsh-docs:data-sources:cloud_link:properties:aws:byoc:connections:auth_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_link:properties:aws:byoc:connections:auth_key", "parent_id": "xcsh-docs:data-sources:cloud_link:properties:aws:byoc:connections", "path": "documentation/data-sources/cloud_link/properties/aws/byoc/connections/auth_key/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3013332200201212-2221123020103131-0033013301231101-2031233110103331-1322031011032202-0131121303312003-3332112101310312-1022031211032101", "registry_path": "docs/guides/data-sources--cloud_link--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws", "byoc", "connections", "auth_key"], "schema_version": 1, "sections": [{"aliases": ["aws byoc connections auth key blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:cloud_link:properties:aws:byoc:connections:auth_key:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws", "byoc", "connections", "auth_key", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws byoc connections auth key clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:cloud_link:properties:aws:byoc:connections:auth_key:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws", "byoc", "connections", "auth_key", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_link/properties/aws/byoc/connections/auth_key/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws.byoc.connections.auth_key

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/)
- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/)
- [aws.byoc](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/byoc/)
- [aws.byoc.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/byoc/connections/)
- aws.byoc.connections.auth_key

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/byoc/connections/auth_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/byoc/connections/auth_key/clear_secret_info/): complete subsection reference.

## Next pages

- [aws.byoc.connections.auth_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/byoc/connections/auth_key/blindfold_secret_info/)
- [aws.byoc.connections.auth_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/byoc/connections/auth_key/clear_secret_info/)
- [aws.byoc.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/byoc/connections/)
- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/)
