---
page_title: "cloudfront.protected_endpoints.flow_label.authentication"
subcategory: ""
description: "cloudfront.protected_endpoints.flow_label.authentication for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 3196, "body_sha256": "sha256:628c2d3b655846333a318e888b8dbf6cf4e4343b6d43ee0a73e5d0e02c874e43", "canonical_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login_mfa", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login_partner", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:logout", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:token_refresh"], "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label", "path": "docs/guides/data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.flow_label.authentication for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.flow_label.authentication

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md)
- [Property reference](data-sources--protected_application--reference.md)
- [cloudfront](data-sources--protected_application--properties--cloudfront.md)
- [cloudfront.protected_endpoints](data-sources--protected_application--properties--cloudfront--protected_endpoints.md)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md)
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

- [login](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login.md): complete subsection reference.

- [login_mfa](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login_mfa.md): complete subsection reference.

- [login_partner](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login_partner.md): complete subsection reference.

- [logout](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--logout.md): complete subsection reference.

- [token_refresh](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--token_refresh.md): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login.md)
- [cloudfront.protected_endpoints.flow_label.authentication.login_mfa](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login_mfa.md)
- [cloudfront.protected_endpoints.flow_label.authentication.login_partner](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login_partner.md)
- [cloudfront.protected_endpoints.flow_label.authentication.logout](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--logout.md)
- [cloudfront.protected_endpoints.flow_label.authentication.token_refresh](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--token_refresh.md)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md)
- [xcsh_protected_application](../data-sources/protected_application.md)
