---
page_title: "ipsec.ike_parameters"
subcategory: ""
description: "IKE configuration parameters required for IPsec Connection type."
xcsh_docs: {"aliases": ["ipsec ike parameters"], "body_bytes": 3532, "body_sha256": "sha256:8b979450423570d32cd88791488a77fe6e21fbc8beabe49490bbaea3d14352d0", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:dpd_disabled", "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:ike_phase1_profile", "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:ike_phase2_profile", "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:initiator", "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:responder", "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:use_default_local_ike_id", "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:use_default_remote_ike_id"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters", "parent_id": "xcsh-docs:data-sources:external_connector:properties:ipsec", "path": "documentation/data-sources/external_connector/properties/ipsec/ike_parameters/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1330002233213122-1233303203212331-0103313203231320-0223301331122010-3123211313221101-1303311322332213-2133122311330211-0022223103030223", "registry_path": "docs/guides/data-sources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ipsec", "ike_parameters"], "schema_version": 1, "sections": [{"aliases": ["ipsec ike parameters dpd disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:dpd_disabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "dpd_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ike parameters dpd keep alive timer"], "anchor": "section", "description": "Configuration parameter for dpd keep alive timer.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ipsec", "ike_parameters", "dpd_keep_alive_timer"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ike parameters ike phase1 profile"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:ike_phase1_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ipsec", "ike_parameters", "ike_phase1_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ike parameters ike phase2 profile"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:ike_phase2_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ipsec", "ike_parameters", "ike_phase2_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ike parameters initiator"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:initiator", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "initiator"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ike parameters responder"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:responder", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "responder"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ike parameters rm hostname"], "anchor": "schema-ipsec--ike_parameters--rm_hostname", "description": "Exclusive with Configure an hostname Remote IKE ID.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "rm_hostname"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipsec ike parameters rm ip address"], "anchor": "section", "description": "IP Address used to specify an IPv4 or IPv6 address.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ipsec", "ike_parameters", "rm_ip_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ike parameters use default local ike id"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:use_default_local_ike_id", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ipsec", "ike_parameters", "use_default_local_ike_id"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ike parameters use default remote ike id"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:use_default_remote_ike_id", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "use_default_remote_ike_id"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/ipsec/ike_parameters/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "IKE configuration parameters required for IPsec Connection type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["external_connectorCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ike_parameters

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/)
- ipsec.ike_parameters

<a id="section"></a>

Type: `"single"`. Computed.

IKE configuration parameters required for IPsec Connection type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dpd_choice": "[\"dpd_disabled\",\"dpd_keep_alive_timer\"]",
  "x-ves-oneof-field-local_ike_id": "[\"use_default_local_ike_id\"]",
  "x-ves-oneof-field-mode_choice": "[\"initiator\",\"responder\"]",
  "x-ves-oneof-field-remote_ike_id": "[\"rm_hostname\",\"rm_ip_address\",\"use_default_remote_ike_id\"]"
}
```

## Direct properties

- [dpd_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/dpd_disabled/): complete subsection reference.

- [dpd_keep_alive_timer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/dpd_keep_alive_timer/): complete subsection reference.

- [ike_phase1_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/ike_phase1_profile/): complete subsection reference.

- [ike_phase2_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/ike_phase2_profile/): complete subsection reference.

- [initiator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/initiator/): complete subsection reference.

- [responder](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/responder/): complete subsection reference.

<a id="schema-ipsec--ike_parameters--rm_hostname"></a>

### rm_hostname property

Type: `"string"`. Computed.

Exclusive with \[rm\_ip\_address use\_default\_remote\_ike\_id\] Configure an hostname Remote IKE
ID.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [rm_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/): complete subsection reference.

- [use_default_local_ike_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/use_default_local_ike_id/): complete subsection reference.

- [use_default_remote_ike_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/use_default_remote_ike_id/): complete subsection reference.
