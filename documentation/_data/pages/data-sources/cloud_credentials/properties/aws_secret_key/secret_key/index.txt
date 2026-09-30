---
page_title: "aws_secret_key.secret_key"
subcategory: "Infrastructure"
description: "aws_secret_key.secret_key for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 1927, "body_sha256": "sha256:390ea50a9675ded6fdd3b36f945226d7538e09efaae28e95f784a8bf6cded4b5", "child_ids": ["xcsh-docs:data-sources:cloud_credentials:properties:aws_secret_key:secret_key:blindfold_secret_info", "xcsh-docs:data-sources:cloud_credentials:properties:aws_secret_key:secret_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_credentials:properties:aws_secret_key:secret_key", "parent_id": "xcsh-docs:data-sources:cloud_credentials:properties:aws_secret_key", "path": "documentation/data-sources/cloud_credentials/properties/aws_secret_key/secret_key/index.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["aws_secret_key", "secret_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_credentials/properties/aws_secret_key/secret_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_secret_key.secret_key for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# aws_secret_key.secret_key

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/)
- [aws_secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/aws_secret_key/)
- aws_secret_key.secret_key

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/aws_secret_key/secret_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/aws_secret_key/secret_key/clear_secret_info/): complete subsection reference.

## Next pages

- [aws_secret_key.secret_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/aws_secret_key/secret_key/blindfold_secret_info/)
- [aws_secret_key.secret_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/aws_secret_key/secret_key/clear_secret_info/)
- [aws_secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/aws_secret_key/)
- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/)
