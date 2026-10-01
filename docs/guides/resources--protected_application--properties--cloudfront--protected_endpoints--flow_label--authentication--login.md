---
page_title: "cloudfront.protected_endpoints.flow_label.authentication.login"
subcategory: ""
description: "cloudfront.protected_endpoints.flow_label.authentication.login for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 2795, "body_sha256": "sha256:09c23f958ee0e67595c73702dc9391dc17dac2a23d98d3036635136331141708", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login:disable_transaction_result", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login:transaction_result"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication", "path": "docs/guides/resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication", "login"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.flow_label.authentication.login for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.flow_label.authentication.login

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- [cloudfront](resources--protected_application--properties--cloudfront.md)
- [cloudfront.protected_endpoints](resources--protected_application--properties--cloudfront--protected_endpoints.md)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication.md)
- cloudfront.protected_endpoints.flow_label.authentication.login

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Upstream description:

Bot Defense Transaction Result.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_transaction_result",
    "transaction_result")}
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
  "x-ves-oneof-field-transaction_result_choice": "[\"disable_transaction_result\",\"transaction_result\"]"
}
```

Terraform syntax:

```terraform
login {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_transaction_result](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login--disable_transaction_result.md): complete subsection reference.

- [transaction_result](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result.md): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login--disable_transaction_result.md)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result.md)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication.md)
- [xcsh_protected_application](../resources/protected_application.md)
