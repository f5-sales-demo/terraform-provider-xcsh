---
page_title: "l7_ddos_protection"
subcategory: "Load Balancing"
description: "L7 DDoS protection is critical for safeguarding web applications, APIs, and services that are exposed to the internet from sophisticated, volumetric, application-level threats. Configure actions, thresholds and policies to apply during L7 DDoS attack."
xcsh_docs: {"aliases": ["l7 ddos protection"], "body_bytes": 5301, "body_sha256": "sha256:e375453d3c4ff512fe2ab60081ce588e6dda988e05f590362f20b7445e610635", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_captcha_challenge", "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_js_challenge", "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_none", "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:ddos_policy_custom", "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:ddos_policy_none", "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:default_rps_threshold", "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:mitigation_block", "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:mitigation_captcha_challenge", "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:mitigation_js_challenge"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/l7_ddos_protection/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-020.md", "relationships": [{"anchor": "schema-l7_ddos_protection--rps_threshold", "enforcement": "provider-schema", "group": "l7_ddos_protection:ConflictingObjectAttributes:default_rps_threshold,rps_threshold", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "l7_ddos_protection:ConflictingObjectAttributes:clientside_action_captcha_challenge,clientside_action_js_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_captcha_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "l7_ddos_protection:ConflictingObjectAttributes:clientside_action_captcha_challenge,clientside_action_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_captcha_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "l7_ddos_protection:ConflictingObjectAttributes:clientside_action_captcha_challenge,clientside_action_js_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_js_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "l7_ddos_protection:ConflictingObjectAttributes:clientside_action_js_challenge,clientside_action_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_js_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "l7_ddos_protection:ConflictingObjectAttributes:clientside_action_captcha_challenge,clientside_action_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "l7_ddos_protection:ConflictingObjectAttributes:clientside_action_js_challenge,clientside_action_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "l7_ddos_protection:ConflictingObjectAttributes:ddos_policy_custom,ddos_policy_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:ddos_policy_custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "l7_ddos_protection:ConflictingObjectAttributes:ddos_policy_custom,ddos_policy_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:ddos_policy_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "l7_ddos_protection:ConflictingObjectAttributes:default_rps_threshold,rps_threshold", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:default_rps_threshold", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "l7_ddos_protection:ConflictingObjectAttributes:mitigation_block,mitigation_captcha_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:mitigation_block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "l7_ddos_protection:ConflictingObjectAttributes:mitigation_block,mitigation_js_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:mitigation_block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "l7_ddos_protection:ConflictingObjectAttributes:mitigation_block,mitigation_captcha_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:mitigation_captcha_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "l7_ddos_protection:ConflictingObjectAttributes:mitigation_captcha_challenge,mitigation_js_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:mitigation_captcha_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "l7_ddos_protection:ConflictingObjectAttributes:mitigation_block,mitigation_js_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:mitigation_js_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "l7_ddos_protection:ConflictingObjectAttributes:mitigation_captcha_challenge,mitigation_js_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:mitigation_js_challenge", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["l7_ddos_protection"], "schema_version": 1, "sections": [{"aliases": ["l7 ddos protection clientside action captcha challenge", "succeeded", "success", "successful"], "anchor": "section", "description": "Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_captcha_challenge", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-l7_ddos_protection--clientside_action_captcha_challenge--cookie_expiry", "enforcement": "provider-schema", "group": "l7_ddos_protection.clientside_action_captcha_challenge:RequiredObjectAttributes:cookie_expiry", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_captcha_challenge", "type": "requires"}], "schema_path": ["l7_ddos_protection", "clientside_action_captcha_challenge"], "syntax": "block", "type": "object"}, {"aliases": ["l7 ddos protection clientside action js challenge"], "anchor": "section", "description": "Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_js_challenge", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-l7_ddos_protection--clientside_action_js_challenge--cookie_expiry", "enforcement": "provider-schema", "group": "l7_ddos_protection.clientside_action_js_challenge:RequiredObjectAttributes:cookie_expiry,js_script_delay", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_js_challenge", "type": "requires"}, {"anchor": "schema-l7_ddos_protection--clientside_action_js_challenge--js_script_delay", "enforcement": "provider-schema", "group": "l7_ddos_protection.clientside_action_js_challenge:RequiredObjectAttributes:cookie_expiry,js_script_delay", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_js_challenge", "type": "requires"}], "schema_path": ["l7_ddos_protection", "clientside_action_js_challenge"], "syntax": "block", "type": "object"}, {"aliases": ["l7 ddos protection clientside action none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_none", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["l7_ddos_protection", "clientside_action_none"], "syntax": "attribute", "type": "object"}, {"aliases": ["l7 ddos protection ddos policy custom"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:ddos_policy_custom", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-l7_ddos_protection--ddos_policy_custom--name", "enforcement": "provider-schema", "group": "l7_ddos_protection.ddos_policy_custom:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:ddos_policy_custom", "type": "requires"}], "schema_path": ["l7_ddos_protection", "ddos_policy_custom"], "syntax": "block", "type": "object"}, {"aliases": ["l7 ddos protection ddos policy none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:ddos_policy_none", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["l7_ddos_protection", "ddos_policy_none"], "syntax": "attribute", "type": "object"}, {"aliases": ["l7 ddos protection default rps threshold"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:default_rps_threshold", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["l7_ddos_protection", "default_rps_threshold"], "syntax": "attribute", "type": "object"}, {"aliases": ["l7 ddos protection mitigation block"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:mitigation_block", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["l7_ddos_protection", "mitigation_block"], "syntax": "attribute", "type": "object"}, {"aliases": ["l7 ddos protection mitigation captcha challenge", "succeeded", "success", "successful"], "anchor": "section", "description": "Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:mitigation_captcha_challenge", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-l7_ddos_protection--mitigation_captcha_challenge--cookie_expiry", "enforcement": "provider-schema", "group": "l7_ddos_protection.mitigation_captcha_challenge:RequiredObjectAttributes:cookie_expiry", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:mitigation_captcha_challenge", "type": "requires"}], "schema_path": ["l7_ddos_protection", "mitigation_captcha_challenge"], "syntax": "block", "type": "object"}, {"aliases": ["l7 ddos protection mitigation js challenge"], "anchor": "section", "description": "Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:mitigation_js_challenge", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-l7_ddos_protection--mitigation_js_challenge--cookie_expiry", "enforcement": "provider-schema", "group": "l7_ddos_protection.mitigation_js_challenge:RequiredObjectAttributes:cookie_expiry,js_script_delay", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:mitigation_js_challenge", "type": "requires"}, {"anchor": "schema-l7_ddos_protection--mitigation_js_challenge--js_script_delay", "enforcement": "provider-schema", "group": "l7_ddos_protection.mitigation_js_challenge:RequiredObjectAttributes:cookie_expiry,js_script_delay", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:mitigation_js_challenge", "type": "requires"}], "schema_path": ["l7_ddos_protection", "mitigation_js_challenge"], "syntax": "block", "type": "object"}, {"aliases": ["l7 ddos protection rps threshold"], "anchor": "schema-l7_ddos_protection--rps_threshold", "description": "Exclusive with Configure custom RPS threshold.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["l7_ddos_protection", "rps_threshold"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/l7_ddos_protection/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "L7 DDoS protection is critical for safeguarding web applications, APIs, and services that are exposed to the internet from sophisticated, volumetric, application-level threats. Configure actions, thresholds and policies to apply during L7 DDoS attack.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# l7_ddos_protection

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- l7_ddos_protection

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

L7 DDoS protection is critical for safeguarding web applications, APIs, and services that are
exposed to the internet from sophisticated, volumetric, application-level threats. Configure
actions, thresholds and policies to apply during L7 DDoS attack. Defaults to \`map\[\]\`. Server
applies default when omitted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("clientside_action_captcha_challenge",
    "clientside_action_js_challenge"),
  validators.ConflictingObjectAttributes("clientside_action_captcha_challenge",
    "clientside_action_none"),
  validators.ConflictingObjectAttributes("clientside_action_js_challenge",
    "clientside_action_none"),
  validators.ConflictingObjectAttributes("ddos_policy_custom",
    "ddos_policy_none"),
  validators.ConflictingObjectAttributes("default_rps_threshold",
    "rps_threshold"),
  validators.ConflictingObjectAttributes("mitigation_block",
    "mitigation_captcha_challenge"),
  validators.ConflictingObjectAttributes("mitigation_block",
    "mitigation_js_challenge"),
  validators.ConflictingObjectAttributes("mitigation_captcha_challenge",
    "mitigation_js_challenge")}
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
  "x-ves-oneof-field-clientside_action_choice": "[\"clientside_action_captcha_challenge\",\"clientside_action_js_challenge\",\"clientside_action_none\"]",
  "x-ves-oneof-field-ddos_policy_choice": "[\"ddos_policy_custom\",\"ddos_policy_none\"]",
  "x-ves-oneof-field-mitigation_action_choice": "[\"mitigation_block\",\"mitigation_captcha_challenge\",\"mitigation_js_challenge\"]",
  "x-ves-oneof-field-rps_threshold_choice": "[\"default_rps_threshold\",\"rps_threshold\"]"
}
```

Terraform syntax:

```terraform
l7_ddos_protection {
  # Configure direct properties listed below.
}
```

## Direct properties

- [clientside_action_captcha_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/l7_ddos_protection/clientside_action_captcha_challenge/): complete subsection reference.

- [clientside_action_js_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/l7_ddos_protection/clientside_action_js_challenge/): complete subsection reference.

- [clientside_action_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/l7_ddos_protection/clientside_action_none/): complete subsection reference.

- [ddos_policy_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/l7_ddos_protection/ddos_policy_custom/): complete subsection reference.

- [ddos_policy_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/l7_ddos_protection/ddos_policy_none/): complete subsection reference.

- [default_rps_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/l7_ddos_protection/default_rps_threshold/): complete subsection reference.

- [mitigation_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/l7_ddos_protection/mitigation_block/): complete subsection reference.

- [mitigation_captcha_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/l7_ddos_protection/mitigation_captcha_challenge/): complete subsection reference.

- [mitigation_js_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/l7_ddos_protection/mitigation_js_challenge/): complete subsection reference.

<a id="schema-l7_ddos_protection--rps_threshold"></a>

### rps_threshold property

Type: `"number"`. Optional.

Exclusive with \[default\_rps\_threshold\] Configure custom RPS threshold.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 50000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 50000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "50000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "50000"
  }
}
```
