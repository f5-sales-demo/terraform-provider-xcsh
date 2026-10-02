---
page_title: "routes.response_cookies_to_add.secret_value"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["routes response cookies to add secret value"], "body_bytes": 2487, "body_sha256": "sha256:595290fbd46392ff82437d26b1e46b8b7f489c2fe3f507c346ca977bbfcd3d21", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:route:properties:routes:response_cookies_to_add:secret_value:blindfold_secret_info", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:secret_value:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:secret_value", "parent_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add", "path": "documentation/resources/route/properties/routes/response_cookies_to_add/secret_value/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0103203110121200-2121220322032313-0233301211311330-2311322312210122-3313130023102202-2101033321222231-0202232321331000-3331200020123000", "registry_path": "docs/guides/resources--route--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add.secret_value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:secret_value:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add.secret_value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:secret_value:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "response_cookies_to_add", "secret_value"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:secret_value:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--response_cookies_to_add--secret_value--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add.secret_value.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:secret_value:blindfold_secret_info", "type": "requires"}], "schema_path": ["routes", "response_cookies_to_add", "secret_value", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:secret_value:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--response_cookies_to_add--secret_value--clear_secret_info--url", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add.secret_value.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:secret_value:clear_secret_info", "type": "requires"}], "schema_path": ["routes", "response_cookies_to_add", "secret_value", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/response_cookies_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["routeCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.response_cookies_to_add.secret_value

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/)
- routes.response_cookies_to_add.secret_value

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/secret_value/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/secret_value/clear_secret_info/): complete subsection reference.

## Next pages

- [routes.response_cookies_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/secret_value/blindfold_secret_info/)
- [routes.response_cookies_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/secret_value/clear_secret_info/)
- [routes.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
