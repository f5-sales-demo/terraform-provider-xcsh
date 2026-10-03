---
page_title: "ipsec.ipsec_tunnel_parameters.segment.refs"
subcategory: ""
description: "Reference to Segment Object."
xcsh_docs: {"aliases": ["ipsec ipsec tunnel parameters segment refs"], "body_bytes": 7191, "body_sha256": "sha256:138964ce4b06069190a758e9154c6ead97481062f96d6aa6bcd1ad5def3fe466", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment:refs", "parent_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment", "path": "documentation/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/refs/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1003032313213122-0230220010230120-2331111030321011-2303003103110013-3300012130130110-3131021102332200-1130312332303122-2122002111022321", "registry_path": "docs/guides/resources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ipsec", "ipsec_tunnel_parameters", "segment", "refs"], "schema_version": 1, "sections": [{"aliases": ["ipsec ipsec tunnel parameters segment refs kind"], "anchor": "schema-ipsec--ipsec_tunnel_parameters--segment--refs--kind", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. \"route\")", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment:refs", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "segment", "refs", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipsec ipsec tunnel parameters segment refs name"], "anchor": "schema-ipsec--ipsec_tunnel_parameters--segment--refs--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment:refs", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "segment", "refs", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipsec ipsec tunnel parameters segment refs namespace"], "anchor": "schema-ipsec--ipsec_tunnel_parameters--segment--refs--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment:refs", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "segment", "refs", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipsec ipsec tunnel parameters segment refs tenant"], "anchor": "schema-ipsec--ipsec_tunnel_parameters--segment--refs--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment:refs", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "segment", "refs", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipsec ipsec tunnel parameters segment refs uid"], "anchor": "schema-ipsec--ipsec_tunnel_parameters--segment--refs--uid", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. Route's) uid.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment:refs", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "segment", "refs", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/refs/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "Reference to Segment Object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["external_connectorCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ipsec_tunnel_parameters.segment.refs

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/)
- [ipsec.ipsec_tunnel_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/)
- [ipsec.ipsec_tunnel_parameters.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/)
- ipsec.ipsec_tunnel_parameters.segment.refs

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Segment. Reference to Segment Object.

Upstream description:

Reference to Segment Object.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
refs {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ipsec--ipsec_tunnel_parameters--segment--refs--kind"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-ipsec--ipsec_tunnel_parameters--segment--refs--name"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-ipsec--ipsec_tunnel_parameters--segment--refs--namespace"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-ipsec--ipsec_tunnel_parameters--segment--refs--tenant"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-ipsec--ipsec_tunnel_parameters--segment--refs--uid"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [ipsec.ipsec_tunnel_parameters.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/)
- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
