---
page_title: "cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result"
subcategory: ""
description: "Bot Defense Transaction ResultType."
xcsh_docs: {"aliases": ["authentication", "cloudfront protected endpoints flow label authentication login transaction result", "credential setup", "credentials", "login", "login result", "sign in"], "body_bytes": 3379, "body_sha256": "sha256:9caf0728bcf6133400dde3b9d5267d287170438da14de65b92104b05165ad827", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login:transaction_result:failure_conditions", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login:transaction_result:success_conditions"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login:transaction_result", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login", "path": "documentation/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/transaction_result/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0013020233033310-3302013320031033-1303312000233302-1213302001321311-3221032102001213-3222212322133223-0002103132310330-0232001010120022", "registry_path": "docs/guides/data-sources--protected_application--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication", "login", "transaction_result"], "schema_version": 1, "sections": [{"aliases": ["failure conditions"], "anchor": "section", "description": "Failure Conditions.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login:transaction_result:failure_conditions", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication", "login", "transaction_result", "failure_conditions"], "syntax": "attribute", "type": "object"}, {"aliases": ["login success", "succeeded", "success", "success conditions", "successful"], "anchor": "section", "description": "Success Conditions.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login:transaction_result:success_conditions", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication", "login", "transaction_result", "success_conditions"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/transaction_result/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Bot Defense Transaction ResultType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/)
- [cloudfront.protected_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/)
- [cloudfront.protected_endpoints.flow_label.authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/)
- [cloudfront.protected_endpoints.flow_label.authentication.login](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/)
- cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result

<a id="section"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result Type. Bot Defense Transaction ResultType.

Upstream description:

Bot Defense Transaction ResultType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

- [failure_conditions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/transaction_result/failure_conditions/): complete subsection reference.

- [success_conditions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/transaction_result/success_conditions/): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/transaction_result/failure_conditions/)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/transaction_result/success_conditions/)
- [cloudfront.protected_endpoints.flow_label.authentication.login](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
