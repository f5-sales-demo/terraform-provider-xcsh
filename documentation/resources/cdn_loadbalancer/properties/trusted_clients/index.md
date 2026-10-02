---
page_title: "trusted_clients"
subcategory: "Load Balancing"
description: "Define rules to skip processing of one or more features such as WAF, Bot Defense etc. For clients."
xcsh_docs: {"aliases": ["trusted clients"], "body_bytes": 12204, "body_sha256": "sha256:0e11938dfc445c062b6fffe07e5e2b335cf24eb0c153908072ef03fd4e18686e", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:bot_skip_processing", "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:http_header", "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:metadata", "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:skip_processing", "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:waf_skip_processing"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/trusted_clients/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-015.md", "relationships": [{"anchor": "schema-trusted_clients--as_number", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:as_number,http_header", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "type": "conflicts"}, {"anchor": "schema-trusted_clients--as_number", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:as_number,ip_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "type": "conflicts"}, {"anchor": "schema-trusted_clients--as_number", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:as_number,ipv6_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "type": "conflicts"}, {"anchor": "schema-trusted_clients--as_number", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:as_number,user_identifier", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "type": "conflicts"}, {"anchor": "schema-trusted_clients--ip_prefix", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:as_number,ip_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "type": "conflicts"}, {"anchor": "schema-trusted_clients--ip_prefix", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:http_header,ip_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "type": "conflicts"}, {"anchor": "schema-trusted_clients--ip_prefix", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:ip_prefix,ipv6_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "type": "conflicts"}, {"anchor": "schema-trusted_clients--ip_prefix", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:ip_prefix,user_identifier", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "type": "conflicts"}, {"anchor": "schema-trusted_clients--ipv6_prefix", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:as_number,ipv6_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "type": "conflicts"}, {"anchor": "schema-trusted_clients--ipv6_prefix", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:http_header,ipv6_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "type": "conflicts"}, {"anchor": "schema-trusted_clients--ipv6_prefix", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:ip_prefix,ipv6_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "type": "conflicts"}, {"anchor": "schema-trusted_clients--ipv6_prefix", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:ipv6_prefix,user_identifier", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "type": "conflicts"}, {"anchor": "schema-trusted_clients--user_identifier", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:as_number,user_identifier", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "type": "conflicts"}, {"anchor": "schema-trusted_clients--user_identifier", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:http_header,user_identifier", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "type": "conflicts"}, {"anchor": "schema-trusted_clients--user_identifier", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:ip_prefix,user_identifier", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "type": "conflicts"}, {"anchor": "schema-trusted_clients--user_identifier", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:ipv6_prefix,user_identifier", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:bot_skip_processing,skip_processing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:bot_skip_processing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:bot_skip_processing,waf_skip_processing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:bot_skip_processing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:as_number,http_header", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:http_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:http_header,ip_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:http_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:http_header,ipv6_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:http_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:http_header,user_identifier", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:http_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:bot_skip_processing,skip_processing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:skip_processing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:skip_processing,waf_skip_processing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:skip_processing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:bot_skip_processing,waf_skip_processing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:waf_skip_processing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "trusted_clients:ConflictingListObjectAttributes:skip_processing,waf_skip_processing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:waf_skip_processing", "type": "conflicts"}, {"anchor": "schema-trusted_clients--actions", "enforcement": "provider-schema", "group": "trusted_clients:RequiredListObjectAttributes:actions", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["trusted_clients"], "schema_version": 1, "sections": [{"aliases": ["actions"], "anchor": "schema-trusted_clients--actions", "description": "Actions that should be taken when client identifier matches the rule.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["trusted_clients", "actions"], "syntax": "attribute", "type": "list"}, {"aliases": ["as number"], "anchor": "schema-trusted_clients--as_number", "description": "Exclusive with RFC 6793 defined 4-byte AS number.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["trusted_clients", "as_number"], "syntax": "attribute", "type": "number"}, {"aliases": ["bot skip processing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:bot_skip_processing", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["trusted_clients", "bot_skip_processing"], "syntax": "attribute", "type": "object"}, {"aliases": ["expiration timestamp"], "anchor": "schema-trusted_clients--expiration_timestamp", "description": "The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["trusted_clients", "expiration_timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["http header"], "anchor": "section", "description": "Request header name and value pairs.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:http_header", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "trusted_clients.http_header:RequiredObjectAttributes:headers", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:http_header:headers", "type": "requires"}], "schema_path": ["trusted_clients", "http_header"], "syntax": "block", "type": "object"}, {"aliases": ["ip prefix"], "anchor": "schema-trusted_clients--ip_prefix", "description": "Exclusive with IPv4 prefix string.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["trusted_clients", "ip_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipv6 prefix"], "anchor": "schema-trusted_clients--ipv6_prefix", "description": "Exclusive with IPv6 prefix string.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["trusted_clients", "ipv6_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-trusted_clients--metadata--name", "enforcement": "provider-schema", "group": "trusted_clients.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:metadata", "type": "requires"}], "schema_path": ["trusted_clients", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["skip processing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:skip_processing", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["trusted_clients", "skip_processing"], "syntax": "attribute", "type": "object"}, {"aliases": ["user identifier"], "anchor": "schema-trusted_clients--user_identifier", "description": "Exclusive with Identify user based on user identifier. User identifier value needs to be copied from security event.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["trusted_clients", "user_identifier"], "syntax": "attribute", "type": "string"}, {"aliases": ["waf skip processing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:trusted_clients:waf_skip_processing", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["trusted_clients", "waf_skip_processing"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/trusted_clients/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Define rules to skip processing of one or more features such as WAF, Bot Defense etc. For clients.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# trusted_clients

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- trusted_clients

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Define rules to skip processing of one or more features such as WAF, Bot Defense etc.

Upstream description:

Define rules to skip processing of one or more features such as WAF, Bot Defense etc. For clients.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("actions"),
  validators.ConflictingListObjectAttributes("as_number",
    "http_header"),
  validators.ConflictingListObjectAttributes("as_number",
    "ip_prefix"),
  validators.ConflictingListObjectAttributes("as_number",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("as_number",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("bot_skip_processing",
    "skip_processing"),
  validators.ConflictingListObjectAttributes("bot_skip_processing",
    "waf_skip_processing"),
  validators.ConflictingListObjectAttributes("http_header",
    "ip_prefix"),
  validators.ConflictingListObjectAttributes("http_header",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("http_header",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("ip_prefix",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("ip_prefix",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("ipv6_prefix",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("skip_processing",
    "waf_skip_processing")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
trusted_clients {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-trusted_clients--actions"></a>

### actions property

Type: `["list", "string"]`. Optional.

\[Enum:
SKIP\_PROCESSING\_WAF|SKIP\_PROCESSING\_BOT|SKIP\_PROCESSING\_MUM|SKIP\_PROCESSING\_IP\_REPUTATION|SKIP\_PROCESSING\_API\_PROTECTION|SKIP\_PROCESSING\_OAS\_VALIDATION|SKIP\_PROCESSING\_DDOS\_PROTECTION|SKIP\_PROCESSING\_THREAT\_MESH|SKIP\_PROCESSING\_MALWARE\_PROTECTION\]
Actions that should be taken when client identifier matches the rule. Possible values are
\`SKIP\_PROCESSING\_WAF\`, \`SKIP\_PROCESSING\_BOT\`, \`SKIP\_PROCESSING\_MUM\`,
\`SKIP\_PROCESSING\_IP\_REPUTATION\`, \`SKIP\_PROCESSING\_API\_PROTECTION\`,
\`SKIP\_PROCESSING\_OAS\_VALIDATION\`, \`SKIP\_PROCESSING\_DDOS\_PROTECTION\`,
\`SKIP\_PROCESSING\_THREAT\_MESH\`, \`SKIP\_PROCESSING\_MALWARE\_PROTECTION\`. Defaults to
\`SKIP\_PROCESSING\_WAF\`.

Upstream description:

Actions that should be taken when client identifier matches the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(10),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-trusted_clients--as_number"></a>

### as_number property

Type: `"number"`. Optional.

Exclusive with \[http\_header ip\_prefix ipv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Upstream description:

Exclusive with \[http\_header ip\_prefix ipv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 401308),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 401308,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  }
}
```

- [bot_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/trusted_clients/bot_skip_processing/): complete subsection reference.

<a id="schema-trusted_clients--expiration_timestamp"></a>

### expiration_timestamp property

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [http_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/trusted_clients/http_header/): complete subsection reference.

<a id="schema-trusted_clients--ip_prefix"></a>

### ip_prefix property

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ipv6\_prefix user\_identifier\] IPv4 prefix string.

Upstream description:

Exclusive with \[as\_number http\_header ipv6\_prefix user\_identifier\] IPv4 prefix string.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="schema-trusted_clients--ipv6_prefix"></a>

### ipv6_prefix property

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/trusted_clients/metadata/): complete subsection reference.

- [skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/trusted_clients/skip_processing/): complete subsection reference.

<a id="schema-trusted_clients--user_identifier"></a>

### user_identifier property

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ip\_prefix ipv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix ipv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [waf_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/trusted_clients/waf_skip_processing/): complete subsection reference.

## Next pages

- [trusted_clients.bot_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/trusted_clients/bot_skip_processing/)
- [trusted_clients.http_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/trusted_clients/http_header/)
- [trusted_clients.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/trusted_clients/metadata/)
- [trusted_clients.skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/trusted_clients/skip_processing/)
- [trusted_clients.waf_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/trusted_clients/waf_skip_processing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
