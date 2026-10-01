---
page_title: "ipsec.ike_parameters"
subcategory: ""
description: "ipsec.ike_parameters for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 4407, "body_sha256": "sha256:84baabc5191c36236213390a4619cd3b0546145b7aea0b98c9d9dc1023910065", "canonical_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters", "child_ids": ["xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:dpd_disabled", "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:ike_phase1_profile", "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:ike_phase2_profile", "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:initiator", "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:responder", "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:use_default_local_ike_id", "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:use_default_remote_ike_id"], "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters", "parent_id": "xcsh-docs:data-sources:external_connector:properties:ipsec", "path": "docs/guides/data-sources--external_connector--properties--ipsec--ike_parameters.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ipsec", "ike_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/ipsec/ike_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ipsec.ike_parameters for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ike_parameters

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md)
- [Property reference](data-sources--external_connector--reference.md)
- [ipsec](data-sources--external_connector--properties--ipsec.md)
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

- [dpd_disabled](data-sources--external_connector--properties--ipsec--ike_parameters--dpd_disabled.md): complete subsection reference.

- [dpd_keep_alive_timer](data-sources--external_connector--properties--ipsec--ike_parameters--dpd_keep_alive_timer.md): complete subsection reference.

- [ike_phase1_profile](data-sources--external_connector--properties--ipsec--ike_parameters--ike_phase1_profile.md): complete subsection reference.

- [ike_phase2_profile](data-sources--external_connector--properties--ipsec--ike_parameters--ike_phase2_profile.md): complete subsection reference.

- [initiator](data-sources--external_connector--properties--ipsec--ike_parameters--initiator.md): complete subsection reference.

- [responder](data-sources--external_connector--properties--ipsec--ike_parameters--responder.md): complete subsection reference.

<a id="schema-ipsec--ike_parameters--rm_hostname"></a>

### rm_hostname property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [rm_ip_address](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address.md): complete subsection reference.

- [use_default_local_ike_id](data-sources--external_connector--properties--ipsec--ike_parameters--use_default_local_ike_id.md): complete subsection reference.

- [use_default_remote_ike_id](data-sources--external_connector--properties--ipsec--ike_parameters--use_default_remote_ike_id.md): complete subsection reference.

## Next pages

- [ipsec.ike_parameters.dpd_disabled](data-sources--external_connector--properties--ipsec--ike_parameters--dpd_disabled.md)
- [ipsec.ike_parameters.dpd_keep_alive_timer](data-sources--external_connector--properties--ipsec--ike_parameters--dpd_keep_alive_timer.md)
- [ipsec.ike_parameters.ike_phase1_profile](data-sources--external_connector--properties--ipsec--ike_parameters--ike_phase1_profile.md)
- [ipsec.ike_parameters.ike_phase2_profile](data-sources--external_connector--properties--ipsec--ike_parameters--ike_phase2_profile.md)
- [ipsec.ike_parameters.initiator](data-sources--external_connector--properties--ipsec--ike_parameters--initiator.md)
- [ipsec.ike_parameters.responder](data-sources--external_connector--properties--ipsec--ike_parameters--responder.md)
- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address.md)
- [ipsec.ike_parameters.use_default_local_ike_id](data-sources--external_connector--properties--ipsec--ike_parameters--use_default_local_ike_id.md)
- [ipsec.ike_parameters.use_default_remote_ike_id](data-sources--external_connector--properties--ipsec--ike_parameters--use_default_remote_ike_id.md)
- [ipsec](data-sources--external_connector--properties--ipsec.md)
- [xcsh_external_connector](../data-sources/external_connector.md)
