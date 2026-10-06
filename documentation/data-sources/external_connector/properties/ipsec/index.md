---
page_title: "ipsec"
subcategory: ""
description: "External Connector with IPsec tunnel."
xcsh_docs: {"aliases": ["ipsec"], "body_bytes": 1032, "body_sha256": "sha256:c2c46748a8eadff1c27fc19131d8b4d4540892c61fc32b28e4dd7aadbdd76b90", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters", "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:ipsec", "parent_id": "xcsh-docs:data-sources:external_connector:reference", "path": "documentation/data-sources/external_connector/properties/ipsec/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233", "registry_path": "docs/guides/data-sources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ipsec"], "schema_version": 1, "sections": [{"aliases": ["ipsec ike parameters"], "anchor": "section", "description": "IKE configuration parameters required for IPsec Connection type.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ipsec", "ike_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ipsec tunnel parameters"], "anchor": "section", "description": "In this section, we will configure the tunnel parameters, source, destination, IP addresses, and segment.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/ipsec/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "External Connector with IPsec tunnel.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/)
- ipsec

<a id="section"></a>

Type: `"single"`. Computed.

IPsec. External Connector with IPsec tunnel.

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

- [ike_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/): complete subsection reference.

- [ipsec_tunnel_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/): complete subsection reference.
