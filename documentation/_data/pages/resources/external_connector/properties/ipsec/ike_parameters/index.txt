---
page_title: "ipsec.ike_parameters"
subcategory: ""
description: "IKE configuration parameters required for IPsec Connection type."
xcsh_docs: {"aliases": ["ipsec ike parameters"], "body_bytes": 6114, "body_sha256": "sha256:ecd7f5547684d0f13d90502c279b22b3e19e90047d9a7e830cc90f515f12418b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_disabled", "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:ike_phase1_profile", "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:ike_phase2_profile", "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:initiator", "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:responder", "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:use_default_local_ike_id", "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:use_default_remote_ike_id"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters", "parent_id": "xcsh-docs:resources:external_connector:properties:ipsec", "path": "documentation/resources/external_connector/properties/ipsec/ike_parameters/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2110112022302230-2103222202112010-1012130322030213-3100331333231022-3132330220301303-2012000231322123-0021133111013121-3222203023211013", "registry_path": "docs/guides/resources--external_connector--reference--group-001.md", "relationships": [{"anchor": "schema-ipsec--ike_parameters--rm_hostname", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:rm_hostname,rm_ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters", "type": "conflicts"}, {"anchor": "schema-ipsec--ike_parameters--rm_hostname", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:rm_hostname,use_default_remote_ike_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:dpd_disabled,dpd_keep_alive_timer", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:dpd_disabled,dpd_keep_alive_timer", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:initiator,responder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:initiator", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:initiator,responder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:responder", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:rm_hostname,rm_ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:rm_ip_address,use_default_remote_ike_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:rm_hostname,use_default_remote_ike_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:use_default_remote_ike_id", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters:ConflictingObjectAttributes:rm_ip_address,use_default_remote_ike_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:use_default_remote_ike_id", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ipsec", "ike_parameters"], "schema_version": 1, "sections": [{"aliases": ["ipsec ike parameters dpd disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "dpd_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ike parameters dpd keep alive timer"], "anchor": "section", "description": "Configuration parameter for dpd keep alive timer.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ipsec--ike_parameters--dpd_keep_alive_timer--timeout", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.dpd_keep_alive_timer:RequiredObjectAttributes:timeout", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "type": "requires"}], "schema_path": ["ipsec", "ike_parameters", "dpd_keep_alive_timer"], "syntax": "block", "type": "object"}, {"aliases": ["ipsec ike parameters ike phase1 profile"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:ike_phase1_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ipsec--ike_parameters--ike_phase1_profile--name", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.ike_phase1_profile:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:ike_phase1_profile", "type": "requires"}], "schema_path": ["ipsec", "ike_parameters", "ike_phase1_profile"], "syntax": "block", "type": "object"}, {"aliases": ["ipsec ike parameters ike phase2 profile"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:ike_phase2_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ipsec--ike_parameters--ike_phase2_profile--name", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.ike_phase2_profile:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:ike_phase2_profile", "type": "requires"}], "schema_path": ["ipsec", "ike_parameters", "ike_phase2_profile"], "syntax": "block", "type": "object"}, {"aliases": ["ipsec ike parameters initiator"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:initiator", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "initiator"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ike parameters responder"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:responder", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "responder"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipsec ike parameters rm hostname"], "anchor": "schema-ipsec--ike_parameters--rm_hostname", "description": "Exclusive with Configure an hostname Remote IKE ID.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "rm_hostname"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipsec ike parameters rm ip address"], "anchor": "section", "description": "IP Address used to specify an IPv4 or IPv6 address.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.rm_ip_address:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.rm_ip_address:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.rm_ip_address:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.rm_ip_address:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.rm_ip_address:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:ipv6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ipsec.ike_parameters.rm_ip_address:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:ipv6", "type": "conflicts"}], "schema_path": ["ipsec", "ike_parameters", "rm_ip_address"], "syntax": "block", "type": "object"}, {"aliases": ["ipsec ike parameters use default local ike id"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:use_default_local_ike_id", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ipsec", "ike_parameters", "use_default_local_ike_id"], "syntax": "block", "type": "object"}, {"aliases": ["ipsec ike parameters use default remote ike id"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:use_default_remote_ike_id", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "use_default_remote_ike_id"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/ipsec/ike_parameters/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "IKE configuration parameters required for IPsec Connection type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["external_connectorCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
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

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

## Next pages

- [ipsec.ike_parameters.dpd_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/dpd_disabled/)
- [ipsec.ike_parameters.dpd_keep_alive_timer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/dpd_keep_alive_timer/)
- [ipsec.ike_parameters.ike_phase1_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/ike_phase1_profile/)
- [ipsec.ike_parameters.ike_phase2_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/ike_phase2_profile/)
- [ipsec.ike_parameters.initiator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/initiator/)
- [ipsec.ike_parameters.responder](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/responder/)
- [ipsec.ike_parameters.rm_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/)
- [ipsec.ike_parameters.use_default_local_ike_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/use_default_local_ike_id/)
- [ipsec.ike_parameters.use_default_remote_ike_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/use_default_remote_ike_id/)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/)
- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
