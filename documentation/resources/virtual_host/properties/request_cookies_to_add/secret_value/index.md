---
page_title: "request_cookies_to_add.secret_value"
subcategory: ""
description: "request_cookies_to_add.secret_value for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 2365, "body_sha256": "sha256:d80c7954d7ac0b279af9a13901dcc8332bdfcc26db16e7a5a16c6829a654edae", "child_ids": ["xcsh-docs:resources:virtual_host:properties:request_cookies_to_add:secret_value:blindfold_secret_info", "xcsh-docs:resources:virtual_host:properties:request_cookies_to_add:secret_value:clear_secret_info"], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:request_cookies_to_add:secret_value", "parent_id": "xcsh-docs:resources:virtual_host:properties:request_cookies_to_add", "path": "documentation/resources/virtual_host/properties/request_cookies_to_add/secret_value/index.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["request_cookies_to_add", "secret_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/request_cookies_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "request_cookies_to_add.secret_value for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# request_cookies_to_add.secret_value

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/)
- request_cookies_to_add.secret_value

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/secret_value/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/secret_value/clear_secret_info/): complete subsection reference.

## Next pages

- [request_cookies_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/secret_value/blindfold_secret_info/)
- [request_cookies_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/secret_value/clear_secret_info/)
- [request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
