---
page_title: "cloudfront.protected_endpoints.flow_label.authentication.token_refresh"
subcategory: ""
description: "cloudfront.protected_endpoints.flow_label.authentication.token_refresh for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 1670, "body_sha256": "sha256:3a0e22b4653876f49db69f77c5b5e4bcc80bb80fc763f60a4cb182d8c6bda882", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:token_refresh", "child_ids": [], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:token_refresh", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication", "path": "docs/guides/resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--token_refresh.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication", "token_refresh"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/token_refresh/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.flow_label.authentication.token_refresh for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.flow_label.authentication.token_refresh

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- [cloudfront](resources--protected_application--properties--cloudfront.md)
- [cloudfront.protected_endpoints](resources--protected_application--properties--cloudfront--protected_endpoints.md)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication.md)
- cloudfront.protected_endpoints.flow_label.authentication.token_refresh

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for token refresh.

Upstream description:

This can be used for messages where no values are needed.

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
token_refresh = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication.md)
- [xcsh_protected_application](../resources/protected_application.md)
