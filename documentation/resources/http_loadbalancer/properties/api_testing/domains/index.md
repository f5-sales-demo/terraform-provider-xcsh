---
page_title: "api_testing.domains"
subcategory: "Load Balancing"
description: "Add and configure testing domains and credentials."
xcsh_docs: {"aliases": ["api testing domains", "authentication", "credential setup", "credentials"], "body_bytes": 4014, "body_sha256": "sha256:5ac65ac9bcd6c75d669f968251ad3b7308b55d0e11963976d418e71faf04d4d5", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing", "path": "documentation/resources/http_loadbalancer/properties/api_testing/domains/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-010.md", "relationships": [{"anchor": "schema-api_testing--domains--domain", "enforcement": "provider-schema", "group": "api_testing.domains:RequiredListObjectAttributes:credentials,domain", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains:RequiredListObjectAttributes:credentials,domain", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_testing", "domains"], "schema_version": 1, "sections": [{"aliases": ["allow destructive methods"], "anchor": "schema-api_testing--domains--allow_destructive_methods", "description": "Enable to allow API Testing to execute against destructive methods. Use with caution as these may modify or DELETE data.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "domains", "allow_destructive_methods"], "syntax": "attribute", "type": "bool"}, {"aliases": ["authentication", "credential setup", "credentials"], "anchor": "section", "description": "Add credentials for API testing to use in the selected environment.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains.credentials:ConflictingListObjectAttributes:admin,standard", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:admin", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains.credentials:ConflictingListObjectAttributes:api_key,basic_auth", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:api_key", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains.credentials:ConflictingListObjectAttributes:api_key,bearer_token", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:api_key", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains.credentials:ConflictingListObjectAttributes:api_key,login_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:api_key", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains.credentials:ConflictingListObjectAttributes:api_key,basic_auth", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:basic_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains.credentials:ConflictingListObjectAttributes:basic_auth,bearer_token", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:basic_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains.credentials:ConflictingListObjectAttributes:basic_auth,login_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:basic_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains.credentials:ConflictingListObjectAttributes:api_key,bearer_token", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains.credentials:ConflictingListObjectAttributes:basic_auth,bearer_token", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains.credentials:ConflictingListObjectAttributes:bearer_token,login_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains.credentials:ConflictingListObjectAttributes:api_key,login_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains.credentials:ConflictingListObjectAttributes:basic_auth,login_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains.credentials:ConflictingListObjectAttributes:bearer_token,login_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains.credentials:ConflictingListObjectAttributes:admin,standard", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:standard", "type": "conflicts"}, {"anchor": "schema-api_testing--domains--credentials--credential_name", "enforcement": "provider-schema", "group": "api_testing.domains.credentials:RequiredListObjectAttributes:credential_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials", "type": "requires"}], "schema_path": ["api_testing", "domains", "credentials"], "syntax": "block", "type": "object"}, {"aliases": ["domain"], "anchor": "schema-api_testing--domains--domain", "description": "Add your testing environment domain. Be aware that running tests on a production domain can impact live applications, as API testing cannot distinguish between production and testing environments.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "domains", "domain"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_testing/domains/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Add and configure testing domains and credentials.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_testing.domains

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/)
- api_testing.domains

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Add and configure testing domains and credentials.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("credentials",
    "domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
domains {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-api_testing--domains--allow_destructive_methods"></a>

### allow_destructive_methods property

Type: `"bool"`. Optional.

Enable to allow API Testing to execute against destructive methods. Use with caution as these may
modify or DELETE data.

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

- [credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/): complete subsection reference.

<a id="schema-api_testing--domains--domain"></a>

### domain property

Type: `"string"`. Optional.

Add your testing environment domain. Be aware that running tests on a production domain can impact
live applications, as API testing cannot distinguish between production and testing environments.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

## Next pages

- [api_testing.domains.credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/)
- [api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
