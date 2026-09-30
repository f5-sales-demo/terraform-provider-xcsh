---
page_title: "ipsec.ike_parameters.rm_ip_address"
subcategory: ""
description: "ipsec.ike_parameters.rm_ip_address for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 2156, "body_sha256": "sha256:b1a6959b953a3d13660b7f8ce3e730ac4b2a5390203d5fca8d09ae094163c904", "canonical_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "child_ids": ["xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:dual_stack", "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:ipv4", "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:ipv6"], "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "parent_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters", "path": "docs/guides/resources--external_connector--properties--ipsec--ike_parameters--rm_ip_address.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ipsec", "ike_parameters", "rm_ip_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ipsec.ike_parameters.rm_ip_address for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ipsec.ike_parameters.rm_ip_address

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md)
- [Property reference](resources--external_connector--reference.md)
- [ipsec](resources--external_connector--properties--ipsec.md)
- [ipsec.ike_parameters](resources--external_connector--properties--ipsec--ike_parameters.md)
- ipsec.ike_parameters.rm_ip_address

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
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
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
rm_ip_address {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dual_stack](resources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--dual_stack.md): complete subsection reference.

- [ipv4](resources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--ipv4.md): complete subsection reference.

- [ipv6](resources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--ipv6.md): complete subsection reference.

## Next pages

- [ipsec.ike_parameters.rm_ip_address.dual_stack](resources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--dual_stack.md)
- [ipsec.ike_parameters.rm_ip_address.ipv4](resources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--ipv4.md)
- [ipsec.ike_parameters.rm_ip_address.ipv6](resources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--ipv6.md)
- [ipsec.ike_parameters](resources--external_connector--properties--ipsec--ike_parameters.md)
- [xcsh_external_connector](../resources/external_connector.md)
