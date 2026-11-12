---
page_title: "custom_anonymization.anonymization_config.cookie"
subcategory: "Security"
description: "Configure anonymization for HTTP Cookies."
xcsh_docs: {"aliases": ["custom anonymization anonymization config cookie"], "body_bytes": 2149, "body_sha256": "sha256:2974d6bbc5a6162c4183e28d88167fc79d6d153d470f5a570f86087b01ea3e34", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "parent_id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config", "path": "documentation/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/cookie/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-0002031121233033-3331213132001322-0131001022112023-2131333200211310-3110013130001123-3123110231001221-3113221002232311-2202223232011211", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_anonymization", "anonymization_config", "cookie"], "schema_version": 1, "sections": [{"aliases": ["custom anonymization anonymization config cookie cookie name"], "anchor": "schema-custom_anonymization--anonymization_config--cookie--cookie_name", "description": "Masks the cookie value. The setting does not mask the cookie name. Wildcard matching can be used by prefixing or suffixing the cookie name with a wildcard asterisk (*), or by using only an asterisk to match any cookie name.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_anonymization", "anonymization_config", "cookie", "cookie_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/cookie/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Configure anonymization for HTTP Cookies.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["app_firewallCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_anonymization.anonymization_config.cookie

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- [custom_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/)
- [custom_anonymization.anonymization_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/)
- custom_anonymization.anonymization_config.cookie

<a id="section"></a>

Type: `"single"`. Computed.

Configure anonymization for HTTP Cookies.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

<a id="schema-custom_anonymization--anonymization_config--cookie--cookie_name"></a>

### cookie_name property

Type: `"string"`. Computed.

Masks the cookie value. The setting does not mask the cookie name. Wildcard matching can be used by
prefixing or suffixing the cookie name with a wildcard asterisk (\*), or by using only an asterisk
to match any cookie name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-29T12:14:42+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```
