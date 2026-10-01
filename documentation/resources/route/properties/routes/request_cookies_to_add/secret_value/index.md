---
page_title: "routes.request_cookies_to_add.secret_value"
subcategory: ""
description: "routes.request_cookies_to_add.secret_value for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 2475, "body_sha256": "sha256:796727cce2603d1e105eb3671193821a112c0806cc0b3dc7e84ebef93273128e", "child_ids": ["xcsh-docs:resources:route:properties:routes:request_cookies_to_add:secret_value:blindfold_secret_info", "xcsh-docs:resources:route:properties:routes:request_cookies_to_add:secret_value:clear_secret_info"], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:request_cookies_to_add:secret_value", "parent_id": "xcsh-docs:resources:route:properties:routes:request_cookies_to_add", "path": "documentation/resources/route/properties/routes/request_cookies_to_add/secret_value/index.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["routes", "request_cookies_to_add", "secret_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/request_cookies_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.request_cookies_to_add.secret_value for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.request_cookies_to_add.secret_value

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/)
- routes.request_cookies_to_add.secret_value

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/secret_value/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/secret_value/clear_secret_info/): complete subsection reference.

## Next pages

- [routes.request_cookies_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/secret_value/blindfold_secret_info/)
- [routes.request_cookies_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/secret_value/clear_secret_info/)
- [routes.request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
