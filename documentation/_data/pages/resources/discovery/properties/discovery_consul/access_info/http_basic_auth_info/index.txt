---
page_title: "discovery_consul.access_info.http_basic_auth_info"
subcategory: ""
description: "Authentication parameters to access Hashicorp Consul."
xcsh_docs: {"aliases": ["discovery consul access info http basic auth info"], "body_bytes": 2292, "body_sha256": "sha256:1253f8e884f43b1c810c36a8412786542cef8ed27ae739e28868944b02c8d892", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info", "path": "documentation/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0001133210323132-2031212030200231-2311331203333003-0221103130123101-1021100123320222-0102311230120010-1211202323201033-1230230013200100", "registry_path": "docs/guides/resources--discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_consul", "access_info", "http_basic_auth_info"], "schema_version": 1, "sections": [{"aliases": ["discovery consul access info http basic auth info passwd url"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_consul.access_info.http_basic_auth_info.passwd_url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_consul.access_info.http_basic_auth_info.passwd_url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url:clear_secret_info", "type": "conflicts"}], "schema_path": ["discovery_consul", "access_info", "http_basic_auth_info", "passwd_url"], "syntax": "block", "type": "object"}, {"aliases": ["discovery consul access info http basic auth info user name"], "anchor": "schema-discovery_consul--access_info--http_basic_auth_info--user_name", "description": "Username in consul.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_consul", "access_info", "http_basic_auth_info", "user_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Authentication parameters to access Hashicorp Consul.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["discoveryCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul.access_info.http_basic_auth_info

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/)
- [discovery_consul.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/)
- discovery_consul.access_info.http_basic_auth_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_basic_auth_info {
  # Configure direct properties listed below.
}
```

## Direct properties

- [passwd_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/): complete subsection reference.

<a id="schema-discovery_consul--access_info--http_basic_auth_info--user_name"></a>

### user_name property

Type: `"string"`. Optional.

User Name. Username in consul.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
