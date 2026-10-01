---
page_title: "virtual_server.http3"
subcategory: ""
description: "virtual_server.http3 for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 4569, "body_sha256": "sha256:81e17e4cb82d6f27b7559897257ce23e4737bc992116ff16efbc7b07121e48f0", "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:http3:client_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:http3_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:http_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:http_server_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:quic_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:server_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:tcp_server_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:udp_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:udp_server_profile"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "documentation/resources/application_profiles/properties/virtual_server/http3/index.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["virtual_server", "http3"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/http3/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.http3 for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.http3

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
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

- [client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/client_ssl_profile/): complete subsection reference.

- [http3_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/http3_profile/): complete subsection reference.

- [http_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/http_client_profile/): complete subsection reference.

- [http_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/http_server_profile/): complete subsection reference.

- [quic_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/quic_profile/): complete subsection reference.

- [server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/server_ssl_profile/): complete subsection reference.

- [tcp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/tcp_server_profile/): complete subsection reference.

- [udp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/udp_client_profile/): complete subsection reference.

- [udp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/udp_server_profile/): complete subsection reference.

## Next pages

- [virtual_server.http3.client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/client_ssl_profile/)
- [virtual_server.http3.http3_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/http3_profile/)
- [virtual_server.http3.http_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/http_client_profile/)
- [virtual_server.http3.http_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/http_server_profile/)
- [virtual_server.http3.quic_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/quic_profile/)
- [virtual_server.http3.server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/server_ssl_profile/)
- [virtual_server.http3.tcp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/tcp_server_profile/)
- [virtual_server.http3.udp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/udp_client_profile/)
- [virtual_server.http3.udp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/udp_server_profile/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
