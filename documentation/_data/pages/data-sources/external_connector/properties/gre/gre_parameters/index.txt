---
page_title: "gre.gre_parameters"
subcategory: ""
description: "GRE configuration parameters required for GRE Connection type."
xcsh_docs: {"aliases": ["gre gre parameters"], "body_bytes": 2738, "body_sha256": "sha256:601e84feb898e15e388593ae86a1654a339c8043b3fda31aee899d341941c766", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:peer_ip_address", "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:segment", "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:site_local_inside_network", "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:site_local_network", "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:tunnel_eps"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters", "parent_id": "xcsh-docs:data-sources:external_connector:properties:gre", "path": "documentation/data-sources/external_connector/properties/gre/gre_parameters/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-3111101010031023-1103010101130011-1021023300000320-0213031013130000-0312022120020323-2100231323133133-3213003201030022-3233200333220233", "registry_path": "docs/guides/data-sources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gre", "gre_parameters"], "schema_version": 1, "sections": [{"aliases": ["gre gre parameters peer ip address"], "anchor": "section", "description": "IPv4 Address in dot-decimal notation.", "document_id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:peer_ip_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["gre", "gre_parameters", "peer_ip_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["gre gre parameters segment"], "anchor": "section", "description": "Reference to Segment Object.", "document_id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:segment", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["gre", "gre_parameters", "segment"], "syntax": "attribute", "type": "object"}, {"aliases": ["gre gre parameters site local inside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:site_local_inside_network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gre", "gre_parameters", "site_local_inside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["gre gre parameters site local network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:site_local_network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gre", "gre_parameters", "site_local_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["gre gre parameters tunnel eps"], "anchor": "section", "description": "Configure tunnel parameters, source, destination, IP addresses.", "document_id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:tunnel_eps", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["gre", "gre_parameters", "tunnel_eps"], "syntax": "attribute", "type": "object"}, {"aliases": ["gre gre parameters tunnel mtu"], "anchor": "schema-gre--gre_parameters--tunnel_mtu", "description": "Configure MTU for the GRE tunnel interface.", "document_id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gre", "gre_parameters", "tunnel_mtu"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/gre/gre_parameters/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "GRE configuration parameters required for GRE Connection type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["external_connectorCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gre.gre_parameters

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/)
- [gre](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/)
- gre.gre_parameters

<a id="section"></a>

Type: `"single"`. Computed.

GRE configuration parameters required for GRE Connection type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tunnel_source_vn": "[\"segment\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

## Direct properties

- [peer_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/peer_ip_address/): complete subsection reference.

- [segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/segment/): complete subsection reference.

- [site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/site_local_inside_network/): complete subsection reference.

- [site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/site_local_network/): complete subsection reference.

- [tunnel_eps](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/tunnel_eps/): complete subsection reference.

<a id="schema-gre--gre_parameters--tunnel_mtu"></a>

### tunnel_mtu property

Type: `"number"`. Computed.

Configure MTU for the GRE tunnel interface.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1370,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 512
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "512",
    "ves.io.schema.rules.uint32.lte": "1370"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "512",
    "ves.io.schema.rules.uint32.lte": "1370"
  }
}
```
