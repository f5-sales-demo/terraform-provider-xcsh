---
page_title: "gre.gre_parameters.peer_ip_address"
subcategory: ""
description: "IPv4 Address in dot-decimal notation."
xcsh_docs: {"aliases": ["gre gre parameters peer ip address"], "body_bytes": 1780, "body_sha256": "sha256:8e22f885b954383fd5345a01ad96e2f9135effddb346be437a128d88612c7bd7", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:peer_ip_address", "parent_id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters", "path": "documentation/data-sources/external_connector/properties/gre/gre_parameters/peer_ip_address/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-0232102201201332-1201210011113212-0123133032101020-1000012321300233-0113022130223121-3221103223321100-1102201211003021-2120121321222231", "registry_path": "docs/guides/data-sources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gre", "gre_parameters", "peer_ip_address"], "schema_version": 1, "sections": [{"aliases": ["gre gre parameters peer ip address addr"], "anchor": "schema-gre--gre_parameters--peer_ip_address--addr", "description": "IPv4 Address in string form with dot-decimal notation.", "document_id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:peer_ip_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gre", "gre_parameters", "peer_ip_address", "addr"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/gre/gre_parameters/peer_ip_address/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "IPv4 Address in dot-decimal notation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["external_connectorCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gre.gre_parameters.peer_ip_address

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/)
- [gre](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/)
- [gre.gre_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/)
- gre.gre_parameters.peer_ip_address

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-gre--gre_parameters--peer_ip_address--addr"></a>

### addr property

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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
      "validatedAt": "2026-10-16T10:31:02+00:00"
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
