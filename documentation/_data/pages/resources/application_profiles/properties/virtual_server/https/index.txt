---
page_title: "virtual_server.https"
subcategory: ""
description: "virtual_server.https for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 5775, "body_sha256": "sha256:7a1cf54f8b880b6d94c39bdbcc2e89f04f94c404030feb8ef18de4016855c86f", "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:https:client_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:https:http2_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:https:http2_server_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:https:http_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:https:http_server_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:https:ocsp_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:https:server_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:https:stream_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:https:tcp_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:https:tcp_server_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:https:websocket_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:https:websocket_server_profile"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:https", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "documentation/resources/application_profiles/properties/virtual_server/https/index.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["virtual_server", "https"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/https/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.https for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.https

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- virtual_server.https

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
https {
  # Configure direct properties listed below.
}
```

## Direct properties

- [client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/client_ssl_profile/): complete subsection reference.

- [http2_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/http2_client_profile/): complete subsection reference.

- [http2_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/http2_server_profile/): complete subsection reference.

- [http_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/http_client_profile/): complete subsection reference.

- [http_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/http_server_profile/): complete subsection reference.

- [ocsp_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/ocsp_profile/): complete subsection reference.

- [server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/server_ssl_profile/): complete subsection reference.

- [stream_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/stream_profile/): complete subsection reference.

- [tcp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/tcp_client_profile/): complete subsection reference.

- [tcp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/tcp_server_profile/): complete subsection reference.

- [websocket_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/websocket_client_profile/): complete subsection reference.

- [websocket_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/websocket_server_profile/): complete subsection reference.

## Next pages

- [virtual_server.https.client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/client_ssl_profile/)
- [virtual_server.https.http2_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/http2_client_profile/)
- [virtual_server.https.http2_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/http2_server_profile/)
- [virtual_server.https.http_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/http_client_profile/)
- [virtual_server.https.http_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/http_server_profile/)
- [virtual_server.https.ocsp_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/ocsp_profile/)
- [virtual_server.https.server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/server_ssl_profile/)
- [virtual_server.https.stream_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/stream_profile/)
- [virtual_server.https.tcp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/tcp_client_profile/)
- [virtual_server.https.tcp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/tcp_server_profile/)
- [virtual_server.https.websocket_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/websocket_client_profile/)
- [virtual_server.https.websocket_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/websocket_server_profile/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
