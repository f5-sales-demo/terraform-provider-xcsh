---
page_title: "webhook.http_config.client_cert_obj.use_tls_obj"
subcategory: ""
description: "Reference to client certificate object."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "webhook http config client cert obj use tls obj"], "body_bytes": 7104, "body_sha256": "sha256:0778cf13a10d85ecfac119c236ba35ab640c9ebd3fcaa67df178a125f2359e8b", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj:use_tls_obj", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj", "path": "documentation/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/use_tls_obj/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3313133313133231-3002013323102012-0102323202333012-2000222010000012-1123202321221331-2310320010313032-3320121212101102-3223023133132310", "registry_path": "docs/guides/resources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["webhook", "http_config", "client_cert_obj", "use_tls_obj"], "schema_version": 1, "sections": [{"aliases": ["kind"], "anchor": "schema-webhook--http_config--client_cert_obj--use_tls_obj--kind", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. \"route\")", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj:use_tls_obj", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "client_cert_obj", "use_tls_obj", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-webhook--http_config--client_cert_obj--use_tls_obj--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj:use_tls_obj", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "client_cert_obj", "use_tls_obj", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-webhook--http_config--client_cert_obj--use_tls_obj--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj:use_tls_obj", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "client_cert_obj", "use_tls_obj", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["tenant"], "anchor": "schema-webhook--http_config--client_cert_obj--use_tls_obj--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj:use_tls_obj", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "client_cert_obj", "use_tls_obj", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["uid"], "anchor": "schema-webhook--http_config--client_cert_obj--use_tls_obj--uid", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. Route's) uid.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj:use_tls_obj", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "client_cert_obj", "use_tls_obj", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/use_tls_obj/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Reference to client certificate object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.client_cert_obj.use_tls_obj

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/)
- [webhook.http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/)
- [webhook.http_config.client_cert_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/)
- webhook.http_config.client_cert_obj.use_tls_obj

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Certificate Object. Reference to client certificate object.

Upstream description:

Reference to client certificate object.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
use_tls_obj {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-webhook--http_config--client_cert_obj--use_tls_obj--kind"></a>

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

<a id="schema-webhook--http_config--client_cert_obj--use_tls_obj--name"></a>

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

<a id="schema-webhook--http_config--client_cert_obj--use_tls_obj--namespace"></a>

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

<a id="schema-webhook--http_config--client_cert_obj--use_tls_obj--tenant"></a>

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

<a id="schema-webhook--http_config--client_cert_obj--use_tls_obj--uid"></a>

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

- [webhook.http_config.client_cert_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
