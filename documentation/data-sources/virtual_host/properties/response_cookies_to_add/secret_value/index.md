---
page_title: "response_cookies_to_add.secret_value"
subcategory: ""
description: "response_cookies_to_add.secret_value for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1996, "body_sha256": "sha256:c7ef7ea849e85f0a714833050ada1850d794e96419ed9541351f1d324d7922f1", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:response_cookies_to_add:secret_value:blindfold_secret_info", "xcsh-docs:data-sources:virtual_host:properties:response_cookies_to_add:secret_value:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:response_cookies_to_add:secret_value", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:response_cookies_to_add", "path": "documentation/data-sources/virtual_host/properties/response_cookies_to_add/secret_value/index.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["response_cookies_to_add", "secret_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/response_cookies_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "response_cookies_to_add.secret_value for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# response_cookies_to_add.secret_value

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/response_cookies_to_add/)
- response_cookies_to_add.secret_value

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/response_cookies_to_add/secret_value/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/response_cookies_to_add/secret_value/clear_secret_info/): complete subsection reference.

## Next pages

- [response_cookies_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/response_cookies_to_add/secret_value/blindfold_secret_info/)
- [response_cookies_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/response_cookies_to_add/secret_value/clear_secret_info/)
- [response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/response_cookies_to_add/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
