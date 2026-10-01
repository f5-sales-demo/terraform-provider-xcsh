---
page_title: "ipsec.ike_parameters.rm_ip_address.dual_stack"
subcategory: ""
description: "ipsec.ike_parameters.rm_ip_address.dual_stack for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1770, "body_sha256": "sha256:baedb5b0bda13f7ae9e9fdb69546602687f0660924503a182332310e95705e6c", "canonical_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:dual_stack", "child_ids": ["xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:dual_stack:ipv4", "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:dual_stack:ipv6"], "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:dual_stack", "parent_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "path": "docs/guides/data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--dual_stack.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ipsec", "ike_parameters", "rm_ip_address", "dual_stack"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/dual_stack/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ipsec.ike_parameters.rm_ip_address.dual_stack for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ike_parameters.rm_ip_address.dual_stack

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md)
- [Property reference](data-sources--external_connector--reference.md)
- [ipsec](data-sources--external_connector--properties--ipsec.md)
- [ipsec.ike_parameters](data-sources--external_connector--properties--ipsec--ike_parameters.md)
- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address.md)
- ipsec.ike_parameters.rm_ip_address.dual_stack

<a id="section"></a>

Type: `"single"`. Computed.

DualStackAddressType represents both IPv4 and IPv6 together.

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

- [ipv4](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--dual_stack--ipv4.md): complete subsection reference.

- [ipv6](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--dual_stack--ipv6.md): complete subsection reference.

## Next pages

- [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--dual_stack--ipv4.md)
- [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--dual_stack--ipv6.md)
- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address.md)
- [xcsh_external_connector](../data-sources/external_connector.md)
