---
page_title: "virtual_server.http3"
subcategory: ""
description: "HTTP/3 profiles."
xcsh_docs: {"aliases": ["virtual server http3"], "body_bytes": 4569, "body_sha256": "sha256:81e17e4cb82d6f27b7559897257ce23e4737bc992116ff16efbc7b07121e48f0", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:http3:client_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:http3_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:http_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:http_server_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:quic_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:server_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:tcp_server_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:udp_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:udp_server_profile"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "documentation/resources/application_profiles/properties/virtual_server/http3/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132", "registry_path": "docs/guides/resources--application_profiles--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "http3"], "schema_version": 1, "sections": [{"aliases": ["client ssl profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:client_ssl_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "http3", "client_ssl_profile"], "syntax": "block", "type": "object"}, {"aliases": ["http3 profile"], "anchor": "section", "description": "Configuration parameter for http3 profile", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:http3_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "http3", "http3_profile"], "syntax": "block", "type": "object"}, {"aliases": ["http client profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:http_client_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "http3", "http_client_profile"], "syntax": "block", "type": "object"}, {"aliases": ["http server profile"], "anchor": "section", "description": "Configuration parameter for http server profile", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:http_server_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "http3", "http_server_profile"], "syntax": "block", "type": "object"}, {"aliases": ["quic profile"], "anchor": "section", "description": "Configuration parameter for quic profile", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:quic_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "http3", "quic_profile"], "syntax": "block", "type": "object"}, {"aliases": ["server ssl profile"], "anchor": "section", "description": "Configuration parameter for server ssl profile", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:server_ssl_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "http3", "server_ssl_profile"], "syntax": "block", "type": "object"}, {"aliases": ["tcp server profile"], "anchor": "section", "description": "Configuration parameter for tcp server profile", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:tcp_server_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "http3", "tcp_server_profile"], "syntax": "block", "type": "object"}, {"aliases": ["udp client profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:udp_client_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "http3", "udp_client_profile"], "syntax": "block", "type": "object"}, {"aliases": ["udp server profile"], "anchor": "section", "description": "Configuration parameter for udp server profile", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:http3:udp_server_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "http3", "udp_server_profile"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/http3/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "HTTP/3 profiles.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
