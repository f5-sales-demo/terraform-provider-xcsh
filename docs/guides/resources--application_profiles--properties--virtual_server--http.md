---
page_title: "virtual_server.http"
subcategory: ""
description: "virtual_server.http for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 4303, "body_sha256": "sha256:f1fdb8d9d6c2ce7a11ce0d749a82625ab7183ec7bf3df01d8a03e388705c1e4f", "canonical_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http", "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:http:client_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http:http2_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http:http2_server_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http:http_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http:http_server_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http:ocsp_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http:server_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http:stream_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http:tcp_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http:tcp_server_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http:websocket_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http:websocket_server_profile"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "docs/guides/resources--application_profiles--properties--virtual_server--http.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "http"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/http/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.http for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.http

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md)
- [Property reference](resources--application_profiles--reference.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- virtual_server.http

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HTTP profiles.

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
http {
  # Configure direct properties listed below.
}
```

## Direct properties

- [client_ssl_profile](resources--application_profiles--properties--virtual_server--http--client_ssl_profile.md): complete subsection reference.

- [http2_client_profile](resources--application_profiles--properties--virtual_server--http--http2_client_profile.md): complete subsection reference.

- [http2_server_profile](resources--application_profiles--properties--virtual_server--http--http2_server_profile.md): complete subsection reference.

- [http_client_profile](resources--application_profiles--properties--virtual_server--http--http_client_profile.md): complete subsection reference.

- [http_server_profile](resources--application_profiles--properties--virtual_server--http--http_server_profile.md): complete subsection reference.

- [ocsp_profile](resources--application_profiles--properties--virtual_server--http--ocsp_profile.md): complete subsection reference.

- [server_ssl_profile](resources--application_profiles--properties--virtual_server--http--server_ssl_profile.md): complete subsection reference.

- [stream_profile](resources--application_profiles--properties--virtual_server--http--stream_profile.md): complete subsection reference.

- [tcp_client_profile](resources--application_profiles--properties--virtual_server--http--tcp_client_profile.md): complete subsection reference.

- [tcp_server_profile](resources--application_profiles--properties--virtual_server--http--tcp_server_profile.md): complete subsection reference.

- [websocket_client_profile](resources--application_profiles--properties--virtual_server--http--websocket_client_profile.md): complete subsection reference.

- [websocket_server_profile](resources--application_profiles--properties--virtual_server--http--websocket_server_profile.md): complete subsection reference.

## Next pages

- [virtual_server.http.client_ssl_profile](resources--application_profiles--properties--virtual_server--http--client_ssl_profile.md)
- [virtual_server.http.http2_client_profile](resources--application_profiles--properties--virtual_server--http--http2_client_profile.md)
- [virtual_server.http.http2_server_profile](resources--application_profiles--properties--virtual_server--http--http2_server_profile.md)
- [virtual_server.http.http_client_profile](resources--application_profiles--properties--virtual_server--http--http_client_profile.md)
- [virtual_server.http.http_server_profile](resources--application_profiles--properties--virtual_server--http--http_server_profile.md)
- [virtual_server.http.ocsp_profile](resources--application_profiles--properties--virtual_server--http--ocsp_profile.md)
- [virtual_server.http.server_ssl_profile](resources--application_profiles--properties--virtual_server--http--server_ssl_profile.md)
- [virtual_server.http.stream_profile](resources--application_profiles--properties--virtual_server--http--stream_profile.md)
- [virtual_server.http.tcp_client_profile](resources--application_profiles--properties--virtual_server--http--tcp_client_profile.md)
- [virtual_server.http.tcp_server_profile](resources--application_profiles--properties--virtual_server--http--tcp_server_profile.md)
- [virtual_server.http.websocket_client_profile](resources--application_profiles--properties--virtual_server--http--websocket_client_profile.md)
- [virtual_server.http.websocket_server_profile](resources--application_profiles--properties--virtual_server--http--websocket_server_profile.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- [xcsh_application_profiles](../resources/application_profiles.md)
