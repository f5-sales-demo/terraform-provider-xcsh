---
page_title: "api_token"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["api token", "authentication", "credential setup", "credentials"], "body_bytes": 1683, "body_sha256": "sha256:15509d84d73cb30010ff386d0cf201e420efd2f9e75f1adf7b3eee0c26162204", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:cminstance:properties:api_token:blindfold_secret_info", "xcsh-docs:data-sources:cminstance:properties:api_token:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cminstance:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cminstance:properties:api_token", "parent_id": "xcsh-docs:data-sources:cminstance:reference", "path": "documentation/data-sources/cminstance/properties/api_token/index.md", "product": "distributed-cloud", "provider_name": "cminstance", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2122011332310321-2012020322322212-2103013120112311-1301000111201000-3210011003323231-1202212322022311-0001312112333103-1011331210231031", "registry_path": "docs/guides/data-sources--cminstance--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_token"], "schema_version": 1, "sections": [{"aliases": ["api token blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:cminstance:properties:api_token:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_token", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["api token clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:cminstance:properties:api_token:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_token", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cminstance/properties/api_token/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["cminstanceCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_token

Breadcrumbs:

- [xcsh_cminstance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/)
- api_token

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/api_token/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/api_token/clear_secret_info/): complete subsection reference.

## Next pages

- [api_token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/api_token/blindfold_secret_info/)
- [api_token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/api_token/clear_secret_info/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/)
- [xcsh_cminstance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/)
