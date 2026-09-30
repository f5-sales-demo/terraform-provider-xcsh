---
page_title: "remote_ip"
subcategory: ""
description: "remote_ip for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1671, "body_sha256": "sha256:63fc34362822a17e0d5095b622cbba4e4905c0665cca97c603ae426cda24c468", "canonical_id": "xcsh-docs:resources:tunnel:properties:remote_ip", "child_ids": ["xcsh-docs:resources:tunnel:properties:remote_ip:endpoints", "xcsh-docs:resources:tunnel:properties:remote_ip:ip"], "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:remote_ip", "parent_id": "xcsh-docs:resources:tunnel:reference", "path": "docs/guides/resources--tunnel--properties--remote_ip.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["remote_ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/remote_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "remote_ip for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# remote_ip

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md)
- [Property reference](resources--tunnel--reference.md)
- remote_ip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - 1. IP
Address - Specifies the remote IP to which tunnel has to be connected 2. Remote endpoint - Is a map
of IP address on per ver node basis.

Upstream description:

Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - &#8203;1.
IP Address - Specifies the remote IP to which tunnel has to be connected &#8203;2. Remote endpoint -
Is a map of IP address on per ver node basis.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("endpoints",
    "ip")}
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
  "x-ves-oneof-field-type": "[\"endpoints\",\"ip\"]"
}
```

Terraform syntax:

```terraform
remote_ip {
  # Configure direct properties listed below.
}
```

## Direct properties

- [endpoints](resources--tunnel--properties--remote_ip--endpoints.md): complete subsection reference.

- [ip](resources--tunnel--properties--remote_ip--ip.md): complete subsection reference.

## Next pages

- [remote_ip.endpoints](resources--tunnel--properties--remote_ip--endpoints.md)
- [remote_ip.ip](resources--tunnel--properties--remote_ip--ip.md)
- [Property reference](resources--tunnel--reference.md)
- [xcsh_tunnel](../resources/tunnel.md)
