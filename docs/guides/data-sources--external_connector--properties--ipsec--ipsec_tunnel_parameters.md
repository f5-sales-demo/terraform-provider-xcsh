---
page_title: "ipsec.ipsec_tunnel_parameters"
subcategory: ""
description: "ipsec.ipsec_tunnel_parameters for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 4358, "body_sha256": "sha256:d20ab7e98cded2683ce1c68e8970a1128b8236122c4e19ae3b0ee24674bc79c7", "canonical_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "child_ids": ["xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:peer_ip_address", "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment", "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_inside_network", "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_network", "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:tunnel_eps"], "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "parent_id": "xcsh-docs:data-sources:external_connector:properties:ipsec", "path": "docs/guides/data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ipsec", "ipsec_tunnel_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ipsec.ipsec_tunnel_parameters for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ipsec_tunnel_parameters

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md)
- [Property reference](data-sources--external_connector--reference.md)
- [ipsec](data-sources--external_connector--properties--ipsec.md)
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

- [peer_ip_address](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--peer_ip_address.md): complete subsection reference.

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

- [segment](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--segment.md): complete subsection reference.

- [site_local_inside_network](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--site_local_inside_network.md): complete subsection reference.

- [site_local_network](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--site_local_network.md): complete subsection reference.

- [tunnel_eps](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--tunnel_eps.md): complete subsection reference.

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

- [ipsec.ipsec_tunnel_parameters.peer_ip_address](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--peer_ip_address.md)
- [ipsec.ipsec_tunnel_parameters.segment](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--segment.md)
- [ipsec.ipsec_tunnel_parameters.site_local_inside_network](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--site_local_inside_network.md)
- [ipsec.ipsec_tunnel_parameters.site_local_network](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--site_local_network.md)
- [ipsec.ipsec_tunnel_parameters.tunnel_eps](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--tunnel_eps.md)
- [ipsec](data-sources--external_connector--properties--ipsec.md)
- [xcsh_external_connector](../data-sources/external_connector.md)
