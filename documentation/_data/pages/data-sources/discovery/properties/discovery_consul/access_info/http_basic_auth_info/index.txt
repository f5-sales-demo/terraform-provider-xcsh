---
page_title: "discovery_consul.access_info.http_basic_auth_info"
subcategory: ""
description: "Authentication parameters to access Hashicorp Consul."
xcsh_docs: {"aliases": ["discovery consul access info http basic auth info"], "body_bytes": 2012, "body_sha256": "sha256:bc77745e0428402055cd2d39614c0f99acd8ed1f76e17c87fd9575e348b0f334", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:http_basic_auth_info", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info", "path": "documentation/data-sources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-3200201013223011-3322202301233100-1122211121001023-1322201313323131-0320133303212223-1112023300300321-3032030112132300-3123232213232003", "registry_path": "docs/guides/data-sources--discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_consul", "access_info", "http_basic_auth_info"], "schema_version": 1, "sections": [{"aliases": ["discovery consul access info http basic auth info passwd url"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_consul", "access_info", "http_basic_auth_info", "passwd_url"], "syntax": "attribute", "type": "object"}, {"aliases": ["discovery consul access info http basic auth info user name"], "anchor": "schema-discovery_consul--access_info--http_basic_auth_info--user_name", "description": "Username in consul.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:http_basic_auth_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_consul", "access_info", "http_basic_auth_info", "user_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Authentication parameters to access Hashicorp Consul.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["discoveryCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul.access_info.http_basic_auth_info

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/)
- [discovery_consul.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/access_info/)
- discovery_consul.access_info.http_basic_auth_info

<a id="section"></a>

Type: `"single"`. Computed.

Authentication parameters to access Hashicorp Consul.

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

- [passwd_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/): complete subsection reference.

<a id="schema-discovery_consul--access_info--http_basic_auth_info--user_name"></a>

### user_name property

Type: `"string"`. Computed.

User Name. Username in consul.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```
