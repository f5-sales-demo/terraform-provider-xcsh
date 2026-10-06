---
page_title: "enable_challenge"
subcategory: "Load Balancing"
description: "Configure auto mitigation i.e risk based challenges for malicious users."
xcsh_docs: {"aliases": ["enable challenge"], "body_bytes": 2932, "body_sha256": "sha256:3c134145ef50ba7bf180f5020bd1c7aebbb812444b7f47bd06ba7aa68cb06e1f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:captcha_challenge_parameters", "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:default_captcha_challenge_parameters", "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:default_js_challenge_parameters", "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:default_mitigation_settings", "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:js_challenge_parameters", "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:malicious_user_mitigation"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/enable_challenge/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-018.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_challenge:ConflictingObjectAttributes:captcha_challenge_parameters,default_captcha_challenge_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:captcha_challenge_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_challenge:ConflictingObjectAttributes:captcha_challenge_parameters,default_captcha_challenge_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:default_captcha_challenge_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_challenge:ConflictingObjectAttributes:default_js_challenge_parameters,js_challenge_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:default_js_challenge_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_challenge:ConflictingObjectAttributes:default_mitigation_settings,malicious_user_mitigation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:default_mitigation_settings", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_challenge:ConflictingObjectAttributes:default_js_challenge_parameters,js_challenge_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:js_challenge_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_challenge:ConflictingObjectAttributes:default_mitigation_settings,malicious_user_mitigation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:malicious_user_mitigation", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_challenge"], "schema_version": 1, "sections": [{"aliases": ["enable challenge captcha challenge parameters", "succeeded", "success", "successful"], "anchor": "section", "description": "Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:captcha_challenge_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-enable_challenge--captcha_challenge_parameters--cookie_expiry", "enforcement": "provider-schema", "group": "enable_challenge.captcha_challenge_parameters:RequiredObjectAttributes:cookie_expiry", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:captcha_challenge_parameters", "type": "requires"}], "schema_path": ["enable_challenge", "captcha_challenge_parameters"], "syntax": "block", "type": "object"}, {"aliases": ["enable challenge default captcha challenge parameters"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:default_captcha_challenge_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_challenge", "default_captcha_challenge_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable challenge default js challenge parameters"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:default_js_challenge_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_challenge", "default_js_challenge_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable challenge default mitigation settings"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:default_mitigation_settings", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_challenge", "default_mitigation_settings"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable challenge js challenge parameters"], "anchor": "section", "description": "Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:js_challenge_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-enable_challenge--js_challenge_parameters--cookie_expiry", "enforcement": "provider-schema", "group": "enable_challenge.js_challenge_parameters:RequiredObjectAttributes:cookie_expiry,js_script_delay", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:js_challenge_parameters", "type": "requires"}, {"anchor": "schema-enable_challenge--js_challenge_parameters--js_script_delay", "enforcement": "provider-schema", "group": "enable_challenge.js_challenge_parameters:RequiredObjectAttributes:cookie_expiry,js_script_delay", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:js_challenge_parameters", "type": "requires"}], "schema_path": ["enable_challenge", "js_challenge_parameters"], "syntax": "block", "type": "object"}, {"aliases": ["enable challenge malicious user mitigation"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:malicious_user_mitigation", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-enable_challenge--malicious_user_mitigation--name", "enforcement": "provider-schema", "group": "enable_challenge.malicious_user_mitigation:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:malicious_user_mitigation", "type": "requires"}], "schema_path": ["enable_challenge", "malicious_user_mitigation"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/enable_challenge/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Configure auto mitigation i.e risk based challenges for malicious users.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_challenge

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- enable_challenge

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure auto mitigation i.e risk based challenges for malicious users.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [captcha_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_challenge/captcha_challenge_parameters/): complete subsection reference.

- [default_captcha_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_challenge/default_captcha_challenge_parameters/): complete subsection reference.

- [default_js_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_challenge/default_js_challenge_parameters/): complete subsection reference.

- [default_mitigation_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_challenge/default_mitigation_settings/): complete subsection reference.

- [js_challenge_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_challenge/js_challenge_parameters/): complete subsection reference.

- [malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_challenge/malicious_user_mitigation/): complete subsection reference.
