---
page_title: "virtual_server.http3.http_server_profile"
subcategory: ""
description: "Configuration parameter for http server profile"
xcsh_docs: {"aliases": ["virtual server http3 http server profile"], "body_bytes": 6033, "body_sha256": "sha256:6baec803642bad542d7107119a4d72d96428e648a585c3294ca64d3506dd7865", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:http_server_profile", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3", "path": "documentation/resources/application_profiles/properties/virtual_server/http3/http_server_profile/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0220231220211221-1300100113032203-1202030023031212-3101012201312312-0221100301100212-0101010110231132-1003301322301023-0020200203102112", "registry_path": "docs/guides/resources--application_profiles--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "http3", "http_server_profile"], "schema_version": 1, "sections": [{"aliases": ["virtual server http3 http server profile kind"], "anchor": "schema-virtual_server--http3--http_server_profile--kind", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. \"route\")", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:http_server_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "kind", "scope_path": ["virtual_server", "http3", "http_server_profile"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["virtual_server", "http3", "http_server_profile", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["virtual server http3 http server profile name"], "anchor": "schema-virtual_server--http3--http_server_profile--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:http_server_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "name", "scope_path": ["virtual_server", "http3", "http_server_profile"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["virtual_server", "http3", "http_server_profile", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["virtual server http3 http server profile namespace"], "anchor": "schema-virtual_server--http3--http_server_profile--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:http_server_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "namespace", "scope_path": ["virtual_server", "http3", "http_server_profile"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["virtual_server", "http3", "http_server_profile", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["virtual server http3 http server profile tenant"], "anchor": "schema-virtual_server--http3--http_server_profile--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:http_server_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "tenant", "scope_path": ["virtual_server", "http3", "http_server_profile"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["virtual_server", "http3", "http_server_profile", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["virtual server http3 http server profile uid"], "anchor": "schema-virtual_server--http3--http_server_profile--uid", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. Route's) uid.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:http_server_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "uid", "scope_path": ["virtual_server", "http3", "http_server_profile"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["virtual_server", "http3", "http_server_profile", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/http3/http_server_profile/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Configuration parameter for http server profile", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["application_profilesCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.http3.http_server_profile

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [virtual_server.http3](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/)
- virtual_server.http3.http_server_profile

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for http server profile.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http_server_profile {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-virtual_server--http3--http_server_profile--kind"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-virtual_server--http3--http_server_profile--name"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-virtual_server--http3--http_server_profile--namespace"></a>

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
  }
}
```

<a id="schema-virtual_server--http3--http_server_profile--tenant"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-virtual_server--http3--http_server_profile--uid"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
