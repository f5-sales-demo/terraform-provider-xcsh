---
page_title: "dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca"
subcategory: ""
description: "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name."
xcsh_docs: {"aliases": ["dynamic proxy https proxy tls params use mtls trusted ca"], "body_bytes": 5059, "body_sha256": "sha256:7d61dab9e8b123181b6549d5c2ad1b78374d8647eb8b7bde1435a367458a94f6", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls:trusted_ca", "parent_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls", "path": "documentation/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/use_mtls/trusted_ca/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0113322102331021-0102312331132203-1131012310302113-3003012200201322-0012303213001132-2323010231111133-2130300201322033-1031010133201001", "registry_path": "docs/guides/data-sources--proxy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls", "trusted_ca"], "schema_version": 1, "sections": [{"aliases": ["dynamic proxy https proxy tls params use mtls trusted ca name"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--use_mtls--trusted_ca--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls:trusted_ca", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "name", "scope_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls", "trusted_ca"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.views.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls", "trusted_ca", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params use mtls trusted ca namespace"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--use_mtls--trusted_ca--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls:trusted_ca", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "namespace", "scope_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls", "trusted_ca"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.views.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls", "trusted_ca", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params use mtls trusted ca tenant"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--use_mtls--trusted_ca--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls:trusted_ca", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "tenant", "scope_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls", "trusted_ca"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.views.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls", "trusted_ca", "tenant"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/use_mtls/trusted_ca/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["proxyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/)
- [dynamic_proxy.https_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/https_proxy/)
- [dynamic_proxy.https_proxy.tls_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/use_mtls/)
- dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-dynamic_proxy--https_proxy--tls_params--use_mtls--trusted_ca--name"></a>

### name property

Type: `"string"`. Computed.

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

<a id="schema-dynamic_proxy--https_proxy--tls_params--use_mtls--trusted_ca--namespace"></a>

### namespace property

Type: `"string"`. Computed.

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

<a id="schema-dynamic_proxy--https_proxy--tls_params--use_mtls--trusted_ca--tenant"></a>

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
