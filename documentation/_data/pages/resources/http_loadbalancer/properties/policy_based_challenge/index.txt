---
page_title: "policy_based_challenge"
subcategory: "Load Balancing"
description: "policy_based_challenge for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 7401, "body_sha256": "sha256:b51dcc3de5e3265cdd261585ecd1b8feb1e6e1d9a4943822239234bc15310521", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:always_enable_captcha_challenge", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:always_enable_js_challenge", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:captcha_challenge_parameters", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:default_captcha_challenge_parameters", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:default_js_challenge_parameters", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:default_mitigation_settings", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:default_temporary_blocking_parameters", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:js_challenge_parameters", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:malicious_user_mitigation", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:no_challenge", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:temporary_user_blocking"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/policy_based_challenge/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["policy_based_challenge"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/policy_based_challenge/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# policy_based_challenge

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- policy_based_challenge

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies the settings for policy rule based challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("always_enable_captcha_challenge",
    "always_enable_js_challenge"),
  validators.ConflictingObjectAttributes("always_enable_captcha_challenge",
    "no_challenge"),
  validators.ConflictingObjectAttributes("always_enable_js_challenge",
    "no_challenge"),
  validators.ConflictingObjectAttributes("captcha_challenge_parameters",
    "default_captcha_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_js_challenge_parameters",
    "js_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_mitigation_settings",
    "malicious_user_mitigation"),
  validators.ConflictingObjectAttributes("default_temporary_blocking_parameters",
    "temporary_user_blocking")}
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
  "x-ves-oneof-field-challenge_choice": "[\"always_enable_captcha_challenge\",\"always_enable_js_challenge\",\"no_challenge\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]",
  "x-ves-oneof-field-temporary_blocking_parameters_choice": "[\"default_temporary_blocking_parameters\",\"temporary_user_blocking\"]"
}
```

Terraform syntax:

```terraform
policy_based_challenge {
  # Configure direct properties listed below.
}
```

## Direct properties

- [always_enable_captcha_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/always_enable_captcha_challenge/): complete subsection reference.

- [always_enable_js_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/always_enable_js_challenge/): complete subsection reference.

- [captcha_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/captcha_challenge_parameters/): complete subsection reference.

- [default_captcha_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/default_captcha_challenge_parameters/): complete subsection reference.

- [default_js_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/default_js_challenge_parameters/): complete subsection reference.

- [default_mitigation_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/default_mitigation_settings/): complete subsection reference.

- [default_temporary_blocking_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/default_temporary_blocking_parameters/): complete subsection reference.

- [js_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/js_challenge_parameters/): complete subsection reference.

- [malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/malicious_user_mitigation/): complete subsection reference.

- [no_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/no_challenge/): complete subsection reference.

- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/): complete subsection reference.

- [temporary_user_blocking](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/temporary_user_blocking/): complete subsection reference.

## Next pages

- [policy_based_challenge.always_enable_captcha_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/always_enable_captcha_challenge/)
- [policy_based_challenge.always_enable_js_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/always_enable_js_challenge/)
- [policy_based_challenge.captcha_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/captcha_challenge_parameters/)
- [policy_based_challenge.default_captcha_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/default_captcha_challenge_parameters/)
- [policy_based_challenge.default_js_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/default_js_challenge_parameters/)
- [policy_based_challenge.default_mitigation_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/default_mitigation_settings/)
- [policy_based_challenge.default_temporary_blocking_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/default_temporary_blocking_parameters/)
- [policy_based_challenge.js_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/js_challenge_parameters/)
- [policy_based_challenge.malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/malicious_user_mitigation/)
- [policy_based_challenge.no_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/no_challenge/)
- [policy_based_challenge.rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/)
- [policy_based_challenge.temporary_user_blocking](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/temporary_user_blocking/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
