---
page_title: "virtual_server.tcp"
subcategory: ""
description: "TCP profiles."
xcsh_docs: {"aliases": ["virtual server tcp"], "body_bytes": 3022, "body_sha256": "sha256:45b9548cdd4623df1baead771de5350c1dfbebd990f23b0a79a062b5365c9e2f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:client_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:ocsp_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:server_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:tcp_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:tcp_server_profile"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "documentation/resources/application_profiles/properties/virtual_server/tcp/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231", "registry_path": "docs/guides/resources--application_profiles--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "tcp"], "schema_version": 1, "sections": [{"aliases": ["client ssl profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:client_ssl_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "client_ssl_profile"], "syntax": "block", "type": "object"}, {"aliases": ["ocsp profile"], "anchor": "section", "description": "Configuration parameter for ocsp profile", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:ocsp_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "ocsp_profile"], "syntax": "block", "type": "object"}, {"aliases": ["server ssl profile"], "anchor": "section", "description": "Configuration parameter for server ssl profile", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:server_ssl_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "server_ssl_profile"], "syntax": "block", "type": "object"}, {"aliases": ["tcp client profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:tcp_client_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "tcp_client_profile"], "syntax": "block", "type": "object"}, {"aliases": ["tcp server profile"], "anchor": "section", "description": "Configuration parameter for tcp server profile", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:tcp_server_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "tcp_server_profile"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/tcp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "TCP profiles.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.tcp

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
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

- [client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/tcp/client_ssl_profile/): complete subsection reference.

- [ocsp_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/tcp/ocsp_profile/): complete subsection reference.

- [server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/tcp/server_ssl_profile/): complete subsection reference.

- [tcp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/tcp/tcp_client_profile/): complete subsection reference.

- [tcp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/tcp/tcp_server_profile/): complete subsection reference.

## Next pages

- [virtual_server.tcp.client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/tcp/client_ssl_profile/)
- [virtual_server.tcp.ocsp_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/tcp/ocsp_profile/)
- [virtual_server.tcp.server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/tcp/server_ssl_profile/)
- [virtual_server.tcp.tcp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/tcp/tcp_client_profile/)
- [virtual_server.tcp.tcp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/tcp/tcp_server_profile/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
