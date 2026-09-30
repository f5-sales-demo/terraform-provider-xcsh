---
page_title: "protocol_policer.protocol.dns"
subcategory: ""
description: "protocol_policer.protocol.dns for xcsh_protocol_policer."
xcsh_docs: {"aliases": [], "body_bytes": 986, "body_sha256": "sha256:9fd67fa7e096f89e11b6868fb9af52a6bb7f0da841d9cca3fc0cc74b59bd62d9", "canonical_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:dns", "child_ids": [], "collection_id": "xcsh-docs:resources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:dns", "parent_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol", "path": "docs/guides/resources--protocol_policer--properties--protocol_policer--protocol--dns.md", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["protocol_policer", "protocol", "dns"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_policer/properties/protocol_policer/protocol/dns/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "protocol_policer.protocol.dns for xcsh_protocol_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# protocol_policer.protocol.dns

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md)
- [Property reference](resources--protocol_policer--reference.md)
- [protocol_policer](resources--protocol_policer--properties--protocol_policer.md)
- [protocol_policer.protocol](resources--protocol_policer--properties--protocol_policer--protocol.md)
- protocol_policer.protocol.dns

<a id="section"></a>

Type: `["object", {}]`. Optional.

Match all DNS packets including UDP and TCP.

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
dns = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [protocol_policer.protocol](resources--protocol_policer--properties--protocol_policer--protocol.md)
- [xcsh_protocol_policer](../resources/protocol_policer.md)
