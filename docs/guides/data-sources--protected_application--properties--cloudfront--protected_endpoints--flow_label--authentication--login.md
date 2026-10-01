---
page_title: "cloudfront.protected_endpoints.flow_label.authentication.login"
subcategory: ""
description: "cloudfront.protected_endpoints.flow_label.authentication.login for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 2523, "body_sha256": "sha256:d2def99e2dac8dab565872c49ec1540390e83c539707040312b6adb808921ab8", "canonical_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login:disable_transaction_result", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login:transaction_result"], "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication", "path": "docs/guides/data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication", "login"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.flow_label.authentication.login for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.flow_label.authentication.login

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md)
- [Property reference](data-sources--protected_application--reference.md)
- [cloudfront](data-sources--protected_application--properties--cloudfront.md)
- [cloudfront.protected_endpoints](data-sources--protected_application--properties--cloudfront--protected_endpoints.md)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication.md)
- cloudfront.protected_endpoints.flow_label.authentication.login

<a id="section"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Upstream description:

Bot Defense Transaction Result.

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

- [disable_transaction_result](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login--disable_transaction_result.md): complete subsection reference.

- [transaction_result](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result.md): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login--disable_transaction_result.md)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result.md)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication.md)
- [xcsh_protected_application](../data-sources/protected_application.md)
