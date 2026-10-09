---
page_title: "cloudfront.protected_endpoints.flow_label.authentication.login"
subcategory: ""
description: "Bot Defense Transaction Result."
xcsh_docs: {"aliases": ["cloudfront protected endpoints flow label authentication login", "login", "login result", "sign in"], "body_bytes": 2130, "body_sha256": "sha256:db156cdc8d0886845eab66270b7f7612195ff2574c75218740d61c35aec6cadd", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login:disable_transaction_result", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login:transaction_result"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication", "path": "documentation/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3002133111222112-1231032002112122-3211131323230030-3013222033101021-0010013323132210-0231021130112001-2123321110123031-3323211011003222", "registry_path": "docs/guides/data-sources--protected_application--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication", "login"], "schema_version": 1, "sections": [{"aliases": ["cloudfront protected endpoints flow label authentication login disable transaction result", "login", "login result", "sign in"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login:disable_transaction_result", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication", "login", "disable_transaction_result"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront protected endpoints flow label authentication login transaction result", "login", "login result", "sign in"], "anchor": "section", "description": "Bot Defense Transaction ResultType.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login:transaction_result", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication", "login", "transaction_result"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Bot Defense Transaction Result.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.flow_label.authentication.login

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/)
- [cloudfront.protected_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/)
- [cloudfront.protected_endpoints.flow_label.authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/)
- cloudfront.protected_endpoints.flow_label.authentication.login

<a id="section"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-transaction_result_choice": "[\"disable_transaction_result\",\"transaction_result\"]"
}
```

## Direct properties

- [disable_transaction_result](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/disable_transaction_result/): complete subsection reference.

- [transaction_result](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/transaction_result/): complete subsection reference.
