---
page_title: "enable_challenge"
subcategory: "Load Balancing"
description: "enable_challenge for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2982, "body_sha256": "sha256:4778285989032e5f2f0a69792faadf42f51922dbc3c332498041ff52517fc4a5", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_challenge", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_challenge:captcha_challenge_parameters", "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_challenge:default_captcha_challenge_parameters", "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_challenge:default_js_challenge_parameters", "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_challenge:default_mitigation_settings", "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_challenge:js_challenge_parameters", "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_challenge:malicious_user_mitigation"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_challenge", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--enable_challenge.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_challenge"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/enable_challenge/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_challenge for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_challenge

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- enable_challenge

<a id="section"></a>

Type: `"single"`. Computed.

Configure auto mitigation i.e risk based challenges for malicious users.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]"
}
```

## Direct properties

- [captcha_challenge_parameters](data-sources--cdn_loadbalancer--properties--enable_challenge--captcha_challenge_parameters.md): complete subsection reference.

- [default_captcha_challenge_parameters](data-sources--cdn_loadbalancer--properties--enable_challenge--default_captcha_challenge_parameters.md): complete subsection reference.

- [default_js_challenge_parameters](data-sources--cdn_loadbalancer--properties--enable_challenge--default_js_challenge_parameters.md): complete subsection reference.

- [default_mitigation_settings](data-sources--cdn_loadbalancer--properties--enable_challenge--default_mitigation_settings.md): complete subsection reference.

- [js_challenge_parameters](data-sources--cdn_loadbalancer--properties--enable_challenge--js_challenge_parameters.md): complete subsection reference.

- [malicious_user_mitigation](data-sources--cdn_loadbalancer--properties--enable_challenge--malicious_user_mitigation.md): complete subsection reference.

## Next pages

- [enable_challenge.captcha_challenge_parameters](data-sources--cdn_loadbalancer--properties--enable_challenge--captcha_challenge_parameters.md)
- [enable_challenge.default_captcha_challenge_parameters](data-sources--cdn_loadbalancer--properties--enable_challenge--default_captcha_challenge_parameters.md)
- [enable_challenge.default_js_challenge_parameters](data-sources--cdn_loadbalancer--properties--enable_challenge--default_js_challenge_parameters.md)
- [enable_challenge.default_mitigation_settings](data-sources--cdn_loadbalancer--properties--enable_challenge--default_mitigation_settings.md)
- [enable_challenge.js_challenge_parameters](data-sources--cdn_loadbalancer--properties--enable_challenge--js_challenge_parameters.md)
- [enable_challenge.malicious_user_mitigation](data-sources--cdn_loadbalancer--properties--enable_challenge--malicious_user_mitigation.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
