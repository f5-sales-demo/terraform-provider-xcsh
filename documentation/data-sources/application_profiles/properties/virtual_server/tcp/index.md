---
page_title: "virtual_server.tcp"
subcategory: ""
description: "TCP profiles."
xcsh_docs: {"aliases": ["virtual server tcp"], "body_bytes": 1777, "body_sha256": "sha256:a696b29ca9e14a69a3593fc82a60917019831e80ff9bdd3227c61df9e87fd0fc", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:client_ssl_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:ocsp_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:server_ssl_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:tcp_client_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:tcp_server_profile"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "documentation/data-sources/application_profiles/properties/virtual_server/tcp/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101", "registry_path": "docs/guides/data-sources--application_profiles--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "tcp"], "schema_version": 1, "sections": [{"aliases": ["virtual server tcp client ssl profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:client_ssl_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "client_ssl_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server tcp ocsp profile"], "anchor": "section", "description": "Configuration parameter for ocsp profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:ocsp_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "ocsp_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server tcp server ssl profile"], "anchor": "section", "description": "Configuration parameter for server ssl profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:server_ssl_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "server_ssl_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server tcp tcp client profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:tcp_client_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "tcp_client_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server tcp tcp server profile"], "anchor": "section", "description": "Configuration parameter for tcp server profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:tcp_server_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "tcp_server_profile"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/tcp/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "TCP profiles.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["application_profilesCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.tcp

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- virtual_server.tcp

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/client_ssl_profile/): complete subsection reference.

- [ocsp_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/ocsp_profile/): complete subsection reference.

- [server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/server_ssl_profile/): complete subsection reference.

- [tcp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/tcp_client_profile/): complete subsection reference.

- [tcp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/tcp_server_profile/): complete subsection reference.
