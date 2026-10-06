---
page_title: "ipsec.ike_parameters"
subcategory: ""
description: "IKE configuration parameters required for IPsec Connection type."
xcsh_docs: {"aliases": ["ipsec ike parameters"], "body_bytes": 4163, "body_sha256": "sha256:e8ade05c21a094d3bd3aeb2e703d473c39e0eff0bcac4b754248ffc2c789e337", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_disabled", "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:ike_phase1_profile", "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:ike_phase2_profile", "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:initiator", "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:responder", "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:use_default_local_ike_id", "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:use_default_remote_ike_id"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters", "parent_id": "xcsh-docs:resources:external_connector:properties:ipsec", "path": "documentation/resources/external_connector/properties/ipsec/ike_parameters/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "resources", "registry_anchor": "canonical-2110112022302230-2103222202112010-1012130322030213-3100331333231022-3132330220301303-2012000231322123-0021133111013121-3222203023211013", "registry_path": "docs/guides/resources--external_connector--reference--group-001.md", "relationships": [{"anchor": "schema-ipsec--ike_parameters--rm_hostname", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:rm_hostname,rm_ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters", "type": "conflicts"}, {"anchor": "schema-ipsec--ike_parameters--rm_hostname", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:rm_hostname,use_default_remote_ike_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:dpd_disabled,dpd_keep_alive_timer", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:dpd_disabled,dpd_keep_alive_timer", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:initiator,responder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:initiator", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:initiator,responder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:responder", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:rm_hostname,rm_ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:rm_ip_address,use_default_remote_ike_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:rm_hostname,use_default_remote_ike_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:use_default_remote_ike_id", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:rm_ip_address,use_default_remote_ike_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:use_default_remote_ike_id", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ipsec", "ike_parameters"], "schema_version": 1, "sections": [{"aliases": ["ipsec ike parameters dpd disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_disabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "dpd_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ike parameters dpd keep alive timer"], "anchor": "section", "description": "Configuration parameter for dpd keep alive timer.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ipsec--ike_parameters--dpd_keep_alive_timer--timeout", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.dpd_keep_alive_timer:RequiredObjectAttributes:timeout", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "type": "requires"}], "schema_path": ["ipsec", "ike_parameters", "dpd_keep_alive_timer"], "syntax": "block", "type": "object"}, {"aliases": ["ipsec ike parameters ike phase1 profile"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:ike_phase1_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ipsec--ike_parameters--ike_phase1_profile--name", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.ike_phase1_profile:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:ike_phase1_profile", "type": "requires"}], "schema_path": ["ipsec", "ike_parameters", "ike_phase1_profile"], "syntax": "block", "type": "object"}, {"aliases": ["ipsec ike parameters ike phase2 profile"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:ike_phase2_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ipsec--ike_parameters--ike_phase2_profile--name", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.ike_phase2_profile:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:ike_phase2_profile", "type": "requires"}], "schema_path": ["ipsec", "ike_parameters", "ike_phase2_profile"], "syntax": "block", "type": "object"}, {"aliases": ["ipsec ike parameters initiator"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:initiator", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "initiator"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ike parameters responder"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:responder", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "responder"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ike parameters rm hostname"], "anchor": "schema-ipsec--ike_parameters--rm_hostname", "description": "Exclusive with Configure an hostname Remote IKE ID.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "rm_hostname"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipsec ike parameters rm ip address"], "anchor": "section", "description": "IP Address used to specify an IPv4 or IPv6 address.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.rm_ip_address:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.rm_ip_address:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.rm_ip_address:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.rm_ip_address:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.rm_ip_address:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:ipv6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.rm_ip_address:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:ipv6", "type": "conflicts"}], "schema_path": ["ipsec", "ike_parameters", "rm_ip_address"], "syntax": "block", "type": "object"}, {"aliases": ["ipsec ike parameters use default local ike id"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:use_default_local_ike_id", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ipsec", "ike_parameters", "use_default_local_ike_id"], "syntax": "block", "type": "object"}, {"aliases": ["ipsec ike parameters use default remote ike id"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:use_default_remote_ike_id", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "use_default_remote_ike_id"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/ipsec/ike_parameters/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "IKE configuration parameters required for IPsec Connection type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["external_connectorCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ike_parameters

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/)
- ipsec.ike_parameters

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IKE configuration parameters required for IPsec Connection type.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("dpd_disabled",
    "dpd_keep_alive_timer"),
  validators.ConflictingObjectAttributes("initiator",
    "responder"),
  validators.ConflictingObjectAttributes("rm_hostname",
    "rm_ip_address"),
  validators.ConflictingObjectAttributes("rm_hostname",
    "use_default_remote_ike_id"),
  validators.ConflictingObjectAttributes("rm_ip_address",
    "use_default_remote_ike_id")}
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
  "x-ves-oneof-field-dpd_choice": "[\"dpd_disabled\",\"dpd_keep_alive_timer\"]",
  "x-ves-oneof-field-local_ike_id": "[\"use_default_local_ike_id\"]",
  "x-ves-oneof-field-mode_choice": "[\"initiator\",\"responder\"]",
  "x-ves-oneof-field-remote_ike_id": "[\"rm_hostname\",\"rm_ip_address\",\"use_default_remote_ike_id\"]"
}
```

Terraform syntax:

```terraform
ike_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dpd_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/dpd_disabled/): complete subsection reference.

- [dpd_keep_alive_timer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/dpd_keep_alive_timer/): complete subsection reference.

- [ike_phase1_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/ike_phase1_profile/): complete subsection reference.

- [ike_phase2_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/ike_phase2_profile/): complete subsection reference.

- [initiator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/initiator/): complete subsection reference.

- [responder](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/responder/): complete subsection reference.

<a id="schema-ipsec--ike_parameters--rm_hostname"></a>

### rm_hostname property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [rm_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/): complete subsection reference.

- [use_default_local_ike_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/use_default_local_ike_id/): complete subsection reference.

- [use_default_remote_ike_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/use_default_remote_ike_id/): complete subsection reference.
