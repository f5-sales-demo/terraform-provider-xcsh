---
page_title: "cloudfront.protected_endpoints.mobile_client"
subcategory: ""
description: "cloudfront.protected_endpoints.mobile_client for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 2078, "body_sha256": "sha256:dec0317e048bf1531a2733011f7257ddc87eb0abe8259498cbb80bd82cad3d48", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:block", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints", "path": "docs/guides/resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "mobile_client"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.mobile_client for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.mobile_client

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- [cloudfront](resources--protected_application--properties--cloudfront.md)
- [cloudfront.protected_endpoints](resources--protected_application--properties--cloudfront--protected_endpoints.md)
- cloudfront.protected_endpoints.mobile_client

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Mobile Client. Mobile client configuration OPTIONS.

Upstream description:

Mobile client configuration OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "continue")}
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
  "x-ves-oneof-field-mitigation": "[\"block\",\"continue\"]"
}
```

Terraform syntax:

```terraform
mobile_client {
  # Configure direct properties listed below.
}
```

## Direct properties

- [block](resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client--block.md): complete subsection reference.

- [continue](resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client--continue.md): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.mobile_client.block](resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client--block.md)
- [cloudfront.protected_endpoints.mobile_client.continue](resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client--continue.md)
- [cloudfront.protected_endpoints](resources--protected_application--properties--cloudfront--protected_endpoints.md)
- [xcsh_protected_application](../resources/protected_application.md)
