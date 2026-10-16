---
page_title: "ipsec.ipsec_tunnel_parameters.peer_ip_address"
subcategory: ""
description: "IPv4 Address in dot-decimal notation."
xcsh_docs: {"aliases": ["ipsec ipsec tunnel parameters peer ip address"], "body_bytes": 2150, "body_sha256": "sha256:fee2069ef70cfc5ef5b13f0ff17bb718694ad45fc90a42b2a21efdab252f3a36", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:peer_ip_address", "parent_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "path": "documentation/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/peer_ip_address/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-3021301113200132-1210020301222122-0132011012332233-0303100330222213-3111333010321220-0302310301201301-0332113022313333-0202300202033212", "registry_path": "docs/guides/resources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ipsec", "ipsec_tunnel_parameters", "peer_ip_address"], "schema_version": 1, "sections": [{"aliases": ["ipsec ipsec tunnel parameters peer ip address addr"], "anchor": "schema-ipsec--ipsec_tunnel_parameters--peer_ip_address--addr", "description": "IPv4 Address in string form with dot-decimal notation.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:peer_ip_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "peer_ip_address", "addr"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/peer_ip_address/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "IPv4 Address in dot-decimal notation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["external_connectorCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ipsec_tunnel_parameters.peer_ip_address

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/)
- [ipsec.ipsec_tunnel_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/)
- ipsec.ipsec_tunnel_parameters.peer_ip_address

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPv4 Address. IPv4 Address in dot-decimal notation.

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
peer_ip_address {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ipsec--ipsec_tunnel_parameters--peer_ip_address--addr"></a>

### addr property

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-15T10:49:30+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```
