---
page_title: "aws_provider.aws_secret_key.secret_key"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["aws provider aws secret key secret key"], "body_bytes": 2324, "body_sha256": "sha256:6d79d9e3ad36797abdd75b46934d0d6c24e6ce2aaa6e0f6886ea9eb7d7514c30", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key:blindfold_secret_info", "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key", "parent_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key", "path": "documentation/data-sources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/index.md", "product": "distributed-cloud", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-2201123032031030-3313132200301120-3003330322021010-0313203203031020-1133033322002321-3133033310333203-1103133012332130-0212002320112023", "registry_path": "docs/guides/data-sources--cloud_user_account--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_provider", "aws_secret_key", "secret_key"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_provider", "aws_secret_key", "secret_key", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_provider", "aws_secret_key", "secret_key", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_secret_key.secret_key

Breadcrumbs:

- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/aws_provider/)
- [aws_provider.aws_secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/aws_provider/aws_secret_key/)
- aws_provider.aws_secret_key.secret_key

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/clear_secret_info/): complete subsection reference.

## Next pages

- [aws_provider.aws_secret_key.secret_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/blindfold_secret_info/)
- [aws_provider.aws_secret_key.secret_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/clear_secret_info/)
- [aws_provider.aws_secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/aws_provider/aws_secret_key/)
- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/)
