---
page_title: "routes.simple_route.advanced_options.response_cookies_to_add.secret_value"
subcategory: "Load Balancing"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["routes simple route advanced options response cookies to add secret value"], "body_bytes": 3307, "body_sha256": "sha256:72d09e4a182d00dc29d148090c1e9742752c2401637eebaf36f7203ebf01ab05", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:secret_value:blindfold_secret_info", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:secret_value:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:secret_value", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add", "path": "documentation/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/response_cookies_to_add/secret_value/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0203212013131210-3112233202311012-1211021022312223-2301111332221021-3300232131111021-2003332002013101-3123100331013113-1132013121310321", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-026.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.response_cookies_to_add.secret_value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:secret_value:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.response_cookies_to_add.secret_value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:secret_value:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "response_cookies_to_add", "secret_value"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:secret_value:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--simple_route--advanced_options--response_cookies_to_add--secret_value--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.response_cookies_to_add.secret_value.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:secret_value:blindfold_secret_info", "type": "requires"}], "schema_path": ["routes", "simple_route", "advanced_options", "response_cookies_to_add", "secret_value", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:secret_value:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--simple_route--advanced_options--response_cookies_to_add--secret_value--clear_secret_info--url", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.response_cookies_to_add.secret_value.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:secret_value:clear_secret_info", "type": "requires"}], "schema_path": ["routes", "simple_route", "advanced_options", "response_cookies_to_add", "secret_value", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/response_cookies_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options.response_cookies_to_add.secret_value

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/)
- [routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/)
- [routes.simple_route.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/)
- [routes.simple_route.advanced_options.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/response_cookies_to_add/)
- routes.simple_route.advanced_options.response_cookies_to_add.secret_value

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/response_cookies_to_add/secret_value/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/response_cookies_to_add/secret_value/clear_secret_info/): complete subsection reference.

## Next pages

- [routes.simple_route.advanced_options.response_cookies_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/response_cookies_to_add/secret_value/blindfold_secret_info/)
- [routes.simple_route.advanced_options.response_cookies_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/response_cookies_to_add/secret_value/clear_secret_info/)
- [routes.simple_route.advanced_options.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/response_cookies_to_add/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
