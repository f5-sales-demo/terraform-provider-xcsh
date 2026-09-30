---
page_title: "remote_ip.endpoints"
subcategory: ""
description: "remote_ip.endpoints for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1187, "body_sha256": "sha256:f2a84657ca0e6d80e4864235f0f549b4308e49d39bda701d3b1e3ef910c54e60", "canonical_id": "xcsh-docs:resources:tunnel:properties:remote_ip:endpoints", "child_ids": ["xcsh-docs:resources:tunnel:properties:remote_ip:endpoints:endpoints"], "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:remote_ip:endpoints", "parent_id": "xcsh-docs:resources:tunnel:properties:remote_ip", "path": "docs/guides/resources--tunnel--properties--remote_ip--endpoints.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["remote_ip", "endpoints"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/remote_ip/endpoints/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "remote_ip.endpoints for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# remote_ip.endpoints

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md)
- [Property reference](resources--tunnel--reference.md)
- [remote_ip](resources--tunnel--properties--remote_ip.md)
- remote_ip.endpoints

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Provides a map of ver node name to remote node attributes Ver node should use these attributes to
configure as remote tunnel.

Upstream description:

Provides a map of ver node name to remote node attributes Ver node should use these attributes to
configure as remote tunnel.

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
endpoints {
  # Configure direct properties listed below.
}
```

## Direct properties

- [endpoints](resources--tunnel--properties--remote_ip--endpoints--endpoints.md): complete subsection reference.

## Next pages

- [remote_ip.endpoints.endpoints](resources--tunnel--properties--remote_ip--endpoints--endpoints.md)
- [remote_ip](resources--tunnel--properties--remote_ip.md)
- [xcsh_tunnel](../resources/tunnel.md)
