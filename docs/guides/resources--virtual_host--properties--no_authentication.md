---
page_title: "no_authentication"
subcategory: ""
description: "no_authentication for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 815, "body_sha256": "sha256:4dbe43112abf35a51e846b0ec8485a6078c4399db13516ac4fb785676902face", "canonical_id": "xcsh-docs:resources:virtual_host:properties:no_authentication", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:no_authentication", "parent_id": "xcsh-docs:resources:virtual_host:reference", "path": "docs/guides/resources--virtual_host--properties--no_authentication.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["no_authentication"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/no_authentication/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "no_authentication for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# no_authentication

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- no_authentication

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no authentication.

Upstream description:

This can be used for messages where no values are needed.

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
no_authentication = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--virtual_host--reference.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
