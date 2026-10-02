---
page_title: "virtual_server.http3"
subcategory: ""
description: "HTTP/3 profiles."
xcsh_docs: {"aliases": ["virtual server http3"], "body_bytes": 4524, "body_sha256": "sha256:cad688bfe967ef883d79556c08272a6ad8ec0dfae14dbc172825d9d356323d94", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3:client_ssl_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3:http3_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3:http_client_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3:http_server_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3:quic_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3:server_ssl_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3:tcp_server_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3:udp_client_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3:udp_server_profile"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "documentation/data-sources/application_profiles/properties/virtual_server/http3/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2233213302301031-1011021100023220-0023001210122232-0332111101220031-0312121002130112-0330200230012032-0313003311311113-0322012000031213", "registry_path": "docs/guides/data-sources--application_profiles--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "http3"], "schema_version": 1, "sections": [{"aliases": ["client ssl profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3:client_ssl_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "http3", "client_ssl_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["http3 profile"], "anchor": "section", "description": "Configuration parameter for http3 profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3:http3_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "http3", "http3_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["http client profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3:http_client_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "http3", "http_client_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["http server profile"], "anchor": "section", "description": "Configuration parameter for http server profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3:http_server_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "http3", "http_server_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["quic profile"], "anchor": "section", "description": "Configuration parameter for quic profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3:quic_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "http3", "quic_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["server ssl profile"], "anchor": "section", "description": "Configuration parameter for server ssl profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3:server_ssl_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "http3", "server_ssl_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["tcp server profile"], "anchor": "section", "description": "Configuration parameter for tcp server profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3:tcp_server_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "http3", "tcp_server_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["udp client profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3:udp_client_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "http3", "udp_client_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["udp server profile"], "anchor": "section", "description": "Configuration parameter for udp server profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3:udp_server_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "http3", "udp_server_profile"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/http3/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "HTTP/3 profiles.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["application_profilesCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.http3

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- virtual_server.http3

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/client_ssl_profile/): complete subsection reference.

- [http3_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http3_profile/): complete subsection reference.

- [http_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http_client_profile/): complete subsection reference.

- [http_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http_server_profile/): complete subsection reference.

- [quic_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/quic_profile/): complete subsection reference.

- [server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/server_ssl_profile/): complete subsection reference.

- [tcp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/tcp_server_profile/): complete subsection reference.

- [udp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/udp_client_profile/): complete subsection reference.

- [udp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/udp_server_profile/): complete subsection reference.

## Next pages

- [virtual_server.http3.client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/client_ssl_profile/)
- [virtual_server.http3.http3_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http3_profile/)
- [virtual_server.http3.http_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http_client_profile/)
- [virtual_server.http3.http_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http_server_profile/)
- [virtual_server.http3.quic_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/quic_profile/)
- [virtual_server.http3.server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/server_ssl_profile/)
- [virtual_server.http3.tcp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/tcp_server_profile/)
- [virtual_server.http3.udp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/udp_client_profile/)
- [virtual_server.http3.udp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/udp_server_profile/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
