---
page_title: "routes.waf_type.app_firewall.app_firewall"
subcategory: ""
description: "References to an Application Firewall configuration object."
xcsh_docs: {"aliases": ["routes waf type app firewall app firewall"], "body_bytes": 5788, "body_sha256": "sha256:813ccdf7de73dba08069a3916153ea778c8366bacd8ae50f66cdc1e3c2758717", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall:app_firewall", "parent_id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall", "path": "documentation/resources/route/properties/routes/waf_type/app_firewall/app_firewall/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2311111213330130-1003101232130330-3221320120310311-3233210300103003-3102313130023232-2302313001002221-0100030100023323-0310113023320310", "registry_path": "docs/guides/resources--route--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "waf_type", "app_firewall", "app_firewall"], "schema_version": 1, "sections": [{"aliases": ["routes waf type app firewall app firewall kind"], "anchor": "schema-routes--waf_type--app_firewall--app_firewall--kind", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. \"route\")", "document_id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall:app_firewall", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "kind", "scope_path": ["routes", "waf_type", "app_firewall", "app_firewall"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["routes", "waf_type", "app_firewall", "app_firewall", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes waf type app firewall app firewall name"], "anchor": "schema-routes--waf_type--app_firewall--app_firewall--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall:app_firewall", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "name", "scope_path": ["routes", "waf_type", "app_firewall", "app_firewall"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["routes", "waf_type", "app_firewall", "app_firewall", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes waf type app firewall app firewall namespace"], "anchor": "schema-routes--waf_type--app_firewall--app_firewall--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall:app_firewall", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "namespace", "scope_path": ["routes", "waf_type", "app_firewall", "app_firewall"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["routes", "waf_type", "app_firewall", "app_firewall", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes waf type app firewall app firewall tenant"], "anchor": "schema-routes--waf_type--app_firewall--app_firewall--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall:app_firewall", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "tenant", "scope_path": ["routes", "waf_type", "app_firewall", "app_firewall"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["routes", "waf_type", "app_firewall", "app_firewall", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes waf type app firewall app firewall uid"], "anchor": "schema-routes--waf_type--app_firewall--app_firewall--uid", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. Route's) uid.", "document_id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall:app_firewall", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "uid", "scope_path": ["routes", "waf_type", "app_firewall", "app_firewall"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["routes", "waf_type", "app_firewall", "app_firewall", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/waf_type/app_firewall/app_firewall/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "References to an Application Firewall configuration object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["routeCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.waf_type.app_firewall.app_firewall

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/)
- [routes.waf_type.app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/app_firewall/)
- routes.waf_type.app_firewall.app_firewall

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

References to an Application Firewall configuration object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1"
  }
}
```

Terraform syntax:

```terraform
app_firewall {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--waf_type--app_firewall--app_firewall--kind"></a>

### kind property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-routes--waf_type--app_firewall--app_firewall--name"></a>

### name property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-routes--waf_type--app_firewall--app_firewall--namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
  }
}
```

<a id="schema-routes--waf_type--app_firewall--app_firewall--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-routes--waf_type--app_firewall--app_firewall--uid"></a>

### uid property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
