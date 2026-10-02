---
page_title: "waf_type.app_firewall.app_firewall"
subcategory: ""
description: "References to an Application Firewall configuration object."
xcsh_docs: {"aliases": ["waf type app firewall app firewall"], "body_bytes": 6557, "body_sha256": "sha256:1e3bfeef05da47c873ef90775c2d65c3a4d849cc24169b67c7102007906b14ad", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:waf_type:app_firewall:app_firewall", "parent_id": "xcsh-docs:resources:virtual_host:properties:waf_type:app_firewall", "path": "documentation/resources/virtual_host/properties/waf_type/app_firewall/app_firewall/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3002131032323211-1313323003223001-1030222330120001-3220101010122031-0213113222211303-3312312200323101-1132121102031200-3312230200131203", "registry_path": "docs/guides/resources--virtual_host--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_type", "app_firewall", "app_firewall"], "schema_version": 1, "sections": [{"aliases": ["kind"], "anchor": "schema-waf_type--app_firewall--app_firewall--kind", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. \"route\")", "document_id": "xcsh-docs:resources:virtual_host:properties:waf_type:app_firewall:app_firewall", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_type", "app_firewall", "app_firewall", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-waf_type--app_firewall--app_firewall--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:resources:virtual_host:properties:waf_type:app_firewall:app_firewall", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_type", "app_firewall", "app_firewall", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-waf_type--app_firewall--app_firewall--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:resources:virtual_host:properties:waf_type:app_firewall:app_firewall", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_type", "app_firewall", "app_firewall", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["tenant"], "anchor": "schema-waf_type--app_firewall--app_firewall--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:resources:virtual_host:properties:waf_type:app_firewall:app_firewall", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_type", "app_firewall", "app_firewall", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["uid"], "anchor": "schema-waf_type--app_firewall--app_firewall--uid", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. Route's) uid.", "document_id": "xcsh-docs:resources:virtual_host:properties:waf_type:app_firewall:app_firewall", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_type", "app_firewall", "app_firewall", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/waf_type/app_firewall/app_firewall/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "References to an Application Firewall configuration object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_type.app_firewall.app_firewall

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/)
- [waf_type.app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/app_firewall/)
- waf_type.app_firewall.app_firewall

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

<a id="schema-waf_type--app_firewall--app_firewall--kind"></a>

### kind property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-waf_type--app_firewall--app_firewall--name"></a>

### name property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-waf_type--app_firewall--app_firewall--namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="schema-waf_type--app_firewall--app_firewall--tenant"></a>

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
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-waf_type--app_firewall--app_firewall--uid"></a>

### uid property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [waf_type.app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/app_firewall/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
