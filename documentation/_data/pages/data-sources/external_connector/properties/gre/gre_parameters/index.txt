---
page_title: "gre.gre_parameters"
subcategory: ""
description: "GRE configuration parameters required for GRE Connection type."
xcsh_docs: {"aliases": ["gre gre parameters"], "body_bytes": 3860, "body_sha256": "sha256:e044ba6b9c3a62e07faeb514395ef1b05a677eb60b5acfcad21e8528d731df25", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:peer_ip_address", "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:segment", "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:site_local_inside_network", "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:site_local_network", "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:tunnel_eps"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters", "parent_id": "xcsh-docs:data-sources:external_connector:properties:gre", "path": "documentation/data-sources/external_connector/properties/gre/gre_parameters/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3111101010031023-1103010101130011-1021023300000320-0213031013130000-0312022120020323-2100231323133133-3213003201030022-3233200333220233", "registry_path": "docs/guides/data-sources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gre", "gre_parameters"], "schema_version": 1, "sections": [{"aliases": ["gre gre parameters peer ip address"], "anchor": "section", "description": "IPv4 Address in dot-decimal notation.", "document_id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:peer_ip_address", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["gre", "gre_parameters", "peer_ip_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["gre gre parameters segment"], "anchor": "section", "description": "Reference to Segment Object.", "document_id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:segment", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["gre", "gre_parameters", "segment"], "syntax": "attribute", "type": "object"}, {"aliases": ["gre gre parameters site local inside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:site_local_inside_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gre", "gre_parameters", "site_local_inside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["gre gre parameters site local network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:site_local_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gre", "gre_parameters", "site_local_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["gre gre parameters tunnel eps"], "anchor": "section", "description": "Configure tunnel parameters, source, destination, IP addresses.", "document_id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:tunnel_eps", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["gre", "gre_parameters", "tunnel_eps"], "syntax": "attribute", "type": "object"}, {"aliases": ["gre gre parameters tunnel mtu"], "anchor": "schema-gre--gre_parameters--tunnel_mtu", "description": "Configure MTU for the GRE tunnel interface.", "document_id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gre", "gre_parameters", "tunnel_mtu"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/gre/gre_parameters/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "GRE configuration parameters required for GRE Connection type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["external_connectorCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [gre.gre_parameters.peer_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/peer_ip_address/)
- [gre.gre_parameters.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/segment/)
- [gre.gre_parameters.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/site_local_inside_network/)
- [gre.gre_parameters.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/site_local_network/)
- [gre.gre_parameters.tunnel_eps](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/tunnel_eps/)
- [gre](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/)
- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
