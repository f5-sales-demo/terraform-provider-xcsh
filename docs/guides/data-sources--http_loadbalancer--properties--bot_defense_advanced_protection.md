---
page_title: "bot_defense_advanced_protection"
subcategory: "Load Balancing"
description: "bot_defense_advanced_protection for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1660, "body_sha256": "sha256:8dc0420ab0319bd9e9fe7865ac8178f59fe971083958e467ee06ab4f84279fe2", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:mobile_only", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "docs/guides/data-sources--http_loadbalancer--properties--bot_defense_advanced_protection.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense_advanced_protection"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense_advanced_protection for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_defense_advanced_protection

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- bot_defense_advanced_protection

<a id="section"></a>

Type: `"single"`. Computed.

Bot Defense Advanced Protection - replaces BotDefenseAdvancedType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-client_type_choice": "[\"both_web_and_mobile\",\"mobile_only\",\"web_only\"]"
}
```

## Direct properties

- [both_web_and_mobile](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile.md): complete subsection reference.

- [mobile_only](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--mobile_only.md): complete subsection reference.

- [web_only](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only.md): complete subsection reference.

## Next pages

- [bot_defense_advanced_protection.both_web_and_mobile](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile.md)
- [bot_defense_advanced_protection.mobile_only](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--mobile_only.md)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
