---
page_title: "rule_list.rules.tls_list.tls_list"
subcategory: "Security"
description: "Domains in SNI for TLS connections."
xcsh_docs: {"aliases": ["rule list rules tls list tls list"], "body_bytes": 5721, "body_sha256": "sha256:22d1fa1e4e9e66477d5e1726f69815f31af0ed371a9ae3ea043ef5198174e043", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:tls_list:tls_list", "parent_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:tls_list", "path": "documentation/data-sources/forward_proxy_policy/properties/rule_list/rules/tls_list/tls_list/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3010320131320220-1022001113232320-3102131330323313-0200102002102232-3133121213313032-1023122302010301-3211303001020120-1312123113010233", "registry_path": "docs/guides/data-sources--forward_proxy_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "tls_list", "tls_list"], "schema_version": 1, "sections": [{"aliases": ["exact value"], "anchor": "schema-rule_list--rules--tls_list--tls_list--exact_value", "description": "Exclusive with Exact domain name.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:tls_list:tls_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "tls_list", "tls_list", "exact_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["regex value"], "anchor": "schema-rule_list--rules--tls_list--tls_list--regex_value", "description": "Exclusive with Regular Expression value for the domain name.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:tls_list:tls_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "tls_list", "tls_list", "regex_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["suffix value"], "anchor": "schema-rule_list--rules--tls_list--tls_list--suffix_value", "description": "Exclusive with Suffix of domain name e.g \"xyz.com\" will match \"*.xyz.com\" and \"xyz.com\"", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:tls_list:tls_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "tls_list", "tls_list", "suffix_value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forward_proxy_policy/properties/rule_list/rules/tls_list/tls_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Domains in SNI for TLS connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.tls_list.tls_list

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/)
- [rule_list.rules.tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/tls_list/)
- rule_list.rules.tls_list.tls_list

<a id="section"></a>

Type: `"list"`. Computed.

TLS Domains. Domains in SNI for TLS connections.

Upstream description:

Domains in SNI for TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-rule_list--rules--tls_list--tls_list--exact_value"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-rule_list--rules--tls_list--tls_list--regex_value"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-rule_list--rules--tls_list--tls_list--suffix_value"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [rule_list.rules.tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/tls_list/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
