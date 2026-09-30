---
page_title: "access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info"
subcategory: ""
description: "access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info for xcsh_lma_region."
xcsh_docs: {"aliases": [], "body_bytes": 2116, "body_sha256": "sha256:2fce0023be8f499b4f50212d1910b329dd6078c43f1c60d5bc3af2801d6aab27", "child_ids": [], "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials:secret_access_key:clear_secret_info", "parent_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials:secret_access_key", "path": "documentation/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/secret_access_key/clear_secret_info/index.md", "provider_name": "lma_region", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["access_logs_s3_params", "aws_credentials", "secret_access_key", "clear_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/secret_access_key/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info for xcsh_lma_region.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info

Breadcrumbs:

- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/)
- [access_logs_s3_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/)
- [access_logs_s3_params.aws_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/)
- [access_logs_s3_params.aws_credentials.secret_access_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/secret_access_key/)
- access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info

<a id="section"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

## Direct properties

<a id="schema-access_logs_s3_params--aws_credentials--secret_access_key--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-access_logs_s3_params--aws_credentials--secret_access_key--clear_secret_info--url"></a>

### url property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

## Next pages

- [access_logs_s3_params.aws_credentials.secret_access_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/secret_access_key/)
- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
