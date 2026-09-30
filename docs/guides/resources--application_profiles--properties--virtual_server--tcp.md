---
page_title: "virtual_server.tcp"
subcategory: ""
description: "virtual_server.tcp for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 2176, "body_sha256": "sha256:be5b0057323488e5d81dcea3b58be98cdc4cd778d95fa3bf2a403e6cec1cb3a8", "canonical_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp", "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:client_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:ocsp_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:server_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:tcp_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:tcp_server_profile"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "docs/guides/resources--application_profiles--properties--virtual_server--tcp.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "tcp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/tcp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.tcp for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# virtual_server.tcp

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md)
- [Property reference](resources--application_profiles--reference.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- virtual_server.tcp

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

TCP profiles.

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
tcp {
  # Configure direct properties listed below.
}
```

## Direct properties

- [client_ssl_profile](resources--application_profiles--properties--virtual_server--tcp--client_ssl_profile.md): complete subsection reference.

- [ocsp_profile](resources--application_profiles--properties--virtual_server--tcp--ocsp_profile.md): complete subsection reference.

- [server_ssl_profile](resources--application_profiles--properties--virtual_server--tcp--server_ssl_profile.md): complete subsection reference.

- [tcp_client_profile](resources--application_profiles--properties--virtual_server--tcp--tcp_client_profile.md): complete subsection reference.

- [tcp_server_profile](resources--application_profiles--properties--virtual_server--tcp--tcp_server_profile.md): complete subsection reference.

## Next pages

- [virtual_server.tcp.client_ssl_profile](resources--application_profiles--properties--virtual_server--tcp--client_ssl_profile.md)
- [virtual_server.tcp.ocsp_profile](resources--application_profiles--properties--virtual_server--tcp--ocsp_profile.md)
- [virtual_server.tcp.server_ssl_profile](resources--application_profiles--properties--virtual_server--tcp--server_ssl_profile.md)
- [virtual_server.tcp.tcp_client_profile](resources--application_profiles--properties--virtual_server--tcp--tcp_client_profile.md)
- [virtual_server.tcp.tcp_server_profile](resources--application_profiles--properties--virtual_server--tcp--tcp_server_profile.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- [xcsh_application_profiles](../resources/application_profiles.md)
