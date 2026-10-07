---
page_title: "virtual_server.udp"
subcategory: ""
description: "UDP profiles."
xcsh_docs: {"aliases": ["virtual server udp"], "body_bytes": 1592, "body_sha256": "sha256:d12bee01d5c8a439d4e5a4c74bd70aa95743db075f3de6bfe5df561b25c72b40", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp:client_ssl_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp:server_ssl_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp:udp_client_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp:udp_server_profile"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "documentation/data-sources/application_profiles/properties/virtual_server/udp/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3111012112013331-0210230112200321-2133130123230231-2102200122231133-2202101332332032-3132132213230031-3232202101331130-1313232102312102", "registry_path": "docs/guides/data-sources--application_profiles--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "udp"], "schema_version": 1, "sections": [{"aliases": ["virtual server udp client ssl profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp:client_ssl_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "udp", "client_ssl_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server udp server ssl profile"], "anchor": "section", "description": "Configuration parameter for server ssl profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp:server_ssl_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "udp", "server_ssl_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server udp udp client profile"], "anchor": "section", "description": "Client-side configuration", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp:udp_client_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "udp", "udp_client_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server udp udp server profile"], "anchor": "section", "description": "Configuration parameter for udp server profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp:udp_server_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "udp", "udp_server_profile"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/udp/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "UDP profiles.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["application_profilesCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
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
