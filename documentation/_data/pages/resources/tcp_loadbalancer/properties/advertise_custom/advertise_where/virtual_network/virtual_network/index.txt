---
page_title: "advertise_custom.advertise_where.virtual_network.virtual_network"
subcategory: "Load Balancing"
description: "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name."
xcsh_docs: {"aliases": ["advertise custom advertise where virtual network virtual network"], "body_bytes": 5141, "body_sha256": "sha256:7b53f8eaca0ed20f7ba34562786d039f9a1229e35d5b765acc9fa9375fb7d6cc", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:virtual_network", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "path": "documentation/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/virtual_network/virtual_network/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0320000010022301-3300003102221232-3203200210312200-0013230101012112-0012033102333001-2000100301113312-3220103030121213-2102232111101311", "registry_path": "docs/guides/resources--tcp_loadbalancer--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advertise_custom", "advertise_where", "virtual_network", "virtual_network"], "schema_version": 1, "sections": [{"aliases": ["advertise custom advertise where virtual network virtual network name"], "anchor": "schema-advertise_custom--advertise_where--virtual_network--virtual_network--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:virtual_network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "name", "scope_path": ["advertise_custom", "advertise_where", "virtual_network", "virtual_network"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.views.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "virtual_network", "virtual_network", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["advertise custom advertise where virtual network virtual network namespace"], "anchor": "schema-advertise_custom--advertise_where--virtual_network--virtual_network--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:virtual_network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "namespace", "scope_path": ["advertise_custom", "advertise_where", "virtual_network", "virtual_network"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.views.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "virtual_network", "virtual_network", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["advertise custom advertise where virtual network virtual network tenant"], "anchor": "schema-advertise_custom--advertise_where--virtual_network--virtual_network--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:virtual_network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "tenant", "scope_path": ["advertise_custom", "advertise_where", "virtual_network", "virtual_network"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.views.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "virtual_network", "virtual_network", "tenant"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/virtual_network/virtual_network/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom.advertise_where.virtual_network.virtual_network

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/)
- [advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/)
- [advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/)
- [advertise_custom.advertise_where.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/virtual_network/)
- advertise_custom.advertise_where.virtual_network.virtual_network

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

Terraform syntax:

```terraform
virtual_network {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-advertise_custom--advertise_where--virtual_network--virtual_network--name"></a>

### name property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-advertise_custom--advertise_where--virtual_network--virtual_network--namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-advertise_custom--advertise_where--virtual_network--virtual_network--tenant"></a>

### tenant property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
