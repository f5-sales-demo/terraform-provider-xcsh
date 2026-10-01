---
page_title: "cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result"
subcategory: ""
description: "cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 3467, "body_sha256": "sha256:9c81280c2a5799594704d7652c146d15f60fbca9e00a16563f2c5bbeb97a1d3e", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login:transaction_result:failure_conditions", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login:transaction_result:success_conditions"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login:transaction_result", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login", "path": "documentation/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/transaction_result/index.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication", "login", "transaction_result"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/transaction_result/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/)
- [cloudfront.protected_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/)
- [cloudfront.protected_endpoints.flow_label.authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/)
- [cloudfront.protected_endpoints.flow_label.authentication.login](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/)
- cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
transaction_result {
  # Configure direct properties listed below.
}
```

## Direct properties

- [failure_conditions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/transaction_result/failure_conditions/): complete subsection reference.

- [success_conditions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/transaction_result/success_conditions/): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/transaction_result/failure_conditions/)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/transaction_result/success_conditions/)
- [cloudfront.protected_endpoints.flow_label.authentication.login](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
