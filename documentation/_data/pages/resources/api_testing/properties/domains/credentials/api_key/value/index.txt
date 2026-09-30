---
page_title: "domains.credentials.api_key.value"
subcategory: ""
description: "domains.credentials.api_key.value for xcsh_api_testing."
xcsh_docs: {"aliases": [], "body_bytes": 2498, "body_sha256": "sha256:e6aff16461ba23480babdba42909dfa781285130d9804a919664110be7b73eaa", "child_ids": ["xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value:blindfold_secret_info", "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value:clear_secret_info"], "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value", "parent_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key", "path": "documentation/resources/api_testing/properties/domains/credentials/api_key/value/index.md", "provider_name": "api_testing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["domains", "credentials", "api_key", "value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/properties/domains/credentials/api_key/value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "domains.credentials.api_key.value for xcsh_api_testing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# domains.credentials.api_key.value

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/)
- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/)
- [domains.credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/)
- [domains.credentials.api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/)
- domains.credentials.api_key.value

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
value {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/value/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/value/clear_secret_info/): complete subsection reference.

## Next pages

- [domains.credentials.api_key.value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/value/blindfold_secret_info/)
- [domains.credentials.api_key.value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/value/clear_secret_info/)
- [domains.credentials.api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/)
- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
