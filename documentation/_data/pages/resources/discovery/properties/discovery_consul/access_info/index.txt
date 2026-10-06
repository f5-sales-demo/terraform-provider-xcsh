---
page_title: "discovery_consul.access_info"
subcategory: ""
description: "Hashicorp Consul API server information."
xcsh_docs: {"aliases": ["discovery consul access info"], "body_bytes": 1306, "body_sha256": "sha256:23805404fbf00840a873f654077269561c4ee469750fcfc34b2afae3c9546fb3", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info", "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_consul", "path": "documentation/resources/discovery/properties/discovery_consul/access_info/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2332031222030011-1013001033031322-3220012120232003-1031221221012300-2223030213002312-1333223111012301-0123302220223011-2032212310123232", "registry_path": "docs/guides/resources--discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_consul", "access_info"], "schema_version": 1, "sections": [{"aliases": ["discovery consul access info connection info"], "anchor": "section", "description": "Configuration details to access discovery service REST API.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-discovery_consul--access_info--connection_info--api_server", "enforcement": "provider-schema", "group": "discovery_consul.access_info.connection_info:RequiredObjectAttributes:api_server", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info", "type": "requires"}], "schema_path": ["discovery_consul", "access_info", "connection_info"], "syntax": "block", "type": "object"}, {"aliases": ["discovery consul access info http basic auth info"], "anchor": "section", "description": "Authentication parameters to access Hashicorp Consul.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_consul", "access_info", "http_basic_auth_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_consul/access_info/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Hashicorp Consul API server information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul.access_info

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/)
- discovery_consul.access_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Hashicorp Consul API server information.

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
access_info {
  # Configure direct properties listed below.
}
```

## Direct properties

- [connection_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/): complete subsection reference.

- [http_basic_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/): complete subsection reference.
