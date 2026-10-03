---
page_title: "enable_challenge"
subcategory: "Load Balancing"
description: "Configure auto mitigation i.e risk based challenges for malicious users."
xcsh_docs: {"aliases": ["enable challenge"], "body_bytes": 3808, "body_sha256": "sha256:c27f289e909b46298b4d90c418b9de502e815aec93a6ba1abd7f067edaf95fe8", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge:captcha_challenge_parameters", "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge:default_captcha_challenge_parameters", "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge:default_js_challenge_parameters", "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge:default_mitigation_settings", "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge:js_challenge_parameters", "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge:malicious_user_mitigation"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "documentation/data-sources/http_loadbalancer/properties/enable_challenge/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-018.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_challenge"], "schema_version": 1, "sections": [{"aliases": ["enable challenge captcha challenge parameters", "succeeded", "success", "successful"], "anchor": "section", "description": "Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge:captcha_challenge_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_challenge", "captcha_challenge_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable challenge default captcha challenge parameters"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge:default_captcha_challenge_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_challenge", "default_captcha_challenge_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable challenge default js challenge parameters"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge:default_js_challenge_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_challenge", "default_js_challenge_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable challenge default mitigation settings"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge:default_mitigation_settings", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_challenge", "default_mitigation_settings"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable challenge js challenge parameters"], "anchor": "section", "description": "Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge:js_challenge_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_challenge", "js_challenge_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable challenge malicious user mitigation"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_challenge:malicious_user_mitigation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_challenge", "malicious_user_mitigation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/enable_challenge/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure auto mitigation i.e risk based challenges for malicious users.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_challenge

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
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

- [captcha_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/enable_challenge/captcha_challenge_parameters/): complete subsection reference.

- [default_captcha_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/enable_challenge/default_captcha_challenge_parameters/): complete subsection reference.

- [default_js_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/enable_challenge/default_js_challenge_parameters/): complete subsection reference.

- [default_mitigation_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/enable_challenge/default_mitigation_settings/): complete subsection reference.

- [js_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/enable_challenge/js_challenge_parameters/): complete subsection reference.

- [malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/enable_challenge/malicious_user_mitigation/): complete subsection reference.

## Next pages

- [enable_challenge.captcha_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/enable_challenge/captcha_challenge_parameters/)
- [enable_challenge.default_captcha_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/enable_challenge/default_captcha_challenge_parameters/)
- [enable_challenge.default_js_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/enable_challenge/default_js_challenge_parameters/)
- [enable_challenge.default_mitigation_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/enable_challenge/default_mitigation_settings/)
- [enable_challenge.js_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/enable_challenge/js_challenge_parameters/)
- [enable_challenge.malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/enable_challenge/malicious_user_mitigation/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
