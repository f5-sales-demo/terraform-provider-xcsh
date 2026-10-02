---
page_title: "discovery_consul.access_info.http_basic_auth_info"
subcategory: ""
description: "Authentication parameters to access Hashicorp Consul."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials", "discovery consul access info http basic auth info"], "body_bytes": 2786, "body_sha256": "sha256:1b8e3398375b61313347e0070d0ef08daf357037ac2c7f5fcf4dc8ca4915c5ac", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info", "path": "documentation/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0001133210323132-2031212030200231-2311331203333003-0221103130123101-1021100123320222-0102311230120010-1211202323201033-1230230013200100", "registry_path": "docs/guides/resources--discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_consul", "access_info", "http_basic_auth_info"], "schema_version": 1, "sections": [{"aliases": ["passwd url"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_consul.access_info.http_basic_auth_info.passwd_url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_consul.access_info.http_basic_auth_info.passwd_url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url:clear_secret_info", "type": "conflicts"}], "schema_path": ["discovery_consul", "access_info", "http_basic_auth_info", "passwd_url"], "syntax": "block", "type": "object"}, {"aliases": ["user name"], "anchor": "schema-discovery_consul--access_info--http_basic_auth_info--user_name", "description": "Username in consul.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_consul", "access_info", "http_basic_auth_info", "user_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Authentication parameters to access Hashicorp Consul.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Upstream description:

Username in consul.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [discovery_consul.access_info.http_basic_auth_info.passwd_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/)
- [discovery_consul.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
