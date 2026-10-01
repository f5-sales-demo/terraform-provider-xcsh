---
page_title: "virtual_server.address_translation"
subcategory: ""
description: "virtual_server.address_translation for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 2296, "body_sha256": "sha256:982a51621363f17aa5d09c60bd72bc7b0d6988860a4ff461edc52e3d5b42d8a8", "canonical_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:address_translation", "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:address_translation:address_translation_disable", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:address_translation:address_translation_enable"], "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:address_translation", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "docs/guides/data-sources--application_profiles--properties--virtual_server--address_translation.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "address_translation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/address_translation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.address_translation for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.address_translation

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md)
- [Property reference](data-sources--application_profiles--reference.md)
- [virtual_server](data-sources--application_profiles--properties--virtual_server.md)
- virtual_server.address_translation

<a id="section"></a>

Type: `"single"`. Computed.

Specifies, when checked (enabled), that the system translates the address of the virtual server.
When cleared (disabled), specifies that the system uses the address without translation. This option
is useful when the system is load balancing devices that have the same IP address.

Upstream description:

Specifies, when checked (enabled), that the system translates the address of the virtual server.
When cleared (disabled), specifies that the system uses the address without translation. This option
is useful when the system is load balancing devices that have the same IP address. The default is
enabled.

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

## Direct properties

- [address_translation_disable](data-sources--application_profiles--properties--virtual_server--address_translation--address_translation_disable.md): complete subsection reference.

- [address_translation_enable](data-sources--application_profiles--properties--virtual_server--address_translation--address_translation_enable.md): complete subsection reference.

## Next pages

- [virtual_server.address_translation.address_translation_disable](data-sources--application_profiles--properties--virtual_server--address_translation--address_translation_disable.md)
- [virtual_server.address_translation.address_translation_enable](data-sources--application_profiles--properties--virtual_server--address_translation--address_translation_enable.md)
- [virtual_server](data-sources--application_profiles--properties--virtual_server.md)
- [xcsh_application_profiles](../data-sources/application_profiles.md)
