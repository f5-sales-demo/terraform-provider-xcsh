---
page_title: "private_adn"
subcategory: ""
description: "X-required Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud support."
xcsh_docs: {"aliases": ["private adn"], "body_bytes": 2349, "body_sha256": "sha256:3570d7b6ecac6ae8a2d331aed535538d1d1748fe1b2e0ff60b5ac4dce58505ce", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:private_adn", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/private_adn/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-1132333020000033-3121313020113130-1113011323131231-1303130201313202-3331202201110330-3232030020101012-0011123211010231-1301321320123330", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-016.md", "relationships": [{"anchor": "schema-private_adn--private_adn", "enforcement": "provider-schema", "group": "private_adn:RequiredObjectAttributes:private_adn", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:private_adn", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["private_adn"], "schema_version": 1, "sections": [{"aliases": ["private adn private adn"], "anchor": "schema-private_adn--private_adn", "description": "Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud support.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:private_adn", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["private_adn", "private_adn"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/private_adn/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "X-required Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud support.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# private_adn

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- private_adn

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

X-required Establish private connectivity with the F5 Distributed Cloud Global Network using a
Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud
support.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("private_adn")}
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
private_adn {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-private_adn--private_adn"></a>

### private_adn property

Type: `"string"`. Optional.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

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
      "validatedAt": "2026-10-15T10:49:30+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```
