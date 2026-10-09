---
page_title: "caching_policy.custom_cache_rule.cdn_cache_rules"
subcategory: "Load Balancing"
description: "Reference to CDN Cache Rule configuration object."
xcsh_docs: {"aliases": ["caching policy custom cache rule cdn cache rules"], "body_bytes": 5440, "body_sha256": "sha256:53053669b968ad6712f7de75a060766ed028280e1db18ddec8e0d237d5732a31", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:custom_cache_rule:cdn_cache_rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:custom_cache_rule", "path": "documentation/resources/http_loadbalancer/properties/caching_policy/custom_cache_rule/cdn_cache_rules/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0320220112122003-3200020033122332-1301313302000102-1122133000313102-1020022313003231-1322102210203301-3010230133213232-3003132230022032", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-013.md", "relationships": [{"anchor": "schema-caching_policy--custom_cache_rule--cdn_cache_rules--name", "enforcement": "provider-schema", "group": "caching_policy.custom_cache_rule.cdn_cache_rules:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:custom_cache_rule:cdn_cache_rules", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["caching_policy", "custom_cache_rule", "cdn_cache_rules"], "schema_version": 1, "sections": [{"aliases": ["caching policy custom cache rule cdn cache rules name"], "anchor": "schema-caching_policy--custom_cache_rule--cdn_cache_rules--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:custom_cache_rule:cdn_cache_rules", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "name", "scope_path": ["caching_policy", "custom_cache_rule", "cdn_cache_rules"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.views.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["caching_policy", "custom_cache_rule", "cdn_cache_rules", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["caching policy custom cache rule cdn cache rules namespace"], "anchor": "schema-caching_policy--custom_cache_rule--cdn_cache_rules--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:custom_cache_rule:cdn_cache_rules", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "namespace", "scope_path": ["caching_policy", "custom_cache_rule", "cdn_cache_rules"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.views.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["caching_policy", "custom_cache_rule", "cdn_cache_rules", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["caching policy custom cache rule cdn cache rules tenant"], "anchor": "schema-caching_policy--custom_cache_rule--cdn_cache_rules--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:custom_cache_rule:cdn_cache_rules", "enum_extraction_complete": true, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "tenant", "scope_path": ["caching_policy", "custom_cache_rule", "cdn_cache_rules"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.views.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["caching_policy", "custom_cache_rule", "cdn_cache_rules", "tenant"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/caching_policy/custom_cache_rule/cdn_cache_rules/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Reference to CDN Cache Rule configuration object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# caching_policy.custom_cache_rule.cdn_cache_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [caching_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/caching_policy/)
- [caching_policy.custom_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/caching_policy/custom_cache_rule/)
- caching_policy.custom_cache_rule.cdn_cache_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Reference to CDN Cache Rule configuration object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
cdn_cache_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-caching_policy--custom_cache_rule--cdn_cache_rules--name"></a>

### name property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="schema-caching_policy--custom_cache_rule--cdn_cache_rules--namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="schema-caching_policy--custom_cache_rule--cdn_cache_rules--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
