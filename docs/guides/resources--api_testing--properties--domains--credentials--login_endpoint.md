---
page_title: "domains.credentials.login_endpoint"
subcategory: ""
description: "domains.credentials.login_endpoint for xcsh_api_testing."
xcsh_docs: {"aliases": [], "body_bytes": 4631, "body_sha256": "sha256:d45e6392c22c81a550f2612d3c235710f588cc185172823ef48ad7d958e86e2e", "canonical_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint", "child_ids": ["xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint:json_payload"], "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint", "parent_id": "xcsh-docs:resources:api_testing:properties:domains:credentials", "path": "docs/guides/resources--api_testing--properties--domains--credentials--login_endpoint.md", "provider_name": "api_testing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["domains", "credentials", "login_endpoint"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/properties/domains/credentials/login_endpoint/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "domains.credentials.login_endpoint for xcsh_api_testing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.credentials.login_endpoint

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md)
- [Property reference](resources--api_testing--reference.md)
- [domains](resources--api_testing--properties--domains.md)
- [domains.credentials](resources--api_testing--properties--domains--credentials.md)
- domains.credentials.login_endpoint

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Login Endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path",
    "token_response_key")}
```

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
login_endpoint {
  # Configure direct properties listed below.
}
```

## Direct properties

- [json_payload](resources--api_testing--properties--domains--credentials--login_endpoint--json_payload.md): complete subsection reference.

<a id="schema-domains--credentials--login_endpoint--method"></a>

### method property

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-domains--credentials--login_endpoint--path"></a>

### path property

Type: `"string"`. Optional.

Path. URL path for the endpoint

Upstream description:

URL path for the endpoint

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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
    "ves.io.schema.rules.string.max_len": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "1024"
  }
}
```

<a id="schema-domains--credentials--login_endpoint--token_response_key"></a>

### token_response_key property

Type: `"string"`. Optional.

Configuration parameter for token response key.

Upstream description:

Configuration parameter for token response key

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

## Next pages

- [domains.credentials.login_endpoint.json_payload](resources--api_testing--properties--domains--credentials--login_endpoint--json_payload.md)
- [domains.credentials](resources--api_testing--properties--domains--credentials.md)
- [xcsh_api_testing](../resources/api_testing.md)
