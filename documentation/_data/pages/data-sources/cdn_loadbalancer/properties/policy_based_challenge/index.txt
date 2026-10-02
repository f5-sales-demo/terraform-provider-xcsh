---
page_title: "policy_based_challenge"
subcategory: "Load Balancing"
description: "Specifies the settings for policy rule based challenge."
xcsh_docs: {"aliases": ["policy based challenge"], "body_bytes": 6581, "body_sha256": "sha256:94b7df1c968655e0b6c84c1f3ed34a9368c64fb6efcf299984d31f8607d95b85", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:always_enable_captcha_challenge", "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:always_enable_js_challenge", "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:captcha_challenge_parameters", "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:default_captcha_challenge_parameters", "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:default_js_challenge_parameters", "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:default_mitigation_settings", "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:default_temporary_blocking_parameters", "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:js_challenge_parameters", "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:malicious_user_mitigation", "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:no_challenge", "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list", "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:temporary_user_blocking"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "documentation/data-sources/cdn_loadbalancer/properties/policy_based_challenge/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_based_challenge"], "schema_version": 1, "sections": [{"aliases": ["always enable captcha challenge"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:always_enable_captcha_challenge", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "always_enable_captcha_challenge"], "syntax": "attribute", "type": "object"}, {"aliases": ["always enable js challenge"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:always_enable_js_challenge", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "always_enable_js_challenge"], "syntax": "attribute", "type": "object"}, {"aliases": ["captcha challenge parameters", "login success", "succeeded", "success", "successful"], "anchor": "section", "description": "Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:captcha_challenge_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_based_challenge", "captcha_challenge_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["default captcha challenge parameters"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:default_captcha_challenge_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "default_captcha_challenge_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["default js challenge parameters"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:default_js_challenge_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "default_js_challenge_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["default mitigation settings"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:default_mitigation_settings", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "default_mitigation_settings"], "syntax": "attribute", "type": "object"}, {"aliases": ["default temporary blocking parameters"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:default_temporary_blocking_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "default_temporary_blocking_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["js challenge parameters"], "anchor": "section", "description": "Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:js_challenge_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_based_challenge", "js_challenge_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["malicious user mitigation"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:malicious_user_mitigation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_based_challenge", "malicious_user_mitigation"], "syntax": "attribute", "type": "object"}, {"aliases": ["no challenge"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:no_challenge", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "no_challenge"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list"], "anchor": "section", "description": "List of challenge rules to be used in policy based challenge.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["temporary user blocking"], "anchor": "section", "description": "Specifies configuration for temporary user blocking resulting from user behavior analysis. When Malicious User Mitigation is enabled from service policy rules, users' accessing the application will be analyzed for malicious activity and the configured mitigation actions will be taken on identified malicious users.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:temporary_user_blocking", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_based_challenge", "temporary_user_blocking"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/policy_based_challenge/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specifies the settings for policy rule based challenge.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- policy_based_challenge

<a id="section"></a>

Type: `"single"`. Computed.

Specifies the settings for policy rule based challenge.

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

## Direct properties

- [always_enable_captcha_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/always_enable_captcha_challenge/): complete subsection reference.

- [always_enable_js_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/always_enable_js_challenge/): complete subsection reference.

- [captcha_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/captcha_challenge_parameters/): complete subsection reference.

- [default_captcha_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/default_captcha_challenge_parameters/): complete subsection reference.

- [default_js_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/default_js_challenge_parameters/): complete subsection reference.

- [default_mitigation_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/default_mitigation_settings/): complete subsection reference.

- [default_temporary_blocking_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/default_temporary_blocking_parameters/): complete subsection reference.

- [js_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/js_challenge_parameters/): complete subsection reference.

- [malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/malicious_user_mitigation/): complete subsection reference.

- [no_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/no_challenge/): complete subsection reference.

- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/): complete subsection reference.

- [temporary_user_blocking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/temporary_user_blocking/): complete subsection reference.

## Next pages

- [policy_based_challenge.always_enable_captcha_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/always_enable_captcha_challenge/)
- [policy_based_challenge.always_enable_js_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/always_enable_js_challenge/)
- [policy_based_challenge.captcha_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/captcha_challenge_parameters/)
- [policy_based_challenge.default_captcha_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/default_captcha_challenge_parameters/)
- [policy_based_challenge.default_js_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/default_js_challenge_parameters/)
- [policy_based_challenge.default_mitigation_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/default_mitigation_settings/)
- [policy_based_challenge.default_temporary_blocking_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/default_temporary_blocking_parameters/)
- [policy_based_challenge.js_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/js_challenge_parameters/)
- [policy_based_challenge.malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/malicious_user_mitigation/)
- [policy_based_challenge.no_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/no_challenge/)
- [policy_based_challenge.rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/)
- [policy_based_challenge.temporary_user_blocking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/temporary_user_blocking/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
