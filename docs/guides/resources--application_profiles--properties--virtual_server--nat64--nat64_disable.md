---
page_title: "virtual_server.nat64.nat64_disable"
subcategory: ""
description: "virtual_server.nat64.nat64_disable for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 1193, "body_sha256": "sha256:4fddc415a9b2f0fe0daea184a7b667f117fb6b9735712c1defcd8b0512048d93", "canonical_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:nat64:nat64_disable", "child_ids": [], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:nat64:nat64_disable", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:nat64", "path": "docs/guides/resources--application_profiles--properties--virtual_server--nat64--nat64_disable.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "nat64", "nat64_disable"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/nat64/nat64_disable/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.nat64.nat64_disable for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.nat64.nat64_disable

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md)
- [Property reference](resources--application_profiles--reference.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- [virtual_server.nat64](resources--application_profiles--properties--virtual_server--nat64.md)
- virtual_server.nat64.nat64_disable

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for nat64 disable.

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
nat64_disable = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [virtual_server.nat64](resources--application_profiles--properties--virtual_server--nat64.md)
- [xcsh_application_profiles](../resources/application_profiles.md)
