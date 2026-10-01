---
page_title: "cloudfront.protected_endpoints.web_mobile_client"
subcategory: ""
description: "cloudfront.protected_endpoints.web_mobile_client for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 3482, "body_sha256": "sha256:6abe55e5304942fc0ac804b0d9b12e7e4ce19a02b2d78cefeb304c94f66580b8", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:block_mobile", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:block_web", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:continue_mobile", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:continue_web", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:redirect_web"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints", "path": "docs/guides/resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "web_mobile_client"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.web_mobile_client for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.web_mobile_client

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- [cloudfront](resources--protected_application--properties--cloudfront.md)
- [cloudfront.protected_endpoints](resources--protected_application--properties--cloudfront--protected_endpoints.md)
- cloudfront.protected_endpoints.web_mobile_client

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Web and Mobile client configuration OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block_mobile",
    "continue_mobile"),
  validators.ConflictingObjectAttributes("block_web",
    "continue_web"),
  validators.ConflictingObjectAttributes("block_web",
    "redirect_web"),
  validators.ConflictingObjectAttributes("continue_web",
    "redirect_web")}
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
  "x-ves-oneof-field-mobile_mitigation": "[\"block_mobile\",\"continue_mobile\"]",
  "x-ves-oneof-field-web_mitigation": "[\"block_web\",\"continue_web\",\"redirect_web\"]"
}
```

Terraform syntax:

```terraform
web_mobile_client {
  # Configure direct properties listed below.
}
```

## Direct properties

- [block_mobile](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--block_mobile.md): complete subsection reference.

- [block_web](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--block_web.md): complete subsection reference.

- [continue_mobile](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--continue_mobile.md): complete subsection reference.

- [continue_web](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--continue_web.md): complete subsection reference.

- [redirect_web](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--redirect_web.md): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.web_mobile_client.block_mobile](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--block_mobile.md)
- [cloudfront.protected_endpoints.web_mobile_client.block_web](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--block_web.md)
- [cloudfront.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--continue_mobile.md)
- [cloudfront.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--continue_web.md)
- [cloudfront.protected_endpoints.web_mobile_client.redirect_web](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--redirect_web.md)
- [cloudfront.protected_endpoints](resources--protected_application--properties--cloudfront--protected_endpoints.md)
- [xcsh_protected_application](../resources/protected_application.md)
