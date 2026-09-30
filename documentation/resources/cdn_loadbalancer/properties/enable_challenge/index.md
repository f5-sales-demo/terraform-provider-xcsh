---
page_title: "enable_challenge"
subcategory: "Load Balancing"
description: "enable_challenge for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4196, "body_sha256": "sha256:69d903fcd90bb93364c203033e57f19ae57bb477ca4143bb988b15a3b41ffc08", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:captcha_challenge_parameters", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:default_captcha_challenge_parameters", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:default_js_challenge_parameters", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:default_mitigation_settings", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:js_challenge_parameters", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:malicious_user_mitigation"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/enable_challenge/index.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["enable_challenge"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/enable_challenge/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_challenge for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# enable_challenge

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- enable_challenge

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure auto mitigation i.e risk based challenges for malicious users.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("captcha_challenge_parameters",
    "default_captcha_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_js_challenge_parameters",
    "js_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_mitigation_settings",
    "malicious_user_mitigation")}
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
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]"
}
```

Terraform syntax:

```terraform
enable_challenge {
  # Configure direct properties listed below.
}
```

## Direct properties

- [captcha_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_challenge/captcha_challenge_parameters/): complete subsection reference.

- [default_captcha_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_challenge/default_captcha_challenge_parameters/): complete subsection reference.

- [default_js_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_challenge/default_js_challenge_parameters/): complete subsection reference.

- [default_mitigation_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_challenge/default_mitigation_settings/): complete subsection reference.

- [js_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_challenge/js_challenge_parameters/): complete subsection reference.

- [malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_challenge/malicious_user_mitigation/): complete subsection reference.

## Next pages

- [enable_challenge.captcha_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_challenge/captcha_challenge_parameters/)
- [enable_challenge.default_captcha_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_challenge/default_captcha_challenge_parameters/)
- [enable_challenge.default_js_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_challenge/default_js_challenge_parameters/)
- [enable_challenge.default_mitigation_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_challenge/default_mitigation_settings/)
- [enable_challenge.js_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_challenge/js_challenge_parameters/)
- [enable_challenge.malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_challenge/malicious_user_mitigation/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
