---
page_title: "cloudfront.mobile_sdk_config"
subcategory: ""
description: "cloudfront.mobile_sdk_config for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 1336, "body_sha256": "sha256:04b860c0c88ca565a7bf41b45432733a505e7463c9fa64078bd29ec7c71f552d", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront", "path": "docs/guides/resources--protected_application--properties--cloudfront--mobile_sdk_config.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudfront", "mobile_sdk_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/mobile_sdk_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.mobile_sdk_config for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.mobile_sdk_config

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- [cloudfront](resources--protected_application--properties--cloudfront.md)
- cloudfront.mobile_sdk_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Mobile SDK Configuration. Mobile SDK configuration.

Upstream description:

Mobile SDK configuration.

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
mobile_sdk_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [mobile_identifier](resources--protected_application--properties--cloudfront--mobile_sdk_config--mobile_identifier.md): complete subsection reference.

## Next pages

- [cloudfront.mobile_sdk_config.mobile_identifier](resources--protected_application--properties--cloudfront--mobile_sdk_config--mobile_identifier.md)
- [cloudfront](resources--protected_application--properties--cloudfront.md)
- [xcsh_protected_application](../resources/protected_application.md)
