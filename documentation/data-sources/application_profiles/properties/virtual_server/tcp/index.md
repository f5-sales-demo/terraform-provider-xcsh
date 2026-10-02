---
page_title: "virtual_server.tcp"
subcategory: ""
description: "TCP profiles."
xcsh_docs: {"aliases": ["virtual server tcp"], "body_bytes": 2955, "body_sha256": "sha256:b89d0ee5409232777171483b5ac1da83c0e227570d778d41f5b634dc8a0a832b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:client_ssl_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:ocsp_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:server_ssl_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:tcp_client_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:tcp_server_profile"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "documentation/data-sources/application_profiles/properties/virtual_server/tcp/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101", "registry_path": "docs/guides/data-sources--application_profiles--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "tcp"], "schema_version": 1, "sections": [{"aliases": ["client ssl profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:client_ssl_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "client_ssl_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["ocsp profile"], "anchor": "section", "description": "Configuration parameter for ocsp profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:ocsp_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "ocsp_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["server ssl profile"], "anchor": "section", "description": "Configuration parameter for server ssl profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:server_ssl_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "server_ssl_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["tcp client profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:tcp_client_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "tcp_client_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["tcp server profile"], "anchor": "section", "description": "Configuration parameter for tcp server profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp:tcp_server_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "tcp", "tcp_server_profile"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/tcp/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "TCP profiles.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["application_profilesCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [virtual_server.tcp.client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/client_ssl_profile/)
- [virtual_server.tcp.ocsp_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/ocsp_profile/)
- [virtual_server.tcp.server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/server_ssl_profile/)
- [virtual_server.tcp.tcp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/tcp_client_profile/)
- [virtual_server.tcp.tcp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/tcp_server_profile/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
