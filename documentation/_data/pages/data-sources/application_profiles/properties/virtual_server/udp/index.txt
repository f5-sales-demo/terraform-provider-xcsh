---
page_title: "virtual_server.udp"
subcategory: ""
description: "UDP profiles."
xcsh_docs: {"aliases": ["virtual server udp"], "body_bytes": 2599, "body_sha256": "sha256:c8353537e4c98437269737d03bf8656d616e3bc7dabbae468e58e32017366671", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp:client_ssl_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp:server_ssl_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp:udp_client_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp:udp_server_profile"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "documentation/data-sources/application_profiles/properties/virtual_server/udp/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3111012112013331-0210230112200321-2133130123230231-2102200122231133-2202101332332032-3132132213230031-3232202101331130-1313232102312102", "registry_path": "docs/guides/data-sources--application_profiles--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "udp"], "schema_version": 1, "sections": [{"aliases": ["client ssl profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp:client_ssl_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "udp", "client_ssl_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["server ssl profile"], "anchor": "section", "description": "Configuration parameter for server ssl profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp:server_ssl_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "udp", "server_ssl_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["udp client profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp:udp_client_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "udp", "udp_client_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["udp server profile"], "anchor": "section", "description": "Configuration parameter for udp server profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp:udp_server_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "udp", "udp_server_profile"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/udp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "UDP profiles.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.udp

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- virtual_server.udp

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/client_ssl_profile/): complete subsection reference.

- [server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/server_ssl_profile/): complete subsection reference.

- [udp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/udp_client_profile/): complete subsection reference.

- [udp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/udp_server_profile/): complete subsection reference.

## Next pages

- [virtual_server.udp.client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/client_ssl_profile/)
- [virtual_server.udp.server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/server_ssl_profile/)
- [virtual_server.udp.udp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/udp_client_profile/)
- [virtual_server.udp.udp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/udp_server_profile/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
