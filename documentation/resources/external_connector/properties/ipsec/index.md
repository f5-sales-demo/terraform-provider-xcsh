---
page_title: "ipsec"
subcategory: ""
description: "External Connector with IPsec tunnel."
xcsh_docs: {"aliases": ["ipsec"], "body_bytes": 1754, "body_sha256": "sha256:cd4bef3d94998a297a916c47446873e2913e321cc2c68dc1248c6458829a267f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters", "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:ipsec", "parent_id": "xcsh-docs:resources:external_connector:reference", "path": "documentation/resources/external_connector/properties/ipsec/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133", "registry_path": "docs/guides/resources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ipsec"], "schema_version": 1, "sections": [{"aliases": ["ike parameters"], "anchor": "section", "description": "IKE configuration parameters required for IPsec Connection type.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ipsec--ike_parameters--rm_hostname", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:rm_hostname,rm_ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters", "type": "conflicts"}, {"anchor": "schema-ipsec--ike_parameters--rm_hostname", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:rm_hostname,use_default_remote_ike_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:dpd_disabled,dpd_keep_alive_timer", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:dpd_disabled,dpd_keep_alive_timer", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:initiator,responder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:initiator", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:initiator,responder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:responder", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:rm_hostname,rm_ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:rm_ip_address,use_default_remote_ike_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:rm_hostname,use_default_remote_ike_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:use_default_remote_ike_id", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:rm_ip_address,use_default_remote_ike_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:use_default_remote_ike_id", "type": "conflicts"}], "schema_path": ["ipsec", "ike_parameters"], "syntax": "block", "type": "object"}, {"aliases": ["ipsec tunnel parameters"], "anchor": "section", "description": "In this section, we will configure the tunnel parameters, source, destination, IP addresses, and segment.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters:ConflictingObjectAttributes:segment,site_local_inside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters:ConflictingObjectAttributes:segment,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters:ConflictingObjectAttributes:segment,site_local_inside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters:ConflictingObjectAttributes:site_local_inside_network,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters:ConflictingObjectAttributes:segment,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters:ConflictingObjectAttributes:site_local_inside_network,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:site_local_network", "type": "conflicts"}, {"anchor": "schema-ipsec--ipsec_tunnel_parameters--psk", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters:RequiredObjectAttributes:psk,tunnel_eps,tunnel_mtu", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "type": "requires"}, {"anchor": "schema-ipsec--ipsec_tunnel_parameters--tunnel_mtu", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters:RequiredObjectAttributes:psk,tunnel_eps,tunnel_mtu", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ipsec_tunnel_parameters:RequiredObjectAttributes:psk,tunnel_eps,tunnel_mtu", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:tunnel_eps", "type": "requires"}], "schema_path": ["ipsec", "ipsec_tunnel_parameters"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/ipsec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "External Connector with IPsec tunnel.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/)
- ipsec

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPsec. External Connector with IPsec tunnel.

Upstream description:

External Connector with IPsec tunnel.

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
ipsec {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ike_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/): complete subsection reference.

- [ipsec_tunnel_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/): complete subsection reference.

## Next pages

- [ipsec.ike_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/)
- [ipsec.ipsec_tunnel_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/)
- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
