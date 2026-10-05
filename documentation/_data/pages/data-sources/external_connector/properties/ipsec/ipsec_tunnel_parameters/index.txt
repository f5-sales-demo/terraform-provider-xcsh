---
page_title: "ipsec.ipsec_tunnel_parameters"
subcategory: ""
description: "In this section, we will configure the tunnel parameters, source, destination, IP addresses, and segment."
xcsh_docs: {"aliases": ["ipsec ipsec tunnel parameters"], "body_bytes": 5105, "body_sha256": "sha256:fef727dcb6dd80c6b023d2240fe419d718cbbba15ac8b5fb1a461d30f0725850", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:peer_ip_address", "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment", "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_inside_network", "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_network", "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:tunnel_eps"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "parent_id": "xcsh-docs:data-sources:external_connector:properties:ipsec", "path": "documentation/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1203031111130212-3100200313212112-2011133112022013-0223221300313313-3333012221300300-3130100330021102-2012202002001221-0023322303313133", "registry_path": "docs/guides/data-sources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ipsec", "ipsec_tunnel_parameters"], "schema_version": 1, "sections": [{"aliases": ["ipsec ipsec tunnel parameters peer ip address"], "anchor": "section", "description": "IPv4 Address in dot-decimal notation.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:peer_ip_address", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "peer_ip_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ipsec tunnel parameters psk"], "anchor": "schema-ipsec--ipsec_tunnel_parameters--psk", "description": "The IKE pre-shared key (PSK) is required to ensure the IKE peers can authenticate one another within IKE phase 1 negotiation.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "psk"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipsec ipsec tunnel parameters segment"], "anchor": "section", "description": "Reference to Segment Object.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "segment"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ipsec tunnel parameters site local inside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_inside_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "site_local_inside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ipsec tunnel parameters site local network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "site_local_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ipsec tunnel parameters tunnel eps"], "anchor": "section", "description": "Configure tunnel parameters, local and remote IP addresses.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:tunnel_eps", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "tunnel_eps"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ipsec tunnel parameters tunnel mtu"], "anchor": "schema-ipsec--ipsec_tunnel_parameters--tunnel_mtu", "description": "The tunnel MTU defines the maximum size of the packet that can be sent through the tunnel without needing to be fragmented.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "tunnel_mtu"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "In this section, we will configure the tunnel parameters, source, destination, IP addresses, and segment.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ipsec_tunnel_parameters

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/)
- ipsec.ipsec_tunnel_parameters

<a id="section"></a>

Type: `"single"`. Computed.

In this section, we will configure the tunnel parameters, source, destination, IP addresses, and
segment.

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

- [peer_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/peer_ip_address/): complete subsection reference.

<a id="schema-ipsec--ipsec_tunnel_parameters--psk"></a>

### psk property

Type: `"string"`. Computed.

The IKE pre-shared key (PSK) is required to ensure the IKE peers can authenticate one another within
IKE phase 1 negotiation.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/): complete subsection reference.

- [site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/site_local_inside_network/): complete subsection reference.

- [site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/site_local_network/): complete subsection reference.

- [tunnel_eps](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/tunnel_eps/): complete subsection reference.

<a id="schema-ipsec--ipsec_tunnel_parameters--tunnel_mtu"></a>

### tunnel_mtu property

Type: `"number"`. Computed.

The tunnel MTU defines the maximum size of the packet that can be sent through the tunnel without
needing to be fragmented.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "512",
    "ves.io.schema.rules.uint32.lte": "1370"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "512",
    "ves.io.schema.rules.uint32.lte": "1370"
  }
}
```

## Next pages

- [ipsec.ipsec_tunnel_parameters.peer_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/peer_ip_address/)
- [ipsec.ipsec_tunnel_parameters.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/)
- [ipsec.ipsec_tunnel_parameters.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/site_local_inside_network/)
- [ipsec.ipsec_tunnel_parameters.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/site_local_network/)
- [ipsec.ipsec_tunnel_parameters.tunnel_eps](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/tunnel_eps/)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/)
- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
