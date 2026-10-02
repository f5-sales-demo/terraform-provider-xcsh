---
page_title: "remote_ip"
subcategory: ""
description: "Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - 1. IP Address - Specifies the remote IP to which tunnel has to be connected 2. Remote endpoint - Is a map of IP address on per ver node basis."
xcsh_docs: {"aliases": ["remote ip"], "body_bytes": 2178, "body_sha256": "sha256:deb16d0b4f4cbab20482888334b74477dec826ec906911058b340b1a53fc9dc3", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:tunnel:properties:remote_ip:endpoints", "xcsh-docs:resources:tunnel:properties:remote_ip:ip"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:remote_ip", "parent_id": "xcsh-docs:resources:tunnel:reference", "path": "documentation/resources/tunnel/properties/remote_ip/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0020131231323103-2233031133231022-3221122302102332-0010301321220022-1200003103313200-0000212120201011-3323033323302300-1213020310102310", "registry_path": "docs/guides/resources--tunnel--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "remote_ip:ConflictingObjectAttributes:endpoints,ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:remote_ip:endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "remote_ip:ConflictingObjectAttributes:endpoints,ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["remote_ip"], "schema_version": 1, "sections": [{"aliases": ["endpoints"], "anchor": "section", "description": "Provides a map of ver node name to remote node attributes Ver node should use these attributes to configure as remote tunnel.", "document_id": "xcsh-docs:resources:tunnel:properties:remote_ip:endpoints", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["remote_ip", "endpoints"], "syntax": "block", "type": "object"}, {"aliases": ["ip"], "anchor": "section", "description": "IP Address used to specify an IPv4 or IPv6 address.", "document_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "remote_ip.ip:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "remote_ip.ip:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "remote_ip.ip:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "remote_ip.ip:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "remote_ip.ip:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:ipv6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "remote_ip.ip:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:ipv6", "type": "conflicts"}], "schema_path": ["remote_ip", "ip"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/remote_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - 1. IP Address - Specifies the remote IP to which tunnel has to be connected 2. Remote endpoint - Is a map of IP address on per ver node basis.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# remote_ip

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
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

- [endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/endpoints/): complete subsection reference.

- [ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/): complete subsection reference.

## Next pages

- [remote_ip.endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/endpoints/)
- [remote_ip.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
