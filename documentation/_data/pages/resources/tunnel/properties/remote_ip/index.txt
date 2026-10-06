---
page_title: "remote_ip"
subcategory: ""
description: "Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - 1. IP Address - Specifies the remote IP to which tunnel has to be connected 2. Remote endpoint - Is a map of IP address on per ver node basis."
xcsh_docs: {"aliases": ["remote ip"], "body_bytes": 1475, "body_sha256": "sha256:fb28e55d79e2333bfaec4710cde71418a36626222080959f4d4669e0e6800374", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:tunnel:properties:remote_ip:endpoints", "xcsh-docs:resources:tunnel:properties:remote_ip:ip"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:remote_ip", "parent_id": "xcsh-docs:resources:tunnel:reference", "path": "documentation/resources/tunnel/properties/remote_ip/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0020131231323103-2233031133231022-3221122302102332-0010301321220022-1200003103313200-0000212120201011-3323033323302300-1213020310102310", "registry_path": "docs/guides/resources--tunnel--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "remote_ip:ConflictingObjectAttributes:endpoints,ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:remote_ip:endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "remote_ip:ConflictingObjectAttributes:endpoints,ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["remote_ip"], "schema_version": 1, "sections": [{"aliases": ["remote ip endpoints"], "anchor": "section", "description": "Provides a map of ver node name to remote node attributes Ver node should use these attributes to configure as remote tunnel.", "document_id": "xcsh-docs:resources:tunnel:properties:remote_ip:endpoints", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["remote_ip", "endpoints"], "syntax": "block", "type": "object"}, {"aliases": ["remote ip ip"], "anchor": "section", "description": "IP Address used to specify an IPv4 or IPv6 address.", "document_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "remote_ip.ip:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "remote_ip.ip:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "remote_ip.ip:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "remote_ip.ip:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "remote_ip.ip:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:ipv6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "remote_ip.ip:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:ipv6", "type": "conflicts"}], "schema_path": ["remote_ip", "ip"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/remote_ip/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - 1. IP Address - Specifies the remote IP to which tunnel has to be connected 2. Remote endpoint - Is a map of IP address on per ver node basis.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
