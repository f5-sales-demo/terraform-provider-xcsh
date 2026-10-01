---
page_title: "routes.simple_route.advanced_options.response_headers_to_add.secret_value"
subcategory: "Load Balancing"
description: "routes.simple_route.advanced_options.response_headers_to_add.secret_value for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2722, "body_sha256": "sha256:52f562b384f8b586d1bc34e6cbb1aa7517d1a645abe2dea559878a9514e2ec7e", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_headers_to_add:secret_value", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_headers_to_add:secret_value:blindfold_secret_info", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_headers_to_add:secret_value:clear_secret_info"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_headers_to_add:secret_value", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_headers_to_add", "path": "docs/guides/resources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_headers_to_add--secret_value.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "response_headers_to_add", "secret_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/response_headers_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.simple_route.advanced_options.response_headers_to_add.secret_value for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options.response_headers_to_add.secret_value

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [routes](resources--http_loadbalancer--properties--routes.md)
- [routes.simple_route](resources--http_loadbalancer--properties--routes--simple_route.md)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- [routes.simple_route.advanced_options.response_headers_to_add](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_headers_to_add.md)
- routes.simple_route.advanced_options.response_headers_to_add.secret_value

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

- [blindfold_secret_info](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_headers_to_add--secret_value--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_headers_to_add--secret_value--clear_secret_info.md): complete subsection reference.

## Next pages

- [routes.simple_route.advanced_options.response_headers_to_add.secret_value.blindfold_secret_info](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_headers_to_add--secret_value--blindfold_secret_info.md)
- [routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_secret_info](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_headers_to_add--secret_value--clear_secret_info.md)
- [routes.simple_route.advanced_options.response_headers_to_add](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_headers_to_add.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
