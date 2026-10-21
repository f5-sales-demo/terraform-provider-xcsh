---
page_title: "ipsec.ipsec_tunnel_parameters.peer_ip_address"
subcategory: ""
description: "IPv4 Address in dot-decimal notation."
xcsh_docs: {"aliases": ["ipsec ipsec tunnel parameters peer ip address"], "body_bytes": 1839, "body_sha256": "sha256:16eb67418a027c84670071bd72edfbd050c94a404f1afd8969befcf4d6264aca", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:peer_ip_address", "parent_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "path": "documentation/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/peer_ip_address/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-2020200020320011-3022011320110023-0112111022113212-0122001013120221-0202002333212101-1013010210133122-3332203113022012-1121332120022011", "registry_path": "docs/guides/data-sources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ipsec", "ipsec_tunnel_parameters", "peer_ip_address"], "schema_version": 1, "sections": [{"aliases": ["ipsec ipsec tunnel parameters peer ip address addr"], "anchor": "schema-ipsec--ipsec_tunnel_parameters--peer_ip_address--addr", "description": "IPv4 Address in string form with dot-decimal notation.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:peer_ip_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "peer_ip_address", "addr"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/peer_ip_address/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "IPv4 Address in dot-decimal notation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["external_connectorCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ipsec_tunnel_parameters.peer_ip_address

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/)
- [ipsec.ipsec_tunnel_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/)
- ipsec.ipsec_tunnel_parameters.peer_ip_address

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

<a id="schema-ipsec--ipsec_tunnel_parameters--peer_ip_address--addr"></a>

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
      "validatedAt": "2026-10-21T11:00:09+00:00"
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
