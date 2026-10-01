---
page_title: "virtual_server.udp"
subcategory: ""
description: "virtual_server.udp for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 2023, "body_sha256": "sha256:626a1e6615b4b4320c9ce004efaa742492d329c060bc1b3015e1f72132765eae", "canonical_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:udp", "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:udp:client_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:udp:server_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:udp:udp_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:udp:udp_server_profile"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:udp", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "docs/guides/resources--application_profiles--properties--virtual_server--udp.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "udp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/udp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.udp for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.udp

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md)
- [Property reference](resources--application_profiles--reference.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- virtual_server.udp

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

UDP profiles.

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
udp {
  # Configure direct properties listed below.
}
```

## Direct properties

- [client_ssl_profile](resources--application_profiles--properties--virtual_server--udp--client_ssl_profile.md): complete subsection reference.

- [server_ssl_profile](resources--application_profiles--properties--virtual_server--udp--server_ssl_profile.md): complete subsection reference.

- [udp_client_profile](resources--application_profiles--properties--virtual_server--udp--udp_client_profile.md): complete subsection reference.

- [udp_server_profile](resources--application_profiles--properties--virtual_server--udp--udp_server_profile.md): complete subsection reference.

## Next pages

- [virtual_server.udp.client_ssl_profile](resources--application_profiles--properties--virtual_server--udp--client_ssl_profile.md)
- [virtual_server.udp.server_ssl_profile](resources--application_profiles--properties--virtual_server--udp--server_ssl_profile.md)
- [virtual_server.udp.udp_client_profile](resources--application_profiles--properties--virtual_server--udp--udp_client_profile.md)
- [virtual_server.udp.udp_server_profile](resources--application_profiles--properties--virtual_server--udp--udp_server_profile.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- [xcsh_application_profiles](../resources/application_profiles.md)
