---
page_title: "virtual_server.tcp"
subcategory: ""
description: "TCP profiles."
xcsh_docs: {"aliases": ["virtual server tcp"], "body_bytes": 1865, "body_sha256": "sha256:81b77e2bec32ff5c5c320434f3cc362d77e2858a6aac40e758a2f5c3ce958fcd", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:client_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:ocsp_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:server_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:tcp_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:tcp_server_profile"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "documentation/resources/application_profiles/properties/virtual_server/tcp/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231", "registry_path": "docs/guides/resources--application_profiles--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "tcp"], "schema_version": 1, "sections": [{"aliases": ["virtual server tcp client ssl profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:client_ssl_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "client_ssl_profile"], "syntax": "block", "type": "object"}, {"aliases": ["virtual server tcp ocsp profile"], "anchor": "section", "description": "Configuration parameter for ocsp profile", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:ocsp_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "ocsp_profile"], "syntax": "block", "type": "object"}, {"aliases": ["virtual server tcp server ssl profile"], "anchor": "section", "description": "Configuration parameter for server ssl profile", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:server_ssl_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "server_ssl_profile"], "syntax": "block", "type": "object"}, {"aliases": ["virtual server tcp tcp client profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:tcp_client_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "tcp_client_profile"], "syntax": "block", "type": "object"}, {"aliases": ["virtual server tcp tcp server profile"], "anchor": "section", "description": "Configuration parameter for tcp server profile", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp:tcp_server_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "tcp_server_profile"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/tcp/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "TCP profiles.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["application_profilesCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
