---
page_title: "cloudfront.protected_endpoints.flow_label.authentication"
subcategory: ""
description: "Bot Defense Flow Label Authentication Category."
xcsh_docs: {"aliases": ["authentication", "cloudfront protected endpoints flow label authentication", "credential setup", "credentials"], "body_bytes": 4020, "body_sha256": "sha256:bd8713def68fbb4913fac20e0eeb3c8c97f269ef84e01259227cb4ba0ed75886", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login_mfa", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login_partner", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:logout", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:token_refresh"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label", "path": "documentation/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201", "registry_path": "docs/guides/data-sources--protected_application--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication"], "schema_version": 1, "sections": [{"aliases": ["login", "login result", "sign in"], "anchor": "section", "description": "Bot Defense Transaction Result.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication", "login"], "syntax": "attribute", "type": "object"}, {"aliases": ["login", "login mfa", "login result", "sign in"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login_mfa", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication", "login_mfa"], "syntax": "attribute", "type": "object"}, {"aliases": ["login", "login partner", "login result", "sign in"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login_partner", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication", "login_partner"], "syntax": "attribute", "type": "object"}, {"aliases": ["logout"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:logout", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication", "logout"], "syntax": "attribute", "type": "object"}, {"aliases": ["token refresh"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:token_refresh", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication", "token_refresh"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Bot Defense Flow Label Authentication Category.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.flow_label.authentication

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/)
- [cloudfront.protected_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/)
- cloudfront.protected_endpoints.flow_label.authentication

<a id="section"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Authentication Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"login\",\"login_mfa\",\"login_partner\",\"logout\",\"token_refresh\"]"
}
```

## Direct properties

- [login](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/): complete subsection reference.

- [login_mfa](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login_mfa/): complete subsection reference.

- [login_partner](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login_partner/): complete subsection reference.

- [logout](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/logout/): complete subsection reference.

- [token_refresh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/token_refresh/): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.flow_label.authentication.login](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/)
- [cloudfront.protected_endpoints.flow_label.authentication.login_mfa](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login_mfa/)
- [cloudfront.protected_endpoints.flow_label.authentication.login_partner](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login_partner/)
- [cloudfront.protected_endpoints.flow_label.authentication.logout](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/logout/)
- [cloudfront.protected_endpoints.flow_label.authentication.token_refresh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/token_refresh/)
- [cloudfront.protected_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
