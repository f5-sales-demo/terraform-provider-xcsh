---
page_title: "bot_defense_advanced_protection.mobile_only"
subcategory: "Load Balancing"
description: "bot_defense_advanced_protection.mobile_only for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1294, "body_sha256": "sha256:6b2e80df5ef0e12f96d15ede4d3ad7fce284119c244b71f955fea13279fa4dce", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:mobile_only", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:mobile_only:mobile"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:mobile_only", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection", "path": "docs/guides/resources--http_loadbalancer--properties--bot_defense_advanced_protection--mobile_only.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense_advanced_protection", "mobile_only"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense_advanced_protection/mobile_only/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense_advanced_protection.mobile_only for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_defense_advanced_protection.mobile_only

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [bot_defense_advanced_protection](resources--http_loadbalancer--properties--bot_defense_advanced_protection.md)
- bot_defense_advanced_protection.mobile_only

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Mobile. Mobile only configuration.

Upstream description:

Mobile only configuration.

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
mobile_only {
  # Configure direct properties listed below.
}
```

## Direct properties

- [mobile](resources--http_loadbalancer--properties--bot_defense_advanced_protection--mobile_only--mobile.md): complete subsection reference.

## Next pages

- [bot_defense_advanced_protection.mobile_only.mobile](resources--http_loadbalancer--properties--bot_defense_advanced_protection--mobile_only--mobile.md)
- [bot_defense_advanced_protection](resources--http_loadbalancer--properties--bot_defense_advanced_protection.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
