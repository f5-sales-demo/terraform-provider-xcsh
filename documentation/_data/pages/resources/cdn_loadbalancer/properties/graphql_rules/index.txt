---
page_title: "graphql_rules"
subcategory: "Load Balancing"
description: "GraphQL is a query language and server-side runtime for APIs which provides a complete and understandable description of the data in API. GraphQL gives clients the power to ask for exactly what they need, makes it easier to evolve APIs over time, and enables powerful developer tools. Policy configuration to analyze"
xcsh_docs: {"aliases": ["graphql rules"], "body_bytes": 8267, "body_sha256": "sha256:78efeed402069a465309f8c4124c0b0ca2fce0705b127aaf5c60c2d96d43726b", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:any_domain", "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:graphql_settings", "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:metadata", "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:method_get", "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:method_post"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/graphql_rules/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-010.md", "relationships": [{"anchor": "schema-graphql_rules--exact_value", "enforcement": "provider-schema", "group": "graphql_rules:ConflictingListObjectAttributes:any_domain,exact_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules", "type": "conflicts"}, {"anchor": "schema-graphql_rules--exact_value", "enforcement": "provider-schema", "group": "graphql_rules:ConflictingListObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules", "type": "conflicts"}, {"anchor": "schema-graphql_rules--suffix_value", "enforcement": "provider-schema", "group": "graphql_rules:ConflictingListObjectAttributes:any_domain,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules", "type": "conflicts"}, {"anchor": "schema-graphql_rules--suffix_value", "enforcement": "provider-schema", "group": "graphql_rules:ConflictingListObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "graphql_rules:ConflictingListObjectAttributes:any_domain,exact_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "graphql_rules:ConflictingListObjectAttributes:any_domain,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "graphql_rules:ConflictingListObjectAttributes:method_get,method_post", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:method_get", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "graphql_rules:ConflictingListObjectAttributes:method_get,method_post", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:method_post", "type": "conflicts"}, {"anchor": "schema-graphql_rules--exact_path", "enforcement": "provider-schema", "group": "graphql_rules:RequiredListObjectAttributes:exact_path", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["graphql_rules"], "schema_version": 1, "sections": [{"aliases": ["any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:any_domain", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["graphql_rules", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["exact path"], "anchor": "schema-graphql_rules--exact_path", "description": "Specifies the exact path to GraphQL endpoint. Default value is /graphql.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["graphql_rules", "exact_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["exact value"], "anchor": "schema-graphql_rules--exact_value", "description": "Exclusive with Exact domain name.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["graphql_rules", "exact_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["graphql settings"], "anchor": "section", "description": "GraphQL configuration.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:graphql_settings", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "graphql_rules.graphql_settings:ConflictingObjectAttributes:disable_introspection,enable_introspection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:graphql_settings:disable_introspection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "graphql_rules.graphql_settings:ConflictingObjectAttributes:disable_introspection,enable_introspection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:graphql_settings:enable_introspection", "type": "conflicts"}, {"anchor": "schema-graphql_rules--graphql_settings--max_batched_queries", "enforcement": "provider-schema", "group": "graphql_rules.graphql_settings:RequiredObjectAttributes:max_batched_queries,max_depth,max_total_length", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:graphql_settings", "type": "requires"}, {"anchor": "schema-graphql_rules--graphql_settings--max_depth", "enforcement": "provider-schema", "group": "graphql_rules.graphql_settings:RequiredObjectAttributes:max_batched_queries,max_depth,max_total_length", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:graphql_settings", "type": "requires"}, {"anchor": "schema-graphql_rules--graphql_settings--max_total_length", "enforcement": "provider-schema", "group": "graphql_rules.graphql_settings:RequiredObjectAttributes:max_batched_queries,max_depth,max_total_length", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:graphql_settings", "type": "requires"}], "schema_path": ["graphql_rules", "graphql_settings"], "syntax": "block", "type": "object"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-graphql_rules--metadata--name", "enforcement": "provider-schema", "group": "graphql_rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:metadata", "type": "requires"}], "schema_path": ["graphql_rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["method get"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:method_get", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["graphql_rules", "method_get"], "syntax": "attribute", "type": "object"}, {"aliases": ["method post"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:method_post", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["graphql_rules", "method_post"], "syntax": "attribute", "type": "object"}, {"aliases": ["suffix value"], "anchor": "schema-graphql_rules--suffix_value", "description": "Exclusive with Suffix of domain name e.g \"xyz.com\" will match \"*.xyz.com\" and \"xyz.com\"", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["graphql_rules", "suffix_value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/graphql_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "GraphQL is a query language and server-side runtime for APIs which provides a complete and understandable description of the data in API. GraphQL gives clients the power to ask for exactly what they need, makes it easier to evolve APIs over time, and enables powerful developer tools. Policy configuration to analyze", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# graphql_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- graphql_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy..

Upstream description:

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy configuration to analyze GraphQL queries and prevent GraphQL tailored attacks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("exact_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "exact_value"),
  validators.ConflictingListObjectAttributes("any_domain",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("method_get",
    "method_post")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
graphql_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/graphql_rules/any_domain/): complete subsection reference.

<a id="schema-graphql_rules--exact_path"></a>

### exact_path property

Type: `"string"`. Optional.

Specifies the exact path to GraphQL endpoint. Defaults to \`/graphql\`.

Upstream description:

Specifies the exact path to GraphQL endpoint. Default value is /graphql.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-graphql_rules--exact_value"></a>

### exact_value property

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [graphql_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/graphql_rules/graphql_settings/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/graphql_rules/metadata/): complete subsection reference.

- [method_get](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/graphql_rules/method_get/): complete subsection reference.

- [method_post](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/graphql_rules/method_post/): complete subsection reference.

<a id="schema-graphql_rules--suffix_value"></a>

### suffix_value property

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

## Next pages

- [graphql_rules.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/graphql_rules/any_domain/)
- [graphql_rules.graphql_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/graphql_rules/graphql_settings/)
- [graphql_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/graphql_rules/metadata/)
- [graphql_rules.method_get](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/graphql_rules/method_get/)
- [graphql_rules.method_post](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/graphql_rules/method_post/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
