---
page_title: "virtual_server.udp"
subcategory: ""
description: "UDP profiles."
xcsh_docs: {"aliases": ["virtual server udp"], "body_bytes": 2672, "body_sha256": "sha256:e62a31325cbffc6ca3f4d3ac35b917344d302891d8559aab78cddb9a753645bd", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:udp:client_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:udp:server_ssl_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:udp:udp_client_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:udp:udp_server_profile"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:udp", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "documentation/resources/application_profiles/properties/virtual_server/udp/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0100212231103000-2031131131020010-3003110221032012-0212301320022111-2133230201233310-2233210103212202-3223201320002311-1322101113222313", "registry_path": "docs/guides/resources--application_profiles--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "udp"], "schema_version": 1, "sections": [{"aliases": ["virtual server udp client ssl profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:udp:client_ssl_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "udp", "client_ssl_profile"], "syntax": "block", "type": "object"}, {"aliases": ["virtual server udp server ssl profile"], "anchor": "section", "description": "Configuration parameter for server ssl profile", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:udp:server_ssl_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "udp", "server_ssl_profile"], "syntax": "block", "type": "object"}, {"aliases": ["virtual server udp udp client profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:udp:udp_client_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "udp", "udp_client_profile"], "syntax": "block", "type": "object"}, {"aliases": ["virtual server udp udp server profile"], "anchor": "section", "description": "Configuration parameter for udp server profile", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:udp:udp_server_profile", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "udp", "udp_server_profile"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/udp/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "UDP profiles.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
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
