---
page_title: "virtual_server.address_translation"
subcategory: ""
description: "virtual_server.address_translation for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 2501, "body_sha256": "sha256:133a9f8b9c260facc528964b6dc241f03bafbb97e487f7d3e112fc6c92986f79", "canonical_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:address_translation", "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:address_translation:address_translation_disable", "xcsh-docs:resources:application_profiles:properties:virtual_server:address_translation:address_translation_enable"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:address_translation", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "docs/guides/resources--application_profiles--properties--virtual_server--address_translation.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "address_translation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/address_translation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.address_translation for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# virtual_server.address_translation

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md)
- [Property reference](resources--application_profiles--reference.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- virtual_server.address_translation

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies, when checked (enabled), that the system translates the address of the virtual server.
When cleared (disabled), specifies that the system uses the address without translation. This option
is useful when the system is load balancing devices that have the same IP address.

Upstream description:

Specifies, when checked (enabled), that the system translates the address of the virtual server.
When cleared (disabled), specifies that the system uses the address without translation. This option
is useful when the system is load balancing devices that have the same IP address. The default is
enabled.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("address_translation_disable",
    "address_translation_enable")}
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
  "x-ves-oneof-field-address_translation_choice": "[\"address_translation_disable\",\"address_translation_enable\"]"
}
```

Terraform syntax:

```terraform
address_translation {
  # Configure direct properties listed below.
}
```

## Direct properties

- [address_translation_disable](resources--application_profiles--properties--virtual_server--address_translation--address_translation_disable.md): complete subsection reference.

- [address_translation_enable](resources--application_profiles--properties--virtual_server--address_translation--address_translation_enable.md): complete subsection reference.

## Next pages

- [virtual_server.address_translation.address_translation_disable](resources--application_profiles--properties--virtual_server--address_translation--address_translation_disable.md)
- [virtual_server.address_translation.address_translation_enable](resources--application_profiles--properties--virtual_server--address_translation--address_translation_enable.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- [xcsh_application_profiles](../resources/application_profiles.md)
