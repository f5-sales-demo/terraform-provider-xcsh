---
page_title: "enable_challenge"
subcategory: "Load Balancing"
description: "Configure auto mitigation i.e risk based challenges for malicious users."
xcsh_docs: {"aliases": ["enable challenge"], "body_bytes": 4295, "body_sha256": "sha256:49e1f584b7003733f2c7de1347a5560d347c7b25453c9a60e75bed3ef9dfa427", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:captcha_challenge_parameters", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:default_captcha_challenge_parameters", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:default_js_challenge_parameters", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:default_mitigation_settings", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:js_challenge_parameters", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:malicious_user_mitigation"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/enable_challenge/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-010.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_challenge:ConflictingObjectAttributes:captcha_challenge_parameters,default_captcha_challenge_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:captcha_challenge_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_challenge:ConflictingObjectAttributes:captcha_challenge_parameters,default_captcha_challenge_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:default_captcha_challenge_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_challenge:ConflictingObjectAttributes:default_js_challenge_parameters,js_challenge_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:default_js_challenge_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_challenge:ConflictingObjectAttributes:default_mitigation_settings,malicious_user_mitigation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:default_mitigation_settings", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_challenge:ConflictingObjectAttributes:default_js_challenge_parameters,js_challenge_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:js_challenge_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_challenge:ConflictingObjectAttributes:default_mitigation_settings,malicious_user_mitigation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:malicious_user_mitigation", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_challenge"], "schema_version": 1, "sections": [{"aliases": ["enable challenge captcha challenge parameters", "succeeded", "success", "successful"], "anchor": "section", "description": "Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:captcha_challenge_parameters", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-enable_challenge--captcha_challenge_parameters--cookie_expiry", "enforcement": "provider-schema", "group": "enable_challenge.captcha_challenge_parameters:RequiredObjectAttributes:cookie_expiry", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:captcha_challenge_parameters", "type": "requires"}], "schema_path": ["enable_challenge", "captcha_challenge_parameters"], "syntax": "block", "type": "object"}, {"aliases": ["enable challenge default captcha challenge parameters"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:default_captcha_challenge_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_challenge", "default_captcha_challenge_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable challenge default js challenge parameters"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:default_js_challenge_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_challenge", "default_js_challenge_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable challenge default mitigation settings"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:default_mitigation_settings", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_challenge", "default_mitigation_settings"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable challenge js challenge parameters"], "anchor": "section", "description": "Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:js_challenge_parameters", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-enable_challenge--js_challenge_parameters--cookie_expiry", "enforcement": "provider-schema", "group": "enable_challenge.js_challenge_parameters:RequiredObjectAttributes:cookie_expiry,js_script_delay", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:js_challenge_parameters", "type": "requires"}, {"anchor": "schema-enable_challenge--js_challenge_parameters--js_script_delay", "enforcement": "provider-schema", "group": "enable_challenge.js_challenge_parameters:RequiredObjectAttributes:cookie_expiry,js_script_delay", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:js_challenge_parameters", "type": "requires"}], "schema_path": ["enable_challenge", "js_challenge_parameters"], "syntax": "block", "type": "object"}, {"aliases": ["enable challenge malicious user mitigation"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:malicious_user_mitigation", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-enable_challenge--malicious_user_mitigation--name", "enforcement": "provider-schema", "group": "enable_challenge.malicious_user_mitigation:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:malicious_user_mitigation", "type": "requires"}], "schema_path": ["enable_challenge", "malicious_user_mitigation"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/enable_challenge/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Configure auto mitigation i.e risk based challenges for malicious users.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
