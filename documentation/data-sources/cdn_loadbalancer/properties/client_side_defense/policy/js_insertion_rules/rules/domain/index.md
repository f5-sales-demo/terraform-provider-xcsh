---
page_title: "client_side_defense.policy.js_insertion_rules.rules.domain"
subcategory: "Load Balancing"
description: "Domains names."
xcsh_docs: {"aliases": ["client side defense policy js insertion rules rules domain"], "body_bytes": 5646, "body_sha256": "sha256:5a1365896d6f938e20bf1a7da880f60e3c2a370bd181c59794c091b41f26e402", "capabilities": ["cdn", "security.client-side-defense"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:domain", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules", "path": "documentation/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/domain/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0331232032213020-3320203122130001-0201311133300031-2321200333311022-0222110302110002-2013020031210203-2213330012122123-2311030210200222", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules", "domain"], "schema_version": 1, "sections": [{"aliases": ["client side defense policy js insertion rules rules domain exact value"], "anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--domain--exact_value", "description": "Exclusive with Exact domain name.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:domain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules", "domain", "exact_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["client side defense policy js insertion rules rules domain regex value"], "anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--domain--regex_value", "description": "Exclusive with Regular Expression value for the domain name.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:domain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules", "domain", "regex_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["client side defense policy js insertion rules rules domain suffix value"], "anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--domain--suffix_value", "description": "Exclusive with Suffix of domain name e.g \"xyz.com\" will match \"*.xyz.com\" and \"xyz.com\"", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:domain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules", "domain", "suffix_value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/domain/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Domains names.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy.js_insertion_rules.rules.domain

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/)
- [client_side_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/)
- [client_side_defense.policy.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/)
- [client_side_defense.policy.js_insertion_rules.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/)
- client_side_defense.policy.js_insertion_rules.rules.domain

<a id="section"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

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

<a id="schema-client_side_defense--policy--js_insertion_rules--rules--domain--exact_value"></a>

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

<a id="schema-client_side_defense--policy--js_insertion_rules--rules--domain--regex_value"></a>

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

<a id="schema-client_side_defense--policy--js_insertion_rules--rules--domain--suffix_value"></a>

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

- [client_side_defense.policy.js_insertion_rules.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
