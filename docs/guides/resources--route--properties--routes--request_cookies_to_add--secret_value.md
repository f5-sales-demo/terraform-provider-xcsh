---
page_title: "routes.request_cookies_to_add.secret_value"
subcategory: ""
description: "routes.request_cookies_to_add.secret_value for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 1878, "body_sha256": "sha256:330ec4251a43c642f364225ca565d7ff071c933b52748632f941b6dea3ad4067", "canonical_id": "xcsh-docs:resources:route:properties:routes:request_cookies_to_add:secret_value", "child_ids": ["xcsh-docs:resources:route:properties:routes:request_cookies_to_add:secret_value:blindfold_secret_info", "xcsh-docs:resources:route:properties:routes:request_cookies_to_add:secret_value:clear_secret_info"], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:request_cookies_to_add:secret_value", "parent_id": "xcsh-docs:resources:route:properties:routes:request_cookies_to_add", "path": "docs/guides/resources--route--properties--routes--request_cookies_to_add--secret_value.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "request_cookies_to_add", "secret_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/request_cookies_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.request_cookies_to_add.secret_value for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.request_cookies_to_add.secret_value

Breadcrumbs:

- [xcsh_route](../resources/route.md)
- [Property reference](resources--route--reference.md)
- [routes](resources--route--properties--routes.md)
- [routes.request_cookies_to_add](resources--route--properties--routes--request_cookies_to_add.md)
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

- [blindfold_secret_info](resources--route--properties--routes--request_cookies_to_add--secret_value--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--route--properties--routes--request_cookies_to_add--secret_value--clear_secret_info.md): complete subsection reference.

## Next pages

- [routes.request_cookies_to_add.secret_value.blindfold_secret_info](resources--route--properties--routes--request_cookies_to_add--secret_value--blindfold_secret_info.md)
- [routes.request_cookies_to_add.secret_value.clear_secret_info](resources--route--properties--routes--request_cookies_to_add--secret_value--clear_secret_info.md)
- [routes.request_cookies_to_add](resources--route--properties--routes--request_cookies_to_add.md)
- [xcsh_route](../resources/route.md)
