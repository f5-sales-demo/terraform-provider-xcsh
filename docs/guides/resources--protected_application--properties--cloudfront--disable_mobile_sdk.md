---
page_title: "cloudfront.disable_mobile_sdk"
subcategory: ""
description: "cloudfront.disable_mobile_sdk for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 1046, "body_sha256": "sha256:37f0695510a109b95c21d4c875a40304776608b17ffe3e9a02e9e9c28db6796e", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudfront:disable_mobile_sdk", "child_ids": [], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:disable_mobile_sdk", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront", "path": "docs/guides/resources--protected_application--properties--cloudfront--disable_mobile_sdk.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudfront", "disable_mobile_sdk"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/disable_mobile_sdk/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.disable_mobile_sdk for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.disable_mobile_sdk

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- [cloudfront](resources--protected_application--properties--cloudfront.md)
- cloudfront.disable_mobile_sdk

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
disable_mobile_sdk = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [cloudfront](resources--protected_application--properties--cloudfront.md)
- [xcsh_protected_application](../resources/protected_application.md)
