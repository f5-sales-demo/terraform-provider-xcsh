---
page_title: "ipsec.ipsec_tunnel_parameters.segment"
subcategory: ""
description: "ipsec.ipsec_tunnel_parameters.segment for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1483, "body_sha256": "sha256:65830f0f833479a18d72cc0a968722180baa58e93a605b46f5c14838f3c5ccce", "canonical_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment", "child_ids": ["xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment:refs"], "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment", "parent_id": "xcsh-docs:resources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "path": "docs/guides/resources--external_connector--properties--ipsec--ipsec_tunnel_parameters--segment.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ipsec", "ipsec_tunnel_parameters", "segment"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ipsec.ipsec_tunnel_parameters.segment for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ipsec.ipsec_tunnel_parameters.segment

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md)
- [Property reference](resources--external_connector--reference.md)
- [ipsec](resources--external_connector--properties--ipsec.md)
- [ipsec.ipsec_tunnel_parameters](resources--external_connector--properties--ipsec--ipsec_tunnel_parameters.md)
- ipsec.ipsec_tunnel_parameters.segment

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Segment Reference Type. Reference to Segment Object.

Upstream description:

Reference to Segment Object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("refs")}
```

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
segment {
  # Configure direct properties listed below.
}
```

## Direct properties

- [refs](resources--external_connector--properties--ipsec--ipsec_tunnel_parameters--segment--refs.md): complete subsection reference.

## Next pages

- [ipsec.ipsec_tunnel_parameters.segment.refs](resources--external_connector--properties--ipsec--ipsec_tunnel_parameters--segment--refs.md)
- [ipsec.ipsec_tunnel_parameters](resources--external_connector--properties--ipsec--ipsec_tunnel_parameters.md)
- [xcsh_external_connector](../resources/external_connector.md)
