---
page_title: "custom_network_config.interface_list"
subcategory: ""
description: "custom_network_config.interface_list for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1438, "body_sha256": "sha256:d6baa34c49b82c79c2e7aeece28897f308705c54cadc4ce92c41ad35012dfa0c", "canonical_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces"], "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list", "parent_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config", "path": "docs/guides/resources--securemesh_site--properties--custom_network_config--interface_list.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "interface_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/custom_network_config/interface_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.interface_list for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md)
- [Property reference](resources--securemesh_site--reference.md)
- [custom_network_config](resources--securemesh_site--properties--custom_network_config.md)
- custom_network_config.interface_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure network interfaces for this Secure Mesh site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interfaces")}
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
interface_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [interfaces](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces.md): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces.md)
- [custom_network_config](resources--securemesh_site--properties--custom_network_config.md)
- [xcsh_securemesh_site](../resources/securemesh_site.md)
