---
page_title: "ipsec"
subcategory: ""
description: "ipsec for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1346, "body_sha256": "sha256:338e01f0097935b86e245a1d982bc9978f6439ceceabb7a081d5399140568254", "canonical_id": "xcsh-docs:resources:external_connector:properties:ipsec", "child_ids": ["xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters", "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters"], "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:ipsec", "parent_id": "xcsh-docs:resources:external_connector:reference", "path": "docs/guides/resources--external_connector--properties--ipsec.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ipsec"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/ipsec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ipsec for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md)
- [Property reference](resources--external_connector--reference.md)
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

- [ike_parameters](resources--external_connector--properties--ipsec--ike_parameters.md): complete subsection reference.

- [ipsec_tunnel_parameters](resources--external_connector--properties--ipsec--ipsec_tunnel_parameters.md): complete subsection reference.

## Next pages

- [ipsec.ike_parameters](resources--external_connector--properties--ipsec--ike_parameters.md)
- [ipsec.ipsec_tunnel_parameters](resources--external_connector--properties--ipsec--ipsec_tunnel_parameters.md)
- [Property reference](resources--external_connector--reference.md)
- [xcsh_external_connector](../resources/external_connector.md)
