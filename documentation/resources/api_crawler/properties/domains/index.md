---
page_title: "domains"
subcategory: ""
description: "API Crawler Configuration."
xcsh_docs: {"aliases": ["domains"], "body_bytes": 3224, "body_sha256": "sha256:4e61620a3c10a8329d3106bf8964eb933fd7013e0bc57749b2d34ad400d60d27", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:api_crawler:properties:domains:simple_login"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_crawler:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_crawler:properties:domains", "parent_id": "xcsh-docs:resources:api_crawler:reference", "path": "documentation/resources/api_crawler/properties/domains/index.md", "product": "distributed-cloud", "provider_name": "api_crawler", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2023202200302200-3220100302000201-1103011010332303-2230201122011301-2110202200310302-0233212012122331-1231130120112020-1311130230231300", "registry_path": "docs/guides/resources--api_crawler--reference--group-001.md", "relationships": [{"anchor": "schema-domains--domain", "enforcement": "provider-schema", "group": "domains:RequiredListObjectAttributes:domain", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:api_crawler:properties:domains", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["domains"], "schema_version": 1, "sections": [{"aliases": ["domain"], "anchor": "schema-domains--domain", "description": "Select the domain to execute API Crawling with given credentials.", "document_id": "xcsh-docs:resources:api_crawler:properties:domains", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "domain"], "syntax": "attribute", "type": "string"}, {"aliases": ["login", "login result", "sign in", "simple login"], "anchor": "section", "description": "Configuration parameter for simple login.", "document_id": "xcsh-docs:resources:api_crawler:properties:domains:simple_login", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["domains", "simple_login"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_crawler/properties/domains/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "API Crawler Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["api_crawlerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains

Breadcrumbs:

- [xcsh_api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/)
- domains

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

API Crawler. API Crawler Configuration.

Upstream description:

API Crawler Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("domain")}
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

<a id="schema-domains--domain"></a>

### domain property

Type: `"string"`. Optional.

Select the domain to execute API Crawling with given credentials.

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

- [simple_login](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/simple_login/): complete subsection reference.

## Next pages

- [domains.simple_login](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/simple_login/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/)
- [xcsh_api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/)
