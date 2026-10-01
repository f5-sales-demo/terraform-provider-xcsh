---
page_title: "virtual_server.http3"
subcategory: ""
description: "virtual_server.http3 for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 3430, "body_sha256": "sha256:ddc20ccbdd7b0cd1afaad4e871da524855bfda81b5f3dbf4fbd1d7e0b3317a80", "canonical_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3", "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:http3:client_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:http3_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:http_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:http_server_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:quic_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:server_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:tcp_server_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:udp_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:udp_server_profile"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "docs/guides/resources--application_profiles--properties--virtual_server--http3.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "http3"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/http3/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.http3 for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.http3

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md)
- [Property reference](resources--application_profiles--reference.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- virtual_server.http3

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HTTP/3 profiles.

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
http3 {
  # Configure direct properties listed below.
}
```

## Direct properties

- [client_ssl_profile](resources--application_profiles--properties--virtual_server--http3--client_ssl_profile.md): complete subsection reference.

- [http3_profile](resources--application_profiles--properties--virtual_server--http3--http3_profile.md): complete subsection reference.

- [http_client_profile](resources--application_profiles--properties--virtual_server--http3--http_client_profile.md): complete subsection reference.

- [http_server_profile](resources--application_profiles--properties--virtual_server--http3--http_server_profile.md): complete subsection reference.

- [quic_profile](resources--application_profiles--properties--virtual_server--http3--quic_profile.md): complete subsection reference.

- [server_ssl_profile](resources--application_profiles--properties--virtual_server--http3--server_ssl_profile.md): complete subsection reference.

- [tcp_server_profile](resources--application_profiles--properties--virtual_server--http3--tcp_server_profile.md): complete subsection reference.

- [udp_client_profile](resources--application_profiles--properties--virtual_server--http3--udp_client_profile.md): complete subsection reference.

- [udp_server_profile](resources--application_profiles--properties--virtual_server--http3--udp_server_profile.md): complete subsection reference.

## Next pages

- [virtual_server.http3.client_ssl_profile](resources--application_profiles--properties--virtual_server--http3--client_ssl_profile.md)
- [virtual_server.http3.http3_profile](resources--application_profiles--properties--virtual_server--http3--http3_profile.md)
- [virtual_server.http3.http_client_profile](resources--application_profiles--properties--virtual_server--http3--http_client_profile.md)
- [virtual_server.http3.http_server_profile](resources--application_profiles--properties--virtual_server--http3--http_server_profile.md)
- [virtual_server.http3.quic_profile](resources--application_profiles--properties--virtual_server--http3--quic_profile.md)
- [virtual_server.http3.server_ssl_profile](resources--application_profiles--properties--virtual_server--http3--server_ssl_profile.md)
- [virtual_server.http3.tcp_server_profile](resources--application_profiles--properties--virtual_server--http3--tcp_server_profile.md)
- [virtual_server.http3.udp_client_profile](resources--application_profiles--properties--virtual_server--http3--udp_client_profile.md)
- [virtual_server.http3.udp_server_profile](resources--application_profiles--properties--virtual_server--http3--udp_server_profile.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- [xcsh_application_profiles](../resources/application_profiles.md)
