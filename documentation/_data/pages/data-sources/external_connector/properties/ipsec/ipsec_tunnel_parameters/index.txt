---
page_title: "ipsec.ipsec_tunnel_parameters"
subcategory: ""
description: "ipsec.ipsec_tunnel_parameters for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 5105, "body_sha256": "sha256:1d2784adf7fbb2cdb3d051c73be53eb1cbea4adfc39056d2691d3a7a5132b6d7", "child_ids": ["xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:peer_ip_address", "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment", "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_inside_network", "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_network", "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:tunnel_eps"], "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "parent_id": "xcsh-docs:data-sources:external_connector:properties:ipsec", "path": "documentation/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/index.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["ipsec", "ipsec_tunnel_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ipsec.ipsec_tunnel_parameters for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
