---
page_title: "l7_ddos_protection"
subcategory: "Load Balancing"
description: "L7 DDoS protection is critical for safeguarding web applications, APIs, and services that are exposed to the internet from sophisticated, volumetric, application-level threats. Configure actions, thresholds and policies to apply during L7 DDoS attack."
xcsh_docs: {"aliases": ["l7 ddos protection"], "body_bytes": 6511, "body_sha256": "sha256:462c7532f6c14cdf77e800f7f70701d940f16b1a173beb9b4aae0b34ce3b46cb", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_captcha_challenge", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_js_challenge", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_none", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:ddos_policy_custom", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:ddos_policy_none", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:default_rps_threshold", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:mitigation_block", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:mitigation_captcha_challenge", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:mitigation_js_challenge"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "documentation/data-sources/http_loadbalancer/properties/l7_ddos_protection/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3033222321011321-1022100112201200-2011110213223303-0111130210011233-3303003112011322-3001020230211012-0110113133233100-1303211333203133", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-020.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["l7_ddos_protection"], "schema_version": 1, "sections": [{"aliases": ["clientside action captcha challenge", "login success", "succeeded", "success", "successful"], "anchor": "section", "description": "Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_captcha_challenge", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["l7_ddos_protection", "clientside_action_captcha_challenge"], "syntax": "attribute", "type": "object"}, {"aliases": ["clientside action js challenge"], "anchor": "section", "description": "Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_js_challenge", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["l7_ddos_protection", "clientside_action_js_challenge"], "syntax": "attribute", "type": "object"}, {"aliases": ["clientside action none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_none", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["l7_ddos_protection", "clientside_action_none"], "syntax": "attribute", "type": "object"}, {"aliases": ["ddos policy custom"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:ddos_policy_custom", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["l7_ddos_protection", "ddos_policy_custom"], "syntax": "attribute", "type": "object"}, {"aliases": ["ddos policy none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:ddos_policy_none", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["l7_ddos_protection", "ddos_policy_none"], "syntax": "attribute", "type": "object"}, {"aliases": ["default rps threshold"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:default_rps_threshold", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["l7_ddos_protection", "default_rps_threshold"], "syntax": "attribute", "type": "object"}, {"aliases": ["mitigation block"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:mitigation_block", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["l7_ddos_protection", "mitigation_block"], "syntax": "attribute", "type": "object"}, {"aliases": ["login success", "mitigation captcha challenge", "succeeded", "success", "successful"], "anchor": "section", "description": "Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:mitigation_captcha_challenge", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["l7_ddos_protection", "mitigation_captcha_challenge"], "syntax": "attribute", "type": "object"}, {"aliases": ["mitigation js challenge"], "anchor": "section", "description": "Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:mitigation_js_challenge", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["l7_ddos_protection", "mitigation_js_challenge"], "syntax": "attribute", "type": "object"}, {"aliases": ["rps threshold"], "anchor": "schema-l7_ddos_protection--rps_threshold", "description": "Exclusive with Configure custom RPS threshold.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["l7_ddos_protection", "rps_threshold"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/l7_ddos_protection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "L7 DDoS protection is critical for safeguarding web applications, APIs, and services that are exposed to the internet from sophisticated, volumetric, application-level threats. Configure actions, thresholds and policies to apply during L7 DDoS attack.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# l7_ddos_protection

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- l7_ddos_protection

<a id="section"></a>

Type: `"single"`. Computed.

L7 DDoS protection is critical for safeguarding web applications, APIs, and services that are
exposed to the internet from sophisticated, volumetric, application-level threats. Configure
actions, thresholds and policies to apply during L7 DDoS attack. Defaults to \`map\[\]\`. Server
applies default when omitted.

Upstream description:

L7 DDoS protection is critical for safeguarding web applications, APIs, and services that are
exposed to the internet from sophisticated, volumetric, application-level threats. Configure
actions, thresholds and policies to apply during L7 DDoS attack.

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

## Direct properties

- [clientside_action_captcha_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/l7_ddos_protection/clientside_action_captcha_challenge/): complete subsection reference.

- [clientside_action_js_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/l7_ddos_protection/clientside_action_js_challenge/): complete subsection reference.

- [clientside_action_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/l7_ddos_protection/clientside_action_none/): complete subsection reference.

- [ddos_policy_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/l7_ddos_protection/ddos_policy_custom/): complete subsection reference.

- [ddos_policy_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/l7_ddos_protection/ddos_policy_none/): complete subsection reference.

- [default_rps_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/l7_ddos_protection/default_rps_threshold/): complete subsection reference.

- [mitigation_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/l7_ddos_protection/mitigation_block/): complete subsection reference.

- [mitigation_captcha_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/l7_ddos_protection/mitigation_captcha_challenge/): complete subsection reference.

- [mitigation_js_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/l7_ddos_protection/mitigation_js_challenge/): complete subsection reference.

<a id="schema-l7_ddos_protection--rps_threshold"></a>

### rps_threshold property

Type: `"number"`. Computed.

Exclusive with \[default\_rps\_threshold\] Configure custom RPS threshold.

Upstream description:

Exclusive with \[default\_rps\_threshold\] Configure custom RPS threshold.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [l7_ddos_protection.clientside_action_captcha_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/l7_ddos_protection/clientside_action_captcha_challenge/)
- [l7_ddos_protection.clientside_action_js_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/l7_ddos_protection/clientside_action_js_challenge/)
- [l7_ddos_protection.clientside_action_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/l7_ddos_protection/clientside_action_none/)
- [l7_ddos_protection.ddos_policy_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/l7_ddos_protection/ddos_policy_custom/)
- [l7_ddos_protection.ddos_policy_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/l7_ddos_protection/ddos_policy_none/)
- [l7_ddos_protection.default_rps_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/l7_ddos_protection/default_rps_threshold/)
- [l7_ddos_protection.mitigation_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/l7_ddos_protection/mitigation_block/)
- [l7_ddos_protection.mitigation_captcha_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/l7_ddos_protection/mitigation_captcha_challenge/)
- [l7_ddos_protection.mitigation_js_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/l7_ddos_protection/mitigation_js_challenge/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
