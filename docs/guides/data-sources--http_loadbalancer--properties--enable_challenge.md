---
page_title: "enable_challenge"
subcategory: "Load Balancing"
description: "enable_challenge for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3000, "body_sha256": "sha256:d972a6be8aa8ab82f68902efc6daa3b90f19d0f00e05049086d9479a35b99618", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge:captcha_challenge_parameters", "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge:default_captcha_challenge_parameters", "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge:default_js_challenge_parameters", "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge:default_mitigation_settings", "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge:js_challenge_parameters", "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge:malicious_user_mitigation"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "docs/guides/data-sources--http_loadbalancer--properties--enable_challenge.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_challenge"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/enable_challenge/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_challenge for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_challenge

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
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

- [captcha_challenge_parameters](data-sources--http_loadbalancer--properties--enable_challenge--captcha_challenge_parameters.md): complete subsection reference.

- [default_captcha_challenge_parameters](data-sources--http_loadbalancer--properties--enable_challenge--default_captcha_challenge_parameters.md): complete subsection reference.

- [default_js_challenge_parameters](data-sources--http_loadbalancer--properties--enable_challenge--default_js_challenge_parameters.md): complete subsection reference.

- [default_mitigation_settings](data-sources--http_loadbalancer--properties--enable_challenge--default_mitigation_settings.md): complete subsection reference.

- [js_challenge_parameters](data-sources--http_loadbalancer--properties--enable_challenge--js_challenge_parameters.md): complete subsection reference.

- [malicious_user_mitigation](data-sources--http_loadbalancer--properties--enable_challenge--malicious_user_mitigation.md): complete subsection reference.

## Next pages

- [enable_challenge.captcha_challenge_parameters](data-sources--http_loadbalancer--properties--enable_challenge--captcha_challenge_parameters.md)
- [enable_challenge.default_captcha_challenge_parameters](data-sources--http_loadbalancer--properties--enable_challenge--default_captcha_challenge_parameters.md)
- [enable_challenge.default_js_challenge_parameters](data-sources--http_loadbalancer--properties--enable_challenge--default_js_challenge_parameters.md)
- [enable_challenge.default_mitigation_settings](data-sources--http_loadbalancer--properties--enable_challenge--default_mitigation_settings.md)
- [enable_challenge.js_challenge_parameters](data-sources--http_loadbalancer--properties--enable_challenge--js_challenge_parameters.md)
- [enable_challenge.malicious_user_mitigation](data-sources--http_loadbalancer--properties--enable_challenge--malicious_user_mitigation.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
