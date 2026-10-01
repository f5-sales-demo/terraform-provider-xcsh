---
page_title: "virtual_server.port_translation"
subcategory: ""
description: "virtual_server.port_translation for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 2590, "body_sha256": "sha256:e2b64191cfc4c9b6838a15d7b9e10d0103d0584f0d623672c50fe5bffa92e1db", "canonical_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation", "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation:port_translation_disable", "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation:port_translation_enable"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "docs/guides/resources--application_profiles--properties--virtual_server--port_translation.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "port_translation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/port_translation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.port_translation for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.port_translation

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md)
- [Property reference](resources--application_profiles--reference.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- virtual_server.port_translation

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies, when checked (enabled), that the system translates the port of the virtual server. When
cleared (disabled), specifies that the system uses the port without translation. Turning off port
translation for a virtual server is useful if you want to use the virtual server to load balance..

Upstream description:

Specifies, when checked (enabled), that the system translates the port of the virtual server. When
cleared (disabled), specifies that the system uses the port without translation. Turning off port
translation for a virtual server is useful if you want to use the virtual server to load balance
connections to any service. The default is enabled.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("port_translation_disable",
    "port_translation_enable")}
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
  "x-ves-oneof-field-port_translation_choice": "[\"port_translation_disable\",\"port_translation_enable\"]"
}
```

Terraform syntax:

```terraform
port_translation {
  # Configure direct properties listed below.
}
```

## Direct properties

- [port_translation_disable](resources--application_profiles--properties--virtual_server--port_translation--port_translation_disable.md): complete subsection reference.

- [port_translation_enable](resources--application_profiles--properties--virtual_server--port_translation--port_translation_enable.md): complete subsection reference.

## Next pages

- [virtual_server.port_translation.port_translation_disable](resources--application_profiles--properties--virtual_server--port_translation--port_translation_disable.md)
- [virtual_server.port_translation.port_translation_enable](resources--application_profiles--properties--virtual_server--port_translation--port_translation_enable.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- [xcsh_application_profiles](../resources/application_profiles.md)
