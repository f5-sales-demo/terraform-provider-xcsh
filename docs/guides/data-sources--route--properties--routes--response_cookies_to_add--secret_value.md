---
page_title: "routes.response_cookies_to_add.secret_value"
subcategory: ""
description: "routes.response_cookies_to_add.secret_value for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 1710, "body_sha256": "sha256:065d7ac4374be3ffc7f36045b42ecc33c508ba9447aa9bca026b852f57787e0c", "canonical_id": "xcsh-docs:data-sources:route:properties:routes:response_cookies_to_add:secret_value", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:response_cookies_to_add:secret_value:blindfold_secret_info", "xcsh-docs:data-sources:route:properties:routes:response_cookies_to_add:secret_value:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:response_cookies_to_add:secret_value", "parent_id": "xcsh-docs:data-sources:route:properties:routes:response_cookies_to_add", "path": "docs/guides/data-sources--route--properties--routes--response_cookies_to_add--secret_value.md", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "response_cookies_to_add", "secret_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/response_cookies_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.response_cookies_to_add.secret_value for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.response_cookies_to_add.secret_value

Breadcrumbs:

- [xcsh_route](../data-sources/route.md)
- [Property reference](data-sources--route--reference.md)
- [routes](data-sources--route--properties--routes.md)
- [routes.response_cookies_to_add](data-sources--route--properties--routes--response_cookies_to_add.md)
- routes.response_cookies_to_add.secret_value

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

- [blindfold_secret_info](data-sources--route--properties--routes--response_cookies_to_add--secret_value--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--route--properties--routes--response_cookies_to_add--secret_value--clear_secret_info.md): complete subsection reference.

## Next pages

- [routes.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--route--properties--routes--response_cookies_to_add--secret_value--blindfold_secret_info.md)
- [routes.response_cookies_to_add.secret_value.clear_secret_info](data-sources--route--properties--routes--response_cookies_to_add--secret_value--clear_secret_info.md)
- [routes.response_cookies_to_add](data-sources--route--properties--routes--response_cookies_to_add.md)
- [xcsh_route](../data-sources/route.md)
