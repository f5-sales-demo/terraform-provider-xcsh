---
page_title: "cloudflare.mobile_sdk_config"
subcategory: ""
description: "cloudflare.mobile_sdk_config for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 1237, "body_sha256": "sha256:d012141269b65f1fb2b27764bd5edac140314e66053d5c886d4054e07dc1535e", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudflare:mobile_sdk_config", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudflare:mobile_sdk_config:mobile_identifier"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:mobile_sdk_config", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare", "path": "docs/guides/resources--protected_application--properties--cloudflare--mobile_sdk_config.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudflare", "mobile_sdk_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/mobile_sdk_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudflare.mobile_sdk_config for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# cloudflare.mobile_sdk_config

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- [cloudflare](resources--protected_application--properties--cloudflare.md)
- cloudflare.mobile_sdk_config

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

- [mobile_identifier](resources--protected_application--properties--cloudflare--mobile_sdk_config--mobile_identifier.md): complete subsection reference.

## Next pages

- [cloudflare.mobile_sdk_config.mobile_identifier](resources--protected_application--properties--cloudflare--mobile_sdk_config--mobile_identifier.md)
- [cloudflare](resources--protected_application--properties--cloudflare.md)
- [xcsh_protected_application](../resources/protected_application.md)
