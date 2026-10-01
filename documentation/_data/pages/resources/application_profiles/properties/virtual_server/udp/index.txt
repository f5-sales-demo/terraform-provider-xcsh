---
page_title: "virtual_server.udp"
subcategory: ""
description: "virtual_server.udp for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 2672, "body_sha256": "sha256:e62a31325cbffc6ca3f4d3ac35b917344d302891d8559aab78cddb9a753645bd", "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:udp:client_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:udp:server_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:udp:udp_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:udp:udp_server_profile"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:udp", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "documentation/resources/application_profiles/properties/virtual_server/udp/index.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["virtual_server", "udp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/udp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.udp for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.udp

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
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

- [client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/udp/client_ssl_profile/): complete subsection reference.

- [server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/udp/server_ssl_profile/): complete subsection reference.

- [udp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/udp/udp_client_profile/): complete subsection reference.

- [udp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/udp/udp_server_profile/): complete subsection reference.

## Next pages

- [virtual_server.udp.client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/udp/client_ssl_profile/)
- [virtual_server.udp.server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/udp/server_ssl_profile/)
- [virtual_server.udp.udp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/udp/udp_client_profile/)
- [virtual_server.udp.udp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/udp/udp_server_profile/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
