---
page_title: "enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match"
subcategory: "Networking"
description: "Domains names."
xcsh_docs: {"aliases": ["enable forward proxy tls intercept policy interception rules domain match"], "body_bytes": 5775, "body_sha256": "sha256:34c0015053f799852aacbefd1e12d2aabd6e8b8d673062d99977054a0a8ea852", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:domain_match", "parent_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules", "path": "documentation/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/domain_match/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0122123032022002-1030232023330031-3103332213321003-2210031130223021-2021011311010330-0200232031130133-0112212001003313-1021222130030312", "registry_path": "docs/guides/data-sources--network_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept", "policy", "interception_rules", "domain_match"], "schema_version": 1, "sections": [{"aliases": ["enable forward proxy tls intercept policy interception rules domain match exact value"], "anchor": "schema-enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match--exact_value", "description": "Exclusive with Exact domain name.", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:domain_match", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "policy", "interception_rules", "domain_match", "exact_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept policy interception rules domain match regex value"], "anchor": "schema-enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match--regex_value", "description": "Exclusive with Regular Expression value for the domain name.", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:domain_match", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "policy", "interception_rules", "domain_match", "regex_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept policy interception rules domain match suffix value"], "anchor": "schema-enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match--suffix_value", "description": "Exclusive with Suffix of domain name e.g \"xyz.com\" will match \"*.xyz.com\" and \"xyz.com\"", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:domain_match", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "policy", "interception_rules", "domain_match", "suffix_value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/domain_match/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Domains names.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/)
- [enable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/)
- [enable_forward_proxy.tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/)
- [enable_forward_proxy.tls_intercept.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/)
- [enable_forward_proxy.tls_intercept.policy.interception_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/)
- enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for domain match.

Upstream description:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

## Direct properties

<a id="schema-enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match--exact_value"></a>

### exact_value property

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match--regex_value"></a>

### regex_value property

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="schema-enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match--suffix_value"></a>

### suffix_value property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

## Next pages

- [enable_forward_proxy.tls_intercept.policy.interception_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/)
- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
