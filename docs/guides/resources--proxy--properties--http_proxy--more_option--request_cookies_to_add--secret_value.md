---
page_title: "http_proxy.more_option.request_cookies_to_add.secret_value"
subcategory: ""
description: "http_proxy.more_option.request_cookies_to_add.secret_value for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2268, "body_sha256": "sha256:440c8ee509369944ec73a0aaa00d53d2782807903271b93b686e1089891d0aaa", "canonical_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:request_cookies_to_add:secret_value", "child_ids": ["xcsh-docs:resources:proxy:properties:http_proxy:more_option:request_cookies_to_add:secret_value:blindfold_secret_info", "xcsh-docs:resources:proxy:properties:http_proxy:more_option:request_cookies_to_add:secret_value:clear_secret_info"], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:request_cookies_to_add:secret_value", "parent_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:request_cookies_to_add", "path": "docs/guides/resources--proxy--properties--http_proxy--more_option--request_cookies_to_add--secret_value.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_proxy", "more_option", "request_cookies_to_add", "secret_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/http_proxy/more_option/request_cookies_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_proxy.more_option.request_cookies_to_add.secret_value for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_proxy.more_option.request_cookies_to_add.secret_value

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [http_proxy](resources--proxy--properties--http_proxy.md)
- [http_proxy.more_option](resources--proxy--properties--http_proxy--more_option.md)
- [http_proxy.more_option.request_cookies_to_add](resources--proxy--properties--http_proxy--more_option--request_cookies_to_add.md)
- http_proxy.more_option.request_cookies_to_add.secret_value

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

- [blindfold_secret_info](resources--proxy--properties--http_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--proxy--properties--http_proxy--more_option--request_cookies_to_add--secret_value--clear_secret_info.md): complete subsection reference.

## Next pages

- [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--properties--http_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info.md)
- [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](resources--proxy--properties--http_proxy--more_option--request_cookies_to_add--secret_value--clear_secret_info.md)
- [http_proxy.more_option.request_cookies_to_add](resources--proxy--properties--http_proxy--more_option--request_cookies_to_add.md)
- [xcsh_proxy](../resources/proxy.md)
