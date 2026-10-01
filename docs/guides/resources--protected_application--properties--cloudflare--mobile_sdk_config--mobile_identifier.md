---
page_title: "cloudflare.mobile_sdk_config.mobile_identifier"
subcategory: ""
description: "cloudflare.mobile_sdk_config.mobile_identifier for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 1551, "body_sha256": "sha256:8052b223a665af88e2c625eaed375d36da6ddda973dfc6d10a959e09e6a29bad", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudflare:mobile_sdk_config:mobile_identifier", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudflare:mobile_sdk_config:mobile_identifier:headers"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:mobile_sdk_config:mobile_identifier", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare:mobile_sdk_config", "path": "docs/guides/resources--protected_application--properties--cloudflare--mobile_sdk_config--mobile_identifier.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudflare", "mobile_sdk_config", "mobile_identifier"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/mobile_sdk_config/mobile_identifier/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudflare.mobile_sdk_config.mobile_identifier for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.mobile_sdk_config.mobile_identifier

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- [cloudflare](resources--protected_application--properties--cloudflare.md)
- [cloudflare.mobile_sdk_config](resources--protected_application--properties--cloudflare--mobile_sdk_config.md)
- cloudflare.mobile_sdk_config.mobile_identifier

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Mobile Traffic Identifier. Mobile traffic identifier type.

Upstream description:

Mobile traffic identifier type.

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
mobile_identifier {
  # Configure direct properties listed below.
}
```

## Direct properties

- [headers](resources--protected_application--properties--cloudflare--mobile_sdk_config--mobile_identifier--headers.md): complete subsection reference.

## Next pages

- [cloudflare.mobile_sdk_config.mobile_identifier.headers](resources--protected_application--properties--cloudflare--mobile_sdk_config--mobile_identifier--headers.md)
- [cloudflare.mobile_sdk_config](resources--protected_application--properties--cloudflare--mobile_sdk_config.md)
- [xcsh_protected_application](../resources/protected_application.md)
