---
page_title: "l7_ddos_protection"
subcategory: "Load Balancing"
description: "l7_ddos_protection for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 5403, "body_sha256": "sha256:c0e19e58be11b2b3c6bcc15d60a408d6c617b631feb17295167881769fc8e781", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_captcha_challenge", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_js_challenge", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_none", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:ddos_policy_custom", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:ddos_policy_none", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:default_rps_threshold", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:mitigation_block", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:mitigation_captcha_challenge", "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection:mitigation_js_challenge"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:l7_ddos_protection", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "docs/guides/data-sources--http_loadbalancer--properties--l7_ddos_protection.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["l7_ddos_protection"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/l7_ddos_protection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "l7_ddos_protection for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# l7_ddos_protection

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
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

- [clientside_action_captcha_challenge](data-sources--http_loadbalancer--properties--l7_ddos_protection--clientside_action_captcha_challenge.md): complete subsection reference.

- [clientside_action_js_challenge](data-sources--http_loadbalancer--properties--l7_ddos_protection--clientside_action_js_challenge.md): complete subsection reference.

- [clientside_action_none](data-sources--http_loadbalancer--properties--l7_ddos_protection--clientside_action_none.md): complete subsection reference.

- [ddos_policy_custom](data-sources--http_loadbalancer--properties--l7_ddos_protection--ddos_policy_custom.md): complete subsection reference.

- [ddos_policy_none](data-sources--http_loadbalancer--properties--l7_ddos_protection--ddos_policy_none.md): complete subsection reference.

- [default_rps_threshold](data-sources--http_loadbalancer--properties--l7_ddos_protection--default_rps_threshold.md): complete subsection reference.

- [mitigation_block](data-sources--http_loadbalancer--properties--l7_ddos_protection--mitigation_block.md): complete subsection reference.

- [mitigation_captcha_challenge](data-sources--http_loadbalancer--properties--l7_ddos_protection--mitigation_captcha_challenge.md): complete subsection reference.

- [mitigation_js_challenge](data-sources--http_loadbalancer--properties--l7_ddos_protection--mitigation_js_challenge.md): complete subsection reference.

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

- [l7_ddos_protection.clientside_action_captcha_challenge](data-sources--http_loadbalancer--properties--l7_ddos_protection--clientside_action_captcha_challenge.md)
- [l7_ddos_protection.clientside_action_js_challenge](data-sources--http_loadbalancer--properties--l7_ddos_protection--clientside_action_js_challenge.md)
- [l7_ddos_protection.clientside_action_none](data-sources--http_loadbalancer--properties--l7_ddos_protection--clientside_action_none.md)
- [l7_ddos_protection.ddos_policy_custom](data-sources--http_loadbalancer--properties--l7_ddos_protection--ddos_policy_custom.md)
- [l7_ddos_protection.ddos_policy_none](data-sources--http_loadbalancer--properties--l7_ddos_protection--ddos_policy_none.md)
- [l7_ddos_protection.default_rps_threshold](data-sources--http_loadbalancer--properties--l7_ddos_protection--default_rps_threshold.md)
- [l7_ddos_protection.mitigation_block](data-sources--http_loadbalancer--properties--l7_ddos_protection--mitigation_block.md)
- [l7_ddos_protection.mitigation_captcha_challenge](data-sources--http_loadbalancer--properties--l7_ddos_protection--mitigation_captcha_challenge.md)
- [l7_ddos_protection.mitigation_js_challenge](data-sources--http_loadbalancer--properties--l7_ddos_protection--mitigation_js_challenge.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
