---
page_title: "ipsec.ipsec_tunnel_parameters"
subcategory: ""
description: "In this section, we will configure the tunnel parameters, source, destination, IP addresses, and segment."
xcsh_docs: {"aliases": ["ipsec ipsec tunnel parameters"], "body_bytes": 4614, "body_sha256": "sha256:93f4fe1b6de02b634ef33d024b97e6a0b356c670796543b1415d4b8c11c776a2", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:peer_ip_address", "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment", "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_inside_network", "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_network", "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:tunnel_eps"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "parent_id": "xcsh-docs:resources:external_connector:properties:ipsec", "path": "documentation/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1111312210033110-1323001002020013-0301303102011303-3323022223203100-3130333022311330-3033000102231322-1330001332110221-0323222300300120", "registry_path": "docs/guides/resources--external_connector--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters:ConflictingObjectAttributes:segment,site_local_inside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters:ConflictingObjectAttributes:segment,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters:ConflictingObjectAttributes:segment,site_local_inside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters:ConflictingObjectAttributes:site_local_inside_network,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters:ConflictingObjectAttributes:segment,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters:ConflictingObjectAttributes:site_local_inside_network,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_network", "type": "conflicts"}, {"anchor": "schema-ipsec--ipsec_tunnel_parameters--psk", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters:RequiredObjectAttributes:psk,tunnel_eps,tunnel_mtu", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "type": "requires"}, {"anchor": "schema-ipsec--ipsec_tunnel_parameters--tunnel_mtu", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters:RequiredObjectAttributes:psk,tunnel_eps,tunnel_mtu", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters:RequiredObjectAttributes:psk,tunnel_eps,tunnel_mtu", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:tunnel_eps", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ipsec", "ipsec_tunnel_parameters"], "schema_version": 1, "sections": [{"aliases": ["ipsec ipsec tunnel parameters peer ip address"], "anchor": "section", "description": "IPv4 Address in dot-decimal notation.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:peer_ip_address", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "peer_ip_address"], "syntax": "block", "type": "object"}, {"aliases": ["ipsec ipsec tunnel parameters psk"], "anchor": "schema-ipsec--ipsec_tunnel_parameters--psk", "description": "The IKE pre-shared key (PSK) is required to ensure the IKE peers can authenticate one another within IKE phase 1 negotiation.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "psk"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipsec ipsec tunnel parameters segment"], "anchor": "section", "description": "Reference to Segment Object.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters.segment:RequiredObjectAttributes:refs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment:refs", "type": "requires"}], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "segment"], "syntax": "block", "type": "object"}, {"aliases": ["ipsec ipsec tunnel parameters site local inside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_inside_network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "site_local_inside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ipsec tunnel parameters site local network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "site_local_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ipsec tunnel parameters tunnel eps"], "anchor": "section", "description": "Configure tunnel parameters, local and remote IP addresses.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:tunnel_eps", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--interface", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters.tunnel_eps:RequiredListObjectAttributes:interface,local_tunnel_ip,node,remote_tunnel_ip", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:tunnel_eps", "type": "requires"}, {"anchor": "schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--local_tunnel_ip", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters.tunnel_eps:RequiredListObjectAttributes:interface,local_tunnel_ip,node,remote_tunnel_ip", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:tunnel_eps", "type": "requires"}, {"anchor": "schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--node", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters.tunnel_eps:RequiredListObjectAttributes:interface,local_tunnel_ip,node,remote_tunnel_ip", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:tunnel_eps", "type": "requires"}, {"anchor": "schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--remote_tunnel_ip", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters.tunnel_eps:RequiredListObjectAttributes:interface,local_tunnel_ip,node,remote_tunnel_ip", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:tunnel_eps", "type": "requires"}], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "tunnel_eps"], "syntax": "block", "type": "object"}, {"aliases": ["ipsec ipsec tunnel parameters tunnel mtu"], "anchor": "schema-ipsec--ipsec_tunnel_parameters--tunnel_mtu", "description": "The tunnel MTU defines the maximum size of the packet that can be sent through the tunnel without needing to be fragmented.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "tunnel_mtu"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "In this section, we will configure the tunnel parameters, source, destination, IP addresses, and segment.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["external_connectorCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ipsec_tunnel_parameters

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/)
- ipsec.ipsec_tunnel_parameters

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

In this section, we will configure the tunnel parameters, source, destination, IP addresses, and
segment.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("psk",
    "tunnel_eps",
    "tunnel_mtu"),
  validators.ConflictingObjectAttributes("segment",
    "site_local_inside_network"),
  validators.ConflictingObjectAttributes("segment",
    "site_local_network"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
```

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

Terraform syntax:

```terraform
ipsec_tunnel_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

- [peer_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/peer_ip_address/): complete subsection reference.

<a id="schema-ipsec--ipsec_tunnel_parameters--psk"></a>

### psk property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/): complete subsection reference.

- [site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/site_local_inside_network/): complete subsection reference.

- [site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/site_local_network/): complete subsection reference.

- [tunnel_eps](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/tunnel_eps/): complete subsection reference.

<a id="schema-ipsec--ipsec_tunnel_parameters--tunnel_mtu"></a>

### tunnel_mtu property

Type: `"number"`. Optional.

The tunnel MTU defines the maximum size of the packet that can be sent through the tunnel without
needing to be fragmented.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(512, 1370),
}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
