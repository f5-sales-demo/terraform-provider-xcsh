---
page_title: "custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies"
subcategory: ""
description: "Ordered List of Enhanced Firewall Policies active."
xcsh_docs: {"aliases": ["custom network config active enhanced firewall policies enhanced firewall policies"], "body_bytes": 6390, "body_sha256": "sha256:f4a2644dad846fe51f43cc48cf1bf503d0d83f7f8078fa8efcf60d5a5e9c8feb", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:active_enhanced_firewall_policies:enhanced_firewall_policies", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:active_enhanced_firewall_policies", "path": "documentation/data-sources/voltstack_site/properties/custom_network_config/active_enhanced_firewall_policies/enhanced_firewall_policies/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3020323100210020-1303120302233212-3311203111133233-3330103303320033-1321132200233323-2030112033300201-2321000203011000-3111222201312030", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "active_enhanced_firewall_policies", "enhanced_firewall_policies"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:active_enhanced_firewall_policies:enhanced_firewall_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "active_enhanced_firewall_policies", "enhanced_firewall_policies", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:active_enhanced_firewall_policies:enhanced_firewall_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "active_enhanced_firewall_policies", "enhanced_firewall_policies", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["tenant"], "anchor": "schema-custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:active_enhanced_firewall_policies:enhanced_firewall_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "active_enhanced_firewall_policies", "enhanced_firewall_policies", "tenant"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_network_config/active_enhanced_firewall_policies/enhanced_firewall_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Ordered List of Enhanced Firewall Policies active.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/)
- [custom_network_config.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/active_enhanced_firewall_policies/)
- custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="section"></a>

Type: `"list"`. Computed.

Ordered List of Enhanced Firewall Policies active.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

## Direct properties

<a id="schema-custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="schema-custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="schema-custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

## Next pages

- [custom_network_config.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/active_enhanced_firewall_policies/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
