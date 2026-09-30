---
page_title: "discovery_consul.access_info"
subcategory: ""
description: "discovery_consul.access_info for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1370, "body_sha256": "sha256:ae6970f30663fb71972b4c051c639f7e2246625c34a40894339ffa7077c90ba1", "canonical_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info", "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info", "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info"], "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_consul", "path": "docs/guides/resources--discovery--properties--discovery_consul--access_info.md", "provider_name": "discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_consul", "access_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_consul/access_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_consul.access_info for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# discovery_consul.access_info

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md)
- [Property reference](resources--discovery--reference.md)
- [discovery_consul](resources--discovery--properties--discovery_consul.md)
- discovery_consul.access_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Hashicorp Consul API server information.

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
access_info {
  # Configure direct properties listed below.
}
```

## Direct properties

- [connection_info](resources--discovery--properties--discovery_consul--access_info--connection_info.md): complete subsection reference.

- [http_basic_auth_info](resources--discovery--properties--discovery_consul--access_info--http_basic_auth_info.md): complete subsection reference.

## Next pages

- [discovery_consul.access_info.connection_info](resources--discovery--properties--discovery_consul--access_info--connection_info.md)
- [discovery_consul.access_info.http_basic_auth_info](resources--discovery--properties--discovery_consul--access_info--http_basic_auth_info.md)
- [discovery_consul](resources--discovery--properties--discovery_consul.md)
- [xcsh_discovery](../resources/discovery.md)
